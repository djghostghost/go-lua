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
	return s.getTable(t, k, false)
}

func (s *luaState) getTable(t, k luaValue, raw bool) LuaType {
	if tbl, ok := t.(*luaTable); ok {
		v := tbl.get(k)
		if raw || v != nil || !tbl.hasMetaField("__index") {
			s.stack.push(v)
			return typeOf(v)
		}
	}
	if !raw {
		if mf := getMetaField(t, "__index", s); mf != nil {
			switch x := mf.(type) {
			case *luaTable:
				return s.getTable(x, k, false)
			case *closure:
				s.stack.push(mf)
				s.stack.push(t)
				s.stack.push(k)
				s.Call(2, 1)
				v := s.stack.get(-1)
				return typeOf(v)
			}
		}
	}
	panic("not a table")
}

func (s *luaState) GetField(idx int, k string) LuaType {
	t := s.stack.get(idx)
	return s.getTable(t, k, false)
}

func (s *luaState) GetI(idx int, i int64) LuaType {
	t := s.stack.get(idx)
	return s.getTable(t, i, false)
}

func (s *luaState) SetTable(idx int) {
	t := s.stack.get(idx)
	v := s.stack.pop()
	k := s.stack.pop()
	s.setTable(t, k, v, false)
}

func (s *luaState) setTable(t, k, v luaValue, raw bool) {
	if tbl, ok := t.(*luaTable); ok {
		if raw || tbl.get(k) != nil || !tbl.hasMetaField("__newindex") {
			tbl.put(k, v)
			return
		}
	}
	if !raw {
		if mf := getMetaField(t, "__newindex", s); mf != nil {
			switch x := mf.(type) {
			case *luaTable:
				s.setTable(x, k, v, false)
				return
			case *closure:
				s.stack.push(mf)
				s.stack.push(t)
				s.stack.push(k)
				s.stack.push(v)
				s.Call(3, 0)
				return
			}
		}
	}
	panic("not a table!")
}

func (s *luaState) SetField(idx int, k string) {
	t := s.stack.get(idx)
	v := s.stack.pop()
	s.setTable(t, k, v, false)
}

func (s *luaState) SetI(idx int, i int64) {
	t := s.stack.get(idx)
	v := s.stack.pop()
	s.setTable(t, i, v, false)
}

func (s *luaState) GetMetaTable(idx int) bool {
	val := s.stack.get(idx)
	if mt := getMetaTable(val, s); mt != nil {
		s.stack.push(mt)
		return true
	} else {
		return false
	}
}
