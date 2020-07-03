package state

import "github.com/djghostghost/go-lua/binchunk"

type closure struct {
	proto *binchunk.Prototype
}

func newLuaClosure(proto *binchunk.Prototype) *closure {
	return &closure{proto: proto}
}
