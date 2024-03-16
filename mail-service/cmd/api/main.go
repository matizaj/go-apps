package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
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

func query(w http.ResponseWriter, r *http.Request) {
	req := r.FormValue("d")
	io.WriteString(w, req)
}

func post(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `
	<form method="get">
		<input type="text" name="q">
		<input type="submit">
	</form>
`)
}
func file(w http.ResponseWriter, r *http.Request) {
	var s string
	fmt.Println("METHOD: ", r.Method)
	if r.Method == http.MethodPost {
		// open
		f, h, err := r.FormFile("q")
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		defer f.Close()

		//fyi
		fmt.Println("file: ", f, "\nheader: ", h)

		bs, err := io.ReadAll(f)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		s = string(bs)
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `
	<form method="post" enctype="multipart/form-data">
		<input type="file" name="q">
		<input type="submit">
	</form>
`+s)
}
func seeOther(w http.ResponseWriter, r *http.Request) {
	fmt.Println("your method: ", r.Method, "\n\n")
	//w.Header().Set("Location", "/")
	//w.WriteHeader(http.StatusSeeOther)
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}
func setCookie(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("my-cookie")
	if err == http.ErrNoCookie {
		fmt.Println("noo cookie")
		cookie = &http.Cookie{
			Name:  "my-cookie",
			Value: "0",
		}
	}
	fmt.Println("after check cookie")

	counter, _ := strconv.Atoi(cookie.Value)
	fmt.Println("cookie: ", counter)

	fmt.Println("increment")
	counter++
	cookie.Value = strconv.Itoa(counter)
	fmt.Println("set new value cookie")
	http.SetCookie(w, cookie)
	io.WriteString(w, cookie.Value)
}

func readCookie(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("my-cookie")
	if err != nil {
		http.Error(w, err.Error(), http.StatusNoContent)
		return
	}

	fmt.Fprintln(w, "COOKIE: ", c)
}
func main() {
	//http.Handle("/", http.FileServer(http.Dir(".")))
	http.Handle("/resources/", http.StripPrefix("/resources", http.FileServer(http.Dir("./assets"))))
	//http.HandleFunc("/", dog)
	http.HandleFunc("/", matt)
	http.HandleFunc("/query", query)
	http.HandleFunc("/post", post)
	http.HandleFunc("/file", file)
	http.HandleFunc("/bar", seeOther)
	http.HandleFunc("/set", setCookie)
	http.HandleFunc("/read", readCookie)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	err := http.ListenAndServe(":5050", nil)
	if err != nil {
		panic(err)
	}
}
