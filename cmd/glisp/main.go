package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/pmukhin/glisp/cmd/glisp/repl"
	"github.com/pmukhin/glisp/pkg/interpreter"
	"github.com/pmukhin/glisp/pkg/parser"
	"github.com/pmukhin/glisp/pkg/scanner"
)

func main() {
	args := os.Args

	if len(args) == 1 {
		usage()
	}

	runner := usage

	switch args[1] {
	case "--repl", "-r":
		runner = repl.Main
	default:
		if strings.HasSuffix(args[1], ".glisp") {
			runner = runFile
		}
	}

	runner()
}

func runFile() {
	filename := os.Args[1]
	bts, err := ioutil.ReadFile(filename)

	if err != nil {
		exit(err.Error())
	}

	scn := scanner.New(strings.Trim(string(bts), "\n"))
	prs := parser.New(scn)

	prg, err := prs.Parse()
	if err != nil {
		exit(err.Error())
	}
	_, err = interpreter.Interpret(prg)
	if err != nil {
		exit(err.Error())
	}
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(-1)
}

func usage() {
	fmt.Println("glisp <file>.glisp | glisp --run")
	os.Exit(0)
}
