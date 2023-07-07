package handler

import (
	"golang_project_ecommerce/pkg/common/response"
	"golang_project_ecommerce/pkg/domain"
	"golang_project_ecommerce/pkg/utils"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func (ph *ProductHandler) AddCoupon(c *gin.Context) {
	var coupon domain.Coupon
	if err := c.ShouldBindJSON(&coupon); err != nil {
		response := response.ErrorResponse(400, "error while getting data from user side", err.Error(), coupon)
		c.JSON(400, response)
		return
	}
	coupon.Validity = time.Now().AddDate(0, 1, 0).Unix()
	couponresp, err := ph.productUsecase.AddCoupon(c, coupon)
	if err != nil {
		response := response.ErrorResponse(400, "failed to add coupon try again", err.Error(), couponresp)
		c.JSON(400, response)
		return
	}
	response := response.SuccessResponse(200, "Successfully add coupon", couponresp)
	c.JSON(200, response)

}

func (ph *ProductHandler) ListCoupon(c *gin.Context) {
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
	couponResp, metadata, err := ph.productUsecase.ListAllCoupons(c, pagination)
	if err != nil {
		response := response.ErrorResponse(400, "Error while geting coupon details", err.Error(), "")
		c.JSON(400, response)
		return
	}
	response := response.SuccessResponse(200, "successfully displayed all coupons", couponResp, metadata)
	c.JSON(200, response)
}
