package data

import (
	"errors"
	"unicode"
)

type user struct {
	name     string
	password string
}

type Users struct {
	Users []user
}

func (u *Users) AddNewUser(login string, password string) {
	u.Users = append(u.Users, user{login, password})
}

func (u Users) CheckLoginInfo(writtenLogin string, writtenPassword string) bool {
	for _, user := range u.Users {
		if user.name == writtenLogin && user.password == writtenPassword {
			return true
		}
	}

	return false
}

func (u Users) CheckRegistrationInfo(writtenLogin string, writtenPassword string, writtenPassword2 string) error {
	if writtenPassword != writtenPassword2 {
		return errors.New("Пароли не совпадают")
	}

	var loginLength, passLength, passUpper, passLower, passDigits bool

	if len(writtenLogin) > 3 && len(writtenLogin) < 21 {
		loginLength = true
	}

	if len(writtenPassword) > 5 {
		passLength = true
	}

	for _, char := range writtenPassword {
		if unicode.IsDigit(char) {
			passDigits = true
		}
		if unicode.IsUpper(char) {
			passUpper = true
		}
		if unicode.IsLower(char) {
			passLower = true
		}
	}

	for _, user := range u.Users {
		if user.name == writtenLogin {
			return errors.New("Аккаунт с таким именем уже существует")
		}
	}

	if loginLength && passLength && passUpper && passLower && passDigits == false {
		return errors.New("Имя аккаунта или пароль не соответствуют условиям")
	}

	return nil
}
