package state

func (s *luaState) CreateTable(nArr, nRec int) {
	t := newLuaTable(nArr, nRec)
	s.stack.push(t)
}

func (s *luaState) NewTable() {
	s.CreateTable(0, 0)
}
