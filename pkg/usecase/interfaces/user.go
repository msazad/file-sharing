package interfaces

import "file-sharing/pkg/utils/models"

type UserUsecase interface {
	Login(user models.UserLogin) (models.UserToken, error)
	SignUp(user models.UserDetails) (models.UserToken, error)
}
