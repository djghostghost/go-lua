package lexer

const (
	TokenEoF         = iota // end of file
	TokenVararg             // ...
	TokenSepSemi            // ,
	TokenSepComma           // ;
	TokenSepDot             // .
	TokenSepColon           // :
	TokenSepLabel           // ::
	TokenSepLParen          // (
	TokenSepRParen          // )
	TokenSepLBracket        // [
	TokenSepRBracket        // ]
	TokenSepLCurly          // {
	TokenSepRCurly          // }
	TokenOpAssign           // =
	TokenOpMinus            // -
	TokenOpWave             // ~ (bnot or bxor)
	TokenOpAdd              //+
	TokenOpMul              //*
	TokenOpDiv              // /
	TokenOpIDiv             // //
	TokenOpPow              // ^
	TokenOpMod              // %
	TokenOpBAnd             // and
	TokenOpBOr              // |
	TokenOpShr              // >>
	TokenOpShl              // <<
	TokenOpConcat           // ..
	TokenOpLt               // <
	TokenOpLe               // <=
	TokenOpGt               // >
	TokenOpGe               // >=
	TokenOpEq               // ==
	TokenOpNe               // ~=
	TokenOpLen              // #
	TokenOpAnd              // and
	TokenOpOr               // or
	TokenOpNot              // not
	TokenKwBreak            // break
	TokenKwDo               // do
	TokenKwElse             // else
	TokenKwElseIf           // elseif
	TokenKwEnd              // end
	TokenKwFalse            // false
	TokenKwFor              // for
	TokenKwFunction         // function
	TokenKwGoto             // goto
	TokenKwIf               // if
	TokenKwIn               // in
	TokenKwLocal            // local
	TokenKwNil              // nil
	TokenKwRepeat           // repeat
	TokenKwReturn           // return
	TokenKwThen             // then
	TokenKwTrue             // true
	TokenKwUtil             // util
	TokenKwWhile            // while
	TokenIdentifier         // identifier
	TokenNumber             // number literal
	TokenString             // string literal
	TokenOpUnm       = TokenOpMinus
	TokenOpSub       = TokenOpMinus
	TokenOpBNot      = TokenOpWave
	TokenOpBXor      = TokenOpWave
)

var KeyWordsMap = map[string]int{
	"and":      TokenOpAnd,
	"break":    TokenKwBreak,
	"do":       TokenKwDo,
	"else":     TokenKwElse,
	"elseif":   TokenKwElseIf,
	"end":      TokenKwEnd,
	"false":    TokenKwFalse,
	"for":      TokenKwFor,
	"function": TokenKwFunction,
	"goto":     TokenKwGoto,
	"if":       TokenKwIf,
	"in":       TokenKwIn,
	"local":    TokenKwLocal,
	"nil":      TokenKwNil,
	"not":      TokenOpNot,
	"or":       TokenOpOr,
	"repeat":   TokenKwRepeat,
	"return":   TokenKwReturn,
	"then":     TokenKwThen,
	"true":     TokenKwTrue,
	"until":    TokenKwUtil,
	"while":    TokenKwWhile,
}
