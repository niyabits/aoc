package main

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	INT    = "INT"
	LPAREN = "LPAREN"
	RPAREN = "RPAREN"

	WHITESPACE = "WHITESPACE"

	// DELIMITER
	COMMA = "COMMA"

	// Keyword
	MULTIPLY = "MULTIPLY"
)

var keywords = map[string]TokenType{
	"mul": MULTIPLY,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}

	return ILLEGAL
}
