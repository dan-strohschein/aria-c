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

type LlvmGen struct {
	Parts []string
	Label_counter int64
	Last_was_term bool
}

type LT_label struct {
	G LlvmGen
	Id int64
}

func _lb() string {
	return "{"
}

func _rb() string {
	return "}"
}

func _nl() string {
	return "\n"
}

func _join(parts []string) string {
	chunks := []string{""}
	_ = chunks
	i := int64(1)
	_ = i
	for (i < int64(len(parts))) {
		chunk := ""
		_ = chunk
		j := int64(0)
		_ = j
		for ((j < int64(64)) && (((i + j)) < int64(len(parts)))) {
			chunk = (chunk + parts[(i + j)])
			j = (j + int64(1))
		}
		chunks = append(chunks, chunk)
		i = (i + int64(64))
	}
	result := ""
	_ = result
	ci := int64(1)
	_ = ci
	for (ci < int64(len(chunks))) {
		result = (result + chunks[ci])
		ci = (ci + int64(1))
	}
	return result
}

func _llvm_escape_str(s string) string {
	result := ""
	_ = result
	i := int64(0)
	_ = i
	for (i < int64(len(s))) {
		ch := string(s[i])
		_ = ch
		if (ch == "\\") {
			result = (result + "\\5C")
		} else if (ch == "\"") {
			result = (result + "\\22")
		} else if (ch == "\n") {
			result = (result + "\\0A")
		} else if (ch == "\r") {
			result = (result + "\\0D")
		} else if (ch == "\t") {
			result = (result + "\\09")
		} else if (ch == "\x00") {
			result = (result + "\\00")
		} else {
			result = (result + ch)
		}
		i = (i + int64(1))
	}
	return result
}

func _get_temp_type(f IrFunc, temp int64) int64 {
	idx := (temp + int64(1))
	_ = idx
	if ((idx >= int64(0)) && (idx < int64(len(f.Temp_types)))) {
		return f.Temp_types[idx]
	}
	return int64(0)
}

func _is_definitely_scalar_op(op string) bool {
	return (((((((((((((((((((((((((((((((((((((op == "Const") || (op == "Add")) || (op == "Sub")) || (op == "Mul")) || (op == "Div")) || (op == "Mod")) || (op == "Neg")) || (op == "Eq")) || (op == "Neq")) || (op == "Lt")) || (op == "Gt")) || (op == "Lte")) || (op == "Gte")) || (op == "And")) || (op == "Or")) || (op == "Not")) || (op == "BitAnd")) || (op == "BitOr")) || (op == "BitXor")) || (op == "Shl")) || (op == "Shr")) || (op == "ArrayLen")) || (op == "StrLen")) || (op == "StrEq")) || (op == "StrCmp")) || (op == "StrContains")) || (op == "StrStartsWith")) || (op == "StrEndsWith")) || (op == "StrIndexOf")) || (op == "MapContains")) || (op == "MapLen")) || (op == "SetContains")) || (op == "SetLen")) || (op == "BranchTrue")) || (op == "BranchFalse")) || (op == "Label")) || (op == "Jump"))
}

func _is_gc_root(f IrFunc, temp int64) bool {
	tt := _get_temp_type(f, temp)
	_ = tt
	if (tt == int64(1)) {
		return false
	}
	if (tt == int64(12)) {
		return false
	}
	if (tt >= int64(100)) {
		return true
	}
	could_be_pointer := false
	_ = could_be_pointer
	found_any := false
	_ = found_any
	ii := int64(1)
	_ = ii
	for (ii < int64(len(f.Insts))) {
		inst := f.Insts[ii]
		_ = inst
		if (inst.Dest == temp) {
			found_any = true
			op := ir_op_name(inst.Op)
			_ = op
			if ((_is_definitely_scalar_op(op) == false) && (op != "ConstStr")) {
				could_be_pointer = true
			}
		}
		ii = (ii + int64(1))
	}
	if found_any {
		return could_be_pointer
	}
	if (temp < f.Param_count) {
		return true
	}
	return false
}

func _gc_root_count_for_func(f IrFunc) int64 {
	count := int64(0)
	_ = count
	i := int64(0)
	_ = i
	for (i < f.Local_count) {
		if _is_gc_root(f, i) {
			count = (count + int64(1))
		}
		i = (i + int64(1))
	}
	return count
}

func _gc_frame_update(g LlvmGen, f IrFunc, temp int64, val string, idx int64) LlvmGen {
	if (_is_gc_root(f, temp) == false) {
		return g
	}
	gc_frame_slots := (_gc_root_count_for_func(f) + int64(2))
	_ = gc_frame_slots
	slot := _gc_slot_for_temp(f, temp)
	_ = slot
	if ((slot < int64(0)) || (gc_frame_slots <= int64(2))) {
		return g
	}
	gg := g
	_ = gg
	gep := ((("%gc_u_" + i2s(idx)) + "_") + i2s(temp))
	_ = gep
	gg = _lg_emit(gg, ((((((gep + " = getelementptr [") + i2s(gc_frame_slots)) + " x i64], [") + i2s(gc_frame_slots)) + " x i64]* %gc_frame, i64 0, i64 ") + i2s((slot + int64(2)))))
	gg = _lg_emit(gg, ((("store i64 " + val) + ", i64* ") + gep))
	return gg
}

func _gc_slot_for_temp(f IrFunc, temp int64) int64 {
	slot := int64(0)
	_ = slot
	i := int64(0)
	_ = i
	for (i < f.Local_count) {
		if _is_gc_root(f, i) {
			if (i == temp) {
				return slot
			}
			slot = (slot + int64(1))
		}
		i = (i + int64(1))
	}
	return (int64(0) - int64(1))
}

func _lg_new() LlvmGen {
	return LlvmGen{Parts: []string{""}, Label_counter: int64(0), Last_was_term: false}
}

func _lg_emit(g LlvmGen, line string) LlvmGen {
	return LlvmGen{Parts: append(g.Parts, (("  " + line) + "\n")), Label_counter: g.Label_counter, Last_was_term: false}
}

func _lg_emit_term(g LlvmGen, line string) LlvmGen {
	return LlvmGen{Parts: append(g.Parts, (("  " + line) + "\n")), Label_counter: g.Label_counter, Last_was_term: true}
}

func _lg_emit_label(g LlvmGen, label string) LlvmGen {
	return LlvmGen{Parts: append(g.Parts, (label + ":\n")), Label_counter: g.Label_counter, Last_was_term: false}
}

func _lg_fresh_label(g LlvmGen) LT_label {
	id := (g.Label_counter + int64(1))
	_ = id
	return LT_label{G: LlvmGen{Parts: g.Parts, Label_counter: id, Last_was_term: g.Last_was_term}, Id: id}
}

func _t(n int64) string {
	if (n < int64(0)) {
		return "%t0"
	}
	return ("%t" + i2s(n))
}

func _v(n int64) string {
	return ("%v" + i2s(n))
}

func _lbl(n int64) string {
	return ("L" + i2s(n))
}

func _fall(n int64) string {
	return ("Lfall" + i2s(n))
}

func _struct_i64_i64() string {
	return ((_lb() + " i64, i64 ") + _rb())
}

func _collect_called_names(m IrModule) []string {
	names := []string{""}
	_ = names
	fi := int64(1)
	_ = fi
	for (fi < int64(len(m.Funcs))) {
		f := m.Funcs[fi]
		_ = f
		ii := int64(1)
		_ = ii
		for (ii < int64(len(f.Insts))) {
			if (ir_op_name(f.Insts[ii].Op) == "Call") {
				cname := f.Insts[ii].S1
				_ = cname
				if (_name_in_list(cname, names) == false) {
					names = append(names, cname)
				}
			}
			ii = (ii + int64(1))
		}
		fi = (fi + int64(1))
	}
	return names
}

func _name_in_list(name string, list []string) bool {
	i := int64(1)
	_ = i
	for (i < int64(len(list))) {
		if (list[i] == name) {
			return true
		}
		i = (i + int64(1))
	}
	return false
}

func _target_triple(target string) string {
	if (target == "darwin-arm64") {
		return "arm64-apple-macosx14.0.0"
	}
	if (target == "darwin-amd64") {
		return "x86_64-apple-macosx14.0.0"
	}
	if (target == "linux-amd64") {
		return "x86_64-unknown-linux-gnu"
	}
	if (target == "linux-arm64") {
		return "aarch64-unknown-linux-gnu"
	}
	if (target == "windows-amd64") {
		return "x86_64-pc-windows-msvc"
	}
	if (target == "wasm32") {
		return "wasm32-unknown-unknown"
	}
	return "arm64-apple-macosx14.0.0"
}

func _target_datalayout(target string) string {
	if ((target == "darwin-arm64") || (target == "native")) {
		return "e-m:o-i64:64-i128:128-n32:64-S128"
	}
	if (target == "darwin-amd64") {
		return "e-m:o-p270:32:32-p271:32:32-p272:64:64-i64:64-f80:128-n8:16:32:64-S128"
	}
	if (target == "linux-amd64") {
		return "e-m:e-p270:32:32-p271:32:32-p272:64:64-i64:64-f80:128-n8:16:32:64-S128"
	}
	if (target == "linux-arm64") {
		return "e-m:e-i8:8:32-i16:16:32-i64:64-i128:128-n32:64-S128"
	}
	if (target == "windows-amd64") {
		return "e-m:w-p270:32:32-p271:32:32-p272:64:64-i64:64-f80:128-n8:16:32:64-S128"
	}
	return "e-m:o-i64:64-i128:128-n32:64-S128"
}

func generate_llvm_ir(m IrModule) string {
	return _generate_llvm_ir_target(m, "native")
}

func generate_llvm_ir_to_file(m IrModule, path string) {
	_generate_llvm_ir_to_file_target(m, path, "native")
	_void := int64(0)
	_ = _void
}

func _generate_llvm_ir_target(m IrModule, target string) string {
	nl := _nl()
	_ = nl
	si64 := _struct_i64_i64()
	_ = si64
	p := []string{""}
	_ = p
	triple := _target_triple(target)
	_ = triple
	layout := _target_datalayout(target)
	_ = layout
	p = append(p, ((("target datalayout = \"" + layout) + "\"") + nl))
	p = append(p, (((("target triple = \"" + triple) + "\"") + nl) + nl))
	si := int64(1)
	_ = si
	for (si < int64(len(m.String_constants))) {
		s := m.String_constants[si]
		_ = s
		escaped := _llvm_escape_str(s)
		_ = escaped
		slen := int64(len(s))
		_ = slen
		p = append(p, ((((((("@.str." + i2s(si)) + " = private unnamed_addr constant [") + i2s((slen + int64(1)))) + " x i8] c\"") + escaped) + "\\00\"") + nl))
		si = (si + int64(1))
	}
	p = append(p, nl)
	p = append(p, ("; Runtime declarations" + nl))
	p = append(p, ("declare void @_aria_exit(i64)" + nl))
	p = append(p, ("declare i64 @_aria_write(i64, i8*, i64)" + nl))
	p = append(p, ("declare void @_aria_println_str(i8*, i64)" + nl))
	p = append(p, ("declare void @_aria_print_str(i8*, i64)" + nl))
	p = append(p, ("declare void @_aria_eprintln_str(i8*, i64)" + nl))
	p = append(p, ("declare i8* @_aria_alloc(i64)" + nl))
	p = append(p, ("declare void @_aria_gc_frame_push(i64, i64)" + nl))
	p = append(p, ("declare void @_aria_gc_frame_pop()" + nl))
	p = append(p, ("declare void @_aria_memcpy(i8*, i8*, i64)" + nl))
	p = append(p, ((("declare " + si64) + " @_aria_read_file(i8*, i64)") + nl))
	p = append(p, ("declare void @_aria_write_file(i8*, i64, i8*, i64)" + nl))
	p = append(p, ("declare void @_aria_append_file(i8*, i64, i8*, i64)" + nl))
	p = append(p, ("declare void @_aria_write_binary_file(i8*, i64, i64*, i64)" + nl))
	p = append(p, ((("declare " + si64) + " @_aria_int_to_str(i64)") + nl))
	p = append(p, ((("declare " + si64) + " @_aria_str_concat(i8*, i64, i8*, i64)") + nl))
	p = append(p, ("declare i64 @_aria_str_eq(i8*, i64, i8*, i64)" + nl))
	p = append(p, ("declare i64 @_aria_str_cmp(i8*, i64, i8*, i64)" + nl))
	p = append(p, ((("declare " + si64) + " @_aria_str_charAt(i8*, i64, i64)") + nl))
	p = append(p, ((("declare " + si64) + " @_aria_str_substring(i8*, i64, i64, i64)") + nl))
	p = append(p, ("declare i64 @_aria_str_contains(i8*, i64, i8*, i64)" + nl))
	p = append(p, ("declare i64 @_aria_str_startsWith(i8*, i64, i8*, i64)" + nl))
	p = append(p, ("declare i64 @_aria_str_endsWith(i8*, i64, i8*, i64)" + nl))
	p = append(p, ("declare i64 @_aria_str_indexOf(i8*, i64, i8*, i64)" + nl))
	p = append(p, ((("declare " + si64) + " @_aria_str_trim(i8*, i64)") + nl))
	p = append(p, ((("declare " + si64) + " @_aria_str_replace(i8*, i64, i8*, i64, i8*, i64)") + nl))
	p = append(p, ((("declare " + si64) + " @_aria_str_toLower(i8*, i64)") + nl))
	p = append(p, ((("declare " + si64) + " @_aria_str_toUpper(i8*, i64)") + nl))
	p = append(p, ("declare i64 @_aria_str_split(i8*, i64, i8*, i64)" + nl))
	p = append(p, ("declare void @_aria_map_set_str(i64, i8*, i64, i64, i64)" + nl))
	p = append(p, ((("declare " + si64) + " @_aria_map_get_str(i64, i8*, i64)") + nl))
	p = append(p, ("declare i64 @_aria_array_new(i64)" + nl))
	p = append(p, ("declare i64 @_aria_array_len(i64)" + nl))
	p = append(p, ("declare i64 @_aria_array_get(i64, i64)" + nl))
	p = append(p, ("declare void @_aria_array_set(i64, i64, i64)" + nl))
	p = append(p, ("declare i64 @_aria_array_append(i64, i64)" + nl))
	p = append(p, ("declare i64 @_aria_array_slice(i64, i64)" + nl))
	p = append(p, ("declare i64 @_aria_list_dir(i8*, i64)" + nl))
	p = append(p, ("declare i64 @_aria_is_dir(i8*, i64)" + nl))
	p = append(p, ("declare i64 @_aria_args_get()" + nl))
	p = append(p, ("declare i64 @_aria_exec(i8*, i64)" + nl))
	p = append(p, ((("declare " + si64) + " @_aria_str_to_int(i8*, i64)") + nl))
	p = append(p, ("declare i64 @_aria_str_to_float(i8*, i64)" + nl))
	p = append(p, ("declare i64 @_aria_map_new(i64)" + nl))
	p = append(p, ("declare void @_aria_map_set(i64, i8*, i64, i64)" + nl))
	p = append(p, ("declare i64 @_aria_map_get(i64, i8*, i64)" + nl))
	p = append(p, ("declare i64 @_aria_map_contains(i64, i8*, i64)" + nl))
	p = append(p, ("declare i64 @_aria_map_len(i64)" + nl))
	p = append(p, ("declare i64 @_aria_map_keys(i64)" + nl))
	p = append(p, ("declare i64 @_aria_set_new(i64)" + nl))
	p = append(p, ("declare void @_aria_set_add(i64, i8*, i64)" + nl))
	p = append(p, ("declare i64 @_aria_set_contains(i64, i8*, i64)" + nl))
	p = append(p, ("declare i64 @_aria_set_len(i64)" + nl))
	p = append(p, ("declare void @_aria_set_remove(i64, i8*, i64)" + nl))
	p = append(p, ("declare i64 @_aria_set_values(i64)" + nl))
	p = append(p, ((("declare " + si64) + " @_aria_getenv(i8*, i64)") + nl))
	p = append(p, ("declare i64 @_aria_tcp_socket()" + nl))
	p = append(p, ("declare i64 @_aria_tcp_bind(i64, i8*, i64, i64)" + nl))
	p = append(p, ("declare i64 @_aria_tcp_listen(i64, i64)" + nl))
	p = append(p, ("declare i64 @_aria_tcp_accept(i64)" + nl))
	p = append(p, ((("declare " + si64) + " @_aria_tcp_read(i64, i64)") + nl))
	p = append(p, ("declare i64 @_aria_tcp_write(i64, i8*, i64)" + nl))
	p = append(p, ("declare void @_aria_tcp_close(i64)" + nl))
	p = append(p, ((("declare " + si64) + " @_aria_tcp_peer_addr(i64)") + nl))
	p = append(p, ("declare i64 @_aria_tcp_set_timeout(i64, i64, i64)" + nl))
	p = append(p, ("declare i64 @_aria_pg_connect(i8*, i64)" + nl))
	p = append(p, ("declare void @_aria_pg_close(i64)" + nl))
	p = append(p, ("declare i64 @_aria_pg_status(i64)" + nl))
	p = append(p, ((("declare " + si64) + " @_aria_pg_error(i64)") + nl))
	p = append(p, ("declare i64 @_aria_pg_exec(i64, i8*, i64)" + nl))
	p = append(p, ("declare i64 @_aria_pg_exec_params(i64, i8*, i64, i64)" + nl))
	p = append(p, ("declare i64 @_aria_pg_result_status(i64)" + nl))
	p = append(p, ((("declare " + si64) + " @_aria_pg_result_error(i64)") + nl))
	p = append(p, ("declare i64 @_aria_pg_nrows(i64)" + nl))
	p = append(p, ("declare i64 @_aria_pg_ncols(i64)" + nl))
	p = append(p, ((("declare " + si64) + " @_aria_pg_field_name(i64, i64)") + nl))
	p = append(p, ((("declare " + si64) + " @_aria_pg_get_value(i64, i64, i64)") + nl))
	p = append(p, ("declare i64 @_aria_pg_is_null(i64, i64, i64)" + nl))
	p = append(p, ("declare void @_aria_pg_clear(i64)" + nl))
	p = append(p, ("declare i64 @_aria_spawn(i64, i64)" + nl))
	p = append(p, ("declare i64 @_aria_task_await(i64)" + nl))
	p = append(p, ("declare i64 @_aria_chan_new(i64)" + nl))
	p = append(p, ("declare i64 @_aria_chan_send(i64, i64)" + nl))
	p = append(p, ("declare i64 @_aria_chan_recv(i64)" + nl))
	p = append(p, ("declare void @_aria_chan_close(i64)" + nl))
	p = append(p, ("declare i64 @_aria_mutex_new()" + nl))
	p = append(p, ("declare void @_aria_mutex_lock(i64)" + nl))
	p = append(p, ("declare void @_aria_mutex_unlock(i64)" + nl))
	p = append(p, ("declare i64 @_aria_rwmutex_new()" + nl))
	p = append(p, ("declare void @_aria_rwmutex_rlock(i64)" + nl))
	p = append(p, ("declare void @_aria_rwmutex_runlock(i64)" + nl))
	p = append(p, ("declare void @_aria_rwmutex_wlock(i64)" + nl))
	p = append(p, ("declare void @_aria_rwmutex_wunlock(i64)" + nl))
	p = append(p, ("declare i64 @_aria_wg_new()" + nl))
	p = append(p, ("declare void @_aria_wg_add(i64, i64)" + nl))
	p = append(p, ("declare void @_aria_wg_done(i64)" + nl))
	p = append(p, ("declare void @_aria_wg_wait(i64)" + nl))
	p = append(p, ("declare i64 @_aria_once_new()" + nl))
	p = append(p, ("declare i64 @_aria_once_call(i64, i64, i64)" + nl))
	p = append(p, ((("declare " + si64) + " @_aria_chan_try_recv(i64)") + nl))
	p = append(p, ((("declare " + si64) + " @_aria_chan_select(i64, i64)") + nl))
	p = append(p, ("declare i64 @_aria_spawn2(i64, i64)" + nl))
	p = append(p, ("declare i64 @_aria_task_await2(i64)" + nl))
	p = append(p, ("declare i64 @_aria_task_done(i64)" + nl))
	p = append(p, ("declare void @_aria_task_cancel(i64)" + nl))
	p = append(p, ("declare i64 @_aria_task_result(i64)" + nl))
	p = append(p, ("declare i64 @_aria_cancel_check(i64)" + nl))
	p = append(p, ("declare i64 @_aria_cancel_new()" + nl))
	p = append(p, ("declare i64 @_aria_cancel_child(i64)" + nl))
	p = append(p, ("declare void @_aria_cancel_trigger(i64)" + nl))
	p = append(p, ("declare i64 @_aria_cancel_is_triggered(i64)" + nl))
	p = append(p, ("declare i64 @_aria_str_char_count(i8*, i64)" + nl))
	p = append(p, ("declare i64 @_aria_str_chars(i8*, i64)" + nl))
	p = append(p, ("declare i64 @_aria_str_graphemes(i8*, i64)" + nl))
	p = append(p, ("declare i64 @_aria_sb_new()" + nl))
	p = append(p, ("declare i64 @_aria_sb_with_capacity(i64)" + nl))
	p = append(p, ("declare void @_aria_sb_append(i64, i8*, i64)" + nl))
	p = append(p, ("declare void @_aria_sb_append_char(i64, i64)" + nl))
	p = append(p, ("declare i64 @_aria_sb_len(i64)" + nl))
	p = append(p, ((("declare " + si64) + " @_aria_sb_build(i64)") + nl))
	p = append(p, ("declare void @_aria_sb_clear(i64)" + nl))
	p = append(p, ((("declare " + si64) + " @_aria_format_int(i64, i8*, i64)") + nl))
	p = append(p, ((("declare " + si64) + " @_aria_format_float(i64, i8*, i64)") + nl))
	p = append(p, ("declare i64 @_aria_gc_collect()" + nl))
	p = append(p, ("declare i64 @_aria_gc_total_bytes()" + nl))
	p = append(p, ("declare i64 @_aria_gc_allocation_count()" + nl))
	p = append(p, ("declare i64 @_aria_arena_new(i64)" + nl))
	p = append(p, ("declare i8* @_aria_arena_alloc(i64, i64)" + nl))
	p = append(p, ("declare void @_aria_arena_reset(i64)" + nl))
	p = append(p, ("declare void @_aria_arena_free(i64)" + nl))
	p = append(p, ("declare i64 @_aria_arena_allocated(i64)" + nl))
	p = append(p, ("declare i64 @_aria_arena_capacity(i64)" + nl))
	p = append(p, ("declare i64 @_aria_pool_new(i64, i64)" + nl))
	p = append(p, ("declare i64 @_aria_pool_get(i64)" + nl))
	p = append(p, ("declare void @_aria_pool_put(i64, i64)" + nl))
	p = append(p, ((("declare " + si64) + " @_aria_float_to_str(i64)") + nl))
	p = append(p, nl)
	ob := "{"
	_ = ob
	cb := "}"
	_ = cb
	p = append(p, (("define void @_aria_panic(i8* %ptr, i64 %len) " + ob) + nl))
	p = append(p, ("  call void @_aria_eprintln_str(i8* %ptr, i64 %len)" + nl))
	p = append(p, ("  call void @_aria_exit(i64 1)" + nl))
	p = append(p, ("  unreachable" + nl))
	p = append(p, (cb + nl))
	p = append(p, nl)
	ei := int64(1)
	_ = ei
	for (ei < int64(len(m.Extern_names))) {
		ename := m.Extern_names[ei]
		_ = ename
		epc := m.Extern_param_counts[ei]
		_ = epc
		ert := m.Extern_ret_types[ei]
		_ = ert
		ret_llvm := "void"
		_ = ret_llvm
		if (ert == "i64") {
			ret_llvm = "i64"
		}
		if (ert == "f64") {
			ret_llvm = "double"
		}
		if (ert == "str") {
			ret_llvm = _str_ret_type()
		}
		params_llvm := ""
		_ = params_llvm
		pi := int64(0)
		_ = pi
		for (pi < epc) {
			if (pi > int64(0)) {
				params_llvm = (params_llvm + ", ")
			}
			params_llvm = (params_llvm + "i64")
			pi = (pi + int64(1))
		}
		p = append(p, ((((((("declare " + ret_llvm) + " @") + ename) + "(") + params_llvm) + ")") + nl))
		ei = (ei + int64(1))
	}
	p = append(p, nl)
	defined_names := []string{""}
	_ = defined_names
	fi := int64(1)
	_ = fi
	for (fi < int64(len(m.Funcs))) {
		f := m.Funcs[fi]
		_ = f
		if (f.Name != "") {
			defined_names = append(defined_names, f.Name)
			if (f.Name == "main") {
				defined_names = append(defined_names, "_aria_main")
			}
			fn_parts := _gen_function(m, f)
			_ = fn_parts
			fpi := int64(1)
			_ = fpi
			for (fpi < int64(len(fn_parts))) {
				p = append(p, fn_parts[fpi])
				fpi = (fpi + int64(1))
			}
			p = append(p, nl)
		}
		fi = (fi + int64(1))
	}
	called_names := _collect_called_names(m)
	_ = called_names
	ci := int64(1)
	_ = ci
	for (ci < int64(len(called_names))) {
		cname := called_names[ci]
		_ = cname
		if (_name_in_list(cname, defined_names) == false) {
			is_extern := false
			_ = is_extern
			exi := int64(1)
			_ = exi
			for (exi < int64(len(m.Extern_names))) {
				if (m.Extern_names[exi] == cname) {
					is_extern = true
				}
				exi = (exi + int64(1))
			}
			if ((strings.HasPrefix(cname, "_aria_") == false) && (is_extern == false)) {
				p = append(p, ((("declare i64 @" + cname) + "(...)") + nl))
			}
		}
		ci = (ci + int64(1))
	}
	lb := _lb()
	_ = lb
	rb := _rb()
	_ = rb
	if (m.Entry_func != "") {
		entry_call := m.Entry_func
		_ = entry_call
		if (entry_call == "main") {
			entry_call = "_aria_main"
		}
		p = append(p, ("; Entry wrapper" + nl))
		p = append(p, (("define void @_aria_entry() " + lb) + nl))
		p = append(p, ("entry:" + nl))
		p = append(p, ((("  call void @" + entry_call) + "()") + nl))
		p = append(p, ("  ret void" + nl))
		p = append(p, ((rb + nl) + nl))
	} else {
		has_main := false
		_ = has_main
		hmi := int64(1)
		_ = hmi
		for (hmi < int64(len(m.Funcs))) {
			if (m.Funcs[hmi].Name == "main") {
				has_main = true
			}
			hmi = (hmi + int64(1))
		}
		if has_main {
			p = append(p, ("; Entry wrapper (fn main)" + nl))
			p = append(p, (("define void @_aria_entry() " + lb) + nl))
			p = append(p, ("entry:" + nl))
			p = append(p, ("  call i64 @_aria_main()" + nl))
			p = append(p, ("  ret void" + nl))
			p = append(p, ((rb + nl) + nl))
		} else {
			p = append(p, (("define void @_aria_entry() " + lb) + nl))
			p = append(p, ("entry:" + nl))
			p = append(p, ("  ret void" + nl))
			p = append(p, ((rb + nl) + nl))
		}
	}
	return _join(p)
}

