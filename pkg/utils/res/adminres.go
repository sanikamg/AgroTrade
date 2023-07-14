package res

import "time"

type SalesReport struct {
	UserID          uint      `json:"user_id"`
	Name            string    `json:"name"`
	Email           string    `json:"email,omitempty"`
	OrderDate       time.Time `json:"order_date,omitempty"`
	OrderTotalPrice float64   `json:"order_total_price,omitempty"`
	OrderStatus     string    `json:"order_status,omitempty"`
	DeliveryStatus  string    `json:"delivery_status,omitempty"`
	PaymentType     string    `json:"payment_type,omitempty"`
	PaymentStatus   string    `json:"payment_status,omitempty"`
}
