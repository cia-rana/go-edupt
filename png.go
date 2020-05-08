package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
)

func SavePNG(filename string, img []color.Color, width, height int) {
	if width < 0 || height < 0 || width*height < len(img) {
		return
	}

	dstImg := image.NewNRGBA(image.Rect(0, 0, width, height))

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			dstImg.Set(x, y, img[width*y+x])
		}
	}

	dstFile, err := os.Create(filename)
	defer func() {
		if err := dstFile.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	if err != nil {
		fmt.Println(err)
		return
	}

	encoder := png.Encoder{CompressionLevel: png.BestSpeed}
	if err := encoder.Encode(dstFile, dstImg); err != nil {
		fmt.Println(err)
		return
	}
}
