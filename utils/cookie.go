package utils

import (
	"net/http"
	"time"
)

func CreateCookie(w http.ResponseWriter, sessionToken string) {
	http.SetCookie(w, &http.Cookie{
		Name:    "session_token",
		Value:   sessionToken,
		Expires: time.Now().Add(2 * time.Hour),
	})
}

func DeleteCookie(w http.ResponseWriter, r *http.Request) {
	cookie := &http.Cookie{
		Name:    "session_token",
		Value:   "",
		Path:    "/",
		Expires: time.Unix(1, 0), // Set expiration time to a past value
	}

	http.SetCookie(w, cookie)
}
