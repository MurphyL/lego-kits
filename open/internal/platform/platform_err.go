package platform

func NewStatusCode(code, phrase, desc string) *EndpointStatusCode {
	return &EndpointStatusCode{code, phrase, desc}
}

// EndpointStatusCode 接口状态码
type EndpointStatusCode struct {
	code, phrase, desc string
}

func (e *EndpointStatusCode) Code() string {
	return e.code
}

func (e *EndpointStatusCode) Phrase() string {
	return e.phrase
}

func (e *EndpointStatusCode) Desc() string {
	return e.desc
}
