#!/usr/bin/env python3
"""
Go -> C transpiler for the Aria bootstrap compiler.
Translates the mechanically-generated Go code into a single C file.

The Go code has very regular patterns produced by the Aria bootstrap compiler:
- Sum types: interface + empty struct variants -> C integer enums
- Structs: Go structs -> C structs
- Functions: package-level only, no methods
- Control flow: if/else, for loops, return
- Type switches: IIFE pattern with switch on type
"""

import re
import sys
import os
from pathlib import Path
from collections import OrderedDict


# ---------------------------------------------------------------------------
# Type mapping
# ---------------------------------------------------------------------------

_fn_ret_types = {}  # Global: function name -> C return type


def go_type_to_c(t, sum_types, struct_names):
    """Convert a Go type string to C type string."""
    t = t.strip()
    if t == 'int64':
        return 'int64_t'
    if t == 'string':
        return 'AriaStr'
    if t == 'bool':
        return 'int64_t'
    if t == 'float64':
        return 'double'
    if t == 'interface{}':
        return 'int64_t'
    if t == 'byte':
        return 'uint8_t'
    # Slice types: []Type
    m = re.match(r'^\[\](.+)$', t)
    if m:
        inner = m.group(1).strip()
        c_inner = go_type_to_c(inner, sum_types, struct_names)
        return f'AriaSlice_{safe_c_name(inner)}'
    # Sum type interfaces -> int64_t (enum tag)
    if t in sum_types:
        return 'int64_t'
    # Known struct -> struct name
    if t in struct_names:
        return t
    # Fallback
    return t


def safe_c_name(go_type):
    """Make a Go type name safe for use in C identifiers."""
    t = go_type.strip()
    if t == 'int64':
        return 'int64_t'
    if t == 'string':
        return 'AriaStr'
    if t == 'bool':
        return 'int64_t'
    if t == 'float64':
        return 'double'
    if t == 'interface{}':
        return 'int64_t'
    if t == 'byte':
        return 'uint8_t'
    m = re.match(r'^\[\](.+)$', t)
    if m:
        return f'AriaSlice_{safe_c_name(m.group(1).strip())}'
    return t


# ---------------------------------------------------------------------------
# Phase 1: Parse Go source to extract types
# ---------------------------------------------------------------------------

def parse_sum_types(all_source):
    """Extract sum type interfaces and their variants."""
    sum_types = OrderedDict()  # interface_name -> [variant_struct_name, ...]

    for name, src in all_source.items():
        # Find interface declarations
        for m in re.finditer(r'^type (\w+) interface \{\s*\n\s*is\w+\(\)\s*\n\}', src, re.MULTILINE):
            iface_name = m.group(1)
            sum_types[iface_name] = []

    # Now find variant structs: type FooBar struct{} \n func (FooBar) isFoo() {}
    for name, src in all_source.items():
        for m in re.finditer(r'^type (\w+) struct\{\}\s*\nfunc \(\1\) (is\w+)\(\) \{\}', src, re.MULTILINE):
            variant_name = m.group(1)
            method_name = m.group(2)
            # Find which interface this belongs to
            for iface_name in sum_types:
                expected_method = f'is{iface_name}'
                if method_name == expected_method:
                    sum_types[iface_name].append(variant_name)
                    break

    return sum_types


def parse_structs(all_source, sum_types):
    """Extract struct definitions (excluding empty sum type variant structs)."""
    structs = OrderedDict()
    variant_names = set()
    for variants in sum_types.values():
        variant_names.update(variants)

    for name, src in all_source.items():
        for m in re.finditer(r'^type (\w+) struct \{(.*?)\n\}', src, re.MULTILINE | re.DOTALL):
            struct_name = m.group(1)
            if struct_name in variant_names:
                continue  # Skip empty variant structs
            fields_text = m.group(2)
            fields = parse_go_fields(fields_text)
            if fields:  # Only non-empty structs
                structs[struct_name] = fields

    return structs


def parse_go_fields(text):
    """Parse Go struct field declarations."""
    fields = []
    for line in text.strip().split('\n'):
        line = line.strip()
        if not line or line.startswith('//'):
            continue
        parts = line.split()
        if len(parts) >= 2:
            field_name = parts[0]
            field_type = ' '.join(parts[1:])
            fields.append((field_name, field_type))
    return fields


def parse_global_vars(all_source):
    """Extract global variable declarations like `var COMPILER_VERSION = "0.5.0"`."""
    globs = []
    for name, src in all_source.items():
        for m in re.finditer(r'^var (\w+) = (.+)$', src, re.MULTILINE):
            var_name = m.group(1)
            var_value = m.group(2).strip()
            if var_name == '_':
                continue
            globs.append((var_name, var_value))
    return globs


# ---------------------------------------------------------------------------
# Phase 2: Parse functions
# ---------------------------------------------------------------------------

def find_matching_brace(src, start):
    """Find the matching closing brace for a function body starting at `start`
    (which should point to the opening brace)."""
    depth = 0
    i = start
    in_string = False
    in_rune = False
    escape = False
    while i < len(src):
        c = src[i]
        if escape:
            escape = False
            i += 1
            continue
        if c == '\\' and (in_string or in_rune):
            escape = True
            i += 1
            continue
        if c == '"' and not in_rune:
            in_string = not in_string
        elif c == '\'' and not in_string:
            in_rune = not in_rune
        elif not in_string and not in_rune:
            if c == '{':
                depth += 1
            elif c == '}':
                depth -= 1
                if depth == 0:
                    return i
        i += 1
    return len(src) - 1


def parse_functions(all_source, sum_types, struct_names):
    """Extract all functions with their full bodies."""
    functions = []
    # Skip these runtime helper functions that use Go-specific features
    skip_fns = {
        'main', '_ariaRange', '_ariaReverse', '_ariaParseInt', '_ariaParseFloat',
        '_ariaReadFile', '_ariaWriteFile', '_ariaAppendFile', '_ariaFileExists',
        '_ariaWriteBinaryFile', '_ariaArgs', '_ariaMapKeys', '_ariaMapValues',
        '_ariaMapEntries', '_ariaMapContains', '_ariaExec', '_ariaGetenv',
        '_ariaListDir', '_ariaIsDir',
        '_ariaTcpSocket', '_ariaTcpBind', '_ariaTcpListen', '_ariaTcpAccept',
        '_ariaTcpRead', '_ariaTcpWrite', '_ariaTcpClose', '_ariaTcpPeerAddr',
        '_ariaTcpSetTimeout',
        '_ariaPgConnect', '_ariaPgClose', '_ariaPgStatus', '_ariaPgError',
        '_ariaPgExec', '_ariaPgExecParams', '_ariaPgResultStatus', '_ariaPgResultError',
        '_ariaPgNrows', '_ariaPgNcols', '_ariaPgFieldName', '_ariaPgGetValue',
        '_ariaPgIsNull', '_ariaPgClear',
        '_ariaSpawn', '_ariaTaskAwait', '_ariaChanNew', '_ariaChanSend',
        '_ariaChanRecv', '_ariaChanClose', '_ariaMutexNew', '_ariaMutexLock',
        '_ariaMutexUnlock',
    }

    # Also skip variant methods like `func (TokenKindTkIdent) isTokenKind() {}`
    for name, src in all_source.items():
        # Find function declarations (not methods)
        for m in re.finditer(r'^func (\w+)\(', src, re.MULTILINE):
            fn_name = m.group(1)
            if fn_name in skip_fns:
                continue

            # Get the full function signature
            sig_start = m.start()
            # Find the opening brace
            brace_pos = src.index('{', m.end())
            sig_text = src[sig_start:brace_pos].strip()

            # Parse signature
            # func name(params) rettype
            sig_m = re.match(r'func (\w+)\((.*?)\)\s*(.*)', sig_text, re.DOTALL)
            if not sig_m:
                continue
            params_text = sig_m.group(2).strip()
            ret_text = sig_m.group(3).strip()

            # Find body
            body_end = find_matching_brace(src, brace_pos)
            body = src[brace_pos + 1:body_end].strip()

            functions.append({
                'name': fn_name,
                'params': params_text,
                'ret': ret_text,
                'body': body,
                'source_file': name,
            })

    return functions


