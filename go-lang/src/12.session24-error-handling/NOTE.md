# Error Handling


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