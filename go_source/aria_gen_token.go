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

type TokenKind interface {
	isTokenKind()
}

type TokenKindTkIdent struct{}
func (TokenKindTkIdent) isTokenKind() {}

type TokenKindTkIntLit struct{}
func (TokenKindTkIntLit) isTokenKind() {}

type TokenKindTkFloatLit struct{}
func (TokenKindTkFloatLit) isTokenKind() {}

type TokenKindTkStringLit struct{}
func (TokenKindTkStringLit) isTokenKind() {}

type TokenKindTkCharLit struct{}
func (TokenKindTkCharLit) isTokenKind() {}

type TokenKindTkBoolTrue struct{}
func (TokenKindTkBoolTrue) isTokenKind() {}

type TokenKindTkBoolFalse struct{}
func (TokenKindTkBoolFalse) isTokenKind() {}

type TokenKindTkRawStringLit struct{}
func (TokenKindTkRawStringLit) isTokenKind() {}

type TokenKindTkTripleStringLit struct{}
func (TokenKindTkTripleStringLit) isTokenKind() {}

type TokenKindTkDurationLit struct{}
func (TokenKindTkDurationLit) isTokenKind() {}

type TokenKindTkSizeLit struct{}
func (TokenKindTkSizeLit) isTokenKind() {}

type TokenKindTkStringStart struct{}
func (TokenKindTkStringStart) isTokenKind() {}

type TokenKindTkStringMiddle struct{}
func (TokenKindTkStringMiddle) isTokenKind() {}

type TokenKindTkStringEnd struct{}
func (TokenKindTkStringEnd) isTokenKind() {}

type TokenKindTkMod struct{}
func (TokenKindTkMod) isTokenKind() {}

type TokenKindTkUse struct{}
func (TokenKindTkUse) isTokenKind() {}

type TokenKindTkPub struct{}
func (TokenKindTkPub) isTokenKind() {}

type TokenKindTkFn struct{}
func (TokenKindTkFn) isTokenKind() {}

type TokenKindTkType struct{}
func (TokenKindTkType) isTokenKind() {}

type TokenKindTkEnum struct{}
func (TokenKindTkEnum) isTokenKind() {}

type TokenKindTkTrait struct{}
func (TokenKindTkTrait) isTokenKind() {}

type TokenKindTkImpl struct{}
func (TokenKindTkImpl) isTokenKind() {}

type TokenKindTkStruct struct{}
func (TokenKindTkStruct) isTokenKind() {}

type TokenKindTkEntry struct{}
func (TokenKindTkEntry) isTokenKind() {}

type TokenKindTkIf struct{}
func (TokenKindTkIf) isTokenKind() {}

type TokenKindTkElse struct{}
func (TokenKindTkElse) isTokenKind() {}

type TokenKindTkMatch struct{}
func (TokenKindTkMatch) isTokenKind() {}

type TokenKindTkFor struct{}
func (TokenKindTkFor) isTokenKind() {}

type TokenKindTkIn struct{}
func (TokenKindTkIn) isTokenKind() {}

type TokenKindTkWhile struct{}
func (TokenKindTkWhile) isTokenKind() {}

type TokenKindTkLoop struct{}
func (TokenKindTkLoop) isTokenKind() {}

type TokenKindTkBreak struct{}
func (TokenKindTkBreak) isTokenKind() {}

type TokenKindTkContinue struct{}
func (TokenKindTkContinue) isTokenKind() {}

type TokenKindTkReturn struct{}
func (TokenKindTkReturn) isTokenKind() {}

type TokenKindTkMut struct{}
func (TokenKindTkMut) isTokenKind() {}

type TokenKindTkConst struct{}
func (TokenKindTkConst) isTokenKind() {}

type TokenKindTkDefer struct{}
func (TokenKindTkDefer) isTokenKind() {}

type TokenKindTkWith struct{}
func (TokenKindTkWith) isTokenKind() {}

type TokenKindTkCatch struct{}
func (TokenKindTkCatch) isTokenKind() {}

type TokenKindTkTest struct{}
func (TokenKindTkTest) isTokenKind() {}

type TokenKindTkAssert struct{}
func (TokenKindTkAssert) isTokenKind() {}

type TokenKindTkDerives struct{}
func (TokenKindTkDerives) isTokenKind() {}

