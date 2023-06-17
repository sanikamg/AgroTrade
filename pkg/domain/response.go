package domain

type ProductResponse struct {
	Product_Id      uint
	ProductName     string
	ProductPrice    uint
	Image           []string
	ProductQuantity uint
	Category_name   string
	DiscountPrice   uint
}
