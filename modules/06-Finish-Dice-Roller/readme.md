# Main Function
Your main function should look like this:

```go
package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	// Create a new Discord session.
	dg, err := discordgo.New("Bot " + os.Getenv("DISCORD_BOT_TOKEN"))
	if err != nil {
		log.Fatal("Error creating Discord session: ", err)
		return
	}

	// Register command handlers - these read messages for prefixes to check if they need to do things. commands could get very complex using many flags
	dg.AddHandler(roll) // uses "!" message prefix

	// Open a websocket connection to Discord.
	err = dg.Open()
	if err != nil {
		log.Fatal("Error opening connection: ", err)
		return
	}

	// Wait here until interrupted.
	fmt.Println("Bot is now running. Press Ctrl+C to exit.")
	<-make(chan struct{})
}
```
# Creating the dice roller
The bot will receive dice roll commands in this form:
  - !1d6: command flag, 1 dice 6 sides
  - !2d4: command flag, 2 dice 4 sides each
# New event handler
We only want to handle channel communication with this function. All command processing will be done with a helper function. This lets us clearly divide what our bot is doing, and debug easily.
  1. Check for the command prefix "!" and trim it
  2. Call the helper function to generate an array of dice rolls
  3. Error check
  4. Construct a message to send to our discord channel
  5. Send the message
```go
func roll(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore messages from the bot itself or other bots to prevent recursive requests
	if m.Author.ID == s.State.User.ID || m.Author.Bot {
		return
	}

	// --------------------NEW STUFF BELOW-----------------------
	if strings.HasPrefix(m.Content, "!") {
		command := strings.TrimPrefix(m.Content, "!")
		result := nil //rollHelper(command)
		if result != nil {
			response := fmt.Sprintf("Rolling %s: ", command)
			for i, roll := range result {
				if i > 0 {
					response += ", "
				}
				response += fmt.Sprintf("%d", roll)
			}
			s.ChannelMessageSend(m.ChannelID, response)
		} else {
			s.ChannelMessageSend(m.ChannelID, "Invalid dice roll command. Use form of How many dice rolled, a 'd' for dice, and how many sides. Ex. !1d6 rolls 1 dice with 6 sides.")
		}
	}
}
```


# rollHelper(command)
- Input: dice roll (e.g. 1d4)
- Output: an array of random roll(s)
```go
  func rollHelper(roll string) []int {
```
  first split the command into two parts: numDice and numSides
  ```go
  parts := strings.Split(roll, "d")
	if len(parts) != 2 {
		return nil
	}

	numDice, err := strconv.Atoi(parts[0])
	if err != nil {
		return nil
	}
	numSides, err := strconv.Atoi(parts[1])
	if err != nil {
		return nil
	}
  ```
  Generate a random seed from the current time:
  ```go
  rand.Seed(time.Now().UnixNano())
  ```
  Initialize an array, and store all rolls:
  ```go
  results := make([]int, numDice)
	for i := 0; i < numDice; i++ {
		results[i] = rand.Intn(numSides) + 1
	}
  ```
  return the results to be sent to the channel
  ```go
    return results
  }
  ```