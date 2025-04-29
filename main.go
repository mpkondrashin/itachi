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
	"fmt"

	"github.com/mpkondrashin/itachi/pkg/generate"
)

func main() {
	fmt.Println("Itachi Samples Generator")
	generate.ExtractAll()
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
