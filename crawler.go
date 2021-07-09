package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

var c *colly.Collector
var parsingResult wordMeaning

func init() {
	c = colly.NewCollector(
		// Visit only domains: hackerspaces.org, wiki.hackerspaces.org
		colly.AllowedDomains("dictionary.cambridge.org"),
	)
	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	settingSearch()
}

func settingSearch() {
	// On every a element which has href attribute call callback
	c.OnHTML(".def-block.ddef_block", func(e *colly.HTMLElement) {
		var newMeaningAndSentence meaningAndSentence
		// meaning
		e.ForEach(".def.ddef_d.db", func(i int, m *colly.HTMLElement) {
			newMeaningAndSentence.Meaning = formatCrawlerResult(m.Text)
		})
		// sentence
		e.ForEach(".def-body.ddef_b .examp.dexamp", func(i int, m *colly.HTMLElement) {
			newMeaningAndSentence.Sentence = append(newMeaningAndSentence.Sentence, formatCrawlerResult(m.Text))
		})
		parsingResult.ResultList = append(parsingResult.ResultList, newMeaningAndSentence)
	})
}

func Search(wordToSearch string) wordMeaning {
	parsingResult = wordMeaning{}
	parsingResult.WordToSearch = wordToSearch
	c.Visit("https://dictionary.cambridge.org/dictionary/english/" + wordToSearch)
	return parsingResult
}
