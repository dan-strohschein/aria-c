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

type Checker struct {
	Tokens []Token
	Pos int64
	Store TypeStore
	Registry TypeRegistry
	Treg TraitRegistry
	Diagnostics DiagnosticBag
	Table SymbolTable
	Current_scope int64
	File string
	Pool NodePool
	Env_names []string
	Env_types []int64
	Env_scopes []int64
	Fn_return_type int64
	Fn_error_type int64
	Fn_current_name string
	In_loop bool
	In_callee_slot bool
}

type CE struct {
	C Checker
	Type_id int64
}

type CheckResult struct {
	Store TypeStore
	Registry TypeRegistry
	Treg TraitRegistry
	Diagnostics DiagnosticBag
}

func _new_checker(tokens []Token, pool NodePool, table SymbolTable, file string) Checker {
	return Checker{Tokens: tokens, Pos: int64(0), Store: new_store(), Registry: new_registry(), Treg: populate_builtins(new_trait_registry()), Diagnostics: new_bag(), Table: table, Current_scope: int64(0), File: file, Pool: pool, Env_names: []string{""}, Env_types: []int64{int64(0)}, Env_scopes: []int64{int64(0)}, Fn_return_type: TY_UNIT, Fn_error_type: int64(0), Fn_current_name: "", In_loop: false, In_callee_slot: false}
}

func _ccur(c Checker) Token {
	if (c.Pos >= int64(len(c.Tokens))) {
		return Token{Kind: TokenKindTkEof{}, Text: "", Line: int64(0), Col: int64(0), Offset: int64(0)}
	}
	return c.Tokens[c.Pos]
}

func _ck2(c Checker) string {
	return token_name(_ccur(c).Kind)
}

func _cadv(c Checker) Checker {
	return Checker{Tokens: c.Tokens, Pos: (c.Pos + int64(1)), Store: c.Store, Registry: c.Registry, Treg: c.Treg, Diagnostics: c.Diagnostics, Table: c.Table, Current_scope: c.Current_scope, File: c.File, Pool: c.Pool, Env_names: c.Env_names, Env_types: c.Env_types, Env_scopes: c.Env_scopes, Fn_return_type: c.Fn_return_type, Fn_error_type: c.Fn_error_type, Fn_current_name: c.Fn_current_name, In_loop: c.In_loop, In_callee_slot: c.In_callee_slot}
}

func _cskip_pub(c Checker) Checker {
	cc := c
	_ = cc
	if (_ck2(cc) == "pub") {
		cc = _cadv(cc)
		if (_ck2(cc) == "(") {
			cc = _cadv(cc)
			if (_ck2(cc) == "IDENT") {
				cc = _cadv(cc)
			}
			if (_ck2(cc) == ")") {
				cc = _cadv(cc)
			}
		}
	}
	return cc
}

func _struct_name_for_tid(store TypeStore, reg TypeRegistry, tid int64) string {
	if (tid <= int64(16)) {
		return ""
	}
	if (tid >= int64(len(store.Types))) {
		return ""
	}
	info := store.Types[tid]
	_ = info
	if (type_kind_name(info.Kind) == "Named") {
		return info.Name
	}
	return ""
}

func _cset_pos(c Checker, pos int64) Checker {
	return Checker{Tokens: c.Tokens, Pos: pos, Store: c.Store, Registry: c.Registry, Treg: c.Treg, Diagnostics: c.Diagnostics, Table: c.Table, Current_scope: c.Current_scope, File: c.File, Pool: c.Pool, Env_names: c.Env_names, Env_types: c.Env_types, Env_scopes: c.Env_scopes, Fn_return_type: c.Fn_return_type, Fn_error_type: c.Fn_error_type, Fn_current_name: c.Fn_current_name, In_loop: c.In_loop, In_callee_slot: c.In_callee_slot}
}

func _cset_store(c Checker, store TypeStore) Checker {
	return Checker{Tokens: c.Tokens, Pos: c.Pos, Store: store, Registry: c.Registry, Treg: c.Treg, Diagnostics: c.Diagnostics, Table: c.Table, Current_scope: c.Current_scope, File: c.File, Pool: c.Pool, Env_names: c.Env_names, Env_types: c.Env_types, Env_scopes: c.Env_scopes, Fn_return_type: c.Fn_return_type, Fn_error_type: c.Fn_error_type, Fn_current_name: c.Fn_current_name, In_loop: c.In_loop, In_callee_slot: c.In_callee_slot}
}

func _cset_reg(c Checker, reg TypeRegistry) Checker {
	return Checker{Tokens: c.Tokens, Pos: c.Pos, Store: c.Store, Registry: reg, Treg: c.Treg, Diagnostics: c.Diagnostics, Table: c.Table, Current_scope: c.Current_scope, File: c.File, Pool: c.Pool, Env_names: c.Env_names, Env_types: c.Env_types, Env_scopes: c.Env_scopes, Fn_return_type: c.Fn_return_type, Fn_error_type: c.Fn_error_type, Fn_current_name: c.Fn_current_name, In_loop: c.In_loop, In_callee_slot: c.In_callee_slot}
}

func _cset_diag(c Checker, diag DiagnosticBag) Checker {
	return Checker{Tokens: c.Tokens, Pos: c.Pos, Store: c.Store, Registry: c.Registry, Treg: c.Treg, Diagnostics: diag, Table: c.Table, Current_scope: c.Current_scope, File: c.File, Pool: c.Pool, Env_names: c.Env_names, Env_types: c.Env_types, Env_scopes: c.Env_scopes, Fn_return_type: c.Fn_return_type, Fn_error_type: c.Fn_error_type, Fn_current_name: c.Fn_current_name, In_loop: c.In_loop, In_callee_slot: c.In_callee_slot}
}

func _cset_fn_ctx(c Checker, ret int64, err int64) Checker {
	return Checker{Tokens: c.Tokens, Pos: c.Pos, Store: c.Store, Registry: c.Registry, Treg: c.Treg, Diagnostics: c.Diagnostics, Table: c.Table, Current_scope: c.Current_scope, File: c.File, Pool: c.Pool, Env_names: c.Env_names, Env_types: c.Env_types, Env_scopes: c.Env_scopes, Fn_return_type: ret, Fn_error_type: err, Fn_current_name: c.Fn_current_name, In_loop: c.In_loop, In_callee_slot: c.In_callee_slot}
}

func _cset_treg(c Checker, treg TraitRegistry) Checker {
	return Checker{Tokens: c.Tokens, Pos: c.Pos, Store: c.Store, Registry: c.Registry, Treg: treg, Diagnostics: c.Diagnostics, Table: c.Table, Current_scope: c.Current_scope, File: c.File, Pool: c.Pool, Env_names: c.Env_names, Env_types: c.Env_types, Env_scopes: c.Env_scopes, Fn_return_type: c.Fn_return_type, Fn_error_type: c.Fn_error_type, Fn_current_name: c.Fn_current_name, In_loop: c.In_loop, In_callee_slot: c.In_callee_slot}
}

func _cset_loop(c Checker, in_loop bool) Checker {
	return Checker{Tokens: c.Tokens, Pos: c.Pos, Store: c.Store, Registry: c.Registry, Treg: c.Treg, Diagnostics: c.Diagnostics, Table: c.Table, Current_scope: c.Current_scope, File: c.File, Pool: c.Pool, Env_names: c.Env_names, Env_types: c.Env_types, Env_scopes: c.Env_scopes, Fn_return_type: c.Fn_return_type, Fn_error_type: c.Fn_error_type, Fn_current_name: c.Fn_current_name, In_loop: in_loop, In_callee_slot: c.In_callee_slot}
}

func _cset_callee_slot(c Checker, in_callee bool) Checker {
	return Checker{Tokens: c.Tokens, Pos: c.Pos, Store: c.Store, Registry: c.Registry, Treg: c.Treg, Diagnostics: c.Diagnostics, Table: c.Table, Current_scope: c.Current_scope, File: c.File, Pool: c.Pool, Env_names: c.Env_names, Env_types: c.Env_types, Env_scopes: c.Env_scopes, Fn_return_type: c.Fn_return_type, Fn_error_type: c.Fn_error_type, Fn_current_name: c.Fn_current_name, In_loop: c.In_loop, In_callee_slot: in_callee}
}

func _cspan(c Checker) Span {
	tok := _ccur(c)
	_ = tok
	return new_span(c.File, tok.Line, tok.Col, tok.Offset, int64(len(tok.Text)))
}

func _cskip_nl(c Checker) Checker {
	cc := c
	_ = cc
	for (_ck2(cc) == "NEWLINE") {
		cc = _cadv(cc)
	}
	return cc
}

func _cerror(c Checker, code string, msg string) Checker {
	span := _cspan(c)
	_ = span
	return _cset_diag(c, bag_add_error(c.Diagnostics, code, msg, span))
}

func _cwarning(c Checker, code string, msg string) Checker {
	span := _cspan(c)
	_ = span
	return _cset_diag(c, bag_add_warning(c.Diagnostics, code, msg, span))
}

func _env_add(c Checker, name string, type_id int64, sc int64) Checker {
	return Checker{Tokens: c.Tokens, Pos: c.Pos, Store: c.Store, Registry: c.Registry, Treg: c.Treg, Diagnostics: c.Diagnostics, Table: c.Table, Current_scope: c.Current_scope, File: c.File, Pool: c.Pool, Env_names: append(c.Env_names, name), Env_types: append(c.Env_types, type_id), Env_scopes: append(c.Env_scopes, sc), Fn_return_type: c.Fn_return_type, Fn_error_type: c.Fn_error_type, Fn_current_name: c.Fn_current_name, In_loop: c.In_loop, In_callee_slot: c.In_callee_slot}
}

func _env_lookup(c Checker, name string) int64 {
	i := (int64(len(c.Env_names)) - int64(1))
	_ = i
	for (i > int64(0)) {
		if (c.Env_names[i] == name) {
			return c.Env_types[i]
		}
		i = (i - int64(1))
	}
	return TY_UNKNOWN
}

func _env_trim(c Checker, saved_len int64) Checker {
	if (saved_len >= int64(len(c.Env_names))) {
		return c
	}
	names := []string{""}
	_ = names
	types := []int64{int64(0)}
	_ = types
	scopes := []int64{int64(0)}
	_ = scopes
	i := int64(1)
	_ = i
	for (i < saved_len) {
		names = append(names, c.Env_names[i])
		types = append(types, c.Env_types[i])
		scopes = append(scopes, c.Env_scopes[i])
		i = (i + int64(1))
	}
	return Checker{Tokens: c.Tokens, Pos: c.Pos, Store: c.Store, Registry: c.Registry, Treg: c.Treg, Diagnostics: c.Diagnostics, Table: c.Table, Current_scope: c.Current_scope, File: c.File, Pool: c.Pool, Env_names: names, Env_types: types, Env_scopes: scopes, Fn_return_type: c.Fn_return_type, Fn_error_type: c.Fn_error_type, Fn_current_name: c.Fn_current_name, In_loop: c.In_loop, In_callee_slot: c.In_callee_slot}
}

func _parse_type_ann(c Checker) CE {
	k := _ck2(c)
	_ = k
	if (k == "IDENT") {
		name := _ccur(c).Text
		_ = name
		if (name == "dyn") {
			cc := _cadv(c)
			_ = cc
			if (_ck2(cc) == "IDENT") {
				trait_name := _ccur(cc).Text
				_ = trait_name
				cc = _cadv(cc)
				tdef := treg_find_trait(cc.Treg, trait_name)
				_ = tdef
				if (tdef.Name == "") {
					cc = _cerror(cc, E0301, (("unknown trait '" + trait_name) + "' in dyn"))
				}
				new_st := store_add(cc.Store, TypeInfo{Kind: TypeKindTyTraitObject{}, Name: trait_name, Type_id: int64(0), Type_id2: int64(0), Param_start: int64(0), Param_count: int64(0)})
				_ = new_st
				cc = _cset_store(cc, new_st)
				return CE{C: cc, Type_id: store_last_id(cc.Store)}
			}
			return CE{C: cc, Type_id: TY_UNKNOWN}
		}
		cc := _cadv(c)
		_ = cc
		if (((_ck2(cc) == ":") && ((cc.Pos + int64(1)) < int64(len(cc.Tokens)))) && (token_name(cc.Tokens[(cc.Pos + int64(1))].Kind) == ":")) {
			cc = _cadv(_cadv(cc))
			if (_ck2(cc) == "IDENT") {
				assoc_name := _ccur(cc).Text
				_ = assoc_name
				cc = _cadv(cc)
				lookup_name := ((name + "_") + assoc_name)
				_ = lookup_name
				if (name == "Self") {
					impl_name := ""
					_ = impl_name
					ui := int64(0)
					_ = ui
					for (ui < int64(len(cc.Fn_current_name))) {
						ch := string(cc.Fn_current_name[ui])
						_ = ch
						if (ch == "_") {
							impl_name = cc.Fn_current_name[int64(0):ui]
						}
						ui = (ui + int64(1))
					}
					if (impl_name != "") {
						lookup_name = ((impl_name + "_") + assoc_name)
					}
				}
				tid := resolve_type_name(cc.Store, lookup_name)
				_ = tid
				if (tid != TY_UNKNOWN) {
					ainfo := store_get(cc.Store, tid)
					_ = ainfo
					if ((type_kind_name(ainfo.Kind) == "Named") && (ainfo.Type_id > int64(0))) {
						return CE{C: cc, Type_id: ainfo.Type_id}
					}
					return CE{C: cc, Type_id: tid}
				}
				return CE{C: cc, Type_id: TY_UNKNOWN}
			}
		}
		if (_ck2(cc) == "[") {
			cc = _cadv(cc)
			arg1 := TY_UNKNOWN
			_ = arg1
			arg2 := TY_UNKNOWN
			_ = arg2
			argc := int64(0)
			_ = argc
			for ((_ck2(cc) != "]") && (_ck2(cc) != "EOF")) {
				pr := _parse_type_ann(cc)
				_ = pr
				cc = pr.C
				if (argc == int64(0)) {
					arg1 = pr.Type_id
				}
				if (argc == int64(1)) {
					arg2 = pr.Type_id
				}
				argc = (argc + int64(1))
				if (_ck2(cc) == ",") {
					cc = _cadv(cc)
				}
			}
			if (_ck2(cc) == "]") {
				cc = _cadv(cc)
			}
			if ((name == "Result") && (argc >= int64(2))) {
				new_store := store_add(cc.Store, mk_result_type(arg1, arg2))
				_ = new_store
				cc = _cset_store(cc, new_store)
				return CE{C: cc, Type_id: store_last_id(cc.Store)}
			}
			if ((name == "Optional") && (argc >= int64(1))) {
				new_store := store_add(cc.Store, mk_optional_type(arg1))
				_ = new_store
				cc = _cset_store(cc, new_store)
				return CE{C: cc, Type_id: store_last_id(cc.Store)}
			}
			tid := resolve_type_name(cc.Store, name)
			_ = tid
			if (tid != TY_UNKNOWN) {
				tinfo := store_get(cc.Store, tid)
				_ = tinfo
				if (((type_kind_name(tinfo.Kind) == "TypeVar") && (tinfo.Param_count > int64(0))) && (argc >= int64(1))) {
					new_store := store_add(cc.Store, mk_applied_type(tid, arg1))
					_ = new_store
					cc = _cset_store(cc, new_store)
					return CE{C: cc, Type_id: store_last_id(cc.Store)}
				}
			}
			return CE{C: cc, Type_id: tid}
		}
		if (_ck2(cc) == "?") {
			cc = _cadv(cc)
			inner := resolve_type_name(cc.Store, name)
			_ = inner
			new_store := store_add(cc.Store, mk_optional_type(inner))
			_ = new_store
			cc = _cset_store(cc, new_store)
			return CE{C: cc, Type_id: store_last_id(cc.Store)}
		}
		if (_ck2(cc) == "!") {
			cc = _cadv(cc)
			ok_type := resolve_type_name(cc.Store, name)
			_ = ok_type
			err_result := _parse_type_ann(cc)
			_ = err_result
			cc = err_result.C
			new_store := store_add(cc.Store, mk_result_type(ok_type, err_result.Type_id))
			_ = new_store
			cc = _cset_store(cc, new_store)
			return CE{C: cc, Type_id: store_last_id(cc.Store)}
		}
		tid := resolve_type_name(cc.Store, name)
		_ = tid
		return CE{C: cc, Type_id: tid}
	}
	if (k == "[") {
		cc := _cadv(c)
		_ = cc
		elem := _parse_type_ann(cc)
		_ = elem
		cc = elem.C
		if (_ck2(cc) == "]") {
			cc = _cadv(cc)
		}
		new_store := store_add(cc.Store, mk_array_type(elem.Type_id))
		_ = new_store
		cc = _cset_store(cc, new_store)
		return CE{C: cc, Type_id: store_last_id(cc.Store)}
	}
	if (k == "fn") {
		cc := _cadv(c)
		_ = cc
		if ((_ck2(cc) == "IDENT") && (_ccur(cc).Text == "once")) {
			cc = _cadv(cc)
		}
		if (_ck2(cc) == "(") {
			cc = _cadv(cc)
			for ((_ck2(cc) != ")") && (_ck2(cc) != "EOF")) {
				pr := _parse_type_ann(cc)
				_ = pr
				cc = pr.C
				if (_ck2(cc) == ",") {
					cc = _cadv(cc)
				}
			}
			if (_ck2(cc) == ")") {
				cc = _cadv(cc)
			}
		}
		ret := TY_UNIT
		_ = ret
		if (_ck2(cc) == "->") {
			cc = _cadv(cc)
			rr := _parse_type_ann(cc)
			_ = rr
			cc = rr.C
			ret = rr.Type_id
		}
		new_st := store_add(cc.Store, mk_fn_type(ret, int64(0), int64(0)))
		_ = new_st
		cc = _cset_store(cc, new_st)
		return CE{C: cc, Type_id: store_last_id(cc.Store)}
	}
	if (k == "(") {
		cc := _cadv(c)
		_ = cc
		elem_ids := ""
		_ = elem_ids
		arity := int64(0)
		_ = arity
		for ((_ck2(cc) != ")") && (_ck2(cc) != "EOF")) {
			pr := _parse_type_ann(cc)
			_ = pr
			cc = pr.C
			if (arity > int64(0)) {
				elem_ids = (elem_ids + ",")
			}
			elem_ids = (elem_ids + i2s(pr.Type_id))
			arity = (arity + int64(1))
			if (_ck2(cc) == ",") {
				cc = _cadv(cc)
			}
		}
		if (_ck2(cc) == ")") {
			cc = _cadv(cc)
		}
		if (arity <= int64(1)) {
			if (arity == int64(1)) {
				first_pr := _parse_type_ann(_cset_pos(c, (c.Pos + int64(1))))
				_ = first_pr
				return CE{C: cc, Type_id: first_pr.Type_id}
			}
			return CE{C: cc, Type_id: TY_UNKNOWN}
		}
		new_store := store_add(cc.Store, mk_tuple_type(elem_ids, arity))
		_ = new_store
		cc = _cset_store(cc, new_store)
		return CE{C: cc, Type_id: store_last_id(cc.Store)}
	}
	return CE{C: c, Type_id: TY_UNKNOWN}
}

func _skip_type_toks(c Checker) Checker {
	cc := c
	_ = cc
	k := _ck2(cc)
	_ = k
	if ((k == "IDENT") || (k == "Self")) {
		cc = _cadv(cc)
		for (_ck2(cc) == ".") {
			cc = _cadv(_cadv(cc))
		}
		if (_ck2(cc) == "[") {
			d := int64(1)
			_ = d
			cc = _cadv(cc)
			for ((d > int64(0)) && (_ck2(cc) != "EOF")) {
				if (_ck2(cc) == "[") {
					d = (d + int64(1))
				}
				if (_ck2(cc) == "]") {
					d = (d - int64(1))
				}
				if (d > int64(0)) {
					cc = _cadv(cc)
				}
			}
			if (_ck2(cc) == "]") {
				cc = _cadv(cc)
			}
		}
		if (_ck2(cc) == "?") {
			cc = _cadv(cc)
		}
		if (_ck2(cc) == "!") {
			cc = _cadv(cc)
			cc = _skip_type_toks(cc)
		}
		return cc
	}
	if (k == "[") {
		cc = _cadv(cc)
		cc = _skip_type_toks(cc)
		if (_ck2(cc) == "]") {
			cc = _cadv(cc)
		}
		return cc
	}
	if (k == "fn") {
		cc = _cadv(cc)
		if ((_ck2(cc) == "IDENT") && (_ccur(cc).Text == "once")) {
			cc = _cadv(cc)
		}
		if (_ck2(cc) == "(") {
			cc = _cadv(cc)
			for ((_ck2(cc) != ")") && (_ck2(cc) != "EOF")) {
				cc = _skip_type_toks(cc)
				if (_ck2(cc) == ",") {
					cc = _cadv(cc)
				}
			}
			if (_ck2(cc) == ")") {
				cc = _cadv(cc)
			}
		}
		if (_ck2(cc) == "->") {
			cc = _cadv(cc)
			cc = _skip_type_toks(cc)
		}
		return cc
	}
	if (k == "(") {
		cc = _cadv(cc)
		for ((_ck2(cc) != ")") && (_ck2(cc) != "EOF")) {
			cc = _skip_type_toks(cc)
			if (_ck2(cc) == ",") {
				cc = _cadv(cc)
			}
		}
		if (_ck2(cc) == ")") {
			cc = _cadv(cc)
		}
		if (_ck2(cc) == "->") {
			cc = _cadv(cc)
			cc = _skip_type_toks(cc)
		}
		return cc
	}
	return cc
}

func _pass1_types(c Checker, index DeclIndex) Checker {
	cc := c
	_ = cc
	di := int64(1)
	_ = di
	for (di < decl_index_len(index)) {
		decl := decl_index_get(index, di)
		_ = decl
		kname := decl_kind_name(decl.Kind)
		_ = kname
		if (kname == "type") {
			cc = _cset_pos(cc, decl.Token_start)
			cc = _cskip_pub(cc)
			if strings.HasPrefix(decl.Name, "_alias:") {
				cc = _register_alias_decl(cc)
			} else if (decl.Node_idx > int64(0)) {
				cc = _register_type_decl_ast(cc, decl.Node_idx)
			} else {
				cc = _register_type_decl(cc)
			}
		} else if (kname == "enum") {
			if (decl.Node_idx > int64(0)) {
				cc = _register_enum_decl_ast(cc, decl.Node_idx)
			} else {
				cc = _cset_pos(cc, decl.Token_start)
				cc = _cskip_pub(cc)
				cc = _register_enum_decl(cc)
			}
		}
		di = (di + int64(1))
	}
	return cc
}

func _parse_derives(c Checker, type_name string) Checker {
	if (_ck2(c) != "derives") {
		return c
	}
	cc := _cadv(c)
	_ = cc
	if (_ck2(cc) != "[") {
		return cc
	}
	cc = _cadv(cc)
	cc = _cskip_nl(cc)
	for ((_ck2(cc) != "]") && (_ck2(cc) != "EOF")) {
		if (_ck2(cc) == "IDENT") {
			tname := _ccur(cc).Text
			_ = tname
			cc = _cadv(cc)
			cc = _cset_treg(cc, treg_add_impl(cc.Treg, _mk_impl(tname, type_name)))
		} else {
			cc = _cadv(cc)
		}
		cc = _cskip_nl(cc)
		if (_ck2(cc) == ",") {
			cc = _cadv(cc)
		}
		cc = _cskip_nl(cc)
	}
	if (_ck2(cc) == "]") {
		cc = _cadv(cc)
	}
	has_transient := treg_has_impl(cc.Treg, type_name, "Transient")
	_ = has_transient
	has_permanent := treg_has_impl(cc.Treg, type_name, "Permanent")
	_ = has_permanent
	has_user := treg_has_impl(cc.Treg, type_name, "UserFault")
	_ = has_user
	has_system := treg_has_impl(cc.Treg, type_name, "SystemFault")
	_ = has_system
	if (has_transient && has_permanent) {
		cc = _cerror(cc, E0301, (("type " + type_name) + " cannot derive both Transient and Permanent"))
	}
	if (has_user && has_system) {
		cc = _cerror(cc, E0301, (("type " + type_name) + " cannot derive both UserFault and SystemFault"))
	}
	return cc
}

