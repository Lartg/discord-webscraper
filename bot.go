package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"net/url"

	"github.com/bwmarrin/discordgo"
	"github.com/gocolly/colly/v2"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file: ", err)
		return
	}

	// Create a new Discord session.
	dg, err := discordgo.New("Bot " + os.Getenv("DISCORD_BOT_TOKEN"))
	if err != nil {
		log.Fatal("Error creating Discord session: ", err)
		return
	}

	// Register messageCreate handler
	dg.AddHandler(messageCreate)

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

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore messages from the bot itself or other bots
	if m.Author.ID == s.State.User.ID || m.Author.Bot {
		return
	}

	// Check if the message starts with "./"
	if strings.HasPrefix(m.Content, "./") {
		query := strings.TrimPrefix(m.Content, "./")
		searchURL := "https://www.google.com/search?q=" + url.QueryEscape(query)

		// Create a new collector
		c := colly.NewCollector()

		// Variable to store the first link
		var firstLink string

		// Set a flag to track if the first link has been found
		firstLinkFound := false

		// Extract the first link from the Google search results
		c.OnHTML("a[href]", func(e *colly.HTMLElement) {
			if !firstLinkFound {
				link := e.Attr("href")
				if strings.HasPrefix(link, "/url?q=") {
					decodedLink, err := url.QueryUnescape(link[7:])
					if err == nil {
						firstLink = decodedLink
						firstLinkFound = true
					}
				}
			}
		})

		// Visit the search URL and scrape the page
		err := c.Visit(searchURL)
		if err != nil {
			log.Println("Error scraping Google search page: ", err)
			return
		}

		// Send the first link from the search results as a message
		if firstLink != "" {
			_, err := s.ChannelMessageSend(m.ChannelID, firstLink)
			if err != nil {
				log.Println("Error sending message: ", err)
			}
		}
	}
}
