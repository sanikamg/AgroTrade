package repository

import (
	"context"
	"errors"
	"golang_project_ecommerce/pkg/domain"
	interfaces "golang_project_ecommerce/pkg/repository/interface"
	"golang_project_ecommerce/pkg/utils/res"

	"gorm.io/gorm"
)

type adminDatabase struct {
	DB *gorm.DB
}

// constructor implements admin interface return admin database struct

func NewadminRepository(DB *gorm.DB) interfaces.AdminRepository {
	return &adminDatabase{DB}
}

func (ad *adminDatabase) FindAdmin(c context.Context, admin domain.AdminDetails) (domain.AdminDetails, error) {
	err := ad.DB.Where("name = ? OR phone=? Or email=?", admin.Name, admin.Phone, admin.Email).First(&admin).Error
	if err != nil {
		return domain.AdminDetails{}, errors.New("error while finding admin based on details")
	}
	return admin, nil
}
func (ad *adminDatabase) AddAdmin(c context.Context, admin domain.AdminDetails) (domain.AdminDetails, error) {
	err := ad.DB.Create(&admin).Error
	if err != nil {
		return domain.AdminDetails{}, errors.New("error while adding admin details to database")
	}
	return admin, nil
}
func (ad *adminDatabase) FindAll(c context.Context) ([]res.AllUsers, error) {
	var users []res.AllUsers
	err := ad.DB.Raw("select username,name,phone,email from users").Scan(&users).Error
	if err != nil {
		return []res.AllUsers{}, errors.New("error while finding all users")
	}
	return users, nil

}