func _mk_newtype_def(name string, underlying int64) StructDef {
	fields := []FieldInfo{FieldInfo{Name: "", Type_id: int64(0)}}
	_ = fields
	fields = append(fields, FieldInfo{Name: "value", Type_id: underlying})
	return new_struct_def(name, fields, false, []string{""})
}

func _register_alias_decl(c Checker) Checker {
	cc := _cadv(c)
	_ = cc
	if (_ck2(cc) != "IDENT") {
		return cc
	}
	name := _ccur(cc).Text
	_ = name
	cc = _cadv(cc)
	if (_ck2(cc) == "=") {
		cc = _cadv(cc)
	}
	cc = _cskip_nl(cc)
	ft := _parse_type_ann(cc)
	_ = ft
	cc = ft.C
	alias_info := TypeInfo{Kind: TypeKindTyNamed{}, Name: name, Type_id: ft.Type_id, Type_id2: int64(1), Param_start: int64(0), Param_count: int64(0)}
	_ = alias_info
	new_st := store_add(cc.Store, alias_info)
	_ = new_st
	cc = _cset_store(cc, new_st)
	return cc
}

func _mk_typevar_from_name(name string) TypeInfo {
	if ((int64(len(name)) > int64(3)) && (name[(int64(len(name)) - int64(3)):int64(len(name))] == "[_]")) {
		base := name[int64(0):(int64(len(name)) - int64(3))]
		_ = base
		return mk_hkt_typevar(base, int64(1))
	}
	return mk_typevar(name)
}

func _strip_hkt_suffix(name string) string {
	if ((int64(len(name)) > int64(3)) && (name[(int64(len(name)) - int64(3)):int64(len(name))] == "[_]")) {
		return name[int64(0):(int64(len(name)) - int64(3))]
	}
	return name
}

func _parse_generic_str(s string, gp_0 string, gp_1 string, gp_2 string, gb_0 string, gb_1 string, gb_2 string) []string {
	slen := int64(len(s))
	_ = slen
	if (slen == int64(0)) {
		return []string{"0", "", "", "", "", "", ""}
	}
	count := int64(0)
	_ = count
	r_gp_0 := ""
	_ = r_gp_0
	r_gp_1 := ""
	_ = r_gp_1
	r_gp_2 := ""
	_ = r_gp_2
	r_gb_0 := ""
	_ = r_gb_0
	r_gb_1 := ""
	_ = r_gb_1
	r_gb_2 := ""
	_ = r_gb_2
	seg_start := int64(0)
	_ = seg_start
	i := int64(0)
	_ = i
	for (i <= slen) {
		is_sep := false
		_ = is_sep
		if (i == slen) {
			is_sep = true
		}
		if (is_sep == false) {
			ch := string(s[i])
			_ = ch
			if (ch == ",") {
				is_sep = true
			}
		}
		if is_sep {
			seg := s[seg_start:i]
			_ = seg
			gname := seg
			_ = gname
			gbound := ""
			_ = gbound
			j := int64(0)
			_ = j
			seglen := int64(len(seg))
			_ = seglen
			found_colon := false
			_ = found_colon
			for (j < seglen) {
				if (found_colon == false) {
					sch := string(seg[j])
					_ = sch
					if (sch == ":") {
						gname = seg[int64(0):j]
						gbound = seg[(j + int64(1)):seglen]
						found_colon = true
					}
				}
				j = (j + int64(1))
			}
			if (count == int64(0)) {
				r_gp_0 = gname
			}
			if (count == int64(0)) {
				r_gb_0 = gbound
			}
			if (count == int64(1)) {
				r_gp_1 = gname
			}
			if (count == int64(1)) {
				r_gb_1 = gbound
			}
			if (count == int64(2)) {
				r_gp_2 = gname
			}
			if (count == int64(2)) {
				r_gb_2 = gbound
			}
			count = (count + int64(1))
			seg_start = (i + int64(1))
		}
		i = (i + int64(1))
	}
	return []string{i2s(count), r_gp_0, r_gp_1, r_gp_2, r_gb_0, r_gb_1, r_gb_2}
}

func _register_type_decl_ast(c Checker, node_idx int64) Checker {
	node := pool_get(c.Pool, node_idx)
	_ = node
	name := node.S1
	_ = name
	cc := c
	_ = cc
	gp_result := _parse_generic_str(node.S2, "", "", "", "", "", "")
	_ = gp_result
	sgp_count := int64(0)
	_ = sgp_count
	if (gp_result[int64(0)] == "1") {
		sgp_count = int64(1)
	}
	if (gp_result[int64(0)] == "2") {
		sgp_count = int64(2)
	}
	if (gp_result[int64(0)] == "3") {
		sgp_count = int64(3)
	}
	sgp_0 := gp_result[int64(1)]
	_ = sgp_0
	sgp_1 := gp_result[int64(2)]
	_ = sgp_1
	sgp_2 := gp_result[int64(3)]
	_ = sgp_2
	if (sgp_count > int64(0)) {
		new_st := store_add(cc.Store, mk_typevar(sgp_0))
		_ = new_st
		cc = _cset_store(cc, new_st)
	}
	if (sgp_count > int64(1)) {
		new_st := store_add(cc.Store, mk_typevar(sgp_1))
		_ = new_st
		cc = _cset_store(cc, new_st)
	}
	if (sgp_count > int64(2)) {
		new_st := store_add(cc.Store, mk_typevar(sgp_2))
		_ = new_st
		cc = _cset_store(cc, new_st)
	}
	lb := "{"
	_ = lb
	if node.B1 {
		variants := []string{""}
		_ = variants
		vf_counts := []int64{int64(0)}
		_ = vf_counts
		vf_names := []string{""}
		_ = vf_names
		vf_types := []int64{int64(0)}
		_ = vf_types
		vf_offsets := []int64{int64(0)}
		_ = vf_offsets
		flat_pos := int64(1)
		_ = flat_pos
		vi := int64(0)
		_ = vi
		for (vi < node.List_count) {
			idx_node := pool_get(cc.Pool, (node.List_start + vi))
			_ = idx_node
			vnode := pool_get(cc.Pool, idx_node.C0)
			_ = vnode
			variants = append(variants, vnode.S1)
			vf_offsets = append(vf_offsets, flat_pos)
			field_count := int64(0)
			_ = field_count
			fi := int64(0)
			_ = fi
			for (fi < vnode.List_count) {
				fidx_node := pool_get(cc.Pool, (vnode.List_start + fi))
				_ = fidx_node
				fnode := pool_get(cc.Pool, fidx_node.C0)
				_ = fnode
				cc = _cset_pos(cc, fnode.C0)
				ft := _parse_type_ann(cc)
				_ = ft
				cc = ft.C
				vf_names = append(vf_names, fnode.S1)
				vf_types = append(vf_types, ft.Type_id)
				field_count = (field_count + int64(1))
				flat_pos = (flat_pos + int64(1))
				fi = (fi + int64(1))
			}
			vf_counts = append(vf_counts, field_count)
			vi = (vi + int64(1))
		}
		new_st := store_add(cc.Store, mk_named_type(name))
		_ = new_st
		cc = _cset_store(cc, new_st)
		def := new_sum_def(name, variants, vf_counts, vf_names, vf_types, vf_offsets)
		_ = def
		cc = _cset_reg(cc, reg_add_struct(cc.Registry, def))
		cc = _cset_pos(cc, c.Pos)
		cc = _scan_to_derives(cc)
		cc = _parse_derives(cc, name)
		return cc
	}
	if ((node.C0 > int64(0)) && (node.List_count == int64(0))) {
		cc = _cset_pos(cc, node.C0)
		ft := _parse_type_ann(cc)
		_ = ft
		cc = ft.C
		new_st := store_add(cc.Store, mk_named_type(name))
		_ = new_st
		cc = _cset_store(cc, new_st)
		def := _mk_newtype_def(name, ft.Type_id)
		_ = def
		cc = _cset_reg(cc, reg_add_struct(cc.Registry, def))
		return cc
	}
	fields := []FieldInfo{FieldInfo{Name: "", Type_id: int64(0)}}
	_ = fields
	fi := int64(0)
	_ = fi
	for (fi < node.List_count) {
		idx_node := pool_get(cc.Pool, (node.List_start + fi))
		_ = idx_node
		fnode := pool_get(cc.Pool, idx_node.C0)
		_ = fnode
		cc = _cset_pos(cc, fnode.C0)
		ft := _parse_type_ann(cc)
		_ = ft
		cc = ft.C
		fields = append(fields, FieldInfo{Name: fnode.S1, Type_id: ft.Type_id})
		fi = (fi + int64(1))
	}
	new_st := store_add(cc.Store, mk_named_type(name))
	_ = new_st
	cc = _cset_store(cc, new_st)
	def := new_struct_def(name, fields, false, []string{""})
	_ = def
	if (sgp_count > int64(0)) {
		def = new_generic_struct_def(name, fields, sgp_count, sgp_0, sgp_1, sgp_2)
	}
	cc = _cset_reg(cc, reg_add_struct(cc.Registry, def))
	cc = _cset_pos(cc, c.Pos)
	cc = _scan_to_derives(cc)
	cc = _parse_derives(cc, name)
	return cc
}

func _scan_to_derives(c Checker) Checker {
	cc := c
	_ = cc
	lb := "{"
	_ = lb
	rb := "}"
	_ = rb
	if ((_ck2(cc) == "type") || (_ck2(cc) == "struct")) {
		cc = _cadv(cc)
	}
	if (_ck2(cc) == "IDENT") {
		cc = _cadv(cc)
	}
	if (_ck2(cc) == "[") {
		cc = _cadv(cc)
		depth := int64(1)
		_ = depth
		for ((depth > int64(0)) && (_ck2(cc) != "EOF")) {
			if (_ck2(cc) == "[") {
				depth = (depth + int64(1))
			}
			if (_ck2(cc) == "]") {
				depth = (depth - int64(1))
			}
			if (depth > int64(0)) {
				cc = _cadv(cc)
			}
		}
		if (_ck2(cc) == "]") {
			cc = _cadv(cc)
		}
	}
	if (_ck2(cc) == "=") {
		cc = _cadv(cc)
	}
	cc = _cskip_nl(cc)
	if (_ck2(cc) == "|") {
		for (_ck2(cc) == "|") {
			cc = _cadv(cc)
			cc = _cskip_nl(cc)
			if (_ck2(cc) == "IDENT") {
				cc = _cadv(cc)
			}
			if (_ck2(cc) == "(") {
				cc = _cadv(cc)
				depth := int64(1)
				_ = depth
				for ((depth > int64(0)) && (_ck2(cc) != "EOF")) {
					if (_ck2(cc) == "(") {
						depth = (depth + int64(1))
					}
					if (_ck2(cc) == ")") {
						depth = (depth - int64(1))
					}
					if (depth > int64(0)) {
						cc = _cadv(cc)
					}
				}
				if (_ck2(cc) == ")") {
					cc = _cadv(cc)
				}
			} else if (_ck2(cc) == lb) {
				cc = _skip_braces2(cc)
			}
			cc = _cskip_nl(cc)
		}
	} else if (_ck2(cc) == lb) {
		cc = _skip_braces2(cc)
	} else {
		for (((_ck2(cc) != "NEWLINE") && (_ck2(cc) != "EOF")) && (_ck2(cc) != "derives")) {
			cc = _cadv(cc)
		}
	}
	cc = _cskip_nl(cc)
	return cc
}

func _register_type_decl(c Checker) Checker {
	cc := _cadv(c)
	_ = cc
	if (_ck2(cc) != "IDENT") {
		return cc
	}
	name := _ccur(cc).Text
	_ = name
	cc = _cadv(cc)
	sgp_count := int64(0)
	_ = sgp_count
	sgp_0 := ""
	_ = sgp_0
	sgp_1 := ""
	_ = sgp_1
	sgp_2 := ""
	_ = sgp_2
	if (_ck2(cc) == "[") {
		cc = _cadv(cc)
		for ((_ck2(cc) != "]") && (_ck2(cc) != "EOF")) {
			if (_ck2(cc) == "IDENT") {
				gname := _ccur(cc).Text
				_ = gname
				if (sgp_count == int64(0)) {
					sgp_0 = gname
				}
				if (sgp_count == int64(1)) {
					sgp_1 = gname
				}
				if (sgp_count == int64(2)) {
					sgp_2 = gname
				}
				sgp_count = (sgp_count + int64(1))
				cc = _cadv(cc)
				if (_ck2(cc) == ":") {
					cc = _cadv(cc)
					if (_ck2(cc) == "IDENT") {
						cc = _cadv(cc)
					}
				}
			} else {
				cc = _cadv(cc)
			}
			if (_ck2(cc) == ",") {
				cc = _cadv(cc)
			}
		}
		if (_ck2(cc) == "]") {
			cc = _cadv(cc)
		}
	}
	if (sgp_count > int64(0)) {
		new_st := store_add(cc.Store, mk_typevar(sgp_0))
		_ = new_st
		cc = _cset_store(cc, new_st)
	}
	if (sgp_count > int64(1)) {
		new_st := store_add(cc.Store, mk_typevar(sgp_1))
		_ = new_st
		cc = _cset_store(cc, new_st)
	}
	if (sgp_count > int64(2)) {
		new_st := store_add(cc.Store, mk_typevar(sgp_2))
		_ = new_st
		cc = _cset_store(cc, new_st)
	}
	lb := "{"
	_ = lb
	if (_ck2(cc) == "=") {
		cc = _cadv(cc)
		cc = _cskip_nl(cc)
		if (_ck2(cc) == "|") {
			variants := []string{""}
			_ = variants
			vf_counts := []int64{int64(0)}
			_ = vf_counts
			vf_names := []string{""}
			_ = vf_names
			vf_types := []int64{int64(0)}
			_ = vf_types
			vf_offsets := []int64{int64(0)}
			_ = vf_offsets
			flat_pos := int64(1)
			_ = flat_pos
			for (_ck2(cc) == "|") {
				cc = _cadv(cc)
				cc = _cskip_nl(cc)
				if (_ck2(cc) == "IDENT") {
					variants = append(variants, _ccur(cc).Text)
					cc = _cadv(cc)
				}
				vf_offsets = append(vf_offsets, flat_pos)
				field_count := int64(0)
				_ = field_count
				if (_ck2(cc) == "(") {
					cc = _cadv(cc)
					cc = _cskip_nl(cc)
					fi := int64(0)
					_ = fi
					for ((_ck2(cc) != ")") && (_ck2(cc) != "EOF")) {
						ft := _parse_type_ann(cc)
						_ = ft
						cc = ft.C
						vf_names = append(vf_names, ("_" + i2s(fi)))
						vf_types = append(vf_types, ft.Type_id)
						field_count = (field_count + int64(1))
						flat_pos = (flat_pos + int64(1))
						fi = (fi + int64(1))
						cc = _cskip_nl(cc)
						if (_ck2(cc) == ",") {
							cc = _cadv(cc)
						}
						cc = _cskip_nl(cc)
					}
					if (_ck2(cc) == ")") {
						cc = _cadv(cc)
					}
				}
				if (_ck2(cc) == lb) {
					cc = _cadv(cc)
					cc = _cskip_nl(cc)
					rb := "}"
					_ = rb
					for ((_ck2(cc) != rb) && (_ck2(cc) != "EOF")) {
						if (_ck2(cc) == "IDENT") {
							fname := _ccur(cc).Text
							_ = fname
							cc = _cadv(cc)
							if (_ck2(cc) == ":") {
								cc = _cadv(cc)
								ft := _parse_type_ann(cc)
								_ = ft
								cc = ft.C
								vf_names = append(vf_names, fname)
								vf_types = append(vf_types, ft.Type_id)
								field_count = (field_count + int64(1))
								flat_pos = (flat_pos + int64(1))
							}
						} else {
							cc = _cadv(cc)
						}
						cc = _cskip_nl(cc)
						if (_ck2(cc) == ",") {
							cc = _cadv(cc)
						}
						cc = _cskip_nl(cc)
					}
					if (_ck2(cc) == rb) {
						cc = _cadv(cc)
					}
				}
				vf_counts = append(vf_counts, field_count)
				cc = _cskip_nl(cc)
			}
			new_st := store_add(cc.Store, mk_named_type(name))
			_ = new_st
			cc = _cset_store(cc, new_st)
			def := new_sum_def(name, variants, vf_counts, vf_names, vf_types, vf_offsets)
			_ = def
			cc = _cset_reg(cc, reg_add_struct(cc.Registry, def))
			cc = _cskip_nl(cc)
			cc = _parse_derives(cc, name)
			return cc
		}
		if (_ck2(cc) != lb) {
			ft := _parse_type_ann(cc)
			_ = ft
			cc = ft.C
			new_st := store_add(cc.Store, mk_named_type(name))
			_ = new_st
			cc = _cset_store(cc, new_st)
			def := _mk_newtype_def(name, ft.Type_id)
			_ = def
			cc = _cset_reg(cc, reg_add_struct(cc.Registry, def))
			return cc
		}
	}
	if (_ck2(cc) == lb) {
		cc = _cadv(cc)
		cc = _cskip_nl(cc)
		fields := []FieldInfo{FieldInfo{Name: "", Type_id: int64(0)}}
		_ = fields
		rb := "}"
		_ = rb
		for ((_ck2(cc) != rb) && (_ck2(cc) != "EOF")) {
			if (_ck2(cc) == "IDENT") {
				fname := _ccur(cc).Text
				_ = fname
				cc = _cadv(cc)
				if (_ck2(cc) == ":") {
					cc = _cadv(cc)
					ft := _parse_type_ann(cc)
					_ = ft
					cc = ft.C
					fields = append(fields, FieldInfo{Name: fname, Type_id: ft.Type_id})
				}
			} else {
				cc = _cadv(cc)
			}
			cc = _cskip_nl(cc)
			if (_ck2(cc) == ",") {
				cc = _cadv(cc)
			}
			cc = _cskip_nl(cc)
		}
		if (_ck2(cc) == rb) {
			cc = _cadv(cc)
		}
		new_st := store_add(cc.Store, mk_named_type(name))
		_ = new_st
		cc = _cset_store(cc, new_st)
		def := new_struct_def(name, fields, false, []string{""})
		_ = def
		if (sgp_count > int64(0)) {
			def = new_generic_struct_def(name, fields, sgp_count, sgp_0, sgp_1, sgp_2)
		}
		cc = _cset_reg(cc, reg_add_struct(cc.Registry, def))
		cc = _cskip_nl(cc)
		cc = _parse_derives(cc, name)
	}
	return cc
}

func _register_enum_decl_ast(c Checker, node_idx int64) Checker {
	node := pool_get(c.Pool, node_idx)
	_ = node
	name := node.S1
	_ = name
	cc := c
	_ = cc
	new_st := store_add(cc.Store, mk_named_type(name))
	_ = new_st
	cc = _cset_store(cc, new_st)
	def := new_struct_def(name, []FieldInfo{FieldInfo{Name: "", Type_id: int64(0)}}, false, []string{""})
	_ = def
	cc = _cset_reg(cc, reg_add_struct(cc.Registry, def))
	return cc
}

func _register_enum_decl(c Checker) Checker {
	cc := _cadv(c)
	_ = cc
	if (_ck2(cc) != "IDENT") {
		return cc
	}
	name := _ccur(cc).Text
	_ = name
	cc = _cadv(cc)
	cc = _skip_braces2(cc)
	new_st := store_add(cc.Store, mk_named_type(name))
	_ = new_st
	cc = _cset_store(cc, new_st)
	def := new_struct_def(name, []FieldInfo{FieldInfo{Name: "", Type_id: int64(0)}}, false, []string{""})
	_ = def
	cc = _cset_reg(cc, reg_add_struct(cc.Registry, def))
	return cc
}

func _skip_braces2(c Checker) Checker {
	lb := "{"
	_ = lb
	rb := "}"
	_ = rb
	if (_ck2(c) != lb) {
		return c
	}
	cc := _cadv(c)
	_ = cc
	d := int64(1)
	_ = d
	for ((d > int64(0)) && (_ck2(cc) != "EOF")) {
		if (_ck2(cc) == lb) {
			d = (d + int64(1))
		}
		if (_ck2(cc) == rb) {
			d = (d - int64(1))
		}
		if (d > int64(0)) {
			cc = _cadv(cc)
		}
	}
	if (_ck2(cc) == rb) {
		cc = _cadv(cc)
	}
	return cc
}

func _pass2_fns(c Checker, index DeclIndex) Checker {
	cc := c
	_ = cc
	di := int64(1)
	_ = di
	for (di < decl_index_len(index)) {
		decl := decl_index_get(index, di)
		_ = decl
		kname := decl_kind_name(decl.Kind)
		_ = kname
		cc = _cset_pos(cc, decl.Token_start)
		cc = _cskip_pub(cc)
		ann_deprecated := ""
		_ = ann_deprecated
		ann_cold := false
		_ = ann_cold
		if (_ck2(cc) == "@") {
			cc = _cadv(cc)
			if (_ck2(cc) == "IDENT") {
				ann_name := _ccur(cc).Text
				_ = ann_name
				cc = _cadv(cc)
				if (ann_name == "deprecated") {
					if (_ck2(cc) == "(") {
						cc = _cadv(cc)
						if (_ck2(cc) == "STRING") {
							ann_deprecated = _ccur(cc).Text
							cc = _cadv(cc)
						}
						if (_ck2(cc) == ")") {
							cc = _cadv(cc)
						}
					}
					if (ann_deprecated == "") {
						ann_deprecated = "deprecated"
					}
				}
				if (ann_name == "cold") {
					ann_cold = true
				}
				if (_ck2(cc) == "(") {
					cc = _cadv(cc)
					for ((_ck2(cc) != ")") && (_ck2(cc) != "EOF")) {
						cc = _cadv(cc)
					}
					if (_ck2(cc) == ")") {
						cc = _cadv(cc)
					}
				}
			}
			cc = _cskip_nl(cc)
		}
		if (kname == "fn") {
			if (decl.Name == "_extern") {
				cc = _register_extern(cc)
			} else if strings.HasPrefix(decl.Name, "_fixture:") {
				_skip := int64(0)
				_ = _skip
			} else {
				next_pos := (cc.Pos + int64(1))
				_ = next_pos
				if ((next_pos < int64(len(cc.Tokens))) && (token_name(cc.Tokens[next_pos].Kind) == "IDENT")) {
					fn_name := cc.Tokens[next_pos].Text
					_ = fn_name
					if (ann_deprecated != "") {
						cc = _cset_reg(cc, reg_add_deprecated(cc.Registry, fn_name, ann_deprecated))
					}
					if ann_cold {
						cc = _cset_reg(cc, reg_add_cold(cc.Registry, fn_name))
					}
				}
				cc = _register_fn_sig(cc)
			}
		} else if (kname == "impl") {
			cc = _register_impl_fns(cc, decl.Node_idx)
		} else if (kname == "trait") {
			if (decl.Node_idx > int64(0)) {
				cc = _register_trait_decl_ast(cc, decl.Node_idx)
			} else {
				cc = _register_trait_decl(cc)
			}
		} else if (kname == "const") {
			if (decl.Node_idx > int64(0)) {
				cc = _register_const_ast(cc, decl.Node_idx)
			} else {
				cc = _register_const(cc)
			}
		}
		di = (di + int64(1))
	}
	return cc
}

