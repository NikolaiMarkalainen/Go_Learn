package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	fileStats := make(map[string][]string)
	// i have counts that holds array of strings and amount of times it appears
	// fileName and the counts held in one place
	// compare this to previous instances ?
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, fileStats)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, fileStats)
			f.Close()
		}
	}
	for line, n := range counts {
		var files string
		for filename, word := range fileStats {
			for _, w := range word {
				if n > 1 && line == w {
					files += filename + " "
					break
				}
			}
		}
		if n > 1 {
			fmt.Printf("%d\t%s \t %s\n", n, line, files)
		}
	}
}
func countLines(f *os.File, counts map[string]int, fileStats map[string][]string) {
	// we count lines here
	input := bufio.NewScanner(f)
	var lines []string
	for input.Scan() {
		counts[input.Text()]++
		lines = append(lines, input.Text())
	}
	fileStats[f.Name()] = lines

	// NOTE: ignoring potential errors from input.Err()
}
