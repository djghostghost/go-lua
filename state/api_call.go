package state

import (
	"fmt"
	"github.com/djghostghost/go-lua/binchunk"
	"github.com/djghostghost/go-lua/vm"
)

func (s *luaState) Load(chunk []byte, chunkName, mode string) int {

	proto := binchunk.UnDump(chunk)
	c := newLuaClosure(proto)
	s.stack.push(c)
	return 0
}

func (s *luaState) Call(nArgs, nResults int) {
	val := s.stack.get(-(nArgs + 1))
	if c, ok := val.(*closure); ok {
		fmt.Printf("call %s<%d,%d>\n", c.proto.Source, c.proto.LineDefined, c.proto.LastLineDefined)
		s.callLuaClosure(nArgs, nResults, c)
	} else {
		panic("not function!")
	}
}

func (s *luaState) callLuaClosure(nArgs, nResults int, c *closure) {
	nRegs := int(c.proto.MaxStackSize)
	nParams := int(c.proto.NumParams)
	isVararg := c.proto.IsVararg == 1

	newStack := newLuaStack(nRegs + 20)
	newStack.closure = c

	funcAndArgs := s.stack.popN(nArgs + 1)
	newStack.pushN(funcAndArgs[1:], nParams)
	newStack.top = nRegs
	if nArgs > nParams && isVararg {
		newStack.varargs = funcAndArgs[nParams+1:]
	}

	s.pushLuaStack(newStack)
	s.runLuaClosure()
	s.popLuaStack()

	if nResults != 0 {
		results := newStack.popN(newStack.top - nRegs)
		s.stack.check(len(results))
		s.stack.pushN(results, nResults)
	}
}

func (s *luaState) runLuaClosure() {
	for {
		inst := vm.Instruction(s.Fetch())
		inst.Execute(s)
		if inst.OpCode() == vm.OP_RETURN {
			break
		}
	}
}
