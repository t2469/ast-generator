package controllers

import (
	"AST-Generator/db"
	"AST-Generator/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SaveSourceCodeRequest struct {
	Language    string `json:"language" binding:"required"`
	Code        string `json:"code" binding:"required"`
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
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
