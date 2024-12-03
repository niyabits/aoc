package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	content, _ := os.ReadFile("input2.txt")
	input := string(content)

	var tokenStream []Token

	l := New(input)
	for tok := l.NextToken(); tok.Type != EOF; tok = l.NextToken() {
		// fmt.Println(tok)
		tokenStream = append(tokenStream, tok)
	}

	fmt.Println(sumSlice(evalMulExpressions(tokenStream)))
}

func evalMulExpressions(tokens []Token) []int {
	var res []int
	i := 0
	j := 1

	for i < len(tokens) {
		if i+5 < len(tokens) &&
			tokens[i].Type == MULTIPLY &&
			tokens[i+1].Type == LPAREN &&
			tokens[i+2].Type == INT &&
			tokens[i+3].Type == COMMA &&
			tokens[i+4].Type == INT &&
			tokens[i+5].Type == RPAREN {

			x, _ := strconv.Atoi(tokens[i+2].Literal)
			y, _ := strconv.Atoi(tokens[i+4].Literal)
			// fmt.Printf("%v: %v%v%v%v%v%v\n", j, tokens[i].Literal, tokens[i+1].Literal, tokens[i+2].Literal, tokens[i+3].Literal, tokens[i+4].Literal, tokens[i+5].Literal)

			res = append(res, x*y)
			i += 6
			j++
		} else {
			i++
		}
	}

	return res
}

func sumSlice(nums []int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}

	return sum
}

type Lexer struct {
	input        string
	position     int  // current char i
	readPosition int  // after curr char: i + 1
	ch           byte // current char under examination
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) NextToken() Token {
	var tok Token

	switch l.ch {
	case '(':
		tok = newToken(LPAREN, '(')
	case ')':
		tok = newToken(RPAREN, ')')
	case ',':
		tok = newToken(COMMA, ',')
	case 0:
		tok.Literal = ""
		tok.Type = EOF
	default:
		if isMul(l.ch) {
			tok.Literal = l.readIdent()
			tok.Type = LookupIdent(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			num := l.readNum()

			if len(num) <= 3 {
				tok.Type = INT
				tok.Literal = num
			} else {
				tok = newToken(ILLEGAL, l.ch)
			}

			return tok
		} else {
			tok = newToken(ILLEGAL, l.ch)
		}
	}

	l.readChar()
	return tok
}

func (l *Lexer) readIdent() string {
	position := l.position

	for isMul(l.ch) {
		l.readChar()
	}

	ident := l.input[position:l.position]
	if strings.HasSuffix(ident, "mul") {
		return "mul"
	}

	return "INVALID"
}

func (l *Lexer) readNum() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}

	return l.input[position:l.position]
}

func isMul(ch byte) bool {
	return ch == 'm' || ch == 'u' || ch == 'l'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func newToken(tokenType TokenType, ch byte) Token {
	return Token{Type: tokenType, Literal: string(ch)}
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}

	l.position = l.readPosition
	l.readPosition += 1
}
