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

type Command interface {
	isCommand()
}

type CommandCmdCheck struct{}
func (CommandCmdCheck) isCommand() {}

type CommandCmdBuild struct{}
func (CommandCmdBuild) isCommand() {}

type CommandCmdRun struct{}
func (CommandCmdRun) isCommand() {}

type CommandCmdTest struct{}
func (CommandCmdTest) isCommand() {}

type CommandCmdFix struct{}
func (CommandCmdFix) isCommand() {}

type CommandCmdExplain struct{}
func (CommandCmdExplain) isCommand() {}

type CommandCmdBench struct{}
func (CommandCmdBench) isCommand() {}

type CommandCmdHelp struct{}
func (CommandCmdHelp) isCommand() {}

type CommandCmdVersion struct{}
func (CommandCmdVersion) isCommand() {}

type CommandCmdUnknown struct{}
func (CommandCmdUnknown) isCommand() {}

type Options struct {
	Command Command
	Files []string
	Format string
	Target string
	Runtime_path string
	Parallel bool
	Output_path string
	Help_requested bool
}

func default_options() Options {
	return Options{Command: CommandCmdHelp{}, Files: []string{""}, Format: "text", Target: "native", Runtime_path: "runtime/runtime.c", Parallel: false, Output_path: "", Help_requested: false}
}

func _opts_set_cmd(o Options, c Command) Options {
	return Options{Command: c, Files: o.Files, Format: o.Format, Target: o.Target, Runtime_path: o.Runtime_path, Parallel: o.Parallel, Output_path: o.Output_path, Help_requested: o.Help_requested}
}

func _opts_set_format(o Options, v string) Options {
	return Options{Command: o.Command, Files: o.Files, Format: v, Target: o.Target, Runtime_path: o.Runtime_path, Parallel: o.Parallel, Output_path: o.Output_path, Help_requested: o.Help_requested}
}

func _opts_set_target(o Options, v string) Options {
	return Options{Command: o.Command, Files: o.Files, Format: o.Format, Target: v, Runtime_path: o.Runtime_path, Parallel: o.Parallel, Output_path: o.Output_path, Help_requested: o.Help_requested}
}

func _opts_set_runtime(o Options, v string) Options {
	return Options{Command: o.Command, Files: o.Files, Format: o.Format, Target: o.Target, Runtime_path: v, Parallel: o.Parallel, Output_path: o.Output_path, Help_requested: o.Help_requested}
}

func _opts_set_parallel(o Options) Options {
	return Options{Command: o.Command, Files: o.Files, Format: o.Format, Target: o.Target, Runtime_path: o.Runtime_path, Parallel: true, Output_path: o.Output_path, Help_requested: o.Help_requested}
}

func _opts_set_output(o Options, v string) Options {
	return Options{Command: o.Command, Files: o.Files, Format: o.Format, Target: o.Target, Runtime_path: o.Runtime_path, Parallel: o.Parallel, Output_path: v, Help_requested: o.Help_requested}
}

func _opts_set_files(o Options, f []string) Options {
	return Options{Command: o.Command, Files: f, Format: o.Format, Target: o.Target, Runtime_path: o.Runtime_path, Parallel: o.Parallel, Output_path: o.Output_path, Help_requested: o.Help_requested}
}

func _opts_set_help(o Options) Options {
	return Options{Command: o.Command, Files: o.Files, Format: o.Format, Target: o.Target, Runtime_path: o.Runtime_path, Parallel: o.Parallel, Output_path: o.Output_path, Help_requested: true}
}

func parse_command(arg string) Command {
	if (arg == "check") {
		return CommandCmdCheck{}
	}
	if (arg == "build") {
		return CommandCmdBuild{}
	}
	if (arg == "run") {
		return CommandCmdRun{}
	}
	if (arg == "test") {
		return CommandCmdTest{}
	}
	if (arg == "fix") {
		return CommandCmdFix{}
	}
	if (arg == "explain") {
		return CommandCmdExplain{}
	}
	if (arg == "bench") {
		return CommandCmdBench{}
	}
	if (arg == "help") {
		return CommandCmdHelp{}
	}
	if (arg == "--help") {
		return CommandCmdHelp{}
	}
	if (arg == "-h") {
		return CommandCmdHelp{}
	}
	if (arg == "version") {
		return CommandCmdVersion{}
	}
	if (arg == "--version") {
		return CommandCmdVersion{}
	}
	if (arg == "-v") {
		return CommandCmdVersion{}
	}
	return CommandCmdUnknown{}
}

func command_to_str(cmd Command) string {
	return func() interface{} {
		switch _tmp1 := (cmd).(type) {
		case CommandCmdCheck:
			return "check"
		case CommandCmdBuild:
			return "build"
		case CommandCmdRun:
			return "run"
		case CommandCmdTest:
			return "test"
		case CommandCmdFix:
			return "fix"
		case CommandCmdExplain:
			return "explain"
		case CommandCmdBench:
			return "bench"
		case CommandCmdHelp:
			return "help"
		case CommandCmdVersion:
			return "version"
		case CommandCmdUnknown:
			return "unknown"
		default:
			_ = _tmp1
		}
		return nil
	}().(string)
}

