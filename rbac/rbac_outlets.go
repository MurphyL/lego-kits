package rbac

type ResourceScope string

const (
	Global  = "global"
	Org     = "org"
	OrgUser = "org_user"
	OrgRole = "org_role"
	OrgDept = "org_dept"
	OrgTag  = "org_tag"
)

type Agent[K any] interface {
	GetRoleById(K) *Role
	GetUserById(K) *User

	BindRolePerm(*Role, *Perm)
	BindRoleUser(*Role, *User)

	GetUsersByRoleId(K) []*User
	GetPermsByRoleId(K) []*Perm
}

type Role struct {
	Scope string
	Name  string
}

type User struct {
	Name string
}

type Perm struct {
	Name string
}

type Vars map[string]string

type Tag struct {
	Vars
	Name string
}

func (t Tag) Extra(key, value string) {
	if nil == t.Vars {
		t.Vars = make(Vars)
	}
	t.Vars[key] = value
}