func _generate_llvm_ir_to_file_target(m IrModule, path string, target string) {
	nl := _nl()
	_ = nl
	si64 := _struct_i64_i64()
	_ = si64
	h := []string{""}
	_ = h
	triple := _target_triple(target)
	_ = triple
	layout := _target_datalayout(target)
	_ = layout
	h = append(h, ((("target datalayout = \"" + layout) + "\"") + nl))
	h = append(h, (((("target triple = \"" + triple) + "\"") + nl) + nl))
	si := int64(1)
	_ = si
	for (si < int64(len(m.String_constants))) {
		s := m.String_constants[si]
		_ = s
		escaped := _llvm_escape_str(s)
		_ = escaped
		slen := int64(len(s))
		_ = slen
		h = append(h, ((((((("@.str." + i2s(si)) + " = private unnamed_addr constant [") + i2s((slen + int64(1)))) + " x i8] c\"") + escaped) + "\\00\"") + nl))
		si = (si + int64(1))
	}
	h = append(h, nl)
	h = append(h, ("; Runtime declarations" + nl))
	h = append(h, ("declare void @_aria_exit(i64)" + nl))
	h = append(h, ("declare i64 @_aria_write(i64, i8*, i64)" + nl))
	h = append(h, ("declare void @_aria_println_str(i8*, i64)" + nl))
	h = append(h, ("declare void @_aria_print_str(i8*, i64)" + nl))
	h = append(h, ("declare void @_aria_eprintln_str(i8*, i64)" + nl))
	h = append(h, ("declare i8* @_aria_alloc(i64)" + nl))
	h = append(h, ("declare void @_aria_gc_frame_push(i64, i64)" + nl))
	h = append(h, ("declare void @_aria_gc_frame_pop()" + nl))
	h = append(h, ("declare void @_aria_memcpy(i8*, i8*, i64)" + nl))
	h = append(h, ((("declare " + si64) + " @_aria_read_file(i8*, i64)") + nl))
	h = append(h, ("declare void @_aria_write_file(i8*, i64, i8*, i64)" + nl))
	h = append(h, ("declare void @_aria_append_file(i8*, i64, i8*, i64)" + nl))
	h = append(h, ("declare void @_aria_write_binary_file(i8*, i64, i64*, i64)" + nl))
	h = append(h, ((("declare " + si64) + " @_aria_int_to_str(i64)") + nl))
	h = append(h, ((("declare " + si64) + " @_aria_str_concat(i8*, i64, i8*, i64)") + nl))
	h = append(h, ("declare i64 @_aria_str_eq(i8*, i64, i8*, i64)" + nl))
	h = append(h, ("declare i64 @_aria_str_cmp(i8*, i64, i8*, i64)" + nl))
	h = append(h, ((("declare " + si64) + " @_aria_str_charAt(i8*, i64, i64)") + nl))
	h = append(h, ((("declare " + si64) + " @_aria_str_substring(i8*, i64, i64, i64)") + nl))
	h = append(h, ("declare i64 @_aria_str_contains(i8*, i64, i8*, i64)" + nl))
	h = append(h, ("declare i64 @_aria_str_startsWith(i8*, i64, i8*, i64)" + nl))
	h = append(h, ("declare i64 @_aria_str_endsWith(i8*, i64, i8*, i64)" + nl))
	h = append(h, ("declare i64 @_aria_str_indexOf(i8*, i64, i8*, i64)" + nl))
	h = append(h, ((("declare " + si64) + " @_aria_str_trim(i8*, i64)") + nl))
	h = append(h, ((("declare " + si64) + " @_aria_str_replace(i8*, i64, i8*, i64, i8*, i64)") + nl))
	h = append(h, ((("declare " + si64) + " @_aria_str_toLower(i8*, i64)") + nl))
	h = append(h, ((("declare " + si64) + " @_aria_str_toUpper(i8*, i64)") + nl))
	h = append(h, ("declare i64 @_aria_str_split(i8*, i64, i8*, i64)" + nl))
	h = append(h, ("declare void @_aria_map_set_str(i64, i8*, i64, i64, i64)" + nl))
	h = append(h, ((("declare " + si64) + " @_aria_map_get_str(i64, i8*, i64)") + nl))
	h = append(h, ("declare i64 @_aria_array_new(i64)" + nl))
	h = append(h, ("declare i64 @_aria_array_len(i64)" + nl))
	h = append(h, ("declare i64 @_aria_array_get(i64, i64)" + nl))
	h = append(h, ("declare void @_aria_array_set(i64, i64, i64)" + nl))
	h = append(h, ("declare i64 @_aria_array_append(i64, i64)" + nl))
	h = append(h, ("declare i64 @_aria_array_slice(i64, i64)" + nl))
	h = append(h, ("declare i64 @_aria_list_dir(i8*, i64)" + nl))
	h = append(h, ("declare i64 @_aria_is_dir(i8*, i64)" + nl))
	h = append(h, ("declare i64 @_aria_args_get()" + nl))
	h = append(h, ("declare i64 @_aria_exec(i8*, i64)" + nl))
	h = append(h, ((("declare " + si64) + " @_aria_str_to_int(i8*, i64)") + nl))
	h = append(h, ("declare i64 @_aria_str_to_float(i8*, i64)" + nl))
	h = append(h, ("declare i64 @_aria_map_new(i64)" + nl))
	h = append(h, ("declare void @_aria_map_set(i64, i8*, i64, i64)" + nl))
	h = append(h, ("declare i64 @_aria_map_get(i64, i8*, i64)" + nl))
	h = append(h, ("declare i64 @_aria_map_contains(i64, i8*, i64)" + nl))
	h = append(h, ("declare i64 @_aria_map_len(i64)" + nl))
	h = append(h, ("declare i64 @_aria_map_keys(i64)" + nl))
	h = append(h, ("declare i64 @_aria_set_new(i64)" + nl))
	h = append(h, ("declare void @_aria_set_add(i64, i8*, i64)" + nl))
	h = append(h, ("declare i64 @_aria_set_contains(i64, i8*, i64)" + nl))
	h = append(h, ("declare i64 @_aria_set_len(i64)" + nl))
	h = append(h, ("declare void @_aria_set_remove(i64, i8*, i64)" + nl))
	h = append(h, ("declare i64 @_aria_set_values(i64)" + nl))
	h = append(h, ((("declare " + si64) + " @_aria_getenv(i8*, i64)") + nl))
	h = append(h, ("declare i64 @_aria_tcp_socket()" + nl))
	h = append(h, ("declare i64 @_aria_tcp_bind(i64, i8*, i64, i64)" + nl))
	h = append(h, ("declare i64 @_aria_tcp_listen(i64, i64)" + nl))
	h = append(h, ("declare i64 @_aria_tcp_accept(i64)" + nl))
	h = append(h, ((("declare " + si64) + " @_aria_tcp_read(i64, i64)") + nl))
	h = append(h, ("declare i64 @_aria_tcp_write(i64, i8*, i64)" + nl))
	h = append(h, ("declare void @_aria_tcp_close(i64)" + nl))
	h = append(h, ((("declare " + si64) + " @_aria_tcp_peer_addr(i64)") + nl))
	h = append(h, ("declare i64 @_aria_tcp_set_timeout(i64, i64, i64)" + nl))
	h = append(h, ("declare i64 @_aria_pg_connect(i8*, i64)" + nl))
	h = append(h, ("declare void @_aria_pg_close(i64)" + nl))
	h = append(h, ("declare i64 @_aria_pg_status(i64)" + nl))
	h = append(h, ((("declare " + si64) + " @_aria_pg_error(i64)") + nl))
	h = append(h, ("declare i64 @_aria_pg_exec(i64, i8*, i64)" + nl))
	h = append(h, ("declare i64 @_aria_pg_exec_params(i64, i8*, i64, i64)" + nl))
	h = append(h, ("declare i64 @_aria_pg_result_status(i64)" + nl))
	h = append(h, ((("declare " + si64) + " @_aria_pg_result_error(i64)") + nl))
	h = append(h, ("declare i64 @_aria_pg_nrows(i64)" + nl))
	h = append(h, ("declare i64 @_aria_pg_ncols(i64)" + nl))
	h = append(h, ((("declare " + si64) + " @_aria_pg_field_name(i64, i64)") + nl))
	h = append(h, ((("declare " + si64) + " @_aria_pg_get_value(i64, i64, i64)") + nl))
	h = append(h, ("declare i64 @_aria_pg_is_null(i64, i64, i64)" + nl))
	h = append(h, ("declare void @_aria_pg_clear(i64)" + nl))
	h = append(h, ("declare i64 @_aria_spawn(i64, i64)" + nl))
	h = append(h, ("declare i64 @_aria_task_await(i64)" + nl))
	h = append(h, ("declare i64 @_aria_chan_new(i64)" + nl))
	h = append(h, ("declare void @_aria_chan_send(i64, i64)" + nl))
	h = append(h, ("declare i64 @_aria_chan_recv(i64)" + nl))
	h = append(h, ("declare void @_aria_chan_close(i64)" + nl))
	h = append(h, ("declare i64 @_aria_chan_try_recv(i64)" + nl))
	h = append(h, ((("declare " + si64) + " @_aria_chan_select(i64, i64)") + nl))
	h = append(h, ("declare i64 @_aria_mutex_new()" + nl))
	h = append(h, ("declare void @_aria_mutex_lock(i64)" + nl))
	h = append(h, ("declare void @_aria_mutex_unlock(i64)" + nl))
	h = append(h, ("declare void @_aria_sleep(i64)" + nl))
	h = append(h, ("declare i64 @_aria_spawn2(i64, i64)" + nl))
	h = append(h, ("declare i64 @_aria_task_await2(i64)" + nl))
	h = append(h, ("declare i64 @_aria_task_done(i64)" + nl))
	h = append(h, ("declare void @_aria_task_cancel(i64)" + nl))
	h = append(h, ("declare i64 @_aria_task_result(i64)" + nl))
	h = append(h, ("declare i64 @_aria_cancel_check()" + nl))
	h = append(h, ("declare i64 @_aria_wg_new()" + nl))
	h = append(h, ("declare void @_aria_wg_add(i64, i64)" + nl))
	h = append(h, ("declare void @_aria_wg_done(i64)" + nl))
	h = append(h, ("declare void @_aria_wg_wait(i64)" + nl))
	h = append(h, ("declare i64 @_aria_once_new()" + nl))
	h = append(h, ("declare void @_aria_once_call(i64, i64, i64)" + nl))
	h = append(h, ("declare i64 @_aria_rwmutex_new()" + nl))
	h = append(h, ("declare void @_aria_rwmutex_rlock(i64)" + nl))
	h = append(h, ("declare void @_aria_rwmutex_runlock(i64)" + nl))
	h = append(h, ("declare void @_aria_rwmutex_wlock(i64)" + nl))
	h = append(h, ("declare void @_aria_rwmutex_wunlock(i64)" + nl))
	h = append(h, ("declare i64 @_aria_cancel_new()" + nl))
	h = append(h, ("declare i64 @_aria_cancel_child(i64)" + nl))
	h = append(h, ("declare void @_aria_cancel_trigger(i64)" + nl))
	h = append(h, ("declare i64 @_aria_cancel_is_triggered(i64)" + nl))
	h = append(h, ("declare i64 @_aria_str_char_count(i8*, i64)" + nl))
	h = append(h, ("declare i64 @_aria_str_chars(i8*, i64)" + nl))
	h = append(h, ("declare i64 @_aria_str_graphemes(i8*, i64)" + nl))
	h = append(h, ("declare i64 @_aria_sb_new()" + nl))
	h = append(h, ("declare i64 @_aria_sb_with_capacity(i64)" + nl))
	h = append(h, ("declare void @_aria_sb_append(i64, i8*, i64)" + nl))
	h = append(h, ("declare void @_aria_sb_append_char(i64, i64)" + nl))
	h = append(h, ("declare i64 @_aria_sb_len(i64)" + nl))
	h = append(h, ((("declare " + si64) + " @_aria_sb_build(i64)") + nl))
	h = append(h, ("declare void @_aria_sb_clear(i64)" + nl))
	h = append(h, ((("declare " + si64) + " @_aria_format_int(i64, i8*, i64)") + nl))
	h = append(h, ((("declare " + si64) + " @_aria_format_float(i64, i8*, i64)") + nl))
	h = append(h, ("declare i64 @_aria_gc_collect()" + nl))
	h = append(h, ("declare i64 @_aria_gc_total_bytes()" + nl))
	h = append(h, ("declare i64 @_aria_gc_allocation_count()" + nl))
	h = append(h, ("declare i64 @_aria_arena_new(i64)" + nl))
	h = append(h, ("declare i8* @_aria_arena_alloc(i64, i64)" + nl))
	h = append(h, ("declare void @_aria_arena_reset(i64)" + nl))
	h = append(h, ("declare void @_aria_arena_free(i64)" + nl))
	h = append(h, ("declare i64 @_aria_arena_allocated(i64)" + nl))
	h = append(h, ("declare i64 @_aria_arena_capacity(i64)" + nl))
	h = append(h, ("declare i64 @_aria_pool_new(i64, i64)" + nl))
	h = append(h, ("declare i64 @_aria_pool_get(i64)" + nl))
	h = append(h, ("declare void @_aria_pool_put(i64, i64)" + nl))
	h = append(h, ((("declare " + si64) + " @_aria_float_to_str(i64)") + nl))
	h = append(h, nl)
	ob := "{"
	_ = ob
	cb := "}"
	_ = cb
	h = append(h, (("define void @_aria_panic(i8* %ptr, i64 %len) " + ob) + nl))
	h = append(h, ("  call void @_aria_eprintln_str(i8* %ptr, i64 %len)" + nl))
	h = append(h, ("  call void @_aria_exit(i64 1)" + nl))
	h = append(h, ("  unreachable" + nl))
	h = append(h, (cb + nl))
	h = append(h, nl)
	ei := int64(1)
	_ = ei
	for (ei < int64(len(m.Extern_names))) {
		ename := m.Extern_names[ei]
		_ = ename
		epc := m.Extern_param_counts[ei]
		_ = epc
		ert := m.Extern_ret_types[ei]
		_ = ert
		ret_llvm := "void"
		_ = ret_llvm
		if (ert == "i64") {
			ret_llvm = "i64"
		}
		if (ert == "f64") {
			ret_llvm = "double"
		}
		if (ert == "str") {
			ret_llvm = _str_ret_type()
		}
		params_llvm := ""
		_ = params_llvm
		pi := int64(0)
		_ = pi
		for (pi < epc) {
			if (pi > int64(0)) {
				params_llvm = (params_llvm + ", ")
			}
			params_llvm = (params_llvm + "i64")
			pi = (pi + int64(1))
		}
		h = append(h, ((((((("declare " + ret_llvm) + " @") + ename) + "(") + params_llvm) + ")") + nl))
		ei = (ei + int64(1))
	}
	h = append(h, nl)
	_ariaWriteFile(path, _join(h))
	defined_names := []string{""}
	_ = defined_names
	fi := int64(1)
	_ = fi
	for (fi < int64(len(m.Funcs))) {
		f := m.Funcs[fi]
		_ = f
		if (f.Name != "") {
			defined_names = append(defined_names, f.Name)
			if (f.Name == "main") {
				defined_names = append(defined_names, "_aria_main")
			}
			fn_parts := _gen_function(m, f)
			_ = fn_parts
			fn_text := (_join(fn_parts) + nl)
			_ = fn_text
			_ariaAppendFile(path, fn_text)
		}
		fi = (fi + int64(1))
	}
	tail := []string{""}
	_ = tail
	called_names := _collect_called_names(m)
	_ = called_names
	ci := int64(1)
	_ = ci
	for (ci < int64(len(called_names))) {
		cname := called_names[ci]
		_ = cname
		if (_name_in_list(cname, defined_names) == false) {
			is_extern := false
			_ = is_extern
			exi := int64(1)
			_ = exi
			for (exi < int64(len(m.Extern_names))) {
				if (m.Extern_names[exi] == cname) {
					is_extern = true
				}
				exi = (exi + int64(1))
			}
			if ((strings.HasPrefix(cname, "_aria_") == false) && (is_extern == false)) {
				tail = append(tail, ((("declare i64 @" + cname) + "(...)") + nl))
			}
		}
		ci = (ci + int64(1))
	}
	lb := _lb()
	_ = lb
	rb := _rb()
	_ = rb
	if (m.Entry_func != "") {
		entry_call := m.Entry_func
		_ = entry_call
		if (entry_call == "main") {
			entry_call = "_aria_main"
		}
		tail = append(tail, ("; Entry wrapper" + nl))
		tail = append(tail, (("define void @_aria_entry() " + lb) + nl))
		tail = append(tail, ("entry:" + nl))
		tail = append(tail, ((("  call void @" + entry_call) + "()") + nl))
		tail = append(tail, ("  ret void" + nl))
		tail = append(tail, ((rb + nl) + nl))
	} else {
		has_main := false
		_ = has_main
		hmi := int64(1)
		_ = hmi
		for (hmi < int64(len(m.Funcs))) {
			if (m.Funcs[hmi].Name == "main") {
				has_main = true
			}
			hmi = (hmi + int64(1))
		}
		if has_main {
			tail = append(tail, ("; Entry wrapper (fn main)" + nl))
			tail = append(tail, (("define void @_aria_entry() " + lb) + nl))
			tail = append(tail, ("entry:" + nl))
			tail = append(tail, ("  call i64 @_aria_main()" + nl))
			tail = append(tail, ("  ret void" + nl))
			tail = append(tail, ((rb + nl) + nl))
		} else {
			tail = append(tail, (("define void @_aria_entry() " + lb) + nl))
			tail = append(tail, ("entry:" + nl))
			tail = append(tail, ("  ret void" + nl))
			tail = append(tail, ((rb + nl) + nl))
		}
	}
	_ariaAppendFile(path, _join(tail))
	_void := int64(0)
	_ = _void
}

