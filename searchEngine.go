// 2026/1/17: added sentences searching.
// TODO:Concurrency implementation

package main

import (
	"context"
	"fmt"
	"strings"

	"github.com/go-ego/gse"
)

type SearchEngine struct {
	ctx       context.Context
	segmenter gse.Segmenter
	wordIndex map[string][]int
	book      []string
	index     int
	wordCount int
}

func newSearchEngine() *SearchEngine {
	var s *SearchEngine

	s = &SearchEngine{
		wordIndex: make(map[string][]int),
		segmenter: s.newSegmenter(),
		book:      []string{},
		index:     0,
		wordCount: 0,
	}

	return s
}

func (s *SearchEngine) startup(ctx context.Context) {
	s.ctx = ctx
	s.readFile()
	s.initialBook()
}

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
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
	line = strings.ReplaceAll(line, "」", "。")
	line = strings.ReplaceAll(line, "「", "。")

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

func (s *SearchEngine) readFile() {

	infile := inputEmbed

	lines := strings.Split(infile, "\n")

	for index, line := range lines {
		sentences := s.splitLine(line)

		for _, sentence := range sentences {
			words := s.segmenter.Cut(sentence)

			for _, word := range words {
				s.index += 1
				s.wordCount += 1
				s.wordIndex[word] = append(s.wordIndex[word], index)
			}
		}
	}

	// printTable(data)

	fmt.Printf("There are %d words in this book\n", s.wordCount)
}

func (s *SearchEngine) initialBook() {
	book := inputEmbed
	sentences := strings.SplitSeq(book, "\n")

	for sentence := range sentences {
		s.book = append(s.book, sentence)
	}
}

func (s *SearchEngine) searchOccurence(target string) ([]int, error) {
	var result []int
	var err error

	result = s.wordIndex[target]

	if result == nil {
		err = fmt.Errorf("Error 404: Word %s not found", target)
	}

	return result, err
}

func (s *SearchEngine) findIntersection(listA []int, listB []int) []int {
	var interSection []int
	seen := make(map[int]bool)

	for _, element := range listA {
		seen[element] = true
	}

	for _, element := range listB {
		if seen[element] {
			interSection = append(interSection, element)
			delete(seen, element)
		}
	}

	return interSection
}

func (s *SearchEngine) SearchInput(input string) []string {

	if input == "" {
		return []string{"Input cannot be empty"}
	}

	fmt.Printf("=== Searching of %s start ===\n", input)

	// var result string

	input = strings.ReplaceAll(input, " ", "")
	fmt.Printf("Input: %s\n", input)

	searchIndex, err := s.searchOccurence(input)

	if err != nil {
		fmt.Println(err)
		fmt.Println("Splitting the input")
		newWords := s.segmenter.Cut(input)

		fmt.Printf("words after split: %s", newWords)

		if len(newWords) <= 1 {
			fmt.Printf("Word %s not found", input)
			return []string{"Not Found"}
		}

		intersection, err := s.searchOccurence(newWords[0])
		checkError(err)

		for i, word := range newWords {
			if i == 0 {
				continue
			}
			temp, err := s.searchOccurence(word)
			checkError(err)

			intersection = s.findIntersection(intersection, temp)

		}

		fmt.Printf("final occurence: %d", intersection)

		if len(intersection) == 0 {
			return []string{"Not Found"}
		}

		searchIndex = intersection
	}

	var searchResult []string

	for _, index := range searchIndex {
		searchResult = append(searchResult, s.book[index])
	}

	// result = fmt.Sprintf("The word %s appears in: %s", input, searchResult)
	return searchResult
}
