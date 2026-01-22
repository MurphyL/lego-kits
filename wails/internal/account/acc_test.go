package account

import (
	"log"
	"testing"
)

type testAcc struct {
	username string
	password string
}

func (acc *testAcc) Username() string {
	return acc.username
}

func (acc *testAcc) Password() string {
	return acc.password
}

func (acc *testAcc) ValidateLogin(ep string) bool {
	return acc.username == "luohao" && acc.password == ep
}

func (acc *testAcc) EncryptPassword() (string, error) {
	return acc.password, nil
}

func TestLogin(t *testing.T) {
	acc := testAcc{username: "luohao", password: "123456"}
	if ok, err := Login(&acc); err == nil {
		log.Println("登录完成：", ok)
	} else {
		log.Println("登录出错：", err.Error())
	}
}
