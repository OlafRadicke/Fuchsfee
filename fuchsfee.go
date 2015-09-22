package main

import (
    "fmt"
    "net/http"
    "./feedfactory"
    "./appconfig"
)


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
    config := appconfig.AppConfig{
      BlogProtocol:"https://",
      BlogDomain:"the-independent-friend.de",
      FeedPort:"8088",
      Author: "Olaf Radicke",
      AuthorMail: "briefkasten@olaf-radicke.de",
    }

    feedfactory.GetAtom(config)
    http.HandleFunc("/new.xml", config.NewHandler)
    http.HandleFunc("/atom.xml", atomHandler)
    http.HandleFunc("/rss.xml", rssHandler)
    http.ListenAndServe(":" + config.FeedPort, nil)
}
