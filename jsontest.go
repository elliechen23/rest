package main

import (  
    "bytes"
    "fmt"
    "io/ioutil"
    "net/http"
)

func callrest() {  
    var jsonStr = []byte(`{
    "ip": "8.8.8.8"
    }`)
    //url := "http://example.com/graph/history"
    url := "http://ip.jsontest.com/"
    fmt.Println("URL:>", url)

    req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
    req.Header.Set("X-Custom-Header", "myvalue")
    req.Header.Set("Content-Type", "application/json")
    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()

    fmt.Println("response Status:", resp.Status)
    fmt.Println("response Headers:", resp.Header)
    body, _ := ioutil.ReadAll(resp.Body)
    fmt.Println("response Body:", string(body))

}


