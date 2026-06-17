package game

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/reggles44/semanti-go/pkg/words"
)

type ChannelGame struct {
	ActiveGame *Game        `json:"game,omitempty"`
	Channel    string       `json:"channel"`
	ChannelID  string       `json:"channel_id"`
	History    []*GameStats `json:"history"`
}

type GameStats struct {
	*Game
	Winner string `json:"author"`
}

func GetOrCreateChannelGame(channelID string) *ChannelGame {
	if chg, exists := saves[channelID]; exists {
		return chg
	} else {
		chg := &ChannelGame{
			ActiveGame: nil,
			Channel:    "",
			ChannelID:  channelID,
		}
		saves[channelID] = chg
		return chg
	}
}

func (chg *ChannelGame) StartNewGame() *discordgo.MessageEmbed {
	if chg.ActiveGame != nil {
		return nil
	}

	w, err := words.GetRandomWord()
	if err != nil {
		log.Print("could not find a random word")
		return nil
	}

	hint := w.GetRandomHint()

	log.Printf("New Game with secret %s", w.Word)
	chg.ActiveGame = &Game{
		Answer:     w.Word,
		Discovered: Discovered{"hint": hint.Word},
		StartTime:  time.Now(),
	}

	return &discordgo.MessageEmbed{
		Title:       "A new game has started",
		Description: fmt.Sprintf("Your starting hint is `%s`", hint.Word),
	}
}

func (chg *ChannelGame) EndActiveGame(winner string, won bool) *discordgo.MessageEmbed {
	if chg.ActiveGame == nil {
		return nil
	}

	finished := chg.ActiveGame
	chg.ActiveGame = nil

	now := time.Now()
	finished.EndTime = now

	chg.History = append(chg.History, &GameStats{finished, winner})
	if len(chg.History) > 7 {
		chg.History = chg.History[1:7]
	}
	chg.ActiveGame = nil

	title := fmt.Sprintf("The game ended.\n The secret word was `%s`.", finished.Answer)
	if won {
		title = fmt.Sprintf("You won!\n The secret word was `%s`.", finished.Answer)
	}

	topWords := make([]*words.Word, 20)
	for _, w := range finished.Words().TopWords {
		if w.Index < 20 {
			topWords[w.Index] = w
		}
	}

	desc := strings.Builder{}
	for _, w := range topWords {
		desc.WriteString(w.String() + "\n")
	}

	embed := &discordgo.MessageEmbed{
		Title:       title,
		Description: desc.String(),
		Footer: &discordgo.MessageEmbedFooter{
			Text: finished.Stats(),
		},
	}

	return embed
}

func (gs *GameStats) String() string {
	s := strings.Builder{}

	s.WriteString(fmt.Sprintf("Winner: %s\n", gs.Winner))
	s.WriteString(fmt.Sprintf("Secret: `%s`\n", gs.Answer))
	s.WriteString(fmt.Sprintf("Total Guesses: %v\n", gs.Guesses))
	s.WriteString(fmt.Sprintf("Discoveries: %v\n", len(gs.Discovered)))
	s.WriteString(fmt.Sprintf("Duration: %s\n", gs.Duration()))

	return s.String()
}
