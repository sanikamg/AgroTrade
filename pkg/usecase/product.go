package usecase

import (
	"context"
	"errors"
	"golang_project_ecommerce/pkg/config"
	"golang_project_ecommerce/pkg/domain"
	interfaces "golang_project_ecommerce/pkg/repository/interface"
	ser "golang_project_ecommerce/pkg/usecase/interface"
	"golang_project_ecommerce/pkg/utils"
	"golang_project_ecommerce/pkg/utils/req"
	"golang_project_ecommerce/pkg/utils/res"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"

	"github.com/google/uuid"
)

type ProductUsecase struct {
	productRepo interfaces.ProductRepository
}

func NewProductUsecase(repo interfaces.ProductRepository) ser.ProductUsecase {
	return &ProductUsecase{
		productRepo: repo,
	}
}

// category
func (pu *ProductUsecase) AddCategory(c context.Context, category domain.Category) (domain.Category, error) {

	_, err := pu.productRepo.FindCategory(c, category)

	if err == nil {
		return domain.Category{}, errors.New("category already exists")
	}
	pu.productRepo.AddCategory(c, category)

	return category, nil
}

func (pu *ProductUsecase) DisplayAllCategory(c context.Context, pagination utils.Pagination) ([]domain.Category, utils.Metadata, error) {

	categories, metadata, err := pu.productRepo.FindAllCategory(c, pagination)
	if err != nil {
		return []domain.Category{}, utils.Metadata{}, errors.New("error while finding all categories")
	}
	return categories, metadata, nil
}

func (pu *ProductUsecase) DeleteCategory(c context.Context, categoryName string) error {
	err := pu.productRepo.FindCategoryByName(c, categoryName)
	if err != nil {
		return errors.New("category doesn't exist")
	}
	err1 := pu.productRepo.DeleteCategory(c, categoryName)
	if err != nil {
		return err1
	}
	return nil
}

// product
func (pu *ProductUsecase) AddProduct(c context.Context, product domain.ProductDetails) (domain.ProductDetails, error) {

	produ, err := pu.productRepo.FindProduct(c, product)
	product.Product_Id = produ.Product_Id
	if err == nil {
		// prod, err := pu.productRepo.AddQuantity(c, product)
		// if err != nil {
		// 	return domain.ProductDetails{}, err
		// }

		// return prod, nil
		return produ, errors.New("product already exist please update product")
	}

	pro, err := pu.productRepo.AddProduct(c, product)
	if err != nil {
		return domain.ProductDetails{}, err
	}
	return pro, nil
}

func (pu *ProductUsecase) GetCategoryByID(c context.Context, categoryId int) (domain.Category, error) {
	category, err := pu.productRepo.GetCategoryByID(c, categoryId)
	if err != nil {
		return domain.Category{}, err
	}
	return category, nil
}

// product management from Admin side
func (pu *ProductUsecase) DeleteProduct(c context.Context, productid uint) error {
	err := pu.productRepo.FindProductById(c, productid)
	if err != nil {
		return errors.New("product doesn't exist")
	}

	err1 := pu.productRepo.DeleteProduct(c, productid)
	if err1 != nil {
		return err1
	}
	return nil
}

func (pu ProductUsecase) UpdateProduct(c context.Context, productup req.UpdateProduct) (domain.ProductDetails, error) {
	err := pu.productRepo.FindProductById(c, productup.ProductId)
	if err != nil {
		return domain.ProductDetails{}, errors.New("please add product,product doesn't exist")
	}

	product, err := pu.productRepo.UpdateProduct(c, productup)
	if err != nil {
		return domain.ProductDetails{}, err
	}
	return product, nil
}

