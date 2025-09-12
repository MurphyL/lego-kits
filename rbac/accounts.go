package rbac

type Account interface {
	Username()
	Password()
}
