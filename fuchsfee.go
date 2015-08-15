package main

import (
    "fmt"
    "./couchdbrest"
    "./jsonconvert"
)




func main() {
    fmt.Printf("Hello, world.\n")

    body, err := couchdbrest.GetLastBlogArticles("127.0.0.1", "10")
    fmt.Println("err: ", err)
    article_list := jsonconvert.JsonToObject(body)
    fmt.Println("============================\n")
    fmt.Println(article_list.Rows[0].Value["title"])
    fmt.Println("============================\n")
    fmt.Println("Anzahl Einträge:", len(article_list.Rows))

}