type TokenKindTkWhere struct{}
func (TokenKindTkWhere) isTokenKind() {}

type TokenKindTkAs struct{}
func (TokenKindTkAs) isTokenKind() {}

type TokenKindTkIs struct{}
func (TokenKindTkIs) isTokenKind() {}

type TokenKindTkSelf struct{}
func (TokenKindTkSelf) isTokenKind() {}

type TokenKindTkSelfType struct{}
func (TokenKindTkSelfType) isTokenKind() {}

type TokenKindTkSpawn struct{}
func (TokenKindTkSpawn) isTokenKind() {}

type TokenKindTkScope struct{}
func (TokenKindTkScope) isTokenKind() {}

type TokenKindTkSelect struct{}
func (TokenKindTkSelect) isTokenKind() {}

type TokenKindTkBench struct{}
func (TokenKindTkBench) isTokenKind() {}

type TokenKindTkBefore struct{}
func (TokenKindTkBefore) isTokenKind() {}

type TokenKindTkAfter struct{}
func (TokenKindTkAfter) isTokenKind() {}

type TokenKindTkYield struct{}
func (TokenKindTkYield) isTokenKind() {}

type TokenKindTkFrom struct{}
func (TokenKindTkFrom) isTokenKind() {}

type TokenKindTkMock struct{}
func (TokenKindTkMock) isTokenKind() {}

type TokenKindTkExtern struct{}
func (TokenKindTkExtern) isTokenKind() {}

type TokenKindTkPlus struct{}
func (TokenKindTkPlus) isTokenKind() {}

type TokenKindTkMinus struct{}
func (TokenKindTkMinus) isTokenKind() {}

type TokenKindTkStar struct{}
func (TokenKindTkStar) isTokenKind() {}

type TokenKindTkSlash struct{}
func (TokenKindTkSlash) isTokenKind() {}

type TokenKindTkPercent struct{}
func (TokenKindTkPercent) isTokenKind() {}

type TokenKindTkAmpersand struct{}
func (TokenKindTkAmpersand) isTokenKind() {}

type TokenKindTkPipe struct{}
func (TokenKindTkPipe) isTokenKind() {}

type TokenKindTkCaret struct{}
func (TokenKindTkCaret) isTokenKind() {}

type TokenKindTkTilde struct{}
func (TokenKindTkTilde) isTokenKind() {}

type TokenKindTkBang struct{}
func (TokenKindTkBang) isTokenKind() {}

type TokenKindTkLShift struct{}
func (TokenKindTkLShift) isTokenKind() {}

type TokenKindTkRShift struct{}
func (TokenKindTkRShift) isTokenKind() {}

type TokenKindTkEqEq struct{}
func (TokenKindTkEqEq) isTokenKind() {}

type TokenKindTkBangEq struct{}
func (TokenKindTkBangEq) isTokenKind() {}

type TokenKindTkLt struct{}
func (TokenKindTkLt) isTokenKind() {}

type TokenKindTkGt struct{}
func (TokenKindTkGt) isTokenKind() {}

type TokenKindTkLtEq struct{}
func (TokenKindTkLtEq) isTokenKind() {}

type TokenKindTkGtEq struct{}
func (TokenKindTkGtEq) isTokenKind() {}

type TokenKindTkAmpAmp struct{}
func (TokenKindTkAmpAmp) isTokenKind() {}

type TokenKindTkPipePipe struct{}
func (TokenKindTkPipePipe) isTokenKind() {}

type TokenKindTkDotDot struct{}
func (TokenKindTkDotDot) isTokenKind() {}

type TokenKindTkDotDotEq struct{}
func (TokenKindTkDotDotEq) isTokenKind() {}

type TokenKindTkPipeGt struct{}
func (TokenKindTkPipeGt) isTokenKind() {}

type TokenKindTkQuestion struct{}
func (TokenKindTkQuestion) isTokenKind() {}

type TokenKindTkQuestionDot struct{}
func (TokenKindTkQuestionDot) isTokenKind() {}

type TokenKindTkQuestionQuestion struct{}
func (TokenKindTkQuestionQuestion) isTokenKind() {}

type TokenKindTkArrow struct{}
func (TokenKindTkArrow) isTokenKind() {}

type TokenKindTkFatArrow struct{}
func (TokenKindTkFatArrow) isTokenKind() {}

