package controllers

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/yardan-org/final-task-pbi-rakamin-fullstack-yardan/app/database"
	"github.com/yardan-org/final-task-pbi-rakamin-fullstack-yardan/app/models"
	"github.com/yardan-org/final-task-pbi-rakamin-fullstack-yardan/app/utils"
)

func ViewPhotoProfile(c *gin.Context) {
	fileName := c.Param("fileName")
	directory := "./photo/"
	filePath := filepath.Join(directory, fileName)

	c.File(filePath)
}

func AddPhotoProfile(c *gin.Context) {
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

	path := utils.GeneratePhotoPath(file)

	if err := c.SaveUploadedFile(file, path); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	photo := models.Photo{
		ID:       uuid.New(),
		Title:    file.Filename,
		Caption:  "photo " + file.Filename,
		PhotoUrl: path,
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

func UpdatePhotoProfile(c *gin.Context) {
	userID, _ := c.Get("userId")
	file, err := c.FormFile("file")

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to upload photo"})
		return
	}

	var photo models.Photo

	if err := database.DB.Where("user_id = ?", userID).First(&photo).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Photo not found"})
		return
	}

	path := utils.GeneratePhotoPath(file)

	if err := c.SaveUploadedFile(file, path); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save photo"})
		return
	}

	if err := os.Remove(photo.PhotoUrl); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete old photo"})
		return
	}

	photo.Title = file.Filename
	photo.Caption = "photo " + file.Filename
	photo.PhotoUrl = path

	if err := database.DB.Save(&photo).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update photo record"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Profile photo updated successfully"})
}

func DeletePhotoProfile(c *gin.Context) {
	userID, _ := c.Get("userId")

	var photo models.Photo

	if err := database.DB.Where("user_id = ?", userID).First(&photo).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Photo not found"})
		return
	}

	fmt.Println("Photo URL : ", photo.PhotoUrl)

	if err := os.Remove(photo.PhotoUrl); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := database.DB.Delete(&photo).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete photo record"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Profile photo deleted successfully"})
}
