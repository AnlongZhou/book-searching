/*
TODO: Searching for a word and return it index
TODO: Wails.io framework
*/
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/go-ego/gse"
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

type Data struct {
	wordIndex map[string][]int
	index     int
}

func newWordIndex() map[string][]int {
	return make(map[string][]int)
}

func newData() Data {
	return Data{
		wordIndex: newWordIndex(),
		index:     0,
	}
}

func printTable(data Data) {
	for word, indices := range data.wordIndex {
		fmt.Printf("%s appears in: %d\n", word, indices)
	}
}

func main() {
	infile, err := os.Open("input.txt")
	checkError(err)

	segmenter := newSegmenter()
	reader := bufio.NewReader(infile)

	data := newData()

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}

		sentences := splitLine(line)

		for _, sentence := range sentences {

			words := segmenter.Cut(sentence)

			for _, word := range words {
				data.index += 1
				data.wordIndex[word] = append(data.wordIndex[word], data.index)
			}
		}

		printTable(data)

		fmt.Println("Finished")

	}
}
