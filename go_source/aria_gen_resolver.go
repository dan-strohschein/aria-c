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

type Resolver struct {
	Tokens []Token
	Pos int64
	Table SymbolTable
	Diagnostics DiagnosticBag
	Current_scope int64
	File string
	Is_public bool
	Pool NodePool
	Index DeclIndex
}

type ResolveResult struct {
	Table SymbolTable
	Diagnostics DiagnosticBag
}

func _new_resolver(tokens []Token, pool NodePool, index DeclIndex, file string) Resolver {
	return Resolver{Tokens: tokens, Pos: int64(0), Table: new_table(), Diagnostics: new_bag(), Current_scope: int64(0), File: file, Is_public: false, Pool: pool, Index: index}
}

func _rcur(r Resolver) Token {
	if (r.Pos >= int64(len(r.Tokens))) {
		return Token{Kind: TokenKindTkEof{}, Text: "", Line: int64(0), Col: int64(0), Offset: int64(0)}
	}
	return r.Tokens[r.Pos]
}

func _rk(r Resolver) string {
	return token_name(_rcur(r).Kind)
}

func _rk2(r Resolver, pos int64) string {
	if (pos < int64(0)) {
		return ""
	}
	if (pos >= int64(len(r.Tokens))) {
		return ""
	}
	return token_name(r.Tokens[pos].Kind)
}

func _radv(r Resolver) Resolver {
	return Resolver{Tokens: r.Tokens, Pos: (r.Pos + int64(1)), Table: r.Table, Diagnostics: r.Diagnostics, Current_scope: r.Current_scope, File: r.File, Is_public: r.Is_public, Pool: r.Pool, Index: r.Index}
}

func _rskip_pub(r Resolver) Resolver {
	rr := r
	_ = rr
	if (_rk(rr) == "pub") {
		rr = _radv(rr)
		if (_rk(rr) == "(") {
			rr = _radv(rr)
			if (_rk(rr) == "IDENT") {
				rr = _radv(rr)
			}
			if (_rk(rr) == ")") {
				rr = _radv(rr)
			}
		}
	}
	return rr
}

func _rset_pos(r Resolver, pos int64) Resolver {
	return Resolver{Tokens: r.Tokens, Pos: pos, Table: r.Table, Diagnostics: r.Diagnostics, Current_scope: r.Current_scope, File: r.File, Is_public: r.Is_public, Pool: r.Pool, Index: r.Index}
}

func _rset_table(r Resolver, table SymbolTable) Resolver {
	return Resolver{Tokens: r.Tokens, Pos: r.Pos, Table: table, Diagnostics: r.Diagnostics, Current_scope: r.Current_scope, File: r.File, Is_public: r.Is_public, Pool: r.Pool, Index: r.Index}
}

func _rset_diag(r Resolver, diag DiagnosticBag) Resolver {
	return Resolver{Tokens: r.Tokens, Pos: r.Pos, Table: r.Table, Diagnostics: diag, Current_scope: r.Current_scope, File: r.File, Is_public: r.Is_public, Pool: r.Pool, Index: r.Index}
}

func _rset_scope(r Resolver, scope_id int64) Resolver {
	return Resolver{Tokens: r.Tokens, Pos: r.Pos, Table: r.Table, Diagnostics: r.Diagnostics, Current_scope: scope_id, File: r.File, Is_public: r.Is_public, Pool: r.Pool, Index: r.Index}
}

func _rset_public(r Resolver, is_public bool) Resolver {
	return Resolver{Tokens: r.Tokens, Pos: r.Pos, Table: r.Table, Diagnostics: r.Diagnostics, Current_scope: r.Current_scope, File: r.File, Is_public: is_public, Pool: r.Pool, Index: r.Index}
}

func _rspan(r Resolver) Span {
	tok := _rcur(r)
	_ = tok
	return new_span(r.File, tok.Line, tok.Col, tok.Offset, int64(len(tok.Text)))
}

func _rskip_nl(r Resolver) Resolver {
	rr := r
	_ = rr
	for (_rk(rr) == "NEWLINE") {
		rr = _radv(rr)
	}
	return rr
}

func _rerror(r Resolver, code string, msg string) Resolver {
	span := _rspan(r)
	_ = span
	return _rset_diag(r, bag_add_error(r.Diagnostics, code, msg, span))
}

func _rwarning(r Resolver, code string, msg string) Resolver {
	span := _rspan(r)
	_ = span
	return _rset_diag(r, bag_add_warning(r.Diagnostics, code, msg, span))
}

func _add_sym(r Resolver, name string, kind SymbolKind, span Span, is_mut bool, is_pub bool) Resolver {
	existing := lookup_local(r.Table, r.Current_scope, name)
	_ = existing
	if (existing.Name != "") {
		q := _q()
		_ = q
		diag := bag_add_error(r.Diagnostics, E0702, ((("duplicate declaration of " + q) + name) + q), span)
		_ = diag
		return _rset_diag(r, diag)
	}
	sym := new_symbol(name, kind, span, r.Current_scope, is_mut, is_pub)
	_ = sym
	return _rset_table(r, add_symbol(r.Table, sym))
}

func _add_sym_warn_shadow(r Resolver, name string, kind SymbolKind, span Span, is_mut bool) Resolver {
	existing := lookup(r.Table, r.Current_scope, name)
	_ = existing
	rr := r
	_ = rr
	if (existing.Name != "") {
		if (sk_eq(existing.Kind, SymbolKindSkBuiltin{}) == false) {
			q := _q()
			_ = q
			rr = _rwarning(rr, W0004, (((("variable " + q) + name) + q) + " shadows previous declaration"))
		}
	}
	sym := new_symbol(name, kind, span, rr.Current_scope, is_mut, false)
	_ = sym
	return _rset_table(rr, add_symbol(rr.Table, sym))
}

func _push_scope(r Resolver, kind ScopeKind, name string) Resolver {
	new_t := add_scope(r.Table, kind, r.Current_scope, name)
	_ = new_t
	new_id := last_scope_id(new_t)
	_ = new_id
	rr := _rset_table(r, new_t)
	_ = rr
	return _rset_scope(rr, new_id)
}

func _pop_scope(r Resolver, parent int64) Resolver {
	return _rset_scope(r, parent)
}

func _builtin_span() Span {
	return new_span("<builtin>", int64(0), int64(0), int64(0), int64(0))
}

func _add_builtin(r Resolver, name string, kind SymbolKind) Resolver {
	sym := new_symbol(name, kind, _builtin_span(), r.Current_scope, false, false)
	_ = sym
	return _rset_table(r, add_symbol(r.Table, sym))
}

