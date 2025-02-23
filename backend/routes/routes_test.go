package routes_test

import (
	"AST-Generator/config"
	"AST-Generator/controllers"
	"AST-Generator/db"
	"AST-Generator/models"
	"AST-Generator/routes"
	"AST-Generator/services"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// setupRouter テスト用のGinルーターを生成し、セッションミドルウェアとルート設定を追加
func setupRouter() *gin.Engine {
	router := gin.Default()
	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("testsession", store))
	routes.RegisterRoutes(router)
	return router
}

// initTestDB in-memory SQLiteを使用してテスト用のDBを初期化し、SourceCodeモデルのマイグレーションを実行
func initTestDB(t *testing.T) {
	t.Helper()
	testDB, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatal("Failed to open test database:", err)
	}
	db.DB = testDB
	if err := db.DB.AutoMigrate(&models.SourceCode{}); err != nil {
		t.Fatal("AutoMigrate error:", err)
	}
}

// TestParseEndpoint /parseエンドポイントの動作確認テスト
func TestParseEndpoint(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := setupRouter()

	// ValidPayload: 正常なペイロードの場合のテスト
	t.Run("ValidPayload", func(t *testing.T) {
		payload := map[string]string{
			"code":     "print('Hello')",
			"language": "python",
		}
		body, _ := json.Marshal(payload)
		req, err := http.NewRequest(http.MethodPost, "/parse", bytes.NewBuffer(body))
		if err != nil {
			t.Fatalf("Failed to create request: %v", err)
		}
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)
		if w.Code != http.StatusOK {
			t.Errorf("Expected status 200, got %d", w.Code)
		}
	})

	// InvalidPayload: 無効なペイロード(nil)の場合のテスト
	t.Run("InvalidPayload", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodPost, "/parse", nil)
		if err != nil {
			t.Fatalf("Failed to create request: %v", err)
		}
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)
		if w.Code != http.StatusBadRequest {
			t.Errorf("Expected status 400 for invalid payload, got %d", w.Code)
		}
	})
}

// TestAuthLogoutEndpoint /auth/logoutエンドポイントが200を返すか確認するテスト
func TestAuthLogoutEndpoint(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := setupRouter()

	req, err := http.NewRequest(http.MethodGet, "/auth/logout", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200 for /auth/logout, got %d", w.Code)
	}
}

