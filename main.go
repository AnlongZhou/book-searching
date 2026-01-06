package main

func main() {
	searchEngine := newSearchEngine()
	segmenter := newSegmenter()

	searchEngine.readFile(searchEngine, segmenter)
	searchEngine.searchInput(searchEngine)
}
