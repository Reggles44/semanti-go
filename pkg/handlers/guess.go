package handlers

import (
	"log"

	"github.com/bwmarrin/discordgo"
	"github.com/reggles44/semanti-go/pkg/game"
)

func guess(chg *game.ChannelGame, s *discordgo.Session, m *discordgo.MessageCreate) {
	if chg.ActiveGame == nil {
		return
	}
	word, err := chg.ActiveGame.Guess(m.Content, m.Author.DisplayName())
	// We won
	if err == game.GameOver {
		embed := chg.EndActiveGame(m.Author.DisplayName(), true)
		s.ChannelMessageSendEmbed(m.ChannelID, embed)

		// Some other error
	} else if err != nil {
		log.Print(err)

		// Found word
	} else if word != nil {
		s.ChannelMessageSend(m.ChannelID, word.Found())
	}
}
