package rbac

type ResourceScope string

const (
	Global = "global"
)

type Role struct {
	Scope ResourceScope
	Name  string
}

type User struct {
	Scope ResourceScope
	Name  string
}

type Perm struct {
	Scope ResourceScope
	Name  string
}

type GetById[K any, R Role | User | Perm] func(K) *R

type Agent[K any] struct {
	GetRoleById GetById[K, Role]
	GetUserById GetById[K, User]
}
