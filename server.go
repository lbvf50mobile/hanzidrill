package main

import (
    "log"
    "net/http"
)

func main() {
    // Serve static/index.html immediately at "/"
    http.Handle("/", http.FileServer(http.Dir("./static")))
    
    // Serve the mp3 files from /data/
    http.Handle("/data/", http.StripPrefix("/data/", http.FileServer(http.Dir("./data"))))

    log.Println("Listening on :8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
