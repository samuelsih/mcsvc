package main

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

type jsonResponse struct {
	HasError bool   `json:"has_error"`
	Message  string `json:"message"`
	Data     any    `json:"data,omitempty"`
}

func (c *Config) readJSON(w http.ResponseWriter, r *http.Request, data any) error {
	var maxReadBytes int64 = 1048576 // 1mb

	r.Body = http.MaxBytesReader(w, r.Body, maxReadBytes)

	dec := json.NewDecoder(r.Body)

	defer r.Body.Close()

	err := dec.Decode(data)
	if err != nil {
		return err
	}

	err = dec.Decode(&struct{}{})
	if err != io.EOF {
		return errors.New("body must have only single JSON value")
	}

	return nil
}

func (c *Config) writeJSON(w http.ResponseWriter, status int, data any, headers ...http.Header) error {
	out, err := json.Marshal(data)
	if err != nil {
		return err
	}

	if len(headers) > 0 {
		for key, header := range headers[0] {
			w.Header()[key] = header
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	_, err = w.Write(out)
	if err != nil {
		return err
	}

	return nil
}

func (c *Config) errorJSON(w http.ResponseWriter, err error, status int) error {
	var payload jsonResponse

	payload.HasError = true
	payload.Message = err.Error()

	return c.writeJSON(w, status, payload)
}