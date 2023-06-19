package handler

import (
	"golang_project_ecommerce/pkg/api/middlware"
	"golang_project_ecommerce/pkg/common/response"
	"golang_project_ecommerce/pkg/domain"
	services "golang_project_ecommerce/pkg/usecase/interface"
	"golang_project_ecommerce/pkg/utils"
	"golang_project_ecommerce/pkg/utils/req"
	"net/http"
	"strconv"

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
	page, _ := strconv.Atoi(c.Query("page"))
	pagesize, _ := strconv.Atoi(c.Query("pagesize"))
	pagination := utils.Pagination{
		Page:     page,
		PageSize: pagesize,
	}
	categories, metadata, err := ph.productUsecase.DisplayAllCategory(c, pagination)
	if err != nil {
		res := response.ErrorResponse(400, "can't find all categories", err.Error(), nil)

		c.JSON(http.StatusBadRequest, res)
		return
	}
	response := response.SuccessResponse(200, "successfully got all cateogries", categories, metadata)
	c.JSON(http.StatusOK, response)
}

//>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>product>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>

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

func (ph *ProductHandler) EditProduct(c *gin.Context) {
	var productup req.UpdateProduct

	if err := c.ShouldBindJSON(&productup); err != nil {
		response := response.ErrorResponse(400, "enter valid details", err.Error(), productup)
		c.JSON(400, response)
		return
	}

	product, err := ph.productUsecase.UpdateProduct(c, productup)
	if err != nil {
		response := response.ErrorResponse(400, "can't update product", err.Error(), productup)
		c.JSON(400, response)
		return
	}
	response := response.SuccessResponse(200, "successfully updated  product", product)
	c.JSON(200, response)
}

func (ph *ProductHandler) DeleteCategory(c *gin.Context) {
	var deletecat req.DeleteCategory

	if err := c.ShouldBindJSON(&deletecat); err != nil {
		response := response.ErrorResponse(400, "Enter valid category name", err.Error(), deletecat)
		c.JSON(400, response)
		return
	}

	err := ph.productUsecase.DeleteCategory(c, deletecat.CategoryName)
	if err != nil {
		response := response.ErrorResponse(400, "can't delete category", err.Error(), deletecat)
		c.JSON(400, response)
		return
	}
	response := response.SuccessResponse(200, "successfully delete category", deletecat)
	c.JSON(200, response)
}

//>>>>>>>>>>>>>>>>>>add image>>>>>>>>>>>>>>>>>>>>>>>>>>>

func (ph *ProductHandler) AddImage(c *gin.Context) {

	pid, err := strconv.Atoi(c.PostForm("product_id"))
	if err != nil {
		response := response.ErrorResponse(400, "Error while fetching product_id", err.Error(), pid)
		c.JSON(400, response)
		return
	}

	form, err := c.MultipartForm()

	if err != nil {
		response := response.ErrorResponse(400, "Error while fetching image file", err.Error(), form)
		c.JSON(400, response)
		return
	}
	files := form.File["image"]
	if len(files) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No image files found"})
		return
	}

	Images, err := ph.productUsecase.AddImage(c, pid, files)
	if err != nil {
		response := response.ErrorResponse(400, "Can't be add images", err.Error(), Images)
		c.JSON(400, response)
		return
	}

	response := response.SuccessResponse(200, "successfully added images", Images)
	c.JSON(200, response)
}

// >>>>>>>>>>>>>>>>>>>>>Get All Products >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
func (ph *ProductHandler) GetAllProducts(c *gin.Context) {
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
	product, metadata, err := ph.productUsecase.FindAllProducts(c, pagination)
	if err != nil {
		response := response.ErrorResponse(400, "error while finding products", err.Error(), product)
		c.JSON(400, response)
	}
	response := response.SuccessResponse(200, "successfully displayed all products", product, metadata)
	c.JSON(200, response)
}

func (ph *ProductHandler) GetAllProductsByCategory(c *gin.Context) {
	category := c.Query("category")
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
	product, metadata, err := ph.productUsecase.FindAllProductsByCategory(c, pagination, category)
	if err != nil {
		response := response.ErrorResponse(400, "error while finding products", err.Error(), product)
		c.JSON(400, response)
	}
	response := response.SuccessResponse(200, "successfully displayed all products", product, metadata)
	c.JSON(200, response)
}

func (ph *ProductHandler) GetProductById(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		response := response.ErrorResponse(400, "Please add id as params", err.Error(), "")
		c.JSON(400, response)
	}
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
	product, metadata, err := ph.productUsecase.FindProductsById(c, pagination, id)
	if err != nil {
		response := response.ErrorResponse(400, "error while finding products", err.Error(), product)
		c.JSON(400, response)
	}
	response := response.SuccessResponse(200, "successfully displayed all products", product, metadata)
	c.JSON(200, response)
}
