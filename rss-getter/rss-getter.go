package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/SlyMarbo/rss"
	"os"
)

const maxFeedItems = 1000

var maxFeeds int
var commandLineFeedURL string
var feedURLFilePath string

func init() {
	flag.StringVar(&commandLineFeedURL, "feed", "", "the url for the RSS or Atom feed to be fetched")
	flag.StringVar(&feedURLFilePath, "file", "", "the filepath to a line separated file containing the urls of RSS feeds desired to be fetched")
	flag.IntVar(&maxFeeds, "limit", 10, "the maximum number of rss feeds allowed to be fetched by this program")
}

func main() {
	flag.Parse()
    slice := compileFeedURLs(commandLineFeedURL, feedURLFilePath, maxFeeds)
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

func GetFeeds(feedURLs []string) []*rss.Item {
    feedItems := make([]*rss.Item, maxFeedItems)
	if len(feedURLs) != 0 {
		for index := 0; index < len(feedURLs); index++ {
			if feedURLs[index] != "" && feedURLs[index] != " " {
				feed, err := rss.Fetch(feedURLs[index])
				if err != nil {
					fmt.Printf("%s", err)
				} else {
                    feedItems = append(feedItems, feed.Items...)
				}
			}
		}
	} else {
		fmt.Println("No feedURL given")
	}
    return feedItems
}

func compileFeedURLs(givenFeedURL string, feedFilePath string, limit int ) []string {
	feedURLSlice := make([]string, limit)
	if givenFeedURL != "" {
		feedURLSlice = append(feedURLSlice, givenFeedURL)
	}
	if feedFilePath != "" {
		feedURLSlice = append(feedURLSlice, GetURLsFromFile(feedFilePath, limit)...)
	}
	return feedURLSlice
}

func GetURLsFromFile(fileName string, limit int) []string {
	slice := make([]string, limit-1)

	file, err := os.Open(fileName)

	if err != nil {
		panic(err.Error())
	}

	defer file.Close()

	reader := bufio.NewReader(file)
	scanner := bufio.NewScanner(reader)

	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		if len(slice) < limit {
			line := scanner.Text()
			if line != "" {
				slice = append(slice, line)
			}
		}
	}

	return slice
}
