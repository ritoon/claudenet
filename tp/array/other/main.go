package main

import (
	"fmt"
	"runtime"
	"unsafe"
)

const nbUsers = 2 << 2

type User struct {
	ID        int
	FirstName string
	LastName  string
}

func main() {
	fmt.Println("nbUsers:", nbUsers)

	initMemoryStart()
	var tbl *[nbUsers]User
	displayInfos("before init array", tbl)
	tbl = new([nbUsers]User)
	displayInfos("after init array", tbl)
	SetData(tbl, 2, "Rob", "Pike")
	displayInfos("after SetData", tbl)
	DeleteData(tbl, 2)
	displayInfos("after DeleteData", tbl)
	SetData(tbl, 2, "Rob", "Pike")
}

func SetData(tbl *[nbUsers]User, id int, fn, ln string) {
	tbl[id] = User{
		ID:        id,
		FirstName: fn,
		LastName:  ln,
	}
}

func DeleteData(tbl *[nbUsers]User, id int) {
	tbl[id] = User{}
}

func displayInfos(msg string, tbl *[nbUsers]User) {
	if tbl != nil {
		displayInfosArray(msg, tbl)
	}
	runtime.GC()
	var stats runtime.MemStats
	runtime.ReadMemStats(&stats)
	fmt.Printf("stats Alloc %v bytes\n", stats.Alloc-ms.StatsAlloc)
	fmt.Printf("stats HeapAlloc %v bytes\n", stats.HeapAlloc-ms.StatsHeapAlloc)
	fmt.Printf("stats TotalAlloc %v bytes\n", stats.TotalAlloc-ms.StatsTotalAlloc)
	fmt.Printf("stats HeapObjects %v\n", stats.HeapObjects-ms.StatsHeapObjects)
	fmt.Printf("stats NumGC %v\n", stats.NumGC-ms.StatsNumGC)
	fmt.Println("===================")
}

func displayInfosArray(msg string, tbl *[nbUsers]User) {
	fmt.Printf("cap : %v len : %v \n", cap(tbl), len(tbl))
	fmt.Println("tbl size:", unsafe.Sizeof(tbl), "bytes")
	fmt.Printf("after %v: %v \n", msg, tbl)
}

type MemoryStart struct {
	StatsAlloc       uint64
	StatsHeapAlloc   uint64
	StatsTotalAlloc  uint64
	StatsHeapObjects uint64
	StatsNumGC       uint32
}

var ms MemoryStart

func initMemoryStart() {
	runtime.GC()
	var stats runtime.MemStats
	runtime.ReadMemStats(&stats)
	ms.StatsAlloc = stats.Alloc
	ms.StatsHeapAlloc = stats.HeapAlloc
	ms.StatsTotalAlloc = stats.TotalAlloc
	ms.StatsHeapObjects = stats.HeapObjects
	ms.StatsNumGC = stats.NumGC
}
