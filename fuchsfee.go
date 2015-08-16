package main

import (
    "fmt"
    "net/http"
    "./couchdbrest"
    "./jsonconvert"
    "./feedfactory"
)





func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func main() {
    fmt.Printf("Hello, world.\n")

    body, err := couchdbrest.GetLastBlogArticles("127.0.0.1", "10")
    fmt.Println("err: ", err)
    article_list := jsonconvert.JsonToObject(body)
    fmt.Println("============================\n")
    fmt.Println(article_list.Rows[0].Value["title"])
    fmt.Println("============================\n")
    fmt.Println("Anzahl Einträge:", len(article_list.Rows))


    feedfactory.GetAtom()


    http.HandleFunc("/", handler)
    http.ListenAndServe(":8088", nil)
}
