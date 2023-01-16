package main

import (
	"fmt"
	"log"
	"os"

	"golang.org/x/sys/windows/registry"
)

func main() {
	filePath, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	AutoRunPath := `SOFTWARE\Microsoft\Windows\CurrentVersion\Run`
	k, err := registry.OpenKey(registry.LOCAL_MACHINE, AutoRunPath, registry.WRITE)
	if err != nil {
		log.Fatal(err)
	}
	defer k.Close()
	if err := k.SetStringValue("itachi", filePath); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Autorun added for %s\n", filePath)
}
