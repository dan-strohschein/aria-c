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

type TypeKind interface {
	isTypeKind()
}

type TypeKindTyI8 struct{}
func (TypeKindTyI8) isTypeKind() {}

type TypeKindTyI16 struct{}
func (TypeKindTyI16) isTypeKind() {}

type TypeKindTyI32 struct{}
func (TypeKindTyI32) isTypeKind() {}

type TypeKindTyI64 struct{}
func (TypeKindTyI64) isTypeKind() {}

type TypeKindTyU8 struct{}
func (TypeKindTyU8) isTypeKind() {}

type TypeKindTyU16 struct{}
func (TypeKindTyU16) isTypeKind() {}

type TypeKindTyU32 struct{}
func (TypeKindTyU32) isTypeKind() {}

type TypeKindTyU64 struct{}
func (TypeKindTyU64) isTypeKind() {}

type TypeKindTyF32 struct{}
func (TypeKindTyF32) isTypeKind() {}

type TypeKindTyF64 struct{}
func (TypeKindTyF64) isTypeKind() {}

type TypeKindTyBool struct{}
func (TypeKindTyBool) isTypeKind() {}

type TypeKindTyStr struct{}
func (TypeKindTyStr) isTypeKind() {}

type TypeKindTyChar struct{}
func (TypeKindTyChar) isTypeKind() {}

type TypeKindTyUnit struct{}
func (TypeKindTyUnit) isTypeKind() {}

type TypeKindTyArray struct{}
func (TypeKindTyArray) isTypeKind() {}

type TypeKindTyFunction struct{}
func (TypeKindTyFunction) isTypeKind() {}

type TypeKindTyNamed struct{}
func (TypeKindTyNamed) isTypeKind() {}

type TypeKindTyResult struct{}
func (TypeKindTyResult) isTypeKind() {}

type TypeKindTyOptional struct{}
func (TypeKindTyOptional) isTypeKind() {}

type TypeKindTyTypeVar struct{}
func (TypeKindTyTypeVar) isTypeKind() {}

type TypeKindTyTraitObject struct{}
func (TypeKindTyTraitObject) isTypeKind() {}

type TypeKindTyTuple struct{}
func (TypeKindTyTuple) isTypeKind() {}

type TypeKindTyNever struct{}
func (TypeKindTyNever) isTypeKind() {}

type TypeKindTyUnknown struct{}
func (TypeKindTyUnknown) isTypeKind() {}

type TypeKindTyNone struct{}
func (TypeKindTyNone) isTypeKind() {}

type TypeKindTyApplied struct{}
func (TypeKindTyApplied) isTypeKind() {}

type TypeInfo struct {
	Kind TypeKind
	Name string
	Type_id int64
	Type_id2 int64
	Param_start int64
	Param_count int64
}

type TypeStore struct {
	Types []TypeInfo
	Name_index StrMap
}

type FieldInfo struct {
	Name string
	Type_id int64
}

type StructDef struct {
	Name string
	Fields []FieldInfo
	Is_sum bool
	Variant_names []string
	Variant_field_counts []int64
	Variant_field_names []string
	Variant_field_types []int64
	Variant_field_offsets []int64
	Generic_count int64
	Generic_name_0 string
	Generic_name_1 string
	Generic_name_2 string
}

type FnSig struct {
	Name string
	Param_names []string
	Param_types []int64
	Return_type int64
	Error_type int64
	Generic_count int64
	Generic_name_0 string
	Generic_name_1 string
	Generic_name_2 string
	Generic_bound_0 string
	Generic_bound_1 string
	Generic_bound_2 string
	Token_start int64
	Has_io bool
	Has_fs bool
	Has_net bool
	Has_ffi bool
	Has_async bool
	User_effects []string
}

type MonoSpec struct {
	Generic_name string
	Specialized_name string
	Type_arg_0 int64
	Type_arg_1 int64
	Type_arg_2 int64
}

type TypeRegistry struct {
	Struct_defs []StructDef
	Fn_sigs []FnSig
	Mono_specs []MonoSpec
	Deprecated_names []string
	Deprecated_msgs []string
	Cold_names []string
	Fn_value_refs []string
	Bound_fn_value_refs []string
}

func type_kind_name(k TypeKind) string {
	return func() interface{} {
		switch _tmp1 := (k).(type) {
		case TypeKindTyI8:
			return "i8"
		case TypeKindTyI16:
			return "i16"
		case TypeKindTyI32:
			return "i32"
		case TypeKindTyI64:
			return "i64"
		case TypeKindTyU8:
			return "u8"
		case TypeKindTyU16:
			return "u16"
		case TypeKindTyU32:
			return "u32"
		case TypeKindTyU64:
			return "u64"
		case TypeKindTyF32:
			return "f32"
		case TypeKindTyF64:
			return "f64"
		case TypeKindTyBool:
			return "bool"
		case TypeKindTyStr:
			return "str"
		case TypeKindTyChar:
			return "char"
		case TypeKindTyUnit:
			return "unit"
		case TypeKindTyArray:
			return "Array"
		case TypeKindTyFunction:
			return "Function"
		case TypeKindTyNamed:
			return "Named"
		case TypeKindTyResult:
			return "Result"
		case TypeKindTyOptional:
			return "Optional"
		case TypeKindTyTypeVar:
			return "TypeVar"
		case TypeKindTyTraitObject:
			return "TraitObject"
		case TypeKindTyTuple:
			return "Tuple"
		case TypeKindTyNever:
			return "Never"
		case TypeKindTyUnknown:
			return "Unknown"
		case TypeKindTyNone:
			return "None"
		case TypeKindTyApplied:
			return "Applied"
		default:
			_ = _tmp1
		}
		return nil
	}().(string)
}

