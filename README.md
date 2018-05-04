## Glisp - a small Lisp implementation in Go

### Example
```lisp
(print (* 5 5 5) 24.5 "Hello World")
```

### Progress

#### Scanner
- [x] Basic expressions
- [ ] Macro expressions
- [ ] Modules & imports
#### Parser
- [x] Atom expressions like Int, String, Float, Rune
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