package belajargolangwebsite

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func ResponseCode(w http.ResponseWriter, r *http.Request)  {
  name := r.URL.Query().Get("name")

  if name == "" {
    w.WriteHeader(http.StatusBadRequest)
    fmt.Fprint(w, "name is empty")
  }else {
    fmt.Fprintf(w,"hello %v", name)
  }
}



func TestResponseCode(t *testing.T)  {
  request := httptest.NewRequest(http.MethodGet, "http://localhost", nil)
  recorder := httptest.NewRecorder()

  ResponseCode(recorder,request)

  response := recorder.Result()
  body, _ := io.ReadAll(response.Body)
  fmt.Println(response.StatusCode)
  fmt.Println(response.Status)
  fmt.Println(string(body))
}


func TestResponseCodeValid(t *testing.T)  {
  request := httptest.NewRequest(http.MethodGet, "http://localhost/?name=eko", nil)
  recorder := httptest.NewRecorder()

  ResponseCode(recorder,request)

  response := recorder.Result()
  body, _ := io.ReadAll(response.Body)
  fmt.Println(response.StatusCode)
  fmt.Println(response.Status)
  fmt.Println(string(body))
}
