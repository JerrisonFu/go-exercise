package network

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
)

func CreateServer() *httptest.Server {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello, World!")
	})
	mux.HandleFunc("/json", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{"message":"Hello"}`)
	})
	mux.HandleFunc("/echo", func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		w.Write(body)
	})
	return httptest.NewServer(mux)
}

func HttpClientExample(baseURL string) (string, error) {
	resp, err := http.Get(baseURL + "/hello")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

func HttpPostExample(baseURL string) (string, error) {
	resp, err := http.Post(baseURL+"/echo", "text/plain", strings.NewReader("Hello POST"))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

func HttpClientWithHeader(baseURL string) int {
	req, _ := http.NewRequest("GET", baseURL+"/hello", nil)
	req.Header.Set("X-Custom-Header", "value")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return 0
	}
	defer resp.Body.Close()

	return resp.StatusCode
}
