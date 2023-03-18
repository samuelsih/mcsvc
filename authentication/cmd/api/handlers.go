package main

import (
	"fmt"
	"net/http"
)

func (c *Config) Authenticate(w http.ResponseWriter, r *http.Request) {
	var payload struct {
		Email string `json:"email"`
		Password string `json:"password"`
	} 

	err := c.readJSON(w, r, &payload)
	if err != nil {
		c.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	user, err := c.Models.User.GetByEmail(payload.Email)
	if err != nil {
		c.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	valid, err := user.PasswordMatches(payload.Password)
	if err != nil || !valid {
		c.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	result := jsonResponse {
		HasError: false,
		Message: fmt.Sprintf("login as %s", user.Email),
		Data: user,
	}

	c.writeJSON(w, http.StatusAccepted, result)
}
