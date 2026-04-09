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

func _selfcheck(source string) CheckResult {
	lex := tokenize(source, "test.aria")
	_ = lex
	toks := get_tokens(lex)
	_ = toks
	pr := build_decl_index(toks)
	_ = pr
	rr := resolve(toks, pr.Index, pr.Pool, "test.aria")
	_ = rr
	return check(toks, pr.Index, pr.Pool, rr.Table, "test.aria")
}

func _sc_ok(result CheckResult) bool {
	return (bag_has_errors(result.Diagnostics) == false)
}

