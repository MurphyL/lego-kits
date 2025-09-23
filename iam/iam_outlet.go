package iam

/* 参考文献 - https://www.woshipm.com/it/4681031.html */

type IntegrateMode string

/*
EIAM 是 Employee Identity and Access Management 的缩写，用于管理企业内部员工，主要解决员工使用的便捷性和企业管理的安全性相关问题。
CIAM 是 Customer Identity and Access Management 的缩写，用于管理企业外部客户，主要解决用户数据的打通和开发成本与标准化相关问题。
RAM 是 Resource and Access Management 的缩写，是云厂商对 IAM 的称呼，指管理企业云资源的IAM，主要用于管理云资源的访问控制。
*/

const (
	/*
		EIAM有以下特点：
		需要集成企业的云应用、本地应用
		需要集成不同的身份源
		SSO和MFA很常用
		不同企业所需的访问控制力度不同
	*/
	EIAM IntegrateMode = "eiam"
	/*
		CIAM有以下特点：
		在用户端常见到的是单点登录和授权登录
		提供通用的组件给开发者直接使用
		更强调高性能和高可用
	*/
	CIAM IntegrateMode = "ciam"
	/*
		云厂商IAM有以下特点：
		强调授权的灵活性和企业管理的安全性
		支持多种类型的账号进行认证或被调用
		一般只关注管理自家的云资源
	*/
	RAM IntegrateMode = "ram"
)

type OrgElement struct {
	Id    int64
	Name  string
	Intro string
}

type Enterprise struct {
	OrgElement
	ParentId              int64
	UnifiedSocialCreditId string // 统一社会信用代码（Unified Social Credit Identifier）
	LegalPersonName       string // 公司法人姓名
	RegisteredAddress     string // 公司注册地址
	TermOfOperation       string // 营业期限
	ContactPersonId       int64  // 联系人 Id
}

type Department struct {
	OrgElement
	ParentId       int64
	LeaderPersonId int64 // 负责人 Id
}

type ServiceProvider interface {
	SetupApp()
}
