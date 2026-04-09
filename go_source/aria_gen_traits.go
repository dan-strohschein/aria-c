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

type TraitDef struct {
	Name string
	Method_names []string
	Parent_traits []string
	Assoc_type_names []string
	Default_starts []int64
	Default_ends []int64
}

type ImplDef struct {
	Trait_name string
	Type_name string
	Assoc_type_names []string
	Assoc_type_values []string
}

type TraitRegistry struct {
	Trait_defs []TraitDef
	Impl_defs []ImplDef
	Impl_pkgs []string
	Trait_index StrMap
	Impl_index StrMap
}

func _sentinel_trait_def() TraitDef {
	return TraitDef{Name: "", Method_names: []string{""}, Parent_traits: []string{""}, Assoc_type_names: []string{""}, Default_starts: []int64{int64(0)}, Default_ends: []int64{int64(0)}}
}

func _sentinel_impl_def() ImplDef {
	return ImplDef{Trait_name: "", Type_name: "", Assoc_type_names: []string{""}, Assoc_type_values: []string{""}}
}

func _mk_impl(tname string, typename string) ImplDef {
	return ImplDef{Trait_name: tname, Type_name: typename, Assoc_type_names: []string{""}, Assoc_type_values: []string{""}}
}

func new_trait_registry() TraitRegistry {
	return TraitRegistry{Trait_defs: []TraitDef{_sentinel_trait_def()}, Impl_defs: []ImplDef{_sentinel_impl_def()}, Impl_pkgs: []string{""}, Trait_index: strmap_new(), Impl_index: strmap_new()}
}

func _impl_key(type_name string, trait_name string) string {
	return ((type_name + ":") + trait_name)
}

func treg_add_trait(reg TraitRegistry, def TraitDef) TraitRegistry {
	idx := int64(len(reg.Trait_defs))
	_ = idx
	return TraitRegistry{Trait_defs: append(reg.Trait_defs, def), Impl_defs: reg.Impl_defs, Impl_pkgs: reg.Impl_pkgs, Trait_index: strmap_put(reg.Trait_index, def.Name, idx), Impl_index: reg.Impl_index}
}

func treg_add_impl(reg TraitRegistry, def ImplDef) TraitRegistry {
	idx := int64(len(reg.Impl_defs))
	_ = idx
	return TraitRegistry{Trait_defs: reg.Trait_defs, Impl_defs: append(reg.Impl_defs, def), Impl_pkgs: append(reg.Impl_pkgs, ""), Trait_index: reg.Trait_index, Impl_index: strmap_put(reg.Impl_index, _impl_key(def.Type_name, def.Trait_name), idx)}
}

func treg_add_impl_pkg(reg TraitRegistry, def ImplDef, pkg string) TraitRegistry {
	idx := int64(len(reg.Impl_defs))
	_ = idx
	return TraitRegistry{Trait_defs: reg.Trait_defs, Impl_defs: append(reg.Impl_defs, def), Impl_pkgs: append(reg.Impl_pkgs, pkg), Trait_index: reg.Trait_index, Impl_index: strmap_put(reg.Impl_index, _impl_key(def.Type_name, def.Trait_name), idx)}
}

func _pkg_from_path(path string) string {
	if (path == "") {
		return ""
	}
	if (path == "<builtin>") {
		return "<builtin>"
	}
	last_slash := (int64(0) - int64(1))
	_ = last_slash
	i := int64(0)
	_ = i
	for (i < int64(len(path))) {
		if (string(path[i]) == "/") {
			last_slash = i
		}
		i = (i + int64(1))
	}
	if (last_slash < int64(0)) {
		return ""
	}
	return path[int64(0):last_slash]
}

func treg_impl_pkg(reg TraitRegistry, type_name string, trait_name string) string {
	i := int64(1)
	_ = i
	for (i < int64(len(reg.Impl_defs))) {
		imp := reg.Impl_defs[i]
		_ = imp
		if ((imp.Type_name == type_name) && (imp.Trait_name == trait_name)) {
			if (i < int64(len(reg.Impl_pkgs))) {
				return reg.Impl_pkgs[i]
			}
			return ""
		}
		i = (i + int64(1))
	}
	return ""
}

