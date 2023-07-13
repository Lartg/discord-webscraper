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
	dg.AddHandler(diceCreate) // uses "!" message prefix

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

/*	----------------------------------------------------------------------------\
This Function will:
	Check if the message starts with "!"
		Get the following message string that corresponds to a dice roll (e.g. 1d6, 2d6, 1d8)
			Send a response message containing roll outcomes.
*/

func diceCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore messages from the bot itself or other bots to prevent recursive requests
	if m.Author.ID == s.State.User.ID || m.Author.Bot {
		return
	}

	// Parse command
	if strings.HasPrefix(m.Content, "!") {
		command := strings.TrimPrefix(m.Content, "!")
		result := rollDice(command)
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
			s.ChannelMessageSend(m.ChannelID, "Invalid dice roll command.")
		}
	}
}

// Our Dice rolling helper functions
func rollDice(roll string) []int {
	parts := strings.Split(roll, "d")
	if len(parts) != 2 {
		return nil
	}

	numDice := parseDiceCount(parts[0])
	numSides := parseDiceSides(parts[1])

	if numDice <= 0 || numSides <= 0 {
		return nil
	}

	rand.Seed(time.Now().UnixNano())

	results := make([]int, numDice)
	for i := 0; i < numDice; i++ {
		results[i] = rand.Intn(numSides) + 1
	}

	return results
}

func parseDiceCount(countStr string) int {
	count, err := strconv.Atoi(countStr)
	if err != nil {
		return 0
	}
	return count
}

func parseDiceSides(sidesStr string) int {
	sides, err := strconv.Atoi(sidesStr)
	if err != nil {
		return 0
	}
	return sides
}
