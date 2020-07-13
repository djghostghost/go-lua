package state

import (
	"github.com/djghostghost/go-lua/api"
	"github.com/djghostghost/go-lua/binchunk"
)

type closure struct {
	proto  *binchunk.Prototype
	goFunc api.GoFunction
}

func newLuaClosure(proto *binchunk.Prototype) *closure {
	return &closure{proto: proto}
}

func newGoClosure(f api.GoFunction) *closure {
	return &closure{goFunc: f}
}
