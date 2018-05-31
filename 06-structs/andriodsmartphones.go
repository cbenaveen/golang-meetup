package main

import "fmt"

type AndriodSmartPhone struct {
	SmartPhoneConfiguration
	phoneNumber   string
	year          int
	appsInstalled []string
}

func (andriodSmartPhone AndriodSmartPhone) acceptIncomingCalls() {
	fmt.Println("Incoming call to ", andriodSmartPhone.phoneNumber)
}

func (andriodSmartPhone AndriodSmartPhone) makeOutgoingCalls(targetNumber string) {
	fmt.Println("Outgoing call from ", andriodSmartPhone.phoneNumber)
}

func (andriodSmartPhone *AndriodSmartPhone) installApp(appName string) {
	andriodSmartPhone.appsInstalled = append(andriodSmartPhone.appsInstalled, appName)
}

func justForFunc() {

}
