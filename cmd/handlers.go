package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gofrs/uuid"
	"github.com/jserva90/Erply-test-task/models"
	"github.com/jserva90/Erply-test-task/utils"
)

type Error struct {
	Code    int
	Message string
}

func (app *application) MainPage(w http.ResponseWriter, r *http.Request) {
	template := template.Must(template.ParseFiles("templates/main.html"))

	if r.URL.Path != "/admin/main" {
		ErrorHandler(w, "not found", http.StatusNotFound)
		return
	}

	switch r.Method {
	case "GET":
		err := template.Execute(w, nil)
		if err != nil {
			ErrorHandler(w, "not found", http.StatusNotFound)
		}
	default:
		ErrorHandler(w, "method not supported", http.StatusMethodNotAllowed)
	}
}

func (app *application) Home(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session_token")
	if err == nil {
		session, err := app.DB.GetSession(cookie.Value)
		if err != nil {
			fmt.Println(err)
			goto ContinueExecution
		}
		decryptedClientCode, err := utils.Decrypt(session.ClientCode, utils.SecretKey)
		if err != nil {
			fmt.Println(err)
			goto ContinueExecution
		}
		decryptedSessionKey, err := utils.Decrypt(session.SessionKey, utils.SecretKey)
		if err != nil {
			fmt.Println(err)
			goto ContinueExecution
		}

		sessionInfo, err := app.getSessionKeyInfo(decryptedClientCode, decryptedSessionKey)
		if err != nil {
			fmt.Println(err)

			goto ContinueExecution
		}

		if utils.IsSessionExpired(*sessionInfo) {
			goto ContinueExecution
		}

		http.Redirect(w, r, "/admin/main", http.StatusSeeOther)
	}

ContinueExecution:
	template := template.Must(template.ParseFiles("templates/index.html"))

	if r.URL.Path != "/" {
		ErrorHandler(w, "not found", http.StatusNotFound)
		return
	}
	switch r.Method {
	case "GET":

		err := template.Execute(w, nil)
		if err != nil {
			ErrorHandler(w, "not found", http.StatusNotFound)
		}
	case "POST":
		clientCode := r.FormValue("clientCode")
		username := r.FormValue("username")
		password := r.FormValue("password")

		res, err := app.verifyUser(clientCode, username, password)
		if err != nil {
			ErrorHandler(w, err.Error(), http.StatusNotFound)
			return
		}

		uuid := uuid.Must(uuid.NewV4()).String()

		err = app.DB.AddSession(clientCode, username, password, res.Records[0].SessionKey, uuid)
		if err != nil {
			ErrorHandler(w, err.Error(), http.StatusNotFound)
			return
		}

		utils.CreateCookie(w, uuid)

		http.Redirect(w, r, "/admin/main", http.StatusSeeOther)
	default:
		ErrorHandler(w, "method not supported", http.StatusMethodNotAllowed)
	}
}

func (app *application) Logout(w http.ResponseWriter, r *http.Request) {
	currentCookie, err := r.Cookie("session_token")
	if err != nil {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	currentSessionToken := currentCookie.Value
	err = app.DB.RemoveSession(currentSessionToken)
	if err != nil {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	utils.DeleteCookie(w, r)

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (app *application) FetchCustomers(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session_token")
	if err != nil {
		return
	}

	decryptedClientCode, decryptedSessionKey, err := app.getClientCodeAndSessionKey(cookie.Value)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	customers, err := app.getCustomers(decryptedClientCode, decryptedSessionKey)
	if err != nil {
		return
	}

	data := struct {
		Customers models.CustomerResponse
	}{
		Customers: *customers,
	}

	tmpl, err := template.ParseFiles("templates/result.html")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func (app *application) SaveCustomer(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/admin/savecustomer" {
		ErrorHandler(w, "not found", http.StatusNotFound)
		return
	}

	template := template.Must(template.ParseFiles("templates/savecustomer.html"))
	switch r.Method {
	case "GET":
		err := template.Execute(w, nil)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
	case "POST":
		cookie, err := r.Cookie("session_token")
		if err != nil {
			return
		}

		decryptedClientCode, decryptedSessionKey, err := app.getClientCodeAndSessionKey(cookie.Value)
		if err != nil {
			ErrorHandler(w, err.Error(), http.StatusUnauthorized)
		}

		fullName := r.FormValue("fullName")
		email := r.FormValue("email")
		phoneNumber := r.FormValue("phone")

		err = app.saveCustomer(decryptedClientCode, decryptedSessionKey, fullName, email, phoneNumber)
		if err != nil {
			ErrorHandler(w, err.Error(), http.StatusMethodNotAllowed)
		}
		http.Redirect(w, r, "/admin/success", http.StatusSeeOther)
	default:
		ErrorHandler(w, "method not supported", http.StatusMethodNotAllowed)
	}
}

func (app *application) Success(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/admin/success" {
		ErrorHandler(w, "not found", http.StatusNotFound)
		return
	}

	template := template.Must(template.ParseFiles("templates/success.html"))

	switch r.Method {
	case "GET":
		err := template.Execute(w, nil)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
	default:
		ErrorHandler(w, "method not supported", http.StatusMethodNotAllowed)
	}
}

func ErrorHandler(w http.ResponseWriter, er string, code int) {
	w.WriteHeader(code)
	e := Error{Message: er, Code: code}
	html, err := template.ParseFiles("templates/error.html")
	if err != nil {
		http.Error(w, "500: Internal Server Error", http.StatusInternalServerError)
		return
	}
	err = html.Execute(w, e)
	if err != nil {
		http.Error(w, "404: Not Found", 404)
		return
	}
}
