package main

import (
	"flag"
	"reverse/asciiArtTemplates"
)

func main() {
	reverseBanner := flag.String("reverse", "file.txt", "Reverses teh file.")
	flag.Parse()

	asciiArtTemplates.ReadTemplates(&Store, "standard")

}

var (
	Store [128][8]string // Переменная для хранения символов из файла
)
