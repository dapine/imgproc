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
	err := os.WriteFile("/tmp/imgproc-" + filename, bytes, 0644)
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

func TestConvertToJpg(t *testing.T) {
	img := getImg(t)
	jpeg, ex, mime := image.Convert(img, "jpeg")
	if (ex != "jpeg") {
		t.Errorf("expected: %s, got: %s", "jpeg", ex)
	}
	if (mime != "image/jpeg") {
		t.Errorf("expected: %s, got: %s", "image/jpeg", mime)
	}
	write(jpeg, "convert-jpeg", t)
}

func TestConvertToPng(t *testing.T) {
	img := getImg(t)
	png, ex, mime := image.Convert(img, "png")
	if (ex != "png") {
		t.Errorf("expected: %s, got: %s", "png", ex)
	}
	if (mime != "image/png") {
		t.Errorf("expected: %s, got: %s", "image/png", mime)
	}
	write(png, "convert-png", t)
}

func TestConvertToGif(t *testing.T) {
	img := getImg(t)
	gif, ex, mime := image.Convert(img, "gif")
	if (ex != "gif") {
		t.Errorf("expected: %s, got: %s", "gif", ex)
	}
	if (mime != "image/gif") {
		t.Errorf("expected: %s, got: %s", "image/gif", mime)
	}
	write(gif, "convert-gif.gif", t)
}

func TestCropCentre(t *testing.T) {
	img := getImg(t)

	newImg := image.Crop(img, 150, 150, "centre")
	write(newImg, "crop-centre", t)
}

func TestCropNorth(t *testing.T) {
	img := getImg(t)

	newImg := image.Crop(img, 150, 150, "north")
	write(newImg, "crop-north", t)
}

func TestCropEast(t *testing.T) {
	img := getImg(t)

	newImg := image.Crop(img, 150, 150, "east")
	write(newImg, "crop-east", t)
}

func TestCropSouth(t *testing.T) {
	img := getImg(t)

	newImg := image.Crop(img, 150, 150, "south")
	write(newImg, "crop-south", t)
}

func TestCropWest(t *testing.T) {
	img := getImg(t)

	newImg := image.Crop(img, 150, 150, "west")
	write(newImg, "crop-west", t)
}

func TestCropSmart(t *testing.T) {
	img := getImg(t)

	newImg := image.Crop(img, 150, 150, "smart")
	write(newImg, "crop-smart", t)
}

func TestEnlarge(t *testing.T) {
	img := getImg(t)

	newImg := image.Enlarge(img, 3000, 0)
	write(newImg, "enlarge", t)
}

func TestExtract(t *testing.T) {
	img := getImg(t)

	newImg := image.Extract(img, 160, 600, 128, 128)
	write(newImg, "extract", t)
}

func TestFlipVertical(t *testing.T) {
	img := getImg(t)

	newImg := image.Flip(img, "vertical")
	write(newImg, "flip-vertical", t)
}

func TestFlipY(t *testing.T) {
	img := getImg(t)

	newImg := image.Flip(img, "y")
	write(newImg, "flip-y", t)
}

func TestFlipHorizontal(t *testing.T) {
	img := getImg(t)

	newImg := image.Flip(img, "horizontal")
	write(newImg, "flip-horizontal", t)
}

func TestFlipX(t *testing.T) {
	img := getImg(t)

	newImg := image.Flip(img, "x")
	write(newImg, "flip-x", t)
}
