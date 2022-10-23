package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	//Defining a boolean flag -l to count lines.
	lines := flag.Bool("l", false, "Count lines")
	//Defining a boolean flag -b to count number of bytes.
	bytes := flag.Bool("b", false, "Count bytes")
	//Parsing the flags provided by the user.
	flag.Parse()
	//Calling the count function to count the words(or lines)
	//received from the standard input and printing it out.
	fmt.Println(count(os.Stdin, *lines, *bytes))
}

func count(r io.Reader, countLines bool, countBytes bool) int {
	scanner := bufio.NewScanner(r)
	//If the Count Lines flag is not set then we want to count words or if Count Bytes is set we define
	//scanner split type to words, and bytes respectively(default is lines).
	if !countLines {
		if countBytes {
			scanner.Split(bufio.ScanBytes)
		} else {
			scanner.Split(bufio.ScanWords)
		}
	}
	c := 0
	for scanner.Scan() {
		c++
	}
	return c
}
