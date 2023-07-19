# [colly/v2](github.com/gocolly/colly/v2) scraping discord bot

#  Install [colly/v2](github.com/gocolly/colly/v2):


 open a terminal in your project root:

```
go get github.com/gocolly/colly/v2
go mod tidy
```
# Tutorial scraper:

##  Open f1Scraper.go, <small>we will only add code where our scraper lives</small>
```go
package f1Sraper

import (
	"strings"

	"github.com/bwmarrin/discordgo"
)

func Scrape(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore messages from the bot itself or other bots to prevent recursive requests
	if m.Author.ID == s.State.User.ID || m.Author.Bot {
		return
	}
	// Parse command
	if strings.HasPrefix(m.Content, "./") {
    /*
    Your
    Scraper
    Will
    Live
    Here
    */
	}
}

```
## 1. Create a new collector
##### <small>*and find a website you'd like to scrape*</small>
```go
if strings.HasPrefix(m.Content, "./") {
  // Website to be scraped
  baseURL := "<website-base-URL>"

  // Create a new collector to scrape our baseURL	
  c := colly.NewCollector()
}
```
- A colly collector is an object with fast methods for making HTTP requests and parsing HTML *(and much more)*

## 2. What will your collector do?


The method
```go
c.onHTML("<css-selector>", func(e *colly.HTMLElement)){
  // cool stuff goes here
}
```

- is a callback function that triggers when colly finds an element matching your custom css-selector. 
  * Great tool to find good selector -> [Selector Gadget](https://selectorgadget.com/)
- Create callback functions before colly visits a site so that colly knows what to do during its visit.

## 3. Send Colly to fetch elements:
After finishing your callback function
```go
err := c.Visit(baseURL)
if err != nil {
	log.Println("Error scraping <site-name>: ", err)
	return
}
```

## 4. Update f1Scraper.go
1. get base URL
2. create collector
3. get css selector for desired element
4. write callback function
5. visit site

```go
func Scrape(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore messages from the bot itself or other bots to prevent recursive requests
  if m.Author.ID == s.State.User.ID || m.Author.Bot {
		return
  }
  // Parse command
  if strings.HasPrefix(m.Content, "./") {
    // Website to be scraped
    baseURL := "https://www.formula1.com"

    // Create a new collector to scrape our baseURL	
    c := colly.NewCollector()

    // Tell colly what to find, then what to do
    c.onHTML(".col-lg-4:nth-child(1)", func(e *colly.HTMLElement){
      // get the href for the first anchor in the card
      anchor := e.ChildAttr("a", "href")
      
      // Print the anchor to see what we get
      fmt.Println(anchor)
    }

    // Send Colly on a visit to our website
    err := c.Visit(baseURL + "/en/latest/all.html")
    if err != nil {
      log.Println("Error scraping f1 latest news: ", err)

      // Remember that the user is on discord
      s.ChannelMessageSendReply(m.ChannelID, "No matching articles found", m.Reference())
	  return
    }
  }
}
```