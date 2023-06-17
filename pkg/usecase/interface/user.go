package interfaces

import (
	"context"
	"golang_project_ecommerce/pkg/domain"
	"golang_project_ecommerce/pkg/utils"
	"golang_project_ecommerce/pkg/utils/req"
	"golang_project_ecommerce/pkg/utils/res"
)

type UserUsecase interface {
	//user signup
	Register(ctx context.Context, user domain.Users) (domain.Users, error)
	//Adduser(ctx context.Context, user domain.Users) (domain.Users, error)
	UpdateStatus(c context.Context, user domain.Users) error
	SendOtpPhn(c context.Context, phn domain.Users) error
	VerifyOtp(c context.Context, phn string, otp string) error
	//user login
	Login(ctx context.Context, user domain.Users) (domain.Users, error)

	//user profile
	FindUserById(c context.Context, id int) (domain.Users, error)
	EditUserDetails(c context.Context, id int, user req.Usereditreq) (domain.Users, error)
	//list all products
	//FindAllProducts(c context.Context, categoryname string) ([]res.AllProducts, error)

	//address
	AddAddress(c context.Context, address domain.Address) (res.AddressResponse, error)
	EditAddress(c context.Context, address domain.Address) (res.AddressResponse, error)
	ListAddresses(c context.Context, pagination utils.Pagination, id uint) ([]res.AddressResponse, utils.Metadata, error)
	DeleteAddress(c context.Context, addressid uint) error

	//forgot password

	SendOtpForgotPass(c context.Context, phn string) error
	VerifyOtpForgotpass(c context.Context, phn string, otp string) error
	ForgotPassword(c context.Context, usrphn string, newpass string) error
}
