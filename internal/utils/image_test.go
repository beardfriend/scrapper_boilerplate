package utils

import (
	"net/http"
	urlM "net/url"
	"os"
	"path/filepath"
	"testing"
)

func TestImageWebpToJpegt(t *testing.T) {
	url := `https://www.costco.co.kr/medias/sys_master/h41/h63/80487999438878.webp`
	response, err := http.Get(url)
	if err != nil {
		t.Error(err)
	}
	defer response.Body.Close()

	parsedURL, err := urlM.Parse(url)
	if err != nil {
		t.Error(err)
	}

	ext := filepath.Ext(parsedURL.Path)

	image := NewImage()
	if err := image.Open(response.Body, ext); err != nil {
		t.Error(err)
	}

	image.Resize(400, 400)

	imageData, _ := image.ToByte()

	// Write bytes to image file

	err = os.WriteFile("output.jpg", imageData, 0o644)
	if err != nil {
		t.Fatal(err)
	}
}

func TestImageResize(t *testing.T) {
	url := `https://www.costco.co.kr/medias/sys_master/images/hb2/h13/153094183649310.jpg`
	response, err := http.Get(url)
	if err != nil {
		t.Error(err)
	}
	defer response.Body.Close()

	parsedURL, err := urlM.Parse(url)
	if err != nil {
		t.Error(err)
	}

	ext := filepath.Ext(parsedURL.Path)

	image := NewImage()
	if err := image.Open(response.Body, ext); err != nil {
		t.Error(err)
	}

	image.Resize(400, 400)

	imageData, _ := image.ToByte()

	// Write bytes to image file
	err = os.WriteFile("output.jpg", imageData, 0o644)
	if err != nil {
		t.Fatal(err)
	}
}
