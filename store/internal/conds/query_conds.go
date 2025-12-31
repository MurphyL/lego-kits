package conds

type Query func(field string, value any) bool
