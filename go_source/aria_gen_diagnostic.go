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

type Severity interface {
	isSeverity()
}

type SeverityError struct{}
func (SeverityError) isSeverity() {}

type SeverityWarning struct{}
func (SeverityWarning) isSeverity() {}

type SeverityInfo struct{}
func (SeverityInfo) isSeverity() {}

type SeverityHint struct{}
func (SeverityHint) isSeverity() {}

type LabelStyle interface {
	isLabelStyle()
}

type LabelStylePrimary struct{}
func (LabelStylePrimary) isLabelStyle() {}

type LabelStyleSecondary struct{}
func (LabelStyleSecondary) isLabelStyle() {}

type Applicability interface {
	isApplicability()
}

type ApplicabilityMachineApplicable struct{}
func (ApplicabilityMachineApplicable) isApplicability() {}

type ApplicabilityMaybeIncorrect struct{}
func (ApplicabilityMaybeIncorrect) isApplicability() {}

type ApplicabilityHasPlaceholders struct{}
func (ApplicabilityHasPlaceholders) isApplicability() {}

type Span struct {
	File string
	Line int64
	Col int64
	Offset int64
	Length int64
}

type Label struct {
	Span Span
	Message string
	Style LabelStyle
}

type Suggestion struct {
	Message string
	Replacement string
	Span Span
	Applicability Applicability
}

type Diagnostic struct {
	Code string
	Severity Severity
	Message string
	Span Span
	Labels []Label
	Notes []string
	Suggestions []Suggestion
}

type DiagnosticBag struct {
	Diagnostics []Diagnostic
	Error_count int64
	Warning_count int64
}

func severity_to_str(s Severity) string {
	return func() interface{} {
		switch _tmp1 := (s).(type) {
		case SeverityError:
			return "error"
		case SeverityWarning:
			return "warning"
		case SeverityInfo:
			return "info"
		case SeverityHint:
			return "hint"
		default:
			_ = _tmp1
		}
		return nil
	}().(string)
}

func label_style_to_str(s LabelStyle) string {
	return func() interface{} {
		switch _tmp2 := (s).(type) {
		case LabelStylePrimary:
			return "primary"
		case LabelStyleSecondary:
			return "secondary"
		default:
			_ = _tmp2
		}
		return nil
	}().(string)
}

func applicability_to_str(a Applicability) string {
	return func() interface{} {
		switch _tmp3 := (a).(type) {
		case ApplicabilityMachineApplicable:
			return "MachineApplicable"
		case ApplicabilityMaybeIncorrect:
			return "MaybeIncorrect"
		case ApplicabilityHasPlaceholders:
			return "HasPlaceholders"
		default:
			_ = _tmp3
		}
		return nil
	}().(string)
}

func new_span(file string, line int64, col int64, offset int64, length int64) Span {
	return Span{File: file, Line: line, Col: col, Offset: offset, Length: length}
}

func span_end_offset(s Span) int64 {
	return (s.Offset + s.Length)
}

func _span_file(s Span) string {
	return s.File
}

func i2s(n int64) string {
	if (n < int64(0)) {
		return ("-" + i2s((int64(0) - n)))
	}
	if (n == int64(0)) {
		return "0"
	}
	if (n < int64(10)) {
		if (n == int64(1)) {
			return "1"
		}
		if (n == int64(2)) {
			return "2"
		}
		if (n == int64(3)) {
			return "3"
		}
		if (n == int64(4)) {
			return "4"
		}
		if (n == int64(5)) {
			return "5"
		}
		if (n == int64(6)) {
			return "6"
		}
		if (n == int64(7)) {
			return "7"
		}
		if (n == int64(8)) {
			return "8"
		}
		return "9"
	}
	return (i2s((n / int64(10))) + i2s((n % int64(10))))
}

func format_span(s Span) string {
	return ((((s.File + ":") + i2s(s.Line)) + ":") + i2s(s.Col))
}

func new_label(span Span, message string, style LabelStyle) Label {
	return Label{Span: span, Message: message, Style: style}
}

func primary_label(span Span, message string) Label {
	return Label{Span: span, Message: message, Style: LabelStylePrimary{}}
}

