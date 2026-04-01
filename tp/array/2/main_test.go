package main

import (
	"io"
	"testing"
)

const testFile = "./data/customers.txt"

func BenchmarkReadFileBuffered(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		if err := ReadFileBuffered(testFile, io.Discard); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkReadFileWithoutBuffer(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		if err := ReadFileWithoutBuffer(testFile, io.Discard); err != nil {
			b.Fatal(err)
		}
	}
}