func _gen_function(m IrModule, f IrFunc) []string {
	nl := _nl()
	_ = nl
	lb := _lb()
	_ = lb
	rb := _rb()
	_ = rb
	si64 := _struct_i64_i64()
	_ = si64
	ret_type := "void"
	_ = ret_type
	is_void := true
	_ = is_void
	if (f.Return_type == int64(12)) {
		ret_type = si64
		is_void = false
	} else if (f.Return_type > int64(0)) {
		ret_type = "i64"
		is_void = false
	}
	has_ret := false
	_ = has_ret
	ii := int64(1)
	_ = ii
	for (ii < int64(len(f.Insts))) {
		op := ir_op_name(f.Insts[ii].Op)
		_ = op
		if (op == "Ret") {
			has_ret = true
		}
		ii = (ii + int64(1))
	}
	if (has_ret && is_void) {
		ret_type = "i64"
		is_void = false
	}
	params := ""
	_ = params
	pi := int64(0)
	_ = pi
	for (pi < f.Param_count) {
		if (pi > int64(0)) {
			params = (params + ", ")
		}
		params = ((params + "i64 %p") + i2s(pi))
		pi = (pi + int64(1))
	}
	p := []string{""}
	_ = p
	cold_attr := ""
	_ = cold_attr
	cni := int64(1)
	_ = cni
	for (cni < int64(len(m.Extern_names))) {
		cni = (cni + int64(1))
	}
	emit_name := f.Name
	_ = emit_name
	if (f.Name == "main") {
		emit_name = "_aria_main"
	}
	p = append(p, (((((((("define " + ret_type) + " @") + emit_name) + "(") + params) + ") ") + lb) + nl))
	p = append(p, ("entry:" + nl))
	max_temp := f.Param_count
	_ = max_temp
	ti := int64(1)
	_ = ti
	for (ti < int64(len(f.Insts))) {
		inst := f.Insts[ti]
		_ = inst
		op := ir_op_name(inst.Op)
		_ = op
		if (inst.Dest > max_temp) {
			max_temp = inst.Dest
		}
		if (((((op != "Const") && (op != "ConstStr")) && (op != "Jump")) && (op != "Label")) && (op != "Nop")) {
			if (inst.Arg1 > max_temp) {
				max_temp = inst.Arg1
			}
		}
		if (((((((((((((((op != "FieldGet") && (op != "FieldSet")) && (op != "ConstStr")) && (op != "Call")) && (op != "CallIndirect")) && (op != "FnRef")) && (op != "BranchTrue")) && (op != "BranchFalse")) && (op != "Nop")) && (op != "Const")) && (op != "Alloc")) && (op != "Jump")) && (op != "Label")) && (op != "TraitObject")) && (op != "VtableCall")) {
			if (inst.Arg2 > max_temp) {
				max_temp = inst.Arg2
			}
		}
		ti = (ti + int64(1))
	}
	max_temp = (max_temp + int64(2))
	ai := int64(0)
	_ = ai
	for (ai <= max_temp) {
		p = append(p, ((("  " + _t(ai)) + " = alloca i64, align 8") + nl))
		ai = (ai + int64(1))
	}
	pi = int64(0)
	for (pi < f.Param_count) {
		p = append(p, (((("  store i64 %p" + i2s(pi)) + ", i64* ") + _t(pi)) + nl))
		pi = (pi + int64(1))
	}
	has_alloc := false
	_ = has_alloc
	gc_root_count := _gc_root_count_for_func(f)
	_ = gc_root_count
	aci := int64(1)
	_ = aci
	for (aci < int64(len(f.Insts))) {
		if (ir_op_name(f.Insts[aci].Op) == "Alloc") {
			has_alloc = true
		}
		aci = (aci + int64(1))
	}
	gc_frame_slots := (gc_root_count + int64(2))
	_ = gc_frame_slots
	p = append(p, ((("  %gc_frame = alloca [" + i2s(gc_frame_slots)) + " x i64], align 8") + nl))
	p = append(p, ((("  %gc_frame_ptr = ptrtoint [" + i2s(gc_frame_slots)) + " x i64]* %gc_frame to i64") + nl))
	if (gc_root_count > int64(0)) {
		zi := int64(0)
		_ = zi
		for (zi < gc_root_count) {
			p = append(p, (((((((("  %gc_init_" + i2s(zi)) + " = getelementptr [") + i2s(gc_frame_slots)) + " x i64], [") + i2s(gc_frame_slots)) + " x i64]* %gc_frame, i64 0, i64 ") + i2s((zi + int64(2)))) + nl))
			p = append(p, (("  store i64 0, i64* %gc_init_" + i2s(zi)) + nl))
			zi = (zi + int64(1))
		}
		p = append(p, ((("  call void @_aria_gc_frame_push(i64 %gc_frame_ptr, i64 " + i2s(gc_root_count)) + ")") + nl))
		pi2 := int64(0)
		_ = pi2
		for (pi2 < f.Param_count) {
			if _is_gc_root(f, pi2) {
				slot := _gc_slot_for_temp(f, pi2)
				_ = slot
				if (slot >= int64(0)) {
					pv := ("%gc_param_" + i2s(pi2))
					_ = pv
					pg := ("%gc_paramg_" + i2s(pi2))
					_ = pg
					p = append(p, (((("  " + pv) + " = load i64, i64* ") + _t(pi2)) + nl))
					p = append(p, (((((((("  " + pg) + " = getelementptr [") + i2s(gc_frame_slots)) + " x i64], [") + i2s(gc_frame_slots)) + " x i64]* %gc_frame, i64 0, i64 ") + i2s((slot + int64(2)))) + nl))
					p = append(p, (((("  store i64 " + pv) + ", i64* ") + pg) + nl))
				}
			}
			pi2 = (pi2 + int64(1))
		}
	}
	g := _lg_new()
	_ = g
	idx := int64(1)
	_ = idx
	for (idx < int64(len(f.Insts))) {
		inst := f.Insts[idx]
		_ = inst
		g = _gen_inst(g, m, f, inst, idx)
		idx = (idx + int64(1))
	}
	if (g.Last_was_term == false) {
		if is_void {
			g = _lg_emit_term(g, "ret void")
		} else if (ret_type == si64) {
			g = _lg_emit_term(g, (("ret " + si64) + " zeroinitializer"))
		} else {
			g = _lg_emit_term(g, "ret i64 0")
		}
	}
	gi := int64(1)
	_ = gi
	for (gi < int64(len(g.Parts))) {
		p = append(p, g.Parts[gi])
		gi = (gi + int64(1))
	}
	p = append(p, (rb + nl))
	return p
}