func _populate_builtins(r Resolver) Resolver {
	rr := r
	_ = rr
	rr = _add_builtin(rr, "i64", SymbolKindSkBuiltin{})
	rr = _add_builtin(rr, "i32", SymbolKindSkBuiltin{})
	rr = _add_builtin(rr, "i16", SymbolKindSkBuiltin{})
	rr = _add_builtin(rr, "i8", SymbolKindSkBuiltin{})
	rr = _add_builtin(rr, "u64", SymbolKindSkBuiltin{})
	rr = _add_builtin(rr, "u32", SymbolKindSkBuiltin{})
	rr = _add_builtin(rr, "u16", SymbolKindSkBuiltin{})
	rr = _add_builtin(rr, "u8", SymbolKindSkBuiltin{})
	rr = _add_builtin(rr, "f64", SymbolKindSkBuiltin{})
	rr = _add_builtin(rr, "f32", SymbolKindSkBuiltin{})
	rr = _add_builtin(rr, "bool", SymbolKindSkBuiltin{})
	rr = _add_builtin(rr, "str", SymbolKindSkBuiltin{})
	rr = _add_builtin(rr, "char", SymbolKindSkBuiltin{})
	rr = _add_builtin(rr, "unit", SymbolKindSkBuiltin{})
	rr = _add_builtin(rr, "print", SymbolKindSkBuiltin{})
	rr = _add_builtin(rr, "println", SymbolKindSkBuiltin{})
	rr = _add_builtin(rr, "eprintln", SymbolKindSkBuiltin{})
	rr = _add_builtin(rr, "assert", SymbolKindSkBuiltin{})
	rr = _add_builtin(rr, "assertEqual", SymbolKindSkBuiltin{})
	rr = _add_builtin(rr, "assertOk", SymbolKindSkBuiltin{})
	rr = _add_builtin(rr, "assertErr", SymbolKindSkBuiltin{})
	rr = _add_builtin(rr, "assertNear", SymbolKindSkBuiltin{})
	rr = _add_builtin(rr, "dbg", SymbolKindSkBuiltin{})
	rr = _add_builtin(rr, "panic", SymbolKindSkBuiltin{})
	rr = _add_builtin(rr, "_ariaReadFile", SymbolKindSkBuiltin{})
	rr = _add_builtin(rr, "_ariaWriteFile", SymbolKindSkBuiltin{})
	rr = _add_builtin(rr, "_ariaAppendFile", SymbolKindSkBuiltin{})
	rr = _add_builtin(rr, "_ariaWriteBinaryFile", SymbolKindSkBuiltin{})
	rr = _add_builtin(rr, "_ariaArgs", SymbolKindSkBuiltin{})
	rr = _add_builtin(rr, "_ariaExec", SymbolKindSkBuiltin{})
	rr = _add_builtin(rr, "_ariaListDir", SymbolKindSkBuiltin{})
	rr = _add_builtin(rr, "_ariaIsDir", SymbolKindSkBuiltin{})
	rr = _add_builtin(rr, "_ariaGetenv", SymbolKindSkBuiltin{})
	rr = _add_builtin(rr, "_ariaTcpSocket", SymbolKindSkBuiltin{})
	rr = _add_builtin(rr, "_ariaTcpBind", SymbolKindSkBuiltin{})
	rr = _add_builtin(rr, "_ariaTcpListen", SymbolKindSkBuiltin{})
	rr = _add_builtin(rr, "_ariaTcpAccept", SymbolKindSkBuiltin{})
	rr = _add_builtin(rr, "_ariaTcpRead", SymbolKindSkBuiltin{})
	rr = _add_builtin(rr, "_ariaTcpWrite", SymbolKindSkBuiltin{})
	rr = _add_builtin(rr, "_ariaTcpClose", SymbolKindSkBuiltin{})
	rr = _add_builtin(rr, "_ariaTcpPeerAddr", SymbolKindSkBuiltin{})
	rr = _add_builtin(rr, "_ariaTcpSetTimeout", SymbolKindSkBuiltin{})
	rr = _add_builtin(rr, "_ariaPgConnect", SymbolKindSkBuiltin{})
	rr = _add_builtin(rr, "_ariaPgClose", SymbolKindSkBuiltin{})
	rr = _add_builtin(rr, "_ariaPgStatus", SymbolKindSkBuiltin{})
	rr = _add_builtin(rr, "_ariaPgError", SymbolKindSkBuiltin{})
	rr = _add_builtin(rr, "_ariaPgExec", SymbolKindSkBuiltin{})
	rr = _add_builtin(rr, "_ariaPgExecParams", SymbolKindSkBuiltin{})
	rr = _add_builtin(rr, "_ariaPgResultStatus", SymbolKindSkBuiltin{})
	rr = _add_builtin(rr, "_ariaPgResultError", SymbolKindSkBuiltin{})
	rr = _add_builtin(rr, "_ariaPgNrows", SymbolKindSkBuiltin{})
	rr = _add_builtin(rr, "_ariaPgNcols", SymbolKindSkBuiltin{})
	rr = _add_builtin(rr, "_ariaPgFieldName", SymbolKindSkBuiltin{})
	rr = _add_builtin(rr, "_ariaPgGetValue", SymbolKindSkBuiltin{})
	rr = _add_builtin(rr, "_ariaPgIsNull", SymbolKindSkBuiltin{})
	rr = _add_builtin(rr, "_ariaPgClear", SymbolKindSkBuiltin{})
	rr = _add_builtin(rr, "_ariaSpawn", SymbolKindSkBuiltin{})
	rr = _add_builtin(rr, "_ariaTaskAwait", SymbolKindSkBuiltin{})
	rr = _add_builtin(rr, "_ariaChanNew", SymbolKindSkBuiltin{})
	rr = _add_builtin(rr, "_ariaChanSend", SymbolKindSkBuiltin{})
	rr = _add_builtin(rr, "_ariaChanRecv", SymbolKindSkBuiltin{})
	rr = _add_builtin(rr, "_ariaChanClose", SymbolKindSkBuiltin{})
	rr = _add_builtin(rr, "_ariaMutexNew", SymbolKindSkBuiltin{})
	rr = _add_builtin(rr, "_ariaMutexLock", SymbolKindSkBuiltin{})
	rr = _add_builtin(rr, "_ariaMutexUnlock", SymbolKindSkBuiltin{})
	rr = _add_builtin(rr, "_ariaRWMutexNew", SymbolKindSkBuiltin{})
	rr = _add_builtin(rr, "_ariaRWMutexRlock", SymbolKindSkBuiltin{})
	rr = _add_builtin(rr, "_ariaRWMutexRunlock", SymbolKindSkBuiltin{})
	rr = _add_builtin(rr, "_ariaRWMutexWlock", SymbolKindSkBuiltin{})
	rr = _add_builtin(rr, "_ariaRWMutexWunlock", SymbolKindSkBuiltin{})
	rr = _add_builtin(rr, "_ariaWgNew", SymbolKindSkBuiltin{})
	rr = _add_builtin(rr, "_ariaWgAdd", SymbolKindSkBuiltin{})
	rr = _add_builtin(rr, "_ariaWgDone", SymbolKindSkBuiltin{})
	rr = _add_builtin(rr, "_ariaWgWait", SymbolKindSkBuiltin{})
	rr = _add_builtin(rr, "_ariaOnceNew", SymbolKindSkBuiltin{})
	rr = _add_builtin(rr, "_ariaOnceCall", SymbolKindSkBuiltin{})
	rr = _add_builtin(rr, "_ariaChanTryRecv", SymbolKindSkBuiltin{})
	rr = _add_builtin(rr, "_ariaChanSelect", SymbolKindSkBuiltin{})
	rr = _add_builtin(rr, "_ariaSpawn2", SymbolKindSkBuiltin{})
	rr = _add_builtin(rr, "_ariaTaskAwait2", SymbolKindSkBuiltin{})
	rr = _add_builtin(rr, "_ariaTaskDone", SymbolKindSkBuiltin{})
	rr = _add_builtin(rr, "_ariaTaskCancel", SymbolKindSkBuiltin{})
	rr = _add_builtin(rr, "_ariaTaskResult", SymbolKindSkBuiltin{})
	rr = _add_builtin(rr, "_ariaCancelCheck", SymbolKindSkBuiltin{})
	rr = _add_builtin(rr, "_ariaCancelNew", SymbolKindSkBuiltin{})
	rr = _add_builtin(rr, "_ariaCancelChild", SymbolKindSkBuiltin{})
	rr = _add_builtin(rr, "_ariaCancelTrigger", SymbolKindSkBuiltin{})
	rr = _add_builtin(rr, "_ariaCancelIsTriggered", SymbolKindSkBuiltin{})
	rr = _add_builtin(rr, "_ariaSbNew", SymbolKindSkBuiltin{})
	rr = _add_builtin(rr, "_ariaSbWithCapacity", SymbolKindSkBuiltin{})
	rr = _add_builtin(rr, "_ariaSbAppend", SymbolKindSkBuiltin{})
	rr = _add_builtin(rr, "_ariaSbAppendChar", SymbolKindSkBuiltin{})
	rr = _add_builtin(rr, "_ariaSbLen", SymbolKindSkBuiltin{})
	rr = _add_builtin(rr, "_ariaSbBuild", SymbolKindSkBuiltin{})
	rr = _add_builtin(rr, "_ariaSbClear", SymbolKindSkBuiltin{})
	rr = _add_builtin(rr, "_ariaGcCollect", SymbolKindSkBuiltin{})
	rr = _add_builtin(rr, "_ariaGcTotalBytes", SymbolKindSkBuiltin{})
	rr = _add_builtin(rr, "_ariaGcAllocationCount", SymbolKindSkBuiltin{})
	rr = _add_builtin(rr, "_ariaArenaNew", SymbolKindSkBuiltin{})
	rr = _add_builtin(rr, "_ariaArenaAlloc", SymbolKindSkBuiltin{})
	rr = _add_builtin(rr, "_ariaArenaReset", SymbolKindSkBuiltin{})
	rr = _add_builtin(rr, "_ariaArenaFree", SymbolKindSkBuiltin{})
	rr = _add_builtin(rr, "_ariaArenaAllocated", SymbolKindSkBuiltin{})
	rr = _add_builtin(rr, "_ariaArenaCapacity", SymbolKindSkBuiltin{})
	rr = _add_builtin(rr, "_ariaPoolNew", SymbolKindSkBuiltin{})
	rr = _add_builtin(rr, "_ariaPoolGet", SymbolKindSkBuiltin{})
	rr = _add_builtin(rr, "_ariaPoolPut", SymbolKindSkBuiltin{})
	rr = _add_builtin(rr, "true", SymbolKindSkBuiltin{})
	rr = _add_builtin(rr, "or", SymbolKindSkBuiltin{})
	rr = _add_builtin(rr, "must", SymbolKindSkBuiltin{})
	rr = _add_builtin(rr, "from", SymbolKindSkBuiltin{})
	rr = _add_builtin(rr, "default", SymbolKindSkBuiltin{})
	rr = _add_builtin(rr, "_", SymbolKindSkBuiltin{})
	rr = _add_builtin(rr, "stack", SymbolKindSkBuiltin{})
	rr = _add_builtin(rr, "arena", SymbolKindSkBuiltin{})
	rr = _add_builtin(rr, "inline", SymbolKindSkBuiltin{})
	rr = _add_builtin(rr, "false", SymbolKindSkBuiltin{})
	rr = _add_builtin(rr, "Ok", SymbolKindSkBuiltin{})
	rr = _add_builtin(rr, "Err", SymbolKindSkBuiltin{})
	rr = _add_builtin(rr, "Some", SymbolKindSkBuiltin{})
	rr = _add_builtin(rr, "None", SymbolKindSkBuiltin{})
	rr = _add_builtin(rr, "Map", SymbolKindSkBuiltin{})
	rr = _add_builtin(rr, "Set", SymbolKindSkBuiltin{})
	rr = _add_builtin(rr, "dyn", SymbolKindSkBuiltin{})
	rr = _add_builtin(rr, "Io", SymbolKindSkBuiltin{})
	rr = _add_builtin(rr, "Fs", SymbolKindSkBuiltin{})
	rr = _add_builtin(rr, "Net", SymbolKindSkBuiltin{})
	rr = _add_builtin(rr, "Ffi", SymbolKindSkBuiltin{})
	rr = _add_builtin(rr, "Async", SymbolKindSkBuiltin{})
	return rr
}

func _is_decl_keyword(k string) bool {
	return (((((((k == "fn") || (k == "type")) || (k == "struct")) || (k == "enum")) || (k == "trait")) || (k == "impl")) || (k == "const"))
}

func _pass1(r Resolver, index DeclIndex) Resolver {
	rr := r
	_ = rr
	di := int64(1)
	_ = di
	for (di < decl_index_len(index)) {
		decl := decl_index_get(index, di)
		_ = decl
		kname := decl_kind_name(decl.Kind)
		_ = kname
		is_pub := decl.Is_pub
		_ = is_pub
		rr = _rset_pos(rr, decl.Token_start)
		if (kname == "fn") {
			if (decl.Name == "_extern") {
				rr = _radv(rr)
				lb2 := "{"
				_ = lb2
				rb2 := "}"
				_ = rb2
				if (_rk(rr) == "STRING") {
					rr = _radv(rr)
				}
				if (_rk(rr) == lb2) {
					rr = _radv(rr)
					rr = _rskip_nl(rr)
					for ((_rk(rr) != rb2) && (_rk(rr) != "EOF")) {
						rr = _rskip_nl(rr)
						if (_rk(rr) == "fn") {
							rr = _radv(rr)
							if (_rk(rr) == "IDENT") {
								ename := _rcur(rr).Text
								_ = ename
								espan := _rspan(rr)
								_ = espan
								rr = _add_sym(rr, ename, SymbolKindSkFunction{}, espan, false, is_pub)
							}
						}
						for (((_rk(rr) != "NEWLINE") && (_rk(rr) != rb2)) && (_rk(rr) != "EOF")) {
							rr = _radv(rr)
						}
						rr = _rskip_nl(rr)
					}
				} else if (_rk(rr) == "fn") {
					rr = _radv(rr)
					if (_rk(rr) == "IDENT") {
						ename := _rcur(rr).Text
						_ = ename
						espan := _rspan(rr)
						_ = espan
						rr = _add_sym(rr, ename, SymbolKindSkFunction{}, espan, false, is_pub)
					}
				}
			} else if strings.HasPrefix(decl.Name, "_fixture:") {
				_skip := int64(0)
				_ = _skip
			} else {
				rr = _rskip_pub(rr)
				if (_rk(rr) == "@") {
					rr = _radv(rr)
					if (_rk(rr) == "IDENT") {
						rr = _radv(rr)
					}
					if (_rk(rr) == "(") {
						rr = _radv(rr)
						for ((_rk(rr) != ")") && (_rk(rr) != "EOF")) {
							rr = _radv(rr)
						}
						if (_rk(rr) == ")") {
							rr = _radv(rr)
						}
					}
					rr = _rskip_nl(rr)
				}
				rr = _radv(rr)
				if (_rk(rr) == "IDENT") {
					name := _rcur(rr).Text
					_ = name
					span := _rspan(rr)
					_ = span
					rr = _add_sym(rr, name, SymbolKindSkFunction{}, span, false, is_pub)
				}
			}
		} else if (kname == "type") {
			if strings.HasPrefix(decl.Name, "_alias:") {
				alias_name := decl.Name[int64(7):int64(len(decl.Name))]
				_ = alias_name
				rr = _radv(rr)
				if (_rk(rr) == "IDENT") {
					span := _rspan(rr)
					_ = span
					rr = _add_sym(rr, alias_name, SymbolKindSkType{}, span, false, is_pub)
				}
			} else {
				rr = _rskip_pub(rr)
				rr = _radv(rr)
				if (_rk(rr) == "IDENT") {
					name := _rcur(rr).Text
					_ = name
					span := _rspan(rr)
					_ = span
					rr = _add_sym(rr, name, SymbolKindSkType{}, span, false, is_pub)
					rr = _radv(rr)
					rr = _pass1_collect_variants(rr)
				}
			}
		} else if (kname == "enum") {
			rr = _rskip_pub(rr)
			rr = _radv(rr)
			if (_rk(rr) == "IDENT") {
				name := _rcur(rr).Text
				_ = name
				span := _rspan(rr)
				_ = span
				rr = _add_sym(rr, name, SymbolKindSkEnum{}, span, false, is_pub)
			}
		} else if (kname == "trait") {
			rr = _rskip_pub(rr)
			rr = _radv(rr)
			if (_rk(rr) == "IDENT") {
				name := _rcur(rr).Text
				_ = name
				span := _rspan(rr)
				_ = span
				rr = _add_sym(rr, name, SymbolKindSkTrait{}, span, false, is_pub)
			}
		} else if (kname == "const") {
			rr = _rskip_pub(rr)
			rr = _radv(rr)
			if (_rk(rr) == "IDENT") {
				name := _rcur(rr).Text
				_ = name
				span := _rspan(rr)
				_ = span
				rr = _add_sym(rr, name, SymbolKindSkConst{}, span, false, is_pub)
			}
		}
		di = (di + int64(1))
	}
	return rr
}

