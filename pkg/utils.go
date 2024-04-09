package utils

import (
	"encoding/json"
	"net/http"
)

// ReadBody reads the request body and decodes it into the provided value.
// It takes a pointer to an http.Request and an interface{} value as parameters.
// The function returns an error if there was an issue decoding the body.
func ReadBody(r *http.Request, v interface{}) error {
	defer r.Body.Close()
	return json.NewDecoder(r.Body).Decode(v)
}

// RespondWithJSON writes a JSON response to the given http.ResponseWriter.
// It takes the HTTP status code, payload interface{}, and writes the JSON representation of the payload to the response.
// If there is an error while marshaling the payload to JSON, it returns an Internal Server Error response.
func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
