package main

import (
	"fmt"
	"io"
	"os"
)

var (
	pngSig  = [8]byte{0x89, 'P', 'N', 'G', 0x0D, 0x0A, 0x1A, 0x0A}
	zipSig4 = [4]byte{'P', 'K', 0x03, 0x04}
	jpgSig3 = [3]byte{0xFF, 0xD8, 0xFF}
	pdfSig  = [5]byte{0x25, 'P', 'D', 'F', 0x2D}
)

func GetMimetype(path string) (string, error) {
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer f.Close()

	return mime(f)
}

func mime(r io.Reader) (string, error) {
	var buf [16]byte
	n, err := io.ReadFull(r, buf[:])
	if err != nil && err != io.EOF && err != io.ErrUnexpectedEOF {
		return "", err
	}

	switch {
	case n >= 8 && [8]byte(buf[:8]) == pngSig:
		return "image/png", nil
		// Ajouter ici le test pour voir si c'est un pdf ou une image jpg
	}
	return "application/octet-stream", nil
}

func main() {
	mt, err := GetMimetype("data/gopher.jpg")
	if err != nil {
		panic(err)
	}
	fmt.Println(mt) // ex: application/pdf
}
