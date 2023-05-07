package http

import (
	"golang_project_ecommerce/pkg/api/handler"
	"golang_project_ecommerce/pkg/api/routes"

	"github.com/gin-gonic/gin"
)

type ServerHTTP struct {
	engine *gin.Engine
}

func NewServerHTTP(userHandler *handler.UserHandler) *ServerHTTP {
	engine := gin.New()

	engine.Use(gin.Logger())

	// set up routes
	routes.UserRoutes(engine.Group("/"), userHandler)

	return &ServerHTTP{engine: engine}
}

func (serverhttp *ServerHTTP) Start() {
	serverhttp.engine.Run(":8081")
}
