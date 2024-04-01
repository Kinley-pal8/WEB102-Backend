package main

import (
    "fmt"
    "log"
    "net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request) {
    if err := r.ParseForm(); err != nil {
        fmt.Fprintf(w, "ParseForm() err: %v", err)
        return
    }

    name := r.FormValue("name")
    password := r.FormValue("password")

    if name == "admin" && password == "admin" {
        // Redirect to the success page
        http.Redirect(w, r, "/success", http.StatusFound)
        return
    }

    // If the form values are not as expected, display an error message
    fmt.Fprintf(w, "Damn, it isn't right")
}

func successHandler(w http.ResponseWriter, r *http.Request) {
    // Render the success page
    fmt.Fprintf(w, "Success!")
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/hello" {
        http.Error(w, "404 not found.", http.StatusNotFound)
        return
    }

    if r.Method != "GET" {
        http.Error(w, "Method is not supported.", http.StatusNotFound)
        return
    }

    fmt.Fprintf(w, "Hello, World!")
}

func main() {
    fileServer := http.FileServer(http.Dir("./"))
    http.Handle("/", http.StripPrefix("/", fileServer))

    http.HandleFunc("/form", formHandler)
    http.HandleFunc("/success", successHandler)
    http.HandleFunc("/hello", helloHandler)

    log.Println("Starting server on :8082")
    err := http.ListenAndServe(":8082", nil)
    if err != nil {
        log.Fatal(err)
    }
}