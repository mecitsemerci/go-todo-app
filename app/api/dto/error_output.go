package dto

import "github.com/gin-gonic/gin"

type ErrorOutput struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Details interface{} `json:"details"`
}

func NewErrorOutput(ctx *gin.Context, status int, message string, err error) {

	output := ErrorOutput{
		Code:    status,
		Message: message,
		Details: map[string]string{
			"internal_message": err.Error(),
		},
	}
	ctx.JSON(status, output)
}
