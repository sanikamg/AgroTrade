package routes

import (
	"golang_project_ecommerce/pkg/api/handler"
	"golang_project_ecommerce/pkg/api/middlware"

	"github.com/gin-gonic/gin"
)

func UserRoutes(api *gin.RouterGroup, userHandler *handler.UserHandler, productHandler *handler.ProductHandler) {
	//login
	login := api.Group("/login")
	{
		login.POST("/", userHandler.UserLogin)

	}
	//signup
	signup := api.Group("/signup")
	{
		signup.POST("/loginorsignup", userHandler.SendOtpPhn)
		signup.GET("/verify_otp", userHandler.VerifyOTP)
		signup.POST("/register", userHandler.Register)
	}
	//logout
	logout := api.Group("/logout")
	{
		logout.POST("/", userHandler.Logout)
	}
	//forgotpassword
	forgotpass := api.Group("/sendotp")
	{
		forgotpass.POST("/", userHandler.SendOtpForgotPass)
	}
	//authentication of forgot password
	api.Use(middlware.AutheticatePhn)
	{
		forgotverification := api.Group("/forgot")
		{
			forgotverification.POST("/otpverification", userHandler.VerifyOTPForgotPass)
			forgotverification.POST("/newpass", userHandler.ForgotPassword)
		}
	}

	api.Use(middlware.AuthenticateUser)
	{
		listproducts := api.Group("/productlist")
		{
			listproducts.GET("/", productHandler.GetAllProducts)
			listproducts.GET("/getallproductsbycategory", productHandler.GetAllProductsByCategory)
		}

		userprofile := api.Group("/userdetails")
		{
			userprofile.GET("/", userHandler.FindUserById)
			userprofile.PATCH("/edit", userHandler.EditUserDetails)
		}

		Address := api.Group("/address")
		{
			Address.POST("/add", userHandler.AddAddress)
			Address.PATCH("/edit", userHandler.EditAddress)
			Address.GET("/listaddress", userHandler.ListAddresses)
			Address.DELETE("/delete", userHandler.DeleteAddress)
		}

		Cart := api.Group("/cart")
		{
			Cart.POST("/add", productHandler.AddProductToCart)
			Cart.GET("/view", productHandler.ViewCart)
			Cart.DELETE("/delete", productHandler.RemoveProductFromCart)
		}

		Order := api.Group("/order")
		{
			Order.POST("/create", productHandler.CreateOrder)
			Order.POST("/payment", productHandler.PlaceOrder)
		}
	}
}
