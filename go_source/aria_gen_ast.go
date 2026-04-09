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

type ExprKind interface {
	isExprKind()
}

type ExprKindEkIntLit struct{}
func (ExprKindEkIntLit) isExprKind() {}

type ExprKindEkFloatLit struct{}
func (ExprKindEkFloatLit) isExprKind() {}

type ExprKindEkStringLit struct{}
func (ExprKindEkStringLit) isExprKind() {}

type ExprKindEkBoolLit struct{}
func (ExprKindEkBoolLit) isExprKind() {}

type ExprKindEkIdent struct{}
func (ExprKindEkIdent) isExprKind() {}

type ExprKindEkBinary struct{}
func (ExprKindEkBinary) isExprKind() {}

type ExprKindEkUnary struct{}
func (ExprKindEkUnary) isExprKind() {}

type ExprKindEkCall struct{}
func (ExprKindEkCall) isExprKind() {}

type ExprKindEkMethodCall struct{}
func (ExprKindEkMethodCall) isExprKind() {}

type ExprKindEkFieldAccess struct{}
func (ExprKindEkFieldAccess) isExprKind() {}

type ExprKindEkIndex struct{}
func (ExprKindEkIndex) isExprKind() {}

type ExprKindEkBlock struct{}
func (ExprKindEkBlock) isExprKind() {}

type ExprKindEkIf struct{}
func (ExprKindEkIf) isExprKind() {}

type ExprKindEkMatch struct{}
func (ExprKindEkMatch) isExprKind() {}

type ExprKindEkClosure struct{}
func (ExprKindEkClosure) isExprKind() {}

type ExprKindEkArray struct{}
func (ExprKindEkArray) isExprKind() {}

type ExprKindEkStruct struct{}
func (ExprKindEkStruct) isExprKind() {}

type ExprKindEkAssign struct{}
func (ExprKindEkAssign) isExprKind() {}

type ExprKindEkBinding struct{}
func (ExprKindEkBinding) isExprKind() {}

type ExprKindEkReturn struct{}
func (ExprKindEkReturn) isExprKind() {}

type ExprKindEkBreak struct{}
func (ExprKindEkBreak) isExprKind() {}

type ExprKindEkContinue struct{}
func (ExprKindEkContinue) isExprKind() {}

type ExprKindEkFor struct{}
func (ExprKindEkFor) isExprKind() {}

type ExprKindEkWhile struct{}
func (ExprKindEkWhile) isExprKind() {}

type ExprKindEkLoop struct{}
func (ExprKindEkLoop) isExprKind() {}

type ExprKindEkDefer struct{}
func (ExprKindEkDefer) isExprKind() {}

type ExprKindEkWith struct{}
func (ExprKindEkWith) isExprKind() {}

type ExprKindEkTuple struct{}
func (ExprKindEkTuple) isExprKind() {}

type ExprKindEkCatch struct{}
func (ExprKindEkCatch) isExprKind() {}

type ExprKindEkPipeline struct{}
func (ExprKindEkPipeline) isExprKind() {}

type ExprKindEkRange struct{}
func (ExprKindEkRange) isExprKind() {}

type ExprKindEkPropagate struct{}
func (ExprKindEkPropagate) isExprKind() {}

type ExprKindEkAssertOk struct{}
func (ExprKindEkAssertOk) isExprKind() {}

type ExprKindEkFnDecl struct{}
func (ExprKindEkFnDecl) isExprKind() {}

type ExprKindEkTypeDecl struct{}
func (ExprKindEkTypeDecl) isExprKind() {}

type ExprKindEkEnumDecl struct{}
func (ExprKindEkEnumDecl) isExprKind() {}

type ExprKindEkTraitDecl struct{}
func (ExprKindEkTraitDecl) isExprKind() {}

type ExprKindEkImplDecl struct{}
func (ExprKindEkImplDecl) isExprKind() {}

type ExprKindEkConstDecl struct{}
func (ExprKindEkConstDecl) isExprKind() {}

