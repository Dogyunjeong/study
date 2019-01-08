# Error Handling
Can use any way of log package, panic, println

## Why go does not have exceptions? [#](https://golang.org/doc/faq#exceptions)
We believe that coupling exceptions to a control structure, as in the try-catch-finally idiom, results in convoluted code. It also tends to encourage programmers to label too many ordinary errors, such as failing to open a file, as exceptional.

Go takes a different approach. For plain error handling, Go's multi-value returns make it easy to report an error without overloading the return value. A canonical error type, coupled with Go's other features, makes error handling pleasant but quite different from that in other languages.

### Example [#](https://blog.golang.org/error-handling-and-go)
If you have written any Go code you have probably encountered the built-in error type. Go code uses error values to indicate an abnormal state. For example, the os.Open function returns a non-nil error value when it fails to open a file.

```
func Open(name string) (file *File, err error)
```

## Error
Error is type in go[#](https://godoc.org/builtin#error). Error interface determine error as a type.

- How it is implemented [#](https://godoc.org/errors)


## How to Printing & logging
message
- fmt.Println()
- log.Println()
- log.Fatalln()
 - os.Exit()
- log.Panicln()
 - deferred functions run
 - can use “recover”


### Func Exit [#](https://godoc.org/os#Exit)
 deferred functions are not run.
 Exit causes the current program to exit with the given status code. Conventionally, code zero indicates success, non-zero an error. The program terminates immediately;

### func panic [#](https://godoc.org/builtin#panic)
Read with `recover`

The panic built-in function stops normal execution of the current goroutine. When a function F calls panic, normal execution of F stops immediately. Any functions whose execution was deferred by F are run in the usual way, and then F returns to its caller. To the caller G, the invocation of F then behaves like a call to panic, terminating G's execution and running any deferred functions. This continues until all functions in the executing goroutine have stopped, in reverse order. At that point, the program is terminated and the error condition is reported, including the value of the argument to panic. This termination sequence is called panicking and can be controlled by the built-in function recover.

### func recover [#](https://godoc.org/builtin#panic)

## Defer, Panic, and Recover [#](https://blog.golang.org/defer-panic-and-recover)

Go has the usual mechanisms for control flow: if, for, switch, goto. It also has the go statement to run code in a separate goroutine. Here I'd like to discuss some of the less common ones: defer, panic, and recover.


### Defer
A defer statement pushes a function call onto a list. The list of saved calls is executed after the surrounding function returns. Defer is commonly used to simplify functions that perform various clean-up actions.
1. A deferred function's arguments are evaluated when the defer statement is evaluated.
2. Deferred function calls are executed in Last In First Out order after the surrounding function returns.
3. Deferred functions may read and assign to the returning function's named return values.
```
func c() (i int) {
    defer func() { i++ }()
    return 1 // will be 2
}
```

### Recover
 Recover is only useful inside deferred functions. 


## Log pakcage

has more flexibility