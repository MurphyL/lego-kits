package idp

import (
	"errors"
	"fmt"
)

// 参考资料 - https://docs.authing.cn/v2/guides/connections/

// Provider - 身份源提供商（，简称 IdP）是负责收集、存储用户身份信息，如用户名、密码等，在用户登录时负责认证用户的服务。使用外部身份服务商可以降低用户管理成本以及降低用户使用成本。
type Provider interface {
	GetIdentity() (Identity, error)
}

// LoginMethod - 登录方式
type LoginMethod string

const (
	LoginMethodPassword     LoginMethod = "password"   // 密码登录
	LoginMethodEmail        LoginMethod = "email_code" // 邮箱验证码登录
	LoginMethodPhone        LoginMethod = "phone_code" // 手机验证码登录
	LoginMethodWechatQrcode LoginMethod = "wechat_qrcode"
	LoginMethodAlipayQrcode LoginMethod = "alipay_qrcode"
)

// Identity - 身份
type Identity interface {
	ValidateLogin(token string, method LoginMethod) bool
	EncryptPassword() (string, error)
}

// Login 登录
func Login(acc Identity, method LoginMethod) (bool, error) {
	if ep, err := acc.EncryptPassword(); err != nil {
		return false, fmt.Errorf("密钥生成出错：%v", err.Error())
	} else {
		if acc.ValidateLogin(ep, method) {
			return true, nil
		}
		return false, errors.New("账号密码验证不通过")
	}
}

// Logout 登出
func Logout(acc Identity) (bool, error) {
	return false, errors.New("no impl")
}

// Register 注册
func Register(acc Identity) (bool, error) {
	return false, errors.New("no impl")
}

// 注销
