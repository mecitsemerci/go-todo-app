package todo

import "errors"

var (
	ErrTaskIDRequired = errors.New("task ID required")
)
