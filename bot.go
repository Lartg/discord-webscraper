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

	// Register command handlers - these read messages for prefixes to check if they need to do things. commands could get very complex using many flags
	dg.AddHandler(webScraperCreate) // uses "./" message prefix
	dice.roll()
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

/* ----------------------------------------------------------------------------
This Function will:
	Check if the message starts with "./"
		Get the command string.
		Scrape a website for matches to the command
			return links to matches
*/

func webScraperCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore messages from the bot itself or other bots to prevent recursive requests
	if m.Author.ID == s.State.User.ID || m.Author.Bot {
		return
	}

	// Parse command
	if strings.HasPrefix(m.Content, "./") {
		query := strings.TrimPrefix(m.Content, "./")
		baseURL := "https://www.formula1.com"

		// Create a new collector to scrape our baseURL
		c := colly.NewCollector()

		//--------------------------------------------------------
		/*
			This function isolates the HTML elements with the css selectors matching the first parameter.
			Then we do cool things and send a response to our discord server messaging channel

			TODO:
			add the response as a thread to the request, and have a reply for no matches
		*/

		c.OnHTML("#article-list .col-12", func(e *colly.HTMLElement) {

			// get article card title
			title := e.ChildTexts("p")

			// if title contains query store anchor
			queryCheck := strings.Contains((title[1]), query)
			if queryCheck {
				// get the anchor for the article
				anchor := e.ChildAttr("a", "href")

				// make a link to the article by adding the anchor to our base URL
				newURL := baseURL + anchor

				// send link to matched article
				s.ChannelMessageSend(m.ChannelID, newURL)
			}
		})

		//--------------------------------------------------------
		// Visit the search URL and scrape the page with our callback
		err := c.Visit(baseURL + "/en/latest/all.html")
		if err != nil {
			log.Println("Error scraping Google search page: ", err)
			return
		}
		fmt.Println("Successfully scraped f1 latest news, if no article sent there may be no matches, will create a message response for this case")

	}
}
