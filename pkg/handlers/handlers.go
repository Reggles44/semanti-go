package handlers

import (
	"log"
	"regexp"

	"github.com/bwmarrin/discordgo"
	"github.com/reggles44/semanti-go/pkg/game"
)

type Handler func(chg *game.ChannelGame, s *discordgo.Session, m *discordgo.MessageCreate)

var handlers = map[string]Handler{
	"$agg":   aggregrate,
	"$end":   end,
	"$help":  help,
	"$hist":  history,
	"$start": start,
	"$stats": stats,
	"$top":   top,
}

var letterRegex = regexp.MustCompile(`^[a-zA-Z]*$`)

func MessageRecieve(s *discordgo.Session, m *discordgo.MessageCreate) {
	channel, err := s.Channel(m.ChannelID)
	if err != nil {
		log.Print(err)
		return
	}

	if channel.Name != "semanti" {
		return
	}

	chg := game.GetOrCreateChannelGame(m.ChannelID)

	if h, exists := handlers[m.Content]; exists {
		go h(chg, s, m)
	} else if !m.Author.Bot && letterRegex.MatchString(m.Content) {
		go guess(chg, s, m)
	}
}
