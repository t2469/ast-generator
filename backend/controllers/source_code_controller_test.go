package controllers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"AST-Generator/config"
	"AST-Generator/controllers"
	"AST-Generator/db"
	"AST-Generator/models"
	"AST-Generator/services"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestGetAllSourceCodesHandler(t *testing.T) {
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

	records := []models.SourceCode{
		{
			Title:       "Title 1",
			Description: "Description 1",
			Language:    "go",
			Code:        "package main\n\nfunc main() { println(\"Hello 1\") }",
		},
		{
			Title:       "Title 2",
			Description: "Description 2",
			Language:    "cpp",
			Code:        "#include <iostream>\n\nint main() { std::cout << \"Hello 2\"; return 0; }",
		},
	}
	for _, rec := range records {
		if err := db.DB.Create(&rec).Error; err != nil {
			t.Fatal("Failed to create test record:", err)
		}
	}

	t.Log("Setting up Gin router with /source_codes endpoint")
	router := gin.Default()
	router.GET("/source_codes", controllers.GetAllSourceCodesHandler)

	t.Log("Creating HTTP GET request")
	req, err := http.NewRequest(http.MethodGet, "/source_codes", nil)
	if err != nil {
		t.Fatal("Failed to create HTTP request:", err)
	}

	t.Log("Sending the request and recording the response")
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	t.Log("Checking response status code")
	if recorder.Code != http.StatusOK {
		t.Fatalf("Expected status code 200, got %d", recorder.Code)
	}

	var fetchedRecords []models.SourceCode
	if err := json.Unmarshal(recorder.Body.Bytes(), &fetchedRecords); err != nil {
		t.Fatal("Error unmarshalling response:", err)
	}

	t.Log("Fetched records:", fetchedRecords)
	if len(fetchedRecords) != len(records) {
		t.Fatalf("Expected %d records, got %d", len(records), len(fetchedRecords))
	}

	t.Log("TestGetAllSourceCodesHandler completed successfully")
}

func TestGetUserSourceCodesHandler(t *testing.T) {
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

	records := []models.SourceCode{
		{
			Title:       "Title 1",
			Description: "Description 1",
			Language:    "go",
			Code:        "package main\nfunc main() { println(\"Hello 1\") }",
			UserID:      "test-user-id",
		},
		{
			Title:       "Title 2",
			Description: "Description 2",
			Language:    "cpp",
			Code:        "#include <iostream>\nint main() { std::cout << \"Hello 2\"; return 0; }",
			UserID:      "test-user-id",
		},
		{
			Title:       "Other Title",
			Description: "Other Description",
			Language:    "python",
			Code:        "print('Hello from other user')",
			UserID:      "other-user-id",
		},
	}
	for _, rec := range records {
		if err := db.DB.Create(&rec).Error; err != nil {
			t.Fatal("Failed to create test record:", err)
		}
	}

	t.Log("Generating test JWT for user 'test-user-id'")
	jwtToken, err := services.GenerateJWT("test-user-id", "test@example.com", "Test User", "https://example.com/test.jpg", config.JwtSecret)
	if err != nil {
		t.Fatal("Failed to generate test JWT:", err)
	}

	t.Log("Setting up Gin router with /source_codes/user endpoint and authentication middleware")
	router := gin.Default()
	router.GET("/source_codes/user", controllers.JWTAuthMiddleware(), controllers.GetUserSourceCodesHandler)

	t.Log("Creating HTTP GET request for /source_codes/user")
	req, err := http.NewRequest(http.MethodGet, "/source_codes/user", nil)
	if err != nil {
		t.Fatal("Failed to create HTTP request:", err)
	}
	req.AddCookie(&http.Cookie{
		Name:  "auth_token",
		Value: jwtToken,
	})

	t.Log("Sending the request and recording the response")
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	t.Log("Checking response status code")
	if recorder.Code != http.StatusOK {
		t.Fatalf("Expected status code 200, got %d", recorder.Code)
	}

	var fetchedRecords []models.SourceCode
	if err := json.Unmarshal(recorder.Body.Bytes(), &fetchedRecords); err != nil {
		t.Fatal("Error unmarshalling response:", err)
	}

	t.Log("Fetched records:", fetchedRecords)
	expectedCount := 2
	if len(fetchedRecords) != expectedCount {
		t.Fatalf("Expected %d records for user 'test-user-id', got %d", expectedCount, len(fetchedRecords))
	}
	for _, r := range fetchedRecords {
		if r.UserID != "test-user-id" {
			t.Fatalf("Expected record UserID 'test-user-id', got %s", r.UserID)
		}
	}

	t.Log("TestGetUserSourceCodesHandler completed successfully")
}

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

	jwtToken, err := services.GenerateJWT("test-user-id", "test@example.com", "Test User", "https://example.com/test.jpg", config.JwtSecret)
	if err != nil {
		t.Fatal("Failed to generate test JWT:", err)
	}

	t.Log("Setting up Gin router with /source_codes/save endpoint with authentication middleware")
	router := gin.Default()
	router.POST("/source_codes/save", controllers.JWTAuthMiddleware(), controllers.SaveSourceCodeHandler)

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
	req.AddCookie(&http.Cookie{
		Name:  "auth_token",
		Value: jwtToken,
	})

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
	if len(records) != 1 {
		t.Fatalf("Expected 1 record, got %d", len(records))
	}
	if records[0].UserID != "test-user-id" {
		t.Fatalf("Expected UserID 'test-user-id', got %v", records[0].UserID)
	}

	t.Log("TestSaveSourceCodeHandler completed successfully")
}

func TestDeleteSourceCodeHandler(t *testing.T) {
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

	source := models.SourceCode{
		Title:       "Test Title",
		Description: "Test Description",
		Language:    "go",
		Code:        "package main\n\nfunc main() { println(\"Hello\") }",
		UserID:      "test-user-id",
	}
	if err := db.DB.Create(&source).Error; err != nil {
		t.Fatal("Failed to create test record:", err)
	}

	jwtToken, err := services.GenerateJWT("test-user-id", "test@example.com", "Test User", "https://example.com/test.jpg", config.JwtSecret)
	if err != nil {
		t.Fatal("Failed to generate test JWT:", err)
	}

	router := gin.Default()
	router.DELETE("/source_codes/:id", controllers.JWTAuthMiddleware(), controllers.DeleteSourceCodeHandler)

	url := "/source_codes/" + strconv.Itoa(int(source.ID))
	t.Log("Creating HTTP DELETE request for URL:", url)
	req, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		t.Fatal("Failed to create HTTP request:", err)
	}
	req.AddCookie(&http.Cookie{
		Name:  "auth_token",
		Value: jwtToken,
	})

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	if recorder.Code != http.StatusOK {
		t.Fatalf("Expected status code 200, got %d", recorder.Code)
	}

	var resp map[string]interface{}
	if err := json.Unmarshal(recorder.Body.Bytes(), &resp); err != nil {
		t.Fatal("Error unmarshalling response:", err)
	}
	t.Log("Response body:", resp)

	if msg, ok := resp["message"].(string); !ok || msg != "Source code deleted successfully" {
		t.Fatalf("Expected message 'Source code deleted successfully', got %v", resp["message"])
	}

	var count int64
	if err := db.DB.Model(&models.SourceCode{}).Where("id = ?", source.ID).Count(&count).Error; err != nil {
		t.Fatal("Failed to count records:", err)
	}
	if count != 0 {
		t.Fatalf("Expected record count 0 after deletion, got %d", count)
	}

	t.Log("TestDeleteSourceCodeHandler completed successfully")
}
