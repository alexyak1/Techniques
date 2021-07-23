package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
       fmt.Fprintf(w, "Here you will see judo techniques, %q", html.EscapeString(r.URL.Path))
    })

    http.HandleFunc("/list", func(w http.ResponseWriter, r *http.Request){
        fmt.Fprintf(w, "First techniq /n Second techniq")
    })

    log.Fatal(http.ListenAndServe(":8787", nil))
}
