package main

import (
	"os"
)

// byteCount grabs the stats for `file` and return the size in bytes
func byteCount(file string) (int64, error) {
	fileInfo, err := os.Stat(file)
	if err != nil {
		return 0, err
	}

	return fileInfo.Size(), nil
}
