package words

import (
	"embed"
	"fmt"
	"log"
	"math/rand/v2"
	"path/filepath"
	"strings"
)

//go:embed data/*
var data embed.FS

var cache map[string]*WordSecret

type WordSecret struct {
	Secret  string
	Matches map[string]*Word
}

type Word struct {
	Word  string
	Index uint16
}

func (w *Word) String() string {
	return fmt.Sprintf("`%s` (#%v)", w.Word, w.Index+2)
}

func (w *Word) Found() string {
	return fmt.Sprintf("`%s` found! #%v", w.Word, w.Index+2)
}

func GetWords(secret string) (*WordSecret, error) {
	if words, exists := cache[secret]; exists {
		return words, nil
	}

	fileName := secret
	if !strings.HasSuffix(secret, ".txt") {
		fileName = secret + ".txt"
	} else {
		secret = strings.Replace(secret, ".txt", "", 1)
	}

	wordData, err := data.ReadFile(filepath.Join("data", fileName))
	if err != nil {
		return nil, err
	}

	ws := &WordSecret{secret, make(map[string]*Word)}

	words := strings.Split(string(wordData), "\n")
	for i, w := range words {
		ws.Matches[w] = &Word{w, uint16(i)}
	}

	return ws, nil
}

func GetRandomWord() (*WordSecret, error) {
	wordDir, err := data.ReadDir("data")
	if err != nil {
		return nil, err
	}

	secret := wordDir[rand.IntN(len(wordDir))].Name()
	log.Printf("GetRandomWord result %s", secret)
	return GetWords(secret)
}

func (w *WordSecret) GetRandomHint() *Word {
	n := uint16(5000 - rand.IntN(100))
	log.Printf("Finding hint with n=%v", n)
	for _, w := range w.Matches {
		if w.Index == n {
			log.Printf("Hint is %v", w.Word)
			return w
		}
	}

	return nil
}
