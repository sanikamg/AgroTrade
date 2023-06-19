package domain

type Users struct {
	User_Id      uint   `gorm:"serial primarykey;autoIncrement:true;unique"`
	Username     string `json:"username" validate:"required,min=3,max=12"`
	Name         string `json:"name" validate:"required,min=3,max=12" `
	Phone        string `json:"phone" gorm:"unique" binding:"required,min=10,max=10"`
	Email        string `json:"email" validate:"required,min=3,max=12" `
	BlockStatus  bool   `json:"block_status" gorm:"not null;default:false"`
	Password     string `json:"password" validate:"required,min=8,max=64" `
	Verification bool   `json:"verification"`
}

type Address struct {
	Address_Id uint `gorm:"serial primarykey;autoIncrement:true;unique"`
	UserId     uint `json:"userid"`
	//User         Users  `gorm:"foreignkey:UserId"`
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
