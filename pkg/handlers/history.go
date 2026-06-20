package handlers

import (
	"github.com/bwmarrin/discordgo"
	"github.com/reggles44/semanti-go/pkg/channel"
)

func history(chg *channel.ChannelInfo, s *discordgo.Session, m *discordgo.MessageCreate) {
	embed := &discordgo.MessageEmbed{
		Title:  "History",
		Fields: []*discordgo.MessageEmbedField{},
	}

	for _, game := range chg.History {
		embed.Fields = append(embed.Fields, &discordgo.MessageEmbedField{
			Value: game.String(),
		})
	}

	s.ChannelMessageSendEmbed(m.ChannelID, embed)
}
