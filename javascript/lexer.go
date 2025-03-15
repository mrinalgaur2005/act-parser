package javascript

import (
	"fmt"
	"regexp"
)

type regexHandler func(lex *lexer, regex *regexp.Regexp)
type regexPattern struct {
	regex   *regexp.Regexp
	handler regexHandler
}
type lexer struct {
	patterns []regexPattern
	Tokens   []Token
	source   string
	pos      int
}

func (lex *lexer) advanceN(n int) {
	lex.pos += n
}

func (lex *lexer) push(token Token) {
	lex.Tokens = append(lex.Tokens, token)
}

func (lex *lexer) at() byte {
	return lex.source[lex.pos]
}

func (lex *lexer) remainder() string {
	return lex.source[lex.pos:]
}

func (lex *lexer) at_eof() bool {
	return lex.pos >= len(lex.source)
}

func Tokenize(source string) []Token {
	lex := createLexer(source)

	for !lex.at_eof() {
		matched := false

		for _, pattern := range lex.patterns {
			loc := pattern.regex.FindStringIndex(lex.remainder())

			if loc != nil && loc[0] == 0 {
				pattern.handler(lex, pattern.regex)
				matched = true
				break
			}
		}
		if !matched {
			panic(fmt.Sprintf("Lexor Error !!! Unrecognized Token near %s \n", lex.remainder()))
		}
	}
	lex.push(NewToken(EOF, "EOF"))
	return lex.Tokens
}

func defaulHandler(Type TokenType, value string) regexHandler {
	return func(lex *lexer, regex *regexp.Regexp) {
		lex.advanceN(len(value))
		lex.push(NewToken(Type, value))
	}
}

func createLexer(source string) *lexer {
	return &lexer{
		pos:    0,
		source: source,
		Tokens: make([]Token, 0),
		patterns: []regexPattern{

			//need to add template literals here

			{regexp.MustCompile(`[0-9]+(\.[0-9]+)?`), numberHandler},
			{regexp.MustCompile(`"(?:[^"\\]|\\.)*"`), stringHandler},
			{regexp.MustCompile(`'(?:[^'\\]|\\.)*'`), stringHandler},

			{regexp.MustCompile(`\/\/.*`), skipHandler},

			{regexp.MustCompile(`/\*[\s\S]*?\*/`), skipHandler},

			{regexp.MustCompile(`\s+`), skipHandler},

			{regexp.MustCompile(`\[`), defaulHandler(OPEN_BRACK, "[")},
			{regexp.MustCompile(`\]`), defaulHandler(CLOSED_BRACK, "]")},
			{regexp.MustCompile(`\{`), defaulHandler(OPEN_CURLY, "{")},
			{regexp.MustCompile(`\}`), defaulHandler(CLOSE_CURLY, "}")},
			{regexp.MustCompile(`\(`), defaulHandler(OPEN_PAREN, "(")},
			{regexp.MustCompile(`\)`), defaulHandler(CLOSE_PAREN, ")")},

			{regexp.MustCompile(`===`), defaulHandler(STRICT_EQUALS, "===")},
			{regexp.MustCompile(`!==`), defaulHandler(STRICT_NOT_EQUALS, "!==")},
			{regexp.MustCompile(`==`), defaulHandler(EQUALS, "==")},
			{regexp.MustCompile(`!=`), defaulHandler(NOT_EQUALS, "!=")},

			{regexp.MustCompile(`\*\*=`), defaulHandler(EXPONENT_EQUALS, "**=")},
			{regexp.MustCompile(`\+=`), defaulHandler(PLUS_EQUALS, "+=")},
			{regexp.MustCompile(`-=`), defaulHandler(MINUS_EQUALS, "-=")},
			{regexp.MustCompile(`\*=`), defaulHandler(MUL_EQUALS, "*=")},
			{regexp.MustCompile(`/=`), defaulHandler(DIV_EQUALS, "/=")},
			{regexp.MustCompile(`%=`), defaulHandler(MOD_EQUALS, "%=")},

			{regexp.MustCompile(`\.\.\.`), defaulHandler(DOT_DOT_DOT, "...")},
			{regexp.MustCompile(`\.\.`), defaulHandler(DOT_DOT, "..")},
			{regexp.MustCompile(`=>`), defaulHandler(ARROW, "=>")},
			{regexp.MustCompile("`"), defaulHandler(BACKTICK, "`")},

			{regexp.MustCompile(`<=`), defaulHandler(LESS_EQUALS, "<=")},
			{regexp.MustCompile(`>=`), defaulHandler(GREATER_EQUALS, ">=")},
			{regexp.MustCompile(`<`), defaulHandler(LESS, "<")},
			{regexp.MustCompile(`>`), defaulHandler(GREATER, ">")},

			{regexp.MustCompile(`=`), defaulHandler(ASSIGNMENT, "=")},

			{regexp.MustCompile(`\?`), defaulHandler(QUESTION, "?")},

			{regexp.MustCompile(`\|\|`), defaulHandler(OR, "||")},
			{regexp.MustCompile(`!`), defaulHandler(NOT, "!")},
			{regexp.MustCompile(`&&`), defaulHandler(AND, "&&")},
			{regexp.MustCompile(`\|`), defaulHandler(BIT_OR, "|")},
			{regexp.MustCompile(`&`), defaulHandler(BIT_AND, "&")},
			{regexp.MustCompile(`\^`), defaulHandler(BIT_XOR, "^")},
			{regexp.MustCompile(`~`), defaulHandler(BIT_NOT, "~")},

			{regexp.MustCompile(`\+\+`), defaulHandler(PLUS_PLUS, "++")},
			{regexp.MustCompile(`--`), defaulHandler(MINUS_MINUS, "--")},
			{regexp.MustCompile(`\*\*`), defaulHandler(EXPONENT, "**")},

			{regexp.MustCompile(`\+`), defaulHandler(PLUS, "+")},
			{regexp.MustCompile(`-`), defaulHandler(DASH, "-")},
			{regexp.MustCompile(`/`), defaulHandler(SLASH, "/")},
			{regexp.MustCompile(`\*`), defaulHandler(STAR, "*")},
			{regexp.MustCompile(`%`), defaulHandler(PERCENT, "%")},

			{regexp.MustCompile(`\?\?`), defaulHandler(NULLISH_COALESCING, "??")},
			{regexp.MustCompile(`\?\.`), defaulHandler(OPTIONAL_CHAINING, "?.")},

			{regexp.MustCompile(`,`), defaulHandler(COMMA, ",")},
			{regexp.MustCompile(`;`), defaulHandler(SEMICOLON, ";")},
			{regexp.MustCompile(`:`), defaulHandler(COLON, ":")},
			{regexp.MustCompile(`\.`), defaulHandler(DOT, ".")},
		},
	}
}

func numberHandler(lex *lexer, regex *regexp.Regexp) {
	match := regex.FindString(lex.remainder())
	lex.push(NewToken(NUMBER, match))
	lex.advanceN(len(match))
}

func skipHandler(lex *lexer, regex *regexp.Regexp) {
	match := regex.FindStringIndex(lex.remainder())
	lex.advanceN(match[1])
}

func stringHandler(lex *lexer, regex *regexp.Regexp) {
	match := regex.FindStringIndex(lex.remainder())
	stringLiteral := lex.remainder()[match[0]:match[1]]

	lex.push(NewToken(STRING, stringLiteral))
	lex.advanceN(len(stringLiteral))

}
