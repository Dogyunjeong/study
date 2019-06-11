# JSON
## Marshal[#](https://godoc.org/encoding/json#example-Marshal)
Marshal returns the JSON encoding of v.


```
type ColorGroup struct {
    ID     int
    Name   string
    Colors []string
}
group := ColorGroup{
    ID:     1,
    Name:   "Reds",
    Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
}
b, err := json.Marshal(group)
if err != nil {
    fmt.Println("error:", err)
}
os.Stdout.Write(b)
```

Output
```
{"ID":1,"Name":"Reds","Colors":["Crimson","Red","Ruby","Maroon"]}
```

 - [palyground] (https://play.golang.org/p/8Bh4yaMP_pB)

## Unmarshal[#](https://godoc.org/encoding/json#example-unMarshal)
Unmarshal parses the JSON-encoded data and stores the result in the value pointed to by v. If v is nil or not a pointer, Unmarshal returns an InvalidUnmarshalError.
```
var jsonBlob = []byte(`[
	{"Name": "Platypus", "Order": "Monotremata"},
	{"Name": "Quoll",    "Order": "Dasyuromorphia"}
]`)
type Animal struct {
    Name  string
    Order string
}
var animals []Animal
err := json.Unmarshal(jsonBlob, &animals)
if err != nil {
    fmt.Println("error:", err)
}
fmt.Printf("%+v", animals)
```
Output
```
[{Name:Platypus Order:Monotremata} {Name:Quoll Order:Dasyuromorphia}]
```

  - [palyground](https://play.golang.org/p/O2AVaI8_6cQ)


# Interface
Understand how to read and use standard library
## Write Interface
 - [palyground](https://play.golang.org/p/uxHcTZcX9sx)


# Sort
 - [playground](https://play.golang.org/p/nogd5ZCix02)