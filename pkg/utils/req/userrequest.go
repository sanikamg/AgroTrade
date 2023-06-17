package req

type Categoryreq struct {
	CategoryName string `json:"categoryname"`
}

type Usereditreq struct {
	Username string `json:"username" `
	Name     string `json:"name" `
	Phone    string `json:"phone"`
	Email    string `json:"email" `
	Password string `json:"password" `
}
