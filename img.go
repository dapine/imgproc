package main

import (
    "log"
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
