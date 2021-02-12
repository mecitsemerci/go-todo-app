package dto

type HealthOutput struct {
	Status  string                 `json:"status" example:"ok"`
	Details map[string]interface{} `json:"details,omitempty"`
}
