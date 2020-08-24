package user

import "blogbk/model"

func Get(u *model.SUser) error {
	return u.GetByUsername()
}

func Auth(u *model.SUser) (bool, error) {
	return u.Auth()
}

func Logout(u *model.SUser) (bool, error) {
	return u.Logout()
}
