package main

import (
	"encoding/json"
	"html/template"
	"net/http"

	"github.com/gofrs/uuid"
	"github.com/jserva90/Erply-test-task/models"
	"github.com/jserva90/Erply-test-task/utils"
)

const (
	mainPath           = "/admin/main"
	fetchCustomerPath  = "/admin/getcustomer"
	fetchCustomersPath = "/admin/getcustomers"
	saveCustomerPath   = "/admin/savecustomer"

	internalServerError  = "Internal Server Error"
	loginErrorMessage    = "Invalid credentials"
	logoutErrorMessage   = "Failed to logout"
	methodNotAllowedMsg  = "Method Not Allowed"
	notFoundMessage      = "Not Found"
	removeSessionMessage = "Failed to remove session"
	sessionExpiredError  = "Session expired"
	unauthorizedMessage  = "Unauthorized"
)

func (app *application) mainPageHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != mainPath {
		handleError(w, notFoundMessage, http.StatusNotFound)
		return
	}

	switch r.Method {
	case "GET":
		renderTemplate(w, "templates/main.html", nil)
	default:
		handleError(w, methodNotAllowedMsg, http.StatusMethodNotAllowed)
	}
}

func (app *application) loginHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session_token")
	if err == nil {
		session, err := app.DB.GetSession(cookie.Value)
		if err != nil {
			goto ContinueExecution
		}
		decryptedClientCode, err := utils.Decrypt(session.ClientCode, utils.GetSecretKey())
		if err != nil {
			goto ContinueExecution
		}
		decryptedSessionKey, err := utils.Decrypt(session.SessionKey, utils.GetSecretKey())
		if err != nil {
			goto ContinueExecution
		}

		sessionInfo, err := app.getSessionKeyInfo(decryptedClientCode, decryptedSessionKey)
		if err != nil {
			goto ContinueExecution
		}

		if utils.IsSessionExpired(*sessionInfo) {
			goto ContinueExecution
		}

		http.Redirect(w, r, "/admin/main", http.StatusSeeOther)
	}

ContinueExecution:
	switch r.Method {
	case http.MethodGet:
		renderTemplate(w, "templates/login.html", nil)
	case http.MethodPost:
		clientCode := r.FormValue("clientCode")
		username := r.FormValue("username")
		password := r.FormValue("password")

		res, err := app.verifyUser(clientCode, username, password)
		if err != nil {
			handleError(w, loginErrorMessage, http.StatusNotFound)
			return
		}

		uuid := uuid.Must(uuid.NewV4()).String()

		err = app.DB.AddSession(clientCode, username, password, res.Records[0].SessionKey, uuid)
		if err != nil {
			handleError(w, notFoundMessage, http.StatusNotFound)
			return
		}

		utils.CreateCookie(w, uuid)

		http.Redirect(w, r, "/admin/main", http.StatusSeeOther)
	default:
		handleError(w, methodNotAllowedMsg, http.StatusMethodNotAllowed)
	}
}

