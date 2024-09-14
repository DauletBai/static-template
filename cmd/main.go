package main

import (
    "log"
    "net/http"
)

func main() {
    mux := http.NewServeMux()

    fileServer := http.FileServer(http.Dir("./assets/"))

    mux.Handle("GET /assets/", http.StripPrefix("/assets", fileServer))

    mux.HandleFunc("GET /{$}", home)
    mux.HandleFunc("GET /post/view/{id}", postView)
    mux.HandleFunc("GET /post/create", postCreate)
    mux.HandleFunc("POST /post/create", postCreatePost)

    log.Print("starting server on :5000")
    
    err := http.ListenAndServe(":5000", mux)
    log.Fatal(err)
}