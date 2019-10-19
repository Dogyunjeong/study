# TEST

Tests must
- be in a file that ends with “_test.go”
- put the file in the same package as the one being tested
- be in a func with an signature “func TestXxx(*testing.T)”
Run a test
```
go test
go test -v  // with printing verbose
go test ./...  // run all test beneath of current path
```
Deal with test failure
- use t.Error to signal failure
Nice idiom
- expected
- got

## Example tests
Examples show up in documentation.
Ref: https://blog.golang.org/examples

## Coverage
Coverage in programming is how much of our code is covered by tests. We can use the “-cover” flag to run coverage analysis on our code. We can use the flag and required file name “-coverprofile <some file name>” to write our coverage analysis to a file.
```
go test -cover
```
```
go test -coverprofile c.out
// show in browser:
go tool cover -html=c.out
```
```
// learn more
go tool cover -h
```

## Refer Testing standard packages

## Linting

- gofmt: formats go code
- go vet: reports suspicious constructs
- golint: reports poor coding style
```
golint ./...
```

# Remember to BET
- Benchmark
- Example
- Test

```
BenchmarkCat(b *testing.B)
ExampleCat()
TestCat(t *testing.T)
```

# Commands
```
godoc -http=:8080

go test
go test -bench .
go test -cover
go test -coverprofile c.out
go tool cover -html=c.out
```

# Bench Marching
