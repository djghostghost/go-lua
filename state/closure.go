package state

import (
	"github.com/djghostghost/go-lua/api"
	"github.com/djghostghost/go-lua/binchunk"
)

type upvalue struct {
	val *luaValue
}

type closure struct {
	proto  *binchunk.Prototype
	goFunc api.GoFunction
	upvals []*upvalue
}

func newLuaClosure(proto *binchunk.Prototype) *closure {
	c := &closure{proto: proto}
	if nUpvals := len(proto.UpValues); nUpvals > 0 {
		c.upvals = make([]*upvalue, nUpvals)
	}
	return c
}

func newGoClosure(f api.GoFunction, nUpvals int) *closure {
	c := &closure{goFunc: f}
	if nUpvals > 0 {
		c.upvals = make([]*upvalue, nUpvals)
	}
	return &closure{goFunc: f}
}
