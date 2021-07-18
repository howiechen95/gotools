package media

import (
	"testing"
)

func TestResizeImage(t *testing.T) {
	fileUri := "C:\\Users\\Administrator\\Pictures\\86446038_p0.jpg"
	//fileUri := "C:\\Users\\Administrator\\Pictures\\Snipaste_2021-06-12_08-40-44.png"
	width, height := uint(400), uint(400)
	ResizeImage(fileUri, width, height)
}
