package api

import (
	"file-sharing/pkg/api/handlers"
	"file-sharing/pkg/routes"

	"github.com/gin-gonic/gin"
)

// http server for the web application
type ServerHTTP struct {
	engine *gin.Engine
}

func NewServerHttp(
	userHandler *handlers.UserHandler,fileHandler *handlers.FileHandler) *ServerHTTP {
	engine := gin.New()
	engine.Use(gin.Logger())

	routes.UserRoutes(engine.Group("/users"), userHandler,fileHandler)

	return &ServerHTTP{
		engine: engine,
	}
}

func (sh *ServerHTTP) Start() {
	sh.engine.Run(":8083")
}
