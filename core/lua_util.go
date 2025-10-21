package core

import (
	"log"

	lua "github.com/yuin/gopher-lua"
)

func NewLuaState() *lua.LState {
	lvm := lua.NewState(lua.Options{
		CallStackSize:       120,
		MinimizeStackMemory: true,
	})
	lvm.SetGlobal("log", lvm.NewFunction(lvmPrintln))
	return lvm
}

func lvmPrintln(lvm *lua.LState) int {
	src := lvm.GetGlobal("script")
	args := lvm.CheckString(1)
	if src == nil {
		log.Println("匿名脚本 -", args)
	} else {
		log.Println(src.String(), "-", args)
	}
	return 0
}
