# Avoiding technical debt
If you want to keep using the same bot, but keep adding features to it, you'll end up with a really messy bot.go. This part of the tutorial will teach you how to store different commands in a reusable library.

## Creating the Library
1. Make a new directory to store commands, make a directory for our dice package, an empty go file
```
mkdir command-library
cd command-library
mkdir dice
cd dice
touch dice.go

```

2. Name the package in dice.go

```go
package dice
```

3. Cut and paste the roll and roll helper function from bot.go to dice.go.
see my command-library for an example.

It is important to rename the roll function to be capital (R) - Roll.()

#### Import the package and register as a handler

```go
// bot.go
package main

import (
	...

	"github.com/<your-username>/<your-repo>/command-library/dice"
  ...
)

func main() {
	...

	// Register command handlers - these read messages for prefixes to check if they need to do things. commands could get very complex using many flags
	dg.AddHandler(dice.Roll) // uses "!" message prefix

	...
}

```
From here on out, whenever we want to make changes to our command handlers, or add commands we will use our command library to do so. This will make our code more modular, and easy to debug.

## Getting ready for module 6

1. create a new package in your command library:
  ```
  cd command-library
  mkdir f1Scraper
  cd f1Scraper
  touch f1Scraper.go

  ```
2. open f1Scraper.go with your fave IDE
3. name the package, and say hello discord upon command. Make sure to use a command flag that is unique from other command handlers in command-library
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
4. Import your package to bot.go and register as a command handler
```go
// bot.go
package main

import (
  ...
  "github.com/<your-username>/<your-repository>/command-library/f1Scraper"
  ...
)
// Register command handlers - these read messages for prefixes to check if they need to do things. commands could get very complex using many flags
dg.AddHandler(f1Scraper.Scrape) // uses "./" message prefix
...
```
5. Test the bot before moving on to module 6