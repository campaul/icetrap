package main

import (
    "fmt"
    "net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "")
}

func main() {
    http.HandleFunc("/", homeHandler)
    fmt.Println("Listening on port 3000")
    http.ListenAndServe(":3000", nil)
}
