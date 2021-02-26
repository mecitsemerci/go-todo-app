package httperror

import (
	"github.com/gin-gonic/gin"
	"github.com/mecitsemerci/go-todo-app/internal/api/dto"
)

//NewError populate error output dto according to given error information
func NewError(ctx *gin.Context, status int, message string, err error) {
	output := dto.ErrorOutput{
		Code:    status,
		Message: message,
	}
	if err != nil {
		output.Details = []string{err.Error()}
	}
	ctx.JSON(status, output)
}
