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

type Parser struct {
	Tokens []Token
	Pos int64
	Pool NodePool
	Diagnostics DiagnosticBag
}

type PR struct {
	P Parser
	E Expr
}

func new_parser(tokens []Token) Parser {
	return Parser{Tokens: tokens, Pos: int64(0), Pool: new_pool(), Diagnostics: new_bag()}
}

func _pidx(p Parser) int64 {
	return pool_size(p.Pool)
}

func _padd(p Parser, e Expr) Parser {
	return Parser{Tokens: p.Tokens, Pos: p.Pos, Pool: pool_add(p.Pool, e), Diagnostics: p.Diagnostics}
}

func _cur(p Parser) Token {
	if (p.Pos >= int64(len(p.Tokens))) {
		return Token{Kind: TokenKindTkEof{}, Text: "", Line: int64(0), Col: int64(0), Offset: int64(0)}
	}
	return p.Tokens[p.Pos]
}

func _ck(p Parser) string {
	return token_name(_cur(p).Kind)
}

func _pk(p Parser) string {
	next := (p.Pos + int64(1))
	_ = next
	if (next >= int64(len(p.Tokens))) {
		return "EOF"
	}
	return token_name(p.Tokens[next].Kind)
}

func _pspan(p Parser) Span {
	tok := _cur(p)
	_ = tok
	return new_span("", tok.Line, tok.Col, tok.Offset, int64(len(tok.Text)))
}

func _adv(p Parser) Parser {
	return Parser{Tokens: p.Tokens, Pos: (p.Pos + int64(1)), Pool: p.Pool, Diagnostics: p.Diagnostics}
}

func _expect(p Parser, expected string) Parser {
	if (_ck(p) == expected) {
		return _adv(p)
	}
	tok := _cur(p)
	_ = tok
	span := new_span("", tok.Line, tok.Col, tok.Offset, int64(len(tok.Text)))
	_ = span
	return Parser{Tokens: p.Tokens, Pos: p.Pos, Pool: p.Pool, Diagnostics: bag_add_error(p.Diagnostics, E0001, ((("expected " + expected) + ", got ") + _ck(p)), span)}
}

func _skip_nl(p Parser) Parser {
	pp := p
	_ = pp
	for (_ck(pp) == "NEWLINE") {
		pp = _adv(pp)
	}
	return pp
}

func _perror(p Parser, code string, msg string) Parser {
	tok := _cur(p)
	_ = tok
	span := new_span("", tok.Line, tok.Col, tok.Offset, int64(len(tok.Text)))
	_ = span
	return Parser{Tokens: p.Tokens, Pos: p.Pos, Pool: p.Pool, Diagnostics: bag_add_error(p.Diagnostics, code, msg, span)}
}

func _is_upper_start(name string) bool {
	nlen := int64(len(name))
	_ = nlen
	if (nlen == int64(0)) {
		return false
	}
	ch := string(name[int64(0)])
	_ = ch
	is_upper := ((((((((((((((((((((((((((ch == "A") || (ch == "B")) || (ch == "C")) || (ch == "D")) || (ch == "E")) || (ch == "F")) || (ch == "G")) || (ch == "H")) || (ch == "I")) || (ch == "J")) || (ch == "K")) || (ch == "L")) || (ch == "M")) || (ch == "N")) || (ch == "O")) || (ch == "P")) || (ch == "Q")) || (ch == "R")) || (ch == "S")) || (ch == "T")) || (ch == "U")) || (ch == "V")) || (ch == "W")) || (ch == "X")) || (ch == "Y")) || (ch == "Z"))
	_ = is_upper
	if (is_upper == false) {
		return false
	}
	i := int64(1)
	_ = i
	for (i < nlen) {
		c := string(name[i])
		_ = c
		is_lower := ((((((((((((((((((((((((((c == "a") || (c == "b")) || (c == "c")) || (c == "d")) || (c == "e")) || (c == "f")) || (c == "g")) || (c == "h")) || (c == "i")) || (c == "j")) || (c == "k")) || (c == "l")) || (c == "m")) || (c == "n")) || (c == "o")) || (c == "p")) || (c == "q")) || (c == "r")) || (c == "s")) || (c == "t")) || (c == "u")) || (c == "v")) || (c == "w")) || (c == "x")) || (c == "y")) || (c == "z"))
		_ = is_lower
		if is_lower {
			return true
		}
		i = (i + int64(1))
	}
	j := int64(0)
	_ = j
	for (j < nlen) {
		if (string(name[j]) == "_") {
			return false
		}
		j = (j + int64(1))
	}
	if (nlen >= int64(2)) {
		return true
	}
	return false
}

func parse(tokens []Token) PR {
	p := new_parser(tokens)
	_ = p
	if (((_ck(p) == "EOF") && (p.Pos == int64(0))) && (int64(len(p.Tokens)) > int64(1))) {
		p = _adv(p)
	}
	p = _skip_nl(p)
	if (_ck(p) == "mod") {
		p = _adv(p)
		p = _adv(p)
		p = _skip_nl(p)
	}
	for (_ck(p) == "use") {
		p = _adv(p)
		if (_ck(p) == "IDENT") {
			p = _adv(p)
		}
		for ((_ck(p) != "NEWLINE") && (_ck(p) != "EOF")) {
			p = _adv(p)
		}
		p = _skip_nl(p)
	}
	decl_indices := []int64{int64(0)}
	_ = decl_indices
	count := int64(0)
	_ = count
	for (_ck(p) != "EOF") {
		p = _skip_nl(p)
		if (_ck(p) == "EOF") {
			list_start := _pidx(p)
			_ = list_start
			di := int64(0)
			_ = di
			for (di < count) {
				p = _padd(p, Expr{Kind: ExprKindEkNone{}, S1: "", S2: "", B1: false, B2: false, C0: decl_indices[(di + int64(1))], C1: int64(0), C2: int64(0), List_start: int64(0), List_count: int64(0), Span: _no_span()})
				di = (di + int64(1))
			}
			return PR{P: p, E: Expr{Kind: ExprKindEkBlock{}, S1: "", S2: "", B1: false, B2: false, C0: int64(0), C1: int64(0), C2: int64(0), List_start: list_start, List_count: count, Span: _no_span()}}
		}
		result := _parse_top(p)
		_ = result
		decl_idx := _pidx(result.P)
		_ = decl_idx
		p = _padd(result.P, result.E)
		decl_indices = append(decl_indices, decl_idx)
		count = (count + int64(1))
		p = _skip_nl(p)
	}
	list_start := _pidx(p)
	_ = list_start
	di := int64(0)
	_ = di
	for (di < count) {
		p = _padd(p, Expr{Kind: ExprKindEkNone{}, S1: "", S2: "", B1: false, B2: false, C0: decl_indices[(di + int64(1))], C1: int64(0), C2: int64(0), List_start: int64(0), List_count: int64(0), Span: _no_span()})
		di = (di + int64(1))
	}
	return PR{P: p, E: Expr{Kind: ExprKindEkBlock{}, S1: "", S2: "", B1: false, B2: false, C0: int64(0), C1: int64(0), C2: int64(0), List_start: list_start, List_count: count, Span: _no_span()}}
}

