package javascript

// Lexer structure
type Lexer struct {
	Source string // Source code
	Size   int    // Size of source
	Char   byte   // Current character
	Index  int    // Current position
}

// Constructor for Lexer
func NewLexer(src string) *Lexer {
	l := &Lexer{
		Source: src,
		Size:   len(src),
		Index:  0,
	}
	if l.Size > 0 {
		l.Char = src[0]
	}
	return l
}

// Advance the lexer
func (l *Lexer) Advance() {
	l.Index++
	if l.Index < l.Size {
		l.Char = l.Source[l.Index]
	} else {
		l.Char = 0 // Null terminator (EOF)
	}
}

// Peek ahead in the source code
func (l *Lexer) Peek(offset int) byte {
	pos := l.Index + offset
	if pos >= l.Size {
		return 0 // Return EOF if out of bounds
	}
	return l.Source[pos]
}

// Advance lexer and return a given token
func (l *Lexer) AdvanceWith(token Token) Token {
	l.Advance()
	return token
}

// Advance lexer while keeping current character as token value
func (l *Lexer) AdvanceCurrent(tokenType TokenType) Token {
	token := NewToken(tokenType, string(l.Char))
	l.Advance()
	return *token
}

// Skip whitespace
func (l *Lexer) SkipWhitespace() {
	for l.Char == ' ' || l.Char == '\t' || l.Char == '\n' || l.Char == '\r' {
		l.Advance()
	}
}

// Skip comments
func (l *Lexer) SkipComment() {
	if l.Char == '/' && l.Peek(1) == '/' { // Single-line comment (//)
		for l.Char != '\n' && l.Char != 0 {
			l.Advance()
		}
	} else if l.Char == '/' && l.Peek(1) == '*' { // Multi-line comment (/* ... */)
		l.Advance() // Skip '/'
		l.Advance() // Skip '*'
		for l.Char != 0 && !(l.Char == '*' && l.Peek(1) == '/') {
			l.Advance()
		}
		l.Advance() // Skip '*'
		l.Advance() // Skip '/'
	}
}

// Parse an identifier (variable names, keywords)
func (l *Lexer) ParseIdentifier() Token {
	start := l.Index
	for (l.Char >= 'a' && l.Char <= 'z') || (l.Char >= 'A' && l.Char <= 'Z') || (l.Char >= '0' && l.Char <= '9') || l.Char == '_' {
		l.Advance()
	}
	return Token{Type: TOKEN_ID, Value: l.Source[start:l.Index]}
}

// Parse a number (integer)
func (l *Lexer) ParseNumber() Token {
	start := l.Index
	for l.Char >= '0' && l.Char <= '9' {
		l.Advance()
	}
	return Token{Type: TOKEN_INT, Value: l.Source[start:l.Index]}
}

// Parse a string
func (l *Lexer) ParseString() Token {
	l.Advance() // Skip the opening quote
	start := l.Index
	for l.Char != '"' && l.Char != 0 {
		l.Advance()
	}
	value := l.Source[start:l.Index]
	l.Advance() // Skip the closing quote
	return Token{Type: TOKEN_STRING, Value: value}
}

// Get the next token
func (l *Lexer) NextToken() Token {
	for l.Char != 0 {
		if l.Char == ' ' || l.Char == '\n' || l.Char == '\t' || l.Char == '\r' {
			l.SkipWhitespace()
			continue
		}

		if l.Char == '/' && (l.Peek(1) == '/' || l.Peek(1) == '*') {
			l.SkipComment()
			continue
		}

		switch l.Char {
		case '=':
			return l.AdvanceCurrent(TOKEN_EQUALS)
		case '(':
			return l.AdvanceCurrent(TOKEN_LPAREN)
		case ')':
			return l.AdvanceCurrent(TOKEN_RPAREN)
		case '{':
			return l.AdvanceCurrent(TOKEN_LBRACE)
		case '}':
			return l.AdvanceCurrent(TOKEN_RBRACE)
		case '[':
			return l.AdvanceCurrent(TOKEN_LBRACKET)
		case ']':
			return l.AdvanceCurrent(TOKEN_RBRACKET)
		case ':':
			return l.AdvanceCurrent(TOKEN_COLON)
		case ',':
			return l.AdvanceCurrent(TOKEN_COMMA)
		case '<':
			return l.AdvanceCurrent(TOKEN_LT)
		case '>':
			return l.AdvanceCurrent(TOKEN_GT)
		case ';':
			return l.AdvanceCurrent(TOKEN_SEMI)
		case '+':
			return l.AdvanceCurrent(TOKEN_PLUS)
		case '-':
			if l.Peek(1) == '>' {
				l.Advance()
				l.Advance()
				return Token{Type: TOKEN_ARROW_RIGHT, Value: "->"}
			}
			return l.AdvanceCurrent(TOKEN_MINUS)
		case '/':
			return l.AdvanceCurrent(TOKEN_DIV)
		case '*':
			return l.AdvanceCurrent(TOKEN_MUL)
		case '"':
			return l.ParseString()
		}

		// Identifiers
		if (l.Char >= 'a' && l.Char <= 'z') || (l.Char >= 'A' && l.Char <= 'Z') || l.Char == '_' {
			return l.ParseIdentifier()
		}

		// Numbers
		if l.Char >= '0' && l.Char <= '9' {
			return l.ParseNumber()
		}

		// Unknown character
		l.Advance()
	}
	return Token{Type: TOKEN_EOF, Value: ""}
}
