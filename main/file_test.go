package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	f "ascii-art-web/func"
)

func TestWelcom(t *testing.T) {
	tests := []struct {
		name         string
		url          string
		method       string
		expectedCode int
	}{
		{
			name:         "valid request",
			url:          "/",
			method:       http.MethodGet,
			expectedCode: http.StatusOK,
		},
		{
			name:         "Invalid URL Path",
			url:          "/invalid",
			method:       http.MethodGet,
			expectedCode: http.StatusNotFound,
		},
		{
			name:         "Invalid method",
			url:          "/",
			method:       http.MethodPost,
			expectedCode: http.StatusMethodNotAllowed,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// creat the request , / GET , "curl -X GET http://localhost:8080/"
			req, err := http.NewRequest(tt.method, tt.url, nil)
			if err != nil {
				t.Fatal(err)
			}
			//[jlsjdflkdjflsdkfjskldfj]
			rr := httptest.NewRecorder()
			// handle -> f.Welcom
			handler := http.HandlerFunc(f.Welcom)
			// send the request and get the response
			// req *http.Re
			// w , responserecorder
			handler.ServeHTTP(rr, req)

			if status := rr.Code; status != tt.expectedCode {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, tt.expectedCode)
			}
		})
	}
}

func TestLast(t *testing.T) {
	tests := []struct {
		name   string
		url    string
		method string
		body   string // used for valid request
		status int
	}{
		{
			name:   "Valid POST request",
			url:    "/ascii-art",
			method: http.MethodPost,
			body:   "ljomla=dgfsadf&banner=standard",
			status: http.StatusOK,
		},

		{
			name:   "Invalid URL path",
			url:    "/ascii-art/invalid",
			method: http.MethodPut,
			status: http.StatusNotFound,
		},
		{
			name:   "Invalid method",
			url:    "/ascii-art",
			method: http.MethodGet,
			status: http.StatusMethodNotAllowed,
		},
		{
			name:   "Bad request",
			url:    "/ascii-art",
			method: http.MethodPost,
			body:   "invalid_body_format",
			status: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var req *http.Request
			// this is for the case of good request , i need to set up the body correctly
			if tt.status == http.StatusOK {
				req = httptest.NewRequest(tt.method, tt.url, strings.NewReader(tt.body))
				req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
			} else {
				req = httptest.NewRequest(tt.method, tt.url, nil)
			}

			rr := httptest.NewRecorder()

			handler := http.HandlerFunc(f.Last)
			handler.ServeHTTP(rr, req)

			if status := rr.Code; status != tt.status {
				t.Errorf("handler returned wrong status code: got %v want %v", status, tt.status)
			}
		})
	}
}
