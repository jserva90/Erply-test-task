package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (app *application) routes() http.Handler {
	mux := chi.NewRouter()

	mux.Get("/", app.Home)
	mux.Post("/", app.Home)

	mux.Route("/admin", func(mux chi.Router) {
		mux.Use(app.authRequired)
		mux.Get("/main", app.MainPage)
		mux.Get("/success", app.Success)
		mux.Post("/logout", app.Logout)
		mux.Post("/getcustomers", app.FetchCustomers)
		mux.Get("/savecustomer", app.SaveCustomer)
		mux.Post("/savecustomer", app.SaveCustomer)
	})

	return mux
}
