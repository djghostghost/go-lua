package state

import . "github.com/djghostghost/go-lua/api"

func (s *luaState) CreateTable(nArr, nRec int) {
	t := newLuaTable(nArr, nRec)
	s.stack.push(t)
}

func (s *luaState) NewTable() {
	s.CreateTable(0, 0)
}

func (s *luaState) GetTable(idx int) LuaType {
	t := s.stack.get(idx)
	k := s.stack.pop()
	return s.getTable(t, k)
}

func (s *luaState) getTable(t, k luaValue) LuaType {
	if tbl, ok := t.(*luaTable); ok {
		v := tbl.get(k)
		s.stack.push(v)
		return typeOf(v)
	}
	panic("not a table")
}

func (s *luaState) GetField(idx int, k string) LuaType {
	t := s.stack.get(idx)
	return s.getTable(t, k)
}

func (s *luaState) GetI(idx int, i int64) LuaType {
	t := s.stack.get(idx)
	return s.getTable(t, i)
}

func (s *luaState) SetTable(idx int) {
	t := s.stack.get(idx)
	v := s.stack.pop()
	k := s.stack.pop()
	s.setTable(t, k, v)
}

func (s *luaState) setTable(t, k, v luaValue) {
	if tbl, ok := t.(*luaTable); ok {
		tbl.put(k, v)
		return
	}
	panic("not a table!")
}

func (s *luaState) SetField(idx int, k string) {
	t := s.stack.get(idx)
	v := s.stack.pop()
	s.setTable(t, k, v)
}

func (s *luaState) SetI(idx int, i int64) {
	t := s.stack.get(idx)
	v := s.stack.pop()
	s.setTable(t, i, v)
}
