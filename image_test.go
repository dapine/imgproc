package main

import (
	"bytes"
	"os"
	"testing"

	"github.com/dapine/imgproc/image"
)

const filename = "./gopher.png"

func getImg(t *testing.T) []byte {
	buf := new(bytes.Buffer)
	file, err := os.Open(filename)
	if err != nil {
		t.Error(err)
	}
	defer file.Close()

	buf.ReadFrom(file)

	return buf.Bytes()
}

func write(bytes []byte, filename string, t *testing.T) {
	err := os.WriteFile("/tmp/" + filename, bytes, 0644)
	if err != nil {
		t.Error(err)
	}
}

func TestResize(t *testing.T) {
	img := getImg(t)

	newImg := image.Resize(img, 50, 50)
	write(newImg, "resize", t)
}

func TestRotate(t *testing.T) {
	img := getImg(t)

	r0 := image.Rotate(img, 0)
	write(r0, "rotate-0", t)
	r90 := image.Rotate(img, 90)
	write(r90, "rotate-90", t)
	r180 := image.Rotate(img, 180)
	write(r180, "rotate-180", t)
	r270 := image.Rotate(img, 270)
	write(r270, "rotate-270", t)
}
