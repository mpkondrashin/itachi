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

const unique = "XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX"

//go:embed malware/*.exe
var malware embed.FS

func extractAll() {
	dir, err := malware.ReadDir("malware")
	if err != nil {
		panic(err)
	}
	for _, each := range dir {
		if !each.Type().IsRegular() {
			continue
		}
		fmt.Println(each.Name())
		f, err := malware.Open("malware/" + each.Name())
		if err != nil {
			panic(err)
		}
		stamp := time.Now().Format("20060102150405")
		targetFileName := stamp + "_" + each.Name()
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
}

func main() {
	extractAll()
}
