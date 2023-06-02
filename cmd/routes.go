package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	_ "github.com/jserva90/Erply-test-task/docs"
	httpSwagger "github.com/swaggo/http-swagger"
)

func (app *application) routes() http.Handler {
	mux := chi.NewRouter()

	mux.Get("/", app.Login)
	mux.Post("/", app.Login)
	mux.Post("/verifyUser", app.verifyUserSwagger)
	mux.Post("/getSessionKeyInfo", app.getSessionKeyInfoSwagger)
	mux.Post("/getCustomers", app.getCustomersSwagger)
	mux.Post("/saveCustomer", app.saveCustomerSwagger)

	mux.Route("/admin", func(mux chi.Router) {
		mux.Use(app.authRequired)
		mux.Get("/main", app.MainPage)
		mux.Get("/success", app.Success)
		mux.Post("/logout", app.Logout)
		mux.Post("/getcustomers", app.FetchCustomers)
		mux.Get("/getcustomer", app.FetchCustomer)
		mux.Post("/getcustomer", app.FetchCustomer)
		mux.Get("/savecustomer", app.SaveCustomer)
		mux.Post("/savecustomer", app.SaveCustomer)
	})

	mux.Mount("/swagger", httpSwagger.WrapHandler)

	return mux
}
