package scanner

import (
	"testing"

	"github.com/pmukhin/glisp/pkg/token"
	"fmt"
	"reflect"
)

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
		for _, tok := range tokens {
			fmt.Println(tok)
		}

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

func do(t *testing.T, input string, expected []token.Token) {
	scn := New(input)
	tokens := make([]token.Token, 0, 256)

	for {
		tok := scn.Next()
		if tok.Type == token.EOF {
			break
		}
		tokens = append(tokens, tok)
	}

	if !reflect.DeepEqual(tokens, expected) {
		t.Errorf("expected %v, got %v", expected, tokens)
	}
}

func TestScanner_Next_ScanDefVar(t *testing.T) {
	do(t, `(defvar int-list '(1 2 3) "a list of ints")`, []token.Token{
		token.New(token.ParenOp, 0, "("),
		token.New(token.Identifier, 1, "defvar"),
		token.New(token.Identifier, 8, "int-list"),
		token.New(token.SingleQuote, 17),
		token.New(token.ParenOp, 18, "("),
		token.New(token.Integer, 19, "1"),
		token.New(token.Integer, 21, "2"),
		token.New(token.Integer, 23, "3"),
		token.New(token.ParenCl, 24, ")"),
		token.New(token.String, 26, "a list of ints"),
		token.New(token.ParenCl, 42, ")"),
	})
}

func TestScanner_Next_4(t *testing.T) {
	do(t, `(+ "test" "b")`, []token.Token{
		token.New(token.ParenOp, 0, "("),
		token.New(token.Identifier, 1, "+"),
		token.New(token.String, 3, "test"),
		token.New(token.String, 10, "b"),
		token.New(token.ParenCl, 13, ")"),
	})
}

func TestScanner_Next_3(t *testing.T) {
	do(t, `(+ 5.545 24)`, []token.Token{
		token.New(token.ParenOp, 0, "("),
		token.New(token.Identifier, 1, "+"),
		token.New(token.Float, 3, "5.545"),
		token.New(token.Integer, 9, "24"),
		token.New(token.ParenCl, 11, ")"),
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

func TestScanner_Next_Vector(t *testing.T) {
	doTest(t, `(print [1 2])`, []token.Type{
		token.ParenOp,
		token.Identifier,
		token.BracketOp,
		token.Integer,
		token.Integer,
		token.BracketCl,
		token.ParenCl,
	})
}
