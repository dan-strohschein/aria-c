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

type IrOp interface {
	isIrOp()
}

type IrOpOpConst struct{}
func (IrOpOpConst) isIrOp() {}

type IrOpOpConstStr struct{}
func (IrOpOpConstStr) isIrOp() {}

type IrOpOpAdd struct{}
func (IrOpOpAdd) isIrOp() {}

type IrOpOpSub struct{}
func (IrOpOpSub) isIrOp() {}

type IrOpOpMul struct{}
func (IrOpOpMul) isIrOp() {}

type IrOpOpDiv struct{}
func (IrOpOpDiv) isIrOp() {}

type IrOpOpMod struct{}
func (IrOpOpMod) isIrOp() {}

type IrOpOpNeg struct{}
func (IrOpOpNeg) isIrOp() {}

type IrOpOpEq struct{}
func (IrOpOpEq) isIrOp() {}

type IrOpOpNeq struct{}
func (IrOpOpNeq) isIrOp() {}

type IrOpOpLt struct{}
func (IrOpOpLt) isIrOp() {}

type IrOpOpGt struct{}
func (IrOpOpGt) isIrOp() {}

type IrOpOpLte struct{}
func (IrOpOpLte) isIrOp() {}

type IrOpOpGte struct{}
func (IrOpOpGte) isIrOp() {}

type IrOpOpAnd struct{}
func (IrOpOpAnd) isIrOp() {}

type IrOpOpOr struct{}
func (IrOpOpOr) isIrOp() {}

type IrOpOpNot struct{}
func (IrOpOpNot) isIrOp() {}

type IrOpOpBitAnd struct{}
func (IrOpOpBitAnd) isIrOp() {}

type IrOpOpBitOr struct{}
func (IrOpOpBitOr) isIrOp() {}

type IrOpOpBitXor struct{}
func (IrOpOpBitXor) isIrOp() {}

type IrOpOpShl struct{}
func (IrOpOpShl) isIrOp() {}

type IrOpOpShr struct{}
func (IrOpOpShr) isIrOp() {}

type IrOpOpCall struct{}
func (IrOpOpCall) isIrOp() {}

type IrOpOpCallIndirect struct{}
func (IrOpOpCallIndirect) isIrOp() {}

type IrOpOpFnRef struct{}
func (IrOpOpFnRef) isIrOp() {}

type IrOpOpRet struct{}
func (IrOpOpRet) isIrOp() {}

type IrOpOpRetVoid struct{}
func (IrOpOpRetVoid) isIrOp() {}

type IrOpOpArg struct{}
func (IrOpOpArg) isIrOp() {}

type IrOpOpLocal struct{}
func (IrOpOpLocal) isIrOp() {}

type IrOpOpStore struct{}
func (IrOpOpStore) isIrOp() {}

type IrOpOpLoad struct{}
func (IrOpOpLoad) isIrOp() {}

type IrOpOpFieldGet struct{}
func (IrOpOpFieldGet) isIrOp() {}

type IrOpOpFieldSet struct{}
func (IrOpOpFieldSet) isIrOp() {}

type IrOpOpAlloc struct{}
func (IrOpOpAlloc) isIrOp() {}

type IrOpOpArrayNew struct{}
func (IrOpOpArrayNew) isIrOp() {}

type IrOpOpArrayGet struct{}
func (IrOpOpArrayGet) isIrOp() {}

type IrOpOpArraySet struct{}
func (IrOpOpArraySet) isIrOp() {}

type IrOpOpArrayLen struct{}
func (IrOpOpArrayLen) isIrOp() {}

type IrOpOpArrayAppend struct{}
func (IrOpOpArrayAppend) isIrOp() {}

type IrOpOpStrLen struct{}
func (IrOpOpStrLen) isIrOp() {}

type IrOpOpStrConcat struct{}
func (IrOpOpStrConcat) isIrOp() {}

type IrOpOpStrEq struct{}
func (IrOpOpStrEq) isIrOp() {}

type IrOpOpStrCmp struct{}
func (IrOpOpStrCmp) isIrOp() {}

type IrOpOpStrCharAt struct{}
func (IrOpOpStrCharAt) isIrOp() {}

type IrOpOpStrSubstring struct{}
func (IrOpOpStrSubstring) isIrOp() {}

type IrOpOpStrContains struct{}
func (IrOpOpStrContains) isIrOp() {}

type IrOpOpStrStartsWith struct{}
func (IrOpOpStrStartsWith) isIrOp() {}

type IrOpOpStrEndsWith struct{}
func (IrOpOpStrEndsWith) isIrOp() {}

