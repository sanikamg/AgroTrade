package interfaces

import (
	"context"
	"golang_project_ecommerce/pkg/domain"
	"golang_project_ecommerce/pkg/utils"
	"golang_project_ecommerce/pkg/utils/res"
)

type AdminUsecase interface {
	AdminSignup(c context.Context, admin domain.AdminDetails) (domain.AdminDetails, error)
	AdminLogin(ctx context.Context, admin domain.AdminDetails) error
	FindAllUsers(c context.Context, pagination utils.Pagination) ([]res.AllUsers, utils.Metadata, error)
	BlockUser(c context.Context, id int) error
	UnBlockUser(c context.Context, id int) error
	FindByUsername(c context.Context, Username string) (domain.AdminDetails, error)
}
