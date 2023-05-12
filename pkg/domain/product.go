package domain

//sample product model
type ProductDetails struct {
	Product_Id      uint     `gorm:"serial primarykey;autoIncrement:true;unique"`
	ProductName     string   `json:"productname" validate:"required,min=3,max=12"`
	ProductPrice    uint     `json:"productprice"`
	ProductQuantity uint     `json:"productquantity"`
	CategoryID      uint     `json:"categoryid"`
	Category        Category `gorm:"foreignkey:CategoryID"`
	DiscountPrice   uint     `json:"discount_price"`
	Image           string   `json:"image" gorm:"not null"`
	AdminId         uint
}

// for all category
type Category struct {
	ID           uint   `gorm:"serial primarykey;autoIncrement:true;unique"`
	CategoryName string `json:"category_name" gorm:"unique;not null"`
}
