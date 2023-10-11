package fetcher

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestFetch(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.Write([]byte("Test content"))
	}))
	defer server.Close()

	res := Fetch(server.URL)
	if res.Error != nil {
		t.Fatalf("Expected no error, but got %v", res.Error)
	}

	if res.Code != 200 {
		t.Fatalf("Expected status code 200, but got %d", res.Code)
	}

	if res.Size != int64(len("Test content")) {
		t.Fatalf("Expected content size %d, but got %d", len("Test content"), res.Size)
	}
}
