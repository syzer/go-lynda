package main

import (
	"os"
	"log"
	"image/jpeg"
	"fmt"
)

func main() {
	f, err := os.Open("./img_processing/img/2016HackZurich.jpg")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	img, err := jpeg.Decode(f)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%T ", img)

	r, g, b, a := img.At(0, 0).RGBA()
	fmt.Println("=>", r, g, b, a)

	bounds := img.Bounds()

	r, g, b, a = img.At(bounds.Dx(), bounds.Dy()).RGBA()
	fmt.Println("=>", r, g, b, a)
}
