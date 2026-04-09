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

func generate_from_source(source string, file string, output_path string, runtime_path string) DiagnosticBag {
	lex := tokenize(source, file)
	_ = lex
	toks := get_tokens(lex)
	_ = toks
	lex_diags := get_diagnostics(lex)
	_ = lex_diags
	if bag_has_errors(lex_diags) {
		return lex_diags
	}
	pr := build_decl_index(toks)
	_ = pr
	if bag_has_errors(pr.Diagnostics) {
		return pr.Diagnostics
	}
	rr := resolve(toks, pr.Index, pr.Pool, file)
	_ = rr
	if bag_has_errors(rr.Diagnostics) {
		return rr.Diagnostics
	}
	cr := check(toks, pr.Index, pr.Pool, rr.Table, file)
	_ = cr
	if bag_has_errors(cr.Diagnostics) {
		return cr.Diagnostics
	}
	m := lower(toks, pr.Index, pr.Pool, cr.Store, cr.Registry, cr.Treg, rr.Table, file)
	_ = m
	return generate_from_module(m, output_path, runtime_path)
}

func generate_from_module(m IrModule, output_path string, runtime_path string) DiagnosticBag {
	ir_errors := validate_ir_module(m)
	_ = ir_errors
	ie := int64(1)
	_ = ie
	for (ie < int64(len(ir_errors))) {
		if strings.HasPrefix(ir_errors[ie], "E") {
			fmt.Println(("IR validation: " + ir_errors[ie]))
		}
		ie = (ie + int64(1))
	}
	ll_path := (output_path + ".ll")
	_ = ll_path
	generate_llvm_ir_to_file(m, ll_path)
	cmd := (((((("clang '" + ll_path) + "' '") + runtime_path) + "' -o '") + output_path) + "' -O1 -Wno-override-module 2>&1")
	_ = cmd
	exit_code := _ariaExec(cmd)
	_ = exit_code
	if (exit_code != int64(0)) {
		bag := new_bag()
		_ = bag
		return bag_add_error(bag, "E9000", (("clang compilation failed (exit code " + i2s(exit_code)) + ")"), new_span("<codegen>", int64(0), int64(0), int64(0), int64(0)))
	}
	return new_bag()
}

func _check_duplicate_decls(index DeclIndex) DiagnosticBag {
	bag := new_bag()
	_ = bag
	i := int64(1)
	_ = i
	for (i < int64(len(index.Decls))) {
		di := index.Decls[i]
		_ = di
		dk := decl_kind_name(di.Kind)
		_ = dk
		if (((((dk == "fn") || (dk == "type")) || (dk == "enum")) || (dk == "trait")) || (dk == "const")) {
			if (int64(len(di.Name)) > int64(0)) {
				j := (i + int64(1))
				_ = j
				for (j < int64(len(index.Decls))) {
					dj := index.Decls[j]
					_ = dj
					if ((decl_kind_name(dj.Kind) == dk) && (dj.Name == di.Name)) {
						bag = bag_add_error(bag, "E0702", (((("duplicate " + dk) + " definition '") + di.Name) + "'"), new_span("<multi-file>", int64(0), int64(0), int64(0), int64(0)))
						j = int64(len(index.Decls))
					}
					j = (j + int64(1))
				}
			}
		}
		i = (i + int64(1))
	}
	return bag
}

