package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(3)

	go readFile(&wg, "/Users/nkanagar/Downloads/pg20417.txt")
	go readFile(&wg, "/Users/nkanagar/Downloads/pg4300.txt")
	go readFile(&wg, "/Users/nkanagar/Downloads/5000-8.txt")

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
