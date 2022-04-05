package auth

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"testing"
)

func TestAuth(t *testing.T) {
	testHandlerSuccessfullRequest := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	})
	// token and secret key are for testing
	os.Setenv("AUTH_SECRET_KEY", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9")
	validToken := "eyJhbGciOiJIUzI1NiJ9.eyJSb2xlIjoiVGVzdGVyIiwiSXNzdWVyIjoiSXNzdWVyIiwiVXNlcm5hbWUiOiJZYWh5YVFhbmRlbCIsImlhdCI6MTY0OTEyMzc2MX0.-UKsQ3rVQ8Eukj7xEGGoIqZBtLXL1jtEBRKQuVT1VAA"
	tests := []struct {
		name             string
		params           string
		header           http.Header
		wantStatus       int
		wantResponseText string
		handler          http.Handler
	}{
		{
			name:             "no authorization provided",
			header:           http.Header{},
			wantStatus:       http.StatusUnauthorized,
			wantResponseText: "Authentication: token required",
			handler:          http.HandlerFunc((func(w http.ResponseWriter, r *http.Request) {})),
		},
		{
			name:             "not bearer token",
			header:           http.Header{"Authorization": []string{"Token Token"}},
			wantStatus:       http.StatusUnauthorized,
			wantResponseText: "Authentication: Bearer token required",
			handler:          http.HandlerFunc((func(w http.ResponseWriter, r *http.Request) {})),
		},
		{
			name:             "invalid jwt token",
			header:           http.Header{"Authorization": []string{"Bearer xx"}},
			wantStatus:       http.StatusUnauthorized,
			wantResponseText: "Error verifying JWT token: token contains an invalid number of segments",
			handler:          http.HandlerFunc((func(w http.ResponseWriter, r *http.Request) {})),
		},
		{
			name:             "successful authorization",
			header:           http.Header{"Authorization": []string{fmt.Sprintf("%s %s", "Bearer", validToken)}},
			wantStatus:       http.StatusAccepted,
			wantResponseText: "",
			handler:          http.HandlerFunc((func(w http.ResponseWriter, r *http.Request) {})),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, _ := http.NewRequest("", "", nil)
			req.Header = tt.header
			rr := httptest.NewRecorder()
			handler := Auth(testHandlerSuccessfullRequest)
			handler.ServeHTTP(rr, req)
			got := rr.Body.String()
			gotStatus := rr.Code
			if !reflect.DeepEqual(got, tt.wantResponseText) {
				t.Errorf("Auth() = %v, want %v", got, tt.wantResponseText)
			}
			if !reflect.DeepEqual(gotStatus, tt.wantStatus) {
				t.Errorf("Auth() = %v, want %v", gotStatus, tt.wantStatus)
			}
		})
	}
}

func TestAuthFailsWithJwtSecretKeyNotSet(t *testing.T) {
	testHandlerSuccessfullRequest := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	})
	os.Setenv("AUTH_SECRET_KEY", "")
	req, _ := http.NewRequest("", fmt.Sprintf(`/%s`, "?tenant_id=102"), nil)
	req.Header = http.Header{"Authorization": []string{"Bearer xx"}}
	rr := httptest.NewRecorder()
	handler := Auth(testHandlerSuccessfullRequest)
	handler.ServeHTTP(rr, req)
	got := rr.Body.String()
	gotStatus := rr.Code
	if !reflect.DeepEqual(got, "secret key not set") {
		t.Errorf("Auth() = %v, want %v", got, "xx")
	}
	if !reflect.DeepEqual(gotStatus, http.StatusUnauthorized) {
		t.Errorf("Auth() = %v, want %v", gotStatus, http.StatusUnauthorized)
	}
}
