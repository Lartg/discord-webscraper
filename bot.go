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
	/* Load environment variables from .env file */
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	} else {
		fmt.Println("Successfully loaded environment")
	}

	/* Create a new Discord session with discordGo.New(list of parameters)*/

	dg, err := discordgo.New("Bot " + os.Getenv("DISCORD_BOT_TOKEN"))
	if err != nil {
		// if somethings wrong we...
		panic(err)
	} else {
		// yay
		fmt.Println("Successfully connected to a Discord server")
	}

	// Call the helper function we write to do the thing
	dg.AddHandler(webScraperCreate)

	// Open a websocket connection to Discord
	// put your bot online to await requests
	err = dg.Open()
	if err != nil {
		log.Fatal("Error opening connection: ", err)
		return
	}

	// Lets you know how to stop the bot from running, draining your computer's resource
	// would love to learn how to get free hosting for this kind of software as a hobbyist
	fmt.Println("Bot is now running. Press Ctrl+C to exit.")
	<-make(chan struct{})
}

//-----------------------------------------------------------------------

func diceCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// write basic function to reply with a random dice coll
	fmt.Println("random number from 1 to whatever dice number they ask for")
}

//-----------------------------------------------------------------------

func webScraperCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore messages from the bot itself or other bots to prevent recursive requests
	if m.Author.ID == s.State.User.ID || m.Author.Bot {
		return
	}

	/*
		Check if the message starts with "./"
			Get the following message string.
				log message and resonse
				Send a response message.
				Optional:
					do colly tutorial
					uncomment line XX to to import the webscraping function you made



	*/
	if strings.HasPrefix(m.Content, "./") {
		query := strings.TrimPrefix(m.Content, "./")
		baseURL := "https://www.formula1.com"

		// Create a new collector
		c := colly.NewCollector()

		//--------------------------------------------------------
		// enter css selector for the article card
		c.OnHTML("#article-list .col-12", func(e *colly.HTMLElement) {
			// get article card title
			// if title contains query store anchor
			title := e.ChildTexts("p")
			queryCheck := strings.Contains((title[1]), query)
			if queryCheck {
				anchor := e.ChildAttr("a", "href")
				newURL := baseURL + anchor
				s.ChannelMessageSend(m.ChannelID, newURL)
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
