package main

import (
	"io"
	"net/http"
	"os"
)

func dog(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, "matt.jpg")
}
func matt(w http.ResponseWriter, r *http.Request) {
	f, err := os.Open("./cmd/api/matt.jpg")
	if err != nil {
		http.Error(w, "file not found :( 404", 404)
		return
	}
	defer f.Close()
	//
	//io.Copy(w, f)
	//fi, err := f.Stat()
	//if err != nil {
	//	http.Error(w, "file not found :( 404", 404)
	//	return
	//}
	//
	//http.ServeContent(w, r, f.Name(), fi.ModTime(), f)

	http.ServeFile(w, r, "./cmd/api/matt.jpg")
}

func main() {
	http.Handle("/", http.FileServer(http.Dir(".")))
	http.Handle("/resources/", http.StripPrefix("/resources", http.FileServer(http.Dir("./assets"))))
	//http.HandleFunc("/", dog)
	http.HandleFunc("/matt", matt)
	err := http.ListenAndServe(":5050", nil)
	if err != nil {
		panic("server cant start")
	}
}
