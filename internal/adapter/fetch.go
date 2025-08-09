package adapter

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

type FetchResult struct {
	Url  string
	Code int
	Body string
}

func FetchUrl(url string) FetchResult {
	client := &http.Client{Timeout: 500 * time.Millisecond}
	res := FetchResult{Url: url, Code: 500, Body: ""}
	resp, err := client.Get(url)
	if err != nil {
		log.Println("ERR [adapter.FetchUrl]: Unable to establish connection", url, err)
		res.Code = 500
		res.Body = err.Error()
		return res
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("ERR [adapter.FetchUrl]: Unable to parse response", url, err)
		res.Code = 500
		res.Body = err.Error()
		return res
	}

	res.Code = resp.StatusCode
	fmt.Println("Delka odpovedi je: ", len(body))
	if len(body) > 350 {
		res.Body = string(body[:320]) + "  ...(shorter version)"
	} else {
		res.Body = string(body)
	}

	return res
}
