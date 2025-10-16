package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// map from line → map of filenames → count
	counts := make(map[string]map[string]int)
	files := os.Args[1:]

	if len(files) == 0 {
		countLines(os.Stdin, counts, "stdin")
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, arg)
			f.Close()
		}
	}

	for line, fileMap := range counts {
		total := 0
		for _, n := range fileMap {
			total += n
		}
		if total > 1 {
			fmt.Printf("%d\t%s\t", total, line)
			for filename := range fileMap {
				fmt.Printf("%s ", filename)
			}
			fmt.Println()
		}
	}
}

func countLines(f *os.File, counts map[string]map[string]int, filename string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		line := input.Text()
		if counts[line] == nil {
			counts[line] = make(map[string]int)
		}
		counts[line][filename]++
	}
	// ignoring potential errors from input.Err()
}