type TokenKindTkLParen struct{}
func (TokenKindTkLParen) isTokenKind() {}

type TokenKindTkRParen struct{}
func (TokenKindTkRParen) isTokenKind() {}

type TokenKindTkLBracket struct{}
func (TokenKindTkLBracket) isTokenKind() {}

type TokenKindTkRBracket struct{}
func (TokenKindTkRBracket) isTokenKind() {}

type TokenKindTkLBrace struct{}
func (TokenKindTkLBrace) isTokenKind() {}

type TokenKindTkRBrace struct{}
func (TokenKindTkRBrace) isTokenKind() {}

type TokenKindTkComma struct{}
func (TokenKindTkComma) isTokenKind() {}

type TokenKindTkColon struct{}
func (TokenKindTkColon) isTokenKind() {}

type TokenKindTkDot struct{}
func (TokenKindTkDot) isTokenKind() {}

type TokenKindTkDotLBrace struct{}
func (TokenKindTkDotLBrace) isTokenKind() {}

type TokenKindTkAt struct{}
func (TokenKindTkAt) isTokenKind() {}

type TokenKindTkEq struct{}
func (TokenKindTkEq) isTokenKind() {}

type TokenKindTkColonEq struct{}
func (TokenKindTkColonEq) isTokenKind() {}

type TokenKindTkNewline struct{}
func (TokenKindTkNewline) isTokenKind() {}

type TokenKindTkEof struct{}
func (TokenKindTkEof) isTokenKind() {}

type TokenKindTkIllegal struct{}
func (TokenKindTkIllegal) isTokenKind() {}

type Token struct {
	Kind TokenKind
	Text string
	Line int64
	Col int64
	Offset int64
}

func new_token(kind TokenKind, text string, line int64, col int64, offset int64) Token {
	return Token{Kind: kind, Text: text, Line: line, Col: col, Offset: offset}
}

