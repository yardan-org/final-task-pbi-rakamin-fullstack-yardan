package controllers

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/yardan-org/final-task-pbi-rakamin-fullstack-yardan/app/database"
	"github.com/yardan-org/final-task-pbi-rakamin-fullstack-yardan/app/models"
)

func UploadPhotoProfile(c *gin.Context) {
	userID, _ := c.Get("userId")
	file, err := c.FormFile("file")

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to upload photo"})
		return
	}

	var user models.User

	if err := database.DB.First(&user, "id = ?", userID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	ext := filepath.Ext(file.Filename)
	newFileName := uuid.New().String() + ext
	savePath := filepath.Join("photos", newFileName)
	savePath = filepath.ToSlash(savePath)

	if err := c.SaveUploadedFile(file, savePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save photo"})
		return
	}

	var photo models.Photo
	photoExists := database.DB.Where("user_id = ?", userID).First(&photo).Error == nil

	if photoExists {

		if err := os.Remove(photo.PhotoUrl); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete old photo"})
			return
		}

		if database.DB.Delete(&photo).Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete photo record"})
			return
		}

	}

	photo = models.Photo{
		ID:       uuid.New(),
		Title:    file.Filename,
		Caption:  "photo " + file.Filename,
		PhotoUrl: savePath,
		UserID:   userID.(uuid.UUID),
	}

	if err := database.DB.Create(&photo).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	user.Photo = &photo

	if err := database.DB.Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user photo"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Profile photo uploaded successfully"})
}