type ExprKindEkUseDecl struct{}
func (ExprKindEkUseDecl) isExprKind() {}

type ExprKindEkParam struct{}
func (ExprKindEkParam) isExprKind() {}

type ExprKindEkField struct{}
func (ExprKindEkField) isExprKind() {}

type ExprKindEkVariant struct{}
func (ExprKindEkVariant) isExprKind() {}

type ExprKindEkListComp struct{}
func (ExprKindEkListComp) isExprKind() {}

type ExprKindEkYield struct{}
func (ExprKindEkYield) isExprKind() {}

type ExprKindEkSelect struct{}
func (ExprKindEkSelect) isExprKind() {}

type ExprKindEkNone struct{}
func (ExprKindEkNone) isExprKind() {}

type ExprKindEkDurationLit struct{}
func (ExprKindEkDurationLit) isExprKind() {}

type ExprKindEkSizeLit struct{}
func (ExprKindEkSizeLit) isExprKind() {}

type ExprKindEkTypeRef struct{}
func (ExprKindEkTypeRef) isExprKind() {}

type ExprKindEkTypeArray struct{}
func (ExprKindEkTypeArray) isExprKind() {}

type ExprKindEkTypeOptional struct{}
func (ExprKindEkTypeOptional) isExprKind() {}

type ExprKindEkTypeResult struct{}
func (ExprKindEkTypeResult) isExprKind() {}

type ExprKindEkTypeFn struct{}
func (ExprKindEkTypeFn) isExprKind() {}

type ExprKindEkTypeTuple struct{}
func (ExprKindEkTypeTuple) isExprKind() {}

type ExprKindEkMapLit struct{}
func (ExprKindEkMapLit) isExprKind() {}

type ExprKindEkSetLit struct{}
func (ExprKindEkSetLit) isExprKind() {}

type Expr struct {
	Kind ExprKind
	S1 string
	S2 string
	B1 bool
	B2 bool
	C0 int64
	C1 int64
	C2 int64
	List_start int64
	List_count int64
	Span Span
}

type NodePool struct {
	Nodes []Expr
}

type DeclKind interface {
	isDeclKind()
}

type DeclKindDkFn struct{}
func (DeclKindDkFn) isDeclKind() {}

type DeclKindDkType struct{}
func (DeclKindDkType) isDeclKind() {}

type DeclKindDkEnum struct{}
func (DeclKindDkEnum) isDeclKind() {}

type DeclKindDkTrait struct{}
func (DeclKindDkTrait) isDeclKind() {}

type DeclKindDkImpl struct{}
func (DeclKindDkImpl) isDeclKind() {}

type DeclKindDkConst struct{}
func (DeclKindDkConst) isDeclKind() {}

type DeclKindDkEntry struct{}
func (DeclKindDkEntry) isDeclKind() {}

type DeclKindDkTest struct{}
func (DeclKindDkTest) isDeclKind() {}

type DeclKindDkUse struct{}
func (DeclKindDkUse) isDeclKind() {}

type DeclKindDkNone struct{}
func (DeclKindDkNone) isDeclKind() {}

type DeclInfo struct {
	Kind DeclKind
	Name string
	Token_start int64
	Body_start int64
	Body_end int64
	Is_pub bool
	Node_idx int64
}

type DeclIndex struct {
	Decls []DeclInfo
}

type ParseResult struct {
	Index DeclIndex
	Pool NodePool
	Diagnostics DiagnosticBag
}

type Program struct {
	Module_name string
	Node_count int64
}