func _register_extern(c Checker) Checker {
	cc := _cadv(c)
	_ = cc
	lb := "{"
	_ = lb
	rb := "}"
	_ = rb
	if (_ck2(cc) == "STRING") {
		cc = _cadv(cc)
	}
	if (_ck2(cc) == lb) {
		cc = _cadv(cc)
		cc = _cskip_nl(cc)
		for ((_ck2(cc) != rb) && (_ck2(cc) != "EOF")) {
			cc = _cskip_nl(cc)
			if (_ck2(cc) == "fn") {
				cc = _register_fn_sig(cc)
			} else {
				cc = _cadv(cc)
			}
			cc = _cskip_nl(cc)
		}
		if (_ck2(cc) == rb) {
			cc = _cadv(cc)
		}
		return cc
	}
	if (_ck2(cc) == "fn") {
		cc = _register_fn_sig(cc)
	}
	return cc
}

func _register_trait_decl_ast(c Checker, node_idx int64) Checker {
	node := pool_get(c.Pool, node_idx)
	_ = node
	trait_name := node.S1
	_ = trait_name
	cc := c
	_ = cc
	s2val := node.S2
	_ = s2val
	s2len := int64(len(s2val))
	_ = s2len
	parents := []string{""}
	_ = parents
	assoc_types := []string{""}
	_ = assoc_types
	semi_pos := -int64(1)
	_ = semi_pos
	si := int64(0)
	_ = si
	for (si < s2len) {
		ch := string(s2val[si])
		_ = ch
		if (ch == ";") {
			semi_pos = si
		}
		si = (si + int64(1))
	}
	super_str := s2val
	_ = super_str
	if (semi_pos >= int64(0)) {
		super_str = s2val[int64(0):semi_pos]
		assoc_str := s2val[(semi_pos + int64(1)):s2len]
		_ = assoc_str
		alen := int64(len(assoc_str))
		_ = alen
		if (alen > int64(0)) {
			astart := int64(0)
			_ = astart
			ai := int64(0)
			_ = ai
			for (ai <= alen) {
				is_sep := false
				_ = is_sep
				if (ai == alen) {
					is_sep = true
				}
				if (is_sep == false) {
					ach := string(assoc_str[ai])
					_ = ach
					if (ach == ",") {
						is_sep = true
					}
				}
				if is_sep {
					if (ai > astart) {
						assoc_types = append(assoc_types, assoc_str[astart:ai])
					}
					astart = (ai + int64(1))
				}
				ai = (ai + int64(1))
			}
		}
	}
	slen := int64(len(super_str))
	_ = slen
	if (slen > int64(0)) {
		pstart := int64(0)
		_ = pstart
		pi := int64(0)
		_ = pi
		for (pi <= slen) {
			is_sep := false
			_ = is_sep
			if (pi == slen) {
				is_sep = true
			}
			if (is_sep == false) {
				pch := string(super_str[pi])
				_ = pch
				if (pch == ",") {
					is_sep = true
				}
			}
			if is_sep {
				if (pi > pstart) {
					parents = append(parents, super_str[pstart:pi])
				}
				pstart = (pi + int64(1))
			}
			pi = (pi + int64(1))
		}
	}
	mnames := []string{""}
	_ = mnames
	mi := int64(0)
	_ = mi
	for (mi < node.List_count) {
		midx_node := pool_get(cc.Pool, (node.List_start + mi))
		_ = midx_node
		method_node := pool_get(cc.Pool, midx_node.C0)
		_ = method_node
		mnames = append(mnames, method_node.S1)
		mi = (mi + int64(1))
	}
	tdef := TraitDef{Name: trait_name, Method_names: mnames, Parent_traits: parents, Assoc_type_names: assoc_types, Default_starts: []int64{int64(0)}, Default_ends: []int64{int64(0)}}
	_ = tdef
	cc = _cset_treg(cc, treg_add_trait(cc.Treg, tdef))
	return cc
}

func _register_trait_decl(c Checker) Checker {
	cc := _cadv(c)
	_ = cc
	if (_ck2(cc) != "IDENT") {
		return cc
	}
	trait_name := _ccur(cc).Text
	_ = trait_name
	cc = _cadv(cc)
	lb := "{"
	_ = lb
	rb := "}"
	_ = rb
	parents := []string{""}
	_ = parents
	if (_ck2(cc) == ":") {
		cc = _cadv(cc)
		parsing_parents := true
		_ = parsing_parents
		for parsing_parents {
			if (_ck2(cc) == "IDENT") {
				parents = append(parents, _ccur(cc).Text)
				cc = _cadv(cc)
				if (_ck2(cc) == ",") {
					cc = _cadv(cc)
				} else {
					parsing_parents = false
				}
			} else {
				parsing_parents = false
			}
		}
	}
	for ((_ck2(cc) != lb) && (_ck2(cc) != "EOF")) {
		cc = _cadv(cc)
	}
	if (_ck2(cc) != lb) {
		return cc
	}
	cc = _cadv(cc)
	cc = _cskip_nl(cc)
	method_name_0 := ""
	_ = method_name_0
	method_name_1 := ""
	_ = method_name_1
	method_name_2 := ""
	_ = method_name_2
	method_name_3 := ""
	_ = method_name_3
	method_count := int64(0)
	_ = method_count
	assoc_types := []string{""}
	_ = assoc_types
	for ((_ck2(cc) != rb) && (_ck2(cc) != "EOF")) {
		cc = _cskip_nl(cc)
		if (_ck2(cc) == "fn") {
			cc = _cadv(cc)
			if (_ck2(cc) == "IDENT") {
				mname := _ccur(cc).Text
				_ = mname
				if (method_count == int64(0)) {
					method_name_0 = mname
				}
				if (method_count == int64(1)) {
					method_name_1 = mname
				}
				if (method_count == int64(2)) {
					method_name_2 = mname
				}
				if (method_count == int64(3)) {
					method_name_3 = mname
				}
				method_count = (method_count + int64(1))
			}
			for (((_ck2(cc) != "NEWLINE") && (_ck2(cc) != rb)) && (_ck2(cc) != "EOF")) {
				if (_ck2(cc) == lb) {
					cc = _skip_braces2(cc)
				} else {
					cc = _cadv(cc)
				}
			}
		} else if (_ck2(cc) == "type") {
			cc = _cadv(cc)
			if (_ck2(cc) == "IDENT") {
				assoc_types = append(assoc_types, _ccur(cc).Text)
			}
			for (((_ck2(cc) != "NEWLINE") && (_ck2(cc) != rb)) && (_ck2(cc) != "EOF")) {
				cc = _cadv(cc)
			}
		} else if ((_ck2(cc) == rb) || (_ck2(cc) == "EOF")) {
		} else {
			cc = _cadv(cc)
		}
		cc = _cskip_nl(cc)
	}
	if (_ck2(cc) == rb) {
		cc = _cadv(cc)
	}
	mnames := []string{""}
	_ = mnames
	if (method_count > int64(0)) {
		mnames = append(mnames, method_name_0)
	}
	if (method_count > int64(1)) {
		mnames = append(mnames, method_name_1)
	}
	if (method_count > int64(2)) {
		mnames = append(mnames, method_name_2)
	}
	if (method_count > int64(3)) {
		mnames = append(mnames, method_name_3)
	}
	tdef := TraitDef{Name: trait_name, Method_names: mnames, Parent_traits: parents, Assoc_type_names: assoc_types, Default_starts: []int64{int64(0)}, Default_ends: []int64{int64(0)}}
	_ = tdef
	cc = _cset_treg(cc, treg_add_trait(cc.Treg, tdef))
	return cc
}

func _register_fn_sig_ast(c Checker, node_idx int64, token_start int64) Checker {
	node := pool_get(c.Pool, node_idx)
	_ = node
	name := node.S1
	_ = name
	cc := c
	_ = cc
	gp_result := _parse_generic_str(node.S2, "", "", "", "", "", "")
	_ = gp_result
	gp_count := int64(0)
	_ = gp_count
	if (gp_result[int64(0)] == "1") {
		gp_count = int64(1)
	}
	if (gp_result[int64(0)] == "2") {
		gp_count = int64(2)
	}
	if (gp_result[int64(0)] == "3") {
		gp_count = int64(3)
	}
	gp_0 := gp_result[int64(1)]
	_ = gp_0
	gp_1 := gp_result[int64(2)]
	_ = gp_1
	gp_2 := gp_result[int64(3)]
	_ = gp_2
	gb_0 := gp_result[int64(4)]
	_ = gb_0
	gb_1 := gp_result[int64(5)]
	_ = gb_1
	gb_2 := gp_result[int64(6)]
	_ = gb_2
	if (gp_count > int64(0)) {
		new_st := store_add(cc.Store, _mk_typevar_from_name(gp_0))
		_ = new_st
		cc = _cset_store(cc, new_st)
	}
	if (gp_count > int64(1)) {
		new_st := store_add(cc.Store, _mk_typevar_from_name(gp_1))
		_ = new_st
		cc = _cset_store(cc, new_st)
	}
	if (gp_count > int64(2)) {
		new_st := store_add(cc.Store, _mk_typevar_from_name(gp_2))
		_ = new_st
		cc = _cset_store(cc, new_st)
	}
	pnames := []string{""}
	_ = pnames
	ptypes := []int64{int64(0)}
	_ = ptypes
	pi := int64(0)
	_ = pi
	for (pi < node.List_count) {
		pidx_node := pool_get(cc.Pool, (node.List_start + pi))
		_ = pidx_node
		pnode := pool_get(cc.Pool, pidx_node.C0)
		_ = pnode
		if pnode.B1 {
			pnames = append(pnames, "self")
			ptypes = append(ptypes, TY_UNKNOWN)
		} else {
			pnames = append(pnames, pnode.S1)
			cc = _cset_pos(cc, pnode.C0)
			pt := _parse_type_ann(cc)
			_ = pt
			cc = pt.C
			ptypes = append(ptypes, pt.Type_id)
		}
		pi = (pi + int64(1))
	}
	ret := TY_UNIT
	_ = ret
	err := int64(0)
	_ = err
	if node.B2 {
		cc = _cset_pos(cc, node.C1)
		rt := _parse_type_ann(cc)
		_ = rt
		cc = rt.C
		rtinfo := store_get(cc.Store, rt.Type_id)
		_ = rtinfo
		if (type_kind_name(rtinfo.Kind) == "Result") {
			ret = rtinfo.Type_id
			err = rtinfo.Type_id2
		} else {
			ret = rt.Type_id
		}
	}
	if ((err == int64(0)) && (node.C2 > int64(0))) {
		cc = _cset_pos(cc, node.C2)
		et := _parse_type_ann(cc)
		_ = et
		cc = et.C
		err = et.Type_id
	}
	cc = _cset_pos(cc, token_start)
	eff_io := false
	_ = eff_io
	eff_fs := false
	_ = eff_fs
	eff_net := false
	_ = eff_net
	eff_ffi := false
	_ = eff_ffi
	eff_async := false
	_ = eff_async
	user_effs := []string{""}
	_ = user_effs
	lb := "{"
	_ = lb
	scan_limit := int64(0)
	_ = scan_limit
	for ((((_ck2(cc) != "=") && (_ck2(cc) != lb)) && (_ck2(cc) != "EOF")) && (scan_limit < int64(200))) {
		if (_ck2(cc) == "with") {
			cc = _cadv(cc)
			if (_ck2(cc) == "[") {
				cc = _cadv(cc)
				for ((_ck2(cc) != "]") && (_ck2(cc) != "EOF")) {
					if (_ck2(cc) == "IDENT") {
						eff := _ccur(cc).Text
						_ = eff
						is_builtin := false
						_ = is_builtin
						if (eff == "Io") {
							eff_io = true
						}
						if (eff == "Fs") {
							eff_fs = true
						}
						if (eff == "Net") {
							eff_net = true
						}
						if (eff == "Ffi") {
							eff_ffi = true
						}
						if (eff == "Async") {
							eff_async = true
						}
						if (((((eff == "Io") || (eff == "Fs")) || (eff == "Net")) || (eff == "Ffi")) || (eff == "Async")) {
							is_builtin = true
						}
						if (is_builtin == false) {
							user_effs = append(user_effs, eff)
						}
						cc = _cadv(cc)
					} else {
						cc = _cadv(cc)
					}
					if (_ck2(cc) == ",") {
						cc = _cadv(cc)
					}
				}
				if (_ck2(cc) == "]") {
					cc = _cadv(cc)
				}
			}
		}
		if (_ck2(cc) == "where") {
			cc = _cadv(cc)
			parsing_where := true
			_ = parsing_where
			for parsing_where {
				if (_ck2(cc) == "IDENT") {
					wname := _ccur(cc).Text
					_ = wname
					cc = _cadv(cc)
					if (_ck2(cc) == ":") {
						cc = _cadv(cc)
						if (_ck2(cc) == "IDENT") {
							wbound := _ccur(cc).Text
							_ = wbound
							cc = _cadv(cc)
							for (_ck2(cc) == "+") {
								cc = _cadv(cc)
								if (_ck2(cc) == "IDENT") {
									wbound = ((wbound + "+") + _ccur(cc).Text)
									cc = _cadv(cc)
								}
							}
							_skip := int64(0)
							_ = _skip
						}
					}
					if (_ck2(cc) == ",") {
						cc = _cadv(cc)
					} else {
						parsing_where = false
					}
				} else {
					parsing_where = false
				}
			}
		}
		cc = _cadv(cc)
		scan_limit = (scan_limit + int64(1))
	}
	sig := new_fn_sig(name, pnames, ptypes, ret, err)
	_ = sig
	if (gp_count > int64(0)) {
		sig = new_generic_fn_sig(name, pnames, ptypes, ret, err, gp_count, _strip_hkt_suffix(gp_0), _strip_hkt_suffix(gp_1), _strip_hkt_suffix(gp_2), gb_0, gb_1, gb_2, token_start)
	}
	if ((((eff_io || eff_fs) || eff_net) || eff_ffi) || eff_async) {
		sig = sig_set_effects(sig, eff_io, eff_fs, eff_net, eff_ffi, eff_async)
	}
	if (int64(len(user_effs)) > int64(1)) {
		sig = sig_set_user_effects(sig, user_effs)
	}
	cc = _cset_reg(cc, reg_add_fn(cc.Registry, sig))
	return cc
}

func _register_impl_fn_sig_ast(c Checker, impl_type string, method_idx int64) Checker {
	node := pool_get(c.Pool, method_idx)
	_ = node
	raw_name := node.S1
	_ = raw_name
	name := ((impl_type + "_") + raw_name)
	_ = name
	cc := c
	_ = cc
	gp_result := _parse_generic_str(node.S2, "", "", "", "", "", "")
	_ = gp_result
	gp_count := int64(0)
	_ = gp_count
	if (gp_result[int64(0)] == "1") {
		gp_count = int64(1)
	}
	if (gp_result[int64(0)] == "2") {
		gp_count = int64(2)
	}
	if (gp_result[int64(0)] == "3") {
		gp_count = int64(3)
	}
	gp_0 := gp_result[int64(1)]
	_ = gp_0
	gp_1 := gp_result[int64(2)]
	_ = gp_1
	gp_2 := gp_result[int64(3)]
	_ = gp_2
	gb_0 := gp_result[int64(4)]
	_ = gb_0
	gb_1 := gp_result[int64(5)]
	_ = gb_1
	gb_2 := gp_result[int64(6)]
	_ = gb_2
	if (gp_count > int64(0)) {
		new_st := store_add(cc.Store, _mk_typevar_from_name(gp_0))
		_ = new_st
		cc = _cset_store(cc, new_st)
	}
	if (gp_count > int64(1)) {
		new_st := store_add(cc.Store, _mk_typevar_from_name(gp_1))
		_ = new_st
		cc = _cset_store(cc, new_st)
	}
	if (gp_count > int64(2)) {
		new_st := store_add(cc.Store, _mk_typevar_from_name(gp_2))
		_ = new_st
		cc = _cset_store(cc, new_st)
	}
	pnames := []string{""}
	_ = pnames
	ptypes := []int64{int64(0)}
	_ = ptypes
	pi := int64(0)
	_ = pi
	for (pi < node.List_count) {
		pidx_node := pool_get(cc.Pool, (node.List_start + pi))
		_ = pidx_node
		pnode := pool_get(cc.Pool, pidx_node.C0)
		_ = pnode
		if pnode.B1 {
			pnames = append(pnames, "self")
			self_type := resolve_type_name(cc.Store, impl_type)
			_ = self_type
			ptypes = append(ptypes, self_type)
		} else {
			pnames = append(pnames, pnode.S1)
			cc = _cset_pos(cc, pnode.C0)
			pt := _parse_type_ann(cc)
			_ = pt
			cc = pt.C
			ptypes = append(ptypes, pt.Type_id)
		}
		pi = (pi + int64(1))
	}
	ret := TY_UNIT
	_ = ret
	err := int64(0)
	_ = err
	if node.B2 {
		cc = _cset_pos(cc, node.C1)
		rt := _parse_type_ann(cc)
		_ = rt
		cc = rt.C
		rtinfo := store_get(cc.Store, rt.Type_id)
		_ = rtinfo
		if (type_kind_name(rtinfo.Kind) == "Result") {
			ret = rtinfo.Type_id
			err = rtinfo.Type_id2
		} else {
			ret = rt.Type_id
		}
	}
	if ((err == int64(0)) && (node.C2 > int64(0))) {
		cc = _cset_pos(cc, node.C2)
		et := _parse_type_ann(cc)
		_ = et
		cc = et.C
		err = et.Type_id
	}
	sig := new_fn_sig(name, pnames, ptypes, ret, err)
	_ = sig
	if (gp_count > int64(0)) {
		tstart := int64(0)
		_ = tstart
		ti := int64(0)
		_ = ti
		for (ti < int64(len(cc.Tokens))) {
			if (cc.Tokens[ti].Offset == node.Span.Offset) {
				tstart = ti
				ti = int64(len(cc.Tokens))
			}
			ti = (ti + int64(1))
		}
		sig = new_generic_fn_sig(name, pnames, ptypes, ret, err, gp_count, _strip_hkt_suffix(gp_0), _strip_hkt_suffix(gp_1), _strip_hkt_suffix(gp_2), gb_0, gb_1, gb_2, tstart)
	}
	cc = _cset_reg(cc, reg_add_fn(cc.Registry, sig))
	return cc
}

func _is_fn_name_kind(tk string) bool {
	if (tk == "IDENT") {
		return true
	}
	if (tk == "test") {
		return true
	}
	if (tk == "check") {
		return true
	}
	if (tk == "match") {
		return true
	}
	if (tk == "type") {
		return true
	}
	if (tk == "where") {
		return true
	}
	if (tk == "bench") {
		return true
	}
	return false
}

func _register_fn_sig(c Checker) Checker {
	token_start := c.Pos
	_ = token_start
	cc := _cadv(c)
	_ = cc
	if (_is_fn_name_kind(_ck2(cc)) == false) {
		return cc
	}
	name := _ccur(cc).Text
	_ = name
	cc = _cadv(cc)
	gp_count := int64(0)
	_ = gp_count
	gp_0 := ""
	_ = gp_0
	gp_1 := ""
	_ = gp_1
	gp_2 := ""
	_ = gp_2
	gb_0 := ""
	_ = gb_0
	gb_1 := ""
	_ = gb_1
	gb_2 := ""
	_ = gb_2
	if (_ck2(cc) == "[") {
		cc = _cadv(cc)
		for ((_ck2(cc) != "]") && (_ck2(cc) != "EOF")) {
			if (_ck2(cc) == "IDENT") {
				gname := _ccur(cc).Text
				_ = gname
				if (gp_count == int64(0)) {
					gp_0 = gname
				}
				if (gp_count == int64(1)) {
					gp_1 = gname
				}
				if (gp_count == int64(2)) {
					gp_2 = gname
				}
				cc = _cadv(cc)
				if (_ck2(cc) == ":") {
					cc = _cadv(cc)
					if (_ck2(cc) == "IDENT") {
						bname := _ccur(cc).Text
						_ = bname
						cc = _cadv(cc)
						for (_ck2(cc) == "+") {
							cc = _cadv(cc)
							if (_ck2(cc) == "IDENT") {
								bname = ((bname + "+") + _ccur(cc).Text)
								cc = _cadv(cc)
							}
						}
						if (gp_count == int64(0)) {
							gb_0 = bname
						}
						if (gp_count == int64(1)) {
							gb_1 = bname
						}
						if (gp_count == int64(2)) {
							gb_2 = bname
						}
					}
				}
				gp_count = (gp_count + int64(1))
			} else {
				cc = _cadv(cc)
			}
			if (_ck2(cc) == ",") {
				cc = _cadv(cc)
			}
		}
		if (_ck2(cc) == "]") {
			cc = _cadv(cc)
		}
	}
	if (gp_count > int64(0)) {
		new_st := store_add(cc.Store, _mk_typevar_from_name(gp_0))
		_ = new_st
		cc = _cset_store(cc, new_st)
	}
	if (gp_count > int64(1)) {
		new_st := store_add(cc.Store, _mk_typevar_from_name(gp_1))
		_ = new_st
		cc = _cset_store(cc, new_st)
	}
	if (gp_count > int64(2)) {
		new_st := store_add(cc.Store, _mk_typevar_from_name(gp_2))
		_ = new_st
		cc = _cset_store(cc, new_st)
	}
	pnames := []string{""}
	_ = pnames
	ptypes := []int64{int64(0)}
	_ = ptypes
	if (_ck2(cc) == "(") {
		cc = _cadv(cc)
		for ((_ck2(cc) != ")") && (_ck2(cc) != "EOF")) {
			is_param := ((_ck2(cc) == "IDENT") || (_ck2(cc) == "self"))
			_ = is_param
			if ((is_param == false) && ((cc.Pos + int64(1)) < int64(len(cc.Tokens)))) {
				if (token_name(cc.Tokens[(cc.Pos + int64(1))].Kind) == ":") {
					is_param = true
				}
			}
			if is_param {
				pn := _ccur(cc).Text
				_ = pn
				cc = _cadv(cc)
				if (_ck2(cc) == ":") {
					cc = _cadv(cc)
					pt := _parse_type_ann(cc)
					_ = pt
					cc = pt.C
					pnames = append(pnames, pn)
					ptypes = append(ptypes, pt.Type_id)
				}
			} else {
				cc = _cadv(cc)
			}
			if (_ck2(cc) == ",") {
				cc = _cadv(cc)
			}
		}
		if (_ck2(cc) == ")") {
			cc = _cadv(cc)
		}
	}
	ret := TY_UNIT
	_ = ret
	err := int64(0)
	_ = err
	if (_ck2(cc) == "->") {
		cc = _cadv(cc)
		rt := _parse_type_ann(cc)
		_ = rt
		cc = rt.C
		rtinfo := store_get(cc.Store, rt.Type_id)
		_ = rtinfo
		if (type_kind_name(rtinfo.Kind) == "Result") {
			ret = rtinfo.Type_id
			err = rtinfo.Type_id2
		} else {
			ret = rt.Type_id
		}
	}
	if ((err == int64(0)) && (_ck2(cc) == "!")) {
		cc = _cadv(cc)
		et := _parse_type_ann(cc)
		_ = et
		cc = et.C
		err = et.Type_id
	}
	eff_io := false
	_ = eff_io
	eff_fs := false
	_ = eff_fs
	eff_net := false
	_ = eff_net
	eff_ffi := false
	_ = eff_ffi
	eff_async := false
	_ = eff_async
	if (_ck2(cc) == "with") {
		cc = _cadv(cc)
		if (_ck2(cc) == "[") {
			cc = _cadv(cc)
			for ((_ck2(cc) != "]") && (_ck2(cc) != "EOF")) {
				if (_ck2(cc) == "IDENT") {
					eff := _ccur(cc).Text
					_ = eff
					if (eff == "Io") {
						eff_io = true
					}
					if (eff == "Fs") {
						eff_fs = true
					}
					if (eff == "Net") {
						eff_net = true
					}
					if (eff == "Ffi") {
						eff_ffi = true
					}
					if (eff == "Async") {
						eff_async = true
					}
					cc = _cadv(cc)
				} else {
					cc = _cadv(cc)
				}
				if (_ck2(cc) == ",") {
					cc = _cadv(cc)
				}
			}
			if (_ck2(cc) == "]") {
				cc = _cadv(cc)
			}
		}
	}
	if (_ck2(cc) == "where") {
		cc = _cadv(cc)
		parsing_where := true
		_ = parsing_where
		for parsing_where {
			if (_ck2(cc) == "IDENT") {
				wname := _ccur(cc).Text
				_ = wname
				cc = _cadv(cc)
				if (_ck2(cc) == ":") {
					cc = _cadv(cc)
					if (_ck2(cc) == "IDENT") {
						wbound := _ccur(cc).Text
						_ = wbound
						cc = _cadv(cc)
						for (_ck2(cc) == "+") {
							cc = _cadv(cc)
							if (_ck2(cc) == "IDENT") {
								wbound = ((wbound + "+") + _ccur(cc).Text)
								cc = _cadv(cc)
							}
						}
						if ((wname == gp_0) && (gb_0 == "")) {
							gb_0 = wbound
						}
						if ((wname == gp_1) && (gb_1 == "")) {
							gb_1 = wbound
						}
						if ((wname == gp_2) && (gb_2 == "")) {
							gb_2 = wbound
						}
					}
				}
				if (_ck2(cc) == ",") {
					cc = _cadv(cc)
				} else {
					parsing_where = false
				}
			} else {
				parsing_where = false
			}
		}
	}
	sig := new_fn_sig(name, pnames, ptypes, ret, err)
	_ = sig
	if (gp_count > int64(0)) {
		sig = new_generic_fn_sig(name, pnames, ptypes, ret, err, gp_count, _strip_hkt_suffix(gp_0), _strip_hkt_suffix(gp_1), _strip_hkt_suffix(gp_2), gb_0, gb_1, gb_2, token_start)
	}
	if ((((eff_io || eff_fs) || eff_net) || eff_ffi) || eff_async) {
		sig = sig_set_effects(sig, eff_io, eff_fs, eff_net, eff_ffi, eff_async)
	}
	cc = _cset_reg(cc, reg_add_fn(cc.Registry, sig))
	lb := "{"
	_ = lb
	for (_ck2(cc) != "EOF") {
		if (_ck2(cc) == lb) {
			cc = _skip_braces2(cc)
			return cc
		}
		if (_ck2(cc) == "NEWLINE") {
			return cc
		}
		cc = _cadv(cc)
	}
	return cc
}

