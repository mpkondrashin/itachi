/***************************************************************
*
* Itachi (c) 2022 by Michael Kondrashin (mkondrashin@gmail.com)
* Copyright under MIT Lincese. Please see LICENSE file for details
*
* main.go - downloader - download and save eicar.com file
*
***************************************************************/

package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

var Unique = "XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX"

func downloadFile(filepath string, url string) (err error) {
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if _, err = io.Copy(out, resp.Body); err != nil {
		return err
	}
	return
}

func main() {
	fmt.Printf("Demo Downloader (%s)\n", Unique)
	if err := downloadFile("eicar.com", "https://secure.eicar.org/eicar.com"); err != nil {
		panic(err)
	}
	fmt.Println("Done")
}