func _is_flag(arg string) bool {
	return ((int64(len(arg)) > int64(0)) && (string(arg[int64(0)]) == "-"))
}

func _parse_flag_value(arg string) string {
	i := int64(0)
	_ = i
	for (i < int64(len(arg))) {
		if (string(arg[i]) == "=") {
			return arg[(i + int64(1)):int64(len(arg))]
		}
		i = (i + int64(1))
	}
	return ""
}

func _validate_format(v string) bool {
	return ((v == "text") || (v == "json"))
}

func parse_args(args []string) Options {
	opts := default_options()
	_ = opts
	cmd_set := false
	_ = cmd_set
	files := []string{""}
	_ = files
	i := int64(2)
	_ = i
	for (i < int64(len(args))) {
		arg := args[i]
		_ = arg
		if _is_flag(arg) {
			if ((arg == "--help") || (arg == "-h")) {
				opts = _opts_set_help(opts)
			} else if strings.HasPrefix(arg, "--format=") {
				fv := _parse_flag_value(arg)
				_ = fv
				if _validate_format(fv) {
					opts = _opts_set_format(opts, fv)
				} else {
					fmt.Println((("error: invalid format '" + fv) + "' (expected 'text' or 'json')"))
				}
			} else if strings.HasPrefix(arg, "--target=") {
				opts = _opts_set_target(opts, _parse_flag_value(arg))
			} else if strings.HasPrefix(arg, "--runtime=") {
				opts = _opts_set_runtime(opts, _parse_flag_value(arg))
			} else if strings.HasPrefix(arg, "--output=") {
				opts = _opts_set_output(opts, _parse_flag_value(arg))
			} else if (arg == "-o") {
				if ((i + int64(1)) < int64(len(args))) {
					i = (i + int64(1))
					opts = _opts_set_output(opts, args[i])
				} else {
					fmt.Println("error: -o requires a path argument")
				}
			} else if (arg == "--parallel") {
				opts = _opts_set_parallel(opts)
			} else if ((arg == "--version") || (arg == "-v")) {
				opts = _opts_set_cmd(opts, CommandCmdVersion{})
			} else {
				fmt.Println((("error: unknown flag '" + arg) + "'"))
			}
		} else if (cmd_set == false) {
			opts = _opts_set_cmd(opts, parse_command(arg))
			cmd_set = true
		} else {
			files = append(files, arg)
		}
		i = (i + int64(1))
	}
	return _opts_set_files(opts, files)
}

func print_version() int64 {
	fmt.Println((("aria " + COMPILER_VERSION) + " (self-hosting)"))
	return int64(0)
}

func _print_command_header(name string, desc string) {
	fmt.Println(((("aria " + name) + " -- ") + desc))
	fmt.Println("")
	fmt.Println((("Usage: aria " + name) + " [flags] <files...>"))
	fmt.Println("")
}

func print_help() int64 {
	fmt.Println((("aria " + COMPILER_VERSION) + " -- the Aria self-hosting compiler"))
	fmt.Println("")
	fmt.Println("Usage: aria <command> [flags] [files...]")
	fmt.Println("")
	fmt.Println("Commands:")
	fmt.Println("  check     Type-check source files")
	fmt.Println("  build     Compile source files to executable")
	fmt.Println("  run       Compile and run source files")
	fmt.Println("  test      Compile and run test blocks")
	fmt.Println("  bench     Compile and run bench blocks")
	fmt.Println("  fix       Apply auto-fix suggestions")
	fmt.Println("  explain   Explain an error code (e.g., aria explain E0100)")
	fmt.Println("")
	fmt.Println("Flags:")
	fmt.Println("  -o PATH              Output path for compiled binary")
	fmt.Println("  --output=PATH        Output path (same as -o)")
	fmt.Println("  --format=text|json   Output format (default: text)")
	fmt.Println("  --target=TARGET      Target platform")
	fmt.Println("  --runtime=PATH       Path to runtime.c (default: runtime/runtime.c)")
	fmt.Println("  --parallel           Run tests in parallel")
	fmt.Println("  --version, -v        Print version")
	fmt.Println("  --help, -h           Print this help")
	fmt.Println("")
	fmt.Println("Run 'aria <command> --help' for command-specific help.")
	return int64(0)
}

func _print_build_help() int64 {
	_print_command_header("build", "Compile source files to an executable")
	fmt.Println("Flags:")
	fmt.Println("  -o PATH              Output binary path (default: derived from input)")
	fmt.Println("  --runtime=PATH       Path to runtime.c")
	fmt.Println("  --target=TARGET      Target platform (default: native)")
	fmt.Println("  --format=text|json   Diagnostic output format")
	return int64(0)
}

