package utils

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestCreateCookie(t *testing.T) {
	recorder := httptest.NewRecorder()

	sessionToken := "abcd1234"
	CreateCookie(recorder, sessionToken)

	cookies := recorder.Result().Cookies()

	found := false
	for _, cookie := range cookies {
		if cookie.Name == "session_token" {
			if cookie.Value != sessionToken {
				t.Errorf("Expected cookie value: %s, got: %s", sessionToken, cookie.Value)
			}

			expectedExpiry := time.Now().Add(2 * time.Hour).Truncate(time.Second)
			cookieExpiry := cookie.Expires.Truncate(time.Second)
			if !cookieExpiry.Equal(expectedExpiry) {
				t.Errorf("Expected cookie expiry: %s, got: %s", expectedExpiry, cookieExpiry)
			}

			found = true
			break
		}
	}

	if !found {
		t.Error("Expected session_token cookie not found")
	}
}

func TestDeleteCookie(t *testing.T) {
	recorder := httptest.NewRecorder()

	request := httptest.NewRequest(http.MethodGet, "/", nil)

	DeleteCookie(recorder, request)

	cookies := recorder.Result().Cookies()

	found := false
	for _, cookie := range cookies {
		if cookie.Name == "session_token" {
			if cookie.Value != "" {
				t.Error("Expected empty cookie value, but got:", cookie.Value)
			}

			if cookie.Path != "/" {
				t.Error("Expected cookie path: '/', but got:", cookie.Path)
			}

			expectedExpiry := time.Unix(1, 0)
			if !cookie.Expires.Equal(expectedExpiry) {
				t.Errorf("Expected cookie expiry: %s, got: %s", expectedExpiry, cookie.Expires)
			}

			found = true
			break
		}
	}

	if !found {
		t.Error("Expected session_token cookie not found")
	}
}
