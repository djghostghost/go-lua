package main

import (
	"fmt"
	"github.com/djghostghost/go-lua/api"
	"github.com/djghostghost/go-lua/binchunk"
	"github.com/djghostghost/go-lua/state"
	"github.com/djghostghost/go-lua/vm"
	"io/ioutil"
	"os"
)

func main() {

	var luacFile string
	if len(os.Args) > 1 {
		luacFile = os.Args[1]
	} else {
		luacFile = "./luac.out"
	}

	data, err := ioutil.ReadFile(luacFile)
	if err != nil {
		panic(err)
	}
	prototype := binchunk.UnDump(data)
	luaMain(prototype)
}

func luaMain(prototype *binchunk.Prototype) {
	nRegs := int(prototype.MaxStackSize)
	ls := state.New(nRegs+8, prototype)
	ls.SetTop(nRegs)
	for {
		pc := ls.PC()
		inst := vm.Instruction(ls.Fetch())
		if inst.OpCode() != vm.OP_RETURN {
			inst.Execute(ls)
			fmt.Printf("[%02d] %s ", pc+1, inst.OpName())
			printStack(ls)
		} else {
			break
		}
	}
}

func printHeader(f *binchunk.Prototype) {
	funcType := "main"
	if f.LineDefined > 0 {
		funcType = "function"
	}
	varargFlag := ""

	if f.IsVararg > 0 {
		varargFlag = "+"
	}

	fmt.Printf("\n %s <%s:%d, %d> (%d instructions)\n", funcType, f.Source, f.LineDefined, f.LastLineDefined, len(f.Code))
	fmt.Printf("%d%s params, %d slots, %d upvalues, ", f.NumParams, varargFlag, f.MaxStackSize, len(f.UpValues))
	fmt.Printf("%d locals, %d constants, %d functions,\n", len(f.LocalVars), len(f.Constants), len(f.SubPrototypes))
}

func printCode(f *binchunk.Prototype) {
	for pc, c := range f.Code {
		line := "-"
		if len(f.LineInfo) > 0 {
			line = fmt.Sprintf("%d", f.LineInfo[pc])
		}

		i := vm.Instruction(c)

		fmt.Printf("\t%d\t[%s]\t%s\t", pc+1, line, i.OpName())
		fmt.Printf("\n")
	}
}

func printDetail(f *binchunk.Prototype) {
	fmt.Printf("constants (%d):\n", len(f.Constants))

	for i, k := range f.Constants {
		fmt.Printf("\t%d\t%s\n", i+1, constantToString(k))
	}

	fmt.Printf("locals (%d):\n", len(f.LocalVars))
	for i, locVar := range f.LocalVars {
		fmt.Printf("\t%d\t%s\t%d\t%d\n",
			i, locVar.VarName, locVar.StartPC+1, locVar.EndPC+1)
	}

	fmt.Printf("upvalues (%d):\n", len(f.UpValues))

	for i, upValue := range f.UpValues {
		fmt.Printf("\t%d\t%s\t%d\t%d\n", i, upValueName(f, i), upValue.InStack, upValue.Idx)
	}
}

func constantToString(k interface{}) string {
	switch k.(type) {
	case nil:
		return "nil"
	case bool:
		return fmt.Sprintf("%t", k)
	case float64:
		return fmt.Sprintf("%g", k)
	case int64:
		return fmt.Sprintf("%d", k)
	case string:
		return fmt.Sprintf("%q", k)
	default:
		return "?"
	}
}

func upValueName(f *binchunk.Prototype, idx int) string {
	if len(f.UpValueNames) > 0 {
		return f.UpValueNames[idx]
	}
	return "-"
}

func printStack(ls api.LuaState) {
	top := ls.GetTop()
	for i := 1; i <= top; i++ {
		t := ls.Type(i)
		switch t {
		case api.LUA_TBOOLEAN:
			fmt.Printf("[%t]", ls.ToBoolean(i))
		case api.LUA_TNUMBER:
			fmt.Printf("[%g]", ls.ToNumber(i))
		case api.LUA_TSTRING:
			fmt.Printf("[%q]", ls.ToString(i))
		default:
			fmt.Printf("[%s]", ls.TypeName(t))
		}
	}
	fmt.Println()
}
