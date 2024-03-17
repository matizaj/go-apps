package main

import (
	"errors"
	"fmt"
	"github.com/matizaj/go-app/mail-service/data/models"
	"github.com/matizaj/go-app/mail-service/helpers"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
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

func session(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("session")
	if err == http.ErrNoCookie {
		fmt.Println("NO cookie session")
		sid := uuid.NewV4()
		c = &http.Cookie{
			Name:     "session",
			Value:    sid.String(),
			HttpOnly: true,
		}
	}

	var u models.User
	if username, ok := dbSession[c.Value]; ok {
		fmt.Println("user found")
		u = dbUsers[username]
		io.WriteString(w, fmt.Sprintf("already there %s", u.Username))
		return
	}
	newUser := models.User{
		Username:  "mat",
		FirstName: "mateusz",
		LastName:  "zajac",
		Password:  "example@com",
	}
	uid := "1"
	dbUsers[uid] = newUser
	dbSession[c.Value] = uid
	fmt.Println("new user %s", newUser.Username)
	fmt.Println("dbSession %s", dbSession)
	fmt.Println("dbUser %s", dbUsers)
	http.SetCookie(w, c)
}

var dbUsers = map[string]models.User{} //userid - user
var dbSession = map[string]string{}    //sid, uid
func signup(w http.ResponseWriter, r *http.Request) {
	if alreadyLoggedIn(r) {
		fmt.Println("already taken ")
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	//get user data
	user, err := helpers.ReadJson(r)
	if err != nil {
		fmt.Println("cant read user data", err)
	}
	// username taken?
	if _, ok := dbUsers[user.Username]; ok {
		http.Error(w, "username already taken", http.StatusForbidden)
		return
	}

	// create session
	sid := uuid.NewV4()
	c := &http.Cookie{
		Name:  "sid",
		Value: sid.String(),
	}

	http.SetCookie(w, c)
	dbSession[c.Value] = user.Username

	// store user in user db
	bp, err := bcrypt.GenerateFromPassword([]byte(user.Password), 12)
	if err != nil {
		http.Error(w, "cant encrypt password", http.StatusInternalServerError)
		return
	}
	user.Password = string(bp)
	dbUsers[user.Username] = *user

	// display both dbs
	fmt.Println("dbUser", dbUsers)
	fmt.Println("dbSession", dbSession)
	// redirect
	http.Redirect(w, r, "/", http.StatusSeeOther)

}

func signin(w http.ResponseWriter, r *http.Request) {
	user := getUser(r)
	payload, err := helpers.ReadJson(r)
	if err != nil {
		http.Error(w, "wrong payload", http.StatusBadRequest)
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(payload.Password))
	if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
		http.Error(w, "wrong password or username", http.StatusUnauthorized)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ok"))
}
func main() {
	//http.Handle("/", http.FileServer(http.Dir(".")))
	http.Handle("/resources/", http.StripPrefix("/resources", http.FileServer(http.Dir("./assets"))))
	//http.HandleFunc("/", dog)
	http.HandleFunc("/", index)
	http.HandleFunc("/query", query)
	http.HandleFunc("/post", post)
	http.HandleFunc("/file", file)
	http.HandleFunc("/bar", seeOther)
	http.HandleFunc("/set", setCookie)
	http.HandleFunc("/read", readCookie)
	http.HandleFunc("/session", session)
	http.HandleFunc("/signup", signup)
	http.HandleFunc("/signin", signin)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	err := http.ListenAndServe(":5050", nil)
	if err != nil {
		panic(err)
	}
}
