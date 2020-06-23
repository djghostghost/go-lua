package vm

type OpCode struct {
	testFlag byte // operator is a test (next instruction must be a jump)
	setAFlag byte // instruction set register A
	argBMode byte // B arg mode
	argCMode byte // C arg mode
	opMode   byte //op mode
	name     string
}

var opcodes = []OpCode{

	{0, 1, OpArgR, OpArgN, IABC, "MOVE    "},
	{0, 1, OpArgK, OpArgN, IABx, "LOADK   "},
	{0, 1, OpArgN, OpArgN, IABx, "LOADKX  "},
	{0, 1, OpArgU, OpArgU, IABC, "LOADBOOL"},
	{0, 1, OpArgU, OpArgN, IABC, "LOADNIL "},
	{0, 1, OpArgU, OpArgN, IABC, "LOADNIL "},
	{0, 1, OpArgU, OpArgN, IABC, "GETUPVAL"},
	{0, 1, OpArgU, OpArgK, IABC, "GETTABUP"},
	{0, 1, OpArgR, OpArgK, IABC, "GETTABLE"},
	{0, 0, OpArgK, OpArgK, IABC, "SETTABUP"},
	{0, 0, OpArgU, OpArgN, IABC, "SETUPVAL"},
	{0, 0, OpArgK, OpArgK, IABC, "SETTABLE"},
	{0, 1, OpArgU, OpArgU, IABC, "NEWTABLE"},
	{0, 1, OpArgR, OpArgK, IABC, "SELF    "},
	{0, 1, OpArgK, OpArgK, IABC, "ADD     "},
	{0, 1, OpArgK, OpArgK, IABC, "SUB     "},
	{0, 1, OpArgK, OpArgK, IABC, "MUL     "},
	{0, 1, OpArgK, OpArgK, IABC, "MOD     "},
	{0, 1, OpArgK, OpArgK, IABC, "POW     "},
	{0, 1, OpArgK, OpArgK, IABC, "DIV     "},
	{0, 1, OpArgK, OpArgK, IABC, "IDIV    "},
	{0, 1, OpArgK, OpArgK, IABC, "BAND    "},
	{0, 1, OpArgK, OpArgK, IABC, "BOR     "},
	{0, 1, OpArgK, OpArgK, IABC, "BXOR    "},
	{0, 1, OpArgK, OpArgK, IABC, "SHL     "},
	{0, 1, OpArgK, OpArgK, IABC, "SHR     "},
	{0, 1, OpArgR, OpArgN, IABC, "UNM     "},
	{0, 1, OpArgR, OpArgN, IABC, "BNOT    "},
	{0, 1, OpArgR, OpArgN, IABC, "NOT     "},
	{0, 1, OpArgR, OpArgN, IABC, "LEN     "},
	{0, 1, OpArgR, OpArgR, IABC, "CONCAT  "},
	{0, 0, OpArgR, OpArgN, IAsBx, "JMP     "},
	{1, 0, OpArgK, OpArgK, IABC, "EQ      "},
	{1, 0, OpArgK, OpArgK, IABC, "LT      "},
	{1, 0, OpArgK, OpArgK, IABC, "LE      "},
	{1, 0, OpArgN, OpArgU, IABC, "TEST    "},
	{1, 1, OpArgR, OpArgU, IABC, "TESTSET "},
	{0, 1, OpArgU, OpArgU, IABC, "CALL"},
	{0, 1, OpArgU, OpArgU, IABC, "TAILCALL"},
	{0, 0, OpArgU, OpArgN, IABC, "RETURN"},
	{0, 1, OpArgR, OpArgN, IAsBx, "FORLOOP"},
	{0, 1, OpArgR, OpArgN, IAsBx, "FORPREP"},
	{0, 0, OpArgN, OpArgU, IABC, "TFORCALL"},
	{0, 1, OpArgR, OpArgN, IAsBx, "TFORLOOP"},
	{0, 0, OpArgU, OpArgU, IABC, "SETLIST "},
	{0, 1, OpArgU, OpArgN, IABx, "CLOSURE "},
	{0, 1, OpArgU, OpArgN, IABC, "VARARG  "},
	{0, 0, OpArgU, OpArgU, IAx, "EXTRAARG"},
}
