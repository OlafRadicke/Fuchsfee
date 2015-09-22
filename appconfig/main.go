package appconfig

import (
    "fmt"
    "net/http"
)

type AppConfig struct {
   BlogProtocol string
   BlogDomain string
   FeedPort string
   Author string
   AuthorMail string
}


func (ac AppConfig) NewHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Printf("Blog domain %s\n", ac.BlogDomain)
    // fmt.Fprintf(w, feedfactory.GetRss())
}
