package domain

import (
	"time"

	"gorm.io/gorm"
)

// sample product model
type ProductDetails struct {
	gorm.Model
	Product_Id      uint     `gorm:"serial primarykey;autoIncrement:true;unique"`
	ProductName     string   `json:"productname" validate:"required,min=3,max=12"`
	ProductPrice    uint     `json:"productprice"`
	ProductQuantity uint     `json:"productquantity"`
	CategoryID      uint     `json:"categoryid"`
	Category        Category `gorm:"foreignkey:CategoryID"`
	DiscountPrice   uint     `json:"discount_price"`
	AdminId         uint
}

// for all category
type Category struct {
	ID           uint   `gorm:"serial primarykey;autoIncrement:true;unique"`
	CategoryName string `json:"category_name" gorm:"unique;not null"`
}

type Image struct {
	ID        uint   `json:"id" gorm:"primaryKey"`
	ProductId uint   `json:"product_id"`
	Image     string `JSON:"Image" `
}

// cart

type Cart_item struct {
	Cart_Id     uint    ` gorm:" serial primaryKey;autoIncrement:true;unique"`
	User_Id     uint    `json:"user_id"   `
	Product_Id  uint    `json:"product_id" `
	Quantity    uint    `json:"quantity"`
	Total_Price float32 `json:"total_price"`
}

// order
type Order struct {
	Order_Id          uint      ` gorm:" serial primaryKey;autoIncrement:true;unique"`
	User_Id           uint      `json:"user_id"  gorm:"not null" `
	Applied_Coupon_id uint      `json:"applied_coupon_id,omitempty"`
	Total_Amount      float64   `json:"total_amount"  gorm:"not null" `
	PaymentMethodID   uint      `json:"paymentmethod_id"  gorm:"not null" `
	Payment_Status    string    `json:"payment_status"`
	Order_Status      string    `json:"order_status"`
	Address_Id        uint      `json:"address_id" `
	OrderDate         time.Time `json:"order_date"`
}

type PaymentMethod struct {
	Method_id     uint    ` gorm:" serial primaryKey;autoIncrement:true;unique"`
	PaymentMethod string  `json:"paymentmethod"`
	MaximumAmount float64 `json:"maximumamount"`
}

//coupon

type Coupon struct {
	gorm.Model

	Coupon_Id uint   `gorm:"serial primaryKey;autoIncrement:true;unique"`
	Coupon    string `json:"coupon"`
	Discount  int    `json:"discount"`
	Validity  int64  `json:"validity"`
}

type Applied_Coupons struct {
	UserID      uint
	Coupon_Code string `json:"coupon_code"`
}

// return
type OrderReturn struct {
	ID           uint      `gorm:"serial primaryKey;autoIncrement:true;unique"`
	OrderID      uint      `json:"order_id" gorm:"not null;unique"`
	RequestDate  time.Time `json:"request_date" gorm:"not null"`
	ReturnReason string    `json:"return_reason" gorm:"not null"`
	RefundAmount float64   `json:"refund_amount" gorm:"not null"`
	IsApproved   bool      `json:"is_approved"`
	ReturnDate   time.Time `json:"return_date"`
	ReturnStatus string    `json:"return_status"`
}
