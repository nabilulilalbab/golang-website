package belajargolangwebsite

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func RequestHeader(w http.ResponseWriter, r *http.Request) {
	contentType := r.Header.Get("content-type")
	fmt.Fprint(w, contentType)
}

func TestRequestHeader(t *testing.T) {
	request := httptest.NewRequest(http.MethodPost, "http://localhost", nil)
	request.Header.Add("content-type", "application/json")
	recorder := httptest.NewRecorder()
	RequestHeader(recorder, request)
	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
}

func ResponseHeader(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("x-powered-by", "nabil ulil albab")
	fmt.Fprint(w, "ok")
}

func TestResponseHeader(t *testing.T) {
	request := httptest.NewRequest(http.MethodPost, "http://localhost", nil)
	recorder := httptest.NewRecorder()
	ResponseHeader(recorder, request)
	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
	fmt.Println(response.Header.Get("x-powered-by"))
}

// Client Request:
// [Request Headers] --> Server
//
// Server Response:
// Server --> [Response Headers] --> Client
//
// Request Header:
//
// Dikirim oleh client ke server
// Berisi informasi tentang request
// Contoh: Content-Type, Authorization, Accept
// Dibaca di sisi server menggunakan r.Header
// Response Header:
//
// Dikirim oleh server ke client
// Berisi metadata tentang response
// Contoh: Content-Type, X-Powered-By, Cache-Control
// Ditulis di sisi server menggunakan w.Header()

func RequestHeader2(w http.ResponseWriter, r *http.Request) {
	contentType := r.Header.Get("content-type")
	fmt.Fprint(w, contentType)
}

func TestRequestHeader2(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost", nil)
	recorder := httptest.NewRecorder()
	request.Header.Add("content-type", "application/json")
	RequestHeader2(recorder, request)
	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
}

func ResponseHeader2(w http.ResponseWriter, r *http.Request)  {
  w.Header().Add("content-type","application/json")
  fmt.Fprint(w, "ok")
}

func TestResponseHeader2(t *testing.T)  {
  request := httptest.NewRequest(http.MethodPost, "http://localhost",nil)
  record := httptest.NewRecorder()
  ResponseHeader2(record,request)
  response := record.Result()
  body ,_ := io.ReadAll(response.Body)
  fmt.Println(string(body))
  fmt.Println(response.Header.Get("content-type"))
}



func RequestHeader3(w http.ResponseWriter, r *http.Request) {
  contentType := r.Header.Get("content-type")
  fmt.Fprint(w, contentType)
}

func TestRequestHeader3(t *testing.T) {
  request := httptest.NewRequest(http.MethodGet,"http://localhost",nil)
  request.Header.Add("content-type","application/json")
  recorder := httptest.NewRecorder()
  RequestHeader3(recorder,request)
  response := recorder.Result()
  body, _ := io.ReadAll(response.Body)
  fmt.Println(string(body))
}

func ResponseHeader3(w http.ResponseWriter, r *http.Request) {
  w.Header().Add("content-type","application/json")
  fmt.Fprint(w,"ok")
}

func TestResponseHeader3(t *testing.T) {
  request := httptest.NewRequest(http.MethodPost,"http://localhost",nil)
  recorder := httptest.NewRecorder()
  ResponseHeader3(recorder,request)
  response := recorder.Result()
  body , _ := io.ReadAll(response.Body)
  fmt.Println(string(body))
   fmt.Println(response.Header.Get("content-type"))
}




