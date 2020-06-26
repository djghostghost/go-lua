package state

import (
	"fmt"
	. "github.com/djghostghost/go-lua/api"
	"github.com/djghostghost/go-lua/number"
)

func (s *luaState) TypeName(tp LuaType) string {
	switch tp {
	case LUA_TNONE:
		return "no value"
	case LUA_TNIL:
		return "nil"
	case LUA_TBOOLEAN:
		return "boolean"
	case LUA_TNUMBER:
		return "number"
	case LUA_TSTRING:
		return "string"
	case LUA_TTABLE:
		return "table"
	case LUA_TFUNCTION:
		return "function"
	case LUA_TTHREAD:
		return "thread"
	default:
		return "userdata"
	}
}

func (s *luaState) Type(idx int) LuaType {
	if s.stack.isValid(idx) {
		val := s.stack.get(idx)
		return typeOf(val)
	}
	return LUA_TNONE
}

func (s *luaState) IsNone(idx int) bool {
	return s.Type(idx) == LUA_TNONE
}

func (s *luaState) IsNil(idx int) bool {
	return s.Type(idx) == LUA_TNIL
}

func (s *luaState) IsNoneOrNil(idx int) bool {
	return s.Type(idx) <= LUA_TNIL
}

func (s *luaState) IsBoolean(idx int) bool {
	return s.Type(idx) == LUA_TBOOLEAN
}

func (s *luaState) IsString(idx int) bool {
	t := s.Type(idx)
	return t == LUA_TSTRING || t == LUA_TNUMBER
}

func (s *luaState) IsNumber(idx int) bool {
	_, ok := s.ToNumberX(idx)
	return ok
}

func (s *luaState) IsInteger(idx int) bool {
	val := s.stack.get(idx)
	_, ok := val.(int64)
	return ok
}

func (s *luaState) ToBoolean(idx int) bool {
	val := s.stack.get(idx)
	return convertToBoolean(val)
}

func convertToBoolean(val luaValue) bool {
	switch x := val.(type) {
	case nil:
		return false
	case bool:
		return x
	default:
		return true
	}
}

func (s *luaState) ToNumber(idx int) float64 {
	n, _ := s.ToNumberX(idx)
	return n
}

func (s *luaState) ToNumberX(idx int) (float64, bool) {
	val := s.stack.get(idx)
	return convertToFloat(val)
}

func (s *luaState) ToInteger(idx int) int64 {
	i, _ := s.ToIntegerX(idx)
	return i
}

func (s *luaState) ToIntegerX(idx int) (int64, bool) {
	val := s.stack.get(idx)
	return convertToInteger(val)
}

func (s *luaState) ToStringX(idx int) (string, bool) {
	val := s.stack.get(idx)
	switch x := val.(type) {
	case string:
		return x, true
	case int64, float64:
		str := fmt.Sprintf("%v", x)
		s.stack.set(idx, str)
		return str, true
	default:
		return "", false
	}
}

func (s *luaState) ToString(idx int) string {
	str, _ := s.ToStringX(idx)
	return str
}

func convertToFloat(val luaValue) (float64, bool) {
	switch x := val.(type) {
	case float64:
		return x, true
	case int64:
		return float64(x), true
	case string:
		return number.ParseFloat(x)
	default:
		return 0, false
	}
}

func convertToInteger(val luaValue) (int64, bool) {
	switch x := val.(type) {
	case int64:
		return x, true
	case float64:
		return number.FloatToInteger(x)
	case string:
		return _stringToInteger(x)
	default:
		return 0, false
	}
}

func _stringToInteger(s string) (int64, bool) {
	if i, ok := number.ParseInteger(s); ok {
		return i, true
	}
	if f, ok := number.ParseFloat(s); ok {
		return number.FloatToInteger(f)
	}
	return 0, false
}
