package state

import "github.com/djghostghost/go-lua/api"

func (s *luaState) SetGlobal(name string) {
	t := s.registry.get(api.LUA_RIDX_GLOABLS)
	v := s.stack.pop()
	s.setTable(t, name, v)
}

func (s *luaState) Register(name string, f api.GoFunction) {
	s.PushGoFunction(f)
	s.SetGlobal(name)
}
