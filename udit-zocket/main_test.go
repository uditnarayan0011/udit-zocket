package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestGetProductByID(t *testing.T) {
	// Setup router and recorder
	r := gin.Default()
	r.GET("/products/:id", getProductByID)

	req, _ := http.NewRequest("GET", "/products/1", nil)
	w := httptest.NewRecorder()

	// Perform the request
	r.ServeHTTP(w, req)

	// Assert status code and response body
	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d but got %d", http.StatusOK, w.Code)
	}

	// Additional checks on the response body can be added
}
