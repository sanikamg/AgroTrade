package handler

import (
	"golang_project_ecommerce/pkg/api/middlware"
	"golang_project_ecommerce/pkg/common/response"
	"golang_project_ecommerce/pkg/domain"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func (pd *ProductHandler) ReturnOrder(c *gin.Context) {
	var returnOrder domain.OrderReturn

	order_id, err := strconv.Atoi(c.Query("orderId"))
	if err != nil {
		response := response.ErrorResponse(400, "Please add order id as params", err.Error(), "")
		c.JSON(400, response)
		return
	}
	//get id from getid
	id, err := middlware.GetId(c, "User_Authorization")
	if err != nil {
		response := response.ErrorResponse(400, "error while getting id from cookie", err.Error(), " ")
		c.JSON(400, response)
		return
	}

	err1 := pd.productUsecase.VerifyOrderID(c, uint(id), uint(order_id))
	if err1 != nil {
		response := response.ErrorResponse(400, "invalid order_id", err1.Error(), id)
		c.JSON(400, response)
		return
	}

	returnOrder.OrderID = uint(order_id)
	returnOrder.RequestDate = time.Now()
	returnOrder.ReturnReason = c.Query("reason")
	returnOrder.ReturnStatus = "Return Requested"
	//finding total amount by orderid
	total_amount, err := pd.productUsecase.FindTotalAmountByOrderId(c, uint(order_id))
	if err != nil {
		response := response.ErrorResponse(400, "failed to find refund amount", err.Error(), "")
		c.JSON(400, response)
		return
	}
	returnOrder.RefundAmount = total_amount
	returnResp, err := pd.productUsecase.ReturnRequest(c, returnOrder)
	if err != nil {
		response := response.ErrorResponse(400, "failed to return order", err.Error(), "")
		c.JSON(400, response)
		return
	}

	response := response.SuccessResponse(200, "successfully requsted to return products", returnResp)
	c.JSON(200, response)

}
