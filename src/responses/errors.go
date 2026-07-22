package responses

import (
	http "net/http"
)


type APIError struct {
	StatusCode int
	ErrorMessage string
}

var (
	ErrBadRequest APIError = APIError{http.StatusBadRequest, "bad request"}

	// Auth
	ErrUnAuthorized APIError = APIError{http.StatusUnauthorized, "not authorized or token expired"}
	ErrInvalidUserCredentials APIError = APIError{http.StatusUnauthorized, "User credentials didn't match"}
	ErrInvalidOrExpiredToken APIError = APIError{http.StatusUnauthorized, "Invalid or expired token"}
	ErrForbiddenRequest APIError = APIError{http.StatusForbidden, "permission denied"}
	ErrNotFound APIError = APIError{http.StatusNotFound, "requested rosource not found"}
	ErrInternalServerError APIError = APIError{http.StatusInternalServerError, "something went wrong"}
)