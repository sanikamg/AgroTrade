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

type CartResponse struct {
	Product_Id  uint    `json:"product_id" `
	Quantity    uint    `json:"quantity"`
	Total_Price float32 `json:"total_price"`
	Image       []string
}

type OrderResponse struct {
	Total_Amount float64 `json:"total_amount"  gorm:"not null" `
	Order_Status string  `json:"order_status"`
	Address_Id   uint    `json:"address_id" `
}

type PaymentResponse struct {
	Total_Amount float64 `json:"total_amount"  gorm:"not null" `
	//Balance_Amount int     `json:"balance_amount"`
	PaymentMethod  string `json:"paymentmethod"  gorm:"not null" `
	Payment_Status string `json:"payment_status"   `
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
