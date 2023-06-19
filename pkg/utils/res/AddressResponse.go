package res

type AddressResponse struct {
	Address_Id   uint   `gorm:"serial primarykey;autoIncrement:true;unique"`
	User_Id      uint   `json:"user_id"`
	Full_name    string `json:"full_name"`
	Phone_number string `json:"phone_number"`
	House_name   string `json:"house_name"`
	Post_office  string `json:"post_office"`
	Area         string `json:"area"`
	Landmark     string `json:"landmark"`
	District     string `json:"district"`
	State        string `json:"state"`
	Pin          string `json:"pin"`
}
