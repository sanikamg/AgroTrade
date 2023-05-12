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

	"github.com/gin-gonic/gin"
)

type AdminHandler struct {
	AdminUsecase services.AdminUsecase
}

func NewAdminHandler(usecase services.AdminUsecase) *AdminHandler {
	return &AdminHandler{
		AdminUsecase: usecase,
	}
}

var admin domain.AdminDetails

func (ah *AdminHandler) AdminSignup(c *gin.Context) {
	if err := c.ShouldBindJSON(&admin); err != nil {
		res := response.ErrorResponse(400, "error while getting admin details", err.Error(), admin)
		c.JSON(http.StatusBadRequest, res)
		return
	}

	if _, err := verification.SendOtp("+91" + admin.Phone); err != nil {
		res := response.ErrorResponse(400, "error while sending otp", err.Error(), admin)
		c.JSON(http.StatusBadRequest, res)
		return
	}

	response := response.SuccessResponse(200, "otp send successfully", nil)
	c.JSON(http.StatusOK, response)

}

func (ad *AdminHandler) VerifyOTP(c *gin.Context) {
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

	err := verification.VerifyOtp("+91"+admin.Phone, body.OTP)

	if err != nil {
		res := response.ErrorResponse(400, "error while verifying otp", err.Error(), nil)
		c.JSON(http.StatusBadRequest, res)
		return
	}

	admin, err := ad.AdminUsecase.AdminSignup(c, admin)

	message := "welcome  " + admin.Name

	if err != nil {
		response := response.ErrorResponse(400, "failed register", err.Error(), "register failed")

		c.JSON(400, response)
		return
	}

	response := response.SuccessResponse(200, "user registration completed  successfully", message)
	c.JSON(http.StatusOK, response)

}

//admin login

func (ad *AdminHandler) AdminLogin(c *gin.Context) {
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

	admin, err := ad.AdminUsecase.FindByUsername(c, body.Username)
	if err != nil {
		response := response.ErrorResponse(400, "Enter valid username", err.Error(), "login failed")

		c.JSON(400, response)
		return
	}

	// check whether the user exists and login usisng usecse function
	if err := ad.AdminUsecase.AdminLogin(c, admin); err != nil {
		response := response.ErrorResponse(400, "failed to login", err.Error(), "login failed")

		c.JSON(400, response)
		return
	}

	message := "Successfully logged in as " + body.Username

	//generate tokenstring with jwt
	tokenString, err := auth.GenerateJWT(int(admin.ID))
	if err != nil {
		response := response.ErrorResponse(400, "failed to login", err.Error(), "login failed")

		c.JSON(400, response)
		return
	}
	//set cookie

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Admin_Authorization", tokenString["accessToken"], 3600*24*30, "/", " ", false, true)

	response := response.SuccessResponse(200, "Successfully logged in", message)

	c.JSON(http.StatusOK, response)

}

// to display all users in admin side

func (ad *AdminHandler) GetAllUsers(c *gin.Context) {
	users, err := ad.AdminUsecase.FindAllUsers(c)
	if err != nil {
		response := response.ErrorResponse(400, "error while finding all users", err.Error(), nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := response.SuccessResponse(200, "successfully displaying all users", nil, users)
	c.JSON(http.StatusOK, response)
}

// to block users
func (ad *AdminHandler) BlockUser(c *gin.Context) {
	var blockuser req.Block
	if err := c.ShouldBindJSON(&blockuser); err != nil {
		response := response.ErrorResponse(400, "error while getting id from admin", err.Error(), nil)
		c.JSON(400, response)
		return
	}

	err := ad.AdminUsecase.BlockUser(c, int(blockuser.UserID))
	if err != nil {
		response := response.ErrorResponse(400, "error while block user", err.Error(), nil)
		c.JSON(400, response)
		return

	}
	response := response.SuccessResponse(200, "successfully blocked user", nil, blockuser)
	c.JSON(http.StatusOK, response)
}

// to unblock user
func (ad *AdminHandler) UnBlockUser(c *gin.Context) {
	var unblockuser req.Block
	if err := c.ShouldBindJSON(&unblockuser); err != nil {
		response := response.ErrorResponse(400, "error while getting id from admin", err.Error(), nil)
		c.JSON(400, response)
		return
	}

	err := ad.AdminUsecase.UnBlockUser(c, int(unblockuser.UserID))
	if err != nil {
		response := response.ErrorResponse(400, "error while unblock user", err.Error(), nil)
		c.JSON(400, response)
		return

	}
	response := response.SuccessResponse(200, "successfully unblocked user", nil, unblockuser)
	c.JSON(http.StatusOK, response)
}

//user logout

func (ad *AdminHandler) Logout(c *gin.Context) {
	c.SetCookie("Admin_Authorization", "", -1, " /", " ", false, true)
	response := response.SuccessResponse(200, "successfully logged out", nil)
	c.JSON(200, response)
}
