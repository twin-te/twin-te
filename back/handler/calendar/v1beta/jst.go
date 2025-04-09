package calendarv1beta

import "time"

var jst *time.Location

func init() {
	jst, _ = time.LoadLocation("Asia/Tokyo") // TODO: Catch error gracefully
}
