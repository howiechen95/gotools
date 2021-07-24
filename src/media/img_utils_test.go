package media

import (
	"fmt"
	"image"
	"image/draw"
	"image/jpeg"
	"os"
	"testing"
)

func TestResizePng(t *testing.T) {
	fileUri := "./test01.png"
	width, height := uint(400), uint(400)
	ResizePng(fileUri, "", width, height)
}

func TestResizeJpg(t *testing.T) {
	fileUri := "./test02.jpg"
	width, height := uint(400), uint(400)
	ResizeJpg(fileUri, "", width, height)
}

// 截图
func TestImg001(t *testing.T) {
	file, err := os.Create("dst.jpg")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	file1, err := os.Open("test02.jpg")
	if err != nil {
		fmt.Println(err)
	}
	defer file1.Close()
	img, _ := jpeg.Decode(file1)

	jpg := image.NewRGBA(image.Rect(0, 0, 100, 100))
	draw.Draw(jpg, img.Bounds().Add(image.Pt(10, 10)), img, img.Bounds().Min, draw.Src) //截取图片的一部分
	jpeg.Encode(file, jpg, nil)
}