type IrOpOpStrIndexOf struct{}
func (IrOpOpStrIndexOf) isIrOp() {}

type IrOpOpStrTrim struct{}
func (IrOpOpStrTrim) isIrOp() {}

type IrOpOpStrReplace struct{}
func (IrOpOpStrReplace) isIrOp() {}

type IrOpOpStrToLower struct{}
func (IrOpOpStrToLower) isIrOp() {}

type IrOpOpStrToUpper struct{}
func (IrOpOpStrToUpper) isIrOp() {}

type IrOpOpStrSplit struct{}
func (IrOpOpStrSplit) isIrOp() {}

type IrOpOpIntToStr struct{}
func (IrOpOpIntToStr) isIrOp() {}

type IrOpOpStrToInt struct{}
func (IrOpOpStrToInt) isIrOp() {}

type IrOpOpMapNew struct{}
func (IrOpOpMapNew) isIrOp() {}

type IrOpOpMapSet struct{}
func (IrOpOpMapSet) isIrOp() {}

type IrOpOpMapGet struct{}
func (IrOpOpMapGet) isIrOp() {}

type IrOpOpMapContains struct{}
func (IrOpOpMapContains) isIrOp() {}

type IrOpOpMapLen struct{}
func (IrOpOpMapLen) isIrOp() {}

type IrOpOpMapKeys struct{}
func (IrOpOpMapKeys) isIrOp() {}

type IrOpOpSetNew struct{}
func (IrOpOpSetNew) isIrOp() {}

type IrOpOpSetAdd struct{}
func (IrOpOpSetAdd) isIrOp() {}

type IrOpOpSetContains struct{}
func (IrOpOpSetContains) isIrOp() {}

type IrOpOpSetLen struct{}
func (IrOpOpSetLen) isIrOp() {}

type IrOpOpSetValues struct{}
func (IrOpOpSetValues) isIrOp() {}

type IrOpOpSetRemove struct{}
func (IrOpOpSetRemove) isIrOp() {}

type IrOpOpJump struct{}
func (IrOpOpJump) isIrOp() {}

type IrOpOpBranchTrue struct{}
func (IrOpOpBranchTrue) isIrOp() {}

type IrOpOpBranchFalse struct{}
func (IrOpOpBranchFalse) isIrOp() {}

type IrOpOpLabel struct{}
func (IrOpOpLabel) isIrOp() {}

type IrOpOpTraitObject struct{}
func (IrOpOpTraitObject) isIrOp() {}

type IrOpOpVtableCall struct{}
func (IrOpOpVtableCall) isIrOp() {}

type IrOpOpArraySlice struct{}
func (IrOpOpArraySlice) isIrOp() {}

type IrOpOpNop struct{}
func (IrOpOpNop) isIrOp() {}

type IrInst struct {
	Op IrOp
	Dest int64
	Arg1 int64
	Arg2 int64
	S1 string
	Type_id int64
}

type IrFunc struct {
	Name string
	Param_count int64
	Insts []IrInst
	Return_type int64
	Local_count int64
	Temp_types []int64
}

type IrModule struct {
	Funcs []IrFunc
	String_constants []string
	Entry_func string
	Test_count int64
	Test_names []string
	Extern_names []string
	Extern_param_counts []int64
	Extern_ret_types []string
	Extern_libs []string
}

