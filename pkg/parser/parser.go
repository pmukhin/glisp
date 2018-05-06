package parser

import (
	"fmt"
	"strconv"

	"github.com/pmukhin/glisp/pkg/ast"
	"github.com/pmukhin/glisp/pkg/scanner"
	"github.com/pmukhin/glisp/pkg/token"
)

type Parser struct {
	tokBackup []token.Token

	tok2infix map[token.Type]func() ast.Expression
	tok2macro map[string]func(token.Token) ast.Expression

	scn       *scanner.Scanner
	currToken token.Token
	error     error
}

func New(scn *scanner.Scanner) *Parser {
	p := new(Parser)
	// init backup
	p.tokBackup = make([]token.Token, 0, 256)

	p.scn = scn
	p.error = nil
	p.currToken = token.Token{Type: token.EOF, Literal: "EOF", Pos: -1}

	p.tok2infix = make(map[token.Type]func() ast.Expression)
	p.tok2infix[token.ParenOp] = p.parseFunctionCall
	p.tok2infix[token.SingleQuote] = p.parseList
	p.tok2infix[token.Integer] = p.parseInteger
	p.tok2infix[token.Float] = p.parseFloat
	p.tok2infix[token.String] = p.parseString
	//p.tok2infix[token.Rune] = p.parseRune
	p.tok2infix[token.Identifier] = p.parseIdentifier
	p.tok2infix[token.BracketOp] = p.parseVector

	p.tok2macro = make(map[string]func(token.Token) ast.Expression)
	p.tok2macro["defvar"] = p.parseDefVar

	p.next()

	return p
}

func (p *Parser) next() {
	p.tokBackup = append(p.tokBackup, p.currToken)

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
	switch p.currToken.Type {
	case token.EOF:
		return nil
		// case token.Semicolon:
	default:
		return p.parseExpressionStatement()
	}
}

func (p *Parser) parseExpressionStatement() ast.Statement {
	stmt := &ast.ExpressionStatement{}
	stmt.Expression = p.parseExpression()

	return stmt
}

func (p *Parser) parseExpression() ast.Expression {
	infixParse, ok := p.tok2infix[p.currToken.Type]
	if !ok {
		p.expectError("no infix parser for %s", p.currToken.Type)
		return nil
	}
	return infixParse()
}

func (p *Parser) parseDefVar(tok token.Token) ast.Expression {
	dve := &ast.DefVarExpression{Token: tok}
	dve.Name = p.parseIdentifier().(*ast.IdentifierExpression)
	dve.Value = p.parseExpression()
	dve.Comment = p.parseString()

	p.assert(token.ParenCl)
	p.next() // eat `)`

	return dve
}

func (p *Parser) parseList() ast.Expression {
	le := &ast.ListExpression{Token: p.currToken}
	p.next() // eat `'`

	p.assert(token.ParenOp)
	p.next() // eat `(`

	le.Elements = p.parseExpressionList()
	p.next() // eat `)`

	return le
}

func (p *Parser) parseVector() ast.Expression {
	ve := &ast.VectorExpression{Token: p.currToken}
	p.next() // eat `[`

	ve.Elements = p.parseExpressionList()

	p.assert(token.BracketCl)
	p.next() // eat `]`

	return ve
}

func (p *Parser) parseExpressionList() []ast.Expression {
	ls := make([]ast.Expression, 0, 8)
	for p.currToken.Type != token.ParenCl &&
		p.currToken.Type != token.BracketCl {
		res := p.parseExpression()
		if res == nil || p.error != nil {
			return nil
		}
		ls = append(ls, res)
	}

	return ls
}

func (p *Parser) parseFunctionCall() ast.Expression {
	prToken := p.currToken // if it's a fun call
	p.next() // eat `(`

	idToken := p.currToken // if it's a macro
	callee := p.parseIdentifier()
	macroFun, ok := p.tok2macro[callee.(*ast.IdentifierExpression).Value]
	if ok {
		return macroFun(idToken)
	}

	fc := &ast.FunctionCall{Token: prToken}

	fc.Callee = callee
	fc.Args = p.parseExpressionList()
	p.next() // eat ')'

	return fc
}

func (p *Parser) parseInteger() ast.Expression {
	ie := &ast.IntegerExpression{Token: p.currToken}
	v, err := strconv.ParseInt(p.currToken.Literal, 10, 64)

	if err != nil {
		p.expectError(err.Error())
		return nil
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
	se := &ast.StringExpression{Token: p.currToken}
	se.Value = p.currToken.Literal
	p.next() // eat String

	return se
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

	if p.error != nil {
		return nil, p.error
	}
	program.Statements = statements

	return program, nil
}

func (p *Parser) TokList() []token.Token {
	return p.tokBackup
}
