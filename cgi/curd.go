package main

import (
	"errors"
)

type ActionEnum string

const (
	ActionCreate ActionEnum = "create"
	ActionUpdate ActionEnum = "update"
	ActionRead   ActionEnum = "read"
	ActionDelete ActionEnum = "delete"
)

func NewActionResult(action ActionEnum, payload any, errs ...error) *ActionResult {
	return &ActionResult{Action: action, Payload: payload, Message: errors.Join(errs...).Error()}
}

type ActionResult struct {
	Action  ActionEnum
	Message string
	Payload any
}

type ActionHandler[T any] interface {
	Create(T) (uint, error) // 增加单条记录
	Update() (uint, error)  // 更新单条记录
	Read() (T, error)       // 单条记录查询
	Delete() (uint, error)  // 单条记录删除

}
