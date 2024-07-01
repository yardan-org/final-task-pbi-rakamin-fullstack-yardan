package utils

import (
	"mime/multipart"
	"path/filepath"

	"github.com/google/uuid"
)

func GeneratePhotoPath(file *multipart.FileHeader) string {
	ext := filepath.Ext(file.Filename)
	newFileName := uuid.New().String() + ext
	savePath := filepath.Join("photo", newFileName)

	return filepath.ToSlash(savePath)
}