func treg_find_trait(reg TraitRegistry, name string) TraitDef {
	idx := strmap_get(reg.Trait_index, name)
	_ = idx
	if ((idx > int64(0)) && (idx < int64(len(reg.Trait_defs)))) {
		return reg.Trait_defs[idx]
	}
	i := int64(1)
	_ = i
	for (i < int64(len(reg.Trait_defs))) {
		if (reg.Trait_defs[i].Name == name) {
			return reg.Trait_defs[i]
		}
		i = (i + int64(1))
	}
	return _sentinel_trait_def()
}

func treg_has_impl(reg TraitRegistry, type_name string, trait_name string) bool {
	idx := strmap_get(reg.Impl_index, _impl_key(type_name, trait_name))
	_ = idx
	if (idx > int64(0)) {
		return true
	}
	i := int64(1)
	_ = i
	for (i < int64(len(reg.Impl_defs))) {
		imp := reg.Impl_defs[i]
		_ = imp
		if ((imp.Type_name == type_name) && (imp.Trait_name == trait_name)) {
			return true
		}
		i = (i + int64(1))
	}
	return false
}

func treg_find_impl(reg TraitRegistry, type_name string, trait_name string) ImplDef {
	idx := strmap_get(reg.Impl_index, _impl_key(type_name, trait_name))
	_ = idx
	if ((idx > int64(0)) && (idx < int64(len(reg.Impl_defs)))) {
		return reg.Impl_defs[idx]
	}
	i := int64(1)
	_ = i
	for (i < int64(len(reg.Impl_defs))) {
		imp := reg.Impl_defs[i]
		_ = imp
		if ((imp.Type_name == type_name) && (imp.Trait_name == trait_name)) {
			return imp
		}
		i = (i + int64(1))
	}
	return _sentinel_impl_def()
}

func treg_find_assoc_value(reg TraitRegistry, type_name string, trait_name string, assoc_name string) string {
	imp := treg_find_impl(reg, type_name, trait_name)
	_ = imp
	if (imp.Trait_name == "") {
		return ""
	}
	i := int64(1)
	_ = i
	for (i < int64(len(imp.Assoc_type_names))) {
		if (imp.Assoc_type_names[i] == assoc_name) {
			return imp.Assoc_type_values[i]
		}
		i = (i + int64(1))
	}
	return ""
}

func treg_types_implementing(reg TraitRegistry, bound string) []string {
	types := []string{""}
	_ = types
	i := int64(1)
	_ = i
	for (i < int64(len(reg.Impl_defs))) {
		imp := reg.Impl_defs[i]
		_ = imp
		if (imp.Trait_name == bound) {
			if (int64(len(imp.Type_name)) > int64(0)) {
				first := imp.Type_name[int64(0):int64(1)]
				_ = first
				if (first != "*") {
					found := false
					_ = found
					j := int64(1)
					_ = j
					for (j < int64(len(types))) {
						if (types[j] == imp.Type_name) {
							found = true
						}
						j = (j + int64(1))
					}
					if (found == false) {
						types = append(types, imp.Type_name)
					}
				}
			}
		}
		i = (i + int64(1))
	}
	return types
}

func _add_numeric_impl(reg TraitRegistry, type_name string) TraitRegistry {
	rr := reg
	_ = rr
	rr = treg_add_impl(rr, _mk_impl("Eq", type_name))
	rr = treg_add_impl(rr, _mk_impl("Ord", type_name))
	rr = treg_add_impl(rr, _mk_impl("Hash", type_name))
	rr = treg_add_impl(rr, _mk_impl("Add", type_name))
	rr = treg_add_impl(rr, _mk_impl("Sub", type_name))
	rr = treg_add_impl(rr, _mk_impl("Mul", type_name))
	rr = treg_add_impl(rr, _mk_impl("Div", type_name))
	rr = treg_add_impl(rr, _mk_impl("Mod", type_name))
	rr = treg_add_impl(rr, _mk_impl("Numeric", type_name))
	return rr
}