func _gen_inst(g LlvmGen, m IrModule, f IrFunc, inst IrInst, idx int64) LlvmGen {
	op := ir_op_name(inst.Op)
	_ = op
	if (op == "Nop") {
		return g
	}
	if (op == "Const") {
		return _lg_emit(g, ((("store i64 " + i2s(inst.Arg1)) + ", i64* ") + _t(inst.Dest)))
	}
	if (op == "ConstStr") {
		str_id := inst.Arg1
		_ = str_id
		str_len := inst.Arg2
		_ = str_len
		if (str_len == (int64(0) - int64(9))) {
			float_text := m.String_constants[str_id]
			_ = float_text
			gg := g
			_ = gg
			fval := ("%fconst_" + i2s(idx))
			_ = fval
			gg = _lg_emit(gg, (((fval + " = bitcast double ") + float_text) + " to i64"))
			gg = _lg_emit(gg, ((("store i64 " + fval) + ", i64* ") + _t(inst.Dest)))
			return gg
		}
		gg := g
		_ = gg
		lbl := _lg_fresh_label(gg)
		_ = lbl
		gg = lbl.G
		tmp_name := ("%str_ptr_" + i2s(lbl.Id))
		_ = tmp_name
		gg = _lg_emit(gg, (((((((tmp_name + " = getelementptr [") + i2s((str_len + int64(1)))) + " x i8], [") + i2s((str_len + int64(1)))) + " x i8]* @.str.") + i2s(str_id)) + ", i64 0, i64 0"))
		tmp_int := ("%str_int_" + i2s(lbl.Id))
		_ = tmp_int
		gg = _lg_emit(gg, (((tmp_int + " = ptrtoint i8* ") + tmp_name) + " to i64"))
		gg = _lg_emit(gg, ((("store i64 " + tmp_int) + ", i64* ") + _t(inst.Dest)))
		gg = _lg_emit(gg, ((("store i64 " + i2s(str_len)) + ", i64* ") + _t((inst.Dest + int64(1)))))
		return gg
	}
	if (op == "Add") {
		if (inst.S1 == "f64") {
			return _gen_float_binop(g, "fadd", inst, idx)
		}
		return _gen_binop(g, "add", inst)
	}
	if (op == "Sub") {
		if (inst.S1 == "f64") {
			return _gen_float_binop(g, "fsub", inst, idx)
		}
		return _gen_binop(g, "sub", inst)
	}
	if (op == "Mul") {
		if (inst.S1 == "f64") {
			return _gen_float_binop(g, "fmul", inst, idx)
		}
		return _gen_binop(g, "mul", inst)
	}
	if (op == "Div") {
		if (inst.S1 == "f64") {
			return _gen_float_binop(g, "fdiv", inst, idx)
		}
		return _gen_binop(g, "sdiv", inst)
	}
	if (op == "Mod") {
		return _gen_binop(g, "srem", inst)
	}
	if (op == "Neg") {
		if (inst.S1 == "f64") {
			gg := g
			_ = gg
			v := ("%fneg_v_" + i2s(idx))
			_ = v
			gg = _lg_emit(gg, ((v + " = load i64, i64* ") + _t(inst.Arg1)))
			dv := ("%fneg_d_" + i2s(idx))
			_ = dv
			gg = _lg_emit(gg, (((dv + " = bitcast i64 ") + v) + " to double"))
			res := ("%fneg_r_" + i2s(idx))
			_ = res
			gg = _lg_emit(gg, ((res + " = fneg double ") + dv))
			ri := ("%fneg_i_" + i2s(idx))
			_ = ri
			gg = _lg_emit(gg, (((ri + " = bitcast double ") + res) + " to i64"))
			gg = _lg_emit(gg, ((("store i64 " + ri) + ", i64* ") + _t(inst.Dest)))
			return gg
		}
		gg := g
		_ = gg
		v := _v(inst.Arg1)
		_ = v
		gg = _lg_emit(gg, ((v + " = load i64, i64* ") + _t(inst.Arg1)))
		res := ("%neg_" + i2s(idx))
		_ = res
		gg = _lg_emit(gg, ((res + " = sub i64 0, ") + v))
		gg = _lg_emit(gg, ((("store i64 " + res) + ", i64* ") + _t(inst.Dest)))
		return gg
	}
	if (op == "Eq") {
		if (inst.S1 == "f64") {
			return _gen_fcmp(g, "oeq", inst, idx)
		}
		return _gen_cmp(g, "eq", inst, idx)
	}
	if (op == "Neq") {
		if (inst.S1 == "f64") {
			return _gen_fcmp(g, "une", inst, idx)
		}
		return _gen_cmp(g, "ne", inst, idx)
	}
	if (op == "Lt") {
		if (inst.S1 == "f64") {
			return _gen_fcmp(g, "olt", inst, idx)
		}
		return _gen_cmp(g, "slt", inst, idx)
	}
	if (op == "Gt") {
		if (inst.S1 == "f64") {
			return _gen_fcmp(g, "ogt", inst, idx)
		}
		return _gen_cmp(g, "sgt", inst, idx)
	}
	if (op == "Lte") {
		if (inst.S1 == "f64") {
			return _gen_fcmp(g, "ole", inst, idx)
		}
		return _gen_cmp(g, "sle", inst, idx)
	}
	if (op == "Gte") {
		if (inst.S1 == "f64") {
			return _gen_fcmp(g, "oge", inst, idx)
		}
		return _gen_cmp(g, "sge", inst, idx)
	}
	if (op == "And") {
		return _gen_binop(g, "and", inst)
	}
	if (op == "Or") {
		return _gen_binop(g, "or", inst)
	}
	if (op == "Not") {
		gg := g
		_ = gg
		v := _v(inst.Arg1)
		_ = v
		gg = _lg_emit(gg, ((v + " = load i64, i64* ") + _t(inst.Arg1)))
		cmp := ("%not_cmp_" + i2s(idx))
		_ = cmp
		gg = _lg_emit(gg, (((cmp + " = icmp eq i64 ") + v) + ", 0"))
		res := ("%not_" + i2s(idx))
		_ = res
		gg = _lg_emit(gg, (((res + " = zext i1 ") + cmp) + " to i64"))
		gg = _lg_emit(gg, ((("store i64 " + res) + ", i64* ") + _t(inst.Dest)))
		return gg
	}
	if (op == "BitAnd") {
		return _gen_binop(g, "and", inst)
	}
	if (op == "BitOr") {
		return _gen_binop(g, "or", inst)
	}
	if (op == "BitXor") {
		return _gen_binop(g, "xor", inst)
	}
	if (op == "Shl") {
		return _gen_binop(g, "shl", inst)
	}
	if (op == "Shr") {
		return _gen_binop(g, "ashr", inst)
	}
	if (op == "Jump") {
		gg := g
		_ = gg
		if (gg.Last_was_term == false) {
			gg = _lg_emit_term(gg, ("br label %" + _lbl(inst.Arg1)))
		}
		return gg
	}
	if (op == "BranchTrue") {
		gg := g
		_ = gg
		v := ("%br_val_" + i2s(idx))
		_ = v
		gg = _lg_emit(gg, ((v + " = load i64, i64* ") + _t(inst.Arg1)))
		cmp := ("%br_cmp_" + i2s(idx))
		_ = cmp
		gg = _lg_emit(gg, (((cmp + " = icmp ne i64 ") + v) + ", 0"))
		lbl := _lg_fresh_label(gg)
		_ = lbl
		gg = lbl.G
		fall := _fall(lbl.Id)
		_ = fall
		gg = _lg_emit_term(gg, ((((("br i1 " + cmp) + ", label %") + _lbl(inst.Arg2)) + ", label %") + fall))
		gg = _lg_emit_label(gg, fall)
		return gg
	}
	if (op == "BranchFalse") {
		gg := g
		_ = gg
		v := ("%br_val_" + i2s(idx))
		_ = v
		gg = _lg_emit(gg, ((v + " = load i64, i64* ") + _t(inst.Arg1)))
		cmp := ("%br_cmp_" + i2s(idx))
		_ = cmp
		gg = _lg_emit(gg, (((cmp + " = icmp eq i64 ") + v) + ", 0"))
		lbl := _lg_fresh_label(gg)
		_ = lbl
		gg = lbl.G
		fall := _fall(lbl.Id)
		_ = fall
		gg = _lg_emit_term(gg, ((((("br i1 " + cmp) + ", label %") + _lbl(inst.Arg2)) + ", label %") + fall))
		gg = _lg_emit_label(gg, fall)
		return gg
	}
	if (op == "Label") {
		gg := g
		_ = gg
		if (gg.Last_was_term == false) {
			gg = _lg_emit_term(gg, ("br label %" + _lbl(inst.Arg1)))
		}
		gg = _lg_emit_label(gg, _lbl(inst.Arg1))
		return gg
	}
	if (op == "Ret") {
		gg := g
		_ = gg
		gc_rc := _gc_root_count_for_func(f)
		_ = gc_rc
		if (gc_rc > int64(0)) {
			gg = _lg_emit(gg, "call void @_aria_gc_frame_pop()")
		}
		si64 := _struct_i64_i64()
		_ = si64
		if (f.Return_type == int64(12)) {
			ptr_v := ("%ret_ptr_" + i2s(idx))
			_ = ptr_v
			len_v := ("%ret_len_" + i2s(idx))
			_ = len_v
			gg = _lg_emit(gg, ((ptr_v + " = load i64, i64* ") + _t(inst.Arg1)))
			gg = _lg_emit(gg, ((len_v + " = load i64, i64* ") + _t((inst.Arg1 + int64(1)))))
			r1 := ("%ret_s1_" + i2s(idx))
			_ = r1
			r2 := ("%ret_s2_" + i2s(idx))
			_ = r2
			gg = _lg_emit(gg, (((((r1 + " = insertvalue ") + si64) + " undef, i64 ") + ptr_v) + ", 0"))
			gg = _lg_emit(gg, (((((((r2 + " = insertvalue ") + si64) + " ") + r1) + ", i64 ") + len_v) + ", 1"))
			gg = _lg_emit_term(gg, ((("ret " + si64) + " ") + r2))
		} else {
			v := ("%ret_v_" + i2s(idx))
			_ = v
			gg = _lg_emit(gg, ((v + " = load i64, i64* ") + _t(inst.Arg1)))
			gg = _lg_emit_term(gg, ("ret i64 " + v))
		}
		return gg
	}
	if (op == "RetVoid") {
		gg := g
		_ = gg
		gc_rcv := _gc_root_count_for_func(f)
		_ = gc_rcv
		if (gc_rcv > int64(0)) {
			gg = _lg_emit(gg, "call void @_aria_gc_frame_pop()")
		}
		return _lg_emit_term(gg, "ret void")
	}
	if (op == "Arg") {
		return g
	}
	if (op == "Local") {
		return g
	}
	if (op == "Store") {
		gg := g
		_ = gg
		v := ("%store_v_" + i2s(idx))
		_ = v
		gg = _lg_emit(gg, ((v + " = load i64, i64* ") + _t(inst.Arg1)))
		gg = _lg_emit(gg, ((("store i64 " + v) + ", i64* ") + _t(inst.Dest)))
		if _is_gc_root(f, inst.Dest) {
			gc_frame_slots := (_gc_root_count_for_func(f) + int64(2))
			_ = gc_frame_slots
			slot := _gc_slot_for_temp(f, inst.Dest)
			_ = slot
			if ((slot >= int64(0)) && (gc_frame_slots > int64(2))) {
				gep := ("%gc_ss_" + i2s(idx))
				_ = gep
				gg = _lg_emit(gg, ((((((gep + " = getelementptr [") + i2s(gc_frame_slots)) + " x i64], [") + i2s(gc_frame_slots)) + " x i64]* %gc_frame, i64 0, i64 ") + i2s((slot + int64(2)))))
				gg = _lg_emit(gg, ((("store i64 " + v) + ", i64* ") + gep))
			}
		}
		return gg
	}
	if (op == "Load") {
		gg := g
		_ = gg
		v := ("%load_v_" + i2s(idx))
		_ = v
		gg = _lg_emit(gg, ((v + " = load i64, i64* ") + _t(inst.Arg1)))
		gg = _lg_emit(gg, ((("store i64 " + v) + ", i64* ") + _t(inst.Dest)))
		gg = _gc_frame_update(gg, f, inst.Dest, v, idx)
		return gg
	}
	if (op == "Call") {
		gg := _gen_call(g, m, f, inst, idx)
		_ = gg
		cv := ("%gc_cr_" + i2s(idx))
		_ = cv
		gg = _lg_emit(gg, ((cv + " = load i64, i64* ") + _t(inst.Dest)))
		gg = _gc_frame_update(gg, f, inst.Dest, cv, idx)
		return gg
	}
	if (op == "FnRef") {
		gg := g
		_ = gg
		raw := ("%fnref_raw_" + i2s(idx))
		_ = raw
		gg = _lg_emit(gg, (((raw + " = bitcast i64 (i64)* @") + inst.S1) + " to i8*"))
		ptr := ("%fnref_ptr_" + i2s(idx))
		_ = ptr
		gg = _lg_emit(gg, (((ptr + " = ptrtoint i8* ") + raw) + " to i64"))
		gg = _lg_emit(gg, ((("store i64 " + ptr) + ", i64* ") + _t(inst.Dest)))
		return gg
	}
	if (op == "CallIndirect") {
		gg := g
		_ = gg
		fn_ptr := ("%icall_ptr_" + i2s(idx))
		_ = fn_ptr
		gg = _lg_emit(gg, ((fn_ptr + " = load i64, i64* ") + _t(inst.Arg1)))
		fn_cast := ("%icall_fn_" + i2s(idx))
		_ = fn_cast
		is_str_ret := (inst.S1 == "str")
		_ = is_str_ret
		ret_ty := "i64"
		_ = ret_ty
		if is_str_ret {
			ret_ty = _struct_i64_i64()
		}
		sig := (ret_ty + " (")
		_ = sig
		si := int64(0)
		_ = si
		for (si < inst.Arg2) {
			if (si > int64(0)) {
				sig = (sig + ", ")
			}
			sig = (sig + "i64")
			si = (si + int64(1))
		}
		sig = (sig + ")*")
		gg = _lg_emit(gg, ((((fn_cast + " = inttoptr i64 ") + fn_ptr) + " to ") + sig))
		args := ""
		_ = args
		ai := int64(0)
		_ = ai
		for (ai < inst.Arg2) {
			av := ((("%icall_a" + i2s(ai)) + "_") + i2s(idx))
			_ = av
			gg = _lg_emit(gg, ((av + " = load i64, i64* ") + _t((inst.Type_id + ai))))
			if (ai > int64(0)) {
				args = (args + ", ")
			}
			args = ((args + "i64 ") + av)
			ai = (ai + int64(1))
		}
		if is_str_ret {
			si64 := _struct_i64_i64()
			_ = si64
			result := ("%icall_r_" + i2s(idx))
			_ = result
			gg = _lg_emit(gg, (((((((result + " = call ") + si64) + " ") + fn_cast) + "(") + args) + ")"))
			rptr := ("%icall_rptr_" + i2s(idx))
			_ = rptr
			rlen := ("%icall_rlen_" + i2s(idx))
			_ = rlen
			gg = _lg_emit(gg, (((((rptr + " = extractvalue ") + si64) + " ") + result) + ", 0"))
			gg = _lg_emit(gg, (((((rlen + " = extractvalue ") + si64) + " ") + result) + ", 1"))
			gg = _lg_emit(gg, ((("store i64 " + rptr) + ", i64* ") + _t(inst.Dest)))
			gg = _lg_emit(gg, ((("store i64 " + rlen) + ", i64* ") + _t((inst.Dest + int64(1)))))
		} else {
			result := ("%icall_r_" + i2s(idx))
			_ = result
			gg = _lg_emit(gg, (((((result + " = call i64 ") + fn_cast) + "(") + args) + ")"))
			gg = _lg_emit(gg, ((("store i64 " + result) + ", i64* ") + _t(inst.Dest)))
			gg = _gc_frame_update(gg, f, inst.Dest, result, idx)
		}
		return gg
	}
	if (op == "Alloc") {
		gg := g
		_ = gg
		size := (inst.Arg1 * int64(8))
		_ = size
		raw := ("%alloc_raw_" + i2s(idx))
		_ = raw
		gg = _lg_emit(gg, (((raw + " = call i8* @_aria_alloc(i64 ") + i2s(size)) + ")"))
		ptr := ("%alloc_ptr_" + i2s(idx))
		_ = ptr
		gg = _lg_emit(gg, (((ptr + " = ptrtoint i8* ") + raw) + " to i64"))
		gg = _lg_emit(gg, ((("store i64 " + ptr) + ", i64* ") + _t(inst.Dest)))
		if _is_gc_root(f, inst.Dest) {
			gc_frame_slots := (_gc_root_count_for_func(f) + int64(2))
			_ = gc_frame_slots
			slot := _gc_slot_for_temp(f, inst.Dest)
			_ = slot
			if (slot >= int64(0)) {
				gep := ("%gc_store_" + i2s(idx))
				_ = gep
				gg = _lg_emit(gg, ((((((gep + " = getelementptr [") + i2s(gc_frame_slots)) + " x i64], [") + i2s(gc_frame_slots)) + " x i64]* %gc_frame, i64 0, i64 ") + i2s((slot + int64(2)))))
				gg = _lg_emit(gg, ((("store i64 " + ptr) + ", i64* ") + gep))
			}
		}
		return gg
	}
	if (op == "TraitObject") {
		gg := g
		_ = gg
		raw := ("%to_raw_" + i2s(idx))
		_ = raw
		gg = _lg_emit(gg, (raw + " = call i8* @_aria_alloc(i64 16)"))
		ptr := ("%to_ptr_" + i2s(idx))
		_ = ptr
		gg = _lg_emit(gg, (((ptr + " = bitcast i8* ") + raw) + " to i64*"))
		d0 := ("%to_d0_" + i2s(idx))
		_ = d0
		gg = _lg_emit(gg, ((d0 + " = load i64, i64* ") + _t(inst.Arg1)))
		gg = _lg_emit(gg, ((("store i64 " + d0) + ", i64* ") + ptr))
		gep1 := ("%to_gep1_" + i2s(idx))
		_ = gep1
		gg = _lg_emit(gg, (((gep1 + " = getelementptr i64, i64* ") + ptr) + ", i64 1"))
		v1 := ("%to_v1_" + i2s(idx))
		_ = v1
		gg = _lg_emit(gg, ((v1 + " = load i64, i64* ") + _t(inst.Arg2)))
		gg = _lg_emit(gg, ((("store i64 " + v1) + ", i64* ") + gep1))
		res := ("%to_res_" + i2s(idx))
		_ = res
		gg = _lg_emit(gg, (((res + " = ptrtoint i64* ") + ptr) + " to i64"))
		gg = _lg_emit(gg, ((("store i64 " + res) + ", i64* ") + _t(inst.Dest)))
		gg = _gc_frame_update(gg, f, inst.Dest, res, idx)
		return gg
	}
	if (op == "VtableCall") {
		gg := g
		_ = gg
		to_ptr := ("%vc_to_" + i2s(idx))
		_ = to_ptr
		gg = _lg_emit(gg, ((to_ptr + " = load i64, i64* ") + _t(inst.Arg1)))
		to_raw := ("%vc_to_raw_" + i2s(idx))
		_ = to_raw
		gg = _lg_emit(gg, (((to_raw + " = inttoptr i64 ") + to_ptr) + " to i64*"))
		data_ptr := ("%vc_data_" + i2s(idx))
		_ = data_ptr
		gg = _lg_emit(gg, ((data_ptr + " = load i64, i64* ") + to_raw))
		vt_gep := ("%vc_vt_gep_" + i2s(idx))
		_ = vt_gep
		gg = _lg_emit(gg, (((vt_gep + " = getelementptr i64, i64* ") + to_raw) + ", i64 1"))
		vt_ptr := ("%vc_vt_" + i2s(idx))
		_ = vt_ptr
		gg = _lg_emit(gg, ((vt_ptr + " = load i64, i64* ") + vt_gep))
		vt_raw := ("%vc_vt_raw_" + i2s(idx))
		_ = vt_raw
		gg = _lg_emit(gg, (((vt_raw + " = inttoptr i64 ") + vt_ptr) + " to i64*"))
		fn_gep := ("%vc_fn_gep_" + i2s(idx))
		_ = fn_gep
		gg = _lg_emit(gg, ((((fn_gep + " = getelementptr i64, i64* ") + vt_raw) + ", i64 ") + i2s(inst.Arg2)))
		fn_ptr := ("%vc_fn_" + i2s(idx))
		_ = fn_ptr
		gg = _lg_emit(gg, ((fn_ptr + " = load i64, i64* ") + fn_gep))
		fn_cast := ("%vc_cast_" + i2s(idx))
		_ = fn_cast
		if (inst.Type_id == int64(12)) {
			si64 := _struct_i64_i64()
			_ = si64
			gg = _lg_emit(gg, (((((fn_cast + " = inttoptr i64 ") + fn_ptr) + " to ") + si64) + " (i64)*"))
			result := ("%vc_result_" + i2s(idx))
			_ = result
			gg = _lg_emit(gg, (((((((result + " = call ") + si64) + " ") + fn_cast) + "(i64 ") + data_ptr) + ")"))
			rptr := ("%vc_rptr_" + i2s(idx))
			_ = rptr
			rlen := ("%vc_rlen_" + i2s(idx))
			_ = rlen
			gg = _lg_emit(gg, (((((rptr + " = extractvalue ") + si64) + " ") + result) + ", 0"))
			gg = _lg_emit(gg, (((((rlen + " = extractvalue ") + si64) + " ") + result) + ", 1"))
			gg = _lg_emit(gg, ((("store i64 " + rptr) + ", i64* ") + _t(inst.Dest)))
			gg = _lg_emit(gg, ((("store i64 " + rlen) + ", i64* ") + _t((inst.Dest + int64(1)))))
		} else {
			gg = _lg_emit(gg, (((fn_cast + " = inttoptr i64 ") + fn_ptr) + " to i64 (i64)*"))
			result := ("%vc_result_" + i2s(idx))
			_ = result
			gg = _lg_emit(gg, (((((result + " = call i64 ") + fn_cast) + "(i64 ") + data_ptr) + ")"))
			gg = _lg_emit(gg, ((("store i64 " + result) + ", i64* ") + _t(inst.Dest)))
			gg = _gc_frame_update(gg, f, inst.Dest, result, idx)
		}
		return gg
	}
	if (op == "FieldGet") {
		gg := g
		_ = gg
		base := ("%fg_base_" + i2s(idx))
		_ = base
		gg = _lg_emit(gg, ((base + " = load i64, i64* ") + _t(inst.Arg1)))
		raw := ("%fg_raw_" + i2s(idx))
		_ = raw
		gg = _lg_emit(gg, (((raw + " = inttoptr i64 ") + base) + " to i64*"))
		gep := ("%fg_gep_" + i2s(idx))
		_ = gep
		gg = _lg_emit(gg, ((((gep + " = getelementptr i64, i64* ") + raw) + ", i64 ") + i2s(inst.Arg2)))
		val := ("%fg_val_" + i2s(idx))
		_ = val
		gg = _lg_emit(gg, ((val + " = load i64, i64* ") + gep))
		gg = _lg_emit(gg, ((("store i64 " + val) + ", i64* ") + _t(inst.Dest)))
		gg = _gc_frame_update(gg, f, inst.Dest, val, idx)
		return gg
	}
	if (op == "FieldSet") {
		gg := g
		_ = gg
		base := ("%fs_base_" + i2s(idx))
		_ = base
		gg = _lg_emit(gg, ((base + " = load i64, i64* ") + _t(inst.Dest)))
		raw := ("%fs_raw_" + i2s(idx))
		_ = raw
		gg = _lg_emit(gg, (((raw + " = inttoptr i64 ") + base) + " to i64*"))
		gep := ("%fs_gep_" + i2s(idx))
		_ = gep
		gg = _lg_emit(gg, ((((gep + " = getelementptr i64, i64* ") + raw) + ", i64 ") + i2s(inst.Arg2)))
		val := ("%fs_val_" + i2s(idx))
		_ = val
		gg = _lg_emit(gg, ((val + " = load i64, i64* ") + _t(inst.Arg1)))
		gg = _lg_emit(gg, ((("store i64 " + val) + ", i64* ") + gep))
		return gg
	}
	if (op == "ArrayNew") {
		gg := g
		_ = gg
		arr := ("%anew_" + i2s(idx))
		_ = arr
		gg = _lg_emit(gg, (arr + " = call i64 @_aria_array_new(i64 8)"))
		gg = _lg_emit(gg, ((("store i64 " + arr) + ", i64* ") + _t(inst.Dest)))
		gg = _gc_frame_update(gg, f, inst.Dest, arr, idx)
		return gg
	}
	if (op == "ArrayLen") {
		gg := g
		_ = gg
		arr_v := ("%alen_arr_" + i2s(idx))
		_ = arr_v
		gg = _lg_emit(gg, ((arr_v + " = load i64, i64* ") + _t(inst.Arg1)))
		len_v := ("%alen_" + i2s(idx))
		_ = len_v
		gg = _lg_emit(gg, (((len_v + " = call i64 @_aria_array_len(i64 ") + arr_v) + ")"))
		gg = _lg_emit(gg, ((("store i64 " + len_v) + ", i64* ") + _t(inst.Dest)))
		return gg
	}
	if (op == "ArrayGet") {
		gg := g
		_ = gg
		arr_v := ("%aget_arr_" + i2s(idx))
		_ = arr_v
		gg = _lg_emit(gg, ((arr_v + " = load i64, i64* ") + _t(inst.Arg1)))
		idx_v := ("%aget_idx_" + i2s(idx))
		_ = idx_v
		gg = _lg_emit(gg, ((idx_v + " = load i64, i64* ") + _t(inst.Arg2)))
		val := ("%aget_" + i2s(idx))
		_ = val
		gg = _lg_emit(gg, (((((val + " = call i64 @_aria_array_get(i64 ") + arr_v) + ", i64 ") + idx_v) + ")"))
		if (inst.Type_id == int64(12)) {
			sp := ("%aget_sp_" + i2s(idx))
			_ = sp
			gg = _lg_emit(gg, (((sp + " = inttoptr i64 ") + val) + " to i64*"))
			sv0 := ("%aget_s0_" + i2s(idx))
			_ = sv0
			gg = _lg_emit(gg, ((sv0 + " = load i64, i64* ") + sp))
			gg = _lg_emit(gg, ((("store i64 " + sv0) + ", i64* ") + _t(inst.Dest)))
			sv1g := ("%aget_s1g_" + i2s(idx))
			_ = sv1g
			gg = _lg_emit(gg, (((sv1g + " = getelementptr i64, i64* ") + sp) + ", i64 1"))
			sv1 := ("%aget_s1_" + i2s(idx))
			_ = sv1
			gg = _lg_emit(gg, ((sv1 + " = load i64, i64* ") + sv1g))
			gg = _lg_emit(gg, ((("store i64 " + sv1) + ", i64* ") + _t((inst.Dest + int64(1)))))
		} else {
			gg = _lg_emit(gg, ((("store i64 " + val) + ", i64* ") + _t(inst.Dest)))
			gg = _gc_frame_update(gg, f, inst.Dest, val, idx)
		}
		return gg
	}
	if (op == "ArraySlice") {
		gg := g
		_ = gg
		asl_arr := ("%aslice_arr_" + i2s(idx))
		_ = asl_arr
		gg = _lg_emit(gg, ((asl_arr + " = load i64, i64* ") + _t(inst.Arg1)))
		asl_idx := ("%aslice_idx_" + i2s(idx))
		_ = asl_idx
		gg = _lg_emit(gg, ((asl_idx + " = load i64, i64* ") + _t(inst.Arg2)))
		asl_res := ("%aslice_" + i2s(idx))
		_ = asl_res
		gg = _lg_emit(gg, (((((asl_res + " = call i64 @_aria_array_slice(i64 ") + asl_arr) + ", i64 ") + asl_idx) + ")"))
		gg = _lg_emit(gg, ((("store i64 " + asl_res) + ", i64* ") + _t(inst.Dest)))
		gg = _gc_frame_update(gg, f, inst.Dest, asl_res, idx)
		return gg
	}
	if (op == "ArraySet") {
		gg := g
		_ = gg
		arr_v := ("%aset_arr_" + i2s(idx))
		_ = arr_v
		gg = _lg_emit(gg, ((arr_v + " = load i64, i64* ") + _t(inst.Dest)))
		idx_v := ("%aset_idx_" + i2s(idx))
		_ = idx_v
		gg = _lg_emit(gg, ((idx_v + " = load i64, i64* ") + _t(inst.Arg1)))
		val := ("%aset_val_" + i2s(idx))
		_ = val
		gg = _lg_emit(gg, ((val + " = load i64, i64* ") + _t(inst.Arg2)))
		gg = _lg_emit(gg, (((((("call void @_aria_array_set(i64 " + arr_v) + ", i64 ") + idx_v) + ", i64 ") + val) + ")"))
		return gg
	}
	if (op == "ArrayAppend") {
		gg := g
		_ = gg
		arr_v := ("%aapp_arr_" + i2s(idx))
		_ = arr_v
		gg = _lg_emit(gg, ((arr_v + " = load i64, i64* ") + _t(inst.Arg1)))
		if (inst.Type_id == int64(12)) {
			sp := ("%aapp_sp_" + i2s(idx))
			_ = sp
			gg = _lg_emit(gg, (sp + " = call i8* @_aria_alloc(i64 16)"))
			spi := ("%aapp_spi_" + i2s(idx))
			_ = spi
			gg = _lg_emit(gg, (((spi + " = bitcast i8* ") + sp) + " to i64*"))
			sv0 := ("%aapp_sv0_" + i2s(idx))
			_ = sv0
			gg = _lg_emit(gg, ((sv0 + " = load i64, i64* ") + _t(inst.Arg2)))
			gg = _lg_emit(gg, ((("store i64 " + sv0) + ", i64* ") + spi))
			sg1 := ("%aapp_sg1_" + i2s(idx))
			_ = sg1
			gg = _lg_emit(gg, (((sg1 + " = getelementptr i64, i64* ") + spi) + ", i64 1"))
			sv1 := ("%aapp_sv1_" + i2s(idx))
			_ = sv1
			gg = _lg_emit(gg, ((sv1 + " = load i64, i64* ") + _t((inst.Arg2 + int64(1)))))
			gg = _lg_emit(gg, ((("store i64 " + sv1) + ", i64* ") + sg1))
			spp := ("%aapp_spp_" + i2s(idx))
			_ = spp
			gg = _lg_emit(gg, (((spp + " = ptrtoint i64* ") + spi) + " to i64"))
			res := ("%aapp_res_" + i2s(idx))
			_ = res
			gg = _lg_emit(gg, (((((res + " = call i64 @_aria_array_append(i64 ") + arr_v) + ", i64 ") + spp) + ")"))
			gg = _lg_emit(gg, ((("store i64 " + res) + ", i64* ") + _t(inst.Dest)))
			gg = _gc_frame_update(gg, f, inst.Dest, res, idx)
		} else {
			val := ("%aapp_val_" + i2s(idx))
			_ = val
			gg = _lg_emit(gg, ((val + " = load i64, i64* ") + _t(inst.Arg2)))
			res := ("%aapp_res_" + i2s(idx))
			_ = res
			gg = _lg_emit(gg, (((((res + " = call i64 @_aria_array_append(i64 ") + arr_v) + ", i64 ") + val) + ")"))
			gg = _lg_emit(gg, ((("store i64 " + res) + ", i64* ") + _t(inst.Dest)))
			gg = _gc_frame_update(gg, f, inst.Dest, res, idx)
		}
		return gg
	}
	if (op == "StrLen") {
		gg := g
		_ = gg
		v := ("%slen_" + i2s(idx))
		_ = v
		gg = _lg_emit(gg, ((v + " = load i64, i64* ") + _t((inst.Arg1 + int64(1)))))
		gg = _lg_emit(gg, ((("store i64 " + v) + ", i64* ") + _t(inst.Dest)))
		return gg
	}
	if (op == "StrConcat") {
		return _gen_str_binop_returning_str(g, "@_aria_str_concat", inst, idx)
	}
	if (op == "StrEq") {
		return _gen_str_binop_returning_i64(g, "@_aria_str_eq", inst, idx)
	}
	if (op == "StrCmp") {
		return _gen_str_binop_returning_i64(g, "@_aria_str_cmp", inst, idx)
	}
	if (op == "StrCharAt") {
		gg := g
		_ = gg
		s_ptr := ("%sca_sp_" + i2s(idx))
		_ = s_ptr
		s_len := ("%sca_sl_" + i2s(idx))
		_ = s_len
		gg = _lg_emit(gg, ((s_ptr + " = load i64, i64* ") + _t(inst.Arg1)))
		gg = _lg_emit(gg, ((s_len + " = load i64, i64* ") + _t((inst.Arg1 + int64(1)))))
		idx_v := ("%sca_idx_" + i2s(idx))
		_ = idx_v
		gg = _lg_emit(gg, ((idx_v + " = load i64, i64* ") + _t(inst.Arg2)))
		s_p := ("%sca_spi_" + i2s(idx))
		_ = s_p
		gg = _lg_emit(gg, (((s_p + " = inttoptr i64 ") + s_ptr) + " to i8*"))
		srt := _str_ret_type()
		_ = srt
		res := ("%sca_res_" + i2s(idx))
		_ = res
		gg = _lg_emit(gg, (((((((((res + " = call ") + srt) + " @_aria_str_charAt(i8* ") + s_p) + ", i64 ") + s_len) + ", i64 ") + idx_v) + ")"))
		gg = _extract_str_result(gg, res, inst.Dest, idx, "sca")
		return gg
	}
	if (op == "StrSubstring") {
		gg := g
		_ = gg
		s_ptr := ("%ssub_sp_" + i2s(idx))
		_ = s_ptr
		s_len := ("%ssub_sl_" + i2s(idx))
		_ = s_len
		gg = _lg_emit(gg, ((s_ptr + " = load i64, i64* ") + _t(inst.Arg1)))
		gg = _lg_emit(gg, ((s_len + " = load i64, i64* ") + _t((inst.Arg1 + int64(1)))))
		start_v := ("%ssub_st_" + i2s(idx))
		_ = start_v
		gg = _lg_emit(gg, ((start_v + " = load i64, i64* ") + _t(inst.Arg2)))
		end_v := ("%ssub_en_" + i2s(idx))
		_ = end_v
		gg = _lg_emit(gg, ((end_v + " = load i64, i64* ") + _t(inst.Type_id)))
		s_p := ("%ssub_spi_" + i2s(idx))
		_ = s_p
		gg = _lg_emit(gg, (((s_p + " = inttoptr i64 ") + s_ptr) + " to i8*"))
		srt := _str_ret_type()
		_ = srt
		res := ("%ssub_res_" + i2s(idx))
		_ = res
		gg = _lg_emit(gg, (((((((((((res + " = call ") + srt) + " @_aria_str_substring(i8* ") + s_p) + ", i64 ") + s_len) + ", i64 ") + start_v) + ", i64 ") + end_v) + ")"))
		gg = _extract_str_result(gg, res, inst.Dest, idx, "ssub")
		return gg
	}
	if (op == "StrContains") {
		return _gen_str_binop_returning_i64(g, "@_aria_str_contains", inst, idx)
	}
	if (op == "StrStartsWith") {
		return _gen_str_binop_returning_i64(g, "@_aria_str_startsWith", inst, idx)
	}
	if (op == "StrEndsWith") {
		return _gen_str_binop_returning_i64(g, "@_aria_str_endsWith", inst, idx)
	}
	if (op == "StrIndexOf") {
		return _gen_str_binop_returning_i64(g, "@_aria_str_indexOf", inst, idx)
	}
	if (op == "StrTrim") {
		gg := g
		_ = gg
		s_ptr := ("%strim_sp_" + i2s(idx))
		_ = s_ptr
		s_len := ("%strim_sl_" + i2s(idx))
		_ = s_len
		gg = _lg_emit(gg, ((s_ptr + " = load i64, i64* ") + _t(inst.Arg1)))
		gg = _lg_emit(gg, ((s_len + " = load i64, i64* ") + _t((inst.Arg1 + int64(1)))))
		s_p := ("%strim_spi_" + i2s(idx))
		_ = s_p
		gg = _lg_emit(gg, (((s_p + " = inttoptr i64 ") + s_ptr) + " to i8*"))
		srt := _str_ret_type()
		_ = srt
		res := ("%strim_res_" + i2s(idx))
		_ = res
		gg = _lg_emit(gg, (((((((res + " = call ") + srt) + " @_aria_str_trim(i8* ") + s_p) + ", i64 ") + s_len) + ")"))
		gg = _extract_str_result(gg, res, inst.Dest, idx, "strim")
		return gg
	}
	if (op == "StrReplace") {
		gg := g
		_ = gg
		s_ptr := ("%srep_sp_" + i2s(idx))
		_ = s_ptr
		s_len := ("%srep_sl_" + i2s(idx))
		_ = s_len
		gg = _lg_emit(gg, ((s_ptr + " = load i64, i64* ") + _t(inst.Arg1)))
		gg = _lg_emit(gg, ((s_len + " = load i64, i64* ") + _t((inst.Arg1 + int64(1)))))
		o_ptr := ("%srep_op_" + i2s(idx))
		_ = o_ptr
		o_len := ("%srep_ol_" + i2s(idx))
		_ = o_len
		gg = _lg_emit(gg, ((o_ptr + " = load i64, i64* ") + _t(inst.Arg2)))
		gg = _lg_emit(gg, ((o_len + " = load i64, i64* ") + _t((inst.Arg2 + int64(1)))))
		n_ptr := ("%srep_np_" + i2s(idx))
		_ = n_ptr
		n_len := ("%srep_nl_" + i2s(idx))
		_ = n_len
		gg = _lg_emit(gg, ((n_ptr + " = load i64, i64* ") + _t(inst.Type_id)))
		gg = _lg_emit(gg, ((n_len + " = load i64, i64* ") + _t((inst.Type_id + int64(1)))))
		sp := ("%srep_spi_" + i2s(idx))
		_ = sp
		op2 := ("%srep_opi_" + i2s(idx))
		_ = op2
		np := ("%srep_npi_" + i2s(idx))
		_ = np
		gg = _lg_emit(gg, (((sp + " = inttoptr i64 ") + s_ptr) + " to i8*"))
		gg = _lg_emit(gg, (((op2 + " = inttoptr i64 ") + o_ptr) + " to i8*"))
		gg = _lg_emit(gg, (((np + " = inttoptr i64 ") + n_ptr) + " to i8*"))
		srt := _str_ret_type()
		_ = srt
		res := ("%srep_res_" + i2s(idx))
		_ = res
		gg = _lg_emit(gg, (((((((((((((((res + " = call ") + srt) + " @_aria_str_replace(i8* ") + sp) + ", i64 ") + s_len) + ", i8* ") + op2) + ", i64 ") + o_len) + ", i8* ") + np) + ", i64 ") + n_len) + ")"))
		gg = _extract_str_result(gg, res, inst.Dest, idx, "srep")
		return gg
	}
	if (op == "StrToLower") {
		gg := g
		_ = gg
		s_ptr := ("%slwr_sp_" + i2s(idx))
		_ = s_ptr
		s_len := ("%slwr_sl_" + i2s(idx))
		_ = s_len
		gg = _lg_emit(gg, ((s_ptr + " = load i64, i64* ") + _t(inst.Arg1)))
		gg = _lg_emit(gg, ((s_len + " = load i64, i64* ") + _t((inst.Arg1 + int64(1)))))
		s_p := ("%slwr_spi_" + i2s(idx))
		_ = s_p
		gg = _lg_emit(gg, (((s_p + " = inttoptr i64 ") + s_ptr) + " to i8*"))
		srt := _str_ret_type()
		_ = srt
		res := ("%slwr_res_" + i2s(idx))
		_ = res
		gg = _lg_emit(gg, (((((((res + " = call ") + srt) + " @_aria_str_toLower(i8* ") + s_p) + ", i64 ") + s_len) + ")"))
		gg = _extract_str_result(gg, res, inst.Dest, idx, "slwr")
		return gg
	}
	if (op == "StrToUpper") {
		gg := g
		_ = gg
		s_ptr := ("%supr_sp_" + i2s(idx))
		_ = s_ptr
		s_len := ("%supr_sl_" + i2s(idx))
		_ = s_len
		gg = _lg_emit(gg, ((s_ptr + " = load i64, i64* ") + _t(inst.Arg1)))
		gg = _lg_emit(gg, ((s_len + " = load i64, i64* ") + _t((inst.Arg1 + int64(1)))))
		s_p := ("%supr_spi_" + i2s(idx))
		_ = s_p
		gg = _lg_emit(gg, (((s_p + " = inttoptr i64 ") + s_ptr) + " to i8*"))
		srt := _str_ret_type()
		_ = srt
		res := ("%supr_res_" + i2s(idx))
		_ = res
		gg = _lg_emit(gg, (((((((res + " = call ") + srt) + " @_aria_str_toUpper(i8* ") + s_p) + ", i64 ") + s_len) + ")"))
		gg = _extract_str_result(gg, res, inst.Dest, idx, "supr")
		return gg
	}
	if (op == "StrSplit") {
		gg := g
		_ = gg
		s_ptr := ("%sspl_sp_" + i2s(idx))
		_ = s_ptr
		s_len := ("%sspl_sl_" + i2s(idx))
		_ = s_len
		gg = _lg_emit(gg, ((s_ptr + " = load i64, i64* ") + _t(inst.Arg1)))
		gg = _lg_emit(gg, ((s_len + " = load i64, i64* ") + _t((inst.Arg1 + int64(1)))))
		d_ptr := ("%sspl_dp_" + i2s(idx))
		_ = d_ptr
		d_len := ("%sspl_dl_" + i2s(idx))
		_ = d_len
		gg = _lg_emit(gg, ((d_ptr + " = load i64, i64* ") + _t(inst.Arg2)))
		gg = _lg_emit(gg, ((d_len + " = load i64, i64* ") + _t((inst.Arg2 + int64(1)))))
		sp := ("%sspl_spi_" + i2s(idx))
		_ = sp
		dp := ("%sspl_dpi_" + i2s(idx))
		_ = dp
		gg = _lg_emit(gg, (((sp + " = inttoptr i64 ") + s_ptr) + " to i8*"))
		gg = _lg_emit(gg, (((dp + " = inttoptr i64 ") + d_ptr) + " to i8*"))
		res := ("%sspl_res_" + i2s(idx))
		_ = res
		gg = _lg_emit(gg, (((((((((res + " = call i64 @_aria_str_split(i8* ") + sp) + ", i64 ") + s_len) + ", i8* ") + dp) + ", i64 ") + d_len) + ")"))
		gg = _lg_emit(gg, ((("store i64 " + res) + ", i64* ") + _t(inst.Dest)))
		return gg
	}
	if (op == "StrToInt") {
		gg := g
		_ = gg
		ptr_v := ("%sti_p_" + i2s(idx))
		_ = ptr_v
		len_v := ("%sti_l_" + i2s(idx))
		_ = len_v
		gg = _lg_emit(gg, ((ptr_v + " = load i64, i64* ") + _t(inst.Arg1)))
		gg = _lg_emit(gg, ((len_v + " = load i64, i64* ") + _t((inst.Arg1 + int64(1)))))
		ptr_cast := ("%sti_pc_" + i2s(idx))
		_ = ptr_cast
		gg = _lg_emit(gg, (((ptr_cast + " = inttoptr i64 ") + ptr_v) + " to i8*"))
		srt := _str_ret_type()
		_ = srt
		res := ("%sti_res_" + i2s(idx))
		_ = res
		gg = _lg_emit(gg, (((((((res + " = call ") + srt) + " @_aria_str_to_int(i8* ") + ptr_cast) + ", i64 ") + len_v) + ")"))
		val := ("%sti_val_" + i2s(idx))
		_ = val
		suc := ("%sti_suc_" + i2s(idx))
		_ = suc
		gg = _lg_emit(gg, (((((val + " = extractvalue ") + srt) + " ") + res) + ", 0"))
		gg = _lg_emit(gg, (((((suc + " = extractvalue ") + srt) + " ") + res) + ", 1"))
		val_int := ("%sti_vi_" + i2s(idx))
		_ = val_int
		gg = _lg_emit(gg, (((val_int + " = ptrtoint i8* ") + val) + " to i64"))
		gg = _lg_emit(gg, ((("store i64 " + val_int) + ", i64* ") + _t(inst.Dest)))
		dest_suc := (inst.Dest + int64(1))
		_ = dest_suc
		gg = _lg_emit(gg, ((("store i64 " + suc) + ", i64* ") + _t(dest_suc)))
		return gg
	}
	if (op == "MapNew") {
		gg := g
		_ = gg
		cap_v := ("%mnew_cap_" + i2s(idx))
		_ = cap_v
		gg = _lg_emit(gg, ((cap_v + " = load i64, i64* ") + _t(inst.Arg1)))
		res := ("%mnew_" + i2s(idx))
		_ = res
		gg = _lg_emit(gg, (((res + " = call i64 @_aria_map_new(i64 ") + cap_v) + ")"))
		gg = _lg_emit(gg, ((("store i64 " + res) + ", i64* ") + _t(inst.Dest)))
		gg = _gc_frame_update(gg, f, inst.Dest, res, idx)
		return gg
	}
	if (op == "MapSet") {
		gg := g
		_ = gg
		map_v := ("%mset_m_" + i2s(idx))
		_ = map_v
		gg = _lg_emit(gg, ((map_v + " = load i64, i64* ") + _t(inst.Arg1)))
		key_p := ("%mset_kp_" + i2s(idx))
		_ = key_p
		key_l := ("%mset_kl_" + i2s(idx))
		_ = key_l
		gg = _lg_emit(gg, ((key_p + " = load i64, i64* ") + _t(inst.Arg2)))
		gg = _lg_emit(gg, ((key_l + " = load i64, i64* ") + _t((inst.Arg2 + int64(1)))))
		key_pc := ("%mset_kpc_" + i2s(idx))
		_ = key_pc
		gg = _lg_emit(gg, (((key_pc + " = inttoptr i64 ") + key_p) + " to i8*"))
		val_v := ("%mset_v_" + i2s(idx))
		_ = val_v
		gg = _lg_emit(gg, ((val_v + " = load i64, i64* ") + _t(inst.Type_id)))
		gg = _lg_emit(gg, (((((((("call void @_aria_map_set(i64 " + map_v) + ", i8* ") + key_pc) + ", i64 ") + key_l) + ", i64 ") + val_v) + ")"))
		return gg
	}
	if (op == "MapGet") {
		gg := g
		_ = gg
		map_v := ("%mget_m_" + i2s(idx))
		_ = map_v
		gg = _lg_emit(gg, ((map_v + " = load i64, i64* ") + _t(inst.Arg1)))
		key_p := ("%mget_kp_" + i2s(idx))
		_ = key_p
		key_l := ("%mget_kl_" + i2s(idx))
		_ = key_l
		gg = _lg_emit(gg, ((key_p + " = load i64, i64* ") + _t(inst.Arg2)))
		gg = _lg_emit(gg, ((key_l + " = load i64, i64* ") + _t((inst.Arg2 + int64(1)))))
		key_pc := ("%mget_kpc_" + i2s(idx))
		_ = key_pc
		gg = _lg_emit(gg, (((key_pc + " = inttoptr i64 ") + key_p) + " to i8*"))
		res := ("%mget_" + i2s(idx))
		_ = res
		gg = _lg_emit(gg, (((((((res + " = call i64 @_aria_map_get(i64 ") + map_v) + ", i8* ") + key_pc) + ", i64 ") + key_l) + ")"))
		gg = _lg_emit(gg, ((("store i64 " + res) + ", i64* ") + _t(inst.Dest)))
		gg = _gc_frame_update(gg, f, inst.Dest, res, idx)
		return gg
	}
	if (op == "MapContains") {
		gg := g
		_ = gg
		map_v := ("%mhas_m_" + i2s(idx))
		_ = map_v
		gg = _lg_emit(gg, ((map_v + " = load i64, i64* ") + _t(inst.Arg1)))
		key_p := ("%mhas_kp_" + i2s(idx))
		_ = key_p
		key_l := ("%mhas_kl_" + i2s(idx))
		_ = key_l
		gg = _lg_emit(gg, ((key_p + " = load i64, i64* ") + _t(inst.Arg2)))
		gg = _lg_emit(gg, ((key_l + " = load i64, i64* ") + _t((inst.Arg2 + int64(1)))))
		key_pc := ("%mhas_kpc_" + i2s(idx))
		_ = key_pc
		gg = _lg_emit(gg, (((key_pc + " = inttoptr i64 ") + key_p) + " to i8*"))
		res := ("%mhas_" + i2s(idx))
		_ = res
		gg = _lg_emit(gg, (((((((res + " = call i64 @_aria_map_contains(i64 ") + map_v) + ", i8* ") + key_pc) + ", i64 ") + key_l) + ")"))
		gg = _lg_emit(gg, ((("store i64 " + res) + ", i64* ") + _t(inst.Dest)))
		return gg
	}
	if (op == "MapLen") {
		gg := g
		_ = gg
		map_v := ("%mlen_m_" + i2s(idx))
		_ = map_v
		gg = _lg_emit(gg, ((map_v + " = load i64, i64* ") + _t(inst.Arg1)))
		res := ("%mlen_" + i2s(idx))
		_ = res
		gg = _lg_emit(gg, (((res + " = call i64 @_aria_map_len(i64 ") + map_v) + ")"))
		gg = _lg_emit(gg, ((("store i64 " + res) + ", i64* ") + _t(inst.Dest)))
		return gg
	}
	if (op == "MapKeys") {
		gg := g
		_ = gg
		map_v := ("%mkeys_m_" + i2s(idx))
		_ = map_v
		gg = _lg_emit(gg, ((map_v + " = load i64, i64* ") + _t(inst.Arg1)))
		res := ("%mkeys_" + i2s(idx))
		_ = res
		gg = _lg_emit(gg, (((res + " = call i64 @_aria_map_keys(i64 ") + map_v) + ")"))
		gg = _lg_emit(gg, ((("store i64 " + res) + ", i64* ") + _t(inst.Dest)))
		gg = _gc_frame_update(gg, f, inst.Dest, res, idx)
		return gg
	}
	if (op == "SetNew") {
		gg := g
		_ = gg
		cap_v := ("%snew_cap_" + i2s(idx))
		_ = cap_v
		gg = _lg_emit(gg, ((cap_v + " = load i64, i64* ") + _t(inst.Arg1)))
		res := ("%snew_" + i2s(idx))
		_ = res
		gg = _lg_emit(gg, (((res + " = call i64 @_aria_set_new(i64 ") + cap_v) + ")"))
		gg = _lg_emit(gg, ((("store i64 " + res) + ", i64* ") + _t(inst.Dest)))
		gg = _gc_frame_update(gg, f, inst.Dest, res, idx)
		return gg
	}
	if (op == "SetAdd") {
		gg := g
		_ = gg
		set_v := ("%sadd_s_" + i2s(idx))
		_ = set_v
		gg = _lg_emit(gg, ((set_v + " = load i64, i64* ") + _t(inst.Arg1)))
		key_p := ("%sadd_kp_" + i2s(idx))
		_ = key_p
		key_l := ("%sadd_kl_" + i2s(idx))
		_ = key_l
		gg = _lg_emit(gg, ((key_p + " = load i64, i64* ") + _t(inst.Arg2)))
		gg = _lg_emit(gg, ((key_l + " = load i64, i64* ") + _t((inst.Arg2 + int64(1)))))
		key_pc := ("%sadd_kpc_" + i2s(idx))
		_ = key_pc
		gg = _lg_emit(gg, (((key_pc + " = inttoptr i64 ") + key_p) + " to i8*"))
		gg = _lg_emit(gg, (((((("call void @_aria_set_add(i64 " + set_v) + ", i8* ") + key_pc) + ", i64 ") + key_l) + ")"))
		return gg
	}
	if (op == "SetContains") {
		gg := g
		_ = gg
		set_v := ("%shas_s_" + i2s(idx))
		_ = set_v
		gg = _lg_emit(gg, ((set_v + " = load i64, i64* ") + _t(inst.Arg1)))
		key_p := ("%shas_kp_" + i2s(idx))
		_ = key_p
		key_l := ("%shas_kl_" + i2s(idx))
		_ = key_l
		gg = _lg_emit(gg, ((key_p + " = load i64, i64* ") + _t(inst.Arg2)))
		gg = _lg_emit(gg, ((key_l + " = load i64, i64* ") + _t((inst.Arg2 + int64(1)))))
		key_pc := ("%shas_kpc_" + i2s(idx))
		_ = key_pc
		gg = _lg_emit(gg, (((key_pc + " = inttoptr i64 ") + key_p) + " to i8*"))
		res := ("%shas_" + i2s(idx))
		_ = res
		gg = _lg_emit(gg, (((((((res + " = call i64 @_aria_set_contains(i64 ") + set_v) + ", i8* ") + key_pc) + ", i64 ") + key_l) + ")"))
		gg = _lg_emit(gg, ((("store i64 " + res) + ", i64* ") + _t(inst.Dest)))
		return gg
	}
	if (op == "SetLen") {
		gg := g
		_ = gg
		set_v := ("%slen_s_" + i2s(idx))
		_ = set_v
		gg = _lg_emit(gg, ((set_v + " = load i64, i64* ") + _t(inst.Arg1)))
		res := ("%slen_" + i2s(idx))
		_ = res
		gg = _lg_emit(gg, (((res + " = call i64 @_aria_set_len(i64 ") + set_v) + ")"))
		gg = _lg_emit(gg, ((("store i64 " + res) + ", i64* ") + _t(inst.Dest)))
		return gg
	}
	if (op == "SetValues") {
		gg := g
		_ = gg
		set_v := ("%svals_s_" + i2s(idx))
		_ = set_v
		gg = _lg_emit(gg, ((set_v + " = load i64, i64* ") + _t(inst.Arg1)))
		res := ("%svals_" + i2s(idx))
		_ = res
		gg = _lg_emit(gg, (((res + " = call i64 @_aria_set_values(i64 ") + set_v) + ")"))
		gg = _lg_emit(gg, ((("store i64 " + res) + ", i64* ") + _t(inst.Dest)))
		gg = _gc_frame_update(gg, f, inst.Dest, res, idx)
		return gg
	}
	if (op == "SetRemove") {
		gg := g
		_ = gg
		set_v := ("%srem_s_" + i2s(idx))
		_ = set_v
		gg = _lg_emit(gg, ((set_v + " = load i64, i64* ") + _t(inst.Arg1)))
		key_p := ("%srem_kp_" + i2s(idx))
		_ = key_p
		key_l := ("%srem_kl_" + i2s(idx))
		_ = key_l
		gg = _lg_emit(gg, ((key_p + " = load i64, i64* ") + _t(inst.Arg2)))
		gg = _lg_emit(gg, ((key_l + " = load i64, i64* ") + _t((inst.Arg2 + int64(1)))))
		key_pc := ("%srem_kpc_" + i2s(idx))
		_ = key_pc
		gg = _lg_emit(gg, (((key_pc + " = inttoptr i64 ") + key_p) + " to i8*"))
		gg = _lg_emit(gg, (((((("call void @_aria_set_remove(i64 " + set_v) + ", i8* ") + key_pc) + ", i64 ") + key_l) + ")"))
		return gg
	}
	if (op == "IntToStr") {
		gg := g
		_ = gg
		v := ("%its_v_" + i2s(idx))
		_ = v
		gg = _lg_emit(gg, ((v + " = load i64, i64* ") + _t(inst.Arg1)))
		srt := _str_ret_type()
		_ = srt
		res := ("%its_res_" + i2s(idx))
		_ = res
		if (inst.S1 == "f64") {
			gg = _lg_emit(gg, (((((res + " = call ") + srt) + " @_aria_float_to_str(i64 ") + v) + ")"))
		} else {
			gg = _lg_emit(gg, (((((res + " = call ") + srt) + " @_aria_int_to_str(i64 ") + v) + ")"))
		}
		gg = _extract_str_result(gg, res, inst.Dest, idx, "its")
		return gg
	}
	return g
}

