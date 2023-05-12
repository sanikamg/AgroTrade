package interfaces

import (
	"context"
	"golang_project_ecommerce/pkg/domain"
	"golang_project_ecommerce/pkg/utils/res"
)

type ProductRepository interface {
	//category
	FindCategory(c context.Context, category domain.Category) (domain.Category, error)
	AddCategory(c context.Context, category domain.Category) (domain.Category, error)
	FindAllCategory(c context.Context) ([]res.AllCategories, error)
	GetCategoryByID(c context.Context, categoryId int) (domain.Category, error)
	//product
	AddProduct(c context.Context, product domain.ProductDetails) (domain.ProductDetails, error)
	FindProductById(c context.Context, productid uint) error
	AddQuantity(c context.Context, product domain.ProductDetails) (domain.ProductDetails, error)
	FindProduct(c context.Context, product domain.ProductDetails) (domain.ProductDetails, error)
	//product management
	DeleteProduct(c context.Context, productid uint) error
}
