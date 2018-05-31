package main

import (
	"fmt"
)

func main() {
	names := [7]string{"Naveen", "Deepak", "Madhave", "Sagar", "Smitha", "Sandeep", "Vinay"}
	//fmt.Println("The names array is", names)

	for i := 0; i < len(names); i++ {
		fmt.Println("Element at position ", i, " is ", names[i])
	}

	for index, value := range names {
		fmt.Println("Array element and index ", index, " is ", value)
	}

	for index := range names {
		fmt.Println("index (alone) ", index)
	}

	for _, value := range names {
		fmt.Println("value (alone) ", value)
	}

}
