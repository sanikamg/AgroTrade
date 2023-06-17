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

	"golang.org/x/crypto/bcrypt"
)

type AdminUsecase struct {
	adminRepo interfaces.AdminRepository
}

func NewadminUsecase(repo interfaces.AdminRepository) ser.AdminUsecase {
	return &AdminUsecase{
		adminRepo: repo,
	}
}

// AdminSignup implements interfaces.AdminUsecase
func (ad *AdminUsecase) AdminSignup(c context.Context, admin domain.AdminDetails) (domain.AdminDetails, error) {

	adn, err := ad.adminRepo.FindAdmin(c, admin)
	if err == nil {
		return adn, errors.New("user already exists")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(admin.Password), 14)
	if err != nil {
		return domain.AdminDetails{}, errors.New("error while hashing")
	}
	admin.Password = string(hash)

	ad.adminRepo.AddAdmin(c, admin)

	return admin, nil

}

func (ad *AdminUsecase) AdminLogin(ctx context.Context, admin domain.AdminDetails) error {
	dbAdmin, dbErr := ad.adminRepo.FindAdmin(ctx, admin)

	//check whether the user exists or valid information
	if dbErr == nil {
		return dbErr
	} else if dbAdmin.ID == 0 {
		return errors.New("user does not exists with this , please register")
	}

	// check password matching

	if bcrypt.CompareHashAndPassword([]byte(dbAdmin.Password), []byte(admin.Password)) != nil {
		return errors.New("password is not correct")
	}
	return nil
}
func (ad *AdminUsecase) FindAllUsers(c context.Context, pagination utils.Pagination) ([]res.AllUsers, utils.Metadata, error) {
	users, metadata, err := ad.adminRepo.FindAllUsers(c, pagination)
	if err != nil {
		return []res.AllUsers{}, utils.Metadata{}, errors.New("error while finding all users")
	}
	return users, metadata, nil
}

// to block an user

func (ad *AdminUsecase) BlockUser(c context.Context, id int) error {
	var status req.BlockStatus
	status.UserID = uint(id)
	status.BlockStatus = true
	err := ad.adminRepo.BlockUser(c, status)
	if err != nil {
		return err
	}
	return nil
}

// to unblock user

func (ad *AdminUsecase) UnBlockUser(c context.Context, id int) error {
	var status req.BlockStatus
	status.UserID = uint(id)
	status.BlockStatus = false
	err := ad.adminRepo.BlockUser(c, status)
	if err != nil {
		return err
	}
	return nil
}

// find admin details by username
func (ad *AdminUsecase) FindByUsername(c context.Context, Username string) (domain.AdminDetails, error) {
	admin, err := ad.adminRepo.FindByUsername(c, Username)
	if err != nil {
		return domain.AdminDetails{}, err
	}
	return admin, nil
}
