package interfaces

import (
	"context"
	"golang_project_ecommerce/pkg/domain"
	"golang_project_ecommerce/pkg/utils"
	"golang_project_ecommerce/pkg/utils/req"
	"golang_project_ecommerce/pkg/utils/res"
)

type UserRepository interface {
	Addusers(ctx context.Context, user domain.Users) (domain.Users, error)
	FindUser(ctx context.Context, user domain.Users) (domain.Users, error)
	IsEmtyUsername(c context.Context, username domain.Users) bool
	UpdateStatus(c context.Context, user domain.Users) error
	FindStatus(c context.Context, phn string) (domain.Users, error)
	FindUserByPhn(c context.Context, phn domain.Users) error
	UpdateUserDetails(c context.Context, user domain.Users) (domain.Users, error)

	FindUserById(c context.Context, id int) (domain.Users, error)
	EditUserDetails(c context.Context, id int, user req.Usereditreq) (domain.Users, error)
	//StoreVerifyDetails(c context.Context, phn string, code int) error

	//list all products
	FindAllProducts(c context.Context, categoryid uint) ([]res.AllProducts, error)

	//get cotegory
	GetCategoryByName(c context.Context, categoryname string) (domain.Category, error)

	//user Address
	AddAddress(c context.Context, address domain.Address) (res.AddressResponse, error)
	FindAddress(c context.Context, addressid uint) error
	EditAddress(c context.Context, address domain.Address) (res.AddressResponse, error)
	ListAddresses(c context.Context, pagination utils.Pagination, id uint) ([]res.AddressResponse, utils.Metadata, error)
	DeleteAddress(c context.Context, addressid uint) error

	//forgot password
	FindUserByPhnNum(c context.Context, phn string) error
	ForgotPassword(c context.Context, usrphn string, newpass string) error
}
