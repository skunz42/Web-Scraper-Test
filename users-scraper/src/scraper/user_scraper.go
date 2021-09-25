package scraper

import (
    "fmt"
    "net/http"
    "io/ioutil"
    //"encoding/json"
    "../../src/auth"
)

func getCommentUsers(c *auth.Client, ids []string, users []string) {
    comments_endpoint_url := "https://oauth.reddit.com/r/golang/comments"

    fmt.Println(ids[0][3:])

    req, _ := http.NewRequest("GET", comments_endpoint_url + "/" + ids[0][3:], nil)
    req.Header.Set("Authorization", "bearer " + c.Access_Token)
    req.Header.Set("User-Agent", c.User_Agent)
    res, _ := c.Http_Client.Do(req)
    defer res.Body.Close()

    body, _ := ioutil.ReadAll(res.Body)

    fmt.Println(string(body))
}

