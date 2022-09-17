package data

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

func CheckLoginInfo(writtenLogin string, writtenPassword string, usr Users) bool {
	for _, user := range usr.Users {
		if user.name == writtenLogin && user.password == writtenPassword {
			return true
		}
	}

	return false
}

func CheckRegistrationInfo(writtenLogin string, writtenPassword string,
	writtenPassword2 string, usr Users) bool {

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
