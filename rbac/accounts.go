package rbac

type LoginAccount interface {
	Username()
	Password()
}
