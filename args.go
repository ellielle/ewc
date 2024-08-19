package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
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

	// flag.Arg(0) will return the filename, if any is provided
	// otherwise, it will attempt to read from stdin when no file is given
	files := flag.Args()

	var text []string

	// TODO: add support for multiple files in most functions, replace files[0] with loop

	// TODO: add prettier help output than PrintDefaults()

	// 	f, err := os.Open(file)
	// 	if err != nil {
	// 		return 0, err
	// 	}
	// 	defer f.Close()
	//
	// 	// create a new scanner and scan each line
	// 	// using a for loop
	// 	// add the line count to the count variable
	// 	scanner := bufio.NewScanner(f)
	// 	count := 0
	// 	for scanner.Scan() {
	// 		count++
	// 	}

	if len(files) == 0 && flag.Arg(0) == "" {
		// create a new scanner to read os.Stdin
		// and check if we have piped data
		fmt.Println("inside pipe code")
		scanner := bufio.NewScanner(os.Stdin)

		// if the text is piped in, read all of it into a []string
		for scanner.Scan() {
			text = append(text, scanner.Text())
		}
		if len(text) == 0 {
			flag.PrintDefaults()
			return
		}
		fmt.Println("byteCount here")

		return
	}

	if len(files) != 0 && flag.Arg(0) != "" {
		// open the given file and read it into a []string
		// 		file, err := os.Open(flag.Arg(0))
		// 		if err != nil {
		// 			log.Fatal(err)
		// 		}
		// 		defer file.Close()
		//
		// 		scanner := bufio.NewScanner(file)
		// 		for scanner.Scan() {
		// 			text = append(text, scanner.Text())
		// 		}

		// get the line count, byte count, and word count
		// and then output them with the file name
		lc, err := lineCount(flag.Arg(0))
		if err != nil {
			log.Fatal(err)
		}
		bc, err := byteCount(flag.Arg(0))
		if err != nil {
			log.Fatal(err)
		}
		wc, err := wordCount(flag.Arg(0))
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%v %v %v %v\n", lc, wc, bc, flag.Arg(0))
		return
	}

	fmt.Println("going down to the flag calls")

	if *help || *helpLong {
		flag.PrintDefaults()
		return
	}

	if *bytes || *bytesLong {
		bc, err := byteCount(flag.Arg(0))
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%v %v\n", bc, flag.Arg(0))
		return
	}

	if *chars || *charsLong {
		cc, err := charCount(flag.Arg(0))
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%v %v\n", cc, flag.Arg(0))
		return
	}

	if *lines || *linesLong {
		lc, err := lineCount(flag.Arg(0))
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%v %v\n", lc, flag.Arg(0))
		return
	}

	if *maxLines || *maxLinesLong {
		// maxLines()
		return
	}

	if *words || *wordsLong {
		wc, err := wordCount(flag.Arg(0))
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%v %v\n", wc, flag.Arg(0))
		return
	}

	if *version || *versionLong {
		// version()
		return
	}
}
