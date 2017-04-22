package main

import (
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"time"
	"unicode"
)

func main() {
	imageName := "in.jpg"
	outputName := "out.jpg"
	if len(os.Args) >= 3 {
		imageName = os.Args[1]
		outputName = os.Args[2]
	} else if len(os.Args) == 2 {
		imageName = os.Args[1]
	}
	img, err := ioutil.ReadFile(imageName)
	if err != nil {
		log.Panicln("ERROR! -", err)
	}
	out := glitch(img)
	ioutil.WriteFile(outputName, out, os.FileMode(0644))
}

func glitch(imageBytes []byte) []byte {
	rand.Seed(time.Now().UTC().UnixNano())
	// arbitrarily assume this 5% will miss the signature & header
	percent := len(imageBytes) * 5 / 100
	// flip the unicode runes upper or lower based on a random number, scaled by image size
	randomint := rand.Intn(percent/2-200) + 200
	for i := range imageBytes {
		if i > percent {
			if i%randomint == 0 {
				startRune := rune(imageBytes[i])
				var endByte byte
				if unicode.IsUpper(startRune) {
					endByte = byte(unicode.ToLower(startRune))
				} else {
					endByte = byte(unicode.ToUpper(startRune))
				}
				imageBytes[i] = endByte
			}
		}
	}
	return imageBytes
}
