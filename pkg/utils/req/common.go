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

//to block user using userid
type Block struct {
	UserID uint `json:"user_id" binding:"required,numeric"`
}

type BlockStatus struct {
	UserID      uint `json:"user_id" binding:"required,numeric"`
	BlockStatus bool `json:"blockstatus"`
}

type ForgotPassword struct {
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

type Phn struct {
	Phone string `json:"phone"`
}

type Pass struct {
	Password string `json:"password"`
}
