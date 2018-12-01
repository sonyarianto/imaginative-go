package routing_with_default_mux

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