func expr_kind_name(k ExprKind) string {
	return func() interface{} {
		switch _tmp1 := (k).(type) {
		case ExprKindEkIntLit:
			return "IntLit"
		case ExprKindEkFloatLit:
			return "FloatLit"
		case ExprKindEkStringLit:
			return "StringLit"
		case ExprKindEkBoolLit:
			return "BoolLit"
		case ExprKindEkIdent:
			return "Ident"
		case ExprKindEkBinary:
			return "Binary"
		case ExprKindEkUnary:
			return "Unary"
		case ExprKindEkCall:
			return "Call"
		case ExprKindEkMethodCall:
			return "MethodCall"
		case ExprKindEkFieldAccess:
			return "FieldAccess"
		case ExprKindEkIndex:
			return "Index"
		case ExprKindEkBlock:
			return "Block"
		case ExprKindEkIf:
			return "If"
		case ExprKindEkMatch:
			return "Match"
		case ExprKindEkClosure:
			return "Closure"
		case ExprKindEkArray:
			return "Array"
		case ExprKindEkStruct:
			return "Struct"
		case ExprKindEkAssign:
			return "Assign"
		case ExprKindEkBinding:
			return "Binding"
		case ExprKindEkReturn:
			return "Return"
		case ExprKindEkBreak:
			return "Break"
		case ExprKindEkContinue:
			return "Continue"
		case ExprKindEkFor:
			return "For"
		case ExprKindEkWhile:
			return "While"
		case ExprKindEkLoop:
			return "Loop"
		case ExprKindEkDefer:
			return "Defer"
		case ExprKindEkWith:
			return "With"
		case ExprKindEkTuple:
			return "Tuple"
		case ExprKindEkCatch:
			return "Catch"
		case ExprKindEkPipeline:
			return "Pipeline"
		case ExprKindEkRange:
			return "Range"
		case ExprKindEkPropagate:
			return "Propagate"
		case ExprKindEkAssertOk:
			return "AssertOk"
		case ExprKindEkFnDecl:
			return "FnDecl"
		case ExprKindEkTypeDecl:
			return "TypeDecl"
		case ExprKindEkEnumDecl:
			return "EnumDecl"
		case ExprKindEkTraitDecl:
			return "TraitDecl"
		case ExprKindEkImplDecl:
			return "ImplDecl"
		case ExprKindEkConstDecl:
			return "ConstDecl"
		case ExprKindEkUseDecl:
			return "UseDecl"
		case ExprKindEkParam:
			return "Param"
		case ExprKindEkField:
			return "Field"
		case ExprKindEkVariant:
			return "Variant"
		case ExprKindEkListComp:
			return "ListComp"
		case ExprKindEkYield:
			return "Yield"
		case ExprKindEkSelect:
			return "Select"
		case ExprKindEkNone:
			return "None"
		case ExprKindEkDurationLit:
			return "DurationLit"
		case ExprKindEkSizeLit:
			return "SizeLit"
		case ExprKindEkTypeRef:
			return "TypeRef"
		case ExprKindEkTypeArray:
			return "TypeArray"
		case ExprKindEkTypeOptional:
			return "TypeOptional"
		case ExprKindEkTypeResult:
			return "TypeResult"
		case ExprKindEkTypeFn:
			return "TypeFn"
		case ExprKindEkTypeTuple:
			return "TypeTuple"
		case ExprKindEkMapLit:
			return "MapLit"
		case ExprKindEkSetLit:
			return "SetLit"
		default:
			_ = _tmp1
		}
		return nil
	}().(string)
}

func _no_span() Span {
	return new_span("", int64(0), int64(0), int64(0), int64(0))
}

func _sentinel_expr() Expr {
	return Expr{Kind: ExprKindEkNone{}, S1: "", S2: "", B1: false, B2: false, C0: int64(0), C1: int64(0), C2: int64(0), List_start: int64(0), List_count: int64(0), Span: _no_span()}
}

func new_pool() NodePool {
	return NodePool{Nodes: []Expr{_sentinel_expr()}}
}

func pool_add(pool NodePool, e Expr) NodePool {
	return NodePool{Nodes: append(pool.Nodes, e)}
}

func pool_get(pool NodePool, idx int64) Expr {
	if ((idx >= int64(0)) && (idx < int64(len(pool.Nodes)))) {
		return pool.Nodes[idx]
	}
	return _sentinel_expr()
}

func pool_size(pool NodePool) int64 {
	return int64(len(pool.Nodes))
}