func tk_eq(a TypeKind, b TypeKind) bool {
	return (type_kind_name(a) == type_kind_name(b))
}

func _sentinel_type() TypeInfo {
	return TypeInfo{Kind: TypeKindTyNone{}, Name: "", Type_id: int64(0), Type_id2: int64(0), Param_start: int64(0), Param_count: int64(0)}
}

func _mk_prim(kind TypeKind) TypeInfo {
	return TypeInfo{Kind: kind, Name: type_kind_name(kind), Type_id: int64(0), Type_id2: int64(0), Param_start: int64(0), Param_count: int64(0)}
}

func type_info_name(ti TypeInfo) string {
	return ti.Name
}

func type_info_kind(ti TypeInfo) TypeKind {
	return ti.Kind
}

func type_info_type_id(ti TypeInfo) int64 {
	return ti.Type_id
}

func mk_array_type(elem_id int64) TypeInfo {
	return TypeInfo{Kind: TypeKindTyArray{}, Name: "", Type_id: elem_id, Type_id2: int64(0), Param_start: int64(0), Param_count: int64(0)}
}

func mk_fn_type(return_id int64, param_start int64, param_count int64) TypeInfo {
	return TypeInfo{Kind: TypeKindTyFunction{}, Name: "", Type_id: return_id, Type_id2: int64(0), Param_start: param_start, Param_count: param_count}
}

func mk_named_type(name string) TypeInfo {
	return TypeInfo{Kind: TypeKindTyNamed{}, Name: name, Type_id: int64(0), Type_id2: int64(0), Param_start: int64(0), Param_count: int64(0)}
}

func mk_result_type(ok_id int64, err_id int64) TypeInfo {
	return TypeInfo{Kind: TypeKindTyResult{}, Name: "", Type_id: ok_id, Type_id2: err_id, Param_start: int64(0), Param_count: int64(0)}
}

func mk_optional_type(inner_id int64) TypeInfo {
	return TypeInfo{Kind: TypeKindTyOptional{}, Name: "", Type_id: inner_id, Type_id2: int64(0), Param_start: int64(0), Param_count: int64(0)}
}

func mk_typevar(name string) TypeInfo {
	return TypeInfo{Kind: TypeKindTyTypeVar{}, Name: name, Type_id: int64(0), Type_id2: int64(0), Param_start: int64(0), Param_count: int64(0)}
}

func mk_hkt_typevar(name string, arity int64) TypeInfo {
	return TypeInfo{Kind: TypeKindTyTypeVar{}, Name: name, Type_id: int64(0), Type_id2: int64(0), Param_start: int64(0), Param_count: arity}
}

func mk_applied_type(constructor_id int64, arg_id int64) TypeInfo {
	return TypeInfo{Kind: TypeKindTyApplied{}, Name: "", Type_id: constructor_id, Type_id2: arg_id, Param_start: int64(0), Param_count: int64(0)}
}

func mk_tuple_type(elem_ids string, arity int64) TypeInfo {
	return TypeInfo{Kind: TypeKindTyTuple{}, Name: elem_ids, Type_id: int64(0), Type_id2: int64(0), Param_start: int64(0), Param_count: arity}
}

func new_store() TypeStore {
	types := []TypeInfo{_sentinel_type()}
	_ = types
	types = append(types, _mk_prim(TypeKindTyI64{}))
	types = append(types, _mk_prim(TypeKindTyI32{}))
	types = append(types, _mk_prim(TypeKindTyI16{}))
	types = append(types, _mk_prim(TypeKindTyI8{}))
	types = append(types, _mk_prim(TypeKindTyU64{}))
	types = append(types, _mk_prim(TypeKindTyU32{}))
	types = append(types, _mk_prim(TypeKindTyU16{}))
	types = append(types, _mk_prim(TypeKindTyU8{}))
	types = append(types, _mk_prim(TypeKindTyF64{}))
	types = append(types, _mk_prim(TypeKindTyF32{}))
	types = append(types, _mk_prim(TypeKindTyBool{}))
	types = append(types, _mk_prim(TypeKindTyStr{}))
	types = append(types, _mk_prim(TypeKindTyChar{}))
	types = append(types, _mk_prim(TypeKindTyUnit{}))
	types = append(types, TypeInfo{Kind: TypeKindTyNever{}, Name: "Never", Type_id: int64(0), Type_id2: int64(0), Param_start: int64(0), Param_count: int64(0)})
	types = append(types, TypeInfo{Kind: TypeKindTyUnknown{}, Name: "Unknown", Type_id: int64(0), Type_id2: int64(0), Param_start: int64(0), Param_count: int64(0)})
	types = append(types, mk_named_type("Chan"))
	types = append(types, mk_named_type("Task"))
	types = append(types, mk_named_type("ChanSend"))
	types = append(types, mk_named_type("ChanRecv"))
	return TypeStore{Types: types, Name_index: strmap_new()}
}

func store_add(store TypeStore, info TypeInfo) TypeStore {
	new_types := append(store.Types, info)
	_ = new_types
	new_idx := (int64(len(new_types)) - int64(1))
	_ = new_idx
	kn := type_kind_name(info.Kind)
	_ = kn
	if ((((kn == "Named") || (kn == "TypeVar"))) && (int64(len(info.Name)) > int64(0))) {
		return TypeStore{Types: new_types, Name_index: strmap_put(store.Name_index, info.Name, new_idx)}
	}
	return TypeStore{Types: new_types, Name_index: store.Name_index}
}

func store_get(store TypeStore, id int64) TypeInfo {
	if ((id >= int64(0)) && (id < int64(len(store.Types)))) {
		return store.Types[id]
	}
	return _sentinel_type()
}

func store_last_id(store TypeStore) int64 {
	return (int64(len(store.Types)) - int64(1))
}

