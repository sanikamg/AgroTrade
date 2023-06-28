package interfaces

import (
	"context"
	"golang_project_ecommerce/pkg/domain"
	"golang_project_ecommerce/pkg/utils"
	"golang_project_ecommerce/pkg/utils/req"
	"golang_project_ecommerce/pkg/utils/res"
)

type ProductRepository interface {
	//category
	FindCategory(c context.Context, category domain.Category) (domain.Category, error)
	AddCategory(c context.Context, category domain.Category) (domain.Category, error)
	FindAllCategory(c context.Context, pagination utils.Pagination) ([]domain.Category, utils.Metadata, error)
	GetCategoryByID(c context.Context, categoryId int) (domain.Category, error)
	FindCategoryByName(c context.Context, categoryName string) error
	//category management
	DeleteCategory(c context.Context, categoryName string) error
	//product
	AddProduct(c context.Context, product domain.ProductDetails) (domain.ProductDetails, error)
	FindProductById(c context.Context, productid uint) error
	AddQuantity(c context.Context, product domain.ProductDetails) (domain.ProductDetails, error)
	FindProduct(c context.Context, product domain.ProductDetails) (domain.ProductDetails, error)
	AddImage(c context.Context, pid int, filename string) (domain.Image, error)
	GetProductByID(c context.Context, ProductId int) (domain.ProductDetails, error)
	FindProductByName(c context.Context, productname string) error
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

	//coupon by admin
	AddCoupon(c context.Context, coupon domain.Coupon) (domain.Coupon, error)
	FindCoupon(c context.Context, coupon domain.Coupon) error

	//payment method by admin
	AddPaymentMethod(c context.Context, payment domain.PaymentMethod) (domain.PaymentMethod, error)
	FindPaymentMethod(c context.Context, payment domain.PaymentMethod) error
	FindPaymentMethodId(c context.Context, method_id uint) (uint, error)
	GetPaymentMethods(c context.Context, pagination utils.Pagination) ([]res.PaymentMethodResponse, utils.Metadata, error)
	UpdatePaymentMethod(c context.Context, payment domain.PaymentMethod) (domain.PaymentMethod, error)
	DeleteMethod(c context.Context, id uint) error

	//order
	GetTotalAmount(c context.Context, userid int) ([]domain.Cart_item, error)
	CreateOrder(c context.Context, order domain.Order) (res.OrderResponse, error)
	UpdateOrderDetails(c context.Context, uporder req.UpdateOrder) (res.OrderResponse, error)
	ListAllOrders(c context.Context, pagination utils.Pagination, usrid uint) ([]res.OrderResponse, utils.Metadata, error)
	GetAllOrders(c context.Context, pagination utils.Pagination) ([]res.OrderResponse, utils.Metadata, error)
	DeleteOrder(c context.Context, order_id uint) error

	//checkout
	ValidateCoupon(c context.Context, CouponId uint) (res.CouponResponse, error)
	FindCouponById(c context.Context, couponId uint) error
	ApplyDiscount(c context.Context, order_id uint) (domain.Order, error)
	PlaceOrder(c context.Context, order domain.Order) (res.PaymentResponse, error)

	FindPaymentMethodIdByOrderId(c context.Context, order_id uint) (uint, error)
	FindTotalAmountByOrderId(c context.Context, order_id uint) (float64, error)
	FindPhnEmailByUsrId(c context.Context, usr_id int) (res.PhnEmailResp, error)

	UpdateOrderStatus(c context.Context, order_id uint, order_status string) (res.OrderResponse, error)
}
