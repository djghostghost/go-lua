package state

import "github.com/djghostghost/go-lua/api"

type luaState struct {
	registry *luaTable
	stack    *luaStack
}

func New() *luaState {
	registry := newLuaTable(0, 0)
	registry.put(api.LUA_RIDX_GLOABLS, newLuaTable(0, 0))

	ls := &luaState{registry: registry}
	ls.pushLuaStack(newLuaStack(api.LUA_MINSTACK, ls))
	return ls
}

func (s *luaState) pushLuaStack(stack *luaStack) {
	stack.prev = s.stack
	s.stack = stack
}

func (s *luaState) popLuaStack() {
	stack := s.stack
	s.stack = stack.prev
	stack.prev = nil
}
