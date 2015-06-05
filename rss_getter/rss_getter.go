package rss_getter

import (
	"bufio"
	"fmt"
	"github.com/SlyMarbo/rss"
	"os"
)

const maxFeedItems = 1000

func GetFeeds(feedURLs []string) []*rss.Item {
	if len(feedURLs) > 0 {
		feedItems := make([]*rss.Item, 0)
		for index := 0; index < len(feedURLs); index++ {
			if feedURLs[index] != "" && feedURLs[index] != " " {
				feed, err := rss.Fetch(feedURLs[index])
				if err != nil {
					fmt.Printf("%s", err)
				} else if len(feedItems) > maxFeedItems {
					fmt.Printf("Hit max feed item limit")
					return feedItems
				} else {
					feedItems = append(feedItems, feed.Items...)
				}
			}
		}
		return feedItems
	} else {
		feedItems := make([]*rss.Item, 0)
		fmt.Println("No feedURLs given")
		return feedItems
	}
}

func CompileFeedURLs(givenFeedURL string, feedFilePath string, limit int) []string {
	if givenFeedURL == "" && feedFilePath == "" {
		//nothing given, return empty
		return make([]string, 0)
	} else {
		feedURLSlice := make([]string, 0)
		if givenFeedURL != "" {
			feedURLSlice = append(feedURLSlice, givenFeedURL)
		}
		if feedFilePath != "" {
			feedURLSlice = append(feedURLSlice, GetURLsFromFile(feedFilePath, limit)...)
		}
		return feedURLSlice
	}
}

func GetURLsFromFile(fileName string, limit int) []string {
	slice := make([]string, 0)

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
