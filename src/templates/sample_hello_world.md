# The standard hello world ritual

## Scenario

It will show how to do hello world in web.

File: `src\examples\hello_world.go`

```go
package main

import (
    "log"
    "io"
    "net/http"
    "github.com/julienschmidt/httprouter"
)

func helloWorld(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    io.WriteString(w, "hello, world")
}

func main() {
    mux := httprouter.New()
    
    mux.GET("/hello-world", helloWorld)
    
    log.Fatal(http.ListenAndServe(":3000", mux))
}
```