func _print_run_help() int64 {
	_print_command_header("run", "Compile and run source files")
	fmt.Println("Flags:")
	fmt.Println("  --runtime=PATH       Path to runtime.c")
	fmt.Println("  --format=text|json   Diagnostic output format")
	return int64(0)
}

func _print_check_help() int64 {
	_print_command_header("check", "Type-check source files without compiling")
	fmt.Println("Flags:")
	fmt.Println("  --format=text|json   Diagnostic output format")
	return int64(0)
}

func _print_test_help() int64 {
	_print_command_header("test", "Compile and run test blocks")
	fmt.Println("Flags:")
	fmt.Println("  -o PATH              Output test binary path")
	fmt.Println("  --runtime=PATH       Path to runtime.c")
	fmt.Println("  --parallel           Run tests in parallel (experimental)")
	fmt.Println("  --format=text|json   Diagnostic output format")
	return int64(0)
}

func _print_bench_help() int64 {
	_print_command_header("bench", "Compile and run bench blocks")
	fmt.Println("Flags:")
	fmt.Println("  --runtime=PATH       Path to runtime.c")
	fmt.Println("  --format=text|json   Diagnostic output format")
	return int64(0)
}

func _ends_with_aria(s string) bool {
	if (int64(len(s)) < int64(6)) {
		return false
	}
	return (s[(int64(len(s)) - int64(5)):int64(len(s))] == ".aria")
}

func _is_stdlib_module(name string) bool {
	return (((name == "json") || (name == "http")) || (name == "db"))
}

func _find_aria_home() string {
	home := _ariaGetenv("ARIA_HOME")
	_ = home
	if (int64(len(home)) > int64(0)) {
		return home
	}
	return "/usr/local/lib/aria"
}

func _stdlib_path(aria_home string, module_name string) string {
	return (((aria_home + "/lib/") + module_name) + ".aria")
}

func _scan_file_for_stdlib_imports(file_path string) []string {
	modules := []string{""}
	_ = modules
	source := _ariaReadFile(file_path)
	_ = source
	i := int64(0)
	_ = i
	for (i < int64(len(source))) {
		if ((i == int64(0)) || (((i > int64(0)) && (string(source[(i - int64(1))]) == "\n")))) {
			j := i
			_ = j
			for ((j < int64(len(source))) && (((string(source[j]) == " ") || (string(source[j]) == "\t")))) {
				j = (j + int64(1))
			}
			if (((j + int64(4)) <= int64(len(source))) && (source[j:(j + int64(4))] == "use ")) {
				k := (j + int64(4))
				_ = k
				for ((k < int64(len(source))) && (((string(source[k]) == " ") || (string(source[k]) == "\t")))) {
					k = (k + int64(1))
				}
				name_start := k
				_ = name_start
				for ((((k < int64(len(source))) && (string(source[k]) != " ")) && (string(source[k]) != "\n")) && (string(source[k]) != "\t")) {
					k = (k + int64(1))
				}
				if (k > name_start) {
					name := source[name_start:k]
					_ = name
					if _is_stdlib_module(name) {
						modules = append(modules, name)
					}
				}
			}
		}
		i = (i + int64(1))
	}
	return modules
}

func _resolve_stdlib_files(files []string) []string {
	aria_home := _find_aria_home()
	_ = aria_home
	extra := []string{""}
	_ = extra
	seen := []string{""}
	_ = seen
	fi := int64(1)
	_ = fi
	for (fi < int64(len(files))) {
		modules := _scan_file_for_stdlib_imports(files[fi])
		_ = modules
		mi := int64(1)
		_ = mi
		for (mi < int64(len(modules))) {
			mod_name := modules[mi]
			_ = mod_name
			already := false
			_ = already
			si := int64(1)
			_ = si
			for (si < int64(len(seen))) {
				if (seen[si] == mod_name) {
					already = true
				}
				si = (si + int64(1))
			}
			if (already == false) {
				local_path := (("lib/" + mod_name) + ".aria")
				_ = local_path
				local_content := _ariaReadFile(local_path)
				_ = local_content
				if (int64(len(local_content)) > int64(0)) {
					extra = append(extra, local_path)
					seen = append(seen, mod_name)
				} else if (aria_home != "/usr/local/lib/aria") {
					lib_path := _stdlib_path(aria_home, mod_name)
					_ = lib_path
					content := _ariaReadFile(lib_path)
					_ = content
					if (int64(len(content)) > int64(0)) {
						extra = append(extra, lib_path)
						seen = append(seen, mod_name)
					}
				}
			}
			mi = (mi + int64(1))
		}
		fi = (fi + int64(1))
	}
	return extra
}

func _merge_file_lists(a []string, b []string) []string {
	result := a
	_ = result
	i := int64(1)
	_ = i
	for (i < int64(len(b))) {
		result = append(result, b[i])
		i = (i + int64(1))
	}
	return result
}

func _resolve_runtime_path(explicit string) string {
	content := _ariaReadFile(explicit)
	_ = content
	if (int64(len(content)) > int64(0)) {
		return explicit
	}
	aria_home := _find_aria_home()
	_ = aria_home
	return (aria_home + "/runtime/runtime.c")
}

