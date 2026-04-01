package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// pathToFile := "./data/customers.txt"
	// ReadFileBuffered(pathToFile, os.Stdout)
	// ReadFileWithoutBuffer(pathToFile, os.Stdout)
}

func ReadFileBuffered(filepath string, w io.Writer) error {
	f, err := os.Open(filepath)
	if err != nil {
		return err
	}
	defer f.Close()

	var buf [16 << 10]byte // tampon fixe sur la pile 16kbits

	for {
		n, err := f.Read(buf[:]) // on lit dans le slice basé sur l'array
		if n > 0 {
			// traiter les n octets lus
			if _, werr := w.Write(buf[:n]); werr != nil {
				return werr
			}
		}
		if err == io.EOF {
			break // fin de fichier
		}
		if err != nil {
			return err // autre erreur I/O
		}
	}
	return nil
}

func ReadFileWithoutBuffer(filepath string, w io.Writer) error {
	data, err := os.ReadFile(filepath)
	if err != nil {
		return fmt.Errorf("ReadFileWithoutBuffer error: %v", err.Error())
	}

	_, err = w.Write(data)
	if err != nil {
		return fmt.Errorf("ReadFileWithoutBuffer error: %v", err.Error())
	}
	return nil
}
