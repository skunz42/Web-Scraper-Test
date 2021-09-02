package main

import (
    "fmt"
    "net/http"
    "time"
    "io/ioutil"
)

func getJSON(url string, client http.Client) string {
    req, _ := http.NewRequest(http.MethodGet, url, nil)
    req.Header.Set("User-Agent", "spacecount-tutorial")

    res, _ := client.Do(req)

    defer res.Body.Close()
    body, _ := ioutil.ReadAll(res.Body)

    return string(body)

}

func main() {
    url := "https://gbfs.capitalbikeshare.com/gbfs/gbfs.json"
    client := http.Client {
        Timeout: time.Second * 10,
    }

    ret := getJSON(url, client)

    fmt.Println(ret)

}
