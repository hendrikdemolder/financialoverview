package hello

import (
    "fmt"
    "http"
)

func init() {
    http.HandleFunc("/api/helloworld", handler)
}

func httphandler(w http.ResponseWriter, r *http.Request){
    fmt.Fprint(w, r.RawURL)
}

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "Hello world!")
    fmt.Fprint(w, r.Method)
}
