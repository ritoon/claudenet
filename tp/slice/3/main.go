package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	tbl1 := []int{10, 20, 30}
	displayInfos("tbl1", tbl1)

	tbl2 := make([]int, len(tbl1), 1_000)
	copy(tbl2, tbl1)
	displayInfos("tbl2", tbl2)

	// ajouter des éléments dans la slice afin de modifier le tableau sous-jacent de tbl1.
	// créer une copie de tbl1 vers tbl2 et ajouter des éléments.
}

func displayInfos(tblName string, tbl []int) {
	hdr := (*reflect.SliceHeader)(unsafe.Pointer(&tbl))
	p := (*int)(unsafe.Pointer(hdr.Data))
	msg := fmt.Sprintf("slice %v array pointer %p", tblName, p)
	msg += fmt.Sprintf(" %v a une capacité %d et une longueur %d", tblName, cap(tbl), len(tbl))
	fmt.Println(msg)
	fmt.Printf("valeurs de %v : %v\n", tblName, tbl)
}
