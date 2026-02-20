package token

// Package token implements types and constants to support tokenizing
// the input source before passing the stream of tokens on to the parser.

const (
	// ILLEGAL represents an illegal token
	ILLEGAL = "ILLEGAL"

	// EOF end of file
	EOF = "EOF"

	//
	// Identifiers + literals
	//

	// IDENT an identifier, e.g: add, foobar, x, y, ...
	IDENT = "IDENT"
	// INT an integer, e.g: 1234
	INT = "INT"

	//
	// Operators
	//

	// ASSIGN the assignment operator
	ASSIGN = "="
	// PLUS the addition operator
	PLUS = "+"
	// MINUS the substraction operator
	MINUS = "-"
	// MULTIPLY the multiplication operator
	MULTIPLY = "*"
	// DIVIDE the division operator
	DIVIDE = "/"

	//
	// Logical operators
	//

	// NOT the not operator
	NOT = "!"

	//
	// Delimiters
	//

	// COMMA a comma
	COMMA = ","
	// SEMICOLON a semi-colon
	SEMICOLON = ";"

	// LPAREN a left paranthesis
	LPAREN = "("
	// RPAREN a right parenthesis
	RPAREN = ")"
	// LBRACE a left brace
	LBRACE = "{"
	// RBRACE a right brace
	RBRACE = "}"

	//
	// Comparision operators
	//

	// LT the less than comparision operator
	LT = "<"
	// GT the greater than comparision operator
	GT = ">"

	// EQ the equality operator
	EQ = "=="
	// NEQ the inequality operator
	NEQ = "!="

	//
	// Keywords
	//

	// FUNCTION the `fn` keyword (function)
	FUNCTION = "FUNCTION"
	// LET the `let` keyword (let)
	LET = "LET"
	// TRUE the `true` keyword (true)
	TRUE = "TRUE"
	// FALSE the `false` keyword (false)
	FALSE = "FALSE"
	// IF the `if` keyword (if)
	IF = "IF"
	// ELSE the `else` keyword (else)
	ELSE = "ELSE"
	// RETURN the `return` keyword (return)
	RETURN = "RETURN"
)

var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}