def parse_params(params_text, sum_types, struct_names):
    """Parse Go function parameters into list of (name, c_type) tuples."""
    if not params_text.strip():
        return []
    params = []
    # Split by comma, but careful about nested types
    parts = split_params(params_text)
    for part in parts:
        part = part.strip()
        if not part:
            continue
        tokens = part.split()
        if len(tokens) >= 2:
            pname = tokens[0]
            ptype = ' '.join(tokens[1:])
            c_type = go_type_to_c(ptype, sum_types, struct_names)
            params.append((pname, c_type))
        elif len(tokens) == 1:
            # unnamed parameter
            c_type = go_type_to_c(tokens[0], sum_types, struct_names)
            params.append(('_', c_type))
    return params


def split_params(text):
    """Split parameters by comma, respecting brackets and braces."""
    parts = []
    depth = 0
    current = []
    in_string = False
    escape = False
    for c in text:
        if escape:
            escape = False
            current.append(c)
            continue
        if c == '\\' and in_string:
            escape = True
            current.append(c)
            continue
        if c == '"':
            in_string = not in_string
            current.append(c)
            continue
        if in_string:
            current.append(c)
            continue
        if c in '([{':
            depth += 1
            current.append(c)
        elif c in ')]}':
            depth -= 1
            current.append(c)
        elif c == ',' and depth == 0:
            parts.append(''.join(current))
            current = []
        else:
            current.append(c)
    if current:
        parts.append(''.join(current))
    return parts


# ---------------------------------------------------------------------------
# Phase 3: Translate function bodies
# ---------------------------------------------------------------------------

def translate_body(body, ret_type, sum_types, struct_names, all_variants):
    """Translate a Go function body to C."""
    lines = body.split('\n')
    result = []
    i = 0
    while i < len(lines):
        line = lines[i]
        translated, skip = translate_line(line, lines, i, ret_type, sum_types, struct_names, all_variants)
        result.append(translated)
        i += 1 + skip
    return '\n'.join(result)


def translate_line(line, lines, idx, ret_type, sum_types, struct_names, all_variants):
    """Translate a single line. Returns (translated_line, lines_to_skip)."""
    stripped = line.strip()
    indent = line[:len(line) - len(line.lstrip())]

    # Empty lines
    if not stripped:
        return '', 0

    # Comments
    if stripped.startswith('//'):
        return f'{indent}{stripped}', 0

    # Skip `_ = x` (unused var suppression)
    if re.match(r'^_ = \w+$', stripped):
        return f'{indent}(void){stripped[4:]};', 0

    # Check for IIFE type-switch pattern (multi-line)
    # return func() interface{} { switch _tmp := (x).(type) { ... } }().(string)
    iife_match = re.match(r'^return func\(\) interface\{\} \{$', stripped)
    if iife_match:
        # Collect the entire IIFE block
        return translate_iife_typeswitch(lines, idx, indent, ret_type, sum_types, struct_names, all_variants)

    # Variable declaration: x := expr
    decl_match = re.match(r'^(\w+) := (.+)$', stripped)
    if decl_match:
        var_name = decl_match.group(1)
        expr = decl_match.group(2)
        c_expr, c_type = translate_expr_with_type(expr, sum_types, struct_names, all_variants)
        # If type inference gave us a questionable result, use __typeof__
        # for variable copies and field accesses where we might be wrong
        if c_type in ('int64_t', 'AriaStr') and expr not in ('true', 'false') and not expr.startswith('"') and not re.match(r'^-?\d+$', expr):
            # Check if RHS is a simple identifier, field access, or function call
            # that we might have inferred incorrectly
            if re.match(r'^[\w.]+$', expr):
                # Simple identifier or field access — use typeof
                return f'{indent}__typeof__({c_expr}) {var_name} = {c_expr};', 0
        return f'{indent}{c_type} {var_name} = {c_expr};', 0

    # Assignment: x = expr
    assign_match = re.match(r'^(\w[\w.]*(?:\[[\w\s+\-*/().]+\])?) = (.+)$', stripped)
    if assign_match and not stripped.startswith('if ') and not stripped.startswith('for ') and not stripped.startswith('return'):
        lhs = assign_match.group(1)
        rhs = assign_match.group(2)
        c_lhs = translate_expr(lhs, sum_types, struct_names, all_variants)
        c_rhs = translate_expr(rhs, sum_types, struct_names, all_variants)
        return f'{indent}{c_lhs} = {c_rhs};', 0

    # Return
    ret_match = re.match(r'^return (.+)$', stripped)
    if ret_match:
        expr = ret_match.group(1)
        c_expr = translate_expr(expr, sum_types, struct_names, all_variants)
        # Wrap bare string literals that didn't get wrapped
        if expr.startswith('"') and not c_expr.startswith('aria_str_lit'):
            c_expr = f'aria_str_lit({expr})'
        return f'{indent}return {c_expr};', 0

    if stripped == 'return':
        return f'{indent}return;', 0

    # If/else-if/else
    if_match = re.match(r'^(} else )?if (.+) \{$', stripped)
    if if_match:
        prefix = if_match.group(1) or ''
        cond = if_match.group(2)
        c_cond = translate_expr(cond, sum_types, struct_names, all_variants)
        c_prefix = '} else ' if prefix else ''
        return f'{indent}{c_prefix}if ({c_cond}) {{', 0

    if stripped == '} else {':
        return f'{indent}}} else {{', 0

    # For loop
    for_match = re.match(r'^for (.+) \{$', stripped)
    if for_match:
        cond = for_match.group(1)
        c_cond = translate_expr(cond, sum_types, struct_names, all_variants)
        return f'{indent}while ({c_cond}) {{', 0

    # Closing brace
    if stripped == '}':
        return f'{indent}}}', 0

    # fmt.Println with single string arg
    println_match = re.match(r'^fmt\.Println\((.+)\)$', stripped)
    if println_match:
        arg = println_match.group(1)
        c_arg = translate_expr(arg, sum_types, struct_names, all_variants)
        return f'{indent}aria_println({c_arg});', 0

    # fmt.Fprintf to stderr
    fprintf_match = re.match(r'^fmt\.Fprintf\(os\.Stderr, (.+)\)$', stripped)
    if fprintf_match:
        args = fprintf_match.group(1)
        c_args = translate_expr(args, sum_types, struct_names, all_variants)
        return f'{indent}aria_eprint({c_args});', 0

    # Bare function call
    call_match = re.match(r'^(\w+)\((.*)$', stripped)
    if call_match:
        fn_name = call_match.group(1)
        rest = call_match.group(2)
        # Reconstruct full call
        full = stripped
        c_full = translate_expr(full, sum_types, struct_names, all_variants)
        return f'{indent}{c_full};', 0

    # Fallback: try to translate as expression statement
    c_expr = translate_expr(stripped, sum_types, struct_names, all_variants)
    return f'{indent}{c_expr};', 0


