package api

import (
	"github.com/Chihaya-Anon123/task_manager/internal/code"
	"github.com/Chihaya-Anon123/task_manager/internal/errs"
	"github.com/Chihaya-Anon123/task_manager/internal/middleware"
	"github.com/Chihaya-Anon123/task_manager/internal/response"
	"github.com/Chihaya-Anon123/task_manager/internal/service"
	"github.com/gin-gonic/gin"
)

func CreateTask(c *gin.Context) {
	userID, ok := middleware.GetCurrentUserID(c)
	if !ok {
		response.HandleError(c, errs.New(code.CodeUnauthorized, "failed to get userID"))
		return
	}

	var input service.CreateTaskInput
	if err := c.ShouldBindJSON(&input); err != nil {
		response.HandleError(c, errs.New(code.CodeInvalidParams, "invalid request"))
		return
	}

	output, err := service.CreateTask(userID, input)
	if err != nil {
		response.HandleError(c, err)
		return
	}

	response.SuccessWithMessage(c, "task created successfully", output)
}
