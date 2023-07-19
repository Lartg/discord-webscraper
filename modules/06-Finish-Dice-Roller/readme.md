# Wrapping up the Dice Roller
## Main Function
 - Your main function should look like this. We will only be changing the event handler.
 - For each command I like to add a new handler.
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
  - the bot will receive dice roll commands in this form:
  - !1d6: 1 dice 6 sides
  - !2d4: 2 dice 4 sides each
## this will be our new event handler
  - We only want to handle channel communication with this function. All command processing will be done with a helper function. This lets us clearly divide what our bot is doing, and debug easily.
  1. Check for our command prefix "!", then trim it, as it is not relevent to our instructions.
  2. we call our helper function, and pass through our command. This will return an array of dice rolls to send out.
  3. Quick error check, and make sure that users are using the right command format.
  4. We construct a message to send to our discord channel.
  5. Send the message.
```go
func roll(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore messages from the bot itself or other bots to prevent recursive requests
	if m.Author.ID == s.State.User.ID || m.Author.Bot {
		return
	}

	// --------------------NEW STUFF BELOW-----------------------
	if strings.HasPrefix(m.Content, "!") {
		command := strings.TrimPrefix(m.Content, "!")
		result := rollHelper(command)
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
## Helper functions

### rollDice(command)
  - This function takes a dice roll (1d6), and will output an array or random rolls
  ```go
  func rollHelper(roll string) []int {
  ```
  - the function first splits the command into two parts: numDice and numSides
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
  - We then generate a random seed based upon the time of the request to ensure that we don't constantly repeat rolls.
  ```go
  rand.Seed(time.Now().UnixNano())
  ```
  - then we make an array, and store as many random rolls (from 1 to the number of sides commanded) as we have dice in the command.
  ```go
  results := make([]int, numDice)
	for i := 0; i < numDice; i++ {
		results[i] = rand.Intn(numSides) + 1
	}
  ```
  - return the results to be sent to the channel
  ```go
    return results
  }
  ```