def translate_iife_typeswitch(lines, start_idx, indent, ret_type, sum_types, struct_names, all_variants):
    """Translate the IIFE type-switch pattern used for match expressions.
    Pattern:
        return func() interface{} {
            switch _tmpN := (expr).(type) {
            case TypeVariant:
                return value
            ...
            default:
                _ = _tmpN
            }
            return nil
        }().(string)
    """
    # Collect all lines until the closing }().(type)
    block_lines = []
    i = start_idx
    depth = 0
    while i < len(lines):
        l = lines[i].strip()
        block_lines.append(lines[i])
        depth += l.count('{') - l.count('}')
        # Check for the closing pattern: }().(string) or }()
        if re.match(r'^\t*\}.*\(\)\.\(\w+\)$', lines[i].strip()) or (depth <= 0 and i > start_idx):
            break
        i += 1

    lines_consumed = i - start_idx

    # Parse the switch variable and cases
    # Find: switch _tmpN := (expr).(type) {
    switch_match = None
    cases = []
    default_body = None
    for bl in block_lines:
        bl_stripped = bl.strip()
        sm = re.match(r'switch _tmp\d+ := \((\w+)\)\.\(type\) \{', bl_stripped)
        if sm:
            switch_match = sm
            continue
        cm = re.match(r'case (\w+):$', bl_stripped)
        if cm:
            cases.append({'variant': cm.group(1), 'lines': []})
            continue
        if bl_stripped == 'default:':
            default_body = []
            continue
        if cases and bl_stripped.startswith('return '):
            val = bl_stripped[7:]
            cases[-1]['return_val'] = val
            continue
        if default_body is not None and bl_stripped not in ('}', ''):
            default_body.append(bl_stripped)

    if not switch_match or not cases:
        # Can't parse; return as comment
        raw = '\n'.join(block_lines)
        return f'{indent}/* TODO: unparsed IIFE type-switch */\n{indent}/* {raw} */', lines_consumed

    switch_var = switch_match.group(1)

    # Generate C switch statement
    result_lines = [f'{indent}switch ({switch_var}) {{']
    for case_item in cases:
        variant = case_item['variant']
        raw_val = case_item.get('return_val', '0')
        c_val = translate_expr(raw_val, sum_types, struct_names, all_variants)
        # Wrap bare string literals in aria_str_lit
        if raw_val.startswith('"') and not c_val.startswith('aria_str_lit'):
            c_val = f'aria_str_lit({raw_val})'
        result_lines.append(f'{indent}  case {variant}: return {c_val};')
    result_lines.append(f'{indent}  default: return aria_str_lit(""); /* unreachable */')
    result_lines.append(f'{indent}}}')

    return '\n'.join(result_lines), lines_consumed


def translate_expr_with_type(expr, sum_types, struct_names, all_variants):
    """Translate an expression and infer its C type."""
    expr = expr.strip()

    # Boolean literals
    if expr == 'true':
        return '1', 'int64_t'
    if expr == 'false':
        return '0', 'int64_t'

    # Integer literals: int64(N)
    m = re.match(r'^int64\(-?\d+\)$', expr)
    if m:
        c_expr = translate_expr(expr, sum_types, struct_names, all_variants)
        return c_expr, 'int64_t'

    # String literal
    if expr.startswith('"') and expr.endswith('"'):
        return f'aria_str_lit({expr})', 'AriaStr'

    # Sum type variant literal: VariantName{}
    m = re.match(r'^(\w+)\{\}$', expr)
    if m and m.group(1) in all_variants:
        return m.group(1), 'int64_t'

    # Struct literal: StructName{Field: val, ...}
    m = re.match(r'^(\w+)\{(.+)\}$', expr, re.DOTALL)
    if m and m.group(1) in struct_names:
        struct_name = m.group(1)
        fields_text = m.group(2)
        c_fields = translate_struct_literal_fields(fields_text, sum_types, struct_names, all_variants)
        return f'({struct_name}){{{c_fields}}}', struct_name

    # Slice literal: []Type{...}
    m = re.match(r'^\[\](\w+)\{(.*)\}$', expr, re.DOTALL)
    if m:
        elem_type = m.group(1)
        elems_text = m.group(2)
        slice_type = f'AriaSlice_{safe_c_name(elem_type)}'
        if not elems_text.strip():
            return f'({slice_type}){{NULL, 0, 0}}', slice_type
        elems = split_params(elems_text)
        c_elems = [translate_expr(e.strip(), sum_types, struct_names, all_variants) for e in elems]
        n = len(c_elems)
        elems_str = ', '.join(c_elems)
        return f'aria_slice_from_{safe_c_name(elem_type)}({n}, {elems_str})', slice_type

    # append(slice, elem)
    m = re.match(r'^append\((.+)\)$', expr)
    if m:
        args = split_params(m.group(1))
        if len(args) == 2:
            c_expr = translate_expr(expr, sum_types, struct_names, all_variants)
            return c_expr, '/* inferred */'

    # len(x) wrapped in int64
    m = re.match(r'^int64\(len\((.+)\)\)$', expr)
    if m:
        c_inner = translate_expr(m.group(1), sum_types, struct_names, all_variants)
        return f'({c_inner}).len', 'int64_t'

    # Default: translate expression and infer type
    c_expr = translate_expr(expr, sum_types, struct_names, all_variants)
    c_type = infer_type(expr, sum_types, struct_names, all_variants)
    return c_expr, c_type


def infer_type(expr, sum_types, struct_names, all_variants):
    """Infer the C type of a Go expression."""
    expr = expr.strip()
    if expr in ('true', 'false'):
        return 'int64_t'
    if re.match(r'^int64\(', expr):
        return 'int64_t'
    if re.match(r'^"', expr):
        return 'AriaStr'
    if re.match(r'^-?\d+$', expr):
        return 'int64_t'
    # Variant literal
    m = re.match(r'^(\w+)\{\}$', expr)
    if m and m.group(1) in all_variants:
        return 'int64_t'
    # Struct literal
    m = re.match(r'^(\w+)\{', expr)
    if m and m.group(1) in struct_names:
        return m.group(1)
    # Slice literal
    m = re.match(r'^\[\](\w+)\{', expr)
    if m:
        return f'AriaSlice_{safe_c_name(m.group(1))}'
    # Known string-returning expressions
    hint = get_expr_type_hint(expr, sum_types, struct_names, all_variants)
    if hint == 'string':
        return 'AriaStr'
    # Expressions with parenthesized inner — check recursively
    if expr.startswith('(') and expr.endswith(')'):
        return infer_type(expr[1:-1], sum_types, struct_names, all_variants)
    # Binary expressions — type from operands
    for op in ['||', '&&', '==', '!=', '<=', '>=', '<', '>']:
        if split_on_operator(expr, op):
            return 'int64_t'
    for op in ['+', '-', '*', '/', '%']:
        parts = split_on_operator(expr, op)
        if parts:
            lhs_hint = get_expr_type_hint(parts[0], sum_types, struct_names, all_variants)
            rhs_hint = get_expr_type_hint(parts[1], sum_types, struct_names, all_variants)
            if lhs_hint == 'string' or rhs_hint == 'string':
                return 'AriaStr'
            return 'int64_t'
    # Field access that returns string
    if '.' in expr:
        hint = get_expr_type_hint(expr, sum_types, struct_names, all_variants)
        if hint:
            return 'AriaStr' if hint == 'string' else hint
    # Function call - use fn_ret_types map first, then heuristics
    m = re.match(r'^(\w+)\(', expr)
    if m:
        fn = m.group(1)
        # Check the return type map first
        global _fn_ret_types
        if fn in _fn_ret_types:
            return _fn_ret_types[fn]
        # Fallback heuristics
        if fn in ('_is_digit', '_is_hex_digit', '_is_octal_digit', '_is_binary_digit',
                   '_is_alpha', '_is_alpha_num', '_is_whitespace', '_at_end',
                   '_in_delimiters', 'bag_has_errors', 'bag_error_count',
                   '_ariaFileExists', '_ariaIsDir', '_ariaParseInt'):
            return 'int64_t'
        if fn in ('_ariaParseFloat',):
            return 'double'
        hint = get_expr_type_hint(expr, sum_types, struct_names, all_variants)
        if hint == 'string':
            return 'AriaStr'
    return 'int64_t'  # default fallback


def get_expr_type_hint(expr, sum_types, struct_names, all_variants):
    """Determine if an expression is a string type (for == and + operators)."""
    expr = expr.strip()
    if expr.startswith('"'):
        return 'string'
    if expr.startswith('string(') or expr.startswith('aria_str_'):
        return 'string'
    # Check for known string-returning functions
    m = re.match(r'^(\w+)\(', expr)
    if m:
        fn = m.group(1)
        if fn in ('_peek', '_peek_next', 'token_name', 'expr_kind_name',
                   'symbol_kind_name', 'scope_kind_name', 'command_to_str',
                   'severity_name', 'label_style_name', 'applicability_name',
                   'ir_op_name', 'type_kind_name', 'decl_kind_name',
                   'i2s', '_parse_flag_value', 'new_token', 'new_lexer'):
            return 'string'
    # Known string fields
    for sf in ('Text', 'Name', 'File', 'Source', 'Format', 'Target',
               'Runtime_path', 'Output_path', 'Module_name', 'Message',
               'S1', 'S2', 'Code'):
        if expr.endswith('.' + sf):
            return 'string'
    # Variable names that are typically strings (only longer, unambiguous names)
    if re.match(r'^(text|name|file|path|source|arg|val|esc|msg|prefix|fv|line_str|formatted|result_str|buf)$', expr):
        return 'string'
    return None


