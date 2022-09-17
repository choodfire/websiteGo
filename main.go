package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type User struct {
	name string
	password string
}

type Users struct {
	Users []User
}

var usr Users

func checkAccount(writtenLogin string, writtenPassword string) bool {
	for _, user := range usr.Users {
		if user.name == writtenLogin && user.password == writtenPassword {
			return true
		}
	}

	return false
}

func checkRegistration(writtenLogin string, writtenPassword string, 
	writtenPassword2 string) bool {
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
	fmt.Println("Main Метод:", r.Method) // получаем информацию о методе запроса

	t, _ := template.ParseFiles("main.html")
	t.Execute(w, nil)
}

func loginPage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Login Метод:", r.Method) // получаем информацию о методе запроса

	switch r.Method {
	case "GET":
		t, _ := template.ParseFiles("login.html")
		t.Execute(w, nil)
	case "POST":
		r.ParseForm()

		if checkAccount(r.FormValue("username"), r.FormValue("password")) == true {
			// redirect to account page 
			// http.Redirect(w, r, "/account/", http.StatusSeeOther)
		} else {
			fmt.Fprintf(w, "Wrong password!")
			// http.Redirect(w, r, "/login/", http.StatusSeeOther)
		}

	default:
		fmt.Fprintf(w, "Поддерживаются только GET и POST методы")
	}
}

func registerPage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Register Метод:", r.Method) // получаем информацию о методе запроса

	switch r.Method {
	case "GET":
		t, _ := template.ParseFiles("register.html")
		t.Execute(w, nil)
	case "POST":
		r.ParseForm()
				
		if checkRegistration(r.FormValue("username"), r.FormValue("password"), r.FormValue("password2")) == true {
			// redirect to account page 
			// http.Redirect(w, r, "/account/", http.StatusSeeOther)
		} else {
			fmt.Fprintf(w, "Wrong credentials!")
			// http.Redirect(w, r, "/register/", http.StatusSeeOther)
		}
		fmt.Println("went POST")

	default:
		fmt.Fprintf(w, "Поддерживаются только GET и POST методы")
	}
}

func accountPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is yout account!")
}

func main() {
	http.HandleFunc("/", mainPage)
	http.HandleFunc("/login", loginPage)
	http.HandleFunc("/register", registerPage)
	http.HandleFunc("/account", accountPage)

	err := http.ListenAndServe(":8080", nil) // устанавливаем порт для прослушивания
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}