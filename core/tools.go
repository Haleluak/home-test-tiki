package core

// IndexOf will return the index of @str within @array or -1 if not existing
func IndexOf(array []string, str string) int {
	for i, v := range array {
		if v == str {
			return i
		}
	}
	return -1
}
