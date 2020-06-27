package state

func (s *luaState) Len(idx int) {
	val := s.stack.get(idx)
	if str, ok := val.(string); ok {
		s.stack.push(int64(len(str)))
	} else {
		panic("length error!")
	}
}
