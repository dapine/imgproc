package image

import (
	"log"

	"github.com/dapine/imgproc/math"
	"github.com/h2non/bimg"
)

func herr(err error) {
    if err != nil {
        log.Println(err)
    }
}

func Resize(bytes []byte, width int64, height int64) []byte {
	img, err := bimg.NewImage(bytes).Resize(int(width), int(height))
	herr(err)

	return img
}

func Rotate(bytes []byte, angle int64) []byte {
	a := math.NearestAngle(angle)
	img, err := bimg.NewImage(bytes).Rotate(a)
	herr(err)

	return img
}
