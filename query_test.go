package belajargolangwebsite

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func SayHello(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		fmt.Fprintf(w, "hello")
	} else {
		fmt.Fprintf(w, "hello %v", name)
	}
}

func TestQueryParamater(t *testing.T)  {
  request := httptest.NewRequest(http.MethodGet, "http://localhost/hello?name=", nil)
  recorder := httptest.NewRecorder()
  SayHello(recorder, request)
  response := recorder.Result()
  body , _ := io.ReadAll(response.Body)
  fmt.Println(string(body))
}
