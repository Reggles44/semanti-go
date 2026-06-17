package words

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"log"
	"math/rand/v2"

	"github.com/james-bowman/nlp"
	"github.com/james-bowman/nlp/measures/pairwise"
	"github.com/sjwhitworth/golearn/base"
	"github.com/sjwhitworth/golearn/knn"
	"gonum.org/v1/gonum/mat"
)

//go:embed words.txt
var data []byte
var words = []string{}
var scanner *nlp.LinearScanIndex

func init() {
	err := json.Unmarshal(data, &words)
	if err != nil {
		log.Fatal(err)
	}

	scanner = nlp.NewLinearScanIndex(pairwise.CosineDistance)
	scanner.Index(mat.NewVecDense(len(words), words), "words")
}

type WordsScores struct {
	Word    string
	Matches map[string]*Word
}

type Word struct {
	Word  string
	Index uint16
	Score float32
}

func (w *Word) String() string {
	return fmt.Sprintf("`%s` (#%v, %.1f)", w.Word, w.Index+2, w.Score)
}

func (w *Word) Found() string {
	return fmt.Sprintf("`%s` found! #%v, (%.1f)", w.Word, w.Index+2, w.Score)
}

func GetWords(topic string) (*WordsScores, error) {
	ws := &WordsScores{topic, make(map[string]*Word)}


	return ws, nil
}

func GetRandomWord() (*WordsScores, error) {
	word := words[rand.IntN(len(words))]
	log.Printf("GetRandomWord result %s", word)
	return GetWords(word)
}

func (w *WordsScores) GetRandomHint() *Word {
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
