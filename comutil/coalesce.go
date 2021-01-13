package comutil

// FirstNotNullString is like sql coalesce on string.
func FirstNotNullString(args ...string) string {
	for _, v := range args {
		if len(v) > 0 {
			return v
		}
	}
	return ""
}
