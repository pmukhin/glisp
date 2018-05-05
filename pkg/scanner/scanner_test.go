package scanner

import (
	"testing"

	"github.com/pmukhin/glisp/pkg/token"
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

	if len(tokens) != len(expected) {
		t.Errorf("wrong number of tokens returned: exp. %d vs %d given",
			len(expected), len(tokens))
		return
	}

	for i := 0; i < len(expected); i++ {
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
	doTestWithLiterals(t, `(+ "test" "b")`, []tokenAndLiteral{
		{token.ParenOp, "("},
		{token.Identifier, "+"},
		{token.String, "test"},
		{token.String, "b"},
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

func TestScanner_Next_ListOfRunes(t *testing.T) {
	doTest(t, `(print '("a" "b" "c"))`, []token.Type{
		token.ParenOp,
		token.Identifier,
		token.SingleQuote,
		token.ParenOp,
		token.String,
		token.String,
		token.String,
		token.ParenCl,
		token.ParenCl,
	})
}

func TestScanner_Next_List(t *testing.T) {
	doTest(t, `(print '(a b))`, []token.Type{
		token.ParenOp,
		token.Identifier,
		token.SingleQuote,
		token.ParenOp,
		token.Identifier,
		token.Identifier,
		token.ParenCl,
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
