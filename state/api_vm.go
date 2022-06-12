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
	stack := s.stack
	proto := stack.closure.proto.SubPrototypes[idx]
	closure := newLuaClosure(proto)
	stack.push(closure)

	for i, uvInfo := range proto.UpValues {
		uvIdx := int(uvInfo.Idx)
		if uvInfo.InStack == 1 {
			if stack.openuvs == nil {
				stack.openuvs = map[int]*upvalue{}
			}
			if openuv, found := stack.openuvs[uvIdx]; found {
				closure.upvals[i] = openuv
			} else {
				closure.upvals[i] = &upvalue{&stack.slots[uvIdx]}
				stack.openuvs[uvIdx] = closure.upvals[i]
			}
		} else {
			closure.upvals[i] = stack.closure.upvals[uvIdx]
		}
	}
}

func (s *luaState) CloseUpvalues(a int) {
	for i, openuv := range s.stack.openuvs {
		if i >= a-1 {
			val := *openuv.val
			openuv.val = &val
			delete(s.stack.openuvs, i)
		}
	}
}
