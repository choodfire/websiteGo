package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"html"
	"html/template"
	"log"
	"net/http"
	"shoppingCart/data"
)

var usr data.Users
var t *template.Template

func mainPage(w http.ResponseWriter, r *http.Request) {
	t.ExecuteTemplate(w, "main.html", nil)
}

func login(w http.ResponseWriter, r *http.Request) {
	t.ExecuteTemplate(w, "login.html", nil)
}

func loginResult(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	username := html.EscapeString(r.FormValue("username"))
	password := html.EscapeString(r.FormValue("password"))

	if usr.CheckLoginInfo(username, password) == true {
		http.Redirect(w, r, "/account", http.StatusSeeOther)
	} else {
		t.ExecuteTemplate(w, "login.html", "Wrong credentials")
	}
}

func register(w http.ResponseWriter, r *http.Request) {
	t.ExecuteTemplate(w, "register.html", nil)
}

func registerResult(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	username := html.EscapeString(r.FormValue("username"))
	password := html.EscapeString(r.FormValue("password"))
	password2 := html.EscapeString(r.FormValue("password2"))

	if err := usr.CheckRegistrationInfo(username, password, password2); err != nil {
		t.ExecuteTemplate(w, "register.html", err.Error())
	} else {
		usr.AddNewUser(username, password)
		t.ExecuteTemplate(w, "login.html", "Registrarion successful")
	}
}

func account(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is your account!")
}

func main() {
	t, _ = template.ParseGlob("./static/html/*.html")

	http.HandleFunc("/", mainPage)
	http.HandleFunc("/login/", login)
	http.HandleFunc("/loginResult/", loginResult)
	http.HandleFunc("/register/", register)
	http.HandleFunc("/registerResult/", registerResult)
	http.HandleFunc("/account/", account)

	err := http.ListenAndServe(":8080", nil) // устанавливаем порт для прослушивания
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
