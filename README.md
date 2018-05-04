## Glisp - a small Lisp implementation in Go

### Examples
#### Value declaration
```lisp
(defval i 64 "just an int equal to 64")
(print i) // 64
```
#### Function declaration
```lisp
(defun main (args)
    "the main function"
    (print (len args)))
(main "one" "two") // 2
```
#### Function call
```lisp
(print (* 5 5)) // 25
```

### Progress

#### Scanner
- [x] Basic expressions
- [ ] Lists
- [ ] Variable expressions
- [ ] Macro expressions
- [ ] Modules & imports
#### Parser
- [x] Atom expressions like Int, String, Float, Rune
- [ ] Lists
- [ ] Macro expressions
- [ ] Modules & imports
#### Evaluation
##### Interpreter
- [x] Simple expressions & internal functions
- [x] REPL
- [ ] Functions and macros
- [ ] Modules & imports
##### LLVM-based compiler
- [ ] Starting...