package game

import (
	"fmt"
	"log"
	"sort"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/reggles44/semanti-go/pkg/words"
)

type Game struct {
	Answer     string     `json:"answer"`
	Discovered Discovered `json:"discovered"`
	Guesses    uint       `json:"guesses"`
	StartTime  time.Time  `json:"startTime,omitempty"`
	EndTime    time.Time  `json:"endTime,omitempty"`
}

func (g *Game) Duration() time.Duration {
	dur := g.EndTime.Sub(g.StartTime)
	return dur.Truncate(time.Second)
}

func (g *Game) Words() *words.Words {
	w, err := words.GetWords(g.Answer)
	if err != nil {
		log.Print(err)
	}
	return w
}

func (g *Game) Stats() string {
	builder := strings.Builder{}

	builder.WriteString(fmt.Sprintf("Guesses %v", g.Guesses))
	builder.WriteString(" • ")
	builder.WriteString(fmt.Sprintf("Discovered %v", len(g.Discovered)))
	builder.WriteString(" • ")
	builder.WriteString(fmt.Sprintf("Duration %s", g.Duration()))

	return builder.String()
}

const (
	TopHeaderFmt = "\n`%-*s | %4s | %5s`"
	TopFmt       = "\n`%-*s | %4v | %5v`"
)

func (g *Game) Top() *discordgo.MessageEmbed {
	ws := []*words.Word{}
	wordLen := 0

	for w, wobj := range g.Words().TopWords {
		if _, guessed := g.Discovered[w]; guessed {
			ws = append(ws, wobj)
			if len(w) > wordLen {
				wordLen = len(w)
			}
		}
	}

	sort.Slice(ws, func(i int, j int) bool {
		return ws[i].N < ws[j].N
	})

	desc := strings.Builder{}
	desc.WriteString(fmt.Sprintf(TopHeaderFmt, wordLen, "Word", "Rank", "Score"))
	for i := 0; i < len(ws) && i < 20; i++ {
		w := ws[i]
		desc.WriteString(fmt.Sprintf(TopFmt, wordLen, w.Word, w.N, w.Score))
	}

	return &discordgo.MessageEmbed{
		Title:       "Top Words",
		Description: desc.String(),
	}
}
