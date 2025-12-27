package ipd

// 参考资料 - https://docs.authing.cn/v2/guides/connections/

// IdentityProvider - 身份源提供商（，简称 IdP）是负责收集、存储用户身份信息，如用户名、密码等，在用户登录时负责认证用户的服务。使用外部身份服务商可以降低用户管理成本以及降低用户使用成本。
type IdentityProvider interface {
	GetIdentity() (Identity, error)
}

// Identity - 身份
type Identity interface {
	DisplayName() string
}
