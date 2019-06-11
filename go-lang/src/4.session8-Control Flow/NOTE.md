# Control flow [#](https://en.wikipedia.org/wiki/Control_flow)

  In computer science, control flow (or flow of control) is the order in which individual [**statements**](https://golang.org/ref/spec#Statements), **instructions** or **function** calls of an imperative program are executed or evaluated. The emphasis on explicit control flow distinguishes an imperative programming language from a declarative programming language.
- how to computer read your program
- sequentially until hit `loop-iterative` and `conditional`

## 1. [Statements](Statements.md)
- Label [#](https://en.wikipedia.org/wiki/Control_flow#labels)
  - goto

<!-- go by example -->
- Loop [#](https://en.wikipedia.org/wiki/Control_flow#Loops) **Go has only for loop**
  - For [#](https://golang.org/ref/spec#For_statements)
- Conditional
  - if statement
  - switch statement
  - select statement [#](https://golang.org/ref/spec#Select_statements)
- Break
- Continue
- All statements in Go [#](https://golang.org/ref/spec#Statements)
- Terminating Statement [#](https://golang.org/ref/spec#Terminating_statements)

## 2. Instructions [#](https://en.wikipedia.org/wiki/Instruction_set_architecture#Instructions)
Machine language is built up from discrete statements or instructions. On the processing architecture, a given instruction may specify:

- particular registers for arithmetic, addressing, or control functions
- particular memory locations or offsets
- particular addressing modes used to interpret the operands

More complex operations are built up by combining these simple instructions, which are executed sequentially, or as otherwise directed by control flow instructions.

### 1) Data handling and memory operations

```
Receive operator
```

#### Address operators [#](https://golang.org/ref/spec#Address_operators)
For an operand x of type T, the address operation &x generates a pointer of type *T to x. The operand must be addressable, that is, either a variable, pointer indirection, or slice indexing operation; or a field selector of an addressable struct operand; or an array indexing operation of an addressable array.
```
&x
&a[f(2)]
&Point{2, 3}
*p
*pf(x)

var x *int = nil
*x   // causes a run-time panic
&*x  // causes a run-time panic
```

### 2) Arithmetic and logic operations

- Add, subtract, multiply, or divide the values of two registers, placing the result in a register, possibly setting one or more condition codes in a status register.
- increment, decrement in some ISAs, saving operand fetch in trivial cases.
- Perform bitwise operations, e.g., taking the conjunction and disjunction of corresponding bits in a pair of registers, taking the negation of each bit in a register.
- Compare two values in registers (for example, to see if one is less, or if they are equal).
- Floating-point instructions for arithmetic on floating-point numbers.
```
Comparison operators, Recieve operator, Arithmetic operators
```

#### Conditional logic operator

Logical operators apply to boolean values and yield a result of the same type as the operands. The right operand is evaluated conditionally.
```
&&    conditional AND    p && q  is  "if p then q else false"
||    conditional OR     p || q  is  "if p then true else q"
!     NOT                !p      is  "not p"
```

### 3) Control flow operations

- Branch to another location in the program and execute instructions there.
- Conditionally branch to another location if a certain condition holds.
- Indirectly branch to another location.
- Call another block of code, while saving the location of the next instruction as a point to return to.

### 4) Coprocessor instructions

- Load/store data to and from a coprocessor, or exchanging with CPU registers.
- Perform coprocessor operations

## 3. Function calls [#](https://en.wikipedia.org/wiki/Subroutine)
In computer programming, a subroutine is a sequence of program instructions that performs a specific task, packaged as a unit. This unit can then be used in programs wherever that particular task should be performed.

- go routine ?


# ETC


<!-- go lang specification -->

## EBNF (Extended Backus–Naur form) [#](https://en.wikipedia.org/wiki/Extended_Backus%E2%80%93Naur_form)
In computer science, extended Backus-Naur form (EBNF) is a family of metasyntax notations, any of which can be used to express a context-free grammar. EBNF is used to make a formal description of a formal language which can be a computer programming language. They are extensions of the basic Backus–Naur form (BNF) metasyntax notation.
```
ForStmt = "for" [ Condition | ForClause | RangeClause ] Block .
Condition = Expression .
```
