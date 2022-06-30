package lexer

import (
	"regexp"
	"strings"
)

const EOF = "EOF"

type Lexer struct {
	chunk     string
	chunkName string
	line      int
}

func NewLexer(chunk string, chunkName string) *Lexer {
	return &Lexer{chunk: chunk, chunkName: chunkName, line: 1}
}

func (l *Lexer) NextToken() (line, kind int, token string) {
	l.skipWhiteSpaces()
	if len(l.chunk) == 0 {
		return l.line, TokenEoF, EOF
	}

	switch l.chunk[0] {
	case ';':
		l.next(1)
		return l.line, TokenSepSemi, ";"
	case ',':
		l.next(1)
		return l.line, TokenSepComma, ","
	case '(':
		l.next(1)
		return l.line, TokenSepLParen, "("
	case ')':
		l.next(1)
		return l.line, TokenSepRParen, ")"
	case '{':
		l.next(1)
		return l.line, TokenSepLCurly, "{"
	case '}':
		l.next(1)
		return l.line, TokenSepRCurly, "}"
	case '+':
		l.next(1)
		return l.line, TokenOpAdd, "+"
	case '-':
		l.next(1)
		return l.line, TokenOpMinus, "-"
	case '*':
		l.next(1)
		return l.line, TokenOpMul, "*"
	case '^':
		l.next(1)
		return l.line, TokenOpPow, "^"
	case '%':
		l.next(1)
		return l.line, TokenOpMod, "%"
	case '&':
		l.next(1)
		return l.line, TokenOpBAnd, "&"
	case '|':
		l.next(1)
		return l.line, TokenOpBOr, "|"
	case '#':
		l.next(1)
		return l.line, TokenOpLen, "#"
	case ':':
		if l.test("::") {
			l.next(2)
			return l.line, TokenSepLabel, "::"
		} else {
			l.next(1)
			return l.line, TokenSepColon, ":"
		}
	case '/':
		if l.test("//") {
			l.next(2)
			return l.line, TokenOpIDiv, "//"
		} else {
			l.next(1)
			return l.line, TokenOpDiv, "/"
		}
	case '~':
		if l.test("~=") {
			l.next(2)
			return l.line, TokenOpNe, "~="
		} else {
			l.next(1)
			return l.line, TokenOpWave, "~"
		}
	case '=':
		if l.test("==") {
			l.next(2)
			return l.line, TokenOpEq, "=="
		} else {
			l.next(1)
			return l.line, TokenOpAssign, "="
		}
	case '<':
		if l.test("<<") {
			l.next(2)
			return l.line, TokenOpShl, "<<"
		} else if l.test("<=") {
			l.next(2)
			return l.line, TokenOpLe, "<="
		} else {
			l.next(1)
			return l.line, TokenOpLt, "<"
		}
	case '>':
		if l.test(">>") {
			l.next(2)
			return l.line, TokenOpShr, ">>"
		} else if l.test(">=") {
			l.next(2)
			return l.line, TokenOpGe, ">="
		} else {
			l.next(1)
			return l.line, TokenOpGt, ">"
		}
	case '.':
		if l.test("...") {
			l.next(3)
			return l.line, TokenVararg, "..."
		} else if l.test("..") {
			l.next(2)
			return l.line, TokenOpConcat, ".."
		} else if len(l.chunk) == 1 || !isDigit(l.chunk[1]) {
			l.next(1)
			return l.line, TokenSepDot, "."
		}

	case '[':
		if l.test("[[") || l.test("[=") {
			return l.line, TokenString, l.scanLongString()
		} else {
			l.next(1)
			return l.line, TokenSepLBracket, "["
		}
	case '\'', '"':
		return l.line, TokenString, l.scanShortString()
	}

}

func (l *Lexer) skipWhiteSpaces() {
	for len(l.chunk) > 0 {
		if l.test("--") {
			l.skipComment()
		} else if l.test("\r\n") || l.test("\n\r") {
			l.next(2)
			l.line += 1
		} else if isNewLine(l.chunk[0]) {
			l.next(1)
			l.line += 1
		} else if isWhiteSpace(l.chunk[0]) {
			l.next(1)
		} else {
			break
		}
	}
}

var reOpeningLongBracket = regexp.MustCompile(`^\[=*\[`)

func (l *Lexer) skipComment() {
	l.next(2) // skip --
	if l.test("[") {
		if reOpeningLongBracket.FindString(l.chunk) != "" {
			l.scanLongString()
			return
		}
	}
	// short comment
	for len(l.chunk) > 0 && !isNewLine(l.chunk[0]) {
		l.next(1)
	}
}

func (l *Lexer) test(str string) bool {
	return strings.HasPrefix(l.chunk, str)
}

func (l *Lexer) next(n int) {
	l.chunk = l.chunk[n:]
}

var WhiteSpace = map[byte]bool{
	'\t': true,
	'\n': true,
	'\v': true,
	'\f': true,
	'\r': true,
	' ':  true,
}

func isWhiteSpace(c byte) bool {
	return WhiteSpace[c]
}

func isNewLine(c byte) bool {
	return c == '\r' || c == '\n'
}

func isDigit(c byte) bool {
	return c >= '0' && c <= '9'
}
