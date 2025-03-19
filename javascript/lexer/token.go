package lexer

import "fmt"

type TokenType int

const (
	EOF TokenType = iota
	NUMBER
	STRING
	IDENTIFIER
	REGEX

	TEMPLATE_EXPR_START
	TEMPLATE_EXPR_END
	TEMPLATE_LITERAL_START
	TEMPLATE_LITERAL_END
	TEMPLATE_EXPR

	OPEN_BRACK
	CLOSED_BRACK
	OPEN_CURLY
	CLOSE_CURLY
	OPEN_PAREN
	CLOSE_PAREN

	ASSIGNMENT
	EQUALS
	STRICT_EQUALS
	NOT
	NOT_EQUALS
	STRICT_NOT_EQUALS

	LESS
	LESS_EQUALS
	GREATER
	GREATER_EQUALS

	OR
	AND
	BIT_OR
	BIT_AND
	BIT_XOR
	BIT_NOT

	SHIFT_ZERO_FILL
	SHIFT_RIGHT

	PLUS
	DASH
	SLASH
	STAR
	PERCENT
	EXPONENT
	PLUS_PLUS
	MINUS_MINUS
	PLUS_EQUALS
	MINUS_EQUALS
	MUL_EQUALS
	DIV_EQUALS
	MOD_EQUALS
	EXPONENT_EQUALS

	NULLISH_COALESCING
	OPTIONAL_CHAINING

	DOT
	DOT_DOT
	DOT_DOT_DOT
	SEMICOLON
	COLON
	QUESTION
	COMMA
	ARROW

	LET
	CONST
	VAR
	CLASS
	NEW
	IMPORT
	FROM
	EXPORT
	FUNCTION
	RETURN
	IF
	ELSE
	SWITCH
	CASE
	DEFAULT
	BREAK
	CONTINUE
	FOR
	WHILE
	DO
	TRY
	CATCH
	FINALLY
	THROW
	ASYNC
	AWAIT
	SUPER
	THIS
	EXTENDS
	STATIC
	DELETE
	YIELD
	DEBUGGER
	TYPEOF
	IN
	INSTANCEOF
	NULL
	UNDEFINED
	TRUE
	FALSE

	BACKTICK
)

var reserved_lu map[string]TokenType = map[string]TokenType{
	"let":        LET,
	"const":      CONST,
	"var":        VAR,
	"class":      CLASS,
	"new":        NEW,
	"import":     IMPORT,
	"from":       FROM,
	"export":     EXPORT,
	"function":   FUNCTION,
	"return":     RETURN,
	"if":         IF,
	"else":       ELSE,
	"switch":     SWITCH,
	"case":       CASE,
	"default":    DEFAULT,
	"break":      BREAK,
	"continue":   CONTINUE,
	"for":        FOR,
	"while":      WHILE,
	"do":         DO,
	"try":        TRY,
	"catch":      CATCH,
	"finally":    FINALLY,
	"throw":      THROW,
	"async":      ASYNC,
	"await":      AWAIT,
	"super":      SUPER,
	"this":       THIS,
	"extends":    EXTENDS,
	"static":     STATIC,
	"delete":     DELETE,
	"yield":      YIELD,
	"debugger":   DEBUGGER,
	"typeof":     TYPEOF,
	"in":         IN,
	"instanceof": INSTANCEOF,
	"null":       NULL,
	"undefined":  UNDEFINED,
	"true":       TRUE,
	"false":      FALSE,
}

type Token struct {
	Type  TokenType
	Value string
}

func NewToken(tokenType TokenType, value string) Token {
	return Token{Type: tokenType, Value: value}
}

