Slack
============

Slack is a simple golang package for handling some common slack things like verifying requests from slash commands


Server
============

The server package abstracts running a slack slash command server. The server is configured like any web server from here[github.com/blend/go-sdk/web], but with the addition of a `SlackSignatureSecret`. The server handles verifying and parsing requests from slack into a stuct giving you the ability to start s slash command sever with only needing to write the handler. 

To use this package, import `github.com/mat285/slack/server`. To create a simple slash command server:

```
webConfig := web.NewConfigFromEnv()
config := &server.Config{
    Config: *webConfig,
    SlackSignatureSecret: os.Getenv(slack.EnvVarSignatureSecret),
    AcknowledgeOnVerify: false,
}
serv := server.New(config)
serv.WithHandler(handler)

err := serv.Start()
if err != nil {
    panic(err)
}
```

In here, the `handler` is a function to handles requests from slack. It has the signature:

```func(*slack.SlashCommandRequest) (*slack.Message, error)```

To respond to the slash command simply return the message that will be returned to slack. If an error is returned, the server will log it, and send a message to the user invoking the command that something went wrong.

