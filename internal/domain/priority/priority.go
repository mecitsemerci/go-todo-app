package priority

import (
	"encoding/json"
	"errors"
)

type Level int

const (
	None     Level = 0
	Normal   Level = 1
	High     Level = 2
	Critical Level = 3
)

func (l Level) String() string {
	return [...]string{"None", "Normal", "High", "Critical"}[l]
}

func (l *Level) UnmarshalJSON(b []byte) error {
	type level Level
	var value = (*level)(l)
	err := json.Unmarshal(b, &value)
	if err != nil {
		return err
	}

	switch *l {
	case None, Normal, High, Critical:
		return nil
	}
	return errors.New("invalid priority level type")
}

func (l *Level) Equal(other *Level) bool {
	return *l == *other
}

func (l Level) Value() int {
	return int(l)
}