func store_find_type(store TypeStore, info TypeInfo) int64 {
	kn := type_kind_name(info.Kind)
	_ = kn
	i := int64(1)
	_ = i
	for (i < int64(len(store.Types))) {
		existing := store.Types[i]
		_ = existing
		if (type_kind_name(existing.Kind) == kn) {
			if ((kn == "Array") && (existing.Type_id == info.Type_id)) {
				return i
			}
			if ((kn == "Optional") && (existing.Type_id == info.Type_id)) {
				return i
			}
			if (((kn == "Result") && (existing.Type_id == info.Type_id)) && (existing.Type_id2 == info.Type_id2)) {
				return i
			}
			if ((kn == "Named") && (existing.Name == info.Name)) {
				return i
			}
			if ((kn == "TypeVar") && (existing.Name == info.Name)) {
				return i
			}
			if (((kn == "Tuple") && (existing.Name == info.Name)) && (existing.Param_count == info.Param_count)) {
				return i
			}
		}
		i = (i + int64(1))
	}
	return (int64(0) - int64(1))
}

func store_size(store TypeStore) int64 {
	return int64(len(store.Types))
}

func format_type(store TypeStore, id int64) string {
	if ((id <= int64(0)) || (id >= int64(len(store.Types)))) {
		return "<none>"
	}
	info := store.Types[id]
	_ = info
	kn := type_kind_name(info.Kind)
	_ = kn
	if ((((((((((((((((kn == "i64") || (kn == "i32")) || (kn == "i16")) || (kn == "i8")) || (kn == "u64")) || (kn == "u32")) || (kn == "u16")) || (kn == "u8")) || (kn == "f64")) || (kn == "f32")) || (kn == "bool")) || (kn == "str")) || (kn == "char")) || (kn == "unit")) || (kn == "Never")) || (kn == "Unknown")) {
		return kn
	}
	if (kn == "Named") {
		return info.Name
	}
	if (kn == "Array") {
		return (("[" + format_type(store, info.Type_id)) + "]")
	}
	if (kn == "Optional") {
		return (format_type(store, info.Type_id) + "?")
	}
	if (kn == "Result") {
		return ((format_type(store, info.Type_id) + " ! ") + format_type(store, info.Type_id2))
	}
	if (kn == "Tuple") {
		return (("(" + info.Name) + ")")
	}
	if (kn == "Function") {
		return "fn(...)"
	}
	if (kn == "TypeVar") {
		return info.Name
	}
	if (kn == "TraitObject") {
		return ("dyn " + info.Name)
	}
	return kn
}

func is_numeric(store TypeStore, id int64) bool {
	if (id <= int64(0)) {
		return false
	}
	k := store.Types[id].Kind
	_ = k
	return (((((((((tk_eq(k, TypeKindTyI64{}) || tk_eq(k, TypeKindTyI32{})) || tk_eq(k, TypeKindTyI16{})) || tk_eq(k, TypeKindTyI8{})) || tk_eq(k, TypeKindTyU64{})) || tk_eq(k, TypeKindTyU32{})) || tk_eq(k, TypeKindTyU16{})) || tk_eq(k, TypeKindTyU8{})) || tk_eq(k, TypeKindTyF64{})) || tk_eq(k, TypeKindTyF32{}))
}

func is_integer(store TypeStore, id int64) bool {
	if (id <= int64(0)) {
		return false
	}
	k := store.Types[id].Kind
	_ = k
	return (((((((tk_eq(k, TypeKindTyI64{}) || tk_eq(k, TypeKindTyI32{})) || tk_eq(k, TypeKindTyI16{})) || tk_eq(k, TypeKindTyI8{})) || tk_eq(k, TypeKindTyU64{})) || tk_eq(k, TypeKindTyU32{})) || tk_eq(k, TypeKindTyU16{})) || tk_eq(k, TypeKindTyU8{}))
}

func is_float(store TypeStore, id int64) bool {
	if (id <= int64(0)) {
		return false
	}
	k := store.Types[id].Kind
	_ = k
	return (tk_eq(k, TypeKindTyF64{}) || tk_eq(k, TypeKindTyF32{}))
}

func types_equal(store TypeStore, a int64, b int64) bool {
	if (a == b) {
		return true
	}
	if ((a <= int64(0)) || (b <= int64(0))) {
		return false
	}
	if ((a >= int64(len(store.Types))) || (b >= int64(len(store.Types)))) {
		return false
	}
	ai := store.Types[a]
	_ = ai
	bi := store.Types[b]
	_ = bi
	ak := type_kind_name(ai.Kind)
	_ = ak
	bk := type_kind_name(bi.Kind)
	_ = bk
	if (ak != bk) {
		return false
	}
	if (ak == "Named") {
		return (ai.Name == bi.Name)
	}
	if (ak == "Array") {
		return types_equal(store, ai.Type_id, bi.Type_id)
	}
	if (ak == "Result") {
		return (types_equal(store, ai.Type_id, bi.Type_id) && types_equal(store, ai.Type_id2, bi.Type_id2))
	}
	if (ak == "Optional") {
		return types_equal(store, ai.Type_id, bi.Type_id)
	}
	if (ak == "Tuple") {
		return ((ai.Name == bi.Name) && (ai.Param_count == bi.Param_count))
	}
	return true
}

