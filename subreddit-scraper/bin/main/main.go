package main

import (
    "fmt"
    "net/http"
    "io/ioutil"
    "encoding/json"
    "../../src/auth"
)

func getSubPosts(c *auth.Client, ids []string, users []string) ([]string, []string) {
    sub_endpoint_url := "https://oauth.reddit.com/r/golang/new"

    req, _ := http.NewRequest("GET", sub_endpoint_url, nil)
    req.Header.Set("Authorization", "bearer " + c.Access_Token)
    req.Header.Set("User-Agent", c.User_Agent)
    res, _ := c.Http_Client.Do(req)
    defer res.Body.Close()

    body, _ := ioutil.ReadAll(res.Body)
    r := &auth.Response{}

    json.Unmarshal(body, r)

    for i := range(r.Data.Children) {
        ids = append(ids, r.Data.Children[i].Data.Name)
        users = append(users, r.Data.Children[i].Data.Author)
    }

    return ids, users
}

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

func main() {
    config_data := auth.MakeClient()

    auth.GetToken(config_data)

    ids := make([]string, 0)
    users := make([]string, 0)

    ids, users = getSubPosts(config_data, ids, users)

    for i := range(users) {
        fmt.Println(users[i])
    }

    getCommentUsers(config_data, ids, users)
}
