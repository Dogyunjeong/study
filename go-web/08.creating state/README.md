## Passing values

### Form Submission
1. thorugh body
2. thorugh url

### URL values
protocol
subdomain
domain
port
path
query
parameter
fragment

### Retrieving value


## Enctype
Form has enctype.

## Redirects
Can set response header manually or using HTTP.Redirect
### 301 move permanently
Browser will store this. on Second try, browser will not hit origin url and hit moved url
Preserve Method
### 303 see other
Always GET
### 307 tempolary move
Preserve Method


# Cookies [#](https://godoc.org/net/http#Cookie)
Domain specified
## Writie cookie on client
## Read cookie
## Controll cookie
with `maxAge` we can controll coookie. if It is  `<= 0` cookie will be deleted