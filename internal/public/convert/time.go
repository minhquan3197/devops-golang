package convert

import "time"

// TimeToISODate func convernt Time to ISODate
func TimeToISODate(payload time.Time) time.Time {
	const (
		layoutISO = "2006-01-02T15:04:05.000Z"
	)
	parseDateString := payload.UTC().Format("2006-01-02T15:04:05.999Z")
	timeConvernt, _ := time.Parse(layoutISO, parseDateString)
	return timeConvernt
}
