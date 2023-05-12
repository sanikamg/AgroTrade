package handler

import (
	"errors"
	"golang_project_ecommerce/pkg/auth"
	"golang_project_ecommerce/pkg/common/response"
	"golang_project_ecommerce/pkg/domain"
	services "golang_project_ecommerce/pkg/usecase/interface"
	"golang_project_ecommerce/pkg/utils"
	"golang_project_ecommerce/pkg/utils/req"
	"golang_project_ecommerce/pkg/verification"
	"net/http"

	//"golang_project_ecommerce/pkg/utils/req"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

type UserHandler struct {
	userUsecase services.UserUsecase
}

func NewUserhandler(usecase services.UserUsecase) *UserHandler {
	return &UserHandler{
		userUsecase: usecase,
	}
}

var user domain.Users

//Register or Signup

func (cr *UserHandler) Register(c *gin.Context) {
	// var user domain.Users
	if err := c.ShouldBindJSON(&user); err != nil {
		res := response.ErrorResponse(400, "error while getting  the data from user side", err.Error(), user)

		c.JSON(http.StatusBadRequest, res)
		return
	}

	if _, err := verification.SendOtp("+91" + user.Phone); err != nil {
		res := response.ErrorResponse(400, "error while sending otp", err.Error(), nil)
		c.JSON(http.StatusBadRequest, res)
		return
	}

	response := response.SuccessResponse(200, "otp send successfully", nil)
	c.JSON(http.StatusOK, response)

}

func (cr *UserHandler) VerifyOTP(c *gin.Context) {
	//bind body details
	var body req.OtpStruct
	if err := c.ShouldBindJSON(&body); err != nil {
		res := response.ErrorResponse(400, "error while getting otp from user", err.Error(), nil)
		utils.ResponseJSON(c, res)
		return
	}

	if body.OTP == " " {
		err := errors.New("please enter otp")
		res := response.ErrorResponse(400, "invalid otp", err.Error(), nil)
		utils.ResponseJSON(c, res)
		return
	}

	// verifying otp

	err := verification.VerifyOtp("+91"+user.Phone, body.OTP)

	if err != nil {
		res := response.ErrorResponse(400, "error while verifying otp", err.Error(), nil)
		c.JSON(http.StatusBadRequest, res)
		return
	}

	user, err := cr.userUsecase.Register(c, user)

	message := "welcome  " + user.Name

	if err != nil {
		response := response.ErrorResponse(400, "failed register", err.Error(), "register failed")

		c.JSON(400, response)
		return
	}

	response := response.SuccessResponse(200, "user registration completed  successfully", message)
	c.JSON(http.StatusOK, response)

}

//end

//userlogin

func (cr *UserHandler) UserLogin(c *gin.Context) {
	// bind body details
	var body req.LoginStruct
	if err := c.ShouldBindJSON(&body); err != nil {
		res := response.ErrorResponse(400, "error while getting  the data from user side", err.Error(), body)

		utils.ResponseJSON(c, res)
		return
	}

	// check any field is empty
	if body.Username == " " && body.Password == " " {
		err := errors.New("enter username and password")
		res := response.ErrorResponse(400, "invalid input", err.Error(), nil)
		utils.ResponseJSON(c, res)
		return
	}

	// copy the values from the body to user
	var user domain.Users
	copier.Copy(&user, &body)

	// check whether the user exists and login usisng usecse function
	user, err := cr.userUsecase.Login(c, user)

	message := "Successfully logged in as " + body.Username

	if err != nil {
		response := response.ErrorResponse(400, "failed to login", err.Error(), "login failed")

		c.JSON(400, response)
		return

	}

	//generate tokenstring with jwt
	tokenString, err := auth.GenerateJWT(int(user.User_Id))
	if err != nil {
		response := response.ErrorResponse(400, "failed to login", err.Error(), "login failed")

		c.JSON(400, response)
		return
	}
	//set cookie

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("User_Authorization", tokenString["accessToken"], 3600*24*30, "/", " ", false, true)

	response := response.SuccessResponse(200, "Successfully logged in", message)

	c.JSON(http.StatusOK, response)

}

//user logout

func (cr *UserHandler) Logout(c *gin.Context) {
	c.SetCookie("User_Authorization", "", -1, "/ ", " ", false, true)
	response := response.SuccessResponse(200, "successfully logged out", nil)
	c.JSON(200, response)
}

// list all products on user side
func (uh *UserHandler) ListProducts(c *gin.Context) {
	var categoryname req.Categoryreq
	if err := c.ShouldBindJSON(&categoryname); err != nil {
		response := response.ErrorResponse(400, "Enter correct details", err.Error(), categoryname)
		c.JSON(400, response)
		return
	}

	products, err := uh.userUsecase.FindAllProducts(c, categoryname.CategoryName)
	if err != nil {
		response := response.ErrorResponse(400, "Error while fetching all products", err.Error(), products)
		c.JSON(400, response)
		return
	}

	response := response.SuccessResponse(200, "successfully displayed all products", products)
	c.JSON(200, response)
}
