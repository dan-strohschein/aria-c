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

type Lexer struct {
	Source string
	File string
	Pos int64
	Line int64
	Col int64
	Tokens []Token
	Diagnostics DiagnosticBag
	Paren_depth int64
	Bracket_depth int64
	Brace_depth int64
	Interp_depth int64
	Interp_brace_depth int64
}

func _is_digit(c string) bool {
	return ((((((((((c == "0") || (c == "1")) || (c == "2")) || (c == "3")) || (c == "4")) || (c == "5")) || (c == "6")) || (c == "7")) || (c == "8")) || (c == "9"))
}

func _is_hex_digit(c string) bool {
	return ((((((((((((_is_digit(c) || (c == "a")) || (c == "b")) || (c == "c")) || (c == "d")) || (c == "e")) || (c == "f")) || (c == "A")) || (c == "B")) || (c == "C")) || (c == "D")) || (c == "E")) || (c == "F"))
}

func _is_octal_digit(c string) bool {
	return ((((((((c == "0") || (c == "1")) || (c == "2")) || (c == "3")) || (c == "4")) || (c == "5")) || (c == "6")) || (c == "7"))
}

func _is_binary_digit(c string) bool {
	return ((c == "0") || (c == "1"))
}

func _is_alpha(c string) bool {
	return (((((((((((((((((((((((((((((((((((((((((((((((((((((c == "a") || (c == "b")) || (c == "c")) || (c == "d")) || (c == "e")) || (c == "f")) || (c == "g")) || (c == "h")) || (c == "i")) || (c == "j")) || (c == "k")) || (c == "l")) || (c == "m")) || (c == "n")) || (c == "o")) || (c == "p")) || (c == "q")) || (c == "r")) || (c == "s")) || (c == "t")) || (c == "u")) || (c == "v")) || (c == "w")) || (c == "x")) || (c == "y")) || (c == "z")) || (c == "A")) || (c == "B")) || (c == "C")) || (c == "D")) || (c == "E")) || (c == "F")) || (c == "G")) || (c == "H")) || (c == "I")) || (c == "J")) || (c == "K")) || (c == "L")) || (c == "M")) || (c == "N")) || (c == "O")) || (c == "P")) || (c == "Q")) || (c == "R")) || (c == "S")) || (c == "T")) || (c == "U")) || (c == "V")) || (c == "W")) || (c == "X")) || (c == "Y")) || (c == "Z")) || (c == "_"))
}

func _is_alpha_num(c string) bool {
	return (_is_alpha(c) || _is_digit(c))
}

func _is_whitespace(c string) bool {
	return (((c == " ") || (c == "\t")) || (c == "\r"))
}

func new_lexer(source string, file string) Lexer {
	return Lexer{Source: source, File: file, Pos: int64(0), Line: int64(1), Col: int64(1), Tokens: []Token{Token{Kind: TokenKindTkEof{}, Text: "", Line: int64(0), Col: int64(0), Offset: int64(0)}}, Diagnostics: new_bag(), Paren_depth: int64(0), Bracket_depth: int64(0), Brace_depth: int64(0), Interp_depth: int64(0), Interp_brace_depth: int64(0)}
}

func _peek(lex Lexer) string {
	if (lex.Pos >= int64(len(lex.Source))) {
		return ""
	}
	return string(lex.Source[lex.Pos])
}

func _peek_next(lex Lexer) string {
	next := (lex.Pos + int64(1))
	_ = next
	if (next >= int64(len(lex.Source))) {
		return ""
	}
	return string(lex.Source[next])
}

func _at_end(lex Lexer) bool {
	return (lex.Pos >= int64(len(lex.Source)))
}

func _advance(lex Lexer) Lexer {
	return Lexer{Source: lex.Source, File: lex.File, Pos: (lex.Pos + int64(1)), Line: lex.Line, Col: (lex.Col + int64(1)), Tokens: lex.Tokens, Diagnostics: lex.Diagnostics, Paren_depth: lex.Paren_depth, Bracket_depth: lex.Bracket_depth, Brace_depth: lex.Brace_depth, Interp_depth: lex.Interp_depth, Interp_brace_depth: lex.Interp_brace_depth}
}

