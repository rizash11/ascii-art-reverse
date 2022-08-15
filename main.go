package main

import (
	"bufio"
	"flag"
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
}

var (
	Store    [128][8]string // Переменная для хранения символов из файла
	AsciiArt [][8]string
)

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
