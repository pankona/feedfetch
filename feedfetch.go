package main

import (
	"fmt"
	rss "github.com/jteeuwen/go-pkg-rss"
	"os"
)

// TODO: parameterize
const url = "http://blog.tok.access-company.com/feed/atom"

func main() {
	timeout := 5

	feed := rss.New(timeout, true, chanHandler, itemHandler)
	err := feed.Fetch(url, nil)

	if err != nil {
		fmt.Fprintf(os.Stderr, "[e] %s: %s", url, err)
		return
	}
}

func chanHandler(feed *rss.Feed, newchannels []*rss.Channel) {
	// fmt.Println(newchannels[0].Title)
}

func itemHandler(feed *rss.Feed, ch *rss.Channel, newitems []*rss.Item) {
	//fmt.Println(ch.Title)
	for _, item := range newitems {
		fmt.Println(item.Title + " by " + item.Author.Name)
		fmt.Println(item.Id)
		//fmt.Println(item.Updated)
		fmt.Println("")
	}
}