func _str_ret_type() string {
	return ((_lb() + " i8*, i64 ") + _rb())
}

func _extract_str_result(g LlvmGen, res_name string, dest int64, idx int64, prefix string) LlvmGen {
	srt := _str_ret_type()
	_ = srt
	gg := g
	_ = gg
	r_ptr := ((("%" + prefix) + "_rp_") + i2s(idx))
	_ = r_ptr
	r_len := ((("%" + prefix) + "_rl_") + i2s(idx))
	_ = r_len
	gg = _lg_emit(gg, (((((r_ptr + " = extractvalue ") + srt) + " ") + res_name) + ", 0"))
	gg = _lg_emit(gg, (((((r_len + " = extractvalue ") + srt) + " ") + res_name) + ", 1"))
	rpi := ((("%" + prefix) + "_rpi_") + i2s(idx))
	_ = rpi
	gg = _lg_emit(gg, (((rpi + " = ptrtoint i8* ") + r_ptr) + " to i64"))
	gg = _lg_emit(gg, ((("store i64 " + rpi) + ", i64* ") + _t(dest)))
	gg = _lg_emit(gg, ((("store i64 " + r_len) + ", i64* ") + _t((dest + int64(1)))))
	return gg
}

func _gen_binop(g LlvmGen, llvm_op string, inst IrInst) LlvmGen {
	gg := g
	_ = gg
	a := ("%binop_a_" + i2s(inst.Dest))
	_ = a
	b := ("%binop_b_" + i2s(inst.Dest))
	_ = b
	gg = _lg_emit(gg, ((a + " = load i64, i64* ") + _t(inst.Arg1)))
	gg = _lg_emit(gg, ((b + " = load i64, i64* ") + _t(inst.Arg2)))
	res := ("%binop_r_" + i2s(inst.Dest))
	_ = res
	gg = _lg_emit(gg, ((((((res + " = ") + llvm_op) + " i64 ") + a) + ", ") + b))
	gg = _lg_emit(gg, ((("store i64 " + res) + ", i64* ") + _t(inst.Dest)))
	return gg
}

