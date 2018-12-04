# Hello World

## Scenario

It will show how to do "hello, world" in Go. There is web and non-web sample for this.

### Sample 1

Display 'hello, world' on console.

Save as `hello-world.go`
```go
package main

import "fmt"

func main() {
    fmt.Println("hello, world")
}
```

Run with `go run hello-world.go`

### Sample 2

Display 'hello, world' on web browser. Using standard mux.

Save as `hello-world-on-web.go`

```go
package main

import (
    "io"
    "log"
    "net/http"
)

func Default(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/" {
        http.NotFound(w, r)
        return
    }

    io.WriteString(w, "hello, world")
}

func main() {
    http.HandleFunc("/", Default)
    
    log.Fatal(http.ListenAndServe(":3000", nil))
}
```

Run with `go run hello-world-on-web.go` and access on your browser at `http://localhost:3000`

### Sample 3

Display 'hello, world' on web browser. Using standard mux. Render with Content Type `text/html` so it aware of HTML tags.

Save as `hello-world-on-web-2.go`

```go
package main

import (
    "io"
    "log"
    "net/http"
)

func Default(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/" {
        http.NotFound(w, r)
        return
    }

    w.Header().Set("Content-Type", "text/html; charset=utf-8")

    io.WriteString(w, "<h1>hello, world</h1>")
}

func main() {
    http.HandleFunc("/", Default)
    
    log.Fatal(http.ListenAndServe(":3000", nil))
}
```

Run with `go run hello-world-on-web-2.go` and access on your browser at `http://localhost:3000`