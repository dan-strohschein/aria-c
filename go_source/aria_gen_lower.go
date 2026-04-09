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

type Lowerer struct {
	Tokens []Token
	Pos int64
	Store TypeStore
	Registry TypeRegistry
	Treg TraitRegistry
	Pool NodePool
	Table SymbolTable
	File string
	Module IrModule
	Current_func IrFunc
	Temp_counter int64
	Label_counter int64
	Env_names []string
	Env_slots []int64
	Loop_start int64
	Loop_end int64
	Const_names []string
	Const_vals []int64
	Const_str_names []string
	Const_str_vals []string
	Temp_types []int64
	Fn_error_type int64
	Defer_starts []int64
	Defer_ends []int64
}

type LT struct {
	L Lowerer
	Temp int64
}

func _new_lowerer(tokens []Token, pool NodePool, store TypeStore, reg TypeRegistry, treg TraitRegistry, table SymbolTable, file string) Lowerer {
	return Lowerer{Tokens: tokens, Pos: int64(0), Store: store, Registry: reg, Treg: treg, Pool: pool, Table: table, File: file, Module: new_ir_module(), Current_func: _sentinel_func(), Temp_counter: int64(0), Label_counter: int64(0), Env_names: []string{""}, Env_slots: []int64{int64(0)}, Loop_start: -int64(1), Loop_end: -int64(1), Const_names: []string{""}, Const_vals: []int64{int64(0)}, Const_str_names: []string{""}, Const_str_vals: []string{""}, Temp_types: []int64{int64(0)}, Fn_error_type: int64(0), Defer_starts: []int64{int64(0)}, Defer_ends: []int64{int64(0)}}
}

func _lcur(l Lowerer) Token {
	if (l.Pos >= int64(len(l.Tokens))) {
		return Token{Kind: TokenKindTkEof{}, Text: "", Line: int64(0), Col: int64(0), Offset: int64(0)}
	}
	return l.Tokens[l.Pos]
}

func _lk(l Lowerer) string {
	return token_name(_lcur(l).Kind)
}

func _ladv(l Lowerer) Lowerer {
	return Lowerer{Tokens: l.Tokens, Pos: (l.Pos + int64(1)), Store: l.Store, Registry: l.Registry, Treg: l.Treg, Pool: l.Pool, Table: l.Table, File: l.File, Module: l.Module, Current_func: l.Current_func, Temp_counter: l.Temp_counter, Label_counter: l.Label_counter, Env_names: l.Env_names, Env_slots: l.Env_slots, Loop_start: l.Loop_start, Loop_end: l.Loop_end, Const_names: l.Const_names, Const_vals: l.Const_vals, Const_str_names: l.Const_str_names, Const_str_vals: l.Const_str_vals, Temp_types: l.Temp_types, Fn_error_type: l.Fn_error_type, Defer_starts: l.Defer_starts, Defer_ends: l.Defer_ends}
}

func _lskip_pub(l Lowerer) Lowerer {
	ll := l
	_ = ll
	if (_lk(ll) == "pub") {
		ll = _ladv(ll)
		if (_lk(ll) == "(") {
			ll = _ladv(ll)
			if (_lk(ll) == "IDENT") {
				ll = _ladv(ll)
			}
			if (_lk(ll) == ")") {
				ll = _ladv(ll)
			}
		}
	}
	return ll
}

func _lset_reg(l Lowerer, reg TypeRegistry) Lowerer {
	return Lowerer{Tokens: l.Tokens, Pos: l.Pos, Store: l.Store, Registry: reg, Treg: l.Treg, Pool: l.Pool, Table: l.Table, File: l.File, Module: l.Module, Current_func: l.Current_func, Temp_counter: l.Temp_counter, Label_counter: l.Label_counter, Env_names: l.Env_names, Env_slots: l.Env_slots, Loop_start: l.Loop_start, Loop_end: l.Loop_end, Const_names: l.Const_names, Const_vals: l.Const_vals, Const_str_names: l.Const_str_names, Const_str_vals: l.Const_str_vals, Temp_types: l.Temp_types, Fn_error_type: l.Fn_error_type, Defer_starts: l.Defer_starts, Defer_ends: l.Defer_ends}
}

func _lset_pos(l Lowerer, pos int64) Lowerer {
	return Lowerer{Tokens: l.Tokens, Pos: pos, Store: l.Store, Registry: l.Registry, Treg: l.Treg, Pool: l.Pool, Table: l.Table, File: l.File, Module: l.Module, Current_func: l.Current_func, Temp_counter: l.Temp_counter, Label_counter: l.Label_counter, Env_names: l.Env_names, Env_slots: l.Env_slots, Loop_start: l.Loop_start, Loop_end: l.Loop_end, Const_names: l.Const_names, Const_vals: l.Const_vals, Const_str_names: l.Const_str_names, Const_str_vals: l.Const_str_vals, Temp_types: l.Temp_types, Fn_error_type: l.Fn_error_type, Defer_starts: l.Defer_starts, Defer_ends: l.Defer_ends}
}

func _lskip_nl(l Lowerer) Lowerer {
	ll := l
	_ = ll
	for (_lk(ll) == "NEWLINE") {
		ll = _ladv(ll)
	}
	return ll
}

func _lnew_temp(l Lowerer) LT {
	temp := l.Temp_counter
	_ = temp
	ll := Lowerer{Tokens: l.Tokens, Pos: l.Pos, Store: l.Store, Registry: l.Registry, Treg: l.Treg, Pool: l.Pool, Table: l.Table, File: l.File, Module: l.Module, Current_func: l.Current_func, Temp_counter: (l.Temp_counter + int64(1)), Label_counter: l.Label_counter, Env_names: l.Env_names, Env_slots: l.Env_slots, Loop_start: l.Loop_start, Loop_end: l.Loop_end, Const_names: l.Const_names, Const_vals: l.Const_vals, Const_str_names: l.Const_str_names, Const_str_vals: l.Const_str_vals, Temp_types: l.Temp_types, Fn_error_type: l.Fn_error_type, Defer_starts: l.Defer_starts, Defer_ends: l.Defer_ends}
	_ = ll
	return LT{L: ll, Temp: temp}
}

func _lnew_label(l Lowerer) LT {
	lbl := l.Label_counter
	_ = lbl
	ll := Lowerer{Tokens: l.Tokens, Pos: l.Pos, Store: l.Store, Registry: l.Registry, Treg: l.Treg, Pool: l.Pool, Table: l.Table, File: l.File, Module: l.Module, Current_func: l.Current_func, Temp_counter: l.Temp_counter, Label_counter: (l.Label_counter + int64(1)), Env_names: l.Env_names, Env_slots: l.Env_slots, Loop_start: l.Loop_start, Loop_end: l.Loop_end, Const_names: l.Const_names, Const_vals: l.Const_vals, Const_str_names: l.Const_str_names, Const_str_vals: l.Const_str_vals, Temp_types: l.Temp_types, Fn_error_type: l.Fn_error_type, Defer_starts: l.Defer_starts, Defer_ends: l.Defer_ends}
	_ = ll
	return LT{L: ll, Temp: lbl}
}

func _lemit(l Lowerer, inst IrInst) Lowerer {
	new_func := func_add_inst(l.Current_func, inst)
	_ = new_func
	return Lowerer{Tokens: l.Tokens, Pos: l.Pos, Store: l.Store, Registry: l.Registry, Treg: l.Treg, Pool: l.Pool, Table: l.Table, File: l.File, Module: l.Module, Current_func: new_func, Temp_counter: l.Temp_counter, Label_counter: l.Label_counter, Env_names: l.Env_names, Env_slots: l.Env_slots, Loop_start: l.Loop_start, Loop_end: l.Loop_end, Const_names: l.Const_names, Const_vals: l.Const_vals, Const_str_names: l.Const_str_names, Const_str_vals: l.Const_str_vals, Temp_types: l.Temp_types, Fn_error_type: l.Fn_error_type, Defer_starts: l.Defer_starts, Defer_ends: l.Defer_ends}
}

func _lset_func(l Lowerer, f IrFunc) Lowerer {
	return Lowerer{Tokens: l.Tokens, Pos: l.Pos, Store: l.Store, Registry: l.Registry, Treg: l.Treg, Pool: l.Pool, Table: l.Table, File: l.File, Module: l.Module, Current_func: f, Temp_counter: int64(0), Label_counter: int64(0), Env_names: []string{""}, Env_slots: []int64{int64(0)}, Loop_start: -int64(1), Loop_end: -int64(1), Const_names: l.Const_names, Const_vals: l.Const_vals, Const_str_names: l.Const_str_names, Const_str_vals: l.Const_str_vals, Temp_types: []int64{int64(0)}, Fn_error_type: int64(0), Defer_starts: []int64{int64(0)}, Defer_ends: []int64{int64(0)}}
}

func _lset_mod(l Lowerer, m IrModule) Lowerer {
	return Lowerer{Tokens: l.Tokens, Pos: l.Pos, Store: l.Store, Registry: l.Registry, Treg: l.Treg, Pool: l.Pool, Table: l.Table, File: l.File, Module: m, Current_func: l.Current_func, Temp_counter: l.Temp_counter, Label_counter: l.Label_counter, Env_names: l.Env_names, Env_slots: l.Env_slots, Loop_start: l.Loop_start, Loop_end: l.Loop_end, Const_names: l.Const_names, Const_vals: l.Const_vals, Const_str_names: l.Const_str_names, Const_str_vals: l.Const_str_vals, Temp_types: l.Temp_types, Fn_error_type: l.Fn_error_type, Defer_starts: l.Defer_starts, Defer_ends: l.Defer_ends}
}

func _lfinish_func(l Lowerer) Lowerer {
	f0 := func_set_locals(l.Current_func, l.Temp_counter)
	_ = f0
	f := func_set_temp_types(f0, l.Temp_types)
	_ = f
	new_mod := mod_add_func(l.Module, f)
	_ = new_mod
	return Lowerer{Tokens: l.Tokens, Pos: l.Pos, Store: l.Store, Registry: l.Registry, Treg: l.Treg, Pool: l.Pool, Table: l.Table, File: l.File, Module: new_mod, Current_func: _sentinel_func(), Temp_counter: int64(0), Label_counter: l.Label_counter, Env_names: []string{""}, Env_slots: []int64{int64(0)}, Loop_start: -int64(1), Loop_end: -int64(1), Const_names: l.Const_names, Const_vals: l.Const_vals, Const_str_names: l.Const_str_names, Const_str_vals: l.Const_str_vals, Temp_types: l.Temp_types, Fn_error_type: l.Fn_error_type, Defer_starts: l.Defer_starts, Defer_ends: l.Defer_ends}
}

func _lset_loop(l Lowerer, start int64, end_l int64) Lowerer {
	return Lowerer{Tokens: l.Tokens, Pos: l.Pos, Store: l.Store, Registry: l.Registry, Treg: l.Treg, Pool: l.Pool, Table: l.Table, File: l.File, Module: l.Module, Current_func: l.Current_func, Temp_counter: l.Temp_counter, Label_counter: l.Label_counter, Env_names: l.Env_names, Env_slots: l.Env_slots, Loop_start: start, Loop_end: end_l, Const_names: l.Const_names, Const_vals: l.Const_vals, Const_str_names: l.Const_str_names, Const_str_vals: l.Const_str_vals, Temp_types: l.Temp_types, Fn_error_type: l.Fn_error_type, Defer_starts: l.Defer_starts, Defer_ends: l.Defer_ends}
}

func _lenv_add(l Lowerer, name string, slot int64) Lowerer {
	return Lowerer{Tokens: l.Tokens, Pos: l.Pos, Store: l.Store, Registry: l.Registry, Treg: l.Treg, Pool: l.Pool, Table: l.Table, File: l.File, Module: l.Module, Current_func: l.Current_func, Temp_counter: l.Temp_counter, Label_counter: l.Label_counter, Env_names: append(l.Env_names, name), Env_slots: append(l.Env_slots, slot), Loop_start: l.Loop_start, Loop_end: l.Loop_end, Const_names: l.Const_names, Const_vals: l.Const_vals, Const_str_names: l.Const_str_names, Const_str_vals: l.Const_str_vals, Temp_types: l.Temp_types, Fn_error_type: l.Fn_error_type, Defer_starts: l.Defer_starts, Defer_ends: l.Defer_ends}
}

func _lenv_lookup(l Lowerer, name string) int64 {
	i := (int64(len(l.Env_names)) - int64(1))
	_ = i
	for (i > int64(0)) {
		if (l.Env_names[i] == name) {
			return l.Env_slots[i]
		}
		i = (i - int64(1))
	}
	return -int64(1)
}

func _lset_error_type(l Lowerer, et int64) Lowerer {
	return Lowerer{Tokens: l.Tokens, Pos: l.Pos, Store: l.Store, Registry: l.Registry, Treg: l.Treg, Pool: l.Pool, Table: l.Table, File: l.File, Module: l.Module, Current_func: l.Current_func, Temp_counter: l.Temp_counter, Label_counter: l.Label_counter, Env_names: l.Env_names, Env_slots: l.Env_slots, Loop_start: l.Loop_start, Loop_end: l.Loop_end, Const_names: l.Const_names, Const_vals: l.Const_vals, Const_str_names: l.Const_str_names, Const_str_vals: l.Const_str_vals, Temp_types: l.Temp_types, Fn_error_type: et, Defer_starts: l.Defer_starts, Defer_ends: l.Defer_ends}
}

func _ladd_const(l Lowerer, name string, val int64) Lowerer {
	return Lowerer{Tokens: l.Tokens, Pos: l.Pos, Store: l.Store, Registry: l.Registry, Treg: l.Treg, Pool: l.Pool, Table: l.Table, File: l.File, Module: l.Module, Current_func: l.Current_func, Temp_counter: l.Temp_counter, Label_counter: l.Label_counter, Env_names: l.Env_names, Env_slots: l.Env_slots, Loop_start: l.Loop_start, Loop_end: l.Loop_end, Const_names: append(l.Const_names, name), Const_vals: append(l.Const_vals, val), Const_str_names: l.Const_str_names, Const_str_vals: l.Const_str_vals, Temp_types: l.Temp_types, Fn_error_type: l.Fn_error_type, Defer_starts: l.Defer_starts, Defer_ends: l.Defer_ends}
}

func _ladd_const_str(l Lowerer, name string, val string) Lowerer {
	return Lowerer{Tokens: l.Tokens, Pos: l.Pos, Store: l.Store, Registry: l.Registry, Treg: l.Treg, Pool: l.Pool, Table: l.Table, File: l.File, Module: l.Module, Current_func: l.Current_func, Temp_counter: l.Temp_counter, Label_counter: l.Label_counter, Env_names: l.Env_names, Env_slots: l.Env_slots, Loop_start: l.Loop_start, Loop_end: l.Loop_end, Const_names: l.Const_names, Const_vals: l.Const_vals, Const_str_names: append(l.Const_str_names, name), Const_str_vals: append(l.Const_str_vals, val), Temp_types: l.Temp_types, Fn_error_type: l.Fn_error_type, Defer_starts: l.Defer_starts, Defer_ends: l.Defer_ends}
}

func _lenv_restore(l Lowerer, names []string, slots []int64) Lowerer {
	return Lowerer{Tokens: l.Tokens, Pos: l.Pos, Store: l.Store, Registry: l.Registry, Treg: l.Treg, Pool: l.Pool, Table: l.Table, File: l.File, Module: l.Module, Current_func: l.Current_func, Temp_counter: l.Temp_counter, Label_counter: l.Label_counter, Env_names: names, Env_slots: slots, Loop_start: l.Loop_start, Loop_end: l.Loop_end, Const_names: l.Const_names, Const_vals: l.Const_vals, Const_str_names: l.Const_str_names, Const_str_vals: l.Const_str_vals, Temp_types: l.Temp_types, Fn_error_type: l.Fn_error_type, Defer_starts: l.Defer_starts, Defer_ends: l.Defer_ends}
}

func _lnew_str_temp(l Lowerer) LT {
	nt := _lnew_temp(l)
	_ = nt
	nt2 := _lnew_temp(nt.L)
	_ = nt2
	return LT{L: nt2.L, Temp: nt.Temp}
}

func _lconst_lookup(l Lowerer, name string) int64 {
	i := int64(1)
	_ = i
	for (i < int64(len(l.Const_names))) {
		if (l.Const_names[i] == name) {
			return l.Const_vals[i]
		}
		i = (i + int64(1))
	}
	return -int64(9999999)
}

func _lconst_str_lookup(l Lowerer, name string) string {
	i := int64(1)
	_ = i
	for (i < int64(len(l.Const_str_names))) {
		if (l.Const_str_names[i] == name) {
			return l.Const_str_vals[i]
		}
		i = (i + int64(1))
	}
	return ""
}

func _lconst_has(l Lowerer, name string) bool {
	i := int64(1)
	_ = i
	for (i < int64(len(l.Const_names))) {
		if (l.Const_names[i] == name) {
			return true
		}
		i = (i + int64(1))
	}
	return false
}

func _lconst_str_has(l Lowerer, name string) bool {
	i := int64(1)
	_ = i
	for (i < int64(len(l.Const_str_names))) {
		if (l.Const_str_names[i] == name) {
			return true
		}
		i = (i + int64(1))
	}
	return false
}

func _array_elem_type(store TypeStore, arr_type int64) int64 {
	if (arr_type == int64(13)) {
		return int64(12)
	}
	if (arr_type <= int64(16)) {
		return int64(0)
	}
	info := store_get(store, arr_type)
	_ = info
	if (type_kind_name(info.Kind) == "Array") {
		return info.Type_id
	}
	return int64(0)
}

func _is_str_array_type(store TypeStore, tid int64) bool {
	info := store_get(store, tid)
	_ = info
	if ((type_kind_name(info.Kind) == "Array") && (info.Type_id == int64(12))) {
		return true
	}
	return false
}

func _is_trait_object_type(store TypeStore, tid int64) bool {
	if (tid <= int64(16)) {
		return false
	}
	if (tid >= int64(len(store.Types))) {
		return false
	}
	info := store.Types[tid]
	_ = info
	return (type_kind_name(info.Kind) == "TraitObject")
}

func _is_named_type(store TypeStore, tid int64) bool {
	if (tid >= TY_STRUCT_BASE) {
		return true
	}
	if (tid <= int64(16)) {
		return false
	}
	info := store_get(store, tid)
	_ = info
	if (type_kind_name(info.Kind) == "Named") {
		return true
	}
	return false
}

func _lower_struct_eq(l Lowerer, a int64, b int64, sname string, result_temp int64) Lowerer {
	ll := l
	_ = ll
	def := reg_find_struct(ll.Registry, sname)
	_ = def
	if ((def.Name == "") || (int64(len(def.Fields)) <= int64(1))) {
		ll = _lemit(ll, new_inst(IrOpOpEq{}, result_temp, a, b, "", int64(0)))
		return ll
	}
	last_cmp := -int64(1)
	_ = last_cmp
	fi := int64(1)
	_ = fi
	field_slot := int64(0)
	_ = field_slot
	for (fi < int64(len(def.Fields))) {
		fld := def.Fields[fi]
		_ = fld
		ft := fld.Type_id
		_ = ft
		if (ft == int64(12)) {
			a_str := _lnew_str_temp(ll)
			_ = a_str
			ll = a_str.L
			ll = _lemit(ll, new_inst(IrOpOpFieldGet{}, a_str.Temp, a, field_slot, fld.Name, int64(0)))
			fidx_len := (field_slot + int64(1))
			_ = fidx_len
			ll = _lemit(ll, new_inst(IrOpOpFieldGet{}, (a_str.Temp + int64(1)), a, fidx_len, fld.Name, int64(0)))
			b_str := _lnew_str_temp(ll)
			_ = b_str
			ll = b_str.L
			ll = _lemit(ll, new_inst(IrOpOpFieldGet{}, b_str.Temp, b, field_slot, fld.Name, int64(0)))
			ll = _lemit(ll, new_inst(IrOpOpFieldGet{}, (b_str.Temp + int64(1)), b, fidx_len, fld.Name, int64(0)))
			cmp := _lnew_temp(ll)
			_ = cmp
			ll = cmp.L
			ll = _lemit(ll, new_inst(IrOpOpStrEq{}, cmp.Temp, a_str.Temp, b_str.Temp, "", int64(0)))
			if (last_cmp >= int64(0)) {
				and_nt := _lnew_temp(ll)
				_ = and_nt
				ll = and_nt.L
				ll = _lemit(ll, new_inst(IrOpOpAnd{}, and_nt.Temp, last_cmp, cmp.Temp, "", int64(0)))
				last_cmp = and_nt.Temp
			} else {
				last_cmp = cmp.Temp
			}
			field_slot = (field_slot + int64(2))
		} else {
			a_val := _lnew_temp(ll)
			_ = a_val
			ll = a_val.L
			ll = _lemit(ll, new_inst(IrOpOpFieldGet{}, a_val.Temp, a, field_slot, fld.Name, int64(0)))
			b_val := _lnew_temp(ll)
			_ = b_val
			ll = b_val.L
			ll = _lemit(ll, new_inst(IrOpOpFieldGet{}, b_val.Temp, b, field_slot, fld.Name, int64(0)))
			cmp := _lnew_temp(ll)
			_ = cmp
			ll = cmp.L
			ll = _lemit(ll, new_inst(IrOpOpEq{}, cmp.Temp, a_val.Temp, b_val.Temp, "", int64(0)))
			if (last_cmp >= int64(0)) {
				and_nt := _lnew_temp(ll)
				_ = and_nt
				ll = and_nt.L
				ll = _lemit(ll, new_inst(IrOpOpAnd{}, and_nt.Temp, last_cmp, cmp.Temp, "", int64(0)))
				last_cmp = and_nt.Temp
			} else {
				last_cmp = cmp.Temp
			}
			field_slot = (field_slot + int64(1))
		}
		fi = (fi + int64(1))
	}
	if (last_cmp >= int64(0)) {
		ll = _lemit(ll, new_inst(IrOpOpStore{}, result_temp, last_cmp, int64(0), "", int64(0)))
	} else {
		ll = _lemit(ll, new_inst(IrOpOpConst{}, result_temp, int64(1), int64(0), "", int64(0)))
	}
	return ll
}

func _is_struct_type(store TypeStore, reg TypeRegistry, tid int64) bool {
	if (tid >= TY_STRUCT_BASE) {
		return true
	}
	if (tid <= int64(16)) {
		return false
	}
	info := store_get(store, tid)
	_ = info
	if (type_kind_name(info.Kind) == "Named") {
		def := reg_find_struct(reg, info.Name)
		_ = def
		if ((def.Name != "") && (def.Is_sum == false)) {
			return true
		}
	}
	return false
}

func _struct_name_for_type(store TypeStore, reg TypeRegistry, tid int64) string {
	if (tid >= TY_STRUCT_BASE) {
		return _struct_name_from_type(reg, tid)
	}
	if (tid <= int64(16)) {
		return ""
	}
	info := store_get(store, tid)
	_ = info
	if (type_kind_name(info.Kind) == "Named") {
		return info.Name
	}
	return ""
}

func _lset_type(l Lowerer, temp int64, type_id int64) Lowerer {
	if (temp < int64(0)) {
		return l
	}
	if (temp > int64(50000)) {
		return l
	}
	types := l.Temp_types
	_ = types
	for (int64(len(types)) <= (temp + int64(1))) {
		types = append(types, int64(0))
	}
	new_types := []int64{int64(0)}
	_ = new_types
	idx := int64(1)
	_ = idx
	for (idx < int64(len(types))) {
		target := (temp + int64(1))
		_ = target
		if (idx == target) {
			new_types = append(new_types, type_id)
		} else {
			new_types = append(new_types, types[idx])
		}
		idx = (idx + int64(1))
	}
	return Lowerer{Tokens: l.Tokens, Pos: l.Pos, Store: l.Store, Registry: l.Registry, Treg: l.Treg, Pool: l.Pool, Table: l.Table, File: l.File, Module: l.Module, Current_func: l.Current_func, Temp_counter: l.Temp_counter, Label_counter: l.Label_counter, Env_names: l.Env_names, Env_slots: l.Env_slots, Loop_start: l.Loop_start, Loop_end: l.Loop_end, Const_names: l.Const_names, Const_vals: l.Const_vals, Const_str_names: l.Const_str_names, Const_str_vals: l.Const_str_vals, Temp_types: new_types, Fn_error_type: l.Fn_error_type, Defer_starts: l.Defer_starts, Defer_ends: l.Defer_ends}
}

func _lget_type(l Lowerer, temp int64) int64 {
	idx := (temp + int64(1))
	_ = idx
	if ((idx >= int64(0)) && (idx < int64(len(l.Temp_types)))) {
		return l.Temp_types[idx]
	}
	return int64(0)
}

func _arith_op(k string) IrOp {
	if (k == "+") {
		return IrOpOpAdd{}
	}
	if (k == "-") {
		return IrOpOpSub{}
	}
	if (k == "*") {
		return IrOpOpMul{}
	}
	if (k == "/") {
		return IrOpOpDiv{}
	}
	if (k == "%") {
		return IrOpOpMod{}
	}
	return IrOpOpAdd{}
}

func _cmp_op(k string) IrOp {
	if (k == "==") {
		return IrOpOpEq{}
	}
	if (k == "!=") {
		return IrOpOpNeq{}
	}
	if (k == "<") {
		return IrOpOpLt{}
	}
	if (k == ">") {
		return IrOpOpGt{}
	}
	if (k == "<=") {
		return IrOpOpLte{}
	}
	if (k == ">=") {
		return IrOpOpGte{}
	}
	return IrOpOpEq{}
}

func _unary_op(k string) IrOp {
	if (k == "!") {
		return IrOpOpNot{}
	}
	return IrOpOpNeg{}
}

func _cmp_or_logic_op(k string) IrOp {
	if (k == "==") {
		return IrOpOpEq{}
	}
	if (k == "!=") {
		return IrOpOpNeq{}
	}
	if (k == "<") {
		return IrOpOpLt{}
	}
	if (k == ">") {
		return IrOpOpGt{}
	}
	if (k == "<=") {
		return IrOpOpLte{}
	}
	if (k == ">=") {
		return IrOpOpGte{}
	}
	if (k == "&&") {
		return IrOpOpAnd{}
	}
	if (k == "||") {
		return IrOpOpOr{}
	}
	return IrOpOpEq{}
}

func _is_float_type(t int64) bool {
	return (t == int64(9))
}

func _is_cont_l(k string) bool {
	return ((((((((((((((((((((((((((k == "+") || (k == "-")) || (k == "*")) || (k == "/")) || (k == "%")) || (k == "==")) || (k == "!=")) || (k == "<")) || (k == ">")) || (k == "<=")) || (k == ">=")) || (k == "&&")) || (k == "||")) || (k == "&")) || (k == "|")) || (k == "^")) || (k == "<<")) || (k == ">>")) || (k == "|>")) || (k == ".")) || (k == "(")) || (k == "[")) || (k == "?")) || (k == "!")) || (k == ",")) || (k == "=>"))
}

func _scan_consts(l Lowerer, index DeclIndex) Lowerer {
	ll := l
	_ = ll
	di := int64(1)
	_ = di
	for (di < decl_index_len(index)) {
		decl := decl_index_get(index, di)
		_ = decl
		kname := decl_kind_name(decl.Kind)
		_ = kname
		if (kname == "const") {
			ll = _lset_pos(ll, decl.Token_start)
			ll = _lskip_pub(ll)
			ll = _ladv(ll)
			if (_lk(ll) == "IDENT") {
				cname := _lcur(ll).Text
				_ = cname
				ll = _ladv(ll)
				if (_lk(ll) == ":") {
					ll = _ladv(ll)
					ll = _skip_type_l(ll)
				}
				if (_lk(ll) == "=") {
					ll = _ladv(ll)
					ll = _lskip_nl(ll)
					if (_lk(ll) == "INT") {
						val_text := _lcur(ll).Text
						_ = val_text
						num := int64(0)
						_ = num
						ci := int64(0)
						_ = ci
						neg := false
						_ = neg
						if ((int64(len(val_text)) > int64(0)) && (string(val_text[int64(0)]) == "-")) {
							neg = true
							ci = int64(1)
						}
						for (ci < int64(len(val_text))) {
							ch := string(val_text[ci])
							_ = ch
							if (ch != "_") {
								digit := int64(0)
								_ = digit
								if (ch == "0") {
									digit = int64(0)
								} else if (ch == "1") {
									digit = int64(1)
								} else if (ch == "2") {
									digit = int64(2)
								} else if (ch == "3") {
									digit = int64(3)
								} else if (ch == "4") {
									digit = int64(4)
								} else if (ch == "5") {
									digit = int64(5)
								} else if (ch == "6") {
									digit = int64(6)
								} else if (ch == "7") {
									digit = int64(7)
								} else if (ch == "8") {
									digit = int64(8)
								} else if (ch == "9") {
									digit = int64(9)
								}
								num = ((num * int64(10)) + digit)
							}
							ci = (ci + int64(1))
						}
						if neg {
							num = (int64(0) - num)
						}
						ll = _ladd_const(ll, cname, num)
						ll = _ladv(ll)
					} else if (_lk(ll) == "STRING") {
						sval := _lcur(ll).Text
						_ = sval
						ll = _ladd_const_str(ll, cname, sval)
						ll = _ladv(ll)
					} else if (_lk(ll) == "-") {
						ll = _ladv(ll)
						if (_lk(ll) == "INT") {
							val_text := _lcur(ll).Text
							_ = val_text
							num := int64(0)
							_ = num
							ci := int64(0)
							_ = ci
							for (ci < int64(len(val_text))) {
								ch := string(val_text[ci])
								_ = ch
								if (ch != "_") {
									digit := int64(0)
									_ = digit
									if (ch == "0") {
										digit = int64(0)
									} else if (ch == "1") {
										digit = int64(1)
									} else if (ch == "2") {
										digit = int64(2)
									} else if (ch == "3") {
										digit = int64(3)
									} else if (ch == "4") {
										digit = int64(4)
									} else if (ch == "5") {
										digit = int64(5)
									} else if (ch == "6") {
										digit = int64(6)
									} else if (ch == "7") {
										digit = int64(7)
									} else if (ch == "8") {
										digit = int64(8)
									} else if (ch == "9") {
										digit = int64(9)
									}
									num = ((num * int64(10)) + digit)
								}
								ci = (ci + int64(1))
							}
							num = (int64(0) - num)
							ll = Lowerer{Tokens: ll.Tokens, Pos: ll.Pos, Store: ll.Store, Registry: ll.Registry, Treg: ll.Treg, Pool: ll.Pool, Table: ll.Table, File: ll.File, Module: ll.Module, Current_func: ll.Current_func, Temp_counter: ll.Temp_counter, Label_counter: ll.Label_counter, Env_names: ll.Env_names, Env_slots: ll.Env_slots, Loop_start: ll.Loop_start, Loop_end: ll.Loop_end, Const_names: append(ll.Const_names, cname), Const_vals: append(ll.Const_vals, num), Const_str_names: ll.Const_str_names, Const_str_vals: ll.Const_str_vals, Temp_types: ll.Temp_types, Fn_error_type: ll.Fn_error_type, Defer_starts: ll.Defer_starts, Defer_ends: ll.Defer_ends}
							ll = _ladv(ll)
						}
					}
				}
			}
		}
		di = (di + int64(1))
	}
	return ll
}

func _lower_bodies(l Lowerer, index DeclIndex) Lowerer {
	ll := l
	_ = ll
	di := int64(1)
	_ = di
	for (di < decl_index_len(index)) {
		decl := decl_index_get(index, di)
		_ = decl
		kname := decl_kind_name(decl.Kind)
		_ = kname
		ll = _lset_pos(ll, decl.Token_start)
		ll = _lskip_pub(ll)
		if (_lk(ll) == "@") {
			ll = _ladv(ll)
			if (_lk(ll) == "IDENT") {
				ll = _ladv(ll)
			}
			if (_lk(ll) == "(") {
				ll = _ladv(ll)
				for ((_lk(ll) != ")") && (_lk(ll) != "EOF")) {
					ll = _ladv(ll)
				}
				if (_lk(ll) == ")") {
					ll = _ladv(ll)
				}
			}
			ll = _lskip_nl(ll)
		}
		k := _lk(ll)
		_ = k
		if ((kname == "fn") && (k == "fn")) {
			if (decl.Name == "_extern") {
				ll = _lower_extern(ll)
			} else if strings.HasPrefix(decl.Name, "_fixture:") {
				ll = _lower_fixture(ll, _lk(ll))
			} else {
				ll = _lower_fn_d(ll, decl)
			}
		} else if (kname == "entry") {
			ll = _lower_entry(ll)
		} else if (kname == "test") {
			if strings.HasPrefix(decl.Name, "_bench:") {
				ll = _lower_bench(ll)
			} else {
				ll = _lower_test(ll)
			}
		} else if (kname == "impl") {
			ll = _lower_impl_d(ll, decl)
		} else {
			ll = _lset_pos(ll, decl.Body_end)
		}
		di = (di + int64(1))
	}
	return ll
}

func _lower_bodies_skip_entry(l Lowerer, index DeclIndex) Lowerer {
	ll := l
	_ = ll
	di := int64(1)
	_ = di
	for (di < decl_index_len(index)) {
		decl := decl_index_get(index, di)
		_ = decl
		kname := decl_kind_name(decl.Kind)
		_ = kname
		ll = _lset_pos(ll, decl.Token_start)
		ll = _lskip_pub(ll)
		if (_lk(ll) == "@") {
			ll = _ladv(ll)
			if (_lk(ll) == "IDENT") {
				ll = _ladv(ll)
			}
			if (_lk(ll) == "(") {
				ll = _ladv(ll)
				for ((_lk(ll) != ")") && (_lk(ll) != "EOF")) {
					ll = _ladv(ll)
				}
				if (_lk(ll) == ")") {
					ll = _ladv(ll)
				}
			}
			ll = _lskip_nl(ll)
		}
		k := _lk(ll)
		_ = k
		if ((kname == "fn") && (k == "fn")) {
			if (decl.Name == "_extern") {
				ll = _lower_extern(ll)
			} else if strings.HasPrefix(decl.Name, "_fixture:") {
				ll = _lower_fixture(ll, _lk(ll))
			} else {
				ll = _lower_fn_d(ll, decl)
			}
		} else if (kname == "entry") {
			ll = _lset_pos(ll, decl.Body_end)
		} else if (kname == "test") {
			if strings.HasPrefix(decl.Name, "_bench:") {
				ll = _lower_bench(ll)
			} else {
				ll = _lower_test(ll)
			}
		} else if (kname == "impl") {
			ll = _lower_impl_d(ll, decl)
		} else {
			ll = _lset_pos(ll, decl.Body_end)
		}
		di = (di + int64(1))
	}
	return ll
}

func _skip_to_next_l(l Lowerer) Lowerer {
	ll := _ladv(l)
	_ = ll
	lb := "{"
	_ = lb
	rb := "}"
	_ = rb
	for (_lk(ll) != "EOF") {
		if (_lk(ll) == lb) {
			d := int64(1)
			_ = d
			ll = _ladv(ll)
			for ((d > int64(0)) && (_lk(ll) != "EOF")) {
				if (_lk(ll) == lb) {
					d = (d + int64(1))
				}
				if (_lk(ll) == rb) {
					d = (d - int64(1))
				}
				if (d > int64(0)) {
					ll = _ladv(ll)
				}
			}
			if (_lk(ll) == rb) {
				ll = _ladv(ll)
			}
			return ll
		}
		if (_lk(ll) == "NEWLINE") {
			ll = _lskip_nl(ll)
			nk := _lk(ll)
			_ = nk
			if (((((((((((nk == "fn") || (nk == "pub")) || (nk == "type")) || (nk == "struct")) || (nk == "enum")) || (nk == "trait")) || (nk == "impl")) || (nk == "const")) || (nk == "entry")) || (nk == "test")) || (nk == "EOF")) {
				return ll
			}
		}
		ll = _ladv(ll)
	}
	return ll
}

func _lower_fn_d(l Lowerer, decl DeclInfo) Lowerer {
	name := decl.Name
	_ = name
	ll := l
	_ = ll
	sig := reg_find_fn(ll.Registry, name)
	_ = sig
	if (sig.Name == "") {
		return _lset_pos(ll, decl.Body_end)
	}
	if fn_is_generic(sig) {
		return _lset_pos(ll, decl.Body_end)
	}
	actual_param_count := int64(0)
	_ = actual_param_count
	pi := int64(1)
	_ = pi
	for (pi < int64(len(sig.Param_names))) {
		if ((pi <= (int64(len(sig.Param_types)) - int64(1))) && (sig.Param_types[pi] == int64(12))) {
			actual_param_count = (actual_param_count + int64(2))
		} else {
			actual_param_count = (actual_param_count + int64(1))
		}
		pi = (pi + int64(1))
	}
	ll = _lset_func(ll, new_ir_func(name, actual_param_count, sig.Return_type))
	if (sig.Error_type > int64(0)) {
		ll = _lset_error_type(ll, sig.Error_type)
	}
	pi = int64(1)
	for (pi < int64(len(sig.Param_names))) {
		if ((pi <= (int64(len(sig.Param_types)) - int64(1))) && (sig.Param_types[pi] == int64(12))) {
			nt := _lnew_str_temp(ll)
			_ = nt
			ll = nt.L
			ll = _lenv_add(ll, sig.Param_names[pi], nt.Temp)
			ll = _lset_type(ll, nt.Temp, int64(12))
		} else {
			nt := _lnew_temp(ll)
			_ = nt
			ll = nt.L
			ll = _lenv_add(ll, sig.Param_names[pi], nt.Temp)
			if ((pi <= (int64(len(sig.Param_types)) - int64(1))) && (sig.Param_types[pi] > int64(0))) {
				pt := sig.Param_types[pi]
				_ = pt
				if _is_str_array_type(ll.Store, pt) {
					ll = _lset_type(ll, nt.Temp, int64(13))
				} else if _is_trait_object_type(ll.Store, pt) {
					ll = _lset_type(ll, nt.Temp, int64(500))
				} else {
					ll = _lset_type(ll, nt.Temp, pt)
				}
			}
		}
		pi = (pi + int64(1))
	}
	if (decl.Node_idx > int64(0)) {
		fn_node := pool_get(ll.Pool, decl.Node_idx)
		_ = fn_node
		if (fn_node.C0 > int64(0)) {
			return _lower_fn_body_ast(ll, fn_node.C0, sig)
		}
	}
	ll = _lset_pos(ll, decl.Body_start)
	return _lower_fn_body_d(ll, sig)
}

func _lower_fn_body_d(l Lowerer, sig FnSig) Lowerer {
	ll := l
	_ = ll
	lb := "{"
	_ = lb
	if (_lk(ll) == "=") {
		ll = _ladv(ll)
		ll = _lskip_nl(ll)
		result := _lower_expr(ll)
		_ = result
		ll = result.L
		ll = _emit_defers(ll)
		if ((ll.Fn_error_type > int64(0)) && (_lget_type(ll, result.Temp) != int64(200))) {
			ret_type := _lget_type(ll, result.Temp)
			_ = ret_type
			ok_slots := int64(2)
			_ = ok_slots
			if (ret_type == int64(12)) {
				ok_slots = int64(3)
			}
			ok_nt := _lnew_temp(ll)
			_ = ok_nt
			ll = ok_nt.L
			ll = _lemit(ll, new_inst(IrOpOpAlloc{}, ok_nt.Temp, ok_slots, int64(0), "_Result", int64(0)))
			tag_nt := _lnew_temp(ll)
			_ = tag_nt
			ll = tag_nt.L
			ll = _lemit(ll, new_inst(IrOpOpConst{}, tag_nt.Temp, int64(0), int64(0), "", int64(0)))
			ll = _lemit(ll, new_inst(IrOpOpFieldSet{}, ok_nt.Temp, tag_nt.Temp, int64(0), "_tag", int64(0)))
			if (ret_type == int64(12)) {
				ll = _lemit(ll, new_inst(IrOpOpFieldSet{}, ok_nt.Temp, result.Temp, int64(1), "_val", int64(0)))
				val_len := (result.Temp + int64(1))
				_ = val_len
				ll = _lemit(ll, new_inst(IrOpOpFieldSet{}, ok_nt.Temp, val_len, int64(2), "_val_len", int64(0)))
			} else {
				ll = _lemit(ll, new_inst(IrOpOpFieldSet{}, ok_nt.Temp, result.Temp, int64(1), "_val", int64(0)))
			}
			ll = _lemit(ll, new_inst(IrOpOpRet{}, int64(0), ok_nt.Temp, int64(0), "", int64(0)))
		} else {
			ret_type := _lget_type(ll, result.Temp)
			_ = ret_type
			ll = _lemit(ll, new_inst(IrOpOpRet{}, int64(0), result.Temp, int64(0), "", ret_type))
		}
		ll = _lfinish_func(ll)
	} else if (_lk(ll) == lb) {
		ll = _lower_fn_body(ll, sig.Return_type)
		ll = _lfinish_func(ll)
	} else {
		ll = _lskip_nl(ll)
		if (_lk(ll) == lb) {
			ll = _lower_fn_body(ll, sig.Return_type)
			ll = _lfinish_func(ll)
		} else {
			ll = _lfinish_func(ll)
		}
	}
	return ll
}

func _find_tok_pos(l Lowerer, offset int64) int64 {
	p := l.Pos
	_ = p
	if ((p < int64(len(l.Tokens))) && (l.Tokens[p].Offset <= offset)) {
		for (p < int64(len(l.Tokens))) {
			if (l.Tokens[p].Offset >= offset) {
				return p
			}
			p = (p + int64(1))
		}
	}
	p = int64(0)
	for (p < int64(len(l.Tokens))) {
		if (l.Tokens[p].Offset >= offset) {
			return p
		}
		p = (p + int64(1))
	}
	return p
}

func _expr_start_offset(pool NodePool, idx int64) int64 {
	if (idx <= int64(0)) {
		return int64(0)
	}
	node := pool_get(pool, idx)
	_ = node
	nk := expr_kind_name(node.Kind)
	_ = nk
	if ((((nk == "Binary") || (nk == "Pipeline")) || (nk == "Range")) || (nk == "Assign")) {
		return _expr_start_offset(pool, node.C0)
	}
	if (((nk == "Unary") || (nk == "Propagate")) || (nk == "AssertOk")) {
		return _expr_start_offset(pool, node.C0)
	}
	if (nk == "Call") {
		return _expr_start_offset(pool, node.C0)
	}
	if (((nk == "MethodCall") || (nk == "FieldAccess")) || (nk == "Index")) {
		return _expr_start_offset(pool, node.C0)
	}
	if (nk == "Catch") {
		return _expr_start_offset(pool, node.C0)
	}
	return node.Span.Offset
}

func _lower_fn_body_ast(l Lowerer, body_idx int64, sig FnSig) Lowerer {
	ll := l
	_ = ll
	body := pool_get(ll.Pool, body_idx)
	_ = body
	nk := expr_kind_name(body.Kind)
	_ = nk
	if (nk == "Block") {
		result := _lower_block_node_result(ll, body)
		_ = result
		ll = result.L
		last_temp := result.Temp
		_ = last_temp
		if (last_temp >= int64(0)) {
			inst_count := int64(len(ll.Current_func.Insts))
			_ = inst_count
			already_returned := false
			_ = already_returned
			if (inst_count > int64(1)) {
				last_op := ir_op_name(ll.Current_func.Insts[(inst_count - int64(1))].Op)
				_ = last_op
				if ((last_op == "Ret") || (last_op == "RetVoid")) {
					already_returned = true
				}
			}
			if (already_returned == false) {
				ll = _emit_defers(ll)
				if ((ll.Fn_error_type > int64(0)) && (_lget_type(ll, last_temp) != int64(200))) {
					ret_type := _lget_type(ll, last_temp)
					_ = ret_type
					ok_slots := int64(2)
					_ = ok_slots
					if (ret_type == int64(12)) {
						ok_slots = int64(3)
					}
					ok_nt := _lnew_temp(ll)
					_ = ok_nt
					ll = ok_nt.L
					ll = _lemit(ll, new_inst(IrOpOpAlloc{}, ok_nt.Temp, ok_slots, int64(0), "_Result", int64(0)))
					tag_nt := _lnew_temp(ll)
					_ = tag_nt
					ll = tag_nt.L
					ll = _lemit(ll, new_inst(IrOpOpConst{}, tag_nt.Temp, int64(0), int64(0), "", int64(0)))
					ll = _lemit(ll, new_inst(IrOpOpFieldSet{}, ok_nt.Temp, tag_nt.Temp, int64(0), "_tag", int64(0)))
					if (ret_type == int64(12)) {
						ll = _lemit(ll, new_inst(IrOpOpFieldSet{}, ok_nt.Temp, last_temp, int64(1), "_val", int64(0)))
						val_len := (last_temp + int64(1))
						_ = val_len
						ll = _lemit(ll, new_inst(IrOpOpFieldSet{}, ok_nt.Temp, val_len, int64(2), "_val_len", int64(0)))
					} else {
						ll = _lemit(ll, new_inst(IrOpOpFieldSet{}, ok_nt.Temp, last_temp, int64(1), "_val", int64(0)))
					}
					ll = _lemit(ll, new_inst(IrOpOpRet{}, int64(0), ok_nt.Temp, int64(0), "", int64(0)))
				} else {
					ret_type := _lget_type(ll, last_temp)
					_ = ret_type
					ll = _lemit(ll, new_inst(IrOpOpRet{}, int64(0), last_temp, int64(0), "", ret_type))
				}
			}
		}
	} else {
		result := _lower_expr_node(ll, body_idx)
		_ = result
		ll = result.L
		ll = _emit_defers(ll)
		if ((ll.Fn_error_type > int64(0)) && (_lget_type(ll, result.Temp) != int64(200))) {
			ret_type := _lget_type(ll, result.Temp)
			_ = ret_type
			ok_slots := int64(2)
			_ = ok_slots
			if (ret_type == int64(12)) {
				ok_slots = int64(3)
			}
			ok_nt := _lnew_temp(ll)
			_ = ok_nt
			ll = ok_nt.L
			ll = _lemit(ll, new_inst(IrOpOpAlloc{}, ok_nt.Temp, ok_slots, int64(0), "_Result", int64(0)))
			tag_nt := _lnew_temp(ll)
			_ = tag_nt
			ll = tag_nt.L
			ll = _lemit(ll, new_inst(IrOpOpConst{}, tag_nt.Temp, int64(0), int64(0), "", int64(0)))
			ll = _lemit(ll, new_inst(IrOpOpFieldSet{}, ok_nt.Temp, tag_nt.Temp, int64(0), "_tag", int64(0)))
			if (ret_type == int64(12)) {
				ll = _lemit(ll, new_inst(IrOpOpFieldSet{}, ok_nt.Temp, result.Temp, int64(1), "_val", int64(0)))
				val_len := (result.Temp + int64(1))
				_ = val_len
				ll = _lemit(ll, new_inst(IrOpOpFieldSet{}, ok_nt.Temp, val_len, int64(2), "_val_len", int64(0)))
			} else {
				ll = _lemit(ll, new_inst(IrOpOpFieldSet{}, ok_nt.Temp, result.Temp, int64(1), "_val", int64(0)))
			}
			ll = _lemit(ll, new_inst(IrOpOpRet{}, int64(0), ok_nt.Temp, int64(0), "", int64(0)))
		} else {
			ret_type := _lget_type(ll, result.Temp)
			_ = ret_type
			ll = _lemit(ll, new_inst(IrOpOpRet{}, int64(0), result.Temp, int64(0), "", ret_type))
		}
	}
	ll = _lfinish_func(ll)
	return ll
}

func _lower_block_node_result(l Lowerer, node Expr) LT {
	ll := l
	_ = ll
	scope_start := int64(len(ll.Env_names))
	_ = scope_start
	last_temp := -int64(1)
	_ = last_temp
	si := int64(0)
	_ = si
	for (si < node.List_count) {
		idx_node := pool_get(ll.Pool, (node.List_start + si))
		_ = idx_node
		child := pool_get(ll.Pool, idx_node.C0)
		_ = child
		child_kind := expr_kind_name(child.Kind)
		_ = child_kind
		is_last := (si == (node.List_count - int64(1)))
		_ = is_last
		tp := _find_tok_pos(ll, _expr_start_offset(ll.Pool, idx_node.C0))
		_ = tp
		ll = _lset_pos(ll, tp)
		if (child_kind == "Return") {
			if child.B1 {
				r := _lower_expr_node(ll, child.C0)
				_ = r
				ll = r.L
				ll = _emit_defers(ll)
				ret_type := _lget_type(ll, r.Temp)
				_ = ret_type
				ll = _lemit(ll, new_inst(IrOpOpRet{}, int64(0), r.Temp, int64(0), "", ret_type))
			} else {
				ll = _emit_defers(ll)
				ll = _lemit(ll, new_inst(IrOpOpRetVoid{}, int64(0), int64(0), int64(0), "", int64(0)))
			}
			last_temp = -int64(1)
		} else if is_last {
			r := _lower_expr_node(ll, idx_node.C0)
			_ = r
			ll = r.L
			last_temp = r.Temp
		} else {
			r := _lower_expr_node(ll, idx_node.C0)
			_ = r
			ll = r.L
		}
		ll = _lskip_nl(ll)
		si = (si + int64(1))
	}
	ll = _emit_drops_scoped(ll, scope_start)
	return LT{L: ll, Temp: last_temp}
}

func _lower_expr_node(l Lowerer, idx int64) LT {
	if (idx <= int64(0)) {
		return LT{L: l, Temp: -int64(1)}
	}
	node := pool_get(l.Pool, idx)
	_ = node
	nk := expr_kind_name(node.Kind)
	_ = nk
	ll := l
	_ = ll
	if (nk == "IntLit") {
		num := _parse_int_literal(node.S1)
		_ = num
		nt := _lnew_temp(ll)
		_ = nt
		ll = nt.L
		ll = _lemit(ll, new_inst(IrOpOpConst{}, nt.Temp, num, int64(0), "", int64(0)))
		ll = _lset_type(ll, nt.Temp, int64(1))
		return LT{L: ll, Temp: nt.Temp}
	}
	if (nk == "FloatLit") {
		val := node.S1
		_ = val
		str_idx := mod_find_string(ll.Module, val)
		_ = str_idx
		if (str_idx == int64(0)) {
			new_mod := mod_add_string(ll.Module, val)
			_ = new_mod
			str_idx = (int64(len(new_mod.String_constants)) - int64(1))
			ll = _lset_mod(ll, new_mod)
		}
		nt := _lnew_temp(ll)
		_ = nt
		ll = nt.L
		ll = _lemit(ll, new_inst(IrOpOpConstStr{}, nt.Temp, str_idx, (int64(0) - int64(9)), "", int64(0)))
		ll = _lset_type(ll, nt.Temp, int64(9))
		return LT{L: ll, Temp: nt.Temp}
	}
	if (nk == "StringLit") {
		val := node.S1
		_ = val
		if ((int64(len(val)) == int64(0)) && (node.Span.Offset > int64(0))) {
			tp := _find_tok_pos(ll, node.Span.Offset)
			_ = tp
			if ((tp < int64(len(ll.Tokens))) && (token_name(ll.Tokens[tp].Kind) == "STRING_START")) {
				ll = _lset_pos(ll, tp)
				return _lower_interpolated_string(ll)
			}
		}
		return _lower_string_const(ll, val)
	}
	if (nk == "BoolLit") {
		nt := _lnew_temp(ll)
		_ = nt
		ll = nt.L
		bval := int64(0)
		_ = bval
		if node.B1 {
			bval = int64(1)
		}
		ll = _lemit(ll, new_inst(IrOpOpConst{}, nt.Temp, bval, int64(0), "", int64(0)))
		return LT{L: ll, Temp: nt.Temp}
	}
	if (nk == "None") {
		nt := _lnew_temp(ll)
		_ = nt
		ll = nt.L
		ll = _lemit(ll, new_inst(IrOpOpAlloc{}, nt.Temp, int64(2), int64(0), "_Optional", int64(0)))
		tag_nt := _lnew_temp(ll)
		_ = tag_nt
		ll = tag_nt.L
		ll = _lemit(ll, new_inst(IrOpOpConst{}, tag_nt.Temp, int64(0), int64(0), "", int64(0)))
		ll = _lemit(ll, new_inst(IrOpOpFieldSet{}, nt.Temp, tag_nt.Temp, int64(0), "_tag", int64(0)))
		ll = _lset_type(ll, nt.Temp, int64(201))
		return LT{L: ll, Temp: nt.Temp}
	}
	if (nk == "Ident") {
		name := node.S1
		_ = name
		slot := _lenv_lookup(ll, name)
		_ = slot
		if (slot >= int64(0)) {
			src_type := _lget_type(ll, slot)
			_ = src_type
			if (src_type == int64(12)) {
				nt := _lnew_str_temp(ll)
				_ = nt
				ll = nt.L
				ll = _lemit(ll, new_inst(IrOpOpLoad{}, nt.Temp, slot, int64(0), "", int64(0)))
				slot_len := (slot + int64(1))
				_ = slot_len
				nt_len := (nt.Temp + int64(1))
				_ = nt_len
				ll = _lemit(ll, new_inst(IrOpOpLoad{}, nt_len, slot_len, int64(0), "", int64(0)))
				ll = _lset_type(ll, nt.Temp, int64(12))
				return LT{L: ll, Temp: nt.Temp}
			}
			nt := _lnew_temp(ll)
			_ = nt
			ll = nt.L
			ll = _lemit(ll, new_inst(IrOpOpLoad{}, nt.Temp, slot, int64(0), "", int64(0)))
			if (src_type != int64(0)) {
				ll = _lset_type(ll, nt.Temp, src_type)
			}
			return LT{L: ll, Temp: nt.Temp}
		}
		cv := _lconst_lookup(ll, name)
		_ = cv
		if (cv != (int64(0) - int64(9999999))) {
			nt := _lnew_temp(ll)
			_ = nt
			ll = nt.L
			ll = _lemit(ll, new_inst(IrOpOpConst{}, nt.Temp, cv, int64(0), "", int64(0)))
			ll = _lset_type(ll, nt.Temp, int64(1))
			return LT{L: ll, Temp: nt.Temp}
		}
		csv := _lconst_str_lookup(ll, name)
		_ = csv
		if (int64(len(csv)) > int64(0)) {
			return _lower_string_const(ll, csv)
		}
		vtag := _variant_tag(ll.Registry, name)
		_ = vtag
		if (vtag >= int64(0)) {
			vparent := _variant_parent(ll.Registry, name)
			_ = vparent
			if sum_has_data(vparent) {
				alloc_sz := sum_alloc_size(vparent)
				_ = alloc_sz
				nt := _lnew_temp(ll)
				_ = nt
				ll = nt.L
				ll = _lemit(ll, new_inst(IrOpOpAlloc{}, nt.Temp, alloc_sz, int64(0), vparent.Name, int64(0)))
				st_id := _struct_type_id(ll.Registry, vparent.Name)
				_ = st_id
				if (st_id > int64(0)) {
					ll = _lset_type(ll, nt.Temp, st_id)
				}
				vtag_nt := _lnew_temp(ll)
				_ = vtag_nt
				ll = vtag_nt.L
				ll = _lemit(ll, new_inst(IrOpOpConst{}, vtag_nt.Temp, vtag, int64(0), "", int64(0)))
				ll = _lemit(ll, new_inst(IrOpOpFieldSet{}, nt.Temp, vtag_nt.Temp, int64(0), "_tag", int64(0)))
				return LT{L: ll, Temp: nt.Temp}
			}
			nt := _lnew_temp(ll)
			_ = nt
			ll = nt.L
			ll = _lemit(ll, new_inst(IrOpOpConst{}, nt.Temp, vtag, int64(0), "", int64(0)))
			return LT{L: ll, Temp: nt.Temp}
		}
		fsig := reg_find_fn(ll.Registry, name)
		_ = fsig
		if (fsig.Name != "") {
			ll = _lset_reg(ll, reg_add_fn_value_ref(ll.Registry, name))
			clos := _lnew_temp(ll)
			_ = clos
			ll = clos.L
			ll = _lemit(ll, new_inst(IrOpOpAlloc{}, clos.Temp, int64(2), int64(0), "_Closure", int64(0)))
			fn_ref := _lnew_temp(ll)
			_ = fn_ref
			ll = fn_ref.L
			ll = _lemit(ll, new_inst(IrOpOpFnRef{}, fn_ref.Temp, int64(0), int64(0), ("_tramp_" + name), int64(0)))
			ll = _lemit(ll, new_inst(IrOpOpFieldSet{}, clos.Temp, fn_ref.Temp, int64(0), "_fn", int64(0)))
			null_env := _lnew_temp(ll)
			_ = null_env
			ll = null_env.L
			ll = _lemit(ll, new_inst(IrOpOpConst{}, null_env.Temp, int64(0), int64(0), "", int64(0)))
			ll = _lemit(ll, new_inst(IrOpOpFieldSet{}, clos.Temp, null_env.Temp, int64(1), "_env", int64(0)))
			return LT{L: ll, Temp: clos.Temp}
		}
		nt := _lnew_temp(ll)
		_ = nt
		ll = nt.L
		return LT{L: ll, Temp: nt.Temp}
	}
	if (nk == "Unary") {
		operand := _lower_expr_node(ll, node.C0)
		_ = operand
		ll = operand.L
		nt := _lnew_temp(ll)
		_ = nt
		ll = nt.L
		op_type := _lget_type(ll, operand.Temp)
		_ = op_type
		if ((node.S1 == "-") && _is_float_type(op_type)) {
			ll = _lemit(ll, new_inst(IrOpOpNeg{}, nt.Temp, operand.Temp, int64(0), "f64", int64(0)))
			ll = _lset_type(ll, nt.Temp, int64(9))
		} else {
			ll = _lemit(ll, new_inst(_unary_op(node.S1), nt.Temp, operand.Temp, int64(0), "", int64(0)))
		}
		return LT{L: ll, Temp: nt.Temp}
	}
	if (nk == "Binary") {
		op := node.S1
		_ = op
		if (op == ":=") {
			lhs := pool_get(ll.Pool, node.C0)
			_ = lhs
			result := _lower_expr_node(ll, node.C1)
			_ = result
			ll = result.L
			if (expr_kind_name(lhs.Kind) == "Ident") {
				ll = _lenv_add(ll, lhs.S1, result.Temp)
			}
			return LT{L: ll, Temp: result.Temp}
		}
		if (op == "=") {
			return _lower_assign_node(ll, node)
		}
		if (((((op == "+=") || (op == "-=")) || (op == "*=")) || (op == "/=")) || (op == "%=")) {
			lhs := pool_get(ll.Pool, node.C0)
			_ = lhs
			right := _lower_expr_node(ll, node.C1)
			_ = right
			ll = right.L
			if (expr_kind_name(lhs.Kind) == "Ident") {
				slot := _lenv_lookup(ll, lhs.S1)
				_ = slot
				if (slot >= int64(0)) {
					cur := _lnew_temp(ll)
					_ = cur
					ll = cur.L
					ll = _lemit(ll, new_inst(IrOpOpLoad{}, cur.Temp, slot, int64(0), "", int64(0)))
					res := _lnew_temp(ll)
					_ = res
					ll = res.L
					arith_str := "+"
					_ = arith_str
					if (op == "-=") {
						arith_str = "-"
					}
					if (op == "*=") {
						arith_str = "*"
					}
					if (op == "/=") {
						arith_str = "/"
					}
					if (op == "%=") {
						arith_str = "%"
					}
					left_type := _lget_type(ll, slot)
					_ = left_type
					if (_is_float_type(left_type) || _is_float_type(_lget_type(ll, right.Temp))) {
						ll = _lemit(ll, new_inst(_arith_op(arith_str), res.Temp, cur.Temp, right.Temp, "f64", int64(0)))
						ll = _lset_type(ll, res.Temp, int64(9))
					} else {
						ll = _lemit(ll, new_inst(_arith_op(arith_str), res.Temp, cur.Temp, right.Temp, "", int64(0)))
					}
					ll = _lemit(ll, new_inst(IrOpOpStore{}, slot, res.Temp, int64(0), "", int64(0)))
					return LT{L: ll, Temp: res.Temp}
				}
			}
			return LT{L: ll, Temp: -int64(1)}
		}
		left := _lower_expr_node(ll, node.C0)
		_ = left
		ll = left.L
		right := _lower_expr_node(ll, node.C1)
		_ = right
		ll = right.L
		left_type := _lget_type(ll, left.Temp)
		_ = left_type
		right_type := _lget_type(ll, right.Temp)
		_ = right_type
		if (((((op == "+") || (op == "-")) || (op == "*")) || (op == "/")) || (op == "%")) {
			if ((op == "+") && (left_type == int64(12))) {
				nt := _lnew_str_temp(ll)
				_ = nt
				ll = nt.L
				ll = _lemit(ll, new_inst(IrOpOpStrConcat{}, nt.Temp, left.Temp, right.Temp, "", int64(0)))
				ll = _lset_type(ll, nt.Temp, int64(12))
				return LT{L: ll, Temp: nt.Temp}
			}
			if (_is_float_type(left_type) || _is_float_type(right_type)) {
				nt := _lnew_temp(ll)
				_ = nt
				ll = nt.L
				ll = _lemit(ll, new_inst(_arith_op(op), nt.Temp, left.Temp, right.Temp, "f64", int64(0)))
				ll = _lset_type(ll, nt.Temp, int64(9))
				return LT{L: ll, Temp: nt.Temp}
			}
			nt := _lnew_temp(ll)
			_ = nt
			ll = nt.L
			ll = _lemit(ll, new_inst(_arith_op(op), nt.Temp, left.Temp, right.Temp, "", int64(0)))
			ll = _lset_type(ll, nt.Temp, int64(1))
			return LT{L: ll, Temp: nt.Temp}
		}
		if ((((((op == "==") || (op == "!=")) || (op == "<")) || (op == ">")) || (op == "<=")) || (op == ">=")) {
			nt := _lnew_temp(ll)
			_ = nt
			ll = nt.L
			if ((((op == "==") || (op == "!="))) && (left_type == int64(12))) {
				ll = _lemit(ll, new_inst(IrOpOpStrEq{}, nt.Temp, left.Temp, right.Temp, "", int64(0)))
				if (op == "!=") {
					not_nt := _lnew_temp(ll)
					_ = not_nt
					ll = not_nt.L
					ll = _lemit(ll, new_inst(IrOpOpNot{}, not_nt.Temp, nt.Temp, int64(0), "", int64(0)))
					ll = _lset_type(ll, not_nt.Temp, int64(1))
					return LT{L: ll, Temp: not_nt.Temp}
				}
				ll = _lset_type(ll, nt.Temp, int64(1))
				return LT{L: ll, Temp: nt.Temp}
			}
			if ((((op == "==") || (op == "!="))) && _is_struct_type(ll.Store, ll.Registry, left_type)) {
				sname := _struct_name_for_type(ll.Store, ll.Registry, left_type)
				_ = sname
				if (sname != "") {
					eq_result := _lower_struct_eq(ll, left.Temp, right.Temp, sname, nt.Temp)
					_ = eq_result
					ll = eq_result
					if (op == "!=") {
						not_nt := _lnew_temp(ll)
						_ = not_nt
						ll = not_nt.L
						ll = _lemit(ll, new_inst(IrOpOpNot{}, not_nt.Temp, nt.Temp, int64(0), "", int64(0)))
						ll = _lset_type(ll, not_nt.Temp, int64(1))
						return LT{L: ll, Temp: not_nt.Temp}
					}
					ll = _lset_type(ll, nt.Temp, int64(1))
					return LT{L: ll, Temp: nt.Temp}
				}
			}
			if (_is_float_type(left_type) || _is_float_type(right_type)) {
				ll = _lemit(ll, new_inst(_cmp_op(op), nt.Temp, left.Temp, right.Temp, "f64", int64(0)))
			} else {
				if (left_type == int64(12)) {
					cmp_nt := _lnew_temp(ll)
					_ = cmp_nt
					ll = cmp_nt.L
					ll = _lemit(ll, new_inst(IrOpOpStrCmp{}, cmp_nt.Temp, left.Temp, right.Temp, "", int64(0)))
					ll = _lset_type(ll, cmp_nt.Temp, int64(1))
					zero_nt := _lnew_temp(ll)
					_ = zero_nt
					ll = zero_nt.L
					ll = _lemit(ll, new_inst(IrOpOpConst{}, zero_nt.Temp, int64(0), int64(0), "", int64(0)))
					ll = _lset_type(ll, zero_nt.Temp, int64(1))
					ll = _lemit(ll, new_inst(_cmp_op(op), nt.Temp, cmp_nt.Temp, zero_nt.Temp, "", int64(0)))
				} else {
					ll = _lemit(ll, new_inst(_cmp_op(op), nt.Temp, left.Temp, right.Temp, "", int64(0)))
				}
			}
			ll = _lset_type(ll, nt.Temp, int64(1))
			return LT{L: ll, Temp: nt.Temp}
		}
		if (op == "&&") {
			nt := _lnew_temp(ll)
			_ = nt
			ll = nt.L
			ll = _lemit(ll, new_inst(IrOpOpAnd{}, nt.Temp, left.Temp, right.Temp, "", int64(0)))
			return LT{L: ll, Temp: nt.Temp}
		}
		if (op == "||") {
			nt := _lnew_temp(ll)
			_ = nt
			ll = nt.L
			ll = _lemit(ll, new_inst(IrOpOpOr{}, nt.Temp, left.Temp, right.Temp, "", int64(0)))
			return LT{L: ll, Temp: nt.Temp}
		}
		if (op == "|>") {
			right_node := pool_get(ll.Pool, node.C1)
			_ = right_node
			if (expr_kind_name(right_node.Kind) == "Ident") {
				fname := right_node.S1
				_ = fname
				fsig := reg_find_fn(ll.Registry, fname)
				_ = fsig
				if (fsig.Name != "") {
					reg_count := int64(1)
					_ = reg_count
					if (left_type == int64(12)) {
						reg_count = int64(2)
					}
					nt_temp := int64(0)
					_ = nt_temp
					if (fsig.Return_type == int64(12)) {
						nt := _lnew_str_temp(ll)
						_ = nt
						ll = nt.L
						nt_temp = nt.Temp
					} else {
						nt := _lnew_temp(ll)
						_ = nt
						ll = nt.L
						nt_temp = nt.Temp
					}
					ll = _lemit(ll, new_inst(IrOpOpCall{}, nt_temp, left.Temp, reg_count, fname, fsig.Return_type))
					if (fsig.Return_type > int64(0)) {
						ll = _lset_type(ll, nt_temp, fsig.Return_type)
					}
					return LT{L: ll, Temp: nt_temp}
				}
			}
			return LT{L: ll, Temp: left.Temp}
		}
		if (((((op == "&") || (op == "|")) || (op == "^")) || (op == "<<")) || (op == ">>")) {
			nt := _lnew_temp(ll)
			_ = nt
			ll = nt.L
			ll = _lemit(ll, new_inst(_arith_op(op), nt.Temp, left.Temp, right.Temp, "", int64(0)))
			ll = _lset_type(ll, nt.Temp, int64(1))
			return LT{L: ll, Temp: nt.Temp}
		}
		return LT{L: ll, Temp: left.Temp}
	}
	if (nk == "Block") {
		return _lower_block_node_result(ll, node)
	}
	if (nk == "If") {
		cond := _lower_expr_node(ll, node.C0)
		_ = cond
		ll = cond.L
		else_lbl := _lnew_label(ll)
		_ = else_lbl
		ll = else_lbl.L
		end_lbl := _lnew_label(ll)
		_ = end_lbl
		ll = end_lbl.L
		result_nt := _lnew_str_temp(ll)
		_ = result_nt
		ll = result_nt.L
		ll = _lemit(ll, new_inst(IrOpOpBranchFalse{}, int64(0), cond.Temp, else_lbl.Temp, "", int64(0)))
		then_result := _lower_expr_node(ll, node.C1)
		_ = then_result
		ll = then_result.L
		if (then_result.Temp >= int64(0)) {
			ll = _lemit(ll, new_inst(IrOpOpStore{}, result_nt.Temp, then_result.Temp, int64(0), "", int64(0)))
			then_len_dst := (result_nt.Temp + int64(1))
			_ = then_len_dst
			then_len_src := (then_result.Temp + int64(1))
			_ = then_len_src
			ll = _lemit(ll, new_inst(IrOpOpStore{}, then_len_dst, then_len_src, int64(0), "", int64(0)))
		}
		ll = _lemit(ll, new_inst(IrOpOpJump{}, int64(0), end_lbl.Temp, int64(0), "", int64(0)))
		ll = _lemit(ll, new_inst(IrOpOpLabel{}, int64(0), else_lbl.Temp, int64(0), "", int64(0)))
		if node.B1 {
			else_result := _lower_expr_node(ll, node.C2)
			_ = else_result
			ll = else_result.L
			if (else_result.Temp >= int64(0)) {
				ll = _lemit(ll, new_inst(IrOpOpStore{}, result_nt.Temp, else_result.Temp, int64(0), "", int64(0)))
				else_len_dst := (result_nt.Temp + int64(1))
				_ = else_len_dst
				else_len_src := (else_result.Temp + int64(1))
				_ = else_len_src
				ll = _lemit(ll, new_inst(IrOpOpStore{}, else_len_dst, else_len_src, int64(0), "", int64(0)))
			}
		}
		ll = _lemit(ll, new_inst(IrOpOpLabel{}, int64(0), end_lbl.Temp, int64(0), "", int64(0)))
		return LT{L: ll, Temp: result_nt.Temp}
	}
	if (nk == "While") {
		loop_lbl := _lnew_label(ll)
		_ = loop_lbl
		ll = loop_lbl.L
		end_lbl := _lnew_label(ll)
		_ = end_lbl
		ll = end_lbl.L
		ll = _lemit(ll, new_inst(IrOpOpLabel{}, int64(0), loop_lbl.Temp, int64(0), "", int64(0)))
		cond := _lower_expr_node(ll, node.C0)
		_ = cond
		ll = cond.L
		ll = _lemit(ll, new_inst(IrOpOpBranchFalse{}, int64(0), cond.Temp, end_lbl.Temp, "", int64(0)))
		prev_start := ll.Loop_start
		_ = prev_start
		prev_end := ll.Loop_end
		_ = prev_end
		ll = _lset_loop(ll, loop_lbl.Temp, end_lbl.Temp)
		body := _lower_expr_node(ll, node.C1)
		_ = body
		ll = body.L
		ll = _lset_loop(ll, prev_start, prev_end)
		ll = _lemit(ll, new_inst(IrOpOpJump{}, int64(0), loop_lbl.Temp, int64(0), "", int64(0)))
		ll = _lemit(ll, new_inst(IrOpOpLabel{}, int64(0), end_lbl.Temp, int64(0), "", int64(0)))
		return LT{L: ll, Temp: -int64(1)}
	}
	if (nk == "Loop") {
		loop_lbl := _lnew_label(ll)
		_ = loop_lbl
		ll = loop_lbl.L
		end_lbl := _lnew_label(ll)
		_ = end_lbl
		ll = end_lbl.L
		ll = _lemit(ll, new_inst(IrOpOpLabel{}, int64(0), loop_lbl.Temp, int64(0), "", int64(0)))
		prev_start := ll.Loop_start
		_ = prev_start
		prev_end := ll.Loop_end
		_ = prev_end
		ll = _lset_loop(ll, loop_lbl.Temp, end_lbl.Temp)
		body := _lower_expr_node(ll, node.C0)
		_ = body
		ll = body.L
		ll = _lset_loop(ll, prev_start, prev_end)
		ll = _lemit(ll, new_inst(IrOpOpJump{}, int64(0), loop_lbl.Temp, int64(0), "", int64(0)))
		ll = _lemit(ll, new_inst(IrOpOpLabel{}, int64(0), end_lbl.Temp, int64(0), "", int64(0)))
		return LT{L: ll, Temp: -int64(1)}
	}
	if (nk == "For") {
		var_name := node.S1
		_ = var_name
		iter_node := pool_get(ll.Pool, node.C0)
		_ = iter_node
		if (expr_kind_name(iter_node.Kind) == "Range") {
			return _lower_for_range_node(ll, var_name, iter_node, node.C1)
		}
		iter := _lower_expr_node(ll, node.C0)
		_ = iter
		ll = iter.L
		iter_type := _lget_type(ll, iter.Temp)
		_ = iter_type
		len_nt := _lnew_temp(ll)
		_ = len_nt
		ll = len_nt.L
		ll = _lemit(ll, new_inst(IrOpOpArrayLen{}, len_nt.Temp, iter.Temp, int64(0), "", int64(0)))
		idx_nt := _lnew_temp(ll)
		_ = idx_nt
		ll = idx_nt.L
		ll = _lemit(ll, new_inst(IrOpOpConst{}, idx_nt.Temp, int64(0), int64(0), "", int64(0)))
		loop_lbl := _lnew_label(ll)
		_ = loop_lbl
		ll = loop_lbl.L
		end_lbl := _lnew_label(ll)
		_ = end_lbl
		ll = end_lbl.L
		inc_lbl := _lnew_label(ll)
		_ = inc_lbl
		ll = inc_lbl.L
		ll = _lemit(ll, new_inst(IrOpOpLabel{}, int64(0), loop_lbl.Temp, int64(0), "", int64(0)))
		cond_nt := _lnew_temp(ll)
		_ = cond_nt
		ll = cond_nt.L
		ll = _lemit(ll, new_inst(IrOpOpLt{}, cond_nt.Temp, idx_nt.Temp, len_nt.Temp, "", int64(0)))
		ll = _lemit(ll, new_inst(IrOpOpBranchFalse{}, int64(0), cond_nt.Temp, end_lbl.Temp, "", int64(0)))
		if (var_name != "") {
			if ((iter_type == int64(13)) || _is_str_array_type(ll.Store, iter_type)) {
				elem_nt := _lnew_str_temp(ll)
				_ = elem_nt
				ll = elem_nt.L
				ll = _lemit(ll, new_inst(IrOpOpArrayGet{}, elem_nt.Temp, iter.Temp, idx_nt.Temp, "", int64(12)))
				ll = _lset_type(ll, elem_nt.Temp, int64(12))
				ll = _lenv_add(ll, var_name, elem_nt.Temp)
			} else {
				elem_nt := _lnew_temp(ll)
				_ = elem_nt
				ll = elem_nt.L
				ll = _lemit(ll, new_inst(IrOpOpArrayGet{}, elem_nt.Temp, iter.Temp, idx_nt.Temp, "", int64(0)))
				elem_type := _array_elem_type(ll.Store, iter_type)
				_ = elem_type
				if (elem_type > int64(0)) {
					ll = _lset_type(ll, elem_nt.Temp, elem_type)
				}
				ll = _lenv_add(ll, var_name, elem_nt.Temp)
			}
		}
		prev_start := ll.Loop_start
		_ = prev_start
		prev_end := ll.Loop_end
		_ = prev_end
		ll = _lset_loop(ll, inc_lbl.Temp, end_lbl.Temp)
		body := _lower_expr_node(ll, node.C1)
		_ = body
		ll = body.L
		ll = _lset_loop(ll, prev_start, prev_end)
		ll = _lemit(ll, new_inst(IrOpOpLabel{}, int64(0), inc_lbl.Temp, int64(0), "", int64(0)))
		one_nt := _lnew_temp(ll)
		_ = one_nt
		ll = one_nt.L
		ll = _lemit(ll, new_inst(IrOpOpConst{}, one_nt.Temp, int64(1), int64(0), "", int64(0)))
		inc_nt := _lnew_temp(ll)
		_ = inc_nt
		ll = inc_nt.L
		ll = _lemit(ll, new_inst(IrOpOpAdd{}, inc_nt.Temp, idx_nt.Temp, one_nt.Temp, "", int64(0)))
		ll = _lemit(ll, new_inst(IrOpOpStore{}, idx_nt.Temp, inc_nt.Temp, int64(0), "", int64(0)))
		ll = _lemit(ll, new_inst(IrOpOpJump{}, int64(0), loop_lbl.Temp, int64(0), "", int64(0)))
		ll = _lemit(ll, new_inst(IrOpOpLabel{}, int64(0), end_lbl.Temp, int64(0), "", int64(0)))
		return LT{L: ll, Temp: -int64(1)}
	}
	if (nk == "Break") {
		if (ll.Loop_end >= int64(0)) {
			ll = _lemit(ll, new_inst(IrOpOpJump{}, int64(0), ll.Loop_end, int64(0), "", int64(0)))
		}
		return LT{L: ll, Temp: -int64(1)}
	}
	if (nk == "Continue") {
		if (ll.Loop_start >= int64(0)) {
			ll = _lemit(ll, new_inst(IrOpOpJump{}, int64(0), ll.Loop_start, int64(0), "", int64(0)))
		}
		return LT{L: ll, Temp: -int64(1)}
	}
	if (nk == "Return") {
		if node.B1 {
			result := _lower_expr_node(ll, node.C0)
			_ = result
			ll = result.L
			ll = _emit_defers(ll)
			if ((ll.Fn_error_type > int64(0)) && (_lget_type(ll, result.Temp) != int64(200))) {
				ret_type := _lget_type(ll, result.Temp)
				_ = ret_type
				ok_slots := int64(2)
				_ = ok_slots
				if (ret_type == int64(12)) {
					ok_slots = int64(3)
				}
				ok_nt := _lnew_temp(ll)
				_ = ok_nt
				ll = ok_nt.L
				ll = _lemit(ll, new_inst(IrOpOpAlloc{}, ok_nt.Temp, ok_slots, int64(0), "_Result", int64(0)))
				tag_nt := _lnew_temp(ll)
				_ = tag_nt
				ll = tag_nt.L
				ll = _lemit(ll, new_inst(IrOpOpConst{}, tag_nt.Temp, int64(0), int64(0), "", int64(0)))
				ll = _lemit(ll, new_inst(IrOpOpFieldSet{}, ok_nt.Temp, tag_nt.Temp, int64(0), "_tag", int64(0)))
				if (ret_type == int64(12)) {
					ll = _lemit(ll, new_inst(IrOpOpFieldSet{}, ok_nt.Temp, result.Temp, int64(1), "_val", int64(0)))
					val_len := (result.Temp + int64(1))
					_ = val_len
					ll = _lemit(ll, new_inst(IrOpOpFieldSet{}, ok_nt.Temp, val_len, int64(2), "_val_len", int64(0)))
				} else {
					ll = _lemit(ll, new_inst(IrOpOpFieldSet{}, ok_nt.Temp, result.Temp, int64(1), "_val", int64(0)))
				}
				ll = _lemit(ll, new_inst(IrOpOpRet{}, int64(0), ok_nt.Temp, int64(0), "", int64(0)))
			} else {
				ret_type := _lget_type(ll, result.Temp)
				_ = ret_type
				ll = _lemit(ll, new_inst(IrOpOpRet{}, int64(0), result.Temp, int64(0), "", ret_type))
			}
		} else {
			ll = _emit_defers(ll)
			ll = _lemit(ll, new_inst(IrOpOpRetVoid{}, int64(0), int64(0), int64(0), "", int64(0)))
		}
		return LT{L: ll, Temp: -int64(1)}
	}
	if (nk == "Binding") {
		result := _lower_expr_node(ll, node.C0)
		_ = result
		ll = result.L
		ll = _lenv_add(ll, node.S1, result.Temp)
		return LT{L: ll, Temp: result.Temp}
	}
	if (nk == "Assign") {
		return _lower_assign_node(ll, node)
	}
	if (nk == "Call") {
		return _lower_call_node(ll, node)
	}
	if (nk == "MethodCall") {
		recv := _lower_expr_node(ll, node.C0)
		_ = recv
		ll = recv.L
		if ((int64(len(node.S2)) > int64(0)) && (((node.S1 == "to") || (node.S1 == "trunc")))) {
			return _lower_conversion(ll, recv.Temp, node.S1, node.S2)
		}
		tp := _find_tok_pos(ll, node.Span.Offset)
		_ = tp
		for ((tp < int64(len(ll.Tokens))) && (token_name(ll.Tokens[tp].Kind) != "(")) {
			tp = (tp + int64(1))
		}
		ll = _lset_pos(ll, tp)
		return _lower_method_call(ll, recv.Temp, node.S1)
	}
	if (nk == "FieldAccess") {
		recv := _lower_expr_node(ll, node.C0)
		_ = recv
		ll = recv.L
		field := node.S1
		_ = field
		left_type := _lget_type(ll, recv.Temp)
		_ = left_type
		sname := _struct_name_for_type(ll.Store, ll.Registry, left_type)
		_ = sname
		fidx := int64(0)
		_ = fidx
		ft := int64(0)
		_ = ft
		if (sname != "") {
			fidx = _field_index(ll.Registry, sname, field)
			ft = _field_type_id(ll, sname, field)
		} else {
			fidx = _field_index_any(ll.Registry, field)
			ft = _field_type_id_any(ll.Registry, field)
		}
		has_field := reg_find_field(ll.Registry, sname, field)
		_ = has_field
		if ((has_field.Name == "") && (sname != "")) {
			qualified := ((sname + "_") + field)
			_ = qualified
			msig := reg_find_fn(ll.Registry, qualified)
			_ = msig
			if (msig.Name != "") {
				ll = _lset_reg(ll, reg_add_bound_fn_value_ref(ll.Registry, qualified))
				env_nt := _lnew_temp(ll)
				_ = env_nt
				ll = env_nt.L
				ll = _lemit(ll, new_inst(IrOpOpAlloc{}, env_nt.Temp, int64(1), int64(0), "_bound_env", int64(0)))
				ll = _lemit(ll, new_inst(IrOpOpFieldSet{}, env_nt.Temp, recv.Temp, int64(0), "_self", int64(0)))
				clos := _lnew_temp(ll)
				_ = clos
				ll = clos.L
				ll = _lemit(ll, new_inst(IrOpOpAlloc{}, clos.Temp, int64(2), int64(0), "_Closure", int64(0)))
				fn_ref := _lnew_temp(ll)
				_ = fn_ref
				ll = fn_ref.L
				ll = _lemit(ll, new_inst(IrOpOpFnRef{}, fn_ref.Temp, int64(0), int64(0), ("_bound_tramp_" + qualified), int64(0)))
				ll = _lemit(ll, new_inst(IrOpOpFieldSet{}, clos.Temp, fn_ref.Temp, int64(0), "_fn", int64(0)))
				ll = _lemit(ll, new_inst(IrOpOpFieldSet{}, clos.Temp, env_nt.Temp, int64(1), "_env", int64(0)))
				return LT{L: ll, Temp: clos.Temp}
			}
		}
		nt_temp := int64(0)
		_ = nt_temp
		if (ft == int64(12)) {
			nt := _lnew_str_temp(ll)
			_ = nt
			ll = nt.L
			nt_temp = nt.Temp
			fidx_len := (fidx + int64(1))
			_ = fidx_len
			nt_len := (nt_temp + int64(1))
			_ = nt_len
			ll = _lemit(ll, new_inst(IrOpOpFieldGet{}, nt_temp, recv.Temp, fidx, field, int64(0)))
			ll = _lemit(ll, new_inst(IrOpOpFieldGet{}, nt_len, recv.Temp, fidx_len, field, int64(0)))
		} else {
			nt := _lnew_temp(ll)
			_ = nt
			ll = nt.L
			nt_temp = nt.Temp
			ll = _lemit(ll, new_inst(IrOpOpFieldGet{}, nt_temp, recv.Temp, fidx, field, int64(0)))
		}
		if (ft > int64(0)) {
			ll = _lset_type(ll, nt_temp, ft)
		}
		return LT{L: ll, Temp: nt_temp}
	}
	if (nk == "Index") {
		container := _lower_expr_node(ll, node.C0)
		_ = container
		ll = container.L
		subscript := _lower_expr_node(ll, node.C1)
		_ = subscript
		ll = subscript.L
		left_type := _lget_type(ll, container.Temp)
		_ = left_type
		elem_type := _array_elem_type(ll.Store, left_type)
		_ = elem_type
		if (elem_type == int64(12)) {
			nt := _lnew_str_temp(ll)
			_ = nt
			ll = nt.L
			ll = _lemit(ll, new_inst(IrOpOpArrayGet{}, nt.Temp, container.Temp, subscript.Temp, "", int64(12)))
			ll = _lset_type(ll, nt.Temp, int64(12))
			return LT{L: ll, Temp: nt.Temp}
		}
		nt := _lnew_temp(ll)
		_ = nt
		ll = nt.L
		ll = _lemit(ll, new_inst(IrOpOpArrayGet{}, nt.Temp, container.Temp, subscript.Temp, "", int64(0)))
		if (elem_type > int64(0)) {
			ll = _lset_type(ll, nt.Temp, elem_type)
		}
		return LT{L: ll, Temp: nt.Temp}
	}
	if (nk == "Array") {
		arr_nt := _lnew_temp(ll)
		_ = arr_nt
		ll = arr_nt.L
		ll = _lemit(ll, new_inst(IrOpOpArrayNew{}, arr_nt.Temp, node.List_count, int64(0), "", int64(0)))
		ai := int64(0)
		_ = ai
		arr_elem_type := int64(0)
		_ = arr_elem_type
		for (ai < node.List_count) {
			idx_node := pool_get(ll.Pool, (node.List_start + ai))
			_ = idx_node
			elem := _lower_expr_node(ll, idx_node.C0)
			_ = elem
			ll = elem.L
			elem_tid := _lget_type(ll, elem.Temp)
			_ = elem_tid
			if (elem_tid > int64(0)) {
				arr_elem_type = elem_tid
			}
			ll = _lemit(ll, new_inst(IrOpOpArrayAppend{}, arr_nt.Temp, arr_nt.Temp, elem.Temp, "", elem_tid))
			ai = (ai + int64(1))
		}
		if (arr_elem_type == int64(12)) {
			ll = _lset_type(ll, arr_nt.Temp, int64(13))
		}
		return LT{L: ll, Temp: arr_nt.Temp}
	}
	if (nk == "MapLit") {
		cap_nt := _lnew_temp(ll)
		_ = cap_nt
		ll = cap_nt.L
		cap := node.List_count
		_ = cap
		if (cap < int64(8)) {
			cap = int64(8)
		}
		ll = _lemit(ll, new_inst(IrOpOpConst{}, cap_nt.Temp, cap, int64(0), "", int64(0)))
		map_nt := _lnew_temp(ll)
		_ = map_nt
		ll = map_nt.L
		ll = _lemit(ll, new_inst(IrOpOpMapNew{}, map_nt.Temp, cap_nt.Temp, int64(0), "", int64(0)))
		ll = _lset_type(ll, map_nt.Temp, int64(300))
		mi := int64(0)
		_ = mi
		for (mi < (node.List_count * int64(2))) {
			key_idx_node := pool_get(ll.Pool, (node.List_start + mi))
			_ = key_idx_node
			val_idx_node := pool_get(ll.Pool, ((node.List_start + mi) + int64(1)))
			_ = val_idx_node
			key := _lower_expr_node(ll, key_idx_node.C0)
			_ = key
			ll = key.L
			val := _lower_expr_node(ll, val_idx_node.C0)
			_ = val
			ll = val.L
			ll = _lemit(ll, new_inst(IrOpOpMapSet{}, int64(0), map_nt.Temp, key.Temp, "", val.Temp))
			mi = (mi + int64(2))
		}
		return LT{L: ll, Temp: map_nt.Temp}
	}
	if (nk == "SetLit") {
		cap_nt := _lnew_temp(ll)
		_ = cap_nt
		ll = cap_nt.L
		cap := node.List_count
		_ = cap
		if (cap < int64(8)) {
			cap = int64(8)
		}
		ll = _lemit(ll, new_inst(IrOpOpConst{}, cap_nt.Temp, cap, int64(0), "", int64(0)))
		set_nt := _lnew_temp(ll)
		_ = set_nt
		ll = set_nt.L
		ll = _lemit(ll, new_inst(IrOpOpSetNew{}, set_nt.Temp, cap_nt.Temp, int64(0), "", int64(0)))
		ll = _lset_type(ll, set_nt.Temp, int64(301))
		si := int64(0)
		_ = si
		for (si < node.List_count) {
			idx_node := pool_get(ll.Pool, (node.List_start + si))
			_ = idx_node
			elem := _lower_expr_node(ll, idx_node.C0)
			_ = elem
			ll = elem.L
			ll = _lemit(ll, new_inst(IrOpOpSetAdd{}, int64(0), set_nt.Temp, elem.Temp, "", int64(0)))
			si = (si + int64(1))
		}
		return LT{L: ll, Temp: set_nt.Temp}
	}
	if (nk == "Struct") {
		sname := node.S1
		_ = sname
		sdef := reg_find_struct(ll.Registry, sname)
		_ = sdef
		alloc_sz := _struct_field_count(ll.Registry, sname)
		_ = alloc_sz
		nt := _lnew_temp(ll)
		_ = nt
		ll = nt.L
		ll = _lemit(ll, new_inst(IrOpOpAlloc{}, nt.Temp, alloc_sz, int64(0), sname, int64(0)))
		st_id := _struct_type_id(ll.Registry, sname)
		_ = st_id
		if (st_id > int64(0)) {
			ll = _lset_type(ll, nt.Temp, st_id)
		}
		fi := int64(0)
		_ = fi
		for (fi < node.List_count) {
			idx_node := pool_get(ll.Pool, (node.List_start + fi))
			_ = idx_node
			fval := _lower_expr_node(ll, idx_node.C0)
			_ = fval
			ll = fval.L
			fname := ""
			_ = fname
			f_slot := fi
			_ = f_slot
			fi_idx := (fi + int64(1))
			_ = fi_idx
			ft := int64(0)
			_ = ft
			if ((sdef.Name != "") && (fi_idx < int64(len(sdef.Fields)))) {
				fname = sdef.Fields[fi_idx].Name
				ft = sdef.Fields[fi_idx].Type_id
				f_slot = _field_index(ll.Registry, sname, fname)
			} else {
				ft = _lget_type(ll, fval.Temp)
			}
			if (ft == int64(12)) {
				ll = _lemit(ll, new_inst(IrOpOpFieldSet{}, nt.Temp, fval.Temp, f_slot, fname, int64(0)))
				f_slot_len := (f_slot + int64(1))
				_ = f_slot_len
				ll = _lemit(ll, new_inst(IrOpOpFieldSet{}, nt.Temp, (fval.Temp + int64(1)), f_slot_len, fname, int64(0)))
			} else {
				ll = _lemit(ll, new_inst(IrOpOpFieldSet{}, nt.Temp, fval.Temp, f_slot, fname, int64(0)))
			}
			fi = (fi + int64(1))
		}
		return LT{L: ll, Temp: nt.Temp}
	}
	if (nk == "Tuple") {
		nt := _lnew_temp(ll)
		_ = nt
		ll = nt.L
		ll = _lemit(ll, new_inst(IrOpOpAlloc{}, nt.Temp, node.List_count, int64(0), "_Tuple", int64(0)))
		ti := int64(0)
		_ = ti
		for (ti < node.List_count) {
			idx_node := pool_get(ll.Pool, (node.List_start + ti))
			_ = idx_node
			elem := _lower_expr_node(ll, idx_node.C0)
			_ = elem
			ll = elem.L
			ll = _lemit(ll, new_inst(IrOpOpFieldSet{}, nt.Temp, elem.Temp, ti, "", int64(0)))
			ti = (ti + int64(1))
		}
		return LT{L: ll, Temp: nt.Temp}
	}
	if (nk == "Pipeline") {
		left := _lower_expr_node(ll, node.C0)
		_ = left
		ll = left.L
		right_node := pool_get(ll.Pool, node.C1)
		_ = right_node
		left_type := _lget_type(ll, left.Temp)
		_ = left_type
		if (expr_kind_name(right_node.Kind) == "Ident") {
			fname := right_node.S1
			_ = fname
			fsig := reg_find_fn(ll.Registry, fname)
			_ = fsig
			if (fsig.Name != "") {
				reg_count := int64(1)
				_ = reg_count
				if (left_type == int64(12)) {
					reg_count = int64(2)
				}
				nt_temp := int64(0)
				_ = nt_temp
				if (fsig.Return_type == int64(12)) {
					nt := _lnew_str_temp(ll)
					_ = nt
					ll = nt.L
					nt_temp = nt.Temp
				} else {
					nt := _lnew_temp(ll)
					_ = nt
					ll = nt.L
					nt_temp = nt.Temp
				}
				ll = _lemit(ll, new_inst(IrOpOpCall{}, nt_temp, left.Temp, reg_count, fname, fsig.Return_type))
				if (fsig.Return_type > int64(0)) {
					ll = _lset_type(ll, nt_temp, fsig.Return_type)
				}
				return LT{L: ll, Temp: nt_temp}
			}
		}
		return LT{L: ll, Temp: left.Temp}
	}
	if (nk == "Range") {
		start := _lower_expr_node(ll, node.C0)
		_ = start
		ll = start.L
		return LT{L: ll, Temp: start.Temp}
	}
	if (nk == "ListComp") {
		arr_nt := _lnew_temp(ll)
		_ = arr_nt
		ll = arr_nt.L
		ll = _lemit(ll, new_inst(IrOpOpArrayNew{}, arr_nt.Temp, int64(0), int64(0), "", int64(0)))
		iter := _lower_expr_node(ll, node.C1)
		_ = iter
		ll = iter.L
		iter_type := _lget_type(ll, iter.Temp)
		_ = iter_type
		len_nt := _lnew_temp(ll)
		_ = len_nt
		ll = len_nt.L
		ll = _lemit(ll, new_inst(IrOpOpArrayLen{}, len_nt.Temp, iter.Temp, int64(0), "", int64(0)))
		idx_nt := _lnew_temp(ll)
		_ = idx_nt
		ll = idx_nt.L
		ll = _lemit(ll, new_inst(IrOpOpConst{}, idx_nt.Temp, int64(0), int64(0), "", int64(0)))
		loop_lbl := _lnew_label(ll)
		_ = loop_lbl
		ll = loop_lbl.L
		end_lbl := _lnew_label(ll)
		_ = end_lbl
		ll = end_lbl.L
		ll = _lemit(ll, new_inst(IrOpOpLabel{}, int64(0), loop_lbl.Temp, int64(0), "", int64(0)))
		cond_nt := _lnew_temp(ll)
		_ = cond_nt
		ll = cond_nt.L
		ll = _lemit(ll, new_inst(IrOpOpLt{}, cond_nt.Temp, idx_nt.Temp, len_nt.Temp, "", int64(0)))
		ll = _lemit(ll, new_inst(IrOpOpBranchFalse{}, int64(0), cond_nt.Temp, end_lbl.Temp, "", int64(0)))
		if ((iter_type == int64(13)) || _is_str_array_type(ll.Store, iter_type)) {
			elem_nt := _lnew_str_temp(ll)
			_ = elem_nt
			ll = elem_nt.L
			ll = _lemit(ll, new_inst(IrOpOpArrayGet{}, elem_nt.Temp, iter.Temp, idx_nt.Temp, "", int64(12)))
			ll = _lset_type(ll, elem_nt.Temp, int64(12))
			ll = _lenv_add(ll, node.S1, elem_nt.Temp)
		} else {
			elem_nt := _lnew_temp(ll)
			_ = elem_nt
			ll = elem_nt.L
			ll = _lemit(ll, new_inst(IrOpOpArrayGet{}, elem_nt.Temp, iter.Temp, idx_nt.Temp, "", int64(0)))
			elem_type := _array_elem_type(ll.Store, iter_type)
			_ = elem_type
			if (elem_type > int64(0)) {
				ll = _lset_type(ll, elem_nt.Temp, elem_type)
			}
			ll = _lenv_add(ll, node.S1, elem_nt.Temp)
		}
		if node.B1 {
			skip_lbl := _lnew_label(ll)
			_ = skip_lbl
			ll = skip_lbl.L
			wcond := _lower_expr_node(ll, node.C2)
			_ = wcond
			ll = wcond.L
			ll = _lemit(ll, new_inst(IrOpOpBranchFalse{}, int64(0), wcond.Temp, skip_lbl.Temp, "", int64(0)))
			mval := _lower_expr_node(ll, node.C0)
			_ = mval
			ll = mval.L
			ll = _lemit(ll, new_inst(IrOpOpArrayAppend{}, arr_nt.Temp, arr_nt.Temp, mval.Temp, "", int64(0)))
			ll = _lemit(ll, new_inst(IrOpOpLabel{}, int64(0), skip_lbl.Temp, int64(0), "", int64(0)))
		} else {
			mval := _lower_expr_node(ll, node.C0)
			_ = mval
			ll = mval.L
			ll = _lemit(ll, new_inst(IrOpOpArrayAppend{}, arr_nt.Temp, arr_nt.Temp, mval.Temp, "", int64(0)))
		}
		one_nt := _lnew_temp(ll)
		_ = one_nt
		ll = one_nt.L
		ll = _lemit(ll, new_inst(IrOpOpConst{}, one_nt.Temp, int64(1), int64(0), "", int64(0)))
		inc_nt := _lnew_temp(ll)
		_ = inc_nt
		ll = inc_nt.L
		ll = _lemit(ll, new_inst(IrOpOpAdd{}, inc_nt.Temp, idx_nt.Temp, one_nt.Temp, "", int64(0)))
		ll = _lemit(ll, new_inst(IrOpOpStore{}, idx_nt.Temp, inc_nt.Temp, int64(0), "", int64(0)))
		ll = _lemit(ll, new_inst(IrOpOpJump{}, int64(0), loop_lbl.Temp, int64(0), "", int64(0)))
		ll = _lemit(ll, new_inst(IrOpOpLabel{}, int64(0), end_lbl.Temp, int64(0), "", int64(0)))
		return LT{L: ll, Temp: arr_nt.Temp}
	}
	if (nk == "Yield") {
		if node.B1 {
			tp := _find_tok_pos(ll, node.Span.Offset)
			_ = tp
			ll = _lset_pos(ll, tp)
			ll = _ladv(ll)
			result := _lower_expr(ll)
			_ = result
			ll = result.L
			return LT{L: ll, Temp: result.Temp}
		}
		return LT{L: ll, Temp: -int64(1)}
	}
	if (nk == "Propagate") {
		operand := _lower_expr_node(ll, node.C0)
		_ = operand
		ll = operand.L
		tag_nt := _lnew_temp(ll)
		_ = tag_nt
		ll = tag_nt.L
		ll = _lemit(ll, new_inst(IrOpOpFieldGet{}, tag_nt.Temp, operand.Temp, int64(0), "_tag", int64(0)))
		const_one := _lnew_temp(ll)
		_ = const_one
		ll = const_one.L
		ll = _lemit(ll, new_inst(IrOpOpConst{}, const_one.Temp, int64(1), int64(0), "", int64(0)))
		is_err := _lnew_temp(ll)
		_ = is_err
		ll = is_err.L
		ll = _lemit(ll, new_inst(IrOpOpEq{}, is_err.Temp, tag_nt.Temp, const_one.Temp, "", int64(0)))
		err_lbl := _lnew_label(ll)
		_ = err_lbl
		ll = err_lbl.L
		ok_lbl := _lnew_label(ll)
		_ = ok_lbl
		ll = ok_lbl.L
		ll = _lemit(ll, new_inst(IrOpOpBranchTrue{}, int64(0), is_err.Temp, err_lbl.Temp, "", int64(0)))
		ok_val := _lnew_temp(ll)
		_ = ok_val
		ll = ok_val.L
		ll = _lemit(ll, new_inst(IrOpOpFieldGet{}, ok_val.Temp, operand.Temp, int64(1), "_val", int64(0)))
		ll = _lemit(ll, new_inst(IrOpOpJump{}, int64(0), ok_lbl.Temp, int64(0), "", int64(0)))
		ll = _lemit(ll, new_inst(IrOpOpLabel{}, int64(0), err_lbl.Temp, int64(0), "", int64(0)))
		ll = _emit_defers(ll)
		ll = _lemit(ll, new_inst(IrOpOpRet{}, int64(0), operand.Temp, int64(0), "", int64(0)))
		ll = _lemit(ll, new_inst(IrOpOpLabel{}, int64(0), ok_lbl.Temp, int64(0), "", int64(0)))
		return LT{L: ll, Temp: ok_val.Temp}
	}
	if (nk == "AssertOk") {
		operand := _lower_expr_node(ll, node.C0)
		_ = operand
		ll = operand.L
		tag_nt := _lnew_temp(ll)
		_ = tag_nt
		ll = tag_nt.L
		ll = _lemit(ll, new_inst(IrOpOpFieldGet{}, tag_nt.Temp, operand.Temp, int64(0), "_tag", int64(0)))
		const_one := _lnew_temp(ll)
		_ = const_one
		ll = const_one.L
		ll = _lemit(ll, new_inst(IrOpOpConst{}, const_one.Temp, int64(1), int64(0), "", int64(0)))
		is_err := _lnew_temp(ll)
		_ = is_err
		ll = is_err.L
		ll = _lemit(ll, new_inst(IrOpOpEq{}, is_err.Temp, tag_nt.Temp, const_one.Temp, "", int64(0)))
		panic_lbl := _lnew_label(ll)
		_ = panic_lbl
		ll = panic_lbl.L
		ok_lbl := _lnew_label(ll)
		_ = ok_lbl
		ll = ok_lbl.L
		ll = _lemit(ll, new_inst(IrOpOpBranchTrue{}, int64(0), is_err.Temp, panic_lbl.Temp, "", int64(0)))
		ok_val := _lnew_temp(ll)
		_ = ok_val
		ll = ok_val.L
		ll = _lemit(ll, new_inst(IrOpOpFieldGet{}, ok_val.Temp, operand.Temp, int64(1), "_val", int64(0)))
		ll = _lemit(ll, new_inst(IrOpOpJump{}, int64(0), ok_lbl.Temp, int64(0), "", int64(0)))
		ll = _lemit(ll, new_inst(IrOpOpLabel{}, int64(0), panic_lbl.Temp, int64(0), "", int64(0)))
		panic_str_idx := mod_find_string(ll.Module, "unwrap failed on Err value")
		_ = panic_str_idx
		if (panic_str_idx == int64(0)) {
			new_mod := mod_add_string(ll.Module, "unwrap failed on Err value")
			_ = new_mod
			panic_str_idx = (int64(len(new_mod.String_constants)) - int64(1))
			ll = _lset_mod(ll, new_mod)
		}
		panic_ptr := _lnew_str_temp(ll)
		_ = panic_ptr
		ll = panic_ptr.L
		ll = _lemit(ll, new_inst(IrOpOpConstStr{}, panic_ptr.Temp, panic_str_idx, int64(25), "", int64(0)))
		ll = _lemit(ll, new_inst(IrOpOpCall{}, int64(0), panic_ptr.Temp, int64(2), "_aria_panic", int64(0)))
		ll = _lemit(ll, new_inst(IrOpOpLabel{}, int64(0), ok_lbl.Temp, int64(0), "", int64(0)))
		return LT{L: ll, Temp: ok_val.Temp}
	}
	if (nk == "Catch") {
		expr := _lower_expr_node(ll, node.C0)
		_ = expr
		ll = expr.L
		tag_nt := _lnew_temp(ll)
		_ = tag_nt
		ll = tag_nt.L
		ll = _lemit(ll, new_inst(IrOpOpFieldGet{}, tag_nt.Temp, expr.Temp, int64(0), "_tag", int64(0)))
		const_one := _lnew_temp(ll)
		_ = const_one
		ll = const_one.L
		ll = _lemit(ll, new_inst(IrOpOpConst{}, const_one.Temp, int64(1), int64(0), "", int64(0)))
		is_err := _lnew_temp(ll)
		_ = is_err
		ll = is_err.L
		ll = _lemit(ll, new_inst(IrOpOpEq{}, is_err.Temp, tag_nt.Temp, const_one.Temp, "", int64(0)))
		result_nt := _lnew_str_temp(ll)
		_ = result_nt
		ll = result_nt.L
		ok_lbl := _lnew_label(ll)
		_ = ok_lbl
		ll = ok_lbl.L
		end_lbl := _lnew_label(ll)
		_ = end_lbl
		ll = end_lbl.L
		ll = _lemit(ll, new_inst(IrOpOpBranchFalse{}, int64(0), is_err.Temp, ok_lbl.Temp, "", int64(0)))
		if (int64(len(node.S1)) > int64(0)) {
			err_val := _lnew_temp(ll)
			_ = err_val
			ll = err_val.L
			ll = _lemit(ll, new_inst(IrOpOpFieldGet{}, err_val.Temp, expr.Temp, int64(1), "_val", int64(0)))
			ll = _lenv_add(ll, node.S1, err_val.Temp)
		}
		handler := _lower_expr_node(ll, node.C1)
		_ = handler
		ll = handler.L
		if (handler.Temp >= int64(0)) {
			ll = _lemit(ll, new_inst(IrOpOpStore{}, result_nt.Temp, handler.Temp, int64(0), "", int64(0)))
		}
		ll = _lemit(ll, new_inst(IrOpOpJump{}, int64(0), end_lbl.Temp, int64(0), "", int64(0)))
		ll = _lemit(ll, new_inst(IrOpOpLabel{}, int64(0), ok_lbl.Temp, int64(0), "", int64(0)))
		ok_val := _lnew_temp(ll)
		_ = ok_val
		ll = ok_val.L
		ll = _lemit(ll, new_inst(IrOpOpFieldGet{}, ok_val.Temp, expr.Temp, int64(1), "_val", int64(0)))
		ll = _lemit(ll, new_inst(IrOpOpStore{}, result_nt.Temp, ok_val.Temp, int64(0), "", int64(0)))
		ll = _lemit(ll, new_inst(IrOpOpLabel{}, int64(0), end_lbl.Temp, int64(0), "", int64(0)))
		return LT{L: ll, Temp: result_nt.Temp}
	}
	if ((nk == "Defer") || (nk == "With")) {
		tp := _find_tok_pos(ll, node.Span.Offset)
		_ = tp
		ll = _lset_pos(ll, tp)
		ll = _lower_stmt(ll)
		return LT{L: ll, Temp: -int64(1)}
	}
	if (nk == "Closure") {
		tp := _find_tok_pos(ll, node.Span.Offset)
		_ = tp
		ll = _lset_pos(ll, tp)
		return _lower_closure(ll)
	}
	if (nk == "Match") {
		tp := _find_tok_pos(ll, node.Span.Offset)
		_ = tp
		ll = _lset_pos(ll, tp)
		return _lower_match(ll)
	}
	nt := _lnew_temp(ll)
	_ = nt
	return LT{L: nt.L, Temp: nt.Temp}
}

func _lower_for_range_node(l Lowerer, var_name string, range_node Expr, body_idx int64) LT {
	ll := l
	_ = ll
	start := _lower_expr_node(ll, range_node.C0)
	_ = start
	ll = start.L
	end_val := _lower_expr_node(ll, range_node.C1)
	_ = end_val
	ll = end_val.L
	limit_temp := end_val.Temp
	_ = limit_temp
	if range_node.B1 {
		one_nt := _lnew_temp(ll)
		_ = one_nt
		ll = one_nt.L
		ll = _lemit(ll, new_inst(IrOpOpConst{}, one_nt.Temp, int64(1), int64(0), "", int64(0)))
		inc_nt := _lnew_temp(ll)
		_ = inc_nt
		ll = inc_nt.L
		ll = _lemit(ll, new_inst(IrOpOpAdd{}, inc_nt.Temp, end_val.Temp, one_nt.Temp, "", int64(0)))
		limit_temp = inc_nt.Temp
	}
	idx_nt := _lnew_temp(ll)
	_ = idx_nt
	ll = idx_nt.L
	ll = _lemit(ll, new_inst(IrOpOpStore{}, idx_nt.Temp, start.Temp, int64(0), "", int64(0)))
	loop_lbl := _lnew_label(ll)
	_ = loop_lbl
	ll = loop_lbl.L
	end_lbl := _lnew_label(ll)
	_ = end_lbl
	ll = end_lbl.L
	inc_lbl := _lnew_label(ll)
	_ = inc_lbl
	ll = inc_lbl.L
	ll = _lemit(ll, new_inst(IrOpOpLabel{}, int64(0), loop_lbl.Temp, int64(0), "", int64(0)))
	cond_nt := _lnew_temp(ll)
	_ = cond_nt
	ll = cond_nt.L
	ll = _lemit(ll, new_inst(IrOpOpLt{}, cond_nt.Temp, idx_nt.Temp, limit_temp, "", int64(0)))
	ll = _lemit(ll, new_inst(IrOpOpBranchFalse{}, int64(0), cond_nt.Temp, end_lbl.Temp, "", int64(0)))
	if (var_name != "") {
		ll = _lenv_add(ll, var_name, idx_nt.Temp)
	}
	prev_start := ll.Loop_start
	_ = prev_start
	prev_end := ll.Loop_end
	_ = prev_end
	ll = _lset_loop(ll, inc_lbl.Temp, end_lbl.Temp)
	body := _lower_expr_node(ll, body_idx)
	_ = body
	ll = body.L
	ll = _lset_loop(ll, prev_start, prev_end)
	ll = _lemit(ll, new_inst(IrOpOpLabel{}, int64(0), inc_lbl.Temp, int64(0), "", int64(0)))
	one_nt := _lnew_temp(ll)
	_ = one_nt
	ll = one_nt.L
	ll = _lemit(ll, new_inst(IrOpOpConst{}, one_nt.Temp, int64(1), int64(0), "", int64(0)))
	inc_nt := _lnew_temp(ll)
	_ = inc_nt
	ll = inc_nt.L
	ll = _lemit(ll, new_inst(IrOpOpAdd{}, inc_nt.Temp, idx_nt.Temp, one_nt.Temp, "", int64(0)))
	ll = _lemit(ll, new_inst(IrOpOpStore{}, idx_nt.Temp, inc_nt.Temp, int64(0), "", int64(0)))
	ll = _lemit(ll, new_inst(IrOpOpJump{}, int64(0), loop_lbl.Temp, int64(0), "", int64(0)))
	ll = _lemit(ll, new_inst(IrOpOpLabel{}, int64(0), end_lbl.Temp, int64(0), "", int64(0)))
	return LT{L: ll, Temp: -int64(1)}
}

func _lower_assign_node(l Lowerer, node Expr) LT {
	ll := l
	_ = ll
	lhs := pool_get(ll.Pool, node.C0)
	_ = lhs
	lhs_nk := expr_kind_name(lhs.Kind)
	_ = lhs_nk
	result := _lower_expr_node(ll, node.C1)
	_ = result
	ll = result.L
	if (lhs_nk == "Ident") {
		slot := _lenv_lookup(ll, lhs.S1)
		_ = slot
		if (slot >= int64(0)) {
			ll = _lemit(ll, new_inst(IrOpOpStore{}, slot, result.Temp, int64(0), "", int64(0)))
			if (_lget_type(ll, result.Temp) == int64(12)) {
				ll = _lemit(ll, new_inst(IrOpOpStore{}, (slot + int64(1)), (result.Temp + int64(1)), int64(0), "", int64(0)))
			}
			rhs_type := _lget_type(ll, result.Temp)
			_ = rhs_type
			if (rhs_type > int64(0)) {
				ll = _lset_type(ll, slot, rhs_type)
			}
		}
	} else if (lhs_nk == "FieldAccess") {
		recv := _lower_expr_node(ll, lhs.C0)
		_ = recv
		ll = recv.L
		left_type := _lget_type(ll, recv.Temp)
		_ = left_type
		sname := _struct_name_for_type(ll.Store, ll.Registry, left_type)
		_ = sname
		fidx := int64(0)
		_ = fidx
		if (sname != "") {
			fidx = _field_index(ll.Registry, sname, lhs.S1)
		} else {
			fidx = _field_index_any(ll.Registry, lhs.S1)
		}
		ll = _lemit(ll, new_inst(IrOpOpFieldSet{}, recv.Temp, result.Temp, fidx, lhs.S1, int64(0)))
		if (_lget_type(ll, result.Temp) == int64(12)) {
			fidx_len := (fidx + int64(1))
			_ = fidx_len
			ll = _lemit(ll, new_inst(IrOpOpFieldSet{}, recv.Temp, (result.Temp + int64(1)), fidx_len, lhs.S1, int64(0)))
		}
	} else if (lhs_nk == "Index") {
		container := _lower_expr_node(ll, lhs.C0)
		_ = container
		ll = container.L
		subscript := _lower_expr_node(ll, lhs.C1)
		_ = subscript
		ll = subscript.L
		ll = _lemit(ll, new_inst(IrOpOpArraySet{}, container.Temp, subscript.Temp, result.Temp, "", int64(0)))
	}
	return LT{L: ll, Temp: result.Temp}
}

func _lower_call_node(l Lowerer, node Expr) LT {
	ll := l
	_ = ll
	name := node.S1
	_ = name
	if ((name == "Ok") || (name == "Err")) {
		if (node.List_count > int64(0)) {
			val := _lower_expr_node(ll, pool_get(ll.Pool, node.List_start).C0)
			_ = val
			ll = val.L
			val_type := _lget_type(ll, val.Temp)
			_ = val_type
			result_slots := int64(2)
			_ = result_slots
			if (val_type == int64(12)) {
				result_slots = int64(3)
			}
			nt := _lnew_temp(ll)
			_ = nt
			ll = nt.L
			ll = _lemit(ll, new_inst(IrOpOpAlloc{}, nt.Temp, result_slots, int64(0), "_Result", int64(0)))
			tag_nt := _lnew_temp(ll)
			_ = tag_nt
			ll = tag_nt.L
			tag_val := int64(0)
			_ = tag_val
			if (name == "Err") {
				tag_val = int64(1)
			}
			ll = _lemit(ll, new_inst(IrOpOpConst{}, tag_nt.Temp, tag_val, int64(0), "", int64(0)))
			ll = _lemit(ll, new_inst(IrOpOpFieldSet{}, nt.Temp, tag_nt.Temp, int64(0), "_tag", int64(0)))
			if (val_type == int64(12)) {
				ll = _lemit(ll, new_inst(IrOpOpFieldSet{}, nt.Temp, val.Temp, int64(1), "_val", int64(0)))
				val_len := (val.Temp + int64(1))
				_ = val_len
				ll = _lemit(ll, new_inst(IrOpOpFieldSet{}, nt.Temp, val_len, int64(2), "_val_len", int64(0)))
			} else {
				ll = _lemit(ll, new_inst(IrOpOpFieldSet{}, nt.Temp, val.Temp, int64(1), "_val", int64(0)))
			}
			ll = _lset_type(ll, nt.Temp, int64(200))
			return LT{L: ll, Temp: nt.Temp}
		}
	}
	if (name == "Some") {
		if (node.List_count > int64(0)) {
			val := _lower_expr_node(ll, pool_get(ll.Pool, node.List_start).C0)
			_ = val
			ll = val.L
			val_type := _lget_type(ll, val.Temp)
			_ = val_type
			opt_slots := int64(2)
			_ = opt_slots
			if (val_type == int64(12)) {
				opt_slots = int64(3)
			}
			nt := _lnew_temp(ll)
			_ = nt
			ll = nt.L
			ll = _lemit(ll, new_inst(IrOpOpAlloc{}, nt.Temp, opt_slots, int64(0), "_Optional", int64(0)))
			tag_nt := _lnew_temp(ll)
			_ = tag_nt
			ll = tag_nt.L
			ll = _lemit(ll, new_inst(IrOpOpConst{}, tag_nt.Temp, int64(1), int64(0), "", int64(0)))
			ll = _lemit(ll, new_inst(IrOpOpFieldSet{}, nt.Temp, tag_nt.Temp, int64(0), "_tag", int64(0)))
			if (val_type == int64(12)) {
				ll = _lemit(ll, new_inst(IrOpOpFieldSet{}, nt.Temp, val.Temp, int64(1), "_val", int64(0)))
				val_len := (val.Temp + int64(1))
				_ = val_len
				ll = _lemit(ll, new_inst(IrOpOpFieldSet{}, nt.Temp, val_len, int64(2), "_val_len", int64(0)))
			} else {
				ll = _lemit(ll, new_inst(IrOpOpFieldSet{}, nt.Temp, val.Temp, int64(1), "_val", int64(0)))
			}
			ll = _lset_type(ll, nt.Temp, int64(201))
			return LT{L: ll, Temp: nt.Temp}
		}
	}
	if (name == "Map") {
		cap_nt := _lnew_temp(ll)
		_ = cap_nt
		ll = cap_nt.L
		ll = _lemit(ll, new_inst(IrOpOpConst{}, cap_nt.Temp, int64(16), int64(0), "", int64(0)))
		nt := _lnew_temp(ll)
		_ = nt
		ll = nt.L
		ll = _lemit(ll, new_inst(IrOpOpMapNew{}, nt.Temp, cap_nt.Temp, int64(0), "", int64(0)))
		ll = _lset_type(ll, nt.Temp, int64(300))
		return LT{L: ll, Temp: nt.Temp}
	}
	if (name == "Set") {
		cap_nt := _lnew_temp(ll)
		_ = cap_nt
		ll = cap_nt.L
		ll = _lemit(ll, new_inst(IrOpOpConst{}, cap_nt.Temp, int64(16), int64(0), "", int64(0)))
		nt := _lnew_temp(ll)
		_ = nt
		ll = nt.L
		ll = _lemit(ll, new_inst(IrOpOpSetNew{}, nt.Temp, cap_nt.Temp, int64(0), "", int64(0)))
		ll = _lset_type(ll, nt.Temp, int64(301))
		return LT{L: ll, Temp: nt.Temp}
	}
	vtag := _variant_tag(ll.Registry, name)
	_ = vtag
	if (vtag > int64(0)) {
		vparent := _variant_parent(ll.Registry, name)
		_ = vparent
		vdc := variant_data_count(vparent, vtag)
		_ = vdc
		if ((vdc > int64(0)) && (node.List_count > int64(0))) {
			alloc_sz := sum_alloc_size(vparent)
			_ = alloc_sz
			nt := _lnew_temp(ll)
			_ = nt
			ll = nt.L
			ll = _lemit(ll, new_inst(IrOpOpAlloc{}, nt.Temp, alloc_sz, int64(0), vparent.Name, int64(0)))
			st_id := _struct_type_id(ll.Registry, vparent.Name)
			_ = st_id
			if (st_id > int64(0)) {
				ll = _lset_type(ll, nt.Temp, st_id)
			}
			vtag_nt := _lnew_temp(ll)
			_ = vtag_nt
			ll = vtag_nt.L
			ll = _lemit(ll, new_inst(IrOpOpConst{}, vtag_nt.Temp, vtag, int64(0), "", int64(0)))
			ll = _lemit(ll, new_inst(IrOpOpFieldSet{}, nt.Temp, vtag_nt.Temp, int64(0), "_tag", int64(0)))
			vfi := int64(0)
			_ = vfi
			for (vfi < node.List_count) {
				val := _lower_expr_node(ll, pool_get(ll.Pool, (node.List_start + vfi)).C0)
				_ = val
				ll = val.L
				vdata_slot := int64(1)
				_ = vdata_slot
				vsi := int64(0)
				_ = vsi
				for (vsi < vfi) {
					vft := variant_field_type(vparent, vtag, vsi)
					_ = vft
					if (vft == int64(12)) {
						vdata_slot = (vdata_slot + int64(2))
					} else {
						vdata_slot = (vdata_slot + int64(1))
					}
					vsi = (vsi + int64(1))
				}
				vft := variant_field_type(vparent, vtag, vfi)
				_ = vft
				if (vft == int64(12)) {
					ll = _lemit(ll, new_inst(IrOpOpFieldSet{}, nt.Temp, val.Temp, vdata_slot, "", int64(0)))
					ll = _lemit(ll, new_inst(IrOpOpFieldSet{}, nt.Temp, (val.Temp + int64(1)), (vdata_slot + int64(1)), "", int64(0)))
				} else {
					ll = _lemit(ll, new_inst(IrOpOpFieldSet{}, nt.Temp, val.Temp, vdata_slot, "", int64(0)))
				}
				vfi = (vfi + int64(1))
			}
			return LT{L: ll, Temp: nt.Temp}
		}
	}
	nt_def := reg_find_struct(ll.Registry, name)
	_ = nt_def
	if ((((nt_def.Name != "") && (nt_def.Is_sum == false)) && _is_newtype(nt_def)) && (node.List_count > int64(0))) {
		val := _lower_expr_node(ll, pool_get(ll.Pool, node.List_start).C0)
		_ = val
		ll = val.L
		nt := _lnew_temp(ll)
		_ = nt
		ll = nt.L
		ll = _lemit(ll, new_inst(IrOpOpAlloc{}, nt.Temp, int64(1), int64(0), name, int64(0)))
		st_id := _struct_type_id(ll.Registry, name)
		_ = st_id
		if (st_id > int64(0)) {
			ll = _lset_type(ll, nt.Temp, st_id)
		}
		ll = _lemit(ll, new_inst(IrOpOpFieldSet{}, nt.Temp, val.Temp, int64(0), "value", int64(0)))
		return LT{L: ll, Temp: nt.Temp}
	}
	if ((name == "assertEqual") && (node.List_count == int64(2))) {
		arg_a := _lower_expr_node(ll, pool_get(ll.Pool, node.List_start).C0)
		_ = arg_a
		ll = arg_a.L
		arg_b := _lower_expr_node(ll, pool_get(ll.Pool, (node.List_start + int64(1))).C0)
		_ = arg_b
		ll = arg_b.L
		at := _lget_type(ll, arg_a.Temp)
		_ = at
		ok_lbl := _lnew_label(ll)
		_ = ok_lbl
		ll = ok_lbl.L
		if (at == int64(12)) {
			eq_nt := _lnew_temp(ll)
			_ = eq_nt
			ll = eq_nt.L
			ll = _lemit(ll, new_inst(IrOpOpStrEq{}, eq_nt.Temp, arg_a.Temp, arg_b.Temp, "", int64(0)))
			ll = _lemit(ll, new_inst(IrOpOpBranchTrue{}, int64(0), eq_nt.Temp, ok_lbl.Temp, "", int64(0)))
		} else {
			eq_nt := _lnew_temp(ll)
			_ = eq_nt
			ll = eq_nt.L
			ll = _lemit(ll, new_inst(IrOpOpEq{}, eq_nt.Temp, arg_a.Temp, arg_b.Temp, "", int64(0)))
			ll = _lemit(ll, new_inst(IrOpOpBranchTrue{}, int64(0), eq_nt.Temp, ok_lbl.Temp, "", int64(0)))
		}
		msg := _lower_string_const(ll, "assertEqual failed")
		_ = msg
		ll = msg.L
		pnt := _lnew_temp(ll)
		_ = pnt
		ll = pnt.L
		ll = _lemit(ll, new_inst(IrOpOpCall{}, pnt.Temp, msg.Temp, int64(2), "_aria_println_str", int64(0)))
		exit_nt := _lnew_temp(ll)
		_ = exit_nt
		ll = exit_nt.L
		ll = _lemit(ll, new_inst(IrOpOpConst{}, exit_nt.Temp, int64(1), int64(0), "", int64(0)))
		ll = _lemit(ll, new_inst(IrOpOpCall{}, int64(0), exit_nt.Temp, int64(1), "_aria_exit", int64(0)))
		ll = _lemit(ll, new_inst(IrOpOpLabel{}, int64(0), ok_lbl.Temp, int64(0), "", int64(0)))
		return LT{L: ll, Temp: int64(0)}
	}
	if ((name == "assertOk") && (node.List_count == int64(1))) {
		arg := _lower_expr_node(ll, pool_get(ll.Pool, node.List_start).C0)
		_ = arg
		ll = arg.L
		tag_nt := _lnew_temp(ll)
		_ = tag_nt
		ll = tag_nt.L
		ll = _lemit(ll, new_inst(IrOpOpFieldGet{}, tag_nt.Temp, arg.Temp, int64(0), "_tag", int64(0)))
		ok_lbl := _lnew_label(ll)
		_ = ok_lbl
		ll = ok_lbl.L
		zero_nt := _lnew_temp(ll)
		_ = zero_nt
		ll = zero_nt.L
		ll = _lemit(ll, new_inst(IrOpOpConst{}, zero_nt.Temp, int64(0), int64(0), "", int64(0)))
		cmp_nt := _lnew_temp(ll)
		_ = cmp_nt
		ll = cmp_nt.L
		ll = _lemit(ll, new_inst(IrOpOpEq{}, cmp_nt.Temp, tag_nt.Temp, zero_nt.Temp, "", int64(0)))
		ll = _lemit(ll, new_inst(IrOpOpBranchTrue{}, int64(0), cmp_nt.Temp, ok_lbl.Temp, "", int64(0)))
		msg := _lower_string_const(ll, "assertOk failed: got Err")
		_ = msg
		ll = msg.L
		pnt := _lnew_temp(ll)
		_ = pnt
		ll = pnt.L
		ll = _lemit(ll, new_inst(IrOpOpCall{}, pnt.Temp, msg.Temp, int64(2), "_aria_println_str", int64(0)))
		exit_nt := _lnew_temp(ll)
		_ = exit_nt
		ll = exit_nt.L
		ll = _lemit(ll, new_inst(IrOpOpConst{}, exit_nt.Temp, int64(1), int64(0), "", int64(0)))
		ll = _lemit(ll, new_inst(IrOpOpCall{}, int64(0), exit_nt.Temp, int64(1), "_aria_exit", int64(0)))
		ll = _lemit(ll, new_inst(IrOpOpLabel{}, int64(0), ok_lbl.Temp, int64(0), "", int64(0)))
		return LT{L: ll, Temp: int64(0)}
	}
	if ((name == "assertErr") && (node.List_count == int64(1))) {
		arg := _lower_expr_node(ll, pool_get(ll.Pool, node.List_start).C0)
		_ = arg
		ll = arg.L
		tag_nt := _lnew_temp(ll)
		_ = tag_nt
		ll = tag_nt.L
		ll = _lemit(ll, new_inst(IrOpOpFieldGet{}, tag_nt.Temp, arg.Temp, int64(0), "_tag", int64(0)))
		ok_lbl := _lnew_label(ll)
		_ = ok_lbl
		ll = ok_lbl.L
		zero_nt := _lnew_temp(ll)
		_ = zero_nt
		ll = zero_nt.L
		ll = _lemit(ll, new_inst(IrOpOpConst{}, zero_nt.Temp, int64(0), int64(0), "", int64(0)))
		cmp_nt := _lnew_temp(ll)
		_ = cmp_nt
		ll = cmp_nt.L
		ll = _lemit(ll, new_inst(IrOpOpNeq{}, cmp_nt.Temp, tag_nt.Temp, zero_nt.Temp, "", int64(0)))
		ll = _lemit(ll, new_inst(IrOpOpBranchTrue{}, int64(0), cmp_nt.Temp, ok_lbl.Temp, "", int64(0)))
		msg := _lower_string_const(ll, "assertErr failed: got Ok")
		_ = msg
		ll = msg.L
		pnt := _lnew_temp(ll)
		_ = pnt
		ll = pnt.L
		ll = _lemit(ll, new_inst(IrOpOpCall{}, pnt.Temp, msg.Temp, int64(2), "_aria_println_str", int64(0)))
		exit_nt := _lnew_temp(ll)
		_ = exit_nt
		ll = exit_nt.L
		ll = _lemit(ll, new_inst(IrOpOpConst{}, exit_nt.Temp, int64(1), int64(0), "", int64(0)))
		ll = _lemit(ll, new_inst(IrOpOpCall{}, int64(0), exit_nt.Temp, int64(1), "_aria_exit", int64(0)))
		ll = _lemit(ll, new_inst(IrOpOpLabel{}, int64(0), ok_lbl.Temp, int64(0), "", int64(0)))
		return LT{L: ll, Temp: int64(0)}
	}
	if ((name == "assertNear") && (node.List_count == int64(3))) {
		arg_a := _lower_expr_node(ll, pool_get(ll.Pool, node.List_start).C0)
		_ = arg_a
		ll = arg_a.L
		arg_b := _lower_expr_node(ll, pool_get(ll.Pool, (node.List_start + int64(1))).C0)
		_ = arg_b
		ll = arg_b.L
		arg_eps := _lower_expr_node(ll, pool_get(ll.Pool, (node.List_start + int64(2))).C0)
		_ = arg_eps
		ll = arg_eps.L
		diff := _lnew_temp(ll)
		_ = diff
		ll = diff.L
		ll = _lemit(ll, new_inst(IrOpOpSub{}, diff.Temp, arg_a.Temp, arg_b.Temp, "", int64(0)))
		neg_lbl := _lnew_label(ll)
		_ = neg_lbl
		ll = neg_lbl.L
		done_lbl := _lnew_label(ll)
		_ = done_lbl
		ll = done_lbl.L
		abs_nt := _lnew_temp(ll)
		_ = abs_nt
		ll = abs_nt.L
		zero_nt := _lnew_temp(ll)
		_ = zero_nt
		ll = zero_nt.L
		ll = _lemit(ll, new_inst(IrOpOpConst{}, zero_nt.Temp, int64(0), int64(0), "", int64(0)))
		lt_nt := _lnew_temp(ll)
		_ = lt_nt
		ll = lt_nt.L
		ll = _lemit(ll, new_inst(IrOpOpLt{}, lt_nt.Temp, diff.Temp, zero_nt.Temp, "", int64(0)))
		ll = _lemit(ll, new_inst(IrOpOpBranchTrue{}, int64(0), lt_nt.Temp, neg_lbl.Temp, "", int64(0)))
		ll = _lemit(ll, new_inst(IrOpOpStore{}, abs_nt.Temp, diff.Temp, int64(0), "", int64(0)))
		ll = _lemit(ll, new_inst(IrOpOpJump{}, int64(0), done_lbl.Temp, int64(0), "", int64(0)))
		ll = _lemit(ll, new_inst(IrOpOpLabel{}, int64(0), neg_lbl.Temp, int64(0), "", int64(0)))
		neg_diff := _lnew_temp(ll)
		_ = neg_diff
		ll = neg_diff.L
		ll = _lemit(ll, new_inst(IrOpOpNeg{}, neg_diff.Temp, diff.Temp, int64(0), "", int64(0)))
		ll = _lemit(ll, new_inst(IrOpOpStore{}, abs_nt.Temp, neg_diff.Temp, int64(0), "", int64(0)))
		ll = _lemit(ll, new_inst(IrOpOpLabel{}, int64(0), done_lbl.Temp, int64(0), "", int64(0)))
		ok_lbl := _lnew_label(ll)
		_ = ok_lbl
		ll = ok_lbl.L
		le_nt := _lnew_temp(ll)
		_ = le_nt
		ll = le_nt.L
		ll = _lemit(ll, new_inst(IrOpOpLte{}, le_nt.Temp, abs_nt.Temp, arg_eps.Temp, "", int64(0)))
		ll = _lemit(ll, new_inst(IrOpOpBranchTrue{}, int64(0), le_nt.Temp, ok_lbl.Temp, "", int64(0)))
		msg := _lower_string_const(ll, "assertNear failed: values not within epsilon")
		_ = msg
		ll = msg.L
		pnt := _lnew_temp(ll)
		_ = pnt
		ll = pnt.L
		ll = _lemit(ll, new_inst(IrOpOpCall{}, pnt.Temp, msg.Temp, int64(2), "_aria_println_str", int64(0)))
		exit_nt := _lnew_temp(ll)
		_ = exit_nt
		ll = exit_nt.L
		ll = _lemit(ll, new_inst(IrOpOpConst{}, exit_nt.Temp, int64(1), int64(0), "", int64(0)))
		ll = _lemit(ll, new_inst(IrOpOpCall{}, int64(0), exit_nt.Temp, int64(1), "_aria_exit", int64(0)))
		ll = _lemit(ll, new_inst(IrOpOpLabel{}, int64(0), ok_lbl.Temp, int64(0), "", int64(0)))
		return LT{L: ll, Temp: int64(0)}
	}
	if ((name == "dbg") && (node.List_count == int64(1))) {
		arg := _lower_expr_node(ll, pool_get(ll.Pool, node.List_start).C0)
		_ = arg
		ll = arg.L
		at := _lget_type(ll, arg.Temp)
		_ = at
		prefix := (("[" + ll.File) + "] ")
		_ = prefix
		prefix_res := _lower_string_const(ll, prefix)
		_ = prefix_res
		ll = prefix_res.L
		if (at == int64(12)) {
			cat := _lnew_str_temp(ll)
			_ = cat
			ll = cat.L
			ll = _lemit(ll, new_inst(IrOpOpStrConcat{}, cat.Temp, prefix_res.Temp, arg.Temp, "", int64(0)))
			pnt := _lnew_temp(ll)
			_ = pnt
			ll = pnt.L
			ll = _lemit(ll, new_inst(IrOpOpCall{}, pnt.Temp, cat.Temp, int64(2), "_aria_println_str", int64(0)))
		} else {
			vstr := _lnew_str_temp(ll)
			_ = vstr
			ll = vstr.L
			ll = _lemit(ll, new_inst(IrOpOpIntToStr{}, vstr.Temp, arg.Temp, int64(0), "", int64(0)))
			cat := _lnew_str_temp(ll)
			_ = cat
			ll = cat.L
			ll = _lemit(ll, new_inst(IrOpOpStrConcat{}, cat.Temp, prefix_res.Temp, vstr.Temp, "", int64(0)))
			pnt := _lnew_temp(ll)
			_ = pnt
			ll = pnt.L
			ll = _lemit(ll, new_inst(IrOpOpCall{}, pnt.Temp, cat.Temp, int64(2), "_aria_println_str", int64(0)))
		}
		return LT{L: ll, Temp: arg.Temp}
	}
	if (((name == "println") || (name == "print")) || (name == "eprintln")) {
		if (node.List_count > int64(0)) {
			arg := _lower_expr_node(ll, pool_get(ll.Pool, node.List_start).C0)
			_ = arg
			ll = arg.L
			print_temp := arg.Temp
			_ = print_temp
			arg_type := _lget_type(ll, arg.Temp)
			_ = arg_type
			if (arg_type == int64(1)) {
				conv_nt := _lnew_str_temp(ll)
				_ = conv_nt
				ll = conv_nt.L
				ll = _lemit(ll, new_inst(IrOpOpIntToStr{}, conv_nt.Temp, arg.Temp, int64(0), "", int64(0)))
				ll = _lset_type(ll, conv_nt.Temp, int64(12))
				print_temp = conv_nt.Temp
			}
			nt := _lnew_temp(ll)
			_ = nt
			ll = nt.L
			ll = _lemit(ll, new_inst(IrOpOpCall{}, nt.Temp, print_temp, int64(2), "_aria_println_str", int64(0)))
			return LT{L: ll, Temp: nt.Temp}
		}
	}
	if (name == "panic") {
		if (node.List_count > int64(0)) {
			arg := _lower_expr_node(ll, pool_get(ll.Pool, node.List_start).C0)
			_ = arg
			ll = arg.L
			nt := _lnew_temp(ll)
			_ = nt
			ll = nt.L
			ll = _lemit(ll, new_inst(IrOpOpCall{}, nt.Temp, arg.Temp, int64(1), "_aria_exit", int64(0)))
			return LT{L: ll, Temp: nt.Temp}
		}
	}
	if (name == "assert") {
		if (node.List_count > int64(0)) {
			cond := _lower_expr_node(ll, pool_get(ll.Pool, node.List_start).C0)
			_ = cond
			ll = cond.L
			ok_lbl := _lnew_label(ll)
			_ = ok_lbl
			ll = ok_lbl.L
			ll = _lemit(ll, new_inst(IrOpOpBranchTrue{}, int64(0), cond.Temp, ok_lbl.Temp, "", int64(0)))
			panic_str_idx := mod_find_string(ll.Module, "assertion failed")
			_ = panic_str_idx
			if (panic_str_idx == int64(0)) {
				new_mod := mod_add_string(ll.Module, "assertion failed")
				_ = new_mod
				panic_str_idx = (int64(len(new_mod.String_constants)) - int64(1))
				ll = _lset_mod(ll, new_mod)
			}
			panic_ptr := _lnew_str_temp(ll)
			_ = panic_ptr
			ll = panic_ptr.L
			ll = _lemit(ll, new_inst(IrOpOpConstStr{}, panic_ptr.Temp, panic_str_idx, int64(16), "", int64(0)))
			ll = _lemit(ll, new_inst(IrOpOpCall{}, int64(0), panic_ptr.Temp, int64(2), "_aria_panic", int64(0)))
			ll = _lemit(ll, new_inst(IrOpOpLabel{}, int64(0), ok_lbl.Temp, int64(0), "", int64(0)))
			nt := _lnew_temp(ll)
			_ = nt
			ll = nt.L
			return LT{L: ll, Temp: nt.Temp}
		}
	}
	arg_temps := []int64{int64(0)}
	_ = arg_temps
	arg_types := []int64{int64(0)}
	_ = arg_types
	ai := int64(0)
	_ = ai
	for (ai < node.List_count) {
		arg := _lower_expr_node(ll, pool_get(ll.Pool, (node.List_start + ai)).C0)
		_ = arg
		ll = arg.L
		arg_temps = append(arg_temps, arg.Temp)
		arg_types = append(arg_types, _lget_type(ll, arg.Temp))
		ai = (ai + int64(1))
	}
	sig := reg_find_fn(ll.Registry, name)
	_ = sig
	call_name := name
	_ = call_name
	call_sig := sig
	_ = call_sig
	if ((sig.Name != "") && fn_is_generic(sig)) {
		call_name = _resolve_generic_call_name(ll.Store, ll.Registry, name, sig, arg_types)
		sp_sig := reg_find_fn(ll.Registry, call_name)
		_ = sp_sig
		if (sp_sig.Name != "") {
			call_sig = sp_sig
		} else if (int64(len(arg_types)) > int64(1)) {
			first_concrete := arg_types[int64(1)]
			_ = first_concrete
			sp := reg_find_mono_by_generic(ll.Registry, name, first_concrete)
			_ = sp
			if (sp.Specialized_name == "") {
				msi := int64(1)
				_ = msi
				for (msi < int64(len(ll.Registry.Mono_specs))) {
					if (ll.Registry.Mono_specs[msi].Generic_name == name) {
						sp = ll.Registry.Mono_specs[msi]
						msi = int64(len(ll.Registry.Mono_specs))
					}
					msi = (msi + int64(1))
				}
			}
			if (sp.Specialized_name != "") {
				call_name = sp.Specialized_name
				sp_sig2 := reg_find_fn(ll.Registry, call_name)
				_ = sp_sig2
				if (sp_sig2.Name != "") {
					call_sig = sp_sig2
				}
			}
		}
	}
	if ((call_sig.Name != "") || strings.HasPrefix(name, "_aria")) {
		reg_count := int64(0)
		_ = reg_count
		ai = int64(1)
		for (ai < int64(len(arg_temps))) {
			if (arg_types[ai] == int64(12)) {
				reg_count = (reg_count + int64(2))
			} else {
				reg_count = (reg_count + int64(1))
			}
			ai = (ai + int64(1))
		}
		already_consecutive := true
		_ = already_consecutive
		if (int64(len(arg_temps)) > int64(1)) {
			expected := arg_temps[int64(1)]
			_ = expected
			ci := int64(1)
			_ = ci
			for (ci < int64(len(arg_temps))) {
				if (arg_temps[ci] != expected) {
					already_consecutive = false
				}
				if (arg_types[ci] == int64(12)) {
					expected = (expected + int64(2))
				} else {
					expected = (expected + int64(1))
				}
				ci = (ci + int64(1))
			}
		}
		fa := int64(0)
		_ = fa
		if (already_consecutive && (int64(len(arg_temps)) > int64(1))) {
			fa = arg_temps[int64(1)]
		} else if (int64(len(arg_temps)) > int64(1)) {
			first_nt := _lnew_temp(ll)
			_ = first_nt
			ll = first_nt.L
			ri := int64(1)
			_ = ri
			for (ri < reg_count) {
				extra := _lnew_temp(ll)
				_ = extra
				ll = extra.L
				ri = (ri + int64(1))
			}
			slot := int64(0)
			_ = slot
			ai = int64(1)
			for (ai < int64(len(arg_temps))) {
				if (arg_types[ai] == int64(12)) {
					ll = _lemit(ll, new_inst(IrOpOpStore{}, (first_nt.Temp + slot), arg_temps[ai], int64(0), "", int64(0)))
					slot_plus := (slot + int64(1))
					_ = slot_plus
					ll = _lemit(ll, new_inst(IrOpOpStore{}, (first_nt.Temp + slot_plus), (arg_temps[ai] + int64(1)), int64(0), "", int64(0)))
					slot = (slot + int64(2))
				} else {
					ll = _lemit(ll, new_inst(IrOpOpStore{}, (first_nt.Temp + slot), arg_temps[ai], int64(0), "", int64(0)))
					slot = (slot + int64(1))
				}
				ai = (ai + int64(1))
			}
			fa = first_nt.Temp
		}
		rt := int64(0)
		_ = rt
		runtime_name := call_name
		_ = runtime_name
		if (call_sig.Name != "") {
			rt = call_sig.Return_type
		} else {
			runtime_name = _builtin_runtime_name(name)
			rt = _builtin_return_type(name)
		}
		nt_temp := int64(0)
		_ = nt_temp
		if (rt == int64(12)) {
			nt := _lnew_str_temp(ll)
			_ = nt
			ll = nt.L
			nt_temp = nt.Temp
		} else {
			nt := _lnew_temp(ll)
			_ = nt
			ll = nt.L
			nt_temp = nt.Temp
		}
		ll = _lemit(ll, new_inst(IrOpOpCall{}, nt_temp, fa, reg_count, runtime_name, rt))
		if (rt == int64(12)) {
			ll = _lset_type(ll, nt_temp, int64(12))
		}
		if (rt > int64(0)) {
			if (rt != int64(12)) {
				ll = _lset_type(ll, nt_temp, rt)
			}
		}
		if (call_sig.Name != "") {
			if (call_sig.Return_type > int64(0)) {
				ll = _lset_type(ll, nt_temp, call_sig.Return_type)
			}
		}
		return LT{L: ll, Temp: nt_temp}
	}
	tp := _find_tok_pos(ll, _expr_start_offset(ll.Pool, node.C0))
	_ = tp
	ll = _lset_pos(ll, tp)
	return _lower_primary(ll)
}

func _builtin_runtime_name(name string) string {
	if (name == "_ariaReadFile") {
		return "_aria_read_file"
	}
	if (name == "_ariaWriteFile") {
		return "_aria_write_file"
	}
	if (name == "_ariaAppendFile") {
		return "_aria_append_file"
	}
	if (name == "_ariaWriteBinaryFile") {
		return "_aria_write_binary_file"
	}
	if (name == "_ariaExec") {
		return "_aria_exec"
	}
	if (name == "_ariaArgs") {
		return "_aria_args"
	}
	if (name == "_ariaListDir") {
		return "_aria_list_dir"
	}
	if (name == "_ariaIsDir") {
		return "_aria_is_dir"
	}
	if (name == "_ariaGetenv") {
		return "_aria_getenv"
	}
	if (name == "_ariaTcpSocket") {
		return "_aria_tcp_socket"
	}
	if (name == "_ariaTcpBind") {
		return "_aria_tcp_bind"
	}
	if (name == "_ariaTcpListen") {
		return "_aria_tcp_listen"
	}
	if (name == "_ariaTcpAccept") {
		return "_aria_tcp_accept"
	}
	if (name == "_ariaTcpRead") {
		return "_aria_tcp_read"
	}
	if (name == "_ariaTcpWrite") {
		return "_aria_tcp_write"
	}
	if (name == "_ariaTcpClose") {
		return "_aria_tcp_close"
	}
	if (name == "_ariaTcpPeerAddr") {
		return "_aria_tcp_peer_addr"
	}
	if (name == "_ariaTcpSetTimeout") {
		return "_aria_tcp_set_timeout"
	}
	if (name == "_ariaPgConnect") {
		return "_aria_pg_connect"
	}
	if (name == "_ariaPgClose") {
		return "_aria_pg_close"
	}
	if (name == "_ariaPgStatus") {
		return "_aria_pg_status"
	}
	if (name == "_ariaPgError") {
		return "_aria_pg_error"
	}
	if (name == "_ariaPgExec") {
		return "_aria_pg_exec"
	}
	if (name == "_ariaPgExecParams") {
		return "_aria_pg_exec_params"
	}
	if (name == "_ariaPgResultStatus") {
		return "_aria_pg_result_status"
	}
	if (name == "_ariaPgResultError") {
		return "_aria_pg_result_error"
	}
	if (name == "_ariaPgNrows") {
		return "_aria_pg_nrows"
	}
	if (name == "_ariaPgNcols") {
		return "_aria_pg_ncols"
	}
	if (name == "_ariaPgFieldName") {
		return "_aria_pg_field_name"
	}
	if (name == "_ariaPgGetValue") {
		return "_aria_pg_get_value"
	}
	if (name == "_ariaPgIsNull") {
		return "_aria_pg_is_null"
	}
	if (name == "_ariaPgClear") {
		return "_aria_pg_clear"
	}
	if (name == "_ariaSpawn") {
		return "_aria_spawn"
	}
	if (name == "_ariaTaskAwait") {
		return "_aria_task_await"
	}
	if (name == "_ariaChanNew") {
		return "_aria_chan_new"
	}
	if (name == "_ariaChanSend") {
		return "_aria_chan_send"
	}
	if (name == "_ariaChanRecv") {
		return "_aria_chan_recv"
	}
	if (name == "_ariaChanClose") {
		return "_aria_chan_close"
	}
	if (name == "_ariaMutexNew") {
		return "_aria_mutex_new"
	}
	if (name == "_ariaMutexLock") {
		return "_aria_mutex_lock"
	}
	if (name == "_ariaMutexUnlock") {
		return "_aria_mutex_unlock"
	}
	if (name == "i2s") {
		return "i2s"
	}
	return name
}

func _builtin_return_type(name string) int64 {
	if (((((((((name == "_ariaReadFile") || (name == "_ariaGetenv")) || (name == "_ariaTcpRead")) || (name == "_ariaTcpPeerAddr")) || (name == "_ariaPgError")) || (name == "_ariaPgResultError")) || (name == "_ariaPgFieldName")) || (name == "_ariaPgGetValue")) || (name == "i2s")) {
		return int64(12)
	}
	if (((((((((((name == "_ariaWriteFile") || (name == "_ariaAppendFile")) || (name == "_ariaWriteBinaryFile")) || (name == "_ariaArgs")) || (name == "_ariaListDir")) || (name == "_ariaTcpClose")) || (name == "_ariaPgClose")) || (name == "_ariaPgClear")) || (name == "_ariaChanClose")) || (name == "_ariaMutexLock")) || (name == "_ariaMutexUnlock")) {
		return int64(0)
	}
	return int64(1)
}

func _lower_fn(l Lowerer) Lowerer {
	ll := _ladv(l)
	_ = ll
	if (_lk(ll) != "IDENT") {
		return ll
	}
	name := _lcur(ll).Text
	_ = name
	ll = _ladv(ll)
	sig := reg_find_fn(ll.Registry, name)
	_ = sig
	if (sig.Name == "") {
		return _skip_to_next_l(ll)
	}
	if fn_is_generic(sig) {
		return _skip_to_next_l(ll)
	}
	actual_param_count := int64(0)
	_ = actual_param_count
	pi := int64(1)
	_ = pi
	for (pi < int64(len(sig.Param_names))) {
		if ((pi <= (int64(len(sig.Param_types)) - int64(1))) && (sig.Param_types[pi] == int64(12))) {
			actual_param_count = (actual_param_count + int64(2))
		} else {
			actual_param_count = (actual_param_count + int64(1))
		}
		pi = (pi + int64(1))
	}
	ll = _lset_func(ll, new_ir_func(name, actual_param_count, sig.Return_type))
	if (sig.Error_type > int64(0)) {
		ll = _lset_error_type(ll, sig.Error_type)
	}
	pi = int64(1)
	for (pi < int64(len(sig.Param_names))) {
		if ((pi <= (int64(len(sig.Param_types)) - int64(1))) && (sig.Param_types[pi] == int64(12))) {
			nt := _lnew_str_temp(ll)
			_ = nt
			ll = nt.L
			ll = _lenv_add(ll, sig.Param_names[pi], nt.Temp)
			ll = _lset_type(ll, nt.Temp, int64(12))
		} else {
			nt := _lnew_temp(ll)
			_ = nt
			ll = nt.L
			ll = _lenv_add(ll, sig.Param_names[pi], nt.Temp)
			if ((pi <= (int64(len(sig.Param_types)) - int64(1))) && (sig.Param_types[pi] > int64(0))) {
				pt := sig.Param_types[pi]
				_ = pt
				if _is_str_array_type(ll.Store, pt) {
					ll = _lset_type(ll, nt.Temp, int64(13))
				} else if _is_trait_object_type(ll.Store, pt) {
					ll = _lset_type(ll, nt.Temp, int64(500))
				} else {
					ll = _lset_type(ll, nt.Temp, pt)
				}
			}
		}
		pi = (pi + int64(1))
	}
	lb := "{"
	_ = lb
	for (((_lk(ll) != "=") && (_lk(ll) != lb)) && (_lk(ll) != "EOF")) {
		ll = _ladv(ll)
	}
	if (_lk(ll) == "=") {
		ll = _ladv(ll)
		ll = _lskip_nl(ll)
		result := _lower_expr(ll)
		_ = result
		ll = result.L
		ll = _emit_defers(ll)
		if ((ll.Fn_error_type > int64(0)) && (_lget_type(ll, result.Temp) != int64(200))) {
			ret_type := _lget_type(ll, result.Temp)
			_ = ret_type
			ok_slots := int64(2)
			_ = ok_slots
			if (ret_type == int64(12)) {
				ok_slots = int64(3)
			}
			ok_nt := _lnew_temp(ll)
			_ = ok_nt
			ll = ok_nt.L
			ll = _lemit(ll, new_inst(IrOpOpAlloc{}, ok_nt.Temp, ok_slots, int64(0), "_Result", int64(0)))
			tag_nt := _lnew_temp(ll)
			_ = tag_nt
			ll = tag_nt.L
			ll = _lemit(ll, new_inst(IrOpOpConst{}, tag_nt.Temp, int64(0), int64(0), "", int64(0)))
			ll = _lemit(ll, new_inst(IrOpOpFieldSet{}, ok_nt.Temp, tag_nt.Temp, int64(0), "_tag", int64(0)))
			if (ret_type == int64(12)) {
				ll = _lemit(ll, new_inst(IrOpOpFieldSet{}, ok_nt.Temp, result.Temp, int64(1), "_val", int64(0)))
				val_len := (result.Temp + int64(1))
				_ = val_len
				ll = _lemit(ll, new_inst(IrOpOpFieldSet{}, ok_nt.Temp, val_len, int64(2), "_val_len", int64(0)))
			} else {
				ll = _lemit(ll, new_inst(IrOpOpFieldSet{}, ok_nt.Temp, result.Temp, int64(1), "_val", int64(0)))
			}
			ll = _lemit(ll, new_inst(IrOpOpRet{}, int64(0), ok_nt.Temp, int64(0), "", int64(0)))
		} else {
			ret_type := _lget_type(ll, result.Temp)
			_ = ret_type
			ll = _lemit(ll, new_inst(IrOpOpRet{}, int64(0), result.Temp, int64(0), "", ret_type))
		}
	} else if (_lk(ll) == lb) {
		ll = _lower_fn_body(ll, sig.Return_type)
	} else {
		ll = _lskip_nl(ll)
		if (_lk(ll) == lb) {
			ll = _lower_fn_body(ll, sig.Return_type)
		}
	}
	ll = _lfinish_func(ll)
	return ll
}

func _lower_entry(l Lowerer) Lowerer {
	ll := _ladv(l)
	_ = ll
	ll = _lset_func(ll, new_ir_func("_main", int64(0), int64(0)))
	lb := "{"
	_ = lb
	if (_lk(ll) == lb) {
		ll = _lower_block(ll)
	}
	nt := _lnew_temp(ll)
	_ = nt
	ll = nt.L
	ll = _lemit(ll, new_inst(IrOpOpConst{}, nt.Temp, int64(0), int64(0), "", int64(0)))
	ll = _lemit(ll, new_inst(IrOpOpCall{}, int64(0), nt.Temp, int64(1), "_aria_exit", int64(0)))
	ll = _lfinish_func(ll)
	new_mod := ll.Module
	_ = new_mod
	new_mod = mod_set_entry(new_mod, "_main")
	return _lset_mod(ll, new_mod)
}

func _lower_test(l Lowerer) Lowerer {
	ll := _ladv(l)
	_ = ll
	test_name := ("_test_" + i2s(ll.Module.Test_count))
	_ = test_name
	test_label := ("test " + i2s(ll.Module.Test_count))
	_ = test_label
	if (_lk(ll) == "STRING") {
		test_label = _lcur(ll).Text
		ll = _ladv(ll)
	} else if (_lk(ll) == "IDENT") {
		test_label = _lcur(ll).Text
		ll = _ladv(ll)
	}
	ll = _lset_func(ll, new_ir_func(test_name, int64(0), int64(0)))
	lb := "{"
	_ = lb
	if (_lk(ll) == lb) {
		ll = _lower_block(ll)
	}
	ll = _lemit(ll, new_inst(IrOpOpRetVoid{}, int64(0), int64(0), int64(0), "", int64(0)))
	ll = _lfinish_func(ll)
	new_mod := mod_add_test(ll.Module, test_label)
	_ = new_mod
	new_mod = mod_add_string(new_mod, test_label)
	return _lset_mod(ll, new_mod)
}

func _lower_extern(l Lowerer) Lowerer {
	ll := _ladv(l)
	_ = ll
	lb := "{"
	_ = lb
	rb := "}"
	_ = rb
	lib_name := ""
	_ = lib_name
	if (_lk(ll) == "STRING") {
		lib_name = _lcur(ll).Text
		ll = _ladv(ll)
	}
	if (_lk(ll) == lb) {
		ll = _ladv(ll)
		ll = _lskip_nl(ll)
		for ((_lk(ll) != rb) && (_lk(ll) != "EOF")) {
			ll = _lskip_nl(ll)
			if (_lk(ll) == "fn") {
				ll = _lower_extern_fn(ll, lib_name)
			} else {
				ll = _ladv(ll)
			}
			ll = _lskip_nl(ll)
		}
		if (_lk(ll) == rb) {
			ll = _ladv(ll)
		}
		return ll
	}
	if (_lk(ll) == "fn") {
		ll = _lower_extern_fn(ll, lib_name)
	}
	return ll
}

func _lower_extern_fn(l Lowerer, lib string) Lowerer {
	ll := _ladv(l)
	_ = ll
	if (_lk(ll) != "IDENT") {
		return ll
	}
	name := _lcur(ll).Text
	_ = name
	ll = _ladv(ll)
	param_count := int64(0)
	_ = param_count
	is_variadic := false
	_ = is_variadic
	if (_lk(ll) == "(") {
		ll = _ladv(ll)
		for ((_lk(ll) != ")") && (_lk(ll) != "EOF")) {
			if (_lk(ll) == ".") {
				ll = _ladv(ll)
				if (_lk(ll) == ".") {
					ll = _ladv(ll)
					if (_lk(ll) == ".") {
						ll = _ladv(ll)
					}
					is_variadic = true
				}
			} else if (_lk(ll) == "IDENT") {
				ll = _ladv(ll)
				if (_lk(ll) == ":") {
					ll = _ladv(ll)
					ll = _skip_type_l(ll)
				}
				param_count = (param_count + int64(1))
			} else {
				ll = _ladv(ll)
			}
			if (_lk(ll) == ",") {
				ll = _ladv(ll)
			}
		}
		if (_lk(ll) == ")") {
			ll = _ladv(ll)
		}
	}
	ret_type := "void"
	_ = ret_type
	if (_lk(ll) == "->") {
		ll = _ladv(ll)
		if (_lk(ll) == "IDENT") {
			rt := _lcur(ll).Text
			_ = rt
			if ((((rt == "i64") || (rt == "i32")) || (rt == "i16")) || (rt == "i8")) {
				ret_type = "i64"
			}
			if ((((rt == "u64") || (rt == "u32")) || (rt == "u16")) || (rt == "u8")) {
				ret_type = "i64"
			}
			if ((rt == "f64") || (rt == "f32")) {
				ret_type = "f64"
			}
			if (rt == "bool") {
				ret_type = "i64"
			}
			if (rt == "str") {
				ret_type = "str"
			}
		}
		ll = _skip_type_l(ll)
	}
	new_mod := mod_add_extern(ll.Module, name, param_count, ret_type, lib)
	_ = new_mod
	return _lset_mod(ll, new_mod)
}

func _lower_bench(l Lowerer) Lowerer {
	ll := _ladv(l)
	_ = ll
	bench_name := ("_bench_" + i2s(ll.Module.Test_count))
	_ = bench_name
	bench_label := ("bench " + i2s(ll.Module.Test_count))
	_ = bench_label
	if (_lk(ll) == "STRING") {
		bench_label = _lcur(ll).Text
		ll = _ladv(ll)
	} else if (_lk(ll) == "IDENT") {
		bench_label = _lcur(ll).Text
		ll = _ladv(ll)
	}
	ll = _lset_func(ll, new_ir_func(bench_name, int64(0), int64(0)))
	lb := "{"
	_ = lb
	if (_lk(ll) == lb) {
		ll = _lower_block(ll)
	}
	ll = _lemit(ll, new_inst(IrOpOpRetVoid{}, int64(0), int64(0), int64(0), "", int64(0)))
	ll = _lfinish_func(ll)
	new_mod := mod_add_test(ll.Module, bench_label)
	_ = new_mod
	new_mod = mod_add_string(new_mod, bench_label)
	return _lset_mod(ll, new_mod)
}

func _lower_fixture(l Lowerer, kind string) Lowerer {
	ll := _ladv(l)
	_ = ll
	fname := "_before"
	_ = fname
	if (kind == "after") {
		fname = "_after"
	}
	ll = _lset_func(ll, new_ir_func(fname, int64(0), int64(0)))
	lb := "{"
	_ = lb
	if (_lk(ll) == lb) {
		ll = _lower_block(ll)
	}
	ll = _lemit(ll, new_inst(IrOpOpRetVoid{}, int64(0), int64(0), int64(0), "", int64(0)))
	ll = _lfinish_func(ll)
	return ll
}

func _lower_impl_d(l Lowerer, decl DeclInfo) Lowerer {
	ll := _ladv(l)
	_ = ll
	is_blanket := false
	_ = is_blanket
	bound_trait := ""
	_ = bound_trait
	if (_lk(ll) == "[") {
		is_blanket = true
		ll = _ladv(ll)
		if (_lk(ll) == "IDENT") {
			ll = _ladv(ll)
		}
		if (_lk(ll) == ":") {
			ll = _ladv(ll)
			if (_lk(ll) == "IDENT") {
				bound_trait = _lcur(ll).Text
				ll = _ladv(ll)
			}
		}
		depth := int64(1)
		_ = depth
		for ((depth > int64(0)) && (_lk(ll) != "EOF")) {
			if (_lk(ll) == "[") {
				depth = (depth + int64(1))
			}
			if (_lk(ll) == "]") {
				depth = (depth - int64(1))
			}
			if (depth > int64(0)) {
				ll = _ladv(ll)
			}
		}
		if (_lk(ll) == "]") {
			ll = _ladv(ll)
		}
	}
	impl_type := ""
	_ = impl_type
	if (_lk(ll) == "IDENT") {
		impl_type = _lcur(ll).Text
		ll = _ladv(ll)
	}
	if (_lk(ll) == "for") {
		ll = _ladv(ll)
		if (_lk(ll) == "IDENT") {
			impl_type = _lcur(ll).Text
			ll = _ladv(ll)
		}
	}
	if is_blanket {
		impl_type = ""
	}
	if (decl.Node_idx > int64(0)) {
		impl_node := pool_get(ll.Pool, decl.Node_idx)
		_ = impl_node
		ll = _lset_pos(ll, decl.Body_start)
		lb2 := "{"
		_ = lb2
		if (_lk(ll) == lb2) {
			ll = _ladv(ll)
		}
		ll = _lskip_nl(ll)
		mi := int64(0)
		_ = mi
		for (mi < impl_node.List_count) {
			midx_node := pool_get(ll.Pool, (impl_node.List_start + mi))
			_ = midx_node
			method_node := pool_get(ll.Pool, midx_node.C0)
			_ = method_node
			raw_name := method_node.S1
			_ = raw_name
			name := ((impl_type + "_") + raw_name)
			_ = name
			sig := reg_find_fn(ll.Registry, name)
			_ = sig
			if (((((sig.Name != "") && fn_is_generic(sig)) && is_blanket) && (bound_trait != "")) && method_node.B1) {
				concrete_types := treg_types_implementing(ll.Treg, bound_trait)
				_ = concrete_types
				ci := int64(1)
				_ = ci
				for (ci < int64(len(concrete_types))) {
					ctype_name := concrete_types[ci]
					_ = ctype_name
					spec_base := ((("_" + raw_name) + "_") + ctype_name)
					_ = spec_base
					existing_sig := reg_find_fn(ll.Registry, spec_base)
					_ = existing_sig
					if (existing_sig.Name == "") {
						ctype_id := _struct_type_id(ll.Registry, ctype_name)
						_ = ctype_id
						sp_ptypes := []int64{int64(0)}
						_ = sp_ptypes
						spi := int64(1)
						_ = spi
						for (spi < int64(len(sig.Param_types))) {
							pt := sig.Param_types[spi]
							_ = pt
							if ((spi == int64(1)) && (sig.Param_names[spi] == "self")) {
								pt = ctype_id
							}
							sp_ptypes = append(sp_ptypes, pt)
							spi = (spi + int64(1))
						}
						sp_sig := new_generic_fn_sig(spec_base, sig.Param_names, sp_ptypes, sig.Return_type, sig.Error_type, sig.Generic_count, sig.Generic_name_0, sig.Generic_name_1, sig.Generic_name_2, sig.Generic_bound_0, sig.Generic_bound_1, sig.Generic_bound_2, sig.Token_start)
						_ = sp_sig
						ll = _lset_reg(ll, reg_add_fn(ll.Registry, sp_sig))
					}
					ci = (ci + int64(1))
				}
			}
			if (((sig.Name != "") && (fn_is_generic(sig) == false)) && method_node.B1) {
				ll = _lower_impl_method_by_name(ll, impl_type, raw_name, sig)
				if (is_blanket && (bound_trait != "")) {
					concrete_types := treg_types_implementing(ll.Treg, bound_trait)
					_ = concrete_types
					ci := int64(1)
					_ = ci
					for (ci < int64(len(concrete_types))) {
						ctype_name := concrete_types[ci]
						_ = ctype_name
						ctype_id := _struct_type_id(ll.Registry, ctype_name)
						_ = ctype_id
						if (ctype_id > int64(0)) {
							ll = _lset_pos(ll, decl.Body_start)
							ll = _lower_blanket_method_specialized(ll, raw_name, sig, ctype_name, ctype_id)
						}
						ci = (ci + int64(1))
					}
				}
			}
			mi = (mi + int64(1))
		}
		ll = _lset_pos(ll, decl.Body_end)
		return ll
	}
	ll = _lower_impl_legacy(ll, impl_type)
	return ll
}

func _lower_blanket_method_specialized(l Lowerer, raw_name string, sig FnSig, concrete_type string, concrete_tid int64) Lowerer {
	ll := l
	_ = ll
	tlen := int64(len(ll.Tokens))
	_ = tlen
	search_pos := ll.Pos
	_ = search_pos
	lb := "{"
	_ = lb
	for (search_pos < tlen) {
		tk := token_name(ll.Tokens[search_pos].Kind)
		_ = tk
		if (tk == "fn") {
			next := (search_pos + int64(1))
			_ = next
			if (next < tlen) {
				if (ll.Tokens[next].Text == raw_name) {
					ll = _lset_pos(ll, (search_pos + int64(2)))
					actual_param_count := int64(0)
					_ = actual_param_count
					pi := int64(1)
					_ = pi
					for (pi < int64(len(sig.Param_names))) {
						if ((pi <= (int64(len(sig.Param_types)) - int64(1))) && (sig.Param_types[pi] == int64(12))) {
							actual_param_count = (actual_param_count + int64(2))
						} else {
							actual_param_count = (actual_param_count + int64(1))
						}
						pi = (pi + int64(1))
					}
					spec_fname := ((("_" + raw_name) + "_") + concrete_type)
					_ = spec_fname
					ll = _lset_func(ll, new_ir_func(spec_fname, actual_param_count, sig.Return_type))
					pi = int64(1)
					for (pi < int64(len(sig.Param_names))) {
						if ((pi <= (int64(len(sig.Param_types)) - int64(1))) && (sig.Param_types[pi] == int64(12))) {
							nt := _lnew_str_temp(ll)
							_ = nt
							ll = nt.L
							ll = _lenv_add(ll, sig.Param_names[pi], nt.Temp)
							ll = _lset_type(ll, nt.Temp, int64(12))
						} else {
							nt := _lnew_temp(ll)
							_ = nt
							ll = nt.L
							ll = _lenv_add(ll, sig.Param_names[pi], nt.Temp)
							if ((pi == int64(1)) && (sig.Param_names[pi] == "self")) {
								ll = _lset_type(ll, nt.Temp, concrete_tid)
							} else if ((pi <= (int64(len(sig.Param_types)) - int64(1))) && (sig.Param_types[pi] > int64(0))) {
								pt := sig.Param_types[pi]
								_ = pt
								if _is_str_array_type(ll.Store, pt) {
									ll = _lset_type(ll, nt.Temp, int64(13))
								} else {
									ll = _lset_type(ll, nt.Temp, pt)
								}
							}
						}
						pi = (pi + int64(1))
					}
					for (((_lk(ll) != "=") && (_lk(ll) != lb)) && (_lk(ll) != "EOF")) {
						ll = _ladv(ll)
					}
					if (_lk(ll) == "=") {
						ll = _ladv(ll)
						ll = _lskip_nl(ll)
						result := _lower_expr(ll)
						_ = result
						ll = result.L
						ret_type := _lget_type(ll, result.Temp)
						_ = ret_type
						ll = _lemit(ll, new_inst(IrOpOpRet{}, int64(0), result.Temp, int64(0), "", ret_type))
					} else if (_lk(ll) == lb) {
						ll = _lower_fn_body(ll, sig.Return_type)
					}
					ll = _lfinish_func(ll)
					return ll
				}
			}
		}
		search_pos = (search_pos + int64(1))
	}
	return ll
}

func _lower_impl_method_by_name(l Lowerer, impl_type string, raw_name string, sig FnSig) Lowerer {
	ll := l
	_ = ll
	tlen := int64(len(ll.Tokens))
	_ = tlen
	search_pos := ll.Pos
	_ = search_pos
	lb := "{"
	_ = lb
	for (search_pos < tlen) {
		tk := token_name(ll.Tokens[search_pos].Kind)
		_ = tk
		if (tk == "fn") {
			next := (search_pos + int64(1))
			_ = next
			if (next < tlen) {
				if (ll.Tokens[next].Text == raw_name) {
					ll = _lset_pos(ll, (search_pos + int64(2)))
					actual_param_count := int64(0)
					_ = actual_param_count
					pi := int64(1)
					_ = pi
					for (pi < int64(len(sig.Param_names))) {
						if ((pi <= (int64(len(sig.Param_types)) - int64(1))) && (sig.Param_types[pi] == int64(12))) {
							actual_param_count = (actual_param_count + int64(2))
						} else {
							actual_param_count = (actual_param_count + int64(1))
						}
						pi = (pi + int64(1))
					}
					ll = _lset_func(ll, new_ir_func(((impl_type + "_") + raw_name), actual_param_count, sig.Return_type))
					pi = int64(1)
					for (pi < int64(len(sig.Param_names))) {
						if ((pi <= (int64(len(sig.Param_types)) - int64(1))) && (sig.Param_types[pi] == int64(12))) {
							nt := _lnew_str_temp(ll)
							_ = nt
							ll = nt.L
							ll = _lenv_add(ll, sig.Param_names[pi], nt.Temp)
							ll = _lset_type(ll, nt.Temp, int64(12))
						} else {
							nt := _lnew_temp(ll)
							_ = nt
							ll = nt.L
							ll = _lenv_add(ll, sig.Param_names[pi], nt.Temp)
							if ((pi <= (int64(len(sig.Param_types)) - int64(1))) && (sig.Param_types[pi] > int64(0))) {
								pt := sig.Param_types[pi]
								_ = pt
								if _is_str_array_type(ll.Store, pt) {
									ll = _lset_type(ll, nt.Temp, int64(13))
								} else {
									ll = _lset_type(ll, nt.Temp, pt)
								}
							}
						}
						pi = (pi + int64(1))
					}
					for (((_lk(ll) != "=") && (_lk(ll) != lb)) && (_lk(ll) != "EOF")) {
						ll = _ladv(ll)
					}
					if (_lk(ll) == "=") {
						ll = _ladv(ll)
						ll = _lskip_nl(ll)
						result := _lower_expr(ll)
						_ = result
						ll = result.L
						ret_type := _lget_type(ll, result.Temp)
						_ = ret_type
						ll = _lemit(ll, new_inst(IrOpOpRet{}, int64(0), result.Temp, int64(0), "", ret_type))
					} else if (_lk(ll) == lb) {
						ll = _lower_fn_body(ll, sig.Return_type)
					}
					ll = _lfinish_func(ll)
					return ll
				}
			}
		}
		search_pos = (search_pos + int64(1))
	}
	return ll
}

func _lower_impl_legacy(l Lowerer, impl_type string) Lowerer {
	ll := l
	_ = ll
	lb := "{"
	_ = lb
	rb := "}"
	_ = rb
	for ((_lk(ll) != lb) && (_lk(ll) != "EOF")) {
		ll = _ladv(ll)
	}
	if (_lk(ll) != lb) {
		return ll
	}
	ll = _ladv(ll)
	ll = _lskip_nl(ll)
	for ((_lk(ll) != rb) && (_lk(ll) != "EOF")) {
		ll = _lskip_nl(ll)
		ll = _lskip_pub(ll)
		if (_lk(ll) == "fn") {
			ll = _lower_impl_fn(ll, impl_type)
		} else if ((_lk(ll) == rb) || (_lk(ll) == "EOF")) {
			return ll
		} else {
			ll = _ladv(ll)
		}
		ll = _lskip_nl(ll)
	}
	if (_lk(ll) == rb) {
		ll = _ladv(ll)
	}
	return ll
}

func _lower_impl(l Lowerer) Lowerer {
	ll := _ladv(l)
	_ = ll
	impl_type := ""
	_ = impl_type
	if (_lk(ll) == "IDENT") {
		impl_type = _lcur(ll).Text
		ll = _ladv(ll)
	}
	if (_lk(ll) == "for") {
		ll = _ladv(ll)
		if (_lk(ll) == "IDENT") {
			impl_type = _lcur(ll).Text
			ll = _ladv(ll)
		}
	}
	lb := "{"
	_ = lb
	rb := "}"
	_ = rb
	for ((_lk(ll) != lb) && (_lk(ll) != "EOF")) {
		ll = _ladv(ll)
	}
	if (_lk(ll) != lb) {
		return ll
	}
	ll = _ladv(ll)
	ll = _lskip_nl(ll)
	for ((_lk(ll) != rb) && (_lk(ll) != "EOF")) {
		ll = _lskip_nl(ll)
		ll = _lskip_pub(ll)
		if (_lk(ll) == "fn") {
			ll = _lower_impl_fn(ll, impl_type)
		} else if ((_lk(ll) == rb) || (_lk(ll) == "EOF")) {
			return ll
		} else {
			ll = _ladv(ll)
		}
		ll = _lskip_nl(ll)
	}
	if (_lk(ll) == rb) {
		ll = _ladv(ll)
	}
	return ll
}

func _lower_impl_fn(l Lowerer, impl_type string) Lowerer {
	ll := _ladv(l)
	_ = ll
	if (_lk(ll) != "IDENT") {
		return ll
	}
	raw_name := _lcur(ll).Text
	_ = raw_name
	name := ((impl_type + "_") + raw_name)
	_ = name
	ll = _ladv(ll)
	sig := reg_find_fn(ll.Registry, name)
	_ = sig
	if (sig.Name == "") {
		return _skip_to_next_l(ll)
	}
	if fn_is_generic(sig) {
		return _skip_to_next_l(ll)
	}
	actual_param_count := int64(0)
	_ = actual_param_count
	pi := int64(1)
	_ = pi
	for (pi < int64(len(sig.Param_names))) {
		if ((pi <= (int64(len(sig.Param_types)) - int64(1))) && (sig.Param_types[pi] == int64(12))) {
			actual_param_count = (actual_param_count + int64(2))
		} else {
			actual_param_count = (actual_param_count + int64(1))
		}
		pi = (pi + int64(1))
	}
	ll = _lset_func(ll, new_ir_func(name, actual_param_count, sig.Return_type))
	pi = int64(1)
	for (pi < int64(len(sig.Param_names))) {
		if ((pi <= (int64(len(sig.Param_types)) - int64(1))) && (sig.Param_types[pi] == int64(12))) {
			nt := _lnew_str_temp(ll)
			_ = nt
			ll = nt.L
			ll = _lenv_add(ll, sig.Param_names[pi], nt.Temp)
			ll = _lset_type(ll, nt.Temp, int64(12))
		} else {
			nt := _lnew_temp(ll)
			_ = nt
			ll = nt.L
			ll = _lenv_add(ll, sig.Param_names[pi], nt.Temp)
			if ((pi <= (int64(len(sig.Param_types)) - int64(1))) && (sig.Param_types[pi] > int64(0))) {
				pt := sig.Param_types[pi]
				_ = pt
				if _is_str_array_type(ll.Store, pt) {
					ll = _lset_type(ll, nt.Temp, int64(13))
				} else if _is_trait_object_type(ll.Store, pt) {
					ll = _lset_type(ll, nt.Temp, int64(500))
				} else {
					ll = _lset_type(ll, nt.Temp, pt)
				}
			}
		}
		pi = (pi + int64(1))
	}
	lb := "{"
	_ = lb
	for (((_lk(ll) != "=") && (_lk(ll) != lb)) && (_lk(ll) != "EOF")) {
		ll = _ladv(ll)
	}
	if (_lk(ll) == "=") {
		ll = _ladv(ll)
		ll = _lskip_nl(ll)
		result := _lower_expr(ll)
		_ = result
		ll = result.L
		ret_type := _lget_type(ll, result.Temp)
		_ = ret_type
		ll = _lemit(ll, new_inst(IrOpOpRet{}, int64(0), result.Temp, int64(0), "", ret_type))
	} else if (_lk(ll) == lb) {
		ll = _lower_fn_body(ll, sig.Return_type)
	}
	ll = _lfinish_func(ll)
	return ll
}

func _lower_fn_body(l Lowerer, return_type int64) Lowerer {
	lb := "{"
	_ = lb
	rb := "}"
	_ = rb
	if (_lk(l) != lb) {
		return l
	}
	ll := _ladv(l)
	_ = ll
	ll = _lskip_nl(ll)
	last_temp := -int64(1)
	_ = last_temp
	for ((_lk(ll) != rb) && (_lk(ll) != "EOF")) {
		k := _lk(ll)
		_ = k
		if (k == "return") {
			ll = _lower_stmt(ll)
			last_temp = -int64(1)
		} else if (((((((((k == "mut") || (k == "break")) || (k == "continue")) || (k == "for")) || (k == "while")) || (k == "loop")) || (k == "assert")) || (k == "defer")) || (k == "with")) {
			ll = _lower_stmt(ll)
		} else if (k == "if") {
			result := _lower_if(ll)
			_ = result
			ll = result.L
			last_temp = result.Temp
		} else if (k == "match") {
			result := _lower_match(ll)
			_ = result
			ll = result.L
			last_temp = result.Temp
		} else {
			result := _lower_expr_result(ll)
			_ = result
			ll = result.L
			last_temp = result.Temp
		}
		ll = _lskip_nl(ll)
	}
	if (_lk(ll) == rb) {
		ll = _ladv(ll)
	}
	if ((return_type > int64(0)) && (last_temp >= int64(0))) {
		inst_count := int64(len(ll.Current_func.Insts))
		_ = inst_count
		already_returned := false
		_ = already_returned
		if (inst_count > int64(1)) {
			last_op := ir_op_name(ll.Current_func.Insts[(inst_count - int64(1))].Op)
			_ = last_op
			if ((last_op == "Ret") || (last_op == "RetVoid")) {
				already_returned = true
			}
		}
		if (already_returned == false) {
			ll = _emit_defers(ll)
			ret_type := _lget_type(ll, last_temp)
			_ = ret_type
			ll = _lemit(ll, new_inst(IrOpOpRet{}, int64(0), last_temp, int64(0), "", ret_type))
		}
	}
	return ll
}

func _lower_expr_result(l Lowerer) LT {
	if (_lk(l) == "IDENT") {
		name := _lcur(l).Text
		_ = name
		next_pos := (l.Pos + int64(1))
		_ = next_pos
		if (next_pos < int64(len(l.Tokens))) {
			next_k := token_name(l.Tokens[next_pos].Kind)
			_ = next_k
			if (next_k == ":=") {
				ll := _ladv(l)
				_ = ll
				ll = _ladv(ll)
				result := _lower_expr(ll)
				_ = result
				ll = result.L
				ll = _lenv_add(ll, name, result.Temp)
				return LT{L: ll, Temp: result.Temp}
			}
			if (next_k == "=") {
				ll := _ladv(l)
				_ = ll
				ll = _ladv(ll)
				result := _lower_expr(ll)
				_ = result
				ll = result.L
				slot := _lenv_lookup(ll, name)
				_ = slot
				if (slot >= int64(0)) {
					ll = _lemit(ll, new_inst(IrOpOpStore{}, slot, result.Temp, int64(0), "", int64(0)))
					result_type := _lget_type(ll, result.Temp)
					_ = result_type
					if (result_type == int64(12)) {
						ll = _lemit(ll, new_inst(IrOpOpStore{}, (slot + int64(1)), (result.Temp + int64(1)), int64(0), "", int64(0)))
					}
				}
				return LT{L: ll, Temp: result.Temp}
			}
		}
	}
	return _lower_expr(l)
}

func _lower_block(l Lowerer) Lowerer {
	lb := "{"
	_ = lb
	rb := "}"
	_ = rb
	if (_lk(l) != lb) {
		return l
	}
	ll := _ladv(l)
	_ = ll
	ll = _lskip_nl(ll)
	scope_start := int64(len(ll.Env_names))
	_ = scope_start
	for ((_lk(ll) != rb) && (_lk(ll) != "EOF")) {
		ll = _lower_stmt(ll)
		ll = _lskip_nl(ll)
	}
	ll = _emit_drops_scoped(ll, scope_start)
	if (_lk(ll) == rb) {
		ll = _ladv(ll)
	}
	return ll
}

func _lower_stmt(l Lowerer) Lowerer {
	k := _lk(l)
	_ = k
	if (k == "mut") {
		return _lower_mut_binding(l)
	}
	if (k == "return") {
		return _lower_return(l)
	}
	if (k == "break") {
		ll := _ladv(l)
		_ = ll
		if (l.Loop_end >= int64(0)) {
			ll = _lemit(ll, new_inst(IrOpOpJump{}, int64(0), l.Loop_end, int64(0), "", int64(0)))
		}
		return ll
	}
	if (k == "continue") {
		ll := _ladv(l)
		_ = ll
		if (l.Loop_start >= int64(0)) {
			ll = _lemit(ll, new_inst(IrOpOpJump{}, int64(0), l.Loop_start, int64(0), "", int64(0)))
		}
		return ll
	}
	if (k == "yield") {
		ll := _ladv(l)
		_ = ll
		result := _lower_expr(ll)
		_ = result
		return result.L
	}
	if (k == "for") {
		return _lower_for(l)
	}
	if (k == "while") {
		return _lower_while(l)
	}
	if (k == "loop") {
		return _lower_loop(l)
	}
	if (k == "if") {
		result := _lower_if(l)
		_ = result
		return result.L
	}
	if (k == "assert") {
		return _lower_assert(l)
	}
	if (k == "mock") {
		ll := _ladv(l)
		_ = ll
		if (_lk(ll) == "IDENT") {
			mock_name := _lcur(ll).Text
			_ = mock_name
			ll = _ladv(ll)
			if (_lk(ll) == "with") {
				ll = _ladv(ll)
				mock_expr := _lower_expr(ll)
				_ = mock_expr
				ll = mock_expr.L
				ll = _lenv_add(ll, ("_mock_" + mock_name), mock_expr.Temp)
			}
		}
		return ll
	}
	if (k == "with") {
		ll := _ladv(l)
		_ = ll
		if (_lk(ll) != "[") {
			res_name := ""
			_ = res_name
			if (_lk(ll) == "IDENT") {
				next := (ll.Pos + int64(1))
				_ = next
				if ((next < int64(len(ll.Tokens))) && (token_name(ll.Tokens[next].Kind) == ":=")) {
					res_name = _lcur(ll).Text
					ll = _ladv(ll)
					ll = _ladv(ll)
				}
			}
			resource := _lower_expr(ll)
			_ = resource
			ll = resource.L
			if (int64(len(res_name)) > int64(0)) {
				ll = _lenv_add(ll, res_name, resource.Temp)
			}
			res_slot := resource.Temp
			_ = res_slot
			obj_type := _lget_type(ll, res_slot)
			_ = obj_type
			sname := _struct_name_for_type(ll.Store, ll.Registry, obj_type)
			_ = sname
			has_drop := false
			_ = has_drop
			drop_fn := ""
			_ = drop_fn
			if (sname != "") {
				drop_fn = (sname + "_drop")
				drop_sig := reg_find_fn(ll.Registry, drop_fn)
				_ = drop_sig
				if (drop_sig.Name != "") {
					has_drop = true
				}
			}
			saved_defer_count := (int64(len(ll.Defer_starts)) - int64(1))
			_ = saved_defer_count
			scope_start := int64(len(ll.Env_names))
			_ = scope_start
			ll = _lower_block(ll)
			if has_drop {
				nt := _lnew_temp(ll)
				_ = nt
				ll = nt.L
				ll = _lemit(ll, new_inst(IrOpOpCall{}, nt.Temp, res_slot, int64(1), drop_fn, int64(0)))
			}
			return ll
		}
		for ((_lk(ll) != "NEWLINE") && (_lk(ll) != "EOF")) {
			ll = _ladv(ll)
		}
		return ll
	}
	if (k == "spawn") {
		result := _lower_spawn(l)
		_ = result
		return result.L
	}
	if (k == "scope") {
		result := _lower_scope(l)
		_ = result
		return result.L
	}
	if (k == "select") {
		result := _lower_select(l)
		_ = result
		return result.L
	}
	if (k == "defer") {
		ll := _ladv(l)
		_ = ll
		defer_start := ll.Pos
		_ = defer_start
		lb2 := "{"
		_ = lb2
		if (_lk(ll) == lb2) {
			ll = _skip_braces_l(ll)
		} else {
			for ((_lk(ll) != "NEWLINE") && (_lk(ll) != "EOF")) {
				ll = _ladv(ll)
			}
		}
		defer_end := ll.Pos
		_ = defer_end
		ll = _lappend_defer(ll, defer_start, defer_end)
		return ll
	}
	return _lower_expr_or_binding(l)
}

func _lower_mut_binding(l Lowerer) Lowerer {
	ll := _ladv(l)
	_ = ll
	if (_lk(ll) != "IDENT") {
		return ll
	}
	name := _lcur(ll).Text
	_ = name
	ll = _ladv(ll)
	is_str_type := false
	_ = is_str_type
	if (_lk(ll) == ":") {
		ll = _ladv(ll)
		if ((_lk(ll) == "IDENT") && (_lcur(ll).Text == "str")) {
			is_str_type = true
		}
		ll = _skip_type_l(ll)
		if (_lk(ll) == "=") {
			ll = _ladv(ll)
		}
	} else if (_lk(ll) == ":=") {
		ll = _ladv(ll)
	}
	result := _lower_expr(ll)
	_ = result
	ll = result.L
	if is_str_type {
		ll = _lset_type(ll, result.Temp, int64(12))
	}
	ll = _lenv_add(ll, name, result.Temp)
	return ll
}

func _lower_return(l Lowerer) Lowerer {
	ll := _ladv(l)
	_ = ll
	rb := "}"
	_ = rb
	if (((_lk(ll) == "NEWLINE") || (_lk(ll) == "EOF")) || (_lk(ll) == rb)) {
		ll = _emit_defers(ll)
		ll = _lemit(ll, new_inst(IrOpOpRetVoid{}, int64(0), int64(0), int64(0), "", int64(0)))
		return ll
	}
	result := _lower_expr(ll)
	_ = result
	ll = result.L
	ll = _emit_defers(ll)
	if ((ll.Fn_error_type > int64(0)) && (_lget_type(ll, result.Temp) != int64(200))) {
		ret_type := _lget_type(ll, result.Temp)
		_ = ret_type
		ok_slots := int64(2)
		_ = ok_slots
		if (ret_type == int64(12)) {
			ok_slots = int64(3)
		}
		ok_nt := _lnew_temp(ll)
		_ = ok_nt
		ll = ok_nt.L
		ll = _lemit(ll, new_inst(IrOpOpAlloc{}, ok_nt.Temp, ok_slots, int64(0), "_Result", int64(0)))
		tag_nt := _lnew_temp(ll)
		_ = tag_nt
		ll = tag_nt.L
		ll = _lemit(ll, new_inst(IrOpOpConst{}, tag_nt.Temp, int64(0), int64(0), "", int64(0)))
		ll = _lemit(ll, new_inst(IrOpOpFieldSet{}, ok_nt.Temp, tag_nt.Temp, int64(0), "_tag", int64(0)))
		if (ret_type == int64(12)) {
			ll = _lemit(ll, new_inst(IrOpOpFieldSet{}, ok_nt.Temp, result.Temp, int64(1), "_val", int64(0)))
			val_len := (result.Temp + int64(1))
			_ = val_len
			ll = _lemit(ll, new_inst(IrOpOpFieldSet{}, ok_nt.Temp, val_len, int64(2), "_val_len", int64(0)))
		} else {
			ll = _lemit(ll, new_inst(IrOpOpFieldSet{}, ok_nt.Temp, result.Temp, int64(1), "_val", int64(0)))
		}
		ll = _lemit(ll, new_inst(IrOpOpRet{}, int64(0), ok_nt.Temp, int64(0), "", int64(0)))
	} else {
		ret_type := _lget_type(ll, result.Temp)
		_ = ret_type
		ll = _lemit(ll, new_inst(IrOpOpRet{}, int64(0), result.Temp, int64(0), "", ret_type))
	}
	return ll
}

func _lower_expr_or_binding(l Lowerer) Lowerer {
	if (_lk(l) == "(") {
		scan := (l.Pos + int64(1))
		_ = scan
		is_tuple_destr := false
		_ = is_tuple_destr
		for (scan < int64(len(l.Tokens))) {
			sk := token_name(l.Tokens[scan].Kind)
			_ = sk
			if (sk == ")") {
				next2 := (scan + int64(1))
				_ = next2
				if ((next2 < int64(len(l.Tokens))) && (token_name(l.Tokens[next2].Kind) == ":=")) {
					is_tuple_destr = true
				}
				scan = int64(len(l.Tokens))
			} else if (((sk == "IDENT") || (sk == ",")) || (sk == "_")) {
				scan = (scan + int64(1))
			} else {
				scan = int64(len(l.Tokens))
			}
		}
		if is_tuple_destr {
			ll := _ladv(l)
			_ = ll
			names := []string{""}
			_ = names
			for ((_lk(ll) != ")") && (_lk(ll) != "EOF")) {
				if (_lk(ll) == "IDENT") {
					names = append(names, _lcur(ll).Text)
				}
				ll = _ladv(ll)
				if (_lk(ll) == ",") {
					ll = _ladv(ll)
				}
			}
			if (_lk(ll) == ")") {
				ll = _ladv(ll)
			}
			if (_lk(ll) == ":=") {
				ll = _ladv(ll)
			}
			result := _lower_expr(ll)
			_ = result
			ll = result.L
			ni := int64(1)
			_ = ni
			for (ni < int64(len(names))) {
				elt_nt := _lnew_temp(ll)
				_ = elt_nt
				ll = elt_nt.L
				ll = _lemit(ll, new_inst(IrOpOpFieldGet{}, elt_nt.Temp, result.Temp, (ni - int64(1)), names[ni], int64(0)))
				ll = _lenv_add(ll, names[ni], elt_nt.Temp)
				ni = (ni + int64(1))
			}
			return ll
		}
	}
	_is_binding_name := (_lk(l) == "IDENT")
	_ = _is_binding_name
	if ((_is_binding_name == false) && ((l.Pos + int64(1)) < int64(len(l.Tokens)))) {
		_nk := token_name(l.Tokens[(l.Pos + int64(1))].Kind)
		_ = _nk
		if ((_nk == ":=") || (_nk == ":")) {
			_ck := _lk(l)
			_ = _ck
			if (((((((((((((((((((((((_ck != "fn") && (_ck != "type")) && (_ck != "struct")) && (_ck != "enum")) && (_ck != "trait")) && (_ck != "impl")) && (_ck != "const")) && (_ck != "entry")) && (_ck != "test")) && (_ck != "if")) && (_ck != "match")) && (_ck != "for")) && (_ck != "while")) && (_ck != "loop")) && (_ck != "return")) && (_ck != "mut")) && (_ck != "break")) && (_ck != "continue")) && (_ck != "defer")) && (_ck != "assert")) && (_ck != "pub")) && (_ck != "use")) && (_ck != "mod")) {
				_is_binding_name = true
			}
		}
	}
	if _is_binding_name {
		name := _lcur(l).Text
		_ = name
		next_pos := (l.Pos + int64(1))
		_ = next_pos
		if (next_pos < int64(len(l.Tokens))) {
			next_k := token_name(l.Tokens[next_pos].Kind)
			_ = next_k
			if (next_k == ":=") {
				ll := _ladv(l)
				_ = ll
				ll = _ladv(ll)
				result := _lower_expr(ll)
				_ = result
				ll = result.L
				ll = _lenv_add(ll, name, result.Temp)
				return ll
			}
			if ((((((next_k == "+") || (next_k == "-")) || (next_k == "*")) || (next_k == "/")) || (next_k == "%"))) {
				eq_pos := (l.Pos + int64(2))
				_ = eq_pos
				if ((eq_pos < int64(len(l.Tokens))) && (token_name(l.Tokens[eq_pos].Kind) == "=")) {
					op_k := next_k
					_ = op_k
					ll := _ladv(l)
					_ = ll
					ll = _ladv(ll)
					ll = _ladv(ll)
					slot := _lenv_lookup(ll, name)
					_ = slot
					if (slot >= int64(0)) {
						cur_nt := _lnew_temp(ll)
						_ = cur_nt
						ll = cur_nt.L
						ll = _lemit(ll, new_inst(IrOpOpLoad{}, cur_nt.Temp, slot, int64(0), "", int64(0)))
						rhs := _lower_expr(ll)
						_ = rhs
						ll = rhs.L
						res_nt := _lnew_temp(ll)
						_ = res_nt
						ll = res_nt.L
						cur_type := _lget_type(ll, slot)
						_ = cur_type
						if (_is_float_type(cur_type) || _is_float_type(_lget_type(ll, rhs.Temp))) {
							ll = _lemit(ll, new_inst(_arith_op(op_k), res_nt.Temp, cur_nt.Temp, rhs.Temp, "f64", int64(0)))
							ll = _lset_type(ll, res_nt.Temp, int64(9))
						} else {
							ll = _lemit(ll, new_inst(_arith_op(op_k), res_nt.Temp, cur_nt.Temp, rhs.Temp, "", int64(0)))
						}
						ll = _lemit(ll, new_inst(IrOpOpStore{}, slot, res_nt.Temp, int64(0), "", int64(0)))
					}
					return ll
				}
			}
			if (next_k == "=") {
				ll := _ladv(l)
				_ = ll
				ll = _ladv(ll)
				result := _lower_expr(ll)
				_ = result
				ll = result.L
				slot := _lenv_lookup(ll, name)
				_ = slot
				if (slot >= int64(0)) {
					ll = _lemit(ll, new_inst(IrOpOpStore{}, slot, result.Temp, int64(0), "", int64(0)))
					result_type := _lget_type(ll, result.Temp)
					_ = result_type
					if (result_type == int64(12)) {
						ll = _lemit(ll, new_inst(IrOpOpStore{}, (slot + int64(1)), (result.Temp + int64(1)), int64(0), "", int64(0)))
					}
					if (result_type > int64(0)) {
						ll = _lset_type(ll, slot, result_type)
					}
				}
				return ll
			}
		}
	}
	if (_lk(l) == "IDENT") {
		rname := _lcur(l).Text
		_ = rname
		rvtag := _variant_tag(l.Registry, rname)
		_ = rvtag
		if (rvtag > int64(0)) {
			rp := (l.Pos + int64(1))
			_ = rp
			lb := "{"
			_ = lb
			rb := "}"
			_ = rb
			if (rp < int64(len(l.Tokens))) {
				rpk := token_name(l.Tokens[rp].Kind)
				_ = rpk
				if ((rpk == "(") || (rpk == lb)) {
					close := ")"
					_ = close
					if (rpk == lb) {
						close = rb
					}
					depth := int64(1)
					_ = depth
					rp = (rp + int64(1))
					for ((rp < int64(len(l.Tokens))) && (depth > int64(0))) {
						rpk2 := token_name(l.Tokens[rp].Kind)
						_ = rpk2
						if (rpk2 == rpk) {
							depth = (depth + int64(1))
						}
						if (rpk2 == close) {
							depth = (depth - int64(1))
						}
						rp = (rp + int64(1))
					}
					if ((rp < int64(len(l.Tokens))) && (token_name(l.Tokens[rp].Kind) == ":=")) {
						ll := _ladv(l)
						_ = ll
						rb_names := []string{""}
						_ = rb_names
						if (_lk(ll) == "(") {
							ll = _ladv(ll)
							ll = _lskip_nl(ll)
							for ((_lk(ll) != ")") && (_lk(ll) != "EOF")) {
								if (_lk(ll) == "IDENT") {
									rb_names = append(rb_names, _lcur(ll).Text)
									ll = _ladv(ll)
								} else {
									ll = _ladv(ll)
								}
								ll = _lskip_nl(ll)
								if (_lk(ll) == ",") {
									ll = _ladv(ll)
								}
								ll = _lskip_nl(ll)
							}
							if (_lk(ll) == ")") {
								ll = _ladv(ll)
							}
						} else if (_lk(ll) == lb) {
							ll = _ladv(ll)
							ll = _lskip_nl(ll)
							for ((_lk(ll) != rb) && (_lk(ll) != "EOF")) {
								if (_lk(ll) == "IDENT") {
									fname := _lcur(ll).Text
									_ = fname
									ll = _ladv(ll)
									if (_lk(ll) == ":") {
										ll = _ladv(ll)
										ll = _lskip_nl(ll)
										if (_lk(ll) == "IDENT") {
											rb_names = append(rb_names, _lcur(ll).Text)
											ll = _ladv(ll)
										}
									} else {
										rb_names = append(rb_names, fname)
									}
								} else {
									ll = _ladv(ll)
								}
								ll = _lskip_nl(ll)
								if (_lk(ll) == ",") {
									ll = _ladv(ll)
								}
								ll = _lskip_nl(ll)
							}
							if (_lk(ll) == rb) {
								ll = _ladv(ll)
							}
						}
						if (_lk(ll) == ":=") {
							ll = _ladv(ll)
						}
						rhs := _lower_expr(ll)
						_ = rhs
						ll = rhs.L
						rhs_tag := _lnew_temp(ll)
						_ = rhs_tag
						ll = rhs_tag.L
						rparent := _variant_parent(ll.Registry, rname)
						_ = rparent
						if ((rparent.Name != "") && sum_has_data(rparent)) {
							ll = _lemit(ll, new_inst(IrOpOpFieldGet{}, rhs_tag.Temp, rhs.Temp, int64(0), "_tag", int64(0)))
						} else {
							ll = _lemit(ll, new_inst(IrOpOpStore{}, rhs_tag.Temp, rhs.Temp, int64(0), "", int64(0)))
						}
						rtag_nt := _lnew_temp(ll)
						_ = rtag_nt
						ll = rtag_nt.L
						ll = _lemit(ll, new_inst(IrOpOpConst{}, rtag_nt.Temp, rvtag, int64(0), "", int64(0)))
						rcmp := _lnew_temp(ll)
						_ = rcmp
						ll = rcmp.L
						ll = _lemit(ll, new_inst(IrOpOpEq{}, rcmp.Temp, rhs_tag.Temp, rtag_nt.Temp, "", int64(0)))
						relse_lbl := _lnew_label(ll)
						_ = relse_lbl
						ll = relse_lbl.L
						rcont_lbl := _lnew_label(ll)
						_ = rcont_lbl
						ll = rcont_lbl.L
						ll = _lemit(ll, new_inst(IrOpOpBranchFalse{}, int64(0), rcmp.Temp, relse_lbl.Temp, "", int64(0)))
						if (int64(len(rb_names)) > int64(1)) {
							rbi := int64(1)
							_ = rbi
							for (rbi < int64(len(rb_names))) {
								rd_slot := int64(1)
								_ = rd_slot
								rsi := int64(1)
								_ = rsi
								for (rsi < rbi) {
									rft := variant_field_type(rparent, rvtag, (rsi - int64(1)))
									_ = rft
									if (rft == int64(12)) {
										rd_slot = (rd_slot + int64(2))
									} else {
										rd_slot = (rd_slot + int64(1))
									}
									rsi = (rsi + int64(1))
								}
								rft := variant_field_type(rparent, rvtag, (rbi - int64(1)))
								_ = rft
								if (rft == int64(12)) {
									rbnt := _lnew_str_temp(ll)
									_ = rbnt
									ll = rbnt.L
									ll = _lemit(ll, new_inst(IrOpOpFieldGet{}, rbnt.Temp, rhs.Temp, rd_slot, "", int64(0)))
									ll = _lemit(ll, new_inst(IrOpOpFieldGet{}, (rbnt.Temp + int64(1)), rhs.Temp, (rd_slot + int64(1)), "", int64(0)))
									ll = _lset_type(ll, rbnt.Temp, int64(12))
									ll = _lenv_add(ll, rb_names[rbi], rbnt.Temp)
								} else {
									rbnt := _lnew_temp(ll)
									_ = rbnt
									ll = rbnt.L
									ll = _lemit(ll, new_inst(IrOpOpFieldGet{}, rbnt.Temp, rhs.Temp, rd_slot, "", int64(0)))
									if (rft != int64(0)) {
										ll = _lset_type(ll, rbnt.Temp, rft)
									}
									ll = _lenv_add(ll, rb_names[rbi], rbnt.Temp)
								}
								rbi = (rbi + int64(1))
							}
						}
						ll = _lemit(ll, new_inst(IrOpOpJump{}, int64(0), rcont_lbl.Temp, int64(0), "", int64(0)))
						ll = _lemit(ll, new_inst(IrOpOpLabel{}, int64(0), relse_lbl.Temp, int64(0), "", int64(0)))
						if (_lk(ll) == "else") {
							ll = _ladv(ll)
						}
						ll = _lskip_nl(ll)
						if (_lk(ll) == "|") {
							ll = _ladv(ll)
							if (_lk(ll) == "IDENT") {
								ename := _lcur(ll).Text
								_ = ename
								ll = _ladv(ll)
								ll = _lenv_add(ll, ename, rhs.Temp)
							}
							if (_lk(ll) == "|") {
								ll = _ladv(ll)
							}
						}
						if (_lk(ll) == lb) {
							ll = _lower_block(ll)
						}
						ll = _lemit(ll, new_inst(IrOpOpLabel{}, int64(0), rcont_lbl.Temp, int64(0), "", int64(0)))
						return ll
					}
				}
			}
		}
	}
	result := _lower_expr(l)
	_ = result
	return result.L
}

func _lower_multiplicative(l Lowerer) LT {
	result := _lower_atom(l)
	_ = result
	ll := result.L
	_ = ll
	left := result.Temp
	_ = left
	running := true
	_ = running
	for running {
		k := _lk(ll)
		_ = k
		if (((k == "*") || (k == "/")) || (k == "%")) {
			ll = _ladv(ll)
			right := _lower_atom(ll)
			_ = right
			ll = right.L
			nt := _lnew_temp(ll)
			_ = nt
			ll = nt.L
			left_type := _lget_type(ll, left)
			_ = left_type
			right_type := _lget_type(ll, right.Temp)
			_ = right_type
			if (_is_float_type(left_type) || _is_float_type(right_type)) {
				ll = _lemit(ll, new_inst(_arith_op(k), nt.Temp, left, right.Temp, "f64", int64(0)))
				ll = _lset_type(ll, nt.Temp, int64(9))
			} else {
				ll = _lemit(ll, new_inst(_arith_op(k), nt.Temp, left, right.Temp, "", int64(0)))
				ll = _lset_type(ll, nt.Temp, int64(1))
			}
			left = nt.Temp
		} else {
			running = false
		}
	}
	return LT{L: ll, Temp: left}
}

func _lower_additive(l Lowerer) LT {
	result := _lower_multiplicative(l)
	_ = result
	ll := result.L
	_ = ll
	left := result.Temp
	_ = left
	running := true
	_ = running
	for running {
		k := _lk(ll)
		_ = k
		if ((k == "+") || (k == "-")) {
			ll = _ladv(ll)
			right := _lower_multiplicative(ll)
			_ = right
			ll = right.L
			left_type := _lget_type(ll, left)
			_ = left_type
			right_type := _lget_type(ll, right.Temp)
			_ = right_type
			if ((k == "+") && (left_type == int64(12))) {
				nt := _lnew_str_temp(ll)
				_ = nt
				ll = nt.L
				ll = _lemit(ll, new_inst(IrOpOpStrConcat{}, nt.Temp, left, right.Temp, "", int64(0)))
				ll = _lset_type(ll, nt.Temp, int64(12))
				left = nt.Temp
			} else if (_is_float_type(left_type) || _is_float_type(right_type)) {
				nt := _lnew_temp(ll)
				_ = nt
				ll = nt.L
				ll = _lemit(ll, new_inst(_arith_op(k), nt.Temp, left, right.Temp, "f64", int64(0)))
				ll = _lset_type(ll, nt.Temp, int64(9))
				left = nt.Temp
			} else {
				nt := _lnew_temp(ll)
				_ = nt
				ll = nt.L
				ll = _lemit(ll, new_inst(_arith_op(k), nt.Temp, left, right.Temp, "", int64(0)))
				ll = _lset_type(ll, nt.Temp, int64(1))
				left = nt.Temp
			}
		} else {
			running = false
		}
	}
	return LT{L: ll, Temp: left}
}

func _lower_comparison(l Lowerer) LT {
	result := _lower_additive(l)
	_ = result
	ll := result.L
	_ = ll
	left := result.Temp
	_ = left
	k := _lk(ll)
	_ = k
	if ((((((k == "==") || (k == "!=")) || (k == "<")) || (k == ">")) || (k == "<=")) || (k == ">=")) {
		ll = _ladv(ll)
		right := _lower_additive(ll)
		_ = right
		ll = right.L
		nt := _lnew_temp(ll)
		_ = nt
		ll = nt.L
		left_type := _lget_type(ll, left)
		_ = left_type
		if ((((k == "==") || (k == "!="))) && (left_type == int64(12))) {
			ll = _lemit(ll, new_inst(IrOpOpStrEq{}, nt.Temp, left, right.Temp, "", int64(0)))
			if (k == "!=") {
				not_nt := _lnew_temp(ll)
				_ = not_nt
				ll = not_nt.L
				ll = _lemit(ll, new_inst(IrOpOpNot{}, not_nt.Temp, nt.Temp, int64(0), "", int64(0)))
				left = not_nt.Temp
				ll = _lset_type(ll, not_nt.Temp, int64(1))
			} else {
				left = nt.Temp
			}
		} else if ((((k == "==") || (k == "!="))) && _is_struct_type(ll.Store, ll.Registry, left_type)) {
			sname := _struct_name_for_type(ll.Store, ll.Registry, left_type)
			_ = sname
			if (sname != "") {
				eq_result := _lower_struct_eq(ll, left, right.Temp, sname, nt.Temp)
				_ = eq_result
				ll = eq_result
				if (k == "!=") {
					not_nt := _lnew_temp(ll)
					_ = not_nt
					ll = not_nt.L
					ll = _lemit(ll, new_inst(IrOpOpNot{}, not_nt.Temp, nt.Temp, int64(0), "", int64(0)))
					left = not_nt.Temp
					ll = _lset_type(ll, not_nt.Temp, int64(1))
				} else {
					left = nt.Temp
				}
			} else {
				ll = _lemit(ll, new_inst(_cmp_op(k), nt.Temp, left, right.Temp, "", int64(0)))
				left = nt.Temp
			}
		} else if (_is_float_type(left_type) || _is_float_type(_lget_type(ll, right.Temp))) {
			ll = _lemit(ll, new_inst(_cmp_op(k), nt.Temp, left, right.Temp, "f64", int64(0)))
			left = nt.Temp
		} else if (left_type == int64(12)) {
			cmp_nt := _lnew_temp(ll)
			_ = cmp_nt
			ll = cmp_nt.L
			ll = _lemit(ll, new_inst(IrOpOpStrCmp{}, cmp_nt.Temp, left, right.Temp, "", int64(0)))
			ll = _lset_type(ll, cmp_nt.Temp, int64(1))
			zero_nt := _lnew_temp(ll)
			_ = zero_nt
			ll = zero_nt.L
			ll = _lemit(ll, new_inst(IrOpOpConst{}, zero_nt.Temp, int64(0), int64(0), "", int64(0)))
			ll = _lset_type(ll, zero_nt.Temp, int64(1))
			ll = _lemit(ll, new_inst(_cmp_op(k), nt.Temp, cmp_nt.Temp, zero_nt.Temp, "", int64(0)))
			left = nt.Temp
		} else {
			ll = _lemit(ll, new_inst(_cmp_op(k), nt.Temp, left, right.Temp, "", int64(0)))
			left = nt.Temp
		}
		ll = _lset_type(ll, nt.Temp, int64(1))
	}
	return LT{L: ll, Temp: left}
}

func _lower_expr(l Lowerer) LT {
	result := _lower_primary(l)
	_ = result
	ll := result.L
	_ = ll
	left := result.Temp
	_ = left
	running := true
	_ = running
	for running {
		k := _lk(ll)
		_ = k
		if ((k == "+") || (k == "-")) {
			ll = _ladv(ll)
			right := _lower_multiplicative(ll)
			_ = right
			ll = right.L
			left_type := _lget_type(ll, left)
			_ = left_type
			right_type := _lget_type(ll, right.Temp)
			_ = right_type
			if ((k == "+") && (left_type == int64(12))) {
				nt := _lnew_str_temp(ll)
				_ = nt
				ll = nt.L
				ll = _lemit(ll, new_inst(IrOpOpStrConcat{}, nt.Temp, left, right.Temp, "", int64(0)))
				ll = _lset_type(ll, nt.Temp, int64(12))
				left = nt.Temp
			} else if (_is_float_type(left_type) || _is_float_type(right_type)) {
				nt := _lnew_temp(ll)
				_ = nt
				ll = nt.L
				ll = _lemit(ll, new_inst(_arith_op(k), nt.Temp, left, right.Temp, "f64", int64(0)))
				ll = _lset_type(ll, nt.Temp, int64(9))
				left = nt.Temp
			} else {
				nt := _lnew_temp(ll)
				_ = nt
				ll = nt.L
				ll = _lemit(ll, new_inst(_arith_op(k), nt.Temp, left, right.Temp, "", int64(0)))
				ll = _lset_type(ll, nt.Temp, int64(1))
				left = nt.Temp
			}
		} else if (((k == "*") || (k == "/")) || (k == "%")) {
			ll = _ladv(ll)
			right := _lower_atom(ll)
			_ = right
			ll = right.L
			nt := _lnew_temp(ll)
			_ = nt
			ll = nt.L
			left_type := _lget_type(ll, left)
			_ = left_type
			right_type := _lget_type(ll, right.Temp)
			_ = right_type
			if (_is_float_type(left_type) || _is_float_type(right_type)) {
				ll = _lemit(ll, new_inst(_arith_op(k), nt.Temp, left, right.Temp, "f64", int64(0)))
				ll = _lset_type(ll, nt.Temp, int64(9))
			} else {
				ll = _lemit(ll, new_inst(_arith_op(k), nt.Temp, left, right.Temp, "", int64(0)))
				ll = _lset_type(ll, nt.Temp, int64(1))
			}
			left = nt.Temp
		} else if ((((((k == "==") || (k == "!=")) || (k == "<")) || (k == ">")) || (k == "<=")) || (k == ">=")) {
			ll = _ladv(ll)
			right := _lower_additive(ll)
			_ = right
			ll = right.L
			nt := _lnew_temp(ll)
			_ = nt
			ll = nt.L
			left_type := _lget_type(ll, left)
			_ = left_type
			if ((((k == "==") || (k == "!="))) && (left_type == int64(12))) {
				ll = _lemit(ll, new_inst(IrOpOpStrEq{}, nt.Temp, left, right.Temp, "", int64(0)))
				if (k == "!=") {
					not_nt := _lnew_temp(ll)
					_ = not_nt
					ll = not_nt.L
					ll = _lemit(ll, new_inst(IrOpOpNot{}, not_nt.Temp, nt.Temp, int64(0), "", int64(0)))
					left = not_nt.Temp
					ll = _lset_type(ll, not_nt.Temp, int64(1))
				} else {
					left = nt.Temp
				}
			} else if ((((k == "==") || (k == "!="))) && _is_struct_type(ll.Store, ll.Registry, left_type)) {
				sname := _struct_name_for_type(ll.Store, ll.Registry, left_type)
				_ = sname
				if (sname != "") {
					eq_result := _lower_struct_eq(ll, left, right.Temp, sname, nt.Temp)
					_ = eq_result
					ll = eq_result
					if (k == "!=") {
						not_nt := _lnew_temp(ll)
						_ = not_nt
						ll = not_nt.L
						ll = _lemit(ll, new_inst(IrOpOpNot{}, not_nt.Temp, nt.Temp, int64(0), "", int64(0)))
						left = not_nt.Temp
						ll = _lset_type(ll, not_nt.Temp, int64(1))
					} else {
						left = nt.Temp
					}
				} else {
					ll = _lemit(ll, new_inst(_cmp_op(k), nt.Temp, left, right.Temp, "", int64(0)))
					left = nt.Temp
				}
			} else if (_is_float_type(left_type) || _is_float_type(_lget_type(ll, right.Temp))) {
				ll = _lemit(ll, new_inst(_cmp_op(k), nt.Temp, left, right.Temp, "f64", int64(0)))
				left = nt.Temp
			} else if (left_type == int64(12)) {
				cmp_nt := _lnew_temp(ll)
				_ = cmp_nt
				ll = cmp_nt.L
				ll = _lemit(ll, new_inst(IrOpOpStrCmp{}, cmp_nt.Temp, left, right.Temp, "", int64(0)))
				ll = _lset_type(ll, cmp_nt.Temp, int64(1))
				zero_nt := _lnew_temp(ll)
				_ = zero_nt
				ll = zero_nt.L
				ll = _lemit(ll, new_inst(IrOpOpConst{}, zero_nt.Temp, int64(0), int64(0), "", int64(0)))
				ll = _lset_type(ll, zero_nt.Temp, int64(1))
				ll = _lemit(ll, new_inst(_cmp_op(k), nt.Temp, cmp_nt.Temp, zero_nt.Temp, "", int64(0)))
				left = nt.Temp
			} else {
				ll = _lemit(ll, new_inst(_cmp_op(k), nt.Temp, left, right.Temp, "", int64(0)))
				left = nt.Temp
			}
			ll = _lset_type(ll, nt.Temp, int64(1))
		} else if (k == "&&") {
			ll = _ladv(ll)
			right := _lower_comparison(ll)
			_ = right
			ll = right.L
			nt := _lnew_temp(ll)
			_ = nt
			ll = nt.L
			ll = _lemit(ll, new_inst(IrOpOpAnd{}, nt.Temp, left, right.Temp, "", int64(0)))
			left = nt.Temp
		} else if (k == "||") {
			ll = _ladv(ll)
			right := _lower_comparison(ll)
			_ = right
			ll = right.L
			nt := _lnew_temp(ll)
			_ = nt
			ll = nt.L
			ll = _lemit(ll, new_inst(IrOpOpOr{}, nt.Temp, left, right.Temp, "", int64(0)))
			left = nt.Temp
		} else if (k == "|>") {
			ll = _ladv(ll)
			if (_lk(ll) == "IDENT") {
				fname := _lcur(ll).Text
				_ = fname
				ll = _ladv(ll)
				pipe_sig := reg_find_fn(ll.Registry, fname)
				_ = pipe_sig
				pipe_call_name := fname
				_ = pipe_call_name
				pipe_call_sig := pipe_sig
				_ = pipe_call_sig
				if ((pipe_sig.Name != "") && fn_is_generic(pipe_sig)) {
					left_type := _lget_type(ll, left)
					_ = left_type
					suffix := _type_name_suffix(ll.Store, left_type)
					_ = suffix
					if (suffix != "unknown") {
						pipe_call_name = ((fname + "_") + suffix)
						sp_sig := reg_find_fn(ll.Registry, pipe_call_name)
						_ = sp_sig
						if (sp_sig.Name != "") {
							pipe_call_sig = sp_sig
						}
					}
				}
				nt_temp := int64(0)
				_ = nt_temp
				if ((pipe_call_sig.Name != "") && (pipe_call_sig.Return_type == int64(12))) {
					nt := _lnew_str_temp(ll)
					_ = nt
					ll = nt.L
					nt_temp = nt.Temp
				} else {
					nt := _lnew_temp(ll)
					_ = nt
					ll = nt.L
					nt_temp = nt.Temp
				}
				ll = _lemit(ll, new_inst(IrOpOpCall{}, nt_temp, left, int64(1), pipe_call_name, int64(0)))
				if ((pipe_call_sig.Name != "") && (pipe_call_sig.Return_type == int64(12))) {
					ll = _lset_type(ll, nt_temp, int64(12))
				} else if ((pipe_call_sig.Name != "") && (pipe_call_sig.Return_type > int64(0))) {
					ll = _lset_type(ll, nt_temp, pipe_call_sig.Return_type)
				}
				left = nt_temp
			}
		} else if (k == ".") {
			ll = _ladv(ll)
			if (_lk(ll) == "IDENT") {
				field := _lcur(ll).Text
				_ = field
				ll = _ladv(ll)
				if (_lk(ll) == "(") {
					mr := _lower_method_call(ll, left, field)
					_ = mr
					ll = mr.L
					left = mr.Temp
				} else if ((_lk(ll) == "[") && (((field == "to") || (field == "trunc")))) {
					ll = _ladv(ll)
					target_type_name := ""
					_ = target_type_name
					if (_lk(ll) == "IDENT") {
						target_type_name = _lcur(ll).Text
						ll = _ladv(ll)
					}
					if (_lk(ll) == "]") {
						ll = _ladv(ll)
					}
					if (_lk(ll) == "(") {
						ll = _ladv(ll)
					}
					if (_lk(ll) == ")") {
						ll = _ladv(ll)
					}
					mr := _lower_conversion(ll, left, field, target_type_name)
					_ = mr
					ll = mr.L
					left = mr.Temp
				} else {
					left_type := _lget_type(ll, left)
					_ = left_type
					sname := _struct_name_for_type(ll.Store, ll.Registry, left_type)
					_ = sname
					fidx := int64(0)
					_ = fidx
					ft := int64(0)
					_ = ft
					if (sname != "") {
						fidx = _field_index(ll.Registry, sname, field)
						ft = _field_type_id(ll, sname, field)
					} else {
						fidx = _field_index_any(ll.Registry, field)
						ft = _field_type_id_any(ll.Registry, field)
					}
					has_fld := reg_find_field(ll.Registry, sname, field)
					_ = has_fld
					if ((has_fld.Name == "") && (sname != "")) {
						qualified := ((sname + "_") + field)
						_ = qualified
						msig := reg_find_fn(ll.Registry, qualified)
						_ = msig
						if (msig.Name != "") {
							ll = _lset_reg(ll, reg_add_bound_fn_value_ref(ll.Registry, qualified))
							env_nt := _lnew_temp(ll)
							_ = env_nt
							ll = env_nt.L
							ll = _lemit(ll, new_inst(IrOpOpAlloc{}, env_nt.Temp, int64(1), int64(0), "_bound_env", int64(0)))
							ll = _lemit(ll, new_inst(IrOpOpFieldSet{}, env_nt.Temp, left, int64(0), "_self", int64(0)))
							clos := _lnew_temp(ll)
							_ = clos
							ll = clos.L
							ll = _lemit(ll, new_inst(IrOpOpAlloc{}, clos.Temp, int64(2), int64(0), "_Closure", int64(0)))
							fn_ref := _lnew_temp(ll)
							_ = fn_ref
							ll = fn_ref.L
							ll = _lemit(ll, new_inst(IrOpOpFnRef{}, fn_ref.Temp, int64(0), int64(0), ("_bound_tramp_" + qualified), int64(0)))
							ll = _lemit(ll, new_inst(IrOpOpFieldSet{}, clos.Temp, fn_ref.Temp, int64(0), "_fn", int64(0)))
							ll = _lemit(ll, new_inst(IrOpOpFieldSet{}, clos.Temp, env_nt.Temp, int64(1), "_env", int64(0)))
							left = clos.Temp
						} else {
							ll = _lemit(ll, new_inst(IrOpOpFieldGet{}, _lnew_temp(ll).Temp, left, fidx, field, int64(0)))
							left = (_lnew_temp(ll).Temp - int64(1))
						}
					} else {
						nt_temp := int64(0)
						_ = nt_temp
						if (ft == int64(12)) {
							nt := _lnew_str_temp(ll)
							_ = nt
							ll = nt.L
							nt_temp = nt.Temp
							fidx_len := (fidx + int64(1))
							_ = fidx_len
							nt_len := (nt_temp + int64(1))
							_ = nt_len
							ll = _lemit(ll, new_inst(IrOpOpFieldGet{}, nt_temp, left, fidx, field, int64(0)))
							ll = _lemit(ll, new_inst(IrOpOpFieldGet{}, nt_len, left, fidx_len, field, int64(0)))
						} else {
							nt := _lnew_temp(ll)
							_ = nt
							ll = nt.L
							nt_temp = nt.Temp
							ll = _lemit(ll, new_inst(IrOpOpFieldGet{}, nt_temp, left, fidx, field, int64(0)))
						}
						if (ft > int64(0)) {
							ll = _lset_type(ll, nt_temp, ft)
						}
						left = nt_temp
					}
				}
			} else if (_lk(ll) == "{") {
				rb2 := "}"
				_ = rb2
				ll = _ladv(ll)
				ll = _lskip_nl(ll)
				left_type := _lget_type(ll, left)
				_ = left_type
				sname := _struct_name_for_type(ll.Store, ll.Registry, left_type)
				_ = sname
				new_nt := _lnew_temp(ll)
				_ = new_nt
				ll = new_nt.L
				if (sname != "") {
					sdef := reg_find_struct(ll.Registry, sname)
					_ = sdef
					alloc_sz := int64(len(sdef.Fields))
					_ = alloc_sz
					ll = _lemit(ll, new_inst(IrOpOpAlloc{}, new_nt.Temp, alloc_sz, int64(0), sname, int64(0)))
					cfi := int64(1)
					_ = cfi
					for (cfi < int64(len(sdef.Fields))) {
						copy_nt := _lnew_temp(ll)
						_ = copy_nt
						ll = copy_nt.L
						fidx2 := (cfi - int64(1))
						_ = fidx2
						ll = _lemit(ll, new_inst(IrOpOpFieldGet{}, copy_nt.Temp, left, fidx2, sdef.Fields[cfi].Name, int64(0)))
						ll = _lemit(ll, new_inst(IrOpOpFieldSet{}, new_nt.Temp, copy_nt.Temp, fidx2, sdef.Fields[cfi].Name, int64(0)))
						cfi = (cfi + int64(1))
					}
				} else {
					ll = _lemit(ll, new_inst(IrOpOpAlloc{}, new_nt.Temp, int64(8), int64(0), "_RecordUpdate", int64(0)))
					ll = _lemit(ll, new_inst(IrOpOpStore{}, new_nt.Temp, left, int64(0), "", int64(0)))
				}
				ll = _lset_type(ll, new_nt.Temp, left_type)
				for ((_lk(ll) != rb2) && (_lk(ll) != "EOF")) {
					if (_lk(ll) == "IDENT") {
						upd_field := _lcur(ll).Text
						_ = upd_field
						ll = _ladv(ll)
						if (_lk(ll) == ":") {
							ll = _ladv(ll)
						}
						upd_val := _lower_expr(ll)
						_ = upd_val
						ll = upd_val.L
						if (sname != "") {
							upd_fidx := _field_index(ll.Registry, sname, upd_field)
							_ = upd_fidx
							ll = _lemit(ll, new_inst(IrOpOpFieldSet{}, new_nt.Temp, upd_val.Temp, upd_fidx, upd_field, int64(0)))
						}
					} else {
						ll = _ladv(ll)
					}
					ll = _lskip_nl(ll)
					if (_lk(ll) == ",") {
						ll = _ladv(ll)
						ll = _lskip_nl(ll)
					}
				}
				if (_lk(ll) == rb2) {
					ll = _ladv(ll)
				}
				left = new_nt.Temp
			} else if (_lk(ll) == "INT") {
				field_idx_str := _lcur(ll).Text
				_ = field_idx_str
				ll = _ladv(ll)
				fidx := int64(0)
				_ = fidx
				if (field_idx_str == "1") {
					fidx = int64(1)
				}
				if (field_idx_str == "2") {
					fidx = int64(2)
				}
				if (field_idx_str == "3") {
					fidx = int64(3)
				}
				if (field_idx_str == "4") {
					fidx = int64(4)
				}
				if (field_idx_str == "5") {
					fidx = int64(5)
				}
				if (field_idx_str == "6") {
					fidx = int64(6)
				}
				if (field_idx_str == "7") {
					fidx = int64(7)
				}
				nt := _lnew_temp(ll)
				_ = nt
				ll = nt.L
				ll = _lemit(ll, new_inst(IrOpOpFieldGet{}, nt.Temp, left, fidx, field_idx_str, int64(0)))
				left = nt.Temp
			} else {
				running = false
			}
		} else if (k == "(") {
			cr := _lower_call_args(ll, left)
			_ = cr
			ll = cr.L
			left = cr.Temp
		} else if (k == "catch") {
			ll = _ladv(ll)
			err_name := "_"
			_ = err_name
			if (_lk(ll) == "|") {
				ll = _ladv(ll)
				if (_lk(ll) == "IDENT") {
					err_name = _lcur(ll).Text
					ll = _ladv(ll)
				} else {
					ll = _ladv(ll)
				}
				if (_lk(ll) == "|") {
					ll = _ladv(ll)
				}
			}
			tag_nt := _lnew_temp(ll)
			_ = tag_nt
			ll = tag_nt.L
			ll = _lemit(ll, new_inst(IrOpOpFieldGet{}, tag_nt.Temp, left, int64(0), "_tag", int64(0)))
			const_one := _lnew_temp(ll)
			_ = const_one
			ll = const_one.L
			ll = _lemit(ll, new_inst(IrOpOpConst{}, const_one.Temp, int64(1), int64(0), "", int64(0)))
			is_err := _lnew_temp(ll)
			_ = is_err
			ll = is_err.L
			ll = _lemit(ll, new_inst(IrOpOpEq{}, is_err.Temp, tag_nt.Temp, const_one.Temp, "", int64(0)))
			catch_lbl := _lnew_label(ll)
			_ = catch_lbl
			ll = catch_lbl.L
			end_lbl := _lnew_label(ll)
			_ = end_lbl
			ll = end_lbl.L
			ll = _lemit(ll, new_inst(IrOpOpBranchTrue{}, int64(0), is_err.Temp, catch_lbl.Temp, "", int64(0)))
			ok_val := _lnew_temp(ll)
			_ = ok_val
			ll = ok_val.L
			ll = _lemit(ll, new_inst(IrOpOpFieldGet{}, ok_val.Temp, left, int64(1), "_val", int64(0)))
			final_nt := _lnew_temp(ll)
			_ = final_nt
			ll = final_nt.L
			ll = _lemit(ll, new_inst(IrOpOpStore{}, final_nt.Temp, ok_val.Temp, int64(0), "", int64(0)))
			ll = _lemit(ll, new_inst(IrOpOpJump{}, int64(0), end_lbl.Temp, int64(0), "", int64(0)))
			ll = _lemit(ll, new_inst(IrOpOpLabel{}, int64(0), catch_lbl.Temp, int64(0), "", int64(0)))
			err_val := _lnew_temp(ll)
			_ = err_val
			ll = err_val.L
			ll = _lemit(ll, new_inst(IrOpOpFieldGet{}, err_val.Temp, left, int64(1), "_val", int64(0)))
			if (err_name != "_") {
				ll = _lenv_add(ll, err_name, err_val.Temp)
			}
			lb := "{"
			_ = lb
			rb := "}"
			_ = rb
			if (_lk(ll) == lb) {
				ll = _ladv(ll)
				ll = _lskip_nl(ll)
				is_typed_catch := false
				_ = is_typed_catch
				peek_p := ll.Pos
				_ = peek_p
				for (peek_p < int64(len(ll.Tokens))) {
					pk := token_name(ll.Tokens[peek_p].Kind)
					_ = pk
					if (pk == "=>") {
						is_typed_catch = true
					}
					if (((pk == "NEWLINE") || (pk == rb)) || (pk == "EOF")) {
						peek_p = int64(len(ll.Tokens))
					}
					peek_p = (peek_p + int64(1))
				}
				if is_typed_catch {
					catch_end := _lnew_label(ll)
					_ = catch_end
					ll = catch_end.L
					err_tag_temp := err_val.Temp
					_ = err_tag_temp
					is_err_sum := false
					_ = is_err_sum
					err_parent := _variant_parent(ll.Registry, "")
					_ = err_parent
					err_tag_ext := _lnew_temp(ll)
					_ = err_tag_ext
					ll = err_tag_ext.L
					ll = _lemit(ll, new_inst(IrOpOpFieldGet{}, err_tag_ext.Temp, err_val.Temp, int64(0), "_tag", int64(0)))
					err_tag_temp = err_tag_ext.Temp
					for ((_lk(ll) != rb) && (_lk(ll) != "EOF")) {
						ll = _lskip_nl(ll)
						if (_lk(ll) == rb) {
							break
						}
						pat := _lcur(ll).Text
						_ = pat
						pat_kind := _lk(ll)
						_ = pat_kind
						ll = _ladv(ll)
						if (_lk(ll) == "(") {
							for ((_lk(ll) != ")") && (_lk(ll) != "EOF")) {
								ll = _ladv(ll)
							}
							if (_lk(ll) == ")") {
								ll = _ladv(ll)
							}
						}
						if (_lk(ll) == "{") {
							for ((_lk(ll) != "}") && (_lk(ll) != "EOF")) {
								ll = _ladv(ll)
							}
							if (_lk(ll) == "}") {
								ll = _ladv(ll)
							}
						}
						if (_lk(ll) == "=>") {
							ll = _ladv(ll)
						}
						ll = _lskip_nl(ll)
						if ((pat == "_") || (((pat_kind == "IDENT") && (_variant_tag(ll.Registry, pat) == int64(0))))) {
							if (pat != "_") {
								ll = _lenv_add(ll, pat, err_val.Temp)
							}
							if (_lk(ll) == "{") {
								ll = _lower_block(ll)
							} else {
								cr := _lower_expr(ll)
								_ = cr
								ll = cr.L
								ll = _lemit(ll, new_inst(IrOpOpStore{}, final_nt.Temp, cr.Temp, int64(0), "", int64(0)))
							}
							ll = _lemit(ll, new_inst(IrOpOpJump{}, int64(0), catch_end.Temp, int64(0), "", int64(0)))
						} else {
							vtag := _variant_tag(ll.Registry, pat)
							_ = vtag
							next_arm := _lnew_label(ll)
							_ = next_arm
							ll = next_arm.L
							vtag_nt := _lnew_temp(ll)
							_ = vtag_nt
							ll = vtag_nt.L
							ll = _lemit(ll, new_inst(IrOpOpConst{}, vtag_nt.Temp, vtag, int64(0), "", int64(0)))
							vcmp := _lnew_temp(ll)
							_ = vcmp
							ll = vcmp.L
							ll = _lemit(ll, new_inst(IrOpOpEq{}, vcmp.Temp, err_tag_temp, vtag_nt.Temp, "", int64(0)))
							ll = _lemit(ll, new_inst(IrOpOpBranchFalse{}, int64(0), vcmp.Temp, next_arm.Temp, "", int64(0)))
							if (_lk(ll) == "{") {
								ll = _lower_block(ll)
							} else {
								cr := _lower_expr(ll)
								_ = cr
								ll = cr.L
								ll = _lemit(ll, new_inst(IrOpOpStore{}, final_nt.Temp, cr.Temp, int64(0), "", int64(0)))
							}
							ll = _lemit(ll, new_inst(IrOpOpJump{}, int64(0), catch_end.Temp, int64(0), "", int64(0)))
							ll = _lemit(ll, new_inst(IrOpOpLabel{}, int64(0), next_arm.Temp, int64(0), "", int64(0)))
						}
						ll = _lskip_nl(ll)
					}
					if (_lk(ll) == rb) {
						ll = _ladv(ll)
					}
					ll = _lemit(ll, new_inst(IrOpOpLabel{}, int64(0), catch_end.Temp, int64(0), "", int64(0)))
				} else {
					catch_last := -int64(1)
					_ = catch_last
					for ((_lk(ll) != rb) && (_lk(ll) != "EOF")) {
						ck := _lk(ll)
						_ = ck
						if (((((ck == "return") || (ck == "mut")) || (ck == "break")) || (ck == "continue")) || (ck == "assert")) {
							ll = _lower_stmt(ll)
						} else {
							cr := _lower_expr_result(ll)
							_ = cr
							ll = cr.L
							catch_last = cr.Temp
						}
						ll = _lskip_nl(ll)
					}
					if (_lk(ll) == rb) {
						ll = _ladv(ll)
					}
					if (catch_last >= int64(0)) {
						ll = _lemit(ll, new_inst(IrOpOpStore{}, final_nt.Temp, catch_last, int64(0), "", int64(0)))
					}
				}
			}
			ll = _lemit(ll, new_inst(IrOpOpLabel{}, int64(0), end_lbl.Temp, int64(0), "", int64(0)))
			left = final_nt.Temp
		} else if ((k == "IDENT") && (_lcur(ll).Text == "or")) {
			ll = _ladv(ll)
			default_val := _lower_expr(ll)
			_ = default_val
			ll = default_val.L
			tag_nt := _lnew_temp(ll)
			_ = tag_nt
			ll = tag_nt.L
			ll = _lemit(ll, new_inst(IrOpOpFieldGet{}, tag_nt.Temp, left, int64(0), "_tag", int64(0)))
			const_one := _lnew_temp(ll)
			_ = const_one
			ll = const_one.L
			ll = _lemit(ll, new_inst(IrOpOpConst{}, const_one.Temp, int64(1), int64(0), "", int64(0)))
			is_err := _lnew_temp(ll)
			_ = is_err
			ll = is_err.L
			ll = _lemit(ll, new_inst(IrOpOpEq{}, is_err.Temp, tag_nt.Temp, const_one.Temp, "", int64(0)))
			or_err_lbl := _lnew_label(ll)
			_ = or_err_lbl
			ll = or_err_lbl.L
			or_end_lbl := _lnew_label(ll)
			_ = or_end_lbl
			ll = or_end_lbl.L
			final_nt := _lnew_temp(ll)
			_ = final_nt
			ll = final_nt.L
			ll = _lemit(ll, new_inst(IrOpOpBranchTrue{}, int64(0), is_err.Temp, or_err_lbl.Temp, "", int64(0)))
			ok_val := _lnew_temp(ll)
			_ = ok_val
			ll = ok_val.L
			ll = _lemit(ll, new_inst(IrOpOpFieldGet{}, ok_val.Temp, left, int64(1), "_val", int64(0)))
			ll = _lemit(ll, new_inst(IrOpOpStore{}, final_nt.Temp, ok_val.Temp, int64(0), "", int64(0)))
			ll = _lemit(ll, new_inst(IrOpOpJump{}, int64(0), or_end_lbl.Temp, int64(0), "", int64(0)))
			ll = _lemit(ll, new_inst(IrOpOpLabel{}, int64(0), or_err_lbl.Temp, int64(0), "", int64(0)))
			ll = _lemit(ll, new_inst(IrOpOpStore{}, final_nt.Temp, default_val.Temp, int64(0), "", int64(0)))
			ll = _lemit(ll, new_inst(IrOpOpLabel{}, int64(0), or_end_lbl.Temp, int64(0), "", int64(0)))
			left = final_nt.Temp
		} else if ((k == "IDENT") && (_lcur(ll).Text == "must")) {
			ll = _ladv(ll)
			msg := _lower_expr(ll)
			_ = msg
			ll = msg.L
			tag_nt := _lnew_temp(ll)
			_ = tag_nt
			ll = tag_nt.L
			ll = _lemit(ll, new_inst(IrOpOpFieldGet{}, tag_nt.Temp, left, int64(0), "_tag", int64(0)))
			const_one := _lnew_temp(ll)
			_ = const_one
			ll = const_one.L
			ll = _lemit(ll, new_inst(IrOpOpConst{}, const_one.Temp, int64(1), int64(0), "", int64(0)))
			is_err := _lnew_temp(ll)
			_ = is_err
			ll = is_err.L
			ll = _lemit(ll, new_inst(IrOpOpEq{}, is_err.Temp, tag_nt.Temp, const_one.Temp, "", int64(0)))
			panic_lbl := _lnew_label(ll)
			_ = panic_lbl
			ll = panic_lbl.L
			ok_lbl := _lnew_label(ll)
			_ = ok_lbl
			ll = ok_lbl.L
			ll = _lemit(ll, new_inst(IrOpOpBranchTrue{}, int64(0), is_err.Temp, panic_lbl.Temp, "", int64(0)))
			ok_val := _lnew_temp(ll)
			_ = ok_val
			ll = ok_val.L
			ll = _lemit(ll, new_inst(IrOpOpFieldGet{}, ok_val.Temp, left, int64(1), "_val", int64(0)))
			ll = _lemit(ll, new_inst(IrOpOpJump{}, int64(0), ok_lbl.Temp, int64(0), "", int64(0)))
			ll = _lemit(ll, new_inst(IrOpOpLabel{}, int64(0), panic_lbl.Temp, int64(0), "", int64(0)))
			ll = _lemit(ll, new_inst(IrOpOpCall{}, int64(0), msg.Temp, int64(2), "_aria_panic", int64(0)))
			ll = _lemit(ll, new_inst(IrOpOpLabel{}, int64(0), ok_lbl.Temp, int64(0), "", int64(0)))
			left = ok_val.Temp
		} else if (k == "??") {
			ll = _ladv(ll)
			default_val := _lower_additive(ll)
			_ = default_val
			ll = default_val.L
			tag_nt := _lnew_temp(ll)
			_ = tag_nt
			ll = tag_nt.L
			ll = _lemit(ll, new_inst(IrOpOpFieldGet{}, tag_nt.Temp, left, int64(0), "_tag", int64(0)))
			some_lbl := _lnew_label(ll)
			_ = some_lbl
			ll = some_lbl.L
			end_lbl := _lnew_label(ll)
			_ = end_lbl
			ll = end_lbl.L
			final_nt := _lnew_temp(ll)
			_ = final_nt
			ll = final_nt.L
			ll = _lemit(ll, new_inst(IrOpOpBranchTrue{}, int64(0), tag_nt.Temp, some_lbl.Temp, "", int64(0)))
			ll = _lemit(ll, new_inst(IrOpOpStore{}, final_nt.Temp, default_val.Temp, int64(0), "", int64(0)))
			ll = _lemit(ll, new_inst(IrOpOpJump{}, int64(0), end_lbl.Temp, int64(0), "", int64(0)))
			ll = _lemit(ll, new_inst(IrOpOpLabel{}, int64(0), some_lbl.Temp, int64(0), "", int64(0)))
			some_val := _lnew_temp(ll)
			_ = some_val
			ll = some_val.L
			ll = _lemit(ll, new_inst(IrOpOpFieldGet{}, some_val.Temp, left, int64(1), "_val", int64(0)))
			ll = _lemit(ll, new_inst(IrOpOpStore{}, final_nt.Temp, some_val.Temp, int64(0), "", int64(0)))
			ll = _lemit(ll, new_inst(IrOpOpLabel{}, int64(0), end_lbl.Temp, int64(0), "", int64(0)))
			left = final_nt.Temp
		} else if (k == "?.") {
			ll = _ladv(ll)
			if (_lk(ll) == "IDENT") {
				field := _lcur(ll).Text
				_ = field
				ll = _ladv(ll)
				tag_nt := _lnew_temp(ll)
				_ = tag_nt
				ll = tag_nt.L
				ll = _lemit(ll, new_inst(IrOpOpFieldGet{}, tag_nt.Temp, left, int64(0), "_tag", int64(0)))
				some_lbl := _lnew_label(ll)
				_ = some_lbl
				ll = some_lbl.L
				end_lbl := _lnew_label(ll)
				_ = end_lbl
				ll = end_lbl.L
				final_nt := _lnew_temp(ll)
				_ = final_nt
				ll = final_nt.L
				ll = _lemit(ll, new_inst(IrOpOpBranchTrue{}, int64(0), tag_nt.Temp, some_lbl.Temp, "", int64(0)))
				none_nt := _lnew_temp(ll)
				_ = none_nt
				ll = none_nt.L
				ll = _lemit(ll, new_inst(IrOpOpAlloc{}, none_nt.Temp, int64(1), int64(0), "_Optional", int64(0)))
				none_tag := _lnew_temp(ll)
				_ = none_tag
				ll = none_tag.L
				ll = _lemit(ll, new_inst(IrOpOpConst{}, none_tag.Temp, int64(0), int64(0), "", int64(0)))
				ll = _lemit(ll, new_inst(IrOpOpFieldSet{}, none_nt.Temp, none_tag.Temp, int64(0), "_tag", int64(0)))
				ll = _lemit(ll, new_inst(IrOpOpStore{}, final_nt.Temp, none_nt.Temp, int64(0), "", int64(0)))
				ll = _lemit(ll, new_inst(IrOpOpJump{}, int64(0), end_lbl.Temp, int64(0), "", int64(0)))
				ll = _lemit(ll, new_inst(IrOpOpLabel{}, int64(0), some_lbl.Temp, int64(0), "", int64(0)))
				inner := _lnew_temp(ll)
				_ = inner
				ll = inner.L
				ll = _lemit(ll, new_inst(IrOpOpFieldGet{}, inner.Temp, left, int64(1), "_val", int64(0)))
				opt_type := _lget_type(ll, left)
				_ = opt_type
				opt_info := store_get(ll.Store, opt_type)
				_ = opt_info
				inner_sname := _struct_name_for_type(ll.Store, ll.Registry, type_info_type_id(opt_info))
				_ = inner_sname
				fidx := int64(0)
				_ = fidx
				if (inner_sname != "") {
					fidx = _field_index(ll.Registry, inner_sname, field)
				}
				field_val := _lnew_temp(ll)
				_ = field_val
				ll = field_val.L
				ll = _lemit(ll, new_inst(IrOpOpFieldGet{}, field_val.Temp, inner.Temp, fidx, field, int64(0)))
				some_result := _lnew_temp(ll)
				_ = some_result
				ll = some_result.L
				ll = _lemit(ll, new_inst(IrOpOpAlloc{}, some_result.Temp, int64(2), int64(0), "_Optional", int64(0)))
				some_tag := _lnew_temp(ll)
				_ = some_tag
				ll = some_tag.L
				ll = _lemit(ll, new_inst(IrOpOpConst{}, some_tag.Temp, int64(1), int64(0), "", int64(0)))
				ll = _lemit(ll, new_inst(IrOpOpFieldSet{}, some_result.Temp, some_tag.Temp, int64(0), "_tag", int64(0)))
				ll = _lemit(ll, new_inst(IrOpOpFieldSet{}, some_result.Temp, field_val.Temp, int64(1), "_val", int64(0)))
				ll = _lemit(ll, new_inst(IrOpOpStore{}, final_nt.Temp, some_result.Temp, int64(0), "", int64(0)))
				ll = _lemit(ll, new_inst(IrOpOpLabel{}, int64(0), end_lbl.Temp, int64(0), "", int64(0)))
				ll = _lset_type(ll, final_nt.Temp, int64(201))
				left = final_nt.Temp
			}
		} else if (k == "[") {
			ll = _ladv(ll)
			idx := _lower_expr(ll)
			_ = idx
			ll = idx.L
			if (_lk(ll) == "]") {
				ll = _ladv(ll)
			}
			left_type := _lget_type(ll, left)
			_ = left_type
			elem_type := _array_elem_type(ll.Store, left_type)
			_ = elem_type
			if (elem_type == int64(12)) {
				nt := _lnew_str_temp(ll)
				_ = nt
				ll = nt.L
				ll = _lemit(ll, new_inst(IrOpOpArrayGet{}, nt.Temp, left, idx.Temp, "", int64(12)))
				ll = _lset_type(ll, nt.Temp, int64(12))
				left = nt.Temp
			} else {
				nt := _lnew_temp(ll)
				_ = nt
				ll = nt.L
				ll = _lemit(ll, new_inst(IrOpOpArrayGet{}, nt.Temp, left, idx.Temp, "", int64(0)))
				if (elem_type > int64(0)) {
					ll = _lset_type(ll, nt.Temp, elem_type)
				}
				left = nt.Temp
			}
		} else if (k == "?") {
			ll = _ladv(ll)
			tag_nt := _lnew_temp(ll)
			_ = tag_nt
			ll = tag_nt.L
			ll = _lemit(ll, new_inst(IrOpOpFieldGet{}, tag_nt.Temp, left, int64(0), "_tag", int64(0)))
			const_one := _lnew_temp(ll)
			_ = const_one
			ll = const_one.L
			ll = _lemit(ll, new_inst(IrOpOpConst{}, const_one.Temp, int64(1), int64(0), "", int64(0)))
			is_err := _lnew_temp(ll)
			_ = is_err
			ll = is_err.L
			ll = _lemit(ll, new_inst(IrOpOpEq{}, is_err.Temp, tag_nt.Temp, const_one.Temp, "", int64(0)))
			err_lbl := _lnew_label(ll)
			_ = err_lbl
			ll = err_lbl.L
			ok_lbl := _lnew_label(ll)
			_ = ok_lbl
			ll = ok_lbl.L
			ll = _lemit(ll, new_inst(IrOpOpBranchTrue{}, int64(0), is_err.Temp, err_lbl.Temp, "", int64(0)))
			ok_val := _lnew_temp(ll)
			_ = ok_val
			ll = ok_val.L
			ll = _lemit(ll, new_inst(IrOpOpFieldGet{}, ok_val.Temp, left, int64(1), "_val", int64(0)))
			ll = _lemit(ll, new_inst(IrOpOpJump{}, int64(0), ok_lbl.Temp, int64(0), "", int64(0)))
			ll = _lemit(ll, new_inst(IrOpOpLabel{}, int64(0), err_lbl.Temp, int64(0), "", int64(0)))
			if (_lk(ll) == "|") {
				ll = _ladv(ll)
				transform_name := "_"
				_ = transform_name
				if (_lk(ll) == "IDENT") {
					transform_name = _lcur(ll).Text
					ll = _ladv(ll)
				} else {
					ll = _ladv(ll)
				}
				if (_lk(ll) == "|") {
					ll = _ladv(ll)
				}
				err_val := _lnew_temp(ll)
				_ = err_val
				ll = err_val.L
				ll = _lemit(ll, new_inst(IrOpOpFieldGet{}, err_val.Temp, left, int64(1), "_val", int64(0)))
				if (transform_name != "_") {
					ll = _lenv_add(ll, transform_name, err_val.Temp)
				}
				transform := _lower_expr(ll)
				_ = transform
				ll = transform.L
				new_err := _lnew_temp(ll)
				_ = new_err
				ll = new_err.L
				ll = _lemit(ll, new_inst(IrOpOpAlloc{}, new_err.Temp, int64(2), int64(0), "_Result", int64(0)))
				ll = _lemit(ll, new_inst(IrOpOpFieldSet{}, new_err.Temp, const_one.Temp, int64(0), "_tag", int64(0)))
				ll = _lemit(ll, new_inst(IrOpOpFieldSet{}, new_err.Temp, transform.Temp, int64(1), "_val", int64(0)))
				ll = _emit_defers(ll)
				ll = _lemit(ll, new_inst(IrOpOpRet{}, int64(0), new_err.Temp, int64(0), "", int64(0)))
			} else {
				ll = _emit_defers(ll)
				ll = _lemit(ll, new_inst(IrOpOpRet{}, int64(0), left, int64(0), "", int64(0)))
			}
			ll = _lemit(ll, new_inst(IrOpOpLabel{}, int64(0), ok_lbl.Temp, int64(0), "", int64(0)))
			left = ok_val.Temp
		} else if (k == "!") {
			ll = _ladv(ll)
			tag_nt := _lnew_temp(ll)
			_ = tag_nt
			ll = tag_nt.L
			ll = _lemit(ll, new_inst(IrOpOpFieldGet{}, tag_nt.Temp, left, int64(0), "_tag", int64(0)))
			const_one := _lnew_temp(ll)
			_ = const_one
			ll = const_one.L
			ll = _lemit(ll, new_inst(IrOpOpConst{}, const_one.Temp, int64(1), int64(0), "", int64(0)))
			is_err := _lnew_temp(ll)
			_ = is_err
			ll = is_err.L
			ll = _lemit(ll, new_inst(IrOpOpEq{}, is_err.Temp, tag_nt.Temp, const_one.Temp, "", int64(0)))
			panic_lbl := _lnew_label(ll)
			_ = panic_lbl
			ll = panic_lbl.L
			ok_lbl := _lnew_label(ll)
			_ = ok_lbl
			ll = ok_lbl.L
			ll = _lemit(ll, new_inst(IrOpOpBranchTrue{}, int64(0), is_err.Temp, panic_lbl.Temp, "", int64(0)))
			ok_val := _lnew_temp(ll)
			_ = ok_val
			ll = ok_val.L
			ll = _lemit(ll, new_inst(IrOpOpFieldGet{}, ok_val.Temp, left, int64(1), "_val", int64(0)))
			ll = _lemit(ll, new_inst(IrOpOpJump{}, int64(0), ok_lbl.Temp, int64(0), "", int64(0)))
			ll = _lemit(ll, new_inst(IrOpOpLabel{}, int64(0), panic_lbl.Temp, int64(0), "", int64(0)))
			panic_str_idx := mod_find_string(ll.Module, "unwrap failed on Err value")
			_ = panic_str_idx
			if (panic_str_idx == int64(0)) {
				new_mod := mod_add_string(ll.Module, "unwrap failed on Err value")
				_ = new_mod
				panic_str_idx = (int64(len(new_mod.String_constants)) - int64(1))
				ll = _lset_mod(ll, new_mod)
			}
			panic_ptr := _lnew_str_temp(ll)
			_ = panic_ptr
			ll = panic_ptr.L
			ll = _lemit(ll, new_inst(IrOpOpConstStr{}, panic_ptr.Temp, panic_str_idx, int64(25), "", int64(0)))
			ll = _lemit(ll, new_inst(IrOpOpCall{}, int64(0), panic_ptr.Temp, int64(2), "_aria_panic", int64(0)))
			ll = _lemit(ll, new_inst(IrOpOpLabel{}, int64(0), ok_lbl.Temp, int64(0), "", int64(0)))
			left = ok_val.Temp
		} else {
			running = false
		}
	}
	return LT{L: ll, Temp: left}
}

func _lower_primary(l Lowerer) LT {
	k := _lk(l)
	_ = k
	if (k == "@") {
		ll := _ladv(l)
		_ = ll
		if (_lk(ll) == "IDENT") {
			ann := _lcur(ll).Text
			_ = ann
			ll = _ladv(ll)
			if (_lk(ll) == "(") {
				ll = _ladv(ll)
				for ((_lk(ll) != ")") && (_lk(ll) != "EOF")) {
					ll = _ladv(ll)
				}
				if (_lk(ll) == ")") {
					ll = _ladv(ll)
				}
			}
			return _lower_primary(ll)
		}
		return _lower_primary(ll)
	}
	if (((k == "-") || (k == "!")) || (k == "~")) {
		ll := _ladv(l)
		_ = ll
		operand := _lower_atom(ll)
		_ = operand
		nt := _lnew_temp(operand.L)
		_ = nt
		op_type := _lget_type(nt.L, operand.Temp)
		_ = op_type
		if ((k == "-") && _is_float_type(op_type)) {
			lll := _lemit(nt.L, new_inst(IrOpOpNeg{}, nt.Temp, operand.Temp, int64(0), "f64", int64(0)))
			_ = lll
			llll := _lset_type(lll, nt.Temp, int64(9))
			_ = llll
			return LT{L: llll, Temp: nt.Temp}
		}
		lll := _lemit(nt.L, new_inst(_unary_op(k), nt.Temp, operand.Temp, int64(0), "", int64(0)))
		_ = lll
		return LT{L: lll, Temp: nt.Temp}
	}
	if (k == "INT") {
		val := _lcur(l).Text
		_ = val
		num := _parse_int_literal(val)
		_ = num
		nt := _lnew_temp(l)
		_ = nt
		ll := _lemit(_ladv(nt.L), new_inst(IrOpOpConst{}, nt.Temp, num, int64(0), "", int64(0)))
		_ = ll
		ll = _lset_type(ll, nt.Temp, int64(1))
		return LT{L: ll, Temp: nt.Temp}
	}
	if (k == "DURATION") {
		val := _lcur(l).Text
		_ = val
		num := _parse_duration(val)
		_ = num
		nt := _lnew_temp(l)
		_ = nt
		ll := _lemit(_ladv(nt.L), new_inst(IrOpOpConst{}, nt.Temp, num, int64(0), "", int64(0)))
		_ = ll
		ll = _lset_type(ll, nt.Temp, int64(1))
		return LT{L: ll, Temp: nt.Temp}
	}
	if (k == "SIZE") {
		val := _lcur(l).Text
		_ = val
		num := _parse_size(val)
		_ = num
		nt := _lnew_temp(l)
		_ = nt
		ll := _lemit(_ladv(nt.L), new_inst(IrOpOpConst{}, nt.Temp, num, int64(0), "", int64(0)))
		_ = ll
		ll = _lset_type(ll, nt.Temp, int64(1))
		return LT{L: ll, Temp: nt.Temp}
	}
	if (k == "FLOAT") {
		val := _lcur(l).Text
		_ = val
		ll := _ladv(l)
		_ = ll
		str_idx := mod_find_string(ll.Module, val)
		_ = str_idx
		if (str_idx == int64(0)) {
			new_mod := mod_add_string(ll.Module, val)
			_ = new_mod
			str_idx = (int64(len(new_mod.String_constants)) - int64(1))
			ll = _lset_mod(ll, new_mod)
		}
		nt := _lnew_temp(ll)
		_ = nt
		ll = nt.L
		ll = _lemit(ll, new_inst(IrOpOpConstStr{}, nt.Temp, str_idx, (int64(0) - int64(9)), "", int64(0)))
		ll = _lset_type(ll, nt.Temp, int64(9))
		return LT{L: ll, Temp: nt.Temp}
	}
	if (k == "STRING") {
		val := _lcur(l).Text
		_ = val
		ll := _ladv(l)
		_ = ll
		str_idx := mod_find_string(ll.Module, val)
		_ = str_idx
		if (str_idx == int64(0)) {
			new_mod := mod_add_string(ll.Module, val)
			_ = new_mod
			str_idx = (int64(len(new_mod.String_constants)) - int64(1))
			ll = _lset_mod(ll, new_mod)
		}
		nt := _lnew_str_temp(ll)
		_ = nt
		ll = nt.L
		ll = _lemit(ll, new_inst(IrOpOpConstStr{}, nt.Temp, str_idx, int64(len(val)), "", int64(0)))
		ll = _lset_type(ll, nt.Temp, int64(12))
		return LT{L: ll, Temp: nt.Temp}
	}
	if (k == "CHAR") {
		val := _lcur(l).Text
		_ = val
		code := int64(0)
		_ = code
		if (int64(len(val)) > int64(0)) {
			ch := string(val[int64(0)])
			_ = ch
			code = _char_to_code(ch)
		}
		nt := _lnew_temp(l)
		_ = nt
		ll := _lemit(_ladv(nt.L), new_inst(IrOpOpConst{}, nt.Temp, code, int64(0), "", int64(0)))
		_ = ll
		ll = _lset_type(ll, nt.Temp, int64(13))
		return LT{L: ll, Temp: nt.Temp}
	}
	if ((k == "RAW_STRING") || (k == "TRIPLE_STRING")) {
		val := _lcur(l).Text
		_ = val
		return _lower_string_const(_ladv(l), val)
	}
	if (k == "DURATION") {
		val := _lcur(l).Text
		_ = val
		dur_ns := _parse_duration(val)
		_ = dur_ns
		nt := _lnew_temp(l)
		_ = nt
		ll := _lemit(_ladv(nt.L), new_inst(IrOpOpConst{}, nt.Temp, dur_ns, int64(0), "", int64(0)))
		_ = ll
		ll = _lset_type(ll, nt.Temp, int64(1))
		return LT{L: ll, Temp: nt.Temp}
	}
	if (k == "SIZE") {
		val := _lcur(l).Text
		_ = val
		size_bytes := _parse_size(val)
		_ = size_bytes
		nt := _lnew_temp(l)
		_ = nt
		ll := _lemit(_ladv(nt.L), new_inst(IrOpOpConst{}, nt.Temp, size_bytes, int64(0), "", int64(0)))
		_ = ll
		ll = _lset_type(ll, nt.Temp, int64(1))
		return LT{L: ll, Temp: nt.Temp}
	}
	if (k == "STRING_START") {
		return _lower_interpolated_string(l)
	}
	if (k == "true") {
		nt := _lnew_temp(l)
		_ = nt
		ll := _lemit(_ladv(nt.L), new_inst(IrOpOpConst{}, nt.Temp, int64(1), int64(0), "", int64(0)))
		_ = ll
		return LT{L: ll, Temp: nt.Temp}
	}
	if (k == "false") {
		nt := _lnew_temp(l)
		_ = nt
		ll := _lemit(_ladv(nt.L), new_inst(IrOpOpConst{}, nt.Temp, int64(0), int64(0), "", int64(0)))
		_ = ll
		return LT{L: ll, Temp: nt.Temp}
	}
	if (k == "if") {
		return _lower_if(l)
	}
	if (k == "match") {
		return _lower_match(l)
	}
	if (k == "fn") {
		return _lower_closure(l)
	}
	if (k == "spawn") {
		return _lower_spawn(l)
	}
	if (k == "scope") {
		return _lower_scope(l)
	}
	if (k == "select") {
		return _lower_select(l)
	}
	if (k == "self") {
		ll := _ladv(l)
		_ = ll
		slot := _lenv_lookup(ll, "self")
		_ = slot
		if (slot >= int64(0)) {
			src_type := _lget_type(ll, slot)
			_ = src_type
			nt := _lnew_temp(ll)
			_ = nt
			ll = nt.L
			ll = _lemit(ll, new_inst(IrOpOpLoad{}, nt.Temp, slot, int64(0), "", int64(0)))
			if (src_type != int64(0)) {
				ll = _lset_type(ll, nt.Temp, src_type)
			}
			return LT{L: ll, Temp: nt.Temp}
		}
	}
	if ((k == "IDENT") && (_lcur(l).Text == "dyn")) {
		ll := _ladv(l)
		_ = ll
		if (_lk(ll) == "IDENT") {
			trait_name := _lcur(ll).Text
			_ = trait_name
			ll = _ladv(ll)
			if (_lk(ll) == "(") {
				ll = _ladv(ll)
				val := _lower_expr(ll)
				_ = val
				ll = val.L
				if (_lk(ll) == ")") {
					ll = _ladv(ll)
				}
				obj_type := _lget_type(ll, val.Temp)
				_ = obj_type
				concrete_name := _struct_name_for_type(ll.Store, ll.Registry, obj_type)
				_ = concrete_name
				if (concrete_name != "") {
					result := _lower_make_trait_object(ll, val.Temp, concrete_name, trait_name)
					_ = result
					return result
				}
				return LT{L: ll, Temp: val.Temp}
			}
		}
	}
	is_kw_fn_call := false
	_ = is_kw_fn_call
	if ((((k == "test") || (k == "check")) || (k == "bench")) || (k == "where")) {
		next_ll := _ladv(l)
		_ = next_ll
		if (_lk(next_ll) == "(") {
			is_kw_fn_call = true
		}
	}
	if ((k == "IDENT") || is_kw_fn_call) {
		name := _lcur(l).Text
		_ = name
		ll := _ladv(l)
		_ = ll
		if (_lk(ll) == ":") {
			next_pos := (ll.Pos + int64(1))
			_ = next_pos
			if ((next_pos < int64(len(ll.Tokens))) && (token_name(ll.Tokens[next_pos].Kind) == ":")) {
				third_pos := (ll.Pos + int64(2))
				_ = third_pos
				if ((third_pos < int64(len(ll.Tokens))) && (token_name(ll.Tokens[third_pos].Kind) == "IDENT")) {
					meth_name := ll.Tokens[third_pos].Text
					_ = meth_name
					qualified := ((name + "_") + meth_name)
					_ = qualified
					ll = _ladv(_ladv(_ladv(ll)))
					fsig := reg_find_fn(ll.Registry, qualified)
					_ = fsig
					if (fsig.Name != "") {
						ll = _lset_reg(ll, reg_add_fn_value_ref(ll.Registry, qualified))
						clos := _lnew_temp(ll)
						_ = clos
						ll = clos.L
						ll = _lemit(ll, new_inst(IrOpOpAlloc{}, clos.Temp, int64(2), int64(0), "_Closure", int64(0)))
						fn_ref := _lnew_temp(ll)
						_ = fn_ref
						ll = fn_ref.L
						ll = _lemit(ll, new_inst(IrOpOpFnRef{}, fn_ref.Temp, int64(0), int64(0), ("_tramp_" + qualified), int64(0)))
						ll = _lemit(ll, new_inst(IrOpOpFieldSet{}, clos.Temp, fn_ref.Temp, int64(0), "_fn", int64(0)))
						null_env := _lnew_temp(ll)
						_ = null_env
						ll = null_env.L
						ll = _lemit(ll, new_inst(IrOpOpConst{}, null_env.Temp, int64(0), int64(0), "", int64(0)))
						ll = _lemit(ll, new_inst(IrOpOpFieldSet{}, clos.Temp, null_env.Temp, int64(1), "_env", int64(0)))
						return LT{L: ll, Temp: clos.Temp}
					}
				}
			}
		}
		if (_lk(ll) == "(") {
			if ((name == "Ok") || (name == "Err")) {
				ll = _ladv(ll)
				ll = _lskip_nl(ll)
				val := _lower_expr(ll)
				_ = val
				ll = val.L
				ll = _lskip_nl(ll)
				if (_lk(ll) == ")") {
					ll = _ladv(ll)
				}
				val_type := _lget_type(ll, val.Temp)
				_ = val_type
				result_slots := int64(2)
				_ = result_slots
				if (val_type == int64(12)) {
					result_slots = int64(3)
				}
				nt := _lnew_temp(ll)
				_ = nt
				ll = nt.L
				ll = _lemit(ll, new_inst(IrOpOpAlloc{}, nt.Temp, result_slots, int64(0), "_Result", int64(0)))
				tag_nt := _lnew_temp(ll)
				_ = tag_nt
				ll = tag_nt.L
				tag_val := int64(0)
				_ = tag_val
				if (name == "Err") {
					tag_val = int64(1)
				}
				ll = _lemit(ll, new_inst(IrOpOpConst{}, tag_nt.Temp, tag_val, int64(0), "", int64(0)))
				ll = _lemit(ll, new_inst(IrOpOpFieldSet{}, nt.Temp, tag_nt.Temp, int64(0), "_tag", int64(0)))
				if (val_type == int64(12)) {
					ll = _lemit(ll, new_inst(IrOpOpFieldSet{}, nt.Temp, val.Temp, int64(1), "_val", int64(0)))
					val_len := (val.Temp + int64(1))
					_ = val_len
					ll = _lemit(ll, new_inst(IrOpOpFieldSet{}, nt.Temp, val_len, int64(2), "_val_len", int64(0)))
				} else {
					ll = _lemit(ll, new_inst(IrOpOpFieldSet{}, nt.Temp, val.Temp, int64(1), "_val", int64(0)))
				}
				ll = _lset_type(ll, nt.Temp, int64(200))
				return LT{L: ll, Temp: nt.Temp}
			}
			if (name == "Some") {
				ll = _ladv(ll)
				ll = _lskip_nl(ll)
				val := _lower_expr(ll)
				_ = val
				ll = val.L
				ll = _lskip_nl(ll)
				if (_lk(ll) == ")") {
					ll = _ladv(ll)
				}
				val_type := _lget_type(ll, val.Temp)
				_ = val_type
				opt_slots := int64(2)
				_ = opt_slots
				if (val_type == int64(12)) {
					opt_slots = int64(3)
				}
				nt := _lnew_temp(ll)
				_ = nt
				ll = nt.L
				ll = _lemit(ll, new_inst(IrOpOpAlloc{}, nt.Temp, opt_slots, int64(0), "_Optional", int64(0)))
				tag_nt := _lnew_temp(ll)
				_ = tag_nt
				ll = tag_nt.L
				ll = _lemit(ll, new_inst(IrOpOpConst{}, tag_nt.Temp, int64(1), int64(0), "", int64(0)))
				ll = _lemit(ll, new_inst(IrOpOpFieldSet{}, nt.Temp, tag_nt.Temp, int64(0), "_tag", int64(0)))
				if (val_type == int64(12)) {
					ll = _lemit(ll, new_inst(IrOpOpFieldSet{}, nt.Temp, val.Temp, int64(1), "_val", int64(0)))
					val_len := (val.Temp + int64(1))
					_ = val_len
					ll = _lemit(ll, new_inst(IrOpOpFieldSet{}, nt.Temp, val_len, int64(2), "_val_len", int64(0)))
				} else {
					ll = _lemit(ll, new_inst(IrOpOpFieldSet{}, nt.Temp, val.Temp, int64(1), "_val", int64(0)))
				}
				ll = _lset_type(ll, nt.Temp, int64(201))
				return LT{L: ll, Temp: nt.Temp}
			}
			if (name == "Map") {
				ll = _ladv(ll)
				if (_lk(ll) == ")") {
					ll = _ladv(ll)
				}
				cap_nt := _lnew_temp(ll)
				_ = cap_nt
				ll = cap_nt.L
				ll = _lemit(ll, new_inst(IrOpOpConst{}, cap_nt.Temp, int64(16), int64(0), "", int64(0)))
				nt := _lnew_temp(ll)
				_ = nt
				ll = nt.L
				ll = _lemit(ll, new_inst(IrOpOpMapNew{}, nt.Temp, cap_nt.Temp, int64(0), "", int64(0)))
				ll = _lset_type(ll, nt.Temp, int64(300))
				return LT{L: ll, Temp: nt.Temp}
			}
			if (name == "Set") {
				ll = _ladv(ll)
				if (_lk(ll) == ")") {
					ll = _ladv(ll)
				}
				cap_nt := _lnew_temp(ll)
				_ = cap_nt
				ll = cap_nt.L
				ll = _lemit(ll, new_inst(IrOpOpConst{}, cap_nt.Temp, int64(16), int64(0), "", int64(0)))
				nt := _lnew_temp(ll)
				_ = nt
				ll = nt.L
				ll = _lemit(ll, new_inst(IrOpOpSetNew{}, nt.Temp, cap_nt.Temp, int64(0), "", int64(0)))
				ll = _lset_type(ll, nt.Temp, int64(301))
				return LT{L: ll, Temp: nt.Temp}
			}
			vtag := _variant_tag(ll.Registry, name)
			_ = vtag
			if (vtag > int64(0)) {
				vparent := _variant_parent(ll.Registry, name)
				_ = vparent
				vdc := variant_data_count(vparent, vtag)
				_ = vdc
				if (vdc > int64(0)) {
					alloc_sz := sum_alloc_size(vparent)
					_ = alloc_sz
					nt := _lnew_temp(ll)
					_ = nt
					ll = nt.L
					ll = _lemit(ll, new_inst(IrOpOpAlloc{}, nt.Temp, alloc_sz, int64(0), vparent.Name, int64(0)))
					st_id := _struct_type_id(ll.Registry, vparent.Name)
					_ = st_id
					if (st_id > int64(0)) {
						ll = _lset_type(ll, nt.Temp, st_id)
					}
					vtag_nt := _lnew_temp(ll)
					_ = vtag_nt
					ll = vtag_nt.L
					ll = _lemit(ll, new_inst(IrOpOpConst{}, vtag_nt.Temp, vtag, int64(0), "", int64(0)))
					ll = _lemit(ll, new_inst(IrOpOpFieldSet{}, nt.Temp, vtag_nt.Temp, int64(0), "_tag", int64(0)))
					ll = _ladv(ll)
					ll = _lskip_nl(ll)
					vfi := int64(0)
					_ = vfi
					for ((_lk(ll) != ")") && (_lk(ll) != "EOF")) {
						val := _lower_expr(ll)
						_ = val
						ll = val.L
						vdata_slot := int64(1)
						_ = vdata_slot
						vsi := int64(0)
						_ = vsi
						for (vsi < vfi) {
							vft := variant_field_type(vparent, vtag, vsi)
							_ = vft
							if (vft == int64(12)) {
								vdata_slot = (vdata_slot + int64(2))
							} else {
								vdata_slot = (vdata_slot + int64(1))
							}
							vsi = (vsi + int64(1))
						}
						vft := variant_field_type(vparent, vtag, vfi)
						_ = vft
						if (vft == int64(12)) {
							ll = _lemit(ll, new_inst(IrOpOpFieldSet{}, nt.Temp, val.Temp, vdata_slot, "", int64(0)))
							ll = _lemit(ll, new_inst(IrOpOpFieldSet{}, nt.Temp, (val.Temp + int64(1)), (vdata_slot + int64(1)), "", int64(0)))
						} else {
							ll = _lemit(ll, new_inst(IrOpOpFieldSet{}, nt.Temp, val.Temp, vdata_slot, "", int64(0)))
						}
						vfi = (vfi + int64(1))
						ll = _lskip_nl(ll)
						if (_lk(ll) == ",") {
							ll = _ladv(ll)
						}
						ll = _lskip_nl(ll)
					}
					if (_lk(ll) == ")") {
						ll = _ladv(ll)
					}
					return LT{L: ll, Temp: nt.Temp}
				}
			}
			nt_def := reg_find_struct(ll.Registry, name)
			_ = nt_def
			if (((nt_def.Name != "") && (nt_def.Is_sum == false)) && _is_newtype(nt_def)) {
				ll = _ladv(ll)
				ll = _lskip_nl(ll)
				val := _lower_expr(ll)
				_ = val
				ll = val.L
				ll = _lskip_nl(ll)
				if (_lk(ll) == ")") {
					ll = _ladv(ll)
				}
				nt := _lnew_temp(ll)
				_ = nt
				ll = nt.L
				ll = _lemit(ll, new_inst(IrOpOpAlloc{}, nt.Temp, int64(1), int64(0), name, int64(0)))
				st_id := _struct_type_id(ll.Registry, name)
				_ = st_id
				if (st_id > int64(0)) {
					ll = _lset_type(ll, nt.Temp, st_id)
				}
				ll = _lemit(ll, new_inst(IrOpOpFieldSet{}, nt.Temp, val.Temp, int64(0), "value", int64(0)))
				return LT{L: ll, Temp: nt.Temp}
			}
			mock_slot := _lenv_lookup(ll, ("_mock_" + name))
			_ = mock_slot
			if ((mock_slot >= int64(0)) && (_lk(ll) == "(")) {
				closure_nt := _lnew_temp(ll)
				_ = closure_nt
				ll = closure_nt.L
				ll = _lemit(ll, new_inst(IrOpOpLoad{}, closure_nt.Temp, mock_slot, int64(0), "", int64(0)))
				fn_ptr_nt := _lnew_temp(ll)
				_ = fn_ptr_nt
				ll = fn_ptr_nt.L
				ll = _lemit(ll, new_inst(IrOpOpFieldGet{}, fn_ptr_nt.Temp, closure_nt.Temp, int64(0), "_fn", int64(0)))
				env_ptr_nt := _lnew_temp(ll)
				_ = env_ptr_nt
				ll = env_ptr_nt.L
				ll = _lemit(ll, new_inst(IrOpOpFieldGet{}, env_ptr_nt.Temp, closure_nt.Temp, int64(1), "_env", int64(0)))
				ll = _ladv(ll)
				ll = _lskip_nl(ll)
				ic_arg_count := int64(0)
				_ = ic_arg_count
				ic_reg_count := int64(0)
				_ = ic_reg_count
				first_arg := -int64(1)
				_ = first_arg
				for ((_lk(ll) != ")") && (_lk(ll) != "EOF")) {
					ll = _lskip_nl(ll)
					arg := _lower_expr(ll)
					_ = arg
					ll = arg.L
					if (ic_arg_count == int64(0)) {
						first_arg = arg.Temp
					}
					atid := _lget_type(ll, arg.Temp)
					_ = atid
					if (atid == int64(12)) {
						ic_reg_count = (ic_reg_count + int64(2))
					} else {
						ic_reg_count = (ic_reg_count + int64(1))
					}
					ic_arg_count = (ic_arg_count + int64(1))
					ll = _lskip_nl(ll)
					if (_lk(ll) == ",") {
						ll = _ladv(ll)
					}
				}
				if (_lk(ll) == ")") {
					ll = _ladv(ll)
				}
				total_args := (ic_reg_count + int64(1))
				_ = total_args
				pack_start := _lnew_temp(ll)
				_ = pack_start
				ll = pack_start.L
				ll = _lemit(ll, new_inst(IrOpOpStore{}, pack_start.Temp, env_ptr_nt.Temp, int64(0), "", int64(0)))
				pack_idx := int64(1)
				_ = pack_idx
				if ((ic_arg_count > int64(0)) && (first_arg >= int64(0))) {
					ai := int64(0)
					_ = ai
					for (ai < ic_arg_count) {
						pack_t := _lnew_temp(ll)
						_ = pack_t
						ll = pack_t.L
						ll = _lemit(ll, new_inst(IrOpOpLoad{}, pack_t.Temp, (first_arg + ai), int64(0), "", int64(0)))
						ll = _lemit(ll, new_inst(IrOpOpStore{}, (pack_start.Temp + pack_idx), pack_t.Temp, int64(0), "", int64(0)))
						pack_idx = (pack_idx + int64(1))
						ai = (ai + int64(1))
					}
				}
				nt := _lnew_temp(ll)
				_ = nt
				ll = nt.L
				ll = _lemit(ll, new_inst(IrOpOpCallIndirect{}, nt.Temp, fn_ptr_nt.Temp, total_args, "", pack_start.Temp))
				return LT{L: ll, Temp: nt.Temp}
			}
			sig := reg_find_fn(ll.Registry, name)
			_ = sig
			if (sig.Name != "") {
				call_name := name
				_ = call_name
				call_sig := sig
				_ = call_sig
				if fn_is_generic(sig) {
					call_sig = sig
				}
				ll = _ladv(ll)
				ll = _lskip_nl(ll)
				arg_temps := []int64{int64(0)}
				_ = arg_temps
				arg_types := []int64{int64(0)}
				_ = arg_types
				param_idx := int64(1)
				_ = param_idx
				for ((_lk(ll) != ")") && (_lk(ll) != "EOF")) {
					ll = _lskip_nl(ll)
					arg := _lower_expr(ll)
					_ = arg
					ll = arg.L
					arg_type := _lget_type(ll, arg.Temp)
					_ = arg_type
					arg_temps = append(arg_temps, arg.Temp)
					arg_types = append(arg_types, arg_type)
					param_idx = (param_idx + int64(1))
					ll = _lskip_nl(ll)
					if (_lk(ll) == ",") {
						ll = _ladv(ll)
					}
				}
				if (_lk(ll) == ")") {
					ll = _ladv(ll)
				}
				if fn_is_generic(sig) {
					call_name = _resolve_generic_call_name(ll.Store, ll.Registry, name, sig, arg_types)
					sp_sig := reg_find_fn(ll.Registry, call_name)
					_ = sp_sig
					if ((sp_sig.Name != "") && (fn_is_generic(sp_sig) == false)) {
						call_sig = sp_sig
					} else if (int64(len(arg_types)) > int64(1)) {
						sp := reg_find_mono_by_generic(ll.Registry, name, arg_types[int64(1)])
						_ = sp
						if (sp.Specialized_name == "") {
							msi := int64(1)
							_ = msi
							for (msi < int64(len(ll.Registry.Mono_specs))) {
								if (ll.Registry.Mono_specs[msi].Generic_name == name) {
									sp = ll.Registry.Mono_specs[msi]
									msi = int64(len(ll.Registry.Mono_specs))
								}
								msi = (msi + int64(1))
							}
						}
						if (sp.Specialized_name != "") {
							call_name = sp.Specialized_name
							sp_sig2 := reg_find_fn(ll.Registry, call_name)
							_ = sp_sig2
							if (sp_sig2.Name != "") {
								call_sig = sp_sig2
							}
						}
					}
				}
				reg_count := int64(0)
				_ = reg_count
				ai := int64(1)
				_ = ai
				for (ai < int64(len(arg_temps))) {
					if (arg_types[ai] == int64(12)) {
						reg_count = (reg_count + int64(2))
					} else {
						reg_count = (reg_count + int64(1))
					}
					ai = (ai + int64(1))
				}
				already_consecutive := true
				_ = already_consecutive
				if (int64(len(arg_temps)) > int64(1)) {
					expected := arg_temps[int64(1)]
					_ = expected
					ci := int64(1)
					_ = ci
					for (ci < int64(len(arg_temps))) {
						if (arg_temps[ci] != expected) {
							already_consecutive = false
						}
						if (arg_types[ci] == int64(12)) {
							expected = (expected + int64(2))
						} else {
							expected = (expected + int64(1))
						}
						ci = (ci + int64(1))
					}
				}
				fa := int64(0)
				_ = fa
				if (already_consecutive && (int64(len(arg_temps)) > int64(1))) {
					fa = arg_temps[int64(1)]
				} else if (int64(len(arg_temps)) > int64(1)) {
					first_nt := _lnew_temp(ll)
					_ = first_nt
					ll = first_nt.L
					ri := int64(1)
					_ = ri
					for (ri < reg_count) {
						extra := _lnew_temp(ll)
						_ = extra
						ll = extra.L
						ri = (ri + int64(1))
					}
					slot := int64(0)
					_ = slot
					ai = int64(1)
					for (ai < int64(len(arg_temps))) {
						if (arg_types[ai] == int64(12)) {
							ll = _lemit(ll, new_inst(IrOpOpStore{}, (first_nt.Temp + slot), arg_temps[ai], int64(0), "", int64(0)))
							slot_plus := (slot + int64(1))
							_ = slot_plus
							ll = _lemit(ll, new_inst(IrOpOpStore{}, (first_nt.Temp + slot_plus), (arg_temps[ai] + int64(1)), int64(0), "", int64(0)))
							slot = (slot + int64(2))
						} else {
							ll = _lemit(ll, new_inst(IrOpOpStore{}, (first_nt.Temp + slot), arg_temps[ai], int64(0), "", int64(0)))
							slot = (slot + int64(1))
						}
						ai = (ai + int64(1))
					}
					fa = first_nt.Temp
				}
				nt_temp := int64(0)
				_ = nt_temp
				if (call_sig.Return_type == int64(12)) {
					nt := _lnew_str_temp(ll)
					_ = nt
					ll = nt.L
					nt_temp = nt.Temp
				} else {
					nt := _lnew_temp(ll)
					_ = nt
					ll = nt.L
					nt_temp = nt.Temp
				}
				ll = _lemit(ll, new_inst(IrOpOpCall{}, nt_temp, fa, reg_count, call_name, call_sig.Return_type))
				if (call_sig.Return_type == int64(12)) {
					ll = _lset_type(ll, nt_temp, int64(12))
				} else if (call_sig.Return_type > int64(0)) {
					ll = _lset_type(ll, nt_temp, call_sig.Return_type)
				}
				return LT{L: ll, Temp: nt_temp}
			}
			if (name == "dbg") {
				ll = _ladv(ll)
				arg := _lower_expr(ll)
				_ = arg
				ll = arg.L
				if (_lk(ll) == ")") {
					ll = _ladv(ll)
				}
				at := _lget_type(ll, arg.Temp)
				_ = at
				prefix := (("[" + ll.File) + "] ")
				_ = prefix
				prefix_res := _lower_string_const(ll, prefix)
				_ = prefix_res
				ll = prefix_res.L
				if (at == int64(12)) {
					cat := _lnew_str_temp(ll)
					_ = cat
					ll = cat.L
					ll = _lemit(ll, new_inst(IrOpOpStrConcat{}, cat.Temp, prefix_res.Temp, arg.Temp, "", int64(0)))
					pnt := _lnew_temp(ll)
					_ = pnt
					ll = pnt.L
					ll = _lemit(ll, new_inst(IrOpOpCall{}, pnt.Temp, cat.Temp, int64(2), "_aria_println_str", int64(0)))
				} else {
					vstr := _lnew_str_temp(ll)
					_ = vstr
					ll = vstr.L
					ll = _lemit(ll, new_inst(IrOpOpIntToStr{}, vstr.Temp, arg.Temp, int64(0), "", int64(0)))
					cat := _lnew_str_temp(ll)
					_ = cat
					ll = cat.L
					ll = _lemit(ll, new_inst(IrOpOpStrConcat{}, cat.Temp, prefix_res.Temp, vstr.Temp, "", int64(0)))
					pnt := _lnew_temp(ll)
					_ = pnt
					ll = pnt.L
					ll = _lemit(ll, new_inst(IrOpOpCall{}, pnt.Temp, cat.Temp, int64(2), "_aria_println_str", int64(0)))
				}
				return LT{L: ll, Temp: arg.Temp}
			}
			if (name == "assertEqual") {
				ll = _ladv(ll)
				arg_a := _lower_expr(ll)
				_ = arg_a
				ll = arg_a.L
				if (_lk(ll) == ",") {
					ll = _ladv(ll)
				}
				ll = _lskip_nl(ll)
				arg_b := _lower_expr(ll)
				_ = arg_b
				ll = arg_b.L
				if (_lk(ll) == ")") {
					ll = _ladv(ll)
				}
				at := _lget_type(ll, arg_a.Temp)
				_ = at
				ok_lbl := _lnew_label(ll)
				_ = ok_lbl
				ll = ok_lbl.L
				if (at == int64(12)) {
					eq_nt := _lnew_temp(ll)
					_ = eq_nt
					ll = eq_nt.L
					ll = _lemit(ll, new_inst(IrOpOpStrEq{}, eq_nt.Temp, arg_a.Temp, arg_b.Temp, "", int64(0)))
					ll = _lemit(ll, new_inst(IrOpOpBranchTrue{}, int64(0), eq_nt.Temp, ok_lbl.Temp, "", int64(0)))
				} else {
					eq_nt := _lnew_temp(ll)
					_ = eq_nt
					ll = eq_nt.L
					ll = _lemit(ll, new_inst(IrOpOpEq{}, eq_nt.Temp, arg_a.Temp, arg_b.Temp, "", int64(0)))
					ll = _lemit(ll, new_inst(IrOpOpBranchTrue{}, int64(0), eq_nt.Temp, ok_lbl.Temp, "", int64(0)))
				}
				msg := _lower_string_const(ll, "assertEqual failed")
				_ = msg
				ll = msg.L
				pnt := _lnew_temp(ll)
				_ = pnt
				ll = pnt.L
				ll = _lemit(ll, new_inst(IrOpOpCall{}, pnt.Temp, msg.Temp, int64(2), "_aria_println_str", int64(0)))
				exit_nt := _lnew_temp(ll)
				_ = exit_nt
				ll = exit_nt.L
				ll = _lemit(ll, new_inst(IrOpOpConst{}, exit_nt.Temp, int64(1), int64(0), "", int64(0)))
				ll = _lemit(ll, new_inst(IrOpOpCall{}, int64(0), exit_nt.Temp, int64(1), "_aria_exit", int64(0)))
				ll = _lemit(ll, new_inst(IrOpOpLabel{}, int64(0), ok_lbl.Temp, int64(0), "", int64(0)))
				return LT{L: ll, Temp: int64(0)}
			}
			if (name == "assertOk") {
				ll = _ladv(ll)
				arg := _lower_expr(ll)
				_ = arg
				ll = arg.L
				if (_lk(ll) == ")") {
					ll = _ladv(ll)
				}
				tag_nt := _lnew_temp(ll)
				_ = tag_nt
				ll = tag_nt.L
				ll = _lemit(ll, new_inst(IrOpOpFieldGet{}, tag_nt.Temp, arg.Temp, int64(0), "_tag", int64(0)))
				ok_lbl := _lnew_label(ll)
				_ = ok_lbl
				ll = ok_lbl.L
				zero_nt := _lnew_temp(ll)
				_ = zero_nt
				ll = zero_nt.L
				ll = _lemit(ll, new_inst(IrOpOpConst{}, zero_nt.Temp, int64(0), int64(0), "", int64(0)))
				cmp_nt := _lnew_temp(ll)
				_ = cmp_nt
				ll = cmp_nt.L
				ll = _lemit(ll, new_inst(IrOpOpEq{}, cmp_nt.Temp, tag_nt.Temp, zero_nt.Temp, "", int64(0)))
				ll = _lemit(ll, new_inst(IrOpOpBranchTrue{}, int64(0), cmp_nt.Temp, ok_lbl.Temp, "", int64(0)))
				msg := _lower_string_const(ll, "assertOk failed: got Err")
				_ = msg
				ll = msg.L
				pnt := _lnew_temp(ll)
				_ = pnt
				ll = pnt.L
				ll = _lemit(ll, new_inst(IrOpOpCall{}, pnt.Temp, msg.Temp, int64(2), "_aria_println_str", int64(0)))
				exit_nt := _lnew_temp(ll)
				_ = exit_nt
				ll = exit_nt.L
				ll = _lemit(ll, new_inst(IrOpOpConst{}, exit_nt.Temp, int64(1), int64(0), "", int64(0)))
				ll = _lemit(ll, new_inst(IrOpOpCall{}, int64(0), exit_nt.Temp, int64(1), "_aria_exit", int64(0)))
				ll = _lemit(ll, new_inst(IrOpOpLabel{}, int64(0), ok_lbl.Temp, int64(0), "", int64(0)))
				return LT{L: ll, Temp: int64(0)}
			}
			if (name == "assertErr") {
				ll = _ladv(ll)
				arg := _lower_expr(ll)
				_ = arg
				ll = arg.L
				if (_lk(ll) == ")") {
					ll = _ladv(ll)
				}
				tag_nt := _lnew_temp(ll)
				_ = tag_nt
				ll = tag_nt.L
				ll = _lemit(ll, new_inst(IrOpOpFieldGet{}, tag_nt.Temp, arg.Temp, int64(0), "_tag", int64(0)))
				ok_lbl := _lnew_label(ll)
				_ = ok_lbl
				ll = ok_lbl.L
				zero_nt := _lnew_temp(ll)
				_ = zero_nt
				ll = zero_nt.L
				ll = _lemit(ll, new_inst(IrOpOpConst{}, zero_nt.Temp, int64(0), int64(0), "", int64(0)))
				cmp_nt := _lnew_temp(ll)
				_ = cmp_nt
				ll = cmp_nt.L
				ll = _lemit(ll, new_inst(IrOpOpNeq{}, cmp_nt.Temp, tag_nt.Temp, zero_nt.Temp, "", int64(0)))
				ll = _lemit(ll, new_inst(IrOpOpBranchTrue{}, int64(0), cmp_nt.Temp, ok_lbl.Temp, "", int64(0)))
				msg := _lower_string_const(ll, "assertErr failed: got Ok")
				_ = msg
				ll = msg.L
				pnt := _lnew_temp(ll)
				_ = pnt
				ll = pnt.L
				ll = _lemit(ll, new_inst(IrOpOpCall{}, pnt.Temp, msg.Temp, int64(2), "_aria_println_str", int64(0)))
				exit_nt := _lnew_temp(ll)
				_ = exit_nt
				ll = exit_nt.L
				ll = _lemit(ll, new_inst(IrOpOpConst{}, exit_nt.Temp, int64(1), int64(0), "", int64(0)))
				ll = _lemit(ll, new_inst(IrOpOpCall{}, int64(0), exit_nt.Temp, int64(1), "_aria_exit", int64(0)))
				ll = _lemit(ll, new_inst(IrOpOpLabel{}, int64(0), ok_lbl.Temp, int64(0), "", int64(0)))
				return LT{L: ll, Temp: int64(0)}
			}
			if (((name == "println") || (name == "print")) || (name == "eprintln")) {
				ll = _ladv(ll)
				arg := _lower_expr(ll)
				_ = arg
				ll = arg.L
				if (_lk(ll) == ")") {
					ll = _ladv(ll)
				}
				print_temp := arg.Temp
				_ = print_temp
				arg_type := _lget_type(ll, arg.Temp)
				_ = arg_type
				if ((arg_type == int64(1)) || (arg_type == int64(0))) {
					if (arg_type == int64(1)) {
						conv_nt := _lnew_str_temp(ll)
						_ = conv_nt
						ll = conv_nt.L
						ll = _lemit(ll, new_inst(IrOpOpIntToStr{}, conv_nt.Temp, arg.Temp, int64(0), "", int64(0)))
						ll = _lset_type(ll, conv_nt.Temp, int64(12))
						print_temp = conv_nt.Temp
					}
				}
				nt := _lnew_temp(ll)
				_ = nt
				ll = nt.L
				ll = _lemit(ll, new_inst(IrOpOpCall{}, nt.Temp, print_temp, int64(2), "_aria_println_str", int64(0)))
				return LT{L: ll, Temp: nt.Temp}
			}
			if (name == "panic") {
				ll = _ladv(ll)
				arg := _lower_expr(ll)
				_ = arg
				ll = arg.L
				if (_lk(ll) == ")") {
					ll = _ladv(ll)
				}
				nt := _lnew_temp(ll)
				_ = nt
				ll = nt.L
				ll = _lemit(ll, new_inst(IrOpOpCall{}, nt.Temp, arg.Temp, int64(1), "_aria_exit", int64(0)))
				return LT{L: ll, Temp: nt.Temp}
			}
			if (name == "_ariaReadFile") {
				ll = _ladv(ll)
				arg := _lower_expr(ll)
				_ = arg
				ll = arg.L
				if (_lk(ll) == ")") {
					ll = _ladv(ll)
				}
				nt := _lnew_str_temp(ll)
				_ = nt
				ll = nt.L
				ll = _lemit(ll, new_inst(IrOpOpCall{}, nt.Temp, arg.Temp, int64(2), "_aria_read_file", int64(12)))
				ll = _lset_type(ll, nt.Temp, int64(12))
				return LT{L: ll, Temp: nt.Temp}
			}
			if (name == "_ariaWriteBinaryFile") {
				ll = _ladv(ll)
				path_arg := _lower_expr(ll)
				_ = path_arg
				ll = path_arg.L
				if (_lk(ll) == ",") {
					ll = _ladv(ll)
				}
				data_arg := _lower_expr(ll)
				_ = data_arg
				ll = data_arg.L
				if (_lk(ll) == ")") {
					ll = _ladv(ll)
				}
				nt := _lnew_temp(ll)
				_ = nt
				ll = nt.L
				ll = _lemit(ll, new_inst(IrOpOpCall{}, nt.Temp, path_arg.Temp, int64(4), "_aria_write_binary_file", int64(0)))
				return LT{L: ll, Temp: nt.Temp}
			}
			if ((name == "_ariaWriteFile") || (name == "_ariaAppendFile")) {
				ll = _ladv(ll)
				path_arg := _lower_expr(ll)
				_ = path_arg
				ll = path_arg.L
				if (_lk(ll) == ",") {
					ll = _ladv(ll)
				}
				content_arg := _lower_expr(ll)
				_ = content_arg
				ll = content_arg.L
				if (_lk(ll) == ")") {
					ll = _ladv(ll)
				}
				p0 := _lnew_temp(ll)
				_ = p0
				ll = p0.L
				ll = _lemit(ll, new_inst(IrOpOpStore{}, p0.Temp, path_arg.Temp, int64(0), "", int64(0)))
				p1 := _lnew_temp(ll)
				_ = p1
				ll = p1.L
				ll = _lemit(ll, new_inst(IrOpOpStore{}, p1.Temp, (path_arg.Temp + int64(1)), int64(0), "", int64(0)))
				c0 := _lnew_temp(ll)
				_ = c0
				ll = c0.L
				ll = _lemit(ll, new_inst(IrOpOpStore{}, c0.Temp, content_arg.Temp, int64(0), "", int64(0)))
				c1 := _lnew_temp(ll)
				_ = c1
				ll = c1.L
				ll = _lemit(ll, new_inst(IrOpOpStore{}, c1.Temp, (content_arg.Temp + int64(1)), int64(0), "", int64(0)))
				nt := _lnew_temp(ll)
				_ = nt
				ll = nt.L
				rt_name := "_aria_write_file"
				_ = rt_name
				if (name == "_ariaAppendFile") {
					rt_name = "_aria_append_file"
				}
				ll = _lemit(ll, new_inst(IrOpOpCall{}, nt.Temp, p0.Temp, int64(4), rt_name, int64(0)))
				return LT{L: ll, Temp: nt.Temp}
			}
			if (name == "_ariaExec") {
				ll = _ladv(ll)
				arg := _lower_expr(ll)
				_ = arg
				ll = arg.L
				if (_lk(ll) == ")") {
					ll = _ladv(ll)
				}
				nt := _lnew_temp(ll)
				_ = nt
				ll = nt.L
				ll = _lemit(ll, new_inst(IrOpOpCall{}, nt.Temp, arg.Temp, int64(2), "_aria_exec", int64(1)))
				ll = _lset_type(ll, nt.Temp, int64(1))
				return LT{L: ll, Temp: nt.Temp}
			}
			if (name == "_ariaArgs") {
				ll = _ladv(ll)
				for ((_lk(ll) != ")") && (_lk(ll) != "EOF")) {
					ll = _ladv(ll)
				}
				if (_lk(ll) == ")") {
					ll = _ladv(ll)
				}
				nt := _lnew_temp(ll)
				_ = nt
				ll = nt.L
				ll = _lemit(ll, new_inst(IrOpOpCall{}, nt.Temp, int64(0), int64(0), "_aria_args", int64(0)))
				ll = _lset_type(ll, nt.Temp, int64(13))
				return LT{L: ll, Temp: nt.Temp}
			}
			if (name == "_ariaListDir") {
				ll = _ladv(ll)
				arg := _lower_expr(ll)
				_ = arg
				ll = arg.L
				if (_lk(ll) == ")") {
					ll = _ladv(ll)
				}
				nt := _lnew_temp(ll)
				_ = nt
				ll = nt.L
				ll = _lemit(ll, new_inst(IrOpOpCall{}, nt.Temp, arg.Temp, int64(2), "_aria_list_dir", int64(0)))
				ll = _lset_type(ll, nt.Temp, int64(13))
				return LT{L: ll, Temp: nt.Temp}
			}
			if (name == "_ariaIsDir") {
				ll = _ladv(ll)
				arg := _lower_expr(ll)
				_ = arg
				ll = arg.L
				if (_lk(ll) == ")") {
					ll = _ladv(ll)
				}
				nt := _lnew_temp(ll)
				_ = nt
				ll = nt.L
				ll = _lemit(ll, new_inst(IrOpOpCall{}, nt.Temp, arg.Temp, int64(2), "_aria_is_dir", int64(1)))
				ll = _lset_type(ll, nt.Temp, int64(1))
				return LT{L: ll, Temp: nt.Temp}
			}
			if (name == "_ariaGetenv") {
				ll = _ladv(ll)
				arg := _lower_expr(ll)
				_ = arg
				ll = arg.L
				if (_lk(ll) == ")") {
					ll = _ladv(ll)
				}
				nt := _lnew_str_temp(ll)
				_ = nt
				ll = nt.L
				ll = _lemit(ll, new_inst(IrOpOpCall{}, nt.Temp, arg.Temp, int64(2), "_aria_getenv", int64(12)))
				ll = _lset_type(ll, nt.Temp, int64(12))
				return LT{L: ll, Temp: nt.Temp}
			}
			if (name == "_ariaTcpSocket") {
				ll = _ladv(ll)
				for ((_lk(ll) != ")") && (_lk(ll) != "EOF")) {
					ll = _ladv(ll)
				}
				if (_lk(ll) == ")") {
					ll = _ladv(ll)
				}
				nt := _lnew_temp(ll)
				_ = nt
				ll = nt.L
				ll = _lemit(ll, new_inst(IrOpOpCall{}, nt.Temp, int64(0), int64(0), "_aria_tcp_socket", int64(1)))
				ll = _lset_type(ll, nt.Temp, int64(1))
				return LT{L: ll, Temp: nt.Temp}
			}
			if (name == "_ariaTcpBind") {
				ll = _ladv(ll)
				fd_arg := _lower_expr(ll)
				_ = fd_arg
				ll = fd_arg.L
				if (_lk(ll) == ",") {
					ll = _ladv(ll)
				}
				addr_arg := _lower_expr(ll)
				_ = addr_arg
				ll = addr_arg.L
				if (_lk(ll) == ",") {
					ll = _ladv(ll)
				}
				port_arg := _lower_expr(ll)
				_ = port_arg
				ll = port_arg.L
				if (_lk(ll) == ")") {
					ll = _ladv(ll)
				}
				nt := _lnew_temp(ll)
				_ = nt
				ll = nt.L
				ll = _lemit(ll, new_inst(IrOpOpCall{}, nt.Temp, fd_arg.Temp, int64(4), "_aria_tcp_bind", int64(1)))
				ll = _lset_type(ll, nt.Temp, int64(1))
				return LT{L: ll, Temp: nt.Temp}
			}
			if (name == "_ariaTcpListen") {
				ll = _ladv(ll)
				fd_arg := _lower_expr(ll)
				_ = fd_arg
				ll = fd_arg.L
				if (_lk(ll) == ",") {
					ll = _ladv(ll)
				}
				backlog_arg := _lower_expr(ll)
				_ = backlog_arg
				ll = backlog_arg.L
				if (_lk(ll) == ")") {
					ll = _ladv(ll)
				}
				nt := _lnew_temp(ll)
				_ = nt
				ll = nt.L
				ll = _lemit(ll, new_inst(IrOpOpCall{}, nt.Temp, fd_arg.Temp, int64(2), "_aria_tcp_listen", int64(1)))
				ll = _lset_type(ll, nt.Temp, int64(1))
				return LT{L: ll, Temp: nt.Temp}
			}
			if (name == "_ariaTcpAccept") {
				ll = _ladv(ll)
				fd_arg := _lower_expr(ll)
				_ = fd_arg
				ll = fd_arg.L
				if (_lk(ll) == ")") {
					ll = _ladv(ll)
				}
				nt := _lnew_temp(ll)
				_ = nt
				ll = nt.L
				ll = _lemit(ll, new_inst(IrOpOpCall{}, nt.Temp, fd_arg.Temp, int64(1), "_aria_tcp_accept", int64(1)))
				ll = _lset_type(ll, nt.Temp, int64(1))
				return LT{L: ll, Temp: nt.Temp}
			}
			if (name == "_ariaTcpRead") {
				ll = _ladv(ll)
				fd_arg := _lower_expr(ll)
				_ = fd_arg
				ll = fd_arg.L
				if (_lk(ll) == ",") {
					ll = _ladv(ll)
				}
				max_arg := _lower_expr(ll)
				_ = max_arg
				ll = max_arg.L
				if (_lk(ll) == ")") {
					ll = _ladv(ll)
				}
				nt := _lnew_str_temp(ll)
				_ = nt
				ll = nt.L
				ll = _lemit(ll, new_inst(IrOpOpCall{}, nt.Temp, fd_arg.Temp, int64(2), "_aria_tcp_read", int64(12)))
				ll = _lset_type(ll, nt.Temp, int64(12))
				return LT{L: ll, Temp: nt.Temp}
			}
			if (name == "_ariaTcpWrite") {
				ll = _ladv(ll)
				fd_arg := _lower_expr(ll)
				_ = fd_arg
				ll = fd_arg.L
				if (_lk(ll) == ",") {
					ll = _ladv(ll)
				}
				data_arg := _lower_expr(ll)
				_ = data_arg
				ll = data_arg.L
				if (_lk(ll) == ")") {
					ll = _ladv(ll)
				}
				nt := _lnew_temp(ll)
				_ = nt
				ll = nt.L
				ll = _lemit(ll, new_inst(IrOpOpCall{}, nt.Temp, fd_arg.Temp, int64(3), "_aria_tcp_write", int64(1)))
				ll = _lset_type(ll, nt.Temp, int64(1))
				return LT{L: ll, Temp: nt.Temp}
			}
			if (name == "_ariaTcpClose") {
				ll = _ladv(ll)
				fd_arg := _lower_expr(ll)
				_ = fd_arg
				ll = fd_arg.L
				if (_lk(ll) == ")") {
					ll = _ladv(ll)
				}
				nt := _lnew_temp(ll)
				_ = nt
				ll = nt.L
				ll = _lemit(ll, new_inst(IrOpOpCall{}, nt.Temp, fd_arg.Temp, int64(1), "_aria_tcp_close", int64(0)))
				return LT{L: ll, Temp: nt.Temp}
			}
			if (name == "_ariaTcpPeerAddr") {
				ll = _ladv(ll)
				fd_arg := _lower_expr(ll)
				_ = fd_arg
				ll = fd_arg.L
				if (_lk(ll) == ")") {
					ll = _ladv(ll)
				}
				nt := _lnew_str_temp(ll)
				_ = nt
				ll = nt.L
				ll = _lemit(ll, new_inst(IrOpOpCall{}, nt.Temp, fd_arg.Temp, int64(1), "_aria_tcp_peer_addr", int64(12)))
				ll = _lset_type(ll, nt.Temp, int64(12))
				return LT{L: ll, Temp: nt.Temp}
			}
			if (name == "_ariaTcpSetTimeout") {
				ll = _ladv(ll)
				fd_arg := _lower_expr(ll)
				_ = fd_arg
				ll = fd_arg.L
				if (_lk(ll) == ",") {
					ll = _ladv(ll)
				}
				kind_arg := _lower_expr(ll)
				_ = kind_arg
				ll = kind_arg.L
				if (_lk(ll) == ",") {
					ll = _ladv(ll)
				}
				ms_arg := _lower_expr(ll)
				_ = ms_arg
				ll = ms_arg.L
				if (_lk(ll) == ")") {
					ll = _ladv(ll)
				}
				nt := _lnew_temp(ll)
				_ = nt
				ll = nt.L
				ll = _lemit(ll, new_inst(IrOpOpCall{}, nt.Temp, fd_arg.Temp, int64(3), "_aria_tcp_set_timeout", int64(1)))
				ll = _lset_type(ll, nt.Temp, int64(1))
				return LT{L: ll, Temp: nt.Temp}
			}
			if (name == "_ariaPgConnect") {
				ll = _ladv(ll)
				arg := _lower_expr(ll)
				_ = arg
				ll = arg.L
				if (_lk(ll) == ")") {
					ll = _ladv(ll)
				}
				nt := _lnew_temp(ll)
				_ = nt
				ll = nt.L
				ll = _lemit(ll, new_inst(IrOpOpCall{}, nt.Temp, arg.Temp, int64(2), "_aria_pg_connect", int64(1)))
				ll = _lset_type(ll, nt.Temp, int64(1))
				return LT{L: ll, Temp: nt.Temp}
			}
			if (name == "_ariaPgClose") {
				ll = _ladv(ll)
				arg := _lower_expr(ll)
				_ = arg
				ll = arg.L
				if (_lk(ll) == ")") {
					ll = _ladv(ll)
				}
				nt := _lnew_temp(ll)
				_ = nt
				ll = nt.L
				ll = _lemit(ll, new_inst(IrOpOpCall{}, nt.Temp, arg.Temp, int64(1), "_aria_pg_close", int64(0)))
				return LT{L: ll, Temp: nt.Temp}
			}
			if (name == "_ariaPgStatus") {
				ll = _ladv(ll)
				arg := _lower_expr(ll)
				_ = arg
				ll = arg.L
				if (_lk(ll) == ")") {
					ll = _ladv(ll)
				}
				nt := _lnew_temp(ll)
				_ = nt
				ll = nt.L
				ll = _lemit(ll, new_inst(IrOpOpCall{}, nt.Temp, arg.Temp, int64(1), "_aria_pg_status", int64(1)))
				ll = _lset_type(ll, nt.Temp, int64(1))
				return LT{L: ll, Temp: nt.Temp}
			}
			if (name == "_ariaPgError") {
				ll = _ladv(ll)
				arg := _lower_expr(ll)
				_ = arg
				ll = arg.L
				if (_lk(ll) == ")") {
					ll = _ladv(ll)
				}
				nt := _lnew_str_temp(ll)
				_ = nt
				ll = nt.L
				ll = _lemit(ll, new_inst(IrOpOpCall{}, nt.Temp, arg.Temp, int64(1), "_aria_pg_error", int64(12)))
				ll = _lset_type(ll, nt.Temp, int64(12))
				return LT{L: ll, Temp: nt.Temp}
			}
			if (name == "_ariaPgExec") {
				ll = _ladv(ll)
				conn_arg := _lower_expr(ll)
				_ = conn_arg
				ll = conn_arg.L
				if (_lk(ll) == ",") {
					ll = _ladv(ll)
				}
				query_arg := _lower_expr(ll)
				_ = query_arg
				ll = query_arg.L
				if (_lk(ll) == ")") {
					ll = _ladv(ll)
				}
				nt := _lnew_temp(ll)
				_ = nt
				ll = nt.L
				ll = _lemit(ll, new_inst(IrOpOpCall{}, nt.Temp, conn_arg.Temp, int64(3), "_aria_pg_exec", int64(1)))
				ll = _lset_type(ll, nt.Temp, int64(1))
				return LT{L: ll, Temp: nt.Temp}
			}
			if (name == "_ariaPgExecParams") {
				ll = _ladv(ll)
				conn_arg := _lower_expr(ll)
				_ = conn_arg
				ll = conn_arg.L
				if (_lk(ll) == ",") {
					ll = _ladv(ll)
				}
				query_arg := _lower_expr(ll)
				_ = query_arg
				ll = query_arg.L
				if (_lk(ll) == ",") {
					ll = _ladv(ll)
				}
				params_arg := _lower_expr(ll)
				_ = params_arg
				ll = params_arg.L
				if (_lk(ll) == ")") {
					ll = _ladv(ll)
				}
				nt := _lnew_temp(ll)
				_ = nt
				ll = nt.L
				ll = _lemit(ll, new_inst(IrOpOpCall{}, nt.Temp, conn_arg.Temp, int64(4), "_aria_pg_exec_params", int64(1)))
				ll = _lset_type(ll, nt.Temp, int64(1))
				return LT{L: ll, Temp: nt.Temp}
			}
			if (name == "_ariaPgResultStatus") {
				ll = _ladv(ll)
				arg := _lower_expr(ll)
				_ = arg
				ll = arg.L
				if (_lk(ll) == ")") {
					ll = _ladv(ll)
				}
				nt := _lnew_temp(ll)
				_ = nt
				ll = nt.L
				ll = _lemit(ll, new_inst(IrOpOpCall{}, nt.Temp, arg.Temp, int64(1), "_aria_pg_result_status", int64(1)))
				ll = _lset_type(ll, nt.Temp, int64(1))
				return LT{L: ll, Temp: nt.Temp}
			}
			if (name == "_ariaPgResultError") {
				ll = _ladv(ll)
				arg := _lower_expr(ll)
				_ = arg
				ll = arg.L
				if (_lk(ll) == ")") {
					ll = _ladv(ll)
				}
				nt := _lnew_str_temp(ll)
				_ = nt
				ll = nt.L
				ll = _lemit(ll, new_inst(IrOpOpCall{}, nt.Temp, arg.Temp, int64(1), "_aria_pg_result_error", int64(12)))
				ll = _lset_type(ll, nt.Temp, int64(12))
				return LT{L: ll, Temp: nt.Temp}
			}
			if (name == "_ariaPgNrows") {
				ll = _ladv(ll)
				arg := _lower_expr(ll)
				_ = arg
				ll = arg.L
				if (_lk(ll) == ")") {
					ll = _ladv(ll)
				}
				nt := _lnew_temp(ll)
				_ = nt
				ll = nt.L
				ll = _lemit(ll, new_inst(IrOpOpCall{}, nt.Temp, arg.Temp, int64(1), "_aria_pg_nrows", int64(1)))
				ll = _lset_type(ll, nt.Temp, int64(1))
				return LT{L: ll, Temp: nt.Temp}
			}
			if (name == "_ariaPgNcols") {
				ll = _ladv(ll)
				arg := _lower_expr(ll)
				_ = arg
				ll = arg.L
				if (_lk(ll) == ")") {
					ll = _ladv(ll)
				}
				nt := _lnew_temp(ll)
				_ = nt
				ll = nt.L
				ll = _lemit(ll, new_inst(IrOpOpCall{}, nt.Temp, arg.Temp, int64(1), "_aria_pg_ncols", int64(1)))
				ll = _lset_type(ll, nt.Temp, int64(1))
				return LT{L: ll, Temp: nt.Temp}
			}
			if (name == "_ariaPgFieldName") {
				ll = _ladv(ll)
				res_arg := _lower_expr(ll)
				_ = res_arg
				ll = res_arg.L
				if (_lk(ll) == ",") {
					ll = _ladv(ll)
				}
				col_arg := _lower_expr(ll)
				_ = col_arg
				ll = col_arg.L
				if (_lk(ll) == ")") {
					ll = _ladv(ll)
				}
				nt := _lnew_str_temp(ll)
				_ = nt
				ll = nt.L
				ll = _lemit(ll, new_inst(IrOpOpCall{}, nt.Temp, res_arg.Temp, int64(2), "_aria_pg_field_name", int64(12)))
				ll = _lset_type(ll, nt.Temp, int64(12))
				return LT{L: ll, Temp: nt.Temp}
			}
			if (name == "_ariaPgGetValue") {
				ll = _ladv(ll)
				res_arg := _lower_expr(ll)
				_ = res_arg
				ll = res_arg.L
				if (_lk(ll) == ",") {
					ll = _ladv(ll)
				}
				row_arg := _lower_expr(ll)
				_ = row_arg
				ll = row_arg.L
				if (_lk(ll) == ",") {
					ll = _ladv(ll)
				}
				col_arg := _lower_expr(ll)
				_ = col_arg
				ll = col_arg.L
				if (_lk(ll) == ")") {
					ll = _ladv(ll)
				}
				nt := _lnew_str_temp(ll)
				_ = nt
				ll = nt.L
				ll = _lemit(ll, new_inst(IrOpOpCall{}, nt.Temp, res_arg.Temp, int64(3), "_aria_pg_get_value", int64(12)))
				ll = _lset_type(ll, nt.Temp, int64(12))
				return LT{L: ll, Temp: nt.Temp}
			}
			if (name == "_ariaPgIsNull") {
				ll = _ladv(ll)
				res_arg := _lower_expr(ll)
				_ = res_arg
				ll = res_arg.L
				if (_lk(ll) == ",") {
					ll = _ladv(ll)
				}
				row_arg := _lower_expr(ll)
				_ = row_arg
				ll = row_arg.L
				if (_lk(ll) == ",") {
					ll = _ladv(ll)
				}
				col_arg := _lower_expr(ll)
				_ = col_arg
				ll = col_arg.L
				if (_lk(ll) == ")") {
					ll = _ladv(ll)
				}
				nt := _lnew_temp(ll)
				_ = nt
				ll = nt.L
				ll = _lemit(ll, new_inst(IrOpOpCall{}, nt.Temp, res_arg.Temp, int64(3), "_aria_pg_is_null", int64(1)))
				ll = _lset_type(ll, nt.Temp, int64(1))
				return LT{L: ll, Temp: nt.Temp}
			}
			if (name == "_ariaPgClear") {
				ll = _ladv(ll)
				arg := _lower_expr(ll)
				_ = arg
				ll = arg.L
				if (_lk(ll) == ")") {
					ll = _ladv(ll)
				}
				nt := _lnew_temp(ll)
				_ = nt
				ll = nt.L
				ll = _lemit(ll, new_inst(IrOpOpCall{}, nt.Temp, arg.Temp, int64(1), "_aria_pg_clear", int64(0)))
				return LT{L: ll, Temp: nt.Temp}
			}
			if (name == "_ariaSpawn") {
				ll = _ladv(ll)
				arg := _lower_expr(ll)
				_ = arg
				ll = arg.L
				if (_lk(ll) == ")") {
					ll = _ladv(ll)
				}
				fn_ptr_nt := _lnew_temp(ll)
				_ = fn_ptr_nt
				ll = fn_ptr_nt.L
				ll = _lemit(ll, new_inst(IrOpOpFieldGet{}, fn_ptr_nt.Temp, arg.Temp, int64(0), "_fn", int64(0)))
				env_ptr_nt := _lnew_temp(ll)
				_ = env_ptr_nt
				ll = env_ptr_nt.L
				ll = _lemit(ll, new_inst(IrOpOpFieldGet{}, env_ptr_nt.Temp, arg.Temp, int64(1), "_env", int64(0)))
				nt := _lnew_temp(ll)
				_ = nt
				ll = nt.L
				ll = _lemit(ll, new_inst(IrOpOpCall{}, nt.Temp, fn_ptr_nt.Temp, int64(2), "_aria_spawn", int64(1)))
				ll = _lset_type(ll, nt.Temp, int64(1))
				return LT{L: ll, Temp: nt.Temp}
			}
			if (name == "_ariaTaskAwait") {
				ll = _ladv(ll)
				arg := _lower_expr(ll)
				_ = arg
				ll = arg.L
				if (_lk(ll) == ")") {
					ll = _ladv(ll)
				}
				nt := _lnew_temp(ll)
				_ = nt
				ll = nt.L
				ll = _lemit(ll, new_inst(IrOpOpCall{}, nt.Temp, arg.Temp, int64(1), "_aria_task_await", int64(1)))
				ll = _lset_type(ll, nt.Temp, int64(1))
				return LT{L: ll, Temp: nt.Temp}
			}
			if (name == "_ariaChanNew") {
				ll = _ladv(ll)
				arg := _lower_expr(ll)
				_ = arg
				ll = arg.L
				if (_lk(ll) == ")") {
					ll = _ladv(ll)
				}
				nt := _lnew_temp(ll)
				_ = nt
				ll = nt.L
				ll = _lemit(ll, new_inst(IrOpOpCall{}, nt.Temp, arg.Temp, int64(1), "_aria_chan_new", int64(1)))
				ll = _lset_type(ll, nt.Temp, int64(17))
				return LT{L: ll, Temp: nt.Temp}
			}
			if (name == "_ariaChanSend") {
				ll = _ladv(ll)
				ch_arg := _lower_expr(ll)
				_ = ch_arg
				ll = ch_arg.L
				if (_lk(ll) == ",") {
					ll = _ladv(ll)
				}
				val_arg := _lower_expr(ll)
				_ = val_arg
				ll = val_arg.L
				if (_lk(ll) == ")") {
					ll = _ladv(ll)
				}
				pack0 := _lnew_temp(ll)
				_ = pack0
				ll = pack0.L
				ll = _lemit(ll, new_inst(IrOpOpStore{}, pack0.Temp, ch_arg.Temp, int64(0), "", int64(0)))
				pack1 := _lnew_temp(ll)
				_ = pack1
				ll = pack1.L
				ll = _lemit(ll, new_inst(IrOpOpStore{}, pack1.Temp, val_arg.Temp, int64(0), "", int64(0)))
				nt := _lnew_temp(ll)
				_ = nt
				ll = nt.L
				ll = _lemit(ll, new_inst(IrOpOpCall{}, nt.Temp, pack0.Temp, int64(2), "_aria_chan_send", int64(1)))
				ll = _lset_type(ll, nt.Temp, int64(1))
				return LT{L: ll, Temp: nt.Temp}
			}
			if (name == "_ariaChanRecv") {
				ll = _ladv(ll)
				arg := _lower_expr(ll)
				_ = arg
				ll = arg.L
				if (_lk(ll) == ")") {
					ll = _ladv(ll)
				}
				nt := _lnew_temp(ll)
				_ = nt
				ll = nt.L
				ll = _lemit(ll, new_inst(IrOpOpCall{}, nt.Temp, arg.Temp, int64(1), "_aria_chan_recv", int64(1)))
				ll = _lset_type(ll, nt.Temp, int64(1))
				return LT{L: ll, Temp: nt.Temp}
			}
			if (name == "_ariaChanClose") {
				ll = _ladv(ll)
				arg := _lower_expr(ll)
				_ = arg
				ll = arg.L
				if (_lk(ll) == ")") {
					ll = _ladv(ll)
				}
				nt := _lnew_temp(ll)
				_ = nt
				ll = nt.L
				ll = _lemit(ll, new_inst(IrOpOpCall{}, nt.Temp, arg.Temp, int64(1), "_aria_chan_close", int64(0)))
				return LT{L: ll, Temp: nt.Temp}
			}
			if (name == "_ariaMutexNew") {
				ll = _ladv(ll)
				for ((_lk(ll) != ")") && (_lk(ll) != "EOF")) {
					ll = _ladv(ll)
				}
				if (_lk(ll) == ")") {
					ll = _ladv(ll)
				}
				nt := _lnew_temp(ll)
				_ = nt
				ll = nt.L
				ll = _lemit(ll, new_inst(IrOpOpCall{}, nt.Temp, int64(0), int64(0), "_aria_mutex_new", int64(1)))
				ll = _lset_type(ll, nt.Temp, int64(1))
				return LT{L: ll, Temp: nt.Temp}
			}
			if (name == "_ariaMutexLock") {
				ll = _ladv(ll)
				arg := _lower_expr(ll)
				_ = arg
				ll = arg.L
				if (_lk(ll) == ")") {
					ll = _ladv(ll)
				}
				nt := _lnew_temp(ll)
				_ = nt
				ll = nt.L
				ll = _lemit(ll, new_inst(IrOpOpCall{}, nt.Temp, arg.Temp, int64(1), "_aria_mutex_lock", int64(0)))
				return LT{L: ll, Temp: nt.Temp}
			}
			if (name == "_ariaMutexUnlock") {
				ll = _ladv(ll)
				arg := _lower_expr(ll)
				_ = arg
				ll = arg.L
				if (_lk(ll) == ")") {
					ll = _ladv(ll)
				}
				nt := _lnew_temp(ll)
				_ = nt
				ll = nt.L
				ll = _lemit(ll, new_inst(IrOpOpCall{}, nt.Temp, arg.Temp, int64(1), "_aria_mutex_unlock", int64(0)))
				return LT{L: ll, Temp: nt.Temp}
			}
			if (name == "_ariaSpawn2") {
				ll = _ladv(ll)
				arg := _lower_expr(ll)
				_ = arg
				ll = arg.L
				if (_lk(ll) == ")") {
					ll = _ladv(ll)
				}
				fn_ptr_nt := _lnew_temp(ll)
				_ = fn_ptr_nt
				ll = fn_ptr_nt.L
				ll = _lemit(ll, new_inst(IrOpOpFieldGet{}, fn_ptr_nt.Temp, arg.Temp, int64(0), "_fn", int64(0)))
				env_ptr_nt := _lnew_temp(ll)
				_ = env_ptr_nt
				ll = env_ptr_nt.L
				ll = _lemit(ll, new_inst(IrOpOpFieldGet{}, env_ptr_nt.Temp, arg.Temp, int64(1), "_env", int64(0)))
				nt := _lnew_temp(ll)
				_ = nt
				ll = nt.L
				ll = _lemit(ll, new_inst(IrOpOpCall{}, nt.Temp, fn_ptr_nt.Temp, int64(2), "_aria_spawn2", int64(1)))
				ll = _lset_type(ll, nt.Temp, int64(1))
				return LT{L: ll, Temp: nt.Temp}
			}
			if ((((name == "_ariaTaskAwait2") || (name == "_ariaTaskDone")) || (name == "_ariaTaskResult")) || (name == "_ariaCancelCheck")) {
				ll = _ladv(ll)
				arg := _lower_expr(ll)
				_ = arg
				ll = arg.L
				if (_lk(ll) == ")") {
					ll = _ladv(ll)
				}
				nt := _lnew_temp(ll)
				_ = nt
				ll = nt.L
				rt_name := "_aria_task_await2"
				_ = rt_name
				if (name == "_ariaTaskDone") {
					rt_name = "_aria_task_done"
				}
				if (name == "_ariaTaskResult") {
					rt_name = "_aria_task_result"
				}
				if (name == "_ariaCancelCheck") {
					rt_name = "_aria_cancel_check"
				}
				ll = _lemit(ll, new_inst(IrOpOpCall{}, nt.Temp, arg.Temp, int64(1), rt_name, int64(1)))
				ll = _lset_type(ll, nt.Temp, int64(1))
				return LT{L: ll, Temp: nt.Temp}
			}
			if (name == "_ariaTaskCancel") {
				ll = _ladv(ll)
				arg := _lower_expr(ll)
				_ = arg
				ll = arg.L
				if (_lk(ll) == ")") {
					ll = _ladv(ll)
				}
				nt := _lnew_temp(ll)
				_ = nt
				ll = nt.L
				ll = _lemit(ll, new_inst(IrOpOpCall{}, nt.Temp, arg.Temp, int64(1), "_aria_task_cancel", int64(0)))
				return LT{L: ll, Temp: nt.Temp}
			}
			if (name == "_ariaChanTryRecv") {
				ll = _ladv(ll)
				arg := _lower_expr(ll)
				_ = arg
				ll = arg.L
				if (_lk(ll) == ")") {
					ll = _ladv(ll)
				}
				nt := _lnew_str_temp(ll)
				_ = nt
				ll = nt.L
				ll = _lemit(ll, new_inst(IrOpOpCall{}, nt.Temp, arg.Temp, int64(1), "_aria_chan_try_recv", int64(0)))
				return LT{L: ll, Temp: nt.Temp}
			}
			if (name == "_ariaChanSelect") {
				ll = _ladv(ll)
				arr_arg := _lower_expr(ll)
				_ = arr_arg
				ll = arr_arg.L
				if (_lk(ll) == ",") {
					ll = _ladv(ll)
				}
				to_arg := _lower_expr(ll)
				_ = to_arg
				ll = to_arg.L
				if (_lk(ll) == ")") {
					ll = _ladv(ll)
				}
				pack0 := _lnew_temp(ll)
				_ = pack0
				ll = pack0.L
				ll = _lemit(ll, new_inst(IrOpOpStore{}, pack0.Temp, arr_arg.Temp, int64(0), "", int64(0)))
				pack1 := _lnew_temp(ll)
				_ = pack1
				ll = pack1.L
				ll = _lemit(ll, new_inst(IrOpOpStore{}, pack1.Temp, to_arg.Temp, int64(0), "", int64(0)))
				nt := _lnew_str_temp(ll)
				_ = nt
				ll = nt.L
				ll = _lemit(ll, new_inst(IrOpOpCall{}, nt.Temp, pack0.Temp, int64(2), "_aria_chan_select", int64(0)))
				return LT{L: ll, Temp: nt.Temp}
			}
			if (name == "_ariaRWMutexNew") {
				ll = _ladv(ll)
				for ((_lk(ll) != ")") && (_lk(ll) != "EOF")) {
					ll = _ladv(ll)
				}
				if (_lk(ll) == ")") {
					ll = _ladv(ll)
				}
				nt := _lnew_temp(ll)
				_ = nt
				ll = nt.L
				ll = _lemit(ll, new_inst(IrOpOpCall{}, nt.Temp, int64(0), int64(0), "_aria_rwmutex_new", int64(1)))
				ll = _lset_type(ll, nt.Temp, int64(1))
				return LT{L: ll, Temp: nt.Temp}
			}
			if ((((name == "_ariaRWMutexRlock") || (name == "_ariaRWMutexRunlock")) || (name == "_ariaRWMutexWlock")) || (name == "_ariaRWMutexWunlock")) {
				ll = _ladv(ll)
				arg := _lower_expr(ll)
				_ = arg
				ll = arg.L
				if (_lk(ll) == ")") {
					ll = _ladv(ll)
				}
				nt := _lnew_temp(ll)
				_ = nt
				ll = nt.L
				rn := "_aria_rwmutex_rlock"
				_ = rn
				if (name == "_ariaRWMutexRunlock") {
					rn = "_aria_rwmutex_runlock"
				}
				if (name == "_ariaRWMutexWlock") {
					rn = "_aria_rwmutex_wlock"
				}
				if (name == "_ariaRWMutexWunlock") {
					rn = "_aria_rwmutex_wunlock"
				}
				ll = _lemit(ll, new_inst(IrOpOpCall{}, nt.Temp, arg.Temp, int64(1), rn, int64(0)))
				return LT{L: ll, Temp: nt.Temp}
			}
			if (name == "_ariaWgNew") {
				ll = _ladv(ll)
				for ((_lk(ll) != ")") && (_lk(ll) != "EOF")) {
					ll = _ladv(ll)
				}
				if (_lk(ll) == ")") {
					ll = _ladv(ll)
				}
				nt := _lnew_temp(ll)
				_ = nt
				ll = nt.L
				ll = _lemit(ll, new_inst(IrOpOpCall{}, nt.Temp, int64(0), int64(0), "_aria_wg_new", int64(1)))
				ll = _lset_type(ll, nt.Temp, int64(1))
				return LT{L: ll, Temp: nt.Temp}
			}
			if (name == "_ariaWgAdd") {
				ll = _ladv(ll)
				h_arg := _lower_expr(ll)
				_ = h_arg
				ll = h_arg.L
				if (_lk(ll) == ",") {
					ll = _ladv(ll)
				}
				d_arg := _lower_expr(ll)
				_ = d_arg
				ll = d_arg.L
				if (_lk(ll) == ")") {
					ll = _ladv(ll)
				}
				pack0 := _lnew_temp(ll)
				_ = pack0
				ll = pack0.L
				ll = _lemit(ll, new_inst(IrOpOpStore{}, pack0.Temp, h_arg.Temp, int64(0), "", int64(0)))
				pack1 := _lnew_temp(ll)
				_ = pack1
				ll = pack1.L
				ll = _lemit(ll, new_inst(IrOpOpStore{}, pack1.Temp, d_arg.Temp, int64(0), "", int64(0)))
				nt := _lnew_temp(ll)
				_ = nt
				ll = nt.L
				ll = _lemit(ll, new_inst(IrOpOpCall{}, nt.Temp, pack0.Temp, int64(2), "_aria_wg_add", int64(0)))
				return LT{L: ll, Temp: nt.Temp}
			}
			if ((name == "_ariaWgDone") || (name == "_ariaWgWait")) {
				ll = _ladv(ll)
				arg := _lower_expr(ll)
				_ = arg
				ll = arg.L
				if (_lk(ll) == ")") {
					ll = _ladv(ll)
				}
				nt := _lnew_temp(ll)
				_ = nt
				ll = nt.L
				wn := "_aria_wg_done"
				_ = wn
				if (name == "_ariaWgWait") {
					wn = "_aria_wg_wait"
				}
				ll = _lemit(ll, new_inst(IrOpOpCall{}, nt.Temp, arg.Temp, int64(1), wn, int64(0)))
				return LT{L: ll, Temp: nt.Temp}
			}
			if (name == "_ariaOnceNew") {
				ll = _ladv(ll)
				for ((_lk(ll) != ")") && (_lk(ll) != "EOF")) {
					ll = _ladv(ll)
				}
				if (_lk(ll) == ")") {
					ll = _ladv(ll)
				}
				nt := _lnew_temp(ll)
				_ = nt
				ll = nt.L
				ll = _lemit(ll, new_inst(IrOpOpCall{}, nt.Temp, int64(0), int64(0), "_aria_once_new", int64(1)))
				ll = _lset_type(ll, nt.Temp, int64(1))
				return LT{L: ll, Temp: nt.Temp}
			}
			if (name == "_ariaOnceCall") {
				ll = _ladv(ll)
				h_arg := _lower_expr(ll)
				_ = h_arg
				ll = h_arg.L
				if (_lk(ll) == ",") {
					ll = _ladv(ll)
				}
				fn_arg := _lower_expr(ll)
				_ = fn_arg
				ll = fn_arg.L
				if (_lk(ll) == ")") {
					ll = _ladv(ll)
				}
				fn_ptr := _lnew_temp(ll)
				_ = fn_ptr
				ll = fn_ptr.L
				ll = _lemit(ll, new_inst(IrOpOpFieldGet{}, fn_ptr.Temp, fn_arg.Temp, int64(0), "_fn", int64(0)))
				env_ptr := _lnew_temp(ll)
				_ = env_ptr
				ll = env_ptr.L
				ll = _lemit(ll, new_inst(IrOpOpFieldGet{}, env_ptr.Temp, fn_arg.Temp, int64(1), "_env", int64(0)))
				p0 := _lnew_temp(ll)
				_ = p0
				ll = p0.L
				ll = _lemit(ll, new_inst(IrOpOpStore{}, p0.Temp, h_arg.Temp, int64(0), "", int64(0)))
				p1 := _lnew_temp(ll)
				_ = p1
				ll = p1.L
				ll = _lemit(ll, new_inst(IrOpOpStore{}, p1.Temp, fn_ptr.Temp, int64(0), "", int64(0)))
				p2 := _lnew_temp(ll)
				_ = p2
				ll = p2.L
				ll = _lemit(ll, new_inst(IrOpOpStore{}, p2.Temp, env_ptr.Temp, int64(0), "", int64(0)))
				nt := _lnew_temp(ll)
				_ = nt
				ll = nt.L
				ll = _lemit(ll, new_inst(IrOpOpCall{}, nt.Temp, p0.Temp, int64(3), "_aria_once_call", int64(1)))
				ll = _lset_type(ll, nt.Temp, int64(1))
				return LT{L: ll, Temp: nt.Temp}
			}
			if (name == "_ariaSbNew") {
				ll = _ladv(ll)
				for ((_lk(ll) != ")") && (_lk(ll) != "EOF")) {
					ll = _ladv(ll)
				}
				if (_lk(ll) == ")") {
					ll = _ladv(ll)
				}
				nt := _lnew_temp(ll)
				_ = nt
				ll = nt.L
				ll = _lemit(ll, new_inst(IrOpOpCall{}, nt.Temp, int64(0), int64(0), "_aria_sb_new", int64(1)))
				ll = _lset_type(ll, nt.Temp, int64(1))
				return LT{L: ll, Temp: nt.Temp}
			}
			if (name == "_ariaSbWithCapacity") {
				ll = _ladv(ll)
				arg := _lower_expr(ll)
				_ = arg
				ll = arg.L
				if (_lk(ll) == ")") {
					ll = _ladv(ll)
				}
				nt := _lnew_temp(ll)
				_ = nt
				ll = nt.L
				ll = _lemit(ll, new_inst(IrOpOpCall{}, nt.Temp, arg.Temp, int64(1), "_aria_sb_with_capacity", int64(1)))
				ll = _lset_type(ll, nt.Temp, int64(1))
				return LT{L: ll, Temp: nt.Temp}
			}
			if (name == "_ariaSbAppend") {
				ll = _ladv(ll)
				h_arg := _lower_expr(ll)
				_ = h_arg
				ll = h_arg.L
				if (_lk(ll) == ",") {
					ll = _ladv(ll)
				}
				s_arg := _lower_expr(ll)
				_ = s_arg
				ll = s_arg.L
				if (_lk(ll) == ")") {
					ll = _ladv(ll)
				}
				p0 := _lnew_temp(ll)
				_ = p0
				ll = p0.L
				ll = _lemit(ll, new_inst(IrOpOpStore{}, p0.Temp, h_arg.Temp, int64(0), "", int64(0)))
				p1 := _lnew_temp(ll)
				_ = p1
				ll = p1.L
				ll = _lemit(ll, new_inst(IrOpOpStore{}, p1.Temp, s_arg.Temp, int64(0), "", int64(0)))
				p2 := _lnew_temp(ll)
				_ = p2
				ll = p2.L
				ll = _lemit(ll, new_inst(IrOpOpStore{}, p2.Temp, (s_arg.Temp + int64(1)), int64(0), "", int64(0)))
				nt := _lnew_temp(ll)
				_ = nt
				ll = nt.L
				ll = _lemit(ll, new_inst(IrOpOpCall{}, nt.Temp, p0.Temp, int64(3), "_aria_sb_append", int64(0)))
				return LT{L: ll, Temp: nt.Temp}
			}
			if (name == "_ariaSbAppendChar") {
				ll = _ladv(ll)
				h_arg := _lower_expr(ll)
				_ = h_arg
				ll = h_arg.L
				if (_lk(ll) == ",") {
					ll = _ladv(ll)
				}
				c_arg := _lower_expr(ll)
				_ = c_arg
				ll = c_arg.L
				if (_lk(ll) == ")") {
					ll = _ladv(ll)
				}
				p0 := _lnew_temp(ll)
				_ = p0
				ll = p0.L
				ll = _lemit(ll, new_inst(IrOpOpStore{}, p0.Temp, h_arg.Temp, int64(0), "", int64(0)))
				p1 := _lnew_temp(ll)
				_ = p1
				ll = p1.L
				ll = _lemit(ll, new_inst(IrOpOpStore{}, p1.Temp, c_arg.Temp, int64(0), "", int64(0)))
				nt := _lnew_temp(ll)
				_ = nt
				ll = nt.L
				ll = _lemit(ll, new_inst(IrOpOpCall{}, nt.Temp, p0.Temp, int64(2), "_aria_sb_append_char", int64(0)))
				return LT{L: ll, Temp: nt.Temp}
			}
			if (name == "_ariaSbLen") {
				ll = _ladv(ll)
				arg := _lower_expr(ll)
				_ = arg
				ll = arg.L
				if (_lk(ll) == ")") {
					ll = _ladv(ll)
				}
				nt := _lnew_temp(ll)
				_ = nt
				ll = nt.L
				ll = _lemit(ll, new_inst(IrOpOpCall{}, nt.Temp, arg.Temp, int64(1), "_aria_sb_len", int64(1)))
				ll = _lset_type(ll, nt.Temp, int64(1))
				return LT{L: ll, Temp: nt.Temp}
			}
			if (name == "_ariaSbBuild") {
				ll = _ladv(ll)
				arg := _lower_expr(ll)
				_ = arg
				ll = arg.L
				if (_lk(ll) == ")") {
					ll = _ladv(ll)
				}
				nt := _lnew_str_temp(ll)
				_ = nt
				ll = nt.L
				ll = _lemit(ll, new_inst(IrOpOpCall{}, nt.Temp, arg.Temp, int64(1), "_aria_sb_build", int64(12)))
				ll = _lset_type(ll, nt.Temp, int64(12))
				return LT{L: ll, Temp: nt.Temp}
			}
			if (name == "_ariaSbClear") {
				ll = _ladv(ll)
				arg := _lower_expr(ll)
				_ = arg
				ll = arg.L
				if (_lk(ll) == ")") {
					ll = _ladv(ll)
				}
				nt := _lnew_temp(ll)
				_ = nt
				ll = nt.L
				ll = _lemit(ll, new_inst(IrOpOpCall{}, nt.Temp, arg.Temp, int64(1), "_aria_sb_clear", int64(0)))
				return LT{L: ll, Temp: nt.Temp}
			}
			if (name == "_ariaGcCollect") {
				ll = _ladv(ll)
				for ((_lk(ll) != ")") && (_lk(ll) != "EOF")) {
					ll = _ladv(ll)
				}
				if (_lk(ll) == ")") {
					ll = _ladv(ll)
				}
				nt := _lnew_temp(ll)
				_ = nt
				ll = nt.L
				ll = _lemit(ll, new_inst(IrOpOpCall{}, nt.Temp, int64(0), int64(0), "_aria_gc_collect", int64(1)))
				ll = _lset_type(ll, nt.Temp, int64(1))
				return LT{L: ll, Temp: nt.Temp}
			}
			if ((name == "_ariaGcTotalBytes") || (name == "_ariaGcAllocationCount")) {
				ll = _ladv(ll)
				for ((_lk(ll) != ")") && (_lk(ll) != "EOF")) {
					ll = _ladv(ll)
				}
				if (_lk(ll) == ")") {
					ll = _ladv(ll)
				}
				nt := _lnew_temp(ll)
				_ = nt
				ll = nt.L
				gc_fn := "_aria_gc_total_bytes"
				_ = gc_fn
				if (name == "_ariaGcAllocationCount") {
					gc_fn = "_aria_gc_allocation_count"
				}
				ll = _lemit(ll, new_inst(IrOpOpCall{}, nt.Temp, int64(0), int64(0), gc_fn, int64(1)))
				ll = _lset_type(ll, nt.Temp, int64(1))
				return LT{L: ll, Temp: nt.Temp}
			}
			if (name == "_ariaArenaNew") {
				ll = _ladv(ll)
				arg := _lower_expr(ll)
				_ = arg
				ll = arg.L
				if (_lk(ll) == ")") {
					ll = _ladv(ll)
				}
				nt := _lnew_temp(ll)
				_ = nt
				ll = nt.L
				ll = _lemit(ll, new_inst(IrOpOpCall{}, nt.Temp, arg.Temp, int64(1), "_aria_arena_new", int64(1)))
				ll = _lset_type(ll, nt.Temp, int64(1))
				return LT{L: ll, Temp: nt.Temp}
			}
			if (name == "_ariaArenaAlloc") {
				ll = _ladv(ll)
				h_arg := _lower_expr(ll)
				_ = h_arg
				ll = h_arg.L
				if (_lk(ll) == ",") {
					ll = _ladv(ll)
				}
				s_arg := _lower_expr(ll)
				_ = s_arg
				ll = s_arg.L
				if (_lk(ll) == ")") {
					ll = _ladv(ll)
				}
				p0 := _lnew_temp(ll)
				_ = p0
				ll = p0.L
				ll = _lemit(ll, new_inst(IrOpOpStore{}, p0.Temp, h_arg.Temp, int64(0), "", int64(0)))
				p1 := _lnew_temp(ll)
				_ = p1
				ll = p1.L
				ll = _lemit(ll, new_inst(IrOpOpStore{}, p1.Temp, s_arg.Temp, int64(0), "", int64(0)))
				nt := _lnew_temp(ll)
				_ = nt
				ll = nt.L
				ll = _lemit(ll, new_inst(IrOpOpCall{}, nt.Temp, p0.Temp, int64(2), "_aria_arena_alloc", int64(1)))
				ll = _lset_type(ll, nt.Temp, int64(1))
				return LT{L: ll, Temp: nt.Temp}
			}
			if ((name == "_ariaArenaReset") || (name == "_ariaArenaFree")) {
				ll = _ladv(ll)
				arg := _lower_expr(ll)
				_ = arg
				ll = arg.L
				if (_lk(ll) == ")") {
					ll = _ladv(ll)
				}
				nt := _lnew_temp(ll)
				_ = nt
				ll = nt.L
				af := "_aria_arena_reset"
				_ = af
				if (name == "_ariaArenaFree") {
					af = "_aria_arena_free"
				}
				ll = _lemit(ll, new_inst(IrOpOpCall{}, nt.Temp, arg.Temp, int64(1), af, int64(0)))
				return LT{L: ll, Temp: nt.Temp}
			}
			if ((name == "_ariaArenaAllocated") || (name == "_ariaArenaCapacity")) {
				ll = _ladv(ll)
				arg := _lower_expr(ll)
				_ = arg
				ll = arg.L
				if (_lk(ll) == ")") {
					ll = _ladv(ll)
				}
				nt := _lnew_temp(ll)
				_ = nt
				ll = nt.L
				af := "_aria_arena_allocated"
				_ = af
				if (name == "_ariaArenaCapacity") {
					af = "_aria_arena_capacity"
				}
				ll = _lemit(ll, new_inst(IrOpOpCall{}, nt.Temp, arg.Temp, int64(1), af, int64(1)))
				ll = _lset_type(ll, nt.Temp, int64(1))
				return LT{L: ll, Temp: nt.Temp}
			}
			if (name == "_ariaPoolNew") {
				ll = _ladv(ll)
				cap_arg := _lower_expr(ll)
				_ = cap_arg
				ll = cap_arg.L
				if (_lk(ll) == ",") {
					ll = _ladv(ll)
				}
				size_arg := _lower_expr(ll)
				_ = size_arg
				ll = size_arg.L
				if (_lk(ll) == ")") {
					ll = _ladv(ll)
				}
				p0 := _lnew_temp(ll)
				_ = p0
				ll = p0.L
				ll = _lemit(ll, new_inst(IrOpOpStore{}, p0.Temp, cap_arg.Temp, int64(0), "", int64(0)))
				p1 := _lnew_temp(ll)
				_ = p1
				ll = p1.L
				ll = _lemit(ll, new_inst(IrOpOpStore{}, p1.Temp, size_arg.Temp, int64(0), "", int64(0)))
				nt := _lnew_temp(ll)
				_ = nt
				ll = nt.L
				ll = _lemit(ll, new_inst(IrOpOpCall{}, nt.Temp, p0.Temp, int64(2), "_aria_pool_new", int64(1)))
				ll = _lset_type(ll, nt.Temp, int64(1))
				return LT{L: ll, Temp: nt.Temp}
			}
			if (name == "_ariaPoolGet") {
				ll = _ladv(ll)
				arg := _lower_expr(ll)
				_ = arg
				ll = arg.L
				if (_lk(ll) == ")") {
					ll = _ladv(ll)
				}
				nt := _lnew_temp(ll)
				_ = nt
				ll = nt.L
				ll = _lemit(ll, new_inst(IrOpOpCall{}, nt.Temp, arg.Temp, int64(1), "_aria_pool_get", int64(1)))
				ll = _lset_type(ll, nt.Temp, int64(1))
				return LT{L: ll, Temp: nt.Temp}
			}
			if (name == "_ariaPoolPut") {
				ll = _ladv(ll)
				h_arg := _lower_expr(ll)
				_ = h_arg
				ll = h_arg.L
				if (_lk(ll) == ",") {
					ll = _ladv(ll)
				}
				o_arg := _lower_expr(ll)
				_ = o_arg
				ll = o_arg.L
				if (_lk(ll) == ")") {
					ll = _ladv(ll)
				}
				p0 := _lnew_temp(ll)
				_ = p0
				ll = p0.L
				ll = _lemit(ll, new_inst(IrOpOpStore{}, p0.Temp, h_arg.Temp, int64(0), "", int64(0)))
				p1 := _lnew_temp(ll)
				_ = p1
				ll = p1.L
				ll = _lemit(ll, new_inst(IrOpOpStore{}, p1.Temp, o_arg.Temp, int64(0), "", int64(0)))
				nt := _lnew_temp(ll)
				_ = nt
				ll = nt.L
				ll = _lemit(ll, new_inst(IrOpOpCall{}, nt.Temp, p0.Temp, int64(2), "_aria_pool_put", int64(0)))
				return LT{L: ll, Temp: nt.Temp}
			}
			if (name == "_ariaCancelNew") {
				ll = _ladv(ll)
				for ((_lk(ll) != ")") && (_lk(ll) != "EOF")) {
					ll = _ladv(ll)
				}
				if (_lk(ll) == ")") {
					ll = _ladv(ll)
				}
				nt := _lnew_temp(ll)
				_ = nt
				ll = nt.L
				ll = _lemit(ll, new_inst(IrOpOpCall{}, nt.Temp, int64(0), int64(0), "_aria_cancel_new", int64(1)))
				ll = _lset_type(ll, nt.Temp, int64(1))
				return LT{L: ll, Temp: nt.Temp}
			}
			if (name == "_ariaCancelChild") {
				ll = _ladv(ll)
				arg := _lower_expr(ll)
				_ = arg
				ll = arg.L
				if (_lk(ll) == ")") {
					ll = _ladv(ll)
				}
				nt := _lnew_temp(ll)
				_ = nt
				ll = nt.L
				ll = _lemit(ll, new_inst(IrOpOpCall{}, nt.Temp, arg.Temp, int64(1), "_aria_cancel_child", int64(1)))
				ll = _lset_type(ll, nt.Temp, int64(1))
				return LT{L: ll, Temp: nt.Temp}
			}
			if (name == "_ariaCancelTrigger") {
				ll = _ladv(ll)
				arg := _lower_expr(ll)
				_ = arg
				ll = arg.L
				if (_lk(ll) == ")") {
					ll = _ladv(ll)
				}
				nt := _lnew_temp(ll)
				_ = nt
				ll = nt.L
				ll = _lemit(ll, new_inst(IrOpOpCall{}, nt.Temp, arg.Temp, int64(1), "_aria_cancel_trigger", int64(0)))
				return LT{L: ll, Temp: nt.Temp}
			}
			if (name == "_ariaCancelIsTriggered") {
				ll = _ladv(ll)
				arg := _lower_expr(ll)
				_ = arg
				ll = arg.L
				if (_lk(ll) == ")") {
					ll = _ladv(ll)
				}
				nt := _lnew_temp(ll)
				_ = nt
				ll = nt.L
				ll = _lemit(ll, new_inst(IrOpOpCall{}, nt.Temp, arg.Temp, int64(1), "_aria_cancel_is_triggered", int64(1)))
				ll = _lset_type(ll, nt.Temp, int64(1))
				return LT{L: ll, Temp: nt.Temp}
			}
			fn_slot := _lenv_lookup(ll, name)
			_ = fn_slot
			if (fn_slot >= int64(0)) {
				closure_nt := _lnew_temp(ll)
				_ = closure_nt
				ll = closure_nt.L
				ll = _lemit(ll, new_inst(IrOpOpLoad{}, closure_nt.Temp, fn_slot, int64(0), "", int64(0)))
				fn_ptr_nt := _lnew_temp(ll)
				_ = fn_ptr_nt
				ll = fn_ptr_nt.L
				ll = _lemit(ll, new_inst(IrOpOpFieldGet{}, fn_ptr_nt.Temp, closure_nt.Temp, int64(0), "_fn", int64(0)))
				env_ptr_nt := _lnew_temp(ll)
				_ = env_ptr_nt
				ll = env_ptr_nt.L
				ll = _lemit(ll, new_inst(IrOpOpFieldGet{}, env_ptr_nt.Temp, closure_nt.Temp, int64(1), "_env", int64(0)))
				ll = _ladv(ll)
				ll = _lskip_nl(ll)
				first_arg := -int64(1)
				_ = first_arg
				ic_arg_count := int64(0)
				_ = ic_arg_count
				ic_reg_count := int64(0)
				_ = ic_reg_count
				arg_is_str_0 := false
				_ = arg_is_str_0
				arg_is_str_1 := false
				_ = arg_is_str_1
				arg_is_str_2 := false
				_ = arg_is_str_2
				arg_is_str_3 := false
				_ = arg_is_str_3
				for ((_lk(ll) != ")") && (_lk(ll) != "EOF")) {
					ll = _lskip_nl(ll)
					arg := _lower_expr(ll)
					_ = arg
					ll = arg.L
					if (ic_arg_count == int64(0)) {
						first_arg = arg.Temp
					}
					atid := _lget_type(ll, arg.Temp)
					_ = atid
					if (atid == int64(12)) {
						if (ic_arg_count == int64(0)) {
							arg_is_str_0 = true
						}
						if (ic_arg_count == int64(1)) {
							arg_is_str_1 = true
						}
						if (ic_arg_count == int64(2)) {
							arg_is_str_2 = true
						}
						if (ic_arg_count == int64(3)) {
							arg_is_str_3 = true
						}
						ic_reg_count = (ic_reg_count + int64(2))
					} else {
						ic_reg_count = (ic_reg_count + int64(1))
					}
					ic_arg_count = (ic_arg_count + int64(1))
					ll = _lskip_nl(ll)
					if (_lk(ll) == ",") {
						ll = _ladv(ll)
					}
				}
				if (_lk(ll) == ")") {
					ll = _ladv(ll)
				}
				is_str_ret := false
				_ = is_str_ret
				fn_tid_pre := _lget_type(ll, fn_slot)
				_ = fn_tid_pre
				if (fn_tid_pre > int64(0)) {
					ti_pre := store_get(ll.Store, fn_tid_pre)
					_ = ti_pre
					if ((type_kind_name(ti_pre.Kind) == "Function") && (type_info_type_id(ti_pre) == int64(12))) {
						is_str_ret = true
					}
				}
				nt_temp := int64(0)
				_ = nt_temp
				if is_str_ret {
					nt_s := _lnew_str_temp(ll)
					_ = nt_s
					ll = nt_s.L
					nt_temp = nt_s.Temp
				} else {
					nt_r := _lnew_temp(ll)
					_ = nt_r
					ll = nt_r.L
					nt_temp = nt_r.Temp
				}
				fa := env_ptr_nt.Temp
				_ = fa
				pack_start := _lnew_temp(ll)
				_ = pack_start
				ll = pack_start.L
				ll = _lemit(ll, new_inst(IrOpOpStore{}, pack_start.Temp, env_ptr_nt.Temp, int64(0), "", int64(0)))
				pack_idx := int64(1)
				_ = pack_idx
				src_off := int64(0)
				_ = src_off
				if ((ic_arg_count > int64(0)) && (first_arg >= int64(0))) {
					ai := int64(0)
					_ = ai
					for (ai < ic_arg_count) {
						this_is_str := false
						_ = this_is_str
						if (ai == int64(0)) {
							this_is_str = arg_is_str_0
						}
						if (ai == int64(1)) {
							this_is_str = arg_is_str_1
						}
						if (ai == int64(2)) {
							this_is_str = arg_is_str_2
						}
						if (ai == int64(3)) {
							this_is_str = arg_is_str_3
						}
						pack_t := _lnew_temp(ll)
						_ = pack_t
						ll = pack_t.L
						ll = _lemit(ll, new_inst(IrOpOpLoad{}, pack_t.Temp, (first_arg + src_off), int64(0), "", int64(0)))
						ll = _lemit(ll, new_inst(IrOpOpStore{}, (pack_start.Temp + pack_idx), pack_t.Temp, int64(0), "", int64(0)))
						pack_idx = (pack_idx + int64(1))
						src_off = (src_off + int64(1))
						if this_is_str {
							pack_t2 := _lnew_temp(ll)
							_ = pack_t2
							ll = pack_t2.L
							ll = _lemit(ll, new_inst(IrOpOpLoad{}, pack_t2.Temp, (first_arg + src_off), int64(0), "", int64(0)))
							ll = _lemit(ll, new_inst(IrOpOpStore{}, (pack_start.Temp + pack_idx), pack_t2.Temp, int64(0), "", int64(0)))
							pack_idx = (pack_idx + int64(1))
							src_off = (src_off + int64(1))
						}
						ai = (ai + int64(1))
					}
				}
				total_args := (ic_reg_count + int64(1))
				_ = total_args
				ret_hint := ""
				_ = ret_hint
				fn_tid := _lget_type(ll, fn_slot)
				_ = fn_tid
				if (fn_tid > int64(0)) {
					ti := store_get(ll.Store, fn_tid)
					_ = ti
					if (type_kind_name(ti.Kind) == "Function") {
						ret_type := type_info_type_id(ti)
						_ = ret_type
						if (ret_type == int64(12)) {
							ret_hint = "str"
						}
						ll = _lset_type(ll, nt_temp, ret_type)
					}
				}
				ll = _lemit(ll, new_inst(IrOpOpCallIndirect{}, nt_temp, fn_ptr_nt.Temp, total_args, ret_hint, pack_start.Temp))
				return LT{L: ll, Temp: nt_temp}
			}
			ll = _ladv(ll)
			for ((_lk(ll) != ")") && (_lk(ll) != "EOF")) {
				ll = _ladv(ll)
			}
			if (_lk(ll) == ")") {
				ll = _ladv(ll)
			}
			nt := _lnew_temp(ll)
			_ = nt
			ll = nt.L
			return LT{L: ll, Temp: nt.Temp}
		}
		lb := "{"
		_ = lb
		if (_lk(ll) == "[") {
			sdef := reg_find_struct(ll.Registry, name)
			_ = sdef
			if ((sdef.Name != "") && struct_is_generic(sdef)) {
				gd := int64(1)
				_ = gd
				ll = _ladv(ll)
				gs_suffix := ""
				_ = gs_suffix
				for ((gd > int64(0)) && (_lk(ll) != "EOF")) {
					if (_lk(ll) == "[") {
						gd = (gd + int64(1))
					}
					if (_lk(ll) == "]") {
						gd = (gd - int64(1))
					}
					if (gd > int64(0)) {
						if ((_lk(ll) == "IDENT") && (gs_suffix == "")) {
							ta_name := _lcur(ll).Text
							_ = ta_name
							ta_id := resolve_type_name(ll.Store, ta_name)
							_ = ta_id
							gs_suffix = _type_name_suffix(ll.Store, ta_id)
						}
						ll = _ladv(ll)
					}
				}
				if (_lk(ll) == "]") {
					ll = _ladv(ll)
				}
				spec_sname := name
				_ = spec_sname
				if (gs_suffix != "") {
					spec_sname = ((name + "_") + gs_suffix)
				}
				if (_lk(ll) == lb) {
					return _lower_struct_lit(ll, spec_sname)
				}
			}
		}
		if (_lk(ll) == lb) {
			sdef := reg_find_struct(ll.Registry, name)
			_ = sdef
			if (sdef.Name != "") {
				return _lower_struct_lit(ll, name)
			}
		}
		if (name == "None") {
			vp := _variant_parent(ll.Registry, "None")
			_ = vp
			if (vp.Name == "") {
				nt := _lnew_temp(ll)
				_ = nt
				ll = nt.L
				ll = _lemit(ll, new_inst(IrOpOpAlloc{}, nt.Temp, int64(1), int64(0), "_Optional", int64(0)))
				tag_nt := _lnew_temp(ll)
				_ = tag_nt
				ll = tag_nt.L
				ll = _lemit(ll, new_inst(IrOpOpConst{}, tag_nt.Temp, int64(0), int64(0), "", int64(0)))
				ll = _lemit(ll, new_inst(IrOpOpFieldSet{}, nt.Temp, tag_nt.Temp, int64(0), "_tag", int64(0)))
				ll = _lset_type(ll, nt.Temp, int64(201))
				return LT{L: ll, Temp: nt.Temp}
			}
		}
		slot := _lenv_lookup(ll, name)
		_ = slot
		if (slot >= int64(0)) {
			src_type := _lget_type(ll, slot)
			_ = src_type
			if (src_type == int64(12)) {
				nt := _lnew_str_temp(ll)
				_ = nt
				ll = nt.L
				ll = _lemit(ll, new_inst(IrOpOpLoad{}, nt.Temp, slot, int64(0), "", int64(0)))
				ll = _lemit(ll, new_inst(IrOpOpLoad{}, (nt.Temp + int64(1)), (slot + int64(1)), int64(0), "", int64(0)))
				ll = _lset_type(ll, nt.Temp, int64(12))
				return LT{L: ll, Temp: nt.Temp}
			} else {
				nt := _lnew_temp(ll)
				_ = nt
				ll = nt.L
				ll = _lemit(ll, new_inst(IrOpOpLoad{}, nt.Temp, slot, int64(0), "", int64(0)))
				if (src_type != int64(0)) {
					ll = _lset_type(ll, nt.Temp, src_type)
				}
				return LT{L: ll, Temp: nt.Temp}
			}
		}
		tag := _variant_tag(ll.Registry, name)
		_ = tag
		if (tag > int64(0)) {
			parent := _variant_parent(ll.Registry, name)
			_ = parent
			if sum_has_data(parent) {
				alloc_sz := sum_alloc_size(parent)
				_ = alloc_sz
				nt := _lnew_temp(ll)
				_ = nt
				ll = nt.L
				ll = _lemit(ll, new_inst(IrOpOpAlloc{}, nt.Temp, alloc_sz, int64(0), parent.Name, int64(0)))
				st_id := _struct_type_id(ll.Registry, parent.Name)
				_ = st_id
				if (st_id > int64(0)) {
					ll = _lset_type(ll, nt.Temp, st_id)
				}
				tag_nt := _lnew_temp(ll)
				_ = tag_nt
				ll = tag_nt.L
				ll = _lemit(ll, new_inst(IrOpOpConst{}, tag_nt.Temp, tag, int64(0), "", int64(0)))
				ll = _lemit(ll, new_inst(IrOpOpFieldSet{}, nt.Temp, tag_nt.Temp, int64(0), "_tag", int64(0)))
				return LT{L: ll, Temp: nt.Temp}
			}
			nt := _lnew_temp(ll)
			_ = nt
			ll = nt.L
			ll = _lemit(ll, new_inst(IrOpOpConst{}, nt.Temp, tag, int64(0), "", int64(0)))
			return LT{L: ll, Temp: nt.Temp}
		}
		if _lconst_has(ll, name) {
			cval := _lconst_lookup(ll, name)
			_ = cval
			nt := _lnew_temp(ll)
			_ = nt
			ll = nt.L
			ll = _lemit(ll, new_inst(IrOpOpConst{}, nt.Temp, cval, int64(0), "", int64(0)))
			ll = _lset_type(ll, nt.Temp, int64(1))
			return LT{L: ll, Temp: nt.Temp}
		}
		if _lconst_str_has(ll, name) {
			sval := _lconst_str_lookup(ll, name)
			_ = sval
			str_idx := mod_find_string(ll.Module, sval)
			_ = str_idx
			if (str_idx == int64(0)) {
				new_mod := mod_add_string(ll.Module, sval)
				_ = new_mod
				str_idx = (int64(len(new_mod.String_constants)) - int64(1))
				ll = Lowerer{Tokens: ll.Tokens, Pos: ll.Pos, Store: ll.Store, Registry: ll.Registry, Treg: ll.Treg, Pool: ll.Pool, Table: ll.Table, File: ll.File, Module: new_mod, Current_func: ll.Current_func, Temp_counter: ll.Temp_counter, Label_counter: ll.Label_counter, Env_names: ll.Env_names, Env_slots: ll.Env_slots, Loop_start: ll.Loop_start, Loop_end: ll.Loop_end, Const_names: ll.Const_names, Const_vals: ll.Const_vals, Const_str_names: ll.Const_str_names, Const_str_vals: ll.Const_str_vals, Temp_types: ll.Temp_types, Fn_error_type: ll.Fn_error_type, Defer_starts: ll.Defer_starts, Defer_ends: ll.Defer_ends}
			}
			nt := _lnew_str_temp(ll)
			_ = nt
			ll = _lemit(nt.L, new_inst(IrOpOpConstStr{}, nt.Temp, str_idx, int64(len(sval)), "", int64(0)))
			ll = _lset_type(ll, nt.Temp, int64(12))
			return LT{L: ll, Temp: nt.Temp}
		}
		fsig := reg_find_fn(ll.Registry, name)
		_ = fsig
		if (fsig.Name != "") {
			ll = _lset_reg(ll, reg_add_fn_value_ref(ll.Registry, name))
			clos := _lnew_temp(ll)
			_ = clos
			ll = clos.L
			ll = _lemit(ll, new_inst(IrOpOpAlloc{}, clos.Temp, int64(2), int64(0), "_Closure", int64(0)))
			fn_ref := _lnew_temp(ll)
			_ = fn_ref
			ll = fn_ref.L
			ll = _lemit(ll, new_inst(IrOpOpFnRef{}, fn_ref.Temp, int64(0), int64(0), ("_tramp_" + name), int64(0)))
			ll = _lemit(ll, new_inst(IrOpOpFieldSet{}, clos.Temp, fn_ref.Temp, int64(0), "_fn", int64(0)))
			null_env := _lnew_temp(ll)
			_ = null_env
			ll = null_env.L
			ll = _lemit(ll, new_inst(IrOpOpConst{}, null_env.Temp, int64(0), int64(0), "", int64(0)))
			ll = _lemit(ll, new_inst(IrOpOpFieldSet{}, clos.Temp, null_env.Temp, int64(1), "_env", int64(0)))
			return LT{L: ll, Temp: clos.Temp}
		}
		nt := _lnew_temp(ll)
		_ = nt
		ll = nt.L
		ll = _lemit(ll, new_inst(IrOpOpConst{}, nt.Temp, int64(0), int64(0), "", int64(0)))
		return LT{L: ll, Temp: nt.Temp}
	}
	if (k == "(") {
		ll := _ladv(l)
		_ = ll
		ll = _lskip_nl(ll)
		result := _lower_expr(ll)
		_ = result
		ll = result.L
		if (_lk(ll) == ",") {
			nt := _lnew_temp(ll)
			_ = nt
			ll = nt.L
			count := int64(1)
			_ = count
			ll = _lemit(ll, new_inst(IrOpOpAlloc{}, nt.Temp, int64(8), int64(0), "_Tuple", int64(0)))
			ll = _lemit(ll, new_inst(IrOpOpFieldSet{}, nt.Temp, result.Temp, int64(0), "0", int64(0)))
			for ((_lk(ll) == ",") && (_lk(ll) != "EOF")) {
				ll = _ladv(ll)
				ll = _lskip_nl(ll)
				if (_lk(ll) != ")") {
					elem := _lower_expr(ll)
					_ = elem
					ll = elem.L
					ll = _lemit(ll, new_inst(IrOpOpFieldSet{}, nt.Temp, elem.Temp, count, i2s(count), int64(0)))
					count = (count + int64(1))
				}
				ll = _lskip_nl(ll)
			}
			if (_lk(ll) == ")") {
				ll = _ladv(ll)
			}
			return LT{L: ll, Temp: nt.Temp}
		}
		for ((_lk(ll) != ")") && (_lk(ll) != "EOF")) {
			ll = _ladv(ll)
		}
		if (_lk(ll) == ")") {
			ll = _ladv(ll)
		}
		return LT{L: ll, Temp: result.Temp}
	}
	if (k == "[") {
		ll := _ladv(l)
		_ = ll
		ll = _lskip_nl(ll)
		if (_lk(ll) == "]") {
			ll = _ladv(ll)
			nt := _lnew_temp(ll)
			_ = nt
			ll = nt.L
			ll = _lemit(ll, new_inst(IrOpOpArrayNew{}, nt.Temp, int64(0), int64(0), "", int64(0)))
			return LT{L: ll, Temp: nt.Temp}
		}
		is_listcomp := false
		_ = is_listcomp
		peek_pos := ll.Pos
		_ = peek_pos
		depth := int64(0)
		_ = depth
		for (peek_pos < int64(len(ll.Tokens))) {
			pk := token_name(ll.Tokens[peek_pos].Kind)
			_ = pk
			if ((pk == "[") || (pk == "(")) {
				depth = (depth + int64(1))
			}
			if (pk == "]") {
				if (depth == int64(0)) {
					peek_pos = int64(len(ll.Tokens))
				} else {
					depth = (depth - int64(1))
				}
			}
			if (pk == ")") {
				if (depth > int64(0)) {
					depth = (depth - int64(1))
				}
			}
			if ((pk == "for") && (depth == int64(0))) {
				is_listcomp = true
			}
			peek_pos = (peek_pos + int64(1))
		}
		if is_listcomp {
			map_expr_pos := ll.Pos
			_ = map_expr_pos
			skip_depth := int64(0)
			_ = skip_depth
			for ((_lk(ll) != "for") || (skip_depth > int64(0))) {
				sk := _lk(ll)
				_ = sk
				if ((sk == "(") || (sk == "[")) {
					skip_depth = (skip_depth + int64(1))
				}
				if ((sk == ")") || (sk == "]")) {
					if (skip_depth > int64(0)) {
						skip_depth = (skip_depth - int64(1))
					}
				}
				if (sk == "EOF") {
					return LT{L: ll, Temp: int64(0)}
				}
				ll = _ladv(ll)
			}
			ll = _ladv(ll)
			lc_var := ""
			_ = lc_var
			if (_lk(ll) == "IDENT") {
				lc_var = _lcur(ll).Text
				ll = _ladv(ll)
			}
			if (_lk(ll) == "in") {
				ll = _ladv(ll)
			}
			iter := _lower_expr(ll)
			_ = iter
			ll = iter.L
			ll = _lskip_nl(ll)
			has_where := false
			_ = has_where
			where_pos := int64(0)
			_ = where_pos
			if ((_lk(ll) == "where") || (_lk(ll) == "if")) {
				has_where = true
				ll = _ladv(ll)
				where_pos = ll.Pos
				for ((_lk(ll) != "]") && (_lk(ll) != "EOF")) {
					ll = _ladv(ll)
				}
			}
			if (_lk(ll) == "]") {
				ll = _ladv(ll)
			}
			end_pos := ll.Pos
			_ = end_pos
			arr_nt := _lnew_temp(ll)
			_ = arr_nt
			ll = arr_nt.L
			ll = _lemit(ll, new_inst(IrOpOpArrayNew{}, arr_nt.Temp, int64(0), int64(0), "", int64(0)))
			len_nt := _lnew_temp(ll)
			_ = len_nt
			ll = len_nt.L
			ll = _lemit(ll, new_inst(IrOpOpArrayLen{}, len_nt.Temp, iter.Temp, int64(0), "", int64(0)))
			idx_nt := _lnew_temp(ll)
			_ = idx_nt
			ll = idx_nt.L
			ll = _lemit(ll, new_inst(IrOpOpConst{}, idx_nt.Temp, int64(0), int64(0), "", int64(0)))
			loop_lbl := _lnew_label(ll)
			_ = loop_lbl
			ll = loop_lbl.L
			end_lbl := _lnew_label(ll)
			_ = end_lbl
			ll = end_lbl.L
			ll = _lemit(ll, new_inst(IrOpOpLabel{}, int64(0), loop_lbl.Temp, int64(0), "", int64(0)))
			cond_nt := _lnew_temp(ll)
			_ = cond_nt
			ll = cond_nt.L
			ll = _lemit(ll, new_inst(IrOpOpLt{}, cond_nt.Temp, idx_nt.Temp, len_nt.Temp, "", int64(0)))
			ll = _lemit(ll, new_inst(IrOpOpBranchFalse{}, int64(0), cond_nt.Temp, end_lbl.Temp, "", int64(0)))
			elem_nt := _lnew_temp(ll)
			_ = elem_nt
			ll = elem_nt.L
			ll = _lemit(ll, new_inst(IrOpOpArrayGet{}, elem_nt.Temp, iter.Temp, idx_nt.Temp, "", int64(0)))
			if (lc_var != "") {
				ll = _lenv_add(ll, lc_var, elem_nt.Temp)
			}
			if has_where {
				ll = _lset_pos(ll, where_pos)
				wcond := _lower_expr(ll)
				_ = wcond
				ll = wcond.L
				skip_lbl := _lnew_label(ll)
				_ = skip_lbl
				ll = skip_lbl.L
				ll = _lemit(ll, new_inst(IrOpOpBranchFalse{}, int64(0), wcond.Temp, skip_lbl.Temp, "", int64(0)))
				ll = _lset_pos(ll, map_expr_pos)
				mval := _lower_expr(ll)
				_ = mval
				ll = mval.L
				ll = _lemit(ll, new_inst(IrOpOpArrayAppend{}, arr_nt.Temp, arr_nt.Temp, mval.Temp, "", int64(0)))
				ll = _lemit(ll, new_inst(IrOpOpLabel{}, int64(0), skip_lbl.Temp, int64(0), "", int64(0)))
			} else {
				ll = _lset_pos(ll, map_expr_pos)
				mval := _lower_expr(ll)
				_ = mval
				ll = mval.L
				ll = _lemit(ll, new_inst(IrOpOpArrayAppend{}, arr_nt.Temp, arr_nt.Temp, mval.Temp, "", int64(0)))
			}
			one_nt := _lnew_temp(ll)
			_ = one_nt
			ll = one_nt.L
			ll = _lemit(ll, new_inst(IrOpOpConst{}, one_nt.Temp, int64(1), int64(0), "", int64(0)))
			inc_nt := _lnew_temp(ll)
			_ = inc_nt
			ll = inc_nt.L
			ll = _lemit(ll, new_inst(IrOpOpAdd{}, inc_nt.Temp, idx_nt.Temp, one_nt.Temp, "", int64(0)))
			ll = _lemit(ll, new_inst(IrOpOpStore{}, idx_nt.Temp, inc_nt.Temp, int64(0), "", int64(0)))
			ll = _lemit(ll, new_inst(IrOpOpJump{}, int64(0), loop_lbl.Temp, int64(0), "", int64(0)))
			ll = _lemit(ll, new_inst(IrOpOpLabel{}, int64(0), end_lbl.Temp, int64(0), "", int64(0)))
			ll = _lset_pos(ll, end_pos)
			return LT{L: ll, Temp: arr_nt.Temp}
		}
		nt := _lnew_temp(ll)
		_ = nt
		ll = nt.L
		ll = _lemit(ll, new_inst(IrOpOpArrayNew{}, nt.Temp, int64(0), int64(0), "", int64(0)))
		arr_elem_tid := int64(0)
		_ = arr_elem_tid
		for ((_lk(ll) != "]") && (_lk(ll) != "EOF")) {
			ll = _lskip_nl(ll)
			elem := _lower_expr(ll)
			_ = elem
			ll = elem.L
			elem_tid := _lget_type(ll, elem.Temp)
			_ = elem_tid
			if (elem_tid > int64(0)) {
				arr_elem_tid = elem_tid
			}
			ll = _lemit(ll, new_inst(IrOpOpArrayAppend{}, nt.Temp, nt.Temp, elem.Temp, "", elem_tid))
			ll = _lskip_nl(ll)
			if (_lk(ll) == ",") {
				ll = _ladv(ll)
			}
		}
		if (_lk(ll) == "]") {
			ll = _ladv(ll)
		}
		if (arr_elem_tid == int64(12)) {
			ll = _lset_type(ll, nt.Temp, int64(13))
		}
		return LT{L: ll, Temp: nt.Temp}
	}
	keyword_name := _lcur(l).Text
	_ = keyword_name
	kw_slot := _lenv_lookup(l, keyword_name)
	_ = kw_slot
	if (kw_slot >= int64(0)) {
		ll := _ladv(l)
		_ = ll
		kw_type := _lget_type(ll, kw_slot)
		_ = kw_type
		if (kw_type == int64(12)) {
			nt := _lnew_str_temp(ll)
			_ = nt
			ll = nt.L
			ll = _lemit(ll, new_inst(IrOpOpLoad{}, nt.Temp, kw_slot, int64(0), "", int64(0)))
			ll = _lemit(ll, new_inst(IrOpOpLoad{}, (nt.Temp + int64(1)), (kw_slot + int64(1)), int64(0), "", int64(0)))
			ll = _lset_type(ll, nt.Temp, int64(12))
			return LT{L: ll, Temp: nt.Temp}
		}
		nt := _lnew_temp(ll)
		_ = nt
		ll = nt.L
		ll = _lemit(ll, new_inst(IrOpOpLoad{}, nt.Temp, kw_slot, int64(0), "", int64(0)))
		if (kw_type != int64(0)) {
			ll = _lset_type(ll, nt.Temp, kw_type)
		}
		return LT{L: ll, Temp: nt.Temp}
	}
	nt := _lnew_temp(l)
	_ = nt
	ll := _ladv(nt.L)
	_ = ll
	ll = _lemit(ll, new_inst(IrOpOpConst{}, nt.Temp, int64(0), int64(0), "", int64(0)))
	return LT{L: ll, Temp: nt.Temp}
}

func _lower_conversion(l Lowerer, obj int64, method string, target string) LT {
	ll := l
	_ = ll
	obj_type := _lget_type(ll, obj)
	_ = obj_type
	if (method == "to") {
		if ((obj_type == int64(12)) && (target == "i64")) {
			result_nt := _lnew_temp(ll)
			_ = result_nt
			ll = result_nt.L
			extra := _lnew_temp(ll)
			_ = extra
			ll = extra.L
			ll = _lemit(ll, new_inst(IrOpOpStrToInt{}, result_nt.Temp, obj, int64(0), "", int64(0)))
			success_nt := _lnew_temp(ll)
			_ = success_nt
			ll = success_nt.L
			ll = _lemit(ll, new_inst(IrOpOpLoad{}, success_nt.Temp, (result_nt.Temp + int64(1)), int64(0), "", int64(0)))
			err_lbl := _lnew_label(ll)
			_ = err_lbl
			ll = err_lbl.L
			ok_lbl := _lnew_label(ll)
			_ = ok_lbl
			ll = ok_lbl.L
			end_lbl := _lnew_label(ll)
			_ = end_lbl
			ll = end_lbl.L
			ll = _lemit(ll, new_inst(IrOpOpBranchTrue{}, int64(0), success_nt.Temp, ok_lbl.Temp, "", int64(0)))
			err_res := _lnew_temp(ll)
			_ = err_res
			ll = err_res.L
			ll = _lemit(ll, new_inst(IrOpOpAlloc{}, err_res.Temp, int64(2), int64(0), "_Result", int64(0)))
			err_tag := _lnew_temp(ll)
			_ = err_tag
			ll = err_tag.L
			ll = _lemit(ll, new_inst(IrOpOpConst{}, err_tag.Temp, int64(1), int64(0), "", int64(0)))
			ll = _lemit(ll, new_inst(IrOpOpFieldSet{}, err_res.Temp, err_tag.Temp, int64(0), "_tag", int64(0)))
			err_val := _lnew_temp(ll)
			_ = err_val
			ll = err_val.L
			ll = _lemit(ll, new_inst(IrOpOpConst{}, err_val.Temp, int64(0), int64(0), "", int64(0)))
			ll = _lemit(ll, new_inst(IrOpOpFieldSet{}, err_res.Temp, err_val.Temp, int64(1), "_val", int64(0)))
			final_nt := _lnew_temp(ll)
			_ = final_nt
			ll = final_nt.L
			ll = _lemit(ll, new_inst(IrOpOpStore{}, final_nt.Temp, err_res.Temp, int64(0), "", int64(0)))
			ll = _lemit(ll, new_inst(IrOpOpJump{}, int64(0), end_lbl.Temp, int64(0), "", int64(0)))
			ll = _lemit(ll, new_inst(IrOpOpLabel{}, int64(0), ok_lbl.Temp, int64(0), "", int64(0)))
			ok_res := _lnew_temp(ll)
			_ = ok_res
			ll = ok_res.L
			ll = _lemit(ll, new_inst(IrOpOpAlloc{}, ok_res.Temp, int64(2), int64(0), "_Result", int64(0)))
			ok_tag := _lnew_temp(ll)
			_ = ok_tag
			ll = ok_tag.L
			ll = _lemit(ll, new_inst(IrOpOpConst{}, ok_tag.Temp, int64(0), int64(0), "", int64(0)))
			ll = _lemit(ll, new_inst(IrOpOpFieldSet{}, ok_res.Temp, ok_tag.Temp, int64(0), "_tag", int64(0)))
			ll = _lemit(ll, new_inst(IrOpOpFieldSet{}, ok_res.Temp, result_nt.Temp, int64(1), "_val", int64(0)))
			ll = _lemit(ll, new_inst(IrOpOpStore{}, final_nt.Temp, ok_res.Temp, int64(0), "", int64(0)))
			ll = _lemit(ll, new_inst(IrOpOpLabel{}, int64(0), end_lbl.Temp, int64(0), "", int64(0)))
			ll = _lset_type(ll, final_nt.Temp, int64(200))
			return LT{L: ll, Temp: final_nt.Temp}
		}
		if ((obj_type == int64(12)) && (target == "f64")) {
			fnt := _lnew_temp(ll)
			_ = fnt
			ll = fnt.L
			ll = _lemit(ll, new_inst(IrOpOpCall{}, fnt.Temp, obj, int64(2), "_aria_str_to_float", int64(0)))
			ll = _lset_type(ll, fnt.Temp, int64(9))
			return LT{L: ll, Temp: fnt.Temp}
		}
		return LT{L: ll, Temp: obj}
	}
	if (method == "trunc") {
		return LT{L: ll, Temp: obj}
	}
	return LT{L: ll, Temp: obj}
}

func _lower_method_call(l Lowerer, obj int64, method string) LT {
	ll := _ladv(l)
	_ = ll
	ll = _lskip_nl(ll)
	if (method == "len") {
		for ((_lk(ll) != ")") && (_lk(ll) != "EOF")) {
			ll = _ladv(ll)
		}
		if (_lk(ll) == ")") {
			ll = _ladv(ll)
		}
		nt := _lnew_temp(ll)
		_ = nt
		ll = nt.L
		obj_type := _lget_type(ll, obj)
		_ = obj_type
		if (obj_type == int64(12)) {
			ll = _lemit(ll, new_inst(IrOpOpStrLen{}, nt.Temp, obj, int64(0), "", int64(0)))
		} else if (obj_type == int64(300)) {
			ll = _lemit(ll, new_inst(IrOpOpMapLen{}, nt.Temp, obj, int64(0), "", int64(0)))
		} else if (obj_type == int64(301)) {
			ll = _lemit(ll, new_inst(IrOpOpSetLen{}, nt.Temp, obj, int64(0), "", int64(0)))
		} else {
			ll = _lemit(ll, new_inst(IrOpOpArrayLen{}, nt.Temp, obj, int64(0), "", int64(0)))
		}
		ll = _lset_type(ll, nt.Temp, int64(1))
		return LT{L: ll, Temp: nt.Temp}
	}
	if ((method == "charCount") && (_lget_type(ll, obj) == int64(12))) {
		if (_lk(ll) == ")") {
			ll = _ladv(ll)
		}
		nt := _lnew_temp(ll)
		_ = nt
		ll = nt.L
		ll = _lemit(ll, new_inst(IrOpOpCall{}, nt.Temp, obj, int64(2), "_aria_str_char_count", int64(1)))
		ll = _lset_type(ll, nt.Temp, int64(1))
		return LT{L: ll, Temp: nt.Temp}
	}
	if ((method == "chars") && (_lget_type(ll, obj) == int64(12))) {
		if (_lk(ll) == ")") {
			ll = _ladv(ll)
		}
		nt := _lnew_temp(ll)
		_ = nt
		ll = nt.L
		ll = _lemit(ll, new_inst(IrOpOpCall{}, nt.Temp, obj, int64(2), "_aria_str_chars", int64(1)))
		ll = _lset_type(ll, nt.Temp, int64(1))
		return LT{L: ll, Temp: nt.Temp}
	}
	if ((method == "graphemes") && (_lget_type(ll, obj) == int64(12))) {
		if (_lk(ll) == ")") {
			ll = _ladv(ll)
		}
		nt := _lnew_temp(ll)
		_ = nt
		ll = nt.L
		ll = _lemit(ll, new_inst(IrOpOpCall{}, nt.Temp, obj, int64(2), "_aria_str_graphemes", int64(13)))
		ll = _lset_type(ll, nt.Temp, int64(13))
		return LT{L: ll, Temp: nt.Temp}
	}
	if ((method == "bytes") && (_lget_type(ll, obj) == int64(12))) {
		if (_lk(ll) == ")") {
			ll = _ladv(ll)
		}
		nt := _lnew_temp(ll)
		_ = nt
		ll = nt.L
		ll = _lemit(ll, new_inst(IrOpOpCall{}, nt.Temp, obj, int64(2), "_aria_str_chars", int64(1)))
		ll = _lset_type(ll, nt.Temp, int64(1))
		return LT{L: ll, Temp: nt.Temp}
	}
	obj_type_ch := _lget_type(ll, obj)
	_ = obj_type_ch
	if ((method == "send") && (((obj_type_ch == int64(17)) || (obj_type_ch == int64(19))))) {
		val := _lower_expr(ll)
		_ = val
		ll = val.L
		if (_lk(ll) == ")") {
			ll = _ladv(ll)
		}
		pack0 := _lnew_temp(ll)
		_ = pack0
		ll = pack0.L
		ll = _lemit(ll, new_inst(IrOpOpStore{}, pack0.Temp, obj, int64(0), "", int64(0)))
		pack1 := _lnew_temp(ll)
		_ = pack1
		ll = pack1.L
		ll = _lemit(ll, new_inst(IrOpOpStore{}, pack1.Temp, val.Temp, int64(0), "", int64(0)))
		nt := _lnew_temp(ll)
		_ = nt
		ll = nt.L
		ll = _lemit(ll, new_inst(IrOpOpCall{}, nt.Temp, pack0.Temp, int64(2), "_aria_chan_send", int64(1)))
		ll = _lset_type(ll, nt.Temp, int64(1))
		return LT{L: ll, Temp: nt.Temp}
	}
	if ((method == "recv") && (((obj_type_ch == int64(17)) || (obj_type_ch == int64(20))))) {
		if (_lk(ll) == ")") {
			ll = _ladv(ll)
		}
		nt := _lnew_temp(ll)
		_ = nt
		ll = nt.L
		ll = _lemit(ll, new_inst(IrOpOpCall{}, nt.Temp, obj, int64(1), "_aria_chan_recv", int64(1)))
		ll = _lset_type(ll, nt.Temp, int64(1))
		return LT{L: ll, Temp: nt.Temp}
	}
	if ((method == "try_recv") && (((obj_type_ch == int64(17)) || (obj_type_ch == int64(20))))) {
		if (_lk(ll) == ")") {
			ll = _ladv(ll)
		}
		nt := _lnew_str_temp(ll)
		_ = nt
		ll = nt.L
		ll = _lemit(ll, new_inst(IrOpOpCall{}, nt.Temp, obj, int64(1), "_aria_chan_try_recv", int64(0)))
		return LT{L: ll, Temp: nt.Temp}
	}
	if ((method == "close") && ((((obj_type_ch == int64(17)) || (obj_type_ch == int64(19))) || (obj_type_ch == int64(20))))) {
		if (_lk(ll) == ")") {
			ll = _ladv(ll)
		}
		nt := _lnew_temp(ll)
		_ = nt
		ll = nt.L
		ll = _lemit(ll, new_inst(IrOpOpCall{}, nt.Temp, obj, int64(1), "_aria_chan_close", int64(0)))
		return LT{L: ll, Temp: nt.Temp}
	}
	if ((method == "await") && (((obj_type_ch == int64(1)) || (obj_type_ch == int64(18))))) {
		if (_lk(ll) == ")") {
			ll = _ladv(ll)
		}
		nt := _lnew_temp(ll)
		_ = nt
		ll = nt.L
		ll = _lemit(ll, new_inst(IrOpOpCall{}, nt.Temp, obj, int64(1), "_aria_task_await", int64(1)))
		ll = _lset_type(ll, nt.Temp, int64(1))
		return LT{L: ll, Temp: nt.Temp}
	}
	if ((method == "done") && (((obj_type_ch == int64(1)) || (obj_type_ch == int64(18))))) {
		if (_lk(ll) == ")") {
			ll = _ladv(ll)
		}
		nt := _lnew_temp(ll)
		_ = nt
		ll = nt.L
		ll = _lemit(ll, new_inst(IrOpOpCall{}, nt.Temp, obj, int64(1), "_aria_task_done", int64(1)))
		ll = _lset_type(ll, nt.Temp, int64(1))
		return LT{L: ll, Temp: nt.Temp}
	}
	if ((method == "cancel") && (((obj_type_ch == int64(1)) || (obj_type_ch == int64(18))))) {
		if (_lk(ll) == ")") {
			ll = _ladv(ll)
		}
		nt := _lnew_temp(ll)
		_ = nt
		ll = nt.L
		ll = _lemit(ll, new_inst(IrOpOpCall{}, nt.Temp, obj, int64(1), "_aria_task_cancel", int64(0)))
		return LT{L: ll, Temp: nt.Temp}
	}
	if ((method == "result") && (((obj_type_ch == int64(1)) || (obj_type_ch == int64(18))))) {
		if (_lk(ll) == ")") {
			ll = _ladv(ll)
		}
		nt := _lnew_temp(ll)
		_ = nt
		ll = nt.L
		ll = _lemit(ll, new_inst(IrOpOpCall{}, nt.Temp, obj, int64(1), "_aria_task_result", int64(1)))
		ll = _lset_type(ll, nt.Temp, int64(1))
		return LT{L: ll, Temp: nt.Temp}
	}
	if ((method == "set") && (obj_type_ch == int64(300))) {
		key := _lower_expr(ll)
		_ = key
		ll = key.L
		if (_lk(ll) == ",") {
			ll = _ladv(ll)
		}
		val := _lower_expr(ll)
		_ = val
		ll = val.L
		if (_lk(ll) == ")") {
			ll = _ladv(ll)
		}
		ll = _lemit(ll, new_inst(IrOpOpMapSet{}, int64(0), obj, key.Temp, "", val.Temp))
		return LT{L: ll, Temp: obj}
	}
	if ((method == "get") && (obj_type_ch == int64(300))) {
		key := _lower_expr(ll)
		_ = key
		ll = key.L
		if (_lk(ll) == ")") {
			ll = _ladv(ll)
		}
		nt := _lnew_temp(ll)
		_ = nt
		ll = nt.L
		ll = _lemit(ll, new_inst(IrOpOpMapGet{}, nt.Temp, obj, key.Temp, "", int64(0)))
		return LT{L: ll, Temp: nt.Temp}
	}
	if ((method == "has") && (obj_type_ch == int64(300))) {
		key := _lower_expr(ll)
		_ = key
		ll = key.L
		if (_lk(ll) == ")") {
			ll = _ladv(ll)
		}
		nt := _lnew_temp(ll)
		_ = nt
		ll = nt.L
		ll = _lemit(ll, new_inst(IrOpOpMapContains{}, nt.Temp, obj, key.Temp, "", int64(0)))
		return LT{L: ll, Temp: nt.Temp}
	}
	if ((method == "keys") && (obj_type_ch == int64(300))) {
		for ((_lk(ll) != ")") && (_lk(ll) != "EOF")) {
			ll = _ladv(ll)
		}
		if (_lk(ll) == ")") {
			ll = _ladv(ll)
		}
		nt := _lnew_temp(ll)
		_ = nt
		ll = nt.L
		ll = _lemit(ll, new_inst(IrOpOpMapKeys{}, nt.Temp, obj, int64(0), "", int64(0)))
		ll = _lset_type(ll, nt.Temp, int64(13))
		return LT{L: ll, Temp: nt.Temp}
	}
	if ((method == "add") && (_lget_type(ll, obj) == int64(301))) {
		key := _lower_expr(ll)
		_ = key
		ll = key.L
		if (_lk(ll) == ")") {
			ll = _ladv(ll)
		}
		ll = _lemit(ll, new_inst(IrOpOpSetAdd{}, int64(0), obj, key.Temp, "", int64(0)))
		return LT{L: ll, Temp: obj}
	}
	if ((method == "contains") && (_lget_type(ll, obj) == int64(301))) {
		key := _lower_expr(ll)
		_ = key
		ll = key.L
		if (_lk(ll) == ")") {
			ll = _ladv(ll)
		}
		nt := _lnew_temp(ll)
		_ = nt
		ll = nt.L
		ll = _lemit(ll, new_inst(IrOpOpSetContains{}, nt.Temp, obj, key.Temp, "", int64(0)))
		ll = _lset_type(ll, nt.Temp, int64(1))
		return LT{L: ll, Temp: nt.Temp}
	}
	if ((method == "remove") && (_lget_type(ll, obj) == int64(301))) {
		key := _lower_expr(ll)
		_ = key
		ll = key.L
		if (_lk(ll) == ")") {
			ll = _ladv(ll)
		}
		ll = _lemit(ll, new_inst(IrOpOpSetRemove{}, int64(0), obj, key.Temp, "", int64(0)))
		return LT{L: ll, Temp: obj}
	}
	if ((method == "values") && (_lget_type(ll, obj) == int64(301))) {
		for ((_lk(ll) != ")") && (_lk(ll) != "EOF")) {
			ll = _ladv(ll)
		}
		if (_lk(ll) == ")") {
			ll = _ladv(ll)
		}
		nt := _lnew_temp(ll)
		_ = nt
		ll = nt.L
		ll = _lemit(ll, new_inst(IrOpOpSetValues{}, nt.Temp, obj, int64(0), "", int64(0)))
		ll = _lset_type(ll, nt.Temp, int64(13))
		return LT{L: ll, Temp: nt.Temp}
	}
	if (method == "filter") {
		obj_type := _lget_type(ll, obj)
		_ = obj_type
		pred := _lower_expr(ll)
		_ = pred
		ll = pred.L
		if (_lk(ll) == ")") {
			ll = _ladv(ll)
		}
		new_arr := _lnew_temp(ll)
		_ = new_arr
		ll = new_arr.L
		cap_nt := _lnew_temp(ll)
		_ = cap_nt
		ll = cap_nt.L
		ll = _lemit(ll, new_inst(IrOpOpConst{}, cap_nt.Temp, int64(8), int64(0), "", int64(0)))
		ll = _lemit(ll, new_inst(IrOpOpArrayNew{}, new_arr.Temp, cap_nt.Temp, int64(0), "", int64(0)))
		len_nt := _lnew_temp(ll)
		_ = len_nt
		ll = len_nt.L
		ll = _lemit(ll, new_inst(IrOpOpArrayLen{}, len_nt.Temp, obj, int64(0), "", int64(0)))
		idx_nt := _lnew_temp(ll)
		_ = idx_nt
		ll = idx_nt.L
		ll = _lemit(ll, new_inst(IrOpOpConst{}, idx_nt.Temp, int64(0), int64(0), "", int64(0)))
		loop_lbl := _lnew_label(ll)
		_ = loop_lbl
		ll = loop_lbl.L
		end_lbl := _lnew_label(ll)
		_ = end_lbl
		ll = end_lbl.L
		body_lbl := _lnew_label(ll)
		_ = body_lbl
		ll = body_lbl.L
		ll = _lemit(ll, new_inst(IrOpOpLabel{}, int64(0), loop_lbl.Temp, int64(0), "", int64(0)))
		cmp_nt := _lnew_temp(ll)
		_ = cmp_nt
		ll = cmp_nt.L
		ll = _lemit(ll, new_inst(IrOpOpLt{}, cmp_nt.Temp, idx_nt.Temp, len_nt.Temp, "", int64(0)))
		ll = _lemit(ll, new_inst(IrOpOpBranchFalse{}, int64(0), cmp_nt.Temp, end_lbl.Temp, "", int64(0)))
		elem_nt := _lnew_temp(ll)
		_ = elem_nt
		ll = elem_nt.L
		ll = _lemit(ll, new_inst(IrOpOpArrayGet{}, elem_nt.Temp, obj, idx_nt.Temp, "", int64(0)))
		fn_ptr := _lnew_temp(ll)
		_ = fn_ptr
		ll = fn_ptr.L
		ll = _lemit(ll, new_inst(IrOpOpFieldGet{}, fn_ptr.Temp, pred.Temp, int64(0), "", int64(0)))
		arg_start := _lnew_temp(ll)
		_ = arg_start
		ll = arg_start.L
		arg_elem := _lnew_temp(ll)
		_ = arg_elem
		ll = arg_elem.L
		env_ptr := _lnew_temp(ll)
		_ = env_ptr
		ll = env_ptr.L
		ll = _lemit(ll, new_inst(IrOpOpFieldGet{}, env_ptr.Temp, pred.Temp, int64(1), "", int64(0)))
		ll = _lemit(ll, new_inst(IrOpOpStore{}, arg_start.Temp, env_ptr.Temp, int64(0), "", int64(0)))
		ll = _lemit(ll, new_inst(IrOpOpStore{}, arg_elem.Temp, elem_nt.Temp, int64(0), "", int64(0)))
		result_nt := _lnew_temp(ll)
		_ = result_nt
		ll = result_nt.L
		ll = _lemit(ll, new_inst(IrOpOpCallIndirect{}, result_nt.Temp, fn_ptr.Temp, int64(2), "", arg_start.Temp))
		skip_lbl := _lnew_label(ll)
		_ = skip_lbl
		ll = skip_lbl.L
		ll = _lemit(ll, new_inst(IrOpOpBranchFalse{}, int64(0), result_nt.Temp, skip_lbl.Temp, "", int64(0)))
		ll = _lemit(ll, new_inst(IrOpOpArrayAppend{}, new_arr.Temp, new_arr.Temp, elem_nt.Temp, "", int64(0)))
		ll = _lemit(ll, new_inst(IrOpOpLabel{}, int64(0), skip_lbl.Temp, int64(0), "", int64(0)))
		one_nt := _lnew_temp(ll)
		_ = one_nt
		ll = one_nt.L
		ll = _lemit(ll, new_inst(IrOpOpConst{}, one_nt.Temp, int64(1), int64(0), "", int64(0)))
		ll = _lemit(ll, new_inst(IrOpOpAdd{}, idx_nt.Temp, idx_nt.Temp, one_nt.Temp, "", int64(0)))
		ll = _lemit(ll, new_inst(IrOpOpJump{}, int64(0), loop_lbl.Temp, int64(0), "", int64(0)))
		ll = _lemit(ll, new_inst(IrOpOpLabel{}, int64(0), end_lbl.Temp, int64(0), "", int64(0)))
		if (obj_type > int64(0)) {
			ll = _lset_type(ll, new_arr.Temp, obj_type)
		}
		return LT{L: ll, Temp: new_arr.Temp}
	}
	if (method == "append") {
		elem := _lower_expr(ll)
		_ = elem
		ll = elem.L
		if (_lk(ll) == ")") {
			ll = _ladv(ll)
		}
		nt := _lnew_temp(ll)
		_ = nt
		ll = nt.L
		obj_type := _lget_type(ll, obj)
		_ = obj_type
		elem_type := _lget_type(ll, elem.Temp)
		_ = elem_type
		if ((elem_type == int64(12)) || (obj_type == int64(13))) {
			ll = _lemit(ll, new_inst(IrOpOpArrayAppend{}, nt.Temp, obj, elem.Temp, "", int64(12)))
			ll = _lset_type(ll, nt.Temp, int64(13))
		} else {
			ll = _lemit(ll, new_inst(IrOpOpArrayAppend{}, nt.Temp, obj, elem.Temp, "", int64(0)))
			if (obj_type > int64(0)) {
				ll = _lset_type(ll, nt.Temp, obj_type)
			}
		}
		return LT{L: ll, Temp: nt.Temp}
	}
	if (method == "contains") {
		arg := _lower_expr(ll)
		_ = arg
		ll = arg.L
		if (_lk(ll) == ")") {
			ll = _ladv(ll)
		}
		nt := _lnew_temp(ll)
		_ = nt
		ll = nt.L
		ll = _lemit(ll, new_inst(IrOpOpStrContains{}, nt.Temp, obj, arg.Temp, "", int64(0)))
		return LT{L: ll, Temp: nt.Temp}
	}
	if (method == "charAt") {
		arg := _lower_expr(ll)
		_ = arg
		ll = arg.L
		if (_lk(ll) == ")") {
			ll = _ladv(ll)
		}
		nt := _lnew_str_temp(ll)
		_ = nt
		ll = nt.L
		ll = _lemit(ll, new_inst(IrOpOpStrCharAt{}, nt.Temp, obj, arg.Temp, "", int64(0)))
		ll = _lset_type(ll, nt.Temp, int64(12))
		return LT{L: ll, Temp: nt.Temp}
	}
	if (method == "startsWith") {
		arg := _lower_expr(ll)
		_ = arg
		ll = arg.L
		if (_lk(ll) == ")") {
			ll = _ladv(ll)
		}
		nt := _lnew_temp(ll)
		_ = nt
		ll = nt.L
		ll = _lemit(ll, new_inst(IrOpOpStrStartsWith{}, nt.Temp, obj, arg.Temp, "", int64(0)))
		return LT{L: ll, Temp: nt.Temp}
	}
	if (method == "endsWith") {
		arg := _lower_expr(ll)
		_ = arg
		ll = arg.L
		if (_lk(ll) == ")") {
			ll = _ladv(ll)
		}
		nt := _lnew_temp(ll)
		_ = nt
		ll = nt.L
		ll = _lemit(ll, new_inst(IrOpOpStrEndsWith{}, nt.Temp, obj, arg.Temp, "", int64(0)))
		return LT{L: ll, Temp: nt.Temp}
	}
	if (method == "indexOf") {
		arg := _lower_expr(ll)
		_ = arg
		ll = arg.L
		if (_lk(ll) == ")") {
			ll = _ladv(ll)
		}
		nt := _lnew_temp(ll)
		_ = nt
		ll = nt.L
		ll = _lemit(ll, new_inst(IrOpOpStrIndexOf{}, nt.Temp, obj, arg.Temp, "", int64(0)))
		ll = _lset_type(ll, nt.Temp, int64(1))
		return LT{L: ll, Temp: nt.Temp}
	}
	if (method == "trim") {
		for ((_lk(ll) != ")") && (_lk(ll) != "EOF")) {
			ll = _ladv(ll)
		}
		if (_lk(ll) == ")") {
			ll = _ladv(ll)
		}
		nt := _lnew_str_temp(ll)
		_ = nt
		ll = nt.L
		ll = _lemit(ll, new_inst(IrOpOpStrTrim{}, nt.Temp, obj, int64(0), "", int64(0)))
		ll = _lset_type(ll, nt.Temp, int64(12))
		return LT{L: ll, Temp: nt.Temp}
	}
	if (method == "replace") {
		old_arg := _lower_expr(ll)
		_ = old_arg
		ll = old_arg.L
		if (_lk(ll) == ",") {
			ll = _ladv(ll)
		}
		new_arg := _lower_expr(ll)
		_ = new_arg
		ll = new_arg.L
		if (_lk(ll) == ")") {
			ll = _ladv(ll)
		}
		nt := _lnew_str_temp(ll)
		_ = nt
		ll = nt.L
		ll = _lemit(ll, new_inst(IrOpOpStrReplace{}, nt.Temp, obj, old_arg.Temp, "", new_arg.Temp))
		ll = _lset_type(ll, nt.Temp, int64(12))
		return LT{L: ll, Temp: nt.Temp}
	}
	if (method == "toLower") {
		for ((_lk(ll) != ")") && (_lk(ll) != "EOF")) {
			ll = _ladv(ll)
		}
		if (_lk(ll) == ")") {
			ll = _ladv(ll)
		}
		nt := _lnew_str_temp(ll)
		_ = nt
		ll = nt.L
		ll = _lemit(ll, new_inst(IrOpOpStrToLower{}, nt.Temp, obj, int64(0), "", int64(0)))
		ll = _lset_type(ll, nt.Temp, int64(12))
		return LT{L: ll, Temp: nt.Temp}
	}
	if (method == "toUpper") {
		for ((_lk(ll) != ")") && (_lk(ll) != "EOF")) {
			ll = _ladv(ll)
		}
		if (_lk(ll) == ")") {
			ll = _ladv(ll)
		}
		nt := _lnew_str_temp(ll)
		_ = nt
		ll = nt.L
		ll = _lemit(ll, new_inst(IrOpOpStrToUpper{}, nt.Temp, obj, int64(0), "", int64(0)))
		ll = _lset_type(ll, nt.Temp, int64(12))
		return LT{L: ll, Temp: nt.Temp}
	}
	if (method == "split") {
		arg := _lower_expr(ll)
		_ = arg
		ll = arg.L
		if (_lk(ll) == ")") {
			ll = _ladv(ll)
		}
		nt := _lnew_temp(ll)
		_ = nt
		ll = nt.L
		ll = _lemit(ll, new_inst(IrOpOpStrSplit{}, nt.Temp, obj, arg.Temp, "", int64(0)))
		ll = _lset_type(ll, nt.Temp, int64(13))
		return LT{L: ll, Temp: nt.Temp}
	}
	if (method == "substring") {
		start := _lower_expr(ll)
		_ = start
		ll = start.L
		if (_lk(ll) == ",") {
			ll = _ladv(ll)
		}
		end := _lower_expr(ll)
		_ = end
		ll = end.L
		if (_lk(ll) == ")") {
			ll = _ladv(ll)
		}
		nt := _lnew_str_temp(ll)
		_ = nt
		ll = nt.L
		ll = _lemit(ll, new_inst(IrOpOpStrSubstring{}, nt.Temp, obj, start.Temp, "", end.Temp))
		ll = _lset_type(ll, nt.Temp, int64(12))
		return LT{L: ll, Temp: nt.Temp}
	}
	if (method == "debug_str") {
		for ((_lk(ll) != ")") && (_lk(ll) != "EOF")) {
			ll = _ladv(ll)
		}
		if (_lk(ll) == ")") {
			ll = _ladv(ll)
		}
		obj_type := _lget_type(ll, obj)
		_ = obj_type
		sname := _struct_name_for_type(ll.Store, ll.Registry, obj_type)
		_ = sname
		if (sname != "") {
			def := reg_find_struct(ll.Registry, sname)
			_ = def
			hdr_idx := mod_find_string(ll.Module, (sname + "{"))
			_ = hdr_idx
			if (hdr_idx == int64(0)) {
				new_mod := mod_add_string(ll.Module, (sname + "{"))
				_ = new_mod
				hdr_idx = (int64(len(new_mod.String_constants)) - int64(1))
				ll = _lset_mod(ll, new_mod)
			}
			hdr_len := (int64(len(sname)) + int64(1))
			_ = hdr_len
			result_nt := _lnew_str_temp(ll)
			_ = result_nt
			ll = result_nt.L
			ll = _lemit(ll, new_inst(IrOpOpConstStr{}, result_nt.Temp, hdr_idx, hdr_len, "", int64(0)))
			ll = _lset_type(ll, result_nt.Temp, int64(12))
			fi := int64(1)
			_ = fi
			field_slot := int64(0)
			_ = field_slot
			for (fi < int64(len(def.Fields))) {
				fld := def.Fields[fi]
				_ = fld
				ft := fld.Type_id
				_ = ft
				if (fi > int64(1)) {
					sep_idx := mod_find_string(ll.Module, ", ")
					_ = sep_idx
					if (sep_idx == int64(0)) {
						new_mod := mod_add_string(ll.Module, ", ")
						_ = new_mod
						sep_idx = (int64(len(new_mod.String_constants)) - int64(1))
						ll = Lowerer{Tokens: ll.Tokens, Pos: ll.Pos, Store: ll.Store, Registry: ll.Registry, Treg: ll.Treg, Pool: ll.Pool, Table: ll.Table, File: ll.File, Module: new_mod, Current_func: ll.Current_func, Temp_counter: ll.Temp_counter, Label_counter: ll.Label_counter, Env_names: ll.Env_names, Env_slots: ll.Env_slots, Loop_start: ll.Loop_start, Loop_end: ll.Loop_end, Const_names: ll.Const_names, Const_vals: ll.Const_vals, Const_str_names: ll.Const_str_names, Const_str_vals: ll.Const_str_vals, Temp_types: ll.Temp_types, Fn_error_type: ll.Fn_error_type, Defer_starts: ll.Defer_starts, Defer_ends: ll.Defer_ends}
					}
					sep_nt := _lnew_str_temp(ll)
					_ = sep_nt
					ll = sep_nt.L
					ll = _lemit(ll, new_inst(IrOpOpConstStr{}, sep_nt.Temp, sep_idx, int64(2), "", int64(0)))
					concat_nt := _lnew_str_temp(ll)
					_ = concat_nt
					ll = concat_nt.L
					ll = _lemit(ll, new_inst(IrOpOpStrConcat{}, concat_nt.Temp, result_nt.Temp, sep_nt.Temp, "", int64(0)))
					result_nt = concat_nt
				}
				fname_str := (fld.Name + ": ")
				_ = fname_str
				fname_idx := mod_find_string(ll.Module, fname_str)
				_ = fname_idx
				if (fname_idx == int64(0)) {
					new_mod := mod_add_string(ll.Module, fname_str)
					_ = new_mod
					fname_idx = (int64(len(new_mod.String_constants)) - int64(1))
					ll = Lowerer{Tokens: ll.Tokens, Pos: ll.Pos, Store: ll.Store, Registry: ll.Registry, Treg: ll.Treg, Pool: ll.Pool, Table: ll.Table, File: ll.File, Module: new_mod, Current_func: ll.Current_func, Temp_counter: ll.Temp_counter, Label_counter: ll.Label_counter, Env_names: ll.Env_names, Env_slots: ll.Env_slots, Loop_start: ll.Loop_start, Loop_end: ll.Loop_end, Const_names: ll.Const_names, Const_vals: ll.Const_vals, Const_str_names: ll.Const_str_names, Const_str_vals: ll.Const_str_vals, Temp_types: ll.Temp_types, Fn_error_type: ll.Fn_error_type, Defer_starts: ll.Defer_starts, Defer_ends: ll.Defer_ends}
				}
				fname_nt := _lnew_str_temp(ll)
				_ = fname_nt
				ll = fname_nt.L
				ll = _lemit(ll, new_inst(IrOpOpConstStr{}, fname_nt.Temp, fname_idx, int64(len(fname_str)), "", int64(0)))
				c1 := _lnew_str_temp(ll)
				_ = c1
				ll = c1.L
				ll = _lemit(ll, new_inst(IrOpOpStrConcat{}, c1.Temp, result_nt.Temp, fname_nt.Temp, "", int64(0)))
				if (ft == int64(12)) {
					fval := _lnew_str_temp(ll)
					_ = fval
					ll = fval.L
					ll = _lemit(ll, new_inst(IrOpOpFieldGet{}, fval.Temp, obj, field_slot, fld.Name, int64(0)))
					fidx_len := (field_slot + int64(1))
					_ = fidx_len
					ll = _lemit(ll, new_inst(IrOpOpFieldGet{}, (fval.Temp + int64(1)), obj, fidx_len, fld.Name, int64(0)))
					c2 := _lnew_str_temp(ll)
					_ = c2
					ll = c2.L
					ll = _lemit(ll, new_inst(IrOpOpStrConcat{}, c2.Temp, c1.Temp, fval.Temp, "", int64(0)))
					result_nt = c2
					field_slot = (field_slot + int64(2))
				} else {
					fval := _lnew_temp(ll)
					_ = fval
					ll = fval.L
					ll = _lemit(ll, new_inst(IrOpOpFieldGet{}, fval.Temp, obj, field_slot, fld.Name, int64(0)))
					fstr := _lnew_str_temp(ll)
					_ = fstr
					ll = fstr.L
					ll = _lemit(ll, new_inst(IrOpOpIntToStr{}, fstr.Temp, fval.Temp, int64(0), "", int64(0)))
					c2 := _lnew_str_temp(ll)
					_ = c2
					ll = c2.L
					ll = _lemit(ll, new_inst(IrOpOpStrConcat{}, c2.Temp, c1.Temp, fstr.Temp, "", int64(0)))
					result_nt = c2
					field_slot = (field_slot + int64(1))
				}
				fi = (fi + int64(1))
			}
			close_idx := mod_find_string(ll.Module, "}")
			_ = close_idx
			if (close_idx == int64(0)) {
				new_mod := mod_add_string(ll.Module, "}")
				_ = new_mod
				close_idx = (int64(len(new_mod.String_constants)) - int64(1))
				ll = _lset_mod(ll, new_mod)
			}
			close_nt := _lnew_str_temp(ll)
			_ = close_nt
			ll = close_nt.L
			ll = _lemit(ll, new_inst(IrOpOpConstStr{}, close_nt.Temp, close_idx, int64(1), "", int64(0)))
			final_nt := _lnew_str_temp(ll)
			_ = final_nt
			ll = final_nt.L
			ll = _lemit(ll, new_inst(IrOpOpStrConcat{}, final_nt.Temp, result_nt.Temp, close_nt.Temp, "", int64(0)))
			ll = _lset_type(ll, final_nt.Temp, int64(12))
			return LT{L: ll, Temp: final_nt.Temp}
		}
		nt := _lnew_str_temp(ll)
		_ = nt
		ll = nt.L
		ll = _lset_type(ll, nt.Temp, int64(12))
		return LT{L: ll, Temp: nt.Temp}
	}
	if ((method == "toString") || (method == "toStr")) {
		for ((_lk(ll) != ")") && (_lk(ll) != "EOF")) {
			ll = _ladv(ll)
		}
		if (_lk(ll) == ")") {
			ll = _ladv(ll)
		}
		nt := _lnew_str_temp(ll)
		_ = nt
		ll = nt.L
		obj_type := _lget_type(ll, obj)
		_ = obj_type
		if (obj_type == int64(12)) {
			return LT{L: ll, Temp: obj}
		}
		if _is_float_type(obj_type) {
			ll = _lemit(ll, new_inst(IrOpOpIntToStr{}, nt.Temp, obj, int64(0), "f64", int64(0)))
		} else {
			ll = _lemit(ll, new_inst(IrOpOpIntToStr{}, nt.Temp, obj, int64(0), "", int64(0)))
		}
		ll = _lset_type(ll, nt.Temp, int64(12))
		return LT{L: ll, Temp: nt.Temp}
	}
	dyn_obj_type := _lget_type(ll, obj)
	_ = dyn_obj_type
	if (dyn_obj_type == int64(500)) {
		method_idx := int64(0)
		_ = method_idx
		ti := int64(1)
		_ = ti
		for (ti < int64(len(ll.Treg.Trait_defs))) {
			tdef := ll.Treg.Trait_defs[ti]
			_ = tdef
			mi2 := int64(1)
			_ = mi2
			for (mi2 < int64(len(tdef.Method_names))) {
				if (tdef.Method_names[mi2] == method) {
					method_idx = (mi2 - int64(1))
				}
				mi2 = (mi2 + int64(1))
			}
			ti = (ti + int64(1))
		}
		extra_arg_temps := []int64{int64(0)}
		_ = extra_arg_temps
		extra_count := int64(0)
		_ = extra_count
		for ((_lk(ll) != ")") && (_lk(ll) != "EOF")) {
			ll = _lskip_nl(ll)
			if (_lk(ll) == ")") {
				return LT{L: ll, Temp: int64(0)}
			}
			arg := _lower_expr(ll)
			_ = arg
			ll = arg.L
			extra_arg_temps = append(extra_arg_temps, arg.Temp)
			extra_count = (extra_count + int64(1))
			ll = _lskip_nl(ll)
			if (_lk(ll) == ",") {
				ll = _ladv(ll)
			}
		}
		if (_lk(ll) == ")") {
			ll = _ladv(ll)
		}
		m_ret := int64(0)
		_ = m_ret
		fi := int64(1)
		_ = fi
		for (fi < int64(len(ll.Registry.Fn_sigs))) {
			sig := ll.Registry.Fn_sigs[fi]
			_ = sig
			slen := int64(len(method))
			_ = slen
			nlen := int64(len(sig.Name))
			_ = nlen
			if (nlen > (slen + int64(1))) {
				suffix := sig.Name[((nlen - slen) - int64(1)):nlen]
				_ = suffix
				if (suffix == ("_" + method)) {
					m_ret = sig.Return_type
					fi = int64(len(ll.Registry.Fn_sigs))
				}
			}
			fi = (fi + int64(1))
		}
		nt_temp := int64(0)
		_ = nt_temp
		if (m_ret == int64(12)) {
			nt := _lnew_str_temp(ll)
			_ = nt
			ll = nt.L
			nt_temp = nt.Temp
			ll = _lset_type(ll, nt.Temp, int64(12))
		} else {
			nt := _lnew_temp(ll)
			_ = nt
			ll = nt.L
			nt_temp = nt.Temp
			if (m_ret > int64(0)) {
				ll = _lset_type(ll, nt.Temp, m_ret)
			}
		}
		ll = _lemit(ll, new_inst(IrOpOpVtableCall{}, nt_temp, obj, method_idx, method, m_ret))
		return LT{L: ll, Temp: nt_temp}
	}
	obj_type := _lget_type(ll, obj)
	_ = obj_type
	obj_sname := _struct_name_for_type(ll.Store, ll.Registry, obj_type)
	_ = obj_sname
	qualified_method := method
	_ = qualified_method
	if (obj_sname != "") {
		qualified_method = ((obj_sname + "_") + method)
	}
	method_sig := reg_find_fn(ll.Registry, qualified_method)
	_ = method_sig
	if (method_sig.Name == "") {
		method_sig = reg_find_fn(ll.Registry, method)
		qualified_method = method
	}
	if ((method_sig.Name == "") && (obj_sname != "")) {
		ii := int64(1)
		_ = ii
		for (ii < int64(len(ll.Treg.Impl_defs))) {
			bimp := ll.Treg.Impl_defs[ii]
			_ = bimp
			if ((int64(len(bimp.Type_name)) > int64(1)) && (string(bimp.Type_name[int64(0)]) == "*")) {
				bound := bimp.Type_name[int64(1):int64(len(bimp.Type_name))]
				_ = bound
				if treg_has_impl(ll.Treg, obj_sname, bound) {
					spec_fn := ((("_" + method) + "_") + obj_sname)
					_ = spec_fn
					spec_fn_check := mod_find_func(ll.Module, spec_fn)
					_ = spec_fn_check
					if (spec_fn_check.Name != "") {
						blanket_fn := ("_" + method)
						_ = blanket_fn
						method_sig = reg_find_fn(ll.Registry, blanket_fn)
						if (method_sig.Name != "") {
							qualified_method = spec_fn
							ii = int64(len(ll.Treg.Impl_defs))
						}
					} else {
						blanket_fn := ("_" + method)
						_ = blanket_fn
						method_sig = reg_find_fn(ll.Registry, blanket_fn)
						if (method_sig.Name != "") {
							qualified_method = blanket_fn
							ii = int64(len(ll.Treg.Impl_defs))
						}
					}
				}
			}
			ii = (ii + int64(1))
		}
	}
	if (method_sig.Name != "") {
		first_arg := obj
		_ = first_arg
		arg_count := int64(1)
		_ = arg_count
		for ((_lk(ll) != ")") && (_lk(ll) != "EOF")) {
			ll = _lskip_nl(ll)
			if (_lk(ll) == ")") {
				return LT{L: _ladv(ll), Temp: obj}
			}
			arg := _lower_expr(ll)
			_ = arg
			ll = arg.L
			arg_count = (arg_count + int64(1))
			ll = _lskip_nl(ll)
			if (_lk(ll) == ",") {
				ll = _ladv(ll)
			}
		}
		if (_lk(ll) == ")") {
			ll = _ladv(ll)
		}
		if fn_is_generic(method_sig) {
			if (obj_sname != "") {
				bspec := ((("_" + method) + "_") + obj_sname)
				_ = bspec
				bsig := reg_find_fn(ll.Registry, bspec)
				_ = bsig
				if ((bsig.Name != "") && fn_is_generic(bsig)) {
					method_sig = bsig
					qualified_method = bspec
				}
			}
			resolved_suffix := ""
			_ = resolved_suffix
			gi := int64(2)
			_ = gi
			for ((gi < int64(len(method_sig.Param_types))) && (gi < (arg_count + int64(1)))) {
				pt := method_sig.Param_types[gi]
				_ = pt
				if ((pt > int64(0)) && (pt < int64(len(ll.Store.Types)))) {
					ptinfo := store_get(ll.Store, pt)
					_ = ptinfo
					if (type_kind_name(ptinfo.Kind) == "TypeVar") {
						at := _lget_type(ll, ((first_arg + gi) - int64(1)))
						_ = at
						suffix := _type_name_suffix_l(ll.Store, ll.Registry, at)
						_ = suffix
						if (resolved_suffix == "") {
							resolved_suffix = suffix
						}
					}
				}
				gi = (gi + int64(1))
			}
			if (resolved_suffix != "") {
				spec_call := ((qualified_method + "_") + resolved_suffix)
				_ = spec_call
				sp := reg_find_fn(ll.Registry, spec_call)
				_ = sp
				if (sp.Name != "") {
					qualified_method = spec_call
					method_sig = sp
				} else {
					at := _lget_type(ll, (first_arg + int64(1)))
					_ = at
					sp_ptypes := []int64{int64(0)}
					_ = sp_ptypes
					spi := int64(1)
					_ = spi
					for (spi < int64(len(method_sig.Param_types))) {
						pt := method_sig.Param_types[spi]
						_ = pt
						if ((pt > int64(0)) && (pt < int64(len(ll.Store.Types)))) {
							pti := store_get(ll.Store, pt)
							_ = pti
							if (type_kind_name(pti.Kind) == "TypeVar") {
								pt = at
							}
						}
						sp_ptypes = append(sp_ptypes, pt)
						spi = (spi + int64(1))
					}
					sp_ret := method_sig.Return_type
					_ = sp_ret
					if ((sp_ret > int64(0)) && (sp_ret < int64(len(ll.Store.Types)))) {
						rti := store_get(ll.Store, sp_ret)
						_ = rti
						if (type_kind_name(rti.Kind) == "TypeVar") {
							sp_ret = at
						}
					}
					new_sig := new_fn_sig(spec_call, method_sig.Param_names, sp_ptypes, sp_ret, method_sig.Error_type)
					_ = new_sig
					ll = _lset_reg(ll, reg_add_fn(ll.Registry, new_sig))
					spec := MonoSpec{Generic_name: qualified_method, Specialized_name: spec_call, Type_arg_0: at, Type_arg_1: int64(0), Type_arg_2: int64(0)}
					_ = spec
					ll = _lset_reg(ll, reg_add_mono(ll.Registry, spec))
					qualified_method = spec_call
					method_sig = new_sig
				}
			}
		}
		nt_temp := int64(0)
		_ = nt_temp
		if (method_sig.Return_type == int64(12)) {
			nt := _lnew_str_temp(ll)
			_ = nt
			ll = nt.L
			nt_temp = nt.Temp
		} else {
			nt := _lnew_temp(ll)
			_ = nt
			ll = nt.L
			nt_temp = nt.Temp
		}
		ll = _lemit(ll, new_inst(IrOpOpCall{}, nt_temp, first_arg, arg_count, qualified_method, int64(0)))
		if (method_sig.Return_type == int64(12)) {
			ll = _lset_type(ll, nt_temp, int64(12))
		} else if (method_sig.Return_type > int64(16)) {
			ret_sname := format_type(ll.Store, method_sig.Return_type)
			_ = ret_sname
			ret_st_id := _struct_type_id(ll.Registry, ret_sname)
			_ = ret_st_id
			if (ret_st_id > int64(0)) {
				ll = _lset_type(ll, nt_temp, ret_st_id)
			}
		}
		return LT{L: ll, Temp: nt_temp}
	}
	fmt.Println(((((((("warning: unresolved method '" + method) + "' on type '") + obj_sname) + "' at ") + ll.File) + ":") + i2s(_lcur(ll).Line)))
	for ((_lk(ll) != ")") && (_lk(ll) != "EOF")) {
		ll = _ladv(ll)
	}
	if (_lk(ll) == ")") {
		ll = _ladv(ll)
	}
	nt := _lnew_temp(ll)
	_ = nt
	ll = nt.L
	ll = _lemit(ll, new_inst(IrOpOpConst{}, nt.Temp, int64(0), int64(0), "", int64(0)))
	return LT{L: ll, Temp: nt.Temp}
}

func _lower_call_args(l Lowerer, callee int64) LT {
	ll := _ladv(l)
	_ = ll
	ll = _lskip_nl(ll)
	for ((_lk(ll) != ")") && (_lk(ll) != "EOF")) {
		ll = _lskip_nl(ll)
		arg := _lower_expr(ll)
		_ = arg
		ll = arg.L
		ll = _lskip_nl(ll)
		if (_lk(ll) == ",") {
			ll = _ladv(ll)
		}
	}
	if (_lk(ll) == ")") {
		ll = _ladv(ll)
	}
	return LT{L: ll, Temp: callee}
}

func _field_index(reg TypeRegistry, struct_name string, field_name string) int64 {
	def := reg_find_struct(reg, struct_name)
	_ = def
	if (def.Name == "") {
		return int64(0)
	}
	slot := int64(0)
	_ = slot
	i := int64(1)
	_ = i
	for (i < int64(len(def.Fields))) {
		if (def.Fields[i].Name == field_name) {
			return slot
		}
		if (def.Fields[i].Type_id == int64(12)) {
			slot = (slot + int64(2))
		} else {
			slot = (slot + int64(1))
		}
		i = (i + int64(1))
	}
	return int64(0)
}

func _field_index_any(reg TypeRegistry, field_name string) int64 {
	si := int64(1)
	_ = si
	for (si < int64(len(reg.Struct_defs))) {
		def := reg.Struct_defs[si]
		_ = def
		if (def.Name != "") {
			if (def.Is_sum == false) {
				fi := int64(1)
				_ = fi
				for (fi < int64(len(def.Fields))) {
					if (def.Fields[fi].Name == field_name) {
						return _field_index(reg, def.Name, field_name)
					}
					fi = (fi + int64(1))
				}
			}
		}
		si = (si + int64(1))
	}
	return int64(0)
}

func _struct_name_from_type(reg TypeRegistry, type_id int64) string {
	idx := ((type_id - TY_STRUCT_BASE) + int64(1))
	_ = idx
	if ((idx >= int64(1)) && (idx < int64(len(reg.Struct_defs)))) {
		return reg.Struct_defs[idx].Name
	}
	return ""
}

func _resolve_generic_call_name(store TypeStore, reg TypeRegistry, base_name string, sig FnSig, arg_types []int64) string {
	resolved_tvs := []string{""}
	_ = resolved_tvs
	resolved_tys := []string{""}
	_ = resolved_tys
	gi := int64(1)
	_ = gi
	for ((gi < int64(len(arg_types))) && (gi < int64(len(sig.Param_types)))) {
		pt := sig.Param_types[gi]
		_ = pt
		at := arg_types[gi]
		_ = at
		if ((pt > int64(0)) && (pt < int64(len(store.Types)))) {
			ptinfo := store_get(store, pt)
			_ = ptinfo
			if (type_kind_name(ptinfo.Kind) == "TypeVar") {
				tv_name := ptinfo.Name
				_ = tv_name
				already := false
				_ = already
				ri := int64(1)
				_ = ri
				for (ri < int64(len(resolved_tvs))) {
					if (resolved_tvs[ri] == tv_name) {
						already = true
					}
					ri = (ri + int64(1))
				}
				if (already == false) {
					suffix := _type_name_suffix_l(store, reg, at)
					_ = suffix
					resolved_tvs = append(resolved_tvs, tv_name)
					resolved_tys = append(resolved_tys, suffix)
				}
			}
		}
		gi = (gi + int64(1))
	}
	result := base_name
	_ = result
	si := int64(1)
	_ = si
	for (si < int64(len(resolved_tys))) {
		if (si == int64(1)) {
			result = ((result + "_") + resolved_tys[si])
		} else {
			result = ((result + "_") + resolved_tys[si])
		}
		si = (si + int64(1))
	}
	return result
}

func _type_name_suffix_l(store TypeStore, reg TypeRegistry, tid int64) string {
	if (tid >= TY_STRUCT_BASE) {
		sname := _struct_name_from_type(reg, tid)
		_ = sname
		if (sname != "") {
			return sname
		}
	}
	return _type_name_suffix(store, tid)
}

func _struct_type_id(reg TypeRegistry, struct_name string) int64 {
	i := int64(1)
	_ = i
	for (i < int64(len(reg.Struct_defs))) {
		if (reg.Struct_defs[i].Name == struct_name) {
			return ((TY_STRUCT_BASE - int64(1)) + i)
		}
		i = (i + int64(1))
	}
	return int64(0)
}

func _field_type_id_from_name(reg TypeRegistry, struct_name string, field_name string) int64 {
	def := reg_find_struct(reg, struct_name)
	_ = def
	if (def.Name == "") {
		return int64(0)
	}
	i := int64(1)
	_ = i
	for (i < int64(len(def.Fields))) {
		if (def.Fields[i].Name == field_name) {
			return def.Fields[i].Type_id
		}
		i = (i + int64(1))
	}
	return int64(0)
}

func _field_type_id(l Lowerer, struct_name string, field_name string) int64 {
	def := reg_find_struct(l.Registry, struct_name)
	_ = def
	if (def.Name == "") {
		return int64(0)
	}
	i := int64(1)
	_ = i
	for (i < int64(len(def.Fields))) {
		if (def.Fields[i].Name == field_name) {
			return def.Fields[i].Type_id
		}
		i = (i + int64(1))
	}
	return int64(0)
}

func _field_type_id_any(reg TypeRegistry, field_name string) int64 {
	si := int64(1)
	_ = si
	for (si < int64(len(reg.Struct_defs))) {
		def := reg.Struct_defs[si]
		_ = def
		if (def.Name != "") {
			if (def.Is_sum == false) {
				fi := int64(1)
				_ = fi
				for (fi < int64(len(def.Fields))) {
					if (def.Fields[fi].Name == field_name) {
						return def.Fields[fi].Type_id
					}
					fi = (fi + int64(1))
				}
			}
		}
		si = (si + int64(1))
	}
	return int64(0)
}

func _struct_field_count(reg TypeRegistry, struct_name string) int64 {
	def := reg_find_struct(reg, struct_name)
	_ = def
	if (def.Name == "") {
		return int64(0)
	}
	slots := int64(0)
	_ = slots
	i := int64(1)
	_ = i
	for (i < int64(len(def.Fields))) {
		if (def.Fields[i].Type_id == int64(12)) {
			slots = (slots + int64(2))
		} else {
			slots = (slots + int64(1))
		}
		i = (i + int64(1))
	}
	return slots
}

func _lower_struct_lit(l Lowerer, name string) LT {
	lb := "{"
	_ = lb
	rb := "}"
	_ = rb
	ll := _ladv(l)
	_ = ll
	ll = _lskip_nl(ll)
	field_count := _struct_field_count(ll.Registry, name)
	_ = field_count
	nt := _lnew_temp(ll)
	_ = nt
	ll = nt.L
	ll = _lemit(ll, new_inst(IrOpOpAlloc{}, nt.Temp, field_count, int64(0), name, int64(0)))
	st_id := _struct_type_id(ll.Registry, name)
	_ = st_id
	if (st_id > int64(0)) {
		ll = _lset_type(ll, nt.Temp, st_id)
	}
	for ((_lk(ll) != rb) && (_lk(ll) != "EOF")) {
		if (_lk(ll) == "IDENT") {
			fname := _lcur(ll).Text
			_ = fname
			ll = _ladv(ll)
			if (_lk(ll) == ":") {
				ll = _ladv(ll)
				val := _lower_expr(ll)
				_ = val
				ll = val.L
				fidx := _field_index(ll.Registry, name, fname)
				_ = fidx
				val_type := _lget_type(ll, val.Temp)
				_ = val_type
				ft := _field_type_id(ll, name, fname)
				_ = ft
				if ((val_type == int64(12)) || (ft == int64(12))) {
					fidx_len := (fidx + int64(1))
					_ = fidx_len
					val_len := (val.Temp + int64(1))
					_ = val_len
					ll = _lemit(ll, new_inst(IrOpOpFieldSet{}, nt.Temp, val.Temp, fidx, fname, int64(0)))
					ll = _lemit(ll, new_inst(IrOpOpFieldSet{}, nt.Temp, val_len, fidx_len, fname, int64(0)))
				} else {
					ll = _lemit(ll, new_inst(IrOpOpFieldSet{}, nt.Temp, val.Temp, fidx, fname, int64(0)))
				}
			}
		} else {
			ll = _ladv(ll)
		}
		ll = _lskip_nl(ll)
		if (_lk(ll) == ",") {
			ll = _ladv(ll)
			ll = _lskip_nl(ll)
		}
	}
	if (_lk(ll) == rb) {
		ll = _ladv(ll)
	}
	return LT{L: ll, Temp: nt.Temp}
}

func _lower_atom(l Lowerer) LT {
	lb := "{"
	_ = lb
	result := _lower_primary(l)
	_ = result
	ll := result.L
	_ = ll
	left := result.Temp
	_ = left
	running := true
	_ = running
	for running {
		k := _lk(ll)
		_ = k
		if (k == ".") {
			ll = _ladv(ll)
			if (_lk(ll) == "IDENT") {
				field := _lcur(ll).Text
				_ = field
				ll = _ladv(ll)
				if (_lk(ll) == "(") {
					mr := _lower_method_call(ll, left, field)
					_ = mr
					ll = mr.L
					left = mr.Temp
				} else if ((_lk(ll) == "[") && (((field == "to") || (field == "trunc")))) {
					ll = _ladv(ll)
					target_type_name := ""
					_ = target_type_name
					if (_lk(ll) == "IDENT") {
						target_type_name = _lcur(ll).Text
						ll = _ladv(ll)
					}
					if (_lk(ll) == "]") {
						ll = _ladv(ll)
					}
					if (_lk(ll) == "(") {
						ll = _ladv(ll)
					}
					if (_lk(ll) == ")") {
						ll = _ladv(ll)
					}
					mr := _lower_conversion(ll, left, field, target_type_name)
					_ = mr
					ll = mr.L
					left = mr.Temp
				} else {
					left_type := _lget_type(ll, left)
					_ = left_type
					sname := _struct_name_for_type(ll.Store, ll.Registry, left_type)
					_ = sname
					fidx := int64(0)
					_ = fidx
					ft := int64(0)
					_ = ft
					if (sname != "") {
						fidx = _field_index(ll.Registry, sname, field)
						ft = _field_type_id(ll, sname, field)
					} else {
						fidx = _field_index_any(ll.Registry, field)
						ft = _field_type_id_any(ll.Registry, field)
					}
					nt_temp := int64(0)
					_ = nt_temp
					if (ft == int64(12)) {
						nt := _lnew_str_temp(ll)
						_ = nt
						ll = nt.L
						nt_temp = nt.Temp
					} else {
						nt := _lnew_temp(ll)
						_ = nt
						ll = nt.L
						nt_temp = nt.Temp
					}
					if (ft == int64(12)) {
						fidx_len := (fidx + int64(1))
						_ = fidx_len
						nt_len := (nt_temp + int64(1))
						_ = nt_len
						ll = _lemit(ll, new_inst(IrOpOpFieldGet{}, nt_temp, left, fidx, field, int64(0)))
						ll = _lemit(ll, new_inst(IrOpOpFieldGet{}, nt_len, left, fidx_len, field, int64(0)))
					} else {
						ll = _lemit(ll, new_inst(IrOpOpFieldGet{}, nt_temp, left, fidx, field, int64(0)))
					}
					if (ft > int64(0)) {
						ll = _lset_type(ll, nt_temp, ft)
					}
					left = nt_temp
				}
			} else {
				running = false
			}
		} else if (k == "[") {
			ll = _ladv(ll)
			idx := _lower_expr(ll)
			_ = idx
			ll = idx.L
			if (_lk(ll) == "]") {
				ll = _ladv(ll)
			}
			left_type := _lget_type(ll, left)
			_ = left_type
			elem_type := _array_elem_type(ll.Store, left_type)
			_ = elem_type
			if (elem_type == int64(12)) {
				nt := _lnew_str_temp(ll)
				_ = nt
				ll = nt.L
				ll = _lemit(ll, new_inst(IrOpOpArrayGet{}, nt.Temp, left, idx.Temp, "", int64(12)))
				ll = _lset_type(ll, nt.Temp, int64(12))
				left = nt.Temp
			} else {
				nt := _lnew_temp(ll)
				_ = nt
				ll = nt.L
				ll = _lemit(ll, new_inst(IrOpOpArrayGet{}, nt.Temp, left, idx.Temp, "", int64(0)))
				if (elem_type > int64(0)) {
					ll = _lset_type(ll, nt.Temp, elem_type)
				}
				left = nt.Temp
			}
		} else if (k == "(") {
			cr := _lower_call_args(ll, left)
			_ = cr
			ll = cr.L
			left = cr.Temp
		} else if (k == "?") {
			ll = _ladv(ll)
			tag_nt := _lnew_temp(ll)
			_ = tag_nt
			ll = tag_nt.L
			ll = _lemit(ll, new_inst(IrOpOpFieldGet{}, tag_nt.Temp, left, int64(0), "_tag", int64(0)))
			const_one := _lnew_temp(ll)
			_ = const_one
			ll = const_one.L
			ll = _lemit(ll, new_inst(IrOpOpConst{}, const_one.Temp, int64(1), int64(0), "", int64(0)))
			is_err := _lnew_temp(ll)
			_ = is_err
			ll = is_err.L
			ll = _lemit(ll, new_inst(IrOpOpEq{}, is_err.Temp, tag_nt.Temp, const_one.Temp, "", int64(0)))
			err_lbl := _lnew_label(ll)
			_ = err_lbl
			ll = err_lbl.L
			ok_lbl := _lnew_label(ll)
			_ = ok_lbl
			ll = ok_lbl.L
			ll = _lemit(ll, new_inst(IrOpOpBranchTrue{}, int64(0), is_err.Temp, err_lbl.Temp, "", int64(0)))
			ok_val := _lnew_temp(ll)
			_ = ok_val
			ll = ok_val.L
			ll = _lemit(ll, new_inst(IrOpOpFieldGet{}, ok_val.Temp, left, int64(1), "_val", int64(0)))
			ll = _lemit(ll, new_inst(IrOpOpJump{}, int64(0), ok_lbl.Temp, int64(0), "", int64(0)))
			ll = _lemit(ll, new_inst(IrOpOpLabel{}, int64(0), err_lbl.Temp, int64(0), "", int64(0)))
			ll = _emit_defers(ll)
			ll = _lemit(ll, new_inst(IrOpOpRet{}, int64(0), left, int64(0), "", int64(0)))
			ll = _lemit(ll, new_inst(IrOpOpLabel{}, int64(0), ok_lbl.Temp, int64(0), "", int64(0)))
			left = ok_val.Temp
		} else if (k == "!") {
			ll = _ladv(ll)
			tag_nt := _lnew_temp(ll)
			_ = tag_nt
			ll = tag_nt.L
			ll = _lemit(ll, new_inst(IrOpOpFieldGet{}, tag_nt.Temp, left, int64(0), "_tag", int64(0)))
			const_one := _lnew_temp(ll)
			_ = const_one
			ll = const_one.L
			ll = _lemit(ll, new_inst(IrOpOpConst{}, const_one.Temp, int64(1), int64(0), "", int64(0)))
			is_err := _lnew_temp(ll)
			_ = is_err
			ll = is_err.L
			ll = _lemit(ll, new_inst(IrOpOpEq{}, is_err.Temp, tag_nt.Temp, const_one.Temp, "", int64(0)))
			panic_lbl := _lnew_label(ll)
			_ = panic_lbl
			ll = panic_lbl.L
			ok_lbl := _lnew_label(ll)
			_ = ok_lbl
			ll = ok_lbl.L
			ll = _lemit(ll, new_inst(IrOpOpBranchTrue{}, int64(0), is_err.Temp, panic_lbl.Temp, "", int64(0)))
			ok_val := _lnew_temp(ll)
			_ = ok_val
			ll = ok_val.L
			ll = _lemit(ll, new_inst(IrOpOpFieldGet{}, ok_val.Temp, left, int64(1), "_val", int64(0)))
			ll = _lemit(ll, new_inst(IrOpOpJump{}, int64(0), ok_lbl.Temp, int64(0), "", int64(0)))
			ll = _lemit(ll, new_inst(IrOpOpLabel{}, int64(0), panic_lbl.Temp, int64(0), "", int64(0)))
			panic_str_idx := mod_find_string(ll.Module, "unwrap failed on Err value")
			_ = panic_str_idx
			if (panic_str_idx == int64(0)) {
				new_mod := mod_add_string(ll.Module, "unwrap failed on Err value")
				_ = new_mod
				panic_str_idx = (int64(len(new_mod.String_constants)) - int64(1))
				ll = _lset_mod(ll, new_mod)
			}
			panic_ptr := _lnew_str_temp(ll)
			_ = panic_ptr
			ll = panic_ptr.L
			ll = _lemit(ll, new_inst(IrOpOpConstStr{}, panic_ptr.Temp, panic_str_idx, int64(25), "", int64(0)))
			ll = _lemit(ll, new_inst(IrOpOpCall{}, int64(0), panic_ptr.Temp, int64(2), "_aria_panic", int64(0)))
			ll = _lemit(ll, new_inst(IrOpOpLabel{}, int64(0), ok_lbl.Temp, int64(0), "", int64(0)))
			left = ok_val.Temp
		} else {
			running = false
		}
	}
	return LT{L: ll, Temp: left}
}

func _lower_if_cond(l Lowerer) LT {
	lb := "{"
	_ = lb
	result := _lower_atom(l)
	_ = result
	ll := result.L
	_ = ll
	left := result.Temp
	_ = left
	running := true
	_ = running
	for running {
		k := _lk(ll)
		_ = k
		if (((k == lb) || (k == "EOF")) || (k == "NEWLINE")) {
			running = false
		} else if ((((((k == "==") || (k == "!=")) || (k == "<")) || (k == ">")) || (k == "<=")) || (k == ">=")) {
			ll = _ladv(ll)
			right := _lower_additive(ll)
			_ = right
			ll = right.L
			nt := _lnew_temp(ll)
			_ = nt
			ll = nt.L
			left_type := _lget_type(ll, left)
			_ = left_type
			if ((((k == "==") || (k == "!="))) && (left_type == int64(12))) {
				ll = _lemit(ll, new_inst(IrOpOpStrEq{}, nt.Temp, left, right.Temp, "", int64(0)))
				if (k == "!=") {
					not_nt := _lnew_temp(ll)
					_ = not_nt
					ll = not_nt.L
					ll = _lemit(ll, new_inst(IrOpOpNot{}, not_nt.Temp, nt.Temp, int64(0), "", int64(0)))
					left = not_nt.Temp
				} else {
					left = nt.Temp
				}
			} else if (_is_float_type(left_type) || _is_float_type(_lget_type(ll, right.Temp))) {
				ll = _lemit(ll, new_inst(_cmp_op(k), nt.Temp, left, right.Temp, "f64", int64(0)))
				left = nt.Temp
			} else if (left_type == int64(12)) {
				cmp_nt := _lnew_temp(ll)
				_ = cmp_nt
				ll = cmp_nt.L
				ll = _lemit(ll, new_inst(IrOpOpStrCmp{}, cmp_nt.Temp, left, right.Temp, "", int64(0)))
				ll = _lset_type(ll, cmp_nt.Temp, int64(1))
				zero_nt := _lnew_temp(ll)
				_ = zero_nt
				ll = zero_nt.L
				ll = _lemit(ll, new_inst(IrOpOpConst{}, zero_nt.Temp, int64(0), int64(0), "", int64(0)))
				ll = _lset_type(ll, zero_nt.Temp, int64(1))
				ll = _lemit(ll, new_inst(_cmp_op(k), nt.Temp, cmp_nt.Temp, zero_nt.Temp, "", int64(0)))
				left = nt.Temp
			} else {
				ll = _lemit(ll, new_inst(_cmp_op(k), nt.Temp, left, right.Temp, "", int64(0)))
				left = nt.Temp
			}
		} else if (k == "&&") {
			ll = _ladv(ll)
			right := _lower_comparison(ll)
			_ = right
			ll = right.L
			nt := _lnew_temp(ll)
			_ = nt
			ll = nt.L
			ll = _lemit(ll, new_inst(IrOpOpAnd{}, nt.Temp, left, right.Temp, "", int64(0)))
			left = nt.Temp
		} else if (k == "||") {
			ll = _ladv(ll)
			right := _lower_comparison(ll)
			_ = right
			ll = right.L
			nt := _lnew_temp(ll)
			_ = nt
			ll = nt.L
			ll = _lemit(ll, new_inst(IrOpOpOr{}, nt.Temp, left, right.Temp, "", int64(0)))
			left = nt.Temp
		} else if ((k == "+") || (k == "-")) {
			ll = _ladv(ll)
			right := _lower_multiplicative(ll)
			_ = right
			ll = right.L
			left_type := _lget_type(ll, left)
			_ = left_type
			right_type := _lget_type(ll, right.Temp)
			_ = right_type
			if ((k == "+") && (left_type == int64(12))) {
				nt := _lnew_str_temp(ll)
				_ = nt
				ll = nt.L
				ll = _lemit(ll, new_inst(IrOpOpStrConcat{}, nt.Temp, left, right.Temp, "", int64(0)))
				ll = _lset_type(ll, nt.Temp, int64(12))
				left = nt.Temp
			} else if (_is_float_type(left_type) || _is_float_type(right_type)) {
				nt := _lnew_temp(ll)
				_ = nt
				ll = nt.L
				ll = _lemit(ll, new_inst(_arith_op(k), nt.Temp, left, right.Temp, "f64", int64(0)))
				ll = _lset_type(ll, nt.Temp, int64(9))
				left = nt.Temp
			} else {
				nt := _lnew_temp(ll)
				_ = nt
				ll = nt.L
				ll = _lemit(ll, new_inst(_arith_op(k), nt.Temp, left, right.Temp, "", int64(0)))
				left = nt.Temp
			}
		} else if (((k == "*") || (k == "/")) || (k == "%")) {
			ll = _ladv(ll)
			right := _lower_atom(ll)
			_ = right
			ll = right.L
			nt := _lnew_temp(ll)
			_ = nt
			ll = nt.L
			left_type := _lget_type(ll, left)
			_ = left_type
			right_type := _lget_type(ll, right.Temp)
			_ = right_type
			if (_is_float_type(left_type) || _is_float_type(right_type)) {
				ll = _lemit(ll, new_inst(_arith_op(k), nt.Temp, left, right.Temp, "f64", int64(0)))
				ll = _lset_type(ll, nt.Temp, int64(9))
			} else {
				ll = _lemit(ll, new_inst(_arith_op(k), nt.Temp, left, right.Temp, "", int64(0)))
			}
			left = nt.Temp
		} else {
			running = false
		}
	}
	for ((_lk(ll) != lb) && (_lk(ll) != "EOF")) {
		ll = _ladv(ll)
	}
	return LT{L: ll, Temp: left}
}

func _lower_if(l Lowerer) LT {
	ll := _ladv(l)
	_ = ll
	lb := "{"
	_ = lb
	rb := "}"
	_ = rb
	cond := _lower_if_cond(ll)
	_ = cond
	ll = cond.L
	else_lbl := _lnew_label(ll)
	_ = else_lbl
	ll = else_lbl.L
	end_lbl := _lnew_label(ll)
	_ = end_lbl
	ll = end_lbl.L
	result_nt := _lnew_str_temp(ll)
	_ = result_nt
	ll = result_nt.L
	ll = _lemit(ll, new_inst(IrOpOpBranchFalse{}, int64(0), cond.Temp, else_lbl.Temp, "", int64(0)))
	then_last := -int64(1)
	_ = then_last
	if (_lk(ll) == lb) {
		ll = _ladv(ll)
		ll = _lskip_nl(ll)
		for ((_lk(ll) != rb) && (_lk(ll) != "EOF")) {
			k := _lk(ll)
			_ = k
			if ((((((((((k == "return") || (k == "mut")) || (k == "break")) || (k == "continue")) || (k == "for")) || (k == "while")) || (k == "loop")) || (k == "assert")) || (k == "defer")) || (k == "with")) {
				ll = _lower_stmt(ll)
			} else if (k == "if") {
				r := _lower_if(ll)
				_ = r
				ll = r.L
				then_last = r.Temp
			} else {
				r := _lower_expr_result(ll)
				_ = r
				ll = r.L
				then_last = r.Temp
			}
			ll = _lskip_nl(ll)
		}
		if (_lk(ll) == rb) {
			ll = _ladv(ll)
		}
		if (then_last >= int64(0)) {
			ll = _lemit(ll, new_inst(IrOpOpStore{}, result_nt.Temp, then_last, int64(0), "", int64(0)))
			then_type := _lget_type(ll, then_last)
			_ = then_type
			if (then_type == int64(12)) {
				then_len_dst := (result_nt.Temp + int64(1))
				_ = then_len_dst
				then_len_src := (then_last + int64(1))
				_ = then_len_src
				ll = _lemit(ll, new_inst(IrOpOpStore{}, then_len_dst, then_len_src, int64(0), "", int64(0)))
			}
			ll = _lset_type(ll, result_nt.Temp, then_type)
		}
	}
	ll = _lemit(ll, new_inst(IrOpOpJump{}, int64(0), end_lbl.Temp, int64(0), "", int64(0)))
	ll = _lemit(ll, new_inst(IrOpOpLabel{}, int64(0), else_lbl.Temp, int64(0), "", int64(0)))
	if (_lk(ll) == "else") {
		ll = _ladv(ll)
		if (_lk(ll) == "if") {
			else_result := _lower_if(ll)
			_ = else_result
			ll = else_result.L
			ll = _lemit(ll, new_inst(IrOpOpStore{}, result_nt.Temp, else_result.Temp, int64(0), "", int64(0)))
			eif_type := _lget_type(ll, else_result.Temp)
			_ = eif_type
			if (eif_type == int64(12)) {
				eif_len_dst := (result_nt.Temp + int64(1))
				_ = eif_len_dst
				eif_len_src := (else_result.Temp + int64(1))
				_ = eif_len_src
				ll = _lemit(ll, new_inst(IrOpOpStore{}, eif_len_dst, eif_len_src, int64(0), "", int64(0)))
			}
			ll = _lset_type(ll, result_nt.Temp, eif_type)
		} else if (_lk(ll) == lb) {
			ll = _ladv(ll)
			ll = _lskip_nl(ll)
			else_last := -int64(1)
			_ = else_last
			for ((_lk(ll) != rb) && (_lk(ll) != "EOF")) {
				k := _lk(ll)
				_ = k
				if ((((((((((k == "return") || (k == "mut")) || (k == "break")) || (k == "continue")) || (k == "for")) || (k == "while")) || (k == "loop")) || (k == "assert")) || (k == "defer")) || (k == "with")) {
					ll = _lower_stmt(ll)
				} else if (k == "if") {
					r := _lower_if(ll)
					_ = r
					ll = r.L
					else_last = r.Temp
				} else {
					r := _lower_expr_result(ll)
					_ = r
					ll = r.L
					else_last = r.Temp
				}
				ll = _lskip_nl(ll)
			}
			if (_lk(ll) == rb) {
				ll = _ladv(ll)
			}
			if (else_last >= int64(0)) {
				ll = _lemit(ll, new_inst(IrOpOpStore{}, result_nt.Temp, else_last, int64(0), "", int64(0)))
				else_type := _lget_type(ll, else_last)
				_ = else_type
				if (else_type == int64(12)) {
					else_len_dst := (result_nt.Temp + int64(1))
					_ = else_len_dst
					else_len_src := (else_last + int64(1))
					_ = else_len_src
					ll = _lemit(ll, new_inst(IrOpOpStore{}, else_len_dst, else_len_src, int64(0), "", int64(0)))
				}
				ll = _lset_type(ll, result_nt.Temp, else_type)
			}
		}
	}
	ll = _lemit(ll, new_inst(IrOpOpLabel{}, int64(0), end_lbl.Temp, int64(0), "", int64(0)))
	return LT{L: ll, Temp: result_nt.Temp}
}

func _lower_while(l Lowerer) Lowerer {
	ll := _ladv(l)
	_ = ll
	lb := "{"
	_ = lb
	loop_lbl := _lnew_label(ll)
	_ = loop_lbl
	ll = loop_lbl.L
	end_lbl := _lnew_label(ll)
	_ = end_lbl
	ll = end_lbl.L
	ll = _lemit(ll, new_inst(IrOpOpLabel{}, int64(0), loop_lbl.Temp, int64(0), "", int64(0)))
	cond := _lower_if_cond(ll)
	_ = cond
	ll = cond.L
	ll = _lemit(ll, new_inst(IrOpOpBranchFalse{}, int64(0), cond.Temp, end_lbl.Temp, "", int64(0)))
	prev_start := ll.Loop_start
	_ = prev_start
	prev_end := ll.Loop_end
	_ = prev_end
	ll = _lset_loop(ll, loop_lbl.Temp, end_lbl.Temp)
	ll = _lower_block(ll)
	ll = _lset_loop(ll, prev_start, prev_end)
	ll = _lemit(ll, new_inst(IrOpOpJump{}, int64(0), loop_lbl.Temp, int64(0), "", int64(0)))
	ll = _lemit(ll, new_inst(IrOpOpLabel{}, int64(0), end_lbl.Temp, int64(0), "", int64(0)))
	return ll
}

func _lower_for_range(l Lowerer, var_name string, inclusive bool) Lowerer {
	ll := l
	_ = ll
	lb := "{"
	_ = lb
	start := _lower_expr(ll)
	_ = start
	ll = start.L
	if ((_lk(ll) == "..") || (_lk(ll) == "..=")) {
		ll = _ladv(ll)
	}
	end := _lower_expr(ll)
	_ = end
	ll = end.L
	for ((_lk(ll) != lb) && (_lk(ll) != "EOF")) {
		ll = _ladv(ll)
	}
	limit_temp := end.Temp
	_ = limit_temp
	if inclusive {
		one_nt := _lnew_temp(ll)
		_ = one_nt
		ll = one_nt.L
		ll = _lemit(ll, new_inst(IrOpOpConst{}, one_nt.Temp, int64(1), int64(0), "", int64(0)))
		inc_nt := _lnew_temp(ll)
		_ = inc_nt
		ll = inc_nt.L
		ll = _lemit(ll, new_inst(IrOpOpAdd{}, inc_nt.Temp, end.Temp, one_nt.Temp, "", int64(0)))
		limit_temp = inc_nt.Temp
	}
	idx_nt := _lnew_temp(ll)
	_ = idx_nt
	ll = idx_nt.L
	ll = _lemit(ll, new_inst(IrOpOpStore{}, idx_nt.Temp, start.Temp, int64(0), "", int64(0)))
	loop_lbl := _lnew_label(ll)
	_ = loop_lbl
	ll = loop_lbl.L
	end_lbl := _lnew_label(ll)
	_ = end_lbl
	ll = end_lbl.L
	inc_lbl := _lnew_label(ll)
	_ = inc_lbl
	ll = inc_lbl.L
	ll = _lemit(ll, new_inst(IrOpOpLabel{}, int64(0), loop_lbl.Temp, int64(0), "", int64(0)))
	cond_nt := _lnew_temp(ll)
	_ = cond_nt
	ll = cond_nt.L
	ll = _lemit(ll, new_inst(IrOpOpLt{}, cond_nt.Temp, idx_nt.Temp, limit_temp, "", int64(0)))
	ll = _lemit(ll, new_inst(IrOpOpBranchFalse{}, int64(0), cond_nt.Temp, end_lbl.Temp, "", int64(0)))
	if (var_name != "") {
		ll = _lenv_add(ll, var_name, idx_nt.Temp)
	}
	prev_start := ll.Loop_start
	_ = prev_start
	prev_end := ll.Loop_end
	_ = prev_end
	ll = Lowerer{Tokens: ll.Tokens, Pos: ll.Pos, Store: ll.Store, Registry: ll.Registry, Treg: ll.Treg, Pool: ll.Pool, Table: ll.Table, File: ll.File, Module: ll.Module, Current_func: ll.Current_func, Temp_counter: ll.Temp_counter, Label_counter: ll.Label_counter, Env_names: ll.Env_names, Env_slots: ll.Env_slots, Loop_start: inc_lbl.Temp, Loop_end: end_lbl.Temp, Const_names: ll.Const_names, Const_vals: ll.Const_vals, Const_str_names: ll.Const_str_names, Const_str_vals: ll.Const_str_vals, Temp_types: ll.Temp_types, Fn_error_type: ll.Fn_error_type, Defer_starts: ll.Defer_starts, Defer_ends: ll.Defer_ends}
	ll = _lower_block(ll)
	ll = _lset_loop(ll, prev_start, prev_end)
	ll = _lemit(ll, new_inst(IrOpOpLabel{}, int64(0), inc_lbl.Temp, int64(0), "", int64(0)))
	one_nt := _lnew_temp(ll)
	_ = one_nt
	ll = one_nt.L
	ll = _lemit(ll, new_inst(IrOpOpConst{}, one_nt.Temp, int64(1), int64(0), "", int64(0)))
	inc_nt := _lnew_temp(ll)
	_ = inc_nt
	ll = inc_nt.L
	ll = _lemit(ll, new_inst(IrOpOpAdd{}, inc_nt.Temp, idx_nt.Temp, one_nt.Temp, "", int64(0)))
	ll = _lemit(ll, new_inst(IrOpOpStore{}, idx_nt.Temp, inc_nt.Temp, int64(0), "", int64(0)))
	ll = _lemit(ll, new_inst(IrOpOpJump{}, int64(0), loop_lbl.Temp, int64(0), "", int64(0)))
	ll = _lemit(ll, new_inst(IrOpOpLabel{}, int64(0), end_lbl.Temp, int64(0), "", int64(0)))
	return ll
}

func _lower_loop(l Lowerer) Lowerer {
	ll := _ladv(l)
	_ = ll
	loop_lbl := _lnew_label(ll)
	_ = loop_lbl
	ll = loop_lbl.L
	end_lbl := _lnew_label(ll)
	_ = end_lbl
	ll = end_lbl.L
	ll = _lemit(ll, new_inst(IrOpOpLabel{}, int64(0), loop_lbl.Temp, int64(0), "", int64(0)))
	prev_start := ll.Loop_start
	_ = prev_start
	prev_end := ll.Loop_end
	_ = prev_end
	ll = _lset_loop(ll, loop_lbl.Temp, end_lbl.Temp)
	ll = _lower_block(ll)
	ll = _lset_loop(ll, prev_start, prev_end)
	ll = _lemit(ll, new_inst(IrOpOpJump{}, int64(0), loop_lbl.Temp, int64(0), "", int64(0)))
	ll = _lemit(ll, new_inst(IrOpOpLabel{}, int64(0), end_lbl.Temp, int64(0), "", int64(0)))
	return ll
}

func _lower_for(l Lowerer) Lowerer {
	ll := _ladv(l)
	_ = ll
	lb := "{"
	_ = lb
	var_name := ""
	_ = var_name
	if (_lk(ll) == "IDENT") {
		var_name = _lcur(ll).Text
		ll = _ladv(ll)
	}
	if (_lk(ll) == "in") {
		ll = _ladv(ll)
	}
	is_range := false
	_ = is_range
	is_inclusive := false
	_ = is_inclusive
	peek := ll.Pos
	_ = peek
	for (peek < int64(len(ll.Tokens))) {
		pk := token_name(ll.Tokens[peek].Kind)
		_ = pk
		if (pk == "..") {
			is_range = true
		}
		if (pk == "..=") {
			is_range = true
		}
		if (pk == "..=") {
			is_inclusive = true
		}
		if (((pk == lb) || (pk == "EOF")) || (pk == "NEWLINE")) {
			peek = int64(len(ll.Tokens))
		}
		peek = (peek + int64(1))
	}
	if is_range {
		return _lower_for_range(ll, var_name, is_inclusive)
	}
	iter := _lower_expr(ll)
	_ = iter
	ll = iter.L
	for ((_lk(ll) != lb) && (_lk(ll) != "EOF")) {
		ll = _ladv(ll)
	}
	iter_type_check := _lget_type(ll, iter.Temp)
	_ = iter_type_check
	is_chan := false
	_ = is_chan
	if ((iter_type_check > int64(0)) && (iter_type_check < int64(len(ll.Store.Types)))) {
		iter_kind := type_kind_name(ll.Store.Types[iter_type_check].Kind)
		_ = iter_kind
		if ((iter_kind == "Named") && (((ll.Store.Types[iter_type_check].Name == "Chan") || (ll.Store.Types[iter_type_check].Name == "Channel")))) {
			is_chan = true
		}
	}
	if is_chan {
		return _lower_for_channel(ll, var_name, iter.Temp)
	}
	len_nt := _lnew_temp(ll)
	_ = len_nt
	ll = len_nt.L
	ll = _lemit(ll, new_inst(IrOpOpArrayLen{}, len_nt.Temp, iter.Temp, int64(0), "", int64(0)))
	idx_nt := _lnew_temp(ll)
	_ = idx_nt
	ll = idx_nt.L
	ll = _lemit(ll, new_inst(IrOpOpConst{}, idx_nt.Temp, int64(0), int64(0), "", int64(0)))
	iter_type := _lget_type(ll, iter.Temp)
	_ = iter_type
	loop_lbl := _lnew_label(ll)
	_ = loop_lbl
	ll = loop_lbl.L
	end_lbl := _lnew_label(ll)
	_ = end_lbl
	ll = end_lbl.L
	inc_lbl := _lnew_label(ll)
	_ = inc_lbl
	ll = inc_lbl.L
	ll = _lemit(ll, new_inst(IrOpOpLabel{}, int64(0), loop_lbl.Temp, int64(0), "", int64(0)))
	cond_nt := _lnew_temp(ll)
	_ = cond_nt
	ll = cond_nt.L
	ll = _lemit(ll, new_inst(IrOpOpLt{}, cond_nt.Temp, idx_nt.Temp, len_nt.Temp, "", int64(0)))
	ll = _lemit(ll, new_inst(IrOpOpBranchFalse{}, int64(0), cond_nt.Temp, end_lbl.Temp, "", int64(0)))
	if (var_name != "") {
		if ((iter_type == int64(13)) || _is_str_array_type(ll.Store, iter_type)) {
			elem_nt := _lnew_str_temp(ll)
			_ = elem_nt
			ll = elem_nt.L
			ll = _lemit(ll, new_inst(IrOpOpArrayGet{}, elem_nt.Temp, iter.Temp, idx_nt.Temp, "", int64(12)))
			ll = _lset_type(ll, elem_nt.Temp, int64(12))
			ll = _lenv_add(ll, var_name, elem_nt.Temp)
		} else {
			elem_nt := _lnew_temp(ll)
			_ = elem_nt
			ll = elem_nt.L
			ll = _lemit(ll, new_inst(IrOpOpArrayGet{}, elem_nt.Temp, iter.Temp, idx_nt.Temp, "", int64(0)))
			ll = _lenv_add(ll, var_name, elem_nt.Temp)
		}
	}
	prev_start := ll.Loop_start
	_ = prev_start
	prev_end := ll.Loop_end
	_ = prev_end
	ll = Lowerer{Tokens: ll.Tokens, Pos: ll.Pos, Store: ll.Store, Registry: ll.Registry, Treg: ll.Treg, Pool: ll.Pool, Table: ll.Table, File: ll.File, Module: ll.Module, Current_func: ll.Current_func, Temp_counter: ll.Temp_counter, Label_counter: ll.Label_counter, Env_names: ll.Env_names, Env_slots: ll.Env_slots, Loop_start: inc_lbl.Temp, Loop_end: end_lbl.Temp, Const_names: ll.Const_names, Const_vals: ll.Const_vals, Const_str_names: ll.Const_str_names, Const_str_vals: ll.Const_str_vals, Temp_types: ll.Temp_types, Fn_error_type: ll.Fn_error_type, Defer_starts: ll.Defer_starts, Defer_ends: ll.Defer_ends}
	ll = _lower_block(ll)
	ll = _lset_loop(ll, prev_start, prev_end)
	ll = _lemit(ll, new_inst(IrOpOpLabel{}, int64(0), inc_lbl.Temp, int64(0), "", int64(0)))
	one_nt := _lnew_temp(ll)
	_ = one_nt
	ll = one_nt.L
	ll = _lemit(ll, new_inst(IrOpOpConst{}, one_nt.Temp, int64(1), int64(0), "", int64(0)))
	inc_nt := _lnew_temp(ll)
	_ = inc_nt
	ll = inc_nt.L
	ll = _lemit(ll, new_inst(IrOpOpAdd{}, inc_nt.Temp, idx_nt.Temp, one_nt.Temp, "", int64(0)))
	ll = _lemit(ll, new_inst(IrOpOpStore{}, idx_nt.Temp, inc_nt.Temp, int64(0), "", int64(0)))
	ll = _lemit(ll, new_inst(IrOpOpJump{}, int64(0), loop_lbl.Temp, int64(0), "", int64(0)))
	ll = _lemit(ll, new_inst(IrOpOpLabel{}, int64(0), end_lbl.Temp, int64(0), "", int64(0)))
	return ll
}

func _lower_for_channel(l Lowerer, var_name string, ch_temp int64) Lowerer {
	ll := l
	_ = ll
	lb := "{"
	_ = lb
	rb := "}"
	_ = rb
	loop_lbl := _lnew_label(ll)
	_ = loop_lbl
	ll = loop_lbl.L
	end_lbl := _lnew_label(ll)
	_ = end_lbl
	ll = end_lbl.L
	ll = _lemit(ll, new_inst(IrOpOpLabel{}, int64(0), loop_lbl.Temp, int64(0), "", int64(0)))
	recv_nt := _lnew_str_temp(ll)
	_ = recv_nt
	ll = recv_nt.L
	ll = _lemit(ll, new_inst(IrOpOpCall{}, recv_nt.Temp, ch_temp, int64(1), "_aria_chan_try_recv", int64(0)))
	status_nt := _lnew_temp(ll)
	_ = status_nt
	ll = status_nt.L
	ll = _lemit(ll, new_inst(IrOpOpLoad{}, status_nt.Temp, (recv_nt.Temp + int64(1)), int64(0), "", int64(0)))
	zero_nt := _lnew_temp(ll)
	_ = zero_nt
	ll = zero_nt.L
	ll = _lemit(ll, new_inst(IrOpOpConst{}, zero_nt.Temp, int64(0), int64(0), "", int64(0)))
	cmp_nt := _lnew_temp(ll)
	_ = cmp_nt
	ll = cmp_nt.L
	ll = _lemit(ll, new_inst(IrOpOpEq{}, cmp_nt.Temp, status_nt.Temp, zero_nt.Temp, "", int64(0)))
	ll = _lemit(ll, new_inst(IrOpOpBranchTrue{}, int64(0), cmp_nt.Temp, end_lbl.Temp, "", int64(0)))
	if (var_name != "") {
		val_nt := _lnew_temp(ll)
		_ = val_nt
		ll = val_nt.L
		ll = _lemit(ll, new_inst(IrOpOpLoad{}, val_nt.Temp, recv_nt.Temp, int64(0), "", int64(0)))
		ll = _lenv_add(ll, var_name, val_nt.Temp)
	}
	prev_start := ll.Loop_start
	_ = prev_start
	prev_end := ll.Loop_end
	_ = prev_end
	ll = _lset_loop(ll, loop_lbl.Temp, end_lbl.Temp)
	ll = _lower_block(ll)
	ll = _lset_loop(ll, prev_start, prev_end)
	ll = _lemit(ll, new_inst(IrOpOpJump{}, int64(0), loop_lbl.Temp, int64(0), "", int64(0)))
	ll = _lemit(ll, new_inst(IrOpOpLabel{}, int64(0), end_lbl.Temp, int64(0), "", int64(0)))
	return ll
}

func _lower_interpolated_string(l Lowerer) LT {
	ll := l
	_ = ll
	start_text := _lcur(ll).Text
	_ = start_text
	start_result := _lower_string_const(ll, start_text)
	_ = start_result
	ll = start_result.L
	ll = _ladv(ll)
	result := start_result.Temp
	_ = result
	running := true
	_ = running
	for running {
		if (((_lk(ll) == "STRING_END") || (_lk(ll) == "STRING_MIDDLE")) || (_lk(ll) == "EOF")) {
			running = false
		} else {
			expr := _lower_expr(ll)
			_ = expr
			ll = expr.L
			expr_type := _lget_type(ll, expr.Temp)
			_ = expr_type
			str_temp := expr.Temp
			_ = str_temp
			if (expr_type != int64(12)) {
				conv_nt := _lnew_str_temp(ll)
				_ = conv_nt
				ll = conv_nt.L
				if _is_float_type(expr_type) {
					ll = _lemit(ll, new_inst(IrOpOpIntToStr{}, conv_nt.Temp, expr.Temp, int64(0), "", int64(0)))
				} else {
					ll = _lemit(ll, new_inst(IrOpOpIntToStr{}, conv_nt.Temp, expr.Temp, int64(0), "", int64(0)))
				}
				ll = _lset_type(ll, conv_nt.Temp, int64(12))
				str_temp = conv_nt.Temp
			}
			cat_nt := _lnew_str_temp(ll)
			_ = cat_nt
			ll = cat_nt.L
			ll = _lemit(ll, new_inst(IrOpOpStrConcat{}, cat_nt.Temp, result, str_temp, "", int64(0)))
			ll = _lset_type(ll, cat_nt.Temp, int64(12))
			result = cat_nt.Temp
		}
	}
	for (_lk(ll) == "STRING_MIDDLE") {
		mid_text := _lcur(ll).Text
		_ = mid_text
		ll = _ladv(ll)
		if (int64(len(mid_text)) > int64(0)) {
			mid_result := _lower_string_const(ll, mid_text)
			_ = mid_result
			ll = mid_result.L
			cat_nt := _lnew_str_temp(ll)
			_ = cat_nt
			ll = cat_nt.L
			ll = _lemit(ll, new_inst(IrOpOpStrConcat{}, cat_nt.Temp, result, mid_result.Temp, "", int64(0)))
			ll = _lset_type(ll, cat_nt.Temp, int64(12))
			result = cat_nt.Temp
		}
		expr_running := true
		_ = expr_running
		for expr_running {
			if (((_lk(ll) == "STRING_END") || (_lk(ll) == "STRING_MIDDLE")) || (_lk(ll) == "EOF")) {
				expr_running = false
			} else {
				expr := _lower_expr(ll)
				_ = expr
				ll = expr.L
				expr_type := _lget_type(ll, expr.Temp)
				_ = expr_type
				str_temp := expr.Temp
				_ = str_temp
				if (expr_type != int64(12)) {
					conv_nt := _lnew_str_temp(ll)
					_ = conv_nt
					ll = conv_nt.L
					if _is_float_type(expr_type) {
						ll = _lemit(ll, new_inst(IrOpOpIntToStr{}, conv_nt.Temp, expr.Temp, int64(0), "", int64(0)))
					} else {
						ll = _lemit(ll, new_inst(IrOpOpIntToStr{}, conv_nt.Temp, expr.Temp, int64(0), "", int64(0)))
					}
					ll = _lset_type(ll, conv_nt.Temp, int64(12))
					str_temp = conv_nt.Temp
				}
				cat_nt := _lnew_str_temp(ll)
				_ = cat_nt
				ll = cat_nt.L
				ll = _lemit(ll, new_inst(IrOpOpStrConcat{}, cat_nt.Temp, result, str_temp, "", int64(0)))
				ll = _lset_type(ll, cat_nt.Temp, int64(12))
				result = cat_nt.Temp
			}
		}
	}
	if (_lk(ll) == "STRING_END") {
		end_text := _lcur(ll).Text
		_ = end_text
		ll = _ladv(ll)
		if (int64(len(end_text)) > int64(0)) {
			end_result := _lower_string_const(ll, end_text)
			_ = end_result
			ll = end_result.L
			cat_nt := _lnew_str_temp(ll)
			_ = cat_nt
			ll = cat_nt.L
			ll = _lemit(ll, new_inst(IrOpOpStrConcat{}, cat_nt.Temp, result, end_result.Temp, "", int64(0)))
			ll = _lset_type(ll, cat_nt.Temp, int64(12))
			result = cat_nt.Temp
		}
	}
	ll = _lset_type(ll, result, int64(12))
	return LT{L: ll, Temp: result}
}

func _lappend_defer(l Lowerer, start int64, end int64) Lowerer {
	return Lowerer{Tokens: l.Tokens, Pos: l.Pos, Store: l.Store, Registry: l.Registry, Treg: l.Treg, Pool: l.Pool, Table: l.Table, File: l.File, Module: l.Module, Current_func: l.Current_func, Temp_counter: l.Temp_counter, Label_counter: l.Label_counter, Env_names: l.Env_names, Env_slots: l.Env_slots, Loop_start: l.Loop_start, Loop_end: l.Loop_end, Const_names: l.Const_names, Const_vals: l.Const_vals, Const_str_names: l.Const_str_names, Const_str_vals: l.Const_str_vals, Temp_types: l.Temp_types, Fn_error_type: l.Fn_error_type, Defer_starts: append(l.Defer_starts, start), Defer_ends: append(l.Defer_ends, end)}
}

func _emit_defers(l Lowerer) Lowerer {
	ll := l
	_ = ll
	ll = _emit_drops(ll)
	defer_count := (int64(len(ll.Defer_starts)) - int64(1))
	_ = defer_count
	if (defer_count == int64(0)) {
		return ll
	}
	di := defer_count
	_ = di
	for (di >= int64(1)) {
		ds := ll.Defer_starts[di]
		_ = ds
		de := ll.Defer_ends[di]
		_ = de
		if ((ds > int64(0)) && (de > ds)) {
			saved_pos := ll.Pos
			_ = saved_pos
			ll = _lset_pos(ll, ds)
			result := _lower_expr(ll)
			_ = result
			ll = result.L
			ll = _lset_pos(ll, saved_pos)
		}
		di = (di - int64(1))
	}
	return ll
}

func _lower_make_trait_object(l Lowerer, data_temp int64, concrete_type string, trait_name string) LT {
	ll := l
	_ = ll
	tdef := treg_find_trait(ll.Treg, trait_name)
	_ = tdef
	method_count := (int64(len(tdef.Method_names)) - int64(1))
	_ = method_count
	if (method_count == int64(0)) {
		return LT{L: ll, Temp: data_temp}
	}
	vt_nt := _lnew_temp(ll)
	_ = vt_nt
	ll = vt_nt.L
	ll = _lemit(ll, new_inst(IrOpOpAlloc{}, vt_nt.Temp, method_count, int64(0), "_vtable", int64(0)))
	vi := int64(1)
	_ = vi
	for (vi < int64(len(tdef.Method_names))) {
		mname := tdef.Method_names[vi]
		_ = mname
		fn_name := ((concrete_type + "_") + mname)
		_ = fn_name
		sig := reg_find_fn(ll.Registry, fn_name)
		_ = sig
		fn_nt := _lnew_temp(ll)
		_ = fn_nt
		ll = fn_nt.L
		if (sig.Name != "") {
			ll = _lemit(ll, new_inst(IrOpOpFnRef{}, fn_nt.Temp, int64(0), int64(0), fn_name, int64(0)))
		} else {
			ll = _lemit(ll, new_inst(IrOpOpConst{}, fn_nt.Temp, int64(0), int64(0), "", int64(0)))
		}
		slot := (vi - int64(1))
		_ = slot
		ll = _lemit(ll, new_inst(IrOpOpFieldSet{}, vt_nt.Temp, fn_nt.Temp, slot, "", int64(0)))
		vi = (vi + int64(1))
	}
	to_nt := _lnew_temp(ll)
	_ = to_nt
	ll = to_nt.L
	ll = _lemit(ll, new_inst(IrOpOpTraitObject{}, to_nt.Temp, data_temp, vt_nt.Temp, trait_name, int64(0)))
	ll = _lset_type(ll, to_nt.Temp, int64(500))
	return LT{L: ll, Temp: to_nt.Temp}
}

func _emit_drops_scoped(l Lowerer, scope_start int64) Lowerer {
	ll := l
	_ = ll
	i := (int64(len(ll.Env_names)) - int64(1))
	_ = i
	for ((i >= scope_start) && (i > int64(0))) {
		slot := ll.Env_slots[i]
		_ = slot
		obj_type := _lget_type(ll, slot)
		_ = obj_type
		sname := _struct_name_for_type(ll.Store, ll.Registry, obj_type)
		_ = sname
		if (sname != "") {
			drop_fn := (sname + "_drop")
			_ = drop_fn
			drop_sig := reg_find_fn(ll.Registry, drop_fn)
			_ = drop_sig
			if (drop_sig.Name != "") {
				nt := _lnew_temp(ll)
				_ = nt
				ll = nt.L
				ll = _lemit(ll, new_inst(IrOpOpCall{}, nt.Temp, slot, int64(1), drop_fn, int64(0)))
			}
		}
		i = (i - int64(1))
	}
	return ll
}

func _emit_drops(l Lowerer) Lowerer {
	return _emit_drops_scoped(l, int64(1))
}

func _lower_closure(l Lowerer) LT {
	ll := _ladv(l)
	_ = ll
	if ((_lk(ll) == "IDENT") && (_lcur(ll).Text == "once")) {
		ll = _ladv(ll)
	}
	closure_name := ("_closure_" + i2s(int64(len(ll.Module.Funcs))))
	_ = closure_name
	param_names := []string{""}
	_ = param_names
	param_types := []int64{int64(0)}
	_ = param_types
	if (_lk(ll) == "(") {
		ll = _ladv(ll)
		for ((_lk(ll) != ")") && (_lk(ll) != "EOF")) {
			if ((_lk(ll) == "IDENT") || (_lk(ll) == "self")) {
				ptext := _lcur(ll).Text
				_ = ptext
				ll = _ladv(ll)
				ptid := int64(0)
				_ = ptid
				if (_lk(ll) == ":") {
					ll = _ladv(ll)
					if (_lk(ll) == "IDENT") {
						resolved := resolve_type_name(ll.Store, _lcur(ll).Text)
						_ = resolved
						if (resolved != TY_UNKNOWN) {
							ptid = resolved
						}
					}
					for (((_lk(ll) != ",") && (_lk(ll) != ")")) && (_lk(ll) != "EOF")) {
						ll = _ladv(ll)
					}
				}
				param_names = append(param_names, ptext)
				param_types = append(param_types, ptid)
			} else {
				ll = _ladv(ll)
			}
			if (_lk(ll) == ",") {
				ll = _ladv(ll)
			}
		}
		if (_lk(ll) == ")") {
			ll = _ladv(ll)
		}
	}
	param_count := (int64(len(param_names)) - int64(1))
	_ = param_count
	ret_type := int64(1)
	_ = ret_type
	lb2 := "{"
	_ = lb2
	if (_lk(ll) == "->") {
		ll = _ladv(ll)
		if (_lk(ll) == "IDENT") {
			rtid := resolve_type_name(ll.Store, _lcur(ll).Text)
			_ = rtid
			if (rtid != TY_UNKNOWN) {
				ret_type = rtid
			}
		}
		for (((_lk(ll) != "=>") && (_lk(ll) != lb2)) && (_lk(ll) != "EOF")) {
			ll = _ladv(ll)
		}
	}
	if (_lk(ll) == "=>") {
		ll = _ladv(ll)
	}
	ll = _lskip_nl(ll)
	body_start := ll.Pos
	_ = body_start
	lb := "{"
	_ = lb
	if (_lk(ll) == lb) {
		ll = _skip_braces_l(ll)
	} else {
		depth := int64(0)
		_ = depth
		found_end := false
		_ = found_end
		for (((_lk(ll) != "NEWLINE") && (_lk(ll) != "EOF")) && (found_end == false)) {
			if (_lk(ll) == "(") {
				depth = (depth + int64(1))
			}
			if (_lk(ll) == ")") {
				if (depth == int64(0)) {
					found_end = true
				} else {
					depth = (depth - int64(1))
				}
			}
			if (found_end == false) {
				if ((_lk(ll) == ",") && (depth == int64(0))) {
					found_end = true
				}
			}
			if (found_end == false) {
				ll = _ladv(ll)
			}
		}
	}
	body_end := ll.Pos
	_ = body_end
	saved_func := ll.Current_func
	_ = saved_func
	saved_temp := ll.Temp_counter
	_ = saved_temp
	saved_label := ll.Label_counter
	_ = saved_label
	saved_env_names := ll.Env_names
	_ = saved_env_names
	saved_env_slots := ll.Env_slots
	_ = saved_env_slots
	saved_types := ll.Temp_types
	_ = saved_types
	cap_names := []string{""}
	_ = cap_names
	cap_slots := []int64{int64(0)}
	_ = cap_slots
	scan_pos := body_start
	_ = scan_pos
	for ((scan_pos < body_end) && (scan_pos < int64(len(ll.Tokens)))) {
		stok := ll.Tokens[scan_pos]
		_ = stok
		if (token_name(stok.Kind) == "IDENT") {
			sname := stok.Text
			_ = sname
			outer_slot := _lenv_lookup(ll, sname)
			_ = outer_slot
			if (outer_slot >= int64(0)) {
				is_param := false
				_ = is_param
				pci := int64(1)
				_ = pci
				for (pci < int64(len(param_names))) {
					if (param_names[pci] == sname) {
						is_param = true
					}
					pci = (pci + int64(1))
				}
				if (is_param == false) {
					already := false
					_ = already
					cci := int64(1)
					_ = cci
					for (cci < int64(len(cap_names))) {
						if (cap_names[cci] == sname) {
							already = true
						}
						cci = (cci + int64(1))
					}
					if (already == false) {
						cap_names = append(cap_names, sname)
						cap_slots = append(cap_slots, outer_slot)
					}
				}
			}
		}
		scan_pos = (scan_pos + int64(1))
	}
	cap_count := (int64(len(cap_names)) - int64(1))
	_ = cap_count
	actual_params := int64(1)
	_ = actual_params
	reg_pi := int64(1)
	_ = reg_pi
	for (reg_pi < int64(len(param_types))) {
		if (param_types[reg_pi] == int64(12)) {
			actual_params = (actual_params + int64(2))
		} else {
			actual_params = (actual_params + int64(1))
		}
		reg_pi = (reg_pi + int64(1))
	}
	ll = _lset_func(ll, new_ir_func(closure_name, actual_params, ret_type))
	env_param := _lnew_temp(ll)
	_ = env_param
	ll = env_param.L
	ll = _lenv_add(ll, "_env", env_param.Temp)
	pi := int64(1)
	_ = pi
	for (pi < int64(len(param_names))) {
		pt := param_types[pi]
		_ = pt
		pn := param_names[pi]
		_ = pn
		if (pt == int64(12)) {
			nt := _lnew_str_temp(ll)
			_ = nt
			ll = nt.L
			ll = _lenv_add(ll, pn, nt.Temp)
			ll = _lset_type(ll, nt.Temp, int64(12))
		} else {
			nt := _lnew_temp(ll)
			_ = nt
			ll = nt.L
			ll = _lenv_add(ll, pn, nt.Temp)
			if (pt > int64(0)) {
				ll = _lset_type(ll, nt.Temp, pt)
			}
		}
		pi = (pi + int64(1))
	}
	if (cap_count > int64(0)) {
		ci := int64(1)
		_ = ci
		for (ci < int64(len(cap_names))) {
			cap_nt := _lnew_temp(ll)
			_ = cap_nt
			ll = cap_nt.L
			ll = _lemit(ll, new_inst(IrOpOpFieldGet{}, cap_nt.Temp, env_param.Temp, (ci - int64(1)), "", int64(0)))
			ll = _lenv_add(ll, cap_names[ci], cap_nt.Temp)
			ci = (ci + int64(1))
		}
	}
	ll = _lset_pos(ll, body_start)
	lb3 := "{"
	_ = lb3
	if (body_start < body_end) {
		if (_lk(ll) == lb3) {
			ll = _lower_fn_body(ll, int64(1))
		} else {
			result := _lower_expr(ll)
			_ = result
			ll = result.L
			ll = _lemit(ll, new_inst(IrOpOpRet{}, int64(0), result.Temp, int64(0), "", int64(0)))
		}
	}
	ll = _lfinish_func(ll)
	ll = Lowerer{Tokens: ll.Tokens, Pos: body_end, Store: ll.Store, Registry: ll.Registry, Treg: ll.Treg, Pool: ll.Pool, Table: ll.Table, File: ll.File, Module: ll.Module, Current_func: saved_func, Temp_counter: saved_temp, Label_counter: saved_label, Env_names: saved_env_names, Env_slots: saved_env_slots, Loop_start: ll.Loop_start, Loop_end: ll.Loop_end, Const_names: ll.Const_names, Const_vals: ll.Const_vals, Const_str_names: ll.Const_str_names, Const_str_vals: ll.Const_str_vals, Temp_types: saved_types, Fn_error_type: int64(0), Defer_starts: []int64{int64(0)}, Defer_ends: []int64{int64(0)}}
	fat_nt := _lnew_temp(ll)
	_ = fat_nt
	ll = fat_nt.L
	ll = _lemit(ll, new_inst(IrOpOpAlloc{}, fat_nt.Temp, int64(2), int64(0), "_closure", int64(0)))
	fnref_nt := _lnew_temp(ll)
	_ = fnref_nt
	ll = fnref_nt.L
	ll = _lemit(ll, new_inst(IrOpOpFnRef{}, fnref_nt.Temp, int64(0), int64(0), closure_name, int64(0)))
	ll = _lemit(ll, new_inst(IrOpOpFieldSet{}, fat_nt.Temp, fnref_nt.Temp, int64(0), "_fn", int64(0)))
	if (cap_count > int64(0)) {
		env_nt := _lnew_temp(ll)
		_ = env_nt
		ll = env_nt.L
		ll = _lemit(ll, new_inst(IrOpOpAlloc{}, env_nt.Temp, cap_count, int64(0), "_env", int64(0)))
		ci := int64(1)
		_ = ci
		for (ci < int64(len(cap_slots))) {
			val_nt := _lnew_temp(ll)
			_ = val_nt
			ll = val_nt.L
			ll = _lemit(ll, new_inst(IrOpOpLoad{}, val_nt.Temp, cap_slots[ci], int64(0), "", int64(0)))
			ll = _lemit(ll, new_inst(IrOpOpFieldSet{}, env_nt.Temp, val_nt.Temp, (ci - int64(1)), "", int64(0)))
			ci = (ci + int64(1))
		}
		ll = _lemit(ll, new_inst(IrOpOpFieldSet{}, fat_nt.Temp, env_nt.Temp, int64(1), "_env", int64(0)))
	} else {
		null_nt := _lnew_temp(ll)
		_ = null_nt
		ll = null_nt.L
		ll = _lemit(ll, new_inst(IrOpOpConst{}, null_nt.Temp, int64(0), int64(0), "", int64(0)))
		ll = _lemit(ll, new_inst(IrOpOpFieldSet{}, fat_nt.Temp, null_nt.Temp, int64(1), "_env", int64(0)))
	}
	return LT{L: ll, Temp: fat_nt.Temp}
}

func _is_newtype(def StructDef) bool {
	if (int64(len(def.Fields)) != int64(2)) {
		return false
	}
	return (def.Fields[int64(1)].Name == "value")
}

func _variant_tag(reg TypeRegistry, name string) int64 {
	si := int64(1)
	_ = si
	for (si < int64(len(reg.Struct_defs))) {
		def := reg.Struct_defs[si]
		_ = def
		if def.Is_sum {
			vi := int64(1)
			_ = vi
			for (vi < int64(len(def.Variant_names))) {
				if (def.Variant_names[vi] == name) {
					return vi
				}
				vi = (vi + int64(1))
			}
		}
		si = (si + int64(1))
	}
	return int64(0)
}

func _variant_parent(reg TypeRegistry, name string) StructDef {
	si := int64(1)
	_ = si
	for (si < int64(len(reg.Struct_defs))) {
		def := reg.Struct_defs[si]
		_ = def
		if def.Is_sum {
			vi := int64(1)
			_ = vi
			for (vi < int64(len(def.Variant_names))) {
				if (def.Variant_names[vi] == name) {
					return def
				}
				vi = (vi + int64(1))
			}
		}
		si = (si + int64(1))
	}
	return _sentinel_struct_def()
}

func _parse_int_literal(val string) int64 {
	num := int64(0)
	_ = num
	i := int64(0)
	_ = i
	neg := false
	_ = neg
	if ((int64(len(val)) > int64(0)) && (string(val[int64(0)]) == "-")) {
		neg = true
		i = int64(1)
	}
	if ((int64(len(val)) > (i + int64(1))) && (string(val[i]) == "0")) {
		p := string(val[(i + int64(1))])
		_ = p
		if ((p == "x") || (p == "X")) {
			num = _parse_hex(val, (i + int64(2)))
			if neg {
				num = (int64(0) - num)
			}
			return num
		}
		if ((p == "o") || (p == "O")) {
			num = _parse_oct(val, (i + int64(2)))
			if neg {
				num = (int64(0) - num)
			}
			return num
		}
		if ((p == "b") || (p == "B")) {
			num = _parse_bin(val, (i + int64(2)))
			if neg {
				num = (int64(0) - num)
			}
			return num
		}
	}
	for (i < int64(len(val))) {
		ch := string(val[i])
		_ = ch
		if (ch != "_") {
			digit := int64(0)
			_ = digit
			if (ch == "0") {
				digit = int64(0)
			} else if (ch == "1") {
				digit = int64(1)
			} else if (ch == "2") {
				digit = int64(2)
			} else if (ch == "3") {
				digit = int64(3)
			} else if (ch == "4") {
				digit = int64(4)
			} else if (ch == "5") {
				digit = int64(5)
			} else if (ch == "6") {
				digit = int64(6)
			} else if (ch == "7") {
				digit = int64(7)
			} else if (ch == "8") {
				digit = int64(8)
			} else if (ch == "9") {
				digit = int64(9)
			}
			num = ((num * int64(10)) + digit)
		}
		i = (i + int64(1))
	}
	if neg {
		num = (int64(0) - num)
	}
	return num
}

func _parse_duration(val string) int64 {
	num_end := int64(0)
	_ = num_end
	for (num_end < int64(len(val))) {
		c := string(val[num_end])
		_ = c
		if ((((c >= "0") && (c <= "9"))) || (c == "_")) {
			num_end = (num_end + int64(1))
		} else {
			break
		}
	}
	num_str := val[int64(0):num_end]
	_ = num_str
	suffix := val[num_end:int64(len(val))]
	_ = suffix
	num := _parse_int_literal(num_str)
	_ = num
	if (suffix == "ns") {
		return num
	}
	if (suffix == "us") {
		return (num * int64(1000))
	}
	if (suffix == "ms") {
		return (num * int64(1000000))
	}
	if (suffix == "s") {
		return (num * int64(1000000000))
	}
	if (suffix == "m") {
		return (num * int64(60000000000))
	}
	if (suffix == "h") {
		return (num * int64(3600000000000))
	}
	return num
}

func _parse_size(val string) int64 {
	num_end := int64(0)
	_ = num_end
	for (num_end < int64(len(val))) {
		c := string(val[num_end])
		_ = c
		if ((((c >= "0") && (c <= "9"))) || (c == "_")) {
			num_end = (num_end + int64(1))
		} else {
			break
		}
	}
	num_str := val[int64(0):num_end]
	_ = num_str
	suffix := val[num_end:int64(len(val))]
	_ = suffix
	num := _parse_int_literal(num_str)
	_ = num
	if (suffix == "b") {
		return num
	}
	if (suffix == "kb") {
		return (num * int64(1024))
	}
	if (suffix == "mb") {
		return (num * int64(1048576))
	}
	if (suffix == "gb") {
		return (num * int64(1073741824))
	}
	return num
}

func _char_to_code(ch string) int64 {
	if (ch == "\n") {
		return int64(10)
	}
	if (ch == "\t") {
		return int64(9)
	}
	if (ch == "\r") {
		return int64(13)
	}
	if (ch == "\x00") {
		return int64(0)
	}
	if (ch == " ") {
		return int64(32)
	}
	if (ch == "!") {
		return int64(33)
	}
	if (ch == "\"") {
		return int64(34)
	}
	if (ch == "#") {
		return int64(35)
	}
	if (ch == "$") {
		return int64(36)
	}
	if (ch == "%") {
		return int64(37)
	}
	if (ch == "&") {
		return int64(38)
	}
	if (ch == "'") {
		return int64(39)
	}
	if (ch == "(") {
		return int64(40)
	}
	if (ch == ")") {
		return int64(41)
	}
	if (ch == "*") {
		return int64(42)
	}
	if (ch == "+") {
		return int64(43)
	}
	if (ch == ",") {
		return int64(44)
	}
	if (ch == "-") {
		return int64(45)
	}
	if (ch == ".") {
		return int64(46)
	}
	if (ch == "/") {
		return int64(47)
	}
	if (ch == "0") {
		return int64(48)
	}
	if (ch == "1") {
		return int64(49)
	}
	if (ch == "2") {
		return int64(50)
	}
	if (ch == "3") {
		return int64(51)
	}
	if (ch == "4") {
		return int64(52)
	}
	if (ch == "5") {
		return int64(53)
	}
	if (ch == "6") {
		return int64(54)
	}
	if (ch == "7") {
		return int64(55)
	}
	if (ch == "8") {
		return int64(56)
	}
	if (ch == "9") {
		return int64(57)
	}
	if (ch == ":") {
		return int64(58)
	}
	if (ch == ";") {
		return int64(59)
	}
	if (ch == "<") {
		return int64(60)
	}
	if (ch == "=") {
		return int64(61)
	}
	if (ch == ">") {
		return int64(62)
	}
	if (ch == "?") {
		return int64(63)
	}
	if (ch == "@") {
		return int64(64)
	}
	if (ch == "A") {
		return int64(65)
	}
	if (ch == "B") {
		return int64(66)
	}
	if (ch == "C") {
		return int64(67)
	}
	if (ch == "D") {
		return int64(68)
	}
	if (ch == "E") {
		return int64(69)
	}
	if (ch == "F") {
		return int64(70)
	}
	if (ch == "G") {
		return int64(71)
	}
	if (ch == "H") {
		return int64(72)
	}
	if (ch == "I") {
		return int64(73)
	}
	if (ch == "J") {
		return int64(74)
	}
	if (ch == "K") {
		return int64(75)
	}
	if (ch == "L") {
		return int64(76)
	}
	if (ch == "M") {
		return int64(77)
	}
	if (ch == "N") {
		return int64(78)
	}
	if (ch == "O") {
		return int64(79)
	}
	if (ch == "P") {
		return int64(80)
	}
	if (ch == "Q") {
		return int64(81)
	}
	if (ch == "R") {
		return int64(82)
	}
	if (ch == "S") {
		return int64(83)
	}
	if (ch == "T") {
		return int64(84)
	}
	if (ch == "U") {
		return int64(85)
	}
	if (ch == "V") {
		return int64(86)
	}
	if (ch == "W") {
		return int64(87)
	}
	if (ch == "X") {
		return int64(88)
	}
	if (ch == "Y") {
		return int64(89)
	}
	if (ch == "Z") {
		return int64(90)
	}
	if (ch == "[") {
		return int64(91)
	}
	if (ch == "\\") {
		return int64(92)
	}
	if (ch == "]") {
		return int64(93)
	}
	if (ch == "^") {
		return int64(94)
	}
	if (ch == "_") {
		return int64(95)
	}
	if (ch == "`") {
		return int64(96)
	}
	if (ch == "a") {
		return int64(97)
	}
	if (ch == "b") {
		return int64(98)
	}
	if (ch == "c") {
		return int64(99)
	}
	if (ch == "d") {
		return int64(100)
	}
	if (ch == "e") {
		return int64(101)
	}
	if (ch == "f") {
		return int64(102)
	}
	if (ch == "g") {
		return int64(103)
	}
	if (ch == "h") {
		return int64(104)
	}
	if (ch == "i") {
		return int64(105)
	}
	if (ch == "j") {
		return int64(106)
	}
	if (ch == "k") {
		return int64(107)
	}
	if (ch == "l") {
		return int64(108)
	}
	if (ch == "m") {
		return int64(109)
	}
	if (ch == "n") {
		return int64(110)
	}
	if (ch == "o") {
		return int64(111)
	}
	if (ch == "p") {
		return int64(112)
	}
	if (ch == "q") {
		return int64(113)
	}
	if (ch == "r") {
		return int64(114)
	}
	if (ch == "s") {
		return int64(115)
	}
	if (ch == "t") {
		return int64(116)
	}
	if (ch == "u") {
		return int64(117)
	}
	if (ch == "v") {
		return int64(118)
	}
	if (ch == "w") {
		return int64(119)
	}
	if (ch == "x") {
		return int64(120)
	}
	if (ch == "y") {
		return int64(121)
	}
	if (ch == "z") {
		return int64(122)
	}
	if (ch == "{") {
		return int64(123)
	}
	if (ch == "|") {
		return int64(124)
	}
	if (ch == "}") {
		return int64(125)
	}
	if (ch == "~") {
		return int64(126)
	}
	return int64(0)
}

func _parse_hex(val string, start int64) int64 {
	num := int64(0)
	_ = num
	i := start
	_ = i
	for (i < int64(len(val))) {
		ch := string(val[i])
		_ = ch
		if (ch != "_") {
			digit := int64(0)
			_ = digit
			if (ch == "0") {
				digit = int64(0)
			} else if (ch == "1") {
				digit = int64(1)
			} else if (ch == "2") {
				digit = int64(2)
			} else if (ch == "3") {
				digit = int64(3)
			} else if (ch == "4") {
				digit = int64(4)
			} else if (ch == "5") {
				digit = int64(5)
			} else if (ch == "6") {
				digit = int64(6)
			} else if (ch == "7") {
				digit = int64(7)
			} else if (ch == "8") {
				digit = int64(8)
			} else if (ch == "9") {
				digit = int64(9)
			} else if ((ch == "a") || (ch == "A")) {
				digit = int64(10)
			} else if ((ch == "b") || (ch == "B")) {
				digit = int64(11)
			} else if ((ch == "c") || (ch == "C")) {
				digit = int64(12)
			} else if ((ch == "d") || (ch == "D")) {
				digit = int64(13)
			} else if ((ch == "e") || (ch == "E")) {
				digit = int64(14)
			} else if ((ch == "f") || (ch == "F")) {
				digit = int64(15)
			}
			num = ((num * int64(16)) + digit)
		}
		i = (i + int64(1))
	}
	return num
}

func _parse_oct(val string, start int64) int64 {
	num := int64(0)
	_ = num
	i := start
	_ = i
	for (i < int64(len(val))) {
		ch := string(val[i])
		_ = ch
		if (ch != "_") {
			digit := int64(0)
			_ = digit
			if (ch == "0") {
				digit = int64(0)
			} else if (ch == "1") {
				digit = int64(1)
			} else if (ch == "2") {
				digit = int64(2)
			} else if (ch == "3") {
				digit = int64(3)
			} else if (ch == "4") {
				digit = int64(4)
			} else if (ch == "5") {
				digit = int64(5)
			} else if (ch == "6") {
				digit = int64(6)
			} else if (ch == "7") {
				digit = int64(7)
			}
			num = ((num * int64(8)) + digit)
		}
		i = (i + int64(1))
	}
	return num
}

func _parse_bin(val string, start int64) int64 {
	num := int64(0)
	_ = num
	i := start
	_ = i
	for (i < int64(len(val))) {
		ch := string(val[i])
		_ = ch
		if (ch != "_") {
			if (ch == "1") {
				num = ((num * int64(2)) + int64(1))
			} else {
				num = (num * int64(2))
			}
		}
		i = (i + int64(1))
	}
	return num
}

func _lower_string_const(l Lowerer, val string) LT {
	ll := l
	_ = ll
	str_idx := mod_find_string(ll.Module, val)
	_ = str_idx
	if (str_idx == int64(0)) {
		new_mod := mod_add_string(ll.Module, val)
		_ = new_mod
		str_idx = (int64(len(new_mod.String_constants)) - int64(1))
		ll = Lowerer{Tokens: ll.Tokens, Pos: ll.Pos, Store: ll.Store, Registry: ll.Registry, Treg: ll.Treg, Pool: ll.Pool, Table: ll.Table, File: ll.File, Module: new_mod, Current_func: ll.Current_func, Temp_counter: ll.Temp_counter, Label_counter: ll.Label_counter, Env_names: ll.Env_names, Env_slots: ll.Env_slots, Loop_start: ll.Loop_start, Loop_end: ll.Loop_end, Const_names: ll.Const_names, Const_vals: ll.Const_vals, Const_str_names: ll.Const_str_names, Const_str_vals: ll.Const_str_vals, Temp_types: ll.Temp_types, Fn_error_type: ll.Fn_error_type, Defer_starts: ll.Defer_starts, Defer_ends: ll.Defer_ends}
	}
	nt := _lnew_str_temp(ll)
	_ = nt
	ll = nt.L
	ll = _lemit(ll, new_inst(IrOpOpConstStr{}, nt.Temp, str_idx, int64(len(val)), "", int64(0)))
	ll = _lset_type(ll, nt.Temp, int64(12))
	return LT{L: ll, Temp: nt.Temp}
}

func _lower_match_subject(l Lowerer) LT {
	k := _lk(l)
	_ = k
	if (k == "IDENT") {
		name := _lcur(l).Text
		_ = name
		ll := _ladv(l)
		_ = ll
		if (_lk(ll) == "(") {
			sig := reg_find_fn(ll.Registry, name)
			_ = sig
			if (sig.Name != "") {
				ll = _ladv(ll)
				ll = _lskip_nl(ll)
				first_arg := -int64(1)
				_ = first_arg
				arg_count := int64(0)
				_ = arg_count
				for ((_lk(ll) != ")") && (_lk(ll) != "EOF")) {
					ll = _lskip_nl(ll)
					arg := _lower_expr(ll)
					_ = arg
					ll = arg.L
					if (arg_count == int64(0)) {
						first_arg = arg.Temp
					}
					arg_count = (arg_count + int64(1))
					ll = _lskip_nl(ll)
					if (_lk(ll) == ",") {
						ll = _ladv(ll)
					}
				}
				if (_lk(ll) == ")") {
					ll = _ladv(ll)
				}
				nt := _lnew_temp(ll)
				_ = nt
				ll = nt.L
				fa := int64(0)
				_ = fa
				if (first_arg >= int64(0)) {
					fa = first_arg
				}
				ll = _lemit(ll, new_inst(IrOpOpCall{}, nt.Temp, fa, arg_count, name, int64(0)))
				return LT{L: ll, Temp: nt.Temp}
			}
		}
		if (_lk(ll) == ".") {
			slot := _lenv_lookup(ll, name)
			_ = slot
			if (slot >= int64(0)) {
				src_type := _lget_type(ll, slot)
				_ = src_type
				base_temp := int64(0)
				_ = base_temp
				if (src_type == int64(12)) {
					nt_base := _lnew_str_temp(ll)
					_ = nt_base
					ll = nt_base.L
					base_temp = nt_base.Temp
				} else {
					nt_base := _lnew_temp(ll)
					_ = nt_base
					ll = nt_base.L
					base_temp = nt_base.Temp
				}
				ll = _lemit(ll, new_inst(IrOpOpLoad{}, base_temp, slot, int64(0), "", int64(0)))
				if (src_type != int64(0)) {
					ll = _lset_type(ll, base_temp, src_type)
				}
				cur := base_temp
				_ = cur
				for (_lk(ll) == ".") {
					ll = _ladv(ll)
					if (_lk(ll) == "IDENT") {
						field := _lcur(ll).Text
						_ = field
						ll = _ladv(ll)
						nt := _lnew_temp(ll)
						_ = nt
						ll = nt.L
						ll = _lemit(ll, new_inst(IrOpOpFieldGet{}, nt.Temp, cur, int64(0), field, int64(0)))
						cur = nt.Temp
					}
				}
				return LT{L: ll, Temp: cur}
			}
		}
		slot := _lenv_lookup(ll, name)
		_ = slot
		if (slot >= int64(0)) {
			src_type := _lget_type(ll, slot)
			_ = src_type
			nt_temp := int64(0)
			_ = nt_temp
			if (src_type == int64(12)) {
				nt := _lnew_str_temp(ll)
				_ = nt
				ll = nt.L
				nt_temp = nt.Temp
			} else {
				nt := _lnew_temp(ll)
				_ = nt
				ll = nt.L
				nt_temp = nt.Temp
			}
			ll = _lemit(ll, new_inst(IrOpOpLoad{}, nt_temp, slot, int64(0), "", int64(0)))
			if (src_type != int64(0)) {
				ll = _lset_type(ll, nt_temp, src_type)
			}
			return LT{L: ll, Temp: nt_temp}
		}
		tag := _variant_tag(ll.Registry, name)
		_ = tag
		if (tag > int64(0)) {
			nt := _lnew_temp(ll)
			_ = nt
			ll = nt.L
			ll = _lemit(ll, new_inst(IrOpOpConst{}, nt.Temp, tag, int64(0), "", int64(0)))
			return LT{L: ll, Temp: nt.Temp}
		}
		nt := _lnew_temp(ll)
		_ = nt
		ll = nt.L
		ll = _lemit(ll, new_inst(IrOpOpConst{}, nt.Temp, int64(0), int64(0), "", int64(0)))
		return LT{L: ll, Temp: nt.Temp}
	}
	if (k == "INT") {
		return _lower_primary(l)
	}
	if (k == "(") {
		return _lower_primary(l)
	}
	nt := _lnew_temp(l)
	_ = nt
	ll := _lemit(_ladv(nt.L), new_inst(IrOpOpConst{}, nt.Temp, int64(0), int64(0), "", int64(0)))
	_ = ll
	return LT{L: ll, Temp: nt.Temp}
}

func _lower_spawn(l Lowerer) LT {
	ll := _ladv(l)
	_ = ll
	lb := "{"
	_ = lb
	is_detach := false
	_ = is_detach
	if (_lk(ll) == ".") {
		next_pos := (ll.Pos + int64(1))
		_ = next_pos
		if ((next_pos < int64(len(ll.Tokens))) && (ll.Tokens[next_pos].Text == "detach")) {
			ll = _ladv(ll)
			ll = _ladv(ll)
			is_detach = true
		}
	}
	if (_lk(ll) == lb) {
		closure := _lower_closure_from_block(ll)
		_ = closure
		ll = closure.L
		fn_ptr := _lnew_temp(ll)
		_ = fn_ptr
		ll = fn_ptr.L
		ll = _lemit(ll, new_inst(IrOpOpFieldGet{}, fn_ptr.Temp, closure.Temp, int64(0), "_fn", int64(0)))
		env_ptr := _lnew_temp(ll)
		_ = env_ptr
		ll = env_ptr.L
		ll = _lemit(ll, new_inst(IrOpOpFieldGet{}, env_ptr.Temp, closure.Temp, int64(1), "_env", int64(0)))
		nt := _lnew_temp(ll)
		_ = nt
		ll = nt.L
		ll = _lemit(ll, new_inst(IrOpOpCall{}, nt.Temp, fn_ptr.Temp, int64(2), "_aria_spawn", int64(1)))
		spawn_type := int64(1)
		_ = spawn_type
		if is_detach {
			spawn_type = int64(18)
		}
		ll = _lset_type(ll, nt.Temp, spawn_type)
		return LT{L: ll, Temp: nt.Temp}
	}
	expr := _lower_expr(ll)
	_ = expr
	ll = expr.L
	fn_ptr := _lnew_temp(ll)
	_ = fn_ptr
	ll = fn_ptr.L
	ll = _lemit(ll, new_inst(IrOpOpFieldGet{}, fn_ptr.Temp, expr.Temp, int64(0), "_fn", int64(0)))
	env_ptr := _lnew_temp(ll)
	_ = env_ptr
	ll = env_ptr.L
	ll = _lemit(ll, new_inst(IrOpOpFieldGet{}, env_ptr.Temp, expr.Temp, int64(1), "_env", int64(0)))
	nt := _lnew_temp(ll)
	_ = nt
	ll = nt.L
	ll = _lemit(ll, new_inst(IrOpOpCall{}, nt.Temp, fn_ptr.Temp, int64(2), "_aria_spawn", int64(1)))
	spawn_type := int64(1)
	_ = spawn_type
	if is_detach {
		spawn_type = int64(18)
	}
	ll = _lset_type(ll, nt.Temp, spawn_type)
	return LT{L: ll, Temp: nt.Temp}
}

func _lower_closure_from_block(l Lowerer) LT {
	ll := l
	_ = ll
	lb := "{"
	_ = lb
	rb := "}"
	_ = rb
	fname := ("_spawn_fn_" + i2s(ll.Temp_counter))
	_ = fname
	saved_func := ll.Current_func
	_ = saved_func
	saved_tc := ll.Temp_counter
	_ = saved_tc
	saved_lc := ll.Label_counter
	_ = saved_lc
	saved_env := ll.Env_names
	_ = saved_env
	saved_slots := ll.Env_slots
	_ = saved_slots
	ll = _lset_func(ll, new_ir_func(fname, int64(1), int64(0)))
	env_nt := _lnew_temp(ll)
	_ = env_nt
	ll = env_nt.L
	ll = _lemit(ll, new_inst(IrOpOpArg{}, env_nt.Temp, int64(0), int64(0), "", int64(0)))
	ll = _ladv(ll)
	ll = _lskip_nl(ll)
	for ((_lk(ll) != rb) && (_lk(ll) != "EOF")) {
		ll = _lower_stmt(ll)
		ll = _lskip_nl(ll)
	}
	if (_lk(ll) == rb) {
		ll = _ladv(ll)
	}
	ret_nt := _lnew_temp(ll)
	_ = ret_nt
	ll = ret_nt.L
	ll = _lemit(ll, new_inst(IrOpOpConst{}, ret_nt.Temp, int64(0), int64(0), "", int64(0)))
	ll = _lemit(ll, new_inst(IrOpOpRet{}, int64(0), ret_nt.Temp, int64(0), "", int64(0)))
	ll = _lfinish_func(ll)
	ll = _lset_func(ll, saved_func)
	fn_ref := _lnew_temp(ll)
	_ = fn_ref
	ll = fn_ref.L
	ll = _lemit(ll, new_inst(IrOpOpFnRef{}, fn_ref.Temp, int64(0), int64(0), fname, int64(0)))
	clos := _lnew_temp(ll)
	_ = clos
	ll = clos.L
	ll = _lemit(ll, new_inst(IrOpOpAlloc{}, clos.Temp, int64(2), int64(0), "_Closure", int64(0)))
	ll = _lemit(ll, new_inst(IrOpOpFieldSet{}, clos.Temp, fn_ref.Temp, int64(0), "_fn", int64(0)))
	null_nt := _lnew_temp(ll)
	_ = null_nt
	ll = null_nt.L
	ll = _lemit(ll, new_inst(IrOpOpConst{}, null_nt.Temp, int64(0), int64(0), "", int64(0)))
	ll = _lemit(ll, new_inst(IrOpOpFieldSet{}, clos.Temp, null_nt.Temp, int64(1), "_env", int64(0)))
	return LT{L: ll, Temp: clos.Temp}
}

func _lower_scope(l Lowerer) LT {
	ll := _ladv(l)
	_ = ll
	lb := "{"
	_ = lb
	rb := "}"
	_ = rb
	if (_lk(ll) != lb) {
		nt := _lnew_temp(ll)
		_ = nt
		ll = nt.L
		ll = _lemit(ll, new_inst(IrOpOpConst{}, nt.Temp, int64(0), int64(0), "", int64(0)))
		return LT{L: ll, Temp: nt.Temp}
	}
	ll = _ladv(ll)
	ll = _lskip_nl(ll)
	task_count := int64(0)
	_ = task_count
	task_0 := int64(0)
	_ = task_0
	task_1 := int64(0)
	_ = task_1
	task_2 := int64(0)
	_ = task_2
	task_3 := int64(0)
	_ = task_3
	task_4 := int64(0)
	_ = task_4
	task_5 := int64(0)
	_ = task_5
	task_6 := int64(0)
	_ = task_6
	task_7 := int64(0)
	_ = task_7
	last_temp := int64(0)
	_ = last_temp
	for ((_lk(ll) != rb) && (_lk(ll) != "EOF")) {
		if (_lk(ll) == "spawn") {
			result := _lower_spawn(ll)
			_ = result
			ll = result.L
			if (task_count == int64(0)) {
				task_0 = result.Temp
			}
			if (task_count == int64(1)) {
				task_1 = result.Temp
			}
			if (task_count == int64(2)) {
				task_2 = result.Temp
			}
			if (task_count == int64(3)) {
				task_3 = result.Temp
			}
			if (task_count == int64(4)) {
				task_4 = result.Temp
			}
			if (task_count == int64(5)) {
				task_5 = result.Temp
			}
			if (task_count == int64(6)) {
				task_6 = result.Temp
			}
			if (task_count == int64(7)) {
				task_7 = result.Temp
			}
			task_count = (task_count + int64(1))
			last_temp = result.Temp
		} else {
			ll = _lower_stmt(ll)
		}
		ll = _lskip_nl(ll)
	}
	if (_lk(ll) == rb) {
		ll = _ladv(ll)
	}
	ti := int64(0)
	_ = ti
	for (ti < task_count) {
		th := int64(0)
		_ = th
		if (ti == int64(0)) {
			th = task_0
		}
		if (ti == int64(1)) {
			th = task_1
		}
		if (ti == int64(2)) {
			th = task_2
		}
		if (ti == int64(3)) {
			th = task_3
		}
		if (ti == int64(4)) {
			th = task_4
		}
		if (ti == int64(5)) {
			th = task_5
		}
		if (ti == int64(6)) {
			th = task_6
		}
		if (ti == int64(7)) {
			th = task_7
		}
		th_type := _lget_type(ll, th)
		_ = th_type
		if (th_type != int64(18)) {
			nt := _lnew_temp(ll)
			_ = nt
			ll = nt.L
			ll = _lemit(ll, new_inst(IrOpOpCall{}, nt.Temp, th, int64(1), "_aria_task_await", int64(1)))
		}
		ti = (ti + int64(1))
	}
	result_nt := _lnew_temp(ll)
	_ = result_nt
	ll = result_nt.L
	ll = _lemit(ll, new_inst(IrOpOpConst{}, result_nt.Temp, int64(0), int64(0), "", int64(0)))
	return LT{L: ll, Temp: result_nt.Temp}
}

func _lower_select(l Lowerer) LT {
	ll := _ladv(l)
	_ = ll
	lb := "{"
	_ = lb
	rb := "}"
	_ = rb
	if (_lk(ll) != lb) {
		nt := _lnew_temp(ll)
		_ = nt
		ll = nt.L
		ll = _lemit(ll, new_inst(IrOpOpConst{}, nt.Temp, int64(0), int64(0), "", int64(0)))
		return LT{L: ll, Temp: nt.Temp}
	}
	ll = _ladv(ll)
	ll = _lskip_nl(ll)
	ch_count := int64(0)
	_ = ch_count
	has_default := false
	_ = has_default
	arr_nt := _lnew_temp(ll)
	_ = arr_nt
	ll = arr_nt.L
	ll = _lemit(ll, new_inst(IrOpOpArrayNew{}, arr_nt.Temp, int64(8), int64(0), "", int64(0)))
	bind_name_0 := ""
	_ = bind_name_0
	bind_name_1 := ""
	_ = bind_name_1
	bind_name_2 := ""
	_ = bind_name_2
	bind_name_3 := ""
	_ = bind_name_3
	arm_start_0 := int64(0)
	_ = arm_start_0
	arm_start_1 := int64(0)
	_ = arm_start_1
	arm_start_2 := int64(0)
	_ = arm_start_2
	arm_start_3 := int64(0)
	_ = arm_start_3
	arm_end_0 := int64(0)
	_ = arm_end_0
	arm_end_1 := int64(0)
	_ = arm_end_1
	arm_end_2 := int64(0)
	_ = arm_end_2
	arm_end_3 := int64(0)
	_ = arm_end_3
	default_start := int64(0)
	_ = default_start
	default_end := int64(0)
	_ = default_end
	for ((_lk(ll) != rb) && (_lk(ll) != "EOF")) {
		ll = _lskip_nl(ll)
		if (_lk(ll) == rb) {
			break
		}
		if ((_lk(ll) == "_") || (((_lk(ll) == "IDENT") && (_lcur(ll).Text == "default")))) {
			has_default = true
			ll = _ladv(ll)
			if (_lk(ll) == "=>") {
				ll = _ladv(ll)
			}
			ll = _lskip_nl(ll)
			default_start = ll.Pos
			if (_lk(ll) == lb) {
				for ((_lk(ll) != rb) && (_lk(ll) != "EOF")) {
					ll = _ladv(ll)
				}
				if (_lk(ll) == rb) {
					ll = _ladv(ll)
				}
			} else {
				for (((_lk(ll) != "NEWLINE") && (_lk(ll) != rb)) && (_lk(ll) != "EOF")) {
					ll = _ladv(ll)
				}
			}
			default_end = ll.Pos
		} else if (_lk(ll) == "IDENT") {
			bname := _lcur(ll).Text
			_ = bname
			ll = _ladv(ll)
			if ((_lk(ll) == "IDENT") && (_lcur(ll).Text == "from")) {
				ll = _ladv(ll)
			}
			ch_expr := _lower_expr(ll)
			_ = ch_expr
			ll = ch_expr.L
			ll = _lemit(ll, new_inst(IrOpOpArrayAppend{}, arr_nt.Temp, arr_nt.Temp, ch_expr.Temp, "", int64(0)))
			if (ch_count == int64(0)) {
				bind_name_0 = bname
			}
			if (ch_count == int64(1)) {
				bind_name_1 = bname
			}
			if (ch_count == int64(2)) {
				bind_name_2 = bname
			}
			if (ch_count == int64(3)) {
				bind_name_3 = bname
			}
			if (_lk(ll) == "=>") {
				ll = _ladv(ll)
			}
			ll = _lskip_nl(ll)
			if (ch_count == int64(0)) {
				arm_start_0 = ll.Pos
			}
			if (ch_count == int64(1)) {
				arm_start_1 = ll.Pos
			}
			if (ch_count == int64(2)) {
				arm_start_2 = ll.Pos
			}
			if (ch_count == int64(3)) {
				arm_start_3 = ll.Pos
			}
			if (_lk(ll) == lb) {
				for ((_lk(ll) != rb) && (_lk(ll) != "EOF")) {
					ll = _ladv(ll)
				}
				if (_lk(ll) == rb) {
					ll = _ladv(ll)
				}
			} else {
				for (((_lk(ll) != "NEWLINE") && (_lk(ll) != rb)) && (_lk(ll) != "EOF")) {
					ll = _ladv(ll)
				}
			}
			if (ch_count == int64(0)) {
				arm_end_0 = ll.Pos
			}
			if (ch_count == int64(1)) {
				arm_end_1 = ll.Pos
			}
			if (ch_count == int64(2)) {
				arm_end_2 = ll.Pos
			}
			if (ch_count == int64(3)) {
				arm_end_3 = ll.Pos
			}
			ch_count = (ch_count + int64(1))
		} else {
			ll = _ladv(ll)
		}
		ll = _lskip_nl(ll)
	}
	if (_lk(ll) == rb) {
		ll = _ladv(ll)
	}
	timeout_nt := _lnew_temp(ll)
	_ = timeout_nt
	ll = timeout_nt.L
	timeout_val := (int64(0) - int64(1))
	_ = timeout_val
	if has_default {
		timeout_val = int64(0)
	}
	ll = _lemit(ll, new_inst(IrOpOpConst{}, timeout_nt.Temp, timeout_val, int64(0), "", int64(0)))
	pack0 := _lnew_temp(ll)
	_ = pack0
	ll = pack0.L
	ll = _lemit(ll, new_inst(IrOpOpStore{}, pack0.Temp, arr_nt.Temp, int64(0), "", int64(0)))
	pack1 := _lnew_temp(ll)
	_ = pack1
	ll = pack1.L
	ll = _lemit(ll, new_inst(IrOpOpStore{}, pack1.Temp, timeout_nt.Temp, int64(0), "", int64(0)))
	sel_result := _lnew_str_temp(ll)
	_ = sel_result
	ll = sel_result.L
	ll = _lemit(ll, new_inst(IrOpOpCall{}, sel_result.Temp, pack0.Temp, int64(2), "_aria_chan_select", int64(0)))
	result_nt := _lnew_temp(ll)
	_ = result_nt
	ll = result_nt.L
	end_lbl := _lnew_label(ll)
	_ = end_lbl
	ll = end_lbl.L
	idx_nt := _lnew_temp(ll)
	_ = idx_nt
	ll = idx_nt.L
	ll = _lemit(ll, new_inst(IrOpOpLoad{}, idx_nt.Temp, sel_result.Temp, int64(0), "", int64(0)))
	val_nt := _lnew_temp(ll)
	_ = val_nt
	ll = val_nt.L
	ll = _lemit(ll, new_inst(IrOpOpLoad{}, val_nt.Temp, (sel_result.Temp + int64(1)), int64(0), "", int64(0)))
	ai := int64(0)
	_ = ai
	for (ai < ch_count) {
		next_lbl := _lnew_label(ll)
		_ = next_lbl
		ll = next_lbl.L
		ai_nt := _lnew_temp(ll)
		_ = ai_nt
		ll = ai_nt.L
		ll = _lemit(ll, new_inst(IrOpOpConst{}, ai_nt.Temp, ai, int64(0), "", int64(0)))
		cmp := _lnew_temp(ll)
		_ = cmp
		ll = cmp.L
		ll = _lemit(ll, new_inst(IrOpOpEq{}, cmp.Temp, idx_nt.Temp, ai_nt.Temp, "", int64(0)))
		ll = _lemit(ll, new_inst(IrOpOpBranchFalse{}, int64(0), cmp.Temp, next_lbl.Temp, "", int64(0)))
		bname := ""
		_ = bname
		if (ai == int64(0)) {
			bname = bind_name_0
		}
		if (ai == int64(1)) {
			bname = bind_name_1
		}
		if (ai == int64(2)) {
			bname = bind_name_2
		}
		if (ai == int64(3)) {
			bname = bind_name_3
		}
		if ((bname != "") && (bname != "_")) {
			ll = _lenv_add(ll, bname, val_nt.Temp)
		}
		arm_s := int64(0)
		_ = arm_s
		arm_e := int64(0)
		_ = arm_e
		if (ai == int64(0)) {
			arm_s = arm_start_0
		}
		if (ai == int64(0)) {
			arm_e = arm_end_0
		}
		if (ai == int64(1)) {
			arm_s = arm_start_1
		}
		if (ai == int64(1)) {
			arm_e = arm_end_1
		}
		if (ai == int64(2)) {
			arm_s = arm_start_2
		}
		if (ai == int64(2)) {
			arm_e = arm_end_2
		}
		if (ai == int64(3)) {
			arm_s = arm_start_3
		}
		if (ai == int64(3)) {
			arm_e = arm_end_3
		}
		if (arm_s > int64(0)) {
			saved_pos := ll.Pos
			_ = saved_pos
			ll = _lset_pos(ll, arm_s)
			if (_lk(ll) == "{") {
				ll = _lower_block(ll)
			} else {
				arm_result := _lower_expr(ll)
				_ = arm_result
				ll = arm_result.L
				ll = _lemit(ll, new_inst(IrOpOpStore{}, result_nt.Temp, arm_result.Temp, int64(0), "", int64(0)))
			}
			ll = _lset_pos(ll, saved_pos)
		}
		ll = _lemit(ll, new_inst(IrOpOpJump{}, int64(0), end_lbl.Temp, int64(0), "", int64(0)))
		ll = _lemit(ll, new_inst(IrOpOpLabel{}, int64(0), next_lbl.Temp, int64(0), "", int64(0)))
		ai = (ai + int64(1))
	}
	if (has_default && (default_start > int64(0))) {
		saved_pos := ll.Pos
		_ = saved_pos
		ll = _lset_pos(ll, default_start)
		if (_lk(ll) == "{") {
			ll = _lower_block(ll)
		} else {
			arm_result := _lower_expr(ll)
			_ = arm_result
			ll = arm_result.L
			ll = _lemit(ll, new_inst(IrOpOpStore{}, result_nt.Temp, arm_result.Temp, int64(0), "", int64(0)))
		}
		ll = _lset_pos(ll, saved_pos)
	}
	ll = _lemit(ll, new_inst(IrOpOpLabel{}, int64(0), end_lbl.Temp, int64(0), "", int64(0)))
	return LT{L: ll, Temp: result_nt.Temp}
}

func _lower_match(l Lowerer) LT {
	ll := _ladv(l)
	_ = ll
	lb := "{"
	_ = lb
	rb := "}"
	_ = rb
	subj := _lower_match_subject(ll)
	_ = subj
	ll = subj.L
	for ((_lk(ll) != lb) && (_lk(ll) != "EOF")) {
		ll = _ladv(ll)
	}
	ll = _lskip_nl(ll)
	if (_lk(ll) != lb) {
		nt := _lnew_temp(ll)
		_ = nt
		return LT{L: nt.L, Temp: nt.Temp}
	}
	ll = _ladv(ll)
	ll = _lskip_nl(ll)
	result_nt := _lnew_str_temp(ll)
	_ = result_nt
	ll = result_nt.L
	end_lbl := _lnew_label(ll)
	_ = end_lbl
	ll = end_lbl.L
	is_data_sum := false
	_ = is_data_sum
	subj_tag_temp := subj.Temp
	_ = subj_tag_temp
	peek_pos := ll.Pos
	_ = peek_pos
	for ((peek_pos < int64(len(ll.Tokens))) && (token_name(ll.Tokens[peek_pos].Kind) == "NEWLINE")) {
		peek_pos = (peek_pos + int64(1))
	}
	if ((peek_pos < int64(len(ll.Tokens))) && (token_name(ll.Tokens[peek_pos].Kind) == "IDENT")) {
		first_pat := ll.Tokens[peek_pos].Text
		_ = first_pat
		at_pos := (peek_pos + int64(1))
		_ = at_pos
		if ((at_pos < int64(len(ll.Tokens))) && (ll.Tokens[at_pos].Text == "@")) {
			pat_pos := (at_pos + int64(1))
			_ = pat_pos
			if ((pat_pos < int64(len(ll.Tokens))) && (token_name(ll.Tokens[pat_pos].Kind) == "IDENT")) {
				first_pat = ll.Tokens[pat_pos].Text
			}
		}
		if (first_pat != "_") {
			pat_parent := _variant_parent(ll.Registry, first_pat)
			_ = pat_parent
			if ((pat_parent.Name != "") && sum_has_data(pat_parent)) {
				is_data_sum = true
				tag_ext := _lnew_temp(ll)
				_ = tag_ext
				ll = tag_ext.L
				ll = _lemit(ll, new_inst(IrOpOpFieldGet{}, tag_ext.Temp, subj.Temp, int64(0), "_tag", int64(0)))
				subj_tag_temp = tag_ext.Temp
			}
		}
	}
	for ((_lk(ll) != rb) && (_lk(ll) != "EOF")) {
		ll = _lskip_nl(ll)
		if ((_lk(ll) == rb) || (_lk(ll) == "EOF")) {
			return LT{L: ll, Temp: result_nt.Temp}
		}
		pattern := _lcur(ll).Text
		_ = pattern
		pattern_kind := _lk(ll)
		_ = pattern_kind
		is_wildcard := (pattern == "_")
		_ = is_wildcard
		is_variant := false
		_ = is_variant
		is_literal := false
		_ = is_literal
		is_binding := false
		_ = is_binding
		is_tuple_pattern := false
		_ = is_tuple_pattern
		is_array_pattern := false
		_ = is_array_pattern
		at_bind_name := ""
		_ = at_bind_name
		if (((((pattern_kind == "INT") || (pattern_kind == "STRING")) || (pattern_kind == "FLOAT")) || (pattern_kind == "true")) || (pattern_kind == "false")) {
			is_literal = true
		}
		if (pattern_kind == "(") {
			is_tuple_pattern = true
		}
		if (pattern_kind == "[") {
			is_array_pattern = true
		}
		if ((pattern_kind == "IDENT") && (is_wildcard == false)) {
			vtag_check := _variant_tag(ll.Registry, pattern)
			_ = vtag_check
			if (vtag_check > int64(0)) {
				is_variant = true
			} else {
				at_pos := (ll.Pos + int64(1))
				_ = at_pos
				if ((at_pos < int64(len(ll.Tokens))) && (token_name(ll.Tokens[at_pos].Kind) == "@")) {
					at_bind_name = pattern
					ll = _ladv(ll)
					ll = _ladv(ll)
					pattern = _lcur(ll).Text
					pattern_kind = _lk(ll)
					is_wildcard = (pattern == "_")
					if (((((pattern_kind == "INT") || (pattern_kind == "STRING")) || (pattern_kind == "FLOAT")) || (pattern_kind == "true")) || (pattern_kind == "false")) {
						is_literal = true
					} else if ((pattern_kind == "IDENT") && (is_wildcard == false)) {
						vtag2 := _variant_tag(ll.Registry, pattern)
						_ = vtag2
						if (vtag2 > int64(0)) {
							is_variant = true
						} else {
							is_binding = true
						}
					}
				} else {
					is_binding = true
				}
			}
		}
		if is_wildcard {
			ll = _ladv(ll)
		} else if ((is_tuple_pattern == false) && (is_array_pattern == false)) {
			ll = _ladv(ll)
		}
		or_count := int64(0)
		_ = or_count
		or_name_0 := ""
		_ = or_name_0
		or_name_1 := ""
		_ = or_name_1
		or_name_2 := ""
		_ = or_name_2
		or_name_3 := ""
		_ = or_name_3
		if (is_variant && (_lk(ll) == "|")) {
			for (_lk(ll) == "|") {
				ll = _ladv(ll)
				ll = _lskip_nl(ll)
				if (_lk(ll) == "IDENT") {
					oname := _lcur(ll).Text
					_ = oname
					if (or_count == int64(0)) {
						or_name_0 = oname
					}
					if (or_count == int64(1)) {
						or_name_1 = oname
					}
					if (or_count == int64(2)) {
						or_name_2 = oname
					}
					if (or_count == int64(3)) {
						or_name_3 = oname
					}
					or_count = (or_count + int64(1))
					ll = _ladv(ll)
				}
				ll = _lskip_nl(ll)
			}
		}
		has_bindings := false
		_ = has_bindings
		bind_name_0 := ""
		_ = bind_name_0
		bind_name_1 := ""
		_ = bind_name_1
		bind_name_2 := ""
		_ = bind_name_2
		bind_name_3 := ""
		_ = bind_name_3
		bind_total := int64(0)
		_ = bind_total
		is_struct_destr := false
		_ = is_struct_destr
		struct_field_name_0 := ""
		_ = struct_field_name_0
		struct_field_name_1 := ""
		_ = struct_field_name_1
		struct_field_name_2 := ""
		_ = struct_field_name_2
		struct_field_name_3 := ""
		_ = struct_field_name_3
		tuple_bind_is_lit_0 := false
		_ = tuple_bind_is_lit_0
		tuple_bind_is_lit_1 := false
		_ = tuple_bind_is_lit_1
		tuple_bind_is_lit_2 := false
		_ = tuple_bind_is_lit_2
		tuple_bind_is_lit_3 := false
		_ = tuple_bind_is_lit_3
		tuple_bind_lit_val_0 := int64(0)
		_ = tuple_bind_lit_val_0
		tuple_bind_lit_val_1 := int64(0)
		_ = tuple_bind_lit_val_1
		tuple_bind_lit_val_2 := int64(0)
		_ = tuple_bind_lit_val_2
		tuple_bind_lit_val_3 := int64(0)
		_ = tuple_bind_lit_val_3
		nested_variant_0 := ""
		_ = nested_variant_0
		nested_variant_1 := ""
		_ = nested_variant_1
		nested_variant_2 := ""
		_ = nested_variant_2
		nested_variant_3 := ""
		_ = nested_variant_3
		nested_bind_0 := ""
		_ = nested_bind_0
		nested_bind_1 := ""
		_ = nested_bind_1
		nested_bind_2 := ""
		_ = nested_bind_2
		nested_bind_3 := ""
		_ = nested_bind_3
		arr_pat_fixed_count := int64(0)
		_ = arr_pat_fixed_count
		arr_pat_rest_name := ""
		_ = arr_pat_rest_name
		arr_pat_has_rest := false
		_ = arr_pat_has_rest
		if is_array_pattern {
			ll = _ladv(ll)
			ll = _lskip_nl(ll)
			for ((_lk(ll) != "]") && (_lk(ll) != "EOF")) {
				if (_lk(ll) == ".") {
					dot_pos := (ll.Pos + int64(1))
					_ = dot_pos
					if ((dot_pos < int64(len(ll.Tokens))) && (token_name(ll.Tokens[dot_pos].Kind) == ".")) {
						ll = _ladv(ll)
						ll = _ladv(ll)
						arr_pat_has_rest = true
						if (_lk(ll) == "IDENT") {
							arr_pat_rest_name = _lcur(ll).Text
							ll = _ladv(ll)
						}
					} else {
						ll = _ladv(ll)
					}
				} else if ((_lk(ll) == "IDENT") || (_lk(ll) == "INT")) {
					btext := _lcur(ll).Text
					_ = btext
					if (bind_total == int64(0)) {
						bind_name_0 = btext
					}
					if (bind_total == int64(1)) {
						bind_name_1 = btext
					}
					if (bind_total == int64(2)) {
						bind_name_2 = btext
					}
					if (bind_total == int64(3)) {
						bind_name_3 = btext
					}
					bind_total = (bind_total + int64(1))
					has_bindings = true
					ll = _ladv(ll)
				} else {
					ll = _ladv(ll)
				}
				ll = _lskip_nl(ll)
				if (_lk(ll) == ",") {
					ll = _ladv(ll)
				}
				ll = _lskip_nl(ll)
			}
			if (_lk(ll) == "]") {
				ll = _ladv(ll)
			}
			arr_pat_fixed_count = bind_total
		} else if is_tuple_pattern {
			ll = _ladv(ll)
			ll = _lskip_nl(ll)
			for ((_lk(ll) != ")") && (_lk(ll) != "EOF")) {
				tk := _lk(ll)
				_ = tk
				if (tk == "INT") {
					lit_v := _parse_int_literal(_lcur(ll).Text)
					_ = lit_v
					if (bind_total == int64(0)) {
						tuple_bind_is_lit_0 = true
					}
					if (bind_total == int64(0)) {
						tuple_bind_lit_val_0 = lit_v
					}
					if (bind_total == int64(0)) {
						bind_name_0 = "_lit_"
					}
					if (bind_total == int64(1)) {
						tuple_bind_is_lit_1 = true
					}
					if (bind_total == int64(1)) {
						tuple_bind_lit_val_1 = lit_v
					}
					if (bind_total == int64(1)) {
						bind_name_1 = "_lit_"
					}
					if (bind_total == int64(2)) {
						tuple_bind_is_lit_2 = true
					}
					if (bind_total == int64(2)) {
						tuple_bind_lit_val_2 = lit_v
					}
					if (bind_total == int64(2)) {
						bind_name_2 = "_lit_"
					}
					if (bind_total == int64(3)) {
						tuple_bind_is_lit_3 = true
					}
					if (bind_total == int64(3)) {
						tuple_bind_lit_val_3 = lit_v
					}
					if (bind_total == int64(3)) {
						bind_name_3 = "_lit_"
					}
					bind_total = (bind_total + int64(1))
					has_bindings = true
					ll = _ladv(ll)
				} else if (tk == "IDENT") {
					btext := _lcur(ll).Text
					_ = btext
					if (btext == "_") {
						if (bind_total == int64(0)) {
							bind_name_0 = "_"
						}
						if (bind_total == int64(1)) {
							bind_name_1 = "_"
						}
						if (bind_total == int64(2)) {
							bind_name_2 = "_"
						}
						if (bind_total == int64(3)) {
							bind_name_3 = "_"
						}
					} else {
						if (bind_total == int64(0)) {
							bind_name_0 = btext
						}
						if (bind_total == int64(1)) {
							bind_name_1 = btext
						}
						if (bind_total == int64(2)) {
							bind_name_2 = btext
						}
						if (bind_total == int64(3)) {
							bind_name_3 = btext
						}
					}
					bind_total = (bind_total + int64(1))
					has_bindings = true
					ll = _ladv(ll)
				} else {
					ll = _ladv(ll)
				}
				ll = _lskip_nl(ll)
				if (_lk(ll) == ",") {
					ll = _ladv(ll)
				}
				ll = _lskip_nl(ll)
			}
			if (_lk(ll) == ")") {
				ll = _ladv(ll)
			}
		} else if (_lk(ll) == "(") {
			ll = _ladv(ll)
			ll = _lskip_nl(ll)
			for ((_lk(ll) != ")") && (_lk(ll) != "EOF")) {
				if (_lk(ll) == "IDENT") {
					btext := _lcur(ll).Text
					_ = btext
					nv_tag := _variant_tag(ll.Registry, btext)
					_ = nv_tag
					if (nv_tag > int64(0)) {
						if (bind_total == int64(0)) {
							nested_variant_0 = btext
						}
						if (bind_total == int64(0)) {
							bind_name_0 = btext
						}
						if (bind_total == int64(1)) {
							nested_variant_1 = btext
						}
						if (bind_total == int64(1)) {
							bind_name_1 = btext
						}
						if (bind_total == int64(2)) {
							nested_variant_2 = btext
						}
						if (bind_total == int64(2)) {
							bind_name_2 = btext
						}
						if (bind_total == int64(3)) {
							nested_variant_3 = btext
						}
						if (bind_total == int64(3)) {
							bind_name_3 = btext
						}
						ll = _ladv(ll)
						if (_lk(ll) == "(") {
							ll = _ladv(ll)
							ll = _lskip_nl(ll)
							if (_lk(ll) == "IDENT") {
								nb := _lcur(ll).Text
								_ = nb
								if (bind_total == int64(0)) {
									nested_bind_0 = nb
								}
								if (bind_total == int64(1)) {
									nested_bind_1 = nb
								}
								if (bind_total == int64(2)) {
									nested_bind_2 = nb
								}
								if (bind_total == int64(3)) {
									nested_bind_3 = nb
								}
								ll = _ladv(ll)
							}
							for ((_lk(ll) != ")") && (_lk(ll) != "EOF")) {
								ll = _ladv(ll)
							}
							if (_lk(ll) == ")") {
								ll = _ladv(ll)
							}
						}
						bind_total = (bind_total + int64(1))
						has_bindings = true
					} else {
						if (bind_total == int64(0)) {
							bind_name_0 = btext
						}
						if (bind_total == int64(1)) {
							bind_name_1 = btext
						}
						if (bind_total == int64(2)) {
							bind_name_2 = btext
						}
						if (bind_total == int64(3)) {
							bind_name_3 = btext
						}
						bind_total = (bind_total + int64(1))
						has_bindings = true
						ll = _ladv(ll)
					}
				} else {
					ll = _ladv(ll)
				}
				ll = _lskip_nl(ll)
				if (_lk(ll) == ",") {
					ll = _ladv(ll)
				}
				ll = _lskip_nl(ll)
			}
			if (_lk(ll) == ")") {
				ll = _ladv(ll)
			}
		} else if ((_lk(ll) == "{") && is_variant) {
			is_struct_destr = true
			ll = _ladv(ll)
			ll = _lskip_nl(ll)
			for ((_lk(ll) != "}") && (_lk(ll) != "EOF")) {
				if (_lk(ll) == "IDENT") {
					fname := _lcur(ll).Text
					_ = fname
					ll = _ladv(ll)
					bname := fname
					_ = bname
					if (_lk(ll) == ":") {
						ll = _ladv(ll)
						ll = _lskip_nl(ll)
						if (_lk(ll) == "IDENT") {
							bname = _lcur(ll).Text
							ll = _ladv(ll)
						}
					}
					if (bind_total == int64(0)) {
						bind_name_0 = bname
					}
					if (bind_total == int64(0)) {
						struct_field_name_0 = fname
					}
					if (bind_total == int64(1)) {
						bind_name_1 = bname
					}
					if (bind_total == int64(1)) {
						struct_field_name_1 = fname
					}
					if (bind_total == int64(2)) {
						bind_name_2 = bname
					}
					if (bind_total == int64(2)) {
						struct_field_name_2 = fname
					}
					if (bind_total == int64(3)) {
						bind_name_3 = bname
					}
					if (bind_total == int64(3)) {
						struct_field_name_3 = fname
					}
					bind_total = (bind_total + int64(1))
					has_bindings = true
				} else {
					ll = _ladv(ll)
				}
				ll = _lskip_nl(ll)
				if (_lk(ll) == ",") {
					ll = _ladv(ll)
				}
				ll = _lskip_nl(ll)
			}
			if (_lk(ll) == "}") {
				ll = _ladv(ll)
			}
		}
		has_guard := false
		_ = has_guard
		guard_start := int64(0)
		_ = guard_start
		guard_end := int64(0)
		_ = guard_end
		if (_lk(ll) == "if") {
			has_guard = true
			ll = _ladv(ll)
			guard_start = ll.Pos
			for ((_lk(ll) != "=>") && (_lk(ll) != "EOF")) {
				ll = _ladv(ll)
			}
			guard_end = ll.Pos
		}
		if (_lk(ll) == "=>") {
			ll = _ladv(ll)
		}
		ll = _lskip_nl(ll)
		if is_array_pattern {
			next_lbl := _lnew_label(ll)
			_ = next_lbl
			ll = next_lbl.L
			alen_nt := _lnew_temp(ll)
			_ = alen_nt
			ll = alen_nt.L
			ll = _lemit(ll, new_inst(IrOpOpArrayLen{}, alen_nt.Temp, subj.Temp, int64(0), "", int64(0)))
			expected_len := _lnew_temp(ll)
			_ = expected_len
			ll = expected_len.L
			ll = _lemit(ll, new_inst(IrOpOpConst{}, expected_len.Temp, arr_pat_fixed_count, int64(0), "", int64(0)))
			len_cmp := _lnew_temp(ll)
			_ = len_cmp
			ll = len_cmp.L
			if arr_pat_has_rest {
				ll = _lemit(ll, new_inst(IrOpOpGte{}, len_cmp.Temp, alen_nt.Temp, expected_len.Temp, "", int64(0)))
			} else {
				ll = _lemit(ll, new_inst(IrOpOpEq{}, len_cmp.Temp, alen_nt.Temp, expected_len.Temp, "", int64(0)))
			}
			ll = _lemit(ll, new_inst(IrOpOpBranchFalse{}, int64(0), len_cmp.Temp, next_lbl.Temp, "", int64(0)))
			ai := int64(0)
			_ = ai
			for (ai < arr_pat_fixed_count) {
				abname := ""
				_ = abname
				if (ai == int64(0)) {
					abname = bind_name_0
				}
				if (ai == int64(1)) {
					abname = bind_name_1
				}
				if (ai == int64(2)) {
					abname = bind_name_2
				}
				if (ai == int64(3)) {
					abname = bind_name_3
				}
				aidx_nt := _lnew_temp(ll)
				_ = aidx_nt
				ll = aidx_nt.L
				ll = _lemit(ll, new_inst(IrOpOpConst{}, aidx_nt.Temp, ai, int64(0), "", int64(0)))
				aelem_nt := _lnew_temp(ll)
				_ = aelem_nt
				ll = aelem_nt.L
				ll = _lemit(ll, new_inst(IrOpOpArrayGet{}, aelem_nt.Temp, subj.Temp, aidx_nt.Temp, "", int64(0)))
				if (abname != "_") {
					ll = _lenv_add(ll, abname, aelem_nt.Temp)
				}
				ai = (ai + int64(1))
			}
			if (arr_pat_has_rest && (arr_pat_rest_name != "")) {
				rest_start := _lnew_temp(ll)
				_ = rest_start
				ll = rest_start.L
				ll = _lemit(ll, new_inst(IrOpOpConst{}, rest_start.Temp, arr_pat_fixed_count, int64(0), "", int64(0)))
				rest_nt := _lnew_temp(ll)
				_ = rest_nt
				ll = rest_nt.L
				ll = _lemit(ll, new_inst(IrOpOpArraySlice{}, rest_nt.Temp, subj.Temp, rest_start.Temp, "", int64(0)))
				ll = _lenv_add(ll, arr_pat_rest_name, rest_nt.Temp)
			}
			if has_guard {
				saved_pos := ll.Pos
				_ = saved_pos
				ll = _lset_pos(ll, guard_start)
				guard_result := _lower_expr(ll)
				_ = guard_result
				ll = guard_result.L
				ll = _lset_pos(ll, saved_pos)
				ll = _lemit(ll, new_inst(IrOpOpBranchFalse{}, int64(0), guard_result.Temp, next_lbl.Temp, "", int64(0)))
			}
			if (_lk(ll) == lb) {
				ll = _lower_block(ll)
				ll = _lemit(ll, new_inst(IrOpOpConst{}, result_nt.Temp, int64(0), int64(0), "", int64(0)))
			} else {
				arm_result := _lower_expr(ll)
				_ = arm_result
				ll = arm_result.L
				ll = _lemit(ll, new_inst(IrOpOpStore{}, result_nt.Temp, arm_result.Temp, int64(0), "", int64(0)))
				result_len := (result_nt.Temp + int64(1))
				_ = result_len
				arm_len := (arm_result.Temp + int64(1))
				_ = arm_len
				ll = _lemit(ll, new_inst(IrOpOpStore{}, result_len, arm_len, int64(0), "", int64(0)))
			}
			ll = _lemit(ll, new_inst(IrOpOpJump{}, int64(0), end_lbl.Temp, int64(0), "", int64(0)))
			ll = _lemit(ll, new_inst(IrOpOpLabel{}, int64(0), next_lbl.Temp, int64(0), "", int64(0)))
		} else if is_tuple_pattern {
			next_lbl := _lnew_label(ll)
			_ = next_lbl
			ll = next_lbl.L
			ti := int64(0)
			_ = ti
			for (ti < bind_total) {
				tbname := ""
				_ = tbname
				if (ti == int64(0)) {
					tbname = bind_name_0
				}
				if (ti == int64(1)) {
					tbname = bind_name_1
				}
				if (ti == int64(2)) {
					tbname = bind_name_2
				}
				if (ti == int64(3)) {
					tbname = bind_name_3
				}
				fnt := _lnew_temp(ll)
				_ = fnt
				ll = fnt.L
				ll = _lemit(ll, new_inst(IrOpOpFieldGet{}, fnt.Temp, subj.Temp, ti, "", int64(0)))
				is_lit_elem := false
				_ = is_lit_elem
				lit_elem_val := int64(0)
				_ = lit_elem_val
				if (ti == int64(0)) {
					is_lit_elem = tuple_bind_is_lit_0
				}
				if (ti == int64(0)) {
					lit_elem_val = tuple_bind_lit_val_0
				}
				if (ti == int64(1)) {
					is_lit_elem = tuple_bind_is_lit_1
				}
				if (ti == int64(1)) {
					lit_elem_val = tuple_bind_lit_val_1
				}
				if (ti == int64(2)) {
					is_lit_elem = tuple_bind_is_lit_2
				}
				if (ti == int64(2)) {
					lit_elem_val = tuple_bind_lit_val_2
				}
				if (ti == int64(3)) {
					is_lit_elem = tuple_bind_is_lit_3
				}
				if (ti == int64(3)) {
					lit_elem_val = tuple_bind_lit_val_3
				}
				if is_lit_elem {
					clit := _lnew_temp(ll)
					_ = clit
					ll = clit.L
					ll = _lemit(ll, new_inst(IrOpOpConst{}, clit.Temp, lit_elem_val, int64(0), "", int64(0)))
					tcmp := _lnew_temp(ll)
					_ = tcmp
					ll = tcmp.L
					ll = _lemit(ll, new_inst(IrOpOpEq{}, tcmp.Temp, fnt.Temp, clit.Temp, "", int64(0)))
					ll = _lemit(ll, new_inst(IrOpOpBranchFalse{}, int64(0), tcmp.Temp, next_lbl.Temp, "", int64(0)))
				} else if (tbname != "_") {
					ll = _lenv_add(ll, tbname, fnt.Temp)
				}
				ti = (ti + int64(1))
			}
			if (at_bind_name != "") {
				ll = _lenv_add(ll, at_bind_name, subj.Temp)
			}
			if has_guard {
				saved_pos := ll.Pos
				_ = saved_pos
				ll = _lset_pos(ll, guard_start)
				guard_result := _lower_expr(ll)
				_ = guard_result
				ll = guard_result.L
				ll = _lset_pos(ll, saved_pos)
				ll = _lemit(ll, new_inst(IrOpOpBranchFalse{}, int64(0), guard_result.Temp, next_lbl.Temp, "", int64(0)))
			}
			if (_lk(ll) == lb) {
				ll = _lower_block(ll)
				ll = _lemit(ll, new_inst(IrOpOpConst{}, result_nt.Temp, int64(0), int64(0), "", int64(0)))
			} else {
				arm_result := _lower_expr(ll)
				_ = arm_result
				ll = arm_result.L
				ll = _lemit(ll, new_inst(IrOpOpStore{}, result_nt.Temp, arm_result.Temp, int64(0), "", int64(0)))
				result_len := (result_nt.Temp + int64(1))
				_ = result_len
				arm_len := (arm_result.Temp + int64(1))
				_ = arm_len
				ll = _lemit(ll, new_inst(IrOpOpStore{}, result_len, arm_len, int64(0), "", int64(0)))
			}
			ll = _lemit(ll, new_inst(IrOpOpJump{}, int64(0), end_lbl.Temp, int64(0), "", int64(0)))
			ll = _lemit(ll, new_inst(IrOpOpLabel{}, int64(0), next_lbl.Temp, int64(0), "", int64(0)))
		} else if (is_wildcard || is_binding) {
			if is_binding {
				ll = _lenv_add(ll, pattern, subj.Temp)
			}
			if (at_bind_name != "") {
				ll = _lenv_add(ll, at_bind_name, subj.Temp)
			}
			if has_guard {
				next_lbl := _lnew_label(ll)
				_ = next_lbl
				ll = next_lbl.L
				saved_pos := ll.Pos
				_ = saved_pos
				ll = _lset_pos(ll, guard_start)
				guard_result := _lower_expr(ll)
				_ = guard_result
				ll = guard_result.L
				ll = _lset_pos(ll, saved_pos)
				ll = _lemit(ll, new_inst(IrOpOpBranchFalse{}, int64(0), guard_result.Temp, next_lbl.Temp, "", int64(0)))
				if (_lk(ll) == lb) {
					ll = _lower_block(ll)
					ll = _lemit(ll, new_inst(IrOpOpConst{}, result_nt.Temp, int64(0), int64(0), "", int64(0)))
				} else {
					arm_result := _lower_expr(ll)
					_ = arm_result
					ll = arm_result.L
					ll = _lemit(ll, new_inst(IrOpOpStore{}, result_nt.Temp, arm_result.Temp, int64(0), "", int64(0)))
					result_len := (result_nt.Temp + int64(1))
					_ = result_len
					arm_len := (arm_result.Temp + int64(1))
					_ = arm_len
					ll = _lemit(ll, new_inst(IrOpOpStore{}, result_len, arm_len, int64(0), "", int64(0)))
				}
				ll = _lemit(ll, new_inst(IrOpOpJump{}, int64(0), end_lbl.Temp, int64(0), "", int64(0)))
				ll = _lemit(ll, new_inst(IrOpOpLabel{}, int64(0), next_lbl.Temp, int64(0), "", int64(0)))
			} else {
				if (_lk(ll) == lb) {
					ll = _lower_block(ll)
					ll = _lemit(ll, new_inst(IrOpOpConst{}, result_nt.Temp, int64(0), int64(0), "", int64(0)))
				} else {
					arm_result := _lower_expr(ll)
					_ = arm_result
					ll = arm_result.L
					ll = _lemit(ll, new_inst(IrOpOpStore{}, result_nt.Temp, arm_result.Temp, int64(0), "", int64(0)))
					result_len := (result_nt.Temp + int64(1))
					_ = result_len
					arm_len := (arm_result.Temp + int64(1))
					_ = arm_len
					ll = _lemit(ll, new_inst(IrOpOpStore{}, result_len, arm_len, int64(0), "", int64(0)))
				}
				ll = _lemit(ll, new_inst(IrOpOpJump{}, int64(0), end_lbl.Temp, int64(0), "", int64(0)))
			}
		} else if is_literal {
			next_lbl := _lnew_label(ll)
			_ = next_lbl
			ll = next_lbl.L
			lit_done := false
			_ = lit_done
			if (pattern_kind == "FLOAT") {
				fll := ll
				_ = fll
				fstr_idx := mod_find_string(fll.Module, pattern)
				_ = fstr_idx
				if (fstr_idx == int64(0)) {
					fnew_mod := mod_add_string(fll.Module, pattern)
					_ = fnew_mod
					fstr_idx = (int64(len(fnew_mod.String_constants)) - int64(1))
					fll = Lowerer{Tokens: fll.Tokens, Pos: fll.Pos, Store: fll.Store, Registry: fll.Registry, Treg: fll.Treg, Pool: fll.Pool, Table: fll.Table, File: fll.File, Module: fnew_mod, Current_func: fll.Current_func, Temp_counter: fll.Temp_counter, Label_counter: fll.Label_counter, Env_names: fll.Env_names, Env_slots: fll.Env_slots, Loop_start: fll.Loop_start, Loop_end: fll.Loop_end, Const_names: fll.Const_names, Const_vals: fll.Const_vals, Const_str_names: fll.Const_str_names, Const_str_vals: fll.Const_str_vals, Temp_types: fll.Temp_types, Fn_error_type: fll.Fn_error_type, Defer_starts: fll.Defer_starts, Defer_ends: fll.Defer_ends}
				}
				ll = fll
				flit_nt := _lnew_temp(ll)
				_ = flit_nt
				ll = flit_nt.L
				ll = _lemit(ll, new_inst(IrOpOpConstStr{}, flit_nt.Temp, fstr_idx, (int64(0) - int64(9)), "", int64(0)))
				ll = _lset_type(ll, flit_nt.Temp, int64(9))
				cmp_nt := _lnew_temp(ll)
				_ = cmp_nt
				ll = cmp_nt.L
				ll = _lemit(ll, new_inst(IrOpOpEq{}, cmp_nt.Temp, subj.Temp, flit_nt.Temp, "f64", int64(0)))
				ll = _lemit(ll, new_inst(IrOpOpBranchFalse{}, int64(0), cmp_nt.Temp, next_lbl.Temp, "", int64(0)))
				lit_done = true
			}
			if ((pattern_kind == "INT") && (lit_done == false)) {
				lit_nt := _lnew_temp(ll)
				_ = lit_nt
				ll = lit_nt.L
				lit_val := _parse_int_literal(pattern)
				_ = lit_val
				ll = _lemit(ll, new_inst(IrOpOpConst{}, lit_nt.Temp, lit_val, int64(0), "", int64(0)))
				cmp_nt := _lnew_temp(ll)
				_ = cmp_nt
				ll = cmp_nt.L
				ll = _lemit(ll, new_inst(IrOpOpEq{}, cmp_nt.Temp, subj.Temp, lit_nt.Temp, "", int64(0)))
				ll = _lemit(ll, new_inst(IrOpOpBranchFalse{}, int64(0), cmp_nt.Temp, next_lbl.Temp, "", int64(0)))
				lit_done = true
			}
			if ((pattern_kind == "true") && (lit_done == false)) {
				lit_nt := _lnew_temp(ll)
				_ = lit_nt
				ll = lit_nt.L
				ll = _lemit(ll, new_inst(IrOpOpConst{}, lit_nt.Temp, int64(1), int64(0), "", int64(0)))
				cmp_nt := _lnew_temp(ll)
				_ = cmp_nt
				ll = cmp_nt.L
				ll = _lemit(ll, new_inst(IrOpOpEq{}, cmp_nt.Temp, subj.Temp, lit_nt.Temp, "", int64(0)))
				ll = _lemit(ll, new_inst(IrOpOpBranchFalse{}, int64(0), cmp_nt.Temp, next_lbl.Temp, "", int64(0)))
				lit_done = true
			}
			if ((pattern_kind == "false") && (lit_done == false)) {
				lit_nt := _lnew_temp(ll)
				_ = lit_nt
				ll = lit_nt.L
				ll = _lemit(ll, new_inst(IrOpOpConst{}, lit_nt.Temp, int64(0), int64(0), "", int64(0)))
				cmp_nt := _lnew_temp(ll)
				_ = cmp_nt
				ll = cmp_nt.L
				ll = _lemit(ll, new_inst(IrOpOpEq{}, cmp_nt.Temp, subj.Temp, lit_nt.Temp, "", int64(0)))
				ll = _lemit(ll, new_inst(IrOpOpBranchFalse{}, int64(0), cmp_nt.Temp, next_lbl.Temp, "", int64(0)))
				lit_done = true
			}
			if ((pattern_kind == "STRING") && (lit_done == false)) {
				lit_result := _lower_string_const(ll, pattern)
				_ = lit_result
				ll = lit_result.L
				cmp_nt := _lnew_temp(ll)
				_ = cmp_nt
				ll = cmp_nt.L
				ll = _lemit(ll, new_inst(IrOpOpStrEq{}, cmp_nt.Temp, subj.Temp, lit_result.Temp, "", int64(0)))
				ll = _lemit(ll, new_inst(IrOpOpBranchFalse{}, int64(0), cmp_nt.Temp, next_lbl.Temp, "", int64(0)))
				lit_done = true
			}
			if (at_bind_name != "") {
				ll = _lenv_add(ll, at_bind_name, subj.Temp)
			}
			if has_guard {
				saved_pos := ll.Pos
				_ = saved_pos
				ll = _lset_pos(ll, guard_start)
				guard_result := _lower_expr(ll)
				_ = guard_result
				ll = guard_result.L
				ll = _lset_pos(ll, saved_pos)
				ll = _lemit(ll, new_inst(IrOpOpBranchFalse{}, int64(0), guard_result.Temp, next_lbl.Temp, "", int64(0)))
			}
			if (_lk(ll) == lb) {
				ll = _lower_block(ll)
				ll = _lemit(ll, new_inst(IrOpOpConst{}, result_nt.Temp, int64(0), int64(0), "", int64(0)))
			} else {
				arm_result := _lower_expr(ll)
				_ = arm_result
				ll = arm_result.L
				ll = _lemit(ll, new_inst(IrOpOpStore{}, result_nt.Temp, arm_result.Temp, int64(0), "", int64(0)))
				result_len := (result_nt.Temp + int64(1))
				_ = result_len
				arm_len := (arm_result.Temp + int64(1))
				_ = arm_len
				ll = _lemit(ll, new_inst(IrOpOpStore{}, result_len, arm_len, int64(0), "", int64(0)))
			}
			ll = _lemit(ll, new_inst(IrOpOpJump{}, int64(0), end_lbl.Temp, int64(0), "", int64(0)))
			ll = _lemit(ll, new_inst(IrOpOpLabel{}, int64(0), next_lbl.Temp, int64(0), "", int64(0)))
		} else {
			tag := _variant_tag(ll.Registry, pattern)
			_ = tag
			next_lbl := _lnew_label(ll)
			_ = next_lbl
			ll = next_lbl.L
			if (or_count > int64(0)) {
				body_lbl := _lnew_label(ll)
				_ = body_lbl
				ll = body_lbl.L
				tag_nt := _lnew_temp(ll)
				_ = tag_nt
				ll = tag_nt.L
				ll = _lemit(ll, new_inst(IrOpOpConst{}, tag_nt.Temp, tag, int64(0), "", int64(0)))
				cmp_nt := _lnew_temp(ll)
				_ = cmp_nt
				ll = cmp_nt.L
				ll = _lemit(ll, new_inst(IrOpOpEq{}, cmp_nt.Temp, subj_tag_temp, tag_nt.Temp, "", int64(0)))
				ll = _lemit(ll, new_inst(IrOpOpBranchTrue{}, int64(0), cmp_nt.Temp, body_lbl.Temp, "", int64(0)))
				oi := int64(0)
				_ = oi
				for (oi < or_count) {
					oname := ""
					_ = oname
					if (oi == int64(0)) {
						oname = or_name_0
					}
					if (oi == int64(1)) {
						oname = or_name_1
					}
					if (oi == int64(2)) {
						oname = or_name_2
					}
					if (oi == int64(3)) {
						oname = or_name_3
					}
					otag := _variant_tag(ll.Registry, oname)
					_ = otag
					otag_nt := _lnew_temp(ll)
					_ = otag_nt
					ll = otag_nt.L
					ll = _lemit(ll, new_inst(IrOpOpConst{}, otag_nt.Temp, otag, int64(0), "", int64(0)))
					ocmp_nt := _lnew_temp(ll)
					_ = ocmp_nt
					ll = ocmp_nt.L
					ll = _lemit(ll, new_inst(IrOpOpEq{}, ocmp_nt.Temp, subj_tag_temp, otag_nt.Temp, "", int64(0)))
					if (oi == (or_count - int64(1))) {
						ll = _lemit(ll, new_inst(IrOpOpBranchFalse{}, int64(0), ocmp_nt.Temp, next_lbl.Temp, "", int64(0)))
					} else {
						ll = _lemit(ll, new_inst(IrOpOpBranchTrue{}, int64(0), ocmp_nt.Temp, body_lbl.Temp, "", int64(0)))
					}
					oi = (oi + int64(1))
				}
				ll = _lemit(ll, new_inst(IrOpOpLabel{}, int64(0), body_lbl.Temp, int64(0), "", int64(0)))
			} else {
				tag_nt := _lnew_temp(ll)
				_ = tag_nt
				ll = tag_nt.L
				ll = _lemit(ll, new_inst(IrOpOpConst{}, tag_nt.Temp, tag, int64(0), "", int64(0)))
				cmp_nt := _lnew_temp(ll)
				_ = cmp_nt
				ll = cmp_nt.L
				ll = _lemit(ll, new_inst(IrOpOpEq{}, cmp_nt.Temp, subj_tag_temp, tag_nt.Temp, "", int64(0)))
				ll = _lemit(ll, new_inst(IrOpOpBranchFalse{}, int64(0), cmp_nt.Temp, next_lbl.Temp, "", int64(0)))
			}
			if (at_bind_name != "") {
				ll = _lenv_add(ll, at_bind_name, subj.Temp)
			}
			if (has_bindings && is_struct_destr) {
				parent := _variant_parent(ll.Registry, pattern)
				_ = parent
				bi := int64(0)
				_ = bi
				for (bi < bind_total) {
					bname := ""
					_ = bname
					sfname := ""
					_ = sfname
					if (bi == int64(0)) {
						bname = bind_name_0
					}
					if (bi == int64(0)) {
						sfname = struct_field_name_0
					}
					if (bi == int64(1)) {
						bname = bind_name_1
					}
					if (bi == int64(1)) {
						sfname = struct_field_name_1
					}
					if (bi == int64(2)) {
						bname = bind_name_2
					}
					if (bi == int64(2)) {
						sfname = struct_field_name_2
					}
					if (bi == int64(3)) {
						bname = bind_name_3
					}
					if (bi == int64(3)) {
						sfname = struct_field_name_3
					}
					field_pos := int64(0)
					_ = field_pos
					found_field := false
					_ = found_field
					if (parent.Name != "") {
						fi := int64(0)
						_ = fi
						scan_done := false
						_ = scan_done
						for ((fi < int64(8)) && (scan_done == false)) {
							vft_name := variant_field_name(parent, tag, fi)
							_ = vft_name
							if (vft_name == "") {
								scan_done = true
							}
							if (vft_name == sfname) {
								field_pos = fi
							}
							if (vft_name == sfname) {
								found_field = true
							}
							if (vft_name == sfname) {
								scan_done = true
							}
							fi = (fi + int64(1))
						}
					}
					data_slot := int64(1)
					_ = data_slot
					si := int64(0)
					_ = si
					for (si < field_pos) {
						ft := variant_field_type(parent, tag, si)
						_ = ft
						if (ft == int64(12)) {
							data_slot = (data_slot + int64(2))
						} else {
							data_slot = (data_slot + int64(1))
						}
						si = (si + int64(1))
					}
					ft := variant_field_type(parent, tag, field_pos)
					_ = ft
					if (ft == int64(12)) {
						bnt := _lnew_str_temp(ll)
						_ = bnt
						ll = bnt.L
						ll = _lemit(ll, new_inst(IrOpOpFieldGet{}, bnt.Temp, subj.Temp, data_slot, "", int64(0)))
						ll = _lemit(ll, new_inst(IrOpOpFieldGet{}, (bnt.Temp + int64(1)), subj.Temp, (data_slot + int64(1)), "", int64(0)))
						ll = _lset_type(ll, bnt.Temp, int64(12))
						ll = _lenv_add(ll, bname, bnt.Temp)
					} else {
						bnt := _lnew_temp(ll)
						_ = bnt
						ll = bnt.L
						ll = _lemit(ll, new_inst(IrOpOpFieldGet{}, bnt.Temp, subj.Temp, data_slot, "", int64(0)))
						if (ft != int64(0)) {
							ll = _lset_type(ll, bnt.Temp, ft)
						}
						ll = _lenv_add(ll, bname, bnt.Temp)
					}
					bi = (bi + int64(1))
				}
			} else if has_bindings {
				parent := _variant_parent(ll.Registry, pattern)
				_ = parent
				bi := int64(0)
				_ = bi
				for (bi < bind_total) {
					bname := ""
					_ = bname
					if (bi == int64(0)) {
						bname = bind_name_0
					}
					if (bi == int64(1)) {
						bname = bind_name_1
					}
					if (bi == int64(2)) {
						bname = bind_name_2
					}
					if (bi == int64(3)) {
						bname = bind_name_3
					}
					data_slot := int64(1)
					_ = data_slot
					si := int64(0)
					_ = si
					for (si < bi) {
						ft := variant_field_type(parent, tag, si)
						_ = ft
						if (ft == int64(12)) {
							data_slot = (data_slot + int64(2))
						} else {
							data_slot = (data_slot + int64(1))
						}
						si = (si + int64(1))
					}
					ft := variant_field_type(parent, tag, bi)
					_ = ft
					nv := ""
					_ = nv
					nb := ""
					_ = nb
					if (bi == int64(0)) {
						nv = nested_variant_0
					}
					if (bi == int64(0)) {
						nb = nested_bind_0
					}
					if (bi == int64(1)) {
						nv = nested_variant_1
					}
					if (bi == int64(1)) {
						nb = nested_bind_1
					}
					if (bi == int64(2)) {
						nv = nested_variant_2
					}
					if (bi == int64(2)) {
						nb = nested_bind_2
					}
					if (bi == int64(3)) {
						nv = nested_variant_3
					}
					if (bi == int64(3)) {
						nb = nested_bind_3
					}
					if (nv != "") {
						inner_nt := _lnew_temp(ll)
						_ = inner_nt
						ll = inner_nt.L
						ll = _lemit(ll, new_inst(IrOpOpFieldGet{}, inner_nt.Temp, subj.Temp, data_slot, "", int64(0)))
						inner_tag := _variant_tag(ll.Registry, nv)
						_ = inner_tag
						inner_tag_nt := _lnew_temp(ll)
						_ = inner_tag_nt
						ll = inner_tag_nt.L
						inner_tag_val := _lnew_temp(ll)
						_ = inner_tag_val
						ll = inner_tag_val.L
						ll = _lemit(ll, new_inst(IrOpOpFieldGet{}, inner_tag_val.Temp, inner_nt.Temp, int64(0), "_tag", int64(0)))
						itag_const := _lnew_temp(ll)
						_ = itag_const
						ll = itag_const.L
						ll = _lemit(ll, new_inst(IrOpOpConst{}, itag_const.Temp, inner_tag, int64(0), "", int64(0)))
						icmp := _lnew_temp(ll)
						_ = icmp
						ll = icmp.L
						ll = _lemit(ll, new_inst(IrOpOpEq{}, icmp.Temp, inner_tag_val.Temp, itag_const.Temp, "", int64(0)))
						ll = _lemit(ll, new_inst(IrOpOpBranchFalse{}, int64(0), icmp.Temp, next_lbl.Temp, "", int64(0)))
						if (nb != "") {
							inner_parent := _variant_parent(ll.Registry, nv)
							_ = inner_parent
							inner_ft := variant_field_type(inner_parent, inner_tag, int64(0))
							_ = inner_ft
							if (inner_ft == int64(12)) {
								ibnt := _lnew_str_temp(ll)
								_ = ibnt
								ll = ibnt.L
								ll = _lemit(ll, new_inst(IrOpOpFieldGet{}, ibnt.Temp, inner_nt.Temp, int64(1), "", int64(0)))
								ll = _lemit(ll, new_inst(IrOpOpFieldGet{}, (ibnt.Temp + int64(1)), inner_nt.Temp, int64(2), "", int64(0)))
								ll = _lset_type(ll, ibnt.Temp, int64(12))
								ll = _lenv_add(ll, nb, ibnt.Temp)
							} else {
								ibnt := _lnew_temp(ll)
								_ = ibnt
								ll = ibnt.L
								ll = _lemit(ll, new_inst(IrOpOpFieldGet{}, ibnt.Temp, inner_nt.Temp, int64(1), "", int64(0)))
								if (inner_ft != int64(0)) {
									ll = _lset_type(ll, ibnt.Temp, inner_ft)
								}
								ll = _lenv_add(ll, nb, ibnt.Temp)
							}
						}
					} else if (ft == int64(12)) {
						bnt := _lnew_str_temp(ll)
						_ = bnt
						ll = bnt.L
						ll = _lemit(ll, new_inst(IrOpOpFieldGet{}, bnt.Temp, subj.Temp, data_slot, "", int64(0)))
						ll = _lemit(ll, new_inst(IrOpOpFieldGet{}, (bnt.Temp + int64(1)), subj.Temp, (data_slot + int64(1)), "", int64(0)))
						ll = _lset_type(ll, bnt.Temp, int64(12))
						ll = _lenv_add(ll, bname, bnt.Temp)
					} else {
						bnt := _lnew_temp(ll)
						_ = bnt
						ll = bnt.L
						ll = _lemit(ll, new_inst(IrOpOpFieldGet{}, bnt.Temp, subj.Temp, data_slot, "", int64(0)))
						if (ft != int64(0)) {
							ll = _lset_type(ll, bnt.Temp, ft)
						}
						ll = _lenv_add(ll, bname, bnt.Temp)
					}
					bi = (bi + int64(1))
				}
			}
			if has_guard {
				saved_pos := ll.Pos
				_ = saved_pos
				ll = _lset_pos(ll, guard_start)
				guard_result := _lower_expr(ll)
				_ = guard_result
				ll = guard_result.L
				ll = _lset_pos(ll, saved_pos)
				ll = _lemit(ll, new_inst(IrOpOpBranchFalse{}, int64(0), guard_result.Temp, next_lbl.Temp, "", int64(0)))
			}
			if (_lk(ll) == lb) {
				ll = _lower_block(ll)
				ll = _lemit(ll, new_inst(IrOpOpConst{}, result_nt.Temp, int64(0), int64(0), "", int64(0)))
			} else {
				arm_result := _lower_expr(ll)
				_ = arm_result
				ll = arm_result.L
				ll = _lemit(ll, new_inst(IrOpOpStore{}, result_nt.Temp, arm_result.Temp, int64(0), "", int64(0)))
				result_len := (result_nt.Temp + int64(1))
				_ = result_len
				arm_len := (arm_result.Temp + int64(1))
				_ = arm_len
				ll = _lemit(ll, new_inst(IrOpOpStore{}, result_len, arm_len, int64(0), "", int64(0)))
			}
			ll = _lemit(ll, new_inst(IrOpOpJump{}, int64(0), end_lbl.Temp, int64(0), "", int64(0)))
			ll = _lemit(ll, new_inst(IrOpOpLabel{}, int64(0), next_lbl.Temp, int64(0), "", int64(0)))
		}
		ll = _lskip_nl(ll)
	}
	if (_lk(ll) == rb) {
		ll = _ladv(ll)
	}
	ll = _lemit(ll, new_inst(IrOpOpLabel{}, int64(0), end_lbl.Temp, int64(0), "", int64(0)))
	return LT{L: ll, Temp: result_nt.Temp}
}

func _skip_braces_l(l Lowerer) Lowerer {
	lb := "{"
	_ = lb
	rb := "}"
	_ = rb
	if (_lk(l) != lb) {
		return l
	}
	ll := _ladv(l)
	_ = ll
	d := int64(1)
	_ = d
	for ((d > int64(0)) && (_lk(ll) != "EOF")) {
		if (_lk(ll) == lb) {
			d = (d + int64(1))
		}
		if (_lk(ll) == rb) {
			d = (d - int64(1))
		}
		if (d > int64(0)) {
			ll = _ladv(ll)
		}
	}
	if (_lk(ll) == rb) {
		ll = _ladv(ll)
	}
	return ll
}

func _lower_assert(l Lowerer) Lowerer {
	ll := _ladv(l)
	_ = ll
	result := _lower_expr(ll)
	_ = result
	ll = result.L
	ok_lbl := _lnew_label(ll)
	_ = ok_lbl
	ll = ok_lbl.L
	ll = _lemit(ll, new_inst(IrOpOpBranchTrue{}, int64(0), result.Temp, ok_lbl.Temp, "", int64(0)))
	nt := _lnew_temp(ll)
	_ = nt
	ll = nt.L
	ll = _lemit(ll, new_inst(IrOpOpConst{}, nt.Temp, int64(1), int64(0), "", int64(0)))
	ll = _lemit(ll, new_inst(IrOpOpCall{}, int64(0), nt.Temp, int64(1), "_aria_exit", int64(0)))
	ll = _lemit(ll, new_inst(IrOpOpLabel{}, int64(0), ok_lbl.Temp, int64(0), "", int64(0)))
	return ll
}

func _skip_type_l(l Lowerer) Lowerer {
	ll := l
	_ = ll
	k := _lk(ll)
	_ = k
	if ((k == "IDENT") || (k == "Self")) {
		ll = _ladv(ll)
		for (_lk(ll) == ".") {
			ll = _ladv(_ladv(ll))
		}
		if (_lk(ll) == "[") {
			d := int64(1)
			_ = d
			ll = _ladv(ll)
			for ((d > int64(0)) && (_lk(ll) != "EOF")) {
				if (_lk(ll) == "[") {
					d = (d + int64(1))
				}
				if (_lk(ll) == "]") {
					d = (d - int64(1))
				}
				if (d > int64(0)) {
					ll = _ladv(ll)
				}
			}
			if (_lk(ll) == "]") {
				ll = _ladv(ll)
			}
		}
		if (_lk(ll) == "?") {
			ll = _ladv(ll)
		}
		if (_lk(ll) == "!") {
			ll = _ladv(ll)
			ll = _skip_type_l(ll)
		}
		return ll
	}
	if (k == "[") {
		ll = _ladv(ll)
		ll = _skip_type_l(ll)
		if (_lk(ll) == "]") {
			ll = _ladv(ll)
		}
		return ll
	}
	if (k == "fn") {
		ll = _ladv(ll)
		if ((_lk(ll) == "IDENT") && (_lcur(ll).Text == "once")) {
			ll = _ladv(ll)
		}
		if (_lk(ll) == "(") {
			ll = _ladv(ll)
			for ((_lk(ll) != ")") && (_lk(ll) != "EOF")) {
				ll = _skip_type_l(ll)
				if (_lk(ll) == ",") {
					ll = _ladv(ll)
				}
			}
			if (_lk(ll) == ")") {
				ll = _ladv(ll)
			}
		}
		if (_lk(ll) == "->") {
			ll = _ladv(ll)
			ll = _skip_type_l(ll)
		}
		return ll
	}
	return ll
}

func _lower_derives(l Lowerer) Lowerer {
	ll := l
	_ = ll
	si := int64(1)
	_ = si
	for (si < int64(len(ll.Registry.Struct_defs))) {
		def := ll.Registry.Struct_defs[si]
		_ = def
		if (((def.Name != "") && (def.Is_sum == false)) && (int64(len(def.Fields)) > int64(1))) {
			if treg_has_impl(ll.Treg, def.Name, "Eq") {
				eq_name := (def.Name + "_eq")
				_ = eq_name
				existing := reg_find_fn(ll.Registry, eq_name)
				_ = existing
				if (existing.Name == "") {
					ll = _lower_derive_eq(ll, def)
				}
			}
			if treg_has_impl(ll.Treg, def.Name, "Clone") {
				clone_name := (def.Name + "_clone")
				_ = clone_name
				existing := reg_find_fn(ll.Registry, clone_name)
				_ = existing
				if (existing.Name == "") {
					ll = _lower_derive_clone(ll, def)
				}
			}
			if treg_has_impl(ll.Treg, def.Name, "Debug") {
				debug_name := (def.Name + "_debug_str")
				_ = debug_name
				existing := reg_find_fn(ll.Registry, debug_name)
				_ = existing
				if (existing.Name == "") {
					ll = _lower_derive_debug(ll, def)
				}
			}
			if treg_has_impl(ll.Treg, def.Name, "Display") {
				display_name := (def.Name + "_display")
				_ = display_name
				existing := reg_find_fn(ll.Registry, display_name)
				_ = existing
				if (existing.Name == "") {
					ll = _lower_derive_display(ll, def)
				}
			}
			if treg_has_impl(ll.Treg, def.Name, "Default") {
				default_name := (def.Name + "_default")
				_ = default_name
				existing := reg_find_fn(ll.Registry, default_name)
				_ = existing
				if (existing.Name == "") {
					ll = _lower_derive_default(ll, def)
				}
			}
			if treg_has_impl(ll.Treg, def.Name, "Ord") {
				cmp_name := (def.Name + "_cmp")
				_ = cmp_name
				existing := reg_find_fn(ll.Registry, cmp_name)
				_ = existing
				if (existing.Name == "") {
					ll = _lower_derive_ord(ll, def)
				}
			}
			if treg_has_impl(ll.Treg, def.Name, "Hash") {
				hash_name := (def.Name + "_hash")
				_ = hash_name
				existing := reg_find_fn(ll.Registry, hash_name)
				_ = existing
				if (existing.Name == "") {
					ll = _lower_derive_hash(ll, def)
				}
			}
		}
		si = (si + int64(1))
	}
	return ll
}

func _lower_derive_eq(l Lowerer, def StructDef) Lowerer {
	fname := (def.Name + "_eq")
	_ = fname
	ll := _lset_func(l, new_ir_func(fname, int64(2), int64(0)))
	_ = ll
	self_nt := _lnew_temp(ll)
	_ = self_nt
	ll = self_nt.L
	ll = _lemit(ll, new_inst(IrOpOpArg{}, self_nt.Temp, int64(0), int64(0), "", int64(0)))
	other_nt := _lnew_temp(ll)
	_ = other_nt
	ll = other_nt.L
	ll = _lemit(ll, new_inst(IrOpOpArg{}, other_nt.Temp, int64(1), int64(0), "", int64(0)))
	result_nt := _lnew_temp(ll)
	_ = result_nt
	ll = result_nt.L
	ll = _lower_struct_eq(ll, self_nt.Temp, other_nt.Temp, def.Name, result_nt.Temp)
	ll = _lemit(ll, new_inst(IrOpOpRet{}, int64(0), result_nt.Temp, int64(0), "", int64(0)))
	ll = _lfinish_func(ll)
	return ll
}

func _lower_derive_clone(l Lowerer, def StructDef) Lowerer {
	fname := (def.Name + "_clone")
	_ = fname
	ll := _lset_func(l, new_ir_func(fname, int64(1), int64(0)))
	_ = ll
	self_nt := _lnew_temp(ll)
	_ = self_nt
	ll = self_nt.L
	ll = _lemit(ll, new_inst(IrOpOpArg{}, self_nt.Temp, int64(0), int64(0), "", int64(0)))
	alloc_size := (int64(len(def.Fields)) - int64(1))
	_ = alloc_size
	new_nt := _lnew_temp(ll)
	_ = new_nt
	ll = new_nt.L
	ll = _lemit(ll, new_inst(IrOpOpAlloc{}, new_nt.Temp, alloc_size, int64(0), def.Name, int64(0)))
	fi := int64(1)
	_ = fi
	slot := int64(0)
	_ = slot
	for (fi < int64(len(def.Fields))) {
		fld := def.Fields[fi]
		_ = fld
		ft := fld.Type_id
		_ = ft
		if (ft == int64(12)) {
			src_val := _lnew_temp(ll)
			_ = src_val
			ll = src_val.L
			ll = _lemit(ll, new_inst(IrOpOpFieldGet{}, src_val.Temp, self_nt.Temp, slot, fld.Name, int64(0)))
			ll = _lemit(ll, new_inst(IrOpOpFieldSet{}, new_nt.Temp, src_val.Temp, slot, fld.Name, int64(0)))
			src_len := _lnew_temp(ll)
			_ = src_len
			ll = src_len.L
			ll = _lemit(ll, new_inst(IrOpOpFieldGet{}, src_len.Temp, self_nt.Temp, (slot + int64(1)), "", int64(0)))
			ll = _lemit(ll, new_inst(IrOpOpFieldSet{}, new_nt.Temp, src_len.Temp, (slot + int64(1)), "", int64(0)))
			slot = (slot + int64(2))
		} else {
			src_val := _lnew_temp(ll)
			_ = src_val
			ll = src_val.L
			ll = _lemit(ll, new_inst(IrOpOpFieldGet{}, src_val.Temp, self_nt.Temp, slot, fld.Name, int64(0)))
			ll = _lemit(ll, new_inst(IrOpOpFieldSet{}, new_nt.Temp, src_val.Temp, slot, fld.Name, int64(0)))
			slot = (slot + int64(1))
		}
		fi = (fi + int64(1))
	}
	ll = _lemit(ll, new_inst(IrOpOpRet{}, int64(0), new_nt.Temp, int64(0), "", int64(0)))
	ll = _lfinish_func(ll)
	return ll
}

func _lower_derive_debug(l Lowerer, def StructDef) Lowerer {
	fname := (def.Name + "_debug_str")
	_ = fname
	ll := _lset_func(l, new_ir_func(fname, int64(1), int64(0)))
	_ = ll
	self_nt := _lnew_temp(ll)
	_ = self_nt
	ll = self_nt.L
	ll = _lemit(ll, new_inst(IrOpOpArg{}, self_nt.Temp, int64(0), int64(0), "", int64(0)))
	prefix := (def.Name + "{")
	_ = prefix
	prefix_res := _lower_string_const(ll, prefix)
	_ = prefix_res
	ll = prefix_res.L
	cur_str := prefix_res.Temp
	_ = cur_str
	fi := int64(1)
	_ = fi
	slot := int64(0)
	_ = slot
	for (fi < int64(len(def.Fields))) {
		fld := def.Fields[fi]
		_ = fld
		ft := fld.Type_id
		_ = ft
		if (fi > int64(1)) {
			sep_res := _lower_string_const(ll, ", ")
			_ = sep_res
			ll = sep_res.L
			cat_nt := _lnew_str_temp(ll)
			_ = cat_nt
			ll = cat_nt.L
			ll = _lemit(ll, new_inst(IrOpOpStrConcat{}, cat_nt.Temp, cur_str, sep_res.Temp, "", int64(0)))
			cur_str = cat_nt.Temp
		}
		flabel := (fld.Name + ": ")
		_ = flabel
		flabel_res := _lower_string_const(ll, flabel)
		_ = flabel_res
		ll = flabel_res.L
		cat2 := _lnew_str_temp(ll)
		_ = cat2
		ll = cat2.L
		ll = _lemit(ll, new_inst(IrOpOpStrConcat{}, cat2.Temp, cur_str, flabel_res.Temp, "", int64(0)))
		cur_str = cat2.Temp
		if (ft == int64(12)) {
			fval := _lnew_str_temp(ll)
			_ = fval
			ll = fval.L
			ll = _lemit(ll, new_inst(IrOpOpFieldGet{}, fval.Temp, self_nt.Temp, slot, fld.Name, int64(0)))
			ll = _lemit(ll, new_inst(IrOpOpFieldGet{}, (fval.Temp + int64(1)), self_nt.Temp, (slot + int64(1)), "", int64(0)))
			cat3 := _lnew_str_temp(ll)
			_ = cat3
			ll = cat3.L
			ll = _lemit(ll, new_inst(IrOpOpStrConcat{}, cat3.Temp, cur_str, fval.Temp, "", int64(0)))
			cur_str = cat3.Temp
			slot = (slot + int64(2))
		} else {
			fval := _lnew_temp(ll)
			_ = fval
			ll = fval.L
			ll = _lemit(ll, new_inst(IrOpOpFieldGet{}, fval.Temp, self_nt.Temp, slot, fld.Name, int64(0)))
			fstr := _lnew_str_temp(ll)
			_ = fstr
			ll = fstr.L
			ll = _lemit(ll, new_inst(IrOpOpIntToStr{}, fstr.Temp, fval.Temp, int64(0), "", int64(0)))
			cat3 := _lnew_str_temp(ll)
			_ = cat3
			ll = cat3.L
			ll = _lemit(ll, new_inst(IrOpOpStrConcat{}, cat3.Temp, cur_str, fstr.Temp, "", int64(0)))
			cur_str = cat3.Temp
			slot = (slot + int64(1))
		}
		fi = (fi + int64(1))
	}
	close_res := _lower_string_const(ll, "}")
	_ = close_res
	ll = close_res.L
	final_cat := _lnew_str_temp(ll)
	_ = final_cat
	ll = final_cat.L
	ll = _lemit(ll, new_inst(IrOpOpStrConcat{}, final_cat.Temp, cur_str, close_res.Temp, "", int64(0)))
	ll = _lemit(ll, new_inst(IrOpOpRet{}, int64(0), final_cat.Temp, int64(0), "", int64(0)))
	ll = _lfinish_func(ll)
	return ll
}

func _lower_derive_display(l Lowerer, def StructDef) Lowerer {
	fname := (def.Name + "_display")
	_ = fname
	ll := _lset_func(l, new_ir_func(fname, int64(1), int64(0)))
	_ = ll
	self_nt := _lnew_temp(ll)
	_ = self_nt
	ll = self_nt.L
	ll = _lemit(ll, new_inst(IrOpOpArg{}, self_nt.Temp, int64(0), int64(0), "", int64(0)))
	empty_res := _lower_string_const(ll, "")
	_ = empty_res
	ll = empty_res.L
	cur_str := empty_res.Temp
	_ = cur_str
	fi := int64(1)
	_ = fi
	slot := int64(0)
	_ = slot
	for (fi < int64(len(def.Fields))) {
		fld := def.Fields[fi]
		_ = fld
		ft := fld.Type_id
		_ = ft
		if (fi > int64(1)) {
			sep_res := _lower_string_const(ll, ", ")
			_ = sep_res
			ll = sep_res.L
			cat_nt := _lnew_str_temp(ll)
			_ = cat_nt
			ll = cat_nt.L
			ll = _lemit(ll, new_inst(IrOpOpStrConcat{}, cat_nt.Temp, cur_str, sep_res.Temp, "", int64(0)))
			cur_str = cat_nt.Temp
		}
		if (ft == int64(12)) {
			fval := _lnew_str_temp(ll)
			_ = fval
			ll = fval.L
			ll = _lemit(ll, new_inst(IrOpOpFieldGet{}, fval.Temp, self_nt.Temp, slot, fld.Name, int64(0)))
			ll = _lemit(ll, new_inst(IrOpOpFieldGet{}, (fval.Temp + int64(1)), self_nt.Temp, (slot + int64(1)), "", int64(0)))
			cat3 := _lnew_str_temp(ll)
			_ = cat3
			ll = cat3.L
			ll = _lemit(ll, new_inst(IrOpOpStrConcat{}, cat3.Temp, cur_str, fval.Temp, "", int64(0)))
			cur_str = cat3.Temp
			slot = (slot + int64(2))
		} else {
			fval := _lnew_temp(ll)
			_ = fval
			ll = fval.L
			ll = _lemit(ll, new_inst(IrOpOpFieldGet{}, fval.Temp, self_nt.Temp, slot, fld.Name, int64(0)))
			fstr := _lnew_str_temp(ll)
			_ = fstr
			ll = fstr.L
			ll = _lemit(ll, new_inst(IrOpOpIntToStr{}, fstr.Temp, fval.Temp, int64(0), "", int64(0)))
			cat3 := _lnew_str_temp(ll)
			_ = cat3
			ll = cat3.L
			ll = _lemit(ll, new_inst(IrOpOpStrConcat{}, cat3.Temp, cur_str, fstr.Temp, "", int64(0)))
			cur_str = cat3.Temp
			slot = (slot + int64(1))
		}
		fi = (fi + int64(1))
	}
	ll = _lemit(ll, new_inst(IrOpOpRet{}, int64(0), cur_str, int64(0), "", int64(0)))
	ll = _lfinish_func(ll)
	return ll
}

func _lower_derive_default(l Lowerer, def StructDef) Lowerer {
	fname := (def.Name + "_default")
	_ = fname
	ll := _lset_func(l, new_ir_func(fname, int64(0), int64(0)))
	_ = ll
	alloc_size := (int64(len(def.Fields)) - int64(1))
	_ = alloc_size
	new_nt := _lnew_temp(ll)
	_ = new_nt
	ll = new_nt.L
	ll = _lemit(ll, new_inst(IrOpOpAlloc{}, new_nt.Temp, alloc_size, int64(0), def.Name, int64(0)))
	fi := int64(1)
	_ = fi
	slot := int64(0)
	_ = slot
	for (fi < int64(len(def.Fields))) {
		fld := def.Fields[fi]
		_ = fld
		ft := fld.Type_id
		_ = ft
		if (ft == int64(12)) {
			empty_res := _lower_string_const(ll, "")
			_ = empty_res
			ll = empty_res.L
			ll = _lemit(ll, new_inst(IrOpOpFieldSet{}, new_nt.Temp, empty_res.Temp, slot, fld.Name, int64(0)))
			ll = _lemit(ll, new_inst(IrOpOpFieldSet{}, new_nt.Temp, (empty_res.Temp + int64(1)), (slot + int64(1)), "", int64(0)))
			slot = (slot + int64(2))
		} else {
			zero_nt := _lnew_temp(ll)
			_ = zero_nt
			ll = zero_nt.L
			ll = _lemit(ll, new_inst(IrOpOpConst{}, zero_nt.Temp, int64(0), int64(0), "", int64(0)))
			ll = _lemit(ll, new_inst(IrOpOpFieldSet{}, new_nt.Temp, zero_nt.Temp, slot, fld.Name, int64(0)))
			slot = (slot + int64(1))
		}
		fi = (fi + int64(1))
	}
	ll = _lemit(ll, new_inst(IrOpOpRet{}, int64(0), new_nt.Temp, int64(0), "", int64(0)))
	ll = _lfinish_func(ll)
	return ll
}

func _lower_derive_ord(l Lowerer, def StructDef) Lowerer {
	fname := (def.Name + "_cmp")
	_ = fname
	ll := _lset_func(l, new_ir_func(fname, int64(2), int64(0)))
	_ = ll
	self_nt := _lnew_temp(ll)
	_ = self_nt
	ll = self_nt.L
	ll = _lemit(ll, new_inst(IrOpOpArg{}, self_nt.Temp, int64(0), int64(0), "", int64(0)))
	other_nt := _lnew_temp(ll)
	_ = other_nt
	ll = other_nt.L
	ll = _lemit(ll, new_inst(IrOpOpArg{}, other_nt.Temp, int64(1), int64(0), "", int64(0)))
	fi := int64(1)
	_ = fi
	slot := int64(0)
	_ = slot
	end_lbl := _lnew_temp(ll)
	_ = end_lbl
	ll = end_lbl.L
	result_nt := _lnew_temp(ll)
	_ = result_nt
	ll = result_nt.L
	for (fi < int64(len(def.Fields))) {
		fld := def.Fields[fi]
		_ = fld
		ft := fld.Type_id
		_ = ft
		if (ft == int64(12)) {
			a_str := _lnew_str_temp(ll)
			_ = a_str
			ll = a_str.L
			ll = _lemit(ll, new_inst(IrOpOpFieldGet{}, a_str.Temp, self_nt.Temp, slot, fld.Name, int64(0)))
			ll = _lemit(ll, new_inst(IrOpOpFieldGet{}, (a_str.Temp + int64(1)), self_nt.Temp, (slot + int64(1)), "", int64(0)))
			b_str := _lnew_str_temp(ll)
			_ = b_str
			ll = b_str.L
			ll = _lemit(ll, new_inst(IrOpOpFieldGet{}, b_str.Temp, other_nt.Temp, slot, fld.Name, int64(0)))
			ll = _lemit(ll, new_inst(IrOpOpFieldGet{}, (b_str.Temp + int64(1)), other_nt.Temp, (slot + int64(1)), "", int64(0)))
			eq_nt := _lnew_temp(ll)
			_ = eq_nt
			ll = eq_nt.L
			ll = _lemit(ll, new_inst(IrOpOpStrEq{}, eq_nt.Temp, a_str.Temp, b_str.Temp, "", int64(0)))
			skip_lbl := _lnew_temp(ll)
			_ = skip_lbl
			ll = skip_lbl.L
			ll = _lemit(ll, new_inst(IrOpOpBranchTrue{}, int64(0), skip_lbl.Temp, eq_nt.Temp, "", int64(0)))
			lt_nt := _lnew_temp(ll)
			_ = lt_nt
			ll = lt_nt.L
			ll = _lemit(ll, new_inst(IrOpOpLt{}, lt_nt.Temp, (a_str.Temp + int64(1)), (b_str.Temp + int64(1)), "", int64(0)))
			neg_lbl := _lnew_temp(ll)
			_ = neg_lbl
			ll = neg_lbl.L
			ll = _lemit(ll, new_inst(IrOpOpBranchTrue{}, int64(0), neg_lbl.Temp, lt_nt.Temp, "", int64(0)))
			one_nt := _lnew_temp(ll)
			_ = one_nt
			ll = one_nt.L
			ll = _lemit(ll, new_inst(IrOpOpConst{}, one_nt.Temp, int64(1), int64(0), "", int64(0)))
			ll = _lemit(ll, new_inst(IrOpOpStore{}, result_nt.Temp, one_nt.Temp, int64(0), "", int64(0)))
			ll = _lemit(ll, new_inst(IrOpOpJump{}, int64(0), end_lbl.Temp, int64(0), "", int64(0)))
			ll = _lemit(ll, new_inst(IrOpOpLabel{}, int64(0), neg_lbl.Temp, int64(0), "", int64(0)))
			neg_one := _lnew_temp(ll)
			_ = neg_one
			ll = neg_one.L
			ll = _lemit(ll, new_inst(IrOpOpConst{}, neg_one.Temp, (int64(0) - int64(1)), int64(0), "", int64(0)))
			ll = _lemit(ll, new_inst(IrOpOpStore{}, result_nt.Temp, neg_one.Temp, int64(0), "", int64(0)))
			ll = _lemit(ll, new_inst(IrOpOpJump{}, int64(0), end_lbl.Temp, int64(0), "", int64(0)))
			ll = _lemit(ll, new_inst(IrOpOpLabel{}, int64(0), skip_lbl.Temp, int64(0), "", int64(0)))
			slot = (slot + int64(2))
		} else {
			a_val := _lnew_temp(ll)
			_ = a_val
			ll = a_val.L
			ll = _lemit(ll, new_inst(IrOpOpFieldGet{}, a_val.Temp, self_nt.Temp, slot, fld.Name, int64(0)))
			b_val := _lnew_temp(ll)
			_ = b_val
			ll = b_val.L
			ll = _lemit(ll, new_inst(IrOpOpFieldGet{}, b_val.Temp, other_nt.Temp, slot, fld.Name, int64(0)))
			lt_nt := _lnew_temp(ll)
			_ = lt_nt
			ll = lt_nt.L
			ll = _lemit(ll, new_inst(IrOpOpLt{}, lt_nt.Temp, a_val.Temp, b_val.Temp, "", int64(0)))
			neg_lbl := _lnew_temp(ll)
			_ = neg_lbl
			ll = neg_lbl.L
			skip1_lbl := _lnew_temp(ll)
			_ = skip1_lbl
			ll = skip1_lbl.L
			ll = _lemit(ll, new_inst(IrOpOpBranchTrue{}, int64(0), neg_lbl.Temp, lt_nt.Temp, "", int64(0)))
			gt_nt := _lnew_temp(ll)
			_ = gt_nt
			ll = gt_nt.L
			ll = _lemit(ll, new_inst(IrOpOpGt{}, gt_nt.Temp, a_val.Temp, b_val.Temp, "", int64(0)))
			pos_lbl := _lnew_temp(ll)
			_ = pos_lbl
			ll = pos_lbl.L
			ll = _lemit(ll, new_inst(IrOpOpBranchTrue{}, int64(0), pos_lbl.Temp, gt_nt.Temp, "", int64(0)))
			ll = _lemit(ll, new_inst(IrOpOpJump{}, int64(0), skip1_lbl.Temp, int64(0), "", int64(0)))
			ll = _lemit(ll, new_inst(IrOpOpLabel{}, int64(0), neg_lbl.Temp, int64(0), "", int64(0)))
			neg_one := _lnew_temp(ll)
			_ = neg_one
			ll = neg_one.L
			ll = _lemit(ll, new_inst(IrOpOpConst{}, neg_one.Temp, (int64(0) - int64(1)), int64(0), "", int64(0)))
			ll = _lemit(ll, new_inst(IrOpOpStore{}, result_nt.Temp, neg_one.Temp, int64(0), "", int64(0)))
			ll = _lemit(ll, new_inst(IrOpOpJump{}, int64(0), end_lbl.Temp, int64(0), "", int64(0)))
			ll = _lemit(ll, new_inst(IrOpOpLabel{}, int64(0), pos_lbl.Temp, int64(0), "", int64(0)))
			one_nt := _lnew_temp(ll)
			_ = one_nt
			ll = one_nt.L
			ll = _lemit(ll, new_inst(IrOpOpConst{}, one_nt.Temp, int64(1), int64(0), "", int64(0)))
			ll = _lemit(ll, new_inst(IrOpOpStore{}, result_nt.Temp, one_nt.Temp, int64(0), "", int64(0)))
			ll = _lemit(ll, new_inst(IrOpOpJump{}, int64(0), end_lbl.Temp, int64(0), "", int64(0)))
			ll = _lemit(ll, new_inst(IrOpOpLabel{}, int64(0), skip1_lbl.Temp, int64(0), "", int64(0)))
			slot = (slot + int64(1))
		}
		fi = (fi + int64(1))
	}
	zero_nt := _lnew_temp(ll)
	_ = zero_nt
	ll = zero_nt.L
	ll = _lemit(ll, new_inst(IrOpOpConst{}, zero_nt.Temp, int64(0), int64(0), "", int64(0)))
	ll = _lemit(ll, new_inst(IrOpOpStore{}, result_nt.Temp, zero_nt.Temp, int64(0), "", int64(0)))
	ll = _lemit(ll, new_inst(IrOpOpLabel{}, int64(0), end_lbl.Temp, int64(0), "", int64(0)))
	ll = _lemit(ll, new_inst(IrOpOpRet{}, int64(0), result_nt.Temp, int64(0), "", int64(0)))
	ll = _lfinish_func(ll)
	return ll
}

func _lower_derive_hash(l Lowerer, def StructDef) Lowerer {
	fname := (def.Name + "_hash")
	_ = fname
	ll := _lset_func(l, new_ir_func(fname, int64(1), int64(0)))
	_ = ll
	self_nt := _lnew_temp(ll)
	_ = self_nt
	ll = self_nt.L
	ll = _lemit(ll, new_inst(IrOpOpArg{}, self_nt.Temp, int64(0), int64(0), "", int64(0)))
	hash_nt := _lnew_temp(ll)
	_ = hash_nt
	ll = hash_nt.L
	ll = _lemit(ll, new_inst(IrOpOpConst{}, hash_nt.Temp, int64(0), int64(0), "", int64(0)))
	fi := int64(1)
	_ = fi
	slot := int64(0)
	_ = slot
	for (fi < int64(len(def.Fields))) {
		fld := def.Fields[fi]
		_ = fld
		ft := fld.Type_id
		_ = ft
		if (ft == int64(12)) {
			flen := _lnew_temp(ll)
			_ = flen
			ll = flen.L
			ll = _lemit(ll, new_inst(IrOpOpFieldGet{}, flen.Temp, self_nt.Temp, (slot + int64(1)), "", int64(0)))
			mul_nt := _lnew_temp(ll)
			_ = mul_nt
			ll = mul_nt.L
			c31 := _lnew_temp(ll)
			_ = c31
			ll = c31.L
			ll = _lemit(ll, new_inst(IrOpOpConst{}, c31.Temp, int64(31), int64(0), "", int64(0)))
			ll = _lemit(ll, new_inst(IrOpOpMul{}, mul_nt.Temp, hash_nt.Temp, c31.Temp, "", int64(0)))
			new_hash := _lnew_temp(ll)
			_ = new_hash
			ll = new_hash.L
			ll = _lemit(ll, new_inst(IrOpOpAdd{}, new_hash.Temp, mul_nt.Temp, flen.Temp, "", int64(0)))
			hash_nt = new_hash
			slot = (slot + int64(2))
		} else {
			fval := _lnew_temp(ll)
			_ = fval
			ll = fval.L
			ll = _lemit(ll, new_inst(IrOpOpFieldGet{}, fval.Temp, self_nt.Temp, slot, fld.Name, int64(0)))
			mul_nt := _lnew_temp(ll)
			_ = mul_nt
			ll = mul_nt.L
			c31 := _lnew_temp(ll)
			_ = c31
			ll = c31.L
			ll = _lemit(ll, new_inst(IrOpOpConst{}, c31.Temp, int64(31), int64(0), "", int64(0)))
			ll = _lemit(ll, new_inst(IrOpOpMul{}, mul_nt.Temp, hash_nt.Temp, c31.Temp, "", int64(0)))
			new_hash := _lnew_temp(ll)
			_ = new_hash
			ll = new_hash.L
			ll = _lemit(ll, new_inst(IrOpOpAdd{}, new_hash.Temp, mul_nt.Temp, fval.Temp, "", int64(0)))
			hash_nt = new_hash
			slot = (slot + int64(1))
		}
		fi = (fi + int64(1))
	}
	ll = _lemit(ll, new_inst(IrOpOpRet{}, int64(0), hash_nt.Temp, int64(0), "", int64(0)))
	ll = _lfinish_func(ll)
	return ll
}

func _lower_trampolines(l Lowerer) Lowerer {
	ll := l
	_ = ll
	i := int64(1)
	_ = i
	for (i < int64(len(ll.Registry.Fn_value_refs))) {
		fname := ll.Registry.Fn_value_refs[i]
		_ = fname
		sig := reg_find_fn(ll.Registry, fname)
		_ = sig
		if (sig.Name != "") {
			tramp_name := ("_tramp_" + fname)
			_ = tramp_name
			existing := mod_find_func(ll.Module, tramp_name)
			_ = existing
			if (existing.Name == "") {
				abi_args := int64(0)
				_ = abi_args
				pi := int64(1)
				_ = pi
				for (pi < int64(len(sig.Param_types))) {
					if (sig.Param_types[pi] == int64(12)) {
						abi_args = (abi_args + int64(2))
					} else {
						abi_args = (abi_args + int64(1))
					}
					pi = (pi + int64(1))
				}
				total_args := (int64(1) + abi_args)
				_ = total_args
				ll = _lset_func(ll, new_ir_func(tramp_name, total_args, sig.Return_type))
				ai := int64(0)
				_ = ai
				for (ai < total_args) {
					arg_temp := _lnew_temp(ll)
					_ = arg_temp
					ll = arg_temp.L
					ll = _lemit(ll, new_inst(IrOpOpArg{}, arg_temp.Temp, ai, int64(0), "", int64(0)))
					ai = (ai + int64(1))
				}
				first_forward := int64(1)
				_ = first_forward
				ret_temp := int64(0)
				_ = ret_temp
				if (sig.Return_type == int64(12)) {
					rt := _lnew_str_temp(ll)
					_ = rt
					ll = rt.L
					ret_temp = rt.Temp
					ll = _lemit(ll, new_inst(IrOpOpCall{}, ret_temp, first_forward, abi_args, fname, sig.Return_type))
					ll = _lemit(ll, new_inst(IrOpOpRet{}, int64(0), ret_temp, int64(0), "", int64(0)))
				} else if ((sig.Return_type == int64(0)) || (sig.Return_type == int64(14))) {
					ll = _lemit(ll, new_inst(IrOpOpCall{}, int64(0), first_forward, abi_args, fname, sig.Return_type))
					ll = _lemit(ll, new_inst(IrOpOpRetVoid{}, int64(0), int64(0), int64(0), "", int64(0)))
				} else {
					rt := _lnew_temp(ll)
					_ = rt
					ll = rt.L
					ret_temp = rt.Temp
					ll = _lemit(ll, new_inst(IrOpOpCall{}, ret_temp, first_forward, abi_args, fname, sig.Return_type))
					ll = _lemit(ll, new_inst(IrOpOpRet{}, int64(0), ret_temp, int64(0), "", int64(0)))
				}
				ll = _lfinish_func(ll)
			}
		}
		i = (i + int64(1))
	}
	bi := int64(1)
	_ = bi
	for (bi < int64(len(ll.Registry.Bound_fn_value_refs))) {
		bfname := ll.Registry.Bound_fn_value_refs[bi]
		_ = bfname
		bsig := reg_find_fn(ll.Registry, bfname)
		_ = bsig
		if (bsig.Name != "") {
			btramp := ("_bound_tramp_" + bfname)
			_ = btramp
			bexist := mod_find_func(ll.Module, btramp)
			_ = bexist
			if (bexist.Name == "") {
				babi := int64(0)
				_ = babi
				bpi := int64(1)
				_ = bpi
				for (bpi < int64(len(bsig.Param_types))) {
					if (bsig.Param_types[bpi] == int64(12)) {
						babi = (babi + int64(2))
					} else {
						babi = (babi + int64(1))
					}
					bpi = (bpi + int64(1))
				}
				extra_args := (babi - int64(1))
				_ = extra_args
				total_in := (int64(1) + extra_args)
				_ = total_in
				ll = _lset_func(ll, new_ir_func(btramp, total_in, bsig.Return_type))
				env_t := _lnew_temp(ll)
				_ = env_t
				ll = env_t.L
				ll = _lemit(ll, new_inst(IrOpOpArg{}, env_t.Temp, int64(0), int64(0), "", int64(0)))
				self_t := _lnew_temp(ll)
				_ = self_t
				ll = self_t.L
				ll = _lemit(ll, new_inst(IrOpOpFieldGet{}, self_t.Temp, env_t.Temp, int64(0), "_self", int64(0)))
				eai := int64(0)
				_ = eai
				for (eai < extra_args) {
					ea_t := _lnew_temp(ll)
					_ = ea_t
					ll = ea_t.L
					ll = _lemit(ll, new_inst(IrOpOpArg{}, ea_t.Temp, (eai + int64(1)), int64(0), "", int64(0)))
					eai = (eai + int64(1))
				}
				if (bsig.Return_type == int64(12)) {
					brt := _lnew_str_temp(ll)
					_ = brt
					ll = brt.L
					ll = _lemit(ll, new_inst(IrOpOpCall{}, brt.Temp, self_t.Temp, babi, bfname, bsig.Return_type))
					ll = _lemit(ll, new_inst(IrOpOpRet{}, int64(0), brt.Temp, int64(0), "", int64(0)))
				} else if ((bsig.Return_type == int64(0)) || (bsig.Return_type == int64(14))) {
					ll = _lemit(ll, new_inst(IrOpOpCall{}, int64(0), self_t.Temp, babi, bfname, bsig.Return_type))
					ll = _lemit(ll, new_inst(IrOpOpRetVoid{}, int64(0), int64(0), int64(0), "", int64(0)))
				} else {
					brt := _lnew_temp(ll)
					_ = brt
					ll = brt.L
					ll = _lemit(ll, new_inst(IrOpOpCall{}, brt.Temp, self_t.Temp, babi, bfname, bsig.Return_type))
					ll = _lemit(ll, new_inst(IrOpOpRet{}, int64(0), brt.Temp, int64(0), "", int64(0)))
				}
				ll = _lfinish_func(ll)
			}
		}
		bi = (bi + int64(1))
	}
	return ll
}

func lower(tokens []Token, index DeclIndex, pool NodePool, store TypeStore, reg TypeRegistry, treg TraitRegistry, table SymbolTable, file string) IrModule {
	l := _new_lowerer(tokens, pool, store, reg, treg, table, file)
	_ = l
	l = _scan_consts(l, index)
	l = _lower_bodies(l, index)
	l = _lower_trampolines(l)
	l = _lower_derives(l)
	l = _lower_mono_pass(l)
	return l.Module
}

func lower_for_tests(tokens []Token, index DeclIndex, pool NodePool, store TypeStore, reg TypeRegistry, treg TraitRegistry, table SymbolTable, file string, parallel bool) IrModule {
	l := _new_lowerer(tokens, pool, store, reg, treg, table, file)
	_ = l
	l = _scan_consts(l, index)
	l = _lower_bodies_skip_entry(l, index)
	if parallel {
		ti := int64(0)
		_ = ti
		for (ti < l.Module.Test_count) {
			test_fn := ("_test_" + i2s(ti))
			_ = test_fn
			test_sig := new_fn_sig(test_fn, []string{""}, []int64{int64(0)}, int64(0), int64(0))
			_ = test_sig
			l = _lset_reg(l, reg_add_fn(l.Registry, test_sig))
			l = _lset_reg(l, reg_add_fn_value_ref(l.Registry, test_fn))
			ti = (ti + int64(1))
		}
	}
	l = _lower_trampolines(l)
	l = _lower_derives(l)
	l = _lower_mono_pass(l)
	l = _generate_test_runner_p(l, parallel)
	return l.Module
}

func _generate_test_runner_p(l Lowerer, parallel bool) Lowerer {
	ll := l
	_ = ll
	test_count := ll.Module.Test_count
	_ = test_count
	if (test_count == int64(0)) {
		ll = _lset_func(ll, new_ir_func("_main", int64(0), int64(0)))
		no_tests_idx := mod_find_string(ll.Module, "no tests found")
		_ = no_tests_idx
		if (no_tests_idx == int64(0)) {
			new_mod := mod_add_string(ll.Module, "no tests found")
			_ = new_mod
			no_tests_idx = (int64(len(new_mod.String_constants)) - int64(1))
			ll = _lset_mod(ll, new_mod)
		}
		nt := _lnew_str_temp(ll)
		_ = nt
		ll = nt.L
		ll = _lemit(ll, new_inst(IrOpOpConstStr{}, nt.Temp, no_tests_idx, int64(14), "", int64(0)))
		ll = _lemit(ll, new_inst(IrOpOpCall{}, int64(0), nt.Temp, int64(2), "_aria_println_str", int64(0)))
		exit_nt := _lnew_temp(ll)
		_ = exit_nt
		ll = exit_nt.L
		ll = _lemit(ll, new_inst(IrOpOpConst{}, exit_nt.Temp, int64(0), int64(0), "", int64(0)))
		ll = _lemit(ll, new_inst(IrOpOpCall{}, int64(0), exit_nt.Temp, int64(1), "_aria_exit", int64(0)))
		ll = _lfinish_func(ll)
		new_mod2 := mod_set_entry(ll.Module, "_main")
		_ = new_mod2
		ll = Lowerer{Tokens: ll.Tokens, Pos: ll.Pos, Store: ll.Store, Registry: ll.Registry, Treg: ll.Treg, Pool: ll.Pool, Table: ll.Table, File: ll.File, Module: new_mod2, Current_func: ll.Current_func, Temp_counter: ll.Temp_counter, Label_counter: ll.Label_counter, Env_names: ll.Env_names, Env_slots: ll.Env_slots, Loop_start: ll.Loop_start, Loop_end: ll.Loop_end, Const_names: ll.Const_names, Const_vals: ll.Const_vals, Const_str_names: ll.Const_str_names, Const_str_vals: ll.Const_str_vals, Temp_types: ll.Temp_types, Fn_error_type: ll.Fn_error_type, Defer_starts: ll.Defer_starts, Defer_ends: ll.Defer_ends}
		return ll
	}
	ll = _lset_func(ll, new_ir_func("_main", int64(0), int64(0)))
	if (parallel && (test_count > int64(0))) {
		first_handle := _lnew_temp(ll)
		_ = first_handle
		ll = first_handle.L
		hi := int64(1)
		_ = hi
		for (hi < test_count) {
			pad := _lnew_temp(ll)
			_ = pad
			ll = pad.L
			hi = (hi + int64(1))
		}
		si := int64(0)
		_ = si
		for (si < test_count) {
			test_fn_name := ("_test_" + i2s(si))
			_ = test_fn_name
			tramp_name := ("_tramp_" + test_fn_name)
			_ = tramp_name
			fn_ref := _lnew_temp(ll)
			_ = fn_ref
			ll = fn_ref.L
			ll = _lemit(ll, new_inst(IrOpOpFnRef{}, fn_ref.Temp, int64(0), int64(0), tramp_name, int64(0)))
			null_env := _lnew_temp(ll)
			_ = null_env
			ll = null_env.L
			ll = _lemit(ll, new_inst(IrOpOpConst{}, null_env.Temp, int64(0), int64(0), "", int64(0)))
			ll = _lemit(ll, new_inst(IrOpOpCall{}, (first_handle.Temp + si), fn_ref.Temp, int64(2), "_aria_spawn", int64(1)))
			si = (si + int64(1))
		}
		ai := int64(0)
		_ = ai
		for (ai < test_count) {
			await_nt := _lnew_temp(ll)
			_ = await_nt
			ll = await_nt.L
			ll = _lemit(ll, new_inst(IrOpOpCall{}, await_nt.Temp, (first_handle.Temp + ai), int64(1), "_aria_task_await", int64(1)))
			ai = (ai + int64(1))
		}
		done_idx := mod_find_string(ll.Module, "all tests passed")
		_ = done_idx
		if (done_idx == int64(0)) {
			new_mod := mod_add_string(ll.Module, "all tests passed")
			_ = new_mod
			done_idx = (int64(len(new_mod.String_constants)) - int64(1))
			ll = _lset_mod(ll, new_mod)
		}
		done_nt := _lnew_str_temp(ll)
		_ = done_nt
		ll = done_nt.L
		ll = _lemit(ll, new_inst(IrOpOpConstStr{}, done_nt.Temp, done_idx, int64(16), "", int64(0)))
		ll = _lemit(ll, new_inst(IrOpOpCall{}, int64(0), done_nt.Temp, int64(2), "_aria_println_str", int64(0)))
		ll = _lfinish_func(ll)
		new_mod_e := mod_set_entry(ll.Module, "_main")
		_ = new_mod_e
		ll = _lset_mod(ll, new_mod_e)
		return ll
	}
	ti := int64(0)
	_ = ti
	for (ti < test_count) {
		test_fn_name := ("_test_" + i2s(ti))
		_ = test_fn_name
		test_label_idx := (ti + int64(1))
		_ = test_label_idx
		label_str_idx := int64(0)
		_ = label_str_idx
		if (test_label_idx < int64(len(ll.Module.Test_names))) {
			label_text := ll.Module.Test_names[test_label_idx]
			_ = label_text
			label_str_idx = mod_find_string(ll.Module, label_text)
		}
		before_fn := mod_find_func(ll.Module, "_before")
		_ = before_fn
		if (before_fn.Name != "") {
			ll = _lemit(ll, new_inst(IrOpOpCall{}, int64(0), int64(0), int64(0), "_before", int64(0)))
		}
		ll = _lemit(ll, new_inst(IrOpOpCall{}, int64(0), int64(0), int64(0), test_fn_name, int64(0)))
		pass_idx := mod_find_string(ll.Module, "PASS: ")
		_ = pass_idx
		if (pass_idx == int64(0)) {
			new_mod := mod_add_string(ll.Module, "PASS: ")
			_ = new_mod
			pass_idx = (int64(len(new_mod.String_constants)) - int64(1))
			ll = _lset_mod(ll, new_mod)
		}
		pass_nt := _lnew_str_temp(ll)
		_ = pass_nt
		ll = pass_nt.L
		ll = _lemit(ll, new_inst(IrOpOpConstStr{}, pass_nt.Temp, pass_idx, int64(6), "", int64(0)))
		if (label_str_idx > int64(0)) {
			label_nt := _lnew_str_temp(ll)
			_ = label_nt
			ll = label_nt.L
			label_text := ll.Module.Test_names[test_label_idx]
			_ = label_text
			ll = _lemit(ll, new_inst(IrOpOpConstStr{}, label_nt.Temp, label_str_idx, int64(len(label_text)), "", int64(0)))
			concat_nt := _lnew_str_temp(ll)
			_ = concat_nt
			ll = concat_nt.L
			ll = _lemit(ll, new_inst(IrOpOpStrConcat{}, concat_nt.Temp, pass_nt.Temp, label_nt.Temp, "", int64(0)))
			ll = _lemit(ll, new_inst(IrOpOpCall{}, int64(0), concat_nt.Temp, int64(2), "_aria_println_str", int64(0)))
		} else {
			ll = _lemit(ll, new_inst(IrOpOpCall{}, int64(0), pass_nt.Temp, int64(2), "_aria_println_str", int64(0)))
		}
		after_fn := mod_find_func(ll.Module, "_after")
		_ = after_fn
		if (after_fn.Name != "") {
			ll = _lemit(ll, new_inst(IrOpOpCall{}, int64(0), int64(0), int64(0), "_after", int64(0)))
		}
		ti = (ti + int64(1))
	}
	done_idx := mod_find_string(ll.Module, "all tests passed")
	_ = done_idx
	if (done_idx == int64(0)) {
		new_mod := mod_add_string(ll.Module, "all tests passed")
		_ = new_mod
		done_idx = (int64(len(new_mod.String_constants)) - int64(1))
		ll = Lowerer{Tokens: ll.Tokens, Pos: ll.Pos, Store: ll.Store, Registry: ll.Registry, Treg: ll.Treg, Pool: ll.Pool, Table: ll.Table, File: ll.File, Module: new_mod, Current_func: ll.Current_func, Temp_counter: ll.Temp_counter, Label_counter: ll.Label_counter, Env_names: ll.Env_names, Env_slots: ll.Env_slots, Loop_start: ll.Loop_start, Loop_end: ll.Loop_end, Const_names: ll.Const_names, Const_vals: ll.Const_vals, Const_str_names: ll.Const_str_names, Const_str_vals: ll.Const_str_vals, Temp_types: ll.Temp_types, Fn_error_type: ll.Fn_error_type, Defer_starts: ll.Defer_starts, Defer_ends: ll.Defer_ends}
	}
	done_nt := _lnew_str_temp(ll)
	_ = done_nt
	ll = done_nt.L
	ll = _lemit(ll, new_inst(IrOpOpConstStr{}, done_nt.Temp, done_idx, int64(16), "", int64(0)))
	ll = _lemit(ll, new_inst(IrOpOpCall{}, int64(0), done_nt.Temp, int64(2), "_aria_println_str", int64(0)))
	exit_nt := _lnew_temp(ll)
	_ = exit_nt
	ll = exit_nt.L
	ll = _lemit(ll, new_inst(IrOpOpConst{}, exit_nt.Temp, int64(0), int64(0), "", int64(0)))
	ll = _lemit(ll, new_inst(IrOpOpCall{}, int64(0), exit_nt.Temp, int64(1), "_aria_exit", int64(0)))
	ll = _lfinish_func(ll)
	new_mod := mod_set_entry(ll.Module, "_main")
	_ = new_mod
	return _lset_mod(ll, new_mod)
}

func _lower_mono_pass(l Lowerer) Lowerer {
	ll := l
	_ = ll
	mi := int64(1)
	_ = mi
	for (mi < int64(len(ll.Registry.Mono_specs))) {
		spec := ll.Registry.Mono_specs[mi]
		_ = spec
		generic_sig := reg_find_fn(ll.Registry, spec.Generic_name)
		_ = generic_sig
		if ((generic_sig.Name != "") && fn_is_generic(generic_sig)) {
			ll = _lset_pos(ll, generic_sig.Token_start)
			ll = _lower_mono_fn(ll, spec)
		}
		mi = (mi + int64(1))
	}
	return ll
}

func _lower_mono_fn(l Lowerer, spec MonoSpec) Lowerer {
	ll := _ladv(l)
	_ = ll
	if (_lk(ll) != "IDENT") {
		return ll
	}
	ll = _ladv(ll)
	if (_lk(ll) == "[") {
		d := int64(1)
		_ = d
		ll = _ladv(ll)
		for ((d > int64(0)) && (_lk(ll) != "EOF")) {
			if (_lk(ll) == "[") {
				d = (d + int64(1))
			}
			if (_lk(ll) == "]") {
				d = (d - int64(1))
			}
			if (d > int64(0)) {
				ll = _ladv(ll)
			}
		}
		if (_lk(ll) == "]") {
			ll = _ladv(ll)
		}
	}
	sig := reg_find_fn(ll.Registry, spec.Specialized_name)
	_ = sig
	if (sig.Name == "") {
		fmt.Println(((((("warning: unresolved generic specialization '" + spec.Specialized_name) + "' at ") + ll.File) + ":") + i2s(_lcur(ll).Line)))
		return ll
	}
	actual_param_count := int64(0)
	_ = actual_param_count
	pi := int64(1)
	_ = pi
	for (pi < int64(len(sig.Param_names))) {
		if ((pi <= (int64(len(sig.Param_types)) - int64(1))) && (sig.Param_types[pi] == int64(12))) {
			actual_param_count = (actual_param_count + int64(2))
		} else {
			actual_param_count = (actual_param_count + int64(1))
		}
		pi = (pi + int64(1))
	}
	ll = _lset_func(ll, new_ir_func(spec.Specialized_name, actual_param_count, sig.Return_type))
	pi = int64(1)
	for (pi < int64(len(sig.Param_names))) {
		if ((pi <= (int64(len(sig.Param_types)) - int64(1))) && (sig.Param_types[pi] == int64(12))) {
			nt := _lnew_str_temp(ll)
			_ = nt
			ll = nt.L
			ll = _lenv_add(ll, sig.Param_names[pi], nt.Temp)
			ll = _lset_type(ll, nt.Temp, int64(12))
		} else {
			nt := _lnew_temp(ll)
			_ = nt
			ll = nt.L
			ll = _lenv_add(ll, sig.Param_names[pi], nt.Temp)
			if ((pi <= (int64(len(sig.Param_types)) - int64(1))) && (sig.Param_types[pi] > int64(0))) {
				pt := sig.Param_types[pi]
				_ = pt
				if _is_str_array_type(ll.Store, pt) {
					ll = _lset_type(ll, nt.Temp, int64(13))
				} else if _is_trait_object_type(ll.Store, pt) {
					ll = _lset_type(ll, nt.Temp, int64(500))
				} else {
					ll = _lset_type(ll, nt.Temp, pt)
				}
			}
		}
		pi = (pi + int64(1))
	}
	if (_lk(ll) == "(") {
		d := int64(1)
		_ = d
		ll = _ladv(ll)
		for ((d > int64(0)) && (_lk(ll) != "EOF")) {
			if (_lk(ll) == "(") {
				d = (d + int64(1))
			}
			if (_lk(ll) == ")") {
				d = (d - int64(1))
			}
			if (d > int64(0)) {
				ll = _ladv(ll)
			}
		}
		if (_lk(ll) == ")") {
			ll = _ladv(ll)
		}
	}
	if (_lk(ll) == "->") {
		ll = _ladv(ll)
		for (((((_lk(ll) == "IDENT") || (_lk(ll) == "[")) || (_lk(ll) == "]")) || (_lk(ll) == "!")) || (_lk(ll) == "?")) {
			ll = _ladv(ll)
		}
	}
	if (_lk(ll) == "!") {
		ll = _ladv(ll)
		for (_lk(ll) == "IDENT") {
			ll = _ladv(ll)
		}
	}
	if (_lk(ll) == "where") {
		ll = _ladv(ll)
		lb2 := "{"
		_ = lb2
		for ((((_lk(ll) != "=") && (_lk(ll) != lb2)) && (_lk(ll) != "EOF")) && (_lk(ll) != "NEWLINE")) {
			ll = _ladv(ll)
		}
	}
	lb := "{"
	_ = lb
	if (_lk(ll) == "=") {
		ll = _ladv(ll)
		ll = _lskip_nl(ll)
		result := _lower_expr(ll)
		_ = result
		ll = result.L
		ll = _emit_defers(ll)
		ret_type := _lget_type(ll, result.Temp)
		_ = ret_type
		ll = _lemit(ll, new_inst(IrOpOpRet{}, int64(0), result.Temp, int64(0), "", ret_type))
	} else if (_lk(ll) == lb) {
		ll = _lower_fn_body(ll, sig.Return_type)
	} else {
		ll = _lskip_nl(ll)
		if (_lk(ll) == lb) {
			ll = _lower_fn_body(ll, sig.Return_type)
		}
	}
	ll = _lfinish_func(ll)
	return ll
}

