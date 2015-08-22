# Slack

Golang client for the Slack API.

## Currently supports:

* files

    files.upload: Upload an image/file

* channels

    channels.list: Lists all channels in a Slack team.

* chat

    chat.postMessage: Sends a message to a channel.

## Example

```go
package main

import (
  "github.com/bluele/slack"
)

const (
  token       = "your-api-token"
  channelName = "general"
)

func main() {
  api := slack.New(token)
  channelID, err := api.LookupChannelID(channelName)
  if err != nil {
    panic(err)
  }
  err = api.ChatPostMessage(channelID, "Hello, world!", nil)
  if err != nil {
    panic(err)
  }
}
```

# Author

**Jun Kimura**

* <http://github.com/bluele>
* <junkxdev@gmail.com>