func _resolve_and_validate_runtime(runtime_path string) string {
	rt := _resolve_runtime_path(runtime_path)
	_ = rt
	content := _ariaReadFile(rt)
	_ = content
	if (int64(len(content)) > int64(0)) {
		return rt
	}
	fmt.Println((("error: runtime not found at '" + rt) + "'"))
	fmt.Println("hint: set ARIA_HOME or use --runtime=PATH")
	return ""
}

func _determine_output_path(raw_files []string, file string, explicit string) string {
	if (int64(len(explicit)) > int64(0)) {
		return explicit
	}
	output := "a.out"
	_ = output
	if ((int64(len(file)) > int64(5)) && (file[(int64(len(file)) - int64(5)):int64(len(file))] == ".aria")) {
		output = file[int64(0):(int64(len(file)) - int64(5))]
	}
	if (int64(len(raw_files)) > int64(1)) {
		raw_input := raw_files[int64(1)]
		_ = raw_input
		if (_ariaIsDir(raw_input) == int64(1)) {
			out := raw_input
			_ = out
			if ((int64(len(out)) > int64(1)) && (string(out[(int64(len(out)) - int64(1))]) == "/")) {
				out = out[int64(0):(int64(len(out)) - int64(1))]
			}
			last_slash := (int64(0) - int64(1))
			_ = last_slash
			si := int64(0)
			_ = si
			for (si < int64(len(out))) {
				if (string(out[si]) == "/") {
					last_slash = si
				}
				si = (si + int64(1))
			}
			if (last_slash >= int64(0)) {
				output = ((out + "/") + out[(last_slash + int64(1)):int64(len(out))])
			} else {
				output = ((out + "/") + out)
			}
		}
	}
	return output
}

func _expand_dir(path string) []string {
	result := []string{""}
	_ = result
	entries := _ariaListDir(path)
	_ = entries
	i := int64(1)
	_ = i
	for (i < int64(len(entries))) {
		ename := entries[i]
		_ = ename
		full := path
		_ = full
		if ((int64(len(path)) > int64(0)) && (string(path[(int64(len(path)) - int64(1))]) != "/")) {
			full = (path + "/")
		}
		full = (full + ename)
		if (_ariaIsDir(full) == int64(1)) {
			sub := _expand_dir(full)
			_ = sub
			si := int64(1)
			_ = si
			for (si < int64(len(sub))) {
				result = append(result, sub[si])
				si = (si + int64(1))
			}
		} else if _ends_with_aria(ename) {
			result = append(result, full)
		}
		i = (i + int64(1))
	}
	return result
}

func _module_priority(path string) int64 {
	base := int64(30)
	_ = base
	if strings.Contains(path, "/lexer/") {
		base = int64(10)
	}
	if strings.Contains(path, "/parser/") {
		base = int64(20)
	}
	if strings.Contains(path, "/diagnostic/") {
		base = int64(30)
	}
	if strings.Contains(path, "/resolver/") {
		base = int64(40)
	}
	if strings.Contains(path, "/checker/") {
		base = int64(50)
	}
	if strings.Contains(path, "/codegen/") {
		base = int64(60)
	}
	if strings.Contains(path, "/main.aria") {
		base = int64(70)
	}
	if strings.Contains(path, "/token.aria") {
		return base
	}
	if strings.Contains(path, "/ast.aria") {
		return base
	}
	if strings.Contains(path, "/types.aria") {
		return base
	}
	if strings.Contains(path, "/scope.aria") {
		return base
	}
	if strings.Contains(path, "/ir.aria") {
		return base
	}
	if strings.Contains(path, "/precedence.aria") {
		return base
	}
	if strings.Contains(path, "/diagnostic.aria") {
		return base
	}
	if strings.Contains(path, "/tests.aria") {
		return (base + int64(5))
	}
	if strings.Contains(path, "_tests.aria") {
		return (base + int64(5))
	}
	return (base + int64(2))
}

func _sort_by_module_priority(files []string) []string {
	sorted := []string{""}
	_ = sorted
	i := int64(1)
	_ = i
	for (i < int64(len(files))) {
		sorted = append(sorted, files[i])
		i = (i + int64(1))
	}
	si := int64(2)
	_ = si
	for (si < int64(len(sorted))) {
		key := sorted[si]
		_ = key
		key_pri := _module_priority(key)
		_ = key_pri
		j := (si - int64(1))
		_ = j
		for ((j >= int64(1)) && (_module_priority(sorted[j]) > key_pri)) {
			sorted = _arr_set(sorted, (j + int64(1)), sorted[j])
			j = (j - int64(1))
		}
		sorted = _arr_set(sorted, (j + int64(1)), key)
		si = (si + int64(1))
	}
	return sorted
}

