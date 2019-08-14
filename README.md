Slack
============

Slack is a simple golang package for handling some common slack things like verifying requests from slash commands


Server
============

The server package abstracts running a slack slash command server. The server is configured like any web server from here[github.com/blend/go-sdk/web], but with the addition of a `SlackSignatureSecret`. 
