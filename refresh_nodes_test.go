package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"

	"testing"
)

const refreshNodesURLPattern = "http://localhost:1234/refresh_nodes%s"

func defaultrefreshNodesURL() string {
	return fmt.Sprintf(refreshNodesURLPattern, "")
}

func TestRefreshNodesHandlerError(t *testing.T) {
	req := httptest.NewRequest("GET", defaultURL(), nil)
	w := httptest.NewRecorder()
	http.HandlerFunc(RefreshNodesHandler).ServeHTTP(w, req)
	resp := w.Result()
	if resp.StatusCode != http.StatusBadRequest {
		t.Errorf("Empty GET didn't return HTTP 400")
	}

	req = httptest.NewRequest("POST", defaultURL(), nil)
	w = httptest.NewRecorder()
	http.HandlerFunc(RefreshNodesHandler).ServeHTTP(w, req)
	resp = w.Result()
	if resp.StatusCode != http.StatusMethodNotAllowed {
		t.Errorf("POST didn't return bad method status")
	}

}
