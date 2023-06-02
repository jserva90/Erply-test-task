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
)

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

	var bodyBuffer bytes.Buffer
	_, err = io.Copy(&bodyBuffer, resp.Body)
	if err != nil {
		return nil, err
	}

	var response models.Response
	err = json.Unmarshal(bodyBuffer.Bytes(), &response)
	if err != nil {
		return nil, err
	}

	if response.Status.ErrorCode != 0 {
		return nil, fmt.Errorf("failed to authenticate")
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
// @Host https://531748.erply.com/api/
// @Router /verifyUser [post]
func (app *application) verifyUserSwagger(w http.ResponseWriter, r *http.Request) {
	clientCode := r.URL.Query().Get("clientCode")
	username := r.URL.Query().Get("username")
	password := r.URL.Query().Get("password")

	apiURL := fmt.Sprintf("https://%s.erply.com/api/", clientCode)
	data := url.Values{}
	data.Set("clientCode", clientCode)
	data.Set("username", username)
	data.Set("password", password)
	data.Set("request", "verifyUser")
	data.Set("sendContentType", "1")

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

	var bodyBuffer bytes.Buffer
	_, err = io.Copy(&bodyBuffer, resp.Body)
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
	data := url.Values{}
	data.Set("clientCode", clientCode)
	data.Set("sessionKey", sessionKey)
	data.Set("request", "getSessionKeyInfo")
	data.Set("sendContentType", "1")

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
	data := url.Values{}
	data.Set("clientCode", clientCode)
	data.Set("sessionKey", sessionKey)
	data.Set("request", "getCustomers")
	data.Set("sendContentType", "1")

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
		return errors.New("could not save customer")
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
