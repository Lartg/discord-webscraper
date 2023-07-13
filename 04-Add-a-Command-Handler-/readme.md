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
    - add the command check below, and feel free to deviate from the tutorial here to parse the cammand however you wish
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

## Sending a Response
- bareboned dice roller to check functionality of the bot before writing your own code
- gives a structure to have some fun trial and error 