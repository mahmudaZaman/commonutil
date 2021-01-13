package comutil

// ContainsStr accepts an array of string and another string under test, it returns true if
// string under exists in the given array.
func ContainsStr(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
