package library

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
)

func ImageProcessor(data string) (*bytes.Buffer, error) {
	// Decode base64 string ke byte array
	imageBytes, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return nil, err
	}

	// decode gambar
	img, format, err := image.Decode(bytes.NewReader(imageBytes))
	if err != nil {
		return nil, err
	}

	// Resize image
	newHeight := img.Bounds().Dy() * 500 / img.Bounds().Dx()
	resized := image.NewRGBA(image.Rect(0, 0, 500, newHeight))
	if err := resize(resized, img); err != nil {
		return nil, err
	}

	// create new varibale bytes to alocate new image
	buf := new(bytes.Buffer)

	// do compress image
	switch format {
	case "jpeg":
		// compress jpeg
		opts := jpeg.Options{
			Quality: 50,
		}
		err = jpeg.Encode(buf, resized, &opts)
		if err != nil {
			return nil, err
		}
	case "png":
		// compress png
		err = png.Encode(buf, resized)
		if err != nil {
			return nil, err
		}
	default:
		// format not valid
		return nil, fmt.Errorf("Format gambar tidak didukung")
	}
	fmt.Println("Gambar berhasil dikompresi")

	return buf, nil
}

// resize resizes src to dst.
func resize(dst *image.RGBA, src image.Image) error {
	sw, sh := src.Bounds().Dx(), src.Bounds().Dy()
	dw, dh := dst.Bounds().Dx(), dst.Bounds().Dy()

	for dy := 0; dy < dh; dy++ {
		for dx := 0; dx < dw; dx++ {
			sx := dx * sw / dw
			sy := dy * sh / dh
			dst.Set(dx, dy, src.At(sx, sy))
		}
	}

	return nil
}
