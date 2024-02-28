package parser

import (
	"github.com/dipeshdulal/monkey-interpreter-go/ast"
	"github.com/dipeshdulal/monkey-interpreter-go/lexer"
	"github.com/dipeshdulal/monkey-interpreter-go/token"
)

type Parser struct {
	l *lexer.Lexer

	curToken  token.Token
	peekToken token.Token
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}

	// two tokens reading, curToken and peekToken are set
	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
  // todo: add code to parse ast
	return nil
}
