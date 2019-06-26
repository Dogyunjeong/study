## Whats GRPC [#](./whatsGRPC.md)
## Whats ProtocolBuffer [#](./whatsPB.md)
## Whats ProtocolBuffer [#](./deepDiveGRPC.md)

## HTTP2
- is binary
- can push messages in parallel over the same TCP
- support header compression
### HTTP/1.1 works
- release 1997
- opens a new TCP connections to a server at each request
- Nowadays a web page load 80 assets on average
- No header compressed
So Http1.1 need open 80 TCP conncections and close with 80 plain text head


## gRPC
- support for streaming APIs
- gRPC is API oriented, instead of resource oriented like REST
- uses HTTP/2
- Easy code definition in over 11 languages

### 4 Types of API in gRPC

#### Unary api
traditional API look like (HTTP REST)

### Scalability in gRPC
- gRPC severs are asynchronous by default
### Security in gRPC
- cross language security
- SSL security built in
- By default gRPC strongly advocates for you to use SSL
- can also provide authentication

### gRPC vs REST
#### REST smaple


### What's an Unary API
- Unary RPC calls are the basic Request/Response that everyon is familiar
