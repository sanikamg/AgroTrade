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
		signup.GET("/verify_otp", adminHandler.VerifyOTP)
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
			category.DELETE("/delete", productHandler.DeleteCategory)
		}

		products := api.Group("/products")
		{
			products.POST("/", productHandler.SaveProduct)
			products.DELETE("/delete", productHandler.RemoveProduct)
			products.PATCH("/edit", productHandler.EditProduct)
			products.POST("/addimage", productHandler.AddImage)
			products.GET("/getallproducts", productHandler.GetAllProducts)
			products.GET("/getallproductsbycategory", productHandler.GetAllProductsByCategory)
			products.GET("/getproductbyid", productHandler.GetProductById)
		}

		user := api.Group("/users")
		{
			user.GET("/", adminHandler.GetAllUsers)
			user.PATCH("/blockuser", adminHandler.BlockUser)
			user.PATCH("/unblockuser", adminHandler.UnBlockUser)
		}

		coupon := api.Group("/coupon")
		{
			coupon.POST("/add", productHandler.AddCoupon)
			coupon.GET("/list", productHandler.ListCoupon)
		}

		paymentmethod := api.Group("/paymentmethod")
		{
			paymentmethod.POST("/add", productHandler.AddpaymentMethod)
			paymentmethod.GET("/view", productHandler.GetAllPaymentMethods)
			paymentmethod.PATCH("/update", productHandler.UpdatePaymentMethod)
			paymentmethod.DELETE("/delete", productHandler.DeleteMethod)
		}

		Dashboard := api.Group("/dashboard")
		{
			Dashboard.GET("/salesreport", productHandler.SalesReport)
			Dashboard.GET("/getallorders", productHandler.GetAllOrders)
			Dashboard.GET("/getallproducts", productHandler.GetAllProducts)
			Dashboard.GET("/salesdata", productHandler.Statistics)
			Dashboard.GET("/pendingdelivery", productHandler.FindPendingDelivery)
		}

	}
}
