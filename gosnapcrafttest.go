package main

import (
    "fmt"
    "log"
    "net/http"
    "github.com/golang/glog"
)

func main() {
    glog.Info("Starting webserver on :8081")
    http.HandleFunc("/", handleMainPage)

    fmt.Println("hello on stdout for new app!")
    if err := http.ListenAndServe(":8081", nil); err != nil {
        log.Fatal("http.ListendAndServer() failed with %s\n", err)
    }
}

func handleMainPage(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/" {
        log.Println("Invalid path reached by the user: r.URL.Path")
        http.NotFound(w, r)
        return
    }

    fmt.Fprintf(w, "Hello from snappy\n")
}
