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