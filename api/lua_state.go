package api

type LuaType = int

type LuaState interface {
	GetTop() int
	AbsIndex(idx int) int
}
