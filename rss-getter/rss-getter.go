package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/SlyMarbo/rss"
	"os"
)

var maxFeeds int
var commandLineFeedURL string
var FeedURLFilePath string

func init() {
	flag.StringVar(&commandLineFeedURL, "feed", "", "the url for the RSS or Atom feed to be fetched")
	flag.StringVar(&FeedURLFilePath, "file", "", "the filepath to a line separated file containing the urls of RSS feeds desired to be fetched")
	flag.IntVar(&maxFeeds, "limit", 10, "the maximum number of rss feeds allowed to be fetched by this program")
}

func main() {
	flag.Parse()
	slice := compileFeedURLs()
	if len(slice) != 0 {
		for index := 0; index < len(slice); index++ {
			if slice[index] != "" && slice[index] != " " {
				feed, err := rss.Fetch(slice[index])
				if err != nil {
					fmt.Printf("%s", err)
				} else {
					fmt.Printf("%v", feed.Items)
				}
			}
		}
	} else {
		fmt.Println("No feedURL given")
	}
}

func compileFeedURLs() []string {
	feedURLSlice := make([]string, maxFeeds)
	if commandLineFeedURL != "" {
		feedURLSlice = append(feedURLSlice, commandLineFeedURL)
	}
	if FeedURLFilePath != "" {
		feedURLSlice = append(feedURLSlice, getLinesFromFile(FeedURLFilePath)...)
	}
	return feedURLSlice
}

func getLinesFromFile(fileName string) []string {
	slice := make([]string, maxFeeds-1)

	file, err := os.Open(fileName)

	if err != nil {
		panic(err.Error())
	}

	defer file.Close()

	reader := bufio.NewReader(file)
	scanner := bufio.NewScanner(reader)

	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		if len(slice) < maxFeeds {
			line := scanner.Text()
			if line != "" {
				slice = append(slice, line)
			}
		}
	}

	return slice
}