func TokenTypeToStr(tokenType TokenType) string {
	switch tokenType {
	case EOF:
		return "EOF"
	case NUMBER:
		return "NUMBER"
	case STRING:
		return "STRING"
	case IDENTIFIER:
		return "IDENTIFIER"
	case REGEX:
		return "REGEX"
	case TEMPLATE_EXPR_START:
		return "TEMPLATE_EXPR_START"
	case TEMPLATE_EXPR_END:
		return "TEMPLATE_EXPR_END"
	case TEMPLATE_LITERAL_START:
		return "TEMPLATE_LITERAL_START"
	case TEMPLATE_LITERAL_END:
		return "TEMPLATE_LITERAL_END"
	case TEMPLATE_EXPR:
		return "TEMPLATE_EXPR"
	case OPEN_BRACK:
		return "OPEN_BRACK"
	case CLOSED_BRACK:
		return "CLOSED_BRACK"
	case OPEN_CURLY:
		return "OPEN_CURLY"
	case CLOSE_CURLY:
		return "CLOSE_CURLY"
	case OPEN_PAREN:
		return "OPEN_PAREN"
	case CLOSE_PAREN:
		return "CLOSE_PAREN"
	case ASSIGNMENT:
		return "ASSIGNMENT"
	case EQUALS:
		return "EQUALS"
	case STRICT_EQUALS:
		return "STRICT_EQUALS"
	case NOT:
		return "NOT"
	case NOT_EQUALS:
		return "NOT_EQUALS"
	case STRICT_NOT_EQUALS:
		return "STRICT_NOT_EQUALS"
	case LESS:
		return "LESS"
	case LESS_EQUALS:
		return "LESS_EQUALS"
	case GREATER:
		return "GREATER"
	case GREATER_EQUALS:
		return "GREATER_EQUALS"
	case OR:
		return "OR"
	case AND:
		return "AND"
	case BIT_OR:
		return "BIT_OR"
	case BIT_AND:
		return "BIT_AND"
	case BIT_XOR:
		return "BIT_XOR"
	case BIT_NOT:
		return "BIT_NOT"
	case SHIFT_ZERO_FILL:
		return "SHIFT_ZERO_FILL"
	case SHIFT_RIGHT:
		return "SHIFT_RIGHT"
	case PLUS:
		return "PLUS"
	case DASH:
		return "DASH"
	case SLASH:
		return "SLASH"
	case STAR:
		return "STAR"
	case PERCENT:
		return "PERCENT"
	case EXPONENT:
		return "EXPONENT"
	case PLUS_PLUS:
		return "PLUS_PLUS"
	case MINUS_MINUS:
		return "MINUS_MINUS"
	case PLUS_EQUALS:
		return "PLUS_EQUALS"
	case MINUS_EQUALS:
		return "MINUS_EQUALS"
	case MUL_EQUALS:
		return "MUL_EQUALS"
	case DIV_EQUALS:
		return "DIV_EQUALS"
	case MOD_EQUALS:
		return "MOD_EQUALS"
	case EXPONENT_EQUALS:
		return "EXPONENT_EQUALS"
	case NULLISH_COALESCING:
		return "NULLISH_COALESCING"
	case OPTIONAL_CHAINING:
		return "OPTIONAL_CHAINING"
	case DOT:
		return "DOT"
	case DOT_DOT:
		return "DOT_DOT"
	case DOT_DOT_DOT:
		return "DOT_DOT_DOT"
	case SEMICOLON:
		return "SEMICOLON"
	case COLON:
		return "COLON"
	case QUESTION:
		return "QUESTION"
	case COMMA:
		return "COMMA"
	case ARROW:
		return "ARROW"
	case LET:
		return "LET"
	case CONST:
		return "CONST"
	case VAR:
		return "VAR"
	case CLASS:
		return "CLASS"
	case NEW:
		return "NEW"
	case IMPORT:
		return "IMPORT"
	case FROM:
		return "FROM"
	case EXPORT:
		return "EXPORT"
	case FUNCTION:
		return "FUNCTION"
	case RETURN:
		return "RETURN"
	case IF:
		return "IF"
	case ELSE:
		return "ELSE"
	case SWITCH:
		return "SWITCH"
	case CASE:
		return "CASE"
	case DEFAULT:
		return "DEFAULT"
	case BREAK:
		return "BREAK"
	case CONTINUE:
		return "CONTINUE"
	case FOR:
		return "FOR"
	case WHILE:
		return "WHILE"
	case DO:
		return "DO"
	case TRY:
		return "TRY"
	case CATCH:
		return "CATCH"
	case FINALLY:
		return "FINALLY"
	case THROW:
		return "THROW"
	case ASYNC:
		return "ASYNC"
	case AWAIT:
		return "AWAIT"
	case SUPER:
		return "SUPER"
	case THIS:
		return "THIS"
	case EXTENDS:
		return "EXTENDS"
	case STATIC:
		return "STATIC"
	case DELETE:
		return "DELETE"
	case YIELD:
		return "YIELD"
	case DEBUGGER:
		return "DEBUGGER"
	case TYPEOF:
		return "TYPEOF"
	case IN:
		return "IN"
	case INSTANCEOF:
		return "INSTANCEOF"
	case NULL:
		return "NULL"
	case UNDEFINED:
		return "UNDEFINED"
	case TRUE:
		return "TRUE"
	case FALSE:
		return "FALSE"
	case BACKTICK:
		return "BACKTICK"
	default:
		return fmt.Sprintf("UNKNOWN(%d)", tokenType)
	}
}

func (token Token) Debug() {
	if token.Type == IDENTIFIER || token.Type == NUMBER || token.Type == STRING || token.Type == TEMPLATE_EXPR {
		fmt.Printf("%s (%s)\n", TokenTypeToStr(token.Type), token.Value)
	} else {
		fmt.Printf("%s ()\n", TokenTypeToStr(token.Type))
	}
}
