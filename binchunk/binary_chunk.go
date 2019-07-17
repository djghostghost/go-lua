package binchunk

type header struct {
	signature       []byte
	version         byte
	format          byte
	luacData        []byte
	cintSize        byte
	sizetSize       byte
	instructionSize byte
	luaIntegerSize  byte
	luaNumberSize   byte
	luacInt         int64
	luacNum         float64
}

//函数基本信息
type ProtoType struct {
	//函数源文件名
	Source string
	//开始行号
	LineDefined string
	//终止行号
	LastLineDefined string
	//固定参数个数
	NumParams string
	//是否是vararg 变长函数
	IsVararg byte
	//寄存器数量
	MaxStackSize byte
	Code         []uint32
	Constants    []interface{}
	UpValues     []UpValue
	Protos       []*ProtoType
	LineInfo     []uint32
	LocalVars    []LocalVar
	UpValueNames []string
}

type UpValue struct {
	Instack byte
	Idx     byte
}

type LocalVar struct {
	//变量名
	VarName string
	//开始指令索引
	StartPC uint32
	//终止指令索引
	EndPC uint32
}

const (
	TAG_NIL       = 0x00
	TAG_BOOLEAN   = 0x01
	TAG_NUMBER    = 0x03
	TAG_INTEGER   = 0x13
	TAG_SHORT_STR = 0x04
	TAG_LONG_STR  = 0x14
)

func UnDump(data []byte) *ProtoType {
	reader := &reader{data}
	reader.checkHeader()
	reader.readByte()
	return reader.readProto("")
}
