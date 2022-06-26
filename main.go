package main

import (
	"fmt"
	"github.com/djghostghost/go-lua/api"
	"github.com/djghostghost/go-lua/state"
	"io/ioutil"
	"os"
)

func main() {

	var luacFile string
	if len(os.Args) > 1 {
		luacFile = os.Args[1]
	} else {
		luacFile = "./lua/luac.out"
	}

	data, err := ioutil.ReadFile(luacFile)
	if err != nil {
		panic(err)
	}
	ls := state.New()
	ls.Register("print", print)
	ls.Register("getmetatable", getMetaTable)
	ls.Register("setmetatable", setMetaTable)

	ls.Register("next", next)
	ls.Register("pairs", pairs)
	ls.Register("ipairs", iPairs)

	ls.Load(data, luacFile, "b")
	ls.Call(0, 0)
}

func print(ls api.LuaState) int {
	nArgs := ls.GetTop()
	for i := 1; i <= nArgs; i++ {
		if ls.IsBoolean(i) {
			fmt.Printf("%t", ls.ToBoolean(i))
		} else if ls.IsString(i) {
			fmt.Print(ls.ToString(i))
		} else {
			fmt.Print(ls.TypeName(ls.Type(i)))
		}

		if i < nArgs {
			fmt.Print("\t")
		}
	}
	fmt.Println()
	return 0
}

func getMetaTable(ls api.LuaState) int {
	if !ls.GetMetaTable(1) {
		ls.PushNil()
	}
	return 1
}

func setMetaTable(ls api.LuaState) int {
	ls.SetMetaTable(1)
	return 1
}

func next(ls api.LuaState) int {
	ls.SetTop(2)
	if ls.Next(1) {
		return 2
	} else {
		ls.PushNil()
		return 1
	}

}

func pairs(ls api.LuaState) int {
	ls.PushGoFunction(next)
	ls.PushValue(1)
	ls.PushNil()
	return 3
}

func iPairs(ls api.LuaState) int {
	ls.PushGoFunction(_iPairsAux)
	ls.PushValue(1)
	ls.PushInteger(0)
	return 3
}

func _iPairsAux(ls api.LuaState) int {
	i := ls.ToInteger(2) + 1
	ls.PushInteger(i)
	if ls.GetI(1, i) == api.LUA_TNIL {
		return 1
	} else {
		return 2
	}
}
