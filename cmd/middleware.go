package main

import "net/http"

func (app *application) authRequired(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("session_token")
		if err != nil || cookie.Value == "" {
			// Redirect to the Home page if the cookie doesn't exist or has no value
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}

		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}