func _register_impl_fns(c Checker, node_idx int64) Checker {
	cc := _cadv(c)
	_ = cc
	is_blanket := false
	_ = is_blanket
	blanket_bound := ""
	_ = blanket_bound
	if (_ck2(cc) == "[") {
		is_blanket = true
		cc = _cadv(cc)
		if (_ck2(cc) == "IDENT") {
			cc = _cadv(cc)
		}
		if (_ck2(cc) == ":") {
			cc = _cadv(cc)
			if (_ck2(cc) == "IDENT") {
				blanket_bound = _ccur(cc).Text
				cc = _cadv(cc)
			}
		}
		depth := int64(1)
		_ = depth
		for ((depth > int64(0)) && (_ck2(cc) != "EOF")) {
			if (_ck2(cc) == "[") {
				depth = (depth + int64(1))
			}
			if (_ck2(cc) == "]") {
				depth = (depth - int64(1))
			}
			if (depth > int64(0)) {
				cc = _cadv(cc)
			}
		}
		if (_ck2(cc) == "]") {
			cc = _cadv(cc)
		}
	}
	first_name := ""
	_ = first_name
	impl_type := ""
	_ = impl_type
	impl_trait := ""
	_ = impl_trait
	if (_ck2(cc) == "IDENT") {
		first_name = _ccur(cc).Text
		impl_type = first_name
		cc = _cadv(cc)
	}
	if (_ck2(cc) == "for") {
		impl_trait = first_name
		cc = _cadv(cc)
		if (_ck2(cc) == "IDENT") {
			impl_type = _ccur(cc).Text
			cc = _cadv(cc)
		}
	}
	if (is_blanket && (impl_trait != "")) {
		cc = _cset_treg(cc, treg_add_impl(cc.Treg, _mk_impl(impl_trait, ("*" + blanket_bound))))
	}
	if ((impl_trait != "") && (impl_type != "")) {
		tdef_check := treg_find_trait(cc.Treg, impl_trait)
		_ = tdef_check
		if (tdef_check.Name == "") {
			q := _q()
			_ = q
			cc = _cerror(cc, E0200, (((("unknown trait " + q) + impl_trait) + q) + " in impl declaration"))
		}
		current_pkg := _pkg_from_path(cc.File)
		_ = current_pkg
		if (((current_pkg != "") && (impl_trait != "")) && (impl_type != "")) {
			trait_pkg := ""
			_ = trait_pkg
			ti := int64(1)
			_ = ti
			for (ti < int64(len(cc.Treg.Impl_defs))) {
				imp := cc.Treg.Impl_defs[ti]
				_ = imp
				if ((imp.Trait_name == impl_trait) && (ti < int64(len(cc.Treg.Impl_pkgs)))) {
					p := cc.Treg.Impl_pkgs[ti]
					_ = p
					if (p != "") {
						trait_pkg = p
						ti = int64(len(cc.Treg.Impl_defs))
					}
				}
				ti = (ti + int64(1))
			}
			type_pkg := ""
			_ = type_pkg
			yi := int64(1)
			_ = yi
			for (yi < int64(len(cc.Treg.Impl_defs))) {
				imp := cc.Treg.Impl_defs[yi]
				_ = imp
				if ((imp.Type_name == impl_type) && (yi < int64(len(cc.Treg.Impl_pkgs)))) {
					p := cc.Treg.Impl_pkgs[yi]
					_ = p
					if (p != "") {
						type_pkg = p
						yi = int64(len(cc.Treg.Impl_defs))
					}
				}
				yi = (yi + int64(1))
			}
			if ((((trait_pkg != "") && (trait_pkg != current_pkg)) && (type_pkg != "")) && (type_pkg != current_pkg)) {
				q := _q()
				_ = q
				cc = _cerror(cc, E0202, (((((((("orphan impl: trait " + q) + impl_trait) + q) + " and type ") + q) + impl_type) + q) + " both defined in other packages"))
			}
		}
		cc = _cset_treg(cc, treg_add_impl_pkg(cc.Treg, _mk_impl(impl_trait, impl_type), current_pkg))
		tdef := treg_find_trait(cc.Treg, impl_trait)
		_ = tdef
		if (tdef.Name != "") {
			pi := int64(1)
			_ = pi
			for (pi < int64(len(tdef.Parent_traits))) {
				parent := tdef.Parent_traits[pi]
				_ = parent
				if (treg_has_impl(cc.Treg, impl_type, parent) == false) {
					q := _q()
					_ = q
					cc = _cerror(cc, E0200, (((((((((((("trait " + q) + impl_trait) + q) + " requires ") + q) + parent) + q) + " but ") + q) + impl_type) + q) + " does not implement it"))
				}
				pi = (pi + int64(1))
			}
		}
	}
	reg_prefix := impl_type
	_ = reg_prefix
	if is_blanket {
		reg_prefix = ""
	}
	if (node_idx > int64(0)) {
		saved_impl_pos := cc.Pos
		_ = saved_impl_pos
		impl_node := pool_get(cc.Pool, node_idx)
		_ = impl_node
		mi := int64(0)
		_ = mi
		for (mi < impl_node.List_count) {
			midx_node := pool_get(cc.Pool, (impl_node.List_start + mi))
			_ = midx_node
			method_idx := midx_node.C0
			_ = method_idx
			method_node := pool_get(cc.Pool, method_idx)
			_ = method_node
			cc = _register_impl_fn_sig_ast(cc, reg_prefix, method_idx)
			mi = (mi + int64(1))
		}
		cc = _cset_pos(cc, saved_impl_pos)
		lb := "{"
		_ = lb
		for ((_ck2(cc) != lb) && (_ck2(cc) != "EOF")) {
			cc = _cadv(cc)
		}
		if (_ck2(cc) != lb) {
			return cc
		}
		cc = _cadv(cc)
		cc = _cskip_nl(cc)
		rb := "}"
		_ = rb
		for ((_ck2(cc) != rb) && (_ck2(cc) != "EOF")) {
			cc = _cskip_nl(cc)
			cc = _cskip_pub(cc)
			if (_ck2(cc) == "fn") {
				cc = _cadv(cc)
				for ((_ck2(cc) != "NEWLINE") && (_ck2(cc) != "EOF")) {
					if (_ck2(cc) == lb) {
						cc = _skip_braces2(cc)
					} else {
						cc = _cadv(cc)
					}
				}
			} else if (_ck2(cc) == "type") {
				cc = _cadv(cc)
				if (_ck2(cc) == "IDENT") {
					assoc_name := _ccur(cc).Text
					_ = assoc_name
					cc = _cadv(cc)
					if (_ck2(cc) == "=") {
						cc = _cadv(cc)
						if (_ck2(cc) == "IDENT") {
							assoc_value := _ccur(cc).Text
							_ = assoc_value
							if (impl_type != "") {
								alias_name := ((impl_type + "_") + assoc_name)
								_ = alias_name
								tid := resolve_type_name(cc.Store, assoc_value)
								_ = tid
								if (tid != TY_UNKNOWN) {
									new_st := store_add(cc.Store, TypeInfo{Kind: TypeKindTyNamed{}, Name: alias_name, Type_id: tid, Type_id2: int64(0), Param_start: int64(0), Param_count: int64(0)})
									_ = new_st
									cc = _cset_store(cc, new_st)
								}
							}
						}
					}
				}
				for (((_ck2(cc) != "NEWLINE") && (_ck2(cc) != rb)) && (_ck2(cc) != "EOF")) {
					cc = _cadv(cc)
				}
			} else if ((_ck2(cc) == rb) || (_ck2(cc) == "EOF")) {
				return cc
			} else {
				cc = _cadv(cc)
			}
			cc = _cskip_nl(cc)
		}
		if (_ck2(cc) == rb) {
			cc = _cadv(cc)
		}
		return cc
	}
	lb := "{"
	_ = lb
	for ((_ck2(cc) != lb) && (_ck2(cc) != "EOF")) {
		cc = _cadv(cc)
	}
	if (_ck2(cc) != lb) {
		return cc
	}
	cc = _cadv(cc)
	cc = _cskip_nl(cc)
	rb := "}"
	_ = rb
	for ((_ck2(cc) != rb) && (_ck2(cc) != "EOF")) {
		cc = _cskip_nl(cc)
		cc = _cskip_pub(cc)
		if (_ck2(cc) == "fn") {
			cc = _register_impl_fn_sig(cc, impl_type)
		} else if (_ck2(cc) == "type") {
			cc = _cadv(cc)
			if (_ck2(cc) == "IDENT") {
				assoc_name := _ccur(cc).Text
				_ = assoc_name
				cc = _cadv(cc)
				if (_ck2(cc) == "=") {
					cc = _cadv(cc)
					if (_ck2(cc) == "IDENT") {
						assoc_value := _ccur(cc).Text
						_ = assoc_value
						if (impl_type != "") {
							alias_name := ((impl_type + "_") + assoc_name)
							_ = alias_name
							tid := resolve_type_name(cc.Store, assoc_value)
							_ = tid
							if (tid != TY_UNKNOWN) {
								new_st := store_add(cc.Store, TypeInfo{Kind: TypeKindTyNamed{}, Name: alias_name, Type_id: tid, Type_id2: int64(0), Param_start: int64(0), Param_count: int64(0)})
								_ = new_st
								cc = _cset_store(cc, new_st)
							}
						}
					}
				}
			}
			for (((_ck2(cc) != "NEWLINE") && (_ck2(cc) != rb)) && (_ck2(cc) != "EOF")) {
				cc = _cadv(cc)
			}
		} else if ((_ck2(cc) == rb) || (_ck2(cc) == "EOF")) {
			return cc
		} else {
			cc = _cadv(cc)
		}
		cc = _cskip_nl(cc)
	}
	if (_ck2(cc) == rb) {
		cc = _cadv(cc)
	}
	return cc
}

func _register_impl_fn_sig(c Checker, impl_type string) Checker {
	token_start := c.Pos
	_ = token_start
	cc := _cadv(c)
	_ = cc
	if (_ck2(cc) != "IDENT") {
		return cc
	}
	raw_name := _ccur(cc).Text
	_ = raw_name
	name := ((impl_type + "_") + raw_name)
	_ = name
	cc = _cadv(cc)
	gp_count := int64(0)
	_ = gp_count
	gp_0 := ""
	_ = gp_0
	gp_1 := ""
	_ = gp_1
	gp_2 := ""
	_ = gp_2
	gb_0 := ""
	_ = gb_0
	gb_1 := ""
	_ = gb_1
	gb_2 := ""
	_ = gb_2
	if (_ck2(cc) == "[") {
		cc = _cadv(cc)
		for ((_ck2(cc) != "]") && (_ck2(cc) != "EOF")) {
			if (_ck2(cc) == "IDENT") {
				gname := _ccur(cc).Text
				_ = gname
				if (gp_count == int64(0)) {
					gp_0 = gname
				}
				if (gp_count == int64(1)) {
					gp_1 = gname
				}
				if (gp_count == int64(2)) {
					gp_2 = gname
				}
				cc = _cadv(cc)
				if (_ck2(cc) == ":") {
					cc = _cadv(cc)
					if (_ck2(cc) == "IDENT") {
						bname := _ccur(cc).Text
						_ = bname
						if (gp_count == int64(0)) {
							gb_0 = bname
						}
						if (gp_count == int64(1)) {
							gb_1 = bname
						}
						if (gp_count == int64(2)) {
							gb_2 = bname
						}
						cc = _cadv(cc)
					}
				}
				gp_count = (gp_count + int64(1))
			} else {
				cc = _cadv(cc)
			}
			if (_ck2(cc) == ",") {
				cc = _cadv(cc)
			}
		}
		if (_ck2(cc) == "]") {
			cc = _cadv(cc)
		}
	}
	if (gp_count > int64(0)) {
		new_st := store_add(cc.Store, _mk_typevar_from_name(gp_0))
		_ = new_st
		cc = _cset_store(cc, new_st)
	}
	if (gp_count > int64(1)) {
		new_st := store_add(cc.Store, _mk_typevar_from_name(gp_1))
		_ = new_st
		cc = _cset_store(cc, new_st)
	}
	if (gp_count > int64(2)) {
		new_st := store_add(cc.Store, _mk_typevar_from_name(gp_2))
		_ = new_st
		cc = _cset_store(cc, new_st)
	}
	pnames := []string{""}
	_ = pnames
	ptypes := []int64{int64(0)}
	_ = ptypes
	if (_ck2(cc) == "(") {
		cc = _cadv(cc)
		for ((_ck2(cc) != ")") && (_ck2(cc) != "EOF")) {
			if (_ck2(cc) == "self") {
				pnames = append(pnames, "self")
				self_type := resolve_type_name(cc.Store, impl_type)
				_ = self_type
				ptypes = append(ptypes, self_type)
				cc = _cadv(cc)
			} else if (_ck2(cc) == "IDENT") {
				pn := _ccur(cc).Text
				_ = pn
				cc = _cadv(cc)
				if (_ck2(cc) == ":") {
					cc = _cadv(cc)
					pt := _parse_type_ann(cc)
					_ = pt
					cc = pt.C
					pnames = append(pnames, pn)
					ptypes = append(ptypes, pt.Type_id)
				}
			} else {
				cc = _cadv(cc)
			}
			if (_ck2(cc) == ",") {
				cc = _cadv(cc)
			}
		}
		if (_ck2(cc) == ")") {
			cc = _cadv(cc)
		}
	}
	ret := TY_UNIT
	_ = ret
	err := int64(0)
	_ = err
	if (_ck2(cc) == "->") {
		cc = _cadv(cc)
		rt := _parse_type_ann(cc)
		_ = rt
		cc = rt.C
		rtinfo := store_get(cc.Store, rt.Type_id)
		_ = rtinfo
		if (type_kind_name(rtinfo.Kind) == "Result") {
			ret = rtinfo.Type_id
			err = rtinfo.Type_id2
		} else {
			ret = rt.Type_id
		}
	}
	if ((err == int64(0)) && (_ck2(cc) == "!")) {
		cc = _cadv(cc)
		et := _parse_type_ann(cc)
		_ = et
		cc = et.C
		err = et.Type_id
	}
	eff_io := false
	_ = eff_io
	eff_fs := false
	_ = eff_fs
	eff_net := false
	_ = eff_net
	eff_ffi := false
	_ = eff_ffi
	eff_async := false
	_ = eff_async
	if (_ck2(cc) == "with") {
		cc = _cadv(cc)
		if (_ck2(cc) == "[") {
			cc = _cadv(cc)
			for ((_ck2(cc) != "]") && (_ck2(cc) != "EOF")) {
				if (_ck2(cc) == "IDENT") {
					eff := _ccur(cc).Text
					_ = eff
					if (eff == "Io") {
						eff_io = true
					}
					if (eff == "Fs") {
						eff_fs = true
					}
					if (eff == "Net") {
						eff_net = true
					}
					if (eff == "Ffi") {
						eff_ffi = true
					}
					if (eff == "Async") {
						eff_async = true
					}
					cc = _cadv(cc)
				} else {
					cc = _cadv(cc)
				}
				if (_ck2(cc) == ",") {
					cc = _cadv(cc)
				}
			}
			if (_ck2(cc) == "]") {
				cc = _cadv(cc)
			}
		}
	}
	if (_ck2(cc) == "where") {
		cc = _cadv(cc)
		parsing_where := true
		_ = parsing_where
		for parsing_where {
			if (_ck2(cc) == "IDENT") {
				wname := _ccur(cc).Text
				_ = wname
				cc = _cadv(cc)
				if (_ck2(cc) == ":") {
					cc = _cadv(cc)
					if (_ck2(cc) == "IDENT") {
						wbound := _ccur(cc).Text
						_ = wbound
						cc = _cadv(cc)
						for (_ck2(cc) == "+") {
							cc = _cadv(cc)
							if (_ck2(cc) == "IDENT") {
								wbound = ((wbound + "+") + _ccur(cc).Text)
								cc = _cadv(cc)
							}
						}
						if ((wname == gp_0) && (gb_0 == "")) {
							gb_0 = wbound
						}
						if ((wname == gp_1) && (gb_1 == "")) {
							gb_1 = wbound
						}
						if ((wname == gp_2) && (gb_2 == "")) {
							gb_2 = wbound
						}
					}
				}
				if (_ck2(cc) == ",") {
					cc = _cadv(cc)
				} else {
					parsing_where = false
				}
			} else {
				parsing_where = false
			}
		}
	}
	sig := new_fn_sig(name, pnames, ptypes, ret, err)
	_ = sig
	if (gp_count > int64(0)) {
		sig = new_generic_fn_sig(name, pnames, ptypes, ret, err, gp_count, _strip_hkt_suffix(gp_0), _strip_hkt_suffix(gp_1), _strip_hkt_suffix(gp_2), gb_0, gb_1, gb_2, token_start)
	}
	if ((((eff_io || eff_fs) || eff_net) || eff_ffi) || eff_async) {
		sig = sig_set_effects(sig, eff_io, eff_fs, eff_net, eff_ffi, eff_async)
	}
	cc = _cset_reg(cc, reg_add_fn(cc.Registry, sig))
	lb := "{"
	_ = lb
	for (_ck2(cc) != "EOF") {
		if (_ck2(cc) == lb) {
			cc = _skip_braces2(cc)
			return cc
		}
		if (_ck2(cc) == "NEWLINE") {
			return cc
		}
		cc = _cadv(cc)
	}
	return cc
}

func _register_const_ast(c Checker, node_idx int64) Checker {
	node := pool_get(c.Pool, node_idx)
	_ = node
	name := node.S1
	_ = name
	cc := c
	_ = cc
	ctype := TY_UNKNOWN
	_ = ctype
	if (node.C1 > int64(0)) {
		cc = _cset_pos(cc, node.C1)
		ct := _parse_type_ann(cc)
		_ = ct
		cc = ct.C
		ctype = ct.Type_id
	}
	cc = _env_add(cc, name, ctype, int64(0))
	return cc
}

func _register_const(c Checker) Checker {
	cc := _cadv(c)
	_ = cc
	if (_ck2(cc) != "IDENT") {
		return cc
	}
	name := _ccur(cc).Text
	_ = name
	cc = _cadv(cc)
	ctype := TY_UNKNOWN
	_ = ctype
	if (_ck2(cc) == ":") {
		cc = _cadv(cc)
		ct := _parse_type_ann(cc)
		_ = ct
		cc = ct.C
		ctype = ct.Type_id
	}
	cc = _env_add(cc, name, ctype, int64(0))
	if (_ck2(cc) == "=") {
		cc = _cadv(cc)
		for ((_ck2(cc) != "NEWLINE") && (_ck2(cc) != "EOF")) {
			cc = _cadv(cc)
		}
	}
	return cc
}

func _pass3_bodies(c Checker, index DeclIndex) Checker {
	cc := c
	_ = cc
	di := int64(1)
	_ = di
	for (di < decl_index_len(index)) {
		decl := decl_index_get(index, di)
		_ = decl
		kname := decl_kind_name(decl.Kind)
		_ = kname
		cc = _cset_pos(cc, decl.Token_start)
		cc = _cskip_pub(cc)
		if (_ck2(cc) == "@") {
			cc = _cadv(cc)
			if (_ck2(cc) == "IDENT") {
				cc = _cadv(cc)
			}
			if (_ck2(cc) == "(") {
				cc = _cadv(cc)
				for ((_ck2(cc) != ")") && (_ck2(cc) != "EOF")) {
					cc = _cadv(cc)
				}
				if (_ck2(cc) == ")") {
					cc = _cadv(cc)
				}
			}
			cc = _cskip_nl(cc)
		}
		if (kname == "fn") {
			if ((decl.Name == "_extern") || strings.HasPrefix(decl.Name, "_fixture:")) {
				cc = _cset_pos(cc, decl.Body_end)
			} else {
				cc = _check_fn_d(cc, decl)
			}
		} else if (kname == "impl") {
			cc = _check_impl_d(cc, decl)
		} else if (kname == "entry") {
			cc = _check_entry(cc)
		} else if (kname == "test") {
			cc = _check_test(cc)
		} else {
			cc = _cset_pos(cc, decl.Body_end)
		}
		di = (di + int64(1))
	}
	return cc
}

func _skip_decl(c Checker) Checker {
	cc := _cadv(c)
	_ = cc
	lb := "{"
	_ = lb
	for (_ck2(cc) != "EOF") {
		if (_ck2(cc) == lb) {
			cc = _skip_braces2(cc)
			cc = _cskip_nl(cc)
			if (_ck2(cc) == "derives") {
				for ((_ck2(cc) != "NEWLINE") && (_ck2(cc) != "EOF")) {
					cc = _cadv(cc)
				}
			}
			return cc
		}
		if (_ck2(cc) == "NEWLINE") {
			cc = _cskip_nl(cc)
			nk := _ck2(cc)
			_ = nk
			if (((((((((((nk == "fn") || (nk == "pub")) || (nk == "type")) || (nk == "struct")) || (nk == "enum")) || (nk == "trait")) || (nk == "impl")) || (nk == "const")) || (nk == "entry")) || (nk == "test")) || (nk == "EOF")) {
				return cc
			}
		}
		cc = _cadv(cc)
	}
	return cc
}

