package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	if response, err := http.Get("https://google.com"); err == nil {
		body, err := ioutil.ReadAll(response.Body)
		fmt.Println("Response Body: ", body)
		fmt.Println("Response Error (if any): ", err)
	} else {
		fmt.Println("error: ", err)
	}
}
