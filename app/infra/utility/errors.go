package utility

import (
	"github.com/gin-gonic/gin"
	"github.com/mecitsemerci/clean-go-todo-api/app/api/dto"
)

func NewErrorOutput(ctx *gin.Context, status int, message string, err error) {

	output := dto.ErrorOutput{
		Code:    status,
		Message: message,
		Details: []string{err.Error()},
	}
	ctx.JSON(status, output)
}

func NewErrorsOutput(ctx *gin.Context, status int, message string, errorList []error) {

	output := dto.ErrorOutput{
		Code:    status,
		Message: message,
	}
	if len(errorList) > 0{
		var errorMessages []string
		for _, err := range errorList{
			errorMessages = append(errorMessages, err.Error())
		}
		output.Details = errorMessages
	}

	ctx.JSON(status, output)
}

