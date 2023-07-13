## Receiving a Command
- Your bot needs a way to know if you want it to do things, as well as instructions for what things to do. Right now we have a bot that does nothing but exist, lets change that by adding an event handler.

- In this tutorial we are going to make a dice roller. This is useful to anyone who likes to play tabletop games virtually with their friends over discord. It creates dice rolls that are public, so no shenanigans, and pretty random.

### Adding an event handler
  #### Define the function
  - Beneath the main function of the program, define a function that we will use as our command handler:

  ```go
  func diceCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore messages from the bot itself or other bots to prevent recursive requests
	if m.Author.ID == s.State.User.ID || m.Author.Bot {
		return
	}
  }
  ```
  This funciton takes 2 parameters:
  1. Our session used to communicate with the Discord API
  2. messages that are read live through our websocket connection

    - the check we have added ensures that bots in the channel will not create recursive requests. You should add this check to any event handler involving the reading of messages. 
  #### Parse messages for a command
    - I like to use message prefixes as flags. Something easy to type, but you wouldn't typically start a message/sentence with. In this tutorial messages that begin with an "!" will be read as a command.
    
  ```go
  func diceCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore messages from the bot itself or other bots to prevent recursive requests
	if m.Author.ID == s.State.User.ID || m.Author.Bot {
		return
	}
  // Parse command
	if strings.HasPrefix(m.Content, "!") {
    s.ChannelMessageSend(m.ChannelID, "Command received")
    }
  }
  ```
  #### Register the event Handler
    - before opening the websocket connection, we register our command handlers, refer to bot.go if this is confusing
  ```go
  // Register command handlers - these read messages for prefixes to check if they need to do things. commands could get very complex using many flags
	dg.AddHandler(diceCreate) // uses "!" message prefix
  ```
## Sending a Response
- Start your bot, and make sure that messages with our command flag are received and a channel message is sent in response.
- Now you have created a discord bot that can recognize, and respond to commands. This now becomes like any simple CLI go program, but the interface is a discord channel instead.