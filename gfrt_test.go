package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRootHandler(t *testing.T) {
	patterns := []struct {
		method             string
		path               string
		expectResponseCode int
	}{
		{
			method:             http.MethodGet,
			path:               "/",
			expectResponseCode: http.StatusOK,
		},
		{
			method:             http.MethodPut,
			path:               "/",
			expectResponseCode: http.StatusNoContent,
		},
		{
			method:             http.MethodGet,
			path:               "/",
			expectResponseCode: http.StatusMovedPermanently,
		},
		{
			method:             http.MethodDelete,
			path:               "/",
			expectResponseCode: http.StatusNoContent,
		},
		{
			method:             http.MethodGet,
			path:               "/",
			expectResponseCode: http.StatusOK,
		},
		{
			method:             http.MethodPost,
			path:               "/",
			expectResponseCode: http.StatusMethodNotAllowed,
		},
	}

	for i, pattern := range patterns {
		rec := httptest.NewRecorder()
		req := httptest.NewRequest(pattern.method, pattern.path, nil)
		rootHandler(rec, req)
		switch {
		case rec.Code != pattern.expectResponseCode:
			t.Errorf("mismatch response code(%d): %d != %d", i, rec.Code, pattern.expectResponseCode)
		}
	}
}