func ir_op_name(op IrOp) string {
	return func() interface{} {
		switch _tmp1 := (op).(type) {
		case IrOpOpConst:
			return "Const"
		case IrOpOpConstStr:
			return "ConstStr"
		case IrOpOpAdd:
			return "Add"
		case IrOpOpSub:
			return "Sub"
		case IrOpOpMul:
			return "Mul"
		case IrOpOpDiv:
			return "Div"
		case IrOpOpMod:
			return "Mod"
		case IrOpOpNeg:
			return "Neg"
		case IrOpOpEq:
			return "Eq"
		case IrOpOpNeq:
			return "Neq"
		case IrOpOpLt:
			return "Lt"
		case IrOpOpGt:
			return "Gt"
		case IrOpOpLte:
			return "Lte"
		case IrOpOpGte:
			return "Gte"
		case IrOpOpAnd:
			return "And"
		case IrOpOpOr:
			return "Or"
		case IrOpOpNot:
			return "Not"
		case IrOpOpBitAnd:
			return "BitAnd"
		case IrOpOpBitOr:
			return "BitOr"
		case IrOpOpBitXor:
			return "BitXor"
		case IrOpOpShl:
			return "Shl"
		case IrOpOpShr:
			return "Shr"
		case IrOpOpCall:
			return "Call"
		case IrOpOpCallIndirect:
			return "CallIndirect"
		case IrOpOpFnRef:
			return "FnRef"
		case IrOpOpRet:
			return "Ret"
		case IrOpOpRetVoid:
			return "RetVoid"
		case IrOpOpArg:
			return "Arg"
		case IrOpOpLocal:
			return "Local"
		case IrOpOpStore:
			return "Store"
		case IrOpOpLoad:
			return "Load"
		case IrOpOpFieldGet:
			return "FieldGet"
		case IrOpOpFieldSet:
			return "FieldSet"
		case IrOpOpAlloc:
			return "Alloc"
		case IrOpOpArrayNew:
			return "ArrayNew"
		case IrOpOpArrayGet:
			return "ArrayGet"
		case IrOpOpArraySet:
			return "ArraySet"
		case IrOpOpArrayLen:
			return "ArrayLen"
		case IrOpOpArrayAppend:
			return "ArrayAppend"
		case IrOpOpStrLen:
			return "StrLen"
		case IrOpOpStrConcat:
			return "StrConcat"
		case IrOpOpStrEq:
			return "StrEq"
		case IrOpOpStrCmp:
			return "StrCmp"
		case IrOpOpStrCharAt:
			return "StrCharAt"
		case IrOpOpStrSubstring:
			return "StrSubstring"
		case IrOpOpStrContains:
			return "StrContains"
		case IrOpOpStrStartsWith:
			return "StrStartsWith"
		case IrOpOpStrEndsWith:
			return "StrEndsWith"
		case IrOpOpStrIndexOf:
			return "StrIndexOf"
		case IrOpOpStrTrim:
			return "StrTrim"
		case IrOpOpStrReplace:
			return "StrReplace"
		case IrOpOpStrToLower:
			return "StrToLower"
		case IrOpOpStrToUpper:
			return "StrToUpper"
		case IrOpOpStrSplit:
			return "StrSplit"
		case IrOpOpIntToStr:
			return "IntToStr"
		case IrOpOpStrToInt:
			return "StrToInt"
		case IrOpOpMapNew:
			return "MapNew"
		case IrOpOpMapSet:
			return "MapSet"
		case IrOpOpMapGet:
			return "MapGet"
		case IrOpOpMapContains:
			return "MapContains"
		case IrOpOpMapLen:
			return "MapLen"
		case IrOpOpMapKeys:
			return "MapKeys"
		case IrOpOpSetNew:
			return "SetNew"
		case IrOpOpSetAdd:
			return "SetAdd"
		case IrOpOpSetContains:
			return "SetContains"
		case IrOpOpSetLen:
			return "SetLen"
		case IrOpOpSetValues:
			return "SetValues"
		case IrOpOpSetRemove:
			return "SetRemove"
		case IrOpOpJump:
			return "Jump"
		case IrOpOpBranchTrue:
			return "BranchTrue"
		case IrOpOpBranchFalse:
			return "BranchFalse"
		case IrOpOpLabel:
			return "Label"
		case IrOpOpTraitObject:
			return "TraitObject"
		case IrOpOpVtableCall:
			return "VtableCall"
		case IrOpOpArraySlice:
			return "ArraySlice"
		case IrOpOpNop:
			return "Nop"
		default:
			_ = _tmp1
		}
		return nil
	}().(string)
}

func _sentinel_inst() IrInst {
	return IrInst{Op: IrOpOpNop{}, Dest: int64(0), Arg1: int64(0), Arg2: int64(0), S1: "", Type_id: int64(0)}
}

func new_inst(op IrOp, dest int64, arg1 int64, arg2 int64, s1 string, tid int64) IrInst {
	return IrInst{Op: op, Dest: dest, Arg1: arg1, Arg2: arg2, S1: s1, Type_id: tid}
}

func _sentinel_func() IrFunc {
	return IrFunc{Name: "", Param_count: int64(0), Insts: []IrInst{_sentinel_inst()}, Return_type: int64(0), Local_count: int64(0), Temp_types: []int64{int64(0)}}
}

func new_ir_func(name string, param_count int64, ret_type int64) IrFunc {
	return IrFunc{Name: name, Param_count: param_count, Insts: []IrInst{_sentinel_inst()}, Return_type: ret_type, Local_count: int64(0), Temp_types: []int64{int64(0)}}
}

