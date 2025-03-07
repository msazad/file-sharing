package interfaces

import "file-sharing/pkg/domain"

type FileRepository interface {
	SaveFileMetadata(file domain.File) error
	GetFileByID(id uint) (domain.File, error)
}