func _add_token(lex Lexer, kind TokenKind, text string, start_line int64, start_col int64, start_offset int64) Lexer {
	tok := new_token(kind, text, start_line, start_col, start_offset)
	_ = tok
	return Lexer{Source: lex.Source, File: lex.File, Pos: lex.Pos, Line: lex.Line, Col: lex.Col, Tokens: append(lex.Tokens, tok), Diagnostics: lex.Diagnostics, Paren_depth: lex.Paren_depth, Bracket_depth: lex.Bracket_depth, Brace_depth: lex.Brace_depth, Interp_depth: lex.Interp_depth, Interp_brace_depth: lex.Interp_brace_depth}
}

func _add_error(lex Lexer, code string, message string) Lexer {
	span := new_span(lex.File, lex.Line, lex.Col, lex.Pos, int64(1))
	_ = span
	return Lexer{Source: lex.Source, File: lex.File, Pos: lex.Pos, Line: lex.Line, Col: lex.Col, Tokens: lex.Tokens, Diagnostics: bag_add_error(lex.Diagnostics, code, message, span), Paren_depth: lex.Paren_depth, Bracket_depth: lex.Bracket_depth, Brace_depth: lex.Brace_depth, Interp_depth: lex.Interp_depth, Interp_brace_depth: lex.Interp_brace_depth}
}

func _in_delimiters(lex Lexer) bool {
	return ((lex.Paren_depth > int64(0)) || (lex.Bracket_depth > int64(0)))
}

func _set_depth(lex Lexer, paren int64, bracket int64, brace int64) Lexer {
	return Lexer{Source: lex.Source, File: lex.File, Pos: lex.Pos, Line: lex.Line, Col: lex.Col, Tokens: lex.Tokens, Diagnostics: lex.Diagnostics, Paren_depth: paren, Bracket_depth: bracket, Brace_depth: brace, Interp_depth: lex.Interp_depth, Interp_brace_depth: lex.Interp_brace_depth}
}

func _set_interp_brace(lex Lexer, bd int64) Lexer {
	return Lexer{Source: lex.Source, File: lex.File, Pos: lex.Pos, Line: lex.Line, Col: lex.Col, Tokens: lex.Tokens, Diagnostics: lex.Diagnostics, Paren_depth: lex.Paren_depth, Bracket_depth: lex.Bracket_depth, Brace_depth: lex.Brace_depth, Interp_depth: lex.Interp_depth, Interp_brace_depth: bd}
}

func _set_interp(lex Lexer, depth int64) Lexer {
	return Lexer{Source: lex.Source, File: lex.File, Pos: lex.Pos, Line: lex.Line, Col: lex.Col, Tokens: lex.Tokens, Diagnostics: lex.Diagnostics, Paren_depth: lex.Paren_depth, Bracket_depth: lex.Bracket_depth, Brace_depth: lex.Brace_depth, Interp_depth: depth, Interp_brace_depth: lex.Interp_brace_depth}
}

func _newline(lex Lexer) Lexer {
	return Lexer{Source: lex.Source, File: lex.File, Pos: (lex.Pos + int64(1)), Line: (lex.Line + int64(1)), Col: int64(1), Tokens: lex.Tokens, Diagnostics: lex.Diagnostics, Paren_depth: lex.Paren_depth, Bracket_depth: lex.Bracket_depth, Brace_depth: lex.Brace_depth, Interp_depth: lex.Interp_depth, Interp_brace_depth: lex.Interp_brace_depth}
}

func _last_token_kind(lex Lexer) TokenKind {
	if (int64(len(lex.Tokens)) <= int64(1)) {
		return TokenKindTkEof{}
	}
	return lex.Tokens[(int64(len(lex.Tokens)) - int64(1))].Kind
}

