# Control flow [#](https://en.wikipedia.org/wiki/Control_flow)

  In computer science, control flow (or flow of control) is the order in which individual [**statements**](https://golang.org/ref/spec#Statements), **instructions** or **function** calls of an imperative program are executed or evaluated. The emphasis on explicit control flow distinguishes an imperative programming language from a declarative programming language.
- how to computer read your program
- sequentially until hit `loop-iterative` and `conditional`

## 1. Statements
### 1) Label [#](https://en.wikipedia.org/wiki/Control_flow#labels)
is it control flow?

<!-- go by example -->
### 2) Loop [#](https://en.wikipedia.org/wiki/Control_flow#Loops)
**Go has only for loop**

  A loop is a sequence of statements which is specified once but which may be carried out several times in succession. The code "inside" the loop (the body of the loop, shown below as xxx) is obeyed a specified number of times, or once for each of a collection of items, or until some condition is met, or indefinitely.


#### For [#](https://golang.org/ref/spec#For_statements)
A "for" statement specifies repeated execution of a block. There are three forms: The iteration may be controlled by a single condition, a "for" clause, or a "range" clause.
```
ForStmt = "for" [ Condition | ForClause | RangeClause ] Block .
Condition = Expression 
```
- The Go for loop is similar to—but not the same as—C's. It unifies for and while and there is no do-while. There are three forms, only one of which has semicolons.

  ```
  // Like a C for
  for init; condition; post { }

  // Like a C while
  for condition { }

  // Like a C for(;;)
  for { }
  ```

- with range

  ```
  s := `"Hello, Play ground`
  for i, v := range s {
    fmt.Printf("at index position %d we have hex %#x\n", i, v) // --> such like "at index position 9 we have hex 0x6c"
  }
  ```

- with if statement like while

  ```
  y := 1
  for {
    y++
    if y > 100 {
      break
    }
    // how about when y is 0. eg  0 / 2, 0%2
    if y%2 != 0 {
      continue
    }
    fmt.Println(y)
  }
  ```


### 3) Conditional

#### if statement
"If" statements specify the conditional execution of two branches according to the value of a boolean expression. If the expression evaluates to true, the "if" branch is executed, otherwise, if present, the "else" branch is executed.
```
IfStmt = "if" [ SimpleStmt ";" ] Expression Block [ "else" ( IfStmt | Block ) ] .
```
- bool
  - true : pre-declare constant
  - false


#### switch statement
"Switch" statements provide multi-way execution. An expression or type specifier is compared to the "cases" inside the "switch" to determine which branch to execute.
```
SwitchStmt = ExprSwitchStmt | TypeSwitchStmt .

ExprSwitchStmt = "switch" [ SimpleStmt ";" ] [ Expression ] "{" { ExprCaseClause } "}" .
ExprCaseClause = ExprSwitchCase ":" StatementList .
ExprSwitchCase = "case" ExpressionList | "default" .
```
- `fallthrough`: will allows keep run next `case` statement without check `condition`
- `default`: If there is no true `case` It will be ran
- multiple case: use with `,`
  ```
  case "Moneypenny", "Bond", "Do No":
  ```
- with `fallthrough`
  ```
  switch {  // without condition is equivalent to true
  case (3 == 3):
    fmt.Println("prints")
    fallthrough
  // it will be not reached without fallthrough
  case (4 == 4):
    fmt.Println("also true, does it print?")
    fallthrough
  case (7 == 9):  // will be ran due to fallthrough
    fmt.Println("not true1")
    fallthrough
  default: // will be ran due to fallthrough
    fmt.Println("this is default")
  }
  ```

#### select statement [#](https://golang.org/ref/spec#Select_statements)
A "select" statement chooses which of a set of possible send or receive operations will proceed. It looks similar to a "switch" statement but with the cases all referring to communication operations.
- eg
  ```
  select {
  case i1 = <-c1:
    print("received ", i1, " from c1\n")
  case c2 <- i2:
    print("sent ", i2, " to c2\n")
  case i3, ok := (<-c3):  // same as: i3, ok := <-c3
    if ok {
      print("received ", i3, " from c3\n")
    } else {
      print("c3 is closed\n")
    }
  case a[f()] = <-c4:
    // same as:
    // case t := <-c4
    //	a[f()] = t
  default:
    print("no communication\n")
  }
  ```

### 4) Break
A "break" statement terminates execution of the innermost "for", "switch", or "select" statement within the same function.

### 5) Continue
A "continue" statement begins the next iteration of the innermost "for" loop at its post statement. The "for" loop must be within the same function.


### 6) All statements in Go [#](https://golang.org/ref/spec#Statements)
```
Statement =
	Declaration | LabeledStmt | SimpleStmt |
	GoStmt | ReturnStmt | BreakStmt | ContinueStmt | GotoStmt |
	FallthroughStmt | Block | IfStmt | SwitchStmt | SelectStmt | ForStmt |
	DeferStmt .

SimpleStmt = EmptyStmt | ExpressionStmt | SendStmt | IncDecStmt | Assignment | ShortVarDecl .
```

### 7) Terminating Statement [#](https://golang.org/ref/spec#Terminating_statements)
A terminating statement prevents execution of all statements that lexically appear after it in the same block. The following statements are terminating:

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
