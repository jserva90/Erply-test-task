package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/jserva90/Erply-test-task/models"
	"github.com/jserva90/Erply-test-task/utils"
)

func (app *application) verifyUser(clientCode, username, password string) (*models.Response, error) {
	apiURL := fmt.Sprintf("https://%s.erply.com/api/", clientCode)

	data := url.Values{
		"clientCode":      {clientCode},
		"username":        {username},
		"password":        {password},
		"request":         {"verifyUser"},
		"sendContentType": {"1"},
	}

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

	bodyBuffer := new(bytes.Buffer)
	_, err = io.Copy(bodyBuffer, resp.Body)
	if err != nil {
		return nil, err
	}

	var response models.Response
	err = json.Unmarshal(bodyBuffer.Bytes(), &response)
	if err != nil {
		return nil, err
	}

	if response.Status.ErrorCode != 0 {
		return nil, fmt.Errorf("authentication failed: %v", response.Status.ErrorCode)
	}

	return &response, nil
}

// verifyUser is a handler function for verifying user credentials with ERPLY API.
// @Summary Verify user credentials
// @Description Verifies the user credentials by making a request to ERPLY API
// @Param clientCode query string true "Client code"
// @Param username query string true "Username"
// @Param password query string true "Password"
// @Produce json
// @Success 200 {object} models.Response
// @Router /verifyUser [post]
func (app *application) verifyUserSwagger(w http.ResponseWriter, r *http.Request) {
	clientCode := r.URL.Query().Get("clientCode")
	username := r.URL.Query().Get("username")
	password := r.URL.Query().Get("password")

	apiURL := fmt.Sprintf("https://%s.erply.com/api/", clientCode)
	data := url.Values{
		"clientCode":      {clientCode},
		"username":        {username},
		"password":        {password},
		"request":         {"verifyUser"},
		"sendContentType": {"1"},
	}

	req, err := http.NewRequest("POST", apiURL, strings.NewReader(data.Encode()))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	bodyBuffer := new(bytes.Buffer)
	_, err = io.Copy(bodyBuffer, resp.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var response models.Response
	err = json.Unmarshal(bodyBuffer.Bytes(), &response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if response.Status.ErrorCode != 0 {
		http.Error(w, "Failed to authenticate", http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (app *application) getSessionKeyInfo(clientCode, sessionKey string) (*models.GetSessionKeyInfoResponse, error) {
	apiURL := fmt.Sprintf("https://%s.erply.com/api/", clientCode)
	data := url.Values{
		"clientCode":      {clientCode},
		"sessionKey":      {sessionKey},
		"request":         {"getSessionKeyInfo"},
		"sendContentType": {"1"},
	}

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

	bodyBuffer := new(bytes.Buffer)
	_, err = io.Copy(bodyBuffer, resp.Body)
	if err != nil {
		return nil, err
	}

	var response models.GetSessionKeyInfoResponse
	err = json.Unmarshal(bodyBuffer.Bytes(), &response)
	if err != nil {
		return nil, err
	}

	if response.Status.ErrorCode != 0 {
		return nil, fmt.Errorf("failed to authenticate")
	}

	return &response, nil
}

// getSessionKeyInfo is a handler function for retrieving session key information from ERPLY API.
// @Summary Get session key information
// @Description Retrieves session key information by making a request to ERPLY API
// @Param clientCode query string true "Client code"
// @Param sessionKey query string true "Session key"
// @Produce json
// @Success 200 {object} models.GetSessionKeyInfoResponse
// @Router /getSessionKeyInfo [post]
func (app *application) getSessionKeyInfoSwagger(w http.ResponseWriter, r *http.Request) {
	clientCode := r.URL.Query().Get("clientCode")
	sessionKey := r.URL.Query().Get("sessionKey")

	apiURL := fmt.Sprintf("https://%s.erply.com/api/", clientCode)
	data := url.Values{
		"clientCode":      {clientCode},
		"sessionKey":      {sessionKey},
		"request":         {"getSessionKeyInfo"},
		"sendContentType": {"1"},
	}

	req, err := http.NewRequest("POST", apiURL, strings.NewReader(data.Encode()))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	var bodyBuffer bytes.Buffer
	_, err = io.Copy(&bodyBuffer, resp.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var response models.GetSessionKeyInfoResponse
	err = json.Unmarshal(bodyBuffer.Bytes(), &response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if response.Status.ErrorCode != 0 {
		http.Error(w, "Failed to authenticate", http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (app *application) getCustomers(clientCode, sessionKey string) (*models.CustomerResponse, error) {
	apiURL := fmt.Sprintf("https://%s.erply.com/api/", clientCode)
	data := url.Values{
		"clientCode":      {clientCode},
		"sessionKey":      {sessionKey},
		"request":         {"getCustomers"},
		"sendContentType": {"1"},
	}

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

	var bodyBuffer bytes.Buffer
	_, err = io.Copy(&bodyBuffer, resp.Body)
	if err != nil {
		return nil, err
	}

	var response models.CustomerResponse
	err = json.Unmarshal(bodyBuffer.Bytes(), &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// getCustomers is a handler function for retrieving customer information from ERPLY API.
// @Summary Get customer information
// @Description Retrieves customer information by making a request to ERPLY API
// @Param clientCode query string true "Client code"
// @Param sessionKey query string true "Session key"
// @Produce json
// @Success 200 {object} models.CustomerResponse
// @Router /getCustomers [post]
func (app *application) getCustomersSwagger(w http.ResponseWriter, r *http.Request) {
	clientCode := r.URL.Query().Get("clientCode")
	sessionKey := r.URL.Query().Get("sessionKey")

	apiURL := fmt.Sprintf("https://%s.erply.com/api/", clientCode)
	data := url.Values{
		"clientCode":      {clientCode},
		"sessionKey":      {sessionKey},
		"request":         {"getCustomers"},
		"sendContentType": {"1"},
	}

	req, err := http.NewRequest("POST", apiURL, strings.NewReader(data.Encode()))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	var bodyBuffer bytes.Buffer
	_, err = io.Copy(&bodyBuffer, resp.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var response models.CustomerResponse
	err = json.Unmarshal(bodyBuffer.Bytes(), &response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (app *application) getCustomerByID(clientCode, sessionKey, customerID string) (*models.CustomerResponse, error) {
	apiURL := fmt.Sprintf("https://%s.erply.com/api/", clientCode)
	data := url.Values{
		"clientCode":      {clientCode},
		"sessionKey":      {sessionKey},
		"customerID":      {customerID},
		"request":         {"getCustomers"},
		"sendContentType": {"1"},
	}

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

	var bodyBuffer bytes.Buffer
	_, err = io.Copy(&bodyBuffer, resp.Body)
	if err != nil {
		return nil, err
	}

	var response models.CustomerResponse
	err = json.Unmarshal(bodyBuffer.Bytes(), &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// getCustomersByID is a handler function for retrieving customer information from ERPLY API by customer ID.
// @Summary Get customer information by customer ID.
// @Description Retrieves customer information either from local database (if customer data in database is less than 10 minutes old) or by making a request to ERPLY API
// @Param clientCode query string true "Client code"
// @Param sessionKey query string true "Session key"
// @Param customerID query string true "Customer ID"
// @Produce json
// @Success 200 {object} models.CustomerResponse
// @Router /getCustomerByID [post]
func (app *application) getCustomerByIDSwagger(w http.ResponseWriter, r *http.Request) {
	clientCode := r.URL.Query().Get("clientCode")
	sessionKey := r.URL.Query().Get("sessionKey")
	customerID := r.URL.Query().Get("customerID")

	var response models.CustomerResponse

	customerDBTimestamp, err := app.DB.GetCustomerAddedTimestampFromDB(customerID, clientCode)
	isDBCustomerExpired := err != nil || utils.IsDatabaseCustomerExpired(customerDBTimestamp)

	if isDBCustomerExpired {
		apiURL := fmt.Sprintf("https://%s.erply.com/api/", clientCode)
		data := url.Values{
			"clientCode":      {clientCode},
			"sessionKey":      {sessionKey},
			"customerID":      {customerID},
			"request":         {"getCustomers"},
			"sendContentType": {"1"},
		}

		req, err := http.NewRequest("POST", apiURL, strings.NewReader(data.Encode()))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()

		var bodyBuffer bytes.Buffer
		_, err = io.Copy(&bodyBuffer, resp.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		err = json.Unmarshal(bodyBuffer.Bytes(), &response)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if len(response.Records) > 0 {
			err = app.DB.AddOrUpdateCustomerToDB(
				response.Records[0].CustomerID,
				clientCode,
				response.Records[0].FullName,
				response.Records[0].Email,
				response.Records[0].Phone,
			)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}

	} else {
		customer, err := app.DB.GetCustomerFromDB(customerID, clientCode)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		response.Records = append(response.Records, *customer)
		response.Status.ResponseStatus = "ok"
		response.Status.RequestFromLocalDB = true
	}

	w.Header().Set("Content-Type", "application/json")

	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (app *application) saveCustomer(clientCode, sessionKey, fullName, email, phoneNumber string) error {
	apiURL := fmt.Sprintf("https://%s.erply.com/api/", clientCode)
	data := url.Values{
		"clientCode":      {clientCode},
		"sessionKey":      {sessionKey},
		"request":         {"saveCustomer"},
		"fullName":        {fullName},
		"email":           {email},
		"phone":           {phoneNumber},
		"sendContentType": {"1"},
	}

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

	var bodyBuffer bytes.Buffer
	_, err = io.Copy(&bodyBuffer, resp.Body)
	if err != nil {
		return err
	}

	var response models.SaveCustomerResponse
	err = json.Unmarshal(bodyBuffer.Bytes(), &response)
	if err != nil {
		return err
	}

	if response.Status.ErrorCode != 0 {
		return errors.New("failed to save customer")
	}

	return nil
}

// saveCustomer is a handler function for saving customer information to ERPLY API.
// @Summary Save customer information
// @Description Saves customer information by making a request to ERPLY API
// @Param clientCode query string true "Client code"
// @Param sessionKey query string true "Session key"
// @Param fullName query string true "Full name"
// @Param email query string true "Email"
// @Param phoneNumber query string true "Phone number"
// @Produce json
// @Success 200 {object} models.SaveCustomerResponse
// @Router /saveCustomer [post]
func (app *application) saveCustomerSwagger(w http.ResponseWriter, r *http.Request) {
	clientCode := r.URL.Query().Get("clientCode")
	sessionKey := r.URL.Query().Get("sessionKey")
	fullName := r.URL.Query().Get("fullName")
	email := r.URL.Query().Get("email")
	phoneNumber := r.URL.Query().Get("phoneNumber")

	apiURL := fmt.Sprintf("https://%s.erply.com/api/", clientCode)
	data := url.Values{
		"clientCode":      {clientCode},
		"sessionKey":      {sessionKey},
		"request":         {"saveCustomer"},
		"fullName":        {fullName},
		"email":           {email},
		"phone":           {phoneNumber},
		"sendContentType": {"1"},
	}

	req, err := http.NewRequest("POST", apiURL, strings.NewReader(data.Encode()))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	var bodyBuffer bytes.Buffer
	_, err = io.Copy(&bodyBuffer, resp.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var response models.SaveCustomerResponse
	err = json.Unmarshal(bodyBuffer.Bytes(), &response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if response.Status.ErrorCode != 0 {
		http.Error(w, "Could not save customer", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
