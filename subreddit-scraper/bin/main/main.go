package main

import (
    "fmt"
    "os"
    "skunz42/web-scraper/src/auth"
    "skunz42/web-scraper/src/scraper"
)

func main() {
    if len(os.Args) != 2 {
        fmt.Println("Please enter a filename")
        os.Exit(1)
    }

    config_data := auth.MakeClient(os.Args[1])

    auth.GetToken(config_data)

    ids := make([]scraper.Listing, 0)

    ids = scraper.GetSubPosts(config_data, ids)

    for i := range(ids) {
        fmt.Println(ids[i].Data.Id)
    }
}
