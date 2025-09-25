package dgt

/* Dynamic Get Toolkit - 动态数据源 */

type DataSource interface {
	NewRequest(string) Request
}

type Request interface {
	Apply() (string, error)
}
