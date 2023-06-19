package verification

import (
	"errors"
	"golang_project_ecommerce/pkg/config"

	"github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/verify/v2"
)

var client *twilio.RestClient

var AccountSid, AuthToken, VerifyServiceSid string

func InitTwilio(cn config.Config) {
	AccountSid, AuthToken, VerifyServiceSid = config.GetTwilioconfig()

	client = twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: AccountSid,
		Password: AuthToken,
	})
}

func SendOtp(phone string) (string, error) {

	params := &openapi.CreateVerificationParams{}
	params.SetTo(phone)
	params.SetChannel("sms")

	_, err := client.VerifyV2.CreateVerification(VerifyServiceSid, params)
	if err != nil {
		return "Error while sending otp", err
	}

	return "otp send successfully", nil
}

// This function waits for you to input the OTP sent to your phone,
// and validates that the code is approved
func VerifyOtp(phone string, code string) error {

	params := &openapi.CreateVerificationCheckParams{}
	params.SetTo(phone)
	params.SetCode(code)

	resp, err := client.VerifyV2.CreateVerificationCheck(VerifyServiceSid, params)

	if err != nil {

		return err

	} else if *resp.Status == "approved" {

		return nil
	} else if *resp.Status == "pending" {

		return errors.New("otp verification failed")
	}

	return nil
}
