package scraper

import (
    "net/http"
    "net/url"
    "io/ioutil"
    "encoding/json"
    "skunz42/web-scraper/src/auth"
)

func GetSubPosts(c *auth.Client, ids []Listing) ([]Listing) {
    sub_endpoint_url := "https://oauth.reddit.com/r/golang/new"

    url_params := url.Values{}
    url_params.Add("limit", "5")

    req, _ := http.NewRequest("GET", sub_endpoint_url + "?" + url_params.Encode(), nil)
    req.Header.Set("Authorization", "bearer " + c.Access_Token)
    req.Header.Set("User-Agent", c.User_Agent)

    res, _ := c.Http_Client.Do(req)
    defer res.Body.Close()

    body, _ := ioutil.ReadAll(res.Body)
    r := &Response{}

    json.Unmarshal(body, r)

    for i := range(r.Data.Children) {
        //f64, _ := r.Data.Children[i].Data.Created.Float64()
        //i64 := int(math.Round(f64))
        //ids = append(ids, r.Data.Children[i].Data.Id)
        //users = append(users, r.Data.Children[i].Data.Author)
        ids = append(ids, r.Data.Children[i])
    }

    return ids
}

