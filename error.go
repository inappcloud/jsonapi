package jsonapi

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

// Error writes the specified error to w.
// If error is a Go standard error, it will write a ErrInternalServer (500) and log the error.
func Error(w http.ResponseWriter, err error) {
	var strErr *StandardError

	switch v := err.(type) {
	case *StandardError:
		strErr = v
	default:
		log.Print(err)
		strErr = ErrInternalServer
	}

	w.WriteHeader(strErr.Status)
	json.NewEncoder(w).Encode(map[string][]*StandardError{"errors": []*StandardError{strErr}})
}

// NotFound is a substitute of the standard NotFoud handler in Goji, but it writes JSON instead of plain responses.
func NotFound(w http.ResponseWriter, r *http.Request) {
	ContentTypeHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(404)
		Error(w, ErrNotFound)
	})).ServeHTTP(w, r)
}

type StandardError struct {
	Id     string `json:"id"`
	Status int    `json:"status"`
	Title  string `json:"title"`
	Detail string `json:"detail"`
}

func (err StandardError) Error() string {
	return err.Detail
}

var (
	ErrBadRequest     = &StandardError{"bad_request", 400, "Bad request", "Request body is not well-formed. It must be JSON."}
	ErrInternalServer = &StandardError{"internal_server_error", 500, "Internal Server Error", "Something went wrong."}
	ErrNoData         = &StandardError{"no_data", 422, "No data", `Key "data" in the top level of the JSON document is missing or contains no data`}
	ErrUnauthorized   = &StandardError{"unauthorized", 401, "Unauthorized", "Access token is invalid."}
	ErrNotFound       = &StandardError{"not_found", 404, "Not found", "Route not found."}
)

func ErrInvalidParams(errors ...string) *StandardError {
	return &StandardError{
		"invalid_params",
		422,
		"Invalid params",
		fmt.Sprintf(`Key "data" in the top level of the JSON document contains invalid params: %s`, strings.Join(errors, ", ")),
	}
}

func ErrCollectionNotFound(coll string) *StandardError {
	return &StandardError{
		"collection_not_found",
		404,
		"Collection not found",
		fmt.Sprintf(`Collection named "%s" has not been found.`, coll),
	}
}

func ErrResourceNotFound(coll string, id string) *StandardError {
	return &StandardError{
		"key_not_found",
		404,
		"Key not found",
		fmt.Sprintf(`Key "%s" in bucket named "%s" has not been found.`, coll, id),
	}
}
