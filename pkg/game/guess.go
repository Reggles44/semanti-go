package game

import "github.com/reggles44/semanti-go/pkg/words"

type Discovered map[string]string

func (g *Game) Guess(word string, guesser string) (*words.Word, error) {
	g.Guesses += 1
	avail := g.Words()

	// Game Over
	if word == g.Answer {
		return nil, GameOver
	}

	// Already Guessed
	if _, alreadyGuessed := g.Discovered[word]; alreadyGuessed {
		return nil, nil
	}

	// Show Match
	if w, match := avail.TopWords[word]; match {
		g.Discovered[word] = guesser
		return w, nil
	}

	// Invalid Guess
	return nil, nil
}
