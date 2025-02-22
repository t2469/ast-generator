package controllers

import (
	"AST-Generator/db"
	"AST-Generator/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type SaveSourceCodeRequest struct {
	Language    string `json:"language" binding:"required"`
	Code        string `json:"code" binding:"required"`
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
}

func GetAllSourceCodesHandler(c *gin.Context) {
	var codes []models.SourceCode
	if err := db.DB.Find(&codes).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch source codes"})
		return
	}
	c.JSON(http.StatusOK, codes)
}

func SaveSourceCodeHandler(c *gin.Context) {
	var req SaveSourceCodeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	record := models.SourceCode{
		Title:       req.Title,
		Description: req.Description,
		Language:    req.Language,
		Code:        req.Code,
	}

	if err := db.DB.Create(&record).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save source code"})
	}

	c.JSON(http.StatusOK, gin.H{"message": "Source code saved successfully"})
}

func DeleteSourceCodeHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	result := db.DB.Delete(&models.SourceCode{}, id)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete source code"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Source code deleted successfully"})
}
