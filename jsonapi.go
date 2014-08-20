// Package jsonapi provides various HTTP handlers to handle JSON API.
package jsonapi

import (
	"encoding/json"
	"net/http"
)

// BodyParserHandler decodes request JSON body to v. If decoding was unsuccessful,
// it writes ErrBadRequest (400) to w.
func BodyParserHandler(v interface{}, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := json.NewDecoder(r.Body).Decode(v)

		if err != nil {
			Error(w, ErrBadRequest)
			return
		}

		if next != nil {
			next.ServeHTTP(w, r)
		}
	})
}

// ContentTypeHandler adds Content-Type header to the response.
func ContentTypeHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")

		if next != nil {
			next.ServeHTTP(w, r)
		}
	})
}
