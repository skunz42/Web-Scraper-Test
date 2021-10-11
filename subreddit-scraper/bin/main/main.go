package main

import (
    "fmt"
    "os"
    "../../src/auth"
    "../../src/scraper"
)

func main() {
    if len(os.Args) != 2 {
        fmt.Println("Please enter a filename")
        os.Exit(1)
    }

    config_data := auth.MakeClient(os.Args[1])

    auth.GetToken(config_data)

    ids := make([]string, 0)
    users := make([]string, 0)

    ids, users = scraper.GetSubPosts(config_data, ids, users)

    for i := range(ids) {
        fmt.Println(ids[i])
    }
}
