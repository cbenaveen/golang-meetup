package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sync"
)

func main() {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(dir)

	var wg sync.WaitGroup

	wg.Add(3)

	go readFile(&wg, "../pg20417.txt")
	go readFile(&wg, "../pg4300.txt")
	go readFile(&wg, "../5000-8.txt")

	//var message = sayHello("Naveen")

	wg.Wait()
}

func sayHello(name string) string {
	return "Hello! name"
}

func readFile(wg *sync.WaitGroup, name string) int {
	defer wg.Done()

	var total = 0

	fileHandle, err := os.Open(name)
	if err != nil {
		fmt.Println("File doesnt exists", name)
		return total
	}

	scanner := bufio.NewScanner(fileHandle)

	for scanner.Scan() {
		total = total + 1
		//fmt.Println(scanner.Text())
	}

	fmt.Println("Total number of lines in file ", name, total)
	return total
}
