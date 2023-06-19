package interfaces

import (
	"context"
	"golang_project_ecommerce/pkg/domain"
	"golang_project_ecommerce/pkg/utils"
	"golang_project_ecommerce/pkg/utils/req"
	"golang_project_ecommerce/pkg/utils/res"
)

type AdminRepository interface {
	FindAdmin(c context.Context, admin domain.AdminDetails) (domain.AdminDetails, error)
	AddAdmin(c context.Context, admin domain.AdminDetails) (domain.AdminDetails, error)
	FindAllUsers(c context.Context, pagination utils.Pagination) ([]res.AllUsers, utils.Metadata, error)

	//user management
	BlockUser(c context.Context, status req.BlockStatus) error
	FindByUsername(c context.Context, Username string) (domain.AdminDetails, error)
}
