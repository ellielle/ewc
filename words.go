package main

import (
	"bufio"
	"os"
)

// wordCount opens `file`, creates a scanner and loops
// over each line. Then it splits the line into words and adding
// the length of the slice to the word count
func wordCount(file string) (int, error) {
	f, err := os.Open(file)
	if err != nil {
		return 0, err
	}
	defer f.Close()

	// set word count to 0, create a new scanner
	// and count the words in each iteration over the scanner
	wc := 0
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		wc++
	}

	return wc, nil
}
