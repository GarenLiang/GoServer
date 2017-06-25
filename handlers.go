package main
import (
  "fmt"
  "github.com/mitchellh/mapstructure"
)
func addChannel(client *Client, data interface{}) {
  var channel Channel
  var message Message
  mapstructure.Decode(data, &channel)
  fmt.Printf("%#v\n", channel)
  // todo with database
  channel.Id = "123"
  message.Name = "channel add"
  message.Data = channel
  client.send <- message
}
