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

static void call_cb(void* cb, Event_t* event) {
	((fevent_hook)(cb))(event);
}
*/
import "C"

//export Print
func Print(s *C.char) {
	log.Printf("FROM GO %s", C.GoString(s))
}

//export EventHook
func EventHook(p unsafe.Pointer, when C.uint8_t, n C.int, s **C.char, cb unsafe.Pointer) {
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
		var event C.Event_t
		event.Kind = C.uint8_t(e.Kind)
		event.When = C.uint64_t(e.When.Unix())
		event.Mask = C.uint16_t(e.Mask)
		event.Keycode = C.uint16_t(e.Keycode)
		event.Rawcode = C.uint16_t(e.Rawcode)
		event.Keychar = C.uint8_t(e.Keychar)
		event.Button = C.uint16_t(e.Button)
		event.Clicks = C.uint16_t(e.Clicks)
		event.X = C.int16_t(e.X)
		event.Y = C.int16_t(e.Y)
		event.Amount = C.uint16_t(e.Amount)
		event.Rotation = C.int32_t(e.Rotation)
		event.Direction = C.uint8_t(e.Direction)
		C.call_cb(cb, &event)
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
