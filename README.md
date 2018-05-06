# Glisp - a small Lisp implementation in Go [WIP]
[![Build Status](https://travis-ci.org/pmukhin/glisp.svg?branch=master)](https://travis-ci.org/pmukhin/glisp)

## Motivation
Fun.

## Examples
### Value declaration
```lisp
(defvar i 64 "just an int equal to 64")
(print i) // 64
```
### Defining a collection
```lisp
(defvar int-list '(1 2 3)
    "a list of ints")  // list
(defvar int-vector [1 2 3]
    "a vector of ints") // vector
```
### Function declaration
```lisp
(defun main (args)
    "the main function"
    (print (len args)))
(main '("one" "two")) // 2
```
### Function call
```lisp
(print (* 5 5)) // 25
```

## Progress

### Scanner
- [x] Basic expressions
- [x] Lists
- [x] Vectors
- [ ] Variable expressions
- [ ] Macro expressions
- [ ] Modules & imports
### Parser
- [x] Atom expressions like Int, String, Float, Rune
- [x] Lists
- [x] Vectors
- [ ] Macro expressions
- [ ] Modules & imports
### Evaluation
#### Interpreter
- [x] Simple expressions & internal functions
- [x] REPL
- [ ] Functions and macros
- [ ] Modules & imports
#### LLVM-based compiler
- [ ] Starting...
