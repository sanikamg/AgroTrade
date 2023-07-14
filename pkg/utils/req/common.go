package req

import (
	"fmt"
	"regexp"
)

// login struct for user and admin

type LoginStruct struct {
	Username string `json:"username" validate:"required,min=3,max=12"`
	Password string `json:"password" validate:"required,min=8,max=64" `
}

//otp struct for otp verification

type OtpStruct struct {
	OTP string `json:"otp" validate:"required,min=6,max=6"`
}

// to block user using userid
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

type Sales struct {
	Sdate string `json:"Sdate" validate:"required,Sdate"`
	Edate string `json:"Edate" validate:"required,Edate"`
}

func (s *Sales) Validate() error {
	dateRegex := `^\d{4}-\d{2}-\d{2}$`
	match, err := regexp.MatchString(dateRegex, s.Sdate)
	if err != nil {
		return err
	}
	if !match {
		return fmt.Errorf("Sdate should be in the yyyy-mm-dd format")
	}

	match, err = regexp.MatchString(dateRegex, s.Edate)
	if err != nil {
		return err
	}
	if !match {
		return fmt.Errorf("Edate should be in the yyyy-mm-dd format")
	}

	return nil
}
