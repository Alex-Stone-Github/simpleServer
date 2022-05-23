package main

import (
    "fmt"
    "net/http"
)

func main() {
    fmt.Println("Hello, World!")
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("<b> Hello, World! </b>"));
    })
    fmt.Println(":5500")
    http.ListenAndServe(":5500", nil)
}
