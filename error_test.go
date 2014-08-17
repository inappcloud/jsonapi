package jsonapi_test

import (
	"errors"
	"net/http/httptest"
	"testing"

	"github.com/inappcloud/jsonapi"
)

func TestStandardError(t *testing.T) {
	w := httptest.NewRecorder()
	err := &jsonapi.StandardError{"im_a_teapot", 418, "I'm a teapot", "Hyper Text Coffee Pot Control Protocol"}
	jsonapi.Error(w, err)
	eq(t, `{"errors":[{"id":"im_a_teapot","status":418,"title":"I'm a teapot","detail":"Hyper Text Coffee Pot Control Protocol"}]}`+"\n", w.Body.String())
}

func TestError(t *testing.T) {
	w := httptest.NewRecorder()
	err := errors.New("Go Error")
	jsonapi.Error(w, err)
	eq(t, `{"errors":[{"id":"internal_server_error","status":500,"title":"Internal Server Error","detail":"Something went wrong."}]}`+"\n", w.Body.String())
}
