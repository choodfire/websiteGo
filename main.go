package main

import (
	"fmt"
	"html/template"
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

func mainPage(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("main.html")
	t.Execute(w, nil)
}

func checkAccount() {}

func loginPage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Метод:", r.Method) // получаем информацию о методе запроса

	switch r.Method {
	case "GET":
		t, _ := template.ParseFiles("login.html")
		t.Execute(w, nil)
	case "POST":
		r.ParseForm()
		
		fmt.Println("Пользователь:", r.Form["username"])
		fmt.Println("Пароль:", r.Form["password"])

		

	default:
		fmt.Fprintf(w, "Поддерживаются только GET и POST методы")
	}
}

func registerPage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Метод:", r.Method) // получаем информацию о методе запроса

	switch r.Method {
	case "GET":
		t, _ := template.ParseFiles("register.html")
		t.Execute(w, nil)
	case "POST":
		r.ParseForm()
		
		fmt.Println("Пользователь:", r.Form["username"])
		fmt.Println("Пароль:", r.Form["password"])
		fmt.Println("Возраст:", r.Form["age"])
		fmt.Println("Фрукт:", r.Form["fruit"])
		fmt.Println("Интересы:", r.Form["interest"])
		//fmt.Println("Интересы:", r.FormValue("interest"))
	default:
		fmt.Fprintf(w, "Поддерживаются только GET и POST методы")
	}
}

func handlers() {
	http.HandleFunc("/", mainPage)
	http.HandleFunc("/register/", registerPage)
	http.HandleFunc("/login/", loginPage)

	http.ListenAndServe(":8080", nil)
}

func main() {

	handlers()
}