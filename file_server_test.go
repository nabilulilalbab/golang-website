package belajargolangwebsite

import (
	"embed"
	"io/fs"
	"net/http"
	"testing"
)




func TestFileServer(t *testing.T)  {
  directory := http.Dir("./resources")
  fileserver := http.FileServer(directory)
  mux := http.NewServeMux()
  // sebelum menggunakan StripPrefix golang akan mengakses ke resources/static/index.html
  // karena struktur folder saya hanya : resources/index.html, dengan adanya path /static golang mengangggap static sebagai folder, makanya membutuhkan StripPrefix
  mux.Handle("/static/", http.StripPrefix("/static",fileserver))
  server := http.Server{
    Addr: "localhost:8080",
    Handler: mux,
  }
  if err := server.ListenAndServe();err != nil {
    panic(err)
  }
}


//go:embed resources
var resources embed.FS

func TestFileServerEmbed(t *testing.T)  {
  // kalau menggunakan embed harus secara eksplisit mengakses /static/resources/index.`js  
  // jadi kita perlu mengakses sub direktorinya terlebih dahulu atau mengakses si resource terlebih dahulu menggunakan methode fs.Sub
  directory, _ := fs.Sub(resources, "resources") 
  fileserver := http.FileServer(http.FS(directory))
  mux := http.NewServeMux()
  mux.Handle("/static/", http.StripPrefix("/static",fileserver))
  server := http.Server{
    Addr: "localhost:8080",
    Handler: mux,
  }
  if err := server.ListenAndServe();err != nil {
    panic(err)
  }
}
