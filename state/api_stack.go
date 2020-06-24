package state

func (s *luaState) GetTop() int {
	return s.stack.top
}

func (s *luaState) AbsIndex(idx int) int {
	return s.stack.absIndex(idx)
}

func (s *luaState) CheckStack(n int) bool {
	s.stack.check(n)
	return true
}

func (s *luaState) Pop(n int) {
	for i := 0; i < n; i++ {
		s.stack.pop()
	}
}

func (s *luaState) Copy(fromIdx, toIdx int) {
	val := s.stack.get(fromIdx)
	s.stack.set(toIdx, val)
}

// 指定索引处的值推入栈顶
func (s *luaState) PushValue(idx int) {
	val := s.stack.get(idx)
	s.stack.push(val)
}

// 将栈顶弹出 然后写入指定位置
func (s *luaState) Replace(idx int) {
	val := s.stack.pop()
	s.stack.set(idx, val)
}

// 将栈顶弹出 然后插入指定位置
func (s *luaState) Insert(idx int) {
	s.Rotate(idx, -1)
	s.Pop(1)
}

func (s *luaState) Rotate(idx, n int) {
	t := s.stack.top - 1
	p := s.stack.absIndex(idx) - 1
	var m int
	if n >= 0 {
		m = t - n
	} else {
		m = p - n - 1
	}
	s.stack.reverse(p, m)
	s.stack.reverse(m+1, t)
	s.stack.reverse(p, t)
}

func (s *luaState) SetTop(idx int) {

	newTop := s.stack.absIndex(idx)
	if newTop < 0 {
		panic("stack underflow!")
	}

	n := s.stack.top - newTop
	if n > 0 {
		s.Pop(n)
	} else if n < 0 {
		for i := 0; i > n; i-- {
			s.PushNil()
		}
	}
}

func (s *luaState) PushNil() {
	s.stack.push(nil)
}
func (s *luaState) PushBoolean(b bool) {
	s.stack.push(b)
}

func (s *luaState) PushInteger(n int64) {
	s.stack.push(n)
}

func (s *luaState) PushNumber(n float64) {
	s.stack.push(n)
}

func (s *luaState) PushString(str string) {
	s.stack.push(str)
}
