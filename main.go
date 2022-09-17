package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	// "time"
)

type User struct {
	name     string
	password string
}

type Users struct {
	Users []User
}

func (u *Users) AddNewUser(login string, password string) {

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

	return true
}

func mainPage(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("main.html")
	t.Execute(w, nil)
}

func login(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("login.html")
	t.Execute(w, nil)
}

func loginResult(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println("Log", r.Form)

	if checkLoginInfo(r.FormValue("username"), r.FormValue("password")) == true {
		// redirect to account page
		http.Redirect(w, r, "/account", http.StatusSeeOther)
	} else {
		// fmt.Fprintf(w, "Wrong password!")
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}
}

func register(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("register.html")
	t.ExecuteTemplate(w, "register.html", nil)
}

func registerResult(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println("Reg", r.Form)

	if checkRegistrationInfo(r.FormValue("username"), r.FormValue("password"), r.FormValue("password2")) == true {
		// fmt.Fprintf(w, "Login successfully!")
		t, _ := template.ParseFiles("login.html")
		t.ExecuteTemplate(w, "login.html", "Registration successful!")
	} else {
		// fmt.Fprintf(w, "Wrong credentials!")
		t, _ := template.ParseFiles("register.html")
		t.ExecuteTemplate(w, "register.html", "Wrong credentials!")
	}
}

func account(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is your account!")
}

func main() {
	http.HandleFunc("/", mainPage)
	http.HandleFunc("/login/", login)
	http.HandleFunc("/register/", register)
	http.HandleFunc("/registerResult/", registerResult)
	http.HandleFunc("/account/", account)

	err := http.ListenAndServe(":8080", nil) // устанавливаем порт для прослушивания
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
