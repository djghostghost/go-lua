package state

type luaTable struct {
	arr  []luaValue
	_map map[luaValue]luaValue
}
