package main

import (
	"os"
	"encoding/csv"
	"github.com/gempir/go-twitch-irc"
)

// {
// 	address = "irc.twitch.tv";
// 	chatnet = "Twitch";
// 	port = "6667";
// 	password = "<YOUR OAUTH TOKEN>";
// 	use_ssl = "no";
// 	ssl_verify = "no";
// 	autoconnect = "yes";
// }

var channels = []string{"loltyler1", "shroud", "pokimane", "doublelift", "Symfuhny", "DrLupo", "Fortnite", "NICKMERCS", "1DrakoNz", "Chap"}

func main() {
	writer := csv.NewWriter(os.Stdout)

	client := twitch.NewClient(os.Getenv("TWITCH_USER"), os.Getenv("TWITCH_OAUTH"))

	client.OnPrivateMessage(func(message twitch.PrivateMessage) {
		writer.Write([]string{message.Channel, message.Time.String(), message.User.Name, message.Message})
	})

	for _, user := range channels {
		client.Join(user)
	}

	if err := client.Connect(); err != nil {
		panic(err)
	}
}
