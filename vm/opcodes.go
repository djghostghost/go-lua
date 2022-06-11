package vm

import "github.com/djghostghost/go-lua/api"

type OpCode struct {
	testFlag byte // operator is a test (next instruction must be a jump)
	setAFlag byte // instruction set register A
	argBMode byte // B arg mode
	argCMode byte // C arg mode
	opMode   byte //op mode
	name     string
	action   func(i Instruction, vm api.LuaVM)
}

var opcodes = []OpCode{

	{0, 1, OpArgR, OpArgN, IABC, "MOVE    ", move},
	{0, 1, OpArgK, OpArgN, IABx, "LOADK   ", loadK},
	{0, 1, OpArgN, OpArgN, IABx, "LOADKX  ", loadKx},
	{0, 1, OpArgU, OpArgU, IABC, "LOADBOOL", loadBool},
	{0, 1, OpArgU, OpArgN, IABC, "LOADNIL ", loadNil},
	{0, 1, OpArgU, OpArgN, IABC, "GETUPVAL", nil},
	{0, 1, OpArgU, OpArgK, IABC, "GETTABUP", getTabUp},
	{0, 1, OpArgR, OpArgK, IABC, "GETTABLE", getTable},
	{0, 0, OpArgK, OpArgK, IABC, "SETTABUP", nil},
	{0, 0, OpArgU, OpArgN, IABC, "SETUPVAL", nil},
	{0, 0, OpArgK, OpArgK, IABC, "SETTABLE", setTable},
	{0, 1, OpArgU, OpArgU, IABC, "NEWTABLE", newTable},
	{0, 1, OpArgR, OpArgK, IABC, "SELF    ", self},
	{0, 1, OpArgK, OpArgK, IABC, "ADD     ", add},
	{0, 1, OpArgK, OpArgK, IABC, "SUB     ", sub},
	{0, 1, OpArgK, OpArgK, IABC, "MUL     ", mul},
	{0, 1, OpArgK, OpArgK, IABC, "MOD     ", mod},
	{0, 1, OpArgK, OpArgK, IABC, "POW     ", pow},
	{0, 1, OpArgK, OpArgK, IABC, "DIV     ", div},
	{0, 1, OpArgK, OpArgK, IABC, "IDIV    ", idiv},
	{0, 1, OpArgK, OpArgK, IABC, "BAND    ", band},
	{0, 1, OpArgK, OpArgK, IABC, "BOR     ", bor},
	{0, 1, OpArgK, OpArgK, IABC, "BXOR    ", bxor},
	{0, 1, OpArgK, OpArgK, IABC, "SHL     ", shl},
	{0, 1, OpArgK, OpArgK, IABC, "SHR     ", shr},
	{0, 1, OpArgR, OpArgN, IABC, "UNM     ", unm},
	{0, 1, OpArgR, OpArgN, IABC, "BNOT    ", bnot},
	{0, 1, OpArgR, OpArgN, IABC, "NOT     ", not},
	{0, 1, OpArgR, OpArgN, IABC, "LEN     ", _len},
	{0, 1, OpArgR, OpArgR, IABC, "CONCAT  ", concat},
	{0, 0, OpArgR, OpArgN, IAsBx, "JMP     ", jmp},
	{1, 0, OpArgK, OpArgK, IABC, "EQ      ", eq},
	{1, 0, OpArgK, OpArgK, IABC, "LT      ", lt},
	{1, 0, OpArgK, OpArgK, IABC, "LE      ", le},
	{1, 0, OpArgN, OpArgU, IABC, "TEST    ", test},
	{1, 1, OpArgR, OpArgU, IABC, "TESTSET ", testSet},
	{0, 1, OpArgU, OpArgU, IABC, "CALL", call},
	{0, 1, OpArgU, OpArgU, IABC, "TAILCALL", tailCall},
	{0, 0, OpArgU, OpArgN, IABC, "RETURN", _return},
	{0, 1, OpArgR, OpArgN, IAsBx, "FORLOOP", forLoop},
	{0, 1, OpArgR, OpArgN, IAsBx, "FORPREP", forPrep},
	{0, 0, OpArgN, OpArgU, IABC, "TFORCALL", nil},
	{0, 1, OpArgR, OpArgN, IAsBx, "TFORLOOP", nil},
	{0, 0, OpArgU, OpArgU, IABC, "SETLIST ", setList},
	{0, 1, OpArgU, OpArgN, IABx, "CLOSURE ", closure},
	{0, 1, OpArgU, OpArgN, IABC, "VARARG  ", vararg},
	{0, 0, OpArgU, OpArgU, IAx, "EXTRAARG", nil},
}
