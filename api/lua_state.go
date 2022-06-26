package api

type LuaType = int
type ArithOp = int
type CompareOp = int

type LuaState interface {
	/* basic stack manipulation */
	GetTop() int
	AbsIndex(idx int) int
	CheckStack(n int) bool
	Pop(n int)
	Copy(fromIdx, toIdx int)
	PushValue(idx int)
	Replace(idx int)
	Insert(idx int)
	Remove(idx int)
	Rotate(idx, n int)
	SetTop(idx int)

	/* access function (stack -> go) */
	TypeName(tp LuaType) string
	Type(idx int) LuaType
	IsNone(idx int) bool
	IsNil(idx int) bool
	IsNoneOrNil(idx int) bool
	IsBoolean(idx int) bool
	IsInteger(idx int) bool
	IsNumber(idx int) bool
	IsString(idx int) bool
	ToBoolean(idx int) bool
	ToInteger(idx int) int64
	ToIntegerX(idx int) (int64, bool)
	ToNumber(idx int) float64
	ToNumberX(idx int) (float64, bool)
	ToString(idx int) string
	ToStringX(idx int) (string, bool)

	/* Push functions (Go -> Stack) */
	PushNil()
	PushBoolean(b bool)
	PushInteger(n int64)
	PushNumber(n float64)
	PushString(s string)

	Arith(op ArithOp)                          // 执行算法和按位计算
	Compare(idx1, idx2 int, op CompareOp) bool //比较计算
	Len(idx int)                               // 执行取长度计算
	Concat(n int)                              // 字符串拼接结算

	/* Get Functions */
	NewTable()
	CreateTable(nArr, nRec int)
	GetTable(idx int) LuaType
	GetField(idx int, k string) LuaType
	GetI(idx int, i int64) LuaType

	/* Set Functions */
	SetTable(idx int)
	SetField(idx int, k string)
	SetI(idx int, n int64)

	Load(chunk []byte, chunkName, mode string) int
	Call(nArgs, nResults int)

	PushGoFunction(f GoFunction)
	IsGoFunction(idx int) bool
	ToGoFunction(idx int) GoFunction

	PushGlobalTable()
	GetGlobal(name string) LuaType
	SetGlobal(name string)
	Register(name string, f GoFunction)

	PushGoClosure(f GoFunction, n int)

	GetMetaTable(idx int) bool
	SetMetaTable(idx int)
	RawLen(idx int) uint
	RawEqual(idx1 int, idx2 int) bool
	RawGet(idx int) LuaType
	RawSet(idx int)
	RawGetI(idx int, i int64) LuaType
	RawSetI(idx int, i int64)
}

type GoFunction func(LuaState) int
