# Installing the [DiscordGo](https://github.com/joho/godotenv) package


```
go get github.com/bwmarrin/discordgo
```

# Authenticating a session
  The function discordgo.New() creates a new instance of a discordgo session. It takes 1 argument: the bot token. We concatenate "Bot " with our token from our .env, as this is the required form to authenticate with the Discord API using [DiscordGo](https://github.com/bwmarrin/discordgo).
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

# Opening a live connection
  Using the session object that we created, the Open() function establishes a websocket connection to discord servers. This will allow us to listen for commands/messages. 

```go
	// Open a websocket connection to Discord.
	err = dg.Open()
	if err != nil {
		log.Fatal("Error opening connection: ", err)
		return
	}
```

# Await command
```go
	// Keep listening until interrupted.
	fmt.Println("Bot is now running. Press Ctrl+C to exit.")
	<-make(chan struct{})
}
```
The printed statement lets us know that the bot is running and how to close it. The line:
```go
	<-make(chan struct{})
```
creates an empty channel. The left arrow `<-` shows that the program is waiting to receive a value from the **empty** structure. Because the structure is empty, the program will listen for commands until interupt.

# Make sure the bot is online
root:
```
go run bot.go
```
Check the discord server that you have invited the bot to. If the bot is online, then you have a websocket connection.
IMAGE OF ONLINE BOT HERE