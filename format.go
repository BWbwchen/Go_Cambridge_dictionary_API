package main

import (
	"regexp"
	"strings"
)

type meaningAndSentence struct {
	Meaning  string   `json:"meaning"`
	Sentence []string `json:"sentence"`
}

type wordMeaning struct {
	WordToSearch string               `json:"word"`
	ResultList   []meaningAndSentence `json:"result"`
}

func formatCrawlerResult(result string) string {
	space := regexp.MustCompile(`\s+`)
	removeSpace := space.ReplaceAllString(result, " ")
	// remove case of [ C ] or [ T ]
	corT := regexp.MustCompile(`\[\s+.\s+\]`)
	removeCorT := corT.ReplaceAllString(removeSpace, "")
	// remove leading space
	s := strings.TrimSpace(removeCorT)
	return s
}
