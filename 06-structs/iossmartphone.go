package main

import "fmt"

type IOSSmartPhone struct {
	smartPhoneConfiguration SmartPhoneConfiguration
	phoneNumber             string
	year                    int
	appsInstalled           []string
}

func (iosSmartPhone IOSSmartPhone) acceptIncomingCalls() {
	fmt.Println("Incoming call to ", iosSmartPhone.phoneNumber)
}

func (iosSmartPhone IOSSmartPhone) makeOutgoingCalls(targetNumber string) {
	fmt.Println("Outgoing call from ", iosSmartPhone.phoneNumber)
}

func (iosSmartPhone *IOSSmartPhone) installApp(appName string) {
	iosSmartPhone.appsInstalled = append(iosSmartPhone.appsInstalled, appName)
	//fmt.Println("from inside Install App ", iosSmartPhone)
}
