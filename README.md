## Glisp - a small Lisp implementation in Go

### Example
```lisp
(mod main)

(import fmt:println)

(defun -fib (n)
    (if ((< n 2) n)
        (+ (-fib (- n 1)) (-fib (- n 2)))))

(defun -runfib (n) (fmt:println (fib n)))

(defun main (args)
    (def args-len (len args))
    (if ((< args-len 2) -usage args-len)
        (-runfib (get-index args 0))))
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