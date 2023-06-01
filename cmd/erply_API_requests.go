package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
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
