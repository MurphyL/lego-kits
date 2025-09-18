package rbac

import (
	"encoding/json"
	"testing"
)

func TestName(t *testing.T) {
	tag := Tag{Name: "测试"}
	tag.Extra("hello", "value")
	data, _ := json.Marshal(tag)
	t.Log(string(data))
}
