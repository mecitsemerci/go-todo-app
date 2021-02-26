package enum

import (
	"encoding/json"
	"errors"
)

// PriorityLevel indicates importance of item
type PriorityLevel uint8

// All PriorityLevel list
const (
	// PriorityNone is default PriorityLevel
	PriorityNone     PriorityLevel = 0
	PriorityNormal   PriorityLevel = 1
	PriorityHigh     PriorityLevel = 2
	PriorityCritical PriorityLevel = 3
)

// UnmarshalJSON validate PriorityLevel when deserialized
func (p *PriorityLevel) UnmarshalJSON(b []byte) error {
	type priority PriorityLevel
	var value = (*priority)(p)
	err := json.Unmarshal(b, &value)
	if err != nil {
		return err
	}

	switch *p {
	case PriorityNone, PriorityNormal, PriorityHigh, PriorityCritical:
		return nil
	}
	return errors.New("invalid priority level type")
}
