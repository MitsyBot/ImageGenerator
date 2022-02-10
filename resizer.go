package main

import (
	"image"
	_ "image/jpeg"
	"image/png"
	"os"
)

import (
	"github.com/disintegration/imaging"
)

func main() {
	path := "./assets/images/badges/bug_hunter.png"
	x := 30
	y := 30
	save_to := "./assets/images/badges/bug_hunter.png"

	img_reader, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer img_reader.Close()

	img, _, err := image.Decode(img_reader)
	if err != nil {
		panic(err)
	}
	print("AVISO: Usando Lanczos para resamplear a imagem. Isto pode demorar um pouco.\n")
	img = imaging.Fill(img, x, y, imaging.Center, imaging.Lanczos)

	outf, err := os.Create(save_to)
	if err != nil {
		panic(err)
	}
	defer outf.Close()

	err = png.Encode(outf, img)
	if err != nil {
		panic(err)
	} else {
		print("pronto\n")
	}
}
