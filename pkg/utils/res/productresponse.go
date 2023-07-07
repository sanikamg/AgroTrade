package res

import "time"

type AllCategories struct {
	ID uint `gorm:"serial primarykey;autoIncrement:true;unique"`

	CategoryName string `json:"category_name" gorm:"unique;not null" binding:"required,min=1,max=30"`
}

type AllProducts struct {
	ProductName     string `json:"productname" validate:"required,min=3,max=12"`
	ProductPrice    uint   `json:"productprice"`
	ProductQuantity uint   `json:"productquantity"`
	Image           string `json:"image" gorm:"not null"`
}

type CartResponse struct {
	Product_Id  uint    `json:"product_id" `
	Quantity    uint    `json:"quantity"`
	Total_Price float32 `json:"total_price"`
	Image       []string
}

type OrderResponse struct {
	Order_id       uint    `json:"order_id"`
	Total_Amount   float64 `json:"total_amount"  gorm:"not null" `
	Order_Status   string  `json:"order_status"`
	Payment_Status string  `json:"payment_status"   `
	Address_Id     uint    `json:"address_id" `
	Payment_Method string  `json:"payment_method"`
}

type PaymentResponse struct {
	Total_Amount float64 `json:"total_amount"  gorm:"not null" `
	//Balance_Amount int     `json:"balance_amount"`
	PaymentMethodId string `json:"payment_method_id"  gorm:"not null" `
	Payment_Status  string `json:"payment_status"   `
	//Payment_Id     string  `json:"payment_id"`
	Order_Status string `json:"order_status"`
	Address_Id   uint   `json:"address_id" `
}

type OrderedItems struct {
	User_id      uint
	Product_id   uint
	Order_Id     string
	Product_Name string
	Image        []string
	Price        int
	Quantity     uint
}

type PaymentMethodResponse struct {
	Method_Id     uint
	PaymentMethod string
	MaximumAmount float64
}

type CouponResponse struct {
	Discount int
	Validity int64
}
type CouponList struct {
	Coupon_id uint
	Discount  int
	Validity  string
}

// razorpay
type ResRazorpayOrder struct {
	RazorpayKey     string      `json:"razorpay_key"`
	UserID          uint        `json:"user_id"`
	AmountToPay     uint        `json:"amount_to_pay"`
	RazorpayOrderID interface{} `json:"razorpay_order_id"`
	Email           string      `json:"email"`
	Phone           string      `json:"phone"`
}

type PhnEmailResp struct {
	Phone string
	Email string
}

type ReturnResponse struct {
	ID           uint      `gorm:"serial primaryKey;autoIncrement:true;unique"`
	OrderID      uint      `json:"order_id" gorm:"not null;unique"`
	RequestDate  time.Time `json:"request_date" gorm:"not null"`
	ReturnReason string    `json:"return_reason" gorm:"not null"`
	RefundAmount float64   `json:"refund_amount" gorm:"not null"`
	ReturnStatus string    `json:"return_status"`
}