func _pass1_collect_variants(r Resolver) Resolver {
	rr := r
	_ = rr
	if (_rk(rr) == "[") {
		d := int64(1)
		_ = d
		rr = _radv(rr)
		for ((d > int64(0)) && (_rk(rr) != "EOF")) {
			if (_rk(rr) == "[") {
				d = (d + int64(1))
			}
			if (_rk(rr) == "]") {
				d = (d - int64(1))
			}
			if (d > int64(0)) {
				rr = _radv(rr)
			}
		}
		if (_rk(rr) == "]") {
			rr = _radv(rr)
		}
	}
	if (_rk(rr) != "=") {
		return rr
	}
	rr = _radv(rr)
	rr = _rskip_nl(rr)
	if (_rk(rr) != "|") {
		return rr
	}
	for (_rk(rr) == "|") {
		rr = _radv(rr)
		rr = _rskip_nl(rr)
		if (_rk(rr) == "IDENT") {
			name := _rcur(rr).Text
			_ = name
			span := _rspan(rr)
			_ = span
			rr = _add_sym(rr, name, SymbolKindSkVariant{}, span, false, false)
		}
		rr = _radv(rr)
		if (_rk(rr) == "(") {
			d := int64(1)
			_ = d
			rr = _radv(rr)
			for ((d > int64(0)) && (_rk(rr) != "EOF")) {
				if (_rk(rr) == "(") {
					d = (d + int64(1))
				}
				if (_rk(rr) == ")") {
					d = (d - int64(1))
				}
				if (d > int64(0)) {
					rr = _radv(rr)
				}
			}
			if (_rk(rr) == ")") {
				rr = _radv(rr)
			}
		}
		lb := "{"
		_ = lb
		rb := "}"
		_ = rb
		if (_rk(rr) == lb) {
			d := int64(1)
			_ = d
			rr = _radv(rr)
			for ((d > int64(0)) && (_rk(rr) != "EOF")) {
				if (_rk(rr) == lb) {
					d = (d + int64(1))
				}
				if (_rk(rr) == rb) {
					d = (d - int64(1))
				}
				if (d > int64(0)) {
					rr = _radv(rr)
				}
			}
			if (_rk(rr) == rb) {
				rr = _radv(rr)
			}
		}
		rr = _rskip_nl(rr)
	}
	return rr
}

func _pass2(r Resolver, index DeclIndex) Resolver {
	rr := r
	_ = rr
	di := int64(1)
	_ = di
	for (di < decl_index_len(index)) {
		decl := decl_index_get(index, di)
		_ = decl
		kname := decl_kind_name(decl.Kind)
		_ = kname
		rr = _rset_pos(rr, decl.Token_start)
		rr = _rskip_pub(rr)
		if (kname == "fn") {
			if (decl.Name == "_extern") {
				rr = _rset_pos(rr, decl.Body_end)
			} else if strings.HasPrefix(decl.Name, "_fixture:") {
				rr = _resolve_test(rr)
			} else if (decl.Node_idx > int64(0)) {
				rr = _resolve_fn_ast(rr, decl.Node_idx)
			} else {
				if (_rk(rr) == "@") {
					rr = _radv(rr)
					if (_rk(rr) == "IDENT") {
						rr = _radv(rr)
					}
					if (_rk(rr) == "(") {
						rr = _radv(rr)
						for ((_rk(rr) != ")") && (_rk(rr) != "EOF")) {
							rr = _radv(rr)
						}
						if (_rk(rr) == ")") {
							rr = _radv(rr)
						}
					}
					rr = _rskip_nl(rr)
				}
				rr = _resolve_fn(rr)
			}
		} else if (kname == "impl") {
			if (decl.Node_idx > int64(0)) {
				rr = _resolve_impl_ast(rr, decl.Node_idx)
			} else {
				rr = _resolve_impl(rr)
			}
		} else if (kname == "const") {
			if (decl.Node_idx > int64(0)) {
				rr = _resolve_const_ast(rr, decl.Node_idx)
			} else {
				rr = _resolve_const(rr)
			}
		} else if (kname == "entry") {
			if (decl.Node_idx > int64(0)) {
				rr = _resolve_entry_ast(rr, decl.Node_idx)
			} else {
				rr = _resolve_entry(rr)
			}
		} else if (kname == "test") {
			if (decl.Node_idx > int64(0)) {
				rr = _resolve_test_ast(rr, decl.Node_idx)
			} else {
				rr = _resolve_test(rr)
			}
		} else {
			rr = _rset_pos(rr, decl.Body_end)
		}
		di = (di + int64(1))
	}
	return rr
}

func _resolve_fn(r Resolver) Resolver {
	parent_scope := r.Current_scope
	_ = parent_scope
	rr := _radv(r)
	_ = rr
	fn_name := _rcur(rr).Text
	_ = fn_name
	rr = _radv(rr)
	rr = _push_scope(rr, ScopeKindScFunction{}, fn_name)
	if (_rk(rr) == "[") {
		rr = _radv(rr)
		for ((_rk(rr) != "]") && (_rk(rr) != "EOF")) {
			if (_rk(rr) == "IDENT") {
				gname := _rcur(rr).Text
				_ = gname
				gspan := _rspan(rr)
				_ = gspan
				rr = _add_sym(rr, gname, SymbolKindSkType{}, gspan, false, false)
				rr = _radv(rr)
				if (_rk(rr) == "[") {
					rr = _radv(rr)
					if ((_rk(rr) == "_") || (((_rk(rr) == "IDENT") && (_rcur(rr).Text == "_")))) {
						rr = _radv(rr)
					}
					if (_rk(rr) == "]") {
						rr = _radv(rr)
					}
				}
				if (_rk(rr) == ":") {
					rr = _radv(rr)
					if (_rk(rr) == "IDENT") {
						rr = _radv(rr)
					}
					if (_rk(rr) == "<") {
						rr = _radv(rr)
						if (_rk(rr) == "IDENT") {
							rr = _radv(rr)
						}
						if (_rk(rr) == "=") {
							rr = _radv(rr)
						}
						if (_rk(rr) == "IDENT") {
							rr = _radv(rr)
						}
						if (_rk(rr) == ">") {
							rr = _radv(rr)
						}
					}
					for (_rk(rr) == "+") {
						rr = _radv(rr)
						if (_rk(rr) == "IDENT") {
							rr = _radv(rr)
						}
						if (_rk(rr) == "<") {
							rr = _radv(rr)
							if (_rk(rr) == "IDENT") {
								rr = _radv(rr)
							}
							if (_rk(rr) == "=") {
								rr = _radv(rr)
							}
							if (_rk(rr) == "IDENT") {
								rr = _radv(rr)
							}
							if (_rk(rr) == ">") {
								rr = _radv(rr)
							}
						}
					}
				}
			} else {
				rr = _radv(rr)
			}
			if (_rk(rr) == ",") {
				rr = _radv(rr)
			}
		}
		if (_rk(rr) == "]") {
			rr = _radv(rr)
		}
	}
	if (_rk(rr) == "(") {
		rr = _radv(rr)
		for ((_rk(rr) != ")") && (_rk(rr) != "EOF")) {
			if (_rk(rr) == "self") {
				pspan := _rspan(rr)
				_ = pspan
				rr = _add_sym_warn_shadow(rr, "self", SymbolKindSkParameter{}, pspan, false)
				rr = _radv(rr)
			} else if (((((((((_rk(rr) == "IDENT") || (_rk(rr) == "before")) || (_rk(rr) == "after")) || (_rk(rr) == "where")) || (_rk(rr) == "scope")) || (_rk(rr) == "select")) || (_rk(rr) == "as")) || (_rk(rr) == "is")) || (_rk(rr) == "catch")) {
				pname := _rcur(rr).Text
				_ = pname
				pspan := _rspan(rr)
				_ = pspan
				rr = _add_sym_warn_shadow(rr, pname, SymbolKindSkParameter{}, pspan, false)
				rr = _radv(rr)
				if (_rk(rr) == ":") {
					rr = _radv(rr)
					rr = _skip_type_tokens(rr)
				}
			} else {
				rr = _radv(rr)
			}
			if (_rk(rr) == ",") {
				rr = _radv(rr)
			}
		}
		if (_rk(rr) == ")") {
			rr = _radv(rr)
		}
	}
	if (_rk(rr) == "->") {
		rr = _radv(rr)
		rr = _skip_type_tokens(rr)
	}
	if (_rk(rr) == "!") {
		rr = _radv(rr)
		rr = _skip_type_tokens(rr)
	}
	if (_rk(rr) == "with") {
		lb := "{"
		_ = lb
		for ((((_rk(rr) != "NEWLINE") && (_rk(rr) != "EOF")) && (_rk(rr) != "=")) && (_rk(rr) != lb)) {
			rr = _radv(rr)
		}
	}
	lb := "{"
	_ = lb
	if (_rk(rr) == "=") {
		rr = _radv(rr)
		rr = _rskip_nl(rr)
		rr = _resolve_expr_until_nl(rr)
	} else if (_rk(rr) == lb) {
		rr = _resolve_block(rr)
	} else {
		rr = _rskip_nl(rr)
		if (_rk(rr) == lb) {
			rr = _resolve_block(rr)
		}
	}
	return _pop_scope(rr, parent_scope)
}