func secondary_label(span Span, message string) Label {
	return Label{Span: span, Message: message, Style: LabelStyleSecondary{}}
}

func new_suggestion(message string, replacement string, span Span, applicability Applicability) Suggestion {
	return Suggestion{Message: message, Replacement: replacement, Span: span, Applicability: applicability}
}

func _sentinel_span() Span {
	return new_span("", int64(0), int64(0), int64(0), int64(0))
}

func new_diagnostic(code string, severity Severity, message string, span Span) Diagnostic {
	return Diagnostic{Code: code, Severity: severity, Message: message, Span: span, Labels: []Label{primary_label(span, message)}, Notes: []string{""}, Suggestions: []Suggestion{new_suggestion("", "", span, ApplicabilityMachineApplicable{})}}
}

func new_error(code string, message string, span Span) Diagnostic {
	return new_diagnostic(code, SeverityError{}, message, span)
}

func new_warning(code string, message string, span Span) Diagnostic {
	return new_diagnostic(code, SeverityWarning{}, message, span)
}

func diag_with_note(diag Diagnostic, note string) Diagnostic {
	return Diagnostic{Code: diag.Code, Severity: diag.Severity, Message: diag.Message, Span: diag.Span, Labels: diag.Labels, Notes: append(diag.Notes, note), Suggestions: diag.Suggestions}
}

func bag_add_error_with_note(bag DiagnosticBag, code string, message string, span Span, note string) DiagnosticBag {
	return bag_add_diagnostic(bag, diag_with_note(new_error(code, message, span), note))
}

func new_bag() DiagnosticBag {
	return DiagnosticBag{Diagnostics: []Diagnostic{new_error("", "", _sentinel_span())}, Error_count: int64(0), Warning_count: int64(0)}
}

func bag_add_diagnostic(bag DiagnosticBag, diag Diagnostic) DiagnosticBag {
	i := int64(1)
	_ = i
	for (i < int64(len(bag.Diagnostics))) {
		existing := bag.Diagnostics[i]
		_ = existing
		if (((existing.Code == diag.Code) && (existing.Span.Line == diag.Span.Line)) && (existing.Span.File == diag.Span.File)) {
			return bag
		}
		i = (i + int64(1))
	}
	new_diagnostics := append(bag.Diagnostics, diag)
	_ = new_diagnostics
	new_errors := bag.Error_count
	_ = new_errors
	new_warnings := bag.Warning_count
	_ = new_warnings
	sev := severity_to_str(diag.Severity)
	_ = sev
	if (sev == "error") {
		new_errors = (new_errors + int64(1))
	}
	if (sev == "warning") {
		new_warnings = (new_warnings + int64(1))
	}
	return DiagnosticBag{Diagnostics: new_diagnostics, Error_count: new_errors, Warning_count: new_warnings}
}

func bag_add_error(bag DiagnosticBag, code string, message string, span Span) DiagnosticBag {
	return bag_add_diagnostic(bag, new_error(code, message, span))
}

func bag_add_warning(bag DiagnosticBag, code string, message string, span Span) DiagnosticBag {
	return bag_add_diagnostic(bag, new_warning(code, message, span))
}

func bag_has_errors(bag DiagnosticBag) bool {
	return (bag.Error_count > int64(0))
}

func bag_error_count(bag DiagnosticBag) int64 {
	return bag.Error_count
}

func bag_warning_count(bag DiagnosticBag) int64 {
	return bag.Warning_count
}

func bag_merge(a DiagnosticBag, b DiagnosticBag) DiagnosticBag {
	result := a
	_ = result
	i := int64(1)
	_ = i
	for (i < int64(len(b.Diagnostics))) {
		result = bag_add_diagnostic(result, b.Diagnostics[i])
		i = (i + int64(1))
	}
	return result
}

func format_summary(bag DiagnosticBag) string {
	return (((i2s(bag.Error_count) + " error(s), ") + i2s(bag.Warning_count)) + " warning(s)")
}

