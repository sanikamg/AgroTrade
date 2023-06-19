package usecase

import (
	"context"
	"errors"
	"golang_project_ecommerce/pkg/domain"
	interfaces "golang_project_ecommerce/pkg/repository/interface"
	ser "golang_project_ecommerce/pkg/usecase/interface"
	"golang_project_ecommerce/pkg/utils"
	"golang_project_ecommerce/pkg/utils/req"
	"golang_project_ecommerce/pkg/utils/res"
	"mime/multipart"
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
	orderresp, err := pu.productRepo.CreateOrder(c, order)
	if err != nil {
		return res.OrderResponse{}, err
	}
	return orderresp, nil
}

func (pu *ProductUsecase) PlaceOrder(c context.Context, order domain.Order) (res.PaymentResponse, error) {
	paymentresp, err := pu.productRepo.PlaceOrder(c, order)
	if err != nil {
		return res.PaymentResponse{}, err
	}
	return paymentresp, nil
}
