package wails

import (
	"errors"
	"strings"

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
	if nil == acc || strings.TrimSpace(acc.Username()) == "" || strings.TrimSpace(acc.Password()) == "" {
		return false, errors.New("账户信息不能为空")
	}
	return account.Login(acc)
}
