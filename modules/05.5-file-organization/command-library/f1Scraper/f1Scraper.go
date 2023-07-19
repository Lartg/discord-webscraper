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
		s.ChannelMessageSend(m.ChannelID, "Hello Discord")
	}
}
