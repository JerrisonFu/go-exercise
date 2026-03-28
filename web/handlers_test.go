package web

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func init() {
	ResetUsers()
}

func TestPing(t *testing.T) {
	r := SetupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Status = %d, want %d", w.Code, http.StatusOK)
	}

	var resp map[string]string
	json.Unmarshal(w.Body.Bytes(), &resp)
	if resp["message"] != "pong" {
		t.Errorf("message = %s, want 'pong'", resp["message"])
	}
}

func TestHello(t *testing.T) {
	r := SetupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/hello?name=Gin", nil)
	r.ServeHTTP(w, req)

	var resp map[string]string
	json.Unmarshal(w.Body.Bytes(), &resp)
	if resp["message"] != "Hello Gin" {
		t.Errorf("message = %s, want 'Hello Gin'", resp["message"])
	}
}

func TestGetUsers(t *testing.T) {
	r := SetupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/users", nil)
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Status = %d, want %d", w.Code, http.StatusOK)
	}

	var users []*User
	json.Unmarshal(w.Body.Bytes(), &users)
	if len(users) != 2 {
		t.Errorf("len(users) = %d, want 2", len(users))
	}
}

func TestGetUserByID(t *testing.T) {
	r := SetupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/users/1", nil)
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Status = %d, want %d", w.Code, http.StatusOK)
	}

	var user User
	json.Unmarshal(w.Body.Bytes(), &user)
	if user.Name != "Alice" {
		t.Errorf("user.Name = %s, want 'Alice'", user.Name)
	}
}

func TestGetUserNotFound(t *testing.T) {
	r := SetupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/users/999", nil)
	r.ServeHTTP(w, req)

	if w.Code != http.StatusNotFound {
		t.Errorf("Status = %d, want %d", w.Code, http.StatusNotFound)
	}
}

func TestCreateUser(t *testing.T) {
	ResetUsers()
	r := SetupRouter()

	body := `{"name":"Charlie","age":35}`
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/users", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	if w.Code != http.StatusCreated {
		t.Errorf("Status = %d, want %d", w.Code, http.StatusCreated)
	}

	var user User
	json.Unmarshal(w.Body.Bytes(), &user)
	if user.Name != "Charlie" {
		t.Errorf("user.Name = %s, want 'Charlie'", user.Name)
	}
}

func TestCreateUserValidation(t *testing.T) {
	ResetUsers()
	r := SetupRouter()

	body := `{"name":""}`
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/users", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("Status = %d, want %d", w.Code, http.StatusBadRequest)
	}
}

func TestUpdateUser(t *testing.T) {
	ResetUsers()
	r := SetupRouter()

	body := `{"name":"Alice Updated","age":31}`
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", "/users/1", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Status = %d, want %d", w.Code, http.StatusOK)
	}

	var user User
	json.Unmarshal(w.Body.Bytes(), &user)
	if user.Name != "Alice Updated" {
		t.Errorf("user.Name = %s, want 'Alice Updated'", user.Name)
	}
}

func TestDeleteUser(t *testing.T) {
	ResetUsers()
	r := SetupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/users/2", nil)
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Status = %d, want %d", w.Code, http.StatusOK)
	}

	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/users/2", nil)
	r.ServeHTTP(w, req)

	if w.Code != http.StatusNotFound {
		t.Errorf("After delete, Status = %d, want %d", w.Code, http.StatusNotFound)
	}
}
