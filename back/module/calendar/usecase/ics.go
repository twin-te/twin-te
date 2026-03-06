package calendarusecase

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"strings"
	"time"

	"cloud.google.com/go/civil"
	"github.com/google/uuid"
	"github.com/twin-te/twin-te/back/appenv"
	calendardomain "github.com/twin-te/twin-te/back/module/calendar/domain"
	timetableappdto "github.com/twin-te/twin-te/back/module/timetable/appdto"
)

var icsTextEscaper = strings.NewReplacer(
	`\`, `\\`,
	`;`, `\;`,
	`,`, `\,`,
	"\n", `\n`,
)

func icsTime(t civil.DateTime) string {
	return "TZID=Asia/Tokyo:" + t.In(calendardomain.JST).Format("20060102T150405")
}

func icsTimeUTC(t civil.DateTime) string {
	return t.In(calendardomain.JST).UTC().Format("20060102T150405Z")
}

func icsDtstamp() string {
	return time.Now().UTC().Format("20060102T150405Z")
}

func icsDay(d time.Weekday) string {
	return strings.ToUpper(d.String()[:2])
}

type errWriter struct {
	w   io.Writer
	err error
}

func (w *errWriter) write(format string, a ...any) {
	if w.err != nil {
		return
	}
	_, w.err = fmt.Fprintf(w.w, format+"\r\n", a...)
}

func (uc *impl) writeICalendar(writer io.Writer, modules []*calendardomain.SchoolCalendarModule, courses []*timetableappdto.RegisteredCourse, isRdateSupported bool) error {
	w := &errWriter{w: writer}

	w.write("BEGIN:VCALENDAR")
	w.write("VERSION:2.0")
	w.write("PRODID:-//Twin:te//Twin:te Calendar Service//EN")

	w.write("BEGIN:VTIMEZONE")
	w.write("TZID:Asia/Tokyo")
	w.write("BEGIN:STANDARD")
	w.write("DTSTART:19700101T000000")
	w.write("TZOFFSETFROM:+0900")
	w.write("TZOFFSETTO:+0900")
	w.write("TZNAME:JST")
	w.write("END:STANDARD")
	w.write("END:VTIMEZONE")

	for _, c := range courses {
		ss := calendardomain.GetSchedules(modules, c.Schedules)
		for _, s := range ss {
			writeCalendarEvent(w, c, s, isRdateSupported)
		}
	}

	w.write("END:VCALENDAR")

	return w.err
}

func generateUID(c *timetableappdto.RegisteredCourse, s calendardomain.Schedule) uuid.UUID {
	ns := uuid.MustParse("7f343367-6ab8-4c2a-9c5f-030dc00e9ac7")
	data := new(bytes.Buffer)
	data.WriteString(c.ID.String())
	binary.Write(data, binary.BigEndian, s.StartTime.In(calendardomain.JST).Unix())
	return uuid.NewSHA1(ns, data.Bytes())
}

func buildDescription(c *timetableappdto.RegisteredCourse) string {
	url := fmt.Sprintf("%s/course/%s", appenv.APP_URL, c.ID)
	if c.Memo != "" {
		return fmt.Sprintf("%s\n\n---\n%s", c.Memo, url)
	}
	return url
}

func writeCalendarEvent(w *errWriter, c *timetableappdto.RegisteredCourse, s calendardomain.Schedule, isRdateSupported bool) {
	w.write("BEGIN:VEVENT")

	w.write("DTSTAMP:%s", icsDtstamp())
	w.write("UID:%s", generateUID(c, s))

	w.write("SUMMARY:%s", icsTextEscaper.Replace(c.Name.String()))
	w.write("LOCATION:%s", icsTextEscaper.Replace(s.Location))
	w.write("DESCRIPTION:%s", icsTextEscaper.Replace(buildDescription(c)))

	w.write("DTSTART;%s", icsTime(s.StartTime))
	w.write("DTEND;%s", icsTime(s.EndTime))
	w.write("RRULE:FREQ=WEEKLY;INTERVAL=1;BYDAY=%s;UNTIL=%s", icsDay(s.Weekday), icsTimeUTC(s.Until))

	for _, t := range s.Exceptions {
		w.write("EXDATE;%s", icsTime(t))
	}
	if isRdateSupported {
		for _, t := range s.Additions {
			w.write("RDATE;%s", icsTime(t))
		}
	} else {
		for _, t := range s.Additions {
			w.write("END:VEVENT")
			w.write("BEGIN:VEVENT")

			newStartTime := s.StartTime
			newStartTime.Date = t.Date
			newEndTime := s.EndTime
			newEndTime.Date = t.Date

			newSchedule := s
			newSchedule.StartTime = newStartTime

			w.write("UID:%s", generateUID(c, newSchedule))
			w.write("DTSTAMP:%s", icsDtstamp())

			w.write("SUMMARY:%s", icsTextEscaper.Replace(c.Name.String()))
			w.write("LOCATION:%s", icsTextEscaper.Replace(s.Location))
			w.write("DESCRIPTION:%s", icsTextEscaper.Replace(buildDescription(c)))

			w.write("DTSTART;%s", icsTime(newStartTime))
			w.write("DTEND;%s", icsTime(newEndTime))
		}
	}

	w.write("END:VEVENT")
}
