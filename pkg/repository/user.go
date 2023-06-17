package repository

import (
	"context"
	"errors"

	"fmt"
	"golang_project_ecommerce/pkg/domain"
	interfaces "golang_project_ecommerce/pkg/repository/interface"
	"golang_project_ecommerce/pkg/utils"
	"golang_project_ecommerce/pkg/utils/req"
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

// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>user signup>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>.....
func (ud *userDatabase) Addusers(ctx context.Context, user domain.Users) (domain.Users, error) {

	err := ud.DB.Create(&user).Error

	if err != nil {
		return domain.Users{}, fmt.Errorf("error adding users: %w", err)
	}
	return user, nil
}

// to check user already exist or not
func (ud *userDatabase) FindUser(ctx context.Context, user domain.Users) (domain.Users, error) {
	err := ud.DB.Where("username = ?  OR email = ?", user.Username, user.Email).First(&user).Error
	if err != nil {
		return domain.Users{}, errors.New("user not found")
	}

	return user, nil
}

// find user by phonenumber
func (ud *userDatabase) FindUserByPhn(c context.Context, phn domain.Users) error {
	err := ud.DB.Where("phone=?", phn.Phone).First(&phn).Error
	if err != nil {
		return errors.New("failed to find user")
	}
	return nil
}

// finding username is empty
func (ud *userDatabase) IsEmtyUsername(c context.Context, username domain.Users) bool {
	ud.DB.Where("phone = ?", username.Phone).First(&username)
	return username.Username == ""
}

//update user status

func (ud *userDatabase) UpdateStatus(c context.Context, user domain.Users) error {
	query := `update users set verification=? where phone=?`
	err := ud.DB.Raw(query, true, user.Phone).Scan(&user).Error
	if err != nil {
		return errors.New("failed to update update status")
	}
	return nil
}

//find status to update details

func (ud *userDatabase) FindStatus(c context.Context, phn string) (domain.Users, error) {
	var usr domain.Users
	query := `select *from users where phone=?`
	err := ud.DB.Raw(query, phn).Scan(&usr).Error
	if err != nil {
		return domain.Users{}, errors.New("failed to find status")
	}
	return usr, nil

}

// complete user profile
func (ud *userDatabase) UpdateUserDetails(c context.Context, user domain.Users) (domain.Users, error) {

	query := `update users set username=?,name=?,email=?,password=? where phone=?`
	err := ud.DB.Raw(query, user.Username, user.Name, user.Email, user.Password, user.Phone).Scan(&user).Error
	if err != nil {
		return domain.Users{}, errors.New("failed to complete user registration")
	}
	return user, nil
}

//>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>end>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>.

// >>>>>>>>>>>>>>>>>>>>>>>>>>> list all products on user side>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>

// for finding allusers
func (ud *userDatabase) FindAllProducts(c context.Context, categoryid uint) ([]res.AllProducts, error) {
	var products []res.AllProducts
	err := ud.DB.Raw("select product_name,product_price,product_quantity from product_details where category_id=? AND deleted_at IS NULL", categoryid).Scan(&products).Error
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

//>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>end>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>..

// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>user management>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>.
func (ud *userDatabase) FindUserById(c context.Context, id int) (domain.Users, error) {
	var user domain.Users
	err := ud.DB.Raw("select *from users where user_id=?", id).Scan(&user).Error
	if err != nil {
		return domain.Users{}, errors.New("user didn't exist")
	}
	return user, nil
}

func (ud *userDatabase) EditUserDetails(c context.Context, id int, user req.Usereditreq) (domain.Users, error) {
	var usr domain.Users
	query := `update users set username=?,name=?,email=?,password=?,phone=? where user_id=?`
	err := ud.DB.Raw(query, user.Username, user.Name, user.Email, user.Password, user.Phone, id).Scan(&usr).Error
	if err != nil {
		return domain.Users{}, errors.New("failed to complete user registration")
	}

	err1 := ud.DB.Raw("select *from users where user_id=?", id).Scan(&usr).Error
	if err1 != nil {
		return domain.Users{}, errors.New("failed to complete user registration")
	}
	return usr, nil
}

// >>>>>>>>>>>>>>>>>>>>>>>address management>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
func (ud *userDatabase) AddAddress(c context.Context, address domain.Address) (res.AddressResponse, error) {
	err := ud.DB.Create(&address).Error
	if err != nil {
		return res.AddressResponse{}, errors.New("failed to add address")
	}
	var ads res.AddressResponse
	query := `select *from addresses where address_id=? `
	err1 := ud.DB.Raw(query, address.Address_Id).Scan(&ads).Error
	if err1 != nil {
		return res.AddressResponse{}, errors.New("failed to fetch address")
	}

	return ads, nil
}
func (ud *userDatabase) FindAddress(c context.Context, addressid uint) error {
	var address domain.Address

	err := ud.DB.Where("address_id=?", addressid).First(&address).Error
	if err != nil {
		return errors.New("address didn't exist")
	}
	return nil
}

func (ud *userDatabase) EditAddress(c context.Context, address domain.Address) (res.AddressResponse, error) {
	var ads res.AddressResponse
	query := `update addresses set full_name=?,phone_number=?,House_name=?,post_office=?,area=?,landmark=?,district=?,state=?,pin=? where address_id=?`
	err := ud.DB.Raw(query, address.Full_name, address.Phone_number, address.House_name, address.Post_office, address.Area, address.Landmark, address.District, address.State, address.Pin, address.Address_Id).Scan(&address).Error
	if err != nil {
		return res.AddressResponse{}, errors.New("failed to update address")
	}
	query2 := `select *from addresses where address_id=? `
	err1 := ud.DB.Raw(query2, address.Address_Id).Scan(&ads).Error
	if err1 != nil {
		return res.AddressResponse{}, errors.New("failed to fetch address")
	}

	return ads, nil
}

func (ud *userDatabase) ListAddresses(c context.Context, pagination utils.Pagination, id uint) ([]res.AddressResponse, utils.Metadata, error) {
	var ads []res.AddressResponse
	var totalRecords int64

	db := ud.DB.Model(&domain.Address{})

	// Count all records
	if err := db.Count(&totalRecords).Error; err != nil {
		return []res.AddressResponse{}, utils.Metadata{}, err
	}
	query := `select *from addresses where user_id=$1 limit $2 offset $3`
	err := db.Raw(query, id, pagination.Limit(), pagination.Offset()).Scan(&ads).Error
	if err != nil {
		return []res.AddressResponse{}, utils.Metadata{}, errors.New("failed to work query")
	}
	// Compute metadata
	metadata := utils.ComputeMetadata(&totalRecords, &pagination.Page, &pagination.PageSize)

	return ads, metadata, nil

}

func (ud *userDatabase) DeleteAddress(c context.Context, addressid uint) error {
	var address domain.Address
	err := ud.DB.Where("address_id=?", addressid).Delete(&address).Error
	if err != nil {
		return errors.New("failed to delete address")
	}
	return nil
}

// <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<forgot password>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>

func (ud *userDatabase) FindUserByPhnNum(c context.Context, phn string) error {
	var user domain.Users

	err := ud.DB.Where("phone=?", phn).First(&user).Error
	if err != nil {
		return errors.New("failed to fetch user account")
	}
	return nil
}

func (ud *userDatabase) ForgotPassword(c context.Context, usrphn string, newpass string) error {
	var user domain.Users
	query := `update users set password=? where phone=?`
	err := ud.DB.Raw(query, newpass, usrphn).Scan(&user).Error
	if err != nil {
		return errors.New("failed to update password")
	}
	return nil
}
