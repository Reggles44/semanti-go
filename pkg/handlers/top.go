package handlers

import (
	"github.com/bwmarrin/discordgo"
	"github.com/reggles44/semanti-go/pkg/game"
)

func top(chg *game.ChannelGame, s *discordgo.Session, m *discordgo.MessageCreate) {
	if chg.ActiveGame != nil {
		s.ChannelMessageSendEmbed(m.ChannelID, chg.ActiveGame.Top())
	} else {
		s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{Title: "No active game"})
	}
}
