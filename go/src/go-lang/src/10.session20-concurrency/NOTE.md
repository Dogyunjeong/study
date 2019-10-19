# Concurrency

Go was built for advantage of multiple cores.
Go is first programming language, natively take advantage of multiple cores

Vidoe: Concurrency is not parlleism: [#](https://www.youtube.com/watch?v=cN_DpYBzKso)


## Concurrency  [#](https://golang.org/doc/effective_go.html#concurrency)

Design pattern, way to code

Concurrent programming in many environments is made difficult by the subtleties required to implement correct access to shared variables.

In computer science, concurrency refers to the ability of different parts or units of a program, algorithm, or problem to be executed out-of-order or in partial order, without affecting the final outcome. This allows for parallel execution of the concurrent units, which can significantly improve overall speed of the execution in multi-processor and multi-core systems. In more technical terms, concurrency refers to the decomposability property of a program, algorithm, or problem into order-independent or partially-ordered components or units [wikipedia](https://en.wikipedia.org/wiki/Concurrency_(computer_science))

### Concurrency vs Parallelism
concurrency is the composition of independently executing processes, while parallelism is the simultaneous execution of (possibly related) computations [concurrency-is-not-parallelism](https://blog.golang.org/concurrency-is-not-parallelism)


### Writing Concurrent code

Primitive concurrent code

```
var wg sync.WaitGroup
func main() {
  wg.Add(1)
  go foo()
  bar()
  wg.Wait()
}
func foo() {
  ...
  wg.Done()
}
func bar() { ... }
```

### go func()
It will launch a go routine
Can use to `function` or `method`. but some built-in function are restricted
When the function terminates, its goroutine alos terminates.
A function literal can be handy in a goroutine invocation

```
func main() {
  go foo()
  bar()
}
func foo() { ... }
func bar() { ... }
```

- main routine is one of go routine
- Muex [#](https://godoc.org/sync#Mutex)
- WaitGroup [#](https://godoc.org/sync#WaitGroup)


## Method sets revisited [#](https://golang.org/ref/spec#Method_sets)
The method set of a type **determines the interfaces that the type implements** and the methods that can be....
- check: https://play.golang.org/p/KK8gjsAWBZ 

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
	// info(c)  // not working as reciever is
	fmt.Println(c.area())
}
```

## Mutiplexed

## Race Condition

A race condition or race hazard is the behavior of an electronics, software, or other system where the output is dependent on the sequence or timing of other uncontrollable events. It becomes a bug when events do not happen in the order the programmer intended.[wikipedia](https://en.wikipedia.org/wiki/Race_condition)

When a light has two switches, two people turn on and off at each switches [#](https://searchstorage.techtarget.com/definition/race-condition)


| Process 1 | Process 2 | Memory Value |
| Read value | - | 0 |
| Flip value | - | 1 |
| - | Read value | 1 |
| - | Flip value | 0 |

| Process 1 | Process 2 | Memory Value |
| Read value | - | 0 |
| - |  Read value | 0 |
| Flip value | - | 1 |
| - | Flip value | 1 |

[palyground](https://play.golang.org/p/z9RLR5s6mTG)


## Mutex

Blocking variable from other go routine access when a goroutine use it

- Mutext: Lock read and write
- RWMutex: Has more flexibility can lock only read or write

```
	var mu sync.Mutex
	go func() {
		mu.Lock() // no one can access counter variable
		v := counter
		runtime.Gosched()
		v++
		counter = v
		mu.Unlock()
	}()
```


## Atomic

One of ways avoid `race condition`
Package atomic provides low-level atomic memory primitives useful for implementing synchronization algorithms.


## Additional

#### func init() {}
It will setup something before run main()

#### Gosched

Gosched yields the processor, allowing other goroutines to run. It does not suspend the current goroutine, so execution resumes automatically.
```
runtime.Gosched() // switch between go routine
```

#### runtime package
```
	fmt.Println("OS\t\t", runtime.GOOS)
	fmt.Println("ARCH\t\t", runtime.GOARCH)
	fmt.Println("CPUs\t\t", runtime.NumCPU())
	fmt.Println("Goroutines\t\t", runtime.NumGoroutine())
```

