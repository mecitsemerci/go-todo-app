package dto

type HealthOutput struct {
	Status       string           `json:"status" example:"healthy"`
	Dependencies *map[string]bool `json:"dependencies"`
}
