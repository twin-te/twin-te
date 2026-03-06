package calendarv1beta

import (
	"bytes"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	calendardomain "github.com/twin-te/twin-te/back/module/calendar/domain"
	shareddomain "github.com/twin-te/twin-te/back/module/shared/domain"
	"github.com/twin-te/twin-te/back/module/shared/domain/idtype"
)

func getYear(c echo.Context) (shareddomain.AcademicYear, error) {
	yearParam := c.QueryParam("year")
	if yearParam == "" {
		now := time.Now().In(calendardomain.JST)
		return shareddomain.NewAcademicYear(now.Year(), now.Month())
	}

	year, err := strconv.Atoi(yearParam)
	if err != nil {
		return 0, err
	}

	return shareddomain.ParseAcademicYear(year)
}

func getTags(c echo.Context) ([]idtype.TagID, error) {
	ss, ok := c.QueryParams()["tags[]"]
	if !ok {
		return nil, nil
	}

	ids := make([]idtype.TagID, len(ss))
	for i, s := range ss {
		id, err := idtype.ParseTagID(s)
		if err != nil {
			return nil, err
		}
		ids[i] = id
	}

	return ids, nil
}

func (h *impl) ICSHandler(c echo.Context) error {
	year, err := getYear(c)
	if err != nil {
		log.Printf("failed to get academic year: %v", err)
		return err
	}

	rdateParam := c.QueryParam("rdate")
	isRdateSupported := rdateParam == "true"

	tags, err := getTags(c)
	if err != nil {
		log.Printf("failed to get tags: %+v", err)
		return err
	}

	ctx := c.Request().Context()
	data, err := h.calendar.ExportTimetableToICal(ctx, year, tags, isRdateSupported)
	if err != nil {
		log.Printf("failed to export timetable to iCal: %+v", err)
		return err
	}

	return c.Stream(http.StatusOK, "text/calendar; charset=utf-8", bytes.NewReader(data))
}
