package convert

import (
	"fmt"
	"time"
)

const (
	timeFullFormat   = "2006-01-02 15:04:05.000"
	timeSimpleFormat = "2006-01-02"
)

// GetToday return ISO format: YYYY-MM-DD
func GetTomorrow() time.Time {
	startTmr := time.Now().In(time.UTC).Add(24 * time.Hour).Format(timeSimpleFormat)
	tmr, err := time.Parse(timeFullFormat, fmt.Sprintf("%s 00:00:00.000", startTmr))
	if err != nil {
		return time.Now()
	}
	return tmr
}

func GetToday() time.Time {
	startTmr := time.Now().In(time.UTC).Format(timeSimpleFormat)
	tmr, err := time.Parse(timeFullFormat, fmt.Sprintf("%s 00:00:00.000", startTmr))
	if err != nil {
		return time.Now()
	}
	return tmr
}

func IsToday(date string) bool {
	return false
}

func StringToTime(val string) (time.Time, error) {
	return time.Parse(timeFullFormat, val)
}

func StringToSimpleTime(val string) (time.Time, error) {
	return time.Parse(timeSimpleFormat, val)
}

func TimeToSimpleTime(timeline time.Time) (time.Time, error) {
	return time.Parse(timeSimpleFormat, timeline.String())
}