func token_name(kind TokenKind) string {
	return func() interface{} {
		switch _tmp1 := (kind).(type) {
		case TokenKindTkIdent:
			return "IDENT"
		case TokenKindTkIntLit:
			return "INT"
		case TokenKindTkFloatLit:
			return "FLOAT"
		case TokenKindTkStringLit:
			return "STRING"
		case TokenKindTkCharLit:
			return "CHAR"
		case TokenKindTkBoolTrue:
			return "true"
		case TokenKindTkBoolFalse:
			return "false"
		case TokenKindTkRawStringLit:
			return "RAW_STRING"
		case TokenKindTkTripleStringLit:
			return "TRIPLE_STRING"
		case TokenKindTkDurationLit:
			return "DURATION"
		case TokenKindTkSizeLit:
			return "SIZE"
		case TokenKindTkStringStart:
			return "STRING_START"
		case TokenKindTkStringMiddle:
			return "STRING_MIDDLE"
		case TokenKindTkStringEnd:
			return "STRING_END"
		case TokenKindTkMod:
			return "mod"
		case TokenKindTkUse:
			return "use"
		case TokenKindTkPub:
			return "pub"
		case TokenKindTkFn:
			return "fn"
		case TokenKindTkType:
			return "type"
		case TokenKindTkEnum:
			return "enum"
		case TokenKindTkTrait:
			return "trait"
		case TokenKindTkImpl:
			return "impl"
		case TokenKindTkStruct:
			return "struct"
		case TokenKindTkEntry:
			return "entry"
		case TokenKindTkIf:
			return "if"
		case TokenKindTkElse:
			return "else"
		case TokenKindTkMatch:
			return "match"
		case TokenKindTkFor:
			return "for"
		case TokenKindTkIn:
			return "in"
		case TokenKindTkWhile:
			return "while"
		case TokenKindTkLoop:
			return "loop"
		case TokenKindTkBreak:
			return "break"
		case TokenKindTkContinue:
			return "continue"
		case TokenKindTkReturn:
			return "return"
		case TokenKindTkMut:
			return "mut"
		case TokenKindTkConst:
			return "const"
		case TokenKindTkDefer:
			return "defer"
		case TokenKindTkWith:
			return "with"
		case TokenKindTkCatch:
			return "catch"
		case TokenKindTkTest:
			return "test"
		case TokenKindTkAssert:
			return "assert"
		case TokenKindTkDerives:
			return "derives"
		case TokenKindTkWhere:
			return "where"
		case TokenKindTkAs:
			return "as"
		case TokenKindTkIs:
			return "is"
		case TokenKindTkSelf:
			return "self"
		case TokenKindTkSelfType:
			return "Self"
		case TokenKindTkSpawn:
			return "spawn"
		case TokenKindTkScope:
			return "scope"
		case TokenKindTkSelect:
			return "select"
		case TokenKindTkBench:
			return "bench"
		case TokenKindTkBefore:
			return "before"
		case TokenKindTkAfter:
			return "after"
		case TokenKindTkYield:
			return "yield"
		case TokenKindTkFrom:
			return "from"
		case TokenKindTkMock:
			return "mock"
		case TokenKindTkExtern:
			return "extern"
		case TokenKindTkPlus:
			return "+"
		case TokenKindTkMinus:
			return "-"
		case TokenKindTkStar:
			return "*"
		case TokenKindTkSlash:
			return "/"
		case TokenKindTkPercent:
			return "%"
		case TokenKindTkAmpersand:
			return "&"
		case TokenKindTkPipe:
			return "|"
		case TokenKindTkCaret:
			return "^"
		case TokenKindTkTilde:
			return "~"
		case TokenKindTkBang:
			return "!"
		case TokenKindTkLShift:
			return "<<"
		case TokenKindTkRShift:
			return ">>"
		case TokenKindTkEqEq:
			return "=="
		case TokenKindTkBangEq:
			return "!="
		case TokenKindTkLt:
			return "<"
		case TokenKindTkGt:
			return ">"
		case TokenKindTkLtEq:
			return "<="
		case TokenKindTkGtEq:
			return ">="
		case TokenKindTkAmpAmp:
			return "&&"
		case TokenKindTkPipePipe:
			return "||"
		case TokenKindTkDotDot:
			return ".."
		case TokenKindTkDotDotEq:
			return "..="
		case TokenKindTkPipeGt:
			return "|>"
		case TokenKindTkQuestion:
			return "?"
		case TokenKindTkQuestionDot:
			return "?."
		case TokenKindTkQuestionQuestion:
			return "??"
		case TokenKindTkArrow:
			return "->"
		case TokenKindTkFatArrow:
			return "=>"
		case TokenKindTkLParen:
			return "("
		case TokenKindTkRParen:
			return ")"
		case TokenKindTkLBracket:
			return "["
		case TokenKindTkRBracket:
			return "]"
		case TokenKindTkLBrace:
			return "{"
		case TokenKindTkRBrace:
			return "}"
		case TokenKindTkComma:
			return ","
		case TokenKindTkColon:
			return ":"
		case TokenKindTkDot:
			return "."
		case TokenKindTkDotLBrace:
			return ".{"
		case TokenKindTkAt:
			return "@"
		case TokenKindTkEq:
			return "="
		case TokenKindTkColonEq:
			return ":="
		case TokenKindTkNewline:
			return "NEWLINE"
		case TokenKindTkEof:
			return "EOF"
		case TokenKindTkIllegal:
			return "ILLEGAL"
		default:
			_ = _tmp1
		}
		return nil
	}().(string)
}

