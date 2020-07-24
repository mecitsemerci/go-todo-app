package utils

import "time"

func UtcNow() time.Time {
	loc, _ := time.LoadLocation("UTC")
	now := time.Now().In(loc)
	return now
}
