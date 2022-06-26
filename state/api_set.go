package state

import "github.com/djghostghost/go-lua/api"

func (s *luaState) SetGlobal(name string) {
	t := s.registry.get(api.LUA_RIDX_GLOABLS)
	v := s.stack.pop()
	s.setTable(t, name, v, false)
}

func (s *luaState) Register(name string, f api.GoFunction) {
	s.PushGoFunction(f)
	s.SetGlobal(name)
}

func (s *luaState) SetMetaTable(idx int) {
	target := s.stack.get(idx)
	mtVal := s.stack.pop()

	if mtVal == nil {
		setMetatable(target, nil, s)
	} else if mt, ok := mtVal.(*luaTable); ok {
		setMetatable(target, mt, s)
	} else {
		panic("table expected")
	}
}

func (s *luaState) RawSet(idx int) {
	t := s.stack.get(idx)
	v := s.stack.pop()
	k := s.stack.pop()
	s.setTable(t, k, v, true)
}

func (s *luaState) RawSetI(idx int, i int64) {
	t := s.stack.get(idx)
	v := s.stack.pop()
	s.setTable(t, i, v, true)
}