func substitute_type(store TypeStore, tid int64, typevar_id int64, concrete_id int64) int64 {
	if (tid == typevar_id) {
		return concrete_id
	}
	if ((tid <= int64(0)) || (tid >= int64(len(store.Types)))) {
		return tid
	}
	info := store.Types[tid]
	_ = info
	kn := type_kind_name(info.Kind)
	_ = kn
	if (kn == "Array") {
		elem := substitute_type(store, info.Type_id, typevar_id, concrete_id)
		_ = elem
		if (elem == info.Type_id) {
			return tid
		}
		return elem
	}
	if (kn == "Applied") {
		new_ctor := substitute_type(store, info.Type_id, typevar_id, concrete_id)
		_ = new_ctor
		new_arg := substitute_type(store, info.Type_id2, typevar_id, concrete_id)
		_ = new_arg
		if ((new_ctor > int64(0)) && (new_ctor < int64(len(store.Types)))) {
			ctor_info := store.Types[new_ctor]
			_ = ctor_info
			if ((type_kind_name(ctor_info.Kind) == "Named") && (ctor_info.Name == "_Array")) {
				return new_arg
			}
		}
		return tid
	}
	return tid
}

func tuple_elem_type(store TypeStore, tuple_id int64, index int64) int64 {
	if ((tuple_id <= int64(0)) || (tuple_id >= int64(len(store.Types)))) {
		return TY_UNKNOWN
	}
	info := store.Types[tuple_id]
	_ = info
	if (type_kind_name(info.Kind) != "Tuple") {
		return TY_UNKNOWN
	}
	if ((index < int64(0)) || (index >= info.Param_count)) {
		return TY_UNKNOWN
	}
	ids := info.Name
	_ = ids
	cur_idx := int64(0)
	_ = cur_idx
	start := int64(0)
	_ = start
	i := int64(0)
	_ = i
	for (i <= int64(len(ids))) {
		at_delim := false
		_ = at_delim
		if (i == int64(len(ids))) {
			at_delim = true
		} else if (string(ids[i]) == ",") {
			at_delim = true
		}
		if at_delim {
			if (cur_idx == index) {
				return _parse_int_from(ids, start, i)
			}
			cur_idx = (cur_idx + int64(1))
			start = (i + int64(1))
		}
		i = (i + int64(1))
	}
	return TY_UNKNOWN
}

func _parse_int_from(s string, start_pos int64, end_pos int64) int64 {
	result := int64(0)
	_ = result
	i := start_pos
	_ = i
	for (i < end_pos) {
		ch := string(s[i])
		_ = ch
		digit := int64(0)
		_ = digit
		if (ch == "1") {
			digit = int64(1)
		}
		if (ch == "2") {
			digit = int64(2)
		}
		if (ch == "3") {
			digit = int64(3)
		}
		if (ch == "4") {
			digit = int64(4)
		}
		if (ch == "5") {
			digit = int64(5)
		}
		if (ch == "6") {
			digit = int64(6)
		}
		if (ch == "7") {
			digit = int64(7)
		}
		if (ch == "8") {
			digit = int64(8)
		}
		if (ch == "9") {
			digit = int64(9)
		}
		result = ((result * int64(10)) + digit)
		i = (i + int64(1))
	}
	return result
}

func _type_name_suffix(store TypeStore, tid int64) string {
	if (tid == TY_I64) {
		return "i64"
	}
	if (tid == TY_I32) {
		return "i32"
	}
	if (tid == TY_I16) {
		return "i16"
	}
	if (tid == TY_I8) {
		return "i8"
	}
	if (tid == TY_U64) {
		return "u64"
	}
	if (tid == TY_U32) {
		return "u32"
	}
	if (tid == TY_U16) {
		return "u16"
	}
	if (tid == TY_U8) {
		return "u8"
	}
	if (tid == TY_F64) {
		return "f64"
	}
	if (tid == TY_F32) {
		return "f32"
	}
	if (tid == TY_BOOL) {
		return "bool"
	}
	if (tid == TY_STR) {
		return "str"
	}
	if (tid == TY_CHAR) {
		return "char"
	}
	if (tid == TY_UNIT) {
		return "unit"
	}
	if ((tid > int64(0)) && (tid < int64(len(store.Types)))) {
		info := store.Types[tid]
		_ = info
		kn := type_kind_name(info.Kind)
		_ = kn
		if (kn == "Named") {
			return info.Name
		}
		if (kn == "Array") {
			return ("arr_" + _type_name_suffix(store, info.Type_id))
		}
	}
	return "unknown"
}

func is_assignable(store TypeStore, src int64, dst int64) bool {
	if types_equal(store, src, dst) {
		return true
	}
	if (src == TY_NEVER) {
		return true
	}
	if ((src == TY_UNKNOWN) || (dst == TY_UNKNOWN)) {
		return true
	}
	if ((((src > int64(0)) && (src < int64(len(store.Types)))) && (dst > int64(0))) && (dst < int64(len(store.Types)))) {
		si := store.Types[src]
		_ = si
		di := store.Types[dst]
		_ = di
		sk := type_kind_name(si.Kind)
		_ = sk
		dk := type_kind_name(di.Kind)
		_ = dk
		if (sk == dk) {
			if ((sk == "Optional") || (sk == "Array")) {
				if ((si.Type_id == TY_UNKNOWN) || (di.Type_id == TY_UNKNOWN)) {
					return true
				}
				return is_assignable(store, si.Type_id, di.Type_id)
			}
			if (sk == "Result") {
				ok_ok := (((si.Type_id == TY_UNKNOWN) || (di.Type_id == TY_UNKNOWN)) || is_assignable(store, si.Type_id, di.Type_id))
				_ = ok_ok
				err_ok := (((si.Type_id2 == TY_UNKNOWN) || (di.Type_id2 == TY_UNKNOWN)) || is_assignable(store, si.Type_id2, di.Type_id2))
				_ = err_ok
				return (ok_ok && err_ok)
			}
		}
	}
	return false
}