func _add_float_impl(reg TraitRegistry, type_name string) TraitRegistry {
	rr := reg
	_ = rr
	rr = treg_add_impl(rr, _mk_impl("Eq", type_name))
	rr = treg_add_impl(rr, _mk_impl("Ord", type_name))
	rr = treg_add_impl(rr, _mk_impl("Add", type_name))
	rr = treg_add_impl(rr, _mk_impl("Sub", type_name))
	rr = treg_add_impl(rr, _mk_impl("Mul", type_name))
	rr = treg_add_impl(rr, _mk_impl("Div", type_name))
	rr = treg_add_impl(rr, _mk_impl("Mod", type_name))
	rr = treg_add_impl(rr, _mk_impl("Numeric", type_name))
	return rr
}

func populate_builtins(reg TraitRegistry) TraitRegistry {
	rr := reg
	_ = rr
	rr = treg_add_trait(rr, TraitDef{Name: "Eq", Method_names: []string{"", "eq"}, Parent_traits: []string{""}, Assoc_type_names: []string{""}, Default_starts: []int64{int64(0)}, Default_ends: []int64{int64(0)}})
	rr = treg_add_trait(rr, TraitDef{Name: "Ord", Method_names: []string{"", "cmp"}, Parent_traits: []string{""}, Assoc_type_names: []string{""}, Default_starts: []int64{int64(0)}, Default_ends: []int64{int64(0)}})
	rr = treg_add_trait(rr, TraitDef{Name: "Hash", Method_names: []string{"", "hash"}, Parent_traits: []string{""}, Assoc_type_names: []string{""}, Default_starts: []int64{int64(0)}, Default_ends: []int64{int64(0)}})
	rr = treg_add_trait(rr, TraitDef{Name: "Add", Method_names: []string{"", "add"}, Parent_traits: []string{""}, Assoc_type_names: []string{""}, Default_starts: []int64{int64(0)}, Default_ends: []int64{int64(0)}})
	rr = treg_add_trait(rr, TraitDef{Name: "Sub", Method_names: []string{"", "sub"}, Parent_traits: []string{""}, Assoc_type_names: []string{""}, Default_starts: []int64{int64(0)}, Default_ends: []int64{int64(0)}})
	rr = treg_add_trait(rr, TraitDef{Name: "Mul", Method_names: []string{"", "mul"}, Parent_traits: []string{""}, Assoc_type_names: []string{""}, Default_starts: []int64{int64(0)}, Default_ends: []int64{int64(0)}})
	rr = treg_add_trait(rr, TraitDef{Name: "Div", Method_names: []string{"", "div"}, Parent_traits: []string{""}, Assoc_type_names: []string{""}, Default_starts: []int64{int64(0)}, Default_ends: []int64{int64(0)}})
	rr = treg_add_trait(rr, TraitDef{Name: "Mod", Method_names: []string{"", "mod"}, Parent_traits: []string{""}, Assoc_type_names: []string{""}, Default_starts: []int64{int64(0)}, Default_ends: []int64{int64(0)}})
	rr = treg_add_trait(rr, TraitDef{Name: "Numeric", Method_names: []string{""}, Parent_traits: []string{""}, Assoc_type_names: []string{""}, Default_starts: []int64{int64(0)}, Default_ends: []int64{int64(0)}})
	rr = treg_add_trait(rr, TraitDef{Name: "Clone", Method_names: []string{"", "clone"}, Parent_traits: []string{""}, Assoc_type_names: []string{""}, Default_starts: []int64{int64(0)}, Default_ends: []int64{int64(0)}})
	rr = treg_add_trait(rr, TraitDef{Name: "Debug", Method_names: []string{"", "debug_str"}, Parent_traits: []string{""}, Assoc_type_names: []string{""}, Default_starts: []int64{int64(0)}, Default_ends: []int64{int64(0)}})
	rr = treg_add_trait(rr, TraitDef{Name: "Display", Method_names: []string{"", "to_str"}, Parent_traits: []string{""}, Assoc_type_names: []string{""}, Default_starts: []int64{int64(0)}, Default_ends: []int64{int64(0)}})
	rr = treg_add_trait(rr, TraitDef{Name: "Convert", Method_names: []string{"", "to"}, Parent_traits: []string{""}, Assoc_type_names: []string{""}, Default_starts: []int64{int64(0)}, Default_ends: []int64{int64(0)}})
	rr = treg_add_trait(rr, TraitDef{Name: "TryConvert", Method_names: []string{"", "try_to"}, Parent_traits: []string{""}, Assoc_type_names: []string{""}, Default_starts: []int64{int64(0)}, Default_ends: []int64{int64(0)}})
	rr = treg_add_trait(rr, TraitDef{Name: "Drop", Method_names: []string{"", "drop"}, Parent_traits: []string{""}, Assoc_type_names: []string{""}, Default_starts: []int64{int64(0)}, Default_ends: []int64{int64(0)}})
	rr = treg_add_trait(rr, TraitDef{Name: "Send", Method_names: []string{""}, Parent_traits: []string{""}, Assoc_type_names: []string{""}, Default_starts: []int64{int64(0)}, Default_ends: []int64{int64(0)}})
	rr = treg_add_trait(rr, TraitDef{Name: "Share", Method_names: []string{""}, Parent_traits: []string{""}, Assoc_type_names: []string{""}, Default_starts: []int64{int64(0)}, Default_ends: []int64{int64(0)}})
	rr = treg_add_trait(rr, TraitDef{Name: "Transient", Method_names: []string{""}, Parent_traits: []string{""}, Assoc_type_names: []string{""}, Default_starts: []int64{int64(0)}, Default_ends: []int64{int64(0)}})
	rr = treg_add_trait(rr, TraitDef{Name: "Permanent", Method_names: []string{""}, Parent_traits: []string{""}, Assoc_type_names: []string{""}, Default_starts: []int64{int64(0)}, Default_ends: []int64{int64(0)}})
	rr = treg_add_trait(rr, TraitDef{Name: "UserFault", Method_names: []string{""}, Parent_traits: []string{""}, Assoc_type_names: []string{""}, Default_starts: []int64{int64(0)}, Default_ends: []int64{int64(0)}})
	rr = treg_add_trait(rr, TraitDef{Name: "SystemFault", Method_names: []string{""}, Parent_traits: []string{""}, Assoc_type_names: []string{""}, Default_starts: []int64{int64(0)}, Default_ends: []int64{int64(0)}})
	rr = treg_add_trait(rr, TraitDef{Name: "Retryable", Method_names: []string{"", "retryAfter", "maxRetries"}, Parent_traits: []string{"", "Transient"}, Assoc_type_names: []string{""}, Default_starts: []int64{int64(0)}, Default_ends: []int64{int64(0)}})
	rr = _add_numeric_impl(rr, "i64")
	rr = _add_numeric_impl(rr, "i32")
	rr = _add_numeric_impl(rr, "i16")
	rr = _add_numeric_impl(rr, "i8")
	rr = _add_numeric_impl(rr, "u64")
	rr = _add_numeric_impl(rr, "u32")
	rr = _add_numeric_impl(rr, "u16")
	rr = _add_numeric_impl(rr, "u8")
	rr = _add_float_impl(rr, "f64")
	rr = _add_float_impl(rr, "f32")
	rr = treg_add_impl(rr, _mk_impl("Eq", "bool"))
	rr = treg_add_impl(rr, _mk_impl("Hash", "bool"))
	rr = treg_add_impl(rr, _mk_impl("Eq", "str"))
	rr = treg_add_impl(rr, _mk_impl("Ord", "str"))
	rr = treg_add_impl(rr, _mk_impl("Hash", "str"))
	rr = treg_add_impl(rr, _mk_impl("Add", "str"))
	rr = treg_add_impl(rr, _mk_impl("Eq", "char"))
	rr = treg_add_impl(rr, _mk_impl("Ord", "char"))
	rr = treg_add_impl(rr, _mk_impl("Hash", "char"))
	rr = treg_add_impl(rr, _mk_impl("Send", "i64"))
	rr = treg_add_impl(rr, _mk_impl("Send", "i32"))
	rr = treg_add_impl(rr, _mk_impl("Send", "i16"))
	rr = treg_add_impl(rr, _mk_impl("Send", "i8"))
	rr = treg_add_impl(rr, _mk_impl("Send", "u64"))
	rr = treg_add_impl(rr, _mk_impl("Send", "u32"))
	rr = treg_add_impl(rr, _mk_impl("Send", "u16"))
	rr = treg_add_impl(rr, _mk_impl("Send", "u8"))
	rr = treg_add_impl(rr, _mk_impl("Send", "f64"))
	rr = treg_add_impl(rr, _mk_impl("Send", "f32"))
	rr = treg_add_impl(rr, _mk_impl("Send", "bool"))
	rr = treg_add_impl(rr, _mk_impl("Send", "str"))
	rr = treg_add_impl(rr, _mk_impl("Send", "char"))
	rr = treg_add_impl(rr, _mk_impl("Share", "i64"))
	rr = treg_add_impl(rr, _mk_impl("Share", "i32"))
	rr = treg_add_impl(rr, _mk_impl("Share", "i16"))
	rr = treg_add_impl(rr, _mk_impl("Share", "i8"))
	rr = treg_add_impl(rr, _mk_impl("Share", "u64"))
	rr = treg_add_impl(rr, _mk_impl("Share", "u32"))
	rr = treg_add_impl(rr, _mk_impl("Share", "u16"))
	rr = treg_add_impl(rr, _mk_impl("Share", "u8"))
	rr = treg_add_impl(rr, _mk_impl("Share", "f64"))
	rr = treg_add_impl(rr, _mk_impl("Share", "f32"))
	rr = treg_add_impl(rr, _mk_impl("Share", "bool"))
	rr = treg_add_impl(rr, _mk_impl("Share", "str"))
	rr = treg_add_impl(rr, _mk_impl("Share", "char"))
	return rr
}