func _scan_ident(lex Lexer) Lexer {
	start := lex.Pos
	_ = start
	start_col := lex.Col
	_ = start_col
	start_line := lex.Line
	_ = start_line
	l := lex
	_ = l
	for ((_at_end(l) == false) && _is_alpha_num(_peek(l))) {
		l = _advance(l)
	}
	text := l.Source[start:l.Pos]
	_ = text
	kind := lookup_keyword(text)
	_ = kind
	return _add_token(l, kind, text, start_line, start_col, start)
}

func _scan_number(lex Lexer) Lexer {
	start := lex.Pos
	_ = start
	start_col := lex.Col
	_ = start_col
	start_line := lex.Line
	_ = start_line
	l := lex
	_ = l
	is_float := false
	_ = is_float
	c := _peek(l)
	_ = c
	if ((c == "0") && (_at_end(l) == false)) {
		next := _peek_next(l)
		_ = next
		if ((next == "x") || (next == "X")) {
			l = _advance(_advance(l))
			for ((_at_end(l) == false) && ((_is_hex_digit(_peek(l)) || (_peek(l) == "_")))) {
				l = _advance(l)
			}
			text := l.Source[start:l.Pos]
			_ = text
			return _add_token(l, TokenKindTkIntLit{}, text, start_line, start_col, start)
		}
		if ((next == "o") || (next == "O")) {
			l = _advance(_advance(l))
			for ((_at_end(l) == false) && ((_is_octal_digit(_peek(l)) || (_peek(l) == "_")))) {
				l = _advance(l)
			}
			text := l.Source[start:l.Pos]
			_ = text
			return _add_token(l, TokenKindTkIntLit{}, text, start_line, start_col, start)
		}
		if ((next == "b") || (next == "B")) {
			l = _advance(_advance(l))
			for ((_at_end(l) == false) && ((_is_binary_digit(_peek(l)) || (_peek(l) == "_")))) {
				l = _advance(l)
			}
			text := l.Source[start:l.Pos]
			_ = text
			return _add_token(l, TokenKindTkIntLit{}, text, start_line, start_col, start)
		}
	}
	for ((_at_end(l) == false) && ((_is_digit(_peek(l)) || (_peek(l) == "_")))) {
		l = _advance(l)
	}
	if ((_at_end(l) == false) && (_peek(l) == ".")) {
		if ((_peek_next(l) != "") && _is_digit(_peek_next(l))) {
			is_float = true
			l = _advance(l)
			for ((_at_end(l) == false) && ((_is_digit(_peek(l)) || (_peek(l) == "_")))) {
				l = _advance(l)
			}
			if ((_at_end(l) == false) && (((_peek(l) == "e") || (_peek(l) == "E")))) {
				l = _advance(l)
				if ((_at_end(l) == false) && (((_peek(l) == "+") || (_peek(l) == "-")))) {
					l = _advance(l)
				}
				for ((_at_end(l) == false) && _is_digit(_peek(l))) {
					l = _advance(l)
				}
			}
		}
	}
	if ((_at_end(l) == false) && (is_float == false)) {
		sc := _peek(l)
		_ = sc
		if ((sc == "n") && (_peek_next(l) == "s")) {
			l = _advance(_advance(l))
			text := l.Source[start:l.Pos]
			_ = text
			return _add_token(l, TokenKindTkDurationLit{}, text, start_line, start_col, start)
		}
		if ((sc == "u") && (_peek_next(l) == "s")) {
			l = _advance(_advance(l))
			text := l.Source[start:l.Pos]
			_ = text
			return _add_token(l, TokenKindTkDurationLit{}, text, start_line, start_col, start)
		}
		if ((sc == "m") && (_peek_next(l) == "s")) {
			l = _advance(_advance(l))
			text := l.Source[start:l.Pos]
			_ = text
			return _add_token(l, TokenKindTkDurationLit{}, text, start_line, start_col, start)
		}
		if (sc == "s") {
			l = _advance(l)
			text := l.Source[start:l.Pos]
			_ = text
			return _add_token(l, TokenKindTkDurationLit{}, text, start_line, start_col, start)
		}
		if ((sc == "m") && (_peek_next(l) != "s")) {
			nx := _peek_next(l)
			_ = nx
			if ((nx == "") || (_is_alpha(nx) == false)) {
				l = _advance(l)
				text := l.Source[start:l.Pos]
				_ = text
				return _add_token(l, TokenKindTkDurationLit{}, text, start_line, start_col, start)
			}
		}
		if (sc == "h") {
			l = _advance(l)
			text := l.Source[start:l.Pos]
			_ = text
			return _add_token(l, TokenKindTkDurationLit{}, text, start_line, start_col, start)
		}
		if ((sc == "k") && (_peek_next(l) == "b")) {
			l = _advance(_advance(l))
			text := l.Source[start:l.Pos]
			_ = text
			return _add_token(l, TokenKindTkSizeLit{}, text, start_line, start_col, start)
		}
		if ((sc == "m") && (_peek_next(l) == "b")) {
			l = _advance(_advance(l))
			text := l.Source[start:l.Pos]
			_ = text
			return _add_token(l, TokenKindTkSizeLit{}, text, start_line, start_col, start)
		}
		if ((sc == "g") && (_peek_next(l) == "b")) {
			l = _advance(_advance(l))
			text := l.Source[start:l.Pos]
			_ = text
			return _add_token(l, TokenKindTkSizeLit{}, text, start_line, start_col, start)
		}
		if (sc == "b") {
			nx := _peek_next(l)
			_ = nx
			if ((nx == "") || (((_is_alpha(nx) == false) && (_is_digit(nx) == false)))) {
				l = _advance(l)
				text := l.Source[start:l.Pos]
				_ = text
				return _add_token(l, TokenKindTkSizeLit{}, text, start_line, start_col, start)
			}
		}
	}
	text := l.Source[start:l.Pos]
	_ = text
	if is_float {
		return _add_token(l, TokenKindTkFloatLit{}, text, start_line, start_col, start)
	}
	return _add_token(l, TokenKindTkIntLit{}, text, start_line, start_col, start)
}

