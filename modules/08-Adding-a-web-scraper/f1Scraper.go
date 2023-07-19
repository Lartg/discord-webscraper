package f1Scraper

import (
	"log"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/gocolly/colly/v2"
)

/* ----------------------------------------------------------------------------
This Function will:
	Check if the message starts with "./"
		Return the most recent f1 news article
*/

func Scrape(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore messages from the bot itself or other bots to prevent recursive requests
	if m.Author.ID == s.State.User.ID || m.Author.Bot {
		return
	}

	// Parse command
	if strings.HasPrefix(m.Content, "./") {
		baseURL := "https://www.formula1.com"

		// Create a new collector to scrape our baseURL
		c := colly.NewCollector()

		//--------------------------------------------------------
		/*
			This callback function isolates the anchor associated with the Newest Article card
			Then adds this anchor to our baseURL and shares the URL with DiscordGo
		*/

		c.OnHTML("#article-list .col-12", func(e *colly.HTMLElement) {
			// get the anchor for the article
			// TODO SEND ONLY THE FIRST_______________
			anchor := e.ChildAttr("a", "href")

			// make a link to the article by adding the anchor to our base URL
			newURL := baseURL + anchor

			// send link to matched article
			s.ChannelMessageSend(m.ChannelID, newURL)

		})

		//--------------------------------------------------------
		// Visit the search URL and scrape the page with our callback
		err := c.Visit(baseURL + "/en/latest/all.html")
		if err != nil {
			log.Println("Error scraping Google search page: ", err)
			return
		}
	}
}
