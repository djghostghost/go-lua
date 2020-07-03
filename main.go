package main

import (
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
	ls.Load(data, luacFile, "b")
	ls.Call(0, 0)
}