func _arr_set(arr []string, idx int64, val string) []string {
	result := []string{""}
	_ = result
	i := int64(1)
	_ = i
	for (i < int64(len(arr))) {
		if (i == idx) {
			result = append(result, val)
		} else {
			result = append(result, arr[i])
		}
		i = (i + int64(1))
	}
	return result
}

func expand_files(files []string) []string {
	result := []string{""}
	_ = result
	i := int64(1)
	_ = i
	for (i < int64(len(files))) {
		file := files[i]
		_ = file
		if (_ariaIsDir(file) == int64(1)) {
			expanded := _expand_dir(file)
			_ = expanded
			ei := int64(1)
			_ = ei
			for (ei < int64(len(expanded))) {
				result = append(result, expanded[ei])
				ei = (ei + int64(1))
			}
		} else {
			result = append(result, file)
		}
		i = (i + int64(1))
	}
	return _sort_by_module_priority(result)
}

func check_source(source string, file string) DiagnosticBag {
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
	return cr.Diagnostics
}

func run_check(raw_files []string, format string) int64 {
	user_files := expand_files(raw_files)
	_ = user_files
	if (int64(len(user_files)) <= int64(1)) {
		fmt.Println("error: no input files")
		return int64(1)
	}
	stdlib_files := _resolve_stdlib_files(user_files)
	_ = stdlib_files
	files := _merge_file_lists(user_files, stdlib_files)
	_ = files
	diags := check_sources(files)
	_ = diags
	if bag_has_errors(diags) {
		_print_diagnostics(diags, format)
		return int64(1)
	}
	if (bag_warning_count(diags) > int64(0)) {
		_print_diagnostics(diags, format)
	}
	wcount := bag_warning_count(diags)
	_ = wcount
	fmt.Println((("check: ok (" + fmt.Sprintf("%v", wcount)) + " warnings)"))
	return int64(0)
}

func _print_diagnostics(diags DiagnosticBag, format string) {
	if (format == "json") {
		fmt.Println(render_json_bag(diags))
	} else {
		di := int64(1)
		_ = di
		for (di < int64(len(diags.Diagnostics))) {
			d := diags.Diagnostics[di]
			_ = d
			fmt.Println(((((((format_span(d.Span) + ": ") + severity_to_str(d.Severity)) + "[") + d.Code) + "]: ") + d.Message))
			di = (di + int64(1))
		}
		fmt.Println(format_summary(diags))
	}
	_unused := int64(0)
	_ = _unused
}

func run_build(raw_files []string, format string, target string, runtime_path string, output_path string) int64 {
	user_files := expand_files(raw_files)
	_ = user_files
	if (int64(len(user_files)) <= int64(1)) {
		fmt.Println("error: no input files")
		return int64(1)
	}
	stdlib_files := _resolve_stdlib_files(user_files)
	_ = stdlib_files
	files := _merge_file_lists(user_files, stdlib_files)
	_ = files
	rt := _resolve_and_validate_runtime(runtime_path)
	_ = rt
	if (rt == "") {
		return int64(1)
	}
	file := files[int64(1)]
	_ = file
	output := _determine_output_path(raw_files, file, output_path)
	_ = output
	diags := generate_from_sources(files, output, rt)
	_ = diags
	if bag_has_errors(diags) {
		_print_diagnostics(diags, format)
		return int64(1)
	}
	if (bag_warning_count(diags) > int64(0)) {
		_print_diagnostics(diags, format)
	}
	fmt.Println(("built: " + output))
	return int64(0)
}

func run_run(raw_files []string, format string, target string, runtime_path string, output_path string) int64 {
	user_files := expand_files(raw_files)
	_ = user_files
	if (int64(len(user_files)) <= int64(1)) {
		fmt.Println("error: no input files")
		return int64(1)
	}
	stdlib_files := _resolve_stdlib_files(user_files)
	_ = stdlib_files
	files := _merge_file_lists(user_files, stdlib_files)
	_ = files
	rt := _resolve_and_validate_runtime(runtime_path)
	_ = rt
	if (rt == "") {
		return int64(1)
	}
	file := files[int64(1)]
	_ = file
	output := _determine_output_path(raw_files, file, output_path)
	_ = output
	diags := generate_from_sources(files, output, rt)
	_ = diags
	if bag_has_errors(diags) {
		_print_diagnostics(diags, format)
		return int64(1)
	}
	if (bag_warning_count(diags) > int64(0)) {
		_print_diagnostics(diags, format)
	}
	exit_code := _ariaExec(output)
	_ = exit_code
	return exit_code
}

func run_test(raw_files []string, format string, runtime_path string, parallel bool) int64 {
	user_files := expand_files(raw_files)
	_ = user_files
	if (int64(len(user_files)) <= int64(1)) {
		fmt.Println("error: no input files")
		return int64(1)
	}
	stdlib_files := _resolve_stdlib_files(user_files)
	_ = stdlib_files
	files := _merge_file_lists(user_files, stdlib_files)
	_ = files
	rt := _resolve_and_validate_runtime(runtime_path)
	_ = rt
	if (rt == "") {
		return int64(1)
	}
	file := files[int64(1)]
	_ = file
	output := file
	_ = output
	if ((int64(len(file)) > int64(5)) && (file[(int64(len(file)) - int64(5)):int64(len(file))] == ".aria")) {
		output = file[int64(0):(int64(len(file)) - int64(5))]
	}
	output = (output + "_test_runner")
	diags := generate_tests_from_sources(files, output, rt, parallel)
	_ = diags
	if bag_has_errors(diags) {
		_print_diagnostics(diags, format)
		return int64(1)
	}
	return _ariaExec(output)
}

