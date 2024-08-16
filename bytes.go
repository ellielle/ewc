package main

import (
	"os"
)

func byteCount(file string) (int64, error) {
	fileInfo, err := os.Stat(file)
	if err != nil {
		return 0, err
	}

	return fileInfo.Size(), nil
}
