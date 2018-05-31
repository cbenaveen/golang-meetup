package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Before Greet method call")
	go greet("Naveen")
	fmt.Println("After Greet method call")

	time.Sleep(10 * time.Millisecond)
}

func greet(name string) {
	fmt.Println("Hello! ", name)
	go welcome("Golang Class")
}

func welcome(name string) {
	fmt.Println("Welcome to ", name)
}
