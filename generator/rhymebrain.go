package main

import (
	"fmt"
	"log"
	"net/url"
	"strings"
	"time"

	"github.com/zpatrick/mashup/mashup"
	"github.com/zpatrick/rclient"
)

// The amount in which two words rhyme is given a score between 0 and 300.
// Words that don't meet MinRequiredScore are not used.
const (
	MaxSleepDuration = time.Second * 30
	MinRequiredScore = 250
)

type RhymeBrainMatch struct {
	Word  string
	Score int
}

func UpdateRhymeMatrix(lines []mashup.Line, rhymeMap map[string][]string) error {
	client := rclient.NewRestClient("http://rhymebrain.com")
	for _, line := range lines {
		word := strings.ToLower(line.LastWord())
		if _, ok := rhymeMap[word]; !ok {
			rhymes, err := getWordRhymes(client, word)
			if err != nil {
				return err
			}

			rhymeMap[word] = rhymes
		}
	}

	return nil
}

func getWordRhymes(client *rclient.RestClient, word string) ([]string, error) {
	q := url.Values{}
	q.Set("function", "getRhymes")
	q.Set("word", word)

	// the rhymbrain api will throttle us once we hit a certain threshold.
	// sleeping at least 1 second each time to try and stay under that limit.
	log.Printf("Finding rhymes for '%s'", word)
	for d := time.Second; d < MaxSleepDuration; d += time.Second {
		time.Sleep(d)

		var matches []RhymeBrainMatch
		if err := client.Get("/talk", &matches, rclient.Query(q)); err != nil {
			if err, ok := err.(*rclient.ResponseError); ok && err.Response.StatusCode == 429 {
				continue
			}

			return nil, err
		}

		rhymes := []string{}
		for _, match := range matches {
			if match.Score >= MinRequiredScore {
				rhymes = append(rhymes, strings.ToLower(match.Word))
			}
		}

		return rhymes, nil
	}

	return nil, fmt.Errorf("Failed to get rhymes for '%s'", word)
}
