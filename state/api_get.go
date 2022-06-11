package state

import "github.com/djghostghost/go-lua/api"

func (s *luaState) GetGlobal(name string) api.LuaType {
	t := s.registry.get(api.LUA_RIDX_GLOABLS)
	return s.getTable(t, name)
}
