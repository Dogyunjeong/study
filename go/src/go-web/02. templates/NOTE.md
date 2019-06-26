# text/template [#](https://godoc.org/text/template)

# Template
os.Create()  io.Copy   string.NewReader

*template is container which holds all the templates

parse: read various files and bring it into our program

### gohtml
Traditionally used. We can use others as well

## use variable in gohtml
1. variable
```
{{ . }}
```
2. loop for data.
```
{{range $index, $element := .}}
  <li>{{$index}} - {{$element}}</li>
{{end}}
```

## composition
1. in `.go`
pass structure e.g
```
  type person struct {
    Name string
    Age  int
  }

	p1 := person{
		Name: "James Bond",
		Age:  42,
	}
```
2. in `.gohtml`
```
<h1>{{.Name}}</h1>
<h2>{{.Age}}</h2>
```

## Use Function and pipeline
There are predefined global functions which can use without declaring with `.Funcs`

1. Use function
in `.go` file
```
<!-- pass fuctions to function map, so template know it -->
func(t *Template) Funcs(funcMap FuncMap)

template.New("").Funcs(template.FuncMap{"uc": strings.ToUpper}).parseFiles()
```
in `.gohtml`
```
{{uc .Name}}
```

2. pipeline
```
{{.Name | uc | something }}
```


## Using methods
can use methods in `.gohtml`
1. in `.go`
```
type person struct {
	Name string
	Age  int
}

func (p person) SomeProcessing() int {
	return 7
}

func (p person) TakesArg(x int) int {
	return x * 2
}

```
2. in `.gohtml`
```
<h3>{{.SomeProcessing}}</h3>
<h3>{{.AgeDbl | .TakesArg}}</h3>
```

## Nested Template
based on what you define not file name
1. in nested template
   - file name is not matter
anything.gothml
```
{{define "polarbear"}}
Something
{{end}}
```
2. upper template
index.gohtml
```
{{template "polarbear"}}
```

#  ETC
What is difference between text/template and html/template [#](https://code.tutsplus.com/tutorials/text-generation-with-go-templates--cms-30441)