func resolve_type_name(store TypeStore, name string) int64 {
	if (name == "i64") {
		return TY_I64
	}
	if (name == "i32") {
		return TY_I32
	}
	if (name == "i16") {
		return TY_I16
	}
	if (name == "i8") {
		return TY_I8
	}
	if (name == "u64") {
		return TY_U64
	}
	if (name == "u32") {
		return TY_U32
	}
	if (name == "u16") {
		return TY_U16
	}
	if (name == "u8") {
		return TY_U8
	}
	if (name == "f64") {
		return TY_F64
	}
	if (name == "f32") {
		return TY_F32
	}
	if (name == "bool") {
		return TY_BOOL
	}
	if (name == "str") {
		return TY_STR
	}
	if (name == "char") {
		return TY_CHAR
	}
	if (name == "unit") {
		return TY_UNIT
	}
	idx := strmap_get(store.Name_index, name)
	_ = idx
	if ((idx > int64(0)) && (idx < int64(len(store.Types)))) {
		info := store.Types[idx]
		_ = info
		if ((type_kind_name(info.Kind) == "Named") && (info.Name == name)) {
			if ((info.Type_id2 == int64(1)) && (info.Type_id > int64(0))) {
				return info.Type_id
			}
			return idx
		}
		if ((type_kind_name(info.Kind) == "TypeVar") && (info.Name == name)) {
			return idx
		}
	}
	i := int64(21)
	_ = i
	for (i < int64(len(store.Types))) {
		info := store.Types[i]
		_ = info
		if ((type_kind_name(info.Kind) == "Named") && (info.Name == name)) {
			if ((info.Type_id2 == int64(1)) && (info.Type_id > int64(0))) {
				return info.Type_id
			}
			return i
		}
		if ((type_kind_name(info.Kind) == "TypeVar") && (info.Name == name)) {
			return i
		}
		i = (i + int64(1))
	}
	return TY_UNKNOWN
}

func _sentinel_field() FieldInfo {
	return FieldInfo{Name: "", Type_id: int64(0)}
}

func _sentinel_struct_def() StructDef {
	return StructDef{Name: "", Fields: []FieldInfo{FieldInfo{Name: "", Type_id: int64(0)}}, Is_sum: false, Variant_names: []string{""}, Variant_field_counts: []int64{int64(0)}, Variant_field_names: []string{""}, Variant_field_types: []int64{int64(0)}, Variant_field_offsets: []int64{int64(0)}, Generic_count: int64(0), Generic_name_0: "", Generic_name_1: "", Generic_name_2: ""}
}

func new_struct_def(name string, fields []FieldInfo, is_sum bool, variants []string) StructDef {
	return StructDef{Name: name, Fields: fields, Is_sum: is_sum, Variant_names: variants, Variant_field_counts: []int64{int64(0)}, Variant_field_names: []string{""}, Variant_field_types: []int64{int64(0)}, Variant_field_offsets: []int64{int64(0)}, Generic_count: int64(0), Generic_name_0: "", Generic_name_1: "", Generic_name_2: ""}
}

func new_generic_struct_def(name string, fields []FieldInfo, gc int64, gn0 string, gn1 string, gn2 string) StructDef {
	return StructDef{Name: name, Fields: fields, Is_sum: false, Variant_names: []string{""}, Variant_field_counts: []int64{int64(0)}, Variant_field_names: []string{""}, Variant_field_types: []int64{int64(0)}, Variant_field_offsets: []int64{int64(0)}, Generic_count: gc, Generic_name_0: gn0, Generic_name_1: gn1, Generic_name_2: gn2}
}

func struct_is_generic(def StructDef) bool {
	return (def.Generic_count > int64(0))
}

func new_sum_def(name string, variants []string, vf_counts []int64, vf_names []string, vf_types []int64, vf_offsets []int64) StructDef {
	return StructDef{Name: name, Fields: []FieldInfo{FieldInfo{Name: "", Type_id: int64(0)}}, Is_sum: true, Variant_names: variants, Variant_field_counts: vf_counts, Variant_field_names: vf_names, Variant_field_types: vf_types, Variant_field_offsets: vf_offsets, Generic_count: int64(0), Generic_name_0: "", Generic_name_1: "", Generic_name_2: ""}
}

func variant_data_count(def StructDef, tag int64) int64 {
	if ((tag <= int64(0)) || (tag >= int64(len(def.Variant_field_counts)))) {
		return int64(0)
	}
	return def.Variant_field_counts[tag]
}

func variant_field_name(def StructDef, tag int64, fi int64) string {
	if ((tag <= int64(0)) || (tag >= int64(len(def.Variant_field_offsets)))) {
		return ""
	}
	off := def.Variant_field_offsets[tag]
	_ = off
	idx := (off + fi)
	_ = idx
	if (idx >= int64(len(def.Variant_field_names))) {
		return ""
	}
	return def.Variant_field_names[idx]
}

func variant_field_type(def StructDef, tag int64, fi int64) int64 {
	if ((tag <= int64(0)) || (tag >= int64(len(def.Variant_field_offsets)))) {
		return int64(0)
	}
	off := def.Variant_field_offsets[tag]
	_ = off
	idx := (off + fi)
	_ = idx
	if (idx >= int64(len(def.Variant_field_types))) {
		return int64(0)
	}
	return def.Variant_field_types[idx]
}

func sum_alloc_size(def StructDef) int64 {
	max_fields := int64(0)
	_ = max_fields
	i := int64(1)
	_ = i
	for (i < int64(len(def.Variant_field_counts))) {
		cnt := def.Variant_field_counts[i]
		_ = cnt
		slots := int64(0)
		_ = slots
		off := def.Variant_field_offsets[i]
		_ = off
		fi := int64(0)
		_ = fi
		for (fi < cnt) {
			idx := (off + fi)
			_ = idx
			if ((idx < int64(len(def.Variant_field_types))) && (def.Variant_field_types[idx] == int64(12))) {
				slots = (slots + int64(2))
			} else {
				slots = (slots + int64(1))
			}
			fi = (fi + int64(1))
		}
		if (slots > max_fields) {
			max_fields = slots
		}
		i = (i + int64(1))
	}
	return (int64(1) + max_fields)
}

