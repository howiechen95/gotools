package media

import (
	"fmt"
	"github.com/nfnt/resize"
	"image"
	"image/jpeg"
	"image/png"
	"log"
	"os"
	"strings"
	"time"
)

func ResizeJpg(srcFile, outFile string, width, height uint) {
	// open "test.jpg"
	file, err := os.Open(srcFile)
	if err != nil {
		log.Fatal(err)
	}

	// decode jpeg into image.Image
	img, err := jpeg.Decode(file)
	if err != nil {
		log.Fatal(err)
	}
	file.Close()

	// resize to width 1000 using Lanczos resampling
	// and preserve aspect ratio
	m := resize.Resize(width, height, img, resize.Lanczos3)

	outFile = getFileName(file.Name(), width, height)
	out, err := os.Create(outFile)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	// write new image to file
	jpeg.Encode(out, m, nil)
}

func ResizePng(srcFile, outFile string, width, height uint) {
	// open "test.jpg"
	file, err := os.Open(srcFile)
	if err != nil {
		log.Fatal(err)
	}

	// decode jpeg into image.Image
	img, err := png.Decode(file)
	if err != nil {
		log.Fatal(err)
	}
	file.Close()

	// resize to width 1000 using Lanczos resampling
	// and preserve aspect ratio
	m := resize.Resize(width, height, img, resize.Lanczos3)
	outFile = getFileName(file.Name(), width, height)
	out, err := os.Create(outFile)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	// write new image to file
	png.Encode(out, m)
}

func getFileName(srcName string, width, height uint) string {
	if srcName == "" {
		return ""
	}
	pathArr := strings.Split(srcName, "/")
	fileName := pathArr[0]
	if len(pathArr) > 1 {
		fileName = pathArr[len(pathArr)-1]
	}
	arr := strings.Split(fileName, ".")
	if len(arr) > 1 {
		return fmt.Sprintf("%s_%d_%dX%d.%s", arr[0], time.Now().Unix(), width, height, arr[len(arr)-1])
	} else {
		return fmt.Sprintf("%s_%d_%dX%d", arr[0], time.Now().Unix(), width, height)
	}
}

func ResizeImage1(srcFile, outFile string, width, height uint) {
	// open "test.jpg"
	file, err := os.Open(srcFile)
	if err != nil {
		log.Fatal(err)
	}

	img, imgType, err := image.Decode(file)
	if err != nil {
		log.Fatal(err)
		return
	}
	file.Close()

	// resize to width 1000 using Lanczos resampling
	// and preserve aspect ratio
	m := resize.Resize(width, height, img, resize.Lanczos3)

	if outFile == "" {
		outFile = fmt.Sprintf("%s_%d.jpg", strings.Split(file.Name(), ".")[0], time.Now().Unix())
	}
	out, err := os.Create(outFile)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	// write new image to file
	jpeg.Encode(out, m, nil)
	if imgType == "jpg" || imgType == "jpeg" {
		jpeg.Encode(out, img, nil)
	} else if imgType == "png" {
		png.Encode(out, img)
	} else {
		log.Fatal("not match type")
	}
}
