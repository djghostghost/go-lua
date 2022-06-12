package vm

import (
	"github.com/djghostghost/go-lua/api"
)

func getTabUp(i Instruction, vm api.LuaVM) {
	a, b, c := i.ABC()
	a += 1
	b += 1
	vm.GetRK(c)
	vm.GetTable(LuaUpvalueIndex(b))
	vm.Replace(a)
}

func setTabUp(i Instruction, vm api.LuaVM) {
	a, b, c := i.ABC()
	a += 1
	vm.GetRK(b)
	vm.GetRK(c)
	vm.SetTable(LuaUpvalueIndex(a))
}

func getUpval(i Instruction, vm api.LuaVM) {
	a, b, _ := i.ABC()
	a += 1
	b += 1
	vm.Copy(LuaUpvalueIndex(b), a)
}

func setUpval(i Instruction, vm api.LuaVM) {
	a, b, _ := i.ABC()
	a += 1
	b += 1
	vm.Copy(a, LuaUpvalueIndex(b))
}

func LuaUpvalueIndex(i int) int {
	return api.LUA_REGISTRY_INDEX - i
}
