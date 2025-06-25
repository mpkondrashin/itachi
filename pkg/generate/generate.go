/***************************************************************
*
* Itachi (c) 2022 by Michael Kondrashin (mkondrashin@gmail.com)
* Copyright under MIT Lincese. Please see LICENSE file for details
*
* generate.go - Itachi generator
*
***************************************************************/

package generate

import (
	"embed"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/google/uuid"
)

// go : embed malware/*.exe

const unique = "XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX"

//go:embed gmw/*.exe
var malware embed.FS

func ExtractAll() {
	dir, err := malware.ReadDir("gmw")
	if err != nil {
		panic(err)
	}
	uniqueStr := uuid.New().String()
	for i, each := range dir {
		if !each.Type().IsRegular() {
			continue
		}
		targetFileName := uniqueStr + "_" + each.Name()
		fmt.Printf("%d: %s\n", i+1, targetFileName)
		f, err := os.Create(targetFileName)
		if err != nil {
			panic(err)
		}
		if err := Extract(f, each.Name()); err != nil {
			f.Close()
			panic(err)
		}
		f.Close()
	}
}

func ExtractFile(w http.ResponseWriter, name string) error {
	uniqueStr := uuid.New().String()
	f, err := malware.Open("gmw/" + name)
	if err != nil {
		return err
	}
	defer f.Close()

	fileInfo, err := f.Stat()
	if err != nil {
		return err
	}
	targetFileName := uniqueStr + "_" + name
	w.Header().Set("Content-Disposition", "attachment; filename="+targetFileName)
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Content-Length", strconv.Itoa(int(fileInfo.Size())))
	return ReplaceString(f, w, unique, uniqueStr)
}

func Extract(w io.Writer, name string) error {
	uniqueStr := uuid.New().String()
	f, err := malware.Open("gmw/" + name)
	if err != nil {
		return err
	}
	defer f.Close()
	return ReplaceString(f, w, unique, uniqueStr)
}

func ExtractEicar(w http.ResponseWriter) error {
	//The following is reverce of eicar.com (X5O!P%@AP[4\PZX54(P^)7CC)7}$EICAR-STANDARD-ANTIVIRUS-TEST-FILE!$H+H*)
	dataReverce := `*H+H$!ELIF-TSET-SURIVITNA-DRADNATS-RACIE$}7)CC7)^P(45XZP\4[PA@%P!O5X`
	w.Header().Set("Content-Disposition", "attachment; filename=eicar.com")
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Content-Length", strconv.Itoa(len(dataReverce)))
	for i := len(dataReverce) - 1; i >= 0; i-- {
		_, err := w.Write([]byte{dataReverce[i]})
		if err != nil {
			return err
		}
	}
	return nil
}

// ReplaceString reads from reader r, replaces all occurrences of old with new,
// and writes the result to writer w. It processes data in chunks to handle large streams efficiently.
func ReplaceString(r io.Reader, w io.Writer, old, new string) error {
	buf := make([]byte, 32*1024) // 32KB chunks
	var leftover []byte

	for {
		n, err := r.Read(buf)
		if n > 0 {
			// Combine leftover with current chunk
			chunk := append(leftover, buf[:n]...)

			// Find a safe split point that doesn't break potential matches
			splitPoint := n
			if len(chunk) > len(old) {
				splitPoint = len(chunk) - len(old) + 1
			}

			// Process main part
			result := strings.Replace(string(chunk[:splitPoint]), old, new, -1)
			_, err := w.Write([]byte(result))
			if err != nil {
				return err
			}

			// Save remainder for next iteration
			leftover = chunk[splitPoint:]
		}

		if err == io.EOF {
			// Process any remaining data
			if len(leftover) > 0 {
				result := strings.Replace(string(leftover), old, new, -1)
				_, err := w.Write([]byte(result))
				if err != nil {
					return err
				}
			}
			return nil
		}
		if err != nil {
			return err
		}
	}
}
