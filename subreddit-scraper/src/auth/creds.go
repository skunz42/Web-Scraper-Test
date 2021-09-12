package auth

import (
    "net/http"
    "time"
    "os"
    "bufio"
)

const file_line_nums = 5

//TODO add in fn
func MakeClient(fn string) *Client {
    c := &Client {
        Http_Client: &http.Client{ Timeout: time.Second * 10, },
    }

    file, _ := os.Open(fn)

    defer file.Close()

    scanner := bufio.NewScanner(file)
    i := 0

    for scanner.Scan() {
        if i >= file_line_nums {
            break
        }

        if i%file_line_nums == 0 {
            setClientId(scanner.Text(), c)
        } else if i%file_line_nums == 1 {
            setClientSecret(scanner.Text(), c)
        } else if i%file_line_nums == 2 {
            setUserAgent(scanner.Text(), c)
        } else if i%file_line_nums == 3 {
            setUsername(scanner.Text(), c)
        } else if i%file_line_nums == 4 {
            setPassword(scanner.Text(), c)
        } else {
            break
        }

        i += 1
    }

    return c
}
