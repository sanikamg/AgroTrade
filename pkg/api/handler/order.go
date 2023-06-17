package handler

import (
	"golang_project_ecommerce/pkg/api/middlware"
	"golang_project_ecommerce/pkg/common/response"
	"golang_project_ecommerce/pkg/domain"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (pd *ProductHandler) CreateOrder(c *gin.Context) {
	var order domain.Order
	addressid, _ := strconv.Atoi(c.Query("addressid"))
	usrid, err := middlware.GetId(c, "User_Authorization")
	if err != nil {
		response := response.ErrorResponse(400, "user authentication failed", err.Error(), "")
		c.JSON(400, response)
		return
	}
	order.User_Id = uint(usrid)
	//get total amount from cart
	totalamnt, err := pd.productUsecase.GetTotalAmount(c, uint(usrid))
	if err != nil {
		response := response.ErrorResponse(400, "failed to get total amount", err.Error(), "")
		c.JSON(400, response)
		return
	}
	order.Total_Amount = totalamnt
	order.Address_Id = uint(addressid)
	order.Order_Status = "payment pending"
	orderresp, err := pd.productUsecase.CreateOrder(c, order)
	if err != nil {
		response := response.ErrorResponse(400, "failed to create order", err.Error(), "try again")
		c.JSON(400, response)
		return
	}
	response := response.SuccessResponse(200, "Successfully placed order please complete payment", orderresp)
	c.JSON(200, response)
}

func (pd *ProductHandler) PlaceOrder(c *gin.Context) {
	var order domain.Order
	orderid, _ := strconv.Atoi(c.Query("orderid"))
	order.Order_Id = uint(orderid)
	payment := c.Query("paymentmethod")
	if payment != "COD" {
		response := response.ErrorResponse(400, "Please enter a valid payment method", " ", "")
		c.JSON(400, response)
		return
	}

	order.PaymentMethod = payment
	if payment == "COD" {
		order.Payment_Status = "Pending"
	}

	order.Order_Status = "order placed"
	paymentResp, err := pd.productUsecase.PlaceOrder(c, order)
	if err != nil {
		response := response.ErrorResponse(400, "failed to place order", err.Error(), "")
		c.JSON(400, response)
		return
	}
	response := response.SuccessResponse(200, "Successfully  complete payment details", paymentResp)
	c.JSON(200, response)
}
