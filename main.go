package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type User struct {
	name     string
	password string
}

type Users struct {
	Users []User
}

func (u *Users) AddNewUser(login string, password string) {
	u.Users = append(u.Users, User{login, password})
}

var usr Users

func checkLoginInfo(writtenLogin string, writtenPassword string) bool {
	fmt.Println(usr)
	for _, user := range usr.Users {
		if user.name == writtenLogin && user.password == writtenPassword {
			return true
		}
	}

	return false
}

func checkRegistrationInfo(writtenLogin string, writtenPassword string,
	writtenPassword2 string) bool {
	fmt.Println(usr)
	if writtenPassword != writtenPassword2 {
		return false
	}

	for _, user := range usr.Users {
		if user.name == writtenLogin {
			return false
		}
	}

	usr.AddNewUser(writtenLogin, writtenPassword)

	return true
}

func mainPage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Main", r.Method)
	t, _ := template.ParseFiles("main.html")
	t.Execute(w, nil)
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Login", r.Method)
	t, _ := template.ParseFiles("login.html")
	t.ExecuteTemplate(w, "login.html", nil)
}

func loginResult(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Login res", r.Method)
	r.ParseForm()
	fmt.Println("Log", r.Form)

	if checkLoginInfo(r.FormValue("username"), r.FormValue("password")) == true {
		http.Redirect(w, r, "/account", http.StatusSeeOther)
	} else {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}
}

func register(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Reg", r.Method)
	t, _ := template.ParseFiles("register.html")
	t.ExecuteTemplate(w, "register.html", nil)
}

func registerResult(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Reg res", r.Method)
	r.ParseForm()
	fmt.Println("Reg", r.Form)

	if checkRegistrationInfo(r.FormValue("username"), 
	r.FormValue("password"), r.FormValue("password2")) == true {
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
