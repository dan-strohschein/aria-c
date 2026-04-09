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

func _resolve_source(source string) ResolveResult {
	lex := tokenize(source, "test.aria")
	_ = lex
	toks := get_tokens(lex)
	_ = toks
	pr := build_decl_index(toks)
	_ = pr
	return resolve(toks, pr.Index, pr.Pool, "test.aria")
}

func _has_errors(result ResolveResult) bool {
	return bag_has_errors(result.Diagnostics)
}

func _err_count(result ResolveResult) int64 {
	return bag_error_count(result.Diagnostics)
}

func _warn_count(result ResolveResult) int64 {
	return bag_warning_count(result.Diagnostics)
}

func _has_code(result ResolveResult, code string) bool {
	i := int64(1)
	_ = i
	for (i < int64(len(result.Diagnostics.Diagnostics))) {
		if (result.Diagnostics.Diagnostics[i].Code == code) {
			return true
		}
		i = (i + int64(1))
	}
	return false
}

func _find_sym(result ResolveResult, name string) Symbol {
	i := int64(1)
	_ = i
	for (i < int64(len(result.Table.Symbols))) {
		if (result.Table.Symbols[i].Name == name) {
			return result.Table.Symbols[i]
		}
		i = (i + int64(1))
	}
	return _sentinel_symbol()
}

