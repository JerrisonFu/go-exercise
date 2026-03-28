package network

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHTTPBasics(t *testing.T) {
	server := CreateServer()
	defer server.Close()

	body, err := HttpClientExample(server.URL)
	if err != nil {
		t.Errorf("HttpClientExample() error = %v", err)
	}
	if body != "Hello, World!" {
		t.Errorf("HttpClientExample() = %s, want 'Hello, World!'", body)
	}
}

func TestHTTPPost(t *testing.T) {
	server := CreateServer()
	defer server.Close()

	body, err := HttpPostExample(server.URL)
	if err != nil {
		t.Errorf("HttpPostExample() error = %v", err)
	}
	if body != "Hello POST" {
		t.Errorf("HttpPostExample() = %s, want 'Hello POST'", body)
	}
}

func TestHTTPHeaders(t *testing.T) {
	server := CreateServer()
	defer server.Close()

	status := HttpClientWithHeader(server.URL)
	if status != 200 {
		t.Errorf("HttpClientWithHeader() = %d, want 200", status)
	}
}

func TestJSONEndpoint(t *testing.T) {
	server := CreateServer()
	defer server.Close()

	resp, err := server.Client().Get(server.URL + "/json")
	if err != nil {
		t.Errorf("GET /json error = %v", err)
	}
	defer resp.Body.Close()

	if resp.Header.Get("Content-Type") != "application/json" {
		t.Errorf("Content-Type = %s, want application/json", resp.Header.Get("Content-Type"))
	}
}

func TestHTTPStatusCodes(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/ok", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	mux.HandleFunc("/err", func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Error", http.StatusInternalServerError)
	})

	server := httptest.NewServer(mux)
	defer server.Close()

	resp, err := http.Get(server.URL + "/ok")
	if err != nil {
		t.Errorf("GET /ok error = %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Status = %d, want %d", resp.StatusCode, http.StatusOK)
	}
	resp.Body.Close()

	resp, _ = http.Get(server.URL + "/err")
	if resp.StatusCode != http.StatusInternalServerError {
		t.Errorf("Status = %d, want %d", resp.StatusCode, http.StatusInternalServerError)
	}
	resp.Body.Close()
}

func TestNetworkInterfaces(t *testing.T) {
	ifaces := NetworkInterfaces()
	if len(ifaces) == 0 {
		t.Error("NetworkInterfaces() should return at least one interface")
	}
}

func TestResolveHost(t *testing.T) {
	host := ResolveHost()
	if host == "" {
		t.Error("ResolveHost() should return non-empty string")
	}
}

func TestEncodeDecodeMessage(t *testing.T) {
	msg := Message{Type: 1, Length: 5, Payload: []byte("hello")}
	encoded := EncodeMessage(msg)
	decoded := DecodeMessage(encoded)

	if decoded.Type != msg.Type {
		t.Errorf("Type = %d, want %d", decoded.Type, msg.Type)
	}
	if decoded.Length != msg.Length {
		t.Errorf("Length = %d, want %d", decoded.Length, msg.Length)
	}
	if string(decoded.Payload) != string(msg.Payload) {
		t.Errorf("Payload = %s, want %s", decoded.Payload, msg.Payload)
	}
}