func _resolve_fn_ast(r Resolver, node_idx int64) Resolver {
	node := pool_get(r.Pool, node_idx)
	_ = node
	parent_scope := r.Current_scope
	_ = parent_scope
	rr := _push_scope(r, ScopeKindScFunction{}, node.S1)
	_ = rr
	rr = _rset_pos(rr, r.Pos)
	if (_rk(rr) == "@") {
		rr = _radv(rr)
		if (_rk(rr) == "IDENT") {
			rr = _radv(rr)
		}
		if (_rk(rr) == "(") {
			rr = _radv(rr)
			for ((_rk(rr) != ")") && (_rk(rr) != "EOF")) {
				rr = _radv(rr)
			}
			if (_rk(rr) == ")") {
				rr = _radv(rr)
			}
		}
		rr = _rskip_nl(rr)
	}
	rr = _radv(rr)
	rr = _radv(rr)
	if (_rk(rr) == "[") {
		rr = _radv(rr)
		for ((_rk(rr) != "]") && (_rk(rr) != "EOF")) {
			if (_rk(rr) == "IDENT") {
				gname := _rcur(rr).Text
				_ = gname
				gspan := _rspan(rr)
				_ = gspan
				rr = _add_sym(rr, gname, SymbolKindSkType{}, gspan, false, false)
				rr = _radv(rr)
				if (_rk(rr) == "[") {
					rr = _radv(rr)
					if ((_rk(rr) == "_") || (((_rk(rr) == "IDENT") && (_rcur(rr).Text == "_")))) {
						rr = _radv(rr)
					}
					if (_rk(rr) == "]") {
						rr = _radv(rr)
					}
				}
				if (_rk(rr) == ":") {
					rr = _radv(rr)
					if (_rk(rr) == "IDENT") {
						rr = _radv(rr)
					}
					if (_rk(rr) == "<") {
						rr = _radv(rr)
						if (_rk(rr) == "IDENT") {
							rr = _radv(rr)
						}
						if (_rk(rr) == "=") {
							rr = _radv(rr)
						}
						if (_rk(rr) == "IDENT") {
							rr = _radv(rr)
						}
						if (_rk(rr) == ">") {
							rr = _radv(rr)
						}
					}
					for (_rk(rr) == "+") {
						rr = _radv(rr)
						if (_rk(rr) == "IDENT") {
							rr = _radv(rr)
						}
						if (_rk(rr) == "<") {
							rr = _radv(rr)
							if (_rk(rr) == "IDENT") {
								rr = _radv(rr)
							}
							if (_rk(rr) == "=") {
								rr = _radv(rr)
							}
							if (_rk(rr) == "IDENT") {
								rr = _radv(rr)
							}
							if (_rk(rr) == ">") {
								rr = _radv(rr)
							}
						}
					}
				}
			} else {
				rr = _radv(rr)
			}
			if (_rk(rr) == ",") {
				rr = _radv(rr)
			}
		}
		if (_rk(rr) == "]") {
			rr = _radv(rr)
		}
	}
	pi := int64(0)
	_ = pi
	for (pi < node.List_count) {
		idx_node := pool_get(rr.Pool, (node.List_start + pi))
		_ = idx_node
		param_node := pool_get(rr.Pool, idx_node.C0)
		_ = param_node
		pname := param_node.S1
		_ = pname
		plen := int64(len(pname))
		_ = plen
		if (plen > int64(0)) {
			rr = _add_sym_warn_shadow(rr, pname, SymbolKindSkParameter{}, param_node.Span, false)
		}
		pi = (pi + int64(1))
	}
	if node.B1 {
		rr = _resolve_expr_node(rr, node.C0)
	}
	return _pop_scope(rr, parent_scope)
}

func _resolve_const_ast(r Resolver, node_idx int64) Resolver {
	node := pool_get(r.Pool, node_idx)
	_ = node
	return _resolve_expr_node(r, node.C0)
}

func _resolve_entry_ast(r Resolver, node_idx int64) Resolver {
	node := pool_get(r.Pool, node_idx)
	_ = node
	parent_scope := r.Current_scope
	_ = parent_scope
	rr := _push_scope(r, ScopeKindScFunction{}, "entry")
	_ = rr
	si := int64(0)
	_ = si
	for (si < node.List_count) {
		idx_node := pool_get(r.Pool, (node.List_start + si))
		_ = idx_node
		rr = _resolve_expr_node(rr, idx_node.C0)
		si = (si + int64(1))
	}
	return _pop_scope(rr, parent_scope)
}

func _resolve_test_ast(r Resolver, node_idx int64) Resolver {
	node := pool_get(r.Pool, node_idx)
	_ = node
	parent_scope := r.Current_scope
	_ = parent_scope
	rr := _push_scope(r, ScopeKindScFunction{}, "test")
	_ = rr
	si := int64(0)
	_ = si
	for (si < node.List_count) {
		idx_node := pool_get(r.Pool, (node.List_start + si))
		_ = idx_node
		rr = _resolve_expr_node(rr, idx_node.C0)
		si = (si + int64(1))
	}
	return _pop_scope(rr, parent_scope)
}

func _resolve_impl_ast(r Resolver, node_idx int64) Resolver {
	node := pool_get(r.Pool, node_idx)
	_ = node
	rr := r
	_ = rr
	mi := int64(0)
	_ = mi
	for (mi < node.List_count) {
		idx_node := pool_get(r.Pool, (node.List_start + mi))
		_ = idx_node
		method_idx := idx_node.C0
		_ = method_idx
		rr = _resolve_fn_ast(rr, method_idx)
		mi = (mi + int64(1))
	}
	return rr
}

func _resolve_impl(r Resolver) Resolver {
	rr := _radv(r)
	_ = rr
	lb := "{"
	_ = lb
	for ((_rk(rr) != lb) && (_rk(rr) != "EOF")) {
		rr = _radv(rr)
	}
	if (_rk(rr) != lb) {
		return rr
	}
	rr = _radv(rr)
	rr = _rskip_nl(rr)
	rb := "}"
	_ = rb
	for ((_rk(rr) != rb) && (_rk(rr) != "EOF")) {
		rr = _rskip_nl(rr)
		rr = _rskip_pub(rr)
		if (_rk(rr) == "fn") {
			rr = _resolve_fn(rr)
		} else if ((_rk(rr) == rb) || (_rk(rr) == "EOF")) {
			return rr
		} else {
			rr = _radv(rr)
		}
		rr = _rskip_nl(rr)
	}
	if (_rk(rr) == rb) {
		rr = _radv(rr)
	}
	return rr
}

func _resolve_const(r Resolver) Resolver {
	rr := _radv(r)
	_ = rr
	if (_rk(rr) == "IDENT") {
		rr = _radv(rr)
	}
	if (_rk(rr) == ":") {
		rr = _radv(rr)
		rr = _skip_type_tokens(rr)
	}
	if (_rk(rr) == "=") {
		rr = _radv(rr)
		rr = _resolve_expr_until_nl(rr)
	}
	return rr
}

func _resolve_entry(r Resolver) Resolver {
	parent_scope := r.Current_scope
	_ = parent_scope
	rr := _radv(r)
	_ = rr
	rr = _push_scope(rr, ScopeKindScFunction{}, "entry")
	rr = _resolve_block(rr)
	return _pop_scope(rr, parent_scope)
}

func _resolve_test(r Resolver) Resolver {
	parent_scope := r.Current_scope
	_ = parent_scope
	rr := _radv(r)
	_ = rr
	if ((_rk(rr) == "STRING") || (_rk(rr) == "IDENT")) {
		rr = _radv(rr)
	}
	rr = _push_scope(rr, ScopeKindScFunction{}, "test")
	rr = _resolve_block(rr)
	return _pop_scope(rr, parent_scope)
}

func _resolve_block(r Resolver) Resolver {
	lb := "{"
	_ = lb
	rb := "}"
	_ = rb
	if (_rk(r) != lb) {
		return r
	}
	parent_scope := r.Current_scope
	_ = parent_scope
	rr := _push_scope(r, ScopeKindScBlock{}, "")
	_ = rr
	rr = _radv(rr)
	rr = _rskip_nl(rr)
	for ((_rk(rr) != rb) && (_rk(rr) != "EOF")) {
		rr = _resolve_stmt(rr)
		rr = _rskip_nl(rr)
	}
	if (_rk(rr) == rb) {
		rr = _radv(rr)
	}
	return _pop_scope(rr, parent_scope)
}

func _resolve_struct_lit(r Resolver) Resolver {
	lb := "{"
	_ = lb
	rb := "}"
	_ = rb
	if (_rk(r) != lb) {
		return r
	}
	rr := _radv(r)
	_ = rr
	rr = _rskip_nl(rr)
	for ((_rk(rr) != rb) && (_rk(rr) != "EOF")) {
		if (_rk(rr) == "IDENT") {
			rr = _radv(rr)
			if (_rk(rr) == ":") {
				rr = _radv(rr)
				rr = _resolve_struct_field_value(rr)
			}
		} else {
			rr = _radv(rr)
		}
		rr = _rskip_nl(rr)
		if (_rk(rr) == ",") {
			rr = _radv(rr)
			rr = _rskip_nl(rr)
		}
	}
	if (_rk(rr) == rb) {
		rr = _radv(rr)
	}
	return rr
}

func _resolve_struct_field_value(r Resolver) Resolver {
	rr := r
	_ = rr
	rb := "}"
	_ = rb
	lb := "{"
	_ = lb
	for ((((_rk(rr) != ",") && (_rk(rr) != rb)) && (_rk(rr) != "NEWLINE")) && (_rk(rr) != "EOF")) {
		k := _rk(rr)
		_ = k
		if (k == "IDENT") {
			if ((rr.Pos > int64(0)) && (token_name(rr.Tokens[(rr.Pos - int64(1))].Kind) == ".")) {
				rr = _radv(rr)
			} else {
				rr = _resolve_ident(rr)
			}
		} else if (k == lb) {
			if ((rr.Pos > int64(0)) && (token_name(rr.Tokens[(rr.Pos - int64(1))].Kind) == "IDENT")) {
				rr = _resolve_struct_lit(rr)
			} else {
				rr = _resolve_block(rr)
			}
		} else if (k == "(") {
			rr = _radv(rr)
			rr = _resolve_paren_contents(rr)
		} else if (k == "[") {
			rr = _radv(rr)
			rr = _resolve_bracket_contents(rr)
		} else {
			rr = _radv(rr)
		}
	}
	return rr
}

func _resolve_paren_contents(r Resolver) Resolver {
	rr := r
	_ = rr
	for ((_rk(rr) != ")") && (_rk(rr) != "EOF")) {
		k := _rk(rr)
		_ = k
		if (k == "IDENT") {
			if ((rr.Pos > int64(0)) && (token_name(rr.Tokens[(rr.Pos - int64(1))].Kind) == ".")) {
				rr = _radv(rr)
			} else {
				rr = _resolve_ident(rr)
			}
		} else if (k == "(") {
			rr = _radv(rr)
			rr = _resolve_paren_contents(rr)
		} else {
			rr = _radv(rr)
		}
	}
	if (_rk(rr) == ")") {
		rr = _radv(rr)
	}
	return rr
}

