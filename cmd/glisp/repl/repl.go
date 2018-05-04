package repl

import (
	"bufio"
	"os"
	"fmt"
	"strings"
	"github.com/pmukhin/glisp/pkg/interpreter"
	"github.com/pmukhin/glisp/pkg/scanner"
	"github.com/pmukhin/glisp/pkg/parser"
)

func Main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Printf("glisp> ")
		bts, err := reader.ReadString('\n')

		if err != nil {
			fmt.Println(err.Error())
			continue
		}

		scn := scanner.New(strings.Trim(string(bts), "\n"))
		prs := parser.New(scn)

		prg, err := prs.Parse()
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		res, err := interpreter.Interpret(prg)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		if res == nil {
			fmt.Println()
			continue
		}

		// trim printed output to void duplication newlines
		fmt.Println(strings.Trim(res.String(), "\n\r"))
	}
}
