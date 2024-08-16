package main

import (
	"bufio"
	"os"
)

func charCount(file string) (int, error) {
	f, err := os.Open(file)
	if err != nil {
		return 0, err
	}
	defer f.Close()

	// set character count to 0, create a new scanner
	// and count the characters in each iteration over the scanner
	cc := 0
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanRunes)

	for scanner.Scan() {
		cc++
	}

	return cc, nil
}
