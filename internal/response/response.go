package response

import (
	"net/http"

	"github.com/Chihaya-Anon123/task_manager/internal/code"
	"github.com/Chihaya-Anon123/task_manager/internal/errs"
	"github.com/Chihaya-Anon123/task_manager/internal/logger"
	"github.com/gin-gonic/gin"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Code:    code.CodeSuccess,
		Message: "success",
		Data:    data,
	})
}

func SuccessWithMessage(c *gin.Context, msg string, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Code:    code.CodeSuccess,
		Message: msg,
		Data:    data,
	})
}

func Fail(c *gin.Context, code int, message string) {
	c.JSON(http.StatusOK, Response{
		Code:    code,
		Message: message,
		Data:    nil,
	})
}

func FailByCode(c *gin.Context, code_ int) {
	c.JSON(http.StatusOK, Response{
		Code:    code_,
		Message: code.GetMessage(code_),
		Data:    nil,
	})
}

func HandleError(c *gin.Context, err error) {
	if err == nil {
		return
	}

	if appErr, ok := err.(*errs.AppError); ok {
		logger.Log.Warnw("business error",
			"code", appErr.Code,
			"message", appErr.Message,
			"path", c.Request.URL.Path,
			"method", c.Request.Method,
		)

		c.JSON(http.StatusOK, Response{
			Code:    appErr.Code,
			Message: appErr.Message,
			Data:    nil,
		})
		return
	}

	logger.Log.Errorw("system error",
		"error", err.Error(),
		"path", c.Request.URL.Path,
		"method", c.Request.Method,
	)

	c.JSON(http.StatusOK, Response{
		Code:    code.CodeInternalServer,
		Message: code.GetMessage(code.CodeInternalServer),
		Data:    nil,
	})
}
