package enum

type PriorityLevel uint8

const (
	PrioryNone     PriorityLevel = 0
	PrioryNormal   PriorityLevel = 1
	PrioryHigh     PriorityLevel = 2
	PrioryCritical PriorityLevel = 3
)
