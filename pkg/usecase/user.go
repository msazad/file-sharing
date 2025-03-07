package usecase

import (
	"errors"
	"file-sharing/pkg/helper"
	"file-sharing/pkg/repository/interfaces"
	"file-sharing/pkg/utils/models"

	"golang.org/x/crypto/bcrypt"
)

type userUsecase struct {
	userRepo interfaces.UserRepository
}

func NewUserUsecase(userRepo interfaces.UserRepository) *userUsecase {
	return &userUsecase{
		userRepo: userRepo,
	}
}

func (usrU *userUsecase) Login(user models.UserLogin) (models.UserToken, error) {
	// check the user already exist or not

	ok := usrU.userRepo.CheckUserAvailability(user.Email)
	if !ok {
		return models.UserToken{}, errors.New("user not exist")
	}
	// Get the user details in order to check password
	user_details, err := usrU.userRepo.FindUserByEmail(user)
	if err != nil {
		return models.UserToken{}, err
	}
	// check the password
	err = bcrypt.CompareHashAndPassword([]byte(user_details.Password), []byte(user.Password))
	if err != nil {
		return models.UserToken{}, errors.New("password incorrect")
	}

	var userResponse models.UserDetailsResponse
	userResponse.Id = int(user_details.Id)
	userResponse.Name = user_details.Name
	userResponse.Email = user_details.Email
	userResponse.Phone = user_details.Phone

	// generate token
	tokenString, err := helper.GenerateUserToken(userResponse)
	if err != nil {
		return models.UserToken{}, errors.New("could't create token for user")
	}
	return models.UserToken{
		User:  userResponse,
		Token: tokenString,
	}, nil
}

func (usrU *userUsecase) SignUp(user models.UserDetails) (models.UserToken, error) {
	// check the user exist or not,if exist show the error(its a signup function)

	userExist := usrU.userRepo.CheckUserAvailability(user.Email)
	if userExist {
		return models.UserToken{}, errors.New("user already exist please sign in")
	}
	if user.Password != user.ConfirmPassword {
		return models.UserToken{}, errors.New("password does't match")
	}
	// hash the password
	hashedPass, err := helper.PasswordHashing(user.Password)
	if err != nil {
		return models.UserToken{}, err
	}
	user.Password = hashedPass
	// insert the user into database
	userData, err := usrU.userRepo.SignUp(user)
	if err != nil {
		return models.UserToken{}, err
	}
	// create jwt token for user
	tokenString, err := helper.GenerateUserToken(userData)
	if err != nil {
		return models.UserToken{}, errors.New("couldn't create token for user due to some internal error")
	}
	// create new wallet for user
	// if _, err := usrU.orderRepo.CreateNewWallet(userData.Id); err != nil {
	// 	return models.UserToken{}, errors.New("error creating new wallet for user")
	// }
	return models.UserToken{
		User:  userData,
		Token: tokenString,
	}, nil

}
