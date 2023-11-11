package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strings"
)

const (
	fileName = "./sherlock.txt"
)

var wordRegexp = regexp.MustCompile(`[a-zA-Z]+`)

// What is the most common world in sherlock.txt?
// Word frequency
func wordFrequency(reader io.Reader) (map[string]int, error) {
	wordToCount := make(map[string]int)
	scanner := bufio.NewScanner(reader)
	linesCount := 0

	for scanner.Scan() {

		currentLine := scanner.Text()
		foundWords := wordRegexp.FindAllString(currentLine, -1)

		fmt.Printf("Current handling line number: %d\n", linesCount)
		fmt.Printf("Current line words found: %d\n\n", len(foundWords))

		for _, word := range foundWords {
			wordLower := strings.ToLower(word)
			currentWordCount, ok := wordToCount[wordLower]
			if ok {
				wordToCount[wordLower] = currentWordCount + 1
			} else {
				wordToCount[wordLower] = 1
			}
		}

		linesCount++
	}

	fmt.Printf("Lines handled: %d\n", linesCount)

	if scanError := scanner.Err(); scanError != nil {
		return nil, scanError
	}

	return wordToCount, nil
}

func findMaxOftenWord(wordsToCount map[string]int) (string, error) {
	if len(wordsToCount) == 0 {
		return "", fmt.Errorf("empty map passed")
	}

	maxCount := 0
	maxWorld := ""
	for word, count := range wordsToCount {
		if count > maxCount {
			maxCount = count
			maxWorld = word
		}
	}

	return maxWorld, nil
}

func main() {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("Can't open file %s, error accure %s", fileName, err)
	}

	res, err := wordFrequency(file)

	if err != nil {
		log.Fatalf("Something went wrong: %s", err)
	}
	resultWord, err := findMaxOftenWord(res)
	if err != nil {
		log.Fatalf("Simething went wrong: %s", err)
	}

	fmt.Printf("Word with maximum accurance: %s(%d)", resultWord, res[resultWord])

	defer func() {
		fileCloseError := file.Close()
		if fileCloseError != nil {
			fmt.Println(fileCloseError)
		}
	}()
}
