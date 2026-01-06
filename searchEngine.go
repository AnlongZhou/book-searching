/*
TODO: Find occurence of a word and sentences contain it
*/
package main

import (
	"bufio"
	"fmt"
	"github.com/go-ego/gse"
	"log"
	"os"
	"strings"
)

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func newSegmenter() gse.Segmenter {

	var segmenter gse.Segmenter
	err := segmenter.LoadDict()

	checkError(err)

	return segmenter

}

func splitLine(line string) []string {

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

type SearchEngine struct {
	wordIndex map[string][]int
	index     int
	wordCount int
}

func newSearchEngine() SearchEngine {
	return SearchEngine{
		wordIndex: make(map[string][]int),
		index:     0,
		wordCount: 0,
	}
}

func printTable(data SearchEngine) {
	for word, indices := range data.wordIndex {
		fmt.Printf("%s appears in: %d\n", word, indices)
	}
}

func (s *SearchEngine) searchOccurence(data SearchEngine, target string) ([]int, error) {
	var result []int
	var err error

	result = data.wordIndex[target]

	if result == nil {
		err = fmt.Errorf("Error 404: Word not found")
	}

	return result, err
}

func (s *SearchEngine) readFile(data SearchEngine, segmenter gse.Segmenter) {

	infile, err := os.Open("input.txt")
	checkError(err)

	for {
		reader := bufio.NewReader(infile)

		line, err := reader.ReadString('\n')
		if err != nil {
			return
		}

		sentences := splitLine(line)

		for _, sentence := range sentences {

			words := segmenter.Cut(sentence)

			for _, word := range words {
				data.index += 1
				data.wordCount += 1
				data.wordIndex[word] = append(data.wordIndex[word], data.index)
			}
		}

		// printTable(data)

		fmt.Printf("There are %d words in this paragraph\n", data.wordCount)
	}
}

func (s *SearchEngine) searchInput(data SearchEngine) {
	for {
		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')
		checkError(err)

		input = strings.TrimSpace(input)

		if input == "" {
			break
		}

		searchResult, err := data.searchOccurence(data, input)

		if err != nil {
			fmt.Println(err)
			continue
		}

		fmt.Printf("\nThe word %s appeared in %d\n", input, searchResult)
		fmt.Printf("Find %d \n\n", len(searchResult))
	}
}