func lookup_keyword(text string) TokenKind {
	if (text == "mod") {
		return TokenKindTkMod{}
	}
	if (text == "use") {
		return TokenKindTkUse{}
	}
	if (text == "pub") {
		return TokenKindTkPub{}
	}
	if (text == "fn") {
		return TokenKindTkFn{}
	}
	if (text == "type") {
		return TokenKindTkType{}
	}
	if (text == "enum") {
		return TokenKindTkEnum{}
	}
	if (text == "trait") {
		return TokenKindTkTrait{}
	}
	if (text == "impl") {
		return TokenKindTkImpl{}
	}
	if (text == "struct") {
		return TokenKindTkStruct{}
	}
	if (text == "entry") {
		return TokenKindTkEntry{}
	}
	if (text == "if") {
		return TokenKindTkIf{}
	}
	if (text == "else") {
		return TokenKindTkElse{}
	}
	if (text == "match") {
		return TokenKindTkMatch{}
	}
	if (text == "for") {
		return TokenKindTkFor{}
	}
	if (text == "in") {
		return TokenKindTkIn{}
	}
	if (text == "while") {
		return TokenKindTkWhile{}
	}
	if (text == "loop") {
		return TokenKindTkLoop{}
	}
	if (text == "break") {
		return TokenKindTkBreak{}
	}
	if (text == "continue") {
		return TokenKindTkContinue{}
	}
	if (text == "return") {
		return TokenKindTkReturn{}
	}
	if (text == "mut") {
		return TokenKindTkMut{}
	}
	if (text == "const") {
		return TokenKindTkConst{}
	}
	if (text == "defer") {
		return TokenKindTkDefer{}
	}
	if (text == "with") {
		return TokenKindTkWith{}
	}
	if (text == "catch") {
		return TokenKindTkCatch{}
	}
	if (text == "test") {
		return TokenKindTkTest{}
	}
	if (text == "assert") {
		return TokenKindTkAssert{}
	}
	if (text == "derives") {
		return TokenKindTkDerives{}
	}
	if (text == "where") {
		return TokenKindTkWhere{}
	}
	if (text == "as") {
		return TokenKindTkAs{}
	}
	if (text == "is") {
		return TokenKindTkIs{}
	}
	if (text == "self") {
		return TokenKindTkSelf{}
	}
	if (text == "Self") {
		return TokenKindTkSelfType{}
	}
	if (text == "spawn") {
		return TokenKindTkSpawn{}
	}
	if (text == "bench") {
		return TokenKindTkBench{}
	}
	if (text == "before") {
		return TokenKindTkBefore{}
	}
	if (text == "after") {
		return TokenKindTkAfter{}
	}
	if (text == "yield") {
		return TokenKindTkYield{}
	}
	if (text == "from") {
		return TokenKindTkFrom{}
	}
	if (text == "mock") {
		return TokenKindTkMock{}
	}
	if (text == "extern") {
		return TokenKindTkExtern{}
	}
	if (text == "scope") {
		return TokenKindTkScope{}
	}
	if (text == "select") {
		return TokenKindTkSelect{}
	}
	if (text == "true") {
		return TokenKindTkBoolTrue{}
	}
	if (text == "false") {
		return TokenKindTkBoolFalse{}
	}
	return TokenKindTkIdent{}
}

func is_keyword(kind TokenKind) bool {
	name := token_name(kind)
	_ = name
	if (int64(len(name)) == int64(0)) {
		return false
	}
	ch := string(name[int64(0)])
	_ = ch
	return ((((((((((((((((((((((((((ch == "a") || (ch == "b")) || (ch == "c")) || (ch == "d")) || (ch == "e")) || (ch == "f")) || (ch == "g")) || (ch == "h")) || (ch == "i")) || (ch == "j")) || (ch == "k")) || (ch == "l")) || (ch == "m")) || (ch == "n")) || (ch == "o")) || (ch == "p")) || (ch == "q")) || (ch == "r")) || (ch == "s")) || (ch == "t")) || (ch == "u")) || (ch == "v")) || (ch == "w")) || (ch == "x")) || (ch == "y")) || (ch == "z"))
}

func terminates_statement(kind TokenKind) bool {
	name := token_name(kind)
	_ = name
	if (name == "IDENT") {
		return true
	}
	if (name == "INT") {
		return true
	}
	if (name == "FLOAT") {
		return true
	}
	if (name == "STRING") {
		return true
	}
	if (name == "RAW_STRING") {
		return true
	}
	if (name == "TRIPLE_STRING") {
		return true
	}
	if (name == "DURATION") {
		return true
	}
	if (name == "SIZE") {
		return true
	}
	if (name == "true") {
		return true
	}
	if (name == "false") {
		return true
	}
	if (name == "STRING_END") {
		return true
	}
	if (name == "break") {
		return true
	}
	if (name == "continue") {
		return true
	}
	if (name == "return") {
		return true
	}
	if (name == ")") {
		return true
	}
	if (name == "]") {
		return true
	}
	if kind_eq(kind, TokenKindTkRBrace{}) {
		return true
	}
	if (name == "?") {
		return true
	}
	if (name == "!") {
		return true
	}
	if (name == "self") {
		return true
	}
	if (name == "Self") {
		return true
	}
	return false
}

func kind_eq(a TokenKind, b TokenKind) bool {
	return (token_name(a) == token_name(b))
}

