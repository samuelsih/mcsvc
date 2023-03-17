package main

import (
	"encoding/json"
	"net/http"
)

type jsonResponse struct {
	HasError bool `json:"has_error"`
	Message string `json:"message"`
	Data any `json:"data,omitempty"`
}

func (c *Config) Broker(w http.ResponseWriter, r *http.Request) {
	payload := jsonResponse{
		HasError: false,
		Message: "Broker test",
	}

	out, _ := json.MarshalIndent(payload, "", "\t")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	w.Write(out)
}