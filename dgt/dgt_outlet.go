package dgt

/* Dynamic Get Toolkit - 动态数据源 */

type DataSource interface {
	NewRequest(string) Request
}

type Request interface {
	Apply() (Response, error)
}

type Response interface {
	Success() bool
	Attrs() map[string]string
	Attr(string) string
	Body() string
}