func describe_code(code string) string {
	if (code == E0001) {
		return "unexpected token"
	}
	if (code == E0002) {
		return "unterminated string literal"
	}
	if (code == E0003) {
		return "unterminated string interpolation"
	}
	if (code == E0004) {
		return "invalid escape sequence"
	}
	if (code == E0005) {
		return "expected expression"
	}
	if (code == E0006) {
		return "expected type"
	}
	if (code == E0007) {
		return "expected pattern"
	}
	if (code == E0008) {
		return "expected identifier"
	}
	if (code == E0009) {
		return "expected declaration"
	}
	if (code == E0010) {
		return "invalid number literal"
	}
	if (code == E0011) {
		return "expected closing delimiter"
	}
	if (code == E0012) {
		return "expected opening delimiter"
	}
	if (code == E0013) {
		return "unexpected end of file"
	}
	if (code == E0014) {
		return "expected newline"
	}
	if (code == E0015) {
		return "expected mod declaration"
	}
	if (code == E0016) {
		return "expected function body"
	}
	if (code == E0017) {
		return "expected match arm"
	}
	if (code == E0018) {
		return "expected block"
	}
	if (code == E0019) {
		return "invalid operator"
	}
	if (code == E0020) {
		return "non-associative operator chaining"
	}
	if (code == E0100) {
		return "type mismatch"
	}
	if (code == E0101) {
		return "cannot infer type"
	}
	if (code == E0102) {
		return "incompatible types in binary operation"
	}
	if (code == E0103) {
		return "invalid unary operation"
	}
	if (code == E0104) {
		return "not callable"
	}
	if (code == E0105) {
		return "wrong number of arguments"
	}
	if (code == E0106) {
		return "return type mismatch"
	}
	if (code == E0107) {
		return "field not found"
	}
	if (code == E0108) {
		return "not a struct type"
	}
	if (code == E0109) {
		return "missing required field"
	}
	if (code == E0110) {
		return "duplicate field"
	}
	if (code == E0200) {
		return "trait not satisfied"
	}
	if (code == E0201) {
		return "missing trait method"
	}
	if (code == E0202) {
		return "orphan impl"
	}
	if (code == E0300) {
		return "missing effect [Io]"
	}
	if (code == E0301) {
		return "missing effect [Fs]"
	}
	if (code == E0302) {
		return "missing effect [Net]"
	}
	if (code == E0303) {
		return "missing effect [Ffi]"
	}
	if (code == E0304) {
		return "missing effect [Async]"
	}
	if (code == E0400) {
		return "non-exhaustive match"
	}
	if (code == E0500) {
		return "immutable binding"
	}
	if (code == E0401) {
		return "unreachable pattern"
	}
	if (code == E0700) {
		return "module not found"
	}
	if (code == E0701) {
		return "unresolved name"
	}
	if (code == E0702) {
		return "duplicate declaration"
	}
	if (code == E0703) {
		return "import cycle"
	}
	if (code == E0704) {
		return "private symbol"
	}
	if (code == W0001) {
		return "unused variable"
	}
	if (code == W0002) {
		return "unused function"
	}
	if (code == W0003) {
		return "unused import"
	}
	if (code == W0004) {
		return "shadowed variable"
	}
	if (code == "W0005") {
		return "use of deprecated function"
	}
	if (code == W0008) {
		return "once closure called multiple times"
	}
	if (code == W0009) {
		return "orphan impl across packages"
	}
	if (code == W0100) {
		return "unreachable code"
	}
	return "unknown error code"
}

func _repeat_char(ch string, count int64) string {
	result := ""
	_ = result
	i := int64(0)
	_ = i
	for (i < count) {
		result = (result + ch)
		i = (i + int64(1))
	}
	return result
}

func _get_source_line(source string, target_line int64) string {
	lines := strings.Split(source, "\n")
	_ = lines
	idx := (target_line - int64(1))
	_ = idx
	if ((idx >= int64(0)) && (idx < int64(len(lines)))) {
		return lines[idx]
	}
	return ""
}

func _line_number_width(line int64) int64 {
	if (line < int64(10)) {
		return int64(1)
	}
	if (line < int64(100)) {
		return int64(2)
	}
	if (line < int64(1000)) {
		return int64(3)
	}
	if (line < int64(10000)) {
		return int64(4)
	}
	return int64(5)
}

