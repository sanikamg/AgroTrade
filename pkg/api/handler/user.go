package handler

import (
	"errors"
	"golang_project_ecommerce/pkg/api/middlware"
	"golang_project_ecommerce/pkg/auth"
	"golang_project_ecommerce/pkg/common/response"
	"golang_project_ecommerce/pkg/domain"
	services "golang_project_ecommerce/pkg/usecase/interface"
	"golang_project_ecommerce/pkg/utils"
	"golang_project_ecommerce/pkg/utils/req"
	"net/http"
	"strconv"

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

// SendOtpToPhone godoc
//
//	@summary		api for user to send otp to phone
//	@description	Enter phone number
//	@tags			SignUp For User
//	@Param			inputs	body	req.Phn	true	"Input Field"
//	@Router			/signup/ [post]
//	@Success		200	{object}	response.Response{}	"error while sending otp"
//	@Failure		400	{object}	response.Response{}	"otp send successfully"
//
// send otp to phn number
func (uh *UserHandler) SendOtpPhn(c *gin.Context) {
	var user domain.Users
	var phone req.Phn
	if err := c.ShouldBindJSON(&phone); err != nil {
		res := response.ErrorResponse(400, "error while getting admin details", err.Error(), admin)
		c.JSON(http.StatusBadRequest, res)
		return
	}

	user.Phone = phone.Phone

	err := uh.userUsecase.SendOtpPhn(c, user)
	if err != nil {
		res := response.ErrorResponse(400, "error while sending otp", err.Error(), nil)
		c.JSON(http.StatusBadRequest, res)
		return
	}

	//generate tokenstring with jwt
	tokenString, err := auth.GenerateJWTPhn(user.Phone)
	if err != nil {
		response := response.ErrorResponse(400, "failed to send otp", err.Error(), "user didn't exist")

		c.JSON(400, response)
		return
	}
	//set cookie

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Signup_Authorization", tokenString["accessToken"], 3600*24*30, "/", " ", false, true)

	response := response.SuccessResponse(200, "otp send successfully", nil)
	c.JSON(http.StatusOK, response)
}

// Verify OTP godoc
//
//	@summary		api for Verify otp of user
//	@description	Enter otp
//	@tags			SignUp For User
//	@Param			inputs	body	req.OtpStruct{}	true	"Input Field"
//	@Router			/signup/verify_otp [post]
//	@Success		200	{object}	response.Response{}	"error while verifying otp"
//	@Failure		400	{object}	response.Response{}	"otp  successfully verified"
//
// verify otp
func (cr *UserHandler) VerifyOTP(c *gin.Context) {
	phonenumber, err := middlware.GetPhn(c, "Signup_Authorization")
	if err != nil {
		response := response.ErrorResponse(400, "error while getting id from cookie", err.Error(), phonenumber)
		c.JSON(400, response)
		return
	}

	var body req.OtpStruct
	if err := c.ShouldBindJSON(&body); err != nil {
		res := response.ErrorResponse(400, "error while getting otp from user", err.Error(), nil)
		utils.ResponseJSON(c, res)
		return
	}

	// verifying otp

	err1 := cr.userUsecase.VerifyOtp(c, phonenumber, body.OTP)

	if err1 != nil {
		res := response.ErrorResponse(400, "error while verifying otp", err.Error(), nil)
		c.JSON(http.StatusBadRequest, res)
		return
	}

	err2 := cr.userUsecase.UpdateStatus(c, user)

	if err2 != nil {
		res := response.ErrorResponse(400, "can't update status", err.Error(), nil)
		c.JSON(http.StatusBadRequest, res)
		return
	}

	response := response.SuccessResponse(200, "OTP verified successfully please update your details", "Continue with register")
	c.JSON(http.StatusOK, response)

}

// Registration godoc
//
//	@summary		api for complete registration
//	@description	Enter user details
//	@tags			SignUp For User
//	@Param			inputs	body	domain.Users{}	true	"Input Field"
//	@Router			/signup/register [post]
//	@Success		200	{object}	response.Response{}	"can't complete registration"
//	@Failure		400	{object}	response.Response{}	"Registration completed please login"
func (cr *UserHandler) Register(c *gin.Context) {
	// var user domain.Users
	if err := c.ShouldBindJSON(&user); err != nil {
		res := response.ErrorResponse(400, "error while getting  the data from user side", err.Error(), user)

		c.JSON(http.StatusBadRequest, res)
		return
	}

	users, err := cr.userUsecase.Register(c, user)
	if err != nil {
		res := response.ErrorResponse(400, "can't complete registration", err.Error(), user)

		c.JSON(http.StatusBadRequest, res)
		return
	}

	response := response.SuccessResponse(200, "Registration completed please login", users)
	c.JSON(http.StatusOK, response)

}

//>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>end>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>

//>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>userlogin>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>

// UserLogin godoc
//
//	@summary		api for user to login
//	@description	Enter user_name  with password
//	@security		ApiKeyAuth
//	@tags			User Login
//	@id				UserLogin
//	@Param			inputs	body	req.LoginStruct{}	true	"Input Field"
//	@Router			/login [post]
//	@Success		200	{object}	response.Response{}	"successfully logged in"
//	@Failure		400	{object}	response.Response{}	"invalid input"
//	@Failure		500	{object}	response.Response{}	"faild to generat JWT"
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

	c.SetCookie("User_Authorization", tokenString["accessToken"], 3600*24*30, "/", "", true, true)

	response := response.SuccessResponse(200, "Successfully logged in", message)

	c.JSON(http.StatusOK, response)

}

