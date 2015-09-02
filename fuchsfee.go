package main

import (
    "fmt"
    "net/http"
    "./feedfactory"
)


func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func rssHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, feedfactory.GetRss())
}

func atomHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, feedfactory.GetAtom())
}

func main() {
    fmt.Printf("Hello, world.\n")



    feedfactory.GetAtom()


    http.HandleFunc("/atom.xml", atomHandler)
    http.HandleFunc("/rss.xml", rssHandler)
    http.ListenAndServe(":8088", nil)
}
