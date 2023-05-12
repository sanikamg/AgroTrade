package interfaces

import (
	"context"
	"golang_project_ecommerce/pkg/domain"
	"golang_project_ecommerce/pkg/utils/res"
)

type UserUsecase interface {
	Register(ctx context.Context, user domain.Users) (domain.Users, error)
	Login(ctx context.Context, user domain.Users) (domain.Users, error)
	FindAllProducts(c context.Context, categoryname string) ([]res.AllProducts, error)
}