def translate_expr(expr, sum_types, struct_names, all_variants):
    """Translate a Go expression to C."""
    expr = expr.strip()
    if not expr:
        return expr

    # nil
    if expr == 'nil':
        return '0'

    # Boolean
    if expr == 'true':
        return '1'
    if expr == 'false':
        return '0'

    # Type cast functions: int64(expr), string(expr), float64(expr)
    # Use balanced paren matching instead of greedy regex
    for cast_name, cast_fmt in [('int64', '((int64_t)({inner}))'),
                                  ('float64', '((double)({inner}))'),
                                  ('string', None),
                                  ('byte', None)]:
        if expr.startswith(cast_name + '('):
            args_str, after = find_matching_paren(expr[len(cast_name) + 1:])
            if args_str is not None and not after:  # Only match if the entire expr is the cast
                if cast_name == 'int64':
                    # int64(len(x)) -> (x).len
                    lm = re.match(r'^len\((.+)\)$', args_str)
                    if lm:
                        c_inner = translate_expr(lm.group(1), sum_types, struct_names, all_variants)
                        return f'({c_inner}).len'
                    c_inner = translate_expr(args_str, sum_types, struct_names, all_variants)
                    return f'((int64_t)({c_inner}))'
                elif cast_name == 'float64':
                    c_inner = translate_expr(args_str, sum_types, struct_names, all_variants)
                    return f'((double)({c_inner}))'
                elif cast_name == 'string':
                    # string(x[idx]) — single char access
                    sm = re.match(r'^(\w+)\[(.+)\]$', args_str)
                    if sm:
                        c_arr = translate_expr(sm.group(1), sum_types, struct_names, all_variants)
                        c_idx = translate_expr(sm.group(2), sum_types, struct_names, all_variants)
                        return f'aria_str_char_at({c_arr}, {c_idx})'
                    c_inner = translate_expr(args_str, sum_types, struct_names, all_variants)
                    return f'aria_str_from_bytes({c_inner})'

    # []byte(expr) — string to bytes
    m = re.match(r'^\[\]byte\((.+)\)$', expr)
    if m:
        inner = m.group(1)
        c_inner = translate_expr(inner, sum_types, struct_names, all_variants)
        return f'aria_str_to_bytes({c_inner})'

    # Variant literal: VariantName{}
    m = re.match(r'^(\w+)\{\}$', expr)
    if m and m.group(1) in all_variants:
        return m.group(1)

    # Struct literal: StructName{Field: val, ...}
    m = re.match(r'^(\w+)\{(.+)\}$', expr, re.DOTALL)
    if m and m.group(1) in struct_names:
        struct_name = m.group(1)
        fields_text = m.group(2)
        c_fields = translate_struct_literal_fields(fields_text, sum_types, struct_names, all_variants)
        return f'({struct_name}){{{c_fields}}}'

    # Slice literal: []Type{...}
    m = re.match(r'^\[\](\w+)\{(.*)\}$', expr, re.DOTALL)
    if m:
        elem_type = m.group(1)
        elems_text = m.group(2)
        slice_type = f'AriaSlice_{safe_c_name(elem_type)}'
        if not elems_text.strip():
            return f'({slice_type}){{NULL, 0, 0}}'
        elems = split_params(elems_text)
        c_elems = [translate_expr(e.strip(), sum_types, struct_names, all_variants) for e in elems]
        n = len(c_elems)
        elems_str = ', '.join(c_elems)
        return f'aria_slice_from_{safe_c_name(elem_type)}({n}, {elems_str})'

    # Slice indexing: x[idx] or x.Field[idx]
    m = re.match(r'^(\w+(?:\.\w+)*)\[(.+)\]$', expr)
    if m:
        obj = m.group(1)
        idx = m.group(2)
        # Check for slice expression: x[a:b]
        if ':' in idx:
            colon_pos = find_colon_in_slice(idx)
            if colon_pos is not None:
                start_expr = idx[:colon_pos]
                end_expr = idx[colon_pos + 1:]
                c_obj = translate_expr(obj, sum_types, struct_names, all_variants)
                c_start = translate_expr(start_expr, sum_types, struct_names, all_variants)
                c_end = translate_expr(end_expr, sum_types, struct_names, all_variants)
                return f'aria_str_slice({c_obj}, {c_start}, {c_end})'
        c_obj = translate_expr(obj, sum_types, struct_names, all_variants)
        c_idx = translate_expr(idx, sum_types, struct_names, all_variants)
        return f'{c_obj}.data[{c_idx}]'

    # append(slice, elem)
    m = re.match(r'^append\((.+)\)$', expr)
    if m:
        args = split_params(m.group(1))
        if len(args) == 2:
            slice_expr = args[0].strip()
            elem_expr = args[1].strip()
            c_slice = translate_expr(slice_expr, sum_types, struct_names, all_variants)
            c_elem = translate_expr(elem_expr, sum_types, struct_names, all_variants)
            # Try to infer slice element type from the slice variable name
            # We'll use a generic macro that works with the slice type
            return f'aria_append_generic({c_slice}, {c_elem})'

    # len(x) without int64 wrapper
    m = re.match(r'^len\((\w+(?:\.\w+)?)\)$', expr)
    if m:
        c_inner = translate_expr(m.group(1), sum_types, struct_names, all_variants)
        return f'({c_inner}).len'

    # fmt.Println(args)
    m = re.match(r'^fmt\.Println\((.+)\)$', expr)
    if m:
        c_arg = translate_expr(m.group(1), sum_types, struct_names, all_variants)
        return f'aria_println({c_arg})'

    # fmt.Fprintf(os.Stderr, args)
    m = re.match(r'^fmt\.Fprintf\(os\.Stderr, (.+)\)$', expr)
    if m:
        c_arg = translate_expr(m.group(1), sum_types, struct_names, all_variants)
        return f'aria_eprint({c_arg})'

    # strings.HasPrefix(a, b)
    m = re.match(r'^strings\.HasPrefix\((.+), (.+)\)$', expr)
    if m:
        c_a = translate_expr(m.group(1), sum_types, struct_names, all_variants)
        c_b = translate_expr(m.group(2), sum_types, struct_names, all_variants)
        return f'aria_str_has_prefix({c_a}, {c_b})'

    # strings.Contains(a, b)
    m = re.match(r'^strings\.Contains\((.+), (.+)\)$', expr)
    if m:
        c_a = translate_expr(m.group(1), sum_types, struct_names, all_variants)
        c_b = translate_expr(m.group(2), sum_types, struct_names, all_variants)
        return f'aria_str_contains({c_a}, {c_b})'

    # strings.ReplaceAll(s, old, new)
    m = re.match(r'^strings\.ReplaceAll\((.+)\)$', expr)
    if m:
        args = split_params(m.group(1))
        if len(args) == 3:
            c_args = [translate_expr(a.strip(), sum_types, struct_names, all_variants) for a in args]
            return f'aria_str_replace_all({c_args[0]}, {c_args[1]}, {c_args[2]})'

    # strings.TrimPrefix
    m = re.match(r'^strings\.TrimPrefix\((.+), (.+)\)$', expr)
    if m:
        c_a = translate_expr(m.group(1), sum_types, struct_names, all_variants)
        c_b = translate_expr(m.group(2), sum_types, struct_names, all_variants)
        return f'aria_str_trim_prefix({c_a}, {c_b})'

    # os.Exit(n)
    m = re.match(r'^os\.Exit\((.+)\)$', expr)
    if m:
        c_n = translate_expr(m.group(1), sum_types, struct_names, all_variants)
        return f'exit({c_n})'

    # Parenthesized expression
    if expr.startswith('(') and find_matching_close_paren(expr, 0) == len(expr) - 1:
        inner = expr[1:-1]
        c_inner = translate_expr(inner, sum_types, struct_names, all_variants)
        return f'({c_inner})'

    # Binary operators
    for op in ['||', '&&', '==', '!=', '<=', '>=', '<<', '>>', '<', '>', '+', '-', '*', '/', '%', '&', '|', '^']:
        parts = split_on_operator(expr, op)
        if parts:
            lhs, rhs = parts
            c_lhs = translate_expr(lhs, sum_types, struct_names, all_variants)
            c_rhs = translate_expr(rhs, sum_types, struct_names, all_variants)
            # Determine if either side is a string
            # Check original Go expression first
            lhs_type = get_expr_type_hint(lhs, sum_types, struct_names, all_variants)
            rhs_type = get_expr_type_hint(rhs, sum_types, struct_names, all_variants)
            # Also check if the translated expression looks like string ops
            if lhs_type != 'string' and _is_c_string_expr(c_lhs):
                lhs_type = 'string'
            if rhs_type != 'string' and _is_c_string_expr(c_rhs):
                rhs_type = 'string'
            # Also use infer_type for recursive cases
            if lhs_type is None:
                if infer_type(lhs, sum_types, struct_names, all_variants) == 'AriaStr':
                    lhs_type = 'string'
            if rhs_type is None:
                if infer_type(rhs, sum_types, struct_names, all_variants) == 'AriaStr':
                    rhs_type = 'string'
            is_str = (lhs_type == 'string' or rhs_type == 'string')
            if op == '==' and is_str:
                return f'aria_str_eq({c_lhs}, {c_rhs})'
            if op == '!=' and is_str:
                return f'(!aria_str_eq({c_lhs}, {c_rhs}))'
            # String concatenation with +
            if op == '+' and is_str:
                return f'aria_str_concat({c_lhs}, {c_rhs})'
            return f'{c_lhs} {op} {c_rhs}'

    # Unary NOT: !expr
    if expr.startswith('!'):
        c_inner = translate_expr(expr[1:], sum_types, struct_names, all_variants)
        return f'!{c_inner}'

    # Unary minus: -expr (when not part of a binary op)
    if expr.startswith('-') and len(expr) > 1 and not expr[1].isdigit():
        c_inner = translate_expr(expr[1:], sum_types, struct_names, all_variants)
        return f'-{c_inner}'

    # Field access: x.Field
    if '.' in expr and not expr.startswith('"') and not expr.startswith('fmt.') and not expr.startswith('os.') and not expr.startswith('strings.'):
        # Make sure it's not a float literal
        if not re.match(r'^\d+\.\d+$', expr):
            parts = expr.split('.', 1)
            if re.match(r'^\w+$', parts[0]) and re.match(r'^\w+', parts[1]):
                c_obj = translate_expr(parts[0], sum_types, struct_names, all_variants)
                rest = parts[1]
                c_rest = translate_field_chain(rest, sum_types, struct_names, all_variants)
                return f'{c_obj}.{c_rest}'

    # Function call: name(args)
    m = re.match(r'^(\w+)\((.*)$', expr)
    if m:
        fn_name = m.group(1)
        rest = m.group(2)
        args_str, after = find_matching_paren(rest)
        if args_str is not None:
            # Handle built-in type casts and functions
            if fn_name == 'int64':
                c_inner = translate_expr(args_str, sum_types, struct_names, all_variants)
                result = f'((int64_t)({c_inner}))'
            elif fn_name == 'len':
                c_inner = translate_expr(args_str, sum_types, struct_names, all_variants)
                result = f'({c_inner}).len'
            elif fn_name == 'append':
                args = split_params(args_str)
                if len(args) == 2:
                    c_s = translate_expr(args[0].strip(), sum_types, struct_names, all_variants)
                    c_e = translate_expr(args[1].strip(), sum_types, struct_names, all_variants)
                    result = f'aria_append_generic({c_s}, {c_e})'
                else:
                    c_args = translate_call_args(args_str, sum_types, struct_names, all_variants)
                    result = f'append({c_args})'
            elif fn_name == 'string':
                c_inner = translate_expr(args_str, sum_types, struct_names, all_variants)
                result = f'aria_str_from_bytes({c_inner})'
            elif fn_name == 'float64':
                c_inner = translate_expr(args_str, sum_types, struct_names, all_variants)
                result = f'((double)({c_inner}))'
            else:
                c_args = translate_call_args(args_str, sum_types, struct_names, all_variants)
                result = f'{fn_name}({c_args})'
            if after:
                result += translate_expr(after, sum_types, struct_names, all_variants)
            return result

    # String literal that wasn't caught earlier (bare string in return)
    if expr.startswith('"') and expr.endswith('"'):
        return f'aria_str_lit({expr})'

    # Simple identifier or number
    return expr


