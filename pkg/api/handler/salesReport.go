package handler

import (
	"fmt"
	"golang_project_ecommerce/pkg/common/response"
	"golang_project_ecommerce/pkg/utils"
	"golang_project_ecommerce/pkg/utils/req"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jung-kurt/gofpdf"
)

func (pd *ProductHandler) SalesReport(c *gin.Context) {

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
	sDate, err := utils.StringToTime(c.Query("startDate"))
	if err != nil {
		response := response.ErrorResponse(400, "Please add start date as params", err.Error(), "")
		c.JSON(400, response)
		return
	}
	eDate, err := utils.StringToTime(c.Query("endDate"))
	if err != nil {
		response := response.ErrorResponse(400, "Please add end date as params", err.Error(), "")
		c.JSON(400, response)
		return
	}
	salesData := req.ReqSalesReport{
		StartDate: sDate,
		EndDate:   eDate,
		Pagination: utils.Pagination{
			Page:     page,
			PageSize: pagesize,
		},
	}
	salesReport, _, _ := pd.productUsecase.SalesReport(c, salesData)
	if salesReport == nil {
		response := response.ErrorResponse(400, "There is no sales report on this period", " ", " ")
		c.JSON(400, response)
		return
	}

	// Create a new PDF document
	pdf := gofpdf.New("P", "mm", "A4", "")

	// Add a new page
	pdf.AddPage()

	// Set the font and font size
	pdf.SetFont("Arial", "i", 12)

	// Add the report title
	pdf.CellFormat(0, 15, "Sales Report", "", 0, "C", false, 0, "")
	pdf.Ln(10)
	// Add the sales report data to the PDF
	i := 1
	for _, sale := range salesReport {

		pdf.CellFormat(0, 15, fmt.Sprint(i)+".", "", 0, "L", false, 0, "")
		pdf.Ln(10)
		pdf.Cell(0, 10, fmt.Sprintf("User ID: %d", sale.UserID))
		pdf.Ln(10)
		pdf.Cell(0, 10, fmt.Sprintf("Name: %s", sale.Name))
		pdf.Ln(10)
		pdf.Cell(0, 10, fmt.Sprintf("Email: %s", sale.Email))
		pdf.Ln(10)
		pdf.Cell(0, 10, fmt.Sprintf("Order Date: %v", sale.OrderDate))
		pdf.Ln(10)
		pdf.Cell(0, 10, fmt.Sprintf("TotalPrice: %v", sale.OrderTotalPrice))
		pdf.Ln(10)
		pdf.Cell(0, 10, fmt.Sprintf("Order Status: %s", sale.OrderStatus))
		pdf.Ln(10)
		pdf.Cell(0, 10, fmt.Sprintf("Payment status: %v", sale.PaymentStatus))
		pdf.Ln(10)
		pdf.Cell(0, 10, fmt.Sprintf("Payment Type: %v", sale.PaymentType))
		pdf.Ln(10)

		// Move to the next line
		pdf.Ln(10)
		i++
	}

	// Generate a temporary file path for the PDF
	pdfFilePath := "salesReport/file.pdf"

	// Save the PDF to the temporary file path
	err = pdf.OutputFileAndClose(pdfFilePath)
	if err != nil {
		response := response.ErrorResponse(500, "Failed to generate PDF", err.Error(), "")
		c.JSON(500, response)
		return
	}

	// Set the appropriate headers for the file download
	c.Header("Content-Disposition", "attachment; filename=sales_report.pdf")
	c.Header("Content-Type", "application/pdf")

	// Serve the PDF file for download
	c.File(pdfFilePath)

	// response := response.SuccessResponse(200, "successfully generated pdf", " ")
	// c.JSON(200, response)
}
