package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"shoppingCart/data"
)

var usr data.Users
var t *template.Template

func mainPage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Main", r.Method)
	t.ExecuteTemplate(w, "main.html", nil)
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Login", r.Method)
	t.ExecuteTemplate(w, "login.html", nil)
}

func loginResult(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Login res", r.Method)
	r.ParseForm()
	fmt.Println("Log", r.Form)

	if usr.CheckLoginInfo(r.FormValue("username"), r.FormValue("password")) == true {
		http.Redirect(w, r, "/account", http.StatusSeeOther)
	} else {
		// http.Redirect(w, r, "/login", http.StatusSeeOther)
		t.ExecuteTemplate(w, "login.html", "Wrong credentials")
	}
}

func register(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Reg", r.Method)
	t.ExecuteTemplate(w, "register.html", nil)
}

func registerResult(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Reg res", r.Method)
	r.ParseForm()
	fmt.Println("Reg", r.Form)

	if err := usr.CheckRegistrationInfo(r.FormValue("username"), r.FormValue("password"), r.FormValue("password2")); err == nil {
		t.ExecuteTemplate(w, "register.html", err.Error())
	} else {
		usr.AddNewUser(r.FormValue("username"), r.FormValue("password"))
		t.ExecuteTemplate(w, "login.html", "Registrarion successful")
	}
}

func account(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Account", r.Method)
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
