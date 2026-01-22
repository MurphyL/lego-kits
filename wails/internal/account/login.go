package account

import (
	"errors"
)

type Account interface {
	Username() string
	Password() string
	ValidateLogin() bool
}

func Login(acc Account) (bool, error) {
	if acc.ValidateLogin() {
		return true, nil
	}
	return false, errors.New("账号密码验证不通过")
}
