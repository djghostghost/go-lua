package state

func (s *luaState) Len(idx int) {
	val := s.stack.get(idx)
	if str, ok := val.(string); ok {
		s.stack.push(int64(len(str)))
	} else if result, ok := callMetaMethod(val, val, "__len", s); ok {
		s.stack.push(result)
	} else if t, ok := val.(*luaTable); ok {
		s.stack.push(int64(t.len()))
	} else {
		panic("length error!")
	}
}

func (s *luaState) Concat(n int) {
	if n == 0 {
		s.stack.push("")
	} else if n >= 2 {
		for i := 1; i < n; i++ {

			if s.IsString(-1) && s.IsString(-2) {
				s2 := s.ToString(-1)
				s1 := s.ToString(-2)
				s.stack.pop()
				s.stack.pop()
				s.stack.push(s1 + s2)

			} else {
				b := s.stack.pop()
				a := s.stack.pop()

				if result, ok := callMetaMethod(a, b, "__concat", s); ok {
					s.stack.push(result)
					continue
				}
				panic("concatenation error!")
			}
		}
	}
}

func (s *luaState) RawLen(idx int) uint {
	val := s.stack.get(idx)
	switch x := val.(type) {
	case string:
		return uint(len(x))
	case *luaTable:
		return uint(x.len())
	default:
		return 0
	}

}
