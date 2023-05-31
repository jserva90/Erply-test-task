package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

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

		sessionInfo, err := app.getSessionInfo(decryptedClientCode, decryptedSessionKey)
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

	// Execute the template and pass the data
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

func (app *application) verifyUser(clientCode, username, password string) (*models.Response, error) {
	apiURL := fmt.Sprintf("https://%s.erply.com/api/", clientCode)
	data := url.Values{}
	data.Set("clientCode", clientCode)
	data.Set("username", username)
	data.Set("password", password)
	data.Set("request", "verifyUser")
	data.Set("sendContentType", "1")

	req, err := http.NewRequest("POST", apiURL, strings.NewReader(data.Encode()))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var response models.Response
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	if response.Status.ErrorCode != 0 {
		return nil, fmt.Errorf("failed to authenticate")
	}

	return &response, nil
}

func (app *application) getSessionInfo(clientCode, sessionKey string) (*models.GetSessionKeyInfoResponse, error) {
	apiURL := fmt.Sprintf("https://%s.erply.com/api/", clientCode)
	data := url.Values{}
	data.Set("clientCode", clientCode)
	data.Set("sessionKey", sessionKey)
	data.Set("request", "getSessionKeyInfo")
	data.Set("sendContentType", "1")

	req, err := http.NewRequest("POST", apiURL, strings.NewReader(data.Encode()))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var response models.GetSessionKeyInfoResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	if response.Status.ErrorCode != 0 {
		return nil, fmt.Errorf("failed to authenticate")
	}

	return &response, nil
}

func (app *application) getCustomers(clientCode, sessionKey string) (*models.CustomerResponse, error) {
	apiURL := fmt.Sprintf("https://%s.erply.com/api/", clientCode)
	data := url.Values{}
	data.Set("clientCode", clientCode)
	data.Set("sessionKey", sessionKey)
	data.Set("request", "getCustomers")
	data.Set("sendContentType", "1")

	req, err := http.NewRequest("POST", apiURL, strings.NewReader(data.Encode()))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var response models.CustomerResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (app *application) saveCustomer(clientCode, sessionKey, fullName, email, phoneNumber string) error {
	apiURL := fmt.Sprintf("https://%s.erply.com/api/", clientCode)
	data := url.Values{}
	data.Set("clientCode", clientCode)
	data.Set("sessionKey", sessionKey)
	data.Set("request", "saveCustomer")
	data.Set("fullName", fullName)
	data.Set("email", email)
	data.Set("phone", phoneNumber)
	data.Set("sendContentType", "1")

	req, err := http.NewRequest("POST", apiURL, strings.NewReader(data.Encode()))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var response models.SaveCustomerResponse
	err = json.Unmarshal([]byte(body), &response)
	if err != nil {
		return err
	}

	if response.Status.ErrorCode != 0 {
		return errors.New("could not save customer")
	}

	return nil
}
