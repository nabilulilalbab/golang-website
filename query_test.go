package belajargolangwebsite

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
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

func MultipleQueryParameter(w http.ResponseWriter,r *http.Request )  {
  first_name := r.URL.Query().Get("first_name")
  last_name := r.URL.Query().Get("last_name")
  fmt.Fprintf(w, "hallo %v %v", first_name, last_name)
}

func TestMultipleQueryParameter(t *testing.T)  {
  request := httptest.NewRequest(http.MethodGet, "http://localhost/hello?first_name=nabiel&last_name=albab",nil)
  record := httptest.NewRecorder()
  MultipleQueryParameter(record,request)
  response := record.Result()
  body , _ := io.ReadAll(response.Body)
  fmt.Println(string(body))
}


type myself func(w http.ResponseWriter, r *http.Request)

func RunTest(a myself, keys []string, values ...string) {
	// Validasi jumlah keys dan values
	if len(keys) != len(values) {
		panic("keys and values length mismatch")
	}

	// Membangun parameter query
	queryParams := url.Values{}
	for i, key := range keys {
		queryParams.Add(key, values[i])
	}

	// Membuat URL dengan parameter query
	baseURL := "http://localhost/hello"
  // params encode untuk conversi map dari urlvalues agar menjadi string url yang valid
	fullURL := fmt.Sprintf("%s?%s", baseURL, queryParams.Encode())

	// Membuat permintaan HTTP untuk pengujian
	request := httptest.NewRequest(http.MethodGet, fullURL, nil)
	record := httptest.NewRecorder()
	a(record, request)

	// Membaca dan mencetak respons
	response := record.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
}

func MultipleQueryParameterValues(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
  fmt.Fprintln(w, strings.Join(query["name"], " "))
  fmt.Fprintln(w, strings.Join(query["age"]," "))
}

func TestMultipleQueryParameterValues(t *testing.T) {
	// Contoh pengujian
	RunTest(MultipleQueryParameterValues, []string{"name", "name","age"}, "Nabiel", "Ulil","30")
}
