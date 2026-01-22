package account

import (
	"errors"
)

type Account interface {
	ValidateLogin() bool
}

func Login(acc Account) (bool, error) {
	if acc.ValidateLogin() {
		return true, nil
	}
	return false, errors.New("账号密码验证不通过")
}

func Logout(acc Account) (bool, error) {
	return false, errors.New("no impl")
}
