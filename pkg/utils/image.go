package utils

import (
	"bytes"
	imgLib "image"
	"image/jpeg"
	"image/png"
	"io"

	"golang.org/x/image/webp"
)

type Image interface {
	Open(data io.Reader, ext string) error
	Resize(newWidth, newHeight int)
	ToByte() ([]byte, error)
}

type image struct {
	file imgLib.Image
	ext  string
}

func NewImage() Image {
	return &image{}
}

func (i *image) Open(data io.Reader, ext string) error {
	var File imgLib.Image
	if ext == ".webp" {
		imageFile, err := webp.Decode(data)
		if err != nil {
			return err
		}
		File = imageFile

	} else {
		imageFile, _, err := imgLib.Decode(data)
		if err != nil {
			return err
		}
		File = imageFile
	}
	i.ext = ext
	i.file = File

	return nil
}

func (i *image) ToByte() ([]byte, error) {
	buf, err := i.encode()
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (i *image) encode() (*bytes.Buffer, error) {
	buf := new(bytes.Buffer)
	switch {
	case i.ext == ".jpg" || i.ext == ".jpeg" || i.ext == ".webp":
		err := jpeg.Encode(buf, i.file, nil)
		if err != nil {
			return nil, err
		}

	case i.ext == ".png":
		err := png.Encode(buf, i.file)
		if err != nil {
			return nil, err
		}
	}
	return buf, nil
}

func (i *image) Resize(newWidth, newHeight int) {
	// 원본 이미지의 bounds 가져오기
	origBounds := i.file.Bounds()

	// 새로운 크기에 맞게 비율 조정
	xRatio := float64(origBounds.Dx()) / float64(newWidth)
	yRatio := float64(origBounds.Dy()) / float64(newHeight)
	var ratio float64
	if xRatio > yRatio {
		ratio = xRatio
	} else {
		ratio = yRatio
	}

	// 새로운 크기로 이미지 크기 조절
	newBounds := imgLib.Rect(0, 0, int(float64(origBounds.Dx())/ratio), int(float64(origBounds.Dy())/ratio))
	newImg := imgLib.NewRGBA(newBounds)
	for y := 0; y < newBounds.Dy(); y++ {
		for x := 0; x < newBounds.Dx(); x++ {
			oldX := int(float64(x) * ratio)
			oldY := int(float64(y) * ratio)
			newImg.Set(x, y, i.file.At(oldX, oldY))
		}
	}

	i.file = newImg
}
