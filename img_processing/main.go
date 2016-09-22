package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"log"
	"os"
	"path/filepath"
)

type pixel struct {
	r, q, b, a uint32
}

func main() {
	imgs := loadImgs("./img_processing/img")
	for i, img := range imgs {
		for j, pixel := range img {
			fmt.Println("Image", i, "\t pixel", j, "\t RGBA", pixel)
			if j == 2 {
				break
			}
		}
	}
}

func getPixels(img image.Image) []pixel {
	bounds := img.Bounds()
	// lol
	pixels := make([]pixel, bounds.Dx()*bounds.Dy())

	for i := 0; i < bounds.Dx()*bounds.Dy(); i++ {
		x := i % bounds.Dx()
		y := i / bounds.Dx()
		r, _, b, a := img.At(x, y).RGBA()
		pixels[i].r = r
		//pixels[i].g = g
		pixels[i].b = b
		pixels[i].a = a
	}

	return pixels
}

func loadImgs(dir string) [][]pixel {
	var images [][]pixel

	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		img := loadImage(path)
		pixels := getPixels(img)
		images = append(images, pixels)
		return nil
	})

	return images
}


func loadImage(filename string) image.Image {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	img, err := jpeg.Decode(f)
	if err != nil {
		log.Fatal(err)
	}

	return img
}
