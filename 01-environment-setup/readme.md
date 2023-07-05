# Environment Set-up

## Installation/Set-up Order
- [Go](https://go.dev/)
- [godotenv](https:///github.com/joho/godotenv@v1.5.1)
- [DiscordAPI](https://discord.com/developers/docs/intro)
- [DiscordGo](https://github.com/bwmarrin/discordgo)


## [Go](https://go.dev/) - The language used in this tutorial
  - ### If unfamiliar with Go
    - take [A tour of Go](https://go.dev/tour/welcome/1)
    - [Installation Instructions](https://go.dev/doc/install) (vary by what machine you are using)
  - ### Starting your first go project
  - Open a terminal in your root directory (keep it open), run: 
    ```
    go mod init <your-project-name>
    ```
    This creates a space for Go to store all project dependencies
  - Create a new file called "bot.go":
  ```
  touch bot.go
  ```
  - Say hello world
  ```go
  package main

  import (
    "fmt"
  )

  func main(){
    fmt.Println("Hello World")
  }
  ```


## [godotenv](https://pkg.go.dev/github.com/joho/godotenv@v1.5.1) - Our first package
  ##### a .env file protects valuable information such as your Discord bot token
  ### Installation
  - in your terminal run
  ```
  go get github.com/joho/godotenv
  ```
  ### Instructions

1. Create a .env file
```
touch .env
```
2. enter these fields in the file to fill with your Discord Bot information
DISCORD_BOT_APP_ID="your-app-id"
DISCORD_BOT_PUBLIC_KEY="your-puclic-key"
DISCORD_BOT_TOKEN="your-bot-token"
DISCORD_BOT_PERMISSIONS_INT="your-permissions-integer"

3. Import your environment variables, editing bot.go
```go
  package main

  import (
    "fmt"

    "github.com/joho/godotenv"    
  )

  func main(){
    // Load environment variables from .env file
    err := godotenv.Load()
    if err != nil {
      panic("Error loading .env file")
    }

    fmt.Println("Hello World")
  }
  ```
## DiscordAPI - [Developer Portal](https://discord.com/developers/docs/intro)
  1. Create an account
  2. Create an Application
  3. Add a bot to the Application, saving what belongs in your .env
  4. Invite your bot to a discord server (I reccoment using a server with no one else in it)

## [DiscordGo] - The go library used to interface with the [DiscordAPI]