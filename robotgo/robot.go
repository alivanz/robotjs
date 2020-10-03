package main

import (
	"log"
)

import "C"

//export Print
func Print(s *C.char) {
	log.Printf("FROM GO %s", C.GoString(s))
}

func main() {
}
