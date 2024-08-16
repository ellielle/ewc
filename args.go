package main

import (
	"flag"
	"fmt"
	"log"
)

func parseArgs() {
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

	// TODO: add support for multiple files in most functions, replace files[0] with loop

	// TODO: add prettier help output than PrintDefaults()

	if len(files) == 0 {
		flag.PrintDefaults()
		return
	}

	if *help || *helpLong {
		flag.PrintDefaults()
		return
	}

	if *bytes || *bytesLong {
		bc, err := byteCount(files[0])
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%v %v\n", bc, files[0])
	}

	if *chars || *charsLong {
		cc, err := charCount(files[0])
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%v %v\n", cc, files[0])
	}

	if *lines || *linesLong {
		lc, err := lineCount(files[0])
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%v %v\n", lc, files[0])

	}

	if *maxLines || *maxLinesLong {
		// maxLines()
	}

	if *words || *wordsLong {
		wc, err := wordCount(files[0])
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%v %v\n", wc, files[0])

	}

	if *version || *versionLong {
		// version()
	}
}
