package api

import (
	"encoding/json"
	"net/http"
)

// coin balance params (I.e. the params the API will take)
type CoinBalanceParams struct {
	Username string
}

type UserDetailsParams struct {
	Username string
}

// Coin Balance Response
type CoinBalanceResponse struct {
	// Success Code, Usually 200
	Code int

	// Account Balance
	Balance int64

	// Currency Type
	Currency string
}

// Users Response
type UserResponse struct {
	// Success Code, Usually 200
	Code int

	// Username
	Username string

	// User Full Name
	FullName string

	// User Country
	Country string

	// User Email
	Email string
}

// Error Response (Response returned when an error occurs)
type Error struct {
	// Error Code
	Code int

	// Error Message
	Message string
}

// Used to produce error response back to the user
func writeError(w http.ResponseWriter, message string, code int) {
	// Creates object using the Error Struct
	resp := Error{
		Code:    code,
		Message: message,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code) // Sets the status code for the response

	json.NewEncoder(w).Encode(resp)
}

var (
	RequestErrorHandler = func(w http.ResponseWriter, err error) {
		writeError(w, err.Error(), http.StatusBadRequest)
	}
	InternalErrorHandler = func(w http.ResponseWriter) {
		// Used to pass a generic error message back to the user (Used for internal/code errors)
		writeError(w, "An Unexpected Error Occurred.", http.StatusInternalServerError)
	}
)