func _no_expr() Expr {
	return _sentinel_expr()
}

func mk_int_lit(value string, span Span) Expr {
	return Expr{Kind: ExprKindEkIntLit{}, S1: value, S2: "", B1: false, B2: false, C0: int64(0), C1: int64(0), C2: int64(0), List_start: int64(0), List_count: int64(0), Span: span}
}

func mk_float_lit(value string, span Span) Expr {
	return Expr{Kind: ExprKindEkFloatLit{}, S1: value, S2: "", B1: false, B2: false, C0: int64(0), C1: int64(0), C2: int64(0), List_start: int64(0), List_count: int64(0), Span: span}
}

func mk_string_lit(value string, span Span) Expr {
	return Expr{Kind: ExprKindEkStringLit{}, S1: value, S2: "", B1: false, B2: false, C0: int64(0), C1: int64(0), C2: int64(0), List_start: int64(0), List_count: int64(0), Span: span}
}

func mk_duration_lit(value string, span Span) Expr {
	return Expr{Kind: ExprKindEkDurationLit{}, S1: value, S2: "", B1: false, B2: false, C0: int64(0), C1: int64(0), C2: int64(0), List_start: int64(0), List_count: int64(0), Span: span}
}

func mk_size_lit(value string, span Span) Expr {
	return Expr{Kind: ExprKindEkSizeLit{}, S1: value, S2: "", B1: false, B2: false, C0: int64(0), C1: int64(0), C2: int64(0), List_start: int64(0), List_count: int64(0), Span: span}
}

func mk_bool_lit(value bool, span Span) Expr {
	return Expr{Kind: ExprKindEkBoolLit{}, S1: "", S2: "", B1: value, B2: false, C0: int64(0), C1: int64(0), C2: int64(0), List_start: int64(0), List_count: int64(0), Span: span}
}

func mk_ident(name string, span Span) Expr {
	return Expr{Kind: ExprKindEkIdent{}, S1: name, S2: "", B1: false, B2: false, C0: int64(0), C1: int64(0), C2: int64(0), List_start: int64(0), List_count: int64(0), Span: span}
}

func mk_binary(op string, left Expr, right Expr, span Span) Expr {
	return Expr{Kind: ExprKindEkBinary{}, S1: op, S2: "", B1: false, B2: false, C0: int64(0), C1: int64(0), C2: int64(0), List_start: int64(0), List_count: int64(0), Span: span}
}

func mk_unary(op string, operand Expr, span Span) Expr {
	return Expr{Kind: ExprKindEkUnary{}, S1: op, S2: "", B1: false, B2: false, C0: int64(0), C1: int64(0), C2: int64(0), List_start: int64(0), List_count: int64(0), Span: span}
}

func mk_call(callee Expr, args_count int64, span Span) Expr {
	return Expr{Kind: ExprKindEkCall{}, S1: callee.S1, S2: "", B1: false, B2: false, C0: int64(0), C1: int64(0), C2: int64(0), List_start: int64(0), List_count: args_count, Span: span}
}

func mk_method_call(method string, span Span) Expr {
	return Expr{Kind: ExprKindEkMethodCall{}, S1: method, S2: "", B1: false, B2: false, C0: int64(0), C1: int64(0), C2: int64(0), List_start: int64(0), List_count: int64(0), Span: span}
}

func mk_field_access(field string, span Span) Expr {
	return Expr{Kind: ExprKindEkFieldAccess{}, S1: field, S2: "", B1: false, B2: false, C0: int64(0), C1: int64(0), C2: int64(0), List_start: int64(0), List_count: int64(0), Span: span}
}

func mk_index(span Span) Expr {
	return Expr{Kind: ExprKindEkIndex{}, S1: "", S2: "", B1: false, B2: false, C0: int64(0), C1: int64(0), C2: int64(0), List_start: int64(0), List_count: int64(0), Span: span}
}

