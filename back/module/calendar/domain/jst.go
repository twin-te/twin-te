package calendardomain

import "time"

var JST *time.Location

func init() {
	var err error
	JST, err = time.LoadLocation("Asia/Tokyo")
	if err != nil {
		panic("failed to load timezone 'Asia/Tokyo': " + err.Error())
	}
}
