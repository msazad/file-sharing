package helper

import (
	"context"
	"errors"
	"file-sharing/pkg/utils/models"
	"fmt"
	"mime/multipart"
	"time"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/golang-jwt/jwt"
	"github.com/jinzhu/copier"
	"golang.org/x/crypto/bcrypt"
)

type AuthCustomClaims struct {
	Id int `json:"id"`
	// Name  string `json:"name"`
	Email string `json:"email"`
	jwt.StandardClaims
}

func GenerateUserToken(user models.UserDetailsResponse) (string, error) {
	claims := &AuthCustomClaims{
		Id:    user.Id,
		Email: user.Email,

		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 48).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte("usersecret"))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func PasswordHashing(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return "", errors.New("internal server error")
	}
	hash := string(hashedPassword)
	return hash, nil
}

func AddImageToS3(file *multipart.FileHeader) (string, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("ap-south-1"))
	if err != nil {
		fmt.Println("configuration error: ", err)
		return "", err
	}
	client := s3.NewFromConfig(cfg)
	uploader := manager.NewUploader(client)
	f, openErr := file.Open()
	if openErr != nil {
		fmt.Println("file open error: ", openErr)
		return "", openErr
	}
	defer f.Close()

	result, uploadErr := uploader.Upload(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String("samplebucket"),
		Key:    aws.String(file.Filename),
		Body:   f,
		ACL:    "public-read",
	})
	if uploadErr != nil {
		fmt.Println("upload error : ", uploadErr)
		return "", uploadErr
	}
	return result.Location, nil
}

func CompareHashAndPassword(a string, b string) error {
	err := bcrypt.CompareHashAndPassword([]byte(a), []byte(b))
	if err != nil {
		return err
	}
	return nil
}

func Copy(a *models.UserDetailsResponse, b *models.UserSignInResponse) (models.UserDetailsResponse, error) {
	if err := copier.Copy(a, b); err != nil {
		return models.UserDetailsResponse{}, err
	}
	return *a, nil
}
