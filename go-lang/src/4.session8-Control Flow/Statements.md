# Statements [#](https://golang.org/ref/spec#Statements)
## 1) Label [#](https://en.wikipedia.org/wiki/Control_flow#labels)
is it control flow?

<!-- go by example -->
## 2) Loop [#](https://en.wikipedia.org/wiki/Control_flow#Loops)
**Go has only for loop**

  A loop is a sequence of statements which is specified once but which may be carried out several times in succession. The code "inside" the loop (the body of the loop, shown below as xxx) is obeyed a specified number of times, or once for each of a collection of items, or until some condition is met, or indefinitely.


### For [#](https://golang.org/ref/spec#For_statements)
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
## 3) Conditional

### if statement
"If" statements specify the conditional execution of two branches according to the value of a boolean expression. If the expression evaluates to true, the "if" branch is executed, otherwise, if present, the "else" branch is executed.
```
IfStmt = "if" [ SimpleStmt ";" ] Expression Block [ "else" ( IfStmt | Block ) ] .
```
- bool
  - true : pre-declare constant
  - false


### switch statement
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

### select statement [#](https://golang.org/ref/spec#Select_statements)
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

## 4) Break
A "break" statement terminates execution of the innermost "for", "switch", or "select" statement within the same function.

## 5) Continue
A "continue" statement begins the next iteration of the innermost "for" loop at its post statement. The "for" loop must be within the same function.


## 6) All statements in Go [#](https://golang.org/ref/spec#Statements)
```
Statement =
	Declaration | LabeledStmt | SimpleStmt |
	GoStmt | ReturnStmt | BreakStmt | ContinueStmt | GotoStmt |
	FallthroughStmt | Block | IfStmt | SwitchStmt | SelectStmt | ForStmt |
	DeferStmt .

SimpleStmt = EmptyStmt | ExpressionStmt | SendStmt | IncDecStmt | Assignment | ShortVarDecl .
```

## 7) Terminating Statement [#](https://golang.org/ref/spec#Terminating_statements)
A terminating statement prevents execution of all statements that lexically appear after it in the same block. The following statements are terminating:
