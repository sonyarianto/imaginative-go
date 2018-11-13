package main

import (
    "log"
    "io"
    "net/http"
    "github.com/julienschmidt/httprouter"
)

func helloWorld(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    w.Header().Set("Content-Type", "text/html; charset=utf-8")

    io.WriteString(w, "<h1>hello, world</h1>")
}

func main() {
    mux := httprouter.New()
    
    mux.GET("/hello-world", helloWorld)
    
    log.Fatal(http.ListenAndServe(":3000", mux))
}