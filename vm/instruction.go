package vm

import (
	"github.com/djghostghost/go-lua/api"
)

type Instruction uint32

func (i Instruction) OpCode() int {
	return int(i & 0x3F)
}

func (i Instruction) ABC() (a, b, c int) {
	a = int((i >> 6) & 0xFF)
	c = int((i >> 14) & 0x1FF)
	b = int((i >> 23) & 0x1FF)
	return
}

func (i Instruction) ABx() (a, bx int) {
	a = int((i >> 6) & 0xFF)
	bx = int(i >> 14)
	return
}

func (i Instruction) AsBx() (a, sbx int) {
	a, bx := i.ABx()
	return a, bx - MAXARG_sBx
}

func (i Instruction) Ax() int {
	return int(i >> 6)
}

func (i Instruction) OpName() string {
	return opcodes[i.OpCode()].name
}

func (i Instruction) OpMode() byte {
	return opcodes[i.OpCode()].opMode
}

func (i Instruction) BMode() byte {
	return opcodes[i.OpCode()].argBMode
}

func (i Instruction) CMode() byte {
	return opcodes[i.OpCode()].argCMode
}

func (i Instruction) Execute(vm api.LuaVM) {
	op := opcodes[i.OpCode()]
	//fmt.Printf("Execute code %s\n", op.name)
	action := op.action

	if action != nil {
		action(i, vm)
	} else {
		panic(i.OpName())
	}
}
