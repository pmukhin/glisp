package parser

import (
	"github.com/pmukhin/glisp/pkg/scanner"
	"github.com/pmukhin/glisp/pkg/token"
	"github.com/pmukhin/glisp/pkg/ast"
	"strconv"
	"fmt"
	"io"
)

type Parser struct {
	tok2infix map[token.Type]func() ast.Expression
	scn       *scanner.Scanner
	currToken token.Token
	error     error
}

func New(scn *scanner.Scanner) *Parser {
	p := new(Parser)
	p.scn = scn
	p.error = nil
	p.currToken = token.Token{Type: token.EOF, Literal: "EOF", Pos: -1}
	p.tok2infix = make(map[token.Type]func() ast.Expression)

	p.tok2infix[token.ParenOp] = p.parseFunctionCall
	p.tok2infix[token.Integer] = p.parseInteger
	p.tok2infix[token.Float] = p.parseFloat
	p.tok2infix[token.String] = p.parseString
	p.tok2infix[token.Rune] = p.parseRune
	p.tok2infix[token.Identifier] = p.parseIdentifier

	p.next()

	return p
}

func (p *Parser) next() {
	tok := p.scn.Next()
	p.currToken = tok
}

func (p *Parser) expectError(msg string, a ...interface{}) {
	p.error = fmt.Errorf(msg, a...)
}

func (p *Parser) assert(typ token.Type) {
	if p.currToken.Type == typ {
		return
	}
	p.expectError("expected token %s, got %s", typ, p.currToken.Type)
}

func (p *Parser) parseIdentifier() ast.Expression {
	p.assert(token.Identifier)
	defer p.next() // eat Identifier

	return &ast.IdentifierExpression{Token: p.currToken, Value: p.currToken.Literal}
}

func (p *Parser) parseStatement() ast.Statement {
	if p.currToken.Type == token.EOF {
		p.error = io.EOF
		return nil
	}

	if p.currToken.Type == token.ParenOp {
		stmt := &ast.ExpressionStatement{}
		stmt.Expression = p.parseExpression()

		return stmt
	}

	panic("unsupported statement type")
}

func (p *Parser) parseExpression() ast.Expression {
	infixParse, ok := p.tok2infix[p.currToken.Type]
	if !ok {
		panic("")
	}
	return infixParse()
}

func (p *Parser) parseFunctionCall() ast.Expression {
	fc := &ast.FunctionCall{Token: p.currToken}
	p.assert(token.ParenOp)
	p.next() // eat `(`

	fc.Callee = p.parseIdentifier()
	for p.currToken.Type != token.ParenCl {
		fc.Args = append(fc.Args, p.parseExpression())
	}
	p.assert(token.ParenCl)
	p.next() // eat `)`

	return fc
}

func (p *Parser) parseInteger() ast.Expression {
	ie := &ast.IntegerExpression{Token: p.currToken}
	v, err := strconv.ParseInt(p.currToken.Literal, 10, 64)

	if err != nil {
		p.expectError(err.Error())
	}

	ie.Value = v
	p.next() // eat Integer

	return ie
}

func (p *Parser) parseFloat() ast.Expression {
	fe := &ast.FloatExpression{Token: p.currToken}
	v, err := strconv.ParseFloat(p.currToken.Literal, 64)

	if err != nil {
		p.expectError(err.Error())
	}

	fe.Value = v
	p.next()

	return fe
}

func (p *Parser) parseString() ast.Expression {
	panic("implement me")
}

func (p *Parser) parseRune() ast.Expression {
	panic("implement me")
}

func (p *Parser) Parse() (*ast.Program, error) {
	program := new(ast.Program)
	statements := make([]ast.Statement, 0, 256)

	for {
		stmt := p.parseStatement()
		if stmt == nil {
			break
		}
		statements = append(statements, stmt)
	}

	if p.error != nil && p.error != io.EOF {
		return nil, p.error
	}
	program.Statements = statements

	return program, nil
}