def _is_c_string_expr(c_expr):
    """Check if a translated C expression is a string type."""
    c_expr = c_expr.strip()
    if c_expr.startswith('aria_str_lit('):
        return True
    if c_expr.startswith('aria_str_concat('):
        return True
    if c_expr.startswith('aria_str_char_at('):
        return True
    if c_expr.startswith('aria_str_slice('):
        return True
    if c_expr.startswith('aria_str_from_'):
        return True
    if c_expr.startswith('aria_str_replace_all('):
        return True
    if c_expr.startswith('aria_str_trim_prefix('):
        return True
    # Parenthesized string expr
    if c_expr.startswith('(') and c_expr.endswith(')'):
        return _is_c_string_expr(c_expr[1:-1])
    return False


def find_matching_close_paren(expr, start):
    """Find matching close paren for open paren at `start`."""
    depth = 0
    in_string = False
    escape = False
    for i in range(start, len(expr)):
        c = expr[i]
        if escape:
            escape = False
            continue
        if c == '\\' and in_string:
            escape = True
            continue
        if c == '"':
            in_string = not in_string
        elif not in_string:
            if c == '(':
                depth += 1
            elif c == ')':
                depth -= 1
                if depth == 0:
                    return i
    return -1


def find_colon_in_slice(idx):
    """Find the colon separating start:end in a slice expression, skipping colons in nested exprs."""
    depth = 0
    for i, c in enumerate(idx):
        if c in '([':
            depth += 1
        elif c in ')]':
            depth -= 1
        elif c == ':' and depth == 0:
            return i
    return None


def translate_field_chain(expr, sum_types, struct_names, all_variants):
    """Translate a chain of field accesses and method calls."""
    # Check for indexed access: Field[idx]
    m = re.match(r'^(\w+)\[(.+)\](.*)$', expr)
    if m:
        field = m.group(1)
        idx = m.group(2)
        rest = m.group(3)
        if ':' in idx:
            parts = idx.split(':', 1)
            c_start = translate_expr(parts[0], sum_types, struct_names, all_variants)
            c_end = translate_expr(parts[1], sum_types, struct_names, all_variants)
            result = f'/* slice */ {field}[/* TODO */]'
        else:
            c_idx = translate_expr(idx, sum_types, struct_names, all_variants)
            result = f'{field}.data[{c_idx}]'
        if rest and rest.startswith('.'):
            result += f'.{translate_field_chain(rest[1:], sum_types, struct_names, all_variants)}'
        return result

    # Simple field
    m = re.match(r'^(\w+)$', expr)
    if m:
        return m.group(1)

    # Field with further chain
    m = re.match(r'^(\w+)\.(.+)$', expr)
    if m:
        return f'{m.group(1)}.{translate_field_chain(m.group(2), sum_types, struct_names, all_variants)}'

    return expr


def find_matching_paren(text):
    """Given text after opening '(', find the matching ')'.
    Returns (args_str, remaining_text) or (None, None) if not found."""
    depth = 1
    i = 0
    in_string = False
    escape = False
    while i < len(text):
        c = text[i]
        if escape:
            escape = False
            i += 1
            continue
        if c == '\\' and in_string:
            escape = True
            i += 1
            continue
        if c == '"':
            in_string = not in_string
        elif not in_string:
            if c == '(':
                depth += 1
            elif c == ')':
                depth -= 1
                if depth == 0:
                    return text[:i], text[i + 1:]
        i += 1
    return None, None


def translate_call_args(args_str, sum_types, struct_names, all_variants):
    """Translate function call arguments."""
    if not args_str.strip():
        return ''
    args = split_params(args_str)
    c_args = []
    for arg in args:
        c_args.append(translate_expr(arg.strip(), sum_types, struct_names, all_variants))
    return ', '.join(c_args)


