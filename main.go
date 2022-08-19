package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"reverse/asciiArtTemplates"
)

func main() {
	reverseBanner := flag.String("reverse", "file.txt", "Reverses teh file.")
	flag.Parse()
	asciiArtTemplates.ReadTemplates(&Store, "standard")

	f, err := os.Open(*reverseBanner)
	asciiArtTemplates.Check("Error opening the file to be reversed:", err)

	ReadAscii(f)

	for _, asciiString := range AsciiArt {
		fmt.Println(Decypher(asciiString))
	}
}

var (
	Store    [128][8]string // Переменная для хранения символов из файла
	AsciiArt [][8]string
)

func Decypher(asciiString [8]string) (regularString string) {

	index := 0

outer:
	for {

		for asI, asChar := range Store {
			found := true

		inner:
			for j := 0; j < 8; j++ {
				switch {
				case asChar[0] == "":
					found = false
					break inner
				case index+len(asChar[j]) > len(asciiString[j]):
					found = false
					break inner
				case asChar[j] != asciiString[j][index:index+len(asChar[j])]:
					found = false
					break inner
				}
			}

			if found {
				regularString = regularString + string(rune(asI))
				index = index + len(asChar[0])
				break
			} else if asI == 127 && !found {
				break outer
			}
		}

	}
	return regularString
}

func ReadAscii(f *os.File) {
	var tmp [8]string

	scanner := bufio.NewScanner(f)
	i := 0

	for scanner.Scan() {
		tmp[i] = scanner.Text()
		i++

		if i == 8 {
			AsciiArt = append(AsciiArt, tmp)
			i = 0
		}
	}
}