func _is_fn_name_token(tk string) bool {
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

func build_decl_index(tokens []Token) ParseResult {
	idx := new_decl_index()
	_ = idx
	pos := int64(0)
	_ = pos
	len := int64(len(tokens))
	_ = len
	if ((pos < len) && (token_name(tokens[pos].Kind) == "EOF")) {
		pos = (pos + int64(1))
	}
	for ((pos < len) && (token_name(tokens[pos].Kind) == "NEWLINE")) {
		pos = (pos + int64(1))
	}
	if ((pos < len) && (token_name(tokens[pos].Kind) == "mod")) {
		pos = (pos + int64(1))
		if ((pos < len) && (token_name(tokens[pos].Kind) == "IDENT")) {
			pos = (pos + int64(1))
		}
		for ((pos < len) && (token_name(tokens[pos].Kind) == "NEWLINE")) {
			pos = (pos + int64(1))
		}
	}
	for ((pos < len) && (token_name(tokens[pos].Kind) == "use")) {
		use_start := pos
		_ = use_start
		name := ""
		_ = name
		pos = (pos + int64(1))
		if ((pos < len) && (token_name(tokens[pos].Kind) == "IDENT")) {
			name = tokens[pos].Text
		}
		for (((pos < len) && (token_name(tokens[pos].Kind) != "NEWLINE")) && (token_name(tokens[pos].Kind) != "EOF")) {
			pos = (pos + int64(1))
		}
		use_end := pos
		_ = use_end
		idx = decl_index_add(idx, DeclInfo{Kind: DeclKindDkUse{}, Name: name, Token_start: use_start, Body_start: use_start, Body_end: use_end, Is_pub: false, Node_idx: int64(0)})
		for ((pos < len) && (token_name(tokens[pos].Kind) == "NEWLINE")) {
			pos = (pos + int64(1))
		}
	}
	for ((pos < len) && (token_name(tokens[pos].Kind) != "EOF")) {
		for ((pos < len) && (token_name(tokens[pos].Kind) == "NEWLINE")) {
			pos = (pos + int64(1))
		}
		if ((pos >= len) || (token_name(tokens[pos].Kind) == "EOF")) {
			return ParseResult{Index: idx, Pool: new_pool(), Diagnostics: new_bag()}
		}
		is_pub := false
		_ = is_pub
		if (token_name(tokens[pos].Kind) == "pub") {
			is_pub = true
			pos = (pos + int64(1))
			if ((pos < len) && (token_name(tokens[pos].Kind) == "(")) {
				pos = (pos + int64(1))
				if ((pos < len) && (token_name(tokens[pos].Kind) == "IDENT")) {
					pos = (pos + int64(1))
				}
				if ((pos < len) && (token_name(tokens[pos].Kind) == ")")) {
					pos = (pos + int64(1))
				}
			}
		}
		if ((pos < len) && (token_name(tokens[pos].Kind) == "@")) {
			pos = (pos + int64(1))
			if ((pos < len) && (token_name(tokens[pos].Kind) == "IDENT")) {
				pos = (pos + int64(1))
			}
			if ((pos < len) && (token_name(tokens[pos].Kind) == "(")) {
				pos = (pos + int64(1))
				for (((pos < len) && (token_name(tokens[pos].Kind) != ")")) && (token_name(tokens[pos].Kind) != "EOF")) {
					pos = (pos + int64(1))
				}
				if ((pos < len) && (token_name(tokens[pos].Kind) == ")")) {
					pos = (pos + int64(1))
				}
			}
			for ((pos < len) && (token_name(tokens[pos].Kind) == "NEWLINE")) {
				pos = (pos + int64(1))
			}
		}
		if ((pos >= len) || (token_name(tokens[pos].Kind) == "EOF")) {
			return ParseResult{Index: idx, Pool: new_pool(), Diagnostics: new_bag()}
		}
		k := token_name(tokens[pos].Kind)
		_ = k
		lb := "{"
		_ = lb
		rb := "}"
		_ = rb
		if (k == "fn") {
			decl_start := pos
			_ = decl_start
			pos = (pos + int64(1))
			name := ""
			_ = name
			if ((pos < len) && _is_fn_name_token(token_name(tokens[pos].Kind))) {
				name = tokens[pos].Text
			}
			pos = _bdi_find_fn_body(tokens, pos, len)
			body_start := pos
			_ = body_start
			if ((pos < len) && (token_name(tokens[pos].Kind) == "=")) {
				pos = (pos + int64(1))
				pos = _bdi_skip_expr(tokens, pos, len)
			} else if ((pos < len) && (token_name(tokens[pos].Kind) == lb)) {
				pos = _bdi_skip_braces(tokens, pos, len)
			}
			idx = decl_index_add(idx, DeclInfo{Kind: DeclKindDkFn{}, Name: name, Token_start: decl_start, Body_start: body_start, Body_end: pos, Is_pub: is_pub, Node_idx: int64(0)})
		} else if ((k == "type") || (k == "struct")) {
			decl_start := pos
			_ = decl_start
			pos = (pos + int64(1))
			name := ""
			_ = name
			if ((pos < len) && (token_name(tokens[pos].Kind) == "IDENT")) {
				name = tokens[pos].Text
				pos = (pos + int64(1))
			}
			if ((pos < len) && (token_name(tokens[pos].Kind) == "[")) {
				pos = _bdi_skip_brackets(tokens, pos, len)
			}
			body_start := pos
			_ = body_start
			if ((pos < len) && (token_name(tokens[pos].Kind) == "=")) {
				pos = (pos + int64(1))
				for ((pos < len) && (token_name(tokens[pos].Kind) == "NEWLINE")) {
					pos = (pos + int64(1))
				}
				if ((pos < len) && (token_name(tokens[pos].Kind) == "|")) {
					for ((pos < len) && (token_name(tokens[pos].Kind) == "|")) {
						pos = (pos + int64(1))
						for ((pos < len) && (token_name(tokens[pos].Kind) == "NEWLINE")) {
							pos = (pos + int64(1))
						}
						if (pos < len) {
							pos = (pos + int64(1))
						}
						if ((pos < len) && (token_name(tokens[pos].Kind) == "(")) {
							pos = _bdi_skip_parens(tokens, pos, len)
						} else if ((pos < len) && (token_name(tokens[pos].Kind) == lb)) {
							pos = _bdi_skip_braces(tokens, pos, len)
						}
						for ((pos < len) && (token_name(tokens[pos].Kind) == "NEWLINE")) {
							pos = (pos + int64(1))
						}
					}
				} else {
					for (((pos < len) && (token_name(tokens[pos].Kind) != "NEWLINE")) && (token_name(tokens[pos].Kind) != "EOF")) {
						pos = (pos + int64(1))
					}
				}
			} else if ((pos < len) && (token_name(tokens[pos].Kind) == lb)) {
				pos = _bdi_skip_braces(tokens, pos, len)
			}
			for ((pos < len) && (token_name(tokens[pos].Kind) == "NEWLINE")) {
				pos = (pos + int64(1))
			}
			if ((pos < len) && (token_name(tokens[pos].Kind) == "derives")) {
				for (((pos < len) && (token_name(tokens[pos].Kind) != "NEWLINE")) && (token_name(tokens[pos].Kind) != "EOF")) {
					pos = (pos + int64(1))
				}
			}
			idx = decl_index_add(idx, DeclInfo{Kind: DeclKindDkType{}, Name: name, Token_start: decl_start, Body_start: body_start, Body_end: pos, Is_pub: is_pub, Node_idx: int64(0)})
		} else if (k == "enum") {
			decl_start := pos
			_ = decl_start
			pos = (pos + int64(1))
			name := ""
			_ = name
			if ((pos < len) && (token_name(tokens[pos].Kind) == "IDENT")) {
				name = tokens[pos].Text
				pos = (pos + int64(1))
			}
			body_start := pos
			_ = body_start
			if ((pos < len) && (token_name(tokens[pos].Kind) == lb)) {
				pos = _bdi_skip_braces(tokens, pos, len)
			}
			idx = decl_index_add(idx, DeclInfo{Kind: DeclKindDkEnum{}, Name: name, Token_start: decl_start, Body_start: body_start, Body_end: pos, Is_pub: is_pub, Node_idx: int64(0)})
		} else if (k == "trait") {
			decl_start := pos
			_ = decl_start
			pos = (pos + int64(1))
			name := ""
			_ = name
			if ((pos < len) && (token_name(tokens[pos].Kind) == "IDENT")) {
				name = tokens[pos].Text
			}
			for (((pos < len) && (token_name(tokens[pos].Kind) != lb)) && (token_name(tokens[pos].Kind) != "EOF")) {
				pos = (pos + int64(1))
			}
			body_start := pos
			_ = body_start
			if ((pos < len) && (token_name(tokens[pos].Kind) == lb)) {
				pos = _bdi_skip_braces(tokens, pos, len)
			}
			idx = decl_index_add(idx, DeclInfo{Kind: DeclKindDkTrait{}, Name: name, Token_start: decl_start, Body_start: body_start, Body_end: pos, Is_pub: is_pub, Node_idx: int64(0)})
		} else if (k == "impl") {
			decl_start := pos
			_ = decl_start
			pos = (pos + int64(1))
			name := ""
			_ = name
			if ((pos < len) && (token_name(tokens[pos].Kind) == "IDENT")) {
				name = tokens[pos].Text
			}
			for (((pos < len) && (token_name(tokens[pos].Kind) != lb)) && (token_name(tokens[pos].Kind) != "EOF")) {
				pos = (pos + int64(1))
			}
			body_start := pos
			_ = body_start
			if ((pos < len) && (token_name(tokens[pos].Kind) == lb)) {
				pos = _bdi_skip_braces(tokens, pos, len)
			}
			idx = decl_index_add(idx, DeclInfo{Kind: DeclKindDkImpl{}, Name: name, Token_start: decl_start, Body_start: body_start, Body_end: pos, Is_pub: is_pub, Node_idx: int64(0)})
		} else if (k == "const") {
			decl_start := pos
			_ = decl_start
			pos = (pos + int64(1))
			name := ""
			_ = name
			if ((pos < len) && (token_name(tokens[pos].Kind) == "IDENT")) {
				name = tokens[pos].Text
			}
			for (((pos < len) && (token_name(tokens[pos].Kind) != "NEWLINE")) && (token_name(tokens[pos].Kind) != "EOF")) {
				pos = (pos + int64(1))
			}
			idx = decl_index_add(idx, DeclInfo{Kind: DeclKindDkConst{}, Name: name, Token_start: decl_start, Body_start: decl_start, Body_end: pos, Is_pub: is_pub, Node_idx: int64(0)})
		} else if (k == "entry") {
			decl_start := pos
			_ = decl_start
			pos = (pos + int64(1))
			for ((pos < len) && (token_name(tokens[pos].Kind) == "NEWLINE")) {
				pos = (pos + int64(1))
			}
			body_start := pos
			_ = body_start
			if ((pos < len) && (token_name(tokens[pos].Kind) == lb)) {
				pos = _bdi_skip_braces(tokens, pos, len)
			}
			idx = decl_index_add(idx, DeclInfo{Kind: DeclKindDkEntry{}, Name: "", Token_start: decl_start, Body_start: body_start, Body_end: pos, Is_pub: false, Node_idx: int64(0)})
		} else if (k == "test") {
			decl_start := pos
			_ = decl_start
			pos = (pos + int64(1))
			name := ""
			_ = name
			if ((pos < len) && (((token_name(tokens[pos].Kind) == "STRING") || (token_name(tokens[pos].Kind) == "IDENT")))) {
				name = tokens[pos].Text
				pos = (pos + int64(1))
			}
			for ((pos < len) && (token_name(tokens[pos].Kind) == "NEWLINE")) {
				pos = (pos + int64(1))
			}
			body_start := pos
			_ = body_start
			if ((pos < len) && (token_name(tokens[pos].Kind) == lb)) {
				pos = _bdi_skip_braces(tokens, pos, len)
			}
			idx = decl_index_add(idx, DeclInfo{Kind: DeclKindDkTest{}, Name: name, Token_start: decl_start, Body_start: body_start, Body_end: pos, Is_pub: false, Node_idx: int64(0)})
		} else if (k == "bench") {
			decl_start := pos
			_ = decl_start
			pos = (pos + int64(1))
			name := ""
			_ = name
			if ((pos < len) && (((token_name(tokens[pos].Kind) == "STRING") || (token_name(tokens[pos].Kind) == "IDENT")))) {
				name = tokens[pos].Text
				pos = (pos + int64(1))
			}
			for ((pos < len) && (token_name(tokens[pos].Kind) == "NEWLINE")) {
				pos = (pos + int64(1))
			}
			body_start := pos
			_ = body_start
			if ((pos < len) && (token_name(tokens[pos].Kind) == lb)) {
				pos = _bdi_skip_braces(tokens, pos, len)
			}
			idx = decl_index_add(idx, DeclInfo{Kind: DeclKindDkTest{}, Name: ("_bench:" + name), Token_start: decl_start, Body_start: body_start, Body_end: pos, Is_pub: false, Node_idx: int64(0)})
		} else if (k == "extern") {
			decl_start := pos
			_ = decl_start
			pos = (pos + int64(1))
			if ((pos < len) && (token_name(tokens[pos].Kind) == "STRING")) {
				pos = (pos + int64(1))
			}
			body_start := pos
			_ = body_start
			if ((pos < len) && (token_name(tokens[pos].Kind) == lb)) {
				pos = _bdi_skip_braces(tokens, pos, len)
			} else if ((pos < len) && (token_name(tokens[pos].Kind) == "fn")) {
				for (((pos < len) && (token_name(tokens[pos].Kind) != "NEWLINE")) && (token_name(tokens[pos].Kind) != "EOF")) {
					pos = (pos + int64(1))
				}
			}
			idx = decl_index_add(idx, DeclInfo{Kind: DeclKindDkFn{}, Name: "_extern", Token_start: decl_start, Body_start: body_start, Body_end: pos, Is_pub: is_pub, Node_idx: int64(0)})
		} else if (((k == "IDENT") && (pos < len)) && (tokens[pos].Text == "alias")) {
			decl_start := pos
			_ = decl_start
			pos = (pos + int64(1))
			name := ""
			_ = name
			if ((pos < len) && (token_name(tokens[pos].Kind) == "IDENT")) {
				name = tokens[pos].Text
			}
			for (((pos < len) && (token_name(tokens[pos].Kind) != "NEWLINE")) && (token_name(tokens[pos].Kind) != "EOF")) {
				pos = (pos + int64(1))
			}
			idx = decl_index_add(idx, DeclInfo{Kind: DeclKindDkType{}, Name: ("_alias:" + name), Token_start: decl_start, Body_start: decl_start, Body_end: pos, Is_pub: is_pub, Node_idx: int64(0)})
		} else if (((k == "before") || (k == "after"))) {
			decl_start := pos
			_ = decl_start
			fixture_kind := k
			_ = fixture_kind
			probe := (pos + int64(1))
			_ = probe
			for ((probe < len) && (token_name(tokens[probe].Kind) == "NEWLINE")) {
				probe = (probe + int64(1))
			}
			if ((probe < len) && (token_name(tokens[probe].Kind) == lb)) {
				pos = (pos + int64(1))
				for ((pos < len) && (token_name(tokens[pos].Kind) == "NEWLINE")) {
					pos = (pos + int64(1))
				}
				body_start := pos
				_ = body_start
				pos = _bdi_skip_braces(tokens, pos, len)
				idx = decl_index_add(idx, DeclInfo{Kind: DeclKindDkFn{}, Name: ("_fixture:" + fixture_kind), Token_start: decl_start, Body_start: body_start, Body_end: pos, Is_pub: false, Node_idx: int64(0)})
			} else {
				pos = (pos + int64(1))
			}
		} else {
			pos = (pos + int64(1))
		}
		for ((pos < len) && (token_name(tokens[pos].Kind) == "NEWLINE")) {
			pos = (pos + int64(1))
		}
	}
	pr := parse(tokens)
	_ = pr
	result := _bridge_ast_to_decl_index(idx, pr.P.Pool, pr.E)
	_ = result
	return ParseResult{Index: result, Pool: pr.P.Pool, Diagnostics: new_bag()}
}

func _bridge_ast_to_decl_index(idx DeclIndex, pool NodePool, root Expr) DeclIndex {
	out := idx
	_ = out
	ni := int64(0)
	_ = ni
	for (ni < root.List_count) {
		idx_node := pool_get(pool, (root.List_start + ni))
		_ = idx_node
		actual_idx := idx_node.C0
		_ = actual_idx
		node := pool_get(pool, actual_idx)
		_ = node
		nk := expr_kind_name(node.Kind)
		_ = nk
		node_name := ""
		_ = node_name
		is_decl := false
		_ = is_decl
		if ((((((nk == "FnDecl") || (nk == "ConstDecl")) || (nk == "TypeDecl")) || (nk == "EnumDecl")) || (nk == "TraitDecl")) || (nk == "ImplDecl")) {
			node_name = node.S1
			is_decl = true
		} else if ((nk == "Block") && (node.S1 == "_entry")) {
			node_name = ""
			is_decl = true
		} else if ((nk == "Block") && strings.HasPrefix(node.S1, "_test:")) {
			slen := int64(len(node.S1))
			_ = slen
			node_name = node.S1[int64(6):slen]
			is_decl = true
		}
		if is_decl {
			di := int64(1)
			_ = di
			for (di < decl_index_len(out)) {
				decl := decl_index_get(out, di)
				_ = decl
				matched := false
				_ = matched
				if ((((nk == "Block") && (node.S1 == "_entry")) && (decl_kind_name(decl.Kind) == "entry")) && (decl.Node_idx == int64(0))) {
					matched = true
				} else if (((((nk == "Block") && strings.HasPrefix(node.S1, "_test:")) && (decl_kind_name(decl.Kind) == "test")) && (decl.Name == node_name)) && (decl.Node_idx == int64(0))) {
					matched = true
				} else if ((((decl.Name == node_name) && (decl.Node_idx == int64(0))) && (decl_kind_name(decl.Kind) != "entry")) && (decl_kind_name(decl.Kind) != "test")) {
					matched = true
				}
				if matched {
					updated := DeclInfo{Kind: decl.Kind, Name: decl.Name, Token_start: decl.Token_start, Body_start: decl.Body_start, Body_end: decl.Body_end, Is_pub: decl.Is_pub, Node_idx: actual_idx}
					_ = updated
					out = _decl_index_set(out, di, updated)
					di = decl_index_len(out)
				} else {
					di = (di + int64(1))
				}
			}
		}
		ni = (ni + int64(1))
	}
	return out
}

func _bdi_skip_braces(tokens []Token, start int64, len int64) int64 {
	lb := "{"
	_ = lb
	rb := "}"
	_ = rb
	if ((start >= len) || (token_name(tokens[start].Kind) != lb)) {
		return start
	}
	pos := (start + int64(1))
	_ = pos
	depth := int64(1)
	_ = depth
	for ((depth > int64(0)) && (pos < len)) {
		if (token_name(tokens[pos].Kind) == lb) {
			depth = (depth + int64(1))
		}
		if (token_name(tokens[pos].Kind) == rb) {
			depth = (depth - int64(1))
		}
		pos = (pos + int64(1))
	}
	return pos
}

func _bdi_skip_brackets(tokens []Token, start int64, len int64) int64 {
	if ((start >= len) || (token_name(tokens[start].Kind) != "[")) {
		return start
	}
	pos := (start + int64(1))
	_ = pos
	depth := int64(1)
	_ = depth
	for ((depth > int64(0)) && (pos < len)) {
		if (token_name(tokens[pos].Kind) == "[") {
			depth = (depth + int64(1))
		}
		if (token_name(tokens[pos].Kind) == "]") {
			depth = (depth - int64(1))
		}
		pos = (pos + int64(1))
	}
	return pos
}

func _bdi_skip_parens(tokens []Token, start int64, len int64) int64 {
	if ((start >= len) || (token_name(tokens[start].Kind) != "(")) {
		return start
	}
	pos := (start + int64(1))
	_ = pos
	depth := int64(1)
	_ = depth
	for ((depth > int64(0)) && (pos < len)) {
		if (token_name(tokens[pos].Kind) == "(") {
			depth = (depth + int64(1))
		}
		if (token_name(tokens[pos].Kind) == ")") {
			depth = (depth - int64(1))
		}
		pos = (pos + int64(1))
	}
	return pos
}

func _bdi_find_fn_body(tokens []Token, start int64, len int64) int64 {
	pos := start
	_ = pos
	lb := "{"
	_ = lb
	if ((pos < len) && _is_fn_name_token(token_name(tokens[pos].Kind))) {
		pos = (pos + int64(1))
	}
	if ((pos < len) && (token_name(tokens[pos].Kind) == "[")) {
		pos = _bdi_skip_brackets(tokens, pos, len)
	}
	if ((pos < len) && (token_name(tokens[pos].Kind) == "(")) {
		pos = _bdi_skip_parens(tokens, pos, len)
	}
	if ((pos < len) && (token_name(tokens[pos].Kind) == "->")) {
		pos = (pos + int64(1))
		pos = _bdi_skip_type_tokens(tokens, pos, len)
	}
	if ((pos < len) && (token_name(tokens[pos].Kind) == "!")) {
		pos = (pos + int64(1))
		pos = _bdi_skip_type_tokens(tokens, pos, len)
	}
	if ((pos < len) && (token_name(tokens[pos].Kind) == "with")) {
		pos = (pos + int64(1))
		for (((((pos < len) && (token_name(tokens[pos].Kind) != "=")) && (token_name(tokens[pos].Kind) != lb)) && (token_name(tokens[pos].Kind) != "NEWLINE")) && (token_name(tokens[pos].Kind) != "EOF")) {
			pos = (pos + int64(1))
		}
	}
	if ((pos < len) && (token_name(tokens[pos].Kind) == "where")) {
		pos = (pos + int64(1))
		for (((((pos < len) && (token_name(tokens[pos].Kind) != "=")) && (token_name(tokens[pos].Kind) != lb)) && (token_name(tokens[pos].Kind) != "NEWLINE")) && (token_name(tokens[pos].Kind) != "EOF")) {
			pos = (pos + int64(1))
		}
	}
	for ((pos < len) && (token_name(tokens[pos].Kind) == "NEWLINE")) {
		pos = (pos + int64(1))
	}
	return pos
}

func _bdi_skip_type_tokens(tokens []Token, start int64, len int64) int64 {
	pos := start
	_ = pos
	if (pos >= len) {
		return pos
	}
	k := token_name(tokens[pos].Kind)
	_ = k
	if ((k == "IDENT") || (k == "Self")) {
		pos = (pos + int64(1))
		for ((pos < len) && (token_name(tokens[pos].Kind) == ".")) {
			nxt := (pos + int64(1))
			_ = nxt
			if ((nxt < len) && (token_name(tokens[nxt].Kind) == ".")) {
				return pos
			}
			pos = (pos + int64(2))
		}
		if ((pos < len) && (token_name(tokens[pos].Kind) == "[")) {
			pos = _bdi_skip_brackets(tokens, pos, len)
		}
		if ((pos < len) && (token_name(tokens[pos].Kind) == "?")) {
			pos = (pos + int64(1))
		}
		if ((pos < len) && (token_name(tokens[pos].Kind) == "!")) {
			pos = (pos + int64(1))
			pos = _bdi_skip_type_tokens(tokens, pos, len)
		}
		return pos
	}
	if (k == "[") {
		pos = _bdi_skip_brackets(tokens, pos, len)
		return pos
	}
	if (k == "(") {
		pos = _bdi_skip_parens(tokens, pos, len)
		if ((pos < len) && (token_name(tokens[pos].Kind) == "->")) {
			pos = (pos + int64(1))
			pos = _bdi_skip_type_tokens(tokens, pos, len)
		}
		return pos
	}
	if (k == "fn") {
		pos = (pos + int64(1))
		if (((pos < len) && (token_name(tokens[pos].Kind) == "IDENT")) && (tokens[pos].Text == "once")) {
			pos = (pos + int64(1))
		}
		if ((pos < len) && (token_name(tokens[pos].Kind) == "(")) {
			pos = _bdi_skip_parens(tokens, pos, len)
		}
		if ((pos < len) && (token_name(tokens[pos].Kind) == "->")) {
			pos = (pos + int64(1))
			pos = _bdi_skip_type_tokens(tokens, pos, len)
		}
		return pos
	}
	return pos
}

func _bdi_skip_expr(tokens []Token, start int64, len int64) int64 {
	pos := start
	_ = pos
	lb := "{"
	_ = lb
	rb := "}"
	_ = rb
	for (pos < len) {
		k := token_name(tokens[pos].Kind)
		_ = k
		if ((k == "NEWLINE") || (k == "EOF")) {
			return pos
		}
		if (k == lb) {
			pos = _bdi_skip_braces(tokens, pos, len)
		} else {
			pos = (pos + int64(1))
		}
	}
	return pos
}

func _parse_top(p Parser) PR {
	k := _ck(p)
	_ = k
	if (k == "fn") {
		return _parse_fn(p)
	}
	if (k == "pub") {
		pp := _adv(p)
		_ = pp
		return _parse_top(pp)
	}
	if ((k == "type") || (k == "struct")) {
		return _parse_type_decl(p)
	}
	if (k == "enum") {
		return _parse_enum_decl(p)
	}
	if (k == "trait") {
		return _parse_trait_decl(p)
	}
	if (k == "impl") {
		return _parse_impl_decl(p)
	}
	if (k == "const") {
		return _parse_const_decl(p)
	}
	if (k == "entry") {
		return _parse_entry(p)
	}
	if (k == "test") {
		return _parse_test(p)
	}
	return _parse_stmt(p)
}

func _parse_fn(p Parser) PR {
	span := _pspan(p)
	_ = span
	pp := _adv(p)
	_ = pp
	name := _cur(pp).Text
	_ = name
	pp = _adv(pp)
	generics := ""
	_ = generics
	if (_ck(pp) == "[") {
		pp = _adv(pp)
		gcount := int64(0)
		_ = gcount
		for ((_ck(pp) != "]") && (_ck(pp) != "EOF")) {
			if (_ck(pp) == "IDENT") {
				gname := _cur(pp).Text
				_ = gname
				pp = _adv(pp)
				if (_ck(pp) == "[") {
					next_pos := (pp.Pos + int64(1))
					_ = next_pos
					if ((next_pos < int64(len(pp.Tokens))) && (_cur(_adv(pp)).Text == "_")) {
						pp = _adv(pp)
						pp = _adv(pp)
						if (_ck(pp) == "]") {
							pp = _adv(pp)
						}
						gname = (gname + "[_]")
					}
				}
				if (gcount > int64(0)) {
					generics = (generics + ",")
				}
				generics = (generics + gname)
				if (_ck(pp) == ":") {
					pp = _adv(pp)
					if (_ck(pp) == "IDENT") {
						generics = ((generics + ":") + _cur(pp).Text)
						pp = _adv(pp)
						if (_ck(pp) == "<") {
							pp = _adv(pp)
							if (_ck(pp) == "IDENT") {
								assoc_n := _cur(pp).Text
								_ = assoc_n
								pp = _adv(pp)
								if (_ck(pp) == "=") {
									pp = _adv(pp)
									if (_ck(pp) == "IDENT") {
										assoc_v := _cur(pp).Text
										_ = assoc_v
										generics = (((((generics + "<") + assoc_n) + "=") + assoc_v) + ">")
										pp = _adv(pp)
									}
								}
							}
							if (_ck(pp) == ">") {
								pp = _adv(pp)
							}
						}
						for (_ck(pp) == "+") {
							pp = _adv(pp)
							if (_ck(pp) == "IDENT") {
								generics = ((generics + "+") + _cur(pp).Text)
								pp = _adv(pp)
							}
						}
					}
				}
				gcount = (gcount + int64(1))
			} else {
				pp = _adv(pp)
			}
			if (_ck(pp) == ",") {
				pp = _adv(pp)
			}
		}
		if (_ck(pp) == "]") {
			pp = _adv(pp)
		}
	}
	pp = _expect(pp, "(")
	param_indices := []int64{int64(0)}
	_ = param_indices
	param_count := int64(0)
	_ = param_count
	for ((_ck(pp) != ")") && (_ck(pp) != "EOF")) {
		is_param := ((_ck(pp) == "IDENT") || (_ck(pp) == "self"))
		_ = is_param
		if (is_param == false) {
			npos := (pp.Pos + int64(1))
			_ = npos
			if (npos < int64(len(pp.Tokens))) {
				nk := token_name(pp.Tokens[npos].Kind)
				_ = nk
				if (nk == ":") {
					is_param = true
				}
			}
		}
		if is_param {
			pname := _cur(pp).Text
			_ = pname
			pspan := _pspan(pp)
			_ = pspan
			pp = _adv(pp)
			if (pname == "self") {
				pidx := _pidx(pp)
				_ = pidx
				pp = _padd(pp, mk_param("self", int64(0), true, pspan))
				param_indices = append(param_indices, pidx)
				param_count = (param_count + int64(1))
			} else if (_ck(pp) == ":") {
				pp = _adv(pp)
				type_tok_pos := pp.Pos
				_ = type_tok_pos
				pp = _skip_type(pp)
				pidx := _pidx(pp)
				_ = pidx
				pp = _padd(pp, mk_param(pname, type_tok_pos, false, pspan))
				param_indices = append(param_indices, pidx)
				param_count = (param_count + int64(1))
			}
		} else {
			pp = _adv(pp)
		}
		if (_ck(pp) == ",") {
			pp = _adv(pp)
		}
	}
	pp = _adv(pp)
	has_ret := false
	_ = has_ret
	ret_tok_pos := int64(0)
	_ = ret_tok_pos
	if (_ck(pp) == "->") {
		pp = _adv(pp)
		has_ret = true
		ret_tok_pos = pp.Pos
		pp = _skip_type(pp)
	}
	err_tok_pos := int64(0)
	_ = err_tok_pos
	if (_ck(pp) == "!") {
		pp = _adv(pp)
		err_tok_pos = pp.Pos
		pp = _skip_type(pp)
	}
	if (_ck(pp) == "with") {
		pp = _adv(pp)
		lb := "{"
		_ = lb
		for ((((_ck(pp) != "NEWLINE") && (_ck(pp) != "EOF")) && (_ck(pp) != "=")) && (_ck(pp) != lb)) {
			pp = _adv(pp)
		}
	}
	if (_ck(pp) == "where") {
		lb := "{"
		_ = lb
		for ((((_ck(pp) != "NEWLINE") && (_ck(pp) != "EOF")) && (_ck(pp) != "=")) && (_ck(pp) != lb)) {
			pp = _adv(pp)
		}
	}
	body := _no_expr()
	_ = body
	lb := "{"
	_ = lb
	if (_ck(pp) == "=") {
		pp = _adv(pp)
		result := _parse_expr(pp, int64(0))
		_ = result
		pp = result.P
		body = result.E
	} else if (_ck(pp) == lb) {
		result := _parse_block(pp)
		_ = result
		pp = result.P
		body = result.E
	} else {
		pp = _skip_nl(pp)
		if (_ck(pp) == lb) {
			result := _parse_block(pp)
			_ = result
			pp = result.P
			body = result.E
		}
	}
	has_body := false
	_ = has_body
	body_idx := int64(0)
	_ = body_idx
	if (expr_kind_name(body.Kind) != "None") {
		has_body = true
		body_idx = _pidx(pp)
		pp = _padd(pp, body)
	}
	list_start := _pidx(pp)
	_ = list_start
	pi := int64(0)
	_ = pi
	for (pi < param_count) {
		pp = _padd(pp, Expr{Kind: ExprKindEkNone{}, S1: "", S2: "", B1: false, B2: false, C0: param_indices[(pi + int64(1))], C1: int64(0), C2: int64(0), List_start: int64(0), List_count: int64(0), Span: _no_span()})
		pi = (pi + int64(1))
	}
	return PR{P: pp, E: Expr{Kind: ExprKindEkFnDecl{}, S1: name, S2: generics, B1: has_body, B2: has_ret, C0: body_idx, C1: ret_tok_pos, C2: err_tok_pos, List_start: list_start, List_count: param_count, Span: span}}
}

func _parse_type(p Parser) PR {
	pp := p
	_ = pp
	k := _ck(pp)
	_ = k
	span := _pspan(pp)
	_ = span
	if ((k == "IDENT") || (k == "Self")) {
		name := _cur(pp).Text
		_ = name
		pp = _adv(pp)
		for ((_ck(pp) == ".") && (_pk(pp) != ".")) {
			pp = _adv(pp)
			name = ((name + ".") + _cur(pp).Text)
			pp = _adv(pp)
		}
		type_arg_idx := int64(0)
		_ = type_arg_idx
		if (_ck(pp) == "[") {
			pp = _adv(pp)
			d := int64(1)
			_ = d
			for ((d > int64(0)) && (_ck(pp) != "EOF")) {
				if (_ck(pp) == "[") {
					d = (d + int64(1))
				}
				if (_ck(pp) == "]") {
					d = (d - int64(1))
				}
				if (d > int64(0)) {
					pp = _adv(pp)
				}
			}
			pp = _adv(pp)
		}
		if (_ck(pp) == "?") {
			pp = _adv(pp)
			base := Expr{Kind: ExprKindEkTypeRef{}, S1: name, S2: "", B1: false, B2: false, C0: int64(0), C1: int64(0), C2: int64(0), List_start: int64(0), List_count: int64(0), Span: span}
			_ = base
			base_idx := _pidx(pp)
			_ = base_idx
			pp = _padd(pp, base)
			return PR{P: pp, E: Expr{Kind: ExprKindEkTypeOptional{}, S1: "", S2: "", B1: false, B2: false, C0: base_idx, C1: int64(0), C2: int64(0), List_start: int64(0), List_count: int64(0), Span: span}}
		}
		if (_ck(pp) == "!") {
			pp = _adv(pp)
			err := _parse_type(pp)
			_ = err
			err_idx := _pidx(err.P)
			_ = err_idx
			pp = _padd(err.P, err.E)
			base := Expr{Kind: ExprKindEkTypeRef{}, S1: name, S2: "", B1: false, B2: false, C0: int64(0), C1: int64(0), C2: int64(0), List_start: int64(0), List_count: int64(0), Span: span}
			_ = base
			base_idx := _pidx(pp)
			_ = base_idx
			pp = _padd(pp, base)
			return PR{P: pp, E: Expr{Kind: ExprKindEkTypeResult{}, S1: "", S2: "", B1: false, B2: false, C0: base_idx, C1: err_idx, C2: int64(0), List_start: int64(0), List_count: int64(0), Span: span}}
		}
		return PR{P: pp, E: Expr{Kind: ExprKindEkTypeRef{}, S1: name, S2: "", B1: false, B2: false, C0: int64(0), C1: int64(0), C2: int64(0), List_start: int64(0), List_count: int64(0), Span: span}}
	}
	if (k == "[") {
		pp = _adv(pp)
		elem := _parse_type(pp)
		_ = elem
		elem_idx := _pidx(elem.P)
		_ = elem_idx
		pp = _padd(elem.P, elem.E)
		pp = _expect(pp, "]")
		return PR{P: pp, E: Expr{Kind: ExprKindEkTypeArray{}, S1: "", S2: "", B1: false, B2: false, C0: elem_idx, C1: int64(0), C2: int64(0), List_start: int64(0), List_count: int64(0), Span: span}}
	}
	if (k == "(") {
		pp = _adv(pp)
		list_start := _pidx(pp)
		_ = list_start
		count := int64(0)
		_ = count
		for ((_ck(pp) != ")") && (_ck(pp) != "EOF")) {
			sub := _parse_type(pp)
			_ = sub
			pp = _padd(sub.P, sub.E)
			count = (count + int64(1))
			if (_ck(pp) == ",") {
				pp = _adv(pp)
			}
		}
		pp = _adv(pp)
		if (_ck(pp) == "->") {
			pp = _adv(pp)
			ret := _parse_type(pp)
			_ = ret
			ret_idx := _pidx(ret.P)
			_ = ret_idx
			pp = _padd(ret.P, ret.E)
			return PR{P: pp, E: Expr{Kind: ExprKindEkTypeFn{}, S1: "", S2: "", B1: false, B2: false, C0: ret_idx, C1: int64(0), C2: int64(0), List_start: list_start, List_count: count, Span: span}}
		}
		return PR{P: pp, E: Expr{Kind: ExprKindEkTypeTuple{}, S1: "", S2: "", B1: false, B2: false, C0: int64(0), C1: int64(0), C2: int64(0), List_start: list_start, List_count: count, Span: span}}
	}
	if (k == "fn") {
		pp = _adv(pp)
		if ((_ck(pp) == "IDENT") && (_cur(pp).Text == "once")) {
			pp = _adv(pp)
		}
		pp = _expect(pp, "(")
		list_start := _pidx(pp)
		_ = list_start
		count := int64(0)
		_ = count
		for ((_ck(pp) != ")") && (_ck(pp) != "EOF")) {
			sub := _parse_type(pp)
			_ = sub
			pp = _padd(sub.P, sub.E)
			count = (count + int64(1))
			if (_ck(pp) == ",") {
				pp = _adv(pp)
			}
		}
		pp = _adv(pp)
		ret_idx := int64(0)
		_ = ret_idx
		if (_ck(pp) == "->") {
			pp = _adv(pp)
			ret := _parse_type(pp)
			_ = ret
			ret_idx = _pidx(ret.P)
			pp = _padd(ret.P, ret.E)
		}
		return PR{P: pp, E: Expr{Kind: ExprKindEkTypeFn{}, S1: "", S2: "", B1: false, B2: false, C0: ret_idx, C1: int64(0), C2: int64(0), List_start: list_start, List_count: count, Span: span}}
	}
	return PR{P: pp, E: Expr{Kind: ExprKindEkTypeRef{}, S1: "", S2: "", B1: false, B2: false, C0: int64(0), C1: int64(0), C2: int64(0), List_start: int64(0), List_count: int64(0), Span: span}}
}

func _skip_type(p Parser) Parser {
	result := _parse_type(p)
	_ = result
	return result.P
}

func _skip_braces(p Parser) Parser {
	lb := "{"
	_ = lb
	rb := "}"
	_ = rb
	if (_ck(p) != lb) {
		return p
	}
	pp := _adv(p)
	_ = pp
	d := int64(1)
	_ = d
	for ((d > int64(0)) && (_ck(pp) != "EOF")) {
		if (_ck(pp) == lb) {
			d = (d + int64(1))
		}
		if (_ck(pp) == rb) {
			d = (d - int64(1))
		}
		if (d > int64(0)) {
			pp = _adv(pp)
		}
	}
	return _adv(pp)
}

func _parse_type_decl(p Parser) PR {
	span := _pspan(p)
	_ = span
	pp := _adv(p)
	_ = pp
	name := _cur(pp).Text
	_ = name
	pp = _adv(pp)
	generics := ""
	_ = generics
	if (_ck(pp) == "[") {
		pp = _adv(pp)
		gcount := int64(0)
		_ = gcount
		for ((_ck(pp) != "]") && (_ck(pp) != "EOF")) {
			if (_ck(pp) == "IDENT") {
				gname := _cur(pp).Text
				_ = gname
				if (gcount > int64(0)) {
					generics = (generics + ",")
				}
				generics = (generics + gname)
				pp = _adv(pp)
				if (_ck(pp) == ":") {
					pp = _adv(pp)
					if (_ck(pp) == "IDENT") {
						generics = ((generics + ":") + _cur(pp).Text)
						pp = _adv(pp)
						if (_ck(pp) == "<") {
							pp = _adv(pp)
							if (_ck(pp) == "IDENT") {
								assoc_n := _cur(pp).Text
								_ = assoc_n
								pp = _adv(pp)
								if (_ck(pp) == "=") {
									pp = _adv(pp)
									if (_ck(pp) == "IDENT") {
										assoc_v := _cur(pp).Text
										_ = assoc_v
										generics = (((((generics + "<") + assoc_n) + "=") + assoc_v) + ">")
										pp = _adv(pp)
									}
								}
							}
							if (_ck(pp) == ">") {
								pp = _adv(pp)
							}
						}
						for (_ck(pp) == "+") {
							pp = _adv(pp)
							if (_ck(pp) == "IDENT") {
								generics = ((generics + "+") + _cur(pp).Text)
								pp = _adv(pp)
							}
						}
					}
				}
				gcount = (gcount + int64(1))
			} else {
				pp = _adv(pp)
			}
			if (_ck(pp) == ",") {
				pp = _adv(pp)
			}
		}
		if (_ck(pp) == "]") {
			pp = _adv(pp)
		}
	}
	lb := "{"
	_ = lb
	has_variants := false
	_ = has_variants
	newtype_tok_pos := int64(0)
	_ = newtype_tok_pos
	child_indices := []int64{int64(0)}
	_ = child_indices
	child_count := int64(0)
	_ = child_count
	if (_ck(pp) == "=") {
		pp = _adv(pp)
		pp = _skip_nl(pp)
	}
	if (_ck(pp) == "|") {
		has_variants = true
		for (_ck(pp) == "|") {
			pp = _adv(pp)
			pp = _skip_nl(pp)
			vname := _cur(pp).Text
			_ = vname
			vspan := _pspan(pp)
			_ = vspan
			pp = _adv(pp)
			vf_indices := []int64{int64(0)}
			_ = vf_indices
			vf_count := int64(0)
			_ = vf_count
			is_tuple := false
			_ = is_tuple
			if (_ck(pp) == "(") {
				is_tuple = true
				pp = _adv(pp)
				fi := int64(0)
				_ = fi
				for ((_ck(pp) != ")") && (_ck(pp) != "EOF")) {
					fspan := _pspan(pp)
					_ = fspan
					type_tok_pos := pp.Pos
					_ = type_tok_pos
					pp = _skip_type(pp)
					fidx := _pidx(pp)
					_ = fidx
					pp = _padd(pp, mk_field(("_" + i2s(fi)), type_tok_pos, fspan))
					vf_indices = append(vf_indices, fidx)
					vf_count = (vf_count + int64(1))
					fi = (fi + int64(1))
					if (_ck(pp) == ",") {
						pp = _adv(pp)
					}
				}
				if (_ck(pp) == ")") {
					pp = _adv(pp)
				}
			} else if (_ck(pp) == lb) {
				pp = _adv(pp)
				pp = _skip_nl(pp)
				rb := "}"
				_ = rb
				for ((_ck(pp) != rb) && (_ck(pp) != "EOF")) {
					if (_ck(pp) == "IDENT") {
						fname := _cur(pp).Text
						_ = fname
						fspan := _pspan(pp)
						_ = fspan
						pp = _adv(pp)
						if (_ck(pp) == ":") {
							pp = _adv(pp)
							type_tok_pos := pp.Pos
							_ = type_tok_pos
							pp = _skip_type(pp)
							fidx := _pidx(pp)
							_ = fidx
							pp = _padd(pp, mk_field(fname, type_tok_pos, fspan))
							vf_indices = append(vf_indices, fidx)
							vf_count = (vf_count + int64(1))
						}
					} else {
						pp = _adv(pp)
					}
					pp = _skip_nl(pp)
					if (_ck(pp) == ",") {
						pp = _adv(pp)
					}
					pp = _skip_nl(pp)
				}
				if (_ck(pp) == rb) {
					pp = _adv(pp)
				}
			}
			vf_list_start := _pidx(pp)
			_ = vf_list_start
			vfi := int64(0)
			_ = vfi
			for (vfi < vf_count) {
				pp = _padd(pp, Expr{Kind: ExprKindEkNone{}, S1: "", S2: "", B1: false, B2: false, C0: vf_indices[(vfi + int64(1))], C1: int64(0), C2: int64(0), List_start: int64(0), List_count: int64(0), Span: _no_span()})
				vfi = (vfi + int64(1))
			}
			vidx := _pidx(pp)
			_ = vidx
			pp = _padd(pp, Expr{Kind: ExprKindEkVariant{}, S1: vname, S2: "", B1: is_tuple, B2: false, C0: int64(0), C1: int64(0), C2: int64(0), List_start: vf_list_start, List_count: vf_count, Span: vspan})
			child_indices = append(child_indices, vidx)
			child_count = (child_count + int64(1))
			pp = _skip_nl(pp)
		}
	} else if (_ck(pp) == lb) {
		pp = _adv(pp)
		pp = _skip_nl(pp)
		rb := "}"
		_ = rb
		for ((_ck(pp) != rb) && (_ck(pp) != "EOF")) {
			if (_ck(pp) == "IDENT") {
				fname := _cur(pp).Text
				_ = fname
				fspan := _pspan(pp)
				_ = fspan
				pp = _adv(pp)
				if (_ck(pp) == ":") {
					pp = _adv(pp)
					type_tok_pos := pp.Pos
					_ = type_tok_pos
					pp = _skip_type(pp)
					fidx := _pidx(pp)
					_ = fidx
					pp = _padd(pp, mk_field(fname, type_tok_pos, fspan))
					child_indices = append(child_indices, fidx)
					child_count = (child_count + int64(1))
				}
			} else {
				pp = _adv(pp)
			}
			pp = _skip_nl(pp)
			if (_ck(pp) == ",") {
				pp = _adv(pp)
			}
			pp = _skip_nl(pp)
		}
		if (_ck(pp) == rb) {
			pp = _adv(pp)
		}
	} else {
		if (((_ck(pp) != "NEWLINE") && (_ck(pp) != "EOF")) && (_ck(pp) != "derives")) {
			newtype_tok_pos = pp.Pos
			pp = _skip_type(pp)
		}
	}
	if (_ck(pp) == "derives") {
		pp = _adv(pp)
		if (_ck(pp) == "[") {
			pp = _adv(pp)
			for ((_ck(pp) != "]") && (_ck(pp) != "EOF")) {
				pp = _adv(pp)
			}
			pp = _adv(pp)
		}
	}
	list_start := _pidx(pp)
	_ = list_start
	ci := int64(0)
	_ = ci
	for (ci < child_count) {
		pp = _padd(pp, Expr{Kind: ExprKindEkNone{}, S1: "", S2: "", B1: false, B2: false, C0: child_indices[(ci + int64(1))], C1: int64(0), C2: int64(0), List_start: int64(0), List_count: int64(0), Span: _no_span()})
		ci = (ci + int64(1))
	}
	is_sum := false
	_ = is_sum
	if has_variants {
		is_sum = true
	}
	glen := int64(len(generics))
	_ = glen
	has_generics := false
	_ = has_generics
	if (glen > int64(0)) {
		has_generics = true
	}
	return PR{P: pp, E: Expr{Kind: ExprKindEkTypeDecl{}, S1: name, S2: generics, B1: is_sum, B2: has_generics, C0: newtype_tok_pos, C1: int64(0), C2: int64(0), List_start: list_start, List_count: child_count, Span: span}}
}

func _parse_enum_decl(p Parser) PR {
	span := _pspan(p)
	_ = span
	pp := _adv(p)
	_ = pp
	name := _cur(pp).Text
	_ = name
	pp = _adv(pp)
	lb := "{"
	_ = lb
	rb := "}"
	_ = rb
	variant_names := ""
	_ = variant_names
	variant_count := int64(0)
	_ = variant_count
	if (_ck(pp) == lb) {
		pp = _adv(pp)
		pp = _skip_nl(pp)
		for ((_ck(pp) != rb) && (_ck(pp) != "EOF")) {
			pp = _skip_nl(pp)
			if (_ck(pp) == "IDENT") {
				if (variant_count > int64(0)) {
					variant_names = (variant_names + ",")
				}
				variant_names = (variant_names + _cur(pp).Text)
				variant_count = (variant_count + int64(1))
				pp = _adv(pp)
			} else if ((_ck(pp) == rb) || (_ck(pp) == "EOF")) {
				_skip := int64(0)
				_ = _skip
			} else {
				pp = _adv(pp)
			}
			pp = _skip_nl(pp)
		}
		if (_ck(pp) == rb) {
			pp = _adv(pp)
		}
	}
	return PR{P: pp, E: Expr{Kind: ExprKindEkEnumDecl{}, S1: name, S2: variant_names, B1: false, B2: false, C0: int64(0), C1: int64(0), C2: int64(0), List_start: int64(0), List_count: variant_count, Span: span}}
}

func _parse_trait_decl(p Parser) PR {
	span := _pspan(p)
	_ = span
	pp := _adv(p)
	_ = pp
	name := _cur(pp).Text
	_ = name
	pp = _adv(pp)
	supers := ""
	_ = supers
	if (_ck(pp) == ":") {
		pp = _adv(pp)
		scount := int64(0)
		_ = scount
		parsing := true
		_ = parsing
		for parsing {
			if (_ck(pp) == "IDENT") {
				if (scount > int64(0)) {
					supers = (supers + ",")
				}
				supers = (supers + _cur(pp).Text)
				scount = (scount + int64(1))
				pp = _adv(pp)
				if (_ck(pp) == ",") {
					pp = _adv(pp)
				} else {
					parsing = false
				}
			} else {
				parsing = false
			}
		}
	}
	lb := "{"
	_ = lb
	for ((_ck(pp) != lb) && (_ck(pp) != "EOF")) {
		pp = _adv(pp)
	}
	if (_ck(pp) != lb) {
		return PR{P: pp, E: Expr{Kind: ExprKindEkTraitDecl{}, S1: name, S2: supers, B1: false, B2: false, C0: int64(0), C1: int64(0), C2: int64(0), List_start: int64(0), List_count: int64(0), Span: span}}
	}
	pp = _adv(pp)
	pp = _skip_nl(pp)
	rb := "}"
	_ = rb
	method_indices := []int64{int64(0)}
	_ = method_indices
	method_count := int64(0)
	_ = method_count
	assoc_types := ""
	_ = assoc_types
	assoc_count := int64(0)
	_ = assoc_count
	for ((_ck(pp) != rb) && (_ck(pp) != "EOF")) {
		pp = _skip_nl(pp)
		if (_ck(pp) == "pub") {
			pp = _adv(pp)
		}
		if (_ck(pp) == "fn") {
			result := _parse_fn(pp)
			_ = result
			method_idx := _pidx(result.P)
			_ = method_idx
			pp = _padd(result.P, result.E)
			method_indices = append(method_indices, method_idx)
			method_count = (method_count + int64(1))
		} else if (_ck(pp) == "type") {
			pp = _adv(pp)
			if (_ck(pp) == "IDENT") {
				if (assoc_count > int64(0)) {
					assoc_types = (assoc_types + ",")
				}
				assoc_types = (assoc_types + _cur(pp).Text)
				assoc_count = (assoc_count + int64(1))
			}
			for (((_ck(pp) != "NEWLINE") && (_ck(pp) != rb)) && (_ck(pp) != "EOF")) {
				pp = _adv(pp)
			}
		} else if ((_ck(pp) == rb) || (_ck(pp) == "EOF")) {
			_skip := int64(0)
			_ = _skip
		} else {
			pp = _adv(pp)
		}
		pp = _skip_nl(pp)
	}
	if (_ck(pp) == rb) {
		pp = _adv(pp)
	}
	list_start := _pidx(pp)
	_ = list_start
	mi := int64(0)
	_ = mi
	for (mi < method_count) {
		pp = _padd(pp, Expr{Kind: ExprKindEkNone{}, S1: "", S2: "", B1: false, B2: false, C0: method_indices[(mi + int64(1))], C1: int64(0), C2: int64(0), List_start: int64(0), List_count: int64(0), Span: _no_span()})
		mi = (mi + int64(1))
	}
	full_s2 := supers
	_ = full_s2
	alen := int64(len(assoc_types))
	_ = alen
	if (alen > int64(0)) {
		full_s2 = ((full_s2 + ";") + assoc_types)
	}
	return PR{P: pp, E: Expr{Kind: ExprKindEkTraitDecl{}, S1: name, S2: full_s2, B1: false, B2: false, C0: int64(0), C1: int64(0), C2: int64(0), List_start: list_start, List_count: method_count, Span: span}}
}

func _parse_impl_decl(p Parser) PR {
	span := _pspan(p)
	_ = span
	pp := _adv(p)
	_ = pp
	impl_type := ""
	_ = impl_type
	if (_ck(pp) == "IDENT") {
		impl_type = _cur(pp).Text
	}
	trait_name := ""
	_ = trait_name
	lb := "{"
	_ = lb
	scan := pp
	_ = scan
	for ((_ck(scan) != lb) && (_ck(scan) != "EOF")) {
		if (_ck(scan) == ":") {
			scan = _adv(scan)
			if (_ck(scan) == "IDENT") {
				trait_name = _cur(scan).Text
			}
		}
		scan = _adv(scan)
	}
	for ((_ck(pp) != lb) && (_ck(pp) != "EOF")) {
		pp = _adv(pp)
	}
	if (_ck(pp) != lb) {
		return PR{P: pp, E: mk_impl_decl(impl_type, trait_name, int64(0), span)}
	}
	pp = _adv(pp)
	pp = _skip_nl(pp)
	rb := "}"
	_ = rb
	method_indices := []int64{int64(0)}
	_ = method_indices
	method_count := int64(0)
	_ = method_count
	for ((_ck(pp) != rb) && (_ck(pp) != "EOF")) {
		pp = _skip_nl(pp)
		if (_ck(pp) == "pub") {
			pp = _adv(pp)
		}
		if (_ck(pp) == "fn") {
			result := _parse_fn(pp)
			_ = result
			method_idx := _pidx(result.P)
			_ = method_idx
			pp = _padd(result.P, result.E)
			method_indices = append(method_indices, method_idx)
			method_count = (method_count + int64(1))
		} else if ((_ck(pp) == rb) || (_ck(pp) == "EOF")) {
			_skip := int64(0)
			_ = _skip
		} else {
			pp = _adv(pp)
		}
		pp = _skip_nl(pp)
	}
	if (_ck(pp) == rb) {
		pp = _adv(pp)
	}
	list_start := _pidx(pp)
	_ = list_start
	mi := int64(0)
	_ = mi
	for (mi < method_count) {
		pp = _padd(pp, Expr{Kind: ExprKindEkNone{}, S1: "", S2: "", B1: false, B2: false, C0: method_indices[(mi + int64(1))], C1: int64(0), C2: int64(0), List_start: int64(0), List_count: int64(0), Span: _no_span()})
		mi = (mi + int64(1))
	}
	return PR{P: pp, E: Expr{Kind: ExprKindEkImplDecl{}, S1: impl_type, S2: trait_name, B1: false, B2: false, C0: int64(0), C1: int64(0), C2: int64(0), List_start: list_start, List_count: method_count, Span: span}}
}

func _parse_const_decl(p Parser) PR {
	span := _pspan(p)
	_ = span
	pp := _adv(p)
	_ = pp
	name := _cur(pp).Text
	_ = name
	pp = _adv(pp)
	type_tok_pos := int64(0)
	_ = type_tok_pos
	if (_ck(pp) == ":") {
		pp = _adv(pp)
		type_tok_pos = pp.Pos
		pp = _skip_type(pp)
	}
	pp = _expect(pp, "=")
	result := _parse_expr(pp, int64(0))
	_ = result
	val_idx := _pidx(result.P)
	_ = val_idx
	pp = _padd(result.P, result.E)
	return PR{P: pp, E: Expr{Kind: ExprKindEkConstDecl{}, S1: name, S2: "", B1: false, B2: false, C0: val_idx, C1: type_tok_pos, C2: int64(0), List_start: int64(0), List_count: int64(0), Span: span}}
}

func _parse_entry(p Parser) PR {
	span := _pspan(p)
	_ = span
	pp := _adv(p)
	_ = pp
	result := _parse_block(pp)
	_ = result
	body_idx := _pidx(result.P)
	_ = body_idx
	pp = _padd(result.P, result.E)
	return PR{P: pp, E: Expr{Kind: ExprKindEkBlock{}, S1: "_entry", S2: "", B1: false, B2: false, C0: body_idx, C1: int64(0), C2: int64(0), List_start: result.E.List_start, List_count: result.E.List_count, Span: span}}
}

func _parse_test(p Parser) PR {
	span := _pspan(p)
	_ = span
	pp := _adv(p)
	_ = pp
	test_name := ""
	_ = test_name
	if ((_ck(pp) == "STRING") || (_ck(pp) == "IDENT")) {
		test_name = _cur(pp).Text
		pp = _adv(pp)
	}
	result := _parse_block(pp)
	_ = result
	body_idx := _pidx(result.P)
	_ = body_idx
	pp = _padd(result.P, result.E)
	return PR{P: pp, E: Expr{Kind: ExprKindEkBlock{}, S1: ("_test:" + test_name), S2: "", B1: false, B2: false, C0: body_idx, C1: int64(0), C2: int64(0), List_start: result.E.List_start, List_count: result.E.List_count, Span: span}}
}

func _parse_stmt(p Parser) PR {
	k := _ck(p)
	_ = k
	if (k == "mut") {
		return _parse_mut(p)
	}
	if (k == "return") {
		return _parse_return(p)
	}
	if (k == "break") {
		return PR{P: _adv(p), E: mk_break(_pspan(p))}
	}
	if (k == "continue") {
		return PR{P: _adv(p), E: mk_continue(_pspan(p))}
	}
	if (k == "yield") {
		span := _pspan(p)
		_ = span
		pp := _adv(p)
		_ = pp
		val := _parse_expr(pp, int64(0))
		_ = val
		val_idx := _pidx(val.P)
		_ = val_idx
		pp = _padd(val.P, val.E)
		return PR{P: pp, E: Expr{Kind: ExprKindEkYield{}, S1: "", S2: "", B1: true, B2: false, C0: val_idx, C1: int64(0), C2: int64(0), List_start: int64(0), List_count: int64(0), Span: span}}
	}
	if (k == "for") {
		return _parse_for(p)
	}
	if (k == "while") {
		return _parse_while(p)
	}
	if (k == "loop") {
		return _parse_loop(p)
	}
	if (k == "defer") {
		span := _pspan(p)
		_ = span
		pp := _adv(p)
		_ = pp
		r := _parse_expr(pp, int64(0))
		_ = r
		defer_idx := _pidx(r.P)
		_ = defer_idx
		pp = _padd(r.P, r.E)
		return PR{P: pp, E: Expr{Kind: ExprKindEkDefer{}, S1: "", S2: "", B1: false, B2: false, C0: defer_idx, C1: int64(0), C2: int64(0), List_start: int64(0), List_count: int64(0), Span: span}}
	}
	if (k == "assert") {
		return _parse_assert(p)
	}
	if (k == "with") {
		return _parse_with(p)
	}
	result := _parse_expr(p, int64(0))
	_ = result
	pp := result.P
	_ = pp
	if (_ck(pp) == ":=") {
		pp = _adv(pp)
		val := _parse_expr(pp, int64(0))
		_ = val
		val_idx := _pidx(val.P)
		_ = val_idx
		pp = _padd(val.P, val.E)
		pp = _skip_nl(pp)
		return PR{P: pp, E: Expr{Kind: ExprKindEkBinding{}, S1: result.E.S1, S2: "", B1: false, B2: false, C0: val_idx, C1: int64(0), C2: int64(0), List_start: int64(0), List_count: int64(0), Span: result.E.Span}}
	}
	if (_ck(pp) == ":") {
		pp = _adv(pp)
		pp = _skip_type(pp)
		pp = _expect(pp, "=")
		val := _parse_expr(pp, int64(0))
		_ = val
		val_idx := _pidx(val.P)
		_ = val_idx
		pp = _padd(val.P, val.E)
		pp = _skip_nl(pp)
		return PR{P: pp, E: Expr{Kind: ExprKindEkBinding{}, S1: result.E.S1, S2: "", B1: false, B2: false, C0: val_idx, C1: int64(0), C2: int64(0), List_start: int64(0), List_count: int64(0), Span: result.E.Span}}
	}
	if (_ck(pp) == "=") {
		lhs_idx := _pidx(pp)
		_ = lhs_idx
		pp = _padd(pp, result.E)
		pp = _adv(pp)
		val := _parse_expr(pp, int64(0))
		_ = val
		rhs_idx := _pidx(val.P)
		_ = rhs_idx
		pp = _padd(val.P, val.E)
		pp = _skip_nl(pp)
		return PR{P: pp, E: Expr{Kind: ExprKindEkAssign{}, S1: "", S2: "", B1: false, B2: false, C0: lhs_idx, C1: rhs_idx, C2: int64(0), List_start: int64(0), List_count: int64(0), Span: result.E.Span}}
	}
	pp = _skip_nl(pp)
	return PR{P: pp, E: result.E}
}

func _parse_mut(p Parser) PR {
	span := _pspan(p)
	_ = span
	pp := _adv(p)
	_ = pp
	name := _cur(pp).Text
	_ = name
	pp = _adv(pp)
	if (_ck(pp) == ":") {
		pp = _adv(pp)
		pp = _skip_type(pp)
		pp = _expect(pp, "=")
	} else {
		pp = _expect(pp, ":=")
	}
	val := _parse_expr(pp, int64(0))
	_ = val
	val_idx := _pidx(val.P)
	_ = val_idx
	pp = _padd(val.P, val.E)
	return PR{P: pp, E: Expr{Kind: ExprKindEkBinding{}, S1: name, S2: "", B1: true, B2: false, C0: val_idx, C1: int64(0), C2: int64(0), List_start: int64(0), List_count: int64(0), Span: span}}
}

func _parse_expr(p Parser, min_bp int64) PR {
	result := _parse_prefix(p)
	_ = result
	pp := result.P
	_ = pp
	left := result.E
	_ = left
	running := true
	_ = running
	for running {
		kind := _cur(pp).Kind
		_ = kind
		post_bp := postfix_binding_power(kind)
		_ = post_bp
		if ((post_bp > int64(0)) && (post_bp >= min_bp)) {
			post := _parse_postfix(pp, left)
			_ = post
			pp = post.P
			left = post.E
		} else {
			bp := infix_binding_power(kind)
			_ = bp
			if ((bp.Left > int64(0)) && (bp.Left >= min_bp)) {
				op := _cur(pp).Text
				_ = op
				span := _pspan(pp)
				_ = span
				left_idx := _pidx(pp)
				_ = left_idx
				pp = _padd(pp, left)
				pp = _adv(pp)
				right := _parse_expr(pp, bp.Right)
				_ = right
				pp = right.P
				right_idx := _pidx(pp)
				_ = right_idx
				pp = _padd(pp, right.E)
				left = Expr{Kind: ExprKindEkBinary{}, S1: op, S2: "", B1: false, B2: false, C0: left_idx, C1: right_idx, C2: int64(0), List_start: int64(0), List_count: int64(0), Span: span}
			} else {
				running = false
			}
		}
	}
	return PR{P: pp, E: left}
}

func _parse_prefix(p Parser) PR {
	k := _ck(p)
	_ = k
	if (k == "@") {
		pp := _adv(p)
		_ = pp
		if (_ck(pp) == "IDENT") {
			pp = _adv(pp)
		}
		if (_ck(pp) == "(") {
			pp = _adv(pp)
			for ((_ck(pp) != ")") && (_ck(pp) != "EOF")) {
				pp = _adv(pp)
			}
			if (_ck(pp) == ")") {
				pp = _adv(pp)
			}
		}
		pp = _skip_nl(pp)
		return _parse_prefix(pp)
	}
	if (((k == "-") || (k == "!")) || (k == "~")) {
		op := _cur(p).Text
		_ = op
		span := _pspan(p)
		_ = span
		pp := _adv(p)
		_ = pp
		bp := prefix_binding_power(_cur(p).Kind)
		_ = bp
		r := _parse_expr(pp, bp)
		_ = r
		operand_idx := _pidx(r.P)
		_ = operand_idx
		pp = _padd(r.P, r.E)
		return PR{P: pp, E: Expr{Kind: ExprKindEkUnary{}, S1: op, S2: "", B1: false, B2: false, C0: operand_idx, C1: int64(0), C2: int64(0), List_start: int64(0), List_count: int64(0), Span: span}}
	}
	if (k == "INT") {
		tok := _cur(p)
		_ = tok
		return PR{P: _adv(p), E: mk_int_lit(tok.Text, _pspan(p))}
	}
	if (k == "FLOAT") {
		tok := _cur(p)
		_ = tok
		return PR{P: _adv(p), E: mk_float_lit(tok.Text, _pspan(p))}
	}
	if (k == "DURATION") {
		tok := _cur(p)
		_ = tok
		return PR{P: _adv(p), E: mk_duration_lit(tok.Text, _pspan(p))}
	}
	if (k == "SIZE") {
		tok := _cur(p)
		_ = tok
		return PR{P: _adv(p), E: mk_size_lit(tok.Text, _pspan(p))}
	}
	if (((k == "STRING") || (k == "RAW_STRING")) || (k == "TRIPLE_STRING")) {
		tok := _cur(p)
		_ = tok
		return PR{P: _adv(p), E: mk_string_lit(tok.Text, _pspan(p))}
	}
	if (k == "STRING_START") {
		span := _pspan(p)
		_ = span
		pp := _adv(p)
		_ = pp
		for ((_ck(pp) != "STRING_END") && (_ck(pp) != "EOF")) {
			if (_ck(pp) == "STRING_MIDDLE") {
				pp = _adv(pp)
			} else {
				r := _parse_expr(pp, int64(0))
				_ = r
				pp = r.P
			}
		}
		if (_ck(pp) == "STRING_END") {
			pp = _adv(pp)
		}
		return PR{P: pp, E: mk_string_lit("", span)}
	}
	if (k == "true") {
		return PR{P: _adv(p), E: mk_bool_lit(true, _pspan(p))}
	}
	if (k == "false") {
		return PR{P: _adv(p), E: mk_bool_lit(false, _pspan(p))}
	}
	if (k == "if") {
		return _parse_if(p)
	}
	if (k == "match") {
		return _parse_match(p)
	}
	if (k == "select") {
		return _parse_select(p)
	}
	if (k == "fn") {
		return _parse_closure_expr(p)
	}
	if (k == "(") {
		return _parse_grouped(p)
	}
	if (k == "[") {
		return _parse_array(p)
	}
	lb := "{"
	_ = lb
	if (k == lb) {
		return _parse_brace_expr(p)
	}
	if (((k == "IDENT") || (k == "Self")) || (k == "self")) {
		name := _cur(p).Text
		_ = name
		span := _pspan(p)
		_ = span
		pp := _adv(p)
		_ = pp
		lb := "{"
		_ = lb
		if ((_ck(pp) == lb) && _is_upper_start(name)) {
			return _parse_struct_lit(pp, name, span)
		}
		if ((_ck(pp) == ":") && (_pk(pp) == ":")) {
			pp = _adv(pp)
			pp = _adv(pp)
			if (_ck(pp) == "IDENT") {
				meth := _cur(pp).Text
				_ = meth
				pp = _adv(pp)
				return PR{P: pp, E: mk_ident(((name + "_") + meth), span)}
			}
		}
		return PR{P: pp, E: mk_ident(name, span)}
	}
	pp := _perror(p, E0005, ("expected expression, got " + k))
	_ = pp
	return PR{P: _adv(pp), E: _no_expr()}
}

func _parse_postfix(p Parser, left Expr) PR {
	k := _ck(p)
	_ = k
	if (k == ".") {
		span := _pspan(p)
		_ = span
		left_idx := _pidx(p)
		_ = left_idx
		pp := _padd(p, left)
		_ = pp
		pp = _adv(pp)
		lb := "{"
		_ = lb
		rb := "}"
		_ = rb
		if (_ck(pp) == lb) {
			pp = _adv(pp)
			pp = _skip_nl(pp)
			list_start := _pidx(pp)
			_ = list_start
			count := int64(0)
			_ = count
			for ((_ck(pp) != rb) && (_ck(pp) != "EOF")) {
				if (_ck(pp) == "IDENT") {
					pp = _adv(pp)
				}
				if (_ck(pp) == ":") {
					pp = _adv(pp)
					r := _parse_expr(pp, int64(0))
					_ = r
					pp = _padd(r.P, r.E)
				}
				count = (count + int64(1))
				pp = _skip_nl(pp)
				if (_ck(pp) == ",") {
					pp = _adv(pp)
					pp = _skip_nl(pp)
				}
			}
			if (_ck(pp) == rb) {
				pp = _adv(pp)
			}
			return PR{P: pp, E: Expr{Kind: ExprKindEkStruct{}, S1: "_update", S2: "", B1: false, B2: false, C0: left_idx, C1: int64(0), C2: int64(0), List_start: list_start, List_count: count, Span: span}}
		}
		field := _cur(pp).Text
		_ = field
		pp = _adv(pp)
		if ((_ck(pp) == "[") && (((field == "to") || (field == "trunc")))) {
			pp = _adv(pp)
			type_arg := ""
			_ = type_arg
			if (_ck(pp) == "IDENT") {
				type_arg = _cur(pp).Text
				pp = _adv(pp)
			}
			if (_ck(pp) == "]") {
				pp = _adv(pp)
			}
			if (_ck(pp) == "(") {
				pp = _adv(pp)
			}
			if (_ck(pp) == ")") {
				pp = _adv(pp)
			}
			return PR{P: pp, E: Expr{Kind: ExprKindEkMethodCall{}, S1: field, S2: type_arg, B1: false, B2: false, C0: left_idx, C1: int64(0), C2: int64(0), List_start: int64(0), List_count: int64(0), Span: span}}
		}
		if (_ck(pp) == "(") {
			pp = _adv(pp)
			mc_arg_indices := []int64{int64(0)}
			_ = mc_arg_indices
			count := int64(0)
			_ = count
			for ((_ck(pp) != ")") && (_ck(pp) != "EOF")) {
				pp = _skip_nl(pp)
				r := _parse_expr(pp, int64(0))
				_ = r
				arg_idx := _pidx(r.P)
				_ = arg_idx
				pp = _padd(r.P, r.E)
				mc_arg_indices = append(mc_arg_indices, arg_idx)
				count = (count + int64(1))
				pp = _skip_nl(pp)
				if (_ck(pp) == ",") {
					pp = _adv(pp)
				}
			}
			pp = _expect(pp, ")")
			list_start := _pidx(pp)
			_ = list_start
			mci := int64(0)
			_ = mci
			for (mci < count) {
				pp = _padd(pp, Expr{Kind: ExprKindEkNone{}, S1: "", S2: "", B1: false, B2: false, C0: mc_arg_indices[(mci + int64(1))], C1: int64(0), C2: int64(0), List_start: int64(0), List_count: int64(0), Span: _no_span()})
				mci = (mci + int64(1))
			}
			return PR{P: pp, E: Expr{Kind: ExprKindEkMethodCall{}, S1: field, S2: "", B1: false, B2: false, C0: left_idx, C1: int64(0), C2: int64(0), List_start: list_start, List_count: count, Span: span}}
		}
		return PR{P: pp, E: Expr{Kind: ExprKindEkFieldAccess{}, S1: field, S2: "", B1: false, B2: false, C0: left_idx, C1: int64(0), C2: int64(0), List_start: int64(0), List_count: int64(0), Span: span}}
	}
	if (k == "(") {
		span := _pspan(p)
		_ = span
		callee_idx := _pidx(p)
		_ = callee_idx
		pp := _padd(p, left)
		_ = pp
		pp = _adv(pp)
		call_arg_indices := []int64{int64(0)}
		_ = call_arg_indices
		count := int64(0)
		_ = count
		for ((_ck(pp) != ")") && (_ck(pp) != "EOF")) {
			pp = _skip_nl(pp)
			r := _parse_expr(pp, int64(0))
			_ = r
			arg_idx := _pidx(r.P)
			_ = arg_idx
			pp = _padd(r.P, r.E)
			call_arg_indices = append(call_arg_indices, arg_idx)
			count = (count + int64(1))
			pp = _skip_nl(pp)
			if (_ck(pp) == ",") {
				pp = _adv(pp)
			}
		}
		pp = _expect(pp, ")")
		list_start := _pidx(pp)
		_ = list_start
		cai := int64(0)
		_ = cai
		for (cai < count) {
			pp = _padd(pp, Expr{Kind: ExprKindEkNone{}, S1: "", S2: "", B1: false, B2: false, C0: call_arg_indices[(cai + int64(1))], C1: int64(0), C2: int64(0), List_start: int64(0), List_count: int64(0), Span: _no_span()})
			cai = (cai + int64(1))
		}
		return PR{P: pp, E: Expr{Kind: ExprKindEkCall{}, S1: left.S1, S2: "", B1: false, B2: false, C0: callee_idx, C1: int64(0), C2: int64(0), List_start: list_start, List_count: count, Span: span}}
	}
	if (k == "[") {
		span := _pspan(p)
		_ = span
		obj_idx := _pidx(p)
		_ = obj_idx
		pp := _padd(p, left)
		_ = pp
		pp = _adv(pp)
		idx_r := _parse_expr(pp, int64(0))
		_ = idx_r
		idx_idx := _pidx(idx_r.P)
		_ = idx_idx
		pp = _padd(idx_r.P, idx_r.E)
		pp = _expect(pp, "]")
		return PR{P: pp, E: Expr{Kind: ExprKindEkIndex{}, S1: "", S2: "", B1: false, B2: false, C0: obj_idx, C1: idx_idx, C2: int64(0), List_start: int64(0), List_count: int64(0), Span: span}}
	}
	if (k == "?") {
		inner_idx := _pidx(p)
		_ = inner_idx
		pp := _padd(p, left)
		_ = pp
		pp = _adv(pp)
		if (_ck(pp) == "|") {
			pp = _adv(pp)
			err_name := "_"
			_ = err_name
			if (_ck(pp) == "IDENT") {
				err_name = _cur(pp).Text
				pp = _adv(pp)
			}
			if (_ck(pp) == "|") {
				pp = _adv(pp)
			}
			body := _parse_expr(pp, int64(0))
			_ = body
			body_idx := _pidx(body.P)
			_ = body_idx
			pp = _padd(body.P, body.E)
			return PR{P: pp, E: Expr{Kind: ExprKindEkPropagate{}, S1: err_name, S2: "", B1: true, B2: false, C0: inner_idx, C1: body_idx, C2: int64(0), List_start: int64(0), List_count: int64(0), Span: _pspan(p)}}
		}
		return PR{P: pp, E: Expr{Kind: ExprKindEkPropagate{}, S1: "", S2: "", B1: false, B2: false, C0: inner_idx, C1: int64(0), C2: int64(0), List_start: int64(0), List_count: int64(0), Span: _pspan(p)}}
	}
	if (k == "!") {
		inner_idx := _pidx(p)
		_ = inner_idx
		pp := _padd(p, left)
		_ = pp
		pp = _adv(pp)
		return PR{P: pp, E: Expr{Kind: ExprKindEkAssertOk{}, S1: "", S2: "", B1: false, B2: false, C0: inner_idx, C1: int64(0), C2: int64(0), List_start: int64(0), List_count: int64(0), Span: _pspan(p)}}
	}
	if (k == "catch") {
		span := _pspan(p)
		_ = span
		inner_idx := _pidx(p)
		_ = inner_idx
		pp := _padd(p, left)
		_ = pp
		pp = _adv(pp)
		if (_ck(pp) == "|") {
			pp = _adv(pp)
			if (_ck(pp) == "IDENT") {
				pp = _adv(pp)
			}
			if (_ck(pp) == "|") {
				pp = _adv(pp)
			}
		}
		body := _parse_block(pp)
		_ = body
		body_idx := _pidx(body.P)
		_ = body_idx
		pp = _padd(body.P, body.E)
		return PR{P: pp, E: Expr{Kind: ExprKindEkCatch{}, S1: "", S2: "", B1: false, B2: false, C0: inner_idx, C1: body_idx, C2: int64(0), List_start: int64(0), List_count: int64(0), Span: span}}
	}
	return PR{P: p, E: left}
}

func _parse_block(p Parser) PR {
	span := _pspan(p)
	_ = span
	lb := "{"
	_ = lb
	rb := "}"
	_ = rb
	if (_ck(p) != lb) {
		pp := _perror(p, E0018, "expected block")
		_ = pp
		return PR{P: pp, E: _no_expr()}
	}
	pp := _adv(p)
	_ = pp
	pp = _skip_nl(pp)
	stmt_indices := []int64{int64(0)}
	_ = stmt_indices
	count := int64(0)
	_ = count
	for ((_ck(pp) != rb) && (_ck(pp) != "EOF")) {
		result := _parse_stmt(pp)
		_ = result
		stmt_idx := _pidx(result.P)
		_ = stmt_idx
		pp = _padd(result.P, result.E)
		stmt_indices = append(stmt_indices, stmt_idx)
		count = (count + int64(1))
		pp = _skip_nl(pp)
	}
	pp = _adv(pp)
	list_start := _pidx(pp)
	_ = list_start
	si := int64(0)
	_ = si
	for (si < count) {
		pp = _padd(pp, Expr{Kind: ExprKindEkNone{}, S1: "", S2: "", B1: false, B2: false, C0: stmt_indices[(si + int64(1))], C1: int64(0), C2: int64(0), List_start: int64(0), List_count: int64(0), Span: _no_span()})
		si = (si + int64(1))
	}
	return PR{P: pp, E: Expr{Kind: ExprKindEkBlock{}, S1: "", S2: "", B1: false, B2: false, C0: int64(0), C1: int64(0), C2: int64(0), List_start: list_start, List_count: count, Span: span}}
}

func _parse_if(p Parser) PR {
	span := _pspan(p)
	_ = span
	pp := _adv(p)
	_ = pp
	cond := _parse_expr(pp, int64(0))
	_ = cond
	cond_idx := _pidx(cond.P)
	_ = cond_idx
	pp = _padd(cond.P, cond.E)
	then_br := _parse_block(pp)
	_ = then_br
	then_idx := _pidx(then_br.P)
	_ = then_idx
	pp = _padd(then_br.P, then_br.E)
	has_else := false
	_ = has_else
	else_idx := int64(0)
	_ = else_idx
	if (_ck(pp) == "else") {
		pp = _adv(pp)
		has_else = true
		if (_ck(pp) == "if") {
			r := _parse_if(pp)
			_ = r
			else_idx = _pidx(r.P)
			pp = _padd(r.P, r.E)
		} else {
			r := _parse_block(pp)
			_ = r
			else_idx = _pidx(r.P)
			pp = _padd(r.P, r.E)
		}
	}
	return PR{P: pp, E: Expr{Kind: ExprKindEkIf{}, S1: "", S2: "", B1: has_else, B2: false, C0: cond_idx, C1: then_idx, C2: else_idx, List_start: int64(0), List_count: int64(0), Span: span}}
}

func _parse_match(p Parser) PR {
	span := _pspan(p)
	_ = span
	pp := _adv(p)
	_ = pp
	subject := _parse_expr(pp, int64(0))
	_ = subject
	subj_idx := _pidx(subject.P)
	_ = subj_idx
	pp = _padd(subject.P, subject.E)
	lb := "{"
	_ = lb
	rb := "}"
	_ = rb
	pp = _skip_nl(pp)
	pp = _expect(pp, lb)
	pp = _skip_nl(pp)
	pat_indices := []int64{int64(0)}
	_ = pat_indices
	body_indices := []int64{int64(0)}
	_ = body_indices
	count := int64(0)
	_ = count
	for ((_ck(pp) != rb) && (_ck(pp) != "EOF")) {
		pat := _parse_pattern(pp)
		_ = pat
		pat_idx := _pidx(pat.P)
		_ = pat_idx
		pp = _padd(pat.P, pat.E)
		for (_ck(pp) == "|") {
			pp = _adv(pp)
			pp = _skip_nl(pp)
			alt := _parse_pattern(pp)
			_ = alt
			pp = _padd(alt.P, alt.E)
		}
		if (_ck(pp) == "if") {
			pp = _adv(pp)
			g := _parse_expr(pp, int64(0))
			_ = g
			pp = _padd(g.P, g.E)
		}
		pp = _expect(pp, "=>")
		body := _parse_expr(pp, int64(0))
		_ = body
		body_idx := _pidx(body.P)
		_ = body_idx
		pp = _padd(body.P, body.E)
		pat_indices = append(pat_indices, pat_idx)
		body_indices = append(body_indices, body_idx)
		count = (count + int64(1))
		pp = _skip_nl(pp)
	}
	pp = _adv(pp)
	list_start := _pidx(pp)
	_ = list_start
	ai := int64(0)
	_ = ai
	for (ai < count) {
		pp = _padd(pp, Expr{Kind: ExprKindEkNone{}, S1: "", S2: "", B1: false, B2: false, C0: pat_indices[(ai + int64(1))], C1: int64(0), C2: int64(0), List_start: int64(0), List_count: int64(0), Span: _no_span()})
		pp = _padd(pp, Expr{Kind: ExprKindEkNone{}, S1: "", S2: "", B1: false, B2: false, C0: body_indices[(ai + int64(1))], C1: int64(0), C2: int64(0), List_start: int64(0), List_count: int64(0), Span: _no_span()})
		ai = (ai + int64(1))
	}
	return PR{P: pp, E: Expr{Kind: ExprKindEkMatch{}, S1: "", S2: "", B1: false, B2: false, C0: subj_idx, C1: int64(0), C2: int64(0), List_start: list_start, List_count: count, Span: span}}
}

func _skip_parens(p Parser) Parser {
	if (_ck(p) != "(") {
		return p
	}
	pp := _adv(p)
	_ = pp
	d := int64(1)
	_ = d
	for ((d > int64(0)) && (_ck(pp) != "EOF")) {
		if (_ck(pp) == "(") {
			d = (d + int64(1))
		}
		if (_ck(pp) == ")") {
			d = (d - int64(1))
		}
		if (d > int64(0)) {
			pp = _adv(pp)
		}
	}
	return _adv(pp)
}

func _parse_pattern(p Parser) PR {
	k := _ck(p)
	_ = k
	span := _pspan(p)
	_ = span
	lb := "{"
	_ = lb
	if (k == "(") {
		pp := _adv(p)
		_ = pp
		list_start := _pidx(pp)
		_ = list_start
		count := int64(0)
		_ = count
		for ((_ck(pp) != ")") && (_ck(pp) != "EOF")) {
			pp = _skip_nl(pp)
			sub := _parse_pattern(pp)
			_ = sub
			pp = _padd(sub.P, sub.E)
			count = (count + int64(1))
			pp = _skip_nl(pp)
			if (_ck(pp) == ",") {
				pp = _adv(pp)
			}
		}
		pp = _expect(pp, ")")
		return PR{P: pp, E: Expr{Kind: ExprKindEkTuple{}, S1: "", S2: "", B1: false, B2: false, C0: int64(0), C1: int64(0), C2: int64(0), List_start: list_start, List_count: count, Span: span}}
	}
	if (k == "[") {
		pp := _adv(p)
		_ = pp
		list_start := _pidx(pp)
		_ = list_start
		count := int64(0)
		_ = count
		for ((_ck(pp) != "]") && (_ck(pp) != "EOF")) {
			pp = _skip_nl(pp)
			if (_ck(pp) == "..") {
				pp = _adv(pp)
				if (_ck(pp) == "IDENT") {
					rest_name := _cur(pp).Text
					_ = rest_name
					pp = _padd(_adv(pp), mk_ident((".." + rest_name), span))
				} else {
					pp = _padd(pp, mk_ident("..", span))
				}
				count = (count + int64(1))
			} else {
				sub := _parse_pattern(pp)
				_ = sub
				pp = _padd(sub.P, sub.E)
				count = (count + int64(1))
			}
			pp = _skip_nl(pp)
			if (_ck(pp) == ",") {
				pp = _adv(pp)
			}
		}
		pp = _expect(pp, "]")
		return PR{P: pp, E: Expr{Kind: ExprKindEkArray{}, S1: "", S2: "", B1: false, B2: false, C0: int64(0), C1: int64(0), C2: int64(0), List_start: list_start, List_count: count, Span: span}}
	}
	if (k == "INT") {
		tok := _cur(p)
		_ = tok
		return PR{P: _adv(p), E: mk_int_lit(tok.Text, span)}
	}
	if (k == "FLOAT") {
		tok := _cur(p)
		_ = tok
		return PR{P: _adv(p), E: mk_float_lit(tok.Text, span)}
	}
	if (((k == "STRING") || (k == "RAW_STRING")) || (k == "TRIPLE_STRING")) {
		tok := _cur(p)
		_ = tok
		return PR{P: _adv(p), E: mk_string_lit(tok.Text, span)}
	}
	if (k == "true") {
		return PR{P: _adv(p), E: mk_bool_lit(true, span)}
	}
	if (k == "false") {
		return PR{P: _adv(p), E: mk_bool_lit(false, span)}
	}
	if (k == "-") {
		pp := _adv(p)
		_ = pp
		if (_ck(pp) == "INT") {
			tok := _cur(pp)
			_ = tok
			return PR{P: _adv(pp), E: mk_int_lit(("-" + tok.Text), span)}
		}
		if (_ck(pp) == "FLOAT") {
			tok := _cur(pp)
			_ = tok
			return PR{P: _adv(pp), E: mk_float_lit(("-" + tok.Text), span)}
		}
		return PR{P: pp, E: mk_ident("-", span)}
	}
	if ((k == "IDENT") || (k == "Self")) {
		name := _cur(p).Text
		_ = name
		pp := _adv(p)
		_ = pp
		if (name == "_") {
			return PR{P: pp, E: mk_ident("_", span)}
		}
		if (_ck(pp) == "@") {
			pp = _adv(pp)
			sub := _parse_pattern(pp)
			_ = sub
			sub_idx := _pidx(sub.P)
			_ = sub_idx
			pp = _padd(sub.P, sub.E)
			return PR{P: pp, E: Expr{Kind: ExprKindEkBinding{}, S1: name, S2: "@", B1: false, B2: false, C0: sub_idx, C1: int64(0), C2: int64(0), List_start: int64(0), List_count: int64(0), Span: span}}
		}
		if (_ck(pp) == "(") {
			pp = _adv(pp)
			list_start := _pidx(pp)
			_ = list_start
			count := int64(0)
			_ = count
			for ((_ck(pp) != ")") && (_ck(pp) != "EOF")) {
				pp = _skip_nl(pp)
				sub := _parse_pattern(pp)
				_ = sub
				pp = _padd(sub.P, sub.E)
				count = (count + int64(1))
				pp = _skip_nl(pp)
				if (_ck(pp) == ",") {
					pp = _adv(pp)
				}
			}
			pp = _expect(pp, ")")
			return PR{P: pp, E: Expr{Kind: ExprKindEkCall{}, S1: name, S2: "", B1: false, B2: false, C0: int64(0), C1: int64(0), C2: int64(0), List_start: list_start, List_count: count, Span: span}}
		}
		if (_ck(pp) == lb) {
			pp = _skip_braces(pp)
			return PR{P: pp, E: Expr{Kind: ExprKindEkCall{}, S1: name, S2: "", B1: false, B2: false, C0: int64(0), C1: int64(0), C2: int64(0), List_start: int64(0), List_count: int64(0), Span: span}}
		}
		if ((_ck(pp) == ".") && (_pk(pp) == "IDENT")) {
			pp = _adv(pp)
			qual := _cur(pp).Text
			_ = qual
			pp = _adv(pp)
			return PR{P: pp, E: mk_ident(((name + ".") + qual), span)}
		}
		if ((_ck(pp) == ":") && (_pk(pp) == ":")) {
			pp = _adv(pp)
			pp = _adv(pp)
			if (_ck(pp) == "IDENT") {
				meth := _cur(pp).Text
				_ = meth
				pp = _adv(pp)
				return PR{P: pp, E: mk_ident(((name + "_") + meth), span)}
			}
		}
		return PR{P: pp, E: mk_ident(name, span)}
	}
	return PR{P: _adv(p), E: mk_ident(_cur(p).Text, span)}
}

func _parse_closure_expr(p Parser) PR {
	span := _pspan(p)
	_ = span
	pp := _adv(p)
	_ = pp
	is_once := false
	_ = is_once
	if ((_ck(pp) == "IDENT") && (_cur(pp).Text == "once")) {
		is_once = true
		pp = _adv(pp)
	}
	pp = _expect(pp, "(")
	params := ""
	_ = params
	for ((_ck(pp) != ")") && (_ck(pp) != "EOF")) {
		if (int64(len(params)) > int64(0)) {
			params = (params + ",")
		}
		params = (params + _cur(pp).Text)
		pp = _adv(pp)
		if (_ck(pp) == ":") {
			pp = _adv(pp)
			pp = _skip_type(pp)
		}
		if (_ck(pp) == ",") {
			pp = _adv(pp)
		}
	}
	pp = _adv(pp)
	if (_ck(pp) == "->") {
		pp = _adv(pp)
		pp = _skip_type(pp)
	}
	lb := "{"
	_ = lb
	body := _no_expr()
	_ = body
	if (_ck(pp) == "=>") {
		pp = _adv(pp)
		result := _parse_expr(pp, int64(0))
		_ = result
		pp = result.P
		body = result.E
	} else if (_ck(pp) == lb) {
		result := _parse_block(pp)
		_ = result
		pp = result.P
		body = result.E
	} else {
		pp = _expect(pp, "=>")
	}
	body_idx := _pidx(pp)
	_ = body_idx
	pp = _padd(pp, body)
	return PR{P: pp, E: Expr{Kind: ExprKindEkClosure{}, S1: params, S2: "", B1: false, B2: is_once, C0: body_idx, C1: int64(0), C2: int64(0), List_start: int64(0), List_count: int64(0), Span: span}}
}

func _parse_grouped(p Parser) PR {
	span := _pspan(p)
	_ = span
	pp := _adv(p)
	_ = pp
	pp = _skip_nl(pp)
	result := _parse_expr(pp, int64(0))
	_ = result
	pp = result.P
	if (_ck(pp) == ",") {
		tup_indices := []int64{int64(0)}
		_ = tup_indices
		first_idx := _pidx(pp)
		_ = first_idx
		pp = _padd(pp, result.E)
		tup_indices = append(tup_indices, first_idx)
		count := int64(1)
		_ = count
		for ((_ck(pp) == ",") && (_ck(pp) != "EOF")) {
			pp = _adv(pp)
			pp = _skip_nl(pp)
			if (_ck(pp) != ")") {
				r := _parse_expr(pp, int64(0))
				_ = r
				elem_idx := _pidx(r.P)
				_ = elem_idx
				pp = _padd(r.P, r.E)
				tup_indices = append(tup_indices, elem_idx)
				count = (count + int64(1))
			}
			pp = _skip_nl(pp)
		}
		if (_ck(pp) == ")") {
			pp = _adv(pp)
		}
		list_start := _pidx(pp)
		_ = list_start
		tei := int64(0)
		_ = tei
		for (tei < count) {
			pp = _padd(pp, Expr{Kind: ExprKindEkNone{}, S1: "", S2: "", B1: false, B2: false, C0: tup_indices[(tei + int64(1))], C1: int64(0), C2: int64(0), List_start: int64(0), List_count: int64(0), Span: _no_span()})
			tei = (tei + int64(1))
		}
		return PR{P: pp, E: Expr{Kind: ExprKindEkTuple{}, S1: "", S2: "", B1: false, B2: false, C0: int64(0), C1: int64(0), C2: int64(0), List_start: list_start, List_count: count, Span: span}}
	}
	for ((_ck(pp) != ")") && (_ck(pp) != "EOF")) {
		pp = _adv(pp)
	}
	pp = _adv(pp)
	return PR{P: pp, E: result.E}
}

func _parse_brace_expr(p Parser) PR {
	span := _pspan(p)
	_ = span
	rb := "}"
	_ = rb
	pp := _adv(p)
	_ = pp
	pp = _skip_nl(pp)
	if (_ck(pp) == rb) {
		pp = _adv(pp)
		return PR{P: pp, E: Expr{Kind: ExprKindEkMapLit{}, S1: "", S2: "", B1: false, B2: false, C0: int64(0), C1: int64(0), C2: int64(0), List_start: int64(0), List_count: int64(0), Span: span}}
	}
	first := _parse_expr(pp, int64(0))
	_ = first
	first_idx := _pidx(first.P)
	_ = first_idx
	pp = _padd(first.P, first.E)
	pp = _skip_nl(pp)
	if (_ck(pp) == ":") {
		pp = _adv(pp)
		val := _parse_expr(pp, int64(0))
		_ = val
		val_idx := _pidx(val.P)
		_ = val_idx
		pp = _padd(val.P, val.E)
		kv_indices := []int64{int64(0)}
		_ = kv_indices
		kv_indices = append(kv_indices, first_idx)
		kv_indices = append(kv_indices, val_idx)
		count := int64(1)
		_ = count
		pp = _skip_nl(pp)
		if (_ck(pp) == ",") {
			pp = _adv(pp)
		}
		for ((_ck(pp) != rb) && (_ck(pp) != "EOF")) {
			pp = _skip_nl(pp)
			kr := _parse_expr(pp, int64(0))
			_ = kr
			ki := _pidx(kr.P)
			_ = ki
			pp = _padd(kr.P, kr.E)
			if (_ck(pp) == ":") {
				pp = _adv(pp)
			}
			vr := _parse_expr(pp, int64(0))
			_ = vr
			vi := _pidx(vr.P)
			_ = vi
			pp = _padd(vr.P, vr.E)
			kv_indices = append(kv_indices, ki)
			kv_indices = append(kv_indices, vi)
			count = (count + int64(1))
			pp = _skip_nl(pp)
			if (_ck(pp) == ",") {
				pp = _adv(pp)
			}
		}
		pp = _adv(pp)
		list_start := _pidx(pp)
		_ = list_start
		mi := int64(0)
		_ = mi
		for (mi < (count * int64(2))) {
			pp = _padd(pp, Expr{Kind: ExprKindEkNone{}, S1: "", S2: "", B1: false, B2: false, C0: kv_indices[(mi + int64(1))], C1: int64(0), C2: int64(0), List_start: int64(0), List_count: int64(0), Span: _no_span()})
			mi = (mi + int64(1))
		}
		return PR{P: pp, E: Expr{Kind: ExprKindEkMapLit{}, S1: "", S2: "", B1: false, B2: false, C0: int64(0), C1: int64(0), C2: int64(0), List_start: list_start, List_count: count, Span: span}}
	}
	if (_ck(pp) == ",") {
		pp = _adv(pp)
		set_indices := []int64{int64(0)}
		_ = set_indices
		set_indices = append(set_indices, first_idx)
		count := int64(1)
		_ = count
		for ((_ck(pp) != rb) && (_ck(pp) != "EOF")) {
			pp = _skip_nl(pp)
			r := _parse_expr(pp, int64(0))
			_ = r
			ei := _pidx(r.P)
			_ = ei
			pp = _padd(r.P, r.E)
			set_indices = append(set_indices, ei)
			count = (count + int64(1))
			pp = _skip_nl(pp)
			if (_ck(pp) == ",") {
				pp = _adv(pp)
			}
		}
		pp = _adv(pp)
		list_start := _pidx(pp)
		_ = list_start
		si := int64(0)
		_ = si
		for (si < count) {
			pp = _padd(pp, Expr{Kind: ExprKindEkNone{}, S1: "", S2: "", B1: false, B2: false, C0: set_indices[(si + int64(1))], C1: int64(0), C2: int64(0), List_start: int64(0), List_count: int64(0), Span: _no_span()})
			si = (si + int64(1))
		}
		return PR{P: pp, E: Expr{Kind: ExprKindEkSetLit{}, S1: "", S2: "", B1: false, B2: false, C0: int64(0), C1: int64(0), C2: int64(0), List_start: list_start, List_count: count, Span: span}}
	}
	pp = _skip_nl(pp)
	if (_ck(pp) == rb) {
		pp = _adv(pp)
	}
	list_start := _pidx(pp)
	_ = list_start
	pp = _padd(pp, Expr{Kind: ExprKindEkNone{}, S1: "", S2: "", B1: false, B2: false, C0: first_idx, C1: int64(0), C2: int64(0), List_start: int64(0), List_count: int64(0), Span: _no_span()})
	return PR{P: pp, E: Expr{Kind: ExprKindEkBlock{}, S1: "", S2: "", B1: false, B2: false, C0: int64(0), C1: int64(0), C2: int64(0), List_start: list_start, List_count: int64(1), Span: span}}
}

func _parse_array(p Parser) PR {
	span := _pspan(p)
	_ = span
	pp := _adv(p)
	_ = pp
	pp = _skip_nl(pp)
	if (_ck(pp) == "]") {
		pp = _adv(pp)
		return PR{P: pp, E: Expr{Kind: ExprKindEkArray{}, S1: "", S2: "", B1: false, B2: false, C0: int64(0), C1: int64(0), C2: int64(0), List_start: int64(0), List_count: int64(0), Span: span}}
	}
	first := _parse_expr(pp, int64(0))
	_ = first
	first_idx := _pidx(first.P)
	_ = first_idx
	pp = _padd(first.P, first.E)
	pp = _skip_nl(pp)
	if (_ck(pp) == "for") {
		pp = _adv(pp)
		var_name := ""
		_ = var_name
		if (_ck(pp) == "IDENT") {
			var_name = _cur(pp).Text
			pp = _adv(pp)
		}
		if (_ck(pp) == "in") {
			pp = _adv(pp)
		}
		iter := _parse_expr(pp, int64(0))
		_ = iter
		iter_idx := _pidx(iter.P)
		_ = iter_idx
		pp = _padd(iter.P, iter.E)
		pp = _skip_nl(pp)
		where_idx := int64(0)
		_ = where_idx
		has_where := false
		_ = has_where
		if ((_ck(pp) == "where") || (_ck(pp) == "if")) {
			pp = _adv(pp)
			wexpr := _parse_expr(pp, int64(0))
			_ = wexpr
			where_idx = _pidx(wexpr.P)
			pp = _padd(wexpr.P, wexpr.E)
			has_where = true
		}
		pp = _skip_nl(pp)
		pp = _expect(pp, "]")
		return PR{P: pp, E: Expr{Kind: ExprKindEkListComp{}, S1: var_name, S2: "", B1: has_where, B2: false, C0: first_idx, C1: iter_idx, C2: where_idx, List_start: int64(0), List_count: int64(0), Span: span}}
	}
	arr_elem_indices := []int64{int64(0)}
	_ = arr_elem_indices
	arr_elem_indices = append(arr_elem_indices, first_idx)
	count := int64(1)
	_ = count
	if (_ck(pp) == ",") {
		pp = _adv(pp)
	}
	for ((_ck(pp) != "]") && (_ck(pp) != "EOF")) {
		pp = _skip_nl(pp)
		r := _parse_expr(pp, int64(0))
		_ = r
		elem_idx := _pidx(r.P)
		_ = elem_idx
		pp = _padd(r.P, r.E)
		arr_elem_indices = append(arr_elem_indices, elem_idx)
		count = (count + int64(1))
		pp = _skip_nl(pp)
		if (_ck(pp) == ",") {
			pp = _adv(pp)
		}
	}
	pp = _expect(pp, "]")
	list_start := _pidx(pp)
	_ = list_start
	aei := int64(0)
	_ = aei
	for (aei < count) {
		pp = _padd(pp, Expr{Kind: ExprKindEkNone{}, S1: "", S2: "", B1: false, B2: false, C0: arr_elem_indices[(aei + int64(1))], C1: int64(0), C2: int64(0), List_start: int64(0), List_count: int64(0), Span: _no_span()})
		aei = (aei + int64(1))
	}
	return PR{P: pp, E: Expr{Kind: ExprKindEkArray{}, S1: "", S2: "", B1: false, B2: false, C0: int64(0), C1: int64(0), C2: int64(0), List_start: list_start, List_count: count, Span: span}}
}

func _parse_struct_lit(p Parser, name string, span Span) PR {
	pp := _adv(p)
	_ = pp
	pp = _skip_nl(pp)
	sf_indices := []int64{int64(0)}
	_ = sf_indices
	count := int64(0)
	_ = count
	rb := "}"
	_ = rb
	for ((_ck(pp) != rb) && (_ck(pp) != "EOF")) {
		pp = _adv(pp)
		if (_ck(pp) == ":") {
			pp = _adv(pp)
			r := _parse_expr(pp, int64(0))
			_ = r
			fval_idx := _pidx(r.P)
			_ = fval_idx
			pp = _padd(r.P, r.E)
			sf_indices = append(sf_indices, fval_idx)
		} else {
			sf_indices = append(sf_indices, int64(0))
		}
		count = (count + int64(1))
		pp = _skip_nl(pp)
		if (_ck(pp) == ",") {
			pp = _adv(pp)
			pp = _skip_nl(pp)
		}
	}
	pp = _adv(pp)
	list_start := _pidx(pp)
	_ = list_start
	sfi := int64(0)
	_ = sfi
	for (sfi < count) {
		pp = _padd(pp, Expr{Kind: ExprKindEkNone{}, S1: "", S2: "", B1: false, B2: false, C0: sf_indices[(sfi + int64(1))], C1: int64(0), C2: int64(0), List_start: int64(0), List_count: int64(0), Span: _no_span()})
		sfi = (sfi + int64(1))
	}
	return PR{P: pp, E: Expr{Kind: ExprKindEkStruct{}, S1: name, S2: "", B1: false, B2: false, C0: int64(0), C1: int64(0), C2: int64(0), List_start: list_start, List_count: count, Span: span}}
}

func _parse_for(p Parser) PR {
	span := _pspan(p)
	_ = span
	pp := _adv(p)
	_ = pp
	pat := _cur(pp).Text
	_ = pat
	pp = _adv(pp)
	pp = _expect(pp, "in")
	iter := _parse_expr(pp, int64(0))
	_ = iter
	iter_idx := _pidx(iter.P)
	_ = iter_idx
	pp = _padd(iter.P, iter.E)
	body := _parse_block(pp)
	_ = body
	body_idx := _pidx(body.P)
	_ = body_idx
	pp = _padd(body.P, body.E)
	return PR{P: pp, E: Expr{Kind: ExprKindEkFor{}, S1: pat, S2: "", B1: false, B2: false, C0: iter_idx, C1: body_idx, C2: int64(0), List_start: int64(0), List_count: int64(0), Span: span}}
}

func _parse_while(p Parser) PR {
	span := _pspan(p)
	_ = span
	pp := _adv(p)
	_ = pp
	cond := _parse_expr(pp, int64(0))
	_ = cond
	cond_idx := _pidx(cond.P)
	_ = cond_idx
	pp = _padd(cond.P, cond.E)
	body := _parse_block(pp)
	_ = body
	body_idx := _pidx(body.P)
	_ = body_idx
	pp = _padd(body.P, body.E)
	return PR{P: pp, E: Expr{Kind: ExprKindEkWhile{}, S1: "", S2: "", B1: false, B2: false, C0: cond_idx, C1: body_idx, C2: int64(0), List_start: int64(0), List_count: int64(0), Span: span}}
}

func _parse_loop(p Parser) PR {
	span := _pspan(p)
	_ = span
	pp := _adv(p)
	_ = pp
	body := _parse_block(pp)
	_ = body
	body_idx := _pidx(body.P)
	_ = body_idx
	pp = _padd(body.P, body.E)
	return PR{P: pp, E: Expr{Kind: ExprKindEkLoop{}, S1: "", S2: "", B1: false, B2: false, C0: body_idx, C1: int64(0), C2: int64(0), List_start: int64(0), List_count: int64(0), Span: span}}
}

func _parse_return(p Parser) PR {
	span := _pspan(p)
	_ = span
	pp := _adv(p)
	_ = pp
	k := _ck(pp)
	_ = k
	rb := "}"
	_ = rb
	if (((k == "NEWLINE") || (k == "EOF")) || (k == rb)) {
		return PR{P: pp, E: mk_return(false, span)}
	}
	val := _parse_expr(pp, int64(0))
	_ = val
	val_idx := _pidx(val.P)
	_ = val_idx
	pp = _padd(val.P, val.E)
	return PR{P: pp, E: Expr{Kind: ExprKindEkReturn{}, S1: "", S2: "", B1: true, B2: false, C0: val_idx, C1: int64(0), C2: int64(0), List_start: int64(0), List_count: int64(0), Span: span}}
}

func _parse_with(p Parser) PR {
	span := _pspan(p)
	_ = span
	pp := _adv(p)
	_ = pp
	name := _cur(pp).Text
	_ = name
	pp = _adv(pp)
	pp = _expect(pp, ":=")
	resource := _parse_expr(pp, int64(0))
	_ = resource
	res_idx := _pidx(resource.P)
	_ = res_idx
	pp = _padd(resource.P, resource.E)
	body := _parse_block(pp)
	_ = body
	body_idx := _pidx(body.P)
	_ = body_idx
	pp = _padd(body.P, body.E)
	return PR{P: pp, E: Expr{Kind: ExprKindEkWith{}, S1: name, S2: "", B1: false, B2: false, C0: res_idx, C1: body_idx, C2: int64(0), List_start: int64(0), List_count: int64(0), Span: span}}
}

func _parse_assert(p Parser) PR {
	span := _pspan(p)
	_ = span
	pp := _adv(p)
	_ = pp
	expr := _parse_expr(pp, int64(0))
	_ = expr
	arg_idx := _pidx(expr.P)
	_ = arg_idx
	pp = _padd(expr.P, expr.E)
	return PR{P: pp, E: Expr{Kind: ExprKindEkCall{}, S1: "assert", S2: "", B1: false, B2: false, C0: int64(0), C1: int64(0), C2: int64(0), List_start: arg_idx, List_count: int64(1), Span: span}}
}

func _parse_select(p Parser) PR {
	span := _pspan(p)
	_ = span
	pp := _adv(p)
	_ = pp
	lb := "{"
	_ = lb
	rb := "}"
	_ = rb
	pp = _skip_nl(pp)
	pp = _expect(pp, lb)
	pp = _skip_nl(pp)
	arm_info_indices := []int64{int64(0)}
	_ = arm_info_indices
	body_indices := []int64{int64(0)}
	_ = body_indices
	count := int64(0)
	_ = count
	for ((_ck(pp) != rb) && (_ck(pp) != "EOF")) {
		pp = _skip_nl(pp)
		if (_ck(pp) == rb) {
			break
		}
		if ((_ck(pp) == "_") || (((_ck(pp) == "IDENT") && (_cur(pp).Text == "default")))) {
			pp = _adv(pp)
			pp = _expect(pp, "=>")
			pp = _skip_nl(pp)
			body := _parse_expr(pp, int64(0))
			_ = body
			body_idx := _pidx(body.P)
			_ = body_idx
			pp = _padd(body.P, body.E)
			info_idx := _pidx(pp)
			_ = info_idx
			pp = _padd(pp, Expr{Kind: ExprKindEkNone{}, S1: "_default", S2: "", B1: false, B2: false, C0: int64(0), C1: int64(0), C2: int64(0), List_start: int64(0), List_count: int64(0), Span: _pspan(pp)})
			arm_info_indices = append(arm_info_indices, info_idx)
			body_indices = append(body_indices, body_idx)
			count = (count + int64(1))
		} else if (_ck(pp) == "IDENT") {
			bname := _cur(pp).Text
			_ = bname
			bspan := _pspan(pp)
			_ = bspan
			pp = _adv(pp)
			if ((_ck(pp) == "IDENT") && (_cur(pp).Text == "from")) {
				pp = _adv(pp)
			}
			ch := _parse_expr(pp, int64(0))
			_ = ch
			ch_idx := _pidx(ch.P)
			_ = ch_idx
			pp = _padd(ch.P, ch.E)
			pp = _expect(pp, "=>")
			pp = _skip_nl(pp)
			body := _parse_expr(pp, int64(0))
			_ = body
			body_idx := _pidx(body.P)
			_ = body_idx
			pp = _padd(body.P, body.E)
			info_idx := _pidx(pp)
			_ = info_idx
			pp = _padd(pp, Expr{Kind: ExprKindEkNone{}, S1: bname, S2: "", B1: false, B2: false, C0: ch_idx, C1: int64(0), C2: int64(0), List_start: int64(0), List_count: int64(0), Span: bspan})
			arm_info_indices = append(arm_info_indices, info_idx)
			body_indices = append(body_indices, body_idx)
			count = (count + int64(1))
		} else {
			pp = _adv(pp)
		}
		pp = _skip_nl(pp)
	}
	if (_ck(pp) == rb) {
		pp = _adv(pp)
	}
	list_start := _pidx(pp)
	_ = list_start
	ai := int64(0)
	_ = ai
	for (ai < count) {
		pp = _padd(pp, Expr{Kind: ExprKindEkNone{}, S1: "", S2: "", B1: false, B2: false, C0: arm_info_indices[(ai + int64(1))], C1: int64(0), C2: int64(0), List_start: int64(0), List_count: int64(0), Span: _no_span()})
		pp = _padd(pp, Expr{Kind: ExprKindEkNone{}, S1: "", S2: "", B1: false, B2: false, C0: body_indices[(ai + int64(1))], C1: int64(0), C2: int64(0), List_start: int64(0), List_count: int64(0), Span: _no_span()})
		ai = (ai + int64(1))
	}
	return PR{P: pp, E: Expr{Kind: ExprKindEkSelect{}, S1: "", S2: "", B1: false, B2: false, C0: int64(0), C1: int64(0), C2: int64(0), List_start: list_start, List_count: count, Span: span}}
}

func parser_diagnostics(p Parser) DiagnosticBag {
	return p.Diagnostics
}

