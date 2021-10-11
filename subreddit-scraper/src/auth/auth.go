package auth

import (
    "net/http"
    "net/url"
    "encoding/json"
    "strings"
)

type Listing struct {
	Kind string
	Data struct {
		Id        string
        Name      string
		Author    string
		Title     string
		Created   json.Number
	}
}

type Response struct {
	Type string
	Data struct {
		Children []Listing
	}
}

func GetToken(c *Client) error {
    endpoint_url := "https://www.reddit.com/api/v1/access_token"
    form := url.Values {
        "grant_type": {"password"},
        "username": {c.Username},
        "password": {c.password},
    }

    req, err := http.NewRequest("POST", endpoint_url, strings.NewReader(form.Encode()))
    if err != nil {
        return err
    }

    req.SetBasicAuth(c.Client_Id, c.client_secret)
    req.Header.Set("User-Agent", c.User_Agent)
    res, err := c.Http_Client.Do(req)
    if err != nil {
        return err
    }

    defer res.Body.Close()

    token_struct := TokenStruct{}
    decoder := json.NewDecoder(res.Body)

    err = decoder.Decode(&token_struct)
    if err != nil {
        return err
    }

    c.Access_Token = token_struct.AccessToken

    return nil
}
