package main

import (
	"bytes"
	"fmt"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io/ioutil"
	"log"
	"os"
	"strings"
	municode "unicode"
	"unicode/utf16"
)

func main() {
	var con []byte

	if len(os.Args) >= 2 {
		file, err := os.Open(os.Args[1])
		if err != nil {
			log.Fatal(err)
		}
		con, err = ioutil.ReadAll(file)
		if err != nil {
			log.Fatal(err)
		}

		if len(os.Args) == 4 {
			con = repUTF16(con, os.Args[2], os.Args[3])
			ioutil.WriteFile(os.Args[1], con, 0666)
		} else {
			printUTF16(con)
		}
	} else {
		fmt.Println(
			"A UTF-16 FILE str print Or str replace Tool.\n" +
				"Usage: " + os.Args[0] + " filename [oldstr] [newstr]")
	}
}

func printUTF16(bs []byte) {
	bsUtf8le, _, _ := transform.Bytes(unicode.UTF16(unicode.LittleEndian, unicode.IgnoreBOM).NewDecoder(), bs)
	var buf []byte
	for _, b := range bsUtf8le {
		if b >= 0x80 {
			continue
		}
		if municode.IsPrint(rune(b)) {
			buf = append(buf, b)
		}
		if b == byte('\n') && len(buf) >= 4 {
			fmt.Println(strings.TrimSpace(string(buf)))
			buf = []byte{}
		}
	}
	fmt.Println(strings.TrimSpace(string(buf)))
}

func repUTF16(bs []byte, sou, des string) []byte {
	souUtf16 := utf16.Encode([]rune(sou))
	desUtf16 := utf16.Encode([]rune(des))
	return bytes.Replace(bs, listUint16ToBytes(souUtf16), listUint16ToBytes(desUtf16), -1)
}

func uint16ToBytes(n uint16) []byte {
	return []byte{
		byte(n),
		byte(n >> 8),
	}
}

func listUint16ToBytes(n []uint16) []byte {
	var t bytes.Buffer
	for _, v := range n {
		t.Write(uint16ToBytes(v))
	}
	return t.Bytes()
}
