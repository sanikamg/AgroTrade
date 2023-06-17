package interfaces

import (
	"context"
	"golang_project_ecommerce/pkg/domain"
	"golang_project_ecommerce/pkg/utils"
	"golang_project_ecommerce/pkg/utils/req"
	"golang_project_ecommerce/pkg/utils/res"
	"mime/multipart"
)

type ProductUsecase interface {
	//category
	AddCategory(c context.Context, category domain.Category) (domain.Category, error)
	DisplayAllCategory(c context.Context, pagination utils.Pagination) ([]domain.Category, utils.Metadata, error)
	GetCategoryByID(c context.Context, categoryId int) (domain.Category, error)
	DeleteCategory(c context.Context, categoryName string) error
	//product
	AddProduct(c context.Context, product domain.ProductDetails) (domain.ProductDetails, error)
	AddImage(c context.Context, pid int, files []*multipart.FileHeader) ([]domain.Image, error)
	GetProductByID(c context.Context, ProductId int) (domain.ProductDetails, error)
	FindAllProducts(c context.Context, pagination utils.Pagination) ([]domain.ProductResponse, utils.Metadata, error)
	FindAllProductsByCategory(c context.Context, pagination utils.Pagination, category string) ([]domain.ProductResponse, utils.Metadata, error)
	FindProductsById(c context.Context, pagination utils.Pagination, id int) ([]domain.ProductResponse, utils.Metadata, error)
	//product management
	DeleteProduct(c context.Context, productid uint) error
	UpdateProduct(c context.Context, productup req.UpdateProduct) (domain.ProductDetails, error)

	//cart
	GetCartByUserID(c context.Context, userId uint) ([]res.CartResponse, error)
	AddToCart(c context.Context, cart domain.Cart_item) (domain.Cart_item, error)
	UpdateCart(c context.Context, cart domain.Cart_item) ([]res.CartResponse, error)
	ListCartItems(c context.Context, pagination utils.Pagination, userid int) ([]res.CartResponse, utils.Metadata, error)
	RemoveProductFromCart(c context.Context, productid uint) error

	//order
	GetTotalAmount(c context.Context, userid uint) (float64, error)
	CreateOrder(c context.Context, order domain.Order) (res.OrderResponse, error)
	PlaceOrder(c context.Context, order domain.Order) (res.PaymentResponse, error)
}