def translate_struct_literal_fields(fields_text, sum_types, struct_names, all_variants):
    """Translate struct literal fields: `Field: val, ...` -> `.Field = val, ...`"""
    parts = split_params(fields_text)
    result = []
    for part in parts:
        part = part.strip()
        if ':' in part:
            colon_idx = part.index(':')
            field_name = part[:colon_idx].strip()
            field_val = part[colon_idx + 1:].strip()
            c_val = translate_expr(field_val, sum_types, struct_names, all_variants)
            result.append(f'.{field_name} = {c_val}')
        else:
            result.append(translate_expr(part, sum_types, struct_names, all_variants))
    return ', '.join(result)


def split_on_operator(expr, op):
    """Try to split expr on a binary operator, respecting parens/brackets/strings.
    Returns (lhs, rhs) or None."""
    depth = 0
    in_string = False
    in_rune = False
    escape = False
    i = 0
    op_len = len(op)
    while i < len(expr):
        c = expr[i]
        if escape:
            escape = False
            i += 1
            continue
        if c == '\\' and (in_string or in_rune):
            escape = True
            i += 1
            continue
        if c == '"' and not in_rune:
            in_string = not in_string
        elif c == '\'' and not in_string:
            in_rune = not in_rune
        elif not in_string and not in_rune:
            if c in '([{':
                depth += 1
            elif c in ')]}':
                depth -= 1
            elif depth == 0 and i > 0 and i + op_len <= len(expr):
                if expr[i:i + op_len] == op:
                    # Make sure we're not matching a longer operator
                    # e.g., don't match '=' when looking at '=='
                    if op == '=' and i + 1 < len(expr) and expr[i + 1] == '=':
                        i += 1
                        continue
                    if op == '!' and i + 1 < len(expr) and expr[i + 1] == '=':
                        i += 1
                        continue
                    if op == '<' and i + 1 < len(expr) and expr[i + 1] in '<=':
                        i += 1
                        continue
                    if op == '>' and i + 1 < len(expr) and expr[i + 1] in '>=':
                        i += 1
                        continue
                    if op == '&' and i + 1 < len(expr) and expr[i + 1] == '&':
                        i += 1
                        continue
                    if op == '|' and i + 1 < len(expr) and expr[i + 1] == '|':
                        i += 1
                        continue
                    lhs = expr[:i].strip()
                    rhs = expr[i + op_len:].strip()
                    if lhs and rhs:
                        return (lhs, rhs)
        i += 1
    return None


def is_likely_string(lhs, rhs):
    """Heuristic: is this comparison/concat likely between strings?"""
    for x in (lhs, rhs):
        x = x.strip()
        if x.startswith('"'):
            return True
        if 'Text' in x or 'Name' in x or 'text' in x or 'name' in x:
            return True
        if 'File' in x or 'file' in x or 'path' in x or 'Path' in x:
            return True
        if 'Format' in x or 'format' in x:
            return True
        if 'message' in x or 'Message' in x:
            return True
        if 'Source' in x or 'source' in x:
            return True
        if 'Kind' in x or 'kind' in x:
            return False  # Kind is typically an enum int
    return False


# ---------------------------------------------------------------------------
# Phase 4: Generate C output
# ---------------------------------------------------------------------------

def collect_slice_types(structs, sum_types, struct_names):
    """Collect all slice types used in struct fields."""
    slice_types = set()
    for sname, fields in structs.items():
        for fname, ftype in fields:
            m = re.match(r'^\[\](.+)$', ftype.strip())
            if m:
                inner = m.group(1).strip()
                slice_types.add(inner)
    return sorted(slice_types)


def topo_sort_structs(structs, sum_types, struct_names):
    """Topologically sort structs so dependencies come first.
    Slice fields are pointer-based so don't create hard dependencies."""
    # Build dependency graph: direct struct field usage (not slices)
    deps = {sname: set() for sname in structs}
    for sname, fields in structs.items():
        for fname, ftype in fields:
            ftype = ftype.strip()
            # Direct struct reference (not slice, not sum type)
            if ftype in struct_names and ftype != sname:
                deps[sname].add(ftype)

    # Kahn's algorithm
    order = []
    in_degree = {s: 0 for s in structs}
    for s, d in deps.items():
        for dep in d:
            if dep in in_degree:
                in_degree[dep] = in_degree.get(dep, 0)  # ensure dep exists

    # Recalculate: in_degree[x] = number of structs that x depends on
    # Actually we want: in_degree[x] = number of structs that depend on x? No.
    # We want to emit dependencies first. So in_degree = count of unsatisfied deps.
    in_degree = {s: len(deps[s]) for s in structs}
    queue = [s for s in structs if in_degree[s] == 0]
    visited = set()
    while queue:
        s = queue.pop(0)
        if s in visited:
            continue
        visited.add(s)
        order.append(s)
        # Find structs that depend on s
        for other, other_deps in deps.items():
            if s in other_deps and other not in visited:
                in_degree[other] -= 1
                if in_degree[other] <= 0:
                    queue.append(other)

    # Add any remaining (circular deps)
    for s in structs:
        if s not in visited:
            order.append(s)

    return order