func mk_block(count int64, span Span) Expr {
	return Expr{Kind: ExprKindEkBlock{}, S1: "", S2: "", B1: false, B2: false, C0: int64(0), C1: int64(0), C2: int64(0), List_start: int64(0), List_count: count, Span: span}
}

func mk_if(has_else bool, span Span) Expr {
	return Expr{Kind: ExprKindEkIf{}, S1: "", S2: "", B1: has_else, B2: false, C0: int64(0), C1: int64(0), C2: int64(0), List_start: int64(0), List_count: int64(0), Span: span}
}

func mk_match(arms_count int64, span Span) Expr {
	return Expr{Kind: ExprKindEkMatch{}, S1: "", S2: "", B1: false, B2: false, C0: int64(0), C1: int64(0), C2: int64(0), List_start: int64(0), List_count: arms_count, Span: span}
}

func mk_array(count int64, span Span) Expr {
	return Expr{Kind: ExprKindEkArray{}, S1: "", S2: "", B1: false, B2: false, C0: int64(0), C1: int64(0), C2: int64(0), List_start: int64(0), List_count: count, Span: span}
}

func mk_struct_lit(name string, field_count int64, span Span) Expr {
	return Expr{Kind: ExprKindEkStruct{}, S1: name, S2: "", B1: false, B2: false, C0: int64(0), C1: int64(0), C2: int64(0), List_start: int64(0), List_count: field_count, Span: span}
}

func mk_map_lit(pair_count int64, span Span) Expr {
	return Expr{Kind: ExprKindEkMapLit{}, S1: "", S2: "", B1: false, B2: false, C0: int64(0), C1: int64(0), C2: int64(0), List_start: int64(0), List_count: pair_count, Span: span}
}

func mk_set_lit(elem_count int64, span Span) Expr {
	return Expr{Kind: ExprKindEkSetLit{}, S1: "", S2: "", B1: false, B2: false, C0: int64(0), C1: int64(0), C2: int64(0), List_start: int64(0), List_count: elem_count, Span: span}
}

func mk_assign(span Span) Expr {
	return Expr{Kind: ExprKindEkAssign{}, S1: "", S2: "", B1: false, B2: false, C0: int64(0), C1: int64(0), C2: int64(0), List_start: int64(0), List_count: int64(0), Span: span}
}

func mk_binding(name string, is_mut bool, span Span) Expr {
	return Expr{Kind: ExprKindEkBinding{}, S1: name, S2: "", B1: is_mut, B2: false, C0: int64(0), C1: int64(0), C2: int64(0), List_start: int64(0), List_count: int64(0), Span: span}
}

func mk_return(has_value bool, span Span) Expr {
	return Expr{Kind: ExprKindEkReturn{}, S1: "", S2: "", B1: has_value, B2: false, C0: int64(0), C1: int64(0), C2: int64(0), List_start: int64(0), List_count: int64(0), Span: span}
}

func mk_break(span Span) Expr {
	return Expr{Kind: ExprKindEkBreak{}, S1: "", S2: "", B1: false, B2: false, C0: int64(0), C1: int64(0), C2: int64(0), List_start: int64(0), List_count: int64(0), Span: span}
}

func mk_continue(span Span) Expr {
	return Expr{Kind: ExprKindEkContinue{}, S1: "", S2: "", B1: false, B2: false, C0: int64(0), C1: int64(0), C2: int64(0), List_start: int64(0), List_count: int64(0), Span: span}
}

func mk_for(pattern_name string, span Span) Expr {
	return Expr{Kind: ExprKindEkFor{}, S1: pattern_name, S2: "", B1: false, B2: false, C0: int64(0), C1: int64(0), C2: int64(0), List_start: int64(0), List_count: int64(0), Span: span}
}

func mk_while(span Span) Expr {
	return Expr{Kind: ExprKindEkWhile{}, S1: "", S2: "", B1: false, B2: false, C0: int64(0), C1: int64(0), C2: int64(0), List_start: int64(0), List_count: int64(0), Span: span}
}

