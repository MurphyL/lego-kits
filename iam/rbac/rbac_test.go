package rbac

import (
	"testing"
)

func TestName(t *testing.T) {
	agent := Agent[int64]{
		GetRoleById: func(i int64) *Role {
			return &Role{Name: "测试角色"}
		},
		GetUserById: func(i int64) *User {
			return &User{Name: "测试用户"}
		},
	}
	v := agent.GetUserById(1)
	t.Log(v)
}
