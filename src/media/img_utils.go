package media

import (
	"fmt"
	"github.com/nfnt/resize"
	"image"
	"image/jpeg"
	"image/png"
	"log"
	"os"
	"time"
)

func ResizeJpg(fileUri string, width, height uint) {
	// open "test.jpg"
	file, err := os.Open(fileUri)
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

	out, err := os.Create(fmt.Sprintf("test_resized_%d.jpg", time.Now().Unix()))
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	// write new image to file
	jpeg.Encode(out, m, nil)
}

func ResizePng(fileUri string, width, height uint) {
	// open "test.jpg"
	file, err := os.Open(fileUri)
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

	out, err := os.Create(fmt.Sprintf("test_resized_%d.png", time.Now().Unix()))
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	// write new image to file
	png.Encode(out, m)
}

func ResizeImage1(fileUri string, width, height uint) {
	// open "test.jpg"
	file, err := os.Open(fileUri)
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

	out, err := os.Create(fmt.Sprintf("test_resized_%d.jpg", time.Now().Unix()))
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
