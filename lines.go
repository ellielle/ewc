package main

import (
	"bufio"
	"os"
)

// lineCount opens `file`, and scans over each line
// while adding to a counter with each line
func lineCount(file string) (int, error) {
	f, err := os.Open(file)
	if err != nil {
		return 0, err
	}
	defer f.Close()

	// create a new scanner and scan each line
	// using a for loop
	// add the line count to the count variable
	scanner := bufio.NewScanner(f)
	count := 0
	for scanner.Scan() {
		count++
	}

	return count, nil
}
