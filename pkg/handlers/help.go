package handlers

import (
	"github.com/bwmarrin/discordgo"
	"github.com/reggles44/semanti-go/pkg/game"
)

var helpEmbed = &discordgo.MessageEmbed{
	URL:         "",
	Type:        "",
	Title:       "Commands",
	Description: "",
	Timestamp:   "",
	Color:       0,
	Footer:      &discordgo.MessageEmbedFooter{},
	Image:       &discordgo.MessageEmbedImage{},
	Thumbnail:   &discordgo.MessageEmbedThumbnail{},
	Video:       &discordgo.MessageEmbedVideo{},
	Provider:    &discordgo.MessageEmbedProvider{},
	Author:      &discordgo.MessageEmbedAuthor{},
	Fields: []*discordgo.MessageEmbedField{
		{Name: "$start", Value: "Start a new game if there isn't one being played", Inline: false},
		{Name: "$top", Value: "Show your best guesses", Inline: false},
		{Name: "$end", Value: "Give up and lose", Inline: false},
		{Name: "$stats", Value: "Show the current game stats", Inline: false},
		{Name: "$hist", Value: "Show the last 20 games' information", Inline: false},
		{Name: "$agg", Value: "Summarize the entire history'", Inline: false},
	},
}

func help(chg *game.ChannelGame, s *discordgo.Session, m *discordgo.MessageCreate) {
	s.ChannelMessageSendEmbed(m.ChannelID, helpEmbed)
}
