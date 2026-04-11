package api

import (
	"github.com/Chihaya-Anon123/task_manager/internal/code"
	"github.com/Chihaya-Anon123/task_manager/internal/errs"
	"github.com/Chihaya-Anon123/task_manager/internal/middleware"
	"github.com/Chihaya-Anon123/task_manager/internal/response"
	"github.com/Chihaya-Anon123/task_manager/internal/service"
	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var input service.RegisterInput
	if err := c.ShouldBindJSON(&input); err != nil {
		response.HandleError(c, errs.New(code.CodeInvalidParams, "invalid request"))
		return
	}

	output, err := service.Register(input)
	if err != nil {
		response.HandleError(c, err)
		return
	}

	response.SuccessWithMessage(c, "register success", output)
}

func Login(c *gin.Context) {
	var input service.LoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		response.HandleError(c, errs.New(code.CodeInvalidParams, "invalid request"))
		return
	}

	output, err := service.Login(input)
	if err != nil {
		response.HandleError(c, err)
		return
	}

	response.SuccessWithMessage(c, "login success", output)
}

func GetMe(c *gin.Context) {
	userID, ok := middleware.GetCurrentUserID(c)
	if !ok {
		response.Fail(c, code.CodeUnauthorized, "failed to get current user")
		return
	}

	username, ok := middleware.GetCurrentUsername(c)
	if !ok {
		response.Fail(c, code.CodeUnauthorized, "failed to get current user")
		return
	}

	response.Success(c, gin.H{
		"user_id":  userID,
		"username": username,
	})
}
