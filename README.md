# YetiRevival
a revival for the old version of PvZ3, currently in progress, any help appreciated

# How can I help?
if you have any experience with any of the following things (server development, game revivals, packet capturing) contact me on discord: funnyguy098

you can get the APK from your own sources or redownload the game if you owned it.

# What are you using, so I know if i can contribute or not?
for this project, I am using golang. I was originally using XAMPP but I need to be able to respond to both POST and GET requests from the client.

# What's done?
The game gets to the Age Check, errors, if you press OK, asks if you'd like to download assets, and if you do this it errors again?

# How do I edit the game to use my custom servers? (on the latest version)
for production (release) unpack the apk, go into assets and open up app_settings.json.
from here, you can change stuff such as serverenvironment (public/prod or dev), or developersettings
but what we are focussing on is:
*GameServerBaseUrl*
*CdnServerBaseUrl*

Change both of these domains to your own custom web server. (You may also need a verified SSL certificate trusted by your device.)

For dev, you will need a hex editor.
First of all, unpack the APK, go into assets and open up app_settings.json.
Now, change the 
*"ServerEnvironment": "prod",* string to say
*"ServerEnvironment": "dev",*
Now, get your Hex Editor, and go to assets/bin/Data/Managed/Metadata and open it up.
Look for both of these URLs.
*https://pvz3-prod.awspopcap.com/*
*https://pvz3-prd-cdn.popcap.com*

Now, simply change them to your own. (The URL **MUST** be the same size or less as the original one)
