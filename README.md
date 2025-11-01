
# Language

# References
- https://craftinginterpreters.com/contents.html
- https://youtu.be/JTjNoejn4iA?si=qmrCcRsNQRGvm4YK

# todo
1. implement
    Source -> [Tokenizer] -> AST
    AST -> [codegen] -> IR
    IR -> [QBE/LLVM] -> ASM
1. seperate keywords and identification as two
1. recursive descent parser

# The Language

- dynamically types, while supporting types

```
x :: Integer = 9
y = 10
print x + y                     // 19
```

```
fn add x y -> return x + y
print add 9 10                  // 19
print add add 9 10 10           // 29
```

```
num = 10
fn isodd num -> return num % 2
if isodd num > print "odd" || print "even" 
```

```
num = 3
if num == 1 > print "one" || if num == 2 > print "two" || print "three"
if num == 1 
    > print "one" 
    || if num == 2 
    > print "two" 
    || print "three";
```
*OR*
```
num = 3
if num == 1 then print "one" else if num == 2 then print "two" else print "three"
if num == 1 
    then print "one" 
    else if num == 2 
    then print "two" 
    else print "three"
fi
```

# my file structure
LangLang 
    cli/
        repl/
            repl.go // has package name repl
    internals/
        lexer/
            Lexer.go // has packagename lexer
            token.go // has packagename lexer
    main.go   // has package name LangLang

# my mod file
module LangLang

go 1.24.5

# Error
in repl/repl.go im getting error import (
	"LangLang/lexer" // could not import LangLang/lexer (no required module provides package "LangLang/lexer")compilerBrokenImport
)
