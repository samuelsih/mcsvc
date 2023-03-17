package main

import (
	"net/http"
)

func (c *Config) Broker(w http.ResponseWriter, r *http.Request) {
	payload := jsonResponse{
		HasError: false,
		Message: "Broker test",
	}

	c.writeJSON(w, http.StatusOK, payload)
}