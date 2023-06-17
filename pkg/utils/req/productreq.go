package req

type UpdateProduct struct {
	ProductId       uint   `json:"productid"`
	ProductName     string `json:"productname"`
	ProductPrice    uint   `json:"productprice"`
	ProductQuantity uint   `json:"productquantity"`
}

type DeleteCategory struct {
	CategoryName string `json:"category_name" gorm:"unique;not null"`
}