func sum_has_data(def StructDef) bool {
	i := int64(1)
	_ = i
	for (i < int64(len(def.Variant_field_counts))) {
		if (def.Variant_field_counts[i] > int64(0)) {
			return true
		}
		i = (i + int64(1))
	}
	return false
}

func _sentinel_fn_sig() FnSig {
	return FnSig{Name: "", Param_names: []string{""}, Param_types: []int64{int64(0)}, Return_type: int64(0), Error_type: int64(0), Generic_count: int64(0), Generic_name_0: "", Generic_name_1: "", Generic_name_2: "", Generic_bound_0: "", Generic_bound_1: "", Generic_bound_2: "", Token_start: int64(0), Has_io: false, Has_fs: false, Has_net: false, Has_ffi: false, Has_async: false, User_effects: []string{""}}
}

func new_fn_sig(name string, pnames []string, ptypes []int64, ret int64, err int64) FnSig {
	return FnSig{Name: name, Param_names: pnames, Param_types: ptypes, Return_type: ret, Error_type: err, Generic_count: int64(0), Generic_name_0: "", Generic_name_1: "", Generic_name_2: "", Generic_bound_0: "", Generic_bound_1: "", Generic_bound_2: "", Token_start: int64(0), Has_io: false, Has_fs: false, Has_net: false, Has_ffi: false, Has_async: false, User_effects: []string{""}}
}

func new_generic_fn_sig(name string, pnames []string, ptypes []int64, ret int64, err int64, gc int64, gn0 string, gn1 string, gn2 string, gb0 string, gb1 string, gb2 string, tstart int64) FnSig {
	return FnSig{Name: name, Param_names: pnames, Param_types: ptypes, Return_type: ret, Error_type: err, Generic_count: gc, Generic_name_0: gn0, Generic_name_1: gn1, Generic_name_2: gn2, Generic_bound_0: gb0, Generic_bound_1: gb1, Generic_bound_2: gb2, Token_start: tstart, Has_io: false, Has_fs: false, Has_net: false, Has_ffi: false, Has_async: false, User_effects: []string{""}}
}

func fn_is_generic(sig FnSig) bool {
	return (sig.Generic_count > int64(0))
}

func fn_has_effect(sig FnSig, effect string) bool {
	if (effect == "Io") {
		return ((sig.Has_io || sig.Has_fs) || sig.Has_net)
	}
	if (effect == "Fs") {
		return sig.Has_fs
	}
	if (effect == "Net") {
		return sig.Has_net
	}
	if (effect == "Ffi") {
		return sig.Has_ffi
	}
	if (effect == "Async") {
		return sig.Has_async
	}
	i := int64(1)
	_ = i
	for (i < int64(len(sig.User_effects))) {
		if (sig.User_effects[i] == effect) {
			return true
		}
		i = (i + int64(1))
	}
	return false
}

func fn_has_effect_t(sig FnSig, effect string, treg TraitRegistry) bool {
	if fn_has_effect(sig, effect) {
		return true
	}
	i := int64(1)
	_ = i
	for (i < int64(len(sig.User_effects))) {
		ue := sig.User_effects[i]
		_ = ue
		tdef := treg_find_trait(treg, ue)
		_ = tdef
		if (tdef.Name != "") {
			pi := int64(1)
			_ = pi
			for (pi < int64(len(tdef.Parent_traits))) {
				if (tdef.Parent_traits[pi] == effect) {
					return true
				}
				pi = (pi + int64(1))
			}
		}
		i = (i + int64(1))
	}
	return false
}

func fn_is_pure(sig FnSig) bool {
	return ((((((sig.Has_io == false) && (sig.Has_fs == false)) && (sig.Has_net == false)) && (sig.Has_ffi == false)) && (sig.Has_async == false)) && (int64(len(sig.User_effects)) <= int64(1)))
}

func _builtin_effect(name string) string {
	if ((((((name == "_ariaReadFile") || (name == "_ariaWriteFile")) || (name == "_ariaAppendFile")) || (name == "_ariaWriteBinaryFile")) || (name == "_ariaListDir")) || (name == "_ariaIsDir")) {
		return "Fs"
	}
	if ((name == "_ariaExec") || (name == "_ariaGetenv")) {
		return "Io"
	}
	if ((((((name == "_ariaTcpConnect") || (name == "_ariaTcpListen")) || (name == "_ariaTcpAccept")) || (name == "_ariaTcpRead")) || (name == "_ariaTcpWrite")) || (name == "_ariaTcpClose")) {
		return "Net"
	}
	if ((((name == "_ariaPgConnect") || (name == "_ariaPgQuery")) || (name == "_ariaPgExec")) || (name == "_ariaPgClose")) {
		return "Net"
	}
	if (((name == "_ariaHttpGet") || (name == "_ariaHttpPost")) || (name == "_ariaHttpRequest")) {
		return "Net"
	}
	if ((((((((name == "_ariaSpawn") || (name == "_ariaTaskAwait")) || (name == "_ariaChanSend")) || (name == "_ariaChanRecv")) || (name == "_ariaChanNew")) || (name == "_ariaMutexNew")) || (name == "_ariaMutexLock")) || (name == "_ariaMutexUnlock")) {
		return "Async"
	}
	if ((((((name == "_ariaSpawn2") || (name == "_ariaTaskAwait2")) || (name == "_ariaTaskDone")) || (name == "_ariaTaskCancel")) || (name == "_ariaTaskResult")) || (name == "_ariaCancelCheck")) {
		return "Async"
	}
	if (((name == "_ariaChanTryRecv") || (name == "_ariaChanSelect")) || (name == "_ariaChanClose")) {
		return "Async"
	}
	if (((((name == "_ariaRWMutexNew") || (name == "_ariaRWMutexRlock")) || (name == "_ariaRWMutexRunlock")) || (name == "_ariaRWMutexWlock")) || (name == "_ariaRWMutexWunlock")) {
		return "Async"
	}
	if ((((name == "_ariaWgNew") || (name == "_ariaWgAdd")) || (name == "_ariaWgDone")) || (name == "_ariaWgWait")) {
		return "Async"
	}
	if ((name == "_ariaOnceNew") || (name == "_ariaOnceCall")) {
		return "Async"
	}
	if ((((name == "_ariaCancelNew") || (name == "_ariaCancelChild")) || (name == "_ariaCancelTrigger")) || (name == "_ariaCancelIsTriggered")) {
		return "Async"
	}
	return ""
}