func (app *application) logoutHandler(w http.ResponseWriter, r *http.Request) {
	currentCookie, err := r.Cookie("session_token")
	if err != nil {
		handleError(w, logoutErrorMessage, http.StatusInternalServerError)
		return
	}

	currentSessionToken := currentCookie.Value
	err = app.DB.RemoveSession(currentSessionToken)
	if err != nil {
		handleError(w, removeSessionMessage, http.StatusInternalServerError)
		return
	}

	utils.DeleteCookie(w, r)

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (app *application) fetchCustomersHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != fetchCustomersPath {
		handleError(w, notFoundMessage, http.StatusNotFound)
		return
	}

	cookie, err := r.Cookie("session_token")
	if err != nil {
		handleError(w, err.Error(), http.StatusUnauthorized)
		return
	}

	decryptedClientCode, decryptedSessionKey, err := app.getClientCodeAndSessionKey(cookie.Value)
	if err != nil {
		handleError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	customers, err := app.getCustomers(decryptedClientCode, decryptedSessionKey)
	if err != nil {
		handleError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := struct {
		Customers models.CustomerResponse
	}{
		Customers: *customers,
	}

	renderTemplate(w, "templates/customers.html", data)
}

func (app *application) fetchCustomerHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != fetchCustomerPath {
		handleError(w, notFoundMessage, http.StatusNotFound)
		return
	}

	switch r.Method {
	case http.MethodGet:
		renderTemplate(w, "templates/customer.html", nil)
	case http.MethodPost:
		cookie, err := r.Cookie("session_token")
		if err != nil {
			handleError(w, err.Error(), http.StatusUnauthorized)
			return
		}

		decryptedClientCode, decryptedSessionKey, err := app.getClientCodeAndSessionKey(cookie.Value)
		if err != nil {
			handleError(w, err.Error(), http.StatusInternalServerError)
			return
		}

		customerID := r.FormValue("customer_id")

		customerDBTimestamp, err := app.DB.GetCustomerAddedTimestampFromDB(customerID, decryptedClientCode)
		isDBCustomerExpired := err != nil || utils.IsDatabaseCustomerExpired(customerDBTimestamp)

		data := struct {
			Customer     models.CustomerRecord
			ErrorMessage string
		}{}

		if isDBCustomerExpired {
			customer, err := app.getCustomerByID(decryptedClientCode, decryptedSessionKey, customerID)
			if err != nil {
				handleError(w, internalServerError, http.StatusInternalServerError)
				return
			}

			if len(customer.Records) == 0 {
				errorMessage := "Customer not found."
				data.Customer = models.CustomerRecord{}
				data.ErrorMessage = errorMessage
				renderTemplate(w, "templates/customer.html", data)
				return
			}

			err = app.DB.AddOrUpdateCustomerToDB(
				customer.Records[0].CustomerID,
				decryptedClientCode,
				customer.Records[0].FullName,
				customer.Records[0].Email,
				customer.Records[0].Phone,
			)
			if err != nil {
				handleError(w, internalServerError, http.StatusInternalServerError)
				return
			}

			data.Customer = customer.Records[0]
		} else {
			customer, err := app.DB.GetCustomerFromDB(customerID, decryptedClientCode)
			if err != nil {
				handleError(w, internalServerError, http.StatusInternalServerError)
				return
			}
			data.Customer = *customer
		}

		renderTemplate(w, "templates/customer.html", data)
	default:
		handleError(w, methodNotAllowedMsg, http.StatusMethodNotAllowed)
	}
}

func (app *application) saveCustomerHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != saveCustomerPath {
		handleError(w, notFoundMessage, http.StatusNotFound)
		return
	}

	switch r.Method {
	case http.MethodGet:
		renderTemplate(w, "templates/savecustomer.html", nil)
	case http.MethodPost:
		cookie, err := r.Cookie("session_token")
		if err != nil {
			handleError(w, err.Error(), http.StatusUnauthorized)
			return
		}

		decryptedClientCode, decryptedSessionKey, err := app.getClientCodeAndSessionKey(cookie.Value)
		if err != nil {
			handleError(w, err.Error(), http.StatusUnauthorized)
			return
		}

		fullName := r.FormValue("fullName")
		email := r.FormValue("email")
		phoneNumber := r.FormValue("phone")

		err = app.saveCustomer(decryptedClientCode, decryptedSessionKey, fullName, email, phoneNumber)
		if err != nil {
			handleError(w, err.Error(), http.StatusMethodNotAllowed)
			return
		}

		successMessage := "Customer saved successfully!"

		data := struct {
			SuccessMessage string
		}{
			SuccessMessage: successMessage,
		}

		renderTemplate(w, "templates/savecustomer.html", data)
	default:
		handleError(w, methodNotAllowedMsg, http.StatusMethodNotAllowed)
	}
}

func renderTemplate(w http.ResponseWriter, templatePath string, data interface{}) {
	tmpl, err := template.ParseFiles(templatePath)
	if err != nil {
		handleError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		handleError(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

type ErrorResponse struct {
	Code    int
	Message string
}

func handleError(w http.ResponseWriter, errMsg string, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	errorResponse := ErrorResponse{
		Message: errMsg,
		Code:    code,
	}

	err := json.NewEncoder(w).Encode(errorResponse)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