def generate_c_runtime():
    """Generate the C runtime header with AriaStr, AriaSlice, etc."""
    return r"""
/* ================================================================
 * Aria Bootstrap Compiler — Generated C code
 * Transpiled from Go by transpile.py
 * ================================================================ */

#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <stdint.h>
#include <stdarg.h>

/* ----------------------------------------------------------------
 * AriaStr — immutable string (ptr + len)
 * ---------------------------------------------------------------- */
typedef struct {
    const char* ptr;
    int64_t len;
} AriaStr;

static AriaStr aria_str_lit(const char* s) {
    AriaStr r;
    r.ptr = s;
    r.len = (int64_t)strlen(s);
    return r;
}

static AriaStr aria_str_from_len(const char* s, int64_t len) {
    char* buf = (char*)malloc(len + 1);
    memcpy(buf, s, len);
    buf[len] = '\0';
    AriaStr r;
    r.ptr = buf;
    r.len = len;
    return r;
}

static AriaStr aria_str_concat(AriaStr a, AriaStr b) {
    int64_t new_len = a.len + b.len;
    char* buf = (char*)malloc(new_len + 1);
    memcpy(buf, a.ptr, a.len);
    memcpy(buf + a.len, b.ptr, b.len);
    buf[new_len] = '\0';
    AriaStr r;
    r.ptr = buf;
    r.len = new_len;
    return r;
}

static int64_t aria_str_eq(AriaStr a, AriaStr b) {
    if (a.len != b.len) return 0;
    return memcmp(a.ptr, b.ptr, a.len) == 0 ? 1 : 0;
}

static int64_t aria_str_has_prefix(AriaStr s, AriaStr prefix) {
    if (prefix.len > s.len) return 0;
    return memcmp(s.ptr, prefix.ptr, prefix.len) == 0 ? 1 : 0;
}

static int64_t aria_str_contains(AriaStr s, AriaStr sub) {
    if (sub.len == 0) return 1;
    if (sub.len > s.len) return 0;
    for (int64_t i = 0; i <= s.len - sub.len; i++) {
        if (memcmp(s.ptr + i, sub.ptr, sub.len) == 0) return 1;
    }
    return 0;
}

static AriaStr aria_str_char_at(AriaStr s, int64_t idx) {
    if (idx < 0 || idx >= s.len) return aria_str_lit("");
    char buf[2] = {s.ptr[idx], '\0'};
    return aria_str_from_len(buf, 1);
}

static AriaStr aria_str_slice(AriaStr s, int64_t start, int64_t end) {
    if (start < 0) start = 0;
    if (end > s.len) end = s.len;
    if (start >= end) return aria_str_lit("");
    return aria_str_from_len(s.ptr + start, end - start);
}

static AriaStr aria_str_replace_all(AriaStr s, AriaStr old, AriaStr new_s) {
    if (old.len == 0) return s;
    // Count occurrences
    int64_t count = 0;
    for (int64_t i = 0; i <= s.len - old.len; i++) {
        if (memcmp(s.ptr + i, old.ptr, old.len) == 0) { count++; i += old.len - 1; }
    }
    if (count == 0) return s;
    int64_t new_len = s.len + count * (new_s.len - old.len);
    char* buf = (char*)malloc(new_len + 1);
    int64_t j = 0;
    for (int64_t i = 0; i < s.len; ) {
        if (i <= s.len - old.len && memcmp(s.ptr + i, old.ptr, old.len) == 0) {
            memcpy(buf + j, new_s.ptr, new_s.len);
            j += new_s.len;
            i += old.len;
        } else {
            buf[j++] = s.ptr[i++];
        }
    }
    buf[new_len] = '\0';
    return aria_str_from_len(buf, new_len);
}

static AriaStr aria_str_trim_prefix(AriaStr s, AriaStr prefix) {
    if (aria_str_has_prefix(s, prefix)) {
        return aria_str_slice(s, prefix.len, s.len);
    }
    return s;
}

static void aria_println(AriaStr s) {
    fwrite(s.ptr, 1, s.len, stdout);
    fputc('\n', stdout);
}

static void aria_eprint(AriaStr s) {
    fwrite(s.ptr, 1, s.len, stderr);
}

static AriaStr aria_str_from_bytes(AriaStr s) {
    return s; /* already a string */
}

/* i2s — int to string. Defined in the Aria code but we need a C fallback. */
static AriaStr aria_i2s(int64_t n) {
    char buf[32];
    snprintf(buf, sizeof(buf), "%lld", (long long)n);
    return aria_str_from_len(buf, strlen(buf));
}

/* ----------------------------------------------------------------
 * AriaSlice — generic growable array (data + len + cap)
 * We use macros to generate typed versions.
 * ---------------------------------------------------------------- */
/* Slice type declaration (just the struct typedef) */
#define ARIA_DECLARE_SLICE(T, NAME) \
    typedef struct { T* data; int64_t len; int64_t cap; } AriaSlice_##NAME;

/* Slice append function (needs complete type) */
#define ARIA_DEFINE_SLICE_FUNCS(T, NAME) \
    static AriaSlice_##NAME aria_append_##NAME(AriaSlice_##NAME s, T elem) { \
        if (s.len >= s.cap) { \
            int64_t new_cap = s.cap < 8 ? 8 : s.cap * 2; \
            T* new_data = (T*)malloc(sizeof(T) * new_cap); \
            if (s.data) memcpy(new_data, s.data, sizeof(T) * s.len); \
            s.data = new_data; \
            s.cap = new_cap; \
        } \
        s.data[s.len++] = elem; \
        return s; \
    }

/* Combined macro for primitive types */
#define ARIA_DEFINE_SLICE(T, NAME) \
    ARIA_DECLARE_SLICE(T, NAME) \
    ARIA_DEFINE_SLICE_FUNCS(T, NAME)

/* ----------------------------------------------------------------
 * File I/O helpers
 * ---------------------------------------------------------------- */
static AriaStr _ariaReadFile(AriaStr path) {
    FILE* f = fopen(path.ptr, "rb");
    if (!f) { fprintf(stderr, "readFile failed: %s\n", path.ptr); exit(1); }
    fseek(f, 0, SEEK_END);
    long sz = ftell(f);
    fseek(f, 0, SEEK_SET);
    char* buf = (char*)malloc(sz + 1);
    fread(buf, 1, sz, f);
    buf[sz] = '\0';
    fclose(f);
    return aria_str_from_len(buf, sz);
}

static void _ariaWriteFile(AriaStr path, AriaStr content) {
    FILE* f = fopen(path.ptr, "w");
    if (!f) { fprintf(stderr, "writeFile failed: %s\n", path.ptr); exit(1); }
    fwrite(content.ptr, 1, content.len, f);
    fclose(f);
}

static int64_t _ariaFileExists(AriaStr path) {
    FILE* f = fopen(path.ptr, "r");
    if (f) { fclose(f); return 1; }
    return 0;
}

static AriaStr _ariaGetenv(AriaStr name) {
    const char* v = getenv(name.ptr);
    if (!v) return aria_str_lit("");
    return aria_str_lit(v);
}

static int64_t _ariaExec(AriaStr command) {
    int r = system(command.ptr);
    return (int64_t)WEXITSTATUS(r);
}

static int64_t _ariaIsDir(AriaStr path) {
    /* Simplified check */
    return 0;
}

static AriaStr _ariaParseInt_str(AriaStr s) {
    /* Stub — the real i2s is in the Aria code */
    return s;
}

static int64_t _ariaParseInt(AriaStr s) {
    return (int64_t)atoll(s.ptr);
}

static double _ariaParseFloat(AriaStr s) {
    return atof(s.ptr);
}

"""