func _pad_left(s string, width int64) string {
	padding := (width - int64(len(s)))
	_ = padding
	if (padding <= int64(0)) {
		return s
	}
	return (_repeat_char(" ", padding) + s)
}

func render_text(diag Diagnostic, source string) string {
	sev := severity_to_str(diag.Severity)
	_ = sev
	nl := "\n"
	_ = nl
	output := (((((sev + "[") + diag.Code) + "]: ") + diag.Message) + nl)
	_ = output
	output = (((output + "  --> ") + format_span(diag.Span)) + nl)
	source_line := _get_source_line(source, diag.Span.Line)
	_ = source_line
	if (int64(len(source_line)) > int64(0)) {
		line_str := i2s(diag.Span.Line)
		_ = line_str
		width := _line_number_width(diag.Span.Line)
		_ = width
		gutter := _repeat_char(" ", width)
		_ = gutter
		output = (((output + gutter) + " |") + nl)
		output = ((((output + _pad_left(line_str, width)) + " | ") + source_line) + nl)
		col_offset := (diag.Span.Col - int64(1))
		_ = col_offset
		underline_len := diag.Span.Length
		_ = underline_len
		if (underline_len < int64(1)) {
			underline_len = int64(1)
		}
		carets := _repeat_char("^", underline_len)
		_ = carets
		spaces := _repeat_char(" ", col_offset)
		_ = spaces
		output = (((((output + gutter) + " | ") + spaces) + carets) + nl)
	}
	i := int64(0)
	_ = i
	for (i < int64(len(diag.Notes))) {
		note := diag.Notes[i]
		_ = note
		if (int64(len(note)) > int64(0)) {
			output = (((output + "  = note: ") + note) + nl)
		}
		i = (i + int64(1))
	}
	return output
}

func _escape_json_str(s string) string {
	result := ""
	_ = result
	i := int64(0)
	_ = i
	for (i < int64(len(s))) {
		ch := string(s[i])
		_ = ch
		if (ch == "\\") {
			result = (result + "\\\\")
		} else if (ch == _q()) {
			result = ((result + "\\") + _q())
		} else {
			result = (result + ch)
		}
		i = (i + int64(1))
	}
	return result
}

func _q() string {
	return "\""
}

func _json_str(key string, value string) string {
	q := _q()
	_ = q
	return ((((((q + key) + q) + ":") + q) + value) + q)
}

func _json_int(key string, value int64) string {
	q := _q()
	_ = q
	val_str := i2s(value)
	_ = val_str
	return ((((q + key) + q) + ":") + val_str)
}

func render_json_diagnostic(diag Diagnostic) string {
	sev := severity_to_str(diag.Severity)
	_ = sev
	msg := _escape_json_str(diag.Message)
	_ = msg
	file := _escape_json_str(diag.Span.File)
	_ = file
	end_off := span_end_offset(diag.Span)
	_ = end_off
	desc := _escape_json_str(describe_code(diag.Code))
	_ = desc
	q := _q()
	_ = q
	lb := "{"
	_ = lb
	rb := "}"
	_ = rb
	json := lb
	_ = json
	json = ((json + _json_str("code", diag.Code)) + ",")
	json = ((json + _json_str("severity", sev)) + ",")
	json = ((json + _json_str("message", msg)) + ",")
	json = ((json + _json_str("description", desc)) + ",")
	json = ((json + _json_str("file", file)) + ",")
	json = ((json + _json_int("line", diag.Span.Line)) + ",")
	json = ((json + _json_int("column", diag.Span.Col)) + ",")
	off_str := i2s(diag.Span.Offset)
	_ = off_str
	end_str := i2s(end_off)
	_ = end_str
	json = ((((((((json + q) + "span") + q) + ":[") + off_str) + ",") + end_str) + "],")
	json = ((((json + q) + "labels") + q) + ":[")
	li := int64(0)
	_ = li
	first_lbl := true
	_ = first_lbl
	for (li < int64(len(diag.Labels))) {
		lbl := diag.Labels[li]
		_ = lbl
		if (int64(len(lbl.Message)) > int64(0)) {
			if first_lbl {
				first_lbl = false
			} else {
				json = (json + ",")
			}
			json = (json + lb)
			json = ((json + _json_str("message", _escape_json_str(lbl.Message))) + ",")
			json = ((json + _json_str("style", label_style_to_str(lbl.Style))) + ",")
			json = ((json + _json_str("file", _escape_json_str(lbl.Span.File))) + ",")
			json = ((json + _json_int("line", lbl.Span.Line)) + ",")
			json = (json + _json_int("column", lbl.Span.Col))
			json = (json + rb)
		}
		li = (li + int64(1))
	}
	json = (json + "],")
	json = ((((json + q) + "suggestions") + q) + ":[")
	si := int64(0)
	_ = si
	first_sug := true
	_ = first_sug
	for (si < int64(len(diag.Suggestions))) {
		sug := diag.Suggestions[si]
		_ = sug
		if (int64(len(sug.Message)) > int64(0)) {
			if first_sug {
				first_sug = false
			} else {
				json = (json + ",")
			}
			json = (json + lb)
			json = ((json + _json_str("message", _escape_json_str(sug.Message))) + ",")
			json = ((json + _json_str("replacement", _escape_json_str(sug.Replacement))) + ",")
			json = (json + _json_str("applicability", applicability_to_str(sug.Applicability)))
			json = (json + rb)
		}
		si = (si + int64(1))
	}
	json = (json + "]")
	json = (json + rb)
	return json
}

