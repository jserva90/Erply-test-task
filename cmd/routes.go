package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	_ "github.com/jserva90/Erply-test-task/docs"
	httpSwagger "github.com/swaggo/http-swagger"
)

func (app *application) routes() http.Handler {
	r := chi.NewRouter()

	corsOptions := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Content-Type"},
		AllowCredentials: true,
	})

	r.Use(corsOptions.Handler)

	r.Route("/", func(r chi.Router) {
		r.Get("/", app.loginHandler)
		r.Post("/", app.loginHandler)
		r.Post("/verifyUser", app.verifyUserSwagger)
		r.Post("/getSessionKeyInfo", app.getSessionKeyInfoSwagger)
		r.Post("/getCustomers", app.getCustomersSwagger)
		r.Post("/getCustomerByID", app.getCustomerByIDSwagger)
		r.Post("/saveCustomer", app.saveCustomerSwagger)
	})

	r.Route("/admin", func(r chi.Router) {
		r.Use(app.authRequired)
		r.Get("/main", app.mainPageHandler)
		r.Post("/logout", app.logoutHandler)
		r.Post("/getcustomers", app.fetchCustomersHandler)
		r.Get("/getcustomer", app.fetchCustomerHandler)
		r.Post("/getcustomer", app.fetchCustomerHandler)
		r.Get("/savecustomer", app.saveCustomerHandler)
		r.Post("/savecustomer", app.saveCustomerHandler)
	})

	r.Mount("/swagger", httpSwagger.WrapHandler)

	return r
}
