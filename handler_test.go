package belajargolangwebsite

import (
	"fmt"
	"net/http"
	"testing"
)

func TestHandler(t *testing.T)  {
  
  var handler http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
    _, err := fmt.Fprint(w,"hallo world")
    if err != nil {
      http.Error(w, err.Error(), http.StatusInternalServerError)
      return
    }
  }

  server := http.Server{
    Addr: "localhost:8080",
    Handler: handler,
  }
  err := server.ListenAndServe()
  if err != nil {
    panic(err)
  }

}
