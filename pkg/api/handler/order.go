package handler

import (
	"golang_project_ecommerce/pkg/api/middlware"
	"golang_project_ecommerce/pkg/common/response"
	"golang_project_ecommerce/pkg/domain"
	"golang_project_ecommerce/pkg/utils"
	"golang_project_ecommerce/pkg/utils/req"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Create order godoc
// @summary api for create order
// @description Enter address id and method id
// @tags Create Order
// @Param page query int true "address_id" format(int32)
// @Param pagesize query int true "paymentmethod_id" format(int32)
// @Router /order/create [post]
// @Success 200 {object} response.Response{} "successfully  created order"
// @Failure 400 {object} response.Response{}  "failed to create order"
func (pd *ProductHandler) CreateOrder(c *gin.Context) {
	var order domain.Order
	addressid, _ := strconv.Atoi(c.Query("address_id"))
	PaymentMetodId, _ := strconv.Atoi(c.Query("paymentmethod_id"))
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
	order.PaymentMethodID = uint(PaymentMetodId)
	order.Payment_Status = "Pending"
	order.Order_Status = "order placed"

	orderresp, err := pd.productUsecase.CreateOrder(c, order)
	if err != nil {
		response := response.ErrorResponse(400, "failed to create order", err.Error(), "try again")
		c.JSON(400, response)
		return
	}
	response := response.SuccessResponse(200, "Successfully placed order please complete payment", orderresp)
	c.JSON(200, response)
}

func (pd *ProductHandler) UpdateOrder(c *gin.Context) {
	var UpdateOrderDetails req.UpdateOrder
	if err := c.ShouldBindJSON(&UpdateOrderDetails); err != nil {
		response := response.ErrorResponse(400, "error while getting data from users", err.Error(), UpdateOrderDetails)
		c.JSON(400, response)
		return
	}

	uporder, err := pd.productUsecase.UpdateOrderDetails(c, UpdateOrderDetails)
	if err != nil {
		response := response.ErrorResponse(400, "error while updating data", err.Error(), UpdateOrderDetails)
		c.JSON(400, response)
		return
	}
	response := response.SuccessResponse(200, "Successfully updated orde", uporder)
	c.JSON(200, response)
}

func (pd *ProductHandler) ListAllOrders(c *gin.Context) {
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		response := response.ErrorResponse(400, "Please add page number as params", err.Error(), "")
		c.JSON(400, response)
	}
	pagesize, err := strconv.Atoi(c.Query("pagesize"))
	if err != nil {
		response := response.ErrorResponse(400, "Please add pages size as params", err.Error(), "")
		c.JSON(400, response)
	}
	pagination := utils.Pagination{
		Page:     page,
		PageSize: pagesize,
	}
	usrid, err := middlware.GetId(c, "User_Authorization")
	if err != nil {
		response := response.ErrorResponse(400, "user authentication failed", err.Error(), "")
		c.JSON(400, response)
		return
	}
	orderResp, metadata, err := pd.productUsecase.ListAllOrders(c, pagination, uint(usrid))
	if err != nil {
		response := response.ErrorResponse(400, "error while finding orderss", err.Error(), orderResp)
		c.JSON(400, response)
	}
	response := response.SuccessResponse(200, "successfully displayed all orders", orderResp, metadata)
	c.JSON(200, response)
}

// to display all orders for admin
func (pd *ProductHandler) GetAllOrders(c *gin.Context) {
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		response := response.ErrorResponse(400, "Please add page number as params", err.Error(), "")
		c.JSON(400, response)
	}
	pagesize, err := strconv.Atoi(c.Query("pagesize"))
	if err != nil {
		response := response.ErrorResponse(400, "Please add pages size as params", err.Error(), "")
		c.JSON(400, response)
	}
	pagination := utils.Pagination{
		Page:     page,
		PageSize: pagesize,
	}
	orderResp, metadata, err := pd.productUsecase.GetAllOrders(c, pagination)
	if err != nil {
		response := response.ErrorResponse(400, "error while finding orders", err.Error(), orderResp)
		c.JSON(400, response)
		return
	}
	response := response.SuccessResponse(200, "successfully displayed all orders", orderResp, metadata)
	c.JSON(200, response)
}

