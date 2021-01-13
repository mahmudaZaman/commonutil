package comutil

import "time"

const (
	apiDateLayout = "2006-01-02T15:04:05Z"
)

// APIDateFormat accepts a date and formats it as per API standard.
func APIDateFormat(date time.Time) string {
	return date.Format(apiDateLayout)
}
