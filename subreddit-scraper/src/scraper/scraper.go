package scraper

import (
    "net/http"
    "io/ioutil"
    "encoding/json"
    "../auth"
)

func GetSubPosts(c *auth.Client, ids []string, users []string) ([]string, []string) {
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

