# Interfacing with  your discord bot
## Installing the [DiscordGo](https://github.com/joho/godotenv) package

Same as before, we go get the package:

```
go get github.com/bwmarrin/discordgo
```

## Establishing a connection

### Authenticating a session
  - The function discordgo.New() creates a new instance of a discordgo session. It takes 1 argument: the bot token. We concatenate "Bot " with our token from our .env, as this is the required form to authenticate with the Discord API as a bot. If you are not creating/using a bot refer to [DiscordGo](https://github.com/bwmarrin/discordgo).
```go
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/bwmarrin/discordgo"
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
```

### Opening a live connection
  - Using the session object that we created, the Open() function establishes a websocket connection to our discord server. This will allow us to listen for commands/messages. 

```go
	// Open a websocket connection to Discord.
	err = dg.Open()
	if err != nil {
		log.Fatal("Error opening connection: ", err)
		return
	}
```

### Await command
```go
	// Keep listening until interrupted.
	fmt.Println("Bot is now running. Press Ctrl+C to exit.")
	<-make(chan struct{})
}
```
- The printed statement lets us know that the bot is running and how to close it. The line:
```go
	<-make(chan struct{})
```
creates an empty channel. The left arrow `<-` shows that the program is waiting to receive a value from the **empty** structure. Because no value will be received, because it is empty and takes no memory, this is a lightweight way of keeping the program open/listening, indefinetely, until interuption.

### Check the discord server that you have invited the bot to. If the bot is online (little green dot by it's name, then we have a websocket connection, and are ready to add a commands that do cool things.