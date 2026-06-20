package handlers

import (
	"github.com/bwmarrin/discordgo"
	"github.com/reggles44/semanti-go/pkg/channel"
)

func start(chg *channel.ChannelInfo, s *discordgo.Session, m *discordgo.MessageCreate) {
	msg := chg.StartNewGame()
	if msg != nil {
		s.ChannelMessageSendEmbed(m.ChannelID, msg)
	} else {
		s.ChannelMessageSend(m.ChannelID, "A game is already active")
	}
}
