package f1Scraper

import (
	"log"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/gocolly/colly/v2"
)

/* ----------------------------------------------------------------------------
This Package will:
	Listen for a message that starts with "./"
		Get the command string.
		Scrape a website for matches to the command
			return links to matches
			return a response if none
*/

func Scrape(s *discordgo.Session, m *discordgo.MessageCreate) {
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

		// Store the messages to be sent as a reply in a slice
		var messages []string

		//--------------------------------------------------------
		/*
			This function isolates the HTML elements with the css selectors matching the first parameter.
			Then we do cool things and append the response to our messages slice.
		*/

		c.OnHTML("#article-list .col-12", func(e *colly.HTMLElement) {
			// get article card title
			title := e.ChildTexts("p")

			// if title contains query store anchor
			queryCheck := strings.Contains(title[1], query)
			if queryCheck {
				// get the anchor for the article
				anchor := e.ChildAttr("a", "href")

				// make a link to the article by adding the anchor to our base URL
				newURL := baseURL + anchor

				// append the link to the messages slice
				messages = append(messages, newURL)
			}
		})

		//--------------------------------------------------------
		// Visit the search URL and scrape the page with our callback
		err := c.Visit(baseURL + "/en/latest/all.html")
		if err != nil {
			log.Println("Error scraping f1 latest news page: ", err)
			return
		}

		// Send all the collected messages as a discord reply
		reply := strings.Join(messages, "\n")
		_, err = s.ChannelMessageSendReply(m.ChannelID, reply, m.Reference())
		if err != nil {
			s.ChannelMessageSendReply(m.ChannelID, "No matching articles found, check spelling?", m.Reference())
		}

	}
}