func _resolve_bracket_contents(r Resolver) Resolver {
	rr := r
	_ = rr
	lb := "{"
	_ = lb
	for ((_rk(rr) != "]") && (_rk(rr) != "EOF")) {
		k := _rk(rr)
		_ = k
		if (k == "IDENT") {
			if ((rr.Pos > int64(0)) && (token_name(rr.Tokens[(rr.Pos - int64(1))].Kind) == ".")) {
				rr = _radv(rr)
			} else {
				rr = _resolve_ident(rr)
			}
		} else if (k == lb) {
			if ((rr.Pos > int64(0)) && (token_name(rr.Tokens[(rr.Pos - int64(1))].Kind) == "IDENT")) {
				rr = _resolve_struct_lit(rr)
			} else {
				rr = _resolve_block(rr)
			}
		} else if (k == "[") {
			rr = _radv(rr)
			rr = _resolve_bracket_contents(rr)
		} else if (k == "(") {
			rr = _radv(rr)
			rr = _resolve_paren_contents(rr)
		} else {
			rr = _radv(rr)
		}
	}
	if (_rk(rr) == "]") {
		rr = _radv(rr)
	}
	return rr
}

func _resolve_stmt(r Resolver) Resolver {
	k := _rk(r)
	_ = k
	if (k == "@") {
		rr := _radv(r)
		_ = rr
		if (_rk(rr) == "IDENT") {
			rr = _radv(rr)
		}
		if (_rk(rr) == "(") {
			rr = _radv(rr)
			for ((_rk(rr) != ")") && (_rk(rr) != "EOF")) {
				rr = _radv(rr)
			}
			if (_rk(rr) == ")") {
				rr = _radv(rr)
			}
		}
		rr = _rskip_nl(rr)
		return _resolve_stmt(rr)
	}
	if (k == "mut") {
		return _resolve_mut_binding(r)
	}
	if (k == "return") {
		rr := _radv(r)
		_ = rr
		rb := "}"
		_ = rb
		if (((_rk(rr) == "NEWLINE") || (_rk(rr) == "EOF")) || (_rk(rr) == rb)) {
			return rr
		}
		return _resolve_expr_until_nl(rr)
	}
	if ((k == "break") || (k == "continue")) {
		return _radv(r)
	}
	if (k == "for") {
		return _resolve_for(r)
	}
	if (k == "while") {
		return _resolve_while(r)
	}
	if (k == "loop") {
		return _resolve_loop(r)
	}
	if (k == "defer") {
		rr := _radv(r)
		_ = rr
		return _resolve_expr_until_nl(rr)
	}
	if (k == "assert") {
		rr := _radv(r)
		_ = rr
		return _resolve_expr_until_nl(rr)
	}
	if (k == "mock") {
		rr := _radv(r)
		_ = rr
		if (_rk(rr) == "IDENT") {
			rr = _radv(rr)
		}
		if (_rk(rr) == "with") {
			rr = _radv(rr)
		}
		return _resolve_expr_until_nl(rr)
	}
	if (k == "if") {
		return _resolve_if(r)
	}
	if (k == "match") {
		return _resolve_match(r)
	}
	if (k == "with") {
		return _resolve_with(r)
	}
	return _resolve_expr_or_binding(r)
}

func _resolve_mut_binding(r Resolver) Resolver {
	rr := _radv(r)
	_ = rr
	if (_rk(rr) == "IDENT") {
		name := _rcur(rr).Text
		_ = name
		span := _rspan(rr)
		_ = span
		rr = _radv(rr)
		if (_rk(rr) == ":") {
			rr = _radv(rr)
			rr = _skip_type_tokens(rr)
			if (_rk(rr) == "=") {
				rr = _radv(rr)
			}
		} else if (_rk(rr) == ":=") {
			rr = _radv(rr)
		}
		rr = _resolve_expr_until_nl(rr)
		rr = _add_sym_warn_shadow(rr, name, SymbolKindSkVariable{}, span, true)
		return rr
	}
	return rr
}

func _resolve_expr_or_binding(r Resolver) Resolver {
	if (_rk(r) == "(") {
		scan := (r.Pos + int64(1))
		_ = scan
		is_tuple_destr := false
		_ = is_tuple_destr
		for (scan < int64(len(r.Tokens))) {
			sk := token_name(r.Tokens[scan].Kind)
			_ = sk
			if (sk == ")") {
				next2 := (scan + int64(1))
				_ = next2
				if ((next2 < int64(len(r.Tokens))) && (token_name(r.Tokens[next2].Kind) == ":=")) {
					is_tuple_destr = true
				}
				scan = int64(len(r.Tokens))
			} else if (((sk == "IDENT") || (sk == ",")) || (sk == "_")) {
				scan = (scan + int64(1))
			} else {
				scan = int64(len(r.Tokens))
			}
		}
		if is_tuple_destr {
			rr := _radv(r)
			_ = rr
			names := []string{""}
			_ = names
			name_count := int64(0)
			_ = name_count
			for ((_rk(rr) != ")") && (_rk(rr) != "EOF")) {
				if (_rk(rr) == "IDENT") {
					names = append(names, _rcur(rr).Text)
					name_count = (name_count + int64(1))
				}
				rr = _radv(rr)
				if (_rk(rr) == ",") {
					rr = _radv(rr)
				}
			}
			if (_rk(rr) == ")") {
				rr = _radv(rr)
			}
			if (_rk(rr) == ":=") {
				rr = _radv(rr)
			}
			rr = _resolve_expr_until_nl(rr)
			ni := int64(1)
			_ = ni
			for (ni < int64(len(names))) {
				span := _rspan(rr)
				_ = span
				rr = _add_sym_warn_shadow(rr, names[ni], SymbolKindSkVariable{}, span, false)
				ni = (ni + int64(1))
			}
			return rr
		}
	}
	if (_rk(r) == "IDENT") {
		name := _rcur(r).Text
		_ = name
		next_pos := (r.Pos + int64(1))
		_ = next_pos
		if (next_pos < int64(len(r.Tokens))) {
			next_k := token_name(r.Tokens[next_pos].Kind)
			_ = next_k
			if (next_k == ":=") {
				span := _rspan(r)
				_ = span
				rr := _radv(r)
				_ = rr
				rr = _radv(rr)
				rr = _resolve_expr_until_nl(rr)
				rr = _add_sym_warn_shadow(rr, name, SymbolKindSkVariable{}, span, false)
				return rr
			}
			if (next_k == ":") {
				rr := _radv(r)
				_ = rr
				rr = _radv(rr)
				saved_pos := rr.Pos
				_ = saved_pos
				rr = _skip_type_tokens(rr)
				if (_rk(rr) == "=") {
					span := _rspan(r)
					_ = span
					rr = _radv(rr)
					rr = _resolve_expr_until_nl(rr)
					rr = _add_sym_warn_shadow(rr, name, SymbolKindSkVariable{}, _rspan(_rset_pos(r, r.Pos)), false)
					return rr
				}
				rr = _rset_pos(rr, r.Pos)
			}
		}
	}
	return _resolve_expr_until_nl(r)
}

func _is_expr_continuation(k string) bool {
	return (((((((((((((((((((((((((((((((((k == "+") || (k == "-")) || (k == "*")) || (k == "/")) || (k == "%")) || (k == "==")) || (k == "!=")) || (k == "<")) || (k == ">")) || (k == "<=")) || (k == ">=")) || (k == "&&")) || (k == "||")) || (k == "&")) || (k == "|")) || (k == "^")) || (k == "~")) || (k == "<<")) || (k == ">>")) || (k == "|>")) || (k == "..")) || (k == "..=")) || (k == ".")) || (k == "(")) || (k == "[")) || (k == "?")) || (k == "!")) || (k == ",")) || (k == "=>")) || (k == "=")) || (k == "?.")) || (k == "??")) || (k == "catch"))
}

func _resolve_expr_until_nl(r Resolver) Resolver {
	rr := r
	_ = rr
	rb := "}"
	_ = rb
	consumed_value := false
	_ = consumed_value
	after_dot := false
	_ = after_dot
	for (((_rk(rr) != "NEWLINE") && (_rk(rr) != "EOF")) && (_rk(rr) != rb)) {
		k := _rk(rr)
		_ = k
		lb := "{"
		_ = lb
		if (consumed_value && (_is_expr_continuation(k) == false)) {
			if ((((((((((((((((((((k == "IDENT") || (k == "INT")) || (k == "FLOAT")) || (k == "STRING")) || (k == "true")) || (k == "false")) || (k == "mut")) || (k == "return")) || (k == "break")) || (k == "continue")) || (k == "for")) || (k == "while")) || (k == "loop")) || (k == "defer")) || (k == "assert")) || (k == "if")) || (k == "match")) || (k == "const")) || (k == "fn")) || (k == "with")) {
				return rr
			}
		}
		if (k == "IDENT") {
			if after_dot {
				rr = _radv(rr)
			} else {
				rr = _resolve_ident(rr)
			}
			after_dot = false
			consumed_value = true
		} else if (k == "if") {
			rr = _resolve_if(rr)
			return rr
		} else if (k == "match") {
			rr = _resolve_match(rr)
			return rr
		} else if (k == lb) {
			is_struct_lit := false
			_ = is_struct_lit
			if (rr.Pos > int64(0)) {
				prev_k := token_name(rr.Tokens[(rr.Pos - int64(1))].Kind)
				_ = prev_k
				if (consumed_value && (prev_k == "IDENT")) {
					is_struct_lit = true
				}
				if (consumed_value && (prev_k == "]")) {
					is_struct_lit = true
				}
				if (prev_k == ".") {
					is_struct_lit = true
				}
			}
			if is_struct_lit {
				rr = _resolve_struct_lit(rr)
			} else {
				rr = _resolve_block(rr)
			}
			consumed_value = true
		} else if (k == "fn") {
			rr = _resolve_closure(rr)
			consumed_value = true
		} else if (((((((k == "INT") || (k == "FLOAT")) || (k == "STRING")) || (k == "true")) || (k == "false")) || (k == "self")) || (k == "Self")) {
			rr = _radv(rr)
			consumed_value = true
		} else if (k == "STRING_START") {
			rr = _radv(rr)
			interp_after_dot := false
			_ = interp_after_dot
			for ((_rk(rr) != "STRING_END") && (_rk(rr) != "EOF")) {
				ik := _rk(rr)
				_ = ik
				interp_handled := false
				_ = interp_handled
				if (ik == "STRING_MIDDLE") {
					rr = _radv(rr)
					interp_after_dot = false
					interp_handled = true
				}
				if ((interp_handled == false) && (ik == "IDENT")) {
					if interp_after_dot {
						rr = _radv(rr)
					} else {
						rr = _resolve_ident(rr)
					}
					interp_after_dot = false
					interp_handled = true
				}
				if ((interp_handled == false) && (ik == "fn")) {
					rr = _resolve_closure(rr)
					interp_after_dot = false
					interp_handled = true
				}
				if (interp_handled == false) {
					if (ik == ".") {
						interp_after_dot = true
					} else {
						interp_after_dot = false
					}
					rr = _radv(rr)
				}
			}
			if (_rk(rr) == "STRING_END") {
				rr = _radv(rr)
			}
			consumed_value = true
		} else if ((k == ")") || (k == "]")) {
			rr = _radv(rr)
			consumed_value = true
		} else if ((k == "?") && (_rk(_radv(rr)) == "|")) {
			rr = _radv(rr)
			rr = _radv(rr)
			if (_rk(rr) == "IDENT") {
				ename := _rcur(rr).Text
				_ = ename
				espan := _rspan(rr)
				_ = espan
				rr = _add_sym_warn_shadow(rr, ename, SymbolKindSkParameter{}, espan, false)
				rr = _radv(rr)
			} else {
				rr = _radv(rr)
			}
			if (_rk(rr) == "|") {
				rr = _radv(rr)
			}
			consumed_value = false
		} else if (k == "catch") {
			rr = _radv(rr)
			if (_rk(rr) == "|") {
				rr = _radv(rr)
				if (_rk(rr) == "IDENT") {
					ename := _rcur(rr).Text
					_ = ename
					espan := _rspan(rr)
					_ = espan
					rr = _add_sym_warn_shadow(rr, ename, SymbolKindSkParameter{}, espan, false)
					rr = _radv(rr)
				} else {
					rr = _radv(rr)
				}
				if (_rk(rr) == "|") {
					rr = _radv(rr)
				}
			}
			lb := "{"
			_ = lb
			if (_rk(rr) == lb) {
				rr = _resolve_block(rr)
			}
			consumed_value = true
		} else {
			if (k == ".") {
				after_dot = true
			} else {
				after_dot = false
			}
			rr = _radv(rr)
			if ((((((((((((((((((((((((((((k == "(") || (k == "[")) || (k == ",")) || (k == ".")) || (k == "+")) || (k == "-")) || (k == "*")) || (k == "/")) || (k == "%")) || (k == "==")) || (k == "!=")) || (k == "<")) || (k == ">")) || (k == "<=")) || (k == ">=")) || (k == "&&")) || (k == "||")) || (k == "&")) || (k == "|")) || (k == "^")) || (k == "~")) || (k == "<<")) || (k == ">>")) || (k == "|>")) || (k == "=>")) || (k == "=")) || (k == ":")) || (k == ":=")) {
				consumed_value = false
			}
		}
	}
	return rr
}

