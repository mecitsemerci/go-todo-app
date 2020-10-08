package check

import "strings"

func IsEmptyOrWhiteSpace(str string) bool {
	if str == "" || len(strings.TrimSpace(str)) == 0 {
		return true
	}

	return false
}
