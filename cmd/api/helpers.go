package main

import (
	"encoding/json"
	"net/http"
)

// envelope is used for composing different types of data (as key-value pairs)
// rendered into json and sent to response writer
type envelope map[string]interface{}

// writeJSON encodes data into JSON and writes it to the response writer,
// sets the content type header to application/json
func (app *application) writeJSON(w http.ResponseWriter, status int, data interface{}) error {
	jsonData, err := json.Marshal(data)

	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_, err = w.Write(jsonData)
	if err != nil {
		return err
	}

	return nil
}
