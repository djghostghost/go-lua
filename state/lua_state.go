package state

import "github.com/djghostghost/go-lua/binchunk"

type luaState struct {
	stack *luaStack
	proto *binchunk.Prototype
	pc    int
}

func New(stackSize int, prototype *binchunk.Prototype) *luaState {
	return &luaState{
		stack: newLuaStack(stackSize),
		proto: prototype,
		pc:    0,
	}
}
