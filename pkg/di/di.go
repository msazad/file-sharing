package di

import (
	"file-sharing/pkg/api"
	"file-sharing/pkg/api/handlers"
	"file-sharing/pkg/config"
	"file-sharing/pkg/db"
	"file-sharing/pkg/repository"
	"file-sharing/pkg/usecase"
)

func InitializeAPI(cfg config.Config)(*api.ServerHTTP,error) {
	gormDB, err := db.ConnectDB(cfg)
	if err != nil {
		return nil, err
	}

	userRepository := repository.NewUserRepository(gormDB)
	userUsecase := usecase.NewUserUsecase(userRepository)
	UserHandler := handlers.NewUserHandler(userUsecase)

	fileRepo := repository.NewFileRepository(gormDB)
	fileUseCase := usecase.NewFileUseCase(fileRepo)
	fileHandler := handlers.NewFileHandler(fileUseCase)

	serverHTTP := api.NewServerHttp(UserHandler,fileHandler)
	return serverHTTP,nil
}
