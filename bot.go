package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/gocolly/colly"
	"github.com/joho/godotenv"
)

func main() {
	// get discord bot env vars
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	discordBotToken := os.Getenv("DISCORD_BOT_TOKEN")

	//-------------------------------------------------
	discord, err := discordgo.New("Bot " + os.Getenv(discordBotToken))
	if err != nil {
		log.Fatal(err)
	}

	discord.AddHandler(messageCreate)

	err = discord.Open()
	if err != nil {
		log.Fatal(err)
	}

	<-make(chan struct{})
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if strings.HasPrefix(m.Content, "//") {
		query := strings.TrimSpace(m.Content[2:])
		results, err := googleSearch(query)
		if err != nil {
			s.ChannelMessageSend(m.ChannelID, "Error performing Google search.")
			return
		}
		s.ChannelMessageSend(m.ChannelID, results)
	}
}

func googleSearch(query string) (string, error) {
	c := colly.NewCollector()

	var results string

	c.OnHTML("h3", func(e *colly.HTMLElement) {
		link := e.ChildAttr("a", "href")
		results += fmt.Sprintf("- [%s](https://www.google.com%s)\n", e.Text, link)
	})

	err := c.Visit("https://www.google.com/search?q=" + query)
	if err != nil {
		return "", err
	}

	return results, nil
}
