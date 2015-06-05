package main

import (
    "fmt"
    "flag"
    "github.com/SlyMarbo/rss"
    . "github.com/ggw215/go-tools/rss_getter"
    . "github.com/ggw215/go-tools/rss_parser"
)

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
    slice := CompileFeedURLs(commandLineFeedURL, feedURLFilePath, maxFeeds)
	feedItems := GetFeeds(slice)
    fmt.Printf("%v", feedItems)
    if len(feedItems) > 0 {
        distributions := getWordDistributionsForFeed(feedItems)
        printWordDistributions(distributions)
    }
}

func getWordDistributionsForFeed( feedItems []*rss.Item) []map[string]int{
    wordDistributions := make([]map[string]int,0)
    for key := range feedItems {
        feedItem := feedItems[key]
        wordDistributions = append(wordDistributions, GetWordCountForRSSItem(feedItem))
    }
    return wordDistributions
}

func printWordDistributions(distributions []map[string]int) {
    for i := range distributions {
        wordDistribution := distributions[i]
        fmt.Println( "")
        fmt.Println( "_________________________")
        fmt.Println( "")
        fmt.Printf("%v", wordDistribution)
        fmt.Println( "")
        fmt.Println( "_________________________")
    }
}