func _scan_char(lex Lexer) Lexer {
	start_line := lex.Line
	_ = start_line
	start_col := lex.Col
	_ = start_col
	start_pos := lex.Pos
	_ = start_pos
	l := _advance(lex)
	_ = l
	val := ""
	_ = val
	if (_peek(l) == "\\") {
		l = _advance(l)
		esc := _peek(l)
		_ = esc
		if (esc == "n") {
			val = "\n"
		} else if (esc == "t") {
			val = "\t"
		} else if (esc == "r") {
			val = "\r"
		} else if (esc == "\\") {
			val = "\\"
		} else if (esc == "'") {
			val = "'"
		} else if (esc == "0") {
			val = "\x00"
		} else {
			val = esc
		}
		l = _advance(l)
	} else {
		val = _peek(l)
		l = _advance(l)
	}
	if (_peek(l) == "'") {
		l = _advance(l)
	}
	return _add_token(l, TokenKindTkCharLit{}, val, start_line, start_col, start_pos)
}

func _scan_string(lex Lexer) Lexer {
	return _scan_string_impl(lex, false)
}

func _scan_string_impl(lex Lexer, is_continuation bool) Lexer {
	start := lex.Pos
	_ = start
	start_col := lex.Col
	_ = start_col
	start_line := lex.Line
	_ = start_line
	l := lex
	_ = l
	if (is_continuation == false) {
		l = _advance(l)
	}
	text := ""
	_ = text
	lb := "{"
	_ = lb
	rb := "}"
	_ = rb
	for ((_at_end(l) == false) && (_peek(l) != "\"")) {
		c := _peek(l)
		_ = c
		if (c == "\\") {
			l = _advance(l)
			if _at_end(l) {
				l = _add_error(l, E0002, "unterminated string literal")
				return l
			}
			esc := _peek(l)
			_ = esc
			if (esc == "n") {
				text = (text + "\n")
			} else if (esc == "r") {
				text = (text + "\r")
			} else if (esc == "t") {
				text = (text + "\t")
			} else if (esc == "\\") {
				text = (text + "\\")
			} else if (esc == "\"") {
				text = (text + "\"")
			} else if (esc == "0") {
				text = (text + "\x00")
			} else if (esc == lb) {
				text = (text + lb)
			} else if (esc == rb) {
				text = (text + rb)
			} else {
				text = (text + esc)
			}
			l = _advance(l)
		} else if (c == lb) {
			if is_continuation {
				l = _add_token(l, TokenKindTkStringMiddle{}, text, start_line, start_col, start)
			} else {
				l = _add_token(l, TokenKindTkStringStart{}, text, start_line, start_col, start)
			}
			l = _advance(l)
			l = _set_interp(l, (l.Interp_depth + int64(1)))
			l = _set_interp_brace(l, l.Brace_depth)
			return l
		} else if (c == "\n") {
			l = _add_error(l, E0002, "unterminated string literal")
			return l
		} else {
			text = (text + c)
			l = _advance(l)
		}
	}
	if _at_end(l) {
		l = _add_error(l, E0002, "unterminated string literal")
		return l
	}
	l = _advance(l)
	return func() interface{} {
		if is_continuation {
			return _add_token(l, TokenKindTkStringEnd{}, text, start_line, start_col, start)
		} else {
			return _add_token(l, TokenKindTkStringLit{}, text, start_line, start_col, start)
		}
	}().(Lexer)
}

