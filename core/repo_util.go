package core

import (
	"log"

	"gorm.io/gorm"
)

type ID interface{ uint64 | string }

func NewRepo[T any](db *gorm.DB) *Repo[T] {
	if err := db.AutoMigrate(new(T)); err != nil {
		log.Println("自动建表出错：", err.Error())
	}
	return &Repo[T]{DB: db}
}

type Repo[T any] struct {
	*gorm.DB
}
