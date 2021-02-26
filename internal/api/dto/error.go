package dto

//ErrorOutput is an output error case
type ErrorOutput struct {
	Code    int      `json:"code"`
	Message string   `json:"message"`
	Details []string `json:"details"`
}
