package shareddomain

import (
	"fmt"
	"time"

	"github.com/twin-te/twinte-back/base"
)

var AllWeekdays = []time.Weekday{
	time.Sunday,
	time.Monday,
	time.Tuesday,
	time.Wednesday,
	time.Thursday,
	time.Friday,
	time.Saturday,
}

func ParseWeekday(s string) (time.Weekday, error) {
	ret, ok := base.FindByString(AllWeekdays, s)
	if ok {
		return ret, nil
	}
	return 0, fmt.Errorf("failed to parse Weekday %#v", s)
}
