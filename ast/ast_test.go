package ast_test

import (
	"testing"

	"github.com/dipeshdulal/monkey-interpreter-go/ast"
	"github.com/dipeshdulal/monkey-interpreter-go/token"
)

func TestString(t *testing.T) {
	program := &ast.Program{
		Statements: []ast.Statement{
			&ast.LetStatement{
				Token: token.Token{Type: token.LET, Literal: "let"},
				Name: &ast.Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "testVar"},
					Value: "testVar",
				},
        Value: &ast.Identifier{
          Token: token.Token{Type: token.IDENT, Literal: "anotherVar"},
          Value: "anotherVar",
        },
			},
		},
	}

  if program.String() != "let myVar = anotherVar;" {
    t.Errorf("program.String() wrong. got=%q", program.String())
  }
}
