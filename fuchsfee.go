package main

import (
    "fmt"
    "net/http"
    "./feedfactory"
)

type AppConfig struct {
   blogProtocol string
   blogDomain string
   feedPort string
}


func (ac AppConfig) newHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Printf("Blog domain %s\n", ac.blogDomain)
    // fmt.Fprintf(w, feedfactory.GetRss())
}

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func rssHandler(w http.ResponseWriter, r *http.Request) {
    // fmt.Fprintf(w, feedfactory.GetRss())
}

func atomHandler(w http.ResponseWriter, r *http.Request) {
    //fmt.Fprintf(w, feedfactory.GetAtom())
}

func main() {
    fmt.Printf("Hello, world.\n")
    config := AppConfig{blogProtocol:"https://", blogDomain:"the-independent-friend.de", feedPort:"8088"}

    feedfactory.GetAtom()
    http.HandleFunc("/new.xml", config.newHandler)
    http.HandleFunc("/atom.xml", atomHandler)
    http.HandleFunc("/rss.xml", rssHandler)
    http.ListenAndServe(":" + config.feedPort, nil)
}