func _gen_cmp(g LlvmGen, cmp_op string, inst IrInst, idx int64) LlvmGen {
	gg := g
	_ = gg
	a := ("%cmp_a_" + i2s(idx))
	_ = a
	b := ("%cmp_b_" + i2s(idx))
	_ = b
	gg = _lg_emit(gg, ((a + " = load i64, i64* ") + _t(inst.Arg1)))
	gg = _lg_emit(gg, ((b + " = load i64, i64* ") + _t(inst.Arg2)))
	cmp := ("%cmp_c_" + i2s(idx))
	_ = cmp
	gg = _lg_emit(gg, ((((((cmp + " = icmp ") + cmp_op) + " i64 ") + a) + ", ") + b))
	res := ("%cmp_r_" + i2s(idx))
	_ = res
	gg = _lg_emit(gg, (((res + " = zext i1 ") + cmp) + " to i64"))
	gg = _lg_emit(gg, ((("store i64 " + res) + ", i64* ") + _t(inst.Dest)))
	return gg
}

func _gen_str_binop_returning_str(g LlvmGen, func_name string, inst IrInst, idx int64) LlvmGen {
	gg := g
	_ = gg
	a_ptr := ("%sbr_ap_" + i2s(idx))
	_ = a_ptr
	a_len := ("%sbr_al_" + i2s(idx))
	_ = a_len
	b_ptr := ("%sbr_bp_" + i2s(idx))
	_ = b_ptr
	b_len := ("%sbr_bl_" + i2s(idx))
	_ = b_len
	gg = _lg_emit(gg, ((a_ptr + " = load i64, i64* ") + _t(inst.Arg1)))
	gg = _lg_emit(gg, ((a_len + " = load i64, i64* ") + _t((inst.Arg1 + int64(1)))))
	gg = _lg_emit(gg, ((b_ptr + " = load i64, i64* ") + _t(inst.Arg2)))
	gg = _lg_emit(gg, ((b_len + " = load i64, i64* ") + _t((inst.Arg2 + int64(1)))))
	a_p := ("%sbr_api_" + i2s(idx))
	_ = a_p
	b_p := ("%sbr_bpi_" + i2s(idx))
	_ = b_p
	gg = _lg_emit(gg, (((a_p + " = inttoptr i64 ") + a_ptr) + " to i8*"))
	gg = _lg_emit(gg, (((b_p + " = inttoptr i64 ") + b_ptr) + " to i8*"))
	srt := _str_ret_type()
	_ = srt
	res := ("%sbr_res_" + i2s(idx))
	_ = res
	gg = _lg_emit(gg, (((((((((((((res + " = call ") + srt) + " ") + func_name) + "(i8* ") + a_p) + ", i64 ") + a_len) + ", i8* ") + b_p) + ", i64 ") + b_len) + ")"))
	gg = _extract_str_result(gg, res, inst.Dest, idx, "sbr")
	return gg
}

func _gen_str_binop_returning_i64(g LlvmGen, func_name string, inst IrInst, idx int64) LlvmGen {
	gg := g
	_ = gg
	a_ptr := ("%sbi_ap_" + i2s(idx))
	_ = a_ptr
	a_len := ("%sbi_al_" + i2s(idx))
	_ = a_len
	b_ptr := ("%sbi_bp_" + i2s(idx))
	_ = b_ptr
	b_len := ("%sbi_bl_" + i2s(idx))
	_ = b_len
	gg = _lg_emit(gg, ((a_ptr + " = load i64, i64* ") + _t(inst.Arg1)))
	gg = _lg_emit(gg, ((a_len + " = load i64, i64* ") + _t((inst.Arg1 + int64(1)))))
	gg = _lg_emit(gg, ((b_ptr + " = load i64, i64* ") + _t(inst.Arg2)))
	gg = _lg_emit(gg, ((b_len + " = load i64, i64* ") + _t((inst.Arg2 + int64(1)))))
	a_p := ("%sbi_api_" + i2s(idx))
	_ = a_p
	b_p := ("%sbi_bpi_" + i2s(idx))
	_ = b_p
	gg = _lg_emit(gg, (((a_p + " = inttoptr i64 ") + a_ptr) + " to i8*"))
	gg = _lg_emit(gg, (((b_p + " = inttoptr i64 ") + b_ptr) + " to i8*"))
	res := ("%sbi_res_" + i2s(idx))
	_ = res
	gg = _lg_emit(gg, (((((((((((res + " = call i64 ") + func_name) + "(i8* ") + a_p) + ", i64 ") + a_len) + ", i8* ") + b_p) + ", i64 ") + b_len) + ")"))
	gg = _lg_emit(gg, ((("store i64 " + res) + ", i64* ") + _t(inst.Dest)))
	return gg
}

func _gen_str_eq_inline(g LlvmGen, inst IrInst, idx int64) LlvmGen {
	gg := g
	_ = gg
	a_ptr := ("%seq_ap_" + i2s(idx))
	_ = a_ptr
	a_len := ("%seq_al_" + i2s(idx))
	_ = a_len
	b_ptr := ("%seq_bp_" + i2s(idx))
	_ = b_ptr
	b_len := ("%seq_bl_" + i2s(idx))
	_ = b_len
	gg = _lg_emit(gg, ((a_ptr + " = load i64, i64* ") + _t(inst.Arg1)))
	gg = _lg_emit(gg, ((a_len + " = load i64, i64* ") + _t((inst.Arg1 + int64(1)))))
	gg = _lg_emit(gg, ((b_ptr + " = load i64, i64* ") + _t(inst.Arg2)))
	gg = _lg_emit(gg, ((b_len + " = load i64, i64* ") + _t((inst.Arg2 + int64(1)))))
	len_eq := ("%seq_leq_" + i2s(idx))
	_ = len_eq
	gg = _lg_emit(gg, ((((len_eq + " = icmp eq i64 ") + a_len) + ", ") + b_len))
	lbl1 := _lg_fresh_label(gg)
	_ = lbl1
	gg = lbl1.G
	lbl2 := _lg_fresh_label(gg)
	_ = lbl2
	gg = lbl2.G
	lbl3 := _lg_fresh_label(gg)
	_ = lbl3
	gg = lbl3.G
	do_cmp_lbl := ("Lseq_cmp_" + i2s(lbl1.Id))
	_ = do_cmp_lbl
	not_eq_lbl := ("Lseq_ne_" + i2s(lbl2.Id))
	_ = not_eq_lbl
	done_lbl := ("Lseq_done_" + i2s(lbl3.Id))
	_ = done_lbl
	gg = _lg_emit_term(gg, ((((("br i1 " + len_eq) + ", label %") + do_cmp_lbl) + ", label %") + not_eq_lbl))
	gg = _lg_emit_label(gg, do_cmp_lbl)
	a_p := ("%seq_api_" + i2s(idx))
	_ = a_p
	b_p := ("%seq_bpi_" + i2s(idx))
	_ = b_p
	gg = _lg_emit(gg, (((a_p + " = inttoptr i64 ") + a_ptr) + " to i8*"))
	gg = _lg_emit(gg, (((b_p + " = inttoptr i64 ") + b_ptr) + " to i8*"))
	res_cmp := ("%seq_rc_" + i2s(idx))
	_ = res_cmp
	gg = _lg_emit(gg, (((((((((res_cmp + " = call i64 @_aria_str_eq(i8* ") + a_p) + ", i64 ") + a_len) + ", i8* ") + b_p) + ", i64 ") + b_len) + ")"))
	gg = _lg_emit_term(gg, ("br label %" + done_lbl))
	gg = _lg_emit_label(gg, not_eq_lbl)
	gg = _lg_emit_term(gg, ("br label %" + done_lbl))
	gg = _lg_emit_label(gg, done_lbl)
	res := ("%seq_res_" + i2s(idx))
	_ = res
	gg = _lg_emit(gg, (((((((res + " = phi i64 [ ") + res_cmp) + ", %") + do_cmp_lbl) + " ], [ 0, %") + not_eq_lbl) + " ]"))
	gg = _lg_emit(gg, ((("store i64 " + res) + ", i64* ") + _t(inst.Dest)))
	return gg
}

