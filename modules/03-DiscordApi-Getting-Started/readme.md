# DiscordAPI - [Developer Portal](https://discord.com/developers/docs/intro)

## New Bot Instructions:
--- 
1. **login or signup** with the [Discord Developer Portal](https://discord.com/developers/applications)
    * verify your email
2. **Create a New Application**: Once logged in, click "New Application"
IMAGE HERE SHOWING LOCATION

3. **Name your bot**, this name will be it's Discord username
IMAGE HERE COMPARING DISCORD APP TO DEV PORTAL

4. **Create a Bot User**: In the left sidebar, click on the "Bot" tab, then click on the "Add Bot" button. Confirm your action in the pop-up window.
IMAGE SHOWING STEPS

5. **Customize Bot User Settings**: Under the "Bot" section, you can configure additional settings for your bot. You can set a custom profile picture, change the username, toggle features like "Public Bot" or "Require OAuth2 Code Grant," etc.
IMAGE SHOWING STEPS

6. **Retrieve Token**: Scroll down to the "Token" section. Click on the "Copy" button to copy your bot token to the clipboard. Keep this token secure and private, as it grants full access and control to your bot. Add to your .env, along with any other keys/tokens/ids.
IMAGE SHOWING STEPS


7. **Invite the Bot to Your Server**: To invite your bot to a server, go to the "OAuth2" tab in the left sidebar. Under the "URL Generator" section, select the necessary bot permissions based on your bot's intended functionality. In this tutorial select "bot", then in the next window select: Read Messages/View Channels, Send Messages, Send Messages in Threads, Embed Links. Then, copy the generated OAuth2 URL and open it in a new browser tab. Select a server where you have the necessary permissions, and follow the authorization flow to add the bot to your chosen server.

IMAGES SHOWING STEPS

---

#### Your bot is now created and added to your server. You can use the bot token to authenticate and interact with the Discord API using DiscordGo. Defer to the [Discord API documentation](https://discord.com/developers/docs/intro) for more information on how to use the API and implement bot functionality.

Onto the next module where we will begin interfacing with our new bot!