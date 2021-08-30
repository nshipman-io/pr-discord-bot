package bot

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/nshipman-io/pr-discord-bot/config"
	"github.com/nshipman-io/pr-discord-bot/github"
	"strings"
)

var BotID string
var discordBot *discordgo.Session

func Start() {

	discordBot, err := discordgo.New("Bot " + config.Token)
	if err != nil {
		fmt.Println(err.Error())
	}

	user, err := discordBot.User("@me")
	if err != nil {
		fmt.Println(err.Error())
	}
	BotID = user.ID

	discordBot.AddHandler(messageHandler)
	discordBot.AddHandler(pullrequestHandler)

	err = discordBot.Open()
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println("Bot is now running")


}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if strings.HasPrefix(m.Content, config.BotPrefix) {
		if m.Author.ID == BotID {
			return
		}
		if m.Content == "!ping" {
			_, _ = s.ChannelMessageSend(m.ChannelID, "pong")
		}

		fmt.Println(m.Content)
	}
}

func pullrequestHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	var message string
	if m.Content == "!get-prs" {
		repos := github.GetOpenPrs()
		for _,v := range repos {
			if len(v.PullRequests) >= 1 {
				for _, vpr := range v.PullRequests {
					url := fmt.Sprintf("https://github.com/%s/%s/pull/%d", v.Owner, v.Name, vpr.Number)
					message = fmt.Sprintf("Please review the open PR for %s.\n%s", v.Name, url)
					_,_ = s.ChannelMessageSend(m.ChannelID, message)
				}
			} else {
				message = fmt.Sprintf("No open PRs for the repo %s/%s", v.Owner, v.Name)
				_,_ = s.ChannelMessageSend(m.ChannelID, message)
			}
		}

	}
}