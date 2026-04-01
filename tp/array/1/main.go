package main

import "fmt"

type IPv4 = [4]byte

// Well-known IPv4 addresses
var (
	IPv4bcast     = IPv4{255, 255, 255, 255} // limited broadcast
	IPv4allsys    = IPv4{224, 0, 0, 1}       // all systems
	IPv4allrouter = IPv4{224, 0, 0, 2}       // all routers
	IPv4zero      = IPv4{0, 0, 0, 0}         // all zeros
)

func main() {
	// var asUint32 uint32 = 0xFFFFFFFF // 4_294_967_295
	var asUint32 uint32 = 0x000000 // 0
	ip := ToArray(asUint32)
	fmt.Println(ip)
}

func ToArray(v uint32) [4]byte {
	return [4]byte{byte(v >> 24), byte(v >> 16), byte(v >> 8), byte(v)}
}

func ToUint32(a [4]byte) uint32 {
	return (uint32(a[0]) << 24) | (uint32(a[1]) << 16) | (uint32(a[2]) << 8) | uint32(a[3])
}
