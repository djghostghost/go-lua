package state

func (s *luaState) PC() int {
	return s.pc
}

func (s *luaState) AddPC(n int) {
	s.pc += n
}

func (s *luaState) Fetch() uint32 {
	i := s.proto.Code[s.pc]
	s.pc++
	return i
}

func (s *luaState) GetConst(idx int) {
	c := s.proto.Constants[idx]
	s.stack.push(c)
}

func (s *luaState) GetRK(rk int) {
	if rk > 0xFF { // constant
		s.GetConst(rk & 0xFF)
	} else { //Register
		s.PushValue(rk + 1)
	}
}
