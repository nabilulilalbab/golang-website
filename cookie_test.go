package belajargolangwebsite

import (
	"fmt"
	"net/http"
	"testing"
)

func SetCookie(w http.ResponseWriter,r *http.Request)  {
  cookie := new(http.Cookie)
  cookie.Name = "X-PZN-Name"
  cookie.Value = r.URL.Query().Get("name")
  cookie.Path = "/"

  http.SetCookie(w,cookie)
  fmt.Fprintf(w,"Success create cookie")
}


func GetCookie(w http.ResponseWriter, r *http.Request)  {
  cookie, err := r.Cookie("X-PZN-Name")
  if err != nil {
    fmt.Fprintf(w,"No cookie")
  }else{
    name := cookie.Value
    fmt.Fprintf(w,"hello %v", name)
  }
}


func TestCookie(t *testing.T)  {
  mux := http.NewServeMux()
  mux.HandleFunc("/set-cookie",SetCookie)
  mux.HandleFunc("/get-cookie",GetCookie)
  server := http.Server{
    Addr: "localhost:8080",
    Handler: mux,
  }
  if err := server.ListenAndServe(); err != nil {
    panic(err)
  } 
}
