package main

import (
	"os"
	"fmt"
	"strings"
	"bufio"
	"github.com/pmukhin/glisp/pkg/scanner"
	"github.com/pmukhin/glisp/pkg/parser"
	"github.com/pmukhin/glisp/pkg/interpreter"
)

func main() {
	args := os.Args

	if len(args) == 1 {
		usage()
	}

	switch args[1] {
	case "--repl", "-r":
		runREPL()
	default:
		if strings.HasSuffix(args[0], ".glisp") {
			runFile()
		}
		usage()
	}
}

func runFile() {
	os.Exit(0)
}

func runREPL() {
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

func usage() {
	fmt.Println("glisp <file.glisp> | glisp --runREPL")
	os.Exit(0)
}
