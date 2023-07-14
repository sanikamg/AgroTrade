package handler

import (
	"golang_project_ecommerce/pkg/common/response"
	"golang_project_ecommerce/pkg/utils"
	"golang_project_ecommerce/pkg/utils/req"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (pd *ProductHandler) Statistics(c *gin.Context) {
	var sales req.Sales
	sales.Sdate = c.Query("startDate")
	sales.Edate = c.Query("endDate")

	sDate, err := utils.StringToTime(sales.Sdate)
	if err != nil {
		response := response.ErrorResponse(400, "Please add start date as params", err.Error(), "")
		c.JSON(400, response)
		return
	}
	eDate, err := utils.StringToTime(sales.Edate)
	if err != nil {
		response := response.ErrorResponse(400, "Please add end date as params", err.Error(), "")
		c.JSON(400, response)
		return
	}

	salesData, err := pd.productUsecase.SalesData(sDate, eDate)
	if err != nil {
		response := response.ErrorResponse(400, "Can't calulate details of sales ", err.Error(), salesData)
		c.JSON(400, response)
		return
	}

	response := response.SuccessResponse(200, "successfully displayed sales data", salesData)
	c.JSON(200, response)

}

func (pd ProductHandler) FindPendingDelivery(c *gin.Context) {
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
	pendings, metadata, err := pd.productUsecase.FindPendingDelivery(c, pagination)
	if err != nil {
		response := response.ErrorResponse(400, "FAiled to find pending deliveries", err.Error(), "")
		c.JSON(400, response)
		return
	}
	response := response.SuccessResponse(200, "Successfully displayed pending deliveries", pendings, metadata)
	c.JSON(200, response)
}
