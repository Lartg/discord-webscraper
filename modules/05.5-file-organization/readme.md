# Avoiding technical debt
If you want to keep using the same bot, but keep adding features to it, you'll end up with a really messy bot.go. This part of the tutorial will teach you how to store different commands in a reusable library.

## Creating the Library
1. Make a new directory to store commands, make a directory for our dice package, an empty go file
```
touch command-library
cd command-library
touch dice
cd dice
touch dice.go

```

2. Name the package in dice.go

```go
package dice
```

3. Cut and paste the roll and roll helper function from bot.go to dice.go.
see my command-library for an example.

It is important to rename the roll function to be capital (R) - Roll.()

#### Import the package and register as a handler

```go
// bot.go
package main

import (
	...

	"github.com/Lartg/discord-webscraper/command-library/dice"
  ...
)

func main() {
	...

	// Register command handlers - these read messages for prefixes to check if they need to do things. commands could get very complex using many flags
	dg.AddHandler(dice.Roll) // uses "!" message prefix

	...
}

```
