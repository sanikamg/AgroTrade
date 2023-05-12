package repository

import (
	"context"
	"errors"
	"golang_project_ecommerce/pkg/domain"
	interfaces "golang_project_ecommerce/pkg/repository/interface"
	"golang_project_ecommerce/pkg/utils/req"
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

// for findin admin
func (ad *adminDatabase) FindAdmin(c context.Context, admin domain.AdminDetails) (domain.AdminDetails, error) {
	err := ad.DB.Where("username=? OR name = ? OR phone=? OR email=?", admin.Username, admin.Name, admin.Phone, admin.Email).First(&admin).Error
	if err != nil {
		return domain.AdminDetails{}, errors.New("admin not found")
	}
	return admin, nil
}

// for adding admin to database
func (ad *adminDatabase) AddAdmin(c context.Context, admin domain.AdminDetails) (domain.AdminDetails, error) {
	err := ad.DB.Create(&admin).Error
	if err != nil {
		return domain.AdminDetails{}, errors.New("error while adding admin details to database")
	}

	return admin, nil
}

// for finding allusers
func (ad *adminDatabase) FindAll(c context.Context) ([]res.AllUsers, error) {
	var users []res.AllUsers
	err := ad.DB.Raw("select user_id as id,username,name,phone,email from users").Scan(&users).Error
	if err != nil {
		return []res.AllUsers{}, errors.New("failed to find all users")
	}
	return users, nil

}

// block user by getting id
func (ad *adminDatabase) BlockUser(c context.Context, status req.BlockStatus) error {
	//find user in that id
	var user domain.Users

	ad.DB.Raw("select *from users where user_id=?", status.UserID).Scan(&user)
	if user.User_Id == 0 {
		return errors.New("user doesn't exist")
	}

	query := `update users set block_status=? where user_id=?`
	err := ad.DB.Raw(query, status.BlockStatus, status.UserID).Scan(&user).Error
	if err != nil {
		return errors.New("failed to update block status")
	}
	return nil
}

func (ad *adminDatabase) FindByUsername(c context.Context, Username string) (domain.AdminDetails, error) {
	var admin domain.AdminDetails

	err := ad.DB.Raw("select *from admin_details where username=?", Username).Scan(&admin).Error
	if err != nil {
		return domain.AdminDetails{}, errors.New("failed find user details")
	}
	return admin, nil
}
