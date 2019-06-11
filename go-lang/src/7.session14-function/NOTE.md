# Function

Package and function allow to do modulus programming
Decoupling code
```
func (r receiver) identifier(parameters) (returns(s)) { ... }
```
### Language specification
A function type denotes the set of all functions with the same parameter and result types. The value of an uninitialized variable of function type is nil.

```
FunctionType   = "func" Signature .
Signature      = Parameters [ Result ] .
Result         = Parameters | Type .
Parameters     = "(" [ ParameterList [ "," ] ] ")" .
ParameterList  = ParameterDecl { "," ParameterDecl } .
ParameterDecl  = [ IdentifierList ] [ "..." ] Type .
```

## Everything in Go is PASS BY VALUE

## Variadic parameter [#](https://golang.org/ref/spec#Passing_arguments_to_..._parameters)

```
func variadicFoo(x ...int) { // x is slice of int
```

- has to be the final parameter
- parameter will be slice of declared type
- `...` is lexical symbol. ref: lexical elements
  - 0 or more values


## Unfurling a slice

```
x := foo(s...)

func foo(x ...string)
```

- It will be same underlying array to x as variadic parameter


## defer
We can defer closing file after open file.
As deferred code will be ran end of the block.


## Methods

```
type person struct {
  first string
  last  string
}

type secretAgent struct {
  person
  ltk bool
}

// func (r receiver) identifier(parameters) (returns(s)) {...}
// when there is receiver, the function will be attached to receiver

func (s secretAgent) speak() {
  fmt.Println("I am", s.first, s.last)
}
```

## Interfaces & polymorphism

Interface is for polymorphism

### polymorphism [#](https://en.wikipedia.org/wiki/Polymorphism_(computer_science))
In programming languages and type theory, polymorphism is the provision of a single interface to entities of different types[1] or the use of a single symbol to represent multiple different types.[


- classically: Interface allows us behave
- a value can be more than one type



### conversion

```
	var x5 hotdog = 42
	var y5 int
	y5 = int(x5)
```

### assertion [#](https://en.wikipedia.org/wiki/Assertion_(software_development))
Find concrete type underlying on it.
```
h.(person).first)
```

### composition [#](https://www.ardanlabs.com/blog/2015/09/composition-with-go.html)


argument
parameter
fmt.sprint


