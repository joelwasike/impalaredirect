package main

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

func signIn(c *gin.Context) {
	url := "http://staging.mam-laka.com/api/?resource=user&action=login"

	// Define a struct to parse the JSON request body
	var credentials struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	// Bind JSON request body to the credentials struct
	if err := c.ShouldBindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email and password are required"})
		return
	}

	// Prepare the body for the external request
	body := map[string]string{
		"email":    credentials.Email,
		"password": credentials.Password,
	}
	jsonData, err := json.Marshal(body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error encoding JSON"})
		return
	}

	// Create the request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating request"})
		return
	}

	// Set headers
	req.Header.Set("Authorization", "Bearer ZjdkMjZhNTJmMzZkNTYwNzI0YTE5MGUwODVjM2UwMjI=")
	req.Header.Set("Content-Type", "application/json")

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error sending request"})
		return
	}
	defer resp.Body.Close()

	// Parse the response JSON into a map
	var responseJSON map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&responseJSON); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error parsing response JSON"})
		return
	}

	// Return the JSON response
	c.JSON(resp.StatusCode, responseJSON)
}
func register(c *gin.Context) {
	url := "http://staging.mam-laka.com/api/?resource=user&action=create"

	// Define a struct to parse the JSON request body
	var registrationData struct {
		Name            string `json:"name" binding:"required"`
		Email           string `json:"email" binding:"required,email"`
		Phone           string `json:"phone" binding:"required"`
		Password        string `json:"password" binding:"required"`
		Occupation      string `json:"occupation" binding:"required"`
		ConfirmPassword string `json:"confirm_password" binding:"required"`
		Country         string `json:"country" binding:"required"`
		Gender          string `json:"gender" binding:"required"`
		TOC             string `json:"toc" binding:"required"`
		BaseCurrency    string `json:"baseCurrency" binding:"required"`
		Pin             string `json:"pin" binding:"required"`
	}

	// Bind JSON request body to the struct
	if err := c.ShouldBindJSON(&registrationData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "All fields are required"})
		return
	}

	// Prepare the body for the external request
	jsonData, err := json.Marshal(registrationData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error encoding JSON"})
		return
	}

	// Create the request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating request"})
		return
	}

	// Set headers
	req.Header.Set("Authorization", "Bearer ZjdkMjZhNTJmMzZkNTYwNzI0YTE5MGUwODVjM2UwMjI=")
	req.Header.Set("Content-Type", "application/json")

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error sending request"})
		return
	}
	defer resp.Body.Close()

	// Parse the response JSON into a map
	var responseJSON map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&responseJSON); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error parsing response JSON"})
		return
	}

	// Return the JSON response
	c.JSON(resp.StatusCode, responseJSON)
}
func fetchbeneficiary(c *gin.Context) {
	url := "http://staging.mam-laka.com/api/?resource=beneficiary&action=read"

	// Define a struct to parse the JSON request body
	var credentials struct {
		Phone string `json:"phone" binding:"required"`
	}

	// Bind JSON request body to the credentials struct
	if err := c.ShouldBindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email and password are required"})
		return
	}

	// Prepare the body for the external request
	body := map[string]string{
		"phone": credentials.Phone,
	}
	jsonData, err := json.Marshal(body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error encoding JSON"})
		return
	}

	// Create the request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating request"})
		return
	}

	// Set headers
	req.Header.Set("Authorization", "Bearer ZjdkMjZhNTJmMzZkNTYwNzI0YTE5MGUwODVjM2UwMjI=")
	req.Header.Set("Content-Type", "application/json")

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error sending request"})
		return
	}
	defer resp.Body.Close()

	// Parse the response JSON into a map
	var responseJSON map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&responseJSON); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error parsing response JSON"})
		return
	}

	// Return the JSON response
	c.JSON(resp.StatusCode, responseJSON)
}
func checkpin(c *gin.Context) {
	url := "http://staging.mam-laka.com/api/?resource=user&action=verifyPin"

	// Define a struct to parse the JSON request body
	var credentials struct {
		Phone string `json:"phone" binding:"required"`
		Pin   string `json:"pin" binding:"required"`
	}

	// Bind JSON request body to the credentials struct
	if err := c.ShouldBindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email and password are required"})
		return
	}

	// Prepare the body for the external request
	body := map[string]string{
		"phone": credentials.Phone,
		"pin":   credentials.Pin,
	}
	jsonData, err := json.Marshal(body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error encoding JSON"})
		return
	}

	// Create the request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating request"})
		return
	}

	// Set headers
	req.Header.Set("Authorization", "Bearer ZjdkMjZhNTJmMzZkNTYwNzI0YTE5MGUwODVjM2UwMjI=")
	req.Header.Set("Content-Type", "application/json")

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error sending request"})
		return
	}
	defer resp.Body.Close()

	// Parse the response JSON into a map
	var responseJSON map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&responseJSON); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error parsing response JSON"})
		return
	}

	// Return the JSON response
	c.JSON(resp.StatusCode, responseJSON)
}
func checkbalance(c *gin.Context) {
	url := "http://staging.mam-laka.com/api/?resource=balance&action=read"

	// Define a struct to parse the JSON request body
	var credentials struct {
		Phone string `json:"phone" binding:"required"`
	}

	// Bind JSON request body to the credentials struct
	if err := c.ShouldBindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email and password are required"})
		return
	}

	// Prepare the body for the external request
	body := map[string]string{
		"phone": credentials.Phone,
	}
	jsonData, err := json.Marshal(body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error encoding JSON"})
		return
	}

	// Create the request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating request"})
		return
	}

	// Set headers
	req.Header.Set("Authorization", "Bearer ZjdkMjZhNTJmMzZkNTYwNzI0YTE5MGUwODVjM2UwMjI=")
	req.Header.Set("Content-Type", "application/json")

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error sending request"})
		return
	}
	defer resp.Body.Close()

	// Parse the response JSON into a map
	var responseJSON map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&responseJSON); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error parsing response JSON"})
		return
	}

	// Return the JSON response
	c.JSON(resp.StatusCode, responseJSON)
}
func transactions(c *gin.Context) {
	url := "http://staging.mam-laka.com/api/?resource=transaction&action=read"

	// Define a struct to parse the JSON request body
	var credentials struct {
		Phone string `json:"phone" binding:"required"`
	}

	// Bind JSON request body to the credentials struct
	if err := c.ShouldBindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email and password are required"})
		return
	}

	// Prepare the body for the external request
	body := map[string]string{
		"phone": credentials.Phone,
	}
	jsonData, err := json.Marshal(body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error encoding JSON"})
		return
	}

	// Create the request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating request"})
		return
	}

	// Set headers
	req.Header.Set("Authorization", "Bearer ZjdkMjZhNTJmMzZkNTYwNzI0YTE5MGUwODVjM2UwMjI=")
	req.Header.Set("Content-Type", "application/json")

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error sending request"})
		return
	}
	defer resp.Body.Close()

	// Parse the response JSON into a map
	var responseJSON map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&responseJSON); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error parsing response JSON"})
		return
	}

	// Return the JSON response
	c.JSON(resp.StatusCode, responseJSON)
}
func forgetpass(c *gin.Context) {
	url := "http://staging.mam-laka.com/api/?resource=password&action=forgot"

	// Define a struct to parse the JSON request body
	var credentials struct {
		Userid string `json:"userId" binding:"required"`
	}

	// Bind JSON request body to the credentials struct
	if err := c.ShouldBindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email and password are required"})
		return
	}

	// Prepare the body for the external request
	body := map[string]string{
		"userId": credentials.Userid,
	}
	jsonData, err := json.Marshal(body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error encoding JSON"})
		return
	}

	// Create the request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating request"})
		return
	}

	// Set headers
	req.Header.Set("Authorization", "Bearer ZjdkMjZhNTJmMzZkNTYwNzI0YTE5MGUwODVjM2UwMjI=")
	req.Header.Set("Content-Type", "application/json")

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error sending request"})
		return
	}
	defer resp.Body.Close()

	// Parse the response JSON into a map
	var responseJSON map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&responseJSON); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error parsing response JSON"})
		return
	}

	// Return the JSON response
	c.JSON(resp.StatusCode, responseJSON)
}
func withdraw(c *gin.Context) {
	url := "http://staging.mam-laka.com/api/?resource=transaction&action=withdraw"

	// Define a struct to parse the JSON request body
	var credentials struct {
		Phone        string `json:"phone" binding:"required"`
		Amount       string `json:"amount" binding:"required"`
		CurrencyCode string `json:"currencyCode" binding:"required"`
	}

	// Bind JSON request body to the credentials struct
	if err := c.ShouldBindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email and password are required"})
		return
	}

	// Prepare the body for the external request
	body := map[string]string{
		"phone":        credentials.Phone,
		"amount":       credentials.Amount,
		"currencyCode": credentials.CurrencyCode,
	}
	jsonData, err := json.Marshal(body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error encoding JSON"})
		return
	}

	// Create the request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating request"})
		return
	}

	// Set headers
	req.Header.Set("Authorization", "Bearer ZjdkMjZhNTJmMzZkNTYwNzI0YTE5MGUwODVjM2UwMjI=")
	req.Header.Set("Content-Type", "application/json")

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error sending request"})
		return
	}
	defer resp.Body.Close()

	// Parse the response JSON into a map
	var responseJSON map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&responseJSON); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error parsing response JSON"})
		return
	}

	// Return the JSON response
	c.JSON(resp.StatusCode, responseJSON)
}
func deposit(c *gin.Context) {
	url := "http://staging.mam-laka.com/api/?resource=transaction&action=deposit	"

	// Define a struct to parse the JSON request body
	var credentials struct {
		Phone         string `json:"phone" binding:"required"`
		Amount        string `json:"amount" binding:"required"`
		Currency      string `json:"currency" binding:"required"`
		SourceOfFunds string `json:"sourceOfFunds" binding:"required"`
		PayerPhone    string `json:"payerPhone" binding:"required"`
	}

	// Bind JSON request body to the credentials struct
	if err := c.ShouldBindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email and password are required"})
		return
	}

	// Prepare the body for the external request
	body := map[string]string{
		"phone":         credentials.Phone,
		"amount":        credentials.Amount,
		"currency":      credentials.Currency,
		"sourceOfFunds": credentials.SourceOfFunds,
		"payerPhone":    credentials.PayerPhone,
	}
	jsonData, err := json.Marshal(body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error encoding JSON"})
		return
	}

	// Create the request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating request"})
		return
	}

	// Set headers
	req.Header.Set("Authorization", "Bearer ZjdkMjZhNTJmMzZkNTYwNzI0YTE5MGUwODVjM2UwMjI=")
	req.Header.Set("Content-Type", "application/json")

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error sending request"})
		return
	}
	defer resp.Body.Close()

	// Parse the response JSON into a map
	var responseJSON map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&responseJSON); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error parsing response JSON"})
		return
	}

	// Return the JSON response
	c.JSON(resp.StatusCode, responseJSON)
}
func sendmoney(c *gin.Context) {
	url := "http://staging.mam-laka.com/api/?resource=transaction&action=create"

	// Define a struct to parse the JSON request body
	var credentials struct {
		Phone           string `json:"phone" binding:"required"`
		RecipientPhone  string `json:"recipientPhone" binding:"required"`
		Amount          string `json:"amount" binding:"required"`
		RecipientName   string `json:"recipientName" binding:"required"`
		SendingCurrency string `json:"sendingCurrency" binding:"required"`
		SourceOfFunds   string `json:"sourceOfFunds" binding:"required"`
		PayerPhone      string `json:"payerPhone" binding:"required"`
	}

	// Bind JSON request body to the credentials struct
	if err := c.ShouldBindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email and password are required"})
		return
	}

	// Prepare the body for the external request
	body := map[string]string{
		"phone":           credentials.Phone,
		"amount":          credentials.Amount,
		"sendingCurrency": credentials.SendingCurrency,
		"sourceOfFunds":   credentials.SourceOfFunds,
		"payerPhone":      credentials.PayerPhone,
		"recipientPhone":  credentials.RecipientPhone,
		"recipientName":   credentials.RecipientName,
	}
	jsonData, err := json.Marshal(body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error encoding JSON"})
		return
	}

	// Create the request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating request"})
		return
	}

	// Set headers
	req.Header.Set("Authorization", "Bearer ZjdkMjZhNTJmMzZkNTYwNzI0YTE5MGUwODVjM2UwMjI=")
	req.Header.Set("Content-Type", "application/json")

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error sending request"})
		return
	}
	defer resp.Body.Close()

	// Parse the response JSON into a map
	var responseJSON map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&responseJSON); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error parsing response JSON"})
		return
	}

	// Return the JSON response
	c.JSON(resp.StatusCode, responseJSON)
}
func buyairtime(c *gin.Context) {
	url := "http://staging.mam-laka.com/api/?resource=utilities&action=sendAirtime"

	// Define a struct to parse the JSON request body
	var credentials struct {
		Phone          string `json:"phone" binding:"required"`
		RecipientPhone string `json:"recipientPhone" binding:"required"`
		Amount         string `json:"amount" binding:"required"`
		Network        string `json:"network" binding:"required"`
		SourceOfFunds  string `json:"sourceOfFunds" binding:"required"`
		PayerPhone     string `json:"payerPhone" binding:"required"`
	}

	// Bind JSON request body to the credentials struct
	if err := c.ShouldBindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email and password are required"})
		return
	}

	// Prepare the body for the external request
	body := map[string]string{
		"phone":          credentials.Phone,
		"amount":         credentials.Amount,
		"network":        credentials.Network,
		"sourceOfFunds":  credentials.SourceOfFunds,
		"payerPhone":     credentials.PayerPhone,
		"recipientPhone": credentials.RecipientPhone,
	}
	jsonData, err := json.Marshal(body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error encoding JSON"})
		return
	}

	// Create the request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating request"})
		return
	}

	// Set headers
	req.Header.Set("Authorization", "Bearer ZjdkMjZhNTJmMzZkNTYwNzI0YTE5MGUwODVjM2UwMjI=")
	req.Header.Set("Content-Type", "application/json")

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error sending request"})
		return
	}
	defer resp.Body.Close()

	// Parse the response JSON into a map
	var responseJSON map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&responseJSON); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error parsing response JSON"})
		return
	}

	// Return the JSON response
	c.JSON(resp.StatusCode, responseJSON)
}
func wallet2wallet(c *gin.Context) {
	url := "http://staging.mam-laka.com/api/?resource=transaction&action=wallet2wallet"

	// Define a struct to parse the JSON request body
	var credentials struct {
		Phone          string `json:"phone" binding:"required"`
		Amount         string `json:"amount" binding:"required"`
		RecipientEmail string `json:"recipientEmail" binding:"required"`
		CurrencyCode   string `json:"currencyCode" binding:"required"`
	}

	// Bind JSON request body to the credentials struct
	if err := c.ShouldBindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email and password are required"})
		return
	}

	// Prepare the body for the external request
	body := map[string]string{
		"phone":          credentials.Phone,
		"amount":         credentials.Amount,
		"recipientEmail": credentials.RecipientEmail,
		"currencyCode":   credentials.CurrencyCode,
	}
	jsonData, err := json.Marshal(body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error encoding JSON"})
		return
	}

	// Create the request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating request"})
		return
	}

	// Set headers
	req.Header.Set("Authorization", "Bearer ZjdkMjZhNTJmMzZkNTYwNzI0YTE5MGUwODVjM2UwMjI=")
	req.Header.Set("Content-Type", "application/json")

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error sending request"})
		return
	}
	defer resp.Body.Close()

	// Parse the response JSON into a map
	var responseJSON map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&responseJSON); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error parsing response JSON"})
		return
	}

	// Return the JSON response
	c.JSON(resp.StatusCode, responseJSON)
}
func convert(c *gin.Context) {
	url := "http://staging.mam-laka.com/api/?resource=balance&action=convert"

	// Define a struct to parse the JSON request body
	var credentials struct {
		Phone                   string `json:"phone" binding:"required"`
		Amount                  string `json:"amount" binding:"required"`
		SourceCurrencyCode      string `json:"sourceCurrencyCode" binding:"required"`
		DestinationCurrencyCode string `json:"destinationCurrencyCode" binding:"required"`
	}

	// Bind JSON request body to the credentials struct
	if err := c.ShouldBindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email and password are required"})
		return
	}

	// Prepare the body for the external request
	body := map[string]string{
		"phone":                   credentials.Phone,
		"amount":                  credentials.Amount,
		"sourceCurrencyCode":      credentials.SourceCurrencyCode,
		"destinationCurrencyCode": credentials.DestinationCurrencyCode,
	}
	jsonData, err := json.Marshal(body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error encoding JSON"})
		return
	}

	// Create the request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating request"})
		return
	}

	// Set headers
	req.Header.Set("Authorization", "Bearer ZjdkMjZhNTJmMzZkNTYwNzI0YTE5MGUwODVjM2UwMjI=")
	req.Header.Set("Content-Type", "application/json")

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error sending request"})
		return
	}
	defer resp.Body.Close()

	// Parse the response JSON into a map
	var responseJSON map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&responseJSON); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error parsing response JSON"})
		return
	}

	// Return the JSON response
	c.JSON(resp.StatusCode, responseJSON)
}
func changepassword(c *gin.Context) {
	url := "http://staging.mam-laka.com/api/?resource=password&action=change"

	// Define a struct to parse the JSON request body
	var credentials struct {
		UserId               string `json:"userId" binding:"required"`
		InputPasswordCurrent string `json:"inputPasswordCurrent" binding:"required"`
		SourceCurrencyCode   string `json:"sourceCurrencyCode" binding:"required"`
		InputPasswordConfirm string `json:"inputPasswordConfirm" binding:"required"`
	}

	// Bind JSON request body to the credentials struct
	if err := c.ShouldBindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email and password are required"})
		return
	}

	// Prepare the body for the external request
	body := map[string]string{
		"userId":               credentials.UserId,
		"inputPasswordCurrent": credentials.InputPasswordCurrent,
		"sourceCurrencyCode":   credentials.SourceCurrencyCode,
		"inputPasswordConfirm": credentials.InputPasswordConfirm,
	}
	jsonData, err := json.Marshal(body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error encoding JSON"})
		return
	}

	// Create the request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating request"})
		return
	}

	// Set headers
	req.Header.Set("Authorization", "Bearer ZjdkMjZhNTJmMzZkNTYwNzI0YTE5MGUwODVjM2UwMjI=")
	req.Header.Set("Content-Type", "application/json")

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error sending request"})
		return
	}
	defer resp.Body.Close()

	// Parse the response JSON into a map
	var responseJSON map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&responseJSON); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error parsing response JSON"})
		return
	}

	// Return the JSON response
	c.JSON(resp.StatusCode, responseJSON)
}
func changepin(c *gin.Context) {
	url := "http://staging.mam-laka.com/api/?resource=user&action=update"

	// Define a struct to parse the JSON request body
	var credentials struct {
		Phone string `json:"phone" binding:"required"`
		Pin   string `json:"pin" binding:"required"`
	}

	// Bind JSON request body to the credentials struct
	if err := c.ShouldBindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email and password are required"})
		return
	}

	// Prepare the body for the external request
	body := map[string]string{
		"phone": credentials.Phone,
		"pin":   credentials.Pin,
	}
	jsonData, err := json.Marshal(body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error encoding JSON"})
		return
	}

	// Create the request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating request"})
		return
	}

	// Set headers
	req.Header.Set("Authorization", "Bearer ZjdkMjZhNTJmMzZkNTYwNzI0YTE5MGUwODVjM2UwMjI=")
	req.Header.Set("Content-Type", "application/json")

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error sending request"})
		return
	}
	defer resp.Body.Close()

	// Parse the response JSON into a map
	var responseJSON map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&responseJSON); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error parsing response JSON"})
		return
	}

	// Return the JSON response
	c.JSON(resp.StatusCode, responseJSON)
}
func forex(c *gin.Context) {
	url := "http://staging.mam-laka.com/api/?resource=transaction&action=forex"

	// Define a struct to parse the JSON request body
	var credentials struct {
		SourceCurrencyCode      string `json:"sourceCurrencyCode" binding:"required"`
		DestinationCurrencyCode string `json:"destinationCurrencyCode" binding:"required"`
	}

	// Bind JSON request body to the credentials struct
	if err := c.ShouldBindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email and password are required"})
		return
	}

	// Prepare the body for the external request
	body := map[string]string{
		"sourceCurrencyCode":      credentials.SourceCurrencyCode,
		"destinationCurrencyCode": credentials.DestinationCurrencyCode,
	}
	jsonData, err := json.Marshal(body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error encoding JSON"})
		return
	}

	// Create the request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating request"})
		return
	}

	// Set headers
	req.Header.Set("Authorization", "Bearer ZjdkMjZhNTJmMzZkNTYwNzI0YTE5MGUwODVjM2UwMjI=")
	req.Header.Set("Content-Type", "application/json")

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error sending request"})
		return
	}
	defer resp.Body.Close()

	// Parse the response JSON into a map
	var responseJSON map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&responseJSON); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error parsing response JSON"})
		return
	}

	// Return the JSON response
	c.JSON(resp.StatusCode, responseJSON)
}
func changebasecurrency(c *gin.Context) {
	url := "http://staging.mam-laka.com/api/?resource=balance&action=update"

	// Define a struct to parse the JSON request body
	var credentials struct {
		Phone        string `json:"phone" binding:"required"`
		BaseCurrency string `json:"baseCurrency" binding:"required"`
	}

	// Bind JSON request body to the credentials struct
	if err := c.ShouldBindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email and password are required"})
		return
	}

	// Prepare the body for the external request
	body := map[string]string{
		"phone":        credentials.Phone,
		"baseCurrency": credentials.BaseCurrency,
	}
	jsonData, err := json.Marshal(body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error encoding JSON"})
		return
	}

	// Create the request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating request"})
		return
	}

	// Set headers
	req.Header.Set("Authorization", "Bearer ZjdkMjZhNTJmMzZkNTYwNzI0YTE5MGUwODVjM2UwMjI=")
	req.Header.Set("Content-Type", "application/json")

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error sending request"})
		return
	}
	defer resp.Body.Close()

	// Parse the response JSON into a map
	var responseJSON map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&responseJSON); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error parsing response JSON"})
		return
	}

	// Return the JSON response
	c.JSON(resp.StatusCode, responseJSON)
}
func addben(c *gin.Context) {
	url := "http://staging.mam-laka.com/api/?resource=beneficiary&action=create"

	// Define a struct to parse the JSON request body
	var credentials struct {
		Phone                   string `json:"phone" binding:"required"`
		BeneficiaryName         string `json:"beneficiaryName" binding:"required"`
		BeneficiaryPhone        string `json:"beneficiaryPhone" binding:"required"`
		BeneficiaryCountry      string `json:"beneficiaryCountry" binding:"required"`
		BeneficiaryRelationship string `json:"beneficiaryRelationship" binding:"required"`
	}

	// Bind JSON request body to the credentials struct
	if err := c.ShouldBindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email and password are required"})
		return
	}

	// Prepare the body for the external request
	body := map[string]string{
		"phone":                   credentials.Phone,
		"beneficiaryName":         credentials.BeneficiaryName,
		"beneficiaryPhone":        credentials.BeneficiaryPhone,
		"beneficiaryCountry":      credentials.BeneficiaryCountry,
		"beneficiaryRelationship": credentials.BeneficiaryRelationship,
	}
	jsonData, err := json.Marshal(body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error encoding JSON"})
		return
	}

	// Create the request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating request"})
		return
	}

	// Set headers
	req.Header.Set("Authorization", "Bearer ZjdkMjZhNTJmMzZkNTYwNzI0YTE5MGUwODVjM2UwMjI=")
	req.Header.Set("Content-Type", "application/json")

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error sending request"})
		return
	}
	defer resp.Body.Close()

	// Parse the response JSON into a map
	var responseJSON map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&responseJSON); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error parsing response JSON"})
		return
	}

	// Return the JSON response
	c.JSON(resp.StatusCode, responseJSON)
}

func main() {
	r := gin.Default()

	r.POST("/login", signIn)
	r.POST("/register", register)
	r.POST("/beneficiaries", fetchbeneficiary)
	r.POST("/verifypin", checkpin)
	r.POST("/checkbalance", checkbalance)
	r.POST("/transactions", transactions)
	r.POST("/withdraw", withdraw)
	r.POST("/addbeneficiary", addben)
	r.POST("/deposit", deposit)
	r.POST("/sendmoney", sendmoney)
	r.POST("/forgetpass", forgetpass)
	r.POST("/wallet2wallet", wallet2wallet)
	r.POST("/convert", convert)
	r.POST("/airtime", buyairtime)
	r.POST("/changebasecurrency", changebasecurrency)
	r.POST("/changepassword", changepassword)
	r.POST("/changepin", changepin)
	r.POST("/forex", forex)

	r.Run(":8080") // Run on port 8080
}