func _gen_call(g LlvmGen, m IrModule, f IrFunc, inst IrInst, idx int64) LlvmGen {
	orig_name := inst.S1
	_ = orig_name
	name := orig_name
	_ = name
	if (name == "main") {
		name = "_aria_main"
	}
	gg := g
	_ = gg
	if (((name == "_aria_println_str") || (name == "_aria_print_str")) || (name == "_aria_eprintln_str")) {
		ptr_v := ("%call_p_" + i2s(idx))
		_ = ptr_v
		len_v := ("%call_l_" + i2s(idx))
		_ = len_v
		gg = _lg_emit(gg, ((ptr_v + " = load i64, i64* ") + _t(inst.Arg1)))
		gg = _lg_emit(gg, ((len_v + " = load i64, i64* ") + _t((inst.Arg1 + int64(1)))))
		ptr_cast := ("%call_pc_" + i2s(idx))
		_ = ptr_cast
		gg = _lg_emit(gg, (((ptr_cast + " = inttoptr i64 ") + ptr_v) + " to i8*"))
		gg = _lg_emit(gg, (((((("call void @" + name) + "(i8* ") + ptr_cast) + ", i64 ") + len_v) + ")"))
		return gg
	}
	if (name == "_aria_exit") {
		v := ("%call_exit_" + i2s(idx))
		_ = v
		gg = _lg_emit(gg, ((v + " = load i64, i64* ") + _t(inst.Arg1)))
		gg = _lg_emit(gg, (("call void @_aria_exit(i64 " + v) + ")"))
		return gg
	}
	if (name == "_aria_panic") {
		ptr_v := ("%call_p_" + i2s(idx))
		_ = ptr_v
		len_v := ("%call_l_" + i2s(idx))
		_ = len_v
		gg = _lg_emit(gg, ((ptr_v + " = load i64, i64* ") + _t(inst.Arg1)))
		gg = _lg_emit(gg, ((len_v + " = load i64, i64* ") + _t((inst.Arg1 + int64(1)))))
		ptr_cast := ("%call_pc_" + i2s(idx))
		_ = ptr_cast
		gg = _lg_emit(gg, (((ptr_cast + " = inttoptr i64 ") + ptr_v) + " to i8*"))
		gg = _lg_emit(gg, (((("call void @_aria_panic(i8* " + ptr_cast) + ", i64 ") + len_v) + ")"))
		gg = _lg_emit_term(gg, "unreachable")
		return gg
	}
	if (name == "_aria_alloc") {
		v := ("%call_alloc_sz_" + i2s(idx))
		_ = v
		gg = _lg_emit(gg, ((v + " = load i64, i64* ") + _t(inst.Arg1)))
		raw := ("%call_alloc_r_" + i2s(idx))
		_ = raw
		gg = _lg_emit(gg, (((raw + " = call i8* @_aria_alloc(i64 ") + v) + ")"))
		res := ("%call_alloc_i_" + i2s(idx))
		_ = res
		gg = _lg_emit(gg, (((res + " = ptrtoint i8* ") + raw) + " to i64"))
		gg = _lg_emit(gg, ((("store i64 " + res) + ", i64* ") + _t(inst.Dest)))
		return gg
	}
	if (name == "_aria_read_file") {
		ptr_v := ("%call_rf_p_" + i2s(idx))
		_ = ptr_v
		len_v := ("%call_rf_l_" + i2s(idx))
		_ = len_v
		gg = _lg_emit(gg, ((ptr_v + " = load i64, i64* ") + _t(inst.Arg1)))
		gg = _lg_emit(gg, ((len_v + " = load i64, i64* ") + _t((inst.Arg1 + int64(1)))))
		ptr_cast := ("%call_rf_pc_" + i2s(idx))
		_ = ptr_cast
		gg = _lg_emit(gg, (((ptr_cast + " = inttoptr i64 ") + ptr_v) + " to i8*"))
		srt := _str_ret_type()
		_ = srt
		res := ("%call_rf_res_" + i2s(idx))
		_ = res
		gg = _lg_emit(gg, (((((((res + " = call ") + srt) + " @_aria_read_file(i8* ") + ptr_cast) + ", i64 ") + len_v) + ")"))
		gg = _extract_str_result(gg, res, inst.Dest, idx, "call_rf")
		return gg
	}
	if (name == "_aria_write_file") {
		pp := ("%call_wf_pp_" + i2s(idx))
		_ = pp
		pl := ("%call_wf_pl_" + i2s(idx))
		_ = pl
		cp := ("%call_wf_cp_" + i2s(idx))
		_ = cp
		cl := ("%call_wf_cl_" + i2s(idx))
		_ = cl
		gg = _lg_emit(gg, ((pp + " = load i64, i64* ") + _t(inst.Arg1)))
		gg = _lg_emit(gg, ((pl + " = load i64, i64* ") + _t((inst.Arg1 + int64(1)))))
		gg = _lg_emit(gg, ((cp + " = load i64, i64* ") + _t((inst.Arg1 + int64(2)))))
		gg = _lg_emit(gg, ((cl + " = load i64, i64* ") + _t((inst.Arg1 + int64(3)))))
		pp_c := ("%call_wf_ppc_" + i2s(idx))
		_ = pp_c
		cp_c := ("%call_wf_cpc_" + i2s(idx))
		_ = cp_c
		gg = _lg_emit(gg, (((pp_c + " = inttoptr i64 ") + pp) + " to i8*"))
		gg = _lg_emit(gg, (((cp_c + " = inttoptr i64 ") + cp) + " to i8*"))
		gg = _lg_emit(gg, (((((((("call void @_aria_write_file(i8* " + pp_c) + ", i64 ") + pl) + ", i8* ") + cp_c) + ", i64 ") + cl) + ")"))
		return gg
	}
	if (name == "_aria_append_file") {
		pp := ("%call_af_pp_" + i2s(idx))
		_ = pp
		pl := ("%call_af_pl_" + i2s(idx))
		_ = pl
		cp := ("%call_af_cp_" + i2s(idx))
		_ = cp
		cl := ("%call_af_cl_" + i2s(idx))
		_ = cl
		gg = _lg_emit(gg, ((pp + " = load i64, i64* ") + _t(inst.Arg1)))
		gg = _lg_emit(gg, ((pl + " = load i64, i64* ") + _t((inst.Arg1 + int64(1)))))
		gg = _lg_emit(gg, ((cp + " = load i64, i64* ") + _t((inst.Arg1 + int64(2)))))
		gg = _lg_emit(gg, ((cl + " = load i64, i64* ") + _t((inst.Arg1 + int64(3)))))
		pp_c := ("%call_af_ppc_" + i2s(idx))
		_ = pp_c
		cp_c := ("%call_af_cpc_" + i2s(idx))
		_ = cp_c
		gg = _lg_emit(gg, (((pp_c + " = inttoptr i64 ") + pp) + " to i8*"))
		gg = _lg_emit(gg, (((cp_c + " = inttoptr i64 ") + cp) + " to i8*"))
		gg = _lg_emit(gg, (((((((("call void @_aria_append_file(i8* " + pp_c) + ", i64 ") + pl) + ", i8* ") + cp_c) + ", i64 ") + cl) + ")"))
		return gg
	}
	if (name == "_aria_write_binary_file") {
		pp := ("%call_wbf_pp_" + i2s(idx))
		_ = pp
		pl := ("%call_wbf_pl_" + i2s(idx))
		_ = pl
		dp := ("%call_wbf_dp_" + i2s(idx))
		_ = dp
		dl := ("%call_wbf_dl_" + i2s(idx))
		_ = dl
		gg = _lg_emit(gg, ((pp + " = load i64, i64* ") + _t(inst.Arg1)))
		gg = _lg_emit(gg, ((pl + " = load i64, i64* ") + _t((inst.Arg1 + int64(1)))))
		gg = _lg_emit(gg, ((dp + " = load i64, i64* ") + _t((inst.Arg1 + int64(2)))))
		gg = _lg_emit(gg, ((dl + " = load i64, i64* ") + _t((inst.Arg1 + int64(3)))))
		pp_c := ("%call_wbf_ppc_" + i2s(idx))
		_ = pp_c
		dp_c := ("%call_wbf_dpc_" + i2s(idx))
		_ = dp_c
		gg = _lg_emit(gg, (((pp_c + " = inttoptr i64 ") + pp) + " to i8*"))
		gg = _lg_emit(gg, (((dp_c + " = inttoptr i64 ") + dp) + " to i64*"))
		gg = _lg_emit(gg, (((((((("call void @_aria_write_binary_file(i8* " + pp_c) + ", i64 ") + pl) + ", i64* ") + dp_c) + ", i64 ") + dl) + ")"))
		return gg
	}
	if (name == "_aria_args") {
		res := ("%call_args_" + i2s(idx))
		_ = res
		gg = _lg_emit(gg, (res + " = call i64 @_aria_args_get()"))
		gg = _lg_emit(gg, ((("store i64 " + res) + ", i64* ") + _t(inst.Dest)))
		return gg
	}
	if (name == "_aria_exec") {
		ptr_v := ("%call_exec_p_" + i2s(idx))
		_ = ptr_v
		len_v := ("%call_exec_l_" + i2s(idx))
		_ = len_v
		gg = _lg_emit(gg, ((ptr_v + " = load i64, i64* ") + _t(inst.Arg1)))
		gg = _lg_emit(gg, ((len_v + " = load i64, i64* ") + _t((inst.Arg1 + int64(1)))))
		ptr_cast := ("%call_exec_pc_" + i2s(idx))
		_ = ptr_cast
		gg = _lg_emit(gg, (((ptr_cast + " = inttoptr i64 ") + ptr_v) + " to i8*"))
		res := ("%call_exec_r_" + i2s(idx))
		_ = res
		gg = _lg_emit(gg, (((((res + " = call i64 @_aria_exec(i8* ") + ptr_cast) + ", i64 ") + len_v) + ")"))
		gg = _lg_emit(gg, ((("store i64 " + res) + ", i64* ") + _t(inst.Dest)))
		return gg
	}
	if (name == "_aria_getenv") {
		ptr_v := ("%call_ge_p_" + i2s(idx))
		_ = ptr_v
		len_v := ("%call_ge_l_" + i2s(idx))
		_ = len_v
		gg = _lg_emit(gg, ((ptr_v + " = load i64, i64* ") + _t(inst.Arg1)))
		gg = _lg_emit(gg, ((len_v + " = load i64, i64* ") + _t((inst.Arg1 + int64(1)))))
		ptr_cast := ("%call_ge_pc_" + i2s(idx))
		_ = ptr_cast
		gg = _lg_emit(gg, (((ptr_cast + " = inttoptr i64 ") + ptr_v) + " to i8*"))
		srt := _str_ret_type()
		_ = srt
		res := ("%call_ge_r_" + i2s(idx))
		_ = res
		gg = _lg_emit(gg, (((((((res + " = call ") + srt) + " @_aria_getenv(i8* ") + ptr_cast) + ", i64 ") + len_v) + ")"))
		gg = _extract_str_result(gg, res, inst.Dest, idx, "call_ge")
		return gg
	}
	if (name == "_aria_tcp_socket") {
		res := ("%call_ts_" + i2s(idx))
		_ = res
		gg = _lg_emit(gg, (res + " = call i64 @_aria_tcp_socket()"))
		gg = _lg_emit(gg, ((("store i64 " + res) + ", i64* ") + _t(inst.Dest)))
		return gg
	}
	if (name == "_aria_tcp_bind") {
		fd_v := ("%call_tb_fd_" + i2s(idx))
		_ = fd_v
		ap := ("%call_tb_ap_" + i2s(idx))
		_ = ap
		al := ("%call_tb_al_" + i2s(idx))
		_ = al
		port_v := ("%call_tb_port_" + i2s(idx))
		_ = port_v
		gg = _lg_emit(gg, ((fd_v + " = load i64, i64* ") + _t(inst.Arg1)))
		gg = _lg_emit(gg, ((ap + " = load i64, i64* ") + _t((inst.Arg1 + int64(1)))))
		gg = _lg_emit(gg, ((al + " = load i64, i64* ") + _t((inst.Arg1 + int64(2)))))
		gg = _lg_emit(gg, ((port_v + " = load i64, i64* ") + _t((inst.Arg1 + int64(3)))))
		apc := ("%call_tb_apc_" + i2s(idx))
		_ = apc
		gg = _lg_emit(gg, (((apc + " = inttoptr i64 ") + ap) + " to i8*"))
		res := ("%call_tb_r_" + i2s(idx))
		_ = res
		gg = _lg_emit(gg, (((((((((res + " = call i64 @_aria_tcp_bind(i64 ") + fd_v) + ", i8* ") + apc) + ", i64 ") + al) + ", i64 ") + port_v) + ")"))
		gg = _lg_emit(gg, ((("store i64 " + res) + ", i64* ") + _t(inst.Dest)))
		return gg
	}
	if (name == "_aria_tcp_listen") {
		fd_v := ("%call_tl_fd_" + i2s(idx))
		_ = fd_v
		bl_v := ("%call_tl_bl_" + i2s(idx))
		_ = bl_v
		gg = _lg_emit(gg, ((fd_v + " = load i64, i64* ") + _t(inst.Arg1)))
		gg = _lg_emit(gg, ((bl_v + " = load i64, i64* ") + _t((inst.Arg1 + int64(1)))))
		res := ("%call_tl_r_" + i2s(idx))
		_ = res
		gg = _lg_emit(gg, (((((res + " = call i64 @_aria_tcp_listen(i64 ") + fd_v) + ", i64 ") + bl_v) + ")"))
		gg = _lg_emit(gg, ((("store i64 " + res) + ", i64* ") + _t(inst.Dest)))
		return gg
	}
	if (name == "_aria_tcp_accept") {
		fd_v := ("%call_ta_fd_" + i2s(idx))
		_ = fd_v
		gg = _lg_emit(gg, ((fd_v + " = load i64, i64* ") + _t(inst.Arg1)))
		res := ("%call_ta_r_" + i2s(idx))
		_ = res
		gg = _lg_emit(gg, (((res + " = call i64 @_aria_tcp_accept(i64 ") + fd_v) + ")"))
		gg = _lg_emit(gg, ((("store i64 " + res) + ", i64* ") + _t(inst.Dest)))
		return gg
	}
	if (name == "_aria_tcp_read") {
		fd_v := ("%call_tr_fd_" + i2s(idx))
		_ = fd_v
		ml_v := ("%call_tr_ml_" + i2s(idx))
		_ = ml_v
		gg = _lg_emit(gg, ((fd_v + " = load i64, i64* ") + _t(inst.Arg1)))
		gg = _lg_emit(gg, ((ml_v + " = load i64, i64* ") + _t((inst.Arg1 + int64(1)))))
		srt := _str_ret_type()
		_ = srt
		res := ("%call_tr_r_" + i2s(idx))
		_ = res
		gg = _lg_emit(gg, (((((((res + " = call ") + srt) + " @_aria_tcp_read(i64 ") + fd_v) + ", i64 ") + ml_v) + ")"))
		gg = _extract_str_result(gg, res, inst.Dest, idx, "call_tr")
		return gg
	}
	if (name == "_aria_tcp_write") {
		fd_v := ("%call_tw_fd_" + i2s(idx))
		_ = fd_v
		ptr_v := ("%call_tw_p_" + i2s(idx))
		_ = ptr_v
		len_v := ("%call_tw_l_" + i2s(idx))
		_ = len_v
		gg = _lg_emit(gg, ((fd_v + " = load i64, i64* ") + _t(inst.Arg1)))
		gg = _lg_emit(gg, ((ptr_v + " = load i64, i64* ") + _t((inst.Arg1 + int64(1)))))
		gg = _lg_emit(gg, ((len_v + " = load i64, i64* ") + _t((inst.Arg1 + int64(2)))))
		ptrc := ("%call_tw_pc_" + i2s(idx))
		_ = ptrc
		gg = _lg_emit(gg, (((ptrc + " = inttoptr i64 ") + ptr_v) + " to i8*"))
		res := ("%call_tw_r_" + i2s(idx))
		_ = res
		gg = _lg_emit(gg, (((((((res + " = call i64 @_aria_tcp_write(i64 ") + fd_v) + ", i8* ") + ptrc) + ", i64 ") + len_v) + ")"))
		gg = _lg_emit(gg, ((("store i64 " + res) + ", i64* ") + _t(inst.Dest)))
		return gg
	}
	if (name == "_aria_tcp_close") {
		fd_v := ("%call_tc_fd_" + i2s(idx))
		_ = fd_v
		gg = _lg_emit(gg, ((fd_v + " = load i64, i64* ") + _t(inst.Arg1)))
		gg = _lg_emit(gg, (("call void @_aria_tcp_close(i64 " + fd_v) + ")"))
		return gg
	}
	if (name == "_aria_tcp_peer_addr") {
		fd_v := ("%call_tpa_fd_" + i2s(idx))
		_ = fd_v
		gg = _lg_emit(gg, ((fd_v + " = load i64, i64* ") + _t(inst.Arg1)))
		srt := _str_ret_type()
		_ = srt
		res := ("%call_tpa_r_" + i2s(idx))
		_ = res
		gg = _lg_emit(gg, (((((res + " = call ") + srt) + " @_aria_tcp_peer_addr(i64 ") + fd_v) + ")"))
		gg = _extract_str_result(gg, res, inst.Dest, idx, "call_tpa")
		return gg
	}
	if (name == "_aria_tcp_set_timeout") {
		fd_v := ("%call_tst_fd_" + i2s(idx))
		_ = fd_v
		k_v := ("%call_tst_k_" + i2s(idx))
		_ = k_v
		ms_v := ("%call_tst_ms_" + i2s(idx))
		_ = ms_v
		gg = _lg_emit(gg, ((fd_v + " = load i64, i64* ") + _t(inst.Arg1)))
		gg = _lg_emit(gg, ((k_v + " = load i64, i64* ") + _t((inst.Arg1 + int64(1)))))
		gg = _lg_emit(gg, ((ms_v + " = load i64, i64* ") + _t((inst.Arg1 + int64(2)))))
		res := ("%call_tst_r_" + i2s(idx))
		_ = res
		gg = _lg_emit(gg, (((((((res + " = call i64 @_aria_tcp_set_timeout(i64 ") + fd_v) + ", i64 ") + k_v) + ", i64 ") + ms_v) + ")"))
		gg = _lg_emit(gg, ((("store i64 " + res) + ", i64* ") + _t(inst.Dest)))
		return gg
	}
	if (name == "_aria_pg_connect") {
		ptr_v := ("%call_pgc_p_" + i2s(idx))
		_ = ptr_v
		len_v := ("%call_pgc_l_" + i2s(idx))
		_ = len_v
		gg = _lg_emit(gg, ((ptr_v + " = load i64, i64* ") + _t(inst.Arg1)))
		gg = _lg_emit(gg, ((len_v + " = load i64, i64* ") + _t((inst.Arg1 + int64(1)))))
		ptrc := ("%call_pgc_pc_" + i2s(idx))
		_ = ptrc
		gg = _lg_emit(gg, (((ptrc + " = inttoptr i64 ") + ptr_v) + " to i8*"))
		res := ("%call_pgc_r_" + i2s(idx))
		_ = res
		gg = _lg_emit(gg, (((((res + " = call i64 @_aria_pg_connect(i8* ") + ptrc) + ", i64 ") + len_v) + ")"))
		gg = _lg_emit(gg, ((("store i64 " + res) + ", i64* ") + _t(inst.Dest)))
		return gg
	}
	if (name == "_aria_pg_close") {
		h_v := ("%call_pgcl_h_" + i2s(idx))
		_ = h_v
		gg = _lg_emit(gg, ((h_v + " = load i64, i64* ") + _t(inst.Arg1)))
		gg = _lg_emit(gg, (("call void @_aria_pg_close(i64 " + h_v) + ")"))
		return gg
	}
	if (name == "_aria_pg_status") {
		h_v := ("%call_pgs_h_" + i2s(idx))
		_ = h_v
		gg = _lg_emit(gg, ((h_v + " = load i64, i64* ") + _t(inst.Arg1)))
		res := ("%call_pgs_r_" + i2s(idx))
		_ = res
		gg = _lg_emit(gg, (((res + " = call i64 @_aria_pg_status(i64 ") + h_v) + ")"))
		gg = _lg_emit(gg, ((("store i64 " + res) + ", i64* ") + _t(inst.Dest)))
		return gg
	}
	if (name == "_aria_pg_error") {
		h_v := ("%call_pge_h_" + i2s(idx))
		_ = h_v
		gg = _lg_emit(gg, ((h_v + " = load i64, i64* ") + _t(inst.Arg1)))
		srt := _str_ret_type()
		_ = srt
		res := ("%call_pge_r_" + i2s(idx))
		_ = res
		gg = _lg_emit(gg, (((((res + " = call ") + srt) + " @_aria_pg_error(i64 ") + h_v) + ")"))
		gg = _extract_str_result(gg, res, inst.Dest, idx, "call_pge")
		return gg
	}
	if (name == "_aria_pg_exec") {
		conn_v := ("%call_pgex_c_" + i2s(idx))
		_ = conn_v
		ptr_v := ("%call_pgex_p_" + i2s(idx))
		_ = ptr_v
		len_v := ("%call_pgex_l_" + i2s(idx))
		_ = len_v
		gg = _lg_emit(gg, ((conn_v + " = load i64, i64* ") + _t(inst.Arg1)))
		gg = _lg_emit(gg, ((ptr_v + " = load i64, i64* ") + _t((inst.Arg1 + int64(1)))))
		gg = _lg_emit(gg, ((len_v + " = load i64, i64* ") + _t((inst.Arg1 + int64(2)))))
		ptrc := ("%call_pgex_pc_" + i2s(idx))
		_ = ptrc
		gg = _lg_emit(gg, (((ptrc + " = inttoptr i64 ") + ptr_v) + " to i8*"))
		res := ("%call_pgex_r_" + i2s(idx))
		_ = res
		gg = _lg_emit(gg, (((((((res + " = call i64 @_aria_pg_exec(i64 ") + conn_v) + ", i8* ") + ptrc) + ", i64 ") + len_v) + ")"))
		gg = _lg_emit(gg, ((("store i64 " + res) + ", i64* ") + _t(inst.Dest)))
		return gg
	}
	if (name == "_aria_pg_exec_params") {
		conn_v := ("%call_pgep_c_" + i2s(idx))
		_ = conn_v
		ptr_v := ("%call_pgep_p_" + i2s(idx))
		_ = ptr_v
		len_v := ("%call_pgep_l_" + i2s(idx))
		_ = len_v
		arr_v := ("%call_pgep_a_" + i2s(idx))
		_ = arr_v
		gg = _lg_emit(gg, ((conn_v + " = load i64, i64* ") + _t(inst.Arg1)))
		gg = _lg_emit(gg, ((ptr_v + " = load i64, i64* ") + _t((inst.Arg1 + int64(1)))))
		gg = _lg_emit(gg, ((len_v + " = load i64, i64* ") + _t((inst.Arg1 + int64(2)))))
		gg = _lg_emit(gg, ((arr_v + " = load i64, i64* ") + _t((inst.Arg1 + int64(3)))))
		ptrc := ("%call_pgep_pc_" + i2s(idx))
		_ = ptrc
		gg = _lg_emit(gg, (((ptrc + " = inttoptr i64 ") + ptr_v) + " to i8*"))
		res := ("%call_pgep_r_" + i2s(idx))
		_ = res
		gg = _lg_emit(gg, (((((((((res + " = call i64 @_aria_pg_exec_params(i64 ") + conn_v) + ", i8* ") + ptrc) + ", i64 ") + len_v) + ", i64 ") + arr_v) + ")"))
		gg = _lg_emit(gg, ((("store i64 " + res) + ", i64* ") + _t(inst.Dest)))
		return gg
	}
	if (name == "_aria_pg_result_status") {
		h_v := ("%call_pgrs_h_" + i2s(idx))
		_ = h_v
		gg = _lg_emit(gg, ((h_v + " = load i64, i64* ") + _t(inst.Arg1)))
		res := ("%call_pgrs_r_" + i2s(idx))
		_ = res
		gg = _lg_emit(gg, (((res + " = call i64 @_aria_pg_result_status(i64 ") + h_v) + ")"))
		gg = _lg_emit(gg, ((("store i64 " + res) + ", i64* ") + _t(inst.Dest)))
		return gg
	}
	if (name == "_aria_pg_result_error") {
		h_v := ("%call_pgre_h_" + i2s(idx))
		_ = h_v
		gg = _lg_emit(gg, ((h_v + " = load i64, i64* ") + _t(inst.Arg1)))
		srt := _str_ret_type()
		_ = srt
		res := ("%call_pgre_r_" + i2s(idx))
		_ = res
		gg = _lg_emit(gg, (((((res + " = call ") + srt) + " @_aria_pg_result_error(i64 ") + h_v) + ")"))
		gg = _extract_str_result(gg, res, inst.Dest, idx, "call_pgre")
		return gg
	}
	if (name == "_aria_pg_nrows") {
		h_v := ("%call_pgnr_h_" + i2s(idx))
		_ = h_v
		gg = _lg_emit(gg, ((h_v + " = load i64, i64* ") + _t(inst.Arg1)))
		res := ("%call_pgnr_r_" + i2s(idx))
		_ = res
		gg = _lg_emit(gg, (((res + " = call i64 @_aria_pg_nrows(i64 ") + h_v) + ")"))
		gg = _lg_emit(gg, ((("store i64 " + res) + ", i64* ") + _t(inst.Dest)))
		return gg
	}
	if (name == "_aria_pg_ncols") {
		h_v := ("%call_pgnc_h_" + i2s(idx))
		_ = h_v
		gg = _lg_emit(gg, ((h_v + " = load i64, i64* ") + _t(inst.Arg1)))
		res := ("%call_pgnc_r_" + i2s(idx))
		_ = res
		gg = _lg_emit(gg, (((res + " = call i64 @_aria_pg_ncols(i64 ") + h_v) + ")"))
		gg = _lg_emit(gg, ((("store i64 " + res) + ", i64* ") + _t(inst.Dest)))
		return gg
	}
	if (name == "_aria_pg_field_name") {
		h_v := ("%call_pgfn_h_" + i2s(idx))
		_ = h_v
		c_v := ("%call_pgfn_c_" + i2s(idx))
		_ = c_v
		gg = _lg_emit(gg, ((h_v + " = load i64, i64* ") + _t(inst.Arg1)))
		gg = _lg_emit(gg, ((c_v + " = load i64, i64* ") + _t((inst.Arg1 + int64(1)))))
		srt := _str_ret_type()
		_ = srt
		res := ("%call_pgfn_r_" + i2s(idx))
		_ = res
		gg = _lg_emit(gg, (((((((res + " = call ") + srt) + " @_aria_pg_field_name(i64 ") + h_v) + ", i64 ") + c_v) + ")"))
		gg = _extract_str_result(gg, res, inst.Dest, idx, "call_pgfn")
		return gg
	}
	if (name == "_aria_pg_get_value") {
		h_v := ("%call_pggv_h_" + i2s(idx))
		_ = h_v
		r_v := ("%call_pggv_r_" + i2s(idx))
		_ = r_v
		c_v := ("%call_pggv_c_" + i2s(idx))
		_ = c_v
		gg = _lg_emit(gg, ((h_v + " = load i64, i64* ") + _t(inst.Arg1)))
		gg = _lg_emit(gg, ((r_v + " = load i64, i64* ") + _t((inst.Arg1 + int64(1)))))
		gg = _lg_emit(gg, ((c_v + " = load i64, i64* ") + _t((inst.Arg1 + int64(2)))))
		srt := _str_ret_type()
		_ = srt
		res := ("%call_pggv_res_" + i2s(idx))
		_ = res
		gg = _lg_emit(gg, (((((((((res + " = call ") + srt) + " @_aria_pg_get_value(i64 ") + h_v) + ", i64 ") + r_v) + ", i64 ") + c_v) + ")"))
		gg = _extract_str_result(gg, res, inst.Dest, idx, "call_pggv")
		return gg
	}
	if (name == "_aria_pg_is_null") {
		h_v := ("%call_pgin_h_" + i2s(idx))
		_ = h_v
		r_v := ("%call_pgin_r_" + i2s(idx))
		_ = r_v
		c_v := ("%call_pgin_c_" + i2s(idx))
		_ = c_v
		gg = _lg_emit(gg, ((h_v + " = load i64, i64* ") + _t(inst.Arg1)))
		gg = _lg_emit(gg, ((r_v + " = load i64, i64* ") + _t((inst.Arg1 + int64(1)))))
		gg = _lg_emit(gg, ((c_v + " = load i64, i64* ") + _t((inst.Arg1 + int64(2)))))
		res := ("%call_pgin_res_" + i2s(idx))
		_ = res
		gg = _lg_emit(gg, (((((((res + " = call i64 @_aria_pg_is_null(i64 ") + h_v) + ", i64 ") + r_v) + ", i64 ") + c_v) + ")"))
		gg = _lg_emit(gg, ((("store i64 " + res) + ", i64* ") + _t(inst.Dest)))
		return gg
	}
	if (name == "_aria_pg_clear") {
		h_v := ("%call_pgcr_h_" + i2s(idx))
		_ = h_v
		gg = _lg_emit(gg, ((h_v + " = load i64, i64* ") + _t(inst.Arg1)))
		gg = _lg_emit(gg, (("call void @_aria_pg_clear(i64 " + h_v) + ")"))
		return gg
	}
	if (name == "_aria_spawn") {
		fp_v := ("%call_sp_fp_" + i2s(idx))
		_ = fp_v
		ep_v := ("%call_sp_ep_" + i2s(idx))
		_ = ep_v
		gg = _lg_emit(gg, ((fp_v + " = load i64, i64* ") + _t(inst.Arg1)))
		gg = _lg_emit(gg, ((ep_v + " = load i64, i64* ") + _t((inst.Arg1 + int64(1)))))
		res := ("%call_sp_r_" + i2s(idx))
		_ = res
		gg = _lg_emit(gg, (((((res + " = call i64 @_aria_spawn(i64 ") + fp_v) + ", i64 ") + ep_v) + ")"))
		gg = _lg_emit(gg, ((("store i64 " + res) + ", i64* ") + _t(inst.Dest)))
		return gg
	}
	if (name == "_aria_task_await") {
		h_v := ("%call_ta_h_" + i2s(idx))
		_ = h_v
		gg = _lg_emit(gg, ((h_v + " = load i64, i64* ") + _t(inst.Arg1)))
		res := ("%call_ta_r_" + i2s(idx))
		_ = res
		gg = _lg_emit(gg, (((res + " = call i64 @_aria_task_await(i64 ") + h_v) + ")"))
		gg = _lg_emit(gg, ((("store i64 " + res) + ", i64* ") + _t(inst.Dest)))
		return gg
	}
	if (name == "_aria_chan_new") {
		cap_v := ("%call_cn_c_" + i2s(idx))
		_ = cap_v
		gg = _lg_emit(gg, ((cap_v + " = load i64, i64* ") + _t(inst.Arg1)))
		res := ("%call_cn_r_" + i2s(idx))
		_ = res
		gg = _lg_emit(gg, (((res + " = call i64 @_aria_chan_new(i64 ") + cap_v) + ")"))
		gg = _lg_emit(gg, ((("store i64 " + res) + ", i64* ") + _t(inst.Dest)))
		return gg
	}
	if (name == "_aria_chan_send") {
		ch_v := ("%call_cs_ch_" + i2s(idx))
		_ = ch_v
		val_v := ("%call_cs_v_" + i2s(idx))
		_ = val_v
		gg = _lg_emit(gg, ((ch_v + " = load i64, i64* ") + _t(inst.Arg1)))
		gg = _lg_emit(gg, ((val_v + " = load i64, i64* ") + _t((inst.Arg1 + int64(1)))))
		res := ("%call_cs_r_" + i2s(idx))
		_ = res
		gg = _lg_emit(gg, (((((res + " = call i64 @_aria_chan_send(i64 ") + ch_v) + ", i64 ") + val_v) + ")"))
		gg = _lg_emit(gg, ((("store i64 " + res) + ", i64* ") + _t(inst.Dest)))
		return gg
	}
	if (name == "_aria_chan_recv") {
		ch_v := ("%call_cr_ch_" + i2s(idx))
		_ = ch_v
		gg = _lg_emit(gg, ((ch_v + " = load i64, i64* ") + _t(inst.Arg1)))
		res := ("%call_cr_r_" + i2s(idx))
		_ = res
		gg = _lg_emit(gg, (((res + " = call i64 @_aria_chan_recv(i64 ") + ch_v) + ")"))
		gg = _lg_emit(gg, ((("store i64 " + res) + ", i64* ") + _t(inst.Dest)))
		return gg
	}
	if (name == "_aria_chan_close") {
		ch_v := ("%call_cc_ch_" + i2s(idx))
		_ = ch_v
		gg = _lg_emit(gg, ((ch_v + " = load i64, i64* ") + _t(inst.Arg1)))
		gg = _lg_emit(gg, (("call void @_aria_chan_close(i64 " + ch_v) + ")"))
		return gg
	}
	if (name == "_aria_mutex_new") {
		res := ("%call_mn_r_" + i2s(idx))
		_ = res
		gg = _lg_emit(gg, (res + " = call i64 @_aria_mutex_new()"))
		gg = _lg_emit(gg, ((("store i64 " + res) + ", i64* ") + _t(inst.Dest)))
		return gg
	}
	if (name == "_aria_mutex_lock") {
		m_v := ("%call_ml_m_" + i2s(idx))
		_ = m_v
		gg = _lg_emit(gg, ((m_v + " = load i64, i64* ") + _t(inst.Arg1)))
		gg = _lg_emit(gg, (("call void @_aria_mutex_lock(i64 " + m_v) + ")"))
		return gg
	}
	if (name == "_aria_mutex_unlock") {
		m_v := ("%call_mu_m_" + i2s(idx))
		_ = m_v
		gg = _lg_emit(gg, ((m_v + " = load i64, i64* ") + _t(inst.Arg1)))
		gg = _lg_emit(gg, (("call void @_aria_mutex_unlock(i64 " + m_v) + ")"))
		return gg
	}
	if (name == "_aria_rwmutex_new") {
		res := ("%call_rwn_" + i2s(idx))
		_ = res
		gg = _lg_emit(gg, (res + " = call i64 @_aria_rwmutex_new()"))
		gg = _lg_emit(gg, ((("store i64 " + res) + ", i64* ") + _t(inst.Dest)))
		return gg
	}
	if ((((name == "_aria_rwmutex_rlock") || (name == "_aria_rwmutex_runlock")) || (name == "_aria_rwmutex_wlock")) || (name == "_aria_rwmutex_wunlock")) {
		h := ("%call_rw_h_" + i2s(idx))
		_ = h
		gg = _lg_emit(gg, ((h + " = load i64, i64* ") + _t(inst.Arg1)))
		gg = _lg_emit(gg, (((("call void @" + name) + "(i64 ") + h) + ")"))
		return gg
	}
	if (name == "_aria_wg_new") {
		res := ("%call_wgn_" + i2s(idx))
		_ = res
		gg = _lg_emit(gg, (res + " = call i64 @_aria_wg_new()"))
		gg = _lg_emit(gg, ((("store i64 " + res) + ", i64* ") + _t(inst.Dest)))
		return gg
	}
	if (name == "_aria_wg_add") {
		h := ("%call_wga_h_" + i2s(idx))
		_ = h
		d := ("%call_wga_d_" + i2s(idx))
		_ = d
		gg = _lg_emit(gg, ((h + " = load i64, i64* ") + _t(inst.Arg1)))
		gg = _lg_emit(gg, ((d + " = load i64, i64* ") + _t((inst.Arg1 + int64(1)))))
		gg = _lg_emit(gg, (((("call void @_aria_wg_add(i64 " + h) + ", i64 ") + d) + ")"))
		return gg
	}
	if ((name == "_aria_wg_done") || (name == "_aria_wg_wait")) {
		h := ("%call_wg_h_" + i2s(idx))
		_ = h
		gg = _lg_emit(gg, ((h + " = load i64, i64* ") + _t(inst.Arg1)))
		gg = _lg_emit(gg, (((("call void @" + name) + "(i64 ") + h) + ")"))
		return gg
	}
	if (name == "_aria_once_new") {
		res := ("%call_on_" + i2s(idx))
		_ = res
		gg = _lg_emit(gg, (res + " = call i64 @_aria_once_new()"))
		gg = _lg_emit(gg, ((("store i64 " + res) + ", i64* ") + _t(inst.Dest)))
		return gg
	}
	if (((name == "_aria_task_done") || (name == "_aria_task_result")) || (name == "_aria_cancel_check")) {
		h := ("%call_td_h_" + i2s(idx))
		_ = h
		gg = _lg_emit(gg, ((h + " = load i64, i64* ") + _t(inst.Arg1)))
		res := ("%call_td_r_" + i2s(idx))
		_ = res
		gg = _lg_emit(gg, (((((res + " = call i64 @") + name) + "(i64 ") + h) + ")"))
		gg = _lg_emit(gg, ((("store i64 " + res) + ", i64* ") + _t(inst.Dest)))
		return gg
	}
	if (name == "_aria_task_cancel") {
		h := ("%call_tc_h_" + i2s(idx))
		_ = h
		gg = _lg_emit(gg, ((h + " = load i64, i64* ") + _t(inst.Arg1)))
		gg = _lg_emit(gg, (("call void @_aria_task_cancel(i64 " + h) + ")"))
		return gg
	}
	if (name == "_aria_spawn2") {
		fp_v := ("%call_sp2_fp_" + i2s(idx))
		_ = fp_v
		ep_v := ("%call_sp2_ep_" + i2s(idx))
		_ = ep_v
		gg = _lg_emit(gg, ((fp_v + " = load i64, i64* ") + _t(inst.Arg1)))
		gg = _lg_emit(gg, ((ep_v + " = load i64, i64* ") + _t((inst.Arg1 + int64(1)))))
		res := ("%call_sp2_r_" + i2s(idx))
		_ = res
		gg = _lg_emit(gg, (((((res + " = call i64 @_aria_spawn2(i64 ") + fp_v) + ", i64 ") + ep_v) + ")"))
		gg = _lg_emit(gg, ((("store i64 " + res) + ", i64* ") + _t(inst.Dest)))
		return gg
	}
	if (name == "_aria_task_await2") {
		h := ("%call_ta2_h_" + i2s(idx))
		_ = h
		gg = _lg_emit(gg, ((h + " = load i64, i64* ") + _t(inst.Arg1)))
		res := ("%call_ta2_r_" + i2s(idx))
		_ = res
		gg = _lg_emit(gg, (((res + " = call i64 @_aria_task_await2(i64 ") + h) + ")"))
		gg = _lg_emit(gg, ((("store i64 " + res) + ", i64* ") + _t(inst.Dest)))
		return gg
	}
	if (name == "_aria_chan_try_recv") {
		ch_v := ("%call_ctr_ch_" + i2s(idx))
		_ = ch_v
		gg = _lg_emit(gg, ((ch_v + " = load i64, i64* ") + _t(inst.Arg1)))
		si := _struct_i64_i64()
		_ = si
		res := ("%call_ctr_r_" + i2s(idx))
		_ = res
		gg = _lg_emit(gg, (((((res + " = call ") + si) + " @_aria_chan_try_recv(i64 ") + ch_v) + ")"))
		val := ("%call_ctr_v_" + i2s(idx))
		_ = val
		gg = _lg_emit(gg, (((((val + " = extractvalue ") + si) + " ") + res) + ", 0"))
		gg = _lg_emit(gg, ((("store i64 " + val) + ", i64* ") + _t(inst.Dest)))
		sta := ("%call_ctr_s_" + i2s(idx))
		_ = sta
		gg = _lg_emit(gg, (((((sta + " = extractvalue ") + si) + " ") + res) + ", 1"))
		gg = _lg_emit(gg, ((("store i64 " + sta) + ", i64* ") + _t((inst.Dest + int64(1)))))
		return gg
	}
	if (name == "_aria_chan_select") {
		arr_v := ("%call_cs_a_" + i2s(idx))
		_ = arr_v
		to_v := ("%call_cs_t_" + i2s(idx))
		_ = to_v
		gg = _lg_emit(gg, ((arr_v + " = load i64, i64* ") + _t(inst.Arg1)))
		gg = _lg_emit(gg, ((to_v + " = load i64, i64* ") + _t((inst.Arg1 + int64(1)))))
		si := _struct_i64_i64()
		_ = si
		res := ("%call_cs_r_" + i2s(idx))
		_ = res
		gg = _lg_emit(gg, (((((((res + " = call ") + si) + " @_aria_chan_select(i64 ") + arr_v) + ", i64 ") + to_v) + ")"))
		idx_v := ("%call_cs_i_" + i2s(idx))
		_ = idx_v
		gg = _lg_emit(gg, (((((idx_v + " = extractvalue ") + si) + " ") + res) + ", 0"))
		gg = _lg_emit(gg, ((("store i64 " + idx_v) + ", i64* ") + _t(inst.Dest)))
		val_v := ("%call_cs_v_" + i2s(idx))
		_ = val_v
		gg = _lg_emit(gg, (((((val_v + " = extractvalue ") + si) + " ") + res) + ", 1"))
		gg = _lg_emit(gg, ((("store i64 " + val_v) + ", i64* ") + _t((inst.Dest + int64(1)))))
		return gg
	}
	if (name == "_aria_cancel_new") {
		res := ("%call_cn_" + i2s(idx))
		_ = res
		gg = _lg_emit(gg, (res + " = call i64 @_aria_cancel_new()"))
		gg = _lg_emit(gg, ((("store i64 " + res) + ", i64* ") + _t(inst.Dest)))
		return gg
	}
	if (name == "_aria_cancel_child") {
		h := ("%call_cc_p_" + i2s(idx))
		_ = h
		gg = _lg_emit(gg, ((h + " = load i64, i64* ") + _t(inst.Arg1)))
		res := ("%call_cc_r_" + i2s(idx))
		_ = res
		gg = _lg_emit(gg, (((res + " = call i64 @_aria_cancel_child(i64 ") + h) + ")"))
		gg = _lg_emit(gg, ((("store i64 " + res) + ", i64* ") + _t(inst.Dest)))
		return gg
	}
	if (name == "_aria_cancel_trigger") {
		h := ("%call_ct_h_" + i2s(idx))
		_ = h
		gg = _lg_emit(gg, ((h + " = load i64, i64* ") + _t(inst.Arg1)))
		gg = _lg_emit(gg, (("call void @_aria_cancel_trigger(i64 " + h) + ")"))
		return gg
	}
	if (name == "_aria_cancel_is_triggered") {
		h := ("%call_cit_h_" + i2s(idx))
		_ = h
		gg = _lg_emit(gg, ((h + " = load i64, i64* ") + _t(inst.Arg1)))
		res := ("%call_cit_r_" + i2s(idx))
		_ = res
		gg = _lg_emit(gg, (((res + " = call i64 @_aria_cancel_is_triggered(i64 ") + h) + ")"))
		gg = _lg_emit(gg, ((("store i64 " + res) + ", i64* ") + _t(inst.Dest)))
		return gg
	}
	args := ""
	_ = args
	ai := int64(0)
	_ = ai
	for (ai < inst.Arg2) {
		if (ai > int64(0)) {
			args = (args + ", ")
		}
		av := ((("%call_a" + i2s(ai)) + "_") + i2s(idx))
		_ = av
		gg = _lg_emit(gg, ((av + " = load i64, i64* ") + _t((inst.Arg1 + ai))))
		args = ((args + "i64 ") + av)
		ai = (ai + int64(1))
	}
	si64 := _struct_i64_i64()
	_ = si64
	target_func := mod_find_func(m, orig_name)
	_ = target_func
	if ((target_func.Name != "") && (target_func.Return_type == int64(12))) {
		res := ("%call_r_" + i2s(idx))
		_ = res
		gg = _lg_emit(gg, (((((((res + " = call ") + si64) + " @") + name) + "(") + args) + ")"))
		r0 := ("%call_r0_" + i2s(idx))
		_ = r0
		r1 := ("%call_r1_" + i2s(idx))
		_ = r1
		gg = _lg_emit(gg, (((((r0 + " = extractvalue ") + si64) + " ") + res) + ", 0"))
		gg = _lg_emit(gg, (((((r1 + " = extractvalue ") + si64) + " ") + res) + ", 1"))
		gg = _lg_emit(gg, ((("store i64 " + r0) + ", i64* ") + _t(inst.Dest)))
		gg = _lg_emit(gg, ((("store i64 " + r1) + ", i64* ") + _t((inst.Dest + int64(1)))))
		return gg
	}
	if (inst.Type_id == int64(12)) {
		res := ("%call_r_" + i2s(idx))
		_ = res
		gg = _lg_emit(gg, (((((((res + " = call ") + si64) + " @") + name) + "(") + args) + ")"))
		r0 := ("%call_r0_" + i2s(idx))
		_ = r0
		r1 := ("%call_r1_" + i2s(idx))
		_ = r1
		gg = _lg_emit(gg, (((((r0 + " = extractvalue ") + si64) + " ") + res) + ", 0"))
		gg = _lg_emit(gg, (((((r1 + " = extractvalue ") + si64) + " ") + res) + ", 1"))
		gg = _lg_emit(gg, ((("store i64 " + r0) + ", i64* ") + _t(inst.Dest)))
		gg = _lg_emit(gg, ((("store i64 " + r1) + ", i64* ") + _t((inst.Dest + int64(1)))))
		return gg
	}
	if ((target_func.Name != "") && (target_func.Return_type == int64(0))) {
		gg = _lg_emit(gg, (((("call void @" + name) + "(") + args) + ")"))
		return gg
	}
	res := ("%call_r_" + i2s(idx))
	_ = res
	gg = _lg_emit(gg, (((((res + " = call i64 @") + name) + "(") + args) + ")"))
	gg = _lg_emit(gg, ((("store i64 " + res) + ", i64* ") + _t(inst.Dest)))
	return gg
}

