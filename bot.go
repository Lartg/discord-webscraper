package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/gocolly/colly/v2"
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
		baseURL := "https://www.formula1.com"

		// Create a new collector
		c := colly.NewCollector()

		//--------------------------------------------------------
		// enter css selector for the article card
		// var anchors []string
		c.OnHTML("#article-list .col-12", func(e *colly.HTMLElement) {
			// get article card title
			// if title contains query store anchor

			title := e.ChildTexts("p")
			queryCheck := strings.Contains((title[1]), query)
			if queryCheck {
				anchor := e.ChildAttr("a", "href")
				fmt.Println(anchor)
			}
		})
		//--------------------------------------------------------
		// Visit the search URL and scrape the page
		err := c.Visit(baseURL + "/en/latest/all.html")
		if err != nil {
			log.Println("Error scraping Google search page: ", err)
			return
		}

		// iterate through links and send a message for each
	}
}