func _scan_raw_string(lex Lexer) Lexer {
	start := lex.Pos
	_ = start
	start_col := lex.Col
	_ = start_col
	start_line := lex.Line
	_ = start_line
	l := _advance(lex)
	_ = l
	hash_count := int64(0)
	_ = hash_count
	for ((_at_end(l) == false) && (_peek(l) == "#")) {
		hash_count = (hash_count + int64(1))
		l = _advance(l)
	}
	if (_at_end(l) || (_peek(l) != "\"")) {
		return _add_error(l, E0002, "expected '\"' after r-string prefix")
	}
	l = _advance(l)
	text := ""
	_ = text
	done := false
	_ = done
	for ((done == false) && (_at_end(l) == false)) {
		c := _peek(l)
		_ = c
		if (c == "\"") {
			match_hashes := int64(0)
			_ = match_hashes
			peek_l := _advance(l)
			_ = peek_l
			for (((match_hashes < hash_count) && (_at_end(peek_l) == false)) && (_peek(peek_l) == "#")) {
				match_hashes = (match_hashes + int64(1))
				peek_l = _advance(peek_l)
			}
			if (match_hashes == hash_count) {
				l = peek_l
				done = true
			} else {
				text = (text + c)
				l = _advance(l)
			}
		} else {
			text = (text + c)
			l = _advance(l)
		}
	}
	return _add_token(l, TokenKindTkRawStringLit{}, text, start_line, start_col, start)
}

func _scan_triple_string(lex Lexer) Lexer {
	start := lex.Pos
	_ = start
	start_col := lex.Col
	_ = start_col
	start_line := lex.Line
	_ = start_line
	l := _advance(_advance(_advance(lex)))
	_ = l
	if ((_at_end(l) == false) && (_peek(l) == "\n")) {
		l = _newline(_advance(l))
	}
	text := ""
	_ = text
	done := false
	_ = done
	for ((done == false) && (_at_end(l) == false)) {
		c := _peek(l)
		_ = c
		if (c == "\"") {
			p1 := _peek_next(l)
			_ = p1
			l2 := _advance(l)
			_ = l2
			if (_at_end(l2) == false) {
				l3 := _advance(l2)
				_ = l3
				if (((p1 == "\"") && (_at_end(l3) == false)) && (_peek(l2) == "\"")) {
					l = _advance(l3)
					done = true
				} else {
					text = (text + c)
					l = _advance(l)
				}
			} else {
				text = (text + c)
				l = _advance(l)
			}
		} else if (c == "\n") {
			text = (text + "\n")
			l = _newline(_advance(l))
		} else {
			text = (text + c)
			l = _advance(l)
		}
	}
	return _add_token(l, TokenKindTkTripleStringLit{}, text, start_line, start_col, start)
}

