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

	//coupon
	AddCoupon(c context.Context, coupon domain.Coupon) (domain.Coupon, error)

	//payment
	AddPaymentMethod(c context.Context, payment domain.PaymentMethod) (domain.PaymentMethod, error)
	GetPaymentMethods(c context.Context, pagination utils.Pagination) ([]res.PaymentMethodResponse, utils.Metadata, error)
	UpdatePaymentMethod(c context.Context, payment domain.PaymentMethod) (domain.PaymentMethod, error)
	DeleteMethod(c context.Context, id uint) error

	//order
	GetTotalAmount(c context.Context, userid uint) (float64, error)
	CreateOrder(c context.Context, order domain.Order) (res.OrderResponse, error)
	UpdateOrderDetails(c context.Context, uporder req.UpdateOrder) (res.OrderResponse, error)
	ListAllOrders(c context.Context, pagination utils.Pagination, usrid uint) ([]res.OrderResponse, utils.Metadata, error)
	GetAllOrders(c context.Context, pagination utils.Pagination) ([]res.OrderResponse, utils.Metadata, error)
	DeleteOrder(c context.Context, order_id uint) error

	ValidateCoupon(c context.Context, CouponId uint) (res.CouponResponse, error)
	ApplyDiscount(c context.Context, CouponResponse res.CouponResponse, order_id uint) (int, error)

	//checkout
	PlaceOrder(c context.Context, order domain.Order) (res.PaymentResponse, error)
	FindPaymentMethodIdByOrderId(c context.Context, order_id uint) (uint, error)
	UpdateOrderStatus(c context.Context, order_id uint) (res.OrderResponse, error)
	FindTotalAmountByOrderId(c context.Context, order_id uint) (float64, error)
	FindPhnEmailByUsrId(c context.Context, usr_id int) (res.PhnEmailResp, error)
	GetRazorpayOrder(c context.Context, userID uint, razorPay req.RazorPayReq) (res.ResRazorpayOrder, error)
}
