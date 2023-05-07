package domain

type Users struct {
	User_Id  uint   `gorm:"serial primarykey;autoIncrement:true;unique"`
	Username string `json:"username" validate:"required,min=3,max=12"`
	Name     string `json:"name" validate:"required,min=3,max=12" `
	Phone    string `json:"phone" gorm:"unique" binding:"required,min=10,max=10"`
	Email    string `json:"email" validate:"required,min=3,max=12" `
	Password string `json:"password" validate:"required,min=8,max=64" `
}
