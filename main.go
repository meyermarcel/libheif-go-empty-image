package main

import (
	"fmt"
	"github.com/disintegration/imaging"
	"github.com/strukturag/libheif/go/heif"
	"image"
	"image/png"
	"os"
	"path/filepath"
	"strconv"
)

func main() {

	file, err := os.Open("sample.png")
	checkErr(err)

	img, err := png.Decode(file)
	checkErr(err)

	widthStart, err := strconv.ParseInt(os.Args[1], 10, 0)
	checkErr(err)
	widthEnd, err := strconv.ParseInt(os.Args[2], 10, 0)
	checkErr(err)

	for width := int(widthStart); width <= int(widthEnd); width++ {
		err := convertPNGtoHEIF(img, width)
		checkErr(err)
	}

	os.Exit(0)
}

func convertPNGtoHEIF(img image.Image, width int) error {

	resizedImg := imaging.Resize(img, width, 360, imaging.Lanczos)
	fmt.Println()
	fmt.Printf("#####  Width: %dpx\n", resizedImg.Bounds().Size().X)
	fmt.Println()

	ctx, err := heif.EncodeFromImage(resizedImg, heif.CompressionHEVC, 75, heif.LosslessModeDisabled, heif.LoggingLevelFull)
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
