# Struct
A struct is a sequence of named elements, called fields, each of which has a name and a type. Field names may be specified explicitly (IdentifierList) or implicitly (EmbeddedField). Within a struct, non-blank field names must be unique.

```
StructType    = "struct" "{" { FieldDecl ";" } "}" .
FieldDecl     = (IdentifierList Type | EmbeddedField) [ Tag ] .
EmbeddedField = [ "*" ] TypeName .
Tag           = string_lit .
```

Create a value of type
To allow to aggregate values different type together
Composite data structure
A struct is an composite data type. (composite data types, aka, aggregate data types, aka, complex data types). **Structs allow us to compose together values of different types.**

- Looks like `object`



## Embedded Struct

```
  type person struct {
    first string
    last  string
    age   int
  }

  type secretAgent struct {
    person
    ltk bool
  }

	sa1 := secretAgent{
		person: person{
			first: "James",
			last:  "Bond",
			age:   32,
		},
		ltk: true,
	}
	fmt.Println(sa1.first, sa1.last, sa1.age, sa1.ltk)
```

- `person` is inner type: inner type promote to outer type
- `secretAgent` is outer type

### Embedded Field
A field declared with a type but no explicit field name is called an embedded field. An embedded field must be specified as a type name T or as a pointer to a non-interface type name *T, and T itself may not be a pointer type. The unqualified type name acts as the field name.

### Promoted
A field or method f of an embedded field in a struct x is called promoted if x.f is a legal selector that denotes that field or method f.

Promoted fields act like ordinary fields of a struct except that they cannot be used as field names in composite literals of the struct.


### Name Collison
```
type person struct {
	first string
	last  string
	age   int
}

type secretAgent struct {
	person
  first string
	ltk bool
}

	fmt.Println(sa1.first, sa1.person.first, sa1.person.last, sa1.age, sa1.ltk)

```

## Anonymous Struct
To prevent code pollution
Struct without name and can use one time.
```
	// anonymous struct
	// As it doesn't have name
	p4 := struct {
		first string
		last  string
		age   int
	}{ // composite literal: {}
		first: "James",
		last:  "Bond",
		age:   32,
	}
	fmt.Println(p4)
```

## Go's 3 goal

- efficient compilation
- efficient execution
- ease programming


Should write code clear and readable

## IS go object oriented language?
YES or NO.
IT doesn't have type hierarchy 


constant of kind 
```
const a = 1
```