package main

import (
	"flag"
	"fmt"
	"log"
)

func main() {
	// set up flags
	// short and long name flags for byte count
	bytes := flag.Bool("c", false, "print the byte counts")
	bytesLong := flag.Bool("bytes", false, "print the byte counts")

	// short and long name flags for char count
	chars := flag.Bool("m", false, "print the character counts")
	charsLong := flag.Bool("chars", false, "print the character counts")

	// short and long name flags for newline count
	lines := flag.Bool("l", false, "print the newline counts")
	linesLong := flag.Bool("lines", false, "print the newline counts")

	// short and long name flags for the max display width
	maxLines := flag.Bool("L", false, "print the max display width")
	maxLinesLong := flag.Bool("max-line-length", false, "print the max display width")

	// short and long name flags for the word count
	words := flag.Bool("w", false, "print the word counts")
	wordsLong := flag.Bool("words", false, "print the word counts")

	// short and long name flags for the help menu
	help := flag.Bool("h", false, "print this help and exit")
	helpLong := flag.Bool("help", false, "print this help and exit")

	// short and long name flags for the version
	version := flag.Bool("v", false, "print version information and quit")
	versionLong := flag.Bool("version", false, "print version information and quit")

	// parse the flags so we can see what we're doing
	flag.Parse()

	// 0, 1 or multiple files
	// it will read from stdin when no file is given
	files := flag.Args()

	if len(files) == 0 {
		flag.PrintDefaults()
	}

	if *bytes || *bytesLong {
		// TODO: add support for multiple files, replace files[0] with loop
		byteCount, err := byteCount(files[0])
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%v %v\n", byteCount, files[0])
	}

	if *chars || *charsLong {
		charCount()
	}

	if *lines || *linesLong {
		// lineCount()
	}

	if *maxLines || *maxLinesLong {
		// maxLines()
	}

	if *words || *wordsLong {
		// wordCount()
	}

	if *help || *helpLong {
		flag.PrintDefaults()
	}

	if *version || *versionLong {
		// version()
	}
}
