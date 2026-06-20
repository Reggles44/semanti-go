package channel

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

const gameFile string = "save.json"

type GamesSaves map[string]*ChannelInfo

var saves GamesSaves

func init() {
	LoadGames()
}

func LoadGames() {
	data, err := os.ReadFile(gameFile)
	if err != nil {
		log.Print(err)
		saves = GamesSaves{}
	}

	err = json.Unmarshal(data, &saves)
	if err != nil {
		log.Print(err)
		saves = GamesSaves{}
	}
}

func SaveGames() error {
	if len(saves) == 0 {
		fmt.Println("Game saves is empty skipping saving")
		return nil
	}

	data, err := json.Marshal(saves)
	if err != nil {
		return err
	}

	return os.WriteFile(gameFile, data, 0644)
}
