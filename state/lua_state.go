package state

type luaState struct {
	registry *luaTable
	stack    *luaStack
}

func New() *luaState {
	return &luaState{
		stack: newLuaStack(20),
	}
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
