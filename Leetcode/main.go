package main

import (
	"fmt"
	"unsafe"
)

type a struct {
	in *int
}

func main(){
	var c a
	fmt.Println(unsafe.Sizeof(c))
}