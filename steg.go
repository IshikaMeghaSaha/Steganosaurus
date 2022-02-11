package main

import (
	"image"
	"image/jpeg"
	"log"
	"math/rand"
	"os"
)

func main() {
	img := image.NewRGBA(image.Rect(0, 0, 100, 200))
	for p := 0; p < 100*200; p++ {
		offset := 4 * p
		img.Pix[0+offset] = uint8(rand.Intn(256))
		img.Pix[1+offset] = uint8(rand.Intn(256))
		img.Pix[2+offset] = uint8(rand.Intn(256))
		img.Pix[3+offset] = 255

		ofile, err := os.Create("image.jpeg")
		if err != nil {
			log.Fatal(err)
		}

		jpeg.Encode(ofile, img, nil)
		err = ofile.Close()
		if err != nil {
			log.Fatal(err)
		}
	}
}
