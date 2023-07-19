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
