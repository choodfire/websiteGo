package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"shoppingCart/data"
)

var usr data.Users

func mainPage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Main", r.Method)
	t, _ := template.ParseFiles("./static/html/main.html")
	t.Execute(w, nil)
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Login", r.Method)
	t, _ := template.ParseFiles("./static/html/login.html")
	t.ExecuteTemplate(w, "login.html", nil)
}

func loginResult(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Login res", r.Method)
	r.ParseForm()
	fmt.Println("Log", r.Form)

	if data.CheckLoginInfo(r.FormValue("username"), r.FormValue("password"), usr) == true {
		http.Redirect(w, r, "/account", http.StatusSeeOther)
	} else {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}
}

func register(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Reg", r.Method)
	t, _ := template.ParseFiles("./static/html/register.html")
	t.ExecuteTemplate(w, "register.html", nil)
}

func registerResult(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Reg res", r.Method)
	r.ParseForm()
	fmt.Println("Reg", r.Form)

	if data.CheckRegistrationInfo(r.FormValue("username"), 
	r.FormValue("password"), r.FormValue("password2"), usr) == true {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	} else {
		http.Redirect(w, r, "/register", http.StatusSeeOther)
	}
}

func account(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Account", r.Method)
	fmt.Fprintf(w, "This is your account!")
}

func main() {
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
