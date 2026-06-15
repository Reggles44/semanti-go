package words

import (
	"embed"
	"encoding/json"
	"fmt"
	"log"
	"math/rand/v2"
	"path/filepath"
	"strings"
)

//go:embed data/*
var data embed.FS

var cachedWords = make(map[string]*Words)

type Words struct {
	Word     string   `json:"word"`
	TopWords TopWords `json:"top_words"`
}

type TopWords map[string]*Word

func (tw *TopWords) UnmarshalJSON(data []byte) error {
	words := []*Word{}
	err := json.Unmarshal(data, &words)
	if err != nil {
		return err
	}

	topWords := make(TopWords)
	for _, w := range words {
		topWords[w.Word] = w
	}

	*tw = topWords

	return nil
}

func GetWords(topic string) (*Words, error) {
	if words, exists := cachedWords[topic]; exists {
		return words, nil
	}

	if !strings.HasSuffix(topic, ".json") {
		topic = fmt.Sprintf("%v.json", topic)
	}

	wordData, err := data.ReadFile(filepath.Join("data", topic))
	if err != nil {
		return nil, err
	}

	var words Words
	err = json.Unmarshal(wordData, &words)
	if err != nil {
		return nil, err
	}

	cachedWords[topic] = &words
	return &words, nil
}

func GetRandomWord() (*Words, error) {
	files, err := data.ReadDir("data")
	if err != nil {
		log.Fatal(err)
	}

	word := files[rand.IntN(len(files))].Name()
	log.Printf("GetRandomWord result %s", word)
	return GetWords(word)
}

func (w *Words) GetRandomHint() *Word {
	n := uint16(5000 - rand.IntN(100))
	log.Printf("Finding hint with n=%v", n)
	for _, w := range w.TopWords {
		if w.N == n {
			log.Printf("Hint is %v", w.Word)
			return w
		}
	}

	return nil
}
