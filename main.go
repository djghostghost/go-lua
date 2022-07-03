package main

import (
	"github.com/djghostghost/go-lua/compiler/lexer"
	"io/ioutil"
	"os"
)

func main() {

	var luacFile string
	if len(os.Args) > 1 {
		luacFile = os.Args[1]
	} else {
		luacFile = "./lua/lexer.out"
	}

	data, err := ioutil.ReadFile(luacFile)

	if err != nil {
		panic(err)
	}

}

func testLexer(chunk, chunkName string) {
	l := lexer.NewLexer(chunk, chunkName)

	for

}