func mk_loop(span Span) Expr {
	return Expr{Kind: ExprKindEkLoop{}, S1: "", S2: "", B1: false, B2: false, C0: int64(0), C1: int64(0), C2: int64(0), List_start: int64(0), List_count: int64(0), Span: span}
}

func mk_closure(params_str string, span Span) Expr {
	return Expr{Kind: ExprKindEkClosure{}, S1: params_str, S2: "", B1: false, B2: false, C0: int64(0), C1: int64(0), C2: int64(0), List_start: int64(0), List_count: int64(0), Span: span}
}

func mk_propagate(span Span) Expr {
	return Expr{Kind: ExprKindEkPropagate{}, S1: "", S2: "", B1: false, B2: false, C0: int64(0), C1: int64(0), C2: int64(0), List_start: int64(0), List_count: int64(0), Span: span}
}

func mk_assert_ok(span Span) Expr {
	return Expr{Kind: ExprKindEkAssertOk{}, S1: "", S2: "", B1: false, B2: false, C0: int64(0), C1: int64(0), C2: int64(0), List_start: int64(0), List_count: int64(0), Span: span}
}

func mk_pipeline(span Span) Expr {
	return Expr{Kind: ExprKindEkPipeline{}, S1: "", S2: "", B1: false, B2: false, C0: int64(0), C1: int64(0), C2: int64(0), List_start: int64(0), List_count: int64(0), Span: span}
}

func mk_range(inclusive bool, span Span) Expr {
	return Expr{Kind: ExprKindEkRange{}, S1: "", S2: "", B1: inclusive, B2: false, C0: int64(0), C1: int64(0), C2: int64(0), List_start: int64(0), List_count: int64(0), Span: span}
}

func mk_deferred(span Span) Expr {
	return Expr{Kind: ExprKindEkDefer{}, S1: "", S2: "", B1: false, B2: false, C0: int64(0), C1: int64(0), C2: int64(0), List_start: int64(0), List_count: int64(0), Span: span}
}

func mk_with(name string, span Span) Expr {
	return Expr{Kind: ExprKindEkWith{}, S1: name, S2: "", B1: false, B2: false, C0: int64(0), C1: int64(0), C2: int64(0), List_start: int64(0), List_count: int64(0), Span: span}
}

func mk_tuple(count int64, span Span) Expr {
	return Expr{Kind: ExprKindEkTuple{}, S1: "", S2: "", B1: false, B2: false, C0: int64(0), C1: int64(0), C2: int64(0), List_start: int64(0), List_count: count, Span: span}
}

func mk_catch(span Span) Expr {
	return Expr{Kind: ExprKindEkCatch{}, S1: "", S2: "", B1: false, B2: false, C0: int64(0), C1: int64(0), C2: int64(0), List_start: int64(0), List_count: int64(0), Span: span}
}

func mk_fn_decl(name string, params string, param_count int64, has_body bool, span Span) Expr {
	return Expr{Kind: ExprKindEkFnDecl{}, S1: name, S2: params, B1: has_body, B2: false, C0: int64(0), C1: int64(0), C2: int64(0), List_start: int64(0), List_count: param_count, Span: span}
}

func mk_type_decl(name string, is_sum bool, count int64, span Span) Expr {
	return Expr{Kind: ExprKindEkTypeDecl{}, S1: name, S2: "", B1: is_sum, B2: false, C0: int64(0), C1: int64(0), C2: int64(0), List_start: int64(0), List_count: count, Span: span}
}

func mk_enum_decl(name string, span Span) Expr {
	return Expr{Kind: ExprKindEkEnumDecl{}, S1: name, S2: "", B1: false, B2: false, C0: int64(0), C1: int64(0), C2: int64(0), List_start: int64(0), List_count: int64(0), Span: span}
}

func mk_trait_decl(name string, method_count int64, span Span) Expr {
	return Expr{Kind: ExprKindEkTraitDecl{}, S1: name, S2: "", B1: false, B2: false, C0: int64(0), C1: int64(0), C2: int64(0), List_start: int64(0), List_count: method_count, Span: span}
}

