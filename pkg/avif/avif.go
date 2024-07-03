package avif

import (
	"image"
	"io"
)

// Version of the avif-go package
const Version = "0.1.0"

// Decode reads an AVIF image from r and returns it as an image.Image.
func Decode(r io.Reader) (image.Image, error) {
	decoder := NewDecoder(r)
	return decoder.Decode()
}

// Encode writes the Image m to w in AVIF format.
func Encode(w io.Writer, m image.Image) error {
	encoder := NewEncoder(w)
	return encoder.Encode(m)
}