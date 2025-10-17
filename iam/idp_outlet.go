package iam

import "github.com/MurphyL/lego-kits/iam/person"

/** 参考资料 - https://docs.authing.cn/v2/guides/connections/ **/

// 身份源提供商（Identity Provider，简称 IdP）是负责收集、存储用户身份信息，如用户名、密码等，在用户登录时负责认证用户的服务。使用外部身份服务商可以降低用户管理成本以及降低用户使用成本。

type IdentityProvider[K any, T person.GenericPerson[K]] interface {
	GetPersonInfo(K) (*T, error)
	CreateIdentity(*T) (*K, error)
}
