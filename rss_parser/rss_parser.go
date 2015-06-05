package rss_parser

import (
    "github.com/SlyMarbo/rss"
    "bytes"
    "strings"
    "fmt"
)

const maxFeeds = 10

var punctuations = []string{ ",", ".", "!", "&", ";", "(" , ")"}
var CommonWords = map[string]int{
    "and": 1,
    "a":   1,
    "the": 1,
    "by": 1,
    "with": 1,
    "had": 1,
    "has": 1,
    "more": 1,
    "if": 1,
    "is": 1,
    "in": 1,
    "to": 1,
    "all": 1,
    "would": 1,
    "they": 1,
    "than": 1,
    "it": 1,
    "as": 1,
    "from": 1,
    "or": 1,
    "so": 1,
    "am": 1,
    "was": 1,
    "like": 1,
    "are": 1,
    "&amp;nbsp;": 1,
    "&amp;": 1,
    "nbsp;": 1,
    "...&lt;div": 1,
    "class=&quot;og_rss_groups&quot;&gt;&lt;/div&gt;": 1,
}

func GetWordCountForRSSItem(feedItem *rss.Item) map[string]int{
    wordDistribution := make(map[string]int)
    itemText := compileRSSItemText(feedItem)
    words := strings.Split(itemText, " ")
    for i := range words {
        word := strings.ToLower(words[i])
        if(CommonWords[word] != 1 && word != "" && word != " ") {
            word = trimPunctuation(word)
            wordDistribution[word] = wordDistribution[word] + 1
        }
    }
    return wordDistribution
}

func compileRSSItemText(feedItem *rss.Item) string {
    var buffer bytes.Buffer
    buffer.WriteString(fmt.Sprint(feedItem.Title, " ", feedItem.Summary, " ", feedItem.Content))
    return buffer.String()
}

func trimPunctuation(word string) string {
    for i := range punctuations {
        punctation := punctuations[i]
        word = strings.Trim(word, punctation)
    }
    return word
}
