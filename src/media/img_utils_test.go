package media

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

func TestResizePng(t *testing.T) {
	fileUri := "./image/test01.png"
	width, height := uint(400), uint(400)
	ResizePng(fileUri, "./image/aa.png", width, height)
}

func TestResizeJpg(t *testing.T) {
	fileUri := "./image/test01.jpg"
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

	file1, err := os.Open("./image/test01.jpg")
	if err != nil {
		fmt.Println(err)
	}
	defer file1.Close()
	img, _ := jpeg.Decode(file1)

	jpg := image.NewRGBA(image.Rect(0, 0, 100, 100))
	draw.Draw(jpg, img.Bounds().Add(image.Pt(10, 10)), img, img.Bounds().Min, draw.Src) //截取图片的一部分
	jpeg.Encode(file, jpg, nil)
}

func TestDraw001(t *testing.T) {
	file, err := os.Create("./image/dst.jpg")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	jpg := image.NewRGBA(image.Rect(0, 0, 1, 1))
	blue := color.RGBA{0, 0, 255, 255}
	//draw.Draw(jpg, img.Bounds().Add(image.Pt(10, 10)), img, img.Bounds().Min, draw.Src) //截取图片的一部分
	draw.Draw(jpg, jpg.Bounds(), &image.Uniform{blue}, image.ZP, draw.Src)
	jpeg.Encode(file, jpg, nil)

}

func TestChangeRGB(t *testing.T) {
	source := "./image/test01.png" //输入图片
	target := "./image/result.png" //输出图片

	ff, _ := ioutil.ReadFile(source) //读取文件
	bbb := bytes.NewBuffer(ff)
	m, _, _ := image.Decode(bbb)
	bounds := m.Bounds()
	dx := bounds.Dx()
	dy := bounds.Dy()

	newRgba := image.NewRGBA(bounds) //new 一个新的图片
	for i := 0; i < dx; i++ {
		for j := 0; j < dy; j++ {
			colorRgb := m.At(i, j)
			r, g, b, a := colorRgb.RGBA()
			r_uint8 := uint8(r >> 8) //转换为 255 值
			g_uint8 := uint8(g >> 8)
			b_uint8 := uint8(b >> 8)
			a_uint8 := uint8(a >> 8)

			r_uint8 = 255 - r_uint8
			g_uint8 = 255 - g_uint8
			b_uint8 = 255 - b_uint8
			newRgba.SetRGBA(i, j, color.RGBA{r_uint8, g_uint8, b_uint8, a_uint8}) //设置像素点

		}
	}

	f, _ := os.Create(target)
	defer f.Close()
	encode(source, f, newRgba)
}

//图片编码 写入
func encode(inputName string, file *os.File, rgba *image.RGBA) {
	if strings.HasSuffix(inputName, "jpg") || strings.HasSuffix(inputName, "jpeg") {
		jpeg.Encode(file, rgba, nil)
	} else if strings.HasSuffix(inputName, "png") {
		png.Encode(file, rgba)
	} else if strings.HasSuffix(inputName, "gif") {
		gif.Encode(file, rgba, nil)
	} else {
		fmt.Errorf("不支持的图片格式")
	}
}

func TestSetRGB(t *testing.T) {
	//source := "./image/test01.png" 		//输入图片
	target := "./image/result.png" //输出图片
	x := 400
	y := 400

	newRgba := image.NewRGBA(image.Rect(0, 0, x, y)) //new 一个新的图片
	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {
			//colorRgb := m.At(i, j)
			//r, g, b, a := colorRgb.RGBA()
			//r_uint8 := uint8(r >> 8)	//转换为 255 值
			//g_uint8 := uint8(g >> 8)
			//b_uint8 := uint8(b >> 8)
			//a_uint8 := uint8(a >> 8)

			//r_uint8 = 255 - r_uint8
			//g_uint8 = 255 - g_uint8
			//b_uint8 = 255 - b_uint8
			p1 := color.RGBA{
				uint8(i % 255),
				uint8(j % 255),
				uint8((i + j) % 255),
				uint8((i + j) % 255),
			}
			//newRgba.SetRGBA(i, j, color.RGBA{r_uint8, g_uint8, b_uint8, a_uint8}) //设置像素点
			newRgba.SetRGBA(i, j, p1) //设置像素点

		}
	}

	f, _ := os.Create(target)
	defer f.Close()
	png.Encode(f, newRgba)
	//encode(source, f, newRgba)
}
