package main

import (
	"fmt"
	"log"

	"golang.org/x/sys/windows/registry"
)

func main() {
	k, err := registry.OpenKey(registry.LOCAL_MACHINE, `SOFTWARE\Microsoft\Windows NT\CurrentVersion`, registry.QUERY_VALUE)
	if err != nil {
		log.Fatal(err)
	}
	defer k.Close()

	s, _, err := k.GetStringValue("SystemRoot")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Windows system root is %q\n", s)
}
