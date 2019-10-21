# POINTER
Pointer is pointing to memory in computer. Each memory block(?) has a memory address.

“As an exception, if the type of x is a named pointer type and (*x).f is a valid selector expression denoting a field (but not a method), x.f is shorthand for (*x).f.”

```
me.first = "Miss Moneypenny" // shorthand of (*me).first = "PaPaPa"
```

### Dereference *p
```
(*p).first = "hello"
```
*p is dereferencing

## Everything in go pass by value
 When you need address you can pass address by pass by value

### & operator
will show adress

```
  fmt.Println(&a)        // --> Like 0xc4200a4000: address of memory
	fmt.Printf("%T\n", a)  // --> int
	fmt.Printf("%T\n", &a) // --> *int
```

### * operator
**dereference** the address. It will make a variable next it to give an value of stored address's memory


```
  b := &a
	fmt.Println(*b)
	fmt.Println(*&a)

  *b =43  // a will be 43
```

## When should use pointer
A Pointer allows you to mutate a value
- When stored huge data and have to pass to use it.
- When have to change value at certain location. can pass address and change value with dereference

```
func main() {
  x := 1
  foo(&x)
  fmt.Println(x) // ---> 42
}
func foo(y *int) {  // param y is passed by value.  value of y is same as an address of x(&x).
  *y = 42
}
```

## Methods set [#](https://golang.org/ref/spec#Method_sets)
  Method sets determine what methods attach to a TYPE. it is exactly
We can attach methods to a type. Those methods attached to a type are known as a method set.

The method set of a type **determines the interfaces that the type implements** and the methods that can be....

[play ground](https://play.golang.org/p/pWFxsg6MSe)
```
type circle struct {
	radius float64
}

type shape interface {
	area() float64
}

func (c *circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func info(s shape) {
	fmt.Println("area", s.area())
}

func main() {
	c := circle{5}
	info(&c)
}

```

- NON-POINTER RECEIVER & NON-POINTER VALUE
  https://play.golang.org/p/2ZU0QX12a8
- NON-POINTER RECEIVER & POINTER VALUE
  https://play.golang.org/p/glWZmm0gY6 
- POINTER RECEIVER & POINTER VALUE
  https://play.golang.org/p/pWFxsg6MSe 
- POINTER RECEIVER & NON-POINTER VALUE
  https://play.golang.org/p/G3lEy-4Mc8 ( code does not run )
- this codes does run - notice the difference -  method set determines the INTERFACES that the type implements
  **Important:**  https://play.golang.org/p/KK8gjsAWBZ 

### Pointer & Non-Pointer receiver
Non-pointer receiver can receive pointers and non-pointers. But Pointer receiver can receive only pointer.

- a NON-POINTER RECEIVER
  works with values that are POINTERS or NON-POINTERS.
- a POINTER RECEIVER
  only works with values that are POINTERS.

|Receivers|Values|
|-----|-----|
|(t T)|T and *T|
|(t *T)| *T|


Mutate