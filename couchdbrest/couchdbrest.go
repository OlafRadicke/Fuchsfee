package couchdbrest

import (
    "bytes"
    "fmt"
    "net/http"
    "io/ioutil"
)

// Do a couchdb request and get back blog article as json.
func GetLastBlogArticles( couchdb_host string, result_limit string)([]byte, error){
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
//        panic(err)
        body := []byte{}
        return body, fmt.Errorf("Resr call failed.")
    }
    defer resp.Body.Close()

    fmt.Println("response Status:", resp.Status)
    fmt.Println("response Headers:", resp.Header)
    body, _ := ioutil.ReadAll(resp.Body)
//     fmt.Println("response Body:", string(body))
    return body, nil
}
