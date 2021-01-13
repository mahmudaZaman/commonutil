package accessory

import "time"

// Func exports all the accessories/utility function.
var Func = &export{
	ContainsStr:        containsStr,
	FirstNotNullString: firstNotNullString,
	APIDateFormat:      apiDateFormat,
	GetUUIDWithoutDash: getUUIDWithoutDash,
}

type export struct {
	ContainsStr        func(s []string, e string) bool
	FirstNotNullString func(args ...string) string
	APIDateFormat      func(date time.Time) string
	GetUUIDWithoutDash func() string
}