def generate_c_file(sum_types, structs, global_vars, functions, all_source, fn_ret_types=None):
    """Generate the complete C file."""
    struct_names = set(structs.keys())
    all_variants = {}  # variant_name -> enum_value
    for iface_name, variants in sum_types.items():
        for i, v in enumerate(variants):
            all_variants[v] = i

    if fn_ret_types is None:
        fn_ret_types = {}

    # Store fn_ret_types globally for use in type inference
    global _fn_ret_types
    _fn_ret_types = fn_ret_types

    out = []
    out.append(generate_c_runtime())

    # Generate sum type enums
    out.append('/* ================================================================')
    out.append(' * Sum type enums')
    out.append(' * ================================================================ */')
    for iface_name, variants in sum_types.items():
        out.append(f'/* {iface_name} */')
        for i, v in enumerate(variants):
            out.append(f'#define {v} ((int64_t){i})')
        out.append('')

    # Collect all slice types needed
    slice_inner_types = set()
    for sname, fields in structs.items():
        for fname, ftype in fields:
            m = re.match(r'^\[\](.+)$', ftype.strip())
            if m:
                slice_inner_types.add(m.group(1).strip())

    # Forward-declare structs
    out.append('/* ================================================================')
    out.append(' * Forward struct declarations')
    out.append(' * ================================================================ */')
    for sname in structs:
        out.append(f'typedef struct {sname} {sname};')
    out.append('')

    # Declare ALL slice types (just the typedef, before struct defs)
    out.append('/* ================================================================')
    out.append(' * Slice type declarations')
    out.append(' * ================================================================ */')
    defined_slices = set()
    # Primitive slice types (full define = typedef + functions)
    for basic in ['int64_t', 'AriaStr', 'double', 'uint8_t']:
        out.append(f'ARIA_DEFINE_SLICE({basic}, {basic})')
        defined_slices.add(basic)
    # Struct slice types (declaration only — functions come after struct defs)
    struct_slice_types = []
    for inner in sorted(slice_inner_types):
        c_inner = go_type_to_c(inner, sum_types, struct_names)
        c_safe = safe_c_name(inner)
        if c_safe not in defined_slices:
            out.append(f'ARIA_DECLARE_SLICE({c_inner}, {c_safe})')
            defined_slices.add(c_safe)
            struct_slice_types.append((c_inner, c_safe))
    out.append('')

    # Struct definitions
    out.append('/* ================================================================')
    out.append(' * Struct definitions')
    out.append(' * ================================================================ */')
    struct_order = topo_sort_structs(structs, sum_types, struct_names)
    for sname in struct_order:
        fields = structs[sname]
        out.append(f'struct {sname} {{')
        for fname, ftype in fields:
            c_type = go_type_to_c(ftype, sum_types, struct_names)
            out.append(f'    {c_type} {fname};')
        out.append('};')
        out.append('')

    # Now define slice FUNCTIONS for struct types (after struct definitions)
    out.append('/* ================================================================')
    out.append(' * Slice functions (struct types)')
    out.append(' * ================================================================ */')
    for c_inner, c_safe in struct_slice_types:
        out.append(f'ARIA_DEFINE_SLICE_FUNCS({c_inner}, {c_safe})')
    out.append('')

    # Nested slice types (e.g., [][]string -> AriaSlice_AriaSlice_AriaStr)
    for sname, fields in structs.items():
        for fname, ftype in fields:
            m = re.match(r'^\[\]\[\](.+)$', ftype.strip())
            if m:
                inner = m.group(1).strip()
                outer_name = f'AriaSlice_{safe_c_name(inner)}'
                slice_of_slice = f'AriaSlice_{outer_name}'
                if slice_of_slice not in defined_slices:
                    out.append(f'ARIA_DECLARE_SLICE({outer_name}, {outer_name})')
                    out.append(f'ARIA_DEFINE_SLICE_FUNCS({outer_name}, {outer_name})')
                    defined_slices.add(slice_of_slice)

    # Generate aria_append_generic macro using _Generic
    # This maps each AriaSlice_X type to aria_append_X
    out.append('/* Generic append dispatch */')
    generic_entries = []
    for c_inner, c_safe in struct_slice_types:
        generic_entries.append(f'AriaSlice_{c_safe}: aria_append_{c_safe}')
    # Also add primitive types
    for basic in ['int64_t', 'AriaStr', 'double', 'uint8_t']:
        generic_entries.append(f'AriaSlice_{basic}: aria_append_{basic}')
    if generic_entries:
        entries_str = ', \\\n    '.join(generic_entries)
        out.append(f'#define aria_append_generic(s, elem) _Generic((s), \\\n    {entries_str})(s, elem)')
    out.append('')

    # Slice helper for struct types — create from varargs
    for inner in sorted(slice_inner_types):
        c_inner = go_type_to_c(inner, sum_types, struct_names)
        c_safe = safe_c_name(inner)
        out.append(f'static AriaSlice_{c_safe} aria_slice_from_{c_safe}(int64_t count, ...) {{')
        out.append(f'    AriaSlice_{c_safe} s = {{NULL, 0, 0}};')
        out.append(f'    va_list args;')
        out.append(f'    va_start(args, count);')
        out.append(f'    for (int64_t i = 0; i < count; i++) {{')
        out.append(f'        s = aria_append_{c_safe}(s, va_arg(args, {c_inner}));')
        out.append(f'    }}')
        out.append(f'    va_end(args);')
        out.append(f'    return s;')
        out.append(f'}}')
        out.append('')

    # Global variables
    out.append('/* ================================================================')
    out.append(' * Global variables')
    out.append(' * ================================================================ */')
    for var_name, var_value in global_vars:
        c_val = translate_expr(var_value, sum_types, struct_names, all_variants)
        if var_value.startswith('"'):
            out.append(f'AriaStr {var_name};  /* initialized in main */')
        elif var_value.startswith('int64(') or re.match(r'^-?\d+$', var_value):
            out.append(f'int64_t {var_name} = {c_val};')
        else:
            out.append(f'int64_t {var_name} = {c_val};')
    out.append('')

    # Forward-declare all functions
    out.append('/* ================================================================')
    out.append(' * Function forward declarations')
    out.append(' * ================================================================ */')
    for fn in functions:
        ret_c = go_type_to_c(fn['ret'], sum_types, struct_names) if fn['ret'] else 'void'
        params = parse_params(fn['params'], sum_types, struct_names)
        params_str = ', '.join(f'{t} {n}' for n, t in params) if params else 'void'
        out.append(f'{ret_c} {fn["name"]}({params_str});')
    out.append('')

    # Function implementations
    out.append('/* ================================================================')
    out.append(' * Function implementations')
    out.append(' * ================================================================ */')
    for fn in functions:
        ret_c = go_type_to_c(fn['ret'], sum_types, struct_names) if fn['ret'] else 'void'
        params = parse_params(fn['params'], sum_types, struct_names)
        params_str = ', '.join(f'{t} {n}' for n, t in params) if params else 'void'
        out.append(f'/* from {fn["source_file"]} */')
        out.append(f'{ret_c} {fn["name"]}({params_str}) {{')
        c_body = translate_body(fn['body'], ret_c, sum_types, struct_names, all_variants)
        out.append(c_body)
        out.append('}')
        out.append('')

    # Generate main()
    out.append('/* ================================================================')
    out.append(' * main()')
    out.append(' * ================================================================ */')
    out.append('int main(int argc, char** argv) {')
    # Initialize string globals
    for var_name, var_value in global_vars:
        if var_value.startswith('"'):
            out.append(f'    {var_name} = aria_str_lit({var_value});')
    out.append('')
    out.append('    /* Build args slice */')
    out.append('    AriaSlice_AriaStr args = {NULL, 0, 0};')
    out.append('    args = aria_append_AriaStr(args, aria_str_lit(""));  /* sentinel */')
    out.append('    for (int i = 0; i < argc; i++) {')
    out.append('        args = aria_append_AriaStr(args, aria_str_lit(argv[i]));')
    out.append('    }')
    out.append('')
    out.append('    if (args.len < 3) {')
    out.append('        print_help();')
    out.append('    } else {')
    out.append('        Options opts = parse_args(args);')
    out.append('        AriaStr cmd = command_to_str(opts.Command);')
    out.append('        if (opts.Help_requested) {')
    out.append('            if (aria_str_eq(cmd, aria_str_lit("build"))) { _print_build_help(); }')
    out.append('            else if (aria_str_eq(cmd, aria_str_lit("run"))) { _print_run_help(); }')
    out.append('            else if (aria_str_eq(cmd, aria_str_lit("check"))) { _print_check_help(); }')
    out.append('            else if (aria_str_eq(cmd, aria_str_lit("test"))) { _print_test_help(); }')
    out.append('            else { print_help(); }')
    out.append('        } else if (aria_str_eq(cmd, aria_str_lit("build"))) {')
    out.append('            run_build(opts.Files, opts.Format, opts.Target, opts.Runtime_path, opts.Output_path);')
    out.append('        } else if (aria_str_eq(cmd, aria_str_lit("run"))) {')
    out.append('            run_run(opts.Files, opts.Format, opts.Target, opts.Runtime_path, opts.Output_path);')
    out.append('        } else if (aria_str_eq(cmd, aria_str_lit("check"))) {')
    out.append('            run_check(opts.Files, opts.Format);')
    out.append('        } else if (aria_str_eq(cmd, aria_str_lit("test"))) {')
    out.append('            run_test(opts.Files, opts.Format, opts.Runtime_path, opts.Parallel);')
    out.append('        } else if (aria_str_eq(cmd, aria_str_lit("version"))) {')
    out.append('            print_version();')
    out.append('        } else { print_help(); }')
    out.append('    }')
    out.append('    return 0;')
    out.append('}')

    return '\n'.join(out)


# ---------------------------------------------------------------------------
# Main
# ---------------------------------------------------------------------------

def main():
    go_dir = Path(__file__).parent / "go_source"
    if not go_dir.exists():
        print("Error: go_source/ directory not found", file=sys.stderr)
        sys.exit(1)

    go_files = sorted(go_dir.glob("*.go"))
    print(f"Found {len(go_files)} Go files", file=sys.stderr)

    # Read all Go source
    all_source = OrderedDict()
    for f in go_files:
        all_source[f.stem] = f.read_text()

    total_lines = sum(s.count('\n') for s in all_source.values())
    print(f"Total: {total_lines} lines of Go", file=sys.stderr)

    # Phase 1: Extract types
    sum_types = parse_sum_types(all_source)
    print(f"Sum types: {len(sum_types)}", file=sys.stderr)
    for name, variants in sum_types.items():
        print(f"  {name}: {len(variants)} variants", file=sys.stderr)

    structs = parse_structs(all_source, sum_types)
    print(f"Structs: {len(structs)}", file=sys.stderr)

    global_vars = parse_global_vars(all_source)
    print(f"Global vars: {len(global_vars)}", file=sys.stderr)

    struct_names = set(structs.keys())

    # Phase 2: Extract functions
    functions = parse_functions(all_source, sum_types, struct_names)
    print(f"Functions: {len(functions)}", file=sys.stderr)

    # Build function return type map
    fn_ret_types = {}
    for fn in functions:
        ret = fn['ret'].strip() if fn['ret'] else ''
        fn_ret_types[fn['name']] = go_type_to_c(ret, sum_types, struct_names) if ret else 'void'

    # Phase 3: Generate C
    c_code = generate_c_file(sum_types, structs, global_vars, functions, all_source, fn_ret_types)

    # Output
    print(c_code)

    c_lines = c_code.count('\n')
    print(f"\nGenerated {c_lines} lines of C", file=sys.stderr)


if __name__ == '__main__':
    main()
