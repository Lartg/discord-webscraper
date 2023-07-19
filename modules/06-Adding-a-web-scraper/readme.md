# [DiscordGo](github.com/bwmarrin/discordgo) and [colly/v2](github.com/gocolly/colly/v2) powered discord bot web scraper tutorial.
- learn to make simple a webscraper that responds to chat message commands via your Discord server
- potential projects include: getting game statistics from various websites, getting popular game build paths, or anything your heart desires

1. Install [colly/v2](github.com/gocolly/colly/v2):


open a terminal in your project root:

```
go get github.com/gocolly/colly/v2
go mod tidy
```

2. create a new package in your command library:
  ```
  cd command-library
  mkdir f1Scraper
  cd f1Scraper
  touch f1Scraper.go

  ```
3. open f1Scraper.go with your fave IDE
4. name the package, and say hello discord upon command. Make sure to use a command flag that is unique from other command handlers in command-library
  ```go
  // f1Scraper.go
  package f1Sraper

  import (
    "fmt"
  )

  func Scrape(s *discordgo.Session, m *discordgo.MessageCreate) {
    // Ignore messages from the bot itself or other bots to prevent recursive requests
    if m.Author.ID == s.State.User.ID || m.Author.Bot {
      return
    }
    // Parse command
    if strings.HasPrefix(m.Content, "./") {
      s.ChannelMessageSend(m.ChannelID, "Hello Discord")
    }
  }
  ```
