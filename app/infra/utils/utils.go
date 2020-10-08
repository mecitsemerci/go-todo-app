package utils

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"strings"
	"time"
)

func UtcNow() time.Time {
	return time.Now().UTC()
}

func IsEmptyOrWhiteSpace(str string) bool {
	if str == "" || len(strings.TrimSpace(str)) == 0 {
		return true
	}

	return false
}

func NewOID() primitive.ObjectID {
	return primitive.NewObjectIDFromTimestamp(UtcNow())
}

func OIDFromStr(str string) (primitive.ObjectID, error) {
	return primitive.ObjectIDFromHex(str)
}
