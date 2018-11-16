package main

import (
    "fmt"
    "log"
    "net/http"
    "github.com/julienschmidt/httprouter"
)

func homeHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    fmt.Fprint(w, "This is path / handled by httprouter!\n")
}

func numberHandler(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
    number := params.ByName("number")

    w.Header().Set("Content-Type", "text/html; charset=utf-8")

    fmt.Fprint(w, "This is number <strong>" + number + "</strong>")
}

func main() {
    router := httprouter.New()
    
    router.GET("/", homeHandler)
    router.GET("/this/:number", numberHandler)

    log.Fatal(http.ListenAndServe(":3000", router))
}