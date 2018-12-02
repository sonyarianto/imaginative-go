package hello_world_on_web

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