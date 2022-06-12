package state

import (
	"github.com/djghostghost/go-lua/api"
	"github.com/djghostghost/go-lua/binchunk"
	"github.com/djghostghost/go-lua/vm"
)

func (s *luaState) Load(chunk []byte, chunkName, mode string) int {
	proto := binchunk.UnDump(chunk)
	c := newLuaClosure(proto)
	s.stack.push(c)
	if len(proto.UpValues) > 0 {
		env := s.registry.get(api.LUA_RIDX_GLOABLS)
		c.upvals[0] = &upvalue{&env}
	}
	return 0
}

func (s *luaState) Call(nArgs, nResults int) {
	val := s.stack.get(-(nArgs + 1))
	if c, ok := val.(*closure); ok {
		if c.proto != nil {
			s.callLuaClosure(nArgs, nResults, c)
		} else {
			s.callGoClosure(nArgs, nResults, c)
		}

	} else {
		panic("not function!")
	}
}

func (s *luaState) callLuaClosure(nArgs, nResults int, c *closure) {
	nRegs := int(c.proto.MaxStackSize)
	nParams := int(c.proto.NumParams)
	isVararg := c.proto.IsVararg == 1

	newStack := newLuaStack(nRegs+api.LUA_MINSTACK, s)
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

func (s *luaState) callGoClosure(nArgs, nResults int, c *closure) {
	newStack := newLuaStack(nArgs+20, s)
	newStack.closure = c

	args := s.stack.popN(nArgs)
	newStack.pushN(args, nArgs)
	s.stack.pop()

	s.pushLuaStack(newStack)
	r := c.goFunc(s)
	s.popLuaStack()

	if nResults != 0 {
		results := newStack.popN(r)

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
