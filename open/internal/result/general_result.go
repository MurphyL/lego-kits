package result

func NewGeneralResult[T any](code uint, message string, payload T) *GeneralResult[T] {
	return &GeneralResult[T]{code, message, payload}
}

func NewPagingPayload[T any](total uint, records T[]) *PagingPayload[T] {
	return &PagingPayload[T]{total, records}
}

type GeneralResult[T any] struct {
	code    uint
	message string
	payload T
}

type PagingPayload[T any] struct {
	total   uint
	records T[]
}

func (g GeneralResult[T]) Code() uint {
	return g.code
}

func (g GeneralResult[T]) Message() string {
	return g.message
}

func (g GeneralResult[T]) Payload() T {
	return g.payload
}