func run_explain(files []string) int64 {
	if (int64(len(files)) <= int64(1)) {
		fmt.Println("Usage: aria explain <error-code>")
		fmt.Println("")
		fmt.Println("Example: aria explain E0100")
		return int64(1)
	}
	code := files[int64(1)]
	_ = code
	desc := describe_code(code)
	_ = desc
	if (desc == "unknown error code") {
		fmt.Println(("Unknown error code: " + code))
		return int64(1)
	}
	fmt.Println(((code + ": ") + desc))
	fmt.Println("")
	first_char := string(code[int64(0)])
	_ = first_char
	if (first_char == "E") {
		num := code[int64(1):int64(len(code))]
		_ = num
		fmt.Println(("Category: " + _error_category(code)))
		fmt.Println("")
		fmt.Println(_error_explanation(code))
	}
	if (first_char == "W") {
		fmt.Println("Category: Warning")
		fmt.Println("")
		fmt.Println("This is a warning -- your code will still compile, but the")
		fmt.Println("indicated issue may cause problems or indicate a mistake.")
	}
	return int64(0)
}

func _error_category(code string) string {
	if strings.HasPrefix(code, "E000") {
		return "Syntax Error"
	}
	if strings.HasPrefix(code, "E001") {
		return "Type Error"
	}
	if strings.HasPrefix(code, "E002") {
		return "Trait Error"
	}
	if strings.HasPrefix(code, "E003") {
		return "Effect Error"
	}
	if strings.HasPrefix(code, "E004") {
		return "Pattern Error"
	}
	if strings.HasPrefix(code, "E005") {
		return "Mutability Error"
	}
	if strings.HasPrefix(code, "E007") {
		return "Module/Import Error"
	}
	return "Compiler Error"
}

func _error_explanation(code string) string {
	if (code == "E0100") {
		return "The types on the left and right side of an expression\ndon't match. Check that the types are compatible."
	}
	if (code == "E0106") {
		return "The function's return type doesn't match what the body\nactually returns. Check the return type annotation."
	}
	if (code == "E0200") {
		return "A type used in a generic context doesn't implement the\nrequired trait. Add a derives clause or impl block."
	}
	if (code == "E0400") {
		return "A match expression doesn't cover all possible cases.\nAdd the missing variant or a wildcard (_) arm."
	}
	if (code == "E0500") {
		return "Attempted to reassign a variable that wasn't declared\nwith 'mut'. Add 'mut' to the binding if reassignment is needed."
	}
	if (code == "E0701") {
		return "A name was used that isn't defined in the current scope.\nCheck spelling, imports, and visibility."
	}
	if (code == "E0704") {
		return "Attempted to access a symbol that is not marked 'pub'\nfrom a different file. Add 'pub' to the declaration or\nmove the access to the same file."
	}
	return "See the Aria language specification for details on this error."
}

func run_fix(raw_files []string, format string) int64 {
	user_files := expand_files(raw_files)
	_ = user_files
	if (int64(len(user_files)) <= int64(1)) {
		fmt.Println("error: no input files")
		return int64(1)
	}
	stdlib_files := _resolve_stdlib_files(user_files)
	_ = stdlib_files
	files := _merge_file_lists(user_files, stdlib_files)
	_ = files
	diags := check_sources(files)
	_ = diags
	fix_count := int64(0)
	_ = fix_count
	di := int64(1)
	_ = di
	for (di < int64(len(diags.Diagnostics))) {
		d := diags.Diagnostics[di]
		_ = d
		if (int64(len(d.Suggestions)) > int64(1)) {
			si := int64(1)
			_ = si
			for (si < int64(len(d.Suggestions))) {
				s := d.Suggestions[si]
				_ = s
				fmt.Println(((((("fix: " + d.Code) + " at ") + format_span(d.Span)) + ": ") + s.Message))
				fix_count = (fix_count + int64(1))
				si = (si + int64(1))
			}
		}
		di = (di + int64(1))
	}
	if (fix_count == int64(0)) {
		fmt.Println("No auto-fixable issues found.")
	} else {
		fmt.Println((fmt.Sprintf("%v", fix_count) + " fix(es) suggested. Use --apply to apply them."))
	}
	return int64(0)
}

