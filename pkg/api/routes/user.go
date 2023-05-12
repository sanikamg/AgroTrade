package routes

import (
	"golang_project_ecommerce/pkg/api/handler"
	"golang_project_ecommerce/pkg/api/middlware"

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
	//logout
	logout := api.Group("/logout")
	{
		logout.POST("/", userHandler.Logout)
	}

	api.Use(middlware.AuthenticateUser)
	{
		listproducts := api.Group("/productlist")
		{
			listproducts.GET("/", userHandler.ListProducts)
		}
	}
}
