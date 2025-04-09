package calendarv1beta

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"strings"
	"time"

	"cloud.google.com/go/civil"
	"github.com/google/uuid"
	timetableappdto "github.com/twin-te/twin-te/back/module/timetable/appdto"
)

var icsTextEscaper = strings.NewReplacer(
	`\`, `\\`,
	`;`, `\;`,
	`,`, `\,`,
	"\n", `\n`,
)

func icsTime(t civil.DateTime) string {
	return "TZID=Asia/Tokyo:" + t.In(jst).Format("20060102T150405")
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

func WriteICalendar(writer io.Writer, modules []Module, courses []*timetableappdto.RegisteredCourse) error {
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
		ss := GetSchedules(modules, c.Schedules)
		for _, s := range ss {
			writeCalendarEvent(w, c, s)
		}
	}

	w.write("END:VCALENDAR")

	return w.err
}

func generateUID(c *timetableappdto.RegisteredCourse, s Schedule) uuid.UUID {
	ns := uuid.MustParse("7f343367-6ab8-4c2a-9c5f-030dc00e9ac7")
	data := new(bytes.Buffer)
	data.WriteString(c.ID.String())
	binary.Write(data, binary.BigEndian, s.StartTime.In(jst).Unix())
	return uuid.NewSHA1(ns, data.Bytes())
}

func buildDescription(c *timetableappdto.RegisteredCourse) string {
	url := fmt.Sprintf("https://app.twinte.net/course/%s", c.ID)
	if c.Memo != "" {
		return fmt.Sprintf("%s\n\n---\n%s", c.Memo, url)
	} else {
		return url
	}
}

func writeCalendarEvent(w *errWriter, c *timetableappdto.RegisteredCourse, s Schedule) {
	w.write("BEGIN:VEVENT")

	w.write("DTSTAMP;%s", icsTime(s.StartTime))
	w.write("UID:%s", generateUID(c, s))

	w.write("SUMMARY:%s", icsTextEscaper.Replace(c.Name.String()))
	w.write("LOCATION:%s", icsTextEscaper.Replace(s.Location))
	w.write("DESCRIPTION:%s", icsTextEscaper.Replace(buildDescription(c)))

	w.write("DTSTART;%s", icsTime(s.StartTime))
	w.write("DTEND;%s", icsTime(s.EndTime))
	w.write("RRULE;TZID=Asia/Tokyo:FREQ=WEEKLY;INTERVAL=1;BYDAY=%s;UNTIL=%s", icsDay(s.Weekday), s.Until.In(jst).Format("20060102T000000"))

	for _, t := range s.Exceptions {
		w.write("EXDATE;%s", icsTime(t))
	}
	for _, t := range s.Additions {
		w.write("RDATE;%s", icsTime(t))
	}

	w.write("END:VEVENT")
}
