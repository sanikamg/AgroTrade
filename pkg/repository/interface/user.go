package interfaces

import (
	"context"
	"golang_project_ecommerce/pkg/domain"
	"golang_project_ecommerce/pkg/utils/res"
)

type UserRepository interface {
	Addusers(ctx context.Context, user domain.Users) (domain.Users, error)
	FindUser(ctx context.Context, user domain.Users) (domain.Users, error)

	//list all products
	FindAllProducts(c context.Context, categoryid uint) ([]res.AllProducts, error)

	//get cotegory
	GetCategoryByName(c context.Context, categoryname string) (domain.Category, error)
}
