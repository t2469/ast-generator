package controllers

import (
	"AST-Generator/db"
	"AST-Generator/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
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

func GetUserSourceCodesHandler(c *gin.Context) {
	userClaims, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	claims, ok := userClaims.(jwt.MapClaims)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse user information"})
		return
	}

	userID, ok := claims["sub"].(string)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "JWT does not contain user ID"})
		return
	}

	var codes []models.SourceCode
	if err := db.DB.Where("user_id = ?", userID).Find(&codes).Error; err != nil {
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

	userClaims, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}
	claims, ok := userClaims.(jwt.MapClaims)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse user information"})
		return
	}
	userID, ok := claims["sub"].(string)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "JWT does not contain user ID"})
		return
	}

	record := models.SourceCode{
		UserID:      userID,
		Title:       req.Title,
		Description: req.Description,
		Language:    req.Language,
		Code:        req.Code,
	}

	if err := db.DB.Create(&record).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save source code"})
		return
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