func op_requires_trait(op string) string {
	if (op == "+") {
		return "Add"
	}
	if (op == "-") {
		return "Sub"
	}
	if (op == "*") {
		return "Mul"
	}
	if (op == "/") {
		return "Div"
	}
	if (op == "%") {
		return "Mod"
	}
	if ((op == "==") || (op == "!=")) {
		return "Eq"
	}
	if ((((op == "<") || (op == ">")) || (op == "<=")) || (op == ">=")) {
		return "Ord"
	}
	return ""
}

func type_supports_op(treg TraitRegistry, store TypeStore, type_id int64, op string) bool {
	if ((op == "&&") || (op == "||")) {
		return (type_id == TY_BOOL)
	}
	if (((((op == "&") || (op == "|")) || (op == "^")) || (op == "<<")) || (op == ">>")) {
		return is_integer(store, type_id)
	}
	trait_name := op_requires_trait(op)
	_ = trait_name
	if (trait_name == "") {
		return true
	}
	type_name := format_type(store, type_id)
	_ = type_name
	return treg_has_impl(treg, type_name, trait_name)
}

func binary_result_type(store TypeStore, left int64, op string) int64 {
	if ((((((op == "==") || (op == "!=")) || (op == "<")) || (op == ">")) || (op == "<=")) || (op == ">=")) {
		return TY_BOOL
	}
	if ((op == "&&") || (op == "||")) {
		return TY_BOOL
	}
	return left
}

