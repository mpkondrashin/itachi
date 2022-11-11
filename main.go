package main

import (
	"embed"
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	"github.com/google/uuid"
)

// go : embed malware/*.exe

const unique = "XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX"

//go:embed gmw/*.exe
var malware embed.FS

func extractAll() {
	dir, err := malware.ReadDir("gmw")
	if err != nil {
		panic(err)
	}
	for i, each := range dir {
		if !each.Type().IsRegular() {
			continue
		}
		f, err := malware.Open("gmw/" + each.Name())
		if err != nil {
			panic(err)
		}
		stamp := time.Now().Format("20060102150405")
		targetFileName := stamp + "_" + each.Name()
		fmt.Printf("%d: %s\n", i+1, targetFileName)
		extract(f, targetFileName)
	}
}

func extract(source io.Reader, targetFileName string) {
	data, err := io.ReadAll(source)
	if err != nil {
		panic(err)
	}
	u := uuid.New().String()
	uniqueData := strings.Replace(string(data), unique, u, 1)
	if err := os.WriteFile(targetFileName, []byte(uniqueData), 0755); err != nil {
		panic(err)
	}
	//fmt.Printf("Sample: %s\n", targetFileName)
}

func main() {
	fmt.Println("Itachi Samples Generator")
	extractAll()
	fmt.Println("Done")
}
