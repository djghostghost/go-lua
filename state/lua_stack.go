package state

import "github.com/djghostghost/go-lua/api"

type luaStack struct {
	slots []luaValue
	top   int

	prev    *luaStack
	closure *closure
	varargs []luaValue
	pc      int
	state   *luaState
	openuvs map[int]*upvalue
}

func newLuaStack(size int, state *luaState) *luaStack {
	return &luaStack{
		slots: make([]luaValue, size),
		top:   0,
		state: state,
	}
}

func (l *luaStack) check(n int) {
	free := len(l.slots) - l.top
	for i := free; i < n; i++ {
		l.slots = append(l.slots, nil)
	}
}

func (l *luaStack) push(val luaValue) {
	if l.top == len(l.slots) {
		panic("stack overflow!")
	}
	l.slots[l.top] = val
	l.top++
}

func (l *luaStack) pop() luaValue {
	if l.top < 1 {
		panic("stack underflow!")
	}
	l.top--
	val := l.slots[l.top]
	l.slots[l.top] = nil
	return val
}

func (l *luaStack) absIndex(idx int) int {
	if idx <= api.LUA_REGISTRY_INDEX {
		return idx
	}
	if idx > 0 {
		return idx
	}
	return idx + l.top + 1
}

func (l *luaStack) isValid(idx int) bool {
	if idx < api.LUA_REGISTRY_INDEX {
		uvIdx := api.LUA_REGISTRY_INDEX - idx - 1
		c := l.closure
		return c != nil && uvIdx < len(c.upvals)
	}

	if idx == api.LUA_REGISTRY_INDEX {
		return true
	}
	absIdx := l.absIndex(idx)
	return absIdx > 0 && absIdx <= l.top
}

func (l *luaStack) get(idx int) luaValue {
	if idx < api.LUA_REGISTRY_INDEX {
		uvIdx := api.LUA_REGISTRY_INDEX - idx - 1
		c := l.closure
		if c == nil || uvIdx >= len(c.upvals) {
			return nil
		}
		return *(c.upvals[uvIdx].val)
	}

	if idx == api.LUA_REGISTRY_INDEX {
		return l.state.registry
	}
	absIdx := l.absIndex(idx)
	if absIdx > 0 && absIdx <= l.top {
		return l.slots[absIdx-1]
	}
	return nil
}

func (l *luaStack) set(idx int, val luaValue) {
	if idx < api.LUA_REGISTRY_INDEX {
		uvIdx := api.LUA_REGISTRY_INDEX - idx - 1
		c := l.closure
		if c != nil && uvIdx < len(c.upvals) {
			*(c.upvals[uvIdx].val) = val
		}
		return
	}

	if idx == api.LUA_REGISTRY_INDEX {
		l.state.registry = val.(*luaTable)
		return
	}
	absIdx := l.absIndex(idx)
	if absIdx > 0 && absIdx <= l.top {
		l.slots[absIdx-1] = val
		return
	}
	panic("invalid index!")
}

func (l *luaStack) reverse(from, to int) {
	slots := l.slots
	for from < to {
		slots[from], slots[to] = slots[to], slots[from]
		from++
		to--
	}
}

func (l *luaStack) popN(n int) []luaValue {
	vals := make([]luaValue, n)
	for i := n - 1; i >= 0; i-- {
		vals[i] = l.pop()
	}
	return vals
}

func (l *luaStack) pushN(vals []luaValue, n int) {
	nVals := len(vals)
	if n < 0 {
		n = nVals
	}
	for i := 0; i < n; i++ {
		if i < nVals {
			l.push(vals[i])
		} else {
			l.push(nil)
		}
	}
}
