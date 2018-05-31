package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
	"sync/atomic"
)

// create uint64 to be used for atomic counting
var maxFileReaders uint64

func main() {
	// list of all files to be processed
	filesToRead := []string{"../5000-8.txt", "../pg4300.txt", "../pg20417.txt"}
	wordCountTable := make(map[string]int)

	// set the value to maxFileReader in atomic way
	atomic.AddUint64(&maxFileReaders, uint64(len(filesToRead)))

	var wg sync.WaitGroup
	wg.Add(1)

	lineByLineDataChannel := make(chan string)
	wordsChannel := make(chan []string)

	go wordCounter(&wg, wordsChannel, wordCountTable)

	go lineParser(lineByLineDataChannel, wordsChannel)

	for _, name := range filesToRead {
		go readFileContentLineByLine(lineByLineDataChannel, &maxFileReaders, name)
	}

	wg.Wait()

	for k, v := range wordCountTable {
		fmt.Println(k, v)
	}
}

func wordCounter(wg *sync.WaitGroup, wordsChannel chan []string, wordCountTable map[string]int) {
	defer wg.Done()

	for words := range wordsChannel {
		for _, word := range words {
			var cleanWord = strings.TrimSpace(word)
			if len(cleanWord) > 0 {
				wordCountTable[cleanWord] = wordCountTable[cleanWord] + 1
			}
		}
	}
}

func lineParser(lineByLineDataChannel chan string, wordsChannel chan []string) {
	for line := range lineByLineDataChannel {
		wordsChannel <- strings.Split(line, " ")
	}
	close(wordsChannel)
}

func readFileContentLineByLine(lineByLineDataChannel chan string,
	maxFileReaders *uint64, name string) {

	fileHandle, err := os.Open(name)

	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("File ", name, " not present")
		} else {
			fmt.Println("Unexpected error happened during file open for file", name)
		}

		return
	}

	defer fileHandle.Close()

	scanner := bufio.NewScanner(fileHandle)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if len(line) > 0 {
			lineByLineDataChannel <- line
		}
	}

	// Once file read complete, reduce the count of maxFileReaders using Atomic Add operation
	// after that immediatly check for the value. If the value is 0, then this function is the
	// last guy who completed the operation. if so, you are privileged to close the channel
	atomic.AddUint64(maxFileReaders, ^uint64(0))
	if atomic.LoadUint64(maxFileReaders) == 0 {
		close(lineByLineDataChannel)
	}
}
