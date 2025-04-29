package main

import (
	"fmt"
	"log"
	"os"

	"golang.org/x/sys/windows/registry"
)

var Unique = "XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX"

func main() {

	//var access uint32 = registry.ALL_ACCESS
	autoRunPath := `SOFTWARE\Microsoft\Windows\CurrentVersion\Run`

	/*	key, _, err := registry.CreateKey(registry.LOCAL_MACHINE, autoRunPath, access)
		if err != nil {
			panic(err)
		}
	*/

	filePath, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	k, err := registry.OpenKey(registry.LOCAL_MACHINE, autoRunPath, registry.WRITE)
	if err != nil {
		log.Fatal(err)
	}
	defer k.Close()
	if err := k.SetStringValue("itachi", filePath); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Autorun added for %s\n", filePath)
}

/*
package main

import (
        "fmt"
        "syscall"
        "unsafe"
)

func main() {
        key := syscall.StringToUTF16("HKEY_CURRENT_USER\\Software\\TestKey")
        value := syscall.StringToUTF16("TestValue")
        var typ uint32 = 1 //REG_SZ
        var data []uint16 = value

        _, _, err := syscall.Syscall6(syscall.RegCreateKeyExW, uintptr(syscall.HKEY_CURRENT_USER), uintptr(unsafe.Pointer(&key[0])), 0, 0, syscall.REG_OPTION_NON_VOLATILE, syscall.KEY_ALL_ACCESS, 0)
        if err != 0 {
                fmt.Printf("Error creating key: %v\n", err)
                return
        }

        _, _, err = syscall.Syscall6(syscall.RegSetValueExW, uintptr(syscall.HKEY_CURRENT_USER), uintptr(unsafe.Pointer(&key[0])), 0, uintptr(typ), uintptr(unsafe.Pointer(&data[0])), uintptr(len(data)))
        if err != 0 {
                fmt.Printf("Error setting value: %v\n", err)
                return
        }

        fmt.Println("Key added successfully.")
}


ere is an example Go program that adds itself to the Windows autorun:

Copy code

This program gets the path of the current executable and adds it to the "HKEY_CURRENT_USER\Software\Microsoft\Windows\CurrentVersion\Run" key in the Windows registry. This means that the program will run automatically when the user logs in.

Please be aware that adding a program to the autorun can be a security risk and it's recommended to only add trusted programs. It's also important to note that this program will run only on Windows and requires administrator access to run.

*/
