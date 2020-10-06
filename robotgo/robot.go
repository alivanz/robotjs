package main

import (
	"log"
	"reflect"
	"unsafe"

	"github.com/go-vgo/robotgo"
	hook "github.com/robotn/gohook"
)

/*
#include <stdint.h>
#include "types.h"

static void call_cb(fevent_hook cb, Event_t* event) {
	cb(event);
}
*/
import "C"

//export Print
func Print(s *C.char) {
	log.Printf("FROM GO %s", C.GoString(s))
}

//export EventHook
func EventHook(p unsafe.Pointer, when C.uint8_t, n C.int, s **C.char, cb C.fevent_hook) {
	keys := make([]string, n)
	var arr []*C.char
	hdr := (*reflect.SliceHeader)(unsafe.Pointer(&arr))
	hdr.Data = uintptr(unsafe.Pointer(s))
	hdr.Len = int(n)
	hdr.Cap = int(n)
	for i, e := range arr {
		keys[i] = C.GoString(e)
	}
	robotgo.EventHook(uint8(when), keys, func(e hook.Event) {
		C.call_cb(cb, nil)
	})
}

//export EventStart
func EventStart() {
	robotgo.EventStart()
}

//export EventEnd
func EventEnd() {
	robotgo.EventEnd()
}

func main() {
}
