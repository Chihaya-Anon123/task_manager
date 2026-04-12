package api

import (
	"strconv"

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

func ListTasks(c *gin.Context) {
	userID, ok := middleware.GetCurrentUserID(c)
	if !ok {
		response.HandleError(c, errs.New(code.CodeUnauthorized, "failed to get userID"))
		return
	}

	var input service.ListTasksInput
	if err := c.ShouldBindQuery(&input); err != nil {
		response.HandleError(c, errs.New(code.CodeInvalidParams, "invalid request"))
		return
	}

	output, err := service.ListTasks(userID, input)
	if err != nil {
		response.HandleError(c, err)
		return
	}

	response.Success(c, output)
}

func GetTaskDetail(c *gin.Context) {
	userID, ok := middleware.GetCurrentUserID(c)
	if !ok {
		response.HandleError(c, errs.New(code.CodeUnauthorized, "failed to get current user"))
		return
	}

	taskID64, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil || taskID64 == 0 {
		response.HandleError(c, errs.New(code.CodeInvalidParams, "invalid task id"))
		return
	}

	output, err := service.GetTaskDetail(uint(taskID64), userID)
	if err != nil {
		response.HandleError(c, err)
		return
	}

	response.Success(c, output)
}

func UpdateTask(c *gin.Context) {
	userID, ok := middleware.GetCurrentUserID(c)
	if !ok {
		response.HandleError(c, errs.New(code.CodeUnauthorized, "failed to get current user"))
		return
	}

	taskID64, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil || taskID64 == 0 {
		response.HandleError(c, errs.New(code.CodeInvalidParams, "invalid task id"))
		return
	}

	var input service.UpdateTaskInput
	if err := c.ShouldBindJSON(&input); err != nil {
		response.HandleError(c, errs.New(code.CodeInvalidParams, "invalid request"))
		return
	}

	output, err := service.UpdateTask(userID, uint(taskID64), input)
	if err != nil {
		response.HandleError(c, err)
		return
	}

	response.SuccessWithMessage(c, "update task successfully", output)
}

func DeleteTask(c *gin.Context) {
	userID, ok := middleware.GetCurrentUserID(c)
	if !ok {
		response.HandleError(c, errs.New(code.CodeUnauthorized, "failed to get current user"))
		return
	}

	taskID64, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil || taskID64 == 0 {
		response.HandleError(c, errs.New(code.CodeInvalidParams, "invalid task id"))
		return
	}

	if err := service.DeleteTask(userID, uint(taskID64)); err != nil {
		response.HandleError(c, err)
		return
	}

	response.SuccessWithMessage(c, "delete task successfully", nil)
}
