package mashup

import (
	"math/rand"
)

type Matrix struct {
	Lines      []Line
	LineMatrix map[int][]int
}

func NewMatrixGenerator(matrix Matrix) Generator {
	return func() Verse {
		history := map[int]bool{}
		firstLine, secondLine := getLinePair(matrix, history)
		thirdLine, fourthLine := getLinePair(matrix, history)
		return Verse{
			Lines: []Line{firstLine, secondLine, thirdLine, fourthLine},
		}
	}
}

func getLinePair(matrix Matrix, history map[int]bool) (Line, Line) {
	for {
		// grab a random line that hasn't already been used
		firstLineIndex := rand.Intn(len(matrix.Lines))
		if history[firstLineIndex] {
			continue
		}

		// grab another line if this one doesn't have any pairs
		pairs := matrix.LineMatrix[firstLineIndex]
		if len(pairs) == 0 {
			continue
		}

		// choose a random pair for the first line
		pairIndex := rand.Intn(len(pairs))
		secondLineIndex := pairs[pairIndex]

		// update our history for both lines
		history[firstLineIndex] = true
		history[secondLineIndex] = true

		return matrix.Lines[firstLineIndex], matrix.Lines[secondLineIndex]
	}
}
