package main

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/go-ego/gse"
)

type SearchEngine struct {
	ctx       context.Context
	wordIndex map[string][]int
	index     int
	wordCount int
}

func newSearchEngine() *SearchEngine {
	var s *SearchEngine

	s = &SearchEngine{
		wordIndex: make(map[string][]int),
		index:     0,
		wordCount: 0,
	}

	return s
}

func (s *SearchEngine) startup(ctx context.Context) {
	s.ctx = ctx
	s.readFile()
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func (s *SearchEngine) newSegmenter() gse.Segmenter {

	var segmenter gse.Segmenter
	err := segmenter.LoadDictEmbed()

	checkError(err)

	return segmenter

}

func (s *SearchEngine) splitLine(line string) []string {

	line = strings.ReplaceAll(line, "？", "。")
	line = strings.ReplaceAll(line, "！", "。")
	line = strings.ReplaceAll(line, "，", "。")
	line = strings.ReplaceAll(line, "?", "。")
	line = strings.ReplaceAll(line, "!", "。")

	sentences := strings.Split(line, "。")
	var clearSentences []string

	for _, s := range sentences {
		if strings.TrimSpace(s) != "" {
			clearSentences = append(clearSentences, s)
		}
	}

	return clearSentences
}

func (s *SearchEngine) printTable(data SearchEngine) {
	for word, indices := range data.wordIndex {
		fmt.Printf("%s appears in: %d\n", word, indices)
	}
}

func (s *SearchEngine) searchOccurence(target string) ([]int, error) {
	var result []int
	var err error

	result = s.wordIndex[target]

	if result == nil {
		err = fmt.Errorf("Error 404: Word not found")
	}

	return result, err
}

func (s *SearchEngine) readFile() {

	infile := inputEmbed

	segmenter := s.newSegmenter()

	sentences := s.splitLine(infile)

	for _, sentence := range sentences {

		words := segmenter.Cut(sentence)

		for _, word := range words {
			s.index += 1
			s.wordCount += 1
			s.wordIndex[word] = append(s.wordIndex[word], s.index)
		}
	}

	// printTable(data)

	fmt.Printf("There are %d words in this paragraph\n", s.wordCount)
}

func (s *SearchEngine) SearchInput(input string) string {

	fmt.Printf("=== Searching of %s start ===\n", input)

	var result string

	input = strings.TrimSpace(input)

	searchResult, err := s.searchOccurence(input)

	if err != nil {
		return fmt.Sprintf("Word: %s not found", input)
	}

	result = fmt.Sprintf("The word %s appeared in %d", input, searchResult)
	return result
}
