package feedfactory

import (
    "time"
    "fmt"
    "github.com/gorilla/feeds"
)

func GetAtom(){



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

}
