package repository

import (
	"context"
	"errors"

	"fmt"
	"golang_project_ecommerce/pkg/domain"
	interfaces "golang_project_ecommerce/pkg/repository/interface"

	"gorm.io/gorm"
)

type userDatabase struct {
	DB *gorm.DB
}

// constructor implement UserRepository interface return userDatabase struct

func NewUserRepository(DB *gorm.DB) interfaces.UserRepository {
	return &userDatabase{DB}
}

func (c *userDatabase) Addusers(ctx context.Context, user domain.Users) (domain.Users, error) {

	err := c.DB.Create(&user).Error

	if err != nil {
		return domain.Users{}, fmt.Errorf("error adding users: %w", err)
	}
	return user, nil
}

func (c *userDatabase) FindUser(ctx context.Context, user domain.Users) (domain.Users, error) {
	err := c.DB.Where("username = ? OR user_id = ? OR email = ?", user.Username, user.User_Id, user.Email).First(&user).Error
	if err != nil {
		return domain.Users{}, errors.New("user not found")
	}

	return user, nil
}