func _find_similar_name(r Resolver, name string) string {
	best := ""
	_ = best
	best_dist := int64(3)
	_ = best_dist
	si := int64(1)
	_ = si
	for (si < int64(len(r.Table.Symbols))) {
		sym := r.Table.Symbols[si]
		_ = sym
		if (sym.Name != "") {
			if (sym.Name != name) {
				dist := _edit_distance(name, sym.Name)
				_ = dist
				if (dist < best_dist) {
					best_dist = dist
					best = sym.Name
				}
				if ((dist == best_dist) && (best == "")) {
					best = sym.Name
				}
			}
		}
		si = (si + int64(1))
	}
	return best
}

func _edit_distance(a string, b string) int64 {
	al := int64(len(a))
	_ = al
	bl := int64(len(b))
	_ = bl
	if (al == int64(0)) {
		return bl
	}
	if (bl == int64(0)) {
		return al
	}
	if ((al > int64(20)) || (bl > int64(20))) {
		return int64(99)
	}
	common := int64(0)
	_ = common
	i := int64(0)
	_ = i
	min_len := al
	_ = min_len
	if (bl < min_len) {
		min_len = bl
	}
	for (i < min_len) {
		if (string(a[i]) == string(b[i])) {
			common = (common + int64(1))
		}
		i = (i + int64(1))
	}
	diff := ((al + bl) - (common * int64(2)))
	_ = diff
	if ((common * int64(3)) >= (min_len * int64(2))) {
		len_diff := (al - bl)
		_ = len_diff
		if (len_diff < int64(0)) {
			len_diff = (int64(0) - len_diff)
		}
		return (len_diff + ((min_len - common)))
	}
	return diff
}

func _stdlib_module_for_name(name string) string {
	if (((name == "json_parse") || (name == "json_emit")) || (name == "json_kind")) {
		return "json"
	}
	if (((name == "json_str_val") || (name == "json_int_val")) || (name == "json_bool_val")) {
		return "json"
	}
	if (((name == "json_array_len") || (name == "json_array_get")) || (name == "json_object_get")) {
		return "json"
	}
	if ((((name == "json_new_null") || (name == "json_new_bool")) || (name == "json_new_int")) || (name == "json_new_string")) {
		return "json"
	}
	if (((name == "json_get_str") || (name == "json_get_int")) || (name == "json_get_bool")) {
		return "json"
	}
	if (((name == "http_response") || (name == "http_response_json")) || (name == "http_add_header")) {
		return "http"
	}
	if (((name == "http_parse_request") || (name == "http_parse_query")) || (name == "http_get_param")) {
		return "http"
	}
	if (((name == "http_serialize_response") || (name == "HttpRequest")) || (name == "HttpResponse")) {
		return "http"
	}
	if ((((name == "db_connect") || (name == "db_connected")) || (name == "db_error")) || (name == "db_close")) {
		return "db"
	}
	if ((((name == "db_exec") || (name == "db_query")) || (name == "db_clear")) || (name == "db_ok")) {
		return "db"
	}
	if (((name == "db_result_error") || (name == "db_rows")) || (name == "db_field_name")) {
		return "db"
	}
	if ((((name == "db_get") || (name == "db_is_null")) || (name == "db_row")) || (name == "db_row_get")) {
		return "db"
	}
	return ""
}

func _resolve_ident(r Resolver) Resolver {
	name := _rcur(r).Text
	_ = name
	found := lookup(r.Table, r.Current_scope, name)
	_ = found
	if (found.Name == "") {
		q := _q()
		_ = q
		suggestion := _find_similar_name(r, name)
		_ = suggestion
		msg := ((("unresolved name " + q) + name) + q)
		_ = msg
		if (suggestion != "") {
			msg = (((((msg + ". Did you mean ") + q) + suggestion) + q) + "?")
		} else {
			sm := _stdlib_module_for_name(name)
			_ = sm
			if (sm != "") {
				msg = ((((((msg + ". Did you forget ") + q) + "use ") + sm) + q) + "?")
			}
		}
		rr := _rerror(r, E0701, msg)
		_ = rr
		return _radv(rr)
	}
	if (found.Is_public == false) {
		ffile := _span_file(found.Span)
		_ = ffile
		if (ffile != "") {
			if (ffile != "<builtin>") {
				if (ffile != r.File) {
					q := _q()
					_ = q
					rr := _rerror(r, E0704, ((((((("symbol " + q) + name) + q) + " is private to module ") + q) + ffile) + q))
					_ = rr
					return _radv(rr)
				}
			}
		}
	}
	return _radv(r)
}

func _resolve_ident_by_name(r Resolver, name string, span Span) Resolver {
	if (int64(len(name)) > int64(0)) {
		c0 := string(name[int64(0)])
		_ = c0
		is_upper := ((c0 >= "A") && (c0 <= "Z"))
		_ = is_upper
		if is_upper {
			ci := int64(1)
			_ = ci
			for (ci < int64(len(name))) {
				if (string(name[ci]) == "_") {
					return r
				}
				ci = (ci + int64(1))
			}
		}
	}
	found := lookup(r.Table, r.Current_scope, name)
	_ = found
	if (found.Name == "") {
		q := _q()
		_ = q
		suggestion := _find_similar_name(r, name)
		_ = suggestion
		msg := ((("unresolved name " + q) + name) + q)
		_ = msg
		if (suggestion != "") {
			msg = (((((msg + ". Did you mean ") + q) + suggestion) + q) + "?")
		} else {
			sm := _stdlib_module_for_name(name)
			_ = sm
			if (sm != "") {
				msg = ((((((msg + ". Did you forget ") + q) + "use ") + sm) + q) + "?")
			}
		}
		espan := span
		_ = espan
		if (espan.File == "") {
			espan = Span{File: r.File, Line: span.Line, Col: span.Col, Offset: span.Offset, Length: span.Length}
		}
		return _rset_diag(r, bag_add_error(r.Diagnostics, E0701, msg, espan))
	}
	if (found.Is_public == false) {
		ffile := _span_file(found.Span)
		_ = ffile
		if (ffile != "") {
			if (ffile != "<builtin>") {
				if (ffile != r.File) {
					q := _q()
					_ = q
					return _rset_diag(r, bag_add_error(r.Diagnostics, E0704, ((((((("symbol " + q) + name) + q) + " is private to module ") + q) + ffile) + q), span))
				}
			}
		}
	}
	return r
}

