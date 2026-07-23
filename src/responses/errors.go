package responses

import (
	"errors"
	"fmt"
	"net/http"
)

var (
	BadRequest               = errors.New("bad request")
	UserCredentialsDnMatch   = errors.New("user credentials didn't match")
	NotAuthorizedOrExpired   = errors.New("not authorized or token expired")
	PermissionDenied         = errors.New("permission denied")
	ResourceNotFound         = errors.New("requested resource not found")
	InternalServerError      = errors.New("something went wrong")
	UnprocessableEntity      = errors.New("unprocessable entity")
	CorruptedTokenData       = errors.New("corrupted token, please relogin")
	InvalidQueryParams       = errors.New("invalid query parameters")

	// DB errors
	DBIneffectiveOperation     = errors.New("DatabaseError: failed to perform operation")
	DBFailedToFetchRecords 	   = errors.New("DatabaseError: failed to fetch records") 
	DBFailedToUpdateRecord     = errors.New("DatabaseError: failed to update record")
)

type APIError struct {
	StatusCode   int
	ErrorMessage error
}

var (
	ErrBadRequest          = APIError{StatusCode: http.StatusBadRequest, ErrorMessage: BadRequest}
	ErrNotFound            = APIError{StatusCode: http.StatusNotFound, ErrorMessage: ResourceNotFound}
	ErrInternalServerError = APIError{StatusCode: http.StatusInternalServerError, ErrorMessage: InternalServerError}
	ErrInvalidQueryParams = APIError{StatusCode: http.StatusBadRequest, ErrorMessage: InvalidQueryParams}

	ErrInvalidUserCredentials = APIError{StatusCode: http.StatusUnauthorized, ErrorMessage: UserCredentialsDnMatch}
	ErrUnAuthorized           = APIError{StatusCode: http.StatusUnauthorized, ErrorMessage: NotAuthorizedOrExpired}
	ErrForbiddenRequest       = APIError{StatusCode: http.StatusForbidden, ErrorMessage: PermissionDenied}
	ErrCourruptedToken        = APIError{StatusCode: http.StatusUnauthorized, ErrorMessage: CorruptedTokenData}
)


type StandardErrorResponse struct {
	Error string `json:"error"`
}

type GinResponseWriter interface {
	JSON(code int, obj any)
}

// WriteAPIError writes the official structured error response to the client.
func WriteAPIError(c GinResponseWriter, apiErr APIError) {
	c.JSON(apiErr.StatusCode,
		StandardErrorResponse{Error: apiErr.ErrorMessage.Error()})
}

func WriteValidationError(c GinResponseWriter, message string) {
	c.JSON(http.StatusUnprocessableEntity, 
	StandardErrorResponse{Error: fmt.Sprintf("Validation Error: %s", message)})
}