func (pu *ProductUsecase) AddImage(c context.Context, pid int, files []*multipart.FileHeader) ([]domain.Image, error) {
	var images []domain.Image
	for _, file := range files {
		// Generate a unique filename for the image

		ext := filepath.Ext(file.Filename)
		filename := uuid.New().String() + ext

		image, err := pu.productRepo.AddImage(c, pid, filename)
		if err != nil {
			return []domain.Image{}, err
		}

		src, err := file.Open()
		if err != nil {
			return []domain.Image{}, err
		}
		defer src.Close()

		// Create the destination file
		dst, err := os.Create(filepath.Join("/home/user/Documents/Project/AgroTrade/images", filename)) // Replace "path/to/save/images" with your desired directory
		if err != nil {
			return []domain.Image{}, err
		}
		defer dst.Close()

		// Copy the uploaded file's content to the destination file
		_, err = io.Copy(dst, src)
		if err != nil {
			return []domain.Image{}, err
		}
		// product, _ := pu.productRepo.GetProductByID(c, pid)
		// image.Product = product

		images = append(images, image)
	}

	return images, nil
}

func (pu *ProductUsecase) GetProductByID(c context.Context, ProductId int) (domain.ProductDetails, error) {
	product, err := pu.productRepo.GetProductByID(c, ProductId)

	if err != nil {
		return domain.ProductDetails{}, err
	}
	return product, nil
}

func (pu *ProductUsecase) FindAllProducts(c context.Context, pagination utils.Pagination) ([]domain.ProductResponse, utils.Metadata, error) {
	product, metadata, err := pu.productRepo.FindAllProducts(c, pagination)
	if err != nil {
		return []domain.ProductResponse{}, utils.Metadata{}, err
	}
	return product, metadata, nil
}
func (pu *ProductUsecase) FindAllProductsByCategory(c context.Context, pagination utils.Pagination, category string) ([]domain.ProductResponse, utils.Metadata, error) {
	product, metadata, err := pu.productRepo.FindAllProductsByCategory(c, pagination, category)
	if err != nil {
		return []domain.ProductResponse{}, utils.Metadata{}, err
	}
	return product, metadata, nil
}

func (pu *ProductUsecase) FindProductsById(c context.Context, pagination utils.Pagination, id int) ([]domain.ProductResponse, utils.Metadata, error) {
	product, metadata, err := pu.productRepo.FindProductsById(c, pagination, id)
	if err != nil {
		return []domain.ProductResponse{}, utils.Metadata{}, err
	}
	return product, metadata, nil
}

// >>>>>>>>>>>>>>>>>>>>>>>>>>>>cart>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
func (pu *ProductUsecase) GetCartByUserID(c context.Context, userId uint) ([]res.CartResponse, error) {
	cart, err := pu.productRepo.GetCartByUserID(c, userId)
	if err != nil {
		return []res.CartResponse{}, err
	}
	return cart, nil
}

func (pu *ProductUsecase) AddToCart(c context.Context, cart domain.Cart_item) (domain.Cart_item, error) {
	cartitem, err := pu.productRepo.AddToCart(c, cart)
	if err != nil {
		return domain.Cart_item{}, nil
	}

	return cartitem, nil
}

func (pu *ProductUsecase) UpdateCart(c context.Context, cart domain.Cart_item) ([]res.CartResponse, error) {
	cartres, err := pu.productRepo.UpdateCart(c, cart)
	if err != nil {
		return []res.CartResponse{}, err
	}
	return cartres, err
}

func (pu *ProductUsecase) ListCartItems(c context.Context, pagination utils.Pagination, userid int) ([]res.CartResponse, utils.Metadata, error) {
	cartitems, metadata, err := pu.productRepo.ListCartItems(c, pagination, userid)
	if err != nil {
		return []res.CartResponse{}, utils.Metadata{}, err
	}
	return cartitems, metadata, nil
}

func (pu *ProductUsecase) RemoveProductFromCart(c context.Context, productid uint) error {
	err := pu.productRepo.RemoveProductFromCart(c, productid)
	if err != nil {
		return err
	}
	return nil
}

//>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>coupon>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>.

func (pu *ProductUsecase) AddCoupon(c context.Context, coupon domain.Coupon) (domain.Coupon, error) {
	err := pu.productRepo.FindCoupon(c, coupon)
	if err == nil {
		return domain.Coupon{}, err
	}
	couponresp, err1 := pu.productRepo.AddCoupon(c, coupon)
	if err1 != nil {
		return domain.Coupon{}, err1
	}
	return couponresp, nil
}

// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>payment method>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>

// add
func (pu *ProductUsecase) AddPaymentMethod(c context.Context, payment domain.PaymentMethod) (domain.PaymentMethod, error) {
	err := pu.productRepo.FindPaymentMethod(c, payment)
	if err == nil {
		return domain.PaymentMethod{}, errors.New("payment method already exists")
	}
	paymentresp, err1 := pu.productRepo.AddPaymentMethod(c, payment)
	if err1 != nil {
		return domain.PaymentMethod{}, errors.New("failed to add payment method")
	}
	return paymentresp, nil
}

// list
func (pu *ProductUsecase) GetPaymentMethods(c context.Context, pagination utils.Pagination) ([]res.PaymentMethodResponse, utils.Metadata, error) {
	paymentresp, metadata, err := pu.productRepo.GetPaymentMethods(c, pagination)
	if err != nil {
		return []res.PaymentMethodResponse{}, utils.Metadata{}, err
	}
	return paymentresp, metadata, nil
}

// update
func (pu *ProductUsecase) UpdatePaymentMethod(c context.Context, payment domain.PaymentMethod) (domain.PaymentMethod, error) {
	//Checking whether the payment id exist
	_, err := pu.productRepo.FindPaymentMethodId(c, payment.Method_id)

	if err != nil {
		return domain.PaymentMethod{}, errors.New("payment method doesn't exists")
	}

	paymentresp, err := pu.productRepo.UpdatePaymentMethod(c, payment)
	if err != nil {
		return domain.PaymentMethod{}, err
	}
	return paymentresp, nil
}

func (pu *ProductUsecase) DeleteMethod(c context.Context, id uint) error {

	err1 := pu.productRepo.DeleteMethod(c, id)
	if err1 != nil {
		return err1
	}
	return nil
}

//>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>Order>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>

func (pu *ProductUsecase) GetTotalAmount(c context.Context, userid uint) (float64, error) {
	var total_amount float64
	total_amount = 0
	cart, err := pu.productRepo.GetTotalAmount(c, int(userid))
	if err != nil {
		return 0, err
	}

	for _, c := range cart {
		total_amount = total_amount + float64(c.Total_Price)
	}
	return total_amount, nil
}

func (pu *ProductUsecase) CreateOrder(c context.Context, order domain.Order) (res.OrderResponse, error) {
	//Checking whether the payment id exist
	_, err := pu.productRepo.FindPaymentMethodId(c, order.PaymentMethodID)

	if err != nil {
		return res.OrderResponse{}, errors.New("payment method doesn't exists")
	}
	orderresp, err := pu.productRepo.CreateOrder(c, order)
	if err != nil {
		return res.OrderResponse{}, err
	}
	return orderresp, nil
}

func (pu *ProductUsecase) UpdateOrderDetails(c context.Context, uporder req.UpdateOrder) (res.OrderResponse, error) {
	//Checking whether the payment id exist
	_, err := pu.productRepo.FindPaymentMethodId(c, uporder.PaymentMethodID)

	if err != nil {
		return res.OrderResponse{}, errors.New("payment method doesn't exists")
	}
	orderup, err := pu.productRepo.UpdateOrderDetails(c, uporder)
	if err != nil {
		return res.OrderResponse{}, err
	}
	return orderup, nil
}

func (pu *ProductUsecase) ListAllOrders(c context.Context, pagination utils.Pagination, usrid uint) ([]res.OrderResponse, utils.Metadata, error) {
	paymentResp, metadata, err := pu.productRepo.ListAllOrders(c, pagination, usrid)
	if err != nil {
		return []res.OrderResponse{}, utils.Metadata{}, errors.New("failed to list products")
	}
	return paymentResp, metadata, err
}