func _resolve_expr_node(r Resolver, idx int64) Resolver {
	if (idx <= int64(0)) {
		return r
	}
	node := pool_get(r.Pool, idx)
	_ = node
	nk := expr_kind_name(node.Kind)
	_ = nk
	if (((((((nk == "IntLit") || (nk == "FloatLit")) || (nk == "StringLit")) || (nk == "BoolLit")) || (nk == "None")) || (nk == "Break")) || (nk == "Continue")) {
		return r
	}
	if (nk == "Ident") {
		return _resolve_ident_by_name(r, node.S1, node.Span)
	}
	if (nk == "Binary") {
		if (node.S1 == ":=") {
			lhs := pool_get(r.Pool, node.C0)
			_ = lhs
			rr := _resolve_expr_node(r, node.C1)
			_ = rr
			if (expr_kind_name(lhs.Kind) == "Ident") {
				rr = _add_sym_warn_shadow(rr, lhs.S1, SymbolKindSkVariable{}, lhs.Span, false)
			}
			return rr
		}
		if (node.S1 == "=") {
			rr := _resolve_expr_node(r, node.C0)
			_ = rr
			return _resolve_expr_node(rr, node.C1)
		}
		rr := _resolve_expr_node(r, node.C0)
		_ = rr
		return _resolve_expr_node(rr, node.C1)
	}
	if ((nk == "Pipeline") || (nk == "Range")) {
		rr := _resolve_expr_node(r, node.C0)
		_ = rr
		return _resolve_expr_node(rr, node.C1)
	}
	if (((nk == "Unary") || (nk == "AssertOk")) || (nk == "Defer")) {
		return _resolve_expr_node(r, node.C0)
	}
	if (nk == "Propagate") {
		rr := _resolve_expr_node(r, node.C0)
		_ = rr
		if node.B1 {
			slen := int64(len(node.S1))
			_ = slen
			if ((slen > int64(0)) && (node.S1 != "_")) {
				rr = _add_sym_warn_shadow(rr, node.S1, SymbolKindSkParameter{}, node.Span, false)
			}
			rr = _resolve_expr_node(rr, node.C1)
		}
		return rr
	}
	if (nk == "Return") {
		if node.B1 {
			return _resolve_expr_node(r, node.C0)
		}
		return r
	}
	if (nk == "Assign") {
		rr := _resolve_expr_node(r, node.C0)
		_ = rr
		return _resolve_expr_node(rr, node.C1)
	}
	if (nk == "FieldAccess") {
		return _resolve_expr_node(r, node.C0)
	}
	if (nk == "Index") {
		rr := _resolve_expr_node(r, node.C0)
		_ = rr
		return _resolve_expr_node(rr, node.C1)
	}
	if (nk == "Call") {
		rr := _resolve_expr_node(r, node.C0)
		_ = rr
		ai := int64(0)
		_ = ai
		for (ai < node.List_count) {
			arg_idx := pool_get(r.Pool, (node.List_start + ai))
			_ = arg_idx
			rr = _resolve_expr_node(rr, arg_idx.C0)
			ai = (ai + int64(1))
		}
		return rr
	}
	if (nk == "MethodCall") {
		rr := _resolve_expr_node(r, node.C0)
		_ = rr
		ai := int64(0)
		_ = ai
		for (ai < node.List_count) {
			arg_idx := pool_get(r.Pool, (node.List_start + ai))
			_ = arg_idx
			rr = _resolve_expr_node(rr, arg_idx.C0)
			ai = (ai + int64(1))
		}
		return rr
	}
	if ((nk == "Array") || (nk == "Tuple")) {
		rr := r
		_ = rr
		ai := int64(0)
		_ = ai
		for (ai < node.List_count) {
			elem_idx := pool_get(r.Pool, (node.List_start + ai))
			_ = elem_idx
			rr = _resolve_expr_node(rr, elem_idx.C0)
			ai = (ai + int64(1))
		}
		return rr
	}
	if (nk == "Struct") {
		rr := r
		_ = rr
		if (node.S1 != "_update") {
			rr = _resolve_ident_by_name(r, node.S1, node.Span)
		}
		fi := int64(0)
		_ = fi
		for (fi < node.List_count) {
			fval_idx := pool_get(r.Pool, (node.List_start + fi))
			_ = fval_idx
			rr = _resolve_expr_node(rr, fval_idx.C0)
			fi = (fi + int64(1))
		}
		return rr
	}
	if (nk == "Block") {
		parent_scope := r.Current_scope
		_ = parent_scope
		rr := _push_scope(r, ScopeKindScBlock{}, "")
		_ = rr
		si := int64(0)
		_ = si
		for (si < node.List_count) {
			idx_node := pool_get(r.Pool, (node.List_start + si))
			_ = idx_node
			rr = _resolve_expr_node(rr, idx_node.C0)
			si = (si + int64(1))
		}
		return _pop_scope(rr, parent_scope)
	}
	if (nk == "If") {
		rr := _resolve_expr_node(r, node.C0)
		_ = rr
		rr = _resolve_expr_node(rr, node.C1)
		if node.B1 {
			rr = _resolve_expr_node(rr, node.C2)
		}
		return rr
	}
	if (nk == "While") {
		rr := _resolve_expr_node(r, node.C0)
		_ = rr
		return _resolve_expr_node(rr, node.C1)
	}
	if (nk == "Loop") {
		return _resolve_expr_node(r, node.C0)
	}
	if (nk == "For") {
		parent_scope := r.Current_scope
		_ = parent_scope
		rr := _push_scope(r, ScopeKindScBlock{}, "for")
		_ = rr
		rr = _add_sym_warn_shadow(rr, node.S1, SymbolKindSkVariable{}, node.Span, false)
		rr = _resolve_expr_node(rr, node.C0)
		rr = _resolve_expr_node(rr, node.C1)
		return _pop_scope(rr, parent_scope)
	}
	if (nk == "Binding") {
		rr := _resolve_expr_node(r, node.C0)
		_ = rr
		rr = _add_sym_warn_shadow(rr, node.S1, SymbolKindSkVariable{}, node.Span, node.B1)
		return rr
	}
	if (nk == "Closure") {
		parent_scope := r.Current_scope
		_ = parent_scope
		rr := _push_scope(r, ScopeKindScFunction{}, "closure")
		_ = rr
		if (int64(len(node.S1)) > int64(0)) {
			rr = _register_param_names(rr, node.S1, node.Span)
		}
		rr = _resolve_expr_node(rr, node.C0)
		return _pop_scope(rr, parent_scope)
	}
	if (nk == "Match") {
		rr := _resolve_expr_node(r, node.C0)
		_ = rr
		ai := int64(0)
		_ = ai
		for (ai < node.List_count) {
			parent_scope := rr.Current_scope
			_ = parent_scope
			rr = _push_scope(rr, ScopeKindScBlock{}, "match_arm")
			pat_idx_node := pool_get(r.Pool, (node.List_start + (ai * int64(2))))
			_ = pat_idx_node
			rr = _resolve_pattern_bindings(rr, pat_idx_node.C0)
			body_idx_node := pool_get(r.Pool, ((node.List_start + (ai * int64(2))) + int64(1)))
			_ = body_idx_node
			rr = _resolve_expr_node(rr, body_idx_node.C0)
			rr = _pop_scope(rr, parent_scope)
			ai = (ai + int64(1))
		}
		return rr
	}
	if (nk == "With") {
		rr := _resolve_expr_node(r, node.C0)
		_ = rr
		slen := int64(len(node.S1))
		_ = slen
		if (slen > int64(0)) {
			rr = _add_sym_warn_shadow(rr, node.S1, SymbolKindSkVariable{}, node.Span, false)
		}
		return _resolve_expr_node(rr, node.C1)
	}
	if (nk == "Catch") {
		rr := _resolve_expr_node(r, node.C0)
		_ = rr
		parent_scope := rr.Current_scope
		_ = parent_scope
		rr = _push_scope(rr, ScopeKindScBlock{}, "catch")
		slen := int64(len(node.S1))
		_ = slen
		if (slen > int64(0)) {
			rr = _add_sym_warn_shadow(rr, node.S1, SymbolKindSkParameter{}, node.Span, false)
		}
		rr = _resolve_expr_node(rr, node.C1)
		return _pop_scope(rr, parent_scope)
	}
	if (nk == "ListComp") {
		parent_scope := r.Current_scope
		_ = parent_scope
		rr := _push_scope(r, ScopeKindScBlock{}, "listcomp")
		_ = rr
		rr = _resolve_expr_node(rr, node.C1)
		rr = _add_sym_warn_shadow(rr, node.S1, SymbolKindSkVariable{}, node.Span, false)
		rr = _resolve_expr_node(rr, node.C0)
		if node.B1 {
			rr = _resolve_expr_node(rr, node.C2)
		}
		return _pop_scope(rr, parent_scope)
	}
	if (nk == "Yield") {
		if node.B1 {
			return _resolve_expr_node(r, node.C0)
		}
		return r
	}
	if (nk == "Select") {
		rr := r
		_ = rr
		ai := int64(0)
		_ = ai
		for (ai < node.List_count) {
			info_idx_node := pool_get(rr.Pool, (node.List_start + (ai * int64(2))))
			_ = info_idx_node
			body_idx_node := pool_get(rr.Pool, ((node.List_start + (ai * int64(2))) + int64(1)))
			_ = body_idx_node
			info := pool_get(rr.Pool, info_idx_node.C0)
			_ = info
			if (info.S1 != "_default") {
				rr = _resolve_expr_node(rr, info.C0)
			}
			parent_scope := rr.Current_scope
			_ = parent_scope
			rr = _push_scope(rr, ScopeKindScBlock{}, "select_arm")
			if ((info.S1 != "_default") && (info.S1 != "_")) {
				rr = _add_sym_warn_shadow(rr, info.S1, SymbolKindSkVariable{}, info.Span, false)
			}
			rr = _resolve_expr_node(rr, body_idx_node.C0)
			rr = _pop_scope(rr, parent_scope)
			ai = (ai + int64(1))
		}
		return rr
	}
	return r
}

func _register_param_names(r Resolver, params string, span Span) Resolver {
	rr := r
	_ = rr
	start := int64(0)
	_ = start
	plen := int64(len(params))
	_ = plen
	i := int64(0)
	_ = i
	for (i < plen) {
		ch := string(params[i])
		_ = ch
		if (ch == ",") {
			if (i > start) {
				name := params[start:i]
				_ = name
				rr = _add_sym_warn_shadow(rr, name, SymbolKindSkParameter{}, span, false)
			}
			start = (i + int64(1))
		}
		i = (i + int64(1))
	}
	if (start < plen) {
		name := params[start:plen]
		_ = name
		rr = _add_sym_warn_shadow(rr, name, SymbolKindSkParameter{}, span, false)
	}
	return rr
}

func _resolve_pattern_bindings(r Resolver, idx int64) Resolver {
	if (idx <= int64(0)) {
		return r
	}
	node := pool_get(r.Pool, idx)
	_ = node
	nk := expr_kind_name(node.Kind)
	_ = nk
	if (nk == "Ident") {
		name := node.S1
		_ = name
		if (name == "_") {
			return r
		}
		existing := lookup(r.Table, r.Current_scope, name)
		_ = existing
		if (existing.Name != "") {
			if ((sk_eq(existing.Kind, SymbolKindSkVariant{}) || sk_eq(existing.Kind, SymbolKindSkType{})) || sk_eq(existing.Kind, SymbolKindSkBuiltin{})) {
				return r
			}
		}
		return _add_sym_warn_shadow(r, name, SymbolKindSkVariable{}, node.Span, false)
	}
	if (nk == "Call") {
		rr := r
		_ = rr
		pi := int64(0)
		_ = pi
		for (pi < node.List_count) {
			rr = _resolve_pattern_bindings(rr, (node.List_start + pi))
			pi = (pi + int64(1))
		}
		return rr
	}
	return r
}

func _resolve_if(r Resolver) Resolver {
	rr := _radv(r)
	_ = rr
	lb := "{"
	_ = lb
	adot := false
	_ = adot
	for ((_rk(rr) != lb) && (_rk(rr) != "EOF")) {
		if (_rk(rr) == "IDENT") {
			if adot {
				rr = _radv(rr)
			} else {
				rr = _resolve_ident(rr)
			}
			adot = false
		} else {
			if (_rk(rr) == ".") {
				adot = true
			} else {
				adot = false
			}
			rr = _radv(rr)
		}
	}
	rr = _resolve_block(rr)
	if (_rk(rr) == "else") {
		rr = _radv(rr)
		if (_rk(rr) == "if") {
			rr = _resolve_if(rr)
		} else {
			rr = _resolve_block(rr)
		}
	}
	return rr
}

