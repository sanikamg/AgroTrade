package res

import "time"

type SalesReport struct {
	UserID          uint      `json:"user_id"`
	Name            string    `json:"name"`
	Email           string    `json:"email"`
	OrderDate       time.Time `json:"order_date"`
	OrderTotalPrice float64   `json:"order_total_price"`

	OrderStatus   string `json:"order_status"`
	PaymentType   string `json:"payment_type"`
	PaymentStatus string `json:"payment_status"`
}
