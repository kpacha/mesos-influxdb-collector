package reader

import (
	"fmt"
	"log"
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
		log.Println("Error reading url:", err.Error())
		t.Fail()
	}

	if !strings.Contains(string(body), response) {
		log.Println("Unexpected response body", string(body))
		t.Fail()
	}
}