func _check_fn_d(c Checker, decl DeclInfo) Checker {
	name := decl.Name
	_ = name
	cc := c
	_ = cc
	sig := reg_find_fn(cc.Registry, name)
	_ = sig
	if fn_is_generic(sig) {
		return _cset_pos(cc, decl.Body_end)
	}
	cc = _cset_fn_ctx(cc, sig.Return_type, sig.Error_type)
	cc = Checker{Tokens: cc.Tokens, Pos: cc.Pos, Store: cc.Store, Registry: cc.Registry, Treg: cc.Treg, Diagnostics: cc.Diagnostics, Table: cc.Table, Current_scope: cc.Current_scope, File: cc.File, Pool: cc.Pool, Env_names: cc.Env_names, Env_types: cc.Env_types, Env_scopes: cc.Env_scopes, Fn_return_type: cc.Fn_return_type, Fn_error_type: cc.Fn_error_type, Fn_current_name: name, In_loop: cc.In_loop, In_callee_slot: cc.In_callee_slot}
	pi := int64(1)
	_ = pi
	for (pi < int64(len(sig.Param_names))) {
		cc = _env_add(cc, sig.Param_names[pi], sig.Param_types[pi], int64(0))
		pi = (pi + int64(1))
	}
	cc = _cset_pos(cc, decl.Body_start)
	lb := "{"
	_ = lb
	if (_ck2(cc) == "=") {
		cc = _cadv(cc)
		cc = _cskip_nl(cc)
		result := _check_expr_full(cc)
		_ = result
		cc = result.C
		if ((sig.Return_type != TY_UNKNOWN) && (sig.Return_type != TY_UNIT)) {
			rt_ok := is_assignable(cc.Store, result.Type_id, sig.Return_type)
			_ = rt_ok
			if ((rt_ok == false) && (result.Type_id != TY_UNKNOWN)) {
				rinfo := store_get(cc.Store, result.Type_id)
				_ = rinfo
				if ((type_kind_name(rinfo.Kind) == "Result") && (sig.Error_type > int64(0))) {
					rt_ok = true
				}
			}
			if ((rt_ok == false) && (result.Type_id != TY_UNKNOWN)) {
				rinfo2 := store_get(cc.Store, result.Type_id)
				_ = rinfo2
				if ((type_kind_name(rinfo2.Kind) == "Array") && (rinfo2.Type_id == TY_UNKNOWN)) {
					rt_ok = true
				}
			}
			if ((rt_ok == false) && (result.Type_id != TY_UNKNOWN)) {
				q := _q()
				_ = q
				cc = _cerror(cc, E0106, ((((((("expected return type " + q) + format_type(cc.Store, sig.Return_type)) + q) + ", got ") + q) + format_type(cc.Store, result.Type_id)) + q))
			}
		}
	} else if (_ck2(cc) == lb) {
		result := _check_block(cc)
		_ = result
		cc = result.C
	} else {
		cc = _cskip_nl(cc)
		if (_ck2(cc) == lb) {
			result := _check_block(cc)
			_ = result
			cc = result.C
		}
	}
	return cc
}

func _check_impl_d(c Checker, decl DeclInfo) Checker {
	cc := _cadv(c)
	_ = cc
	impl_type := ""
	_ = impl_type
	if (_ck2(cc) == "IDENT") {
		impl_type = _ccur(cc).Text
		cc = _cadv(cc)
	}
	if (_ck2(cc) == "for") {
		cc = _cadv(cc)
		if (_ck2(cc) == "IDENT") {
			impl_type = _ccur(cc).Text
			cc = _cadv(cc)
		}
	}
	if (decl.Node_idx > int64(0)) {
		impl_node := pool_get(cc.Pool, decl.Node_idx)
		_ = impl_node
		mi := int64(0)
		_ = mi
		for (mi < impl_node.List_count) {
			midx_node := pool_get(cc.Pool, (impl_node.List_start + mi))
			_ = midx_node
			method_node := pool_get(cc.Pool, midx_node.C0)
			_ = method_node
			cc = _check_impl_fn_ast(cc, impl_type, method_node)
			mi = (mi + int64(1))
		}
		cc = _cset_pos(cc, decl.Body_end)
		return cc
	}
	lb := "{"
	_ = lb
	for ((_ck2(cc) != lb) && (_ck2(cc) != "EOF")) {
		cc = _cadv(cc)
	}
	if (_ck2(cc) != lb) {
		return cc
	}
	cc = _cadv(cc)
	cc = _cskip_nl(cc)
	rb := "}"
	_ = rb
	for ((_ck2(cc) != rb) && (_ck2(cc) != "EOF")) {
		cc = _cskip_nl(cc)
		cc = _cskip_pub(cc)
		if (_ck2(cc) == "fn") {
			cc = _check_impl_fn(cc, impl_type)
		} else if ((_ck2(cc) == rb) || (_ck2(cc) == "EOF")) {
			return cc
		} else {
			cc = _cadv(cc)
		}
		cc = _cskip_nl(cc)
	}
	if (_ck2(cc) == rb) {
		cc = _cadv(cc)
	}
	return cc
}

func _check_impl_fn_ast(c Checker, impl_type string, method_node Expr) Checker {
	raw_name := method_node.S1
	_ = raw_name
	name := ((impl_type + "_") + raw_name)
	_ = name
	cc := c
	_ = cc
	sig := reg_find_fn(cc.Registry, name)
	_ = sig
	if fn_is_generic(sig) {
		return cc
	}
	if (method_node.B1 == false) {
		return cc
	}
	cc = _cset_fn_ctx(cc, sig.Return_type, sig.Error_type)
	pi := int64(1)
	_ = pi
	for (pi < int64(len(sig.Param_names))) {
		cc = _env_add(cc, sig.Param_names[pi], sig.Param_types[pi], int64(0))
		pi = (pi + int64(1))
	}
	tlen := int64(len(cc.Tokens))
	_ = tlen
	search_pos := cc.Pos
	_ = search_pos
	lb := "{"
	_ = lb
	for (search_pos < tlen) {
		tk := token_name(cc.Tokens[search_pos].Kind)
		_ = tk
		if (tk == "fn") {
			next := (search_pos + int64(1))
			_ = next
			if (next < tlen) {
				if (cc.Tokens[next].Text == raw_name) {
					cc = _cset_pos(cc, (search_pos + int64(2)))
					for (((_ck2(cc) != "=") && (_ck2(cc) != lb)) && (_ck2(cc) != "EOF")) {
						cc = _cadv(cc)
					}
					if (_ck2(cc) == "=") {
						cc = _cadv(cc)
						cc = _cskip_nl(cc)
						result := _check_expr_full(cc)
						_ = result
						cc = result.C
					} else if (_ck2(cc) == lb) {
						result := _check_block(cc)
						_ = result
						cc = result.C
					}
					return cc
				}
			}
		}
		search_pos = (search_pos + int64(1))
	}
	return cc
}

func _check_impl_fn(c Checker, impl_type string) Checker {
	cc := _cadv(c)
	_ = cc
	if (_ck2(cc) != "IDENT") {
		return cc
	}
	raw_name := _ccur(cc).Text
	_ = raw_name
	name := ((impl_type + "_") + raw_name)
	_ = name
	cc = _cadv(cc)
	sig := reg_find_fn(cc.Registry, name)
	_ = sig
	if fn_is_generic(sig) {
		return _skip_decl(cc)
	}
	cc = _cset_fn_ctx(cc, sig.Return_type, sig.Error_type)
	pi := int64(1)
	_ = pi
	for (pi < int64(len(sig.Param_names))) {
		cc = _env_add(cc, sig.Param_names[pi], sig.Param_types[pi], int64(0))
		pi = (pi + int64(1))
	}
	lb := "{"
	_ = lb
	for (((_ck2(cc) != "=") && (_ck2(cc) != lb)) && (_ck2(cc) != "EOF")) {
		cc = _cadv(cc)
	}
	if (_ck2(cc) == "=") {
		cc = _cadv(cc)
		cc = _cskip_nl(cc)
		result := _check_expr_full(cc)
		_ = result
		cc = result.C
	} else if (_ck2(cc) == lb) {
		result := _check_block(cc)
		_ = result
		cc = result.C
	}
	return cc
}

func _check_entry(c Checker) Checker {
	cc := _cadv(c)
	_ = cc
	cc = _cset_fn_ctx(cc, TY_UNIT, int64(0))
	cc = Checker{Tokens: cc.Tokens, Pos: cc.Pos, Store: cc.Store, Registry: cc.Registry, Treg: cc.Treg, Diagnostics: cc.Diagnostics, Table: cc.Table, Current_scope: cc.Current_scope, File: cc.File, Pool: cc.Pool, Env_names: cc.Env_names, Env_types: cc.Env_types, Env_scopes: cc.Env_scopes, Fn_return_type: cc.Fn_return_type, Fn_error_type: cc.Fn_error_type, Fn_current_name: "_entry", In_loop: cc.In_loop, In_callee_slot: cc.In_callee_slot}
	result := _check_block(cc)
	_ = result
	return result.C
}

func _check_test(c Checker) Checker {
	cc := _cadv(c)
	_ = cc
	if ((_ck2(cc) == "STRING") || (_ck2(cc) == "IDENT")) {
		cc = _cadv(cc)
	}
	cc = _cset_fn_ctx(cc, TY_UNIT, int64(0))
	result := _check_block(cc)
	_ = result
	return result.C
}

func _is_cont(k string) bool {
	return (((((((((((((((((((((((((((((((((k == "+") || (k == "-")) || (k == "*")) || (k == "/")) || (k == "%")) || (k == "==")) || (k == "!=")) || (k == "<")) || (k == ">")) || (k == "<=")) || (k == ">=")) || (k == "&&")) || (k == "||")) || (k == "&")) || (k == "|")) || (k == "^")) || (k == "~")) || (k == "<<")) || (k == ">>")) || (k == "|>")) || (k == "..")) || (k == "..=")) || (k == ".")) || (k == "(")) || (k == "[")) || (k == "?")) || (k == "!")) || (k == ",")) || (k == "=>")) || (k == "=")) || (k == "?.")) || (k == "??")) || (k == "catch"))
}

func _check_block(c Checker) CE {
	lb := "{"
	_ = lb
	rb := "}"
	_ = rb
	if (_ck2(c) != lb) {
		return CE{C: c, Type_id: TY_UNIT}
	}
	cc := _cadv(c)
	_ = cc
	cc = _cskip_nl(cc)
	saved_env_len := int64(len(cc.Env_names))
	_ = saved_env_len
	last_type := TY_UNIT
	_ = last_type
	for ((_ck2(cc) != rb) && (_ck2(cc) != "EOF")) {
		result := _check_stmt(cc)
		_ = result
		cc = result.C
		last_type = result.Type_id
		cc = _cskip_nl(cc)
	}
	if (_ck2(cc) == rb) {
		cc = _cadv(cc)
	}
	cc = _env_trim(cc, saved_env_len)
	return CE{C: cc, Type_id: last_type}
}

func _check_stmt(c Checker) CE {
	k := _ck2(c)
	_ = k
	if (k == "mut") {
		return _check_mut_binding(c)
	}
	if (k == "return") {
		return _check_return(c)
	}
	if (k == "break") {
		return CE{C: _cadv(c), Type_id: TY_NEVER}
	}
	if (k == "continue") {
		return CE{C: _cadv(c), Type_id: TY_NEVER}
	}
	if (k == "for") {
		return _check_for(c)
	}
	if (k == "while") {
		return _check_while(c)
	}
	if (k == "loop") {
		return _check_loop(c)
	}
	if (k == "if") {
		return _check_if(c)
	}
	if (k == "match") {
		return _check_match(c)
	}
	if (k == "defer") {
		cc := _cadv(c)
		_ = cc
		result := _check_expr_full(cc)
		_ = result
		return CE{C: result.C, Type_id: TY_UNIT}
	}
	if (k == "assert") {
		cc := _cadv(c)
		_ = cc
		result := _check_expr_full(cc)
		_ = result
		return CE{C: result.C, Type_id: TY_UNIT}
	}
	if (k == "with") {
		cc := _cadv(c)
		_ = cc
		if (_ck2(cc) != "[") {
			res_name := ""
			_ = res_name
			if (_ck2(cc) == "IDENT") {
				next_pos := (cc.Pos + int64(1))
				_ = next_pos
				if ((next_pos < int64(len(cc.Tokens))) && (token_name(cc.Tokens[next_pos].Kind) == ":=")) {
					res_name = _ccur(cc).Text
					cc = _cadv(cc)
					cc = _cadv(cc)
				}
			}
			result := _check_expr_full(cc)
			_ = result
			cc = result.C
			if (int64(len(res_name)) > int64(0)) {
				cc = _env_add(cc, res_name, result.Type_id, int64(0))
			}
			body := _check_block(cc)
			_ = body
			return CE{C: body.C, Type_id: body.Type_id}
		}
	}
	return _check_expr_or_binding(c)
}

func _check_mut_binding(c Checker) CE {
	cc := _cadv(c)
	_ = cc
	if (_ck2(cc) != "IDENT") {
		return CE{C: cc, Type_id: TY_UNIT}
	}
	name := _ccur(cc).Text
	_ = name
	cc = _cadv(cc)
	ann_type := TY_UNKNOWN
	_ = ann_type
	if (_ck2(cc) == ":") {
		cc = _cadv(cc)
		at := _parse_type_ann(cc)
		_ = at
		cc = at.C
		ann_type = at.Type_id
		if (_ck2(cc) == "=") {
			cc = _cadv(cc)
		}
	} else if (_ck2(cc) == ":=") {
		cc = _cadv(cc)
	}
	result := _check_expr_full(cc)
	_ = result
	cc = result.C
	final_type := result.Type_id
	_ = final_type
	if (ann_type != TY_UNKNOWN) {
		if ((result.Type_id != TY_UNKNOWN) && (is_assignable(cc.Store, result.Type_id, ann_type) == false)) {
			q := _q()
			_ = q
			cc = _cerror(cc, E0100, ((((((("type mismatch in binding: expected " + q) + format_type(cc.Store, ann_type)) + q) + ", got ") + q) + format_type(cc.Store, result.Type_id)) + q))
		}
		final_type = ann_type
	}
	cc = _env_add(cc, name, final_type, int64(0))
	return CE{C: cc, Type_id: TY_UNIT}
}

func _check_return(c Checker) CE {
	cc := _cadv(c)
	_ = cc
	rb := "}"
	_ = rb
	if (((_ck2(cc) == "NEWLINE") || (_ck2(cc) == "EOF")) || (_ck2(cc) == rb)) {
		return CE{C: cc, Type_id: TY_NEVER}
	}
	result := _check_expr_full(cc)
	_ = result
	cc = result.C
	if ((cc.Fn_return_type != TY_UNKNOWN) && (cc.Fn_return_type != TY_UNIT)) {
		type_ok := is_assignable(cc.Store, result.Type_id, cc.Fn_return_type)
		_ = type_ok
		if ((type_ok == false) && (result.Type_id != TY_UNKNOWN)) {
			rinfo := store_get(cc.Store, result.Type_id)
			_ = rinfo
			if ((type_kind_name(rinfo.Kind) == "Result") && (cc.Fn_error_type > int64(0))) {
				type_ok = true
			}
		}
		if ((type_ok == false) && (result.Type_id != TY_UNKNOWN)) {
			rinfo := store_get(cc.Store, result.Type_id)
			_ = rinfo
			if ((type_kind_name(rinfo.Kind) == "Array") && (rinfo.Type_id == TY_UNKNOWN)) {
				type_ok = true
			}
		}
		if ((type_ok == false) && (result.Type_id != TY_UNKNOWN)) {
			q := _q()
			_ = q
			cc = _cerror(cc, E0106, ((((((("return type mismatch: expected " + q) + format_type(cc.Store, cc.Fn_return_type)) + q) + ", got ") + q) + format_type(cc.Store, result.Type_id)) + q))
		}
	}
	return CE{C: cc, Type_id: TY_NEVER}
}

func _check_expr_or_binding(c Checker) CE {
	if (_ck2(c) == "(") {
		scan := (c.Pos + int64(1))
		_ = scan
		is_tuple_destr := false
		_ = is_tuple_destr
		for (scan < int64(len(c.Tokens))) {
			sk := token_name(c.Tokens[scan].Kind)
			_ = sk
			if (sk == ")") {
				next2 := (scan + int64(1))
				_ = next2
				if ((next2 < int64(len(c.Tokens))) && (token_name(c.Tokens[next2].Kind) == ":=")) {
					is_tuple_destr = true
				}
				scan = int64(len(c.Tokens))
			} else if (((sk == "IDENT") || (sk == ",")) || (sk == "_")) {
				scan = (scan + int64(1))
			} else {
				scan = int64(len(c.Tokens))
			}
		}
		if is_tuple_destr {
			cc := _cadv(c)
			_ = cc
			names := []string{""}
			_ = names
			for ((_ck2(cc) != ")") && (_ck2(cc) != "EOF")) {
				if (_ck2(cc) == "IDENT") {
					names = append(names, _ccur(cc).Text)
				}
				cc = _cadv(cc)
				if (_ck2(cc) == ",") {
					cc = _cadv(cc)
				}
			}
			if (_ck2(cc) == ")") {
				cc = _cadv(cc)
			}
			if (_ck2(cc) == ":=") {
				cc = _cadv(cc)
			}
			result := _check_expr_full(cc)
			_ = result
			cc = result.C
			ni := int64(1)
			_ = ni
			for (ni < int64(len(names))) {
				elem_t := tuple_elem_type(cc.Store, result.Type_id, (ni - int64(1)))
				_ = elem_t
				cc = _env_add(cc, names[ni], elem_t, int64(0))
				ni = (ni + int64(1))
			}
			return CE{C: cc, Type_id: TY_UNIT}
		}
	}
	if (_ck2(c) == "IDENT") {
		name := _ccur(c).Text
		_ = name
		next_pos := (c.Pos + int64(1))
		_ = next_pos
		if (next_pos < int64(len(c.Tokens))) {
			next_k := token_name(c.Tokens[next_pos].Kind)
			_ = next_k
			if (next_k == ":=") {
				cc := _cadv(c)
				_ = cc
				cc = _cadv(cc)
				result := _check_expr_full(cc)
				_ = result
				cc = result.C
				cc = _env_add(cc, name, result.Type_id, int64(0))
				return CE{C: cc, Type_id: TY_UNIT}
			}
			if (next_k == "=") {
				sym := lookup(c.Table, c.Current_scope, name)
				_ = sym
				if ((sym.Name != "") && (sym.Is_mutable == false)) {
					q := _q()
					_ = q
					cc := _cerror(c, "E0500", ((("cannot assign to immutable binding " + q) + name) + q))
					_ = cc
					cc = _cadv(cc)
					cc = _cadv(cc)
					result := _check_expr_full(cc)
					_ = result
					return CE{C: result.C, Type_id: TY_UNIT}
				}
			}
			if (next_k == ":") {
				cc := _cadv(c)
				_ = cc
				cc = _cadv(cc)
				saved_pos := cc.Pos
				_ = saved_pos
				at := _parse_type_ann(cc)
				_ = at
				cc = at.C
				if (_ck2(cc) == "=") {
					cc = _cadv(cc)
					result := _check_expr_full(cc)
					_ = result
					cc = result.C
					cc = _env_add(cc, name, at.Type_id, int64(0))
					return CE{C: cc, Type_id: TY_UNIT}
				}
				cc = _cset_pos(cc, c.Pos)
			}
		}
	}
	return _check_expr_full(c)
}