func (ph *ProductHandler) CancelOrder(c *gin.Context) {
	order_id, err := strconv.Atoi(c.Query("order_id"))
	if err != nil {
		response := response.ErrorResponse(400, "Please add id as params", err.Error(), order_id)
		c.JSON(400, response)
		return
	}

	err1 := ph.productUsecase.DeleteOrder(c, uint(order_id))
	if err1 != nil {
		response := response.ErrorResponse(400, "can't delete order", err.Error(), "")
		c.JSON(400, response)
		return
	}
	response := response.SuccessResponse(200, "successfully deleted order")
	c.JSON(200, response)
}

//>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>Check out>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>

func (pd *ProductHandler) PlaceOrder(c *gin.Context) {
	var order domain.Order
	order_id, _ := strconv.Atoi(c.Query("order_id"))
	coupon_id, _ := strconv.Atoi(c.Query("coupon_id"))
	order.Order_Id = uint(order_id)
	order.Applied_Coupon_id = uint(coupon_id)
	couponResp, err := pd.productUsecase.ValidateCoupon(c, order.Applied_Coupon_id)
	if err != nil {
		response := response.ErrorResponse(400, "Invalid coupon", err.Error(), "try with a valid coupon")
		c.JSON(400, response)
		return
	} else {
		totalamnt, err := pd.productUsecase.ApplyDiscount(c, couponResp, uint(order_id))
		if err != nil {
			response := response.ErrorResponse(400, "Add more quantity", err.Error(), "try again")
			c.JSON(400, response)
			return
		}
		order.Total_Amount = float64(totalamnt)
	}
	paymentResp, err := pd.productUsecase.PlaceOrder(c, order)
	if err != nil {
		response := response.ErrorResponse(400, "failed to place order", err.Error(), "")
		c.JSON(400, response)
		return
	}
	response := response.SuccessResponse(200, "Successfully  placed order complete payment details", paymentResp)
	c.JSON(200, response)
}

func (pd *ProductHandler) CheckOut(c *gin.Context) {
	var razorPay req.RazorPayReq
	order_id, err := strconv.Atoi(c.Query("order_id"))
	if err != nil {
		response := response.ErrorResponse(400, "Please add order_id  as params", err.Error(), "")
		c.JSON(400, response)
		return
	}
	payment_method_id, err := pd.productUsecase.FindPaymentMethodIdByOrderId(c, uint(order_id))
	if err != nil {
		response := response.ErrorResponse(400, "failed to find payment method", err.Error(), "")
		c.JSON(400, response)
		return
	}
	if payment_method_id == 1 {
		orderREsp, err := pd.productUsecase.UpdateOrderStatus(c, uint(order_id))
		if err != nil {
			response := response.ErrorResponse(400, "failed to place order", err.Error(), "")
			c.JSON(400, response)
			return
		}
		response := response.SuccessResponse(200, "Successfully  confirmed order", orderREsp)
		c.JSON(200, response)
	} else {
		id, err := middlware.GetId(c, "User_Authorization")
		if err != nil {
			response := response.ErrorResponse(400, "error while getting id from cookie", err.Error(), id)
			c.JSON(400, response)
			return
		}
		totalAmount, err := pd.productUsecase.FindTotalAmountByOrderId(c, uint(order_id))
		if err != nil {
			response := response.ErrorResponse(400, "error while getting total amount", err.Error(), id)
			c.JSON(400, response)
			return
		}
		razorPay.Total_Amount = totalAmount
		phnEmail, err := pd.productUsecase.FindPhnEmailByUsrId(c, int(id))
		if err != nil {
			response := response.ErrorResponse(400, "error while getting details", err.Error(), id)
			c.JSON(400, response)
			return
		}
		razorPay.Email = phnEmail.Email
		razorPay.Phone = phnEmail.Phone

		razorpayOrder, err := pd.productUsecase.GetRazorpayOrder(c, uint(id), razorPay)
		if err != nil {
			response := response.ErrorResponse(400, "faild to create razorpay order ", err.Error(), nil)
			c.JSON(400, response)
			return
		}
		c.HTML(200, "payment.html", razorpayOrder)
		// response := response.SuccessResponse(200, "Payment Done Successfully", razorpayOrder)
		// c.JSON(200, response)
	}

}
