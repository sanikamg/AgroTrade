package usecase

import (
	"context"
	"errors"
	"fmt"
	"golang_project_ecommerce/pkg/domain"
	interfaces "golang_project_ecommerce/pkg/repository/interface"
	ser "golang_project_ecommerce/pkg/usecase/interface"
	"golang_project_ecommerce/pkg/utils/res"
)

type ProductUsecase struct {
	productRepo interfaces.ProductRepository
}

func NewProductUsecase(repo interfaces.ProductRepository) ser.ProductUsecase {
	return &ProductUsecase{
		productRepo: repo,
	}
}

func (pu *ProductUsecase) AddCategory(c context.Context, category domain.Category) (domain.Category, error) {

	_, err := pu.productRepo.FindCategory(c, category)

	if err == nil {
		return domain.Category{}, errors.New("category already exists")
	}
	pu.productRepo.AddCategory(c, category)

	return category, nil
}

func (pu *ProductUsecase) DisplayAllCategory(c context.Context) ([]res.AllCategories, error) {

	categories, err := pu.productRepo.FindAllCategory(c)
	if err != nil {
		return []res.AllCategories{}, errors.New("error while finding all categories")
	}
	return categories, nil
}

func (pu *ProductUsecase) AddProduct(c context.Context, product domain.ProductDetails) (domain.ProductDetails, error) {
	fmt.Println(product)
	produ, err := pu.productRepo.FindProduct(c, product)
	product.Product_Id = produ.Product_Id
	if err == nil {
		prod, err := pu.productRepo.AddQuantity(c, product)
		if err != nil {
			return domain.ProductDetails{}, err
		}

		return prod, nil
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
	err = pu.productRepo.DeleteProduct(c, productid)
	if err != nil {
		return err
	}
	return nil
}
