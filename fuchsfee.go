package main

import (
    "bytes"
    "fmt"
    "net/http"
    "io/ioutil"
    "encoding/json"
)

// Do a couchdb request and get back blog article as json.
func getLastBlogArticles( couchdb_host string, result_limit string)([]byte, int){
    var string_buffer bytes.Buffer
    string_buffer.WriteString("http://")
    string_buffer.WriteString(couchdb_host)
    string_buffer.WriteString(":5984/tuxerjoch/_design/blog_article/_view/all?descending=true&limit=")
    string_buffer.WriteString(result_limit)

    fmt.Println("URL:  ", string_buffer.String())

    req, err := http.NewRequest("GET", string_buffer.String(), nil)
    req.Header.Set("Content-Type", "application/json")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        fmt.Println(err)
        panic(err)
    }
    defer resp.Body.Close()

    fmt.Println("response Status:", resp.Status)
    fmt.Println("response Headers:", resp.Header)
    body, _ := ioutil.ReadAll(resp.Body)
//     fmt.Println("response Body:", string(body))
    return body, 0
}

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
func json_to_object(json_body []byte)(*ArticleList){
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


func main() {
    fmt.Printf("Hello, world.\n")

    body, err := getLastBlogArticles("127.0.0.1", "10")
    fmt.Println("err: ", err)
    article_list := json_to_object(body)
    fmt.Println("============================\n")
    fmt.Println(article_list.Rows[0].Value["title"])

}
