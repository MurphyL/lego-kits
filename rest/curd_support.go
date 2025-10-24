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
