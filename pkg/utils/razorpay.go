package utils

import (
	"fmt"
	"golang_project_ecommerce/pkg/config"

	razorpay "github.com/razorpay/razorpay-go"
)

func GenerateRazorpayOrder(razorPayAmount uint, recieptIdOptional string) (razorpayOrderID interface{}, err error) {
	// get razor pay key and secret
	razorpayKey, razorpaySecret := config.GetRazorPayConfig()

	//create a razorpay client
	client := razorpay.NewClient(razorpayKey, razorpaySecret)

	data := map[string]interface{}{
		"amount":   razorPayAmount,
		"currency": "INR",
		"receipt":  recieptIdOptional,
	}
	// create an order on razor pay
	razorpayRes, err := client.Order.Create(data, nil)
	if err != nil {
		return razorpayOrderID, fmt.Errorf("failed to create razorpay order for amount %v", razorPayAmount)
	}

	razorpayOrderID = razorpayRes["id"]

	return razorpayOrderID, nil
}
