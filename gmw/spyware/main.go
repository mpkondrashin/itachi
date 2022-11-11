package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

var Unique = "XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX"

func main() {
	fmt.Printf("Demo Spyware (%s)\n", Unique)
	url := "http://wrs21.winshipway.com/"
	fmt.Printf("Get: %s\n", url)
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	html, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Got\n%s\n", html)
}
