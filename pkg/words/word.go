package words

import "fmt"

type Word struct {
	Word  string  `json:"word"`
	N     uint16  `json:"n"`
	Score float32 `json:"score"`
}

func (w *Word) String() string {
	return fmt.Sprintf("`%s` (#%v, %.1f)", w.Word, w.N+2, w.Score)
}

func (w *Word) Found() string {
	return fmt.Sprintf("`%s` found! #%v, (%.1f)", w.Word, w.N+2, w.Score)
}
