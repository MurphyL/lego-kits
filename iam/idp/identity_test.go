package idp

import (
	"log"
	"testing"
)

type testAccount struct {
	username, password string
}

func (a *testAccount) ValidateLogin(token string, method LoginMethod) bool {
	switch method {
	case LoginMethodPassword:
		return a.password == token
	default:
		return false
	}
}

func (a *testAccount) EncryptPassword() (string, error) {
	return a.password, nil
}

func TestLoginByPassword(t *testing.T) {
	acc := &testAccount{username: "uname", password: "123456"}
	if ok, err := Login(acc, LoginMethodPassword); err == nil {
		log.Println("登录完成：", ok)
	} else {
		log.Panicln("登录出错：", err.Error())
	}
}
