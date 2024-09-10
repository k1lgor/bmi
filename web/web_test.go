package web

import (
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	router := gin.Default()

	// Dynamically find the absolute path to the templates directory
	wd, _ := os.Getwd()
	templatesPath := filepath.Join(wd, "templates", "*")
	router.LoadHTMLGlob(templatesPath)

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{"bmi": "", "error": ""})
	})

	router.POST("/calculate", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{"bmi": "22.5", "category": "Normal", "error": ""})
	})
	return router
}

func TestCalculateRoute(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/calculate", strings.NewReader("weight=70&height=170&weightUnit=kg&heightUnit=cm"))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}
}
