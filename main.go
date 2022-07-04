package main

import (
	"fmt"
	"github.com/djghostghost/go-lua/compiler/lexer"
	"io/ioutil"
	"os"
)

func main() {

	var luacFile string
	if len(os.Args) > 1 {
		luacFile = os.Args[1]
	} else {
		luacFile = "./lua/hello.lua"
	}

	data, err := ioutil.ReadFile(luacFile)

	if err != nil {
		panic(err)
	}
	testLexer(string(data), luacFile)
}

func testLexer(chunk, chunkName string) {
	l := lexer.NewLexer(chunk, chunkName)

	for {
		line, kind, token := l.NextToken()
		fmt.Printf("[%2d] [%-10s] %s\n", line, kindToCategory(kind), token)
		if kind == lexer.TokenEoF {
			break
		}
	}
}

func kindToCategory(kind int) string {
	switch {
	case kind < lexer.TokenSepSemi:
		return "other"
	case kind <= lexer.TokenSepRCurly:
		return "separator"
	case kind <= lexer.TokenOpNot:
		return "operator"
	case kind <= lexer.TokenKwWhile:
		return "keyword"
	case kind == lexer.TokenIdentifier:
		return "identifier"
	case kind == lexer.TokenNumber:
		return "number"
	case kind == lexer.TokenString:
		return "string"
	default:
		return "other"
	}
}
