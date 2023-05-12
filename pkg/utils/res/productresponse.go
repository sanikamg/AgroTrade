package res

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
