/***************************************************************
*
* Itachi (c) 2022 by Mikhail Kondrashin (mkondrashin@gmail.com)
*
* main.go - dropper - write eicar.com to locally
*
***************************************************************/

package main

import (
	"fmt"
	"os"
)

var reicar = "*H+H$!ELIF-TSET-SURIVITNA-DRADNATS-RACIE$}7)CC7)^P(45XZP\\4[PA@%P!O5X"

const fileName = "eicar.com"

var Unique = "XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX"

func main() {
	fmt.Printf("Demo Dropper (%s)\n", Unique)
	f, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	for i := len(reicar) - 1; i >= 0; i-- {
		if _, err := f.Write([]byte(string(reicar[i]))); err != nil {
			panic(err)
		}
	}
	fmt.Println("Done")
}
