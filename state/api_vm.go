package state

func (s *luaState) PC() int {
	return s.stack.pc
}

func (s *luaState) AddPC(n int) {
	s.stack.pc += n
}

func (s *luaState) Fetch() uint32 {
	i := s.stack.closure.proto.Code[s.stack.pc]
	s.stack.pc++
	return i
}

func (s *luaState) GetConst(idx int) {
	c := s.stack.closure.proto.Constants[idx]
	s.stack.push(c)
}

func (s *luaState) GetRK(rk int) {
	if rk > 0xFF { // constant
		s.GetConst(rk & 0xFF)
	} else { //Register
		s.PushValue(rk + 1)
	}
}

func (s *luaState) RegisterCount() int {
	return int(s.stack.closure.proto.MaxStackSize)
}

func (s *luaState) LoadVararg(n int) {
	if n < 0 {
		n = len(s.stack.varargs)
	}
	s.stack.check(n)
	s.stack.pushN(s.stack.varargs, n)
}

func (s *luaState) LoadProto(idx int) {
	proto := s.stack.closure.proto.SubPrototypes[idx]
	closure := newLuaClosure(proto)
	s.stack.push(closure)
}
