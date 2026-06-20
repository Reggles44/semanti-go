package handlers

import (
	"github.com/bwmarrin/discordgo"
	"github.com/reggles44/semanti-go/pkg/channel"
)

func stats(chg *channel.ChannelInfo, s *discordgo.Session, m *discordgo.MessageCreate) {
	if chg.ActiveGame != nil {
		s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{Title: "Stats", Description: chg.ActiveGame.Stats()})
	} else {
		s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{Title: "No active game"})
	}
}
