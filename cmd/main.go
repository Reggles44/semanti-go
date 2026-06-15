package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/reggles44/semanti-go/pkg/game"
	"github.com/reggles44/semanti-go/pkg/handlers"
)

var Token string

func init() {
	flag.StringVar(&Token, "t", "", "Bot Token")
	flag.Parse()
}

func main() {
	bot, err := discordgo.New("Bot " + Token)
	if err != nil {
		log.Fatal("Failed to start bot", err)
		os.Exit(1)
	}

	bot.AddHandler(handlers.MessageRecieve)

	// Open a websocket connection to Discord and begin listening.
	err = bot.Open()
	if err != nil {
		log.Fatal("error opening connection", err)
		os.Exit(1)
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	// Cleanly close down the Discord session.
	bot.Close()
	game.SaveGames()
}
