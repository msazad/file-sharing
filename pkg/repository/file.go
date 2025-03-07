package repository

import (
	"file-sharing/pkg/domain"
	"file-sharing/pkg/repository/interfaces"

	"gorm.io/gorm"
)

type fileRepository struct {
	DB *gorm.DB
}

func NewFileRepository(DB *gorm.DB) interfaces.FileRepository {
	return &fileRepository{DB: DB}
}

func (r *fileRepository) SaveFileMetadata(file domain.File) error {
	return r.DB.Create(&file).Error
}

func (r *fileRepository) GetFileByID(id uint) (domain.File, error) {
	var file domain.File
	err := r.DB.First(&file, id).Error
	return file, err
}
