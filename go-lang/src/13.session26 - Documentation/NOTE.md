#
Go can easily generate documentation from the comment in code.
The comment will turn into your documentation

## Writing documentation
Documentation is a huge part of making software accessible and maintainable. Of course it must be well-written and accurate, but it also must be easy to write and to maintain. Ideally, it should be coupled to the code itself so the documentation evolves along with the code. The easier it is for programmers to produce good documentation, the better for everyone.

### Best example
Look at `godoc.org`'s pakcage documentation and look the package code

## Go command
- go doc
- godoc
- godoc.org


### Run Documentation server
```
godoc -http :8080
```
It include godoc.org documentation and your own codes doc

### See specific package
```
$go doc fmt
$ go doc fmt.Println

```


## `godoc`
go doc prints the documentation for a package, const, func, type, var, or method


## add a repo at godoc.org
- searching package at repo
- If not able to find, searching with package git hub url
- Can search with package name

