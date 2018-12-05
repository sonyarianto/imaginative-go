
# URL router a.k.a HTTP request multiplexer

## Scenario

It will show how to do URL routing using several multiplexer (mux) libraries. Note that this only show the simple feature of routing on each mux library but we think you already can use it for any purpose.

For more detail of each mux library sample, maybe we will put it on separate article.

### Sample 1 (using net/http)

Using default mux of `net/http`.

Save as `web-routing-default-mux.go`

<script src="https://gist.github.com/sonyarianto/45feb2da543ba038a6ae7413f496666e.js"></script>

Run with `go run web-routing-default-mux.go` and access on your browser at `http://localhost:3000`

### Sample 2 (using gorilla/mux)

Using `gorilla/mux`. More detail at [gorilla/mux](https://github.com/gorilla/mux) website.

First install the package.

```
go get -u github.com/gorilla/mux
```

Save as `web-routing-gorilla-mux.go`

<script src="https://gist.github.com/sonyarianto/47b2e2fd4c3103bf699c3c3b1b86040f.js"></script>

Run with `go run web-routing-gorilla-mux.go` and access on your browser at `http://localhost:3000`.

### Sample 3 (using httprouter)

Using `httprouter` mux. More detail at [httprouter](https://github.com/julienschmidt/httprouter) website.

First install the package.

```
go get -u github.com/julienschmidt/httprouter
```

Save as `web-routing-httprouter-mux.go`

```go
<script src="https://gist.github.com/sonyarianto/2e608ceea148e371f72f4bf5eca0f309.js"></script>
```

Run with `go run web-routing-httprouter-mux.go` and access on your browser at `http://localhost:3000`.

### Sample 4 (using Goji)

Using `Goji` mux. More detail at [Goji](https://github.com/goji/goji) website.

First install the package.

```
go get -u goji.io
```

Save as `web-routing-goji-mux.go`

<script src="https://gist.github.com/sonyarianto/a2cdf7f40e77ed44fe492c9cc1a2d306.js"></script>

Run with `go run web-routing-goji-mux.go` and access on your browser at `http://localhost:3000`.