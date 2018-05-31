package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
)

func main() {
	wordCountTable := make(map[string]int)
	var wg sync.WaitGroup
	var mux sync.Mutex

	wg.Add(3)

	go readFile(&wg, &mux, wordCountTable, "../pg20417.txt")
	go readFile(&wg, &mux, wordCountTable, "../pg4300.txt")
	go readFile(&wg, &mux, wordCountTable, "../5000-8.txt")

	wg.Wait()

	for k, v := range wordCountTable {
		fmt.Println(k, v)
	}
}

func readFile(wg *sync.WaitGroup, mux *sync.Mutex, wordCountTable map[string]int, name string) int {
	defer wg.Done()

	var total = 0

	fileHandle, err := os.Open(name)
	if err != nil {
		fmt.Println("File doesnt exists", name)
		return total
	}

	scanner := bufio.NewScanner(fileHandle)

	for scanner.Scan() {
		var line = strings.Trim(scanner.Text(), " ")
		if len(line) > 0 {
			var words = strings.Split(line, " ")
			mux.Lock()
			for _, word := range words {
				if len(strings.Trim(word, "")) > 0 {
					var wordCount = wordCountTable[word]
					wordCount = 1 + wordCount
					wordCountTable[word] = wordCount
				}
			}
			mux.Unlock()
		}
	}

	fmt.Println("Total number of lines in file ", name, total)
	return total
}
