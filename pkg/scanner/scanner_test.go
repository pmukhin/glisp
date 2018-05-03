package scanner

import (
	"testing"
	"glisp/pkg/token"
)

type tokenAndLiteral struct {
	typ token.Type
	lit string
}

// test template
func doTest(t *testing.T, input string, expected []token.Type) {
	scn := New(input)
	tokens := make([]token.Token, 0, 256)

	for {
		tok := scn.Next()
		if tok.Type == token.EOF {
			break
		}
		tokens = append(tokens, tok)
	}

	for i := 0; i < len(tokens); i++ {
		if tokens[i].Type != expected[i] {
			t.Errorf(
				"%d: expected token of type %v, got %v in pos %d",
				i,
				expected[i],
				tokens[i].Type,
				tokens[i].Pos,
			)
		}
	}
}

// test template with literals
func doTestWithLiterals(t *testing.T, input string, expected []tokenAndLiteral) {
	scn := New(input)
	tokens := make([]token.Token, 0, 256)

	for {
		tok := scn.Next()
		if tok.Type == token.EOF {
			break
		}
		tokens = append(tokens, tok)
	}

	for i := 0; i < len(tokens); i++ {
		if tokens[i].Type != expected[i].typ {
			t.Errorf(
				"%d: expected token of type %v, got %v in pos %d",
				i,
				expected[i].typ,
				tokens[i].Type,
				tokens[i].Pos,
			)
		}
		if tokens[i].Literal != expected[i].lit {
			t.Errorf(
				"%d: expected token of literal %v, got %v in pos %d",
				i,
				expected[i].lit,
				tokens[i].Literal,
				tokens[i].Pos,
			)
		}
	}
}

func TestScanner_Next_4(t *testing.T) {
	doTestWithLiterals(t, `(+ "test" 'a')`, []tokenAndLiteral{
		{token.ParenOp, "("},
		{token.Identifier, "+"},
		{token.String, "test"},
		{token.Rune, "a"},
		{token.ParenCl, ")"},
	})
}

func TestScanner_Next_3(t *testing.T) {
	doTestWithLiterals(t, `(+ 5.545 24)`, []tokenAndLiteral{
		{token.ParenOp, "("},
		{token.Identifier, "+"},
		{token.Float, "5.545"},
		{token.Integer, "24"},
		{token.ParenCl, ")"},
	})
}

func TestScanner_Next_2(t *testing.T) {
	doTest(t, `(+ "abc" "def")`, []token.Type{
		token.ParenOp,
		token.Identifier,
		token.String,
		token.String,
		token.ParenCl,
	})
}

func TestScanner_Next_1(t *testing.T) {
	input := `
(defun fib (n) 
    (case n 
        (< 2 n)
        (+ (fib (- n 1)) (fib (- n 2)))))
`
	doTest(t, input, []token.Type{
		// (defun fib (n)
		token.ParenOp,
		token.Identifier,
		token.Identifier,
		token.ParenOp,
		token.Identifier,
		token.ParenCl,
		// (case n
		token.ParenOp,
		token.Identifier,
		token.Identifier,
		// (< 2 n)
		token.ParenOp,
		token.Identifier,
		token.Integer,
		token.Identifier,
		token.ParenCl,
		// (+ (fib (- n 1)) (fib (- n 2)) )
		token.ParenOp,
		// function identifier
		token.Identifier,
		// first arg
		// (fib (- n 1))
		token.ParenOp,
		token.Identifier,
		token.ParenOp,
		token.Identifier,
		token.Identifier,
		token.Integer,
		token.ParenCl,
		token.ParenCl,
		// second arg
		// (fib (- n 2))
		token.ParenOp,
		token.Identifier,
		token.ParenOp,
		token.Identifier,
		token.Identifier,
		token.Integer,
		token.ParenCl,
		token.ParenCl,
		// closing
		token.ParenCl,
		// ))
		token.ParenCl,
		token.ParenCl,
	})
}
