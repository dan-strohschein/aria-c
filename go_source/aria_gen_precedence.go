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

type BindingPower struct {
	Left int64
	Right int64
}

func no_binding_power() BindingPower {
	return BindingPower{Left: int64(0), Right: int64(0)}
}

func infix_binding_power(kind TokenKind) BindingPower {
	name := token_name(kind)
	_ = name
	if (name == ":=") {
		return BindingPower{Left: int64(10), Right: int64(10)}
	}
	if (name == "=") {
		return BindingPower{Left: int64(10), Right: int64(10)}
	}
	if (name == "|>") {
		return BindingPower{Left: int64(20), Right: int64(21)}
	}
	if (name == "??") {
		return BindingPower{Left: int64(30), Right: int64(31)}
	}
	if (name == "||") {
		return BindingPower{Left: int64(40), Right: int64(41)}
	}
	if (name == "&&") {
		return BindingPower{Left: int64(50), Right: int64(51)}
	}
	if (name == "==") {
		return BindingPower{Left: int64(60), Right: int64(60)}
	}
	if (name == "!=") {
		return BindingPower{Left: int64(60), Right: int64(60)}
	}
	if (name == "<") {
		return BindingPower{Left: int64(60), Right: int64(60)}
	}
	if (name == ">") {
		return BindingPower{Left: int64(60), Right: int64(60)}
	}
	if (name == "<=") {
		return BindingPower{Left: int64(60), Right: int64(60)}
	}
	if (name == ">=") {
		return BindingPower{Left: int64(60), Right: int64(60)}
	}
	if (name == "..") {
		return BindingPower{Left: int64(70), Right: int64(70)}
	}
	if (name == "..=") {
		return BindingPower{Left: int64(70), Right: int64(70)}
	}
	if (name == "|") {
		return BindingPower{Left: int64(80), Right: int64(81)}
	}
	if (name == "^") {
		return BindingPower{Left: int64(90), Right: int64(91)}
	}
	if (name == "&") {
		return BindingPower{Left: int64(100), Right: int64(101)}
	}
	if (name == "<<") {
		return BindingPower{Left: int64(110), Right: int64(111)}
	}
	if (name == ">>") {
		return BindingPower{Left: int64(110), Right: int64(111)}
	}
	if (name == "+") {
		return BindingPower{Left: int64(120), Right: int64(121)}
	}
	if (name == "-") {
		return BindingPower{Left: int64(120), Right: int64(121)}
	}
	if (name == "*") {
		return BindingPower{Left: int64(130), Right: int64(131)}
	}
	if (name == "/") {
		return BindingPower{Left: int64(130), Right: int64(131)}
	}
	if (name == "%") {
		return BindingPower{Left: int64(130), Right: int64(131)}
	}
	return no_binding_power()
}

func prefix_binding_power(kind TokenKind) int64 {
	name := token_name(kind)
	_ = name
	if (name == "-") {
		return int64(140)
	}
	if (name == "!") {
		return int64(140)
	}
	if (name == "~") {
		return int64(140)
	}
	return int64(0)
}

func postfix_binding_power(kind TokenKind) int64 {
	name := token_name(kind)
	_ = name
	if (name == "?") {
		return int64(150)
	}
	if (name == "!") {
		return int64(150)
	}
	if (name == "catch") {
		return int64(150)
	}
	if (name == ".") {
		return int64(160)
	}
	if (name == "?.") {
		return int64(160)
	}
	if (name == "(") {
		return int64(160)
	}
	if (name == "[") {
		return int64(160)
	}
	return int64(0)
}

func is_infix_op(kind TokenKind) bool {
	bp := infix_binding_power(kind)
	_ = bp
	return (bp.Left > int64(0))
}

func is_prefix_op(kind TokenKind) bool {
	return (prefix_binding_power(kind) > int64(0))
}

func is_postfix_op(kind TokenKind) bool {
	return (postfix_binding_power(kind) > int64(0))
}

func is_comparison_op(kind TokenKind) bool {
	name := token_name(kind)
	_ = name
	return ((((((name == "==") || (name == "!=")) || (name == "<")) || (name == ">")) || (name == "<=")) || (name == ">="))
}

func is_assignment_op(kind TokenKind) bool {
	name := token_name(kind)
	_ = name
	return ((name == "=") || (name == ":="))
}

