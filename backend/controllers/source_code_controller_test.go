package controllers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"AST-Generator/controllers"
	"AST-Generator/db"
	"AST-Generator/models"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestSaveSourceCodeHandler(t *testing.T) {
	t.Log("Setting Gin to test mode")
	gin.SetMode(gin.TestMode)

	t.Log("Initializing in-memory test DB")
	testDB, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatal("Failed to open test database:", err)
	}
	db.DB = testDB

	t.Log("Running AutoMigrate for SourceCode model")
	if err := db.DB.AutoMigrate(&models.SourceCode{}); err != nil {
		t.Fatal("AutoMigrate error:", err)
	}

	t.Log("Setting up Gin router with /source_codes/save endpoint")
	router := gin.Default()
	router.POST("/source_codes/save", controllers.SaveSourceCodeHandler)

	t.Log("Creating test payload with title and description")
	payload := struct {
		Language    string `json:"language" binding:"required"`
		Code        string `json:"code" binding:"required"`
		Title       string `json:"title" binding:"required"`
		Description string `json:"description"`
	}{
		Language:    "go",
		Code:        "package main\n\nfunc main() {\n    println(\"Hello World\")\n}",
		Title:       "Hello World Example",
		Description: "This is a sample code that prints Hello World.",
	}
	jsonData, err := json.Marshal(payload)
	if err != nil {
		t.Fatal("JSON marshal error:", err)
	}

	t.Log("Creating HTTP POST request")
	req, err := http.NewRequest(http.MethodPost, "/source_codes/save", bytes.NewBuffer(jsonData))
	if err != nil {
		t.Fatal("Failed to create HTTP request:", err)
	}
	req.Header.Set("Content-Type", "application/json")

	t.Log("Sending the request and recording the response")
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	t.Log("Checking response status code")
	if recorder.Code != http.StatusOK {
		t.Fatalf("Expected status code 200, got %d", recorder.Code)
	}

	var resp map[string]interface{}
	if err := json.Unmarshal(recorder.Body.Bytes(), &resp); err != nil {
		t.Fatal("Error unmarshalling response:", err)
	}

	t.Log("Response body:", resp)

	if msg, ok := resp["message"].(string); !ok || msg != "Source code saved successfully" {
		t.Fatalf("Expected message 'Source code saved successfully', got %v", resp["message"])
	}

	var records []models.SourceCode
	if err := db.DB.Find(&records).Error; err != nil {
		t.Fatal("Failed to query saved records:", err)
	}
	t.Log("Saved records:", records)

	t.Log("Test completed successfully")
}
