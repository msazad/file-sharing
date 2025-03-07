package models

type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// User Signup
type UserDetails struct {
	Name            string `json:"name"`
	Email           string `json:"email" validate:"required"`
	Username        string `json:"username"`
	Phone           string `json:"phone"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
}

type UserToken struct {
	User  UserDetailsResponse
	Token string
}

// user details shown after loggin
type UserDetailsResponse struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email" validate:"email"`
	Phone string `json:"phone"`
}

type UserSignInResponse struct{
	Id       uint   `json:"id"`
	UserID   uint   `json:"user_id"`
	Name     string `json:"name"`
	Email    string `json:"email" validate:"email"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
}