func _check_expr_full(c Checker) CE {
	result := _check_primary(c)
	_ = result
	cc := result.C
	_ = cc
	left_type := result.Type_id
	_ = left_type
	running := true
	_ = running
	for running {
		k := _ck2(cc)
		_ = k
		if (k == ".") {
			cc = _cadv(cc)
			if (_ck2(cc) != "IDENT") {
				running = false
			} else {
				field := _ccur(cc).Text
				_ = field
				cc = _cadv(cc)
				if ((_ck2(cc) == "[") && (((field == "to") || (field == "trunc")))) {
					cc = _cadv(cc)
					conv_type := TY_UNKNOWN
					_ = conv_type
					if (_ck2(cc) == "IDENT") {
						ct := _parse_type_ann(cc)
						_ = ct
						cc = ct.C
						conv_type = ct.Type_id
					}
					if (_ck2(cc) == "]") {
						cc = _cadv(cc)
					}
					if (_ck2(cc) == "(") {
						cc = _cadv(cc)
					}
					if (_ck2(cc) == ")") {
						cc = _cadv(cc)
					}
					if (field == "to") {
						new_st := store_add(cc.Store, mk_result_type(conv_type, TY_STR))
						_ = new_st
						cc = _cset_store(cc, new_st)
						left_type = store_last_id(cc.Store)
					} else {
						left_type = conv_type
					}
				} else if (_ck2(cc) == "(") {
					mr := _check_method_call(cc, left_type, field)
					_ = mr
					cc = mr.C
					left_type = mr.Type_id
				} else {
					fr := _check_field(cc, left_type, field)
					_ = fr
					cc = fr.C
					left_type = fr.Type_id
				}
			}
		} else if (k == "(") {
			cr := _check_call_args(cc, left_type)
			_ = cr
			cc = cr.C
			left_type = cr.Type_id
		} else if (k == "[") {
			cc = _cadv(cc)
			idx := _check_expr_full(cc)
			_ = idx
			cc = idx.C
			if (_ck2(cc) == "]") {
				cc = _cadv(cc)
			}
			info := store_get(cc.Store, left_type)
			_ = info
			if (type_kind_name(info.Kind) == "Array") {
				left_type = info.Type_id
			} else if (type_kind_name(info.Kind) == "str") {
				left_type = TY_STR
			} else {
				left_type = TY_UNKNOWN
			}
		} else if (k == "?") {
			cc = _cadv(cc)
			info := store_get(cc.Store, left_type)
			_ = info
			if (type_kind_name(info.Kind) == "Result") {
				left_type = info.Type_id
			} else if (type_kind_name(info.Kind) == "Optional") {
				left_type = info.Type_id
			}
		} else if (k == "!") {
			cc = _cadv(cc)
			info := store_get(cc.Store, left_type)
			_ = info
			if (type_kind_name(info.Kind) == "Result") {
				left_type = info.Type_id
			}
		} else if (k == "catch") {
			cc = _cadv(cc)
			if (_ck2(cc) == "|") {
				cc = _cadv(cc)
				if ((_ck2(cc) == "IDENT") || (_ck2(cc) == "_")) {
					cc = _cadv(cc)
				}
				if (_ck2(cc) == "|") {
					cc = _cadv(cc)
				}
			}
			cb := _check_block(cc)
			_ = cb
			cc = cb.C
			info := store_get(cc.Store, left_type)
			_ = info
			if (type_kind_name(info.Kind) == "Result") {
				left_type = info.Type_id
			}
		} else if ((k == "IDENT") && (_ccur(cc).Text == "or")) {
			cc = _cadv(cc)
			right := _check_primary(cc)
			_ = right
			cc = right.C
			info := store_get(cc.Store, left_type)
			_ = info
			if (type_kind_name(info.Kind) == "Result") {
				left_type = info.Type_id
			} else {
				left_type = right.Type_id
			}
		} else if ((k == "IDENT") && (_ccur(cc).Text == "must")) {
			cc = _cadv(cc)
			right := _check_primary(cc)
			_ = right
			cc = right.C
			info := store_get(cc.Store, left_type)
			_ = info
			if (type_kind_name(info.Kind) == "Result") {
				left_type = info.Type_id
			}
		} else if (k == "??") {
			cc = _cadv(cc)
			right := _check_primary(cc)
			_ = right
			cc = right.C
			info := store_get(cc.Store, left_type)
			_ = info
			if (type_kind_name(info.Kind) == "Optional") {
				left_type = info.Type_id
			} else {
				left_type = right.Type_id
			}
		} else if (k == "?.") {
			cc = _cadv(cc)
			if (_ck2(cc) == "IDENT") {
				cc = _cadv(cc)
			}
		} else if ((((((((((((((((((k == "+") || (k == "-")) || (k == "*")) || (k == "/")) || (k == "%")) || (k == "==")) || (k == "!=")) || (k == "<")) || (k == ">")) || (k == "<=")) || (k == ">=")) || (k == "&&")) || (k == "||")) || (k == "&")) || (k == "|")) || (k == "^")) || (k == "<<")) || (k == ">>")) {
			op := _ccur(cc).Text
			_ = op
			cc = _cadv(cc)
			right := _check_primary(cc)
			_ = right
			cc = right.C
			rr_running := true
			_ = rr_running
			for rr_running {
				rk := _ck2(cc)
				_ = rk
				if (((((rk == ".") || (rk == "(")) || (rk == "[")) || (rk == "?")) || (rk == "!")) {
					post := _check_postfix_chain(cc, right.Type_id)
					_ = post
					cc = post.C
					right = CE{C: cc, Type_id: post.Type_id}
				} else {
					rr_running = false
				}
			}
			if (((type_supports_op(cc.Treg, cc.Store, left_type, op) == false) && (left_type != TY_UNKNOWN)) && (left_type != TY_BOOL)) {
				q := _q()
				_ = q
				cc = _cerror(cc, E0102, ((((((("operator " + q) + op) + q) + " not supported for type ") + q) + format_type(cc.Store, left_type)) + q))
			}
			if (((op != "&&") && (op != "||")) && (left_type != TY_BOOL)) {
				if ((left_type != TY_UNKNOWN) && (right.Type_id != TY_UNKNOWN)) {
					if (types_equal(cc.Store, left_type, right.Type_id) == false) {
						is_str_concat := ((left_type == TY_STR) && (right.Type_id == TY_STR))
						_ = is_str_concat
						is_cmp := ((((((op == "==") || (op == "!=")) || (op == "<")) || (op == ">")) || (op == "<=")) || (op == ">="))
						_ = is_cmp
						if ((is_str_concat == false) && is_cmp) {
							q := _q()
							_ = q
							cc = _cerror(cc, E0102, ((((((("cannot compare " + q) + format_type(cc.Store, left_type)) + q) + " with ") + q) + format_type(cc.Store, right.Type_id)) + q))
						}
					}
				}
			}
			left_type = binary_result_type(cc.Store, left_type, op)
		} else if (k == "|>") {
			cc = _cadv(cc)
			if (_ck2(cc) == "IDENT") {
				fname := _ccur(cc).Text
				_ = fname
				cc = _cadv(cc)
				sig := reg_find_fn(cc.Registry, fname)
				_ = sig
				if (sig.Name != "") {
					if fn_is_generic(sig) {
						pipe_ret := sig.Return_type
						_ = pipe_ret
						tv_id := int64(0)
						_ = tv_id
						if (sig.Generic_count > int64(0)) {
							tv_id = resolve_type_name(cc.Store, sig.Generic_name_0)
						}
						if ((tv_id > int64(0)) && (pipe_ret == tv_id)) {
							pipe_ret = left_type
						}
						spec_name := ((fname + "_") + _type_name_suffix(cc.Store, left_type))
						_ = spec_name
						existing := reg_find_mono(cc.Registry, spec_name)
						_ = existing
						if (existing.Specialized_name == "") {
							sp_ptypes := []int64{int64(0)}
							_ = sp_ptypes
							spi := int64(1)
							_ = spi
							for (spi < int64(len(sig.Param_types))) {
								pt := sig.Param_types[spi]
								_ = pt
								if ((pt == tv_id) && (tv_id > int64(0))) {
									pt = left_type
								}
								sp_ptypes = append(sp_ptypes, pt)
								spi = (spi + int64(1))
							}
							sp_sig := new_fn_sig(spec_name, sig.Param_names, sp_ptypes, pipe_ret, sig.Error_type)
							_ = sp_sig
							cc = _cset_reg(cc, reg_add_fn(cc.Registry, sp_sig))
							spec := MonoSpec{Generic_name: sig.Name, Specialized_name: spec_name, Type_arg_0: left_type, Type_arg_1: int64(0), Type_arg_2: int64(0)}
							_ = spec
							cc = _cset_reg(cc, reg_add_mono(cc.Registry, spec))
						}
						left_type = pipe_ret
					} else {
						left_type = sig.Return_type
					}
				} else {
					left_type = TY_UNKNOWN
				}
			}
		} else if (k == "=") {
			cc = _cadv(cc)
			rhs := _check_expr_full(cc)
			_ = rhs
			cc = rhs.C
			if (((left_type != TY_UNKNOWN) && (rhs.Type_id != TY_UNKNOWN)) && (left_type != TY_UNIT)) {
				if (is_assignable(cc.Store, rhs.Type_id, left_type) == false) {
					q := _q()
					_ = q
					cc = _cerror(cc, E0100, ((((((("assignment type mismatch: expected " + q) + format_type(cc.Store, left_type)) + q) + ", got ") + q) + format_type(cc.Store, rhs.Type_id)) + q))
				}
			}
			left_type = TY_UNIT
			running = false
		} else {
			if (_is_cont(k) == false) {
				running = false
			} else {
				running = false
			}
		}
	}
	return CE{C: cc, Type_id: left_type}
}

func _check_postfix_chain(c Checker, base_type int64) CE {
	cc := c
	_ = cc
	cur_type := base_type
	_ = cur_type
	running := true
	_ = running
	for running {
		k := _ck2(cc)
		_ = k
		if (k == ".") {
			cc = _cadv(cc)
			if (_ck2(cc) == "IDENT") {
				field := _ccur(cc).Text
				_ = field
				cc = _cadv(cc)
				if (_ck2(cc) == "(") {
					mr := _check_method_call(cc, cur_type, field)
					_ = mr
					cc = mr.C
					cur_type = mr.Type_id
				} else {
					fr := _check_field(cc, cur_type, field)
					_ = fr
					cc = fr.C
					cur_type = fr.Type_id
				}
			} else if (_ck2(cc) == "INT") {
				field_idx_str := _ccur(cc).Text
				_ = field_idx_str
				cc = _cadv(cc)
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
				elem_tid := tuple_elem_type(cc.Store, cur_type, fidx)
				_ = elem_tid
				if (elem_tid != TY_UNKNOWN) {
					cur_type = elem_tid
				} else {
					cur_type = TY_UNKNOWN
				}
			} else if (_ck2(cc) == "{") {
				rb := "}"
				_ = rb
				cc = _cadv(cc)
				cc = _cskip_nl(cc)
				for ((_ck2(cc) != rb) && (_ck2(cc) != "EOF")) {
					if (_ck2(cc) == "IDENT") {
						cc = _cadv(cc)
					}
					if (_ck2(cc) == ":") {
						cc = _cadv(cc)
					}
					ur := _check_expr_full(cc)
					_ = ur
					cc = ur.C
					cc = _cskip_nl(cc)
					if (_ck2(cc) == ",") {
						cc = _cadv(cc)
					}
					cc = _cskip_nl(cc)
				}
				if (_ck2(cc) == rb) {
					cc = _cadv(cc)
				}
			} else {
				running = false
			}
		} else if (k == "(") {
			cr := _check_call_args(cc, cur_type)
			_ = cr
			cc = cr.C
			cur_type = cr.Type_id
		} else if (k == "[") {
			cc = _cadv(cc)
			idx := _check_expr_full(cc)
			_ = idx
			cc = idx.C
			if (_ck2(cc) == "]") {
				cc = _cadv(cc)
			}
			info := store_get(cc.Store, cur_type)
			_ = info
			if (type_kind_name(info.Kind) == "Array") {
				cur_type = info.Type_id
			} else {
				cur_type = TY_UNKNOWN
			}
		} else if ((k == "?") || (k == "!")) {
			cc = _cadv(cc)
		} else {
			running = false
		}
	}
	return CE{C: cc, Type_id: cur_type}
}

func _check_primary(c Checker) CE {
	k := _ck2(c)
	_ = k
	if (((k == "-") || (k == "!")) || (k == "~")) {
		cc := _cadv(c)
		_ = cc
		result := _check_primary(cc)
		_ = result
		if (k == "!") {
			return CE{C: result.C, Type_id: TY_BOOL}
		}
		return result
	}
	if (k == "INT") {
		return CE{C: _cadv(c), Type_id: TY_I64}
	}
	if (k == "FLOAT") {
		return CE{C: _cadv(c), Type_id: TY_F64}
	}
	if (k == "STRING") {
		return CE{C: _cadv(c), Type_id: TY_STR}
	}
	if (k == "CHAR") {
		return CE{C: _cadv(c), Type_id: TY_CHAR}
	}
	if (k == "STRING_START") {
		cc := _cadv(c)
		_ = cc
		for ((_ck2(cc) != "STRING_END") && (_ck2(cc) != "EOF")) {
			cc = _cadv(cc)
		}
		if (_ck2(cc) == "STRING_END") {
			cc = _cadv(cc)
		}
		return CE{C: cc, Type_id: TY_STR}
	}
	if ((k == "true") || (k == "false")) {
		return CE{C: _cadv(c), Type_id: TY_BOOL}
	}
	if (k == "if") {
		return _check_if(c)
	}
	if (k == "match") {
		return _check_match(c)
	}
	if (k == "IDENT") {
		name := _ccur(c).Text
		_ = name
		cc := _cadv(c)
		_ = cc
		if ((_ck2(cc) == "(") && (name == "Some")) {
			cr := _check_call_args(cc, TY_UNKNOWN)
			_ = cr
			cc = cr.C
			new_st := store_add(cc.Store, mk_optional_type(cr.Type_id))
			_ = new_st
			cc = _cset_store(cc, new_st)
			return CE{C: cc, Type_id: store_last_id(cc.Store)}
		}
		if ((_ck2(cc) == "(") && (((name == "Ok") || (name == "Err")))) {
			cr := _check_call_args(cc, TY_UNKNOWN)
			_ = cr
			cc = cr.C
			if (name == "Ok") {
				if (cc.Fn_error_type > int64(0)) {
					new_st := store_add(cc.Store, mk_result_type(cr.Type_id, cc.Fn_error_type))
					_ = new_st
					cc = _cset_store(cc, new_st)
					return CE{C: cc, Type_id: store_last_id(cc.Store)}
				}
				return CE{C: cc, Type_id: cr.Type_id}
			}
			if (name == "Err") {
				if (cc.Fn_return_type > int64(0)) {
					new_st := store_add(cc.Store, mk_result_type(cc.Fn_return_type, cr.Type_id))
					_ = new_st
					cc = _cset_store(cc, new_st)
					return CE{C: cc, Type_id: store_last_id(cc.Store)}
				}
				return CE{C: cc, Type_id: cr.Type_id}
			}
		}
		if (_ck2(cc) == "(") {
			conv_type := resolve_type_name(cc.Store, name)
			_ = conv_type
			if ((conv_type != TY_UNKNOWN) && (conv_type <= TY_CHAR)) {
				cr := _check_call_args(cc, TY_UNKNOWN)
				_ = cr
				return CE{C: cr.C, Type_id: conv_type}
			}
		}
		if (_ck2(cc) == "(") {
			sig := reg_find_fn(cc.Registry, name)
			_ = sig
			if (sig.Name != "") {
				if reg_is_deprecated(cc.Registry, name) {
					dep_msg := reg_deprecated_msg(cc.Registry, name)
					_ = dep_msg
					cc = _cwarning(cc, "W0005", ((("'" + name) + "' is deprecated: ") + dep_msg))
				}
				if ((((fn_is_pure(sig) == false) && (cc.Fn_current_name != "")) && (cc.Fn_current_name != "_entry")) && (cc.Fn_current_name != "_test")) {
					caller_sig := reg_find_fn(cc.Registry, cc.Fn_current_name)
					_ = caller_sig
					if (caller_sig.Name != "") {
						if (sig.Has_io && (fn_has_effect_t(caller_sig, "Io", cc.Treg) == false)) {
							cc = _cerror(cc, "E0300", (((("function '" + name) + "' requires effect [Io] but '") + cc.Fn_current_name) + "' does not declare [Io]"))
						}
						if (sig.Has_fs && (fn_has_effect_t(caller_sig, "Fs", cc.Treg) == false)) {
							cc = _cerror(cc, "E0301", (((("function '" + name) + "' requires effect [Fs] but '") + cc.Fn_current_name) + "' does not declare [Fs]"))
						}
						if (sig.Has_net && (fn_has_effect_t(caller_sig, "Net", cc.Treg) == false)) {
							cc = _cerror(cc, "E0302", (((("function '" + name) + "' requires effect [Net] but '") + cc.Fn_current_name) + "' does not declare [Net]"))
						}
						if (sig.Has_ffi && (fn_has_effect_t(caller_sig, "Ffi", cc.Treg) == false)) {
							cc = _cerror(cc, "E0303", (((("function '" + name) + "' requires effect [Ffi] but '") + cc.Fn_current_name) + "' does not declare [Ffi]"))
						}
						if (sig.Has_async && (fn_has_effect(caller_sig, "Async") == false)) {
							cc = _cerror(cc, "E0304", (((("function '" + name) + "' requires effect [Async] but '") + cc.Fn_current_name) + "' does not declare [Async]"))
						}
						uei := int64(1)
						_ = uei
						for (uei < int64(len(sig.User_effects))) {
							ue := sig.User_effects[uei]
							_ = ue
							if (fn_has_effect_t(caller_sig, ue, cc.Treg) == false) {
								cc = _cerror(cc, "E0305", (((((((("function '" + name) + "' requires effect [") + ue) + "] but '") + cc.Fn_current_name) + "' does not declare [") + ue) + "]"))
							}
							uei = (uei + int64(1))
						}
					}
				}
				if fn_is_generic(sig) {
					result := _check_generic_call(cc, sig)
					_ = result
					return result
				}
				cr := _check_call_args_sig(cc, sig)
				_ = cr
				return cr
			}
			if (((((((name == "println") || (name == "print")) || (name == "eprintln"))) && (cc.Fn_current_name != "")) && (cc.Fn_current_name != "_entry")) && (cc.Fn_current_name != "_test")) {
				caller_sig := reg_find_fn(cc.Registry, cc.Fn_current_name)
				_ = caller_sig
				if ((caller_sig.Name != "") && (fn_has_effect_t(caller_sig, "Io", cc.Treg) == false)) {
					cc = _cwarning(cc, "W0100", (((name + " requires effect [Io] but ") + cc.Fn_current_name) + " is pure or missing [Io]"))
				}
			}
			if (((cc.Fn_current_name != "") && (cc.Fn_current_name != "_entry")) && (cc.Fn_current_name != "_test")) {
				builtin_eff := _builtin_effect(name)
				_ = builtin_eff
				if (builtin_eff != "") {
					caller_sig := reg_find_fn(cc.Registry, cc.Fn_current_name)
					_ = caller_sig
					if ((caller_sig.Name != "") && (fn_has_effect_t(caller_sig, builtin_eff, cc.Treg) == false)) {
						eff_code := "E0300"
						_ = eff_code
						if (builtin_eff == "Fs") {
							eff_code = "E0301"
						}
						if (builtin_eff == "Net") {
							eff_code = "E0302"
						}
						if (builtin_eff == "Ffi") {
							eff_code = "E0303"
						}
						if (builtin_eff == "Async") {
							eff_code = "E0304"
						}
						cc = _cwarning(cc, eff_code, (((((((("'" + name) + "' requires effect [") + builtin_eff) + "] but '") + cc.Fn_current_name) + "' does not declare [") + builtin_eff) + "]"))
					}
				}
			}
			cr := _check_call_args(cc, TY_UNKNOWN)
			_ = cr
			return CE{C: cr.C, Type_id: TY_UNKNOWN}
		}
		lb := "{"
		_ = lb
		if (_ck2(cc) == "[") {
			sdef := reg_find_struct(cc.Registry, name)
			_ = sdef
			if ((sdef.Name != "") && struct_is_generic(sdef)) {
				result := _check_generic_struct_lit(cc, name, sdef)
				_ = result
				return result
			}
		}
		if (_ck2(cc) == lb) {
			tid := resolve_type_name(cc.Store, name)
			_ = tid
			if (tid != TY_UNKNOWN) {
				sr := _check_struct_lit(cc, name, tid)
				_ = sr
				return sr
			}
		}
		if (name == "None") {
			has_user_none := false
			_ = has_user_none
			ui := int64(1)
			_ = ui
			for (ui < int64(len(cc.Registry.Struct_defs))) {
				udef := cc.Registry.Struct_defs[ui]
				_ = udef
				if udef.Is_sum {
					uvi := int64(1)
					_ = uvi
					for (uvi < int64(len(udef.Variant_names))) {
						if (udef.Variant_names[uvi] == "None") {
							has_user_none = true
						}
						uvi = (uvi + int64(1))
					}
				}
				ui = (ui + int64(1))
			}
			if (has_user_none == false) {
				new_st := store_add(cc.Store, mk_optional_type(TY_UNKNOWN))
				_ = new_st
				cc = _cset_store(cc, new_st)
				return CE{C: cc, Type_id: store_last_id(cc.Store)}
			}
		}
		env_type := _env_lookup(cc, name)
		_ = env_type
		if (env_type != TY_UNKNOWN) {
			return CE{C: cc, Type_id: env_type}
		}
		sig := reg_find_fn(cc.Registry, name)
		_ = sig
		if (sig.Name != "") {
			return CE{C: cc, Type_id: sig.Return_type}
		}
		if (name == "true") {
			return CE{C: cc, Type_id: TY_BOOL}
		}
		if (name == "false") {
			return CE{C: cc, Type_id: TY_BOOL}
		}
		return CE{C: cc, Type_id: TY_UNKNOWN}
	}
	if (k == "(") {
		cc := _cadv(c)
		_ = cc
		cc = _cskip_nl(cc)
		result := _check_expr_full(cc)
		_ = result
		cc = result.C
		cc = _cskip_nl(cc)
		if (_ck2(cc) == ",") {
			elem_ids := i2s(result.Type_id)
			_ = elem_ids
			arity := int64(1)
			_ = arity
			for ((_ck2(cc) == ",") && (_ck2(cc) != "EOF")) {
				cc = _cadv(cc)
				cc = _cskip_nl(cc)
				if (_ck2(cc) != ")") {
					er := _check_expr_full(cc)
					_ = er
					cc = er.C
					elem_ids = ((elem_ids + ",") + i2s(er.Type_id))
					arity = (arity + int64(1))
				}
				cc = _cskip_nl(cc)
			}
			if (_ck2(cc) == ")") {
				cc = _cadv(cc)
			}
			new_store := store_add(cc.Store, mk_tuple_type(elem_ids, arity))
			_ = new_store
			cc = _cset_store(cc, new_store)
			return CE{C: cc, Type_id: store_last_id(cc.Store)}
		}
		for ((_ck2(cc) != ")") && (_ck2(cc) != "EOF")) {
			cc = _cadv(cc)
		}
		if (_ck2(cc) == ")") {
			cc = _cadv(cc)
		}
		return CE{C: cc, Type_id: result.Type_id}
	}
	if (k == "[") {
		cc := _cadv(c)
		_ = cc
		cc = _cskip_nl(cc)
		elem_type := TY_UNKNOWN
		_ = elem_type
		count := int64(0)
		_ = count
		for ((_ck2(cc) != "]") && (_ck2(cc) != "EOF")) {
			cc = _cskip_nl(cc)
			result := _check_expr_full(cc)
			_ = result
			cc = result.C
			if (count == int64(0)) {
				elem_type = result.Type_id
			} else if ((elem_type != TY_UNKNOWN) && (result.Type_id != TY_UNKNOWN)) {
				if (is_assignable(cc.Store, result.Type_id, elem_type) == false) {
					q := _q()
					_ = q
					cc = _cerror(cc, E0100, ((((((("array element type mismatch: expected " + q) + format_type(cc.Store, elem_type)) + q) + ", got ") + q) + format_type(cc.Store, result.Type_id)) + q))
				}
			}
			count = (count + int64(1))
			cc = _cskip_nl(cc)
			if (_ck2(cc) == ",") {
				cc = _cadv(cc)
			}
		}
		if (_ck2(cc) == "]") {
			cc = _cadv(cc)
		}
		new_st := store_add(cc.Store, mk_array_type(elem_type))
		_ = new_st
		cc = _cset_store(cc, new_st)
		return CE{C: cc, Type_id: store_last_id(cc.Store)}
	}
	if (k == "fn") {
		cc := _cadv(c)
		_ = cc
		if ((_ck2(cc) == "IDENT") && (_ccur(cc).Text == "once")) {
			cc = _cadv(cc)
		}
		if (_ck2(cc) == "(") {
			cc = _cadv(cc)
			for ((_ck2(cc) != ")") && (_ck2(cc) != "EOF")) {
				cc = _cadv(cc)
			}
			if (_ck2(cc) == ")") {
				cc = _cadv(cc)
			}
		}
		if (_ck2(cc) == "->") {
			cc = _cadv(cc)
			cc = _skip_type_toks(cc)
		}
		if (_ck2(cc) == "=>") {
			cc = _cadv(cc)
			result := _check_expr_full(cc)
			_ = result
			cc = result.C
		}
		return CE{C: cc, Type_id: TY_UNKNOWN}
	}
	return CE{C: _cadv(c), Type_id: TY_UNKNOWN}
}

func _check_call_args(c Checker, callee_type int64) CE {
	cc := _cadv(c)
	_ = cc
	cc = _cskip_nl(cc)
	arg_count := int64(0)
	_ = arg_count
	for ((_ck2(cc) != ")") && (_ck2(cc) != "EOF")) {
		cc = _cskip_nl(cc)
		result := _check_expr_full(cc)
		_ = result
		cc = result.C
		arg_count = (arg_count + int64(1))
		cc = _cskip_nl(cc)
		if (_ck2(cc) == ",") {
			cc = _cadv(cc)
		}
	}
	if (_ck2(cc) == ")") {
		cc = _cadv(cc)
	}
	return CE{C: cc, Type_id: callee_type}
}

func _check_call_args_sig(c Checker, sig FnSig) CE {
	cc := _cadv(c)
	_ = cc
	cc = _cskip_nl(cc)
	arg_idx := int64(1)
	_ = arg_idx
	for ((_ck2(cc) != ")") && (_ck2(cc) != "EOF")) {
		cc = _cskip_nl(cc)
		result := _check_expr_full(cc)
		_ = result
		cc = result.C
		if (arg_idx < int64(len(sig.Param_types))) {
			expected := sig.Param_types[arg_idx]
			_ = expected
			if ((((expected != TY_UNKNOWN) && (expected > int64(0))) && (result.Type_id != TY_UNKNOWN)) && (result.Type_id > int64(0))) {
				exp_info := store_get(cc.Store, expected)
				_ = exp_info
				res_info := store_get(cc.Store, result.Type_id)
				_ = res_info
				exp_kn := type_kind_name(exp_info.Kind)
				_ = exp_kn
				res_kn := type_kind_name(res_info.Kind)
				_ = res_kn
				should_check := true
				_ = should_check
				if (exp_kn == "TypeVar") {
					should_check = false
				}
				if (res_kn == "TypeVar") {
					should_check = false
				}
				if (res_kn == "Unknown") {
					should_check = false
				}
				if (exp_kn == "Unknown") {
					should_check = false
				}
				if ((res_info.Type_id == TY_UNKNOWN) && ((((res_kn == "Optional") || (res_kn == "Result")) || (res_kn == "Array")))) {
					should_check = false
				}
				if ((exp_info.Type_id == TY_UNKNOWN) && ((((exp_kn == "Optional") || (exp_kn == "Result")) || (exp_kn == "Array")))) {
					should_check = false
				}
				if should_check {
					ok := is_assignable(cc.Store, result.Type_id, expected)
					_ = ok
					if ((ok == false) && (exp_kn == "TraitObject")) {
						res_name := _struct_name_for_tid(cc.Store, cc.Registry, result.Type_id)
						_ = res_name
						if ((res_name != "") && treg_has_impl(cc.Treg, res_name, exp_info.Name)) {
							ok = true
						}
					}
					if (ok == false) {
						q := _q()
						_ = q
						pname := ""
						_ = pname
						if (arg_idx < int64(len(sig.Param_names))) {
							pname = sig.Param_names[arg_idx]
						}
						cc = _cerror(cc, E0100, ((((((("argument " + q) + pname) + q) + " expects ") + format_type(cc.Store, expected)) + ", got ") + format_type(cc.Store, result.Type_id)))
					}
				}
			}
		}
		arg_idx = (arg_idx + int64(1))
		cc = _cskip_nl(cc)
		if (_ck2(cc) == ",") {
			cc = _cadv(cc)
		}
	}
	if (_ck2(cc) == ")") {
		cc = _cadv(cc)
	}
	return CE{C: cc, Type_id: sig.Return_type}
}

