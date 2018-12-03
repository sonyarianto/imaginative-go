
# URL router a.k.a HTTP request multiplexer

## Scenario

It will show how to do URL routing using several multiplexer (mux) libraries. Note that this only show the simple feature of routing on each mux library but we think you already can use it for any purpose.

For more detail of each mux library sample, maybe we will put it on separate article.

### Sample 1
Using default mux of `net/http`.

Save as `web-routing-default-mux.go`

```go
package main

import (
    "io"
    "log"
    "net/http"
)

func serviceDefault(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/" {
        http.NotFound(w, r)
        return
    }

    io.WriteString(w, "this is the / path")
}

func service1(w http.ResponseWriter, r *http.Request) {
    io.WriteString(w, "this is the /service1 path")
}

func service2(w http.ResponseWriter, r *http.Request) {
    io.WriteString(w, "this is the /service2 path")
}

func main() {
    http.HandleFunc("/", serviceDefault)
    http.HandleFunc("/service1", service1)
    http.HandleFunc("/service2", service2)

    log.Fatal(http.ListenAndServe(":3000", nil))
}
```

Run with `go run web-routing-default-mux.go` and access on your browser at `http://localhost:3000`

### Sample 2
Using `gorilla/mux`. 

Save as `web-routing-gorilla-mux.go`

```go
package main

import (
    "fmt"
    "github.com/gorilla/mux"
    "log"
    "net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "This is the home page")
}

func Article(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)

    fmt.Fprintf(w, "Article id is %s", vars["id"])
}

func Number(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)

    fmt.Fprintf(w, "This is number %s", vars["id"])
}

func main() {
    r := mux.NewRouter()

    r.HandleFunc("/", Home)
    r.HandleFunc("/article/{id}", Article)      // {id} can be anything
    r.HandleFunc("/number/{id:[0-9]+}", Number) // only match numeric pattern

    log.Fatal(http.ListenAndServe(":3000", r))
}
```

Run with `go run web-routing-gorilla-mux.go` and access on your browser at `http://localhost:3000`.

More detail at [gorilla/mux](https://github.com/gorilla/mux) website.

### Sample 3
Using `httprouter` mux.

Save as `web-routing-httprouter-mux.go`

```go
package main

import (
    "fmt"
    "github.com/julienschmidt/httprouter"
    "log"
    "net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    fmt.Fprint(w, "This is path / handled by httprouter!\n")
}

func numberHandler(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
    number := params.ByName("number")

    w.Header().Set("Content-Type", "text/html; charset=utf-8")

    fmt.Fprint(w, "This is number <strong>"+number+"</strong>")
}

func main() {
    router := httprouter.New()

    router.GET("/", homeHandler)
    router.GET("/this/:number", numberHandler)

    log.Fatal(http.ListenAndServe(":3000", router))
}
```

Run with `go run web-routing-httprouter-mux.go` and access on your browser at `http://localhost:3000`.

More detail at [httprouter](https://github.com/julienschmidt/httprouter) website.

### Sample 4
Using `Goji` mux.

Save as `web-routing-goji-mux.go`

```go
package main

import (
    "fmt"
    "goji.io"
    "goji.io/pat"
    "net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "This is the %s!", "homepage")
}

func hello(w http.ResponseWriter, r *http.Request) {
    name := pat.Param(r, "name")
    fmt.Fprintf(w, "Hello, %s!", name)
}

func main() {
    mux := goji.NewMux()

    mux.HandleFunc(pat.Get("/"), home)
    mux.HandleFunc(pat.Get("/hello/:name"), hello)

    http.ListenAndServe("localhost:3000", mux)
}
```

Run with `go run web-routing-goji-mux.go` and access on your browser at `http://localhost:3000`.

More detail at [Goji](https://github.com/goji/goji) website.