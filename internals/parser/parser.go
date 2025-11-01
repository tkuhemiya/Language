package parser

import (
	"Language/internals/ast"
	"Language/internals/lexer"
	Token "Language/internals/token"
)

type Parser struct {
	l         *lexer.Lexer
	curToken  Token.Token
	peekToken Token.Token
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}
	// Read two tokens, so curToken and peekToken are both set
	p.nextToken()
	p.nextToken()
	return p
}
func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}
func (p *Parser) ParseProgram() *ast.Program {
	return nil
}