func func_add_inst(f IrFunc, inst IrInst) IrFunc {
	return IrFunc{Name: f.Name, Param_count: f.Param_count, Insts: append(f.Insts, inst), Return_type: f.Return_type, Local_count: f.Local_count, Temp_types: f.Temp_types}
}

func func_set_locals(f IrFunc, count int64) IrFunc {
	return IrFunc{Name: f.Name, Param_count: f.Param_count, Insts: f.Insts, Return_type: f.Return_type, Local_count: count, Temp_types: f.Temp_types}
}

func func_set_temp_types(f IrFunc, types []int64) IrFunc {
	return IrFunc{Name: f.Name, Param_count: f.Param_count, Insts: f.Insts, Return_type: f.Return_type, Local_count: f.Local_count, Temp_types: types}
}

func func_inst_count(f IrFunc) int64 {
	return (int64(len(f.Insts)) - int64(1))
}

func new_ir_module() IrModule {
	return IrModule{Funcs: []IrFunc{_sentinel_func()}, String_constants: []string{""}, Entry_func: "", Test_count: int64(0), Test_names: []string{""}, Extern_names: []string{""}, Extern_param_counts: []int64{int64(0)}, Extern_ret_types: []string{""}, Extern_libs: []string{""}}
}

func mod_add_func(m IrModule, f IrFunc) IrModule {
	return IrModule{Funcs: append(m.Funcs, f), String_constants: m.String_constants, Entry_func: m.Entry_func, Test_count: m.Test_count, Test_names: m.Test_names, Extern_names: m.Extern_names, Extern_param_counts: m.Extern_param_counts, Extern_ret_types: m.Extern_ret_types, Extern_libs: m.Extern_libs}
}

func mod_add_string(m IrModule, s string) IrModule {
	return IrModule{Funcs: m.Funcs, String_constants: append(m.String_constants, s), Entry_func: m.Entry_func, Test_count: m.Test_count, Test_names: m.Test_names, Extern_names: m.Extern_names, Extern_param_counts: m.Extern_param_counts, Extern_ret_types: m.Extern_ret_types, Extern_libs: m.Extern_libs}
}

func mod_add_extern(m IrModule, name string, param_count int64, ret_type string, lib string) IrModule {
	return IrModule{Funcs: m.Funcs, String_constants: m.String_constants, Entry_func: m.Entry_func, Test_count: m.Test_count, Test_names: m.Test_names, Extern_names: append(m.Extern_names, name), Extern_param_counts: append(m.Extern_param_counts, param_count), Extern_ret_types: append(m.Extern_ret_types, ret_type), Extern_libs: append(m.Extern_libs, lib)}
}

func mod_set_entry(m IrModule, name string) IrModule {
	return IrModule{Funcs: m.Funcs, String_constants: m.String_constants, Entry_func: name, Test_count: m.Test_count, Test_names: m.Test_names, Extern_names: m.Extern_names, Extern_param_counts: m.Extern_param_counts, Extern_ret_types: m.Extern_ret_types, Extern_libs: m.Extern_libs}
}

func mod_add_test(m IrModule, name string) IrModule {
	return IrModule{Funcs: m.Funcs, String_constants: m.String_constants, Entry_func: m.Entry_func, Test_count: (m.Test_count + int64(1)), Test_names: append(m.Test_names, name), Extern_names: m.Extern_names, Extern_param_counts: m.Extern_param_counts, Extern_ret_types: m.Extern_ret_types, Extern_libs: m.Extern_libs}
}

func mod_find_string(m IrModule, s string) int64 {
	i := int64(1)
	_ = i
	for (i < int64(len(m.String_constants))) {
		if (m.String_constants[i] == s) {
			return i
		}
		i = (i + int64(1))
	}
	return int64(0)
}

func mod_func_count(m IrModule) int64 {
	return (int64(len(m.Funcs)) - int64(1))
}

func mod_find_func(m IrModule, name string) IrFunc {
	i := int64(1)
	_ = i
	for (i < int64(len(m.Funcs))) {
		if (m.Funcs[i].Name == name) {
			return m.Funcs[i]
		}
		i = (i + int64(1))
	}
	return _sentinel_func()
}

func _is_jump_op(op IrOp) bool {
	return (((ir_op_name(op) == "Jump") || (ir_op_name(op) == "BranchTrue")) || (ir_op_name(op) == "BranchFalse"))
}

