package main

import (
	"fmt"
)

func listAllApps(smartPhone SmartPhone) {
	for i, k := range smartPhone.getAllApps() {
		fmt.Println(i, k)
	}
}

func makeOutgoingCall(smartPhone SmartPhone,
	targetPhoneNumber string) {
	smartPhone.makeCall(targetPhoneNumber)
}

func installApp(smartPhone SmartPhone,
	appName string) {
	smartPhone.installApp(appName)
}

type SmartPhone interface {
	makeCall(targetPhoneNumber string)
	installApp(appName string) (bool, error)
	getAllApps() []string
}

type IOSSmartphone struct {
	phoneNumber   string
	cpu           string
	screenSize    string
	ram           string
	storage       string
	year          int
	appsInstalled []string
}

func (iosSmartphone IOSSmartphone) getAllApps() []string {
	return iosSmartphone.appsInstalled
}

func (iosSmartphone IOSSmartphone) makeCall(targetPhoneNumber string) {
	fmt.Println("Calling to target number ", targetPhoneNumber,
		" from source number", iosSmartphone.phoneNumber)
}

func (iosSmartphone IOSSmartphone) installApp(appname string) (bool, error) {
	iosSmartphone.appsInstalled = append(iosSmartphone.appsInstalled, appname)
	fmt.Println("Added a", appname, " to phonenumber ", iosSmartphone.phoneNumber)
	fmt.Println("All apps", iosSmartphone.appsInstalled)
	return true, nil
}

func (androidphone Androidphone) installApp(appname string) (bool, error) {
	androidphone.appsInstalled = append(androidphone.appsInstalled, appname)
	fmt.Println("Added a", appname, " to phonenumber ", androidphone.phoneNumber)
	return true, nil
}

func (androidphone Androidphone) getAllApps() []string {
	return androidphone.appsInstalled
}

type Androidphone struct {
	phoneNumber   string
	cpu           string
	screenSize    string
	ram           string
	storage       string
	year          int
	appsInstalled []string
}

func (androidphone Androidphone) makeCall(targetPhoneNumber string) {
	fmt.Println("Calling to target number ", targetPhoneNumber,
		" from source number", androidphone.phoneNumber)
}

func main() {
	iphone6 := IOSSmartphone{"0001", "AppleAB", "5inch", "16gb", "32gb", 2017, make([]string, 1)}
	makeOutgoingCall(iphone6, "234567")
	installApp(&iphone6, "IOS Email")
	installApp(&iphone6, "IOS Email2")
	installApp(&iphone6, "IOS Email3")

	redmi4 := Androidphone{"0002", "Snapdragon", "5.5inch", "16gb", "32gb", 2017, make([]string, 1)}
	makeOutgoingCall(redmi4, "234567")
	installApp(&redmi4, "Google EMail")

	listAllApps(&iphone6)
	listAllApps(&redmi4)
}
