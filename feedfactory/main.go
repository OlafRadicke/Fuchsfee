package feedfactory

import (
    "time"
    "fmt"
    "github.com/gorilla/feeds"
    "../couchdbrest"
    "../jsonconvert"
)



func convertToFeed()(*feeds.Feed){


    body, err := couchdbrest.GetLastBlogArticles("127.0.0.1", "10")
    fmt.Println("err: ", err)
    article_list := jsonconvert.JsonToObject(body)
    fmt.Println("============================\n")
    fmt.Println(article_list.Rows[0].Value["title"])
    fmt.Println("============================\n")
    fmt.Println("Anzahl Einträge:", len(article_list.Rows))

    now := time.Now()
    feed := &feeds.Feed{
        Title:       "THE INDEPENDENT FRIEND",
        Link:        &feeds.Link{Href: "https://the-independent-friend.de/"},
        Description: "Weblog von Olaf Radicke",
        Author:      &feeds.Author{"Olaf Radicke", "briefkasten@olaf-radicke.de"},
        Created:     now,
    }

    feed.Items = []*feeds.Item{}
    for index,element := range article_list.Rows {
        fmt.Println("============================\n")
        fmt.Println( "No.:" , index ,  " Value: " , element.Value["title"] )
        fmt.Println("============================\n")

        article_title := element.Value["title"]
        feed.Items = append(
            feed.Items,
            &feeds.Item{
                Title:       article_title.(string),
                Link:        &feeds.Link{Href: "http://jmoiron.net/blog/limiting-concurrency-in-go/"},
                Description: "A discussion on controlled parallelism in golang",
                Author:      &feeds.Author{"Jason Moiron", "jmoiron@jmoiron.net"},
                Created:     now,
            } )
    }

    return feed
}

func GetAtom()(string){
    var feed = convertToFeed()
    atom, _ := feed.ToAtom()
    fmt.Println("============================\n")
    fmt.Println("atom:", atom)
    fmt.Println("============================\n")
    return atom
}

func GetRss()(string){
    var feedObject = convertToFeed()
    rss, _ := feedObject.ToRss()
    fmt.Println("atom:", rss)
    return rss
}
