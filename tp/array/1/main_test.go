package main

import (
	"math/rand"
	"reflect"
	"testing"
	"time"
)

// Génère des jeux d'entrée réutilisables (évite de mesurer le coût du rand dans le hot loop).
func makeInputs32(n int) []uint32 {
	r := rand.New(rand.NewSource(42))
	in := make([]uint32, n)
	for i := range in {
		in[i] = r.Uint32()
	}
	return in
}

func makeInputs4(n int) []IPv4 {
	r := rand.New(rand.NewSource(1337))
	in := make([]IPv4, n)
	for i := range in {
		v := r.Uint32()
		in[i] = IPv4{byte(v >> 24), byte(v >> 16), byte(v >> 8), byte(v)}
	}
	return in
}

func BenchmarkToArray(b *testing.B) {
	const N = 1 << 16
	inputs := makeInputs32(N)

	b.ReportAllocs()
	b.SetBytes(4) // on "traite" 4 octets par op (résultat)
	for i := 0; i < b.N; i++ {
		_ = ToArray(inputs[i&(N-1)])
	}
}

func BenchmarkToUint32(b *testing.B) {
	const N = 1 << 16
	inputs := makeInputs4(N)

	b.ReportAllocs()
	b.SetBytes(4) // on "traite" 4 octets d'entrée par op
	for i := 0; i < b.N; i++ {
		_ = ToUint32(inputs[i&(N-1)])
	}
}

// --- Tests de correction (rapides) ---

func TestRoundTrip(t *testing.T) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < 10000; i++ {
		v := r.Uint32()
		a := ToArray(v)
		got := ToUint32(a)
		if got != v {
			t.Fatalf("roundtrip failed: v=%d, got=%d, a=%v", v, got, a)
		}
	}
}

func TestKnownVectors(t *testing.T) {
	cases := []struct {
		v   uint32
		arr IPv4
	}{
		{0, IPv4{0, 0, 0, 0}},
		{1, IPv4{0, 0, 0, 1}},
		{0xFF_FF_FF_FF, IPv4{255, 255, 255, 255}},
		{0xC0_A8_00_01, IPv4{192, 168, 0, 1}}, // 192.168.0.1
	}
	for _, c := range cases {

		if got := ToArray(c.v); !reflect.DeepEqual(got, c.arr) {
			t.Fatalf("ToArray(%08x) = %v, want %v", c.v, got, c.arr)
		}

		if got := ToUint32(c.arr); got != c.v {
			t.Fatalf("ToUint32(%v) = %08x, want %08x", c.arr, got, c.v)
		}
	}
}
