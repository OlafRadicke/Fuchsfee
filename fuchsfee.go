package main

import "fmt"
import "net/http"
import "io/ioutil"
// import "bytes"

func main() {
    fmt.Printf("Hello, world.\n")

    url := "http://127.0.0.1:5984/tuxerjoch/_design/blog_article/_view/all?descending=true&limit=10"
    fmt.Println("URL:>", url)

//     var jsonStr = []byte(`{"title":"Buy cheese and bread for breakfast."}`)
//     req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
//     req.Header.Set("X-Custom-Header", "myvalue")

    req, err := http.NewRequest("GET", url, nil)
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
    fmt.Println("response Body:", string(body))
}
