package main

import (
    "time"
    "fmt"
    "net/http"
    "github.com/gorilla/feeds"
    "./couchdbrest"
    "./jsonconvert"
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



/////////////////////////////////////////////////////////////////////////////

    now := time.Now()
    feed := &feeds.Feed{
        Title:       "jmoiron.net blog",
        Link:        &feeds.Link{Href: "http://jmoiron.net/blog"},
        Description: "discussion about tech, footie, photos",
        Author:      &feeds.Author{"Jason Moiron", "jmoiron@jmoiron.net"},
        Created:     now,
    }

    feed.Items = []*feeds.Item{
        &feeds.Item{
            Title:       "Limiting Concurrency in Go",
            Link:        &feeds.Link{Href: "http://jmoiron.net/blog/limiting-concurrency-in-go/"},
            Description: "A discussion on controlled parallelism in golang",
            Author:      &feeds.Author{"Jason Moiron", "jmoiron@jmoiron.net"},
            Created:     now,
        },
        &feeds.Item{
            Title:       "Logic-less Template Redux",
            Link:        &feeds.Link{Href: "http://jmoiron.net/blog/logicless-template-redux/"},
            Description: "More thoughts on logicless templates",
            Created:     now,
        },
        &feeds.Item{
            Title:       "Idiomatic Code Reuse in Go",
            Link:        &feeds.Link{Href: "http://jmoiron.net/blog/idiomatic-code-reuse-in-go/"},
            Description: "How to use interfaces <em>effectively</em>",
            Created:     now,
        },
    }

    atom, _ := feed.ToAtom()
    rss, _ := feed.ToRss()

    fmt.Println(atom, "\n", rss)

    fmt.Println("============================\n")
    fmt.Println("atom:", atom)
////////////////////////////////////////////////////////////////////////////



    http.HandleFunc("/", handler)
    http.ListenAndServe(":8088", nil)
}
