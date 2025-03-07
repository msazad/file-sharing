package interfaces

import "mime/multipart"

type FileUseCase interface {
	UploadFile(userID uint, fileHeader *multipart.FileHeader) (string, error)
}
