# Handler [# goDoc](https://godoc.org/net/http#Handler)

Implementing with interface.
So Passing type is important

# Request [# goDoc](https://godoc.org/net/http#Request)

## ParseForm [# goDoc](https://godoc.org/net/http#Request.ParseForm)
ParseForm populates r.Form and r.PostForm.
(r.Form and r.PostForm is available after invocing ParseForm)
For all requests, ParseForm parses the raw query from the URL and updates r.Form.

## FormValue [# goDoc](https://godoc.org/net/http#Request.FormValue)
FormValue returns the first value for the named component of the query. POST and PUT body parameters take precedence over URL query string values.

# etc
### go intterface
https://medium.com/golangspec/interfaces-in-go-part-i-4ae53a97479c

### Go variable naming
as shot as possible, clear, concise