//user logout

func (cr *UserHandler) Logout(c *gin.Context) {
	c.SetCookie("User_Authorization", "", -1, "/ ", " ", false, true)
	response := response.SuccessResponse(200, "successfully logged out", nil)
	c.JSON(200, response)
}

//>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>list all products on user side>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
// func (uh *UserHandler) ListProducts(c *gin.Context) {

// 	products, err := uh.userUsecase.FindAllProducts(c, categoryname.CategoryName)
// 	if err != nil {
// 		response := response.ErrorResponse(400, "Error while fetching all products", err.Error(), products)
// 		c.JSON(400, response)
// 		return
// 	}

// 	response := response.SuccessResponse(200, "successfully displayed all products", products)
// 	c.JSON(200, response)
// }

// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>user management>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>.
func (uu *UserHandler) FindUserById(c *gin.Context) {
	//get id from getid
	id, err := middlware.GetId(c, "User_Authorization")
	if err != nil {
		response := response.ErrorResponse(400, "error while getting id from cookie", err.Error(), id)
		c.JSON(400, response)
		return
	}

	user, err := uu.userUsecase.FindUserById(c, int(id))
	if err != nil {
		response := response.ErrorResponse(400, "failed to find user", err.Error(), user)
		c.JSON(400, response)
		return
	}

	response := response.SuccessResponse(200, "successfully displayed user details", user)
	c.JSON(200, response)
}

func (uh *UserHandler) EditUserDetails(c *gin.Context) {
	var usr req.Usereditreq
	id, err := middlware.GetId(c, "User_Authorization")
	if err != nil {
		response := response.ErrorResponse(400, "error while getting id from cookie", err.Error(), id)
		c.JSON(400, response)
		return
	}

	//bind details by shouldbyjson
	if err := c.ShouldBindJSON(&usr); err != nil {
		res := response.ErrorResponse(400, "error while getting  the data from user side", err.Error(), usr)

		utils.ResponseJSON(c, res)
		return
	}

	user, err := uh.userUsecase.EditUserDetails(c, int(id), usr)
	if err != nil {
		response := response.ErrorResponse(400, "failed to edit user details", err.Error(), user)
		c.JSON(400, response)
		return
	}

	response := response.SuccessResponse(200, "successfully updated user details", user)
	c.JSON(200, response)
}

//>>>>>>>>>>>>>>>>>>>>>>>>address of user>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>

func (uh *UserHandler) AddAddress(c *gin.Context) {
	var address domain.Address
	id, err := middlware.GetId(c, "User_Authorization")
	if err != nil {
		response := response.ErrorResponse(400, "error while getting id from cookie", err.Error(), id)
		c.JSON(400, response)
		return
	}

	//bind details by shouldbyjson
	if err := c.ShouldBindJSON(&address); err != nil {
		res := response.ErrorResponse(400, "error while getting  the data from user side", err.Error(), address)

		utils.ResponseJSON(c, res)
		return
	}
	address.UserId = uint(id)
	ads, err := uh.userUsecase.AddAddress(c, address)
	if err != nil {
		response := response.ErrorResponse(400, "can't add address", err.Error(), ads)
		c.JSON(400, response)
		return
	}
	response := response.SuccessResponse(200, "successfully added address", ads)
	c.JSON(200, response)

}

func (uh *UserHandler) EditAddress(c *gin.Context) {
	var address domain.Address
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		res := response.ErrorResponse(400, "error while getting  the id ", err.Error(), address)

		utils.ResponseJSON(c, res)
		return
	}
	//bind details by shouldbyjson
	if err := c.ShouldBindJSON(&address); err != nil {
		res := response.ErrorResponse(400, "error while getting  the data from user side", err.Error(), address)

		utils.ResponseJSON(c, res)
		return
	}
	address.Address_Id = uint(id)
	ads, err := uh.userUsecase.EditAddress(c, address)
	if err != nil {
		response := response.ErrorResponse(400, "can't add address", err.Error(), ads)
		c.JSON(400, response)
		return
	}
	response := response.SuccessResponse(200, "successfully added address", ads)
	c.JSON(200, response)

}

