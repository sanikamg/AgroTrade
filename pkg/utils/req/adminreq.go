package req

type DeleteId struct {
	ProductID uint `json:"productid" binding:"required,numeric"`
}
