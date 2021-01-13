package accessory

import "time"

const (
	apiDateLayout = "2006-01-02T15:04:05Z"
)

func apiDateFormat(date time.Time) string {
	return date.Format(apiDateLayout)
}
