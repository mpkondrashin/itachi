package main

import (
	"fmt"
	"os"
	"syscall"
	"unsafe"
)

func main() {
	path, _ := os.Executable()
	key := syscall.StringToUTF16("HKEY_CURRENT_USER\\Software\\Microsoft\\Windows\\CurrentVersion\\Run")
	value := syscall.StringToUTF16(path)
	var typ uint32 = 1 //REG_SZ
	var data []uint16 = value

	_, _, err := syscall.Syscall6(syscall.RegCreateKeyExW,
		uintptr(syscall.HKEY_CURRENT_USER),
		uintptr(unsafe.Pointer(&key[0])),
		0,
		0,
		syscall.REG_OPTION_NON_VOLATILE,
		syscall.KEY_ALL_ACCESS,
		0)
	if err != 0 {
		fmt.Printf("Error creating key: %v\n", err)
		return
	}

	_, _, err = syscall.Syscall6(syscall.RegSetValueExW,
		uintptr(syscall.HKEY_CURRENT_USER),
		uintptr(unsafe.Pointer(&key[0])),
		0,
		uintptr(typ),
		uintptr(unsafe.Pointer(&data[0])),
		uintptr(len(data)))
	if err != 0 {
		fmt.Printf("Error setting value: %v\n", err)
		return
	}
	fmt.Println("%s added to autorun successfully.", path)
}
