package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/zpatrick/mashup/mashup"
)

const RhymeFilePath = "rhyme_map.json"
const MatrixFilePath = "matrix.json"

func main() {
	songs := map[string]int{
		"Thrift Shop":        86538,
		"Tootsee Roll":       7159,
		"Hot Boyz":           33171,
		"Can't Hold Us":      57234,
		"Expression":         41347,
		"Gangsta's Paradise": 1352,
		"Flava In Ya Ear":    30077,
		"Big Poppa":          191,
		"No Hands":           1051,
		"The Humpty Dance":   3606,

		// todo: others
		"Lose Yourself": 207,
	}

	bytes, err := ioutil.ReadFile(RhymeFilePath)
	if err != nil {
		log.Printf("Failed to load rhyme map file '%s': %s", RhymeFilePath, err.Error())
		os.Exit(1)
	}

	var rhymeMap map[string][]string
	if err := json.Unmarshal(bytes, &rhymeMap); err != nil {
		log.Printf("Failed to unmarshal rhyme matrix: %s", err.Error())
		os.Exit(1)
	}

	lines := []mashup.Line{}
	for title, songID := range songs {
		log.Printf("Getting lyrics for %s", title)
		songLines, err := GetSongLyrics(songID)
		if err != nil {
			log.Printf("Failed to get song lyrics for '%s': %s", title, err.Error())
			os.Exit(1)
		}

		log.Printf("Getting rhymes for %s", title)
		if err := UpdateRhymeMatrix(songLines, rhymeMap); err != nil {
			log.Printf("Failed to update rhyme matrix: %s", err.Error())
			os.Exit(1)
		}

		data, err := json.Marshal(rhymeMap)
		if err != nil {
			log.Printf("Failed to marshal rhyme map: %s", err.Error())
			os.Exit(1)
		}

		if err := ioutil.WriteFile(RhymeFilePath, data, 0644); err != nil {
			log.Printf("Failed to write rhyme map file '%s': %s", RhymeFilePath, err.Error())
			os.Exit(1)
		}

		lines = append(lines, songLines...)
	}

	lineMatrix := map[int][]int{}
	for i := 0; i < len(lines); i++ {
		lineMatrix[i] = []int{}
		wordA := strings.ToLower(lines[i].LastWord())

		for j := 0; j < len(lines); j++ {
			if i == j {
				continue
			}

			wordB := strings.ToLower(lines[j].LastWord())
			for _, word := range rhymeMap[wordA] {
				if strings.ToLower(word) == wordB {
					lineMatrix[i] = append(lineMatrix[i], j)
					break
				}
			}
		}
	}

	matrix := mashup.Matrix{
		Lines:      lines,
		LineMatrix: lineMatrix,
	}

	data, err := json.Marshal(matrix)
	if err != nil {
		log.Printf("Failed to marshal matrix: %s", err.Error())
		os.Exit(1)
	}

	if err := ioutil.WriteFile(MatrixFilePath, data, 0644); err != nil {
		log.Printf("Failed to write matri map file '%s': %s", MatrixFilePath, err.Error())
		os.Exit(1)
	}
}
