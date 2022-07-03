package lexer

import (
	"bytes"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

const EOF = "EOF"

var reOpeningLongBracket = regexp.MustCompile(`^\[=*\[`)

var reIdentifier = regexp.MustCompile(`^[_\d\w]+`)
var reNumber = regexp.MustCompile(`^0[xX][0-9a-fA-F]*(\.[0-9a-fA-F]*)?([pP][+\-]?[0-9]+)?|^[0-9]*(\.[0-9]*)?([eE][+\-]?[0-9]+)?`)
var reShortStr = regexp.MustCompile(`(?s)(^'(\\\\|\\'|\\\n|\\z\s*|[^'\n])*')|(^"(\\\\|\\"|\\\n|\\z\s*|[^"\n])*")`)

var reDecEscapeSeq = regexp.MustCompile(`^\\[0-9]{1,3}`)
var reHexEscapeSeq = regexp.MustCompile(`^\\x[0-9a-fA-F]{2}`)
var reUnicodeEscapeSeq = regexp.MustCompile(`^\\u\{[0-9a-fA-F]+\}`)

type Lexer struct {
	chunk     string
	chunkName string
	line      int

	nextToken     string
	nextTokenKind int
	nextTokenLine int
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

	c := l.chunk[0]
	if c == '.' || isDigit(c) {
		token := l.scanNumber()
		return l.line, TokenNumber, token
	}

	if c == '_' || isLatter(c) {
		token := l.scanIdentifier()
		if kind, found := KeyWordsMap[token]; found {
			return l.line, kind, token
		} else {
			return l.line, TokenIdentifier, token
		}
	}

	if l.nextTokenLine > 0 {
		line = l.nextTokenLine
		kind = l.nextTokenKind
		token = l.nextToken
		l.line = l.nextTokenLine
		l.nextTokenLine = 0
		return
	}
	l.error("unexpected symbol near %q", c)
	return
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

var reNewLine = regexp.MustCompile("\r\n|\n\r|\n\r")

func (l *Lexer) scanLongString() string {
	openingLongBracket := reOpeningLongBracket.FindString(l.chunk)
	if openingLongBracket == "" {
		l.error("invalid long string delimiter near '%s", l.chunk[0:2])
	}

	closingLongBracket := strings.Replace(openingLongBracket, "[", "]", -1)
	closingBracketIndex := strings.Index(l.chunk, closingLongBracket)

	if closingBracketIndex < 0 {
		l.error("unfinished long string or comment")
	}

	str := l.chunk[len(openingLongBracket):closingBracketIndex]
	l.next(closingBracketIndex + len(closingLongBracket))

	str = reNewLine.ReplaceAllString(str, "\n")
	l.line += strings.Count(str, "\n")

	i := 0
	for ; i < len(str); i++ {
		if str[i] != '\n' {
			break
		}
	}

	if i != 0 {
		str = str[i:]
	}
	return str
}

func (l *Lexer) scanShortString() string {
	if str := reShortStr.FindString(l.chunk); str != "" {
		l.next(len(str))
		str = str[1 : len(str)-1]
		if strings.Index(str, `n`) >= 0 {
			l.line += len(reNewLine.FindAllString(str, -1))
			str = l.escape(str)
		}
		return str
	}
	l.error("unfinished string")
	return ""
}

func (l *Lexer) scanIdentifier() string {
	return l.scan(reIdentifier)
}

func (l *Lexer) error(f string, a ...interface{}) {
	err := fmt.Sprintf(f, a)
	err = fmt.Sprintf("%s:%d: %s", l.chunkName, l.line, err)
	panic(err)
}

func (l *Lexer) escape(str string) string {
	buf := bytes.Buffer{}
	for len(str) > 0 {

		if str[0] != '\\' {
			buf.WriteByte(str[0])
			str = str[1:]
			continue
		}
		if len(str) == 1 {
			l.error("unfinished string")
		}

		switch str[1] {
		case 'a':
			buf.WriteByte('\a')
			str = str[2:]
			continue
		case 'b':
			buf.WriteByte('\b')
			str = str[2:]
			continue
		case 'f':
			buf.WriteByte('\f')
			str = str[2:]
			continue
		case 'n':
			buf.WriteByte('\n')
			str = str[2:]
			continue
		case '\n':
			buf.WriteByte('\n')
			str = str[2:]
			continue
		case 'r':
			buf.WriteByte('\r')
			str = str[2:]
			continue
		case 't':
			buf.WriteByte('\t')
			str = str[2:]
			continue
		case 'v':
			buf.WriteByte('\v')
			str = str[2:]
			continue
		case '"':
			buf.WriteByte('"')
			str = str[2:]
			continue
		case '\'':
			buf.WriteByte('\'')
			str = str[2:]
			continue
		case '\\':
			buf.WriteByte('\\')
			str = str[2:]
			continue
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			if found := reDecEscapeSeq.FindString(str); found != "" {
				d, _ := strconv.ParseInt(found[1:], 10, 32)
				if d <= 0xFF {
					buf.WriteByte(byte(d))
					str = str[len(found):]
					continue
				}
				l.error("decimal escape too large near '%s'", found)
			}
		case 'x':
			if found := reHexEscapeSeq.FindString(str); found != "" {
				d, _ := strconv.ParseInt(found[2:], 16, 32)
				buf.WriteByte(byte(d))
				str = str[len(found):]
				continue
			}
		case 'u':
			if found := reUnicodeEscapeSeq.FindString(str); found != "" {
				d, err := strconv.ParseInt(found[3:len(found)-1], 16, 32)
				if err == nil && d < 0x10FFFF {
					buf.WriteRune(rune(d))
					str = str[len(found):]
					continue
				}
				l.error("UTF-8 value too large near '%s'", found)
			}
		case 'z':
			str = str[2:]
			for len(str) > 0 && isWhiteSpace(str[0]) {
				str = str[1:]
			}
			continue
		}

	}
	return buf.String()
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

func (l *Lexer) scan(re *regexp.Regexp) string {
	if token := re.FindString(l.chunk); token != "" {
		l.next(len(token))
		return token
	}
	panic("unreachable")
}

func (l *Lexer) scanNumber() string {
	return l.scan(reNumber)
}

func isLatter(c byte) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z')
}

func (l *Lexer) LookAhead() int {
	if l.nextTokenLine > 0 {
		return l.nextTokenKind
	}

	currentLine := l.line
	line, kind, token := l.NextToken()
	l.line = currentLine
	l.nextToken = token
	l.nextTokenKind = kind
	l.nextTokenLine = line
	return kind
}

func (l *Lexer) NextTokenOfKind(kind int) (line int, token string) {
	line, _kind, token := l.NextToken()
	if kind != _kind {
		l.error("syntax error near '%s'", token)
	}
	return line, token
}

func (l *Lexer) NextIdentifier() (line int, token string) {
	return l.NextTokenOfKind(TokenIdentifier)
}

func (l *Lexer) Line() int {
	return l.line
}