func _parse_single_bound(b string) []string {
	blen := int64(len(b))
	_ = blen
	if (blen == int64(0)) {
		return []string{"", "", ""}
	}
	lt := (int64(0) - int64(1))
	_ = lt
	i := int64(0)
	_ = i
	for (i < blen) {
		if (string(b[i]) == "<") {
			lt = i
			i = blen
		} else {
			i = (i + int64(1))
		}
	}
	if (lt < int64(0)) {
		return []string{b, "", ""}
	}
	trait_n := b[int64(0):lt]
	_ = trait_n
	gt := (int64(0) - int64(1))
	_ = gt
	j := (lt + int64(1))
	_ = j
	for (j < blen) {
		if (string(b[j]) == ">") {
			gt = j
			j = blen
		} else {
			j = (j + int64(1))
		}
	}
	if (gt < int64(0)) {
		return []string{trait_n, "", ""}
	}
	inner := b[(lt + int64(1)):gt]
	_ = inner
	eq := (int64(0) - int64(1))
	_ = eq
	k := int64(0)
	_ = k
	for (k < int64(len(inner))) {
		if (string(inner[k]) == "=") {
			eq = k
			k = int64(len(inner))
		} else {
			k = (k + int64(1))
		}
	}
	if (eq < int64(0)) {
		return []string{trait_n, "", ""}
	}
	assoc_n := inner[int64(0):eq]
	_ = assoc_n
	assoc_v := inner[(eq + int64(1)):int64(len(inner))]
	_ = assoc_v
	return []string{trait_n, assoc_n, assoc_v}
}

func _check_multi_bound_sig(c Checker, type_name string, bound string, param_name string, sig FnSig) Checker {
	cc := c
	_ = cc
	start := int64(0)
	_ = start
	i := int64(0)
	_ = i
	for (i <= int64(len(bound))) {
		at_sep := false
		_ = at_sep
		if (i == int64(len(bound))) {
			at_sep = true
		}
		if ((i < int64(len(bound))) && (string(bound[i]) == "+")) {
			at_sep = true
		}
		if (at_sep && (i > start)) {
			single := bound[start:i]
			_ = single
			parts := _parse_single_bound(single)
			_ = parts
			trait_n := parts[int64(0)]
			_ = trait_n
			assoc_n := parts[int64(1)]
			_ = assoc_n
			assoc_v := parts[int64(2)]
			_ = assoc_v
			if (treg_has_impl(cc.Treg, type_name, trait_n) == false) {
				cc = _cerror(cc, E0200, ((((("type " + type_name) + " does not satisfy bound ") + trait_n) + " for ") + param_name))
			} else if (assoc_n != "") {
				is_gp := (((assoc_v == sig.Generic_name_0) || (assoc_v == sig.Generic_name_1)) || (assoc_v == sig.Generic_name_2))
				_ = is_gp
				if ((is_gp == false) && _is_concrete_type_name(cc.Store, assoc_v)) {
					alias_name := ((type_name + "_") + assoc_n)
					_ = alias_name
					alias_tid := resolve_type_name(cc.Store, alias_name)
					_ = alias_tid
					expected_tid := resolve_type_name(cc.Store, assoc_v)
					_ = expected_tid
					if ((alias_tid != TY_UNKNOWN) && (expected_tid != TY_UNKNOWN)) {
						alias_info := store_get(cc.Store, alias_tid)
						_ = alias_info
						actual_tid := alias_info.Type_id
						_ = actual_tid
						if (actual_tid != expected_tid) {
							cc = _cerror(cc, E0200, ((((((((((("type " + type_name) + " has ") + trait_n) + "::") + assoc_n) + " = ") + format_type(cc.Store, actual_tid)) + " but expected ") + assoc_v) + " for ") + param_name))
						}
					}
				}
			}
			start = (i + int64(1))
		}
		i = (i + int64(1))
	}
	return cc
}

func _check_multi_bound(c Checker, type_name string, bound string, param_name string) Checker {
	cc := c
	_ = cc
	start := int64(0)
	_ = start
	i := int64(0)
	_ = i
	for (i <= int64(len(bound))) {
		at_sep := false
		_ = at_sep
		if (i == int64(len(bound))) {
			at_sep = true
		}
		if ((i < int64(len(bound))) && (string(bound[i]) == "+")) {
			at_sep = true
		}
		if (at_sep && (i > start)) {
			single := bound[start:i]
			_ = single
			parts := _parse_single_bound(single)
			_ = parts
			trait_n := parts[int64(0)]
			_ = trait_n
			assoc_n := parts[int64(1)]
			_ = assoc_n
			assoc_v := parts[int64(2)]
			_ = assoc_v
			if (treg_has_impl(cc.Treg, type_name, trait_n) == false) {
				cc = _cerror(cc, E0200, ((((("type " + type_name) + " does not satisfy bound ") + trait_n) + " for ") + param_name))
			} else if (assoc_n != "") {
				if _is_concrete_type_name(cc.Store, assoc_v) {
					alias_name := ((type_name + "_") + assoc_n)
					_ = alias_name
					alias_tid := resolve_type_name(cc.Store, alias_name)
					_ = alias_tid
					expected_tid := resolve_type_name(cc.Store, assoc_v)
					_ = expected_tid
					if ((alias_tid != TY_UNKNOWN) && (expected_tid != TY_UNKNOWN)) {
						alias_info := store_get(cc.Store, alias_tid)
						_ = alias_info
						actual_tid := alias_info.Type_id
						_ = actual_tid
						if (actual_tid != expected_tid) {
							cc = _cerror(cc, E0200, ((((((((((("type " + type_name) + " has ") + trait_n) + "::") + assoc_n) + " = ") + format_type(cc.Store, actual_tid)) + " but expected ") + assoc_v) + " for ") + param_name))
						}
					}
				}
			}
			start = (i + int64(1))
		}
		i = (i + int64(1))
	}
	return cc
}

func _is_concrete_type_name(store TypeStore, name string) bool {
	if (name == "") {
		return false
	}
	if (((((name == "i64") || (name == "f64")) || (name == "bool")) || (name == "str")) || (name == "char")) {
		return true
	}
	tid := resolve_type_name(store, name)
	_ = tid
	if (tid == TY_UNKNOWN) {
		return false
	}
	return true
}

func _derive_from_assoc_bound(c Checker, sig FnSig, target_name string, c0 int64, c1 int64, c2 int64, t0 int64, t1 int64, t2 int64) int64 {
	result := int64(0)
	_ = result
	if ((sig.Generic_bound_0 != "") && (c0 > int64(0))) {
		result = _scan_bound_for_target(c, sig.Generic_bound_0, target_name, c0)
		if (result > int64(0)) {
			return result
		}
	}
	if ((sig.Generic_bound_1 != "") && (c1 > int64(0))) {
		result = _scan_bound_for_target(c, sig.Generic_bound_1, target_name, c1)
		if (result > int64(0)) {
			return result
		}
	}
	if ((sig.Generic_bound_2 != "") && (c2 > int64(0))) {
		result = _scan_bound_for_target(c, sig.Generic_bound_2, target_name, c2)
		if (result > int64(0)) {
			return result
		}
	}
	return int64(0)
}

func _scan_bound_for_target(c Checker, bound string, target_name string, known_concrete int64) int64 {
	type_name := format_type(c.Store, known_concrete)
	_ = type_name
	start := int64(0)
	_ = start
	i := int64(0)
	_ = i
	for (i <= int64(len(bound))) {
		at_sep := false
		_ = at_sep
		if (i == int64(len(bound))) {
			at_sep = true
		}
		if ((i < int64(len(bound))) && (string(bound[i]) == "+")) {
			at_sep = true
		}
		if (at_sep && (i > start)) {
			single := bound[start:i]
			_ = single
			parts := _parse_single_bound(single)
			_ = parts
			trait_n := parts[int64(0)]
			_ = trait_n
			assoc_n := parts[int64(1)]
			_ = assoc_n
			assoc_v := parts[int64(2)]
			_ = assoc_v
			if ((assoc_n != "") && (assoc_v == target_name)) {
				alias_name := ((type_name + "_") + assoc_n)
				_ = alias_name
				alias_tid := resolve_type_name(c.Store, alias_name)
				_ = alias_tid
				if (alias_tid != TY_UNKNOWN) {
					alias_info := store_get(c.Store, alias_tid)
					_ = alias_info
					return alias_info.Type_id
				}
			}
			start = (i + int64(1))
		}
		i = (i + int64(1))
	}
	return int64(0)
}

func _check_generic_call(c Checker, sig FnSig) CE {
	cc := _cadv(c)
	_ = cc
	cc = _cskip_nl(cc)
	arg_type_0 := int64(0)
	_ = arg_type_0
	arg_type_1 := int64(0)
	_ = arg_type_1
	arg_type_2 := int64(0)
	_ = arg_type_2
	arg_type_3 := int64(0)
	_ = arg_type_3
	arg_type_4 := int64(0)
	_ = arg_type_4
	arg_type_5 := int64(0)
	_ = arg_type_5
	arg_type_6 := int64(0)
	_ = arg_type_6
	arg_type_7 := int64(0)
	_ = arg_type_7
	arg_count := int64(0)
	_ = arg_count
	for ((_ck2(cc) != ")") && (_ck2(cc) != "EOF")) {
		cc = _cskip_nl(cc)
		result := _check_expr_full(cc)
		_ = result
		cc = result.C
		if (arg_count == int64(0)) {
			arg_type_0 = result.Type_id
		}
		if (arg_count == int64(1)) {
			arg_type_1 = result.Type_id
		}
		if (arg_count == int64(2)) {
			arg_type_2 = result.Type_id
		}
		if (arg_count == int64(3)) {
			arg_type_3 = result.Type_id
		}
		if (arg_count == int64(4)) {
			arg_type_4 = result.Type_id
		}
		if (arg_count == int64(5)) {
			arg_type_5 = result.Type_id
		}
		if (arg_count == int64(6)) {
			arg_type_6 = result.Type_id
		}
		if (arg_count == int64(7)) {
			arg_type_7 = result.Type_id
		}
		arg_count = (arg_count + int64(1))
		cc = _cskip_nl(cc)
		if (_ck2(cc) == ",") {
			cc = _cadv(cc)
		}
	}
	if (_ck2(cc) == ")") {
		cc = _cadv(cc)
	}
	tv_id_0 := int64(0)
	_ = tv_id_0
	tv_id_1 := int64(0)
	_ = tv_id_1
	tv_id_2 := int64(0)
	_ = tv_id_2
	concrete_0 := int64(0)
	_ = concrete_0
	concrete_1 := int64(0)
	_ = concrete_1
	concrete_2 := int64(0)
	_ = concrete_2
	if (sig.Generic_count > int64(0)) {
		tv_id_0 = resolve_type_name(cc.Store, sig.Generic_name_0)
	}
	if (sig.Generic_count > int64(1)) {
		tv_id_1 = resolve_type_name(cc.Store, sig.Generic_name_1)
	}
	if (sig.Generic_count > int64(2)) {
		tv_id_2 = resolve_type_name(cc.Store, sig.Generic_name_2)
	}
	pi := int64(1)
	_ = pi
	for (pi < int64(len(sig.Param_types))) {
		pt := sig.Param_types[pi]
		_ = pt
		at := int64(0)
		_ = at
		ai := (pi - int64(1))
		_ = ai
		if (ai == int64(0)) {
			at = arg_type_0
		}
		if (ai == int64(1)) {
			at = arg_type_1
		}
		if (ai == int64(2)) {
			at = arg_type_2
		}
		if (ai == int64(3)) {
			at = arg_type_3
		}
		if (ai == int64(4)) {
			at = arg_type_4
		}
		if (ai == int64(5)) {
			at = arg_type_5
		}
		if (ai == int64(6)) {
			at = arg_type_6
		}
		if (ai == int64(7)) {
			at = arg_type_7
		}
		if ((at > int64(0)) && (at != TY_UNKNOWN)) {
			if ((pt == tv_id_0) && (tv_id_0 > int64(0))) {
				concrete_0 = at
			}
			if ((pt == tv_id_1) && (tv_id_1 > int64(0))) {
				concrete_1 = at
			}
			if ((pt == tv_id_2) && (tv_id_2 > int64(0))) {
				concrete_2 = at
			}
			if ((((pt > int64(0)) && (pt < int64(len(cc.Store.Types)))) && (at > int64(0))) && (at < int64(len(cc.Store.Types)))) {
				pti := store_get(cc.Store, pt)
				_ = pti
				ati := store_get(cc.Store, at)
				_ = ati
				if ((type_kind_name(pti.Kind) == "Function") && (type_kind_name(ati.Kind) == "Function")) {
					pret := pti.Type_id
					_ = pret
					aret := ati.Type_id
					_ = aret
					if ((((pret == tv_id_0) && (tv_id_0 > int64(0))) && (aret > int64(0))) && (aret != TY_UNKNOWN)) {
						concrete_0 = aret
					}
					if ((((pret == tv_id_1) && (tv_id_1 > int64(0))) && (aret > int64(0))) && (aret != TY_UNKNOWN)) {
						concrete_1 = aret
					}
					if ((((pret == tv_id_2) && (tv_id_2 > int64(0))) && (aret > int64(0))) && (aret != TY_UNKNOWN)) {
						concrete_2 = aret
					}
					fpi := int64(0)
					_ = fpi
					for ((fpi < pti.Param_count) && (fpi < ati.Param_count)) {
						fpt := cc.Store.Types[(pti.Param_start + fpi)].Type_id
						_ = fpt
						fat := cc.Store.Types[(ati.Param_start + fpi)].Type_id
						_ = fat
						if ((fat > int64(0)) && (fat != TY_UNKNOWN)) {
							if ((fpt == tv_id_0) && (tv_id_0 > int64(0))) {
								concrete_0 = fat
							}
							if ((fpt == tv_id_1) && (tv_id_1 > int64(0))) {
								concrete_1 = fat
							}
							if ((fpt == tv_id_2) && (tv_id_2 > int64(0))) {
								concrete_2 = fat
							}
						}
						fpi = (fpi + int64(1))
					}
				}
			}
			if ((pt > int64(0)) && (pt < int64(len(cc.Store.Types)))) {
				pti := store_get(cc.Store, pt)
				_ = pti
				if (type_kind_name(pti.Kind) == "Applied") {
					constructor_tv := pti.Type_id
					_ = constructor_tv
					arg_tv := pti.Type_id2
					_ = arg_tv
					at_info := store_get(cc.Store, at)
					_ = at_info
					at_kind := type_kind_name(at_info.Kind)
					_ = at_kind
					if (at_kind == "Array") {
						arr_ctor := resolve_type_name(cc.Store, "_Array")
						_ = arr_ctor
						if (arr_ctor == TY_UNKNOWN) {
							new_st := store_add(cc.Store, TypeInfo{Kind: TypeKindTyNamed{}, Name: "_Array", Type_id: int64(0), Type_id2: int64(0), Param_start: int64(0), Param_count: int64(0)})
							_ = new_st
							cc = _cset_store(cc, new_st)
							arr_ctor = store_last_id(cc.Store)
						}
						if ((constructor_tv == tv_id_0) && (tv_id_0 > int64(0))) {
							concrete_0 = arr_ctor
						}
						if ((constructor_tv == tv_id_1) && (tv_id_1 > int64(0))) {
							concrete_1 = arr_ctor
						}
						if ((constructor_tv == tv_id_2) && (tv_id_2 > int64(0))) {
							concrete_2 = arr_ctor
						}
						elem_tid := at_info.Type_id
						_ = elem_tid
						if ((arg_tv == tv_id_0) && (tv_id_0 > int64(0))) {
							concrete_0 = elem_tid
						}
						if ((arg_tv == tv_id_1) && (tv_id_1 > int64(0))) {
							concrete_1 = elem_tid
						}
						if ((arg_tv == tv_id_2) && (tv_id_2 > int64(0))) {
							concrete_2 = elem_tid
						}
					}
				}
			}
		}
		pi = (pi + int64(1))
	}
	if ((concrete_0 == int64(0)) && (sig.Generic_name_0 != "")) {
		concrete_0 = _derive_from_assoc_bound(cc, sig, sig.Generic_name_0, concrete_0, concrete_1, concrete_2, tv_id_0, tv_id_1, tv_id_2)
	}
	if ((concrete_1 == int64(0)) && (sig.Generic_name_1 != "")) {
		concrete_1 = _derive_from_assoc_bound(cc, sig, sig.Generic_name_1, concrete_0, concrete_1, concrete_2, tv_id_0, tv_id_1, tv_id_2)
	}
	if ((concrete_2 == int64(0)) && (sig.Generic_name_2 != "")) {
		concrete_2 = _derive_from_assoc_bound(cc, sig, sig.Generic_name_2, concrete_0, concrete_1, concrete_2, tv_id_0, tv_id_1, tv_id_2)
	}
	spec_name := sig.Name
	_ = spec_name
	if ((sig.Generic_count > int64(0)) && (concrete_0 > int64(0))) {
		spec_name = ((spec_name + "_") + _type_name_suffix(cc.Store, concrete_0))
	}
	if ((sig.Generic_count > int64(1)) && (concrete_1 > int64(0))) {
		spec_name = ((spec_name + "_") + _type_name_suffix(cc.Store, concrete_1))
	}
	if ((sig.Generic_count > int64(2)) && (concrete_2 > int64(0))) {
		spec_name = ((spec_name + "_") + _type_name_suffix(cc.Store, concrete_2))
	}
	existing := reg_find_mono(cc.Registry, spec_name)
	_ = existing
	if (existing.Specialized_name == "") {
		sp_ptypes := []int64{int64(0)}
		_ = sp_ptypes
		pi = int64(1)
		for (pi < int64(len(sig.Param_types))) {
			pt := sig.Param_types[pi]
			_ = pt
			if ((pt == tv_id_0) && (tv_id_0 > int64(0))) {
				pt = concrete_0
			}
			if ((pt == tv_id_1) && (tv_id_1 > int64(0))) {
				pt = concrete_1
			}
			if ((pt == tv_id_2) && (tv_id_2 > int64(0))) {
				pt = concrete_2
			}
			sp_ptypes = append(sp_ptypes, pt)
			pi = (pi + int64(1))
		}
		sp_ret := sig.Return_type
		_ = sp_ret
		if ((sp_ret == tv_id_0) && (tv_id_0 > int64(0))) {
			sp_ret = concrete_0
		}
		if ((sp_ret == tv_id_1) && (tv_id_1 > int64(0))) {
			sp_ret = concrete_1
		}
		if ((sp_ret == tv_id_2) && (tv_id_2 > int64(0))) {
			sp_ret = concrete_2
		}
		sp_sig := new_fn_sig(spec_name, sig.Param_names, sp_ptypes, sp_ret, sig.Error_type)
		_ = sp_sig
		cc = _cset_reg(cc, reg_add_fn(cc.Registry, sp_sig))
		spec := MonoSpec{Generic_name: sig.Name, Specialized_name: spec_name, Type_arg_0: concrete_0, Type_arg_1: concrete_1, Type_arg_2: concrete_2}
		_ = spec
		cc = _cset_reg(cc, reg_add_mono(cc.Registry, spec))
	}
	if (((sig.Generic_bound_0 != "") && (concrete_0 > int64(0))) && (concrete_0 != TY_UNKNOWN)) {
		cname := format_type(cc.Store, concrete_0)
		_ = cname
		cc = _check_multi_bound_sig(cc, cname, sig.Generic_bound_0, sig.Generic_name_0, sig)
	}
	if (((sig.Generic_bound_1 != "") && (concrete_1 > int64(0))) && (concrete_1 != TY_UNKNOWN)) {
		cname := format_type(cc.Store, concrete_1)
		_ = cname
		cc = _check_multi_bound_sig(cc, cname, sig.Generic_bound_1, sig.Generic_name_1, sig)
	}
	if (((sig.Generic_bound_2 != "") && (concrete_2 > int64(0))) && (concrete_2 != TY_UNKNOWN)) {
		cname := format_type(cc.Store, concrete_2)
		_ = cname
		cc = _check_multi_bound_sig(cc, cname, sig.Generic_bound_2, sig.Generic_name_2, sig)
	}
	ret_type := sig.Return_type
	_ = ret_type
	if ((ret_type == tv_id_0) && (tv_id_0 > int64(0))) {
		ret_type = concrete_0
	}
	if ((ret_type == tv_id_1) && (tv_id_1 > int64(0))) {
		ret_type = concrete_1
	}
	if ((ret_type == tv_id_2) && (tv_id_2 > int64(0))) {
		ret_type = concrete_2
	}
	if (ret_type == int64(0)) {
		ret_type = TY_UNKNOWN
	}
	return CE{C: cc, Type_id: ret_type}
}

func _check_method_call(c Checker, obj_type int64, method string) CE {
	cc := _cadv(c)
	_ = cc
	cc = _cskip_nl(cc)
	for ((_ck2(cc) != ")") && (_ck2(cc) != "EOF")) {
		cc = _cskip_nl(cc)
		result := _check_expr_full(cc)
		_ = result
		cc = result.C
		cc = _cskip_nl(cc)
		if (_ck2(cc) == ",") {
			cc = _cadv(cc)
		}
	}
	if (_ck2(cc) == ")") {
		cc = _cadv(cc)
	}
	if ((method == "toStr") || (method == "toString")) {
		return CE{C: cc, Type_id: TY_STR}
	}
	if (method == "len") {
		return CE{C: cc, Type_id: TY_I64}
	}
	if (method == "contains") {
		return CE{C: cc, Type_id: TY_BOOL}
	}
	if ((method == "starts_with") || (method == "startsWith")) {
		return CE{C: cc, Type_id: TY_BOOL}
	}
	if ((method == "ends_with") || (method == "endsWith")) {
		return CE{C: cc, Type_id: TY_BOOL}
	}
	if (method == "indexOf") {
		return CE{C: cc, Type_id: TY_I64}
	}
	if (method == "is_empty") {
		return CE{C: cc, Type_id: TY_BOOL}
	}
	if ((((((((method == "trim") || (method == "to_upper")) || (method == "to_lower")) || (method == "replace")) || (method == "substring")) || (method == "charAt")) || (method == "toUpper")) || (method == "toLower")) {
		return CE{C: cc, Type_id: TY_STR}
	}
	if (method == "split") {
		new_st := store_add(cc.Store, mk_array_type(TY_STR))
		_ = new_st
		cc = _cset_store(cc, new_st)
		return CE{C: cc, Type_id: store_last_id(cc.Store)}
	}
	if (method == "append") {
		return CE{C: cc, Type_id: obj_type}
	}
	return CE{C: cc, Type_id: TY_UNKNOWN}
}

func _check_field(c Checker, obj_type int64, field string) CE {
	info := store_get(c.Store, obj_type)
	_ = info
	if (type_kind_name(info.Kind) == "Named") {
		fi := reg_find_field(c.Registry, info.Name, field)
		_ = fi
		if (fi.Name != "") {
			return CE{C: c, Type_id: fi.Type_id}
		}
		qualified := ((info.Name + "_") + field)
		_ = qualified
		msig := reg_find_fn(c.Registry, qualified)
		_ = msig
		if (msig.Name != "") {
			cc := c
			_ = cc
			new_st := store_add(cc.Store, mk_fn_type(msig.Return_type, int64(0), int64(0)))
			_ = new_st
			cc = _cset_store(cc, new_st)
			return CE{C: cc, Type_id: store_last_id(cc.Store)}
		}
		if (obj_type != TY_UNKNOWN) {
			q := _q()
			_ = q
			cc := _cerror(c, E0107, ((((((("field " + q) + field) + q) + " not found on type ") + q) + info.Name) + q))
			_ = cc
			return CE{C: cc, Type_id: TY_UNKNOWN}
		}
	}
	return CE{C: c, Type_id: TY_UNKNOWN}
}

