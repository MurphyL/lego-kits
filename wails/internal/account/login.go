package account

import (
	"errors"
	"fmt"
)

type LoginMethod string

type Account interface {
	ValidateLogin(ep string) bool
	EncryptPassword() (string, error)
}

// 登录
func Login(acc Account) (bool, error) {
	if ep, err := acc.EncryptPassword(); err != nil {
		return false, fmt.Errorf("密钥生成出错：%v", err.Error())
	} else {
		if acc.ValidateLogin(ep) {
			return true, nil
		}
		return false, errors.New("账号密码验证不通过")
	}
}

// 登出
func Logout(acc Account) (bool, error) {
	return false, errors.New("no impl")
}

// 注册
func Register(acc Account) (bool, error) {
	return false, errors.New("no impl")
}

// 注销
