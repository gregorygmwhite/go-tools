#A collection of go src code for various tasks

##Useage

####Clone the Repo into GOPATH/src/github/ggw215/
###OR
####Just import the tools you need into your go code
```
import ( 
  . "github.com/ggw215/go-tools/rss_getter"
  . "github.com/ggw215/go-tools/rss_parser"
)
```

###RSS Getter

####Dependencies
*https://github.com/SlyMarbo/rss (for the rss and rss.Item structs)

####GetFeeds
#####Parameters
* a slice of strings that are urls of rss feeds
#####Returns
* a slice of RSS Items of type rss.Item

####CompileFeedURLs
####Parameters:
* String representing one rss feed URL (use "" if you don't want to use it)
* String representing the filepath to a file of line separated urls
* Int representing the number of feeds you want to limit yourself to

#####Returns 
* a slice of strings that represent a collection of urls

####GetURLsFromFile
#####Parameters
* String: filepath to file of line separated urls
* Int: Limit on number of urls you want returned
#####Returns
* A slice of strings representing a collection or URLs

