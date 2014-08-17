package jsonapi

import (
	"encoding/json"
	"net/http"
)

func Handler(v interface{}, next http.Handler) http.Handler {
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

func ContentTypeHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")

		if next != nil {
			next.ServeHTTP(w, r)
		}
	})
}