func _scan_comment(lex Lexer) Lexer {
	l := lex
	_ = l
	for ((_at_end(l) == false) && (_peek(l) != "\n")) {
		l = _advance(l)
	}
	return l
}

func tokenize(source string, file string) Lexer {
	lex := new_lexer(source, file)
	_ = lex
	for (_at_end(lex) == false) {
		c := _peek(lex)
		_ = c
		if (c == "\n") {
			if (_in_delimiters(lex) == false) {
				last := _last_token_kind(lex)
				_ = last
				if terminates_statement(last) {
					lex = _add_token(lex, TokenKindTkNewline{}, "\n", lex.Line, lex.Col, lex.Pos)
				}
			}
			lex = _newline(lex)
		} else if _is_whitespace(c) {
			lex = _advance(lex)
		} else if ((c == "/") && (_peek_next(lex) == "/")) {
			lex = _scan_comment(lex)
		} else if ((c == "r") && (_at_end(lex) == false)) {
			nx := _peek_next(lex)
			_ = nx
			if ((nx == "\"") || (nx == "#")) {
				lex = _scan_raw_string(lex)
			} else {
				lex = _scan_ident(lex)
			}
		} else if _is_alpha(c) {
			lex = _scan_ident(lex)
		} else if _is_digit(c) {
			lex = _scan_number(lex)
		} else if (c == "\"") {
			if (_peek_next(lex) == "\"") {
				p3 := (lex.Pos + int64(2))
				_ = p3
				is_triple := false
				_ = is_triple
				if ((p3 < int64(len(lex.Source))) && (string(lex.Source[p3]) == "\"")) {
					is_triple = true
				}
				if is_triple {
					lex = _scan_triple_string(lex)
				} else {
					lex = _scan_string(lex)
				}
			} else {
				lex = _scan_string(lex)
			}
		} else if (c == "'") {
			lex = _scan_char(lex)
		} else if (c == "+") {
			lex = _add_token(lex, TokenKindTkPlus{}, "+", lex.Line, lex.Col, lex.Pos)
			lex = _advance(lex)
		} else if (c == "-") {
			if (_peek_next(lex) == ">") {
				lex = _add_token(lex, TokenKindTkArrow{}, "->", lex.Line, lex.Col, lex.Pos)
				lex = _advance(_advance(lex))
			} else {
				lex = _add_token(lex, TokenKindTkMinus{}, "-", lex.Line, lex.Col, lex.Pos)
				lex = _advance(lex)
			}
		} else if (c == "*") {
			lex = _add_token(lex, TokenKindTkStar{}, "*", lex.Line, lex.Col, lex.Pos)
			lex = _advance(lex)
		} else if (c == "/") {
			lex = _add_token(lex, TokenKindTkSlash{}, "/", lex.Line, lex.Col, lex.Pos)
			lex = _advance(lex)
		} else if (c == "%") {
			lex = _add_token(lex, TokenKindTkPercent{}, "%", lex.Line, lex.Col, lex.Pos)
			lex = _advance(lex)
		} else if (c == "=") {
			if (_peek_next(lex) == "=") {
				lex = _add_token(lex, TokenKindTkEqEq{}, "==", lex.Line, lex.Col, lex.Pos)
				lex = _advance(_advance(lex))
			} else if (_peek_next(lex) == ">") {
				lex = _add_token(lex, TokenKindTkFatArrow{}, "=>", lex.Line, lex.Col, lex.Pos)
				lex = _advance(_advance(lex))
			} else {
				lex = _add_token(lex, TokenKindTkEq{}, "=", lex.Line, lex.Col, lex.Pos)
				lex = _advance(lex)
			}
		} else if (c == "!") {
			if (_peek_next(lex) == "=") {
				lex = _add_token(lex, TokenKindTkBangEq{}, "!=", lex.Line, lex.Col, lex.Pos)
				lex = _advance(_advance(lex))
			} else {
				lex = _add_token(lex, TokenKindTkBang{}, "!", lex.Line, lex.Col, lex.Pos)
				lex = _advance(lex)
			}
		} else if (c == "<") {
			if (_peek_next(lex) == "=") {
				lex = _add_token(lex, TokenKindTkLtEq{}, "<=", lex.Line, lex.Col, lex.Pos)
				lex = _advance(_advance(lex))
			} else if (_peek_next(lex) == "<") {
				lex = _add_token(lex, TokenKindTkLShift{}, "<<", lex.Line, lex.Col, lex.Pos)
				lex = _advance(_advance(lex))
			} else {
				lex = _add_token(lex, TokenKindTkLt{}, "<", lex.Line, lex.Col, lex.Pos)
				lex = _advance(lex)
			}
		} else if (c == ">") {
			if (_peek_next(lex) == "=") {
				lex = _add_token(lex, TokenKindTkGtEq{}, ">=", lex.Line, lex.Col, lex.Pos)
				lex = _advance(_advance(lex))
			} else if (_peek_next(lex) == ">") {
				lex = _add_token(lex, TokenKindTkRShift{}, ">>", lex.Line, lex.Col, lex.Pos)
				lex = _advance(_advance(lex))
			} else {
				lex = _add_token(lex, TokenKindTkGt{}, ">", lex.Line, lex.Col, lex.Pos)
				lex = _advance(lex)
			}
		} else if (c == "&") {
			if (_peek_next(lex) == "&") {
				lex = _add_token(lex, TokenKindTkAmpAmp{}, "&&", lex.Line, lex.Col, lex.Pos)
				lex = _advance(_advance(lex))
			} else {
				lex = _add_token(lex, TokenKindTkAmpersand{}, "&", lex.Line, lex.Col, lex.Pos)
				lex = _advance(lex)
			}
		} else if (c == "|") {
			if (_peek_next(lex) == "|") {
				lex = _add_token(lex, TokenKindTkPipePipe{}, "||", lex.Line, lex.Col, lex.Pos)
				lex = _advance(_advance(lex))
			} else if (_peek_next(lex) == ">") {
				lex = _add_token(lex, TokenKindTkPipeGt{}, "|>", lex.Line, lex.Col, lex.Pos)
				lex = _advance(_advance(lex))
			} else {
				lex = _add_token(lex, TokenKindTkPipe{}, "|", lex.Line, lex.Col, lex.Pos)
				lex = _advance(lex)
			}
		} else if (c == "^") {
			lex = _add_token(lex, TokenKindTkCaret{}, "^", lex.Line, lex.Col, lex.Pos)
			lex = _advance(lex)
		} else if (c == "~") {
			lex = _add_token(lex, TokenKindTkTilde{}, "~", lex.Line, lex.Col, lex.Pos)
			lex = _advance(lex)
		} else if (c == "?") {
			if (_peek_next(lex) == ".") {
				lex = _add_token(lex, TokenKindTkQuestionDot{}, "?.", lex.Line, lex.Col, lex.Pos)
				lex = _advance(_advance(lex))
			} else if (_peek_next(lex) == "?") {
				lex = _add_token(lex, TokenKindTkQuestionQuestion{}, "??", lex.Line, lex.Col, lex.Pos)
				lex = _advance(_advance(lex))
			} else {
				lex = _add_token(lex, TokenKindTkQuestion{}, "?", lex.Line, lex.Col, lex.Pos)
				lex = _advance(lex)
			}
		} else if (c == ".") {
			lb := "{"
			_ = lb
			if (_peek_next(lex) == lb) {
				lex = _add_token(lex, TokenKindTkDotLBrace{}, ".{", lex.Line, lex.Col, lex.Pos)
				lex = _advance(_advance(lex))
			} else if (_peek_next(lex) == ".") {
				lex = _advance(_advance(lex))
				if ((_at_end(lex) == false) && (_peek(lex) == "=")) {
					lex = _add_token(lex, TokenKindTkDotDotEq{}, "..=", lex.Line, (lex.Col - int64(2)), (lex.Pos - int64(2)))
					lex = _advance(lex)
				} else {
					lex = _add_token(lex, TokenKindTkDotDot{}, "..", lex.Line, (lex.Col - int64(2)), (lex.Pos - int64(2)))
				}
			} else {
				lex = _add_token(lex, TokenKindTkDot{}, ".", lex.Line, lex.Col, lex.Pos)
				lex = _advance(lex)
			}
		} else if (c == ":") {
			if (_peek_next(lex) == "=") {
				lex = _add_token(lex, TokenKindTkColonEq{}, ":=", lex.Line, lex.Col, lex.Pos)
				lex = _advance(_advance(lex))
			} else {
				lex = _add_token(lex, TokenKindTkColon{}, ":", lex.Line, lex.Col, lex.Pos)
				lex = _advance(lex)
			}
		} else if (c == ",") {
			lex = _add_token(lex, TokenKindTkComma{}, ",", lex.Line, lex.Col, lex.Pos)
			lex = _advance(lex)
		} else if (c == "@") {
			lex = _add_token(lex, TokenKindTkAt{}, "@", lex.Line, lex.Col, lex.Pos)
			lex = _advance(lex)
		} else if (c == "(") {
			lex = _add_token(lex, TokenKindTkLParen{}, "(", lex.Line, lex.Col, lex.Pos)
			lex = _set_depth(lex, (lex.Paren_depth + int64(1)), lex.Bracket_depth, lex.Brace_depth)
			lex = _advance(lex)
		} else if (c == ")") {
			lex = _add_token(lex, TokenKindTkRParen{}, ")", lex.Line, lex.Col, lex.Pos)
			if (lex.Paren_depth > int64(0)) {
				lex = _set_depth(lex, (lex.Paren_depth - int64(1)), lex.Bracket_depth, lex.Brace_depth)
			}
			lex = _advance(lex)
		} else if (c == "[") {
			lex = _add_token(lex, TokenKindTkLBracket{}, "[", lex.Line, lex.Col, lex.Pos)
			lex = _set_depth(lex, lex.Paren_depth, (lex.Bracket_depth + int64(1)), lex.Brace_depth)
			lex = _advance(lex)
		} else if (c == "]") {
			lex = _add_token(lex, TokenKindTkRBracket{}, "]", lex.Line, lex.Col, lex.Pos)
			if (lex.Bracket_depth > int64(0)) {
				lex = _set_depth(lex, lex.Paren_depth, (lex.Bracket_depth - int64(1)), lex.Brace_depth)
			}
			lex = _advance(lex)
		} else if (c == "{") {
			lex = _add_token(lex, TokenKindTkLBrace{}, "{", lex.Line, lex.Col, lex.Pos)
			lex = _set_depth(lex, lex.Paren_depth, lex.Bracket_depth, (lex.Brace_depth + int64(1)))
			lex = _advance(lex)
		} else if (c == "}") {
			if ((lex.Interp_depth > int64(0)) && (lex.Brace_depth == lex.Interp_brace_depth)) {
				lex = _advance(lex)
				lex = _set_interp(lex, (lex.Interp_depth - int64(1)))
				lex = _scan_string_impl(lex, true)
			} else {
				lex = _add_token(lex, TokenKindTkRBrace{}, "}", lex.Line, lex.Col, lex.Pos)
				if (lex.Brace_depth > int64(0)) {
					lex = _set_depth(lex, lex.Paren_depth, lex.Bracket_depth, (lex.Brace_depth - int64(1)))
				}
				lex = _advance(lex)
			}
		} else {
			lex = _add_error(lex, E0001, "unexpected character")
			lex = _add_token(lex, TokenKindTkIllegal{}, c, lex.Line, lex.Col, lex.Pos)
			lex = _advance(lex)
		}
	}
	lex = _add_token(lex, TokenKindTkEof{}, "", lex.Line, lex.Col, lex.Pos)
	return lex
}

func get_tokens(lex Lexer) []Token {
	return lex.Tokens
}

func get_diagnostics(lex Lexer) DiagnosticBag {
	return lex.Diagnostics
}

