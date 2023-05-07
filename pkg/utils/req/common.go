package req

// login struct for user and admin

type LoginStruct struct {
	Username string `json:"username" validate:"required,min=3,max=12"`
	Password string `json:"password" validate:"required,min=8,max=64" `
}

//otp struct for otp verification

type OtpStruct struct {
	OTP string `json:"otp" validate:"required,min=6,max=6"`
}
