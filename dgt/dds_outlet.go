package dgt

import (
	"errors"
)

/* Dynamic Get Toolkit - 动态数据源 */

type Request interface {
	Exec() (string, error)
}

func Apply(request Request) (string, error) {
	var err error
	var ret string
	if nil == request {
		err = errors.New("请求参数为空")
	} else {
		ret, err = request.Exec()
	}
	return ret, err
}
