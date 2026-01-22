package wails

import (
	"errors"

	"github.com/MurphyL/lego-kits/wails/internal/account"
)

func NewApp() *WailsApp {
	return &WailsApp{}
}

type WailsApp struct {
	EnableRegistration     bool
	EnablePasswordRecovery bool
}

func (wa *WailsApp) Login(acc account.Account) (bool, error) {
	if nil == acc {
		return false, errors.New("登录信息缺失")
	}
	return account.Login(acc)
}

func (wa *WailsApp) Logout(acc account.Account) (bool, error) {
	return account.Login(acc)
}
