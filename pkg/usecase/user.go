package usecase

import (
	"context"
	"errors"
	"fmt"
	"golang_project_ecommerce/pkg/domain"
	interfaces "golang_project_ecommerce/pkg/repository/interface"
	ser "golang_project_ecommerce/pkg/usecase/interface"
	"golang_project_ecommerce/pkg/utils/res"

	"golang.org/x/crypto/bcrypt"
)

type UserUsecase struct {
	userRepo interfaces.UserRepository
}

func NewUserUsecase(repo interfaces.UserRepository) ser.UserUsecase {
	return &UserUsecase{
		userRepo: repo,
	}
}

func (c *UserUsecase) Register(ctx context.Context, user domain.Users) (domain.Users, error) {

	usr, err := c.userRepo.FindUser(ctx, user)
	if err == nil {
		return usr, errors.New("user already exist")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)

	if err != nil {
		return domain.Users{}, errors.New("bcrypt failed:" + err.Error())
	}
	user.Password = string(hash)

	c.userRepo.Addusers(ctx, user)

	return user, err

}

func (c *UserUsecase) Login(ctx context.Context, user domain.Users) (domain.Users, error) {
	dbUser, dbErr := c.userRepo.FindUser(ctx, user)

	//check whether the user exists or valid information
	if dbErr != nil {
		return domain.Users{}, dbErr
	} else if dbUser.User_Id == 0 {
		return domain.Users{}, errors.New("user does not exists with this , please register")
	}

	//checking block status

	fmt.Println("user", dbUser)

	if dbUser.BlockStatus {
		return domain.Users{}, errors.New("blocked user trying to login")
	}

	// check password matching

	if bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(user.Password)) != nil {
		return domain.Users{}, errors.New("password is not correct")
	}

	return dbUser, nil
}

//list all product on user side

func (uu *UserUsecase) FindAllProducts(c context.Context, categoryname string) ([]res.AllProducts, error) {
	category, err := uu.userRepo.GetCategoryByName(c, categoryname)
	if err != nil {
		return []res.AllProducts{}, err
	}

	products, err := uu.userRepo.FindAllProducts(c, category.ID)
	if err != nil {
		return []res.AllProducts{}, err
	}

	return products, nil
}
