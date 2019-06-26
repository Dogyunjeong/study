# Server

request comes in, decide what to do.

clinet - request - server - response

certain url determine what specified job

1. Synonymous
- url
- multiplex
- router
- mux
- servermux
- request router
- http
- http router
- http request router

2. osi model
  - tcp:
  - http: http is protocol written by IETF. http works on tcp

## Request and Response

### HTTP request
 http is protocol written by IETF

- request line
- headers
- optional message body

Request-line

### HTTP response
- status line
- headers
- optional message body

## TCP Server -Write to Connection

When request comes in, request is just text. HTTP is a rule, protocol which parse request. So request must be formatted as HTTP

### net/http package
http package provides HTTP client and server implementations