func _gen_float_binop(g LlvmGen, llvm_op string, inst IrInst, idx int64) LlvmGen {
	gg := g
	_ = gg
	ai := ("%fb_ai_" + i2s(idx))
	_ = ai
	bi := ("%fb_bi_" + i2s(idx))
	_ = bi
	gg = _lg_emit(gg, ((ai + " = load i64, i64* ") + _t(inst.Arg1)))
	gg = _lg_emit(gg, ((bi + " = load i64, i64* ") + _t(inst.Arg2)))
	af := ("%fb_af_" + i2s(idx))
	_ = af
	bf := ("%fb_bf_" + i2s(idx))
	_ = bf
	gg = _lg_emit(gg, (((af + " = bitcast i64 ") + ai) + " to double"))
	gg = _lg_emit(gg, (((bf + " = bitcast i64 ") + bi) + " to double"))
	rf := ("%fb_rf_" + i2s(idx))
	_ = rf
	gg = _lg_emit(gg, ((((((rf + " = ") + llvm_op) + " double ") + af) + ", ") + bf))
	ri := ("%fb_ri_" + i2s(idx))
	_ = ri
	gg = _lg_emit(gg, (((ri + " = bitcast double ") + rf) + " to i64"))
	gg = _lg_emit(gg, ((("store i64 " + ri) + ", i64* ") + _t(inst.Dest)))
	return gg
}

func _gen_fcmp(g LlvmGen, cmp_op string, inst IrInst, idx int64) LlvmGen {
	gg := g
	_ = gg
	ai := ("%fc_ai_" + i2s(idx))
	_ = ai
	bi := ("%fc_bi_" + i2s(idx))
	_ = bi
	gg = _lg_emit(gg, ((ai + " = load i64, i64* ") + _t(inst.Arg1)))
	gg = _lg_emit(gg, ((bi + " = load i64, i64* ") + _t(inst.Arg2)))
	af := ("%fc_af_" + i2s(idx))
	_ = af
	bf := ("%fc_bf_" + i2s(idx))
	_ = bf
	gg = _lg_emit(gg, (((af + " = bitcast i64 ") + ai) + " to double"))
	gg = _lg_emit(gg, (((bf + " = bitcast i64 ") + bi) + " to double"))
	cmp := ("%fc_c_" + i2s(idx))
	_ = cmp
	gg = _lg_emit(gg, ((((((cmp + " = fcmp ") + cmp_op) + " double ") + af) + ", ") + bf))
	res := ("%fc_r_" + i2s(idx))
	_ = res
	gg = _lg_emit(gg, (((res + " = zext i1 ") + cmp) + " to i64"))
	gg = _lg_emit(gg, ((("store i64 " + res) + ", i64* ") + _t(inst.Dest)))
	return gg
}