// TestAuthCurrentUserEndpoint /auth/current_userエンドポイントに未認証でアクセスした際に401が返るか確認するテスト
func TestAuthCurrentUserEndpoint(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := setupRouter()

	req, err := http.NewRequest(http.MethodGet, "/auth/current_user", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	if w.Code != http.StatusUnauthorized {
		t.Errorf("Expected status 401 for unauthenticated /auth/current_user, got %d", w.Code)
	}
}

// TestSourceCodesEndpoints /source_codes関連の各エンドポイントをテスト
func TestSourceCodesEndpoints(t *testing.T) {
	gin.SetMode(gin.TestMode)
	initTestDB(t)

	{
		records := []models.SourceCode{
			{
				Title: "Title 1", Description: "Description 1", Language: "go",
				Code: "package main\nfunc main() { println(\"Hello 1\") }",
			},
			{
				Title: "Title 2", Description: "Description 2", Language: "cpp",
				Code: "#include <iostream>\nint main() { std::cout << \"Hello 2\"; return 0; }",
			},
		}
		for _, rec := range records {
			if err := db.DB.Create(&rec).Error; err != nil {
				t.Fatal("Failed to create test record:", err)
			}
		}

		r := gin.Default()
		r.GET("/source_codes", controllers.GetAllSourceCodesHandler)
		req, err := http.NewRequest(http.MethodGet, "/source_codes", nil)
		if err != nil {
			t.Fatalf("Failed to create request: %v", err)
		}
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		if w.Code != http.StatusOK {
			t.Fatalf("Expected status 200 for /source_codes, got %d", w.Code)
		}
		var fetched []models.SourceCode
		if err := json.Unmarshal(w.Body.Bytes(), &fetched); err != nil {
			t.Fatalf("Error unmarshalling response: %v", err)
		}
		if len(fetched) != len(records) {
			t.Fatalf("Expected %d records, got %d", len(records), len(fetched))
		}
	}

	// GET /source_codes/user のテスト（JWT認証必須）
	{
		db.DB.Exec("DELETE FROM source_codes")
		records := []models.SourceCode{
			{
				Title: "User Title 1", Description: "Description 1", Language: "go",
				Code:   "package main\nfunc main() { println(\"Hello User 1\") }",
				UserID: "test-user-id",
			},
			{
				Title: "User Title 2", Description: "Description 2", Language: "cpp",
				Code:   "#include <iostream>\nint main() { std::cout << \"Hello User 2\"; return 0; }",
				UserID: "test-user-id",
			},
			{
				Title: "Other Title", Description: "Other Description", Language: "python",
				Code:   "print('Hello from other user')",
				UserID: "other-user-id",
			},
		}
		for _, rec := range records {
			if err := db.DB.Create(&rec).Error; err != nil {
				t.Fatal("Failed to create test record:", err)
			}
		}
		jwtToken, err := services.GenerateJWT("test-user-id", "test@example.com", "Test User", "https://example.com/test.jpg", config.JwtSecret)
		if err != nil {
			t.Fatal("Failed to generate JWT:", err)
		}
		r := gin.Default()

		// /source_codes/userはJWT認証が必要
		r.GET("/source_codes/user", controllers.JWTAuthMiddleware(), controllers.GetUserSourceCodesHandler)
		req, err := http.NewRequest(http.MethodGet, "/source_codes/user", nil)
		if err != nil {
			t.Fatalf("Failed to create request: %v", err)
		}
		req.AddCookie(&http.Cookie{Name: "auth_token", Value: jwtToken})
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		// Expect 200 OK
		if w.Code != http.StatusOK {
			t.Fatalf("Expected status 200 for /source_codes/user, got %d", w.Code)
		}
		var fetched []models.SourceCode
		if err := json.Unmarshal(w.Body.Bytes(), &fetched); err != nil {
			t.Fatalf("Error unmarshalling response: %v", err)
		}
		if len(fetched) != 2 {
			t.Fatalf("Expected 2 records for user 'test-user-id', got %d", len(fetched))
		}
		for _, r := range fetched {
			if r.UserID != "test-user-id" {
				t.Fatalf("Expected record UserID 'test-user-id', got %s", r.UserID)
			}
		}
	}

	// POST /source_codes/save のテスト（JWT認証必須）
	{
		// 既存のレコードを削除
		db.DB.Exec("DELETE FROM source_codes")
		jwtToken, err := services.GenerateJWT("test-user-id", "test@example.com", "Test User", "https://example.com/test.jpg", config.JwtSecret)
		if err != nil {
			t.Fatal("Failed to generate JWT:", err)
		}
		r := gin.Default()
		// /source_codes/saveはJWT認証が必要
		r.POST("/source_codes/save", controllers.JWTAuthMiddleware(), controllers.SaveSourceCodeHandler)
		// 新規レコード作成用のペイロードを作成
		payload := map[string]string{
			"language":    "go",
			"code":        "package main\n\nfunc main() { println(\"Hello World\") }",
			"title":       "Hello World Example",
			"description": "This is a sample code that prints Hello World.",
		}
		jsonData, _ := json.Marshal(payload)
		req, err := http.NewRequest(http.MethodPost, "/source_codes/save", bytes.NewBuffer(jsonData))
		if err != nil {
			t.Fatalf("Failed to create request: %v", err)
		}
		req.Header.Set("Content-Type", "application/json")
		req.AddCookie(&http.Cookie{Name: "auth_token", Value: jwtToken})
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		// Expect 200 OK
		if w.Code != http.StatusOK {
			t.Fatalf("Expected status 200 for /source_codes/save, got %d", w.Code)
		}
		var resp map[string]interface{}
		if err := json.Unmarshal(w.Body.Bytes(), &resp); err != nil {
			t.Fatalf("Error unmarshalling response: %v", err)
		}
		if msg, ok := resp["message"].(string); !ok || msg != "Source code saved successfully" {
			t.Fatalf("Expected message 'Source code saved successfully', got %v", resp["message"])
		}
		var recs []models.SourceCode
		if err := db.DB.Find(&recs).Error; err != nil {
			t.Fatal("Failed to query saved records:", err)
		}
		if len(recs) != 1 {
			t.Fatalf("Expected 1 record, got %d", len(recs))
		}
		if recs[0].UserID != "test-user-id" {
			t.Fatalf("Expected record UserID 'test-user-id', got %v", recs[0].UserID)
		}
	}

	// DELETE /source_codes/:id のテスト（JWT認証必須）
	{
		// 既存のレコードを削除して、新規テストレコードを作成
		db.DB.Exec("DELETE FROM source_codes")
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
			t.Fatal("Failed to generate JWT:", err)
		}
		r := gin.Default()

		// /source_codes/:idはJWT認証が必要
		r.DELETE("/source_codes/:id", controllers.JWTAuthMiddleware(), controllers.DeleteSourceCodeHandler)
		url := "/source_codes/" + strconv.Itoa(int(source.ID))
		req, err := http.NewRequest(http.MethodDelete, url, nil)
		if err != nil {
			t.Fatalf("Failed to create request: %v", err)
		}
		req.AddCookie(&http.Cookie{Name: "auth_token", Value: jwtToken})
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		if w.Code != http.StatusOK {
			t.Fatalf("Expected status 200 for DELETE /source_codes/:id, got %d", w.Code)
		}
		var resp map[string]interface{}
		if err := json.Unmarshal(w.Body.Bytes(), &resp); err != nil {
			t.Fatalf("Error unmarshalling response: %v", err)
		}
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
	}
}