func (pu *ProductUsecase) GetAllOrders(c context.Context, pagination utils.Pagination) ([]res.OrderResponse, utils.Metadata, error) {
	paymentResp, metadata, err := pu.productRepo.GetAllOrders(c, pagination)
	if err != nil {
		return []res.OrderResponse{}, utils.Metadata{}, errors.New("failed to list products")
	}
	return paymentResp, metadata, err
}

func (pu *ProductUsecase) DeleteOrder(c context.Context, order_id uint) error {
	err := pu.productRepo.DeleteOrder(c, order_id)
	if err != nil {
		return err
	}
	return nil
}

// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>..checkout..>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>

func (pu *ProductUsecase) PlaceOrder(c context.Context, order domain.Order) (res.PaymentResponse, error) {
	paymentresp, err := pu.productRepo.PlaceOrder(c, order)
	if err != nil {
		return res.PaymentResponse{}, err
	}
	return paymentresp, nil
}

func (pu *ProductUsecase) ValidateCoupon(c context.Context, CouponId uint) (res.CouponResponse, error) {
	err := pu.productRepo.FindCouponById(c, CouponId)
	if err != nil {
		return res.CouponResponse{}, err
	}
	couponResp, err := pu.productRepo.ValidateCoupon(c, CouponId)
	if err != nil {
		return res.CouponResponse{}, err
	}
	return couponResp, nil
}

func (pu *ProductUsecase) ApplyDiscount(c context.Context, CouponResponse res.CouponResponse, order_id uint) (int, error) {
	order, err := pu.productRepo.ApplyDiscount(c, order_id)
	if err != nil {
		return 0, nil
	}
	if CouponResponse.Quantity > 2 {
		totalamnt := order.Total_Amount - float64(CouponResponse.Discount)
		return int(totalamnt), nil
	}
	return 0, errors.New("Quantity should be more than 5")
}

func (pu *ProductUsecase) FindPaymentMethodIdByOrderId(c context.Context, order_id uint) (uint, error) {
	method_id, err := pu.productRepo.FindPaymentMethodIdByOrderId(c, order_id)
	if err != nil {
		return 0, err
	}
	return method_id, nil
}

func (pu *ProductUsecase) UpdateOrderStatus(c context.Context, order_id uint) (res.OrderResponse, error) {
	order_status := "order confirmed"
	orderResp, err := pu.productRepo.UpdateOrderStatus(c, order_id, order_status)
	if err != nil {
		return res.OrderResponse{}, err
	}
	return orderResp, nil
}

func (pu *ProductUsecase) FindTotalAmountByOrderId(c context.Context, order_id uint) (float64, error) {
	totalAmount, err := pu.productRepo.FindTotalAmountByOrderId(c, order_id)
	if err != nil {
		return 0, err
	}
	return totalAmount, nil
}

func (pu *ProductUsecase) FindPhnEmailByUsrId(c context.Context, usr_id int) (res.PhnEmailResp, error) {
	phnEmail, err := pu.productRepo.FindPhnEmailByUsrId(c, usr_id)
	if err != nil {
		return res.PhnEmailResp{}, err
	}
	return phnEmail, nil
}

// generate razorpay order
func (pu *ProductUsecase) GetRazorpayOrder(c context.Context, userID uint, razorPay req.RazorPayReq) (res.ResRazorpayOrder, error) {
	var razorpayOrder res.ResRazorpayOrder

	//razorpay amount is caluculate on pisa for india so make the actual price into paisa
	razorPayAmount := uint(razorPay.Total_Amount * 100)
	razopayOrderId, err := utils.GenerateRazorpayOrder(razorPayAmount, "test reciept")
	if err != nil {
		return razorpayOrder, err
	}

	// set all details on razopay order
	razorpayOrder.AmountToPay = uint(razorPay.Total_Amount)

	razorpayOrder.RazorpayKey, _ = config.GetRazorPayConfig()

	razorpayOrder.UserID = userID
	razorpayOrder.RazorpayOrderID = razopayOrderId

	razorpayOrder.Email = razorPay.Email
	razorpayOrder.Phone = razorPay.Phone

	return razorpayOrder, nil
}
