package belajargolangwebsite

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)


func RequestHeader(w http.ResponseWriter, r *http.Request)  {
  contentType := r.Header.Get("content-type")
  fmt.Fprint(w,contentType)
}

func TestRequestHeader(t *testing.T)  {
  request := httptest.NewRequest(http.MethodPost, "http://localhost",nil)
  request.Header.Add("content-type","application/json")
  recorder := httptest.NewRecorder()
  RequestHeader(recorder,request)
  response := recorder.Result()
  body, _ := io.ReadAll(response.Body)
  fmt.Println(string(body))
}


func ResponseHeader(w http.ResponseWriter, r *http.Request)  {
  w.Header().Add("x-powered-by","nabil ulil albab")
  fmt.Fprint(w, "ok")
}
func TestResponseHeader(t *testing.T)  {
  request := httptest.NewRequest(http.MethodPost,"http://localhost",nil)
  recorder := httptest.NewRecorder()
  ResponseHeader(recorder,request)
  response := recorder.Result()
  body, _ := io.ReadAll(response.Body)
  fmt.Println(string(body))
  fmt.Println(response.Header.Get("x-powered-by"))
}
