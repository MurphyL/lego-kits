package open

import (
	"github.com/MurphyL/lego-kits/open/internal/third_party"
)

/* 开放平台 - Self Holding */

func NewResultWithCode[T any](code uint, message string, payload T) PlatformResult[T] {
	return third_party.NewResultWithCode[T](code, message, payload)
}

func NewPagingWithCode[T any](code uint, total uint, records []T) PlatformResult[*third_party.PagingPayload[T]] {
	return third_party.NewResultWithCode[*third_party.PagingPayload[T]](code, "OK", third_party.NewPagingPayload[T](total, records))
}

type PlatformResult[T any] interface {
	Code() uint
	Message() string
	Payload() T
}