func _is_def_op(op IrOp) bool {
	name := ir_op_name(op)
	_ = name
	if (name == "Const") {
		return true
	}
	if (name == "ConstStr") {
		return true
	}
	if (name == "Add") {
		return true
	}
	if (name == "Sub") {
		return true
	}
	if (name == "Mul") {
		return true
	}
	if (name == "Div") {
		return true
	}
	if (name == "Mod") {
		return true
	}
	if (name == "Neg") {
		return true
	}
	if (name == "Eq") {
		return true
	}
	if (name == "Neq") {
		return true
	}
	if (name == "Lt") {
		return true
	}
	if (name == "Gt") {
		return true
	}
	if (name == "Lte") {
		return true
	}
	if (name == "Gte") {
		return true
	}
	if (name == "And") {
		return true
	}
	if (name == "Or") {
		return true
	}
	if (name == "Not") {
		return true
	}
	if (name == "Call") {
		return true
	}
	if (name == "CallIndirect") {
		return true
	}
	if (name == "FnRef") {
		return true
	}
	if (name == "Load") {
		return true
	}
	if (name == "FieldGet") {
		return true
	}
	if (name == "Alloc") {
		return true
	}
	if (name == "ArrayNew") {
		return true
	}
	if (name == "ArrayGet") {
		return true
	}
	if (name == "ArrayLen") {
		return true
	}
	if (name == "ArrayAppend") {
		return true
	}
	if (name == "MapNew") {
		return true
	}
	if (name == "SetNew") {
		return true
	}
	if (name == "IntToStr") {
		return true
	}
	if (name == "StrToInt") {
		return true
	}
	return false
}

func _validate_func(f IrFunc) []string {
	errors := []string{""}
	_ = errors
	labels := []int64{int64(0)}
	_ = labels
	ii := int64(1)
	_ = ii
	for (ii < int64(len(f.Insts))) {
		inst := f.Insts[ii]
		_ = inst
		if (ir_op_name(inst.Op) == "Label") {
			labels = append(labels, inst.Arg1)
		}
		ii = (ii + int64(1))
	}
	ji := int64(1)
	_ = ji
	for (ji < int64(len(f.Insts))) {
		inst := f.Insts[ji]
		_ = inst
		if _is_jump_op(inst.Op) {
			op_name := ir_op_name(inst.Op)
			_ = op_name
			target := inst.Arg1
			_ = target
			if ((op_name == "BranchTrue") || (op_name == "BranchFalse")) {
				target = inst.Arg2
			}
			found := false
			_ = found
			li := int64(1)
			_ = li
			for (li < int64(len(labels))) {
				if (labels[li] == target) {
					found = true
				}
				li = (li + int64(1))
			}
			if (found == false) {
				errors = append(errors, (((("E9002: jump to undefined label L" + _ir_i2s(target)) + " in function '") + f.Name) + "'"))
			}
		}
		ji = (ji + int64(1))
	}
	return errors
}

func _ir_i2s(n int64) string {
	if (n == int64(0)) {
		return "0"
	}
	result := ""
	_ = result
	v := n
	_ = v
	if (v < int64(0)) {
		v = (int64(0) - v)
	}
	for (v > int64(0)) {
		d := (v - (((v / int64(10))) * int64(10)))
		_ = d
		ch := "0"
		_ = ch
		if (d == int64(1)) {
			ch = "1"
		}
		if (d == int64(2)) {
			ch = "2"
		}
		if (d == int64(3)) {
			ch = "3"
		}
		if (d == int64(4)) {
			ch = "4"
		}
		if (d == int64(5)) {
			ch = "5"
		}
		if (d == int64(6)) {
			ch = "6"
		}
		if (d == int64(7)) {
			ch = "7"
		}
		if (d == int64(8)) {
			ch = "8"
		}
		if (d == int64(9)) {
			ch = "9"
		}
		result = (ch + result)
		v = (v / int64(10))
	}
	if (n < int64(0)) {
		result = ("-" + result)
	}
	return result
}

func validate_ir_module(m IrModule) []string {
	all_errors := []string{""}
	_ = all_errors
	if (int64(len(m.Entry_func)) > int64(0)) {
		ef := mod_find_func(m, m.Entry_func)
		_ = ef
		if (ef.Name == "") {
			all_errors = append(all_errors, (("E9004: entry function '" + m.Entry_func) + "' not found in module"))
		}
	}
	fi := int64(1)
	_ = fi
	for (fi < int64(len(m.Funcs))) {
		func_errors := _validate_func(m.Funcs[fi])
		_ = func_errors
		ei := int64(1)
		_ = ei
		for (ei < int64(len(func_errors))) {
			all_errors = append(all_errors, func_errors[ei])
			ei = (ei + int64(1))
		}
		fi = (fi + int64(1))
	}
	return all_errors
}

