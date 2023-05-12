package handler

import (
	"golang_project_ecommerce/pkg/api/middlware"
	"golang_project_ecommerce/pkg/common/response"
	"golang_project_ecommerce/pkg/domain"
	services "golang_project_ecommerce/pkg/usecase/interface"
	"golang_project_ecommerce/pkg/utils/req"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	productUsecase services.ProductUsecase
}

func NewProductHandler(usecase services.ProductUsecase) *ProductHandler {
	return &ProductHandler{
		productUsecase: usecase,
	}
}

//category

func (ph ProductHandler) SaveCategory(c *gin.Context) {
	var category domain.Category
	if err := c.ShouldBindJSON(&category); err != nil {
		res := response.ErrorResponse(400, "error while getting  the data from admin side", err.Error(), category)

		c.JSON(http.StatusBadRequest, res)
		return
	}
	category, err := ph.productUsecase.AddCategory(c, category)
	if err != nil {
		res := response.ErrorResponse(400, "error while getting  the data from the user side", err.Error(), category)

		c.JSON(http.StatusBadRequest, res)
		return
	}
	message := category.CategoryName + "category added successfully"
	response := response.SuccessResponse(200, "Category added", message)
	c.JSON(http.StatusOK, response)
}

func (ph ProductHandler) GetAllCategory(c *gin.Context) {
	categories, err := ph.productUsecase.DisplayAllCategory(c)
	if err != nil {
		res := response.ErrorResponse(400, "can't find all categories", err.Error(), nil)

		c.JSON(http.StatusBadRequest, res)
		return
	}
	response := response.SuccessResponse(200, "successfully got all cateogries", categories)
	c.JSON(http.StatusOK, response)
}

//product

func (ph *ProductHandler) SaveProduct(c *gin.Context) {
	var product domain.ProductDetails

	//get id from getid
	id, err := middlware.GetId(c, "Admin_Authorization")
	if err != nil {
		response := response.ErrorResponse(400, "error while getting id from cookie", err.Error(), product)
		c.JSON(400, response)
		return
	}

	product.AdminId = uint(id)

	if err := c.ShouldBindJSON(&product); err != nil {
		response := response.ErrorResponse(400, "error entering details", err.Error(), product)
		c.JSON(400, response)
		return
	}

	// Get the category ID from the request body
	categoryID := product.CategoryID

	// Check if the category exists in the database
	category, err := ph.productUsecase.GetCategoryByID(c, int(categoryID))
	if err != nil {
		response := response.ErrorResponse(400, "can't find category", err.Error(), product)
		c.JSON(400, response)
		return
	}

	// Set the Category field to the category object
	product.Category = category

	productDetails, err := ph.productUsecase.AddProduct(c, product)
	if err != nil {
		response := response.ErrorResponse(400, "can't add product", err.Error(), product)
		c.JSON(400, response)
		return
	}
	response := response.SuccessResponse(200, "successfully added product", productDetails)
	c.JSON(200, response)
}

//product management by admin side

func (ph *ProductHandler) RemoveProduct(c *gin.Context) {
	var product req.DeleteId
	if err := c.ShouldBindJSON(&product); err != nil {
		response := response.ErrorResponse(400, "enter a valid product id", err.Error(), product)
		c.JSON(400, response)
		return
	}

	err := ph.productUsecase.DeleteProduct(c, product.ProductID)
	if err != nil {
		response := response.ErrorResponse(400, "can't be delete", err.Error(), product)
		c.JSON(400, response)
		return
	}
	response := response.SuccessResponse(200, "successfully deleted product", product)
	c.JSON(200, response)
}
