package mashup

import "strings"

type Line struct {
	SongID int
	Text   string
}

func NewLine(songID int, text string) Line {
	return Line{
		SongID: songID,
		Text:   text,
	}
}

func (l Line) LastWord() string {
	words := strings.Split(l.Text, " ")
	return words[len(words)-1]
}
