package main

import (
	"fmt"
	"net/http"
)

func main() {
    fmt.Println("Hello World?")
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprint(w, "Hello World!")
    })

    err := http.ListenAndServe(":80", nil)

    if err != nil {
        fmt.Printf("Failed to start server: %s\n", err)
        return
    }

    fmt.Println("Listening on port 80")
}
