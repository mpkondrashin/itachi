/***************************************************************
*
* Itachi (c) 2022 by Michael Kondrashin (mkondrashin@gmail.com)
* Copyright under MIT Lincese. Please see LICENSE file for details
*
* main.go - main Itachi file
*
***************************************************************/

package main

import (
	"embed"
	"fmt"
	"io"
	"os"
	"strings"

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
	uniqueStr := uuid.New().String()
	for i, each := range dir {
		if !each.Type().IsRegular() {
			continue
		}
		f, err := malware.Open("gmw/" + each.Name())
		if err != nil {
			panic(err)
		}
		//stamp := time.Now().Format("20060102150405")
		targetFileName := uniqueStr + "_" + each.Name()
		fmt.Printf("%d: %s\n", i+1, targetFileName)
		extract(f, uniqueStr, targetFileName)
	}
}

func extract(source io.Reader, uniqueStr string, targetFileName string) {
	data, err := io.ReadAll(source)
	if err != nil {
		panic(err)
	}
	uniqueData := strings.Replace(string(data), unique, uniqueStr, 1)
	if err := os.WriteFile(targetFileName, []byte(uniqueData), 0755); err != nil {
		panic(err)
	}
	//fmt.Printf("Sample: %s\n", targetFileName)
}

func main() {
	fmt.Println("Itachi Samples Generator")
	extractAll()
	fmt.Println("Done")
	fmt.Print(warning)
}

var warning = `
!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
!!!!!!!!!!!!!!!!! WARNING !!!!!!!!!!!!!!!!!!!
!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
!!                                         !!
!!  Do not run these executables locally!  !!
!!                                         !!
!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
`
