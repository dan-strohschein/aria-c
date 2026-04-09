# aria-c -- C Bootstrap Compiler for Aria

The C bootstrap compiler (gen0) for the [Aria programming language](https://github.com/dan-strohschein/aria). This compiler enables the Aria self-hosting chain without depending on the Go bootstrap compiler.

## What This Is

`aria-c` is a single-file C compiler for Aria, produced by a two-stage pipeline:

1. **Go source** (`go_source/`): The Aria bootstrap compiler (`aria-compiler-go`) compiles the self-hosting Aria compiler source into Go. These Go files are a mechanical translation -- not handwritten.
2. **C transpiler** (`transpile.py`): A Python script translates the generated Go code into a single C file (`aria_gen.c`), mapping Go patterns (interfaces, type switches, slices) to C equivalents.
3. **C compiler** (`gen0_c`): Compile `aria_gen.c` with any C compiler to get the gen0 binary.

## The Self-Hosting Chain

```
aria-c/gen0_c  (C binary, compiled from aria_gen.c)
    |
    |  reads Aria source, emits LLVM IR
    v
gen1           (native binary, compiled by gen0)
    |
    |  reads same Aria source, emits LLVM IR
    v
gen2           (native binary, compiled by gen1)
    |
    |  reads same Aria source, emits LLVM IR
    v
gen3           (native binary, compiled by gen2)

gen2.ll == gen3.ll   (byte-identical -- codegen fidelity achieved)
```

Once gen2 and gen3 produce identical output, the bootstrap compiler is no longer needed. The Aria compiler is fully self-hosting.

## Building

### Prerequisites

- A C compiler (clang or gcc)
- Python 3 (only if regenerating `aria_gen.c` from Go source)

### Build gen0

```bash
cc -O1 -o gen0_c aria_gen.c -Wno-everything
```

### Use gen0 to build gen1

```bash
# From the aria/ repo directory:
ARIA_GC_THRESHOLD=off ../aria-c/gen0_c build \
  src/main.aria src/lexer/token.aria src/lexer/lexer.aria \
  src/parser/ast.aria src/parser/parser.aria src/parser/precedence.aria \
  src/resolver/scope.aria src/resolver/resolver.aria \
  src/checker/types.aria src/checker/traits.aria src/checker/checker.aria \
  src/codegen/ir.aria src/codegen/llvm.aria src/codegen/lower.aria \
  src/codegen/codegen.aria src/diagnostic/diagnostic.aria \
  -o gen1 --runtime=runtime/runtime.c

clang gen1.ll runtime/runtime.c -o gen1 -O1 -Wno-override-module -lpthread
```

### Regenerate aria_gen.c (optional)

If the Go source files are updated:

```bash
python3 transpile.py
```

This reads `go_source/*.go` and produces `aria_gen.c`.

## Project Structure

```
aria-c/
├── aria_gen.c       # Generated C source (~38K lines)
├── transpile.py     # Go -> C transpiler
├── go_source/       # Generated Go files from aria-compiler-go
│   ├── aria_gen_main.go
│   ├── aria_gen_lexer.go
│   ├── aria_gen_parser.go
│   ├── aria_gen_resolver.go
│   ├── aria_gen_checker.go
│   ├── aria_gen_lower.go
│   ├── aria_gen_llvm.go
│   └── ...
└── gen0_c           # Compiled binary (not checked in)
```

## Why C?

The Go bootstrap compiler (`aria-compiler-go`) transpiles Aria to Go, then calls `go build`. This works but introduces a dependency on the Go toolchain. By transpiling the generated Go to C, we get a bootstrap compiler that only needs a C compiler -- available everywhere.

The C code is not meant to be human-maintained. It's a mechanical translation that preserves the exact semantics of the Aria compiler. The real source of truth is the Aria source code in the `aria` repo.

## Related Repositories

- [aria](https://github.com/dan-strohschein/aria) -- The self-hosting Aria compiler (written in Aria)
- [aria-compiler-go](https://github.com/dan-strohschein/aria-compiler-go) -- The Go bootstrap compiler
- [aria-docs](https://github.com/dan-strohschein/aria-docs) -- Language specification and design documents
