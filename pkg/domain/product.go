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
	Order_Id uint ` gorm:" serial primaryKey;autoIncrement:true;unique"`
	User_Id  uint `json:"user_id"  gorm:"not null" `

	//Applied_Coupons string `json:"applied_coupons"  `
	//Discount        uint   `json:"discount"`
	Total_Amount float64 `json:"total_amount"  gorm:"not null" `
	//Balance_Amount int     `json:"balance_amount"`
	PaymentMethod  string `json:"paymentmethod"  gorm:"not null" `
	Payment_Status string `json:"payment_status"   `
	//Payment_Id     string  `json:"payment_id"`
	Order_Status string `json:"order_status"`
	Address_Id   uint   `json:"address_id" `
}

type PaymentMethod struct {
	COD bool
}

//coupon

type Coupon struct {
	Created_At time.Time
	Coupon_Id  uint   `gorm:"serial primaryKey;autoIncrement:true;unique"`
	Coupon     string `json:"coupon"`
	Discount   int    `json:"discount"`
	Quantity   int    `json:"quantity"`
	Validity   int64  `json:"validity"`
}

type Applied_Coupons struct {
	UserID      uint
	Coupon_Code string `json:"coupon_code"`
}
