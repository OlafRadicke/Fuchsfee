package feedfactory

import (
    "time"
    "fmt"
    "github.com/gorilla/feeds"
    "../couchdbrest"
    "../jsonconvert"
    "../appconfig"
)


func convertToFeed(config appconfig.AppConfig)(*feeds.Feed, error){


    body, err := couchdbrest.GetLastBlogArticles("127.0.0.1", "10")
    fmt.Println("err: ", err)

    if err != nil {
        fmt.Println("CouchDB rest call faild. Is database all ready running?")
//        body := []byte{}
        return nil, fmt.Errorf("CouchDB rest call faild. Is database all ready running?")
    }
    article_list, err := jsonconvert.JsonToObject(body)
    fmt.Println("============================\n")
    fmt.Println(article_list.Rows[0].Value["title"])
    fmt.Println("============================\n")
    fmt.Println("Anzahl Einträge:", len(article_list.Rows))

    now := time.Now()
    feed := &feeds.Feed{
        Title:       "THE INDEPENDENT FRIEND",
        Link:        &feeds.Link{Href: "https://the-independent-friend.de/"},
        Description: "Weblog von Olaf Radicke",
        Author:      &feeds.Author{config.Author, config.AuthorMail},
        Created:     now,
    }

    feed.Items = []*feeds.Item{}
    for index,element := range article_list.Rows {
        fmt.Println("============================\n")
        fmt.Println( "No.:" , index ,  " Value: " , element.Value["title"] )
        fmt.Println("============================\n")

        articleTitle := element.Value["title"]
        feed.Items = append(
            feed.Items,
            &feeds.Item{
                Title:       articleTitle.(string),
                Link:        &feeds.Link{Href: config.BlogProtocol + config.BlogDomain},
                Description: "A discussion on controlled parallelism in golang",
                Author:      &feeds.Author{config.Author, config.AuthorMail},
                Created:     now,
            } )
    }

    return feed, nil
}

func GetAtom(config appconfig.AppConfig)(string, error){
    var feed, err = convertToFeed(config)
    if err != nil {
        fmt.Println("Can't create feed!")
        return "", fmt.Errorf("Can't create feed!")
    }
    atom, _ := feed.ToAtom()
//    fmt.Println("============================\n")
//    fmt.Println("atom:", atom)
//    fmt.Println("============================\n")
    return atom, nil
}

func GetRss(config appconfig.AppConfig)(string){
    var feedObject, _ = convertToFeed(config)
    rss, _ := feedObject.ToRss()
//    fmt.Println("atom:", rss)
    return rss
}
