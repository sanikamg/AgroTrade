package interfaces

import (
	"context"
	"golang_project_ecommerce/pkg/domain"
	"golang_project_ecommerce/pkg/utils/res"
)

type ProductUsecase interface {
	//category
	AddCategory(c context.Context, category domain.Category) (domain.Category, error)
	DisplayAllCategory(c context.Context) ([]res.AllCategories, error)
	GetCategoryByID(c context.Context, categoryId int) (domain.Category, error)

	//product
	AddProduct(c context.Context, product domain.ProductDetails) (domain.ProductDetails, error)

	//product management
	DeleteProduct(c context.Context, productid uint) error
}
