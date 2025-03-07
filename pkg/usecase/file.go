package usecase

import (
	"file-sharing/pkg/domain"
	"file-sharing/pkg/repository/interfaces"
	interfacess "file-sharing/pkg/usecase/interfaces"
	"fmt"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"
)

type fileUseCase struct {
	repo interfaces.FileRepository
}

func NewFileUseCase(repo interfaces.FileRepository) interfacess.FileUseCase {
	return &fileUseCase{repo: repo}
}

func (uc *fileUseCase) UploadFile(userID uint, fileHeader *multipart.FileHeader) (string, error) {
	filePath := filepath.Join("uploads", fmt.Sprintf("%d-%s", time.Now().Unix(), fileHeader.Filename))

	// Open the file
	src, err := fileHeader.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	// Save file concurrently
	done := make(chan error)
	go func() {
		// Create destination file
		dst, err := os.Create(filePath)
		if err != nil {
			done <- err
			return
		}
		defer dst.Close()

		// Copy file content
		_, err = dst.ReadFrom(src)
		done <- err
	}()

	if err := <-done; err != nil {
		return "", err
	}

	// Save metadata
	file := domain.File{
		UserID:    userID,
		FileName:  fileHeader.Filename,
		FileSize:  fileHeader.Size,
		FileType:  fileHeader.Header.Get("Content-Type"),
		FileURL:   "http://localhost:8083/" + filePath,
		CreatedAt: time.Now(),
	}

	err = uc.repo.SaveFileMetadata(file)
	if err != nil {
		return "", err
	}

	return file.FileURL, nil
}
