package platform

func NewStatus(code, phrase, desc string) *EndpointStatus {
	return &EndpointStatus{code, phrase, desc}
}

// EndpointStatus 接口状态码
type EndpointStatus struct {
	code, phrase, desc string
}

func (e *EndpointStatus) Code() string {
	return e.code
}

func (e *EndpointStatus) Phrase() string {
	return e.phrase
}

func (e *EndpointStatus) Desc() string {
	return e.desc
}
