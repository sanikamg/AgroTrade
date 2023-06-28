package req

type UpdateProduct struct {
	ProductId       uint   `json:"productid"`
	ProductName     string `json:"productname"`
	ProductPrice    uint   `json:"productprice"`
	ProductQuantity uint   `json:"productquantity"`
	Categoryid      uint   `json:"categoryid"`
}

type DeleteCategory struct {
	CategoryName string `json:"category_name" gorm:"unique;not null"`
}

type UpdateOrder struct {
	Order_Id        uint ` json:"order_id"`
	PaymentMethodID uint `json:"paymentmethod_id"  gorm:"not null" `
	Address_Id      uint `json:"address_id" `
}

type ReqRazorpayVeification struct {
	RazorpayPaymentID string `json:"razorpay_payment_id"`
	RazorpayOrderID   string `json:"razorpay_order_id"`
	RazorpaySignature string `json:"razorpay_signature"`
	UserID            uint   `json:"user_id"`
}
type RazorPayReq struct {
	Total_Amount float64
	Email        string
	Phone        string
}
