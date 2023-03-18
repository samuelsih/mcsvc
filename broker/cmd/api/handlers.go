package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
)

type RequestPayload struct {
	Action string      `json:"action"`
	Auth   AuthPayload `json:"auth,omitempty"`
}

type AuthPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (c *Config) Broker(w http.ResponseWriter, r *http.Request) {
	payload := jsonResponse{
		HasError: false,
		Message:  "Broker test",
	}

	c.writeJSON(w, http.StatusOK, payload)
}

func (c *Config) HandleSubmission(w http.ResponseWriter, r *http.Request) {
	var payload RequestPayload

	err := c.readJSON(w, r, &payload)
	if err != nil {
		c.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	switch payload.Action {
	case "auth":
		c.authenticate(w, payload.Auth)
	default:
		c.errorJSON(w, errors.New("unknown action"), http.StatusBadRequest)
		return
	}
}

func (c *Config) authenticate(w http.ResponseWriter, payload AuthPayload) {
	jsonData, _ := json.MarshalIndent(payload, "", "\t")
	request, err := http.NewRequest(http.MethodPost, "http://auth-service/authenticate", bytes.NewBuffer(jsonData))
	if err != nil {
		c.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		c.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	defer response.Body.Close()

	if response.StatusCode == http.StatusUnauthorized {
		c.errorJSON(w, errors.New("invalid credentials"), http.StatusUnauthorized)
		return
	}

	if response.StatusCode != http.StatusAccepted {
		c.errorJSON(w, err, http.StatusInternalServerError)
		return
	}

	var jsonFromService jsonResponse

	err = json.NewDecoder(response.Body).Decode(&jsonFromService)
	if response.StatusCode != http.StatusAccepted {
		c.errorJSON(w, err, http.StatusInternalServerError)
		return
	}

	if jsonFromService.HasError {
		c.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	resultPayload := jsonResponse{}
	resultPayload.HasError = false
	resultPayload.Message = "Authenticated!"
	resultPayload.Data = jsonFromService.Data

	c.writeJSON(w, http.StatusAccepted, resultPayload)
}
