package api

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHealthCheck(t *testing.T) {
	r := SetupAPIRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/health", nil)
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Status = %d, want %d", w.Code, http.StatusOK)
	}

	var resp Response
	json.Unmarshal(w.Body.Bytes(), &resp)
	if !resp.Success {
		t.Error("Health check should succeed")
	}
}

func TestListProductsUnauthorized(t *testing.T) {
	r := SetupAPIRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/products", nil)
	r.ServeHTTP(w, req)

	if w.Code != http.StatusUnauthorized {
		t.Errorf("Status = %d, want %d", w.Code, http.StatusUnauthorized)
	}
}

func TestListProducts(t *testing.T) {
	r := SetupAPIRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/products", nil)
	req.Header.Set("Authorization", "Bearer valid-token")
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Status = %d, want %d", w.Code, http.StatusOK)
	}

	var resp PaginatedResponse
	json.Unmarshal(w.Body.Bytes(), &resp)
	if !resp.Success {
		t.Error("List products should succeed")
	}
	if resp.TotalCount != 2 {
		t.Errorf("TotalCount = %d, want 2", resp.TotalCount)
	}
}

func TestGetProduct(t *testing.T) {
	r := SetupAPIRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/products/1", nil)
	req.Header.Set("Authorization", "Bearer valid-token")
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Status = %d, want %d", w.Code, http.StatusOK)
	}

	var resp Response
	json.Unmarshal(w.Body.Bytes(), &resp)
	if !resp.Success {
		t.Error("Get product should succeed")
	}
}

func TestGetProductNotFound(t *testing.T) {
	r := SetupAPIRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/products/999", nil)
	req.Header.Set("Authorization", "Bearer valid-token")
	r.ServeHTTP(w, req)

	if w.Code != http.StatusNotFound {
		t.Errorf("Status = %d, want %d", w.Code, http.StatusNotFound)
	}
}

func TestCreateProduct(t *testing.T) {
	r := SetupAPIRouter()

	body := `{"name":"Keyboard","description":"Mechanical keyboard","price":79.99,"stock":20}`
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/v1/products", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer valid-token")
	r.ServeHTTP(w, req)

	if w.Code != http.StatusCreated {
		t.Errorf("Status = %d, want %d", w.Code, http.StatusCreated)
	}

	var resp Response
	json.Unmarshal(w.Body.Bytes(), &resp)
	if !resp.Success {
		t.Error("Create product should succeed")
	}
}

func TestUpdateProduct(t *testing.T) {
	r := SetupAPIRouter()

	body := `{"name":"Laptop Pro","description":"Updated laptop","price":1299.99,"stock":5}`
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", "/api/v1/products/1", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer valid-token")
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Status = %d, want %d", w.Code, http.StatusOK)
	}
}

func TestDeleteProduct(t *testing.T) {
	r := SetupAPIRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/api/v1/products/2", nil)
	req.Header.Set("Authorization", "Bearer valid-token")
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Status = %d, want %d", w.Code, http.StatusOK)
	}
}

func TestBuyProduct(t *testing.T) {
	r := SetupAPIRouter()

	body := `{"quantity":2}`
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/v1/products/1/buy", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer valid-token")
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Status = %d, want %d", w.Code, http.StatusOK)
	}

	var resp Response
	json.Unmarshal(w.Body.Bytes(), &resp)
	if !resp.Success {
		t.Error("Buy product should succeed")
	}
}

func TestBuyProductInsufficientStock(t *testing.T) {
	r := SetupAPIRouter()

	body := `{"quantity":1000}`
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/v1/products/1/buy", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer valid-token")
	r.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("Status = %d, want %d", w.Code, http.StatusBadRequest)
	}
}

func TestPagination(t *testing.T) {
	r := SetupAPIRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/products?page=1&page_size=1", nil)
	req.Header.Set("Authorization", "Bearer valid-token")
	r.ServeHTTP(w, req)

	var resp PaginatedResponse
	json.Unmarshal(w.Body.Bytes(), &resp)
	if resp.PageSize != 1 {
		t.Errorf("PageSize = %d, want 1", resp.PageSize)
	}
	if resp.TotalPages < 1 {
		t.Errorf("TotalPages = %d, want >= 1", resp.TotalPages)
	}
}

func TestCORSHeaders(t *testing.T) {
	r := SetupAPIRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("OPTIONS", "/api/v1/products", nil)
	r.ServeHTTP(w, req)

	if w.Header().Get("Access-Control-Allow-Origin") != "*" {
		t.Error("CORS header should be set")
	}
}
