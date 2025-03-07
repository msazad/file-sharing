package interfaces

import "file-sharing/pkg/utils/models"

type UserRepository interface {
	SignUp(user models.UserDetails) (models.UserDetailsResponse, error)
	FindUserByEmail(user models.UserLogin) (models.UserSignInResponse, error)
	CheckUserAvailability(email string) bool
}