func sig_set_effects(sig FnSig, io bool, fs bool, net bool, ffi bool, async_e bool) FnSig {
	return FnSig{Name: sig.Name, Param_names: sig.Param_names, Param_types: sig.Param_types, Return_type: sig.Return_type, Error_type: sig.Error_type, Generic_count: sig.Generic_count, Generic_name_0: sig.Generic_name_0, Generic_name_1: sig.Generic_name_1, Generic_name_2: sig.Generic_name_2, Generic_bound_0: sig.Generic_bound_0, Generic_bound_1: sig.Generic_bound_1, Generic_bound_2: sig.Generic_bound_2, Token_start: sig.Token_start, Has_io: io, Has_fs: fs, Has_net: net, Has_ffi: ffi, Has_async: async_e, User_effects: sig.User_effects}
}

func sig_set_user_effects(sig FnSig, effects []string) FnSig {
	return FnSig{Name: sig.Name, Param_names: sig.Param_names, Param_types: sig.Param_types, Return_type: sig.Return_type, Error_type: sig.Error_type, Generic_count: sig.Generic_count, Generic_name_0: sig.Generic_name_0, Generic_name_1: sig.Generic_name_1, Generic_name_2: sig.Generic_name_2, Generic_bound_0: sig.Generic_bound_0, Generic_bound_1: sig.Generic_bound_1, Generic_bound_2: sig.Generic_bound_2, Token_start: sig.Token_start, Has_io: sig.Has_io, Has_fs: sig.Has_fs, Has_net: sig.Has_net, Has_ffi: sig.Has_ffi, Has_async: sig.Has_async, User_effects: effects}
}

func _sentinel_mono_spec() MonoSpec {
	return MonoSpec{Generic_name: "", Specialized_name: "", Type_arg_0: int64(0), Type_arg_1: int64(0), Type_arg_2: int64(0)}
}

func new_registry() TypeRegistry {
	return TypeRegistry{Struct_defs: []StructDef{_sentinel_struct_def()}, Fn_sigs: []FnSig{_sentinel_fn_sig()}, Mono_specs: []MonoSpec{_sentinel_mono_spec()}, Deprecated_names: []string{""}, Deprecated_msgs: []string{""}, Cold_names: []string{""}, Fn_value_refs: []string{""}, Bound_fn_value_refs: []string{""}}
}

func reg_add_struct(reg TypeRegistry, def StructDef) TypeRegistry {
	return TypeRegistry{Struct_defs: append(reg.Struct_defs, def), Fn_sigs: reg.Fn_sigs, Mono_specs: reg.Mono_specs, Deprecated_names: reg.Deprecated_names, Deprecated_msgs: reg.Deprecated_msgs, Cold_names: reg.Cold_names, Fn_value_refs: reg.Fn_value_refs, Bound_fn_value_refs: reg.Bound_fn_value_refs}
}

func reg_add_fn(reg TypeRegistry, sig FnSig) TypeRegistry {
	return TypeRegistry{Struct_defs: reg.Struct_defs, Fn_sigs: append(reg.Fn_sigs, sig), Mono_specs: reg.Mono_specs, Deprecated_names: reg.Deprecated_names, Deprecated_msgs: reg.Deprecated_msgs, Cold_names: reg.Cold_names, Fn_value_refs: reg.Fn_value_refs, Bound_fn_value_refs: reg.Bound_fn_value_refs}
}

func reg_add_mono(reg TypeRegistry, spec MonoSpec) TypeRegistry {
	return TypeRegistry{Struct_defs: reg.Struct_defs, Fn_sigs: reg.Fn_sigs, Mono_specs: append(reg.Mono_specs, spec), Deprecated_names: reg.Deprecated_names, Deprecated_msgs: reg.Deprecated_msgs, Cold_names: reg.Cold_names, Fn_value_refs: reg.Fn_value_refs, Bound_fn_value_refs: reg.Bound_fn_value_refs}
}

func reg_add_deprecated(reg TypeRegistry, name string, msg string) TypeRegistry {
	return TypeRegistry{Struct_defs: reg.Struct_defs, Fn_sigs: reg.Fn_sigs, Mono_specs: reg.Mono_specs, Deprecated_names: append(reg.Deprecated_names, name), Deprecated_msgs: append(reg.Deprecated_msgs, msg), Cold_names: reg.Cold_names, Fn_value_refs: reg.Fn_value_refs, Bound_fn_value_refs: reg.Bound_fn_value_refs}
}

func reg_add_cold(reg TypeRegistry, name string) TypeRegistry {
	return TypeRegistry{Struct_defs: reg.Struct_defs, Fn_sigs: reg.Fn_sigs, Mono_specs: reg.Mono_specs, Deprecated_names: reg.Deprecated_names, Deprecated_msgs: reg.Deprecated_msgs, Cold_names: append(reg.Cold_names, name), Fn_value_refs: reg.Fn_value_refs, Bound_fn_value_refs: reg.Bound_fn_value_refs}
}

