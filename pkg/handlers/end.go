package handlers

import (
	"github.com/bwmarrin/discordgo"
	"github.com/reggles44/semanti-go/pkg/channel"
)

func end(chg *channel.ChannelInfo, s *discordgo.Session, m *discordgo.MessageCreate) {
	embed := chg.EndActiveGame("YOU GAVE UP", false)
	if embed != nil {
		s.ChannelMessageSendEmbed(m.ChannelID, embed)
	}
}
