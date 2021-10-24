package main

import (
    "fmt"
    "net/http"
    "os"
    "strconv"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Entering Dungeon.\n")
}

func main() {
    port, _ := strconv.Atoi(os.Args[1])
    fmt.Printf("Starting server at Port %d", port)
    http.HandleFunc("/", handler)
    http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