func reg_add_fn_value_ref(reg TypeRegistry, fname string) TypeRegistry {
	i := int64(1)
	_ = i
	for (i < int64(len(reg.Fn_value_refs))) {
		if (reg.Fn_value_refs[i] == fname) {
			return reg
		}
		i = (i + int64(1))
	}
	return TypeRegistry{Struct_defs: reg.Struct_defs, Fn_sigs: reg.Fn_sigs, Mono_specs: reg.Mono_specs, Deprecated_names: reg.Deprecated_names, Deprecated_msgs: reg.Deprecated_msgs, Cold_names: reg.Cold_names, Fn_value_refs: append(reg.Fn_value_refs, fname), Bound_fn_value_refs: reg.Bound_fn_value_refs}
}

func reg_add_bound_fn_value_ref(reg TypeRegistry, fname string) TypeRegistry {
	i := int64(1)
	_ = i
	for (i < int64(len(reg.Bound_fn_value_refs))) {
		if (reg.Bound_fn_value_refs[i] == fname) {
			return reg
		}
		i = (i + int64(1))
	}
	return TypeRegistry{Struct_defs: reg.Struct_defs, Fn_sigs: reg.Fn_sigs, Mono_specs: reg.Mono_specs, Deprecated_names: reg.Deprecated_names, Deprecated_msgs: reg.Deprecated_msgs, Cold_names: reg.Cold_names, Fn_value_refs: reg.Fn_value_refs, Bound_fn_value_refs: append(reg.Bound_fn_value_refs, fname)}
}

func reg_is_deprecated(reg TypeRegistry, name string) bool {
	i := int64(1)
	_ = i
	for (i < int64(len(reg.Deprecated_names))) {
		if (reg.Deprecated_names[i] == name) {
			return true
		}
		i = (i + int64(1))
	}
	return false
}

func reg_deprecated_msg(reg TypeRegistry, name string) string {
	i := int64(1)
	_ = i
	for (i < int64(len(reg.Deprecated_names))) {
		if (reg.Deprecated_names[i] == name) {
			return reg.Deprecated_msgs[i]
		}
		i = (i + int64(1))
	}
	return ""
}

func reg_is_cold(reg TypeRegistry, name string) bool {
	i := int64(1)
	_ = i
	for (i < int64(len(reg.Cold_names))) {
		if (reg.Cold_names[i] == name) {
			return true
		}
		i = (i + int64(1))
	}
	return false
}

func reg_find_mono(reg TypeRegistry, specialized_name string) MonoSpec {
	i := int64(1)
	_ = i
	for (i < int64(len(reg.Mono_specs))) {
		if (reg.Mono_specs[i].Specialized_name == specialized_name) {
			return reg.Mono_specs[i]
		}
		i = (i + int64(1))
	}
	return _sentinel_mono_spec()
}

func reg_find_mono_by_generic(reg TypeRegistry, generic_name string, type_arg_0 int64) MonoSpec {
	i := int64(1)
	_ = i
	for (i < int64(len(reg.Mono_specs))) {
		sp := reg.Mono_specs[i]
		_ = sp
		if ((sp.Generic_name == generic_name) && (sp.Type_arg_0 == type_arg_0)) {
			return sp
		}
		i = (i + int64(1))
	}
	return _sentinel_mono_spec()
}

func reg_find_struct(reg TypeRegistry, name string) StructDef {
	i := int64(1)
	_ = i
	for (i < int64(len(reg.Struct_defs))) {
		if (reg.Struct_defs[i].Name == name) {
			return reg.Struct_defs[i]
		}
		i = (i + int64(1))
	}
	return _sentinel_struct_def()
}

func reg_find_fn(reg TypeRegistry, name string) FnSig {
	i := int64(1)
	_ = i
	for (i < int64(len(reg.Fn_sigs))) {
		if (reg.Fn_sigs[i].Name == name) {
			return reg.Fn_sigs[i]
		}
		i = (i + int64(1))
	}
	return _sentinel_fn_sig()
}

func reg_find_field(reg TypeRegistry, struct_name string, field_name string) FieldInfo {
	def := reg_find_struct(reg, struct_name)
	_ = def
	if (def.Name == "") {
		return _sentinel_field()
	}
	i := int64(1)
	_ = i
	for (i < int64(len(def.Fields))) {
		if (def.Fields[i].Name == field_name) {
			return def.Fields[i]
		}
		i = (i + int64(1))
	}
	return _sentinel_field()
}

func reg_has_variant(reg TypeRegistry, struct_name string, variant_name string) bool {
	def := reg_find_struct(reg, struct_name)
	_ = def
	if (def.Name == "") {
		return false
	}
	if (def.Is_sum == false) {
		return false
	}
	i := int64(1)
	_ = i
	for (i < int64(len(def.Variant_names))) {
		if (def.Variant_names[i] == variant_name) {
			return true
		}
		i = (i + int64(1))
	}
	return false
}

var TY_NONE = int64(0)

var TY_I64 = int64(1)

var TY_I32 = int64(2)

var TY_I16 = int64(3)

var TY_I8 = int64(4)

var TY_U64 = int64(5)

var TY_U32 = int64(6)

var TY_U16 = int64(7)

var TY_U8 = int64(8)

var TY_F64 = int64(9)

var TY_F32 = int64(10)

var TY_BOOL = int64(11)

var TY_STR = int64(12)

var TY_CHAR = int64(13)

var TY_UNIT = int64(14)

var TY_NEVER = int64(15)

var TY_UNKNOWN = int64(16)

var TY_CHAN = int64(17)

var TY_TASK = int64(18)

var TY_CHAN_SEND = int64(19)

var TY_CHAN_RECV = int64(20)

var TY_STRUCT_BASE = int64(10000)

