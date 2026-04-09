package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"runtime/debug"
	"strings"
	"strconv"
)

var _ = fmt.Sprintf
var _ = os.Exit
var _ = exec.Command
var _ = runtime.KeepAlive
var _ = debug.SetGCPercent
var _ = strings.Contains
var _ = strconv.Itoa

type SymbolKind interface {
	isSymbolKind()
}

type SymbolKindSkFunction struct{}
func (SymbolKindSkFunction) isSymbolKind() {}

type SymbolKindSkType struct{}
func (SymbolKindSkType) isSymbolKind() {}

type SymbolKindSkEnum struct{}
func (SymbolKindSkEnum) isSymbolKind() {}

type SymbolKindSkTrait struct{}
func (SymbolKindSkTrait) isSymbolKind() {}

type SymbolKindSkConst struct{}
func (SymbolKindSkConst) isSymbolKind() {}

type SymbolKindSkVariable struct{}
func (SymbolKindSkVariable) isSymbolKind() {}

type SymbolKindSkParameter struct{}
func (SymbolKindSkParameter) isSymbolKind() {}

type SymbolKindSkVariant struct{}
func (SymbolKindSkVariant) isSymbolKind() {}

type SymbolKindSkModule struct{}
func (SymbolKindSkModule) isSymbolKind() {}

type SymbolKindSkBuiltin struct{}
func (SymbolKindSkBuiltin) isSymbolKind() {}

type SymbolKindSkNone struct{}
func (SymbolKindSkNone) isSymbolKind() {}

type Symbol struct {
	Name string
	Kind SymbolKind
	Span Span
	Scope_id int64
	Is_mutable bool
	Is_public bool
}

type ScopeKind interface {
	isScopeKind()
}

type ScopeKindScUniverse struct{}
func (ScopeKindScUniverse) isScopeKind() {}

type ScopeKindScModule struct{}
func (ScopeKindScModule) isScopeKind() {}

type ScopeKindScFunction struct{}
func (ScopeKindScFunction) isScopeKind() {}

type ScopeKindScBlock struct{}
func (ScopeKindScBlock) isScopeKind() {}

type ScopeKindScNone struct{}
func (ScopeKindScNone) isScopeKind() {}

type Scope struct {
	Kind ScopeKind
	Parent int64
	Name string
	Depth int64
}

type SymbolTable struct {
	Scopes []Scope
	Symbols []Symbol
}

type StrMap struct {
	Keys []string
	Vals []int64
	Count int64
}

func symbol_kind_name(k SymbolKind) string {
	return func() interface{} {
		switch _tmp1 := (k).(type) {
		case SymbolKindSkFunction:
			return "Function"
		case SymbolKindSkType:
			return "Type"
		case SymbolKindSkEnum:
			return "Enum"
		case SymbolKindSkTrait:
			return "Trait"
		case SymbolKindSkConst:
			return "Const"
		case SymbolKindSkVariable:
			return "Variable"
		case SymbolKindSkParameter:
			return "Parameter"
		case SymbolKindSkVariant:
			return "Variant"
		case SymbolKindSkModule:
			return "Module"
		case SymbolKindSkBuiltin:
			return "Builtin"
		case SymbolKindSkNone:
			return "None"
		default:
			_ = _tmp1
		}
		return nil
	}().(string)
}

func sk_eq(a SymbolKind, b SymbolKind) bool {
	na := symbol_kind_name(a)
	_ = na
	nb := symbol_kind_name(b)
	_ = nb
	return (na == nb)
}

func _sentinel_symbol() Symbol {
	return Symbol{Name: "", Kind: SymbolKindSkNone{}, Span: new_span("", int64(0), int64(0), int64(0), int64(0)), Scope_id: int64(0), Is_mutable: false, Is_public: false}
}

func new_symbol(name string, kind SymbolKind, span Span, scope_id int64, is_mutable bool, is_public bool) Symbol {
	return Symbol{Name: name, Kind: kind, Span: span, Scope_id: scope_id, Is_mutable: is_mutable, Is_public: is_public}
}

func scope_kind_name(k ScopeKind) string {
	return func() interface{} {
		switch _tmp2 := (k).(type) {
		case ScopeKindScUniverse:
			return "Universe"
		case ScopeKindScModule:
			return "Module"
		case ScopeKindScFunction:
			return "Function"
		case ScopeKindScBlock:
			return "Block"
		case ScopeKindScNone:
			return "None"
		default:
			_ = _tmp2
		}
		return nil
	}().(string)
}

func _sentinel_scope() Scope {
	return Scope{Kind: ScopeKindScNone{}, Parent: int64(0), Name: "", Depth: int64(0)}
}

func new_scope(kind ScopeKind, parent int64, name string, depth int64) Scope {
	return Scope{Kind: kind, Parent: parent, Name: name, Depth: depth}
}

func new_table() SymbolTable {
	return SymbolTable{Scopes: []Scope{_sentinel_scope()}, Symbols: []Symbol{_sentinel_symbol()}}
}

func add_scope(table SymbolTable, kind ScopeKind, parent int64, name string) SymbolTable {
	depth := int64(0)
	_ = depth
	if ((parent > int64(0)) && (parent < int64(len(table.Scopes)))) {
		depth = (table.Scopes[parent].Depth + int64(1))
	}
	s := new_scope(kind, parent, name, depth)
	_ = s
	return SymbolTable{Scopes: append(table.Scopes, s), Symbols: table.Symbols}
}

func last_scope_id(table SymbolTable) int64 {
	return (int64(len(table.Scopes)) - int64(1))
}

func add_symbol(table SymbolTable, sym Symbol) SymbolTable {
	return SymbolTable{Scopes: table.Scopes, Symbols: append(table.Symbols, sym)}
}

