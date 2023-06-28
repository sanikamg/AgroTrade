package http

import (
	_ "golang_project_ecommerce/cmd/api/docs"
	"golang_project_ecommerce/pkg/api/handler"
	"golang_project_ecommerce/pkg/api/routes"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type ServerHTTP struct {
	engine *gin.Engine
}

func NewServerHTTP(userHandler *handler.UserHandler, adminHandler *handler.AdminHandler, productHandler *handler.ProductHandler) *ServerHTTP {
	engine := gin.New()

	engine.Use(gin.Logger())
	engine.LoadHTMLGlob("/home/user/Documents/Project/AgroTrade/views/*.html")
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// set up routes
	routes.UserRoutes(engine.Group("/"), userHandler, productHandler)
	routes.AdminRoutes(engine.Group("/"), adminHandler, productHandler)

	return &ServerHTTP{engine: engine}
}

func (serverhttp *ServerHTTP) Start() {
	serverhttp.engine.Run(":8081")
}