func run_bench(raw_files []string, format string, runtime_path string) int64 {
	user_files := expand_files(raw_files)
	_ = user_files
	if (int64(len(user_files)) <= int64(1)) {
		fmt.Println("error: no input files")
		return int64(1)
	}
	stdlib_files := _resolve_stdlib_files(user_files)
	_ = stdlib_files
	files := _merge_file_lists(user_files, stdlib_files)
	_ = files
	rt := _resolve_and_validate_runtime(runtime_path)
	_ = rt
	if (rt == "") {
		return int64(1)
	}
	file := files[int64(1)]
	_ = file
	output := file
	_ = output
	if ((int64(len(file)) > int64(5)) && (file[(int64(len(file)) - int64(5)):int64(len(file))] == ".aria")) {
		output = file[int64(0):(int64(len(file)) - int64(5))]
	}
	output = (output + "_bench_runner")
	diags := generate_bench_from_sources(files, output, rt)
	_ = diags
	if bag_has_errors(diags) {
		_print_diagnostics(diags, format)
		return int64(1)
	}
	return _ariaExec(output)
}

var COMPILER_VERSION = "0.5.0"

func main() {
	debug.SetMemoryLimit(8 * 1024 * 1024 * 1024)
	debug.SetGCPercent(-1)
	args := _ariaArgs()
	_ = args
	if (int64(len(args)) < int64(3)) {
		print_help()
	} else {
		opts := parse_args(args)
		_ = opts
		cmd := command_to_str(opts.Command)
		_ = cmd
		if opts.Help_requested {
			if (cmd == "build") {
				_print_build_help()
			} else if (cmd == "run") {
				_print_run_help()
			} else if (cmd == "check") {
				_print_check_help()
			} else if (cmd == "test") {
				_print_test_help()
			} else if (cmd == "bench") {
				_print_bench_help()
			} else {
				print_help()
			}
		} else if (cmd == "build") {
			run_build(opts.Files, opts.Format, opts.Target, opts.Runtime_path, opts.Output_path)
		} else if (cmd == "run") {
			run_run(opts.Files, opts.Format, opts.Target, opts.Runtime_path, opts.Output_path)
		} else if (cmd == "check") {
			run_check(opts.Files, opts.Format)
		} else if (cmd == "test") {
			run_test(opts.Files, opts.Format, opts.Runtime_path, opts.Parallel)
		} else if (cmd == "bench") {
			run_bench(opts.Files, opts.Format, opts.Runtime_path)
		} else if (cmd == "fix") {
			run_fix(opts.Files, opts.Format)
		} else if (cmd == "explain") {
			run_explain(opts.Files)
		} else if (cmd == "version") {
			print_version()
		} else {
			print_help()
		}
	}
}


// --- Aria runtime helpers ---

func _ariaRange(start, end int64, inclusive bool) []int64 {
	var result []int64
	if inclusive {
		for i := start; i <= end; i++ {
			result = append(result, i)
		}
	} else {
		for i := start; i < end; i++ {
			result = append(result, i)
		}
	}
	return result
}

func _ariaReverse[T any](s []T) []T {
	result := make([]T, len(s))
	for i, v := range s {
		result[len(s)-1-i] = v
	}
	return result
}

func _ariaParseInt(s string) int64 {
	v, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		panic(fmt.Sprintf("parseInt failed: %s", s))
	}
	return v
}

func _ariaParseFloat(s string) float64 {
	v, err := strconv.ParseFloat(s, 64)
	if err != nil {
		panic(fmt.Sprintf("parseFloat failed: %s", s))
	}
	return v
}

func _ariaReadFile(path string) string {
	data, err := os.ReadFile(path)
	if err != nil {
		panic(fmt.Sprintf("readFile failed: %v", err))
	}
	return string(data)
}

func _ariaWriteFile(path string, content string) {
	err := os.WriteFile(path, []byte(content), 0644)
	if err != nil {
		panic(fmt.Sprintf("writeFile failed: %v", err))
	}
}

func _ariaAppendFile(path string, content string) {
	f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(fmt.Sprintf("appendFile failed: %v", err))
	}
	defer f.Close()
	_, err = f.WriteString(content)
	if err != nil {
		panic(fmt.Sprintf("appendFile write failed: %v", err))
	}
}

func _ariaFileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func _ariaWriteBinaryFile(path string, bytes interface{}) {
	data := make([]byte, 0)
	switch v := bytes.(type) {
	case []int64:
		// Skip sentinel at index 0
		for i, b := range v {
			if i == 0 { continue }
			data = append(data, byte(b&0xFF))
		}
	case []interface{}:
		// Skip sentinel at index 0
		for i, b := range v {
			if i == 0 { continue }
			switch bv := b.(type) {
			case int64:
				data = append(data, byte(bv&0xFF))
			case int:
				data = append(data, byte(bv&0xFF))
			default:
				data = append(data, 0)
			}
		}
	}
	err := os.WriteFile(path, data, 0755)
	if err != nil {
		panic(fmt.Sprintf("writeBinaryFile failed: %v", err))
	}
}

func _ariaArgs() []string {
	// Returns os.Args as a sentinel array (element 0 = "")
	result := []string{""}
	result = append(result, os.Args...)
	return result
}

