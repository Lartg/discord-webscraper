package dice

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
)

/*	----------------------------------------------------------------------------\
This Function will:
	Check if the message starts with "!"
		Get the following message string that corresponds to a dice roll (e.g. 1d6, 2d6, 1d8)
			Send a response message containing roll outcomes.
*/

func Roll(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore messages from the bot itself or other bots to prevent recursive requests
	if m.Author.ID == s.State.User.ID || m.Author.Bot {
		return
	}

	// --------------------NEW STUFF BELOW-----------------------
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
			s.ChannelMessageSend(m.ChannelID, "Invalid dice roll command. Use form of How many dice rolled, a 'd' for dice, and how many sides. Ex. 1d6 rolls 1 dice with 6 sides.")
		}
	}
}

// Our Dice rolling helper function
func rollDice(roll string) []int {
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

	rand.Seed(time.Now().UnixNano())

	results := make([]int, numDice)
	for i := 0; i < numDice; i++ {
		results[i] = rand.Intn(numSides) + 1
	}

	return results
}
