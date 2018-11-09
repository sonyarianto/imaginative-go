package main

import (
    "io"
    "log"
    "net/http"
    "github.com/julienschmidt/httprouter"
    "github.com/urfave/negroni"
    muxg "github.com/gorilla/mux"
)

func one(w http.ResponseWriter, r *http.Request) {
    io.WriteString(w, "one")
}

func main() {
    muxNetHttp := http.NewServeMux()
    muxHttpRouter := httprouter.New()
    muxGorilla := muxg.NewRouter()
    
    n := negroni.Classic()
    
    // registers the handler function for the given pattern
    muxNetHttp.HandleFunc("/1", one)
    muxHttpRouter.HandleFunc("/2", two)
    
    log.Fatal(http.ListenAndServe(":9999", mux))
    log.Fatal(http.ListenAndServe(":10000", muxHttpRouter))
}