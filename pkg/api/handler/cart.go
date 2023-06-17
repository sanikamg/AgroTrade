package handler

import (
	"golang_project_ecommerce/pkg/api/middlware"
	"golang_project_ecommerce/pkg/common/response"
	"golang_project_ecommerce/pkg/domain"
	"golang_project_ecommerce/pkg/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>cart>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
func (ph *ProductHandler) AddProductToCart(c *gin.Context) {
	var cart domain.Cart_item
	usrid, err := middlware.GetId(c, "User_Authorization")
	if err != nil {
		response := response.ErrorResponse(400, "user authentication failed", err.Error(), "")
		c.JSON(400, response)
		return
	}
	cart.User_Id = uint(usrid)
	productId, err := strconv.Atoi(c.Query("product_id"))
	if err != nil {
		response := response.ErrorResponse(400, "Please give product id as params", err.Error(), "")
		c.JSON(400, response)
		return
	}
	cart.Product_Id = uint(productId)
	qty, err := strconv.Atoi(c.Query("quantity"))
	if err != nil {
		response := response.ErrorResponse(400, "Please give quantity as params", err.Error(), "")
		c.JSON(400, response)
		return
	}
	cart.Quantity = uint(qty)
	products, err := ph.productUsecase.GetProductByID(c, int(cart.Product_Id))
	if err != nil {
		response := response.ErrorResponse(400, "Product doesn't exist", err.Error(), "")
		c.JSON(400, response)
		return
	}
	cart.Total_Price = float32(products.ProductPrice) * float32(cart.Quantity)
	count := 0
	userCart, err := ph.productUsecase.GetCartByUserID(c, cart.User_Id)
	if err != nil {
		response := response.ErrorResponse(400, "failed to get cart", err.Error(), " ")
		c.JSON(400, response)
		return
	}
	for _, ct := range userCart {
		if cart.Product_Id == ct.Product_Id {
			count++
			cart.Quantity = ct.Quantity + cart.Quantity
			cart.Total_Price = float32(cart.Quantity) * float32(products.ProductPrice)
			cartItem, err := ph.productUsecase.UpdateCart(c, cart)
			if err != nil {
				response := response.ErrorResponse(400, "failed to add product into cart try again", err.Error(), " ")
				c.JSON(400, response)
				return
			}
			response := response.SuccessResponse(200, "Successfully updated  your cart", cartItem)
			c.JSON(200, response)
			return
		}
	}
	if count == 0 {
		cartitem, err := ph.productUsecase.AddToCart(c, cart)
		if err != nil {
			response := response.ErrorResponse(400, "failed to add product into cart try again", err.Error(), " ")
			c.JSON(400, response)
			return
		}
		response := response.SuccessResponse(200, "Succefully added product to your cart", cartitem)
		c.JSON(200, response)
		return
	}

}

func (ph *ProductHandler) ViewCart(c *gin.Context) {
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
	user_id := int(usrid)

	cartitems, metadata, err := ph.productUsecase.ListCartItems(c, pagination, user_id)
	if err != nil {
		response := response.ErrorResponse(400, "error while finding products", err.Error(), cartitems)
		c.JSON(400, response)
	}
	response := response.SuccessResponse(200, "successfully displayed all products", cartitems, metadata)
	c.JSON(200, response)

}

func (ph *ProductHandler) RemoveProductFromCart(c *gin.Context) {
	product_id, err := strconv.Atoi(c.Query("productid"))
	if err != nil {
		response := response.ErrorResponse(400, "Please add productid as params", err.Error(), "")
		c.JSON(400, response)
	}
	productId := uint(product_id)
	err1 := ph.productUsecase.RemoveProductFromCart(c, productId)
	if err1 != nil {
		response := response.ErrorResponse(400, "can't delete product", err.Error(), "")
		c.JSON(400, response)
	}
	response := response.SuccessResponse(200, "successfully delete product", " ")
	c.JSON(200, response)

}
