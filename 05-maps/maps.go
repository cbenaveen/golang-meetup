package main

import (
	"fmt"
)

func main() {
	fmt.Println("Map intro")

	companyDatabase := make(map[string]int)
	companyDatabase["Cisco"] = 1984
	companyDatabase["Google"] = 1998
	companyDatabase["FB"] = 2004
	companyDatabase["Amazon"] = 1994

	amazonYear, exists := companyDatabase["Amazon"]
	if exists {
		fmt.Println("Amazon founded year ", amazonYear)
	} else {
		fmt.Println("Amazon details not available", amazonYear)
	}

	twitterYear, exists := companyDatabase["Twitter"]
	if exists {
		fmt.Println("twitter founded year ", twitterYear)
	} else {
		fmt.Println("twitter details not available", twitterYear)
	}

	for key, value := range companyDatabase {
		fmt.Println("Company ", key, " was founded year ", value)
	}

	companySharePrice := map[string]int{
		"CSCO": 41,
	}

	for key, value := range companySharePrice {
		fmt.Println("Company Symbole", key, " share price ", value)
	}

}
