/*
This script downloads the raw text of a book from Project Gutenberg
and calculates its "readability index".

Example usage (103 is "Around the World in 80 Days"):

  go run readability.go 103

  readability 103
*/

package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"net/http"
	"os"
	"strconv"
	"strings"
	"unicode"
)

// ARI displays the "automated readability index" of the text that can
// be read from the provided body.
func ARI(body io.ReadCloser) {

	scanner := bufio.NewScanner(body)

	var nSentences int
	var nWords int
	var nCharacters int

	for scanner.Scan() {

		line := scanner.Bytes()

		nSentences += strings.Count(string(line), ".")
		nSentences += strings.Count(string(line), "?")
		nSentences += strings.Count(string(line), "!")

		for _, x := range line {
			if unicode.IsLetter(rune(x)) {
				nCharacters++
			}
		}

		nWords += strings.Count(string(line), " ")
	}

	ari := 4.71 * float64(nCharacters) / float64(nWords)
	ari += 0.5 * float64(nWords) / float64(nSentences)
	ari -= 21.43

	fmt.Printf("Number of sentences:  %10d\n", nSentences)
	fmt.Printf("Number of characters: %10d\n", nCharacters)
	fmt.Printf("Number of words:      %10d\n", nWords)
	fmt.Printf("ARI:                           %.0f\n", math.Ceil(ari))
}

func main() {

	if len(os.Args) != 2 {
		os.Stderr.WriteString("usage: readability id_number\n")
		os.Exit(0)
	}

	// Get the id for the book.
	id, err := strconv.Atoi(os.Args[1])
	if err != nil {
		panic(err)
	}

	// https://www.gutenberg.org/files/19513/19513.txt
	ur := fmt.Sprintf("https://www.gutenberg.org/cache/epub/%d/pg%d.txt", id, id)
	resp, err := http.Get(ur)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	ARI(resp.Body)
}