func _resolve_match(r Resolver) Resolver {
	rr := _radv(r)
	_ = rr
	lb := "{"
	_ = lb
	rb := "}"
	_ = rb
	adot := false
	_ = adot
	for ((_rk(rr) != lb) && (_rk(rr) != "EOF")) {
		if (_rk(rr) == "IDENT") {
			if adot {
				rr = _radv(rr)
			} else {
				rr = _resolve_ident(rr)
			}
			adot = false
		} else {
			if (_rk(rr) == ".") {
				adot = true
			} else {
				adot = false
			}
			rr = _radv(rr)
		}
	}
	if (_rk(rr) != lb) {
		return rr
	}
	rr = _radv(rr)
	rr = _rskip_nl(rr)
	for ((_rk(rr) != rb) && (_rk(rr) != "EOF")) {
		parent_scope := rr.Current_scope
		_ = parent_scope
		rr = _push_scope(rr, ScopeKindScBlock{}, "match_arm")
		if (_rk(rr) == "IDENT") {
			pat_name := _rcur(rr).Text
			_ = pat_name
			if (pat_name != "_") {
				existing := lookup(rr.Table, rr.Current_scope, pat_name)
				_ = existing
				is_variant := false
				_ = is_variant
				if (existing.Name != "") {
					is_variant = sk_eq(existing.Kind, SymbolKindSkVariant{})
				}
				if (is_variant == false) {
					span := _rspan(rr)
					_ = span
					rr = _add_sym_warn_shadow(rr, pat_name, SymbolKindSkVariable{}, span, false)
				}
			}
			rr = _radv(rr)
		} else if ((((_rk(rr) == "INT") || (_rk(rr) == "STRING")) || (_rk(rr) == "true")) || (_rk(rr) == "false")) {
			rr = _radv(rr)
		}
		for (_rk(rr) == "|") {
			rr = _radv(rr)
			rr = _rskip_nl(rr)
			if (_rk(rr) == "IDENT") {
				rr = _radv(rr)
			}
			rr = _rskip_nl(rr)
		}
		if (_rk(rr) == "(") {
			rr = _radv(rr)
			rr = _rskip_nl(rr)
			for ((_rk(rr) != ")") && (_rk(rr) != "EOF")) {
				if (_rk(rr) == "IDENT") {
					bname := _rcur(rr).Text
					_ = bname
					span := _rspan(rr)
					_ = span
					rr = _add_sym_warn_shadow(rr, bname, SymbolKindSkVariable{}, span, false)
					rr = _radv(rr)
				} else {
					rr = _radv(rr)
				}
				rr = _rskip_nl(rr)
				if (_rk(rr) == ",") {
					rr = _radv(rr)
				}
				rr = _rskip_nl(rr)
			}
			if (_rk(rr) == ")") {
				rr = _radv(rr)
			}
		}
		if (_rk(rr) == "if") {
			rr = _radv(rr)
			for ((_rk(rr) != "=>") && (_rk(rr) != "EOF")) {
				if (_rk(rr) == "IDENT") {
					rr = _resolve_ident(rr)
				} else {
					rr = _radv(rr)
				}
			}
		}
		for ((_rk(rr) != "=>") && (_rk(rr) != "EOF")) {
			rr = _radv(rr)
		}
		if (_rk(rr) == "=>") {
			rr = _radv(rr)
		}
		rr = _rskip_nl(rr)
		if (_rk(rr) == lb) {
			rr = _resolve_block(rr)
		} else {
			rr = _resolve_expr_until_nl(rr)
		}
		rr = _pop_scope(rr, parent_scope)
		rr = _rskip_nl(rr)
	}
	if (_rk(rr) == rb) {
		rr = _radv(rr)
	}
	return rr
}

func _resolve_for(r Resolver) Resolver {
	parent_scope := r.Current_scope
	_ = parent_scope
	rr := _radv(r)
	_ = rr
	rr = _push_scope(rr, ScopeKindScBlock{}, "for")
	if (_rk(rr) == "IDENT") {
		name := _rcur(rr).Text
		_ = name
		span := _rspan(rr)
		_ = span
		rr = _add_sym_warn_shadow(rr, name, SymbolKindSkVariable{}, span, false)
		rr = _radv(rr)
	}
	if (_rk(rr) == "in") {
		rr = _radv(rr)
	}
	lb := "{"
	_ = lb
	for ((_rk(rr) != lb) && (_rk(rr) != "EOF")) {
		if (_rk(rr) == "IDENT") {
			rr = _resolve_ident(rr)
		} else {
			rr = _radv(rr)
		}
	}
	rr = _resolve_block(rr)
	return _pop_scope(rr, parent_scope)
}

func _resolve_while(r Resolver) Resolver {
	rr := _radv(r)
	_ = rr
	lb := "{"
	_ = lb
	adot := false
	_ = adot
	for ((_rk(rr) != lb) && (_rk(rr) != "EOF")) {
		if (_rk(rr) == "IDENT") {
			if adot {
				rr = _radv(rr)
			} else {
				rr = _resolve_ident(rr)
			}
			adot = false
		} else {
			if (_rk(rr) == ".") {
				adot = true
			} else {
				adot = false
			}
			rr = _radv(rr)
		}
	}
	return _resolve_block(rr)
}

func _resolve_loop(r Resolver) Resolver {
	rr := _radv(r)
	_ = rr
	return _resolve_block(rr)
}

func _resolve_with(r Resolver) Resolver {
	rr := _radv(r)
	_ = rr
	res_name := ""
	_ = res_name
	res_span := _rspan(rr)
	_ = res_span
	if (_rk(rr) == "IDENT") {
		next := (rr.Pos + int64(1))
		_ = next
		if ((next < int64(len(rr.Tokens))) && (token_name(rr.Tokens[next].Kind) == ":=")) {
			res_name = _rcur(rr).Text
			res_span = _rspan(rr)
			rr = _radv(rr)
			rr = _radv(rr)
		}
	}
	lb := "{"
	_ = lb
	for (((_rk(rr) != lb) && (_rk(rr) != "NEWLINE")) && (_rk(rr) != "EOF")) {
		if (_rk(rr) == "IDENT") {
			rr = _resolve_ident(rr)
		} else {
			rr = _radv(rr)
		}
	}
	rr = _rskip_nl(rr)
	if (int64(len(res_name)) > int64(0)) {
		rr = _add_sym_warn_shadow(rr, res_name, SymbolKindSkVariable{}, res_span, false)
	}
	return _resolve_block(rr)
}

func _resolve_closure(r Resolver) Resolver {
	parent_scope := r.Current_scope
	_ = parent_scope
	rr := _radv(r)
	_ = rr
	if ((_rk(rr) == "IDENT") && (_rcur(rr).Text == "once")) {
		rr = _radv(rr)
	}
	rr = _push_scope(rr, ScopeKindScFunction{}, "closure")
	if (_rk(rr) == "(") {
		rr = _radv(rr)
		for ((_rk(rr) != ")") && (_rk(rr) != "EOF")) {
			if (_rk(rr) == "IDENT") {
				pname := _rcur(rr).Text
				_ = pname
				pspan := _rspan(rr)
				_ = pspan
				rr = _add_sym_warn_shadow(rr, pname, SymbolKindSkParameter{}, pspan, false)
				rr = _radv(rr)
				if (_rk(rr) == ":") {
					rr = _radv(rr)
					rr = _skip_type_tokens(rr)
				}
			}
			if (_rk(rr) == ",") {
				rr = _radv(rr)
			}
		}
		if (_rk(rr) == ")") {
			rr = _radv(rr)
		}
	}
	if (_rk(rr) == "->") {
		rr = _radv(rr)
		rr = _skip_type_tokens(rr)
	}
	if (_rk(rr) == "=>") {
		rr = _radv(rr)
		rr = _resolve_expr_until_nl(rr)
	}
	return _pop_scope(rr, parent_scope)
}

func _skip_type_tokens(r Resolver) Resolver {
	rr := r
	_ = rr
	k := _rk(rr)
	_ = k
	if ((k == "IDENT") || (k == "Self")) {
		rr = _radv(rr)
		for (_rk(rr) == ".") {
			rr = _radv(_radv(rr))
		}
		if (_rk(rr) == "[") {
			d := int64(1)
			_ = d
			rr = _radv(rr)
			for ((d > int64(0)) && (_rk(rr) != "EOF")) {
				if (_rk(rr) == "[") {
					d = (d + int64(1))
				}
				if (_rk(rr) == "]") {
					d = (d - int64(1))
				}
				if (d > int64(0)) {
					rr = _radv(rr)
				}
			}
			if (_rk(rr) == "]") {
				rr = _radv(rr)
			}
		}
		if (_rk(rr) == "?") {
			rr = _radv(rr)
		}
		if (_rk(rr) == "!") {
			rr = _radv(rr)
			rr = _skip_type_tokens(rr)
		}
		return rr
	}
	if (k == "[") {
		rr = _radv(rr)
		rr = _skip_type_tokens(rr)
		if (_rk(rr) == "]") {
			rr = _radv(rr)
		}
		return rr
	}
	if (k == "(") {
		rr = _radv(rr)
		for ((_rk(rr) != ")") && (_rk(rr) != "EOF")) {
			rr = _skip_type_tokens(rr)
			if (_rk(rr) == ",") {
				rr = _radv(rr)
			}
		}
		if (_rk(rr) == ")") {
			rr = _radv(rr)
		}
		if (_rk(rr) == "->") {
			rr = _radv(rr)
			rr = _skip_type_tokens(rr)
		}
		return rr
	}
	if (k == "fn") {
		rr = _radv(rr)
		if (_rk(rr) == "(") {
			rr = _radv(rr)
			for ((_rk(rr) != ")") && (_rk(rr) != "EOF")) {
				rr = _skip_type_tokens(rr)
				if (_rk(rr) == ",") {
					rr = _radv(rr)
				}
			}
			if (_rk(rr) == ")") {
				rr = _radv(rr)
			}
		}
		if (_rk(rr) == "->") {
			rr = _radv(rr)
			rr = _skip_type_tokens(rr)
		}
		return rr
	}
	return rr
}

func resolve(tokens []Token, index DeclIndex, pool NodePool, file string) ResolveResult {
	r := _new_resolver(tokens, pool, index, file)
	_ = r
	r = _push_scope(r, ScopeKindScUniverse{}, "universe")
	r = _populate_builtins(r)
	r = _push_scope(r, ScopeKindScModule{}, "module")
	ui := int64(1)
	_ = ui
	for (ui < int64(len(r.Tokens))) {
		if (token_name(r.Tokens[ui].Kind) == "use") {
			next := (ui + int64(1))
			_ = next
			if ((next < int64(len(r.Tokens))) && (token_name(r.Tokens[next].Kind) == "IDENT")) {
				mod_name := r.Tokens[next].Text
				_ = mod_name
				existing := lookup_local(r.Table, r.Current_scope, mod_name)
				_ = existing
				if (existing.Name == "") {
					sym := new_symbol(mod_name, SymbolKindSkModule{}, _builtin_span(), r.Current_scope, false, false)
					_ = sym
					r = _rset_table(r, add_symbol(r.Table, sym))
				}
			}
		}
		ui = (ui + int64(1))
	}
	r = _pass1(r, index)
	r = _pass2(r, index)
	return ResolveResult{Table: r.Table, Diagnostics: r.Diagnostics}
}

