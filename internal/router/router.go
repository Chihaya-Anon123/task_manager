package router

import (
	"errors"

	"github.com/Chihaya-Anon123/task_manager/internal/api"
	"github.com/Chihaya-Anon123/task_manager/internal/config"
	"github.com/Chihaya-Anon123/task_manager/internal/errs"
	"github.com/Chihaya-Anon123/task_manager/internal/middleware"
	"github.com/Chihaya-Anon123/task_manager/internal/response"
	"github.com/gin-gonic/gin"
)

func SetupRouter(cfg config.JWTConfig) *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.GET("/health", func(c *gin.Context) {
		response.Success(c, "server and database are ready")
	})

	r.GET("/test/app-error", func(c *gin.Context) {
		err := errs.ErrInvalidParams
		response.HandleError(c, err)
	})

	r.GET("/test/system-error", func(c *gin.Context) {
		err := errors.New("sql: connection is broken")
		response.HandleError(c, err)
	})

	apiV1 := r.Group("/api/v1")
	{
		authGroup := apiV1.Group("/auth")
		{
			authGroup.POST("/register", api.Register)
			authGroup.POST("/login", api.Login)
		}

		userGroup := apiV1.Group("/user")
		userGroup.Use(middleware.JWTAuth(cfg))
		{
			userGroup.GET("/me", api.GetMe)
		}

		taskGroup := apiV1.Group("/tasks")
		taskGroup.Use(middleware.JWTAuth(cfg))
		{
			taskGroup.POST("", api.CreateTask)
			taskGroup.GET("", api.ListTasks)
			taskGroup.GET("/:id", api.GetTaskDetail)
			taskGroup.PUT("/:id", api.UpdateTask)
			taskGroup.DELETE("/:id", api.DeleteTask)
		}
	}
	return r
}
