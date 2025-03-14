package javascript

import "fmt"

// Define token types using iota
type TokenType int

const (
	TOKEN_ID TokenType = iota
	TOKEN_EQUALS
	TOKEN_LPAREN
	TOKEN_RPAREN
	TOKEN_LBRACE
	TOKEN_RBRACE
	TOKEN_LBRACKET
	TOKEN_RBRACKET
	TOKEN_COLON
	TOKEN_COMMA
	TOKEN_LT
	TOKEN_GT
	TOKEN_ARROW_RIGHT
	TOKEN_INT
	TOKEN_STRING
	TOKEN_STATEMENT
	TOKEN_SEMI
	TOKEN_PLUS
	TOKEN_MINUS
	TOKEN_DIV
	TOKEN_MUL
	TOKEN_EOF
)

// Token structure
type Token struct {
	Type  TokenType
	Value string
}

// Constructor for Token
func NewToken(tokenType TokenType, value string) *Token {
	return &Token{Type: tokenType, Value: value}
}

// Convert token type to string
func TokenTypeToStr(tokenType TokenType) string {
	switch tokenType {
	case TOKEN_ID:
		return "TOKEN_ID"
	case TOKEN_EQUALS:
		return "TOKEN_EQUALS"
	case TOKEN_LPAREN:
		return "TOKEN_LPAREN"
	case TOKEN_RPAREN:
		return "TOKEN_RPAREN"
	case TOKEN_LBRACE:
		return "TOKEN_LBRACE"
	case TOKEN_RBRACE:
		return "TOKEN_RBRACE"
	case TOKEN_LBRACKET:
		return "TOKEN_LBRACKET"
	case TOKEN_RBRACKET:
		return "TOKEN_RBRACKET"
	case TOKEN_COLON:
		return "TOKEN_COLON"
	case TOKEN_COMMA:
		return "TOKEN_COMMA"
	case TOKEN_LT:
		return "TOKEN_LT"
	case TOKEN_GT:
		return "TOKEN_GT"
	case TOKEN_ARROW_RIGHT:
		return "TOKEN_ARROW_RIGHT"
	case TOKEN_INT:
		return "TOKEN_INT"
	case TOKEN_STRING:
		return "TOKEN_STRING"
	case TOKEN_STATEMENT:
		return "TOKEN_STATEMENT"
	case TOKEN_SEMI:
		return "TOKEN_SEMI"
	case TOKEN_PLUS:
		return "TOKEN_PLUS"
	case TOKEN_MINUS:
		return "TOKEN_MINUS"
	case TOKEN_DIV:
		return "TOKEN_DIV"
	case TOKEN_MUL:
		return "TOKEN_MUL"
	case TOKEN_EOF:
		return "TOKEN_EOF"
	default:
		return "UNKNOWN"
	}
}

// Convert token to string representation
func (t Token) String() string {
	return fmt.Sprintf("<type=`%s`, int_type=`%d`, value=`%s`>", TokenTypeToStr(t.Type), t.Type, t.Value)
}
