package handlers

import (
	"net/http"
	"strconv"

	services "file-sharing/pkg/usecase/interfaces"

	"github.com/gin-gonic/gin"
)

type FileHandler struct {
	fileUseCase services.FileUseCase
}

func NewFileHandler(fileUseCase services.FileUseCase) *FileHandler {
	return &FileHandler{fileUseCase: fileUseCase}
}

func (h *FileHandler) UploadFile(c *gin.Context) {
	userID, _ := strconv.ParseUint(c.Request.Header.Get("User-ID"), 10, 64)

	// Get file from request
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid file"})
		return
	}

	// Upload file
	fileURL, err := h.fileUseCase.UploadFile(uint(userID), file)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upload file"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "File uploaded successfully", "url": fileURL})
}
