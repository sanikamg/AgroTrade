package routes

import (
	"golang_project_ecommerce/pkg/api/handler"
	"golang_project_ecommerce/pkg/api/middlware"

	"github.com/gin-gonic/gin"
)

func AdminRoutes(api *gin.RouterGroup, adminHandler *handler.AdminHandler, productHandler *handler.ProductHandler) {
	//login
	login := api.Group("/adminlogin")
	{
		login.POST("/", adminHandler.AdminLogin)

	}
	//signup
	signup := api.Group("/adminsignup")
	{
		signup.POST("/", adminHandler.AdminSignup)
		signup.POST("/verify_otp", adminHandler.VerifyOTP)
	}
	//logout
	logout := api.Group("/adminlogout")
	{
		logout.POST("/", adminHandler.Logout)
	}

	api.Use(middlware.AuthenticateAdmin)
	{
		category := api.Group("/category")
		{
			category.POST("/", productHandler.SaveCategory)
			category.GET("/allcategory", productHandler.GetAllCategory)
		}

		products := api.Group("/products")
		{
			products.POST("/", productHandler.SaveProduct)
			products.DELETE("/delete", productHandler.RemoveProduct)
		}

		user := api.Group("/users")
		{
			user.GET("/", adminHandler.GetAllUsers)
			user.PATCH("/blockuser", adminHandler.BlockUser)
			user.PATCH("/unblockuser", adminHandler.UnBlockUser)
		}
	}
}
