package main

import (
    "fmt"
    "log"
    "net/http"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
    })

    // Check for errors from ListenAndServe
    if err := http.ListenAndServe(":80", nil); err != nil {
        log.Fatalf("Error starting server: %s", err)
    }
}
