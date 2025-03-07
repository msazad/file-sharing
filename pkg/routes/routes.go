package routes

import (
	"file-sharing/pkg/api/handlers"
	"file-sharing/pkg/api/middleware"

	"github.com/gin-gonic/gin"
)

func UserRoutes(engine *gin.RouterGroup, userHandler *handlers.UserHandler,fileHandler *handlers.FileHandler) {
	engine.POST("/signup", userHandler.SignUp)
	engine.POST("/login", userHandler.Login)
	engine.Use(middleware.UserAuthMiddleware)
	{
		engine.GET("/upload",fileHandler.UploadFile)
	}
}
