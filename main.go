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
