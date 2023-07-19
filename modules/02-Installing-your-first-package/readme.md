

## [godotenv](https://pkg.go.dev/github.com/joho/godotenv@v1.5.1) - Our first package
  ###### a .env file protects information such as a Discord bot token
  ### Installation
  - in your project root:
  ```
  go get github.com/joho/godotenv
  ```
  ### Instructions

1. Create a .env file
```
touch .env
```
2. enter these fields in the file to fill with your Discord Bot information
```
DISCORD_BOT_APP_ID="your-app-id"
DISCORD_BOT_PUBLIC_KEY="your-puclic-key"
DISCORD_BOT_TOKEN="your-bot-token"
DISCORD_BOT_PERMISSIONS_INT="your-permissions-integer"
```

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

Now that we can protect your data, we are ready to head to the [Discord Developer Portal!](https://discord.com/developers/docs/intro)

