package main

import (
	"github.com/MurphyL/lego-kits/store/internal/engine"
)

func NewStore(filepath string, withOpts ...Option) (Store, error) {
	if db, err := engine.NewTinyDB(filepath); err == nil {
		for _, withOpt := range withOpts {
			withOpt(db)
		}
		return db, nil
	} else {
		return nil, err
	}
}

type Store interface {
	Collection(name string) engine.Collection
	Truncate() bool
}

// Option 选项类型
type Option func(*engine.TinyDB)