func _ariaMapKeys[K comparable, V any](m map[K]V) []K {
	keys := make([]K, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

func _ariaMapValues[K comparable, V any](m map[K]V) []V {
	vals := make([]V, 0, len(m))
	for _, v := range m {
		vals = append(vals, v)
	}
	return vals
}

type _ariaMapEntry[K comparable, V any] struct {
	Key   K
	Value V
}

func _ariaMapEntries[K comparable, V any](m map[K]V) []_ariaMapEntry[K, V] {
	entries := make([]_ariaMapEntry[K, V], 0, len(m))
	for k, v := range m {
		entries = append(entries, _ariaMapEntry[K, V]{k, v})
	}
	return entries
}

func _ariaMapContains[K comparable, V any](m map[K]V, key K) bool {
	_, ok := m[key]
	return ok
}

func _ariaExec(command string) int64 {
	cmd := exec.Command("sh", "-c", command)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			return int64(exitErr.ExitCode())
		}
		return 1
	}
	return 0
}

// Environment
func _ariaGetenv(name string) string { return os.Getenv(name) }

func _ariaListDir(path string) []string {
	entries, err := os.ReadDir(path)
	if err != nil {
		return []string{""}
	}
	result := []string{""}
	for _, e := range entries {
		result = append(result, e.Name())
	}
	return result
}

func _ariaIsDir(path string) int64 {
	info, err := os.Stat(path)
	if err != nil {
		return 0
	}
	if info.IsDir() {
		return 1
	}
	return 0
}

// TCP networking stubs (bootstrap uses LLVM backend for actual networking)
func _ariaTcpSocket() int64 { panic("TCP not supported in bootstrap") }
func _ariaTcpBind(fd int64, addr string, port int64) int64 { panic("TCP not supported in bootstrap") }
func _ariaTcpListen(fd int64, backlog int64) int64 { panic("TCP not supported in bootstrap") }
func _ariaTcpAccept(fd int64) int64 { panic("TCP not supported in bootstrap") }
func _ariaTcpRead(fd int64, maxLen int64) string { panic("TCP not supported in bootstrap") }
func _ariaTcpWrite(fd int64, data string) int64 { panic("TCP not supported in bootstrap") }
func _ariaTcpClose(fd int64) { panic("TCP not supported in bootstrap") }
func _ariaTcpPeerAddr(fd int64) string { panic("TCP not supported in bootstrap") }
func _ariaTcpSetTimeout(fd int64, kind int64, ms int64) int64 { panic("TCP not supported in bootstrap") }

// PostgreSQL stubs (bootstrap uses LLVM backend for actual PG access)
func _ariaPgConnect(connstr string) int64 { panic("PostgreSQL not supported in bootstrap") }
func _ariaPgClose(conn int64) { panic("PostgreSQL not supported in bootstrap") }
func _ariaPgStatus(conn int64) int64 { panic("PostgreSQL not supported in bootstrap") }
func _ariaPgError(conn int64) string { panic("PostgreSQL not supported in bootstrap") }
func _ariaPgExec(conn int64, query string) int64 { panic("PostgreSQL not supported in bootstrap") }
func _ariaPgExecParams(conn int64, query string, params []string) int64 { panic("PostgreSQL not supported in bootstrap") }
func _ariaPgResultStatus(result int64) int64 { panic("PostgreSQL not supported in bootstrap") }
func _ariaPgResultError(result int64) string { panic("PostgreSQL not supported in bootstrap") }
func _ariaPgNrows(result int64) int64 { panic("PostgreSQL not supported in bootstrap") }
func _ariaPgNcols(result int64) int64 { panic("PostgreSQL not supported in bootstrap") }
func _ariaPgFieldName(result int64, col int64) string { panic("PostgreSQL not supported in bootstrap") }
func _ariaPgGetValue(result int64, row int64, col int64) string { panic("PostgreSQL not supported in bootstrap") }
func _ariaPgIsNull(result int64, row int64, col int64) int64 { panic("PostgreSQL not supported in bootstrap") }
func _ariaPgClear(result int64) { panic("PostgreSQL not supported in bootstrap") }

// Concurrency stubs
func _ariaSpawn(closure interface{}) int64 { panic("concurrency not supported in bootstrap") }
func _ariaTaskAwait(handle int64) int64 { panic("concurrency not supported in bootstrap") }
func _ariaChanNew(capacity int64) int64 { panic("concurrency not supported in bootstrap") }
func _ariaChanSend(ch int64, value int64) int64 { panic("concurrency not supported in bootstrap") }
func _ariaChanRecv(ch int64) int64 { panic("concurrency not supported in bootstrap") }
func _ariaChanClose(ch int64) { panic("concurrency not supported in bootstrap") }
func _ariaMutexNew() int64 { panic("concurrency not supported in bootstrap") }
func _ariaMutexLock(handle int64) { panic("concurrency not supported in bootstrap") }
func _ariaMutexUnlock(handle int64) { panic("concurrency not supported in bootstrap") }