func render_json_bag(bag DiagnosticBag) string {
	q := _q()
	_ = q
	lb := "{"
	_ = lb
	rb := "}"
	_ = rb
	json := ((((lb + q) + "diagnostics") + q) + ":[")
	_ = json
	first := true
	_ = first
	i := int64(1)
	_ = i
	for (i < int64(len(bag.Diagnostics))) {
		if first {
			first = false
		} else {
			json = (json + ",")
		}
		json = (json + render_json_diagnostic(bag.Diagnostics[i]))
		i = (i + int64(1))
	}
	json = (json + "],")
	json = (((((json + q) + "summary") + q) + ":") + lb)
	json = ((json + _json_int("errors", bag.Error_count)) + ",")
	json = (json + _json_int("warnings", bag.Warning_count))
	json = ((json + rb) + rb)
	return json
}

var E0001 = "E0001"

var E0002 = "E0002"

var E0003 = "E0003"

var E0004 = "E0004"

var E0005 = "E0005"

var E0006 = "E0006"

var E0007 = "E0007"

var E0008 = "E0008"

var E0009 = "E0009"

var E0010 = "E0010"

var E0011 = "E0011"

var E0012 = "E0012"

var E0013 = "E0013"

var E0014 = "E0014"

var E0015 = "E0015"

var E0016 = "E0016"

var E0017 = "E0017"

var E0018 = "E0018"

var E0019 = "E0019"

var E0020 = "E0020"

var E0100 = "E0100"

var E0101 = "E0101"

var E0102 = "E0102"

var E0103 = "E0103"

var E0104 = "E0104"

var E0105 = "E0105"

var E0106 = "E0106"

var E0107 = "E0107"

var E0108 = "E0108"

var E0109 = "E0109"

var E0110 = "E0110"

var E0111 = "E0111"

var E0112 = "E0112"

var E0113 = "E0113"

var E0114 = "E0114"

var E0200 = "E0200"

var E0201 = "E0201"

var E0202 = "E0202"

var E0203 = "E0203"

var E0204 = "E0204"

var E0205 = "E0205"

var E0300 = "E0300"

var E0301 = "E0301"

var E0302 = "E0302"

var E0303 = "E0303"

var E0304 = "E0304"

var E0400 = "E0400"

var E0401 = "E0401"

var E0402 = "E0402"

var E0403 = "E0403"

var E0500 = "E0500"

var E0700 = "E0700"

var E0701 = "E0701"

var E0702 = "E0702"

var E0703 = "E0703"

var E0704 = "E0704"

var W0001 = "W0001"

var W0002 = "W0002"

var W0003 = "W0003"

var W0004 = "W0004"

var W0005 = "W0005"

var W0008 = "W0008"

var W0009 = "W0009"

var W0100 = "W0100"

