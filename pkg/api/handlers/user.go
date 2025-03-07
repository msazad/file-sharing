package handlers

import (
	"file-sharing/pkg/utils/models"
	"file-sharing/pkg/utils/response"
	"net/http"

	"github.com/gin-gonic/gin"
	
	services"file-sharing/pkg/usecase/interfaces"
)

type UserHandler struct {
	userusecase services.UserUsecase
}

// Constructor function
func NewUserHandler(userUsecase services.UserUsecase) *UserHandler {
	return &UserHandler{
		userusecase: userUsecase,
	}
}

func (uH *UserHandler) SignUp(c *gin.Context) {
	var user models.UserDetails

	if err := c.BindJSON(&user); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "fields provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	userToken, err := uH.userusecase.SignUp(user)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "couldn't signup user", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "successfully signed up", userToken, nil)
	c.JSON(http.StatusOK, successRes)
}

func (uH *UserHandler) Login(c *gin.Context) {
	var user models.UserLogin
	if err := c.BindJSON(&user); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "fields provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	userToken, err := uH.userusecase.Login(user)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "user couldn't login", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "user successfully logged in", userToken, nil)
	// c.SetCookie("Authorization",userToken.Token,3600,"/","primemart.online",true,false)
	c.SetCookie("Authorization", userToken.Token, 3600*24*30, "", "", false, true)

	c.JSON(http.StatusOK, successRes)
}