func scope_count(table SymbolTable) int64 {
	return int64(len(table.Scopes))
}

func symbol_count(table SymbolTable) int64 {
	return int64(len(table.Symbols))
}

func lookup_local(table SymbolTable, scope_id int64, name string) Symbol {
	i := int64(1)
	_ = i
	for (i < int64(len(table.Symbols))) {
		sym := table.Symbols[i]
		_ = sym
		if (sym.Scope_id == scope_id) {
			if (sym.Name == name) {
				return sym
			}
		}
		i = (i + int64(1))
	}
	return _sentinel_symbol()
}

func lookup(table SymbolTable, scope_id int64, name string) Symbol {
	cur := scope_id
	_ = cur
	for (cur > int64(0)) {
		found := lookup_local(table, cur, name)
		_ = found
		if (found.Name != "") {
			return found
		}
		if (cur < int64(len(table.Scopes))) {
			cur = table.Scopes[cur].Parent
		} else {
			cur = int64(0)
		}
	}
	return _sentinel_symbol()
}

func strmap_new() StrMap {
	return StrMap{Keys: []string{""}, Vals: []int64{int64(0)}, Count: int64(0)}
}

func _char_ord(c string) int64 {
	if (c == "a") {
		return int64(97)
	}
	if (c == "b") {
		return int64(98)
	}
	if (c == "c") {
		return int64(99)
	}
	if (c == "d") {
		return int64(100)
	}
	if (c == "e") {
		return int64(101)
	}
	if (c == "f") {
		return int64(102)
	}
	if (c == "g") {
		return int64(103)
	}
	if (c == "h") {
		return int64(104)
	}
	if (c == "i") {
		return int64(105)
	}
	if (c == "j") {
		return int64(106)
	}
	if (c == "k") {
		return int64(107)
	}
	if (c == "l") {
		return int64(108)
	}
	if (c == "m") {
		return int64(109)
	}
	if (c == "n") {
		return int64(110)
	}
	if (c == "o") {
		return int64(111)
	}
	if (c == "p") {
		return int64(112)
	}
	if (c == "q") {
		return int64(113)
	}
	if (c == "r") {
		return int64(114)
	}
	if (c == "s") {
		return int64(115)
	}
	if (c == "t") {
		return int64(116)
	}
	if (c == "u") {
		return int64(117)
	}
	if (c == "v") {
		return int64(118)
	}
	if (c == "w") {
		return int64(119)
	}
	if (c == "x") {
		return int64(120)
	}
	if (c == "y") {
		return int64(121)
	}
	if (c == "z") {
		return int64(122)
	}
	if (c == "A") {
		return int64(65)
	}
	if (c == "B") {
		return int64(66)
	}
	if (c == "C") {
		return int64(67)
	}
	if (c == "D") {
		return int64(68)
	}
	if (c == "E") {
		return int64(69)
	}
	if (c == "F") {
		return int64(70)
	}
	if (c == "G") {
		return int64(71)
	}
	if (c == "H") {
		return int64(72)
	}
	if (c == "I") {
		return int64(73)
	}
	if (c == "J") {
		return int64(74)
	}
	if (c == "K") {
		return int64(75)
	}
	if (c == "L") {
		return int64(76)
	}
	if (c == "M") {
		return int64(77)
	}
	if (c == "N") {
		return int64(78)
	}
	if (c == "O") {
		return int64(79)
	}
	if (c == "P") {
		return int64(80)
	}
	if (c == "Q") {
		return int64(81)
	}
	if (c == "R") {
		return int64(82)
	}
	if (c == "S") {
		return int64(83)
	}
	if (c == "T") {
		return int64(84)
	}
	if (c == "U") {
		return int64(85)
	}
	if (c == "V") {
		return int64(86)
	}
	if (c == "W") {
		return int64(87)
	}
	if (c == "X") {
		return int64(88)
	}
	if (c == "Y") {
		return int64(89)
	}
	if (c == "Z") {
		return int64(90)
	}
	if (c == "0") {
		return int64(48)
	}
	if (c == "1") {
		return int64(49)
	}
	if (c == "2") {
		return int64(50)
	}
	if (c == "3") {
		return int64(51)
	}
	if (c == "4") {
		return int64(52)
	}
	if (c == "5") {
		return int64(53)
	}
	if (c == "6") {
		return int64(54)
	}
	if (c == "7") {
		return int64(55)
	}
	if (c == "8") {
		return int64(56)
	}
	if (c == "9") {
		return int64(57)
	}
	if (c == "_") {
		return int64(95)
	}
	if (c == ":") {
		return int64(58)
	}
	if (c == ".") {
		return int64(46)
	}
	if (c == "-") {
		return int64(45)
	}
	if (c == "/") {
		return int64(47)
	}
	return int64(42)
}

func _str_hash(s string) int64 {
	h := int64(5381)
	_ = h
	i := int64(0)
	_ = i
	for (i < int64(len(s))) {
		h = ((h * int64(33)) + _char_ord(string(s[i])))
		i = (i + int64(1))
	}
	if (h < int64(0)) {
		h = (int64(0) - h)
	}
	return h
}

func strmap_put(m StrMap, key string, val int64) StrMap {
	return StrMap{Keys: append(m.Keys, key), Vals: append(m.Vals, val), Count: (m.Count + int64(1))}
}

func strmap_get(m StrMap, key string) int64 {
	i := int64(1)
	_ = i
	for (i < int64(len(m.Keys))) {
		if (m.Keys[i] == key) {
			return m.Vals[i]
		}
		i = (i + int64(1))
	}
	return (int64(0) - int64(1))
}

func strmap_has(m StrMap, key string) bool {
	return (strmap_get(m, key) >= int64(0))
}

