package state

import "github.com/djghostghost/go-lua/api"

func (s *luaState) GetGlobal(name string) api.LuaType {
	t := s.registry.get(api.LUA_RIDX_GLOABLS)
	return s.getTable(t, name, false)
}

func (s *luaState) RawGet(idx int) api.LuaType {
	t := s.stack.get(idx)
	k := s.stack.pop()
	return s.getTable(t, k, true)
}

func (s *luaState) RawGetI(idx int, i int64) api.LuaType {
	t := s.stack.get(idx)
	return s.getTable(t, i, true)
}
