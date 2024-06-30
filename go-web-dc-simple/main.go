package main

import (
    "fmt"
    "net/http"
    "path/filepath"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        http.ServeFile(w, r, filepath.Join("main", "index.html"))
    })

    http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello, World")
    })

    http.HandleFunc("/sawadde", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "สวัสดีชาวโลก !!")
    })

    http.HandleFunc("/about/", func(w http.ResponseWriter, r *http.Request) {
        http.ServeFile(w, r, filepath.Join("about", "index.html"))
    })

    http.Handle("/about/pic/", http.StripPrefix("/about/pic/", http.FileServer(http.Dir("./about/pic"))))

    fmt.Println("Starting server at port 8000")
    if err := http.ListenAndServe(":8000", nil); err != nil {
        fmt.Println(err)
    }
}
