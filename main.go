package main

import (
	"fmt"
	"log"
	"os"
	"image"
	//"image/color"
	_ "image/jpeg"
)

func main() {
	imgFileName := "test.jpg"
	reader, err := os.Open(imgFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer reader.Close()

	m, _, err := image.Decode(reader)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(m.Bounds())
	fmt.Println(m.At(50, 70))
}