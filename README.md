# kataribe playground

Playground for [`kataribe`](https://github.com/matsuu/kataribe)

## Build Server

```sh
docker-compose up
```

And listen on localhost:8080

## Send HTTP Requests

Edit `client/main.go`:

```go
const (
  concurrency = 4
  duration    = 10 * time.Second
  host        = "http://localhost:8080"
)

var requestLines = []request.RequestLine{
  request.NewRequestLine("GET", host, "/"),
  request.NewRequestLine("GET", host, "/users"),
  request.NewRequestLine("GET", host, "/users/:id"),
}
```

And send HTTP requests:

```sh
cd client && make
```

## Play `kataribe`

```sh
cat logs/access.log | kataribe -f kataribe.toml
```
