package routes

import (
	"golang_project_ecommerce/pkg/api/handler"

	"github.com/gin-gonic/gin"
)

func UserRoutes(api *gin.RouterGroup, userHandler *handler.UserHandler) {
	//login
	login := api.Group("/login")
	{
		login.POST("/", userHandler.UserLogin)

	}
	//signup
	signup := api.Group("/signup")
	{
		signup.POST("/", userHandler.Register)
		signup.POST("/verify_otp", userHandler.VerifyOTP)
	}
}
