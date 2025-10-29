package lexer

import Token "LangLang/internals/token"

type Lexer struct {
	input        string // loads the whole file
	readPosition int    // index of reading position
	chPosition   int    // index of ch
	ch           byte   // current char
}

func NewLexer(input string) *Lexer {
	lexer := &Lexer{
		input:        input,
		readPosition: 0,
		chPosition:   0,
		ch:           0,
	}
	lexer.readChar()
	return lexer
}

func (l *Lexer) readChar() {
	l.chPosition = l.readPosition
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.readPosition += 1
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition+1]
	}
}

func (l *Lexer) NextToken() Token.Token {
	var token Token.Token

	l.skipWhiteSpace()

	switch l.ch {
	case '=':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			token = Token.Token{Type: Token.EQUAL, Literal: string(ch) + string(l.ch)}
		} else {
			token = Token.NewToken(Token.ASSIGN, l.ch)
		}

	case '!':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			token = Token.Token{Type: Token.NOTE, Literal: string(ch) + string(l.ch)}
		} else {
			token = Token.NewToken(Token.NOT, l.ch)
		}

	case '<':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			token = Token.Token{Type: Token.LDE, Literal: string(ch) + string(l.ch)}
		} else {
			token = Token.NewToken(Token.LD, l.ch)
		}

	case '>':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			token = Token.Token{Type: Token.GDE, Literal: string(ch) + string(l.ch)}
		} else {
			token = Token.NewToken(Token.GD, l.ch)
		}

	case '+':
		token = Token.NewToken(Token.PLUS, l.ch)
	case '-':
		token = Token.NewToken(Token.MINUS, l.ch)
	case '*':
		token = Token.NewToken(Token.MULT, l.ch)
	case '/':
		token = Token.NewToken(Token.DIVIDE, l.ch)
	case '^':
		token = Token.NewToken(Token.POW, l.ch)
	case ';':
		token = Token.NewToken(Token.SEMICOLON, l.ch)
	case '(':
		token = Token.NewToken(Token.LPAREN, l.ch)
	case ')':
		token = Token.NewToken(Token.RPAREN, l.ch)
	case ',':
		token = Token.NewToken(Token.COMMA, l.ch)
	case '{':
		token = Token.NewToken(Token.LBRACE, l.ch)
	case '}':
		token = Token.NewToken(Token.RBRACE, l.ch)
	case '[':
		token = Token.NewToken(Token.LBRACK, l.ch)
	case ']':
		token = Token.NewToken(Token.RBRACK, l.ch)
	case 0:
		token.Literal = ""
		token.Type = Token.EOF

	default:
		if isLetter(l.ch) {
			token.Literal = l.readIdentifier()
			token.Type = lookupIdent(token.Literal)
			return token
		} else if isDigit(l.ch) {
			token.Literal = l.readDigit()
			token.Type = Token.INT
			return token
		} else {
			token = Token.NewToken(Token.ILLEGAL, l.ch)
		}
	}

	l.readChar()
	return token
}

func (l *Lexer) skipWhiteSpace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func isLetter(char byte) bool {
	return ('a' <= char && char <= 'z') || ('A' <= char && char <= 'Z') || (char == '_')
}

func isDigit(char byte) bool {
	return '0' <= char && char <= '9'
}

func (l *Lexer) readIdentifier() string {
	position := l.chPosition
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.chPosition]
}

func (l *Lexer) readDigit() string {
	position := l.chPosition
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.chPosition]
}

func lookupIdent(ident string) Token.TokenType {
	if token, ok := Token.KeywordsTable[ident]; ok {
		return token
	}
	return Token.IDENT
}
