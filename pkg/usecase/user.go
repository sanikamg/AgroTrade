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
	"golang_project_ecommerce/pkg/verification"

	"golang.org/x/crypto/bcrypt"
)

type UserUsecase struct {
	userRepo interfaces.UserRepository
}

func NewUserUsecase(repo interfaces.UserRepository) ser.UserUsecase {
	return &UserUsecase{
		userRepo: repo,
	}
} //>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>login/sign up>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
// func generateOTPCode() string {
// 	// Generate a random 6-digit OTP code
// 	return fmt.Sprintf("%06d", rand.Intn(1000000))
// }

func (uu *UserUsecase) SendOtpPhn(c context.Context, phn domain.Users) error {
	err := uu.userRepo.FindUserByPhn(c, phn)
	if err == nil {
		if uu.userRepo.IsEmtyUsername(c, phn) {
			return errors.New("user verification already completed please complete registration")
		}
		return errors.New("user already exists please login")
	}
	// Generate OTP code

	if _, err1 := verification.SendOtp("+91" + phn.Phone); err1 != nil {

		return errors.New("can't send otp")
	}

	return nil
}

// func (c *UserUsecase) Adduser(ctx context.Context, user domain.Users) (domain.Users, error) {

// 	c.userRepo.Addusers(ctx, user)

// 	return user, nil

// }

//verify otp

func (uu *UserUsecase) VerifyOtp(c context.Context, phn string, otp string) error {
	var usr domain.Users
	err := verification.VerifyOtp("+91"+phn, otp)
	if err != nil {
		return errors.New("failed to verify otp")
	}
	usr.Phone = phn
	_, er := uu.userRepo.Addusers(c, usr)
	if er != nil {
		return errors.New("can't add user ")
	}
	return nil
}

func (uu *UserUsecase) UpdateStatus(c context.Context, user domain.Users) error {

	err := uu.userRepo.UpdateStatus(c, user)
	if err != nil {
		return err
	}
	return nil
}

func (uu *UserUsecase) Register(ctx context.Context, user domain.Users) (domain.Users, error) {
	//to hash password
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)

	if err != nil {
		return domain.Users{}, errors.New("bcrypt failed:" + err.Error())
	}
	user.Password = string(hash)
	usr, err := uu.userRepo.FindStatus(ctx, user.Phone)
	if err != nil {
		return domain.Users{}, err
	}
	user.Verification = usr.Verification
	user.User_Id = usr.User_Id

	if user.Verification {
		usr, err := uu.userRepo.UpdateUserDetails(ctx, user)
		if err != nil {
			return domain.Users{}, err
		}

		return usr, nil
	}
	return domain.Users{}, errors.New("enter correct details")
}

// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>.login>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
func (c *UserUsecase) Login(ctx context.Context, user domain.Users) (domain.Users, error) {
	dbUser, dbErr := c.userRepo.FindUser(ctx, user)

	//check whether the user exists or valid information
	if dbErr != nil {
		return domain.Users{}, dbErr
	} else if dbUser.User_Id == 0 {
		return domain.Users{}, errors.New("user does not exists with this , please register")
	}

	//checking block status

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

// func (uu *UserUsecase) FindAllProducts(c context.Context, categoryname string) ([]res.AllProducts, error) {
// 	category, err := uu.userRepo.GetCategoryByName(c, categoryname)
// 	if err != nil {
// 		return []res.AllProducts{}, err
// 	}

// 	products, err := uu.userRepo.FindAllProducts(c, category.ID)
// 	if err != nil {
// 		return []res.AllProducts{}, err
// 	}

// 	return products, nil
// }

// >>>>>>>>>>>>>>>>>>>>>>>>>>user profile>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>..
func (ur *UserUsecase) FindUserById(c context.Context, id int) (domain.Users, error) {
	user, err := ur.userRepo.FindUserById(c, id)
	if err != nil {
		return domain.Users{}, err
	}
	return user, nil
}

func (ur *UserUsecase) EditUserDetails(c context.Context, id int, user req.Usereditreq) (domain.Users, error) {
	_, err := ur.userRepo.FindUserById(c, id)
	if err != nil {
		return domain.Users{}, errors.New("user doesn't exist invalid id")
	}
	usr, err := ur.userRepo.EditUserDetails(c, id, user)
	if err != nil {
		return domain.Users{}, err
	}

	return usr, nil
}

// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>user address>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
func (ur *UserUsecase) AddAddress(c context.Context, address domain.Address) (res.AddressResponse, error) {
	ads, err := ur.userRepo.AddAddress(c, address)
	if err != nil {
		return res.AddressResponse{}, err
	}

	return ads, nil
}

func (ur *UserUsecase) EditAddress(c context.Context, address domain.Address) (res.AddressResponse, error) {
	err := ur.userRepo.FindAddress(c, address.Address_Id)
	if err != nil {
		return res.AddressResponse{}, err
	}
	ads, err := ur.userRepo.EditAddress(c, address)
	if err != nil {
		return res.AddressResponse{}, err
	}
	return ads, nil
}

func (uu *UserUsecase) ListAddresses(c context.Context, pagination utils.Pagination, id uint) ([]res.AddressResponse, utils.Metadata, error) {
	address, metadata, err := uu.userRepo.ListAddresses(c, pagination, id)
	if err != nil {
		return []res.AddressResponse{}, utils.Metadata{}, err
	}
	return address, metadata, nil
}

func (ur *UserUsecase) DeleteAddress(c context.Context, addressid uint) error {
	err := ur.userRepo.FindAddress(c, addressid)
	if err != nil {
		return err
	}
	err1 := ur.userRepo.DeleteAddress(c, addressid)
	if err1 != nil {
		return err
	}
	return nil
}

// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>forgot password>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>

func (uu *UserUsecase) SendOtpForgotPass(c context.Context, phn string) error {

	err := uu.userRepo.FindUserByPhnNum(c, phn)
	if err != nil {
		return err
	}
	// Generate OTP code

	if _, err1 := verification.SendOtp("+91" + phn); err1 != nil {

		return errors.New("can't send otp")
	}

	return nil
}

func (uu *UserUsecase) VerifyOtpForgotpass(c context.Context, phn string, otp string) error {

	err := verification.VerifyOtp("+91"+phn, otp)
	if err != nil {
		return errors.New("failed to verify otp")
	}

	return nil
}

func (ur *UserUsecase) ForgotPassword(c context.Context, usrphn string, newpass string) error {
	//to hash password
	hash, err := bcrypt.GenerateFromPassword([]byte(newpass), 14)
	if err != nil {
		return errors.New("failed to hash password")
	}
	newpass = string(hash)
	err1 := ur.userRepo.ForgotPassword(c, usrphn, newpass)
	if err1 != nil {
		return err1
	}
	return nil
}
