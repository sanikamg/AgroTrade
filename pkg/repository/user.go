package repository

import (
	"context"
	"errors"

	"fmt"
	"golang_project_ecommerce/pkg/domain"
	interfaces "golang_project_ecommerce/pkg/repository/interface"
	"golang_project_ecommerce/pkg/utils/res"

	"gorm.io/gorm"
)

type userDatabase struct {
	DB *gorm.DB
}

// constructor implement UserRepository interface return userDatabase struct

func NewUserRepository(DB *gorm.DB) interfaces.UserRepository {
	return &userDatabase{DB}
}

func (ud *userDatabase) Addusers(ctx context.Context, user domain.Users) (domain.Users, error) {

	err := ud.DB.Create(&user).Error

	if err != nil {
		return domain.Users{}, fmt.Errorf("error adding users: %w", err)
	}
	return user, nil
}

func (ud *userDatabase) FindUser(ctx context.Context, user domain.Users) (domain.Users, error) {
	err := ud.DB.Where("username = ? OR user_id = ? OR email = ?", user.Username, user.User_Id, user.Email).First(&user).Error
	if err != nil {
		return domain.Users{}, errors.New("user not found")
	}

	return user, nil
}

// list all mproducts on user side
// for finding allusers
func (ud *userDatabase) FindAllProducts(c context.Context, categoryid uint) ([]res.AllProducts, error) {
	var products []res.AllProducts
	err := ud.DB.Raw("select product_name,product_price,product_quantity,image from product_details where category_id=? ", categoryid).Scan(&products).Error
	if err != nil {
		return []res.AllProducts{}, errors.New("failed to find all products")
	}
	return products, nil

}

// category
func (pd *userDatabase) GetCategoryByName(c context.Context, categoryname string) (domain.Category, error) {
	var category domain.Category
	query := `select * from categories where category_name=?`
	err := pd.DB.Raw(query, categoryname).First(&category).Error
	if err != nil {
		return domain.Category{}, errors.New("failed to find category name(first letter must be capital)")
	}

	return category, nil
}