func _merge_file_tokens(base []Token, new_toks []Token) []Token {
	result := []Token{Token{Kind: TokenKindTkEof{}, Text: "", Line: int64(0), Col: int64(0), Offset: int64(0)}}
	_ = result
	i := int64(1)
	_ = i
	for (i < int64(len(base))) {
		tk := base[i]
		_ = tk
		if (token_name(tk.Kind) == "EOF") {
		} else {
			result = append(result, tk)
		}
		i = (i + int64(1))
	}
	j := int64(1)
	_ = j
	for (j < int64(len(new_toks))) {
		nk := token_name(new_toks[j].Kind)
		_ = nk
		if (nk == "mod") {
			j = (j + int64(1))
			if ((j < int64(len(new_toks))) && (token_name(new_toks[j].Kind) == "IDENT")) {
				j = (j + int64(1))
			}
			for ((j < int64(len(new_toks))) && (token_name(new_toks[j].Kind) == "NEWLINE")) {
				j = (j + int64(1))
			}
		} else if (nk == "use") {
			j = (j + int64(1))
			for (((j < int64(len(new_toks))) && (token_name(new_toks[j].Kind) != "NEWLINE")) && (token_name(new_toks[j].Kind) != "EOF")) {
				j = (j + int64(1))
			}
			for ((j < int64(len(new_toks))) && (token_name(new_toks[j].Kind) == "NEWLINE")) {
				j = (j + int64(1))
			}
		} else if (nk == "NEWLINE") {
			j = (j + int64(1))
		} else {
			return _merge_remaining(result, new_toks, j)
		}
	}
	result = append(result, Token{Kind: TokenKindTkEof{}, Text: "", Line: int64(0), Col: int64(0), Offset: int64(0)})
	return result
}

func _merge_remaining(result []Token, new_toks []Token, start int64) []Token {
	out := result
	_ = out
	j := start
	_ = j
	for (j < int64(len(new_toks))) {
		out = append(out, new_toks[j])
		j = (j + int64(1))
	}
	return out
}

func generate_from_sources(files []string, output_path string, runtime_path string) DiagnosticBag {
	merged := []Token{Token{Kind: TokenKindTkEof{}, Text: "", Line: int64(0), Col: int64(0), Offset: int64(0)}}
	_ = merged
	first := true
	_ = first
	fi := int64(1)
	_ = fi
	for (fi < int64(len(files))) {
		file := files[fi]
		_ = file
		source := _ariaReadFile(file)
		_ = source
		lex := tokenize(source, file)
		_ = lex
		toks := get_tokens(lex)
		_ = toks
		lex_diags := get_diagnostics(lex)
		_ = lex_diags
		if bag_has_errors(lex_diags) {
			return lex_diags
		}
		if first {
			merged = toks
			first = false
		} else {
			merged = _merge_file_tokens(merged, toks)
		}
		fi = (fi + int64(1))
	}
	file := files[int64(1)]
	_ = file
	pr := build_decl_index(merged)
	_ = pr
	if bag_has_errors(pr.Diagnostics) {
		return pr.Diagnostics
	}
	dup_bag := _check_duplicate_decls(pr.Index)
	_ = dup_bag
	if bag_has_errors(dup_bag) {
		return dup_bag
	}
	rr := resolve(merged, pr.Index, pr.Pool, file)
	_ = rr
	if bag_has_errors(rr.Diagnostics) {
		return rr.Diagnostics
	}
	cr := check(merged, pr.Index, pr.Pool, rr.Table, file)
	_ = cr
	if bag_has_errors(cr.Diagnostics) {
		return cr.Diagnostics
	}
	m := lower(merged, pr.Index, pr.Pool, cr.Store, cr.Registry, cr.Treg, rr.Table, file)
	_ = m
	return generate_from_module(m, output_path, runtime_path)
}

func check_sources(files []string) DiagnosticBag {
	merged := []Token{Token{Kind: TokenKindTkEof{}, Text: "", Line: int64(0), Col: int64(0), Offset: int64(0)}}
	_ = merged
	first := true
	_ = first
	fi := int64(1)
	_ = fi
	for (fi < int64(len(files))) {
		file := files[fi]
		_ = file
		source := _ariaReadFile(file)
		_ = source
		lex := tokenize(source, file)
		_ = lex
		toks := get_tokens(lex)
		_ = toks
		lex_diags := get_diagnostics(lex)
		_ = lex_diags
		if bag_has_errors(lex_diags) {
			return lex_diags
		}
		if first {
			merged = toks
			first = false
		} else {
			merged = _merge_file_tokens(merged, toks)
		}
		fi = (fi + int64(1))
	}
	file := files[int64(1)]
	_ = file
	pr := build_decl_index(merged)
	_ = pr
	if bag_has_errors(pr.Diagnostics) {
		return pr.Diagnostics
	}
	rr := resolve(merged, pr.Index, pr.Pool, file)
	_ = rr
	if bag_has_errors(rr.Diagnostics) {
		return rr.Diagnostics
	}
	cr := check(merged, pr.Index, pr.Pool, rr.Table, file)
	_ = cr
	return cr.Diagnostics
}

