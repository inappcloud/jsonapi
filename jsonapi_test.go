package jsonapi_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"

	"github.com/inappcloud/jsonapi"
)

func err(err *jsonapi.StandardError) string {
	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(map[string][]*jsonapi.StandardError{"errors": []*jsonapi.StandardError{err}})
	return buf.String()
}

func eq(tb testing.TB, exp, act interface{}) {
	if !reflect.DeepEqual(exp, act) {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\033[31m%s:%d:\n\n\texp: %#v\n\n\tgot: %#v\033[39m\n\n", filepath.Base(file), line, exp, act)
		tb.FailNow()
	}
}

type TestPost struct {
	Id    string `json:"id"`
	Title string `json:"title"`
}

type PostRequest struct {
	Posts []TestPost `json:"data"`
}

func TestHandler(t *testing.T) {
	body := `{"data":[{"id":"1","title":"foobar"}]}`
	w := httptest.NewRecorder()
	r, _ := http.NewRequest("POST", "/", bytes.NewBufferString(body))

	pr := new(PostRequest)
	jsonapi.Handler(pr, nil).ServeHTTP(w, r)

	eq(t, 1, len(pr.Posts))
	eq(t, "1", pr.Posts[0].Id)
	eq(t, "foobar", pr.Posts[0].Title)

	//////////////////////////////////////////////

	body = ``
	w = httptest.NewRecorder()
	r, _ = http.NewRequest("POST", "/", bytes.NewBufferString(body))

	pr = new(PostRequest)
	jsonapi.Handler(pr, nil).ServeHTTP(w, r)

	eq(t, 400, w.Code)
	eq(t, err(jsonapi.ErrBadRequest), w.Body.String())
}