func (uh *UserHandler) ListAddresses(c *gin.Context) {
	page, _ := strconv.Atoi(c.Query("page"))
	pagesize, _ := strconv.Atoi(c.Query("pagesize"))
	pagination := utils.Pagination{
		Page:     page,
		PageSize: pagesize,
	}
	id, err := middlware.GetId(c, "User_Authorization")
	if err != nil {
		response := response.ErrorResponse(400, "error while getting id from cookie", err.Error(), id)
		c.JSON(400, response)
		return
	}

	address, metadata, err := uh.userUsecase.ListAddresses(c, pagination, uint(id))
	if err != nil {
		response := response.ErrorResponse(400, "error while finding address", err.Error(), address)
		c.JSON(400, response)
		return
	}
	response := response.SuccessResponse(200, "successfully displayed all address", address, metadata)
	c.JSON(200, response)

}

func (uh *UserHandler) DeleteAddress(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("addressid"))
	err := uh.userUsecase.DeleteAddress(c, uint(id))
	if err != nil {
		response := response.ErrorResponse(400, "error while deleting address", err.Error(), "")
		c.JSON(400, response)
		return
	}
	response := response.SuccessResponse(200, "successfully deleted address", "")
	c.JSON(200, response)

}

// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>forgotpassword>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
// 1.sendotpto phn
func (uh *UserHandler) SendOtpForgotPass(c *gin.Context) {
	var phn req.Phn
	if err := c.ShouldBindJSON(&phn); err != nil {
		res := response.ErrorResponse(400, "error while getting  the data from user side", err.Error(), phn)

		c.JSON(http.StatusBadRequest, res)
		return
	}
	phnNo := phn.Phone

	err := uh.userUsecase.SendOtpForgotPass(c, phnNo)
	if err != nil {
		res := response.ErrorResponse(400, "error while sending otp", err.Error(), nil)
		c.JSON(http.StatusBadRequest, res)
		return
	}

	//generate tokenstring with jwt
	tokenString, err := auth.GenerateJWTPhn(phnNo)
	if err != nil {
		response := response.ErrorResponse(400, "failed to send otp", err.Error(), "user didn't exist")

		c.JSON(400, response)
		return
	}
	//set cookie

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Phone_Authorization", tokenString["accessToken"], 3600*24*30, "/", " ", false, true)

	response := response.SuccessResponse(200, "otp send successfully", nil)
	c.JSON(http.StatusOK, response)
}

func (cr *UserHandler) VerifyOTPForgotPass(c *gin.Context) {
	//bind body details
	var otp req.OtpStruct
	phonenumber, err := middlware.GetPhn(c, "Phone_Authorization")
	if err != nil {
		response := response.ErrorResponse(400, "error while getting id from cookie", err.Error(), phonenumber)
		c.JSON(400, response)
		return
	}

	if err := c.ShouldBindJSON(&otp); err != nil {
		res := response.ErrorResponse(400, "error while getting  the data from user side", err.Error(), otp)

		c.JSON(http.StatusBadRequest, res)
		return
	}
	code := otp.OTP
	if phonenumber == " " || code == " " {
		err := errors.New("please enter otp")
		res := response.ErrorResponse(400, "invalid otp", err.Error(), nil)
		utils.ResponseJSON(c, res)
		return
	}

	// verifying otp

	err1 := cr.userUsecase.VerifyOtpForgotpass(c, phonenumber, code)

	if err1 != nil {
		res := response.ErrorResponse(400, "error while verifying otp", err1.Error(), " ")
		c.JSON(http.StatusBadRequest, res)
		return
	}

	response := response.SuccessResponse(200, "OTP verified successfully please update your password")
	c.JSON(http.StatusOK, response)

}
func (ur *UserHandler) ForgotPassword(c *gin.Context) {
	var pass req.Pass
	usrphn, err := middlware.GetPhn(c, "Phone_Authorization")
	if err != nil {
		response := response.ErrorResponse(400, "error while getting id from cookie", err.Error(), usrphn)
		c.JSON(400, response)
		return
	}

	if err := c.ShouldBindJSON(&pass); err != nil {
		res := response.ErrorResponse(400, "error while getting  the data from user side", err.Error(), pass)

		c.JSON(http.StatusBadRequest, res)
		return
	}
	newpass := pass.Password
	err1 := ur.userUsecase.ForgotPassword(c, usrphn, newpass)
	if err1 != nil {
		res := response.ErrorResponse(400, "error while getting  updating password", err.Error(), "try again")

		c.JSON(http.StatusBadRequest, res)
		return
	}

	c.SetCookie("Phone_Authorization", "", -1, " /", " ", false, true)
	response := response.SuccessResponse(200, "successsfully updated password please login")
	c.JSON(200, response)

}
