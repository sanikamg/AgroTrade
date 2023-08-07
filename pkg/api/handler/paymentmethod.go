package handler

import (
	"golang_project_ecommerce/pkg/common/response"
	"golang_project_ecommerce/pkg/domain"
	"golang_project_ecommerce/pkg/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

// PaymentMethod godoc
//
//	@summary		api for add payment method by admin
//	@description	Enter payment method and maximum amount
//	@tags			payment method
//	@Param			inputs	body	domain.PaymentMethod{}	true	"Input Field"
//	@Router			/paymentmethod/add [post]
//	@Success		200	{object}	response.Response{}	"successfully  added payment method"
//	@Failure		400	{object}	response.Response{}	"failed to add payment method"
func (ph *ProductHandler) AddpaymentMethod(c *gin.Context) {
	var payment domain.PaymentMethod
	if err := c.ShouldBindJSON(&payment); err != nil {
		response := response.ErrorResponse(400, "error while fetching data from user", err.Error(), payment)
		c.JSON(400, response)
		return
	}
	paymentresp, err1 := ph.productUsecase.AddPaymentMethod(c, payment)
	if err1 != nil {
		response := response.ErrorResponse(400, "can't add payment method", err1.Error(), paymentresp)
		c.JSON(400, response)
		return
	}
	response := response.SuccessResponse(200, "successfully added payment method", paymentresp)
	c.JSON(200, response)
}

//	 PaymentMethods godoc
//
//		@Summary		Get all paymentmethods
//		@Description	Get all products
//		@tags			payment method
//		@Param			page		query	int	true	"Page"		format(int32)
//		@Param			pagesize	query	int	true	"Page Size"	format(int32)
//		@Router			/paymentmethod/view [get]
//		@Success		200	{object}	response.Response{}	"successfully  displayed all prioducts"
//		@Failure		400	{object}	response.Response{}	"ferror while getting data"
//
// to list all payment methods
func (ph *ProductHandler) GetAllPaymentMethods(c *gin.Context) {
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		response := response.ErrorResponse(400, "Please add page number as params", err.Error(), "")
		c.JSON(400, response)
		return
	}
	pagesize, err := strconv.Atoi(c.Query("pagesize"))
	if err != nil {
		response := response.ErrorResponse(400, "Please add pages size as params", err.Error(), "")
		c.JSON(400, response)
		return
	}
	pagination := utils.Pagination{
		Page:     page,
		PageSize: pagesize,
	}
	paymentResp, metadata, err := ph.productUsecase.GetPaymentMethods(c, pagination)
	if err != nil {
		response := response.ErrorResponse(400, "Error while getting data", err.Error(), "")
		c.JSON(400, response)
		return
	}
	response := response.SuccessResponse(200, "successfully dispalyed all products", paymentResp, metadata)
	c.JSON(200, response)

}

// PaymentMethods godoc
//
//	@summary		api for update payment method by admin
//	@description	Enter payment method and maximum amount with id
//	@tags			payment method
//	@Param			inputs	body	domain.PaymentMethod{}	true	"Input Field"
//	@Router			/paymentmethod/update [patch]
//	@Success		200	{object}	response.Response{}	"successfully  updated payment method"
//	@Failure		400	{object}	response.Response{}	"failed to updatepayment method"
func (ph *ProductHandler) UpdatePaymentMethod(c *gin.Context) {
	var payment domain.PaymentMethod
	if err := c.BindJSON(&payment); err != nil {
		response := response.ErrorResponse(400, "Error while getting data from admin side", err.Error(), payment)
		c.JSON(400, response)
		return
	}
	paymentresp, err := ph.productUsecase.UpdatePaymentMethod(c, payment)
	if err != nil {
		response := response.ErrorResponse(400, "can't update data", err.Error(), "")
		c.JSON(400, response)
		return
	}

	response := response.SuccessResponse(200, "successfully updated product", paymentresp)
	c.JSON(200, response)
}

// PaymentMethods godoc
//
//	@Summary		delete paymentmethod
//	@Description	Delete payment methods
//	@tags			payment method
//	@Param			page	query	int	true	"id"	format(int32)
//	@Router			/paymentmethod/delete [delete]
//	@Success		200	{object}	response.Response{}	"successfully  deleted method"
//	@Failure		400	{object}	response.Response{}	"failed to delete method"
//
// to list all payment methods
func (ph *ProductHandler) DeleteMethod(c *gin.Context) {
	method_id, err := strconv.Atoi(c.Query("method_id"))
	if err != nil {
		response := response.ErrorResponse(400, "Please add id as params", err.Error(), method_id)
		c.JSON(400, response)
		return
	}

	err1 := ph.productUsecase.DeleteMethod(c, uint(method_id))
	if err1 != nil {
		response := response.ErrorResponse(400, "can't delete payment method", err.Error(), "")
		c.JSON(400, response)
		return
	}
	response := response.SuccessResponse(200, "successfully deleted method")
	c.JSON(200, response)
}
