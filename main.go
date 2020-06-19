package main

import (
	"fmt"
	"github.com/strukturag/libheif/go/heif"
	"image"
	"image/draw"
	"image/png"
	"os"
	"path/filepath"
	"strconv"
)

func main() {

	widthStart, err := strconv.ParseInt(os.Args[1], 10, 0)
	checkErr(err)
	widthEnd, err := strconv.ParseInt(os.Args[2], 10, 0)
	checkErr(err)

	for width := int(widthStart); width <= int(widthEnd); width++ {
		filename := "sample-width" + strconv.Itoa(width) + ".png"
		file, err := os.Open(filename)
		checkErr(err)
		img, err := png.Decode(file)
		checkErr(err)
		err = convertPNGtoHEIF(img, width)
		checkErr(err)
	}

	os.Exit(0)
}

func convertPNGtoHEIF(img image.Image, width int) error {

	fmt.Println()
	fmt.Printf("#####  Width: %dpx\n", img.Bounds().Size().X)
	fmt.Println()

	b := img.Bounds()
	imgNRGBA := image.NewNRGBA(image.Rect(0, 0, b.Dx(), b.Dy()))
	draw.Draw(imgNRGBA, imgNRGBA.Bounds(), img, b.Min, draw.Src)

	ctx, err := heif.EncodeFromImage(imgNRGBA, heif.CompressionHEVC, 75, heif.LosslessModeDisabled, heif.LoggingLevelFull)
	if err != nil {
		return err
	}

	filename := "sample-width" + strconv.Itoa(width) + ".heif"

	dstPath := filepath.Join(".", filename)
	err = ctx.WriteToFile(dstPath)
	if err != nil {
		return err
	}
	return nil
}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