func _check_struct_lit(c Checker, name string, type_id int64) CE {
	lb := "{"
	_ = lb
	rb := "}"
	_ = rb
	cc := _cadv(c)
	_ = cc
	cc = _cskip_nl(cc)
	for ((_ck2(cc) != rb) && (_ck2(cc) != "EOF")) {
		if (_ck2(cc) == "IDENT") {
			field_name := _ccur(cc).Text
			_ = field_name
			cc = _cadv(cc)
			if (_ck2(cc) == ":") {
				cc = _cadv(cc)
				result := _check_expr_full(cc)
				_ = result
				cc = result.C
				fi := reg_find_field(cc.Registry, name, field_name)
				_ = fi
				if (((fi.Name != "") && (result.Type_id != TY_UNKNOWN)) && (fi.Type_id != TY_UNKNOWN)) {
					if (types_equal(cc.Store, result.Type_id, fi.Type_id) == false) {
						q := _q()
						_ = q
						cc = _cerror(cc, E0100, ((((((("field " + q) + field_name) + q) + " expects ") + format_type(cc.Store, fi.Type_id)) + ", got ") + format_type(cc.Store, result.Type_id)))
					}
				}
			}
		} else {
			cc = _cadv(cc)
		}
		cc = _cskip_nl(cc)
		if (_ck2(cc) == ",") {
			cc = _cadv(cc)
			cc = _cskip_nl(cc)
		}
	}
	if (_ck2(cc) == rb) {
		cc = _cadv(cc)
	}
	return CE{C: cc, Type_id: type_id}
}

func _check_generic_struct_lit(c Checker, name string, sdef StructDef) CE {
	cc := _cadv(c)
	_ = cc
	ta_0 := int64(0)
	_ = ta_0
	ta_1 := int64(0)
	_ = ta_1
	ta_2 := int64(0)
	_ = ta_2
	ta_count := int64(0)
	_ = ta_count
	for ((_ck2(cc) != "]") && (_ck2(cc) != "EOF")) {
		ta := _parse_type_ann(cc)
		_ = ta
		cc = ta.C
		if (ta_count == int64(0)) {
			ta_0 = ta.Type_id
		}
		if (ta_count == int64(1)) {
			ta_1 = ta.Type_id
		}
		if (ta_count == int64(2)) {
			ta_2 = ta.Type_id
		}
		ta_count = (ta_count + int64(1))
		if (_ck2(cc) == ",") {
			cc = _cadv(cc)
		}
	}
	if (_ck2(cc) == "]") {
		cc = _cadv(cc)
	}
	spec_name := name
	_ = spec_name
	if (ta_count > int64(0)) {
		spec_name = ((spec_name + "_") + _type_name_suffix(cc.Store, ta_0))
	}
	if (ta_count > int64(1)) {
		spec_name = ((spec_name + "_") + _type_name_suffix(cc.Store, ta_1))
	}
	if (ta_count > int64(2)) {
		spec_name = ((spec_name + "_") + _type_name_suffix(cc.Store, ta_2))
	}
	existing := reg_find_struct(cc.Registry, spec_name)
	_ = existing
	if (existing.Name == "") {
		tv_0 := int64(0)
		_ = tv_0
		tv_1 := int64(0)
		_ = tv_1
		tv_2 := int64(0)
		_ = tv_2
		if (sdef.Generic_count > int64(0)) {
			tv_0 = resolve_type_name(cc.Store, sdef.Generic_name_0)
		}
		if (sdef.Generic_count > int64(1)) {
			tv_1 = resolve_type_name(cc.Store, sdef.Generic_name_1)
		}
		if (sdef.Generic_count > int64(2)) {
			tv_2 = resolve_type_name(cc.Store, sdef.Generic_name_2)
		}
		sp_fields := []FieldInfo{FieldInfo{Name: "", Type_id: int64(0)}}
		_ = sp_fields
		fi := int64(1)
		_ = fi
		for (fi < int64(len(sdef.Fields))) {
			fld := sdef.Fields[fi]
			_ = fld
			ft := fld.Type_id
			_ = ft
			if ((ft == tv_0) && (tv_0 > int64(0))) {
				ft = ta_0
			}
			if ((ft == tv_1) && (tv_1 > int64(0))) {
				ft = ta_1
			}
			if ((ft == tv_2) && (tv_2 > int64(0))) {
				ft = ta_2
			}
			sp_fields = append(sp_fields, FieldInfo{Name: fld.Name, Type_id: ft})
			fi = (fi + int64(1))
		}
		sp_def := new_struct_def(spec_name, sp_fields, false, []string{""})
		_ = sp_def
		cc = _cset_reg(cc, reg_add_struct(cc.Registry, sp_def))
		new_st := store_add(cc.Store, mk_named_type(spec_name))
		_ = new_st
		cc = _cset_store(cc, new_st)
	}
	spec_tid := resolve_type_name(cc.Store, spec_name)
	_ = spec_tid
	lb := "{"
	_ = lb
	if (_ck2(cc) == lb) {
		sr := _check_struct_lit(cc, spec_name, spec_tid)
		_ = sr
		return sr
	}
	return CE{C: cc, Type_id: spec_tid}
}

func _check_if(c Checker) CE {
	cc := _cadv(c)
	_ = cc
	lb := "{"
	_ = lb
	cond_type := TY_UNKNOWN
	_ = cond_type
	for ((_ck2(cc) != lb) && (_ck2(cc) != "EOF")) {
		result := _check_primary(cc)
		_ = result
		cc = result.C
		cond_type = result.Type_id
		k := _ck2(cc)
		_ = k
		if ((((((((k == "==") || (k == "!=")) || (k == "<")) || (k == ">")) || (k == "<=")) || (k == ">=")) || (k == "&&")) || (k == "||")) {
			cc = _cadv(cc)
			right := _check_primary(cc)
			_ = right
			cc = right.C
			cond_type = TY_BOOL
		} else if ((_is_cont(k) == false) || (k == lb)) {
		} else {
			cc = _cadv(cc)
		}
	}
	then_result := _check_block(cc)
	_ = then_result
	cc = then_result.C
	result_type := then_result.Type_id
	_ = result_type
	if (_ck2(cc) == "else") {
		cc = _cadv(cc)
		else_type := TY_UNKNOWN
		_ = else_type
		if (_ck2(cc) == "if") {
			else_result := _check_if(cc)
			_ = else_result
			cc = else_result.C
			else_type = else_result.Type_id
		} else {
			else_result := _check_block(cc)
			_ = else_result
			cc = else_result.C
			else_type = else_result.Type_id
		}
		if ((((((result_type != TY_UNKNOWN) && (else_type != TY_UNKNOWN)) && (result_type != TY_UNIT)) && (else_type != TY_UNIT)) && (result_type != TY_NEVER)) && (else_type != TY_NEVER)) {
			if ((is_assignable(cc.Store, else_type, result_type) == false) && (is_assignable(cc.Store, result_type, else_type) == false)) {
				q := _q()
				_ = q
				cc = _cerror(cc, E0100, ((((((("if/else branch type mismatch: then returns " + q) + format_type(cc.Store, result_type)) + q) + ", else returns ") + q) + format_type(cc.Store, else_type)) + q))
			}
		}
		if (result_type == TY_NEVER) {
			result_type = else_type
		}
	}
	return CE{C: cc, Type_id: result_type}
}

func _check_match(c Checker) CE {
	cc := _cadv(c)
	_ = cc
	lb := "{"
	_ = lb
	rb := "}"
	_ = rb
	subj_type := TY_UNKNOWN
	_ = subj_type
	for ((_ck2(cc) != lb) && (_ck2(cc) != "EOF")) {
		result := _check_primary(cc)
		_ = result
		cc = result.C
		subj_type = result.Type_id
		if (_is_cont(_ck2(cc)) && (_ck2(cc) != lb)) {
			cc = _cadv(cc)
		}
	}
	cc = _cskip_nl(cc)
	if (_ck2(cc) != lb) {
		return CE{C: cc, Type_id: TY_UNKNOWN}
	}
	cc = _cadv(cc)
	cc = _cskip_nl(cc)
	has_wildcard := false
	_ = has_wildcard
	covered := []string{""}
	_ = covered
	has_true := false
	_ = has_true
	has_false := false
	_ = has_false
	arm_type := TY_UNKNOWN
	_ = arm_type
	arm_count := int64(0)
	_ = arm_count
	for ((_ck2(cc) != rb) && (_ck2(cc) != "EOF")) {
		arm_is_guarded := false
		_ = arm_is_guarded
		if (_ck2(cc) == "IDENT") {
			pat_name := _ccur(cc).Text
			_ = pat_name
			real_name := pat_name
			_ = real_name
			atp := (cc.Pos + int64(1))
			_ = atp
			if (atp < int64(len(cc.Tokens))) {
				if (token_name(cc.Tokens[atp].Kind) == "@") {
					atp2 := (atp + int64(1))
					_ = atp2
					if (atp2 < int64(len(cc.Tokens))) {
						if (token_name(cc.Tokens[atp2].Kind) == "IDENT") {
							real_name = cc.Tokens[atp2].Text
						}
					}
				}
			}
			if (pat_name == "_") {
				scan := cc
				_ = scan
				found_guard := false
				_ = found_guard
				for ((_ck2(scan) != "=>") && (_ck2(scan) != "EOF")) {
					if (_ck2(scan) == "if") {
						found_guard = true
					}
					scan = _cadv(scan)
				}
				if (found_guard == false) {
					has_wildcard = true
				}
				arm_is_guarded = found_guard
			} else if (real_name == "true") {
				has_true = true
			} else if (real_name == "false") {
				has_false = true
			} else {
				is_variant_name := false
				_ = is_variant_name
				si := int64(1)
				_ = si
				for (si < int64(len(cc.Registry.Struct_defs))) {
					sdef := cc.Registry.Struct_defs[si]
					_ = sdef
					if sdef.Is_sum {
						vvi := int64(1)
						_ = vvi
						for (vvi < int64(len(sdef.Variant_names))) {
							if (sdef.Variant_names[vvi] == real_name) {
								is_variant_name = true
							}
							vvi = (vvi + int64(1))
						}
					}
					si = (si + int64(1))
				}
				if (is_variant_name == false) {
					scan := cc
					_ = scan
					found_guard := false
					_ = found_guard
					for ((_ck2(scan) != "=>") && (_ck2(scan) != "EOF")) {
						if (_ck2(scan) == "if") {
							found_guard = true
						}
						scan = _cadv(scan)
					}
					if (found_guard == false) {
						has_wildcard = true
					}
					arm_is_guarded = found_guard
				}
				if (arm_is_guarded == false) {
					covered = append(covered, real_name)
				}
			}
		}
		for ((_ck2(cc) != "=>") && (_ck2(cc) != "EOF")) {
			if (_ck2(cc) == "|") {
				cc = _cadv(cc)
				if ((_ck2(cc) == "IDENT") && (arm_is_guarded == false)) {
					covered = append(covered, _ccur(cc).Text)
				}
			}
			cc = _cadv(cc)
		}
		if (_ck2(cc) == "=>") {
			cc = _cadv(cc)
		}
		cc = _cskip_nl(cc)
		if (_ck2(cc) == lb) {
			body := _check_block(cc)
			_ = body
			cc = body.C
			if (arm_count == int64(0)) {
				arm_type = body.Type_id
			}
		} else {
			body := _check_expr_full(cc)
			_ = body
			cc = body.C
			if (arm_count == int64(0)) {
				arm_type = body.Type_id
			}
		}
		arm_count = (arm_count + int64(1))
		cc = _cskip_nl(cc)
	}
	if (_ck2(cc) == rb) {
		cc = _cadv(cc)
	}
	if (has_wildcard == false) {
		info := store_get(cc.Store, subj_type)
		_ = info
		kind_name := type_kind_name(info.Kind)
		_ = kind_name
		if (kind_name == "Named") {
			iname := type_info_name(info)
			_ = iname
			def := reg_find_struct(cc.Registry, iname)
			_ = def
			if def.Is_sum {
				if (int64(len(def.Variant_names)) > int64(1)) {
					vi := int64(1)
					_ = vi
					first_missing := ""
					_ = first_missing
					for (vi < int64(len(def.Variant_names))) {
						vname := def.Variant_names[vi]
						_ = vname
						found := false
						_ = found
						ci := int64(1)
						_ = ci
						for (ci < int64(len(covered))) {
							if (covered[ci] == vname) {
								found = true
							}
							ci = (ci + int64(1))
						}
						if (found == false) {
							if (first_missing == "") {
								first_missing = vname
							}
						}
						vi = (vi + int64(1))
					}
					if (first_missing != "") {
						cc = _cerror(cc, E0400, (("non-exhaustive match: variant '" + first_missing) + "' not covered"))
					}
				}
			}
			if ((iname == "Optional") || (iname == "Option")) {
				has_some := false
				_ = has_some
				has_none := false
				_ = has_none
				oi := int64(1)
				_ = oi
				for (oi < int64(len(covered))) {
					if (covered[oi] == "Some") {
						has_some = true
					}
					if (covered[oi] == "None") {
						has_none = true
					}
					oi = (oi + int64(1))
				}
				if (((has_some == false) || (has_none == false))) {
					missing := "Some"
					_ = missing
					if has_some {
						missing = "None"
					}
					cc = _cerror(cc, E0400, (("non-exhaustive match: variant '" + missing) + "' not covered"))
				}
			}
			if (iname == "Result") {
				has_ok := false
				_ = has_ok
				has_err := false
				_ = has_err
				ri := int64(1)
				_ = ri
				for (ri < int64(len(covered))) {
					if (covered[ri] == "Ok") {
						has_ok = true
					}
					if (covered[ri] == "Err") {
						has_err = true
					}
					ri = (ri + int64(1))
				}
				if (((has_ok == false) || (has_err == false))) {
					missing := "Ok"
					_ = missing
					if has_ok {
						missing = "Err"
					}
					cc = _cerror(cc, E0400, (("non-exhaustive match: variant '" + missing) + "' not covered"))
				}
			}
		}
		if ((kind_name == "bool") || (subj_type == TY_BOOL)) {
			if (((has_true == false) || (has_false == false))) {
				missing := "true"
				_ = missing
				if has_true {
					missing = "false"
				}
				cc = _cerror(cc, E0400, (("non-exhaustive match: case '" + missing) + "' not covered"))
			}
		}
		if (kind_name == "Optional") {
			has_some := false
			_ = has_some
			has_none := false
			_ = has_none
			oi := int64(1)
			_ = oi
			for (oi < int64(len(covered))) {
				if (covered[oi] == "Some") {
					has_some = true
				}
				if (covered[oi] == "None") {
					has_none = true
				}
				oi = (oi + int64(1))
			}
			if (((has_some == false) || (has_none == false))) {
				missing := "Some"
				_ = missing
				if has_some {
					missing = "None"
				}
				cc = _cerror(cc, E0400, (("non-exhaustive match: variant '" + missing) + "' not covered"))
			}
		}
		if (kind_name == "Result") {
			has_ok := false
			_ = has_ok
			has_err := false
			_ = has_err
			ri := int64(1)
			_ = ri
			for (ri < int64(len(covered))) {
				if (covered[ri] == "Ok") {
					has_ok = true
				}
				if (covered[ri] == "Err") {
					has_err = true
				}
				ri = (ri + int64(1))
			}
			if (((has_ok == false) || (has_err == false))) {
				missing := "Ok"
				_ = missing
				if has_ok {
					missing = "Err"
				}
				cc = _cerror(cc, E0400, (("non-exhaustive match: variant '" + missing) + "' not covered"))
			}
		}
	}
	return CE{C: cc, Type_id: arm_type}
}

func _check_for(c Checker) CE {
	cc := _cadv(c)
	_ = cc
	lb := "{"
	_ = lb
	var_name := ""
	_ = var_name
	if (_ck2(cc) == "IDENT") {
		var_name = _ccur(cc).Text
		cc = _cadv(cc)
	}
	if (_ck2(cc) == "in") {
		cc = _cadv(cc)
	}
	iter_type := TY_UNKNOWN
	_ = iter_type
	for ((_ck2(cc) != lb) && (_ck2(cc) != "EOF")) {
		result := _check_expr_full(cc)
		_ = result
		cc = result.C
		iter_type = result.Type_id
	}
	elem_type := TY_UNKNOWN
	_ = elem_type
	info := store_get(cc.Store, iter_type)
	_ = info
	if (type_kind_name(info.Kind) == "Array") {
		elem_type = info.Type_id
	}
	if (var_name != "") {
		cc = _env_add(cc, var_name, elem_type, int64(0))
	}
	cc = _cset_loop(cc, true)
	body := _check_block(cc)
	_ = body
	cc = body.C
	cc = _cset_loop(cc, false)
	return CE{C: cc, Type_id: TY_UNIT}
}

func _check_while(c Checker) CE {
	cc := _cadv(c)
	_ = cc
	lb := "{"
	_ = lb
	for ((_ck2(cc) != lb) && (_ck2(cc) != "EOF")) {
		result := _check_primary(cc)
		_ = result
		cc = result.C
		if (_is_cont(_ck2(cc)) && (_ck2(cc) != lb)) {
			cc = _cadv(cc)
		}
	}
	cc = _cset_loop(cc, true)
	body := _check_block(cc)
	_ = body
	cc = body.C
	cc = _cset_loop(cc, false)
	return CE{C: cc, Type_id: TY_UNIT}
}

func _check_loop(c Checker) CE {
	cc := _cadv(c)
	_ = cc
	body := _check_block(cc)
	_ = body
	return CE{C: body.C, Type_id: TY_UNIT}
}

func _check_once_consumption(c Checker, start int64, end_tok int64) Checker {
	cc := c
	_ = cc
	i := start
	_ = i
	once_names := []string{""}
	_ = once_names
	once_positions := []int64{int64(0)}
	_ = once_positions
	for ((i < end_tok) && (i < int64(len(cc.Tokens)))) {
		tk := token_name(cc.Tokens[i].Kind)
		_ = tk
		if ((tk == "IDENT") && ((i + int64(2)) < end_tok)) {
			next := (i + int64(1))
			_ = next
			nextk := token_name(cc.Tokens[next].Kind)
			_ = nextk
			if (nextk == ":=") {
				rhs := (i + int64(2))
				_ = rhs
				if (((rhs + int64(1)) < end_tok) && (token_name(cc.Tokens[rhs].Kind) == "fn")) {
					rhs2 := (rhs + int64(1))
					_ = rhs2
					if ((token_name(cc.Tokens[rhs2].Kind) == "IDENT") && (cc.Tokens[rhs2].Text == "once")) {
						has_unchecked := false
						_ = has_unchecked
						back := (i - int64(1))
						_ = back
						for ((back >= start) && (token_name(cc.Tokens[back].Kind) == "NEWLINE")) {
							back = (back - int64(1))
						}
						if (((back >= (start + int64(1))) && (token_name(cc.Tokens[back].Kind) == "IDENT")) && (cc.Tokens[back].Text == "once_unchecked")) {
							if (token_name(cc.Tokens[(back - int64(1))].Kind) == "@") {
								has_unchecked = true
							}
						}
						if (has_unchecked == false) {
							once_names = append(once_names, cc.Tokens[i].Text)
							once_positions = append(once_positions, i)
						}
					}
				}
			}
		}
		i = (i + int64(1))
	}
	oi := int64(1)
	_ = oi
	for (oi < int64(len(once_names))) {
		oname := once_names[oi]
		_ = oname
		opos := once_positions[oi]
		_ = opos
		call_count := int64(0)
		_ = call_count
		first_call_line := int64(0)
		_ = first_call_line
		second_call_line := int64(0)
		_ = second_call_line
		j := (opos + int64(1))
		_ = j
		for ((j < end_tok) && ((j + int64(1)) < int64(len(cc.Tokens)))) {
			tk := token_name(cc.Tokens[j].Kind)
			_ = tk
			if ((tk == "IDENT") && (cc.Tokens[j].Text == oname)) {
				nextk := token_name(cc.Tokens[(j + int64(1))].Kind)
				_ = nextk
				if (nextk == "(") {
					call_count = (call_count + int64(1))
					if (call_count == int64(1)) {
						first_call_line = cc.Tokens[j].Line
					}
					if (call_count == int64(2)) {
						second_call_line = cc.Tokens[j].Line
					}
				}
			}
			j = (j + int64(1))
		}
		if (call_count > int64(1)) {
			q := _q()
			_ = q
			line_str := i2s(first_call_line)
			_ = line_str
			msg := (((((("once closure " + q) + oname) + q) + " called multiple times (first at line ") + line_str) + ")")
			_ = msg
			cc = _cwarning(cc, "W0008", msg)
		}
		oi = (oi + int64(1))
	}
	return cc
}

func _pass4_once_check(c Checker, index DeclIndex) Checker {
	cc := c
	_ = cc
	i := int64(1)
	_ = i
	for (i < decl_index_len(index)) {
		decl := decl_index_get(index, i)
		_ = decl
		kn := decl_kind_name(decl.Kind)
		_ = kn
		if (((kn == "fn") || (kn == "entry")) || (kn == "test")) {
			if ((decl.Body_start > int64(0)) && (decl.Body_end > decl.Body_start)) {
				cc = _check_once_consumption(cc, decl.Body_start, decl.Body_end)
			}
		}
		i = (i + int64(1))
	}
	return cc
}

func _scan_effect_decls(c Checker) Checker {
	cc := c
	_ = cc
	i := int64(0)
	_ = i
	for (i < (int64(len(cc.Tokens)) - int64(1))) {
		if ((token_name(cc.Tokens[i].Kind) == "IDENT") && (cc.Tokens[i].Text == "effect")) {
			next := (i + int64(1))
			_ = next
			if ((next < int64(len(cc.Tokens))) && (token_name(cc.Tokens[next].Kind) == "IDENT")) {
				eff_name := cc.Tokens[next].Text
				_ = eff_name
				parents := []string{""}
				_ = parents
				imp := (next + int64(1))
				_ = imp
				if (((imp < int64(len(cc.Tokens))) && (token_name(cc.Tokens[imp].Kind) == "IDENT")) && (cc.Tokens[imp].Text == "implies")) {
					pi := (imp + int64(1))
					_ = pi
					for ((pi < int64(len(cc.Tokens))) && (token_name(cc.Tokens[pi].Kind) == "IDENT")) {
						parents = append(parents, cc.Tokens[pi].Text)
						pi = (pi + int64(1))
						if ((pi < int64(len(cc.Tokens))) && (token_name(cc.Tokens[pi].Kind) == ",")) {
							pi = (pi + int64(1))
						}
					}
				}
				tdef := TraitDef{Name: eff_name, Method_names: []string{""}, Parent_traits: parents, Assoc_type_names: []string{""}, Default_starts: []int64{int64(0)}, Default_ends: []int64{int64(0)}}
				_ = tdef
				cc = _cset_treg(cc, treg_add_trait(cc.Treg, tdef))
			}
		}
		i = (i + int64(1))
	}
	return cc
}

func check(tokens []Token, index DeclIndex, pool NodePool, table SymbolTable, file string) CheckResult {
	c := _new_checker(tokens, pool, table, file)
	_ = c
	c = _scan_effect_decls(c)
	c = _pass1_types(c, index)
	c = _pass2_fns(c, index)
	c = _pass3_bodies(c, index)
	c = _pass4_once_check(c, index)
	return CheckResult{Store: c.Store, Registry: c.Registry, Treg: c.Treg, Diagnostics: c.Diagnostics}
}

