package calendarv1beta

import (
	"log"
	"time"
)

var jst *time.Location

func init() {
	var err error
	jst, err = time.LoadLocation("Asia/Tokyo")
	if err != nil {
		log.Printf("Failed to load timezone 'Asia/Tokyo': %v. Falling back to Local.", err)
		jst = time.Local
	}
}
