package main

import (
	"fmt"
)

func main() {
	// first type of initializing user defined type
	iphone6 := IOSSmartPhone{
		SmartPhoneConfiguration{"Apple A8", "5Inch", "16GB", "32GB"},
		"123456", 2017, make([]string, 1)}
	iphone6.smartPhoneConfiguration.storage = "32GB"
	fmt.Println("Before app install", iphone6)

	iphone6.installApp("GMail")
	fmt.Println("After app install", iphone6)

	iphone7 := IOSSmartPhone{
		smartPhoneConfiguration: SmartPhoneConfiguration{
			ram:        "16GB",
			storage:    "64GB",
			screenSize: "5.5Inch",
			cpu:        "Apple A10 Fusion",
		},
		phoneNumber:   "3333333",
		year:          2017,
		appsInstalled: make([]string, 1),
	}

	fmt.Println(iphone7)
}
