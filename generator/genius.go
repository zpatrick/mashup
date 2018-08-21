package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/zpatrick/mashup/mashup"
)

func GetSongLyrics(songID int) ([]mashup.Line, error) {
	songURL := fmt.Sprintf("https://genius.com/songs/%d", songID)
	response, err := http.Get(songURL)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		return nil, err
	}

	catalog := map[string]bool{}
	text := doc.Find(".lyrics").First().Text()
	for _, line := range strings.Split(text, "\n") {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "[") || strings.HasSuffix(line, "]") || line == "" {
			continue
		}

		catalog[line] = true
	}

	lines := make([]mashup.Line, 0, len(catalog))
	for text := range catalog {
		lines = append(lines, mashup.NewLine(songID, text))
	}

	return lines, nil
}
