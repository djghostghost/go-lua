package state

import (
	"github.com/djghostghost/go-lua/api"
)

func (s *luaState) PushNil() {
	s.stack.push(nil)
}
func (s *luaState) PushBoolean(b bool) {
	s.stack.push(b)
}

func (s *luaState) PushInteger(n int64) {
	s.stack.push(n)
}

func (s *luaState) PushNumber(n float64) {
	s.stack.push(n)
}

func (s *luaState) PushString(str string) {
	s.stack.push(str)
}

func (s *luaState) PushGoFunction(f api.GoFunction) {
	s.stack.push(newGoClosure(f))
}

func (s *luaState) PushGlobalTable() {
	global := s.registry.get(api.LUA_RIDX_GLOABLS)
	s.stack.push(global)
}

func (s *luaState) PushGoClosure(f api.GoFunction, n int) {
	closure := newGoClosure(f)
	for i := n; i > 0; i-- {
		_ = s.stack.pop()
		//closure.upvals[n-1] = &binchunk.UpValue{&val}
	}
	s.stack.push(closure)
}
