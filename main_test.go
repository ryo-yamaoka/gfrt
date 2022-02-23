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
			path:               "/feed",
			expectResponseCode: http.StatusOK,
		},
		{
			method:             http.MethodPut,
			path:               "/feed",
			expectResponseCode: http.StatusNoContent,
		},
		{
			method:             http.MethodGet,
			path:               "/feed",
			expectResponseCode: http.StatusMovedPermanently,
		},
		{
			method:             http.MethodDelete,
			path:               "/feed",
			expectResponseCode: http.StatusNoContent,
		},
		{
			method:             http.MethodGet,
			path:               "/feed",
			expectResponseCode: http.StatusOK,
		},
		{
			method:             http.MethodPost,
			path:               "/feed",
			expectResponseCode: http.StatusMethodNotAllowed,
		},
	}

	for i, pattern := range patterns {
		rec := httptest.NewRecorder()
		req := httptest.NewRequest(pattern.method, pattern.path, nil)
		feedHandler(rec, req)
		switch {
		case rec.Code != pattern.expectResponseCode:
			t.Errorf("mismatch response code(%d): %d != %d", i, rec.Code, pattern.expectResponseCode)
		}
	}
}

func TestExampleArticle1Handler(t *testing.T) {
	patterns := []struct {
		method             string
		path               string
		expectResponseCode int
	}{
		{
			method:             http.MethodGet,
			path:               "/example1",
			expectResponseCode: http.StatusOK,
		},
		{
			method:             http.MethodPut,
			path:               "/example1",
			expectResponseCode: http.StatusMethodNotAllowed,
		},
		{
			method:             http.MethodDelete,
			path:               "/example1",
			expectResponseCode: http.StatusMethodNotAllowed,
		},
		{
			method:             http.MethodPost,
			path:               "/example1",
			expectResponseCode: http.StatusMethodNotAllowed,
		},
		{
			method:             http.MethodPost,
			path:               "/example1",
			expectResponseCode: http.StatusMethodNotAllowed,
		},
	}

	for i, pattern := range patterns {
		rec := httptest.NewRecorder()
		req := httptest.NewRequest(pattern.method, pattern.path, nil)
		exampleArticle1Handler(rec, req)
		if rec.Code != pattern.expectResponseCode {
			t.Errorf("mismatch response code(%d): %d != %d", i, rec.Code, pattern.expectResponseCode)
		}
	}
}
