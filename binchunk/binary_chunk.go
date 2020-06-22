package binchunk

type header struct {
	signature       [4]byte
	version         byte
	format          byte
	luacData        [6]byte
	cintSize        byte
	sizetSize       byte
	instructionSize byte
	luaIntegerSize  byte
	luaNumberSize   byte
	luacInt         int64
	luacNum         float64
}

//函数基本信息
type Prototype struct {
	//函数源文件名
	Source string
	//开始行号
	LineDefined uint32
	//终止行号
	LastLineDefined uint32
	//固定参数个数
	NumParams byte
	//是否是vararg 变长函数
	IsVararg byte
	//寄存器数量
	MaxStackSize  byte
	Code          []uint32
	Constants     []interface{}
	UpValues      []UpValue
	SubPrototypes []*Prototype
	LineInfo      []uint32
	LocalVars     []LocalVar
	UpValueNames  []string
}

type UpValue struct {
	InStack byte
	Idx     byte
}

type LocalVar struct {
	VarName string //变量名
	StartPC uint32 //开始指令索引
	EndPC   uint32 //终止指令索引
}

const (
	TAG_NIL       = 0x00
	TAG_BOOLEAN   = 0x01
	TAG_NUMBER    = 0x03
	TAG_INTEGER   = 0x13
	TAG_SHORT_STR = 0x04
	TAG_LONG_STR  = 0x14
)

func UnDump(data []byte) *Prototype {
	reader := &reader{data}
	reader.checkHeader()
	reader.readByte()
	return reader.readPrototype("")
}