func generate_tests_from_sources(files []string, output_path string, runtime_path string, parallel bool) DiagnosticBag {
	merged := []Token{Token{Kind: TokenKindTkEof{}, Text: "", Line: int64(0), Col: int64(0), Offset: int64(0)}}
	_ = merged
	first := true
	_ = first
	fi := int64(1)
	_ = fi
	for (fi < int64(len(files))) {
		file := files[fi]
		_ = file
		source := _ariaReadFile(file)
		_ = source
		lex := tokenize(source, file)
		_ = lex
		toks := get_tokens(lex)
		_ = toks
		lex_diags := get_diagnostics(lex)
		_ = lex_diags
		if bag_has_errors(lex_diags) {
			return lex_diags
		}
		if first {
			merged = toks
			first = false
		} else {
			merged = _merge_file_tokens(merged, toks)
		}
		fi = (fi + int64(1))
	}
	file := files[int64(1)]
	_ = file
	pr := build_decl_index(merged)
	_ = pr
	if bag_has_errors(pr.Diagnostics) {
		return pr.Diagnostics
	}
	rr := resolve(merged, pr.Index, pr.Pool, file)
	_ = rr
	if bag_has_errors(rr.Diagnostics) {
		return rr.Diagnostics
	}
	cr := check(merged, pr.Index, pr.Pool, rr.Table, file)
	_ = cr
	if bag_has_errors(cr.Diagnostics) {
		return cr.Diagnostics
	}
	m := lower_for_tests(merged, pr.Index, pr.Pool, cr.Store, cr.Registry, cr.Treg, rr.Table, file, parallel)
	_ = m
	return generate_from_module(m, output_path, runtime_path)
}

func generate_bench_from_sources(files []string, output_path string, runtime_path string) DiagnosticBag {
	merged := []Token{Token{Kind: TokenKindTkEof{}, Text: "", Line: int64(0), Col: int64(0), Offset: int64(0)}}
	_ = merged
	first := true
	_ = first
	fi := int64(1)
	_ = fi
	for (fi < int64(len(files))) {
		file := files[fi]
		_ = file
		source := _ariaReadFile(file)
		_ = source
		lex := tokenize(source, file)
		_ = lex
		toks := get_tokens(lex)
		_ = toks
		lex_diags := get_diagnostics(lex)
		_ = lex_diags
		if bag_has_errors(lex_diags) {
			return lex_diags
		}
		if first {
			merged = toks
			first = false
		} else {
			merged = _merge_file_tokens(merged, toks)
		}
		fi = (fi + int64(1))
	}
	file := files[int64(1)]
	_ = file
	pr := build_decl_index(merged)
	_ = pr
	if bag_has_errors(pr.Diagnostics) {
		return pr.Diagnostics
	}
	rr := resolve(merged, pr.Index, pr.Pool, file)
	_ = rr
	if bag_has_errors(rr.Diagnostics) {
		return rr.Diagnostics
	}
	cr := check(merged, pr.Index, pr.Pool, rr.Table, file)
	_ = cr
	if bag_has_errors(cr.Diagnostics) {
		return cr.Diagnostics
	}
	m := lower_for_tests(merged, pr.Index, pr.Pool, cr.Store, cr.Registry, cr.Treg, rr.Table, file, false)
	_ = m
	return generate_from_module(m, output_path, runtime_path)
}