func mk_impl_decl(type_name string, trait_name string, method_count int64, span Span) Expr {
	return Expr{Kind: ExprKindEkImplDecl{}, S1: type_name, S2: trait_name, B1: false, B2: false, C0: int64(0), C1: int64(0), C2: int64(0), List_start: int64(0), List_count: method_count, Span: span}
}

func mk_const_decl(name string, span Span) Expr {
	return Expr{Kind: ExprKindEkConstDecl{}, S1: name, S2: "", B1: false, B2: false, C0: int64(0), C1: int64(0), C2: int64(0), List_start: int64(0), List_count: int64(0), Span: span}
}

func mk_use_decl(module string, span Span) Expr {
	return Expr{Kind: ExprKindEkUseDecl{}, S1: module, S2: "", B1: false, B2: false, C0: int64(0), C1: int64(0), C2: int64(0), List_start: int64(0), List_count: int64(0), Span: span}
}

func mk_param(name string, type_tok_pos int64, is_self bool, span Span) Expr {
	return Expr{Kind: ExprKindEkParam{}, S1: name, S2: "", B1: is_self, B2: false, C0: type_tok_pos, C1: int64(0), C2: int64(0), List_start: int64(0), List_count: int64(0), Span: span}
}

func mk_field(name string, type_tok_pos int64, span Span) Expr {
	return Expr{Kind: ExprKindEkField{}, S1: name, S2: "", B1: false, B2: false, C0: type_tok_pos, C1: int64(0), C2: int64(0), List_start: int64(0), List_count: int64(0), Span: span}
}

func mk_variant(name string, is_tuple bool, span Span) Expr {
	return Expr{Kind: ExprKindEkVariant{}, S1: name, S2: "", B1: is_tuple, B2: false, C0: int64(0), C1: int64(0), C2: int64(0), List_start: int64(0), List_count: int64(0), Span: span}
}

func decl_kind_name(k DeclKind) string {
	return func() interface{} {
		switch _tmp2 := (k).(type) {
		case DeclKindDkFn:
			return "fn"
		case DeclKindDkType:
			return "type"
		case DeclKindDkEnum:
			return "enum"
		case DeclKindDkTrait:
			return "trait"
		case DeclKindDkImpl:
			return "impl"
		case DeclKindDkConst:
			return "const"
		case DeclKindDkEntry:
			return "entry"
		case DeclKindDkTest:
			return "test"
		case DeclKindDkUse:
			return "use"
		case DeclKindDkNone:
			return "none"
		default:
			_ = _tmp2
		}
		return nil
	}().(string)
}

func _sentinel_decl() DeclInfo {
	return DeclInfo{Kind: DeclKindDkNone{}, Name: "", Token_start: int64(0), Body_start: int64(0), Body_end: int64(0), Is_pub: false, Node_idx: int64(0)}
}

func new_decl_index() DeclIndex {
	return DeclIndex{Decls: []DeclInfo{_sentinel_decl()}}
}

func decl_index_add(idx DeclIndex, d DeclInfo) DeclIndex {
	return DeclIndex{Decls: append(idx.Decls, d)}
}

func decl_index_len(idx DeclIndex) int64 {
	return int64(len(idx.Decls))
}

func decl_index_get(idx DeclIndex, i int64) DeclInfo {
	if ((i >= int64(0)) && (i < int64(len(idx.Decls)))) {
		return idx.Decls[i]
	}
	return _sentinel_decl()
}

func _decl_index_set(idx DeclIndex, i int64, d DeclInfo) DeclIndex {
	decls := idx.Decls
	_ = decls
	dlen := int64(len(decls))
	_ = dlen
	new_decls := []DeclInfo{_sentinel_decl()}
	_ = new_decls
	j := int64(1)
	_ = j
	for (j < dlen) {
		if (j == i) {
			new_decls = append(new_decls, d)
		} else {
			new_decls = append(new_decls, decls[j])
		}
		j = (j + int64(1))
	}
	return DeclIndex{Decls: new_decls}
}

