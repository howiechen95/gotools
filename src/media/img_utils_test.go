package media

import (
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
