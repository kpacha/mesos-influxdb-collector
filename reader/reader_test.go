package reader

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestReadUrl(t *testing.T) {
	response := "Hello, client"
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, response)
	}))
	defer ts.Close()

	body, err := ReadUrl(ts.URL)
	if err != nil {
		t.Fatal("Error reading url:", err.Error())
	}

	if !strings.Contains(string(body), response) {
		t.Fatal("Unexpected response body", string(body))
	}
}
