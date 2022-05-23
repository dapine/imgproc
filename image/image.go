package image

import (
	"log"

	"github.com/dapine/imgproc/math"
	"github.com/h2non/bimg"
)

type ImageType bimg.ImageType

var stringToImageType = map[string]ImageType {
	"jpeg": 1,
	"jpg": 1,
	"webp": 2,
	"png": 3,
	"tiff": 4,
	"gif": 5,
	"pdf": 6,
	"svg": 7,
	"magick": 8,
	"miff": 8,
	"heif": 9,
	"avif": 10,
}

var imageTypeToExtension = map[ImageType]string {
	1: "jpeg",
	2: "webp",
	3: "png",
	4: "tiff",
	5: "gif",
	6: "pdf",
	7: "svg",
	8: "miff",
	9: "heif",
	10: "avif",
}

var imageTypeToMime = map[ImageType]string {
	1: "image/jpeg",
	2: "image/webp",
	3: "image/png",
	4: "image/tiff",
	5: "image/gif",
	6: "application/pdf",
	7: "image/svg+xml",
	8: "image/x-miff",
	9: "image/heif",
	10: "image/avif",
}

var gravityMap = map[string]bimg.Gravity {
	"centre": 0,
	"north": 1,
	"east": 2,
	"south": 3,
	"west": 4,
	"smart": 5,
}

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

// Convert takes the image bytes and target format as extension or MIME type
func Convert(bytes []byte, imgType string) ([]byte, string, string) {
	it := NewImageType(imgType)
	img, err := bimg.NewImage(bytes).Convert(bimg.ImageType(it))
	herr(err)

	ex := it.ToExtension()
	mime := it.ToMime()

	return img, ex, mime
}

func Crop(bytes []byte, width, height int64, gravity string) []byte {
	gr := newGravity(gravity)

	img, err := bimg.NewImage(bytes).Crop(int(width), int(height), bimg.Gravity(gr))
	herr(err)

	return img
}

func (imgType ImageType) ToExtension() string {
	ex := imageTypeToExtension[imgType]

	return ex
}

func (imgType ImageType) ToMime() string {
	mime := imageTypeToMime[imgType]

	return mime
}

func NewImageType(itype string) ImageType {
	it := stringToImageType[itype]

	return it
}

func newGravity(gravity string) bimg.Gravity {
	gr := gravityMap[gravity]

	return gr
}
