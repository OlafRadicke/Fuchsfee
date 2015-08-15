package jsonconvert

import (
    "fmt"
    "encoding/json"
)


// Object to represent the value part in the json tree.
type Article []struct {
    Value map[string]interface{} `json:"value"`
}

// Object to represent the rows part in the json tree.
type ArticleList struct {
    Rows Article `json:"rows"`
}

// Converting json to go objecth
// Object hierarchy: "rows"/"value"/"title"
func JsonToObject(json_body []byte)(*ArticleList){
    article_list := &ArticleList{}
    if err := json.Unmarshal(json_body, &article_list); err != nil {
        panic(err)
    }
    fmt.Println("============================")
    fmt.Println(article_list.Rows[0])
    fmt.Println("============================\n")
    fmt.Println(article_list.Rows[0].Value["title"])
    return article_list
}
