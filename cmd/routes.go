package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	_ "github.com/jserva90/Erply-test-task/docs"
	httpSwagger "github.com/swaggo/http-swagger"
)

func (app *application) routes() http.Handler {
	mux := chi.NewRouter()

	corsOptions := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // Adjust the allowed origins accordingly
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Content-Type"},
		AllowCredentials: true,
	})

	mux.Use(corsOptions.Handler)

	mux.Get("/", app.LoginHandler)
	mux.Post("/", app.LoginHandler)
	mux.Post("/verifyUser", app.verifyUserSwagger)
	mux.Post("/getSessionKeyInfo", app.getSessionKeyInfoSwagger)
	mux.Post("/getCustomers", app.getCustomersSwagger)
	mux.Post("/getCustomerByID", app.getCustomerByIDSwagger)
	mux.Post("/saveCustomer", app.saveCustomerSwagger)

	mux.Route("/admin", func(mux chi.Router) {
		mux.Use(app.authRequired)
		mux.Get("/main", app.MainPageHandler)
		mux.Get("/success", app.SuccessHandler)
		mux.Post("/logout", app.LogoutHandler)
		mux.Post("/getcustomers", app.FetchCustomersHandler)
		mux.Get("/getcustomer", app.FetchCustomerHandler)
		mux.Post("/getcustomer", app.FetchCustomerHandler)
		mux.Get("/savecustomer", app.SaveCustomerHandler)
		mux.Post("/savecustomer", app.SaveCustomerHandler)
	})

	mux.Mount("/swagger", httpSwagger.WrapHandler)

	